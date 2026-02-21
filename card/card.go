package card

import (
	"fmt"
	"github.com/GamerMina/Homework/types"
)

const withdrawLimit = 20_000_00
const depositLimit = 50_000_00

func IssueCard(currency types.Currrency, color string, name string) types.Card {
	return types.Card{
		Id:       1000,
		Pan:      "5058 0000 0000 00x01",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
}

func Deposit(card types.Card, amount types.Money) {
	if card.Active != true {
		return
	}
	if amount > depositLimit {
		return
	}
	card.Balance += amount
	//по идеи я должен дать указатель на card но пока что нам негде хранить информацию
}

func Withdraw(card types.Card, amount types.Money) types.Card {
	if amount < 0 {
		return card
	}
	if amount > withdrawLimit {
		return card
	}
	if card.Active == false {
		return card
	}
	if card.Balance < amount {
		return card
	}
	card.Balance -= amount
	return card
	//по идеи я должен дать указатель на card но пока что нам негде хранить информацию
}
func AddBonus(card types.Card, percent int, daysInMonth int, daysInYear int) {
	if card.Active != true {
		return
	}
	if card.Balance > 0 {
		return
	}
	// Считаем бонус
	bonus := card.MinBalance * types.Money(percent) * types.Money(daysInMonth) / types.Money(daysInYear)
	if bonus > 5000 {
		bonus = 5000
	}
	card.Balance += bonus
}
func Total(cards []types.Card) types.Money {
	var total types.Money = 0
	for i := 0; i < len(cards); i++ {
		if cards[i].Active == false {
			return total //0  or return err
		}
		if cards[i].Balance < 0 {
			continue
		}
		total += cards[i].Balance
	}
	return total
}
func MaxPaymentReturn(payment []types.Payment) {
	max := MaxPayment(payment)
	fmt.Println(max.Id)
}
func MaxPayment(payment []types.Payment) types.Payment {
	var max types.Payment
	max.Amount = 0
	for i := 0; i < len(payment); i++ {
		if max.Amount < payment[i].Amount {
			max = payment[i]
		}
	}
	return max
}
