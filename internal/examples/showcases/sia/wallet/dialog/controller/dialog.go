package controller

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/therecipe/qt/core"

	"github.com/NebulousLabs/Sia/types"

	maincontroller "github.com/therecipe/qt/internal/examples/showcases/sia/controller"
	_ "github.com/therecipe/qt/internal/examples/showcases/sia/view/controller"
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

// currencyUnits converts a types.Currency to a string with human-readable
// units. The unit used will be the largest unit that results in a value
// greater than 1. The value is rounded to 4 significant digits.
func currencyUnits(c types.Currency) string {
	pico := types.SiacoinPrecision.Div64(1e12)
	if c.Cmp(pico) < 0 {
		return c.String() + " H"
	}

	// iterate until we find a unit greater than c
	mag := pico
	unit := ""
	for _, unit = range []string{"pS", "nS", "uS", "mS", "SC", "KS", "MS", "GS", "TS"} {
		if c.Cmp(mag.Mul64(1e3)) < 0 {
			break
		} else if unit != "TS" {
			// don't want to perform this multiply on the last iter; that
			// would give us 1.235 TS instead of 1235 TS
			mag = mag.Mul64(1e3)
		}
	}

	num := new(big.Rat).SetInt(c.Big())
	denom := new(big.Rat).SetInt(mag.Big())
	res, _ := new(big.Rat).Mul(num, denom.Inv(denom)).Float64()

	return fmt.Sprintf("%.4g %s", res, unit)
}

// parseCurrency converts a siacoin amount to base units.
func parseCurrency(amount string) (string, error) {
	units := []string{"pS", "nS", "uS", "mS", "SC", "KS", "MS", "GS", "TS"}
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
