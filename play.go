package superspades

import "errors"

func (g Game) Play(card Card) (Game, error) {
    if g.CurrentPlayer != g.cardPositions[card.Order()] {
        return g, errors.New(MsgNotHeld)
    }

    return g, nil
}
