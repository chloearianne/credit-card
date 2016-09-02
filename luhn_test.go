package main

import "testing"

type Card struct {
	Number string
	Valid  bool
}

var cards = []Card{
	Card{
		Number: "4111111111111111",
		Valid:  true,
	},
	Card{
		Number: "5454545454545454",
		Valid:  true,
	},
	Card{
		Number: "4024007166637172",
		Valid:  true,
	},
	Card{
		Number: "1234567890123456",
		Valid:  false,
	},
	Card{
		Number: "",
		Valid:  false,
	},
	Card{
		Number: "12345678901239081209482039482309123456",
		Valid:  false,
	},
	Card{
		Number: "not a number",
		Valid:  false,
	},
	Card{
		Number: "-120394809235823",
		Valid:  false,
	},
	Card{
		Number: "545454545t445454",
		Valid:  false,
	},
}

func TestLuhn(t *testing.T) {
	for _, card := range cards {
		if validCard(card.Number) != card.Valid {
			t.Errorf("Card %s returned wrong validity", card.Number)
		}
	}
}
