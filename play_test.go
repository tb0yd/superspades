package superspades

import (
    "testing"
)

func TestPlay(t *testing.T) {
    // test Play by simulating an incomplete hand. Player 1 led,
    // Player 3's turn. Cards are: 2H, 3H, 4H, 5H in that order.
    g := NewGame()
    g.cardPositions[0] = cardOnTableForPlayer1
    g.cardPositions[1] = cardOnTableForPlayer2
    g.cardPositions[2] = cardInHand3
    g.cardPositions[3] = cardInHand4
    g.CurrentPlayer = 3

    g, err := g.Play(0)

    if err != nil {
        t.Errorf("got error: %s\n", err.Error())
    }

    if g.CurrentPlayer != 4 {
        t.Errorf("current player should be 4, was %d\n", g.CurrentPlayer)
    }

    if g.Books[3] != 0 {
        t.Errorf("Books[3] should be 0, was %d\b", g.Books[3])
    }

    cardsOnTable := 0
    for i := 0; i < 52; i++ {
        if g.cardPositions[i] >= cardOnTableForPlayer1 && g.cardPositions[i] <= cardOnTableForPlayer4 {
            cardsOnTable++
        }
    }
    if cardsOnTable != 3 {
        t.Errorf("should be 3 cards on table, were %d\n", cardsOnTable)
    }

    g, err = g.Play(0)

    if err != nil {
        t.Errorf("got error: %s\n", err.Error())
    }

    if g.CurrentPlayer != 4 {
        t.Errorf("current player should be 4, was %d\n", g.CurrentPlayer)
    }

    if g.Books[3] != 1 {
        t.Errorf("Books[3] should be 1, was %d\b", g.Books[3])
    }

    cardsOnTable = 0
    for i := 0; i < 52; i++ {
        if g.cardPositions[i] >= cardOnTableForPlayer1 && g.cardPositions[i] <= cardOnTableForPlayer4 {
            cardsOnTable++
        }
    }
    if cardsOnTable > 0 {
        t.Errorf("should be no cards on table, were %d\n", cardsOnTable)
    }

    // simulate an end game
    // player 4's go, has ace of spades
    g = Game{
        4,
        [4]uint8{4,14,14,4},
        [4]uint8{4,4,4,4},
        [52]uint8{
            cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,
            cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,
            cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,
            cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardInBook,cardOnTableForPlayer1,cardOnTableForPlayer2,cardOnTableForPlayer3,cardInHand4,
        },
    }

    g, err = g.Play(0)

    if err != nil {
        t.Errorf("got error: %s\n", err.Error())
    }

    if g.CurrentPlayer != 0 {
        t.Errorf("current player should be 0, was %d\n", g.CurrentPlayer)
    }

    if g.Books[3] != 5 {
        t.Errorf("Books[3] should be 5, was %d\b", g.Books[3])
    }

    cardsOnTable = 0
    for i := 0; i < 52; i++ {
        if g.cardPositions[i] >= cardOnTableForPlayer1 && g.cardPositions[i] <= cardOnTableForPlayer4 {
            cardsOnTable++
        }
    }
    if cardsOnTable > 0 {
        t.Errorf("should be no cards on table, were %d\n", cardsOnTable)
    }

}
