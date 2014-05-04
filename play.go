package superspades

import (
    "errors"
   // "fmt"
)

const MsgNotHeld = "Card not held by player."

func (g Game) Play(card Card) (Game, error) {
    // zero-index
    zxCurrentPlayer := g.CurrentPlayer - 1

    if zxCurrentPlayer != g.cardPositions[card.Order()] - cardInHand1 {
        return g, errors.New(MsgNotHeld)
    }

    // play on table in slot for player
    g.cardPositions[card.Order()] = cardOnTableForPlayer1 + zxCurrentPlayer

    // next player in circle (1-indexed)
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

    for i := uint8(0); i < 52; i++ {
        if g.cardPositions[i] >= cardOnTableForPlayer1 && g.cardPositions[i] <= cardOnTableForPlayer4 {
            cards[g.cardPositions[i] - cardOnTableForPlayer1] = card{}.InOrder(i)
            g.cardPositions[i] = cardInBook
        }
    }

    // zero-index
    zxCurrentPlayer := g.CurrentPlayer - 1

    // CurrentPlayer should be the leading player because it made a full circle
    zxWinner := winningCardNum(cards, zxCurrentPlayer)

    // add book
    g.Books[zxWinner]++

    // one-index
    g.CurrentPlayer = zxWinner + 1

    return g
}
