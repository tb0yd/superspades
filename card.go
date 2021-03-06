package superspades

import (
    "strconv"
    //"fmt"
)

// A simple representation of a single card. The unicode symbol for the
// suit (♥, ♦, ♣, ♠), and a string containing the number or letter value
// (2-10, J, Q, K, A)
type Card struct {
    Suit string
    Value string
}

// Card value which represents no card.
var NilCard Card = Card{}

func sameSuit(card1 int, card2 int) bool {
    return (card1 / 13) == (card2 / 13)
}

func beats(card1 int, card2 int) bool {
    if sameSuit(card1,card2) {
        return (card1 % 13) > (card2 % 13)

    } else if isSpades(card1) {
        return true

    } else if isSpades(card2) {
        return false
    }

    return false
}

// True if the card is a spade.
func isSpades(card int) bool {
    return card / 13 == 3
}

func intToCard(order int) Card {
    suitNo := order / 13
    valueNo:= order % 13

    suit := ""
    switch suitNo {
    case 0:
        suit = "♥"
    case 1:
        suit = "♦"
    case 2:
        suit = "♣"
    case 3:
        suit = "♠"
    }

    value := ""
    switch {
    case valueNo >= 0 && valueNo <= 8:
        value = strconv.Itoa(int(valueNo) + 2)
    case valueNo == 9:
        value = "J"
    case valueNo == 10:
        value = "Q"
    case valueNo == 11:
        value = "K"
    case valueNo == 12:
        value = "A"
    }

    return Card{suit, value}
}

