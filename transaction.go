package main

type Transaction struct {
	From      []byte
	To        []byte
	Amount    float64
	Signature []byte
}

func (t *Transaction) SetSignature() {
	t.Signature = []byte("signature")
}

func NewTransaction(from, to []byte, amount float64) *Transaction {
	return &Transaction{
		From:      from,
		To:        to,
		Amount:    amount,
		Signature: []byte{},
	}
}
