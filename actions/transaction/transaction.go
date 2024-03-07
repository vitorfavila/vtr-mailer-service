package transaction

import (
	"errors"
	"example/vtr-mailer-service/stores"
	"example/vtr-mailer-service/structs"
)

const (
	Created         = iota // 0
	ProcessPending         // 1
	Processing             // 2
	ProcessFinished        // 3
	Sent                   // 4
	Failed                 // 5
)

func GetTransaction(transactionId int64) (structs.Transaction, error) {
	for _, transaction := range stores.Transactions {
		if transaction.Id == transactionId {
			return transaction, nil
		}
	}

	return structs.Transaction{}, errors.New("unable to find transaction")
}

func AddTransaction(transaction structs.Transaction) {
	stores.Transactions = append(stores.Transactions, transaction)
}

func UpdateTransaction(transaction structs.Transaction) error {
	idx := 0

	for index, trans := range stores.Transactions {
		if trans.Id == transaction.Id {
			idx = index
		}
	}

	if idx == 0 {
		return errors.New("unable to find transaction")
	}

	stores.Transactions[idx] = transaction

	return nil
}
