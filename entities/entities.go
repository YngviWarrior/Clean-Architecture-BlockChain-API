package entities

type Entities interface {
	NewTransaction() *Transaction
}
