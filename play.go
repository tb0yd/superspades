package superspades

import (
    "errors"
    //"fmt"
)

// Play the card in the current player's hand indexed by cardNo.
// If the current player doesn't hold
// that many cards, an error is thrown.
func (g Game) Play(cardNo int) (Game, error) {
    // zero-index
    zxCurrentPlayer := g.CurrentPlayer - 1
    hand := g.Hand(int(zxCurrentPlayer))

    if cardNo >= len(hand) {
        return g, errors.New("Not enough cards in hand.")
    }

    card := hand[cardNo]

    // play on table in slot for player
    g.cardPositions[cardToInt(card)] = cardOnTableForPlayer1 + zxCurrentPlayer

    // next player in circle (0-indexed)
    zxNextPlayer := (zxCurrentPlayer + 1) % 4

    for i := 0; i < 52; i++ {
        // assumes plays will always be made in sequence
        if g.cardPositions[i] == cardOnTableForPlayer1 + zxNextPlayer {
            return g.collectBook(), nil
        }
    }

    // one-index
    g.CurrentPlayer = zxNextPlayer + 1

    return g, nil
}

func (g Game) collectBook() Game {
    var cards [4]Card

    for i := 0; i < 52; i++ {
        if g.cardPositions[i] >= cardOnTableForPlayer1 && g.cardPositions[i] <= cardOnTableForPlayer4 {
            cards[g.cardPositions[i] - cardOnTableForPlayer1] = intToCard(i)
            g.cardPositions[i] = cardInBook
        }
    }

    // zero-index
    zxCurrentPlayer := g.CurrentPlayer - 1

    // CurrentPlayer should be the leading player because it made a full circle
    zxWinner := winningCardNum(cards, zxCurrentPlayer)

    // add book
    g.Books[zxWinner]++

    // check if all cards have been played
    if g.IsOver() {
        g.CurrentPlayer = 0
        return g
    }

    // one-index
    g.CurrentPlayer = zxWinner + 1
    return g
}
