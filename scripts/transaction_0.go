package main

import (
	"fmt"
	"log"
)

// --- Data Structures ---

// TransactionType defines the type of a financial transaction.
type TransactionType string

const (
	TypeDeposit        TransactionType = "deposit"
	TypeWithdraw       TransactionType = "withdraw"
	TypeGift           TransactionType = "gift"
	TypeReward         TransactionType = "reward"
	TypeCancelDeposit  TransactionType = "cancel_deposit"
	TypeCancelWithdraw TransactionType = "cancel_withdraw"
	TypeCancelGift     TransactionType = "cancel_gift"
	TypeCancelReward   TransactionType = "cancel_reward"
)

// Transaction represents a single financial event.
type Transaction struct {
	ID             string
	Type           TransactionType
	AmountSpent    float64 // Money going out from the user's perspective
	AmountReceived float64 // Money coming in to the user's perspective
}

// User represents a user with a balance.
type User struct {
	ID      string
	Balance float64
}

// FakeDB simulates a database for storing transactions.
type FakeDB struct {
	transactions []Transaction
}

// Insert adds a transaction to our fake database.
func (db *FakeDB) Insert(t Transaction) {
	log.Printf("SUCCESS: Inserting transaction %s into DB.", t.ID)
	db.transactions = append(db.transactions, t)
}

// Count returns the number of transactions in the database.
func (db *FakeDB) Count() int {
	return len(db.transactions)
}

// --- Business Logic ---

// Processor handles the logic for processing transactions.
type Processor struct{}

// ProcessTransaction is responsible for updating the user's balance and saving the transaction.
// It contains a logical flaw.
func (p *Processor) ProcessTransaction(user *User, db *FakeDB, tx Transaction) {
	log.Printf("--- Processing transaction %s (%s) ---", tx.ID, tx.Type)

	// "If a transaction results in a zero net change to the user's
	// balance, it should NOT be written to the database."

	positiveTxTypes := map[TransactionType]bool{
		TypeDeposit:        true,
		TypeReward:         true,
		TypeCancelWithdraw: true,
		TypeCancelGift:     true,
	}

	if _, ok := positiveTxTypes[tx.Type]; ok && tx.AmountReceived == 0 {
		log.Printf("INFO: Transaction %s is skipped.", tx.ID)
		return // Exit without updating balance or inserting into DB.
	}

	// Update user balance
	user.Balance += tx.AmountReceived
	user.Balance -= tx.AmountSpent
	log.Printf("Balance updated. Received: %.2f, Spent: %.2f, New Balance: %.2f", tx.AmountReceived, tx.AmountSpent, user.Balance)

	// Insert into database
	db.Insert(tx)
}

// --- Main Execution ---

func main() {
	// Setup
	user := &User{ID: "user-123", Balance: 1000.0}
	db := &FakeDB{}
	processor := &Processor{}

	// A series of transactions to be processed
	transactions := []Transaction{
		{ID: "tx-001", Type: TypeDeposit, AmountReceived: 200.0},
		{ID: "tx-002", Type: TypeWithdraw, AmountSpent: 50.0},
		{ID: "tx-003", Type: TypeDeposit, AmountReceived: 0.0},
		{ID: "tx-004", Type: TypeWithdraw, AmountSpent: 0.0},
		{ID: "tx-005", Type: TypeGift, AmountSpent: 10.0},
		{ID: "tx-007", Type: TypeReward, AmountReceived: 5.0},
	}

	// Process all transactions
	for _, tx := range transactions {
		processor.ProcessTransaction(user, db, tx)
		fmt.Println()
	}
}
