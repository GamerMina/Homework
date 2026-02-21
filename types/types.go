package types

const (
	TJS Currrency = "TJS"
	USD Currrency = "USD"
	RUB Currrency = "RUB"
)

type DepositInfo struct {
	Deposit        float64
	IntRatePerYear float64
	TimeMonths     float64
	Currency       string
	StartOfPayment string
}

type Card struct {
	Id         int
	Pan        PAN
	Balance    Money
	Currency   Currrency
	Color      string
	Name       string
	Active     bool
	MinBalance Money
}
type Currrency string
type Money int64
type PAN string
type PaymentCategory string
type PaymentStatus string
type Phone string

type Payment struct {
	ID        string
	AccountID int
	Amount    Money
	Category  PaymentCategory
	Status    PaymentStatus
}
type Account struct {
	ID      int64
	Phone   Phone
	Balance Money
}

const (
	PaymentStatsOk         PaymentStatus = "OK"
	PaymentStatsFail       PaymentStatus = "Fail"
	PaymentStatsInProgress PaymentStatus = "InProgress"
)
