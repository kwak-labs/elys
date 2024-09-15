package indexer

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
)

// AppI defines the interface that the app must implement
type AppI interface {
	InterfaceRegistry() types.InterfaceRegistry
}

// TX Processor, This function is implemented on every TX type interface
type Processor interface {
	Process(*LMDBManager, BaseTransaction) (Response, error)
}

type queueItem struct {
	ctx  sdk.Context
	proc Processor
}

type Response struct {
}

var (
	txChan           chan queueItem
	database         *LMDBManager
	totalIndexLength uint64
	once             sync.Once
	workerDone       chan struct{}
	app              AppI
	workerReady      sync.WaitGroup
	dbReady          chan struct{}
)

// Init initializes the indexer with a single worker and stores the app interface
func Init(a AppI) {
	once.Do(func() {
		app = a
		dbReady = make(chan struct{})
		workerReady.Add(1)

		go func() {
			db, err := NewLMDBManager("./lmdb-data", &totalIndexLength)
			if err != nil {
				panic(err)
			}
			database = db
			close(dbReady) // Signal that the database is ready
		}()

		txChan = make(chan queueItem, 10000)
		workerDone = make(chan struct{})

		go func() {
			<-dbReady // Wait for the database to be ready
			go worker()
			workerReady.Done() // Signal that the worker is ready
		}()

		// Wait for both the database and the worker to be ready
		<-dbReady
		workerReady.Wait()
	})
}

// StopIndexer stops the indexer worker
func StopIndexer() {
	close(txChan)
	<-workerDone
}

// worker processes transactions from the channel
func worker() {
	defer close(workerDone)
	for item := range txChan {
		processTransactionInternal(item.ctx, item.proc)
	}
}

// QueueTransaction sends the transaction context and processor to the worker
// This function will block if the channel is full
func QueueTransaction(ctx sdk.Context, proc Processor) {
	item := queueItem{ctx: ctx, proc: proc}
	select {
	case txChan <- item:
		fmt.Println("Processing")
		// Transaction sent to channel for processing
	default:
		// Channel is full, wait and retry
		fmt.Println("Transaction indexer channel is full, waiting to enqueue...")
		txChan <- item // This will block until there's space in the channel
	}
}

// processTransactionInternal contains the actual processing logic
// May move this into seperate file in the future
func processTransactionInternal(ctx sdk.Context, proc Processor) {
	var txBytes []byte = ctx.TxBytes()

	if len(txBytes) == 0 {
		fmt.Println("No transaction bytes found in context")
		return
	}

	// Get the sha256 checksum of the data, this returns the TX Hash
	txChecksum := sha256.Sum256(txBytes)
	txHash := hex.EncodeToString(txChecksum[:]) // Encode into hexdecimal

	// General TX information
	blockHeight := ctx.BlockHeight()
	blockTime := ctx.BlockTime() // time.Time
	gasUsed := ctx.GasMeter().GasConsumed()

	// Create a ProtoCodec using the app's InterfaceRegistry
	// Maybe create this as some top level variable, not sure
	txConfig := tx.NewTxConfig(codec.NewProtoCodec(app.InterfaceRegistry()), tx.DefaultSignModes)
	decodedTx, err := txConfig.TxDecoder()(txBytes)

	//  Hopefully this never hits
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	msg := decodedTx.GetMsgs()[0]
	sender := msg.GetSigners()[0]

	// Cast to sdk.FeeTx to access fee related methods
	feeTx, ok := decodedTx.(sdk.FeeTx)

	if !ok {
		fmt.Println("tx is not a sdk.FeeTx")
		return
	}
	memoTx, ok := decodedTx.(sdk.TxWithMemo)

	if !ok {
		fmt.Println("tx is not a sdk.TxWithMemo")
		return
	}

	memo := memoTx.GetMemo()
	fees := feeTx.GetFee()
	gasLimit := feeTx.GetGas()

	var feeDetails []FeeDetail
	for _, fee := range fees {
		feeDetails = append(feeDetails, FeeDetail{
			Amount: fee.Amount.String(),
			Denom:  fee.Denom,
		})
	}

	res, err := proc.Process(database, BaseTransaction{
		BlockTime:   blockTime,
		Author:      sender.String(),
		BlockHeight: blockHeight,
		TxHash:      txHash,
		TxType:      "d",
		Fees:        feeDetails,
		GasUsed:     strconv.FormatUint(gasUsed, 10),
		GasLimit:    strconv.FormatUint(gasLimit, 10),
		Memo:        memo,
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}

// retryProcessing attempts to reprocess a transaction after a delay
func retryProcessing(ctx sdk.Context, proc Processor) {
	go func() {
		time.Sleep(5 * time.Second) // Wait for 5 seconds before retrying
		QueueTransaction(ctx, proc)
	}()
}
