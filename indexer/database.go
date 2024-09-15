package indexer

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"

	"github.com/bmatsuo/lmdb-go/lmdb"
)

// LMDBManager handles LMDB operations for storing and retrieving transactions
type LMDBManager struct {
	env              *lmdb.Env
	txDB             lmdb.DBI
	addressDB        lmdb.DBI
	txCountDB        lmdb.DBI
	path             string
	totalIndexLength *uint64
}

// NewLMDBManager creates a new LMDB manager
func NewLMDBManager(path string, totalIndexLength *uint64) (*LMDBManager, error) {
	// Create directory if it doesn't exist
	if err := os.MkdirAll(path, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory: %v", err)
	}

	// Set up LMDB environment
	env, err := lmdb.NewEnv()
	if err != nil {
		return nil, err
	}

	if err := env.SetMaxDBs(3); err != nil {
		return nil, err
	}

	// Start with 1GB, we'll increase it later if needed
	if err := env.SetMapSize(1 << 30); err != nil {
		return nil, err
	}

	if err := env.Open(path, 0, 0644); err != nil {
		return nil, err
	}

	manager := &LMDBManager{env: env, path: path, totalIndexLength: totalIndexLength}

	// Initialize databases
	err = env.Update(func(txn *lmdb.Txn) error {
		var err error
		if manager.txDB, err = txn.CreateDBI("txs"); err != nil {
			return err
		}
		if manager.addressDB, err = txn.CreateDBI("addresses"); err != nil {
			return err
		}
		if manager.txCountDB, err = txn.CreateDBI("txcount"); err != nil {
			return err
		}

		// Get the current transaction count
		countBytes, err := txn.Get(manager.txCountDB, []byte("count"))
		if err == nil {
			*totalIndexLength = binary.LittleEndian.Uint64(countBytes)
		} else if lmdb.IsNotFound(err) {
			*totalIndexLength = 0
		} else {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return manager, nil
}

// CheckAndResizeIfNeeded increases the database size if it's getting full
func (m *LMDBManager) CheckAndResizeIfNeeded() error {
	info, err := m.env.Info()
	if err != nil {
		return err
	}

	usedSpace := uint64(info.LastPNO) * uint64(os.Getpagesize())
	availableSpace := uint64(info.MapSize) - usedSpace

	// Double the size if less than 20% is available
	if availableSpace < uint64(info.MapSize)/5 {
		newSize := info.MapSize * 2
		if err := m.env.SetMapSize(newSize); err != nil {
			// If resizing fails, close and reopen the environment
			m.env.Close()
			if env, err := lmdb.NewEnv(); err == nil {
				if err := env.SetMaxDBs(3); err == nil {
					if err := env.SetMapSize(newSize); err == nil {
						if err := env.Open(m.path, 0, 0644); err == nil {
							m.env = env
							fmt.Printf("Resized database to %d bytes\n", newSize)
							return nil
						}
					}
				}
			}
			return err
		}
	}

	return nil
}

// ProcessNewTx adds a new transaction to the database
func (m *LMDBManager) ProcessNewTx(tx interface{}, address string) error {
	if err := m.CheckAndResizeIfNeeded(); err != nil {
		return err
	}

	return m.env.Update(func(txn *lmdb.Txn) error {
		count, err := m.incrementTxCount(txn)
		if err != nil {
			return err
		}
		fmt.Printf("New Count: %d\n", count)

		// Store the transaction
		txBytes, err := json.Marshal(tx)
		if err != nil {
			return err
		}

		indexBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(indexBytes, count)
		if err := txn.Put(m.txDB, indexBytes, txBytes, 0); err != nil {
			return err
		}

		// Update the address index
		var indices []uint64
		indicesBytes, err := txn.Get(m.addressDB, []byte(address))
		if err != nil {
			if !lmdb.IsNotFound(err) {
				return err
			}
			indices = []uint64{}
		} else {
			if err := json.Unmarshal(indicesBytes, &indices); err != nil {
				return err
			}
		}

		indices = append(indices, count)
		newIndicesBytes, err := json.Marshal(indices)
		if err != nil {
			return err
		}

		return txn.Put(m.addressDB, []byte(address), newIndicesBytes, 0)
	})
}

// incrementTxCount increases the transaction count and returns the new value
func (m *LMDBManager) incrementTxCount(txn *lmdb.Txn) (uint64, error) {
	var count uint64
	countBytes, err := txn.Get(m.txCountDB, []byte("count"))
	if err != nil {
		if !lmdb.IsNotFound(err) {
			return 0, fmt.Errorf("error retrieving count: %v", err)
		}
		count = 0
	} else {
		count = binary.BigEndian.Uint64(countBytes)
	}

	count++
	countBytes = make([]byte, 8)
	binary.BigEndian.PutUint64(countBytes, count)
	if err := txn.Put(m.txCountDB, []byte("count"), countBytes, 0); err != nil {
		return 0, fmt.Errorf("error storing new count: %v", err)
	}

	*m.totalIndexLength = count
	return count, nil
}

// GetTxCount returns the total number of transactions
func (m *LMDBManager) GetTxCount() uint64 {
	return *m.totalIndexLength
}

// GetTxByIndex retrieves a transaction by its index
func (m *LMDBManager) GetTxByIndex(index uint64) (interface{}, error) {
	var tx interface{}
	err := m.env.View(func(txn *lmdb.Txn) error {
		indexBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(indexBytes, index)
		txBytes, err := txn.Get(m.txDB, indexBytes)
		if err != nil {
			return err
		}
		return json.Unmarshal(txBytes, &tx)
	})
	return tx, err
}

// GetTxsByAddress retrieves all transactions for a given address
func (m *LMDBManager) GetTxsByAddress(address string) ([]interface{}, error) {
	var txs []interface{}
	err := m.env.View(func(txn *lmdb.Txn) error {
		indicesBytes, err := txn.Get(m.addressDB, []byte(address))
		if err == lmdb.NotFound {
			return nil
		} else if err != nil {
			return err
		}

		var indices []uint64
		if err := json.Unmarshal(indicesBytes, &indices); err != nil {
			return err
		}

		for _, index := range indices {
			tx, err := m.GetTxByIndex(index)
			if err != nil {
				return err
			}
			txs = append(txs, tx)
		}
		return nil
	})
	return txs, err
}

// Close shuts down the LMDB environment
func (m *LMDBManager) Close() error {
	m.env.Close()
	return nil
}
