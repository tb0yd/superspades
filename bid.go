package superspades

import (
    "errors"
    "fmt"
)

const msgBadBid = "Cannot bid %d."
const msgExactly1 = "Exactly one person on the team must bid nil."

// Make a bid. It is assumed that the one bidding
// is the current player. If the bid is out of range (1-14),
// an error is thrown. 14 is the nil bid.
func (g Game) Bid(amt uint8) (Game, error) {
    // zero-index
    zxCurrentPlayer := g.CurrentPlayer - 1

    if amt < 1 || amt > 14 {
        return g, errors.New(fmt.Sprintf(msgBadBid, amt))
    }

    zxPartner := (zxCurrentPlayer + 2) % 4

    if g.Bids[zxPartner] != 0 {
        if (amt == 14 && g.Bids[zxPartner] == 14) || (amt != 14 && g.Bids[zxPartner] != 14) {
            return g, errors.New(msgExactly1)
        }
    }

    if amt != 14 {
        g.Bids[zxPartner] = 14
    }

    g.Bids[zxCurrentPlayer] = amt

    for i,j := 0,zxCurrentPlayer; i < 4; i,j = i+1, (j+1)%4 {
        if g.Bids[j] == 0 {

            // one-index
            g.CurrentPlayer = j + 1

            return g, nil
        }
    }

    // all players have bid
    g.CurrentPlayer = 0
    return g, nil
}
