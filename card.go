package superspades

import "github.com/jteeuwen/deck"

// Utility functions for cards. You can define your own card datatype.
//
// Beats() is non-exclusive -- it's possible that card1.Beats(card2) could
// be false, and card2.Beats(card1) could also be false. In this case, it
// would depend on who threw what first, which the cards themselves don't
// know. This is also the case if you compare two identical cards. Two
// cards will never return true compared against each other, though.
//
// Order() is an integer value, 0 through 51, unique for that card.
type Card interface {
    SameSuit(Card) bool
    Beats(Card) bool
    IsSpades() bool
    Order() uint8
    InOrder(uint8) Card

    // private methods to implement the above
    suit() uint8
    value() uint8
}

type card struct {
    card deck.Card
}

func (card card) suit() uint8 {
    return uint8(card.card.Suit())
}

func (card card) value() uint8 {
    // deck.Card.Value() makes Ace into a 0, 2 into 1, etc...
    // We want value to be in order of what beats what
    return (card.card.Value() + 12) % 13
}


func (card1 card) SameSuit(card2 Card) bool {
    return card1.suit() == card2.suit()
}

func (card1 card) Beats(card2 Card) bool {
    if card1.SameSuit(card2) {
        return card1.value() > card2.value()

    } else if card1.IsSpades() {
        return true

    } else if card2.IsSpades() {
        return false
    }

    return false
}

// True if the card is a spade.
func (card card) IsSpades() bool {
    return card.suit() == uint8(deck.Spades)
}

func (card card) Order() uint8 {
    return (card.suit() * 13) + card.value()
}

func (c card) InOrder(order uint8) Card {
    return card{deck.Card(uint8((order / 13) << 4) | ((order + 1) % 13))}
}

