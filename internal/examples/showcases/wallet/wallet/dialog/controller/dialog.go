package controller

import (
	"errors"
	"math/big"
	"strings"

	"github.com/therecipe/qt/core"

	maincontroller "github.com/therecipe/qt/internal/examples/showcases/wallet/controller"
	_ "github.com/therecipe/qt/internal/examples/showcases/wallet/view/controller"
)

func init() {
	if maincontroller.Controller != nil {
	}
}

var Controller *dialogController

type dialogController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(cident string) `signal:"show"`
	_ func(bool)          `signal:"blur,->(controller.Controller)"`

	_ func(password string) *core.QVariant     `slot:"unlock,auto"`
	_ func() string                            `slot:"receive,auto"`
	_ func(amount, dest string) *core.QVariant `slot:"send,auto"`
	_ func(seed string) *core.QVariant         `slot:"recover,auto"`

	_ func() bool `slot:"isLocked,->(maincontroller.Controller)"`
}

func (c *dialogController) init() {
	Controller = c
}

func (c *dialogController) receive() string {
	return ""
}

func (c *dialogController) send(amount, dest string) *core.QVariant {
	return core.NewQVariant()
}

func (c *dialogController) unlock(password string) *core.QVariant {
	return core.NewQVariant()
}

func (c *dialogController) recover(seed string) *core.QVariant {
	return core.NewQVariant()
}

// parseCurrency converts a coin amount to base units.
func parseCurrency(amount string) (string, error) {
	units := []string{"pS", "nS", "uS", "mS", "C", "KS", "MS", "GS", "TS"}
	for i, unit := range units {
		if strings.HasSuffix(amount, unit) {
			// scan into big.Rat
			r, ok := new(big.Rat).SetString(strings.TrimSuffix(amount, unit))
			if !ok {
				return "", errors.New("malformed amount")
			}
			// convert units
			exp := 24 + 3*(int64(i)-4)
			mag := new(big.Int).Exp(big.NewInt(10), big.NewInt(exp), nil)
			r.Mul(r, new(big.Rat).SetInt(mag))
			// r must be an integer at this point
			if !r.IsInt() {
				return "", errors.New("non-integer number of hastings")
			}
			return r.RatString(), nil
		}
	}
	// check for hastings separately
	if strings.HasSuffix(amount, "H") {
		return strings.TrimSuffix(amount, "H"), nil
	}

	return "", errors.New("amount is missing units; run 'wallet --help' for a list of units")
}
