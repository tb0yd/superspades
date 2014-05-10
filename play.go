package superspades

import (
    "errors"
    //"fmt"
)

// Play the card in the current player's PlayableCards(), indexed by cardNo.
// If the current player doesn't have
// that many cards playable, an error is thrown.
func (g Game) Play(cardNo int) (Game, error) {
    // zero-index
    zxCurrentPlayer := g.CurrentPlayer - 1
    playable := g.PlayableCards()

    if cardNo >= len(playable) {
        return g, errors.New("Not enough playable cards.")
    }

    card := playable[cardNo]

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

// Returns a slice of cards in the current player's hand. Only
// cards which the player is allowed to play will appear.
func (g Game) PlayableCards() []Card {
    var cardsArray [13]Card
    i := 0
    hand := g.Hand(g.CurrentPlayer-1)
    handInts := make([]int, len(hand))

    for j, card := range hand {
        handInts[j] = cardToInt(card)
    }

    for _, card := range hand {
        if canPlayMove(cardToInt(card), handInts, g.spadesBroken(), g.leading(), g.lastSuit()) {
            cardsArray[i] = card
            i++
        }
    }

    return cardsArray[0:i]
}

func (g Game) spadesBroken() bool {
    for i := 39; i < 52; i++ {
        if g.cardPositions[i] > cardInHand4 {
            return true
        }
    }

    return false
}

func (g Game) leading() bool {
    for i := 0; i < 52; i++ {
        if g.cardPositions[i] >= cardOnTableForPlayer1 && g.cardPositions[i] <= cardOnTableForPlayer4 {
            return false
        }
    }

    return true
}

func (g Game) lastSuit() int {
    var oxOnTable [4]int

    for i := 0; i < 52; i++ {
        if g.cardPositions[i] >= cardOnTableForPlayer1 && g.cardPositions[i] <= cardOnTableForPlayer4 {
            oxOnTable[g.cardPositions[i] - cardOnTableForPlayer1] = i
        }
    }

    for i,j := (g.CurrentPlayer+1) % 4,0; j < 4; i,j = (i+1) % 4, j + 1 {
        zxPrevPlayer := (i+3) % 4

        if oxOnTable[zxPrevPlayer] == 0 && oxOnTable[i] != 0 {
            return (oxOnTable[i] - 1) / 13
        }
    }

    return 4
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

    // next player in circle (0-indexed)
    zxNextPlayer := (zxCurrentPlayer + 1) % 4

    // zxNextPlayer should also be the leading player because it made a full circle
    zxWinner := winningCardNum(cards, zxNextPlayer)

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
