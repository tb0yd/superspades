package superspades

import "errors"

func (g Game) Play(card Card) (Game, error) {
    if state, msg := g.State(); state != GameInPlay {
        return g, errors.New(msg)
    }

    if g.CurrentPlayer != g.cardPositions[card.Order()] {
        return g, errors.New(MsgNotHeld)
    }

    return g, nil
}
