package types

// Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и т.д).
type Money int64

// Category представляет собой категорию, в которой был совершен платёж (авто, аптека, рестораны и т.д).
type Category string

// Status представляет статус плятежей.
type Status string


// Payment представляет информацию о платеже.
type Payment struct {
	ID       int
	Amount   Money
	Category Category
}