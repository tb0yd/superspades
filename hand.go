package superspades

// Returns a slice of Cards representing the cards in player n's hand.
func (g Game) Hand(n uint8) []Card {
    var cardArray [13]Card

    c := 0

    for i, pos := range g.cardPositions {
        if pos == cardInHand1 + n {
            cardArray[c] = intToCard(i)
            c++
        }
    }

    return cardArray[0:c]
}

// Returns all the cards on the table, indexed by player.
func (g Game) Table() (cards [4]Card) {
    for i, pos := range g.cardPositions {
        if pos >= cardOnTableForPlayer1 && cardOnTableForPlayer4 >= pos {
            cards[pos - cardOnTableForPlayer1] = intToCard(i)
        }
    }

    return
}
