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
	OnQueue                // 4
	Sent                   // 5
	FailOnProcess          // 6
	FailOnSend             // 7
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
	idx := -1

	for index := range stores.Transactions {
		if stores.Transactions[index].Id == transaction.Id {
			idx = index
			break
		}
	}

	if idx < 0 {
		return errors.New("unable to find transaction")
	}

	stores.Transactions[idx] = transaction

	return nil
}

func UpdateTransactionStatus(transactionId int64, status structs.TransactionStatus) structs.Transaction {
	trans, err := GetTransaction(transactionId)
	if err != nil {
		// TODO - Logar erro
	}

	trans.Status = status
	UpdateTransaction(trans)

	return trans
}
