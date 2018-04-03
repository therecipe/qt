package controller

import (
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"

	"github.com/NebulousLabs/Sia/types"

	"github.com/therecipe/qt/internal/examples/showcases/sia/controller"
	"github.com/therecipe/qt/internal/examples/showcases/sia/wallet/model"
)

var Controller *walletController

type walletController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ *model.WalletModel `property:"model"`

	_ func(ID string) `signal:"doubleClicked"`
}

func (c *walletController) init() {
	Controller = c

	c.SetModel(model.NewWalletModel(nil))

	c.ConnectDoubleClicked(c.doubleClicked)

	go c.loop()
}

func (c *walletController) doubleClicked(ID string) {
	gui.QDesktopServices_OpenUrl(core.NewQUrl3(fmt.Sprintf("https://explore.sia.tech/hash.html?hash=%v", ID), 0))
}

func (c *walletController) loop() {
	for range time.NewTicker(1 * time.Second).C {

		wtg, err := controller.Client.WalletTransactionsGet(0, 10000000)
		if err != nil && !DEMO {
			println(err.Error())
		} else {

			transactions := make([]model.Transaction, 0)
			var balance float64

			for _, txn := range append(wtg.ConfirmedTransactions, wtg.UnconfirmedTransactions...) {
				transaction := model.Transaction{ID: txn.TransactionID.String(), Type: "Transaction"}

				// Determine the number of outgoing siacoins and siafunds.
				var outgoingSiacoins types.Currency
				var outgoingSiafunds types.Currency
				for _, input := range txn.Inputs {
					if input.FundType == types.SpecifierSiacoinInput && input.WalletAddress {
						outgoingSiacoins = outgoingSiacoins.Add(input.Value)
					}
					if input.FundType == types.SpecifierSiafundInput && input.WalletAddress {
						outgoingSiafunds = outgoingSiafunds.Add(input.Value)
					}
				}

				// Determine the number of incoming siacoins and siafunds.
				var incomingSiacoins types.Currency
				var incomingSiafunds types.Currency
				for _, output := range txn.Outputs {
					if output.FundType == types.SpecifierMinerPayout {
						incomingSiacoins = incomingSiacoins.Add(output.Value)
					}
					if output.FundType == types.SpecifierSiacoinOutput && output.WalletAddress {
						incomingSiacoins = incomingSiacoins.Add(output.Value)
					}
					if output.FundType == types.SpecifierSiafundOutput && output.WalletAddress {
						incomingSiafunds = incomingSiafunds.Add(output.Value)
					}
				}

				// Convert the siacoins to float.
				incomingSiacoinsFloat, _ := new(big.Rat).SetFrac(incomingSiacoins.Big(), types.SiacoinPrecision.Big()).Float64()
				outgoingSiacoinsFloat, _ := new(big.Rat).SetFrac(outgoingSiacoins.Big(), types.SiacoinPrecision.Big()).Float64()
				amount := incomingSiacoinsFloat - outgoingSiacoinsFloat

				if amount == 0 {
					continue
				}

				if controller.Controller.Height() >= txn.ConfirmationHeight {
					transaction.Status = "Confirmed"
					transaction.Date = time.Unix(int64(txn.ConfirmationTimestamp), 0).Format(time.RFC822Z)
				} else {
					transaction.Status = "Not Confirmed"
				}

				transaction.Amount = fmt.Sprintf("%1.5f SC", amount)

				balance += amount
				transaction.Total = fmt.Sprintf("%1.5f SC", balance)

				//TODO: siafunds
				// For siafunds, need to avoid having a negative types.Currency.
				if incomingSiafunds.Cmp(outgoingSiafunds) >= 0 {
					//fmt.Sprintf("%14v SF", incomingSiafunds.Sub(outgoingSiafunds))
				} else {
					//fmt.Sprintf("-%14v SF", outgoingSiafunds.Sub(incomingSiafunds))
				}

				transactions = append(transactions, transaction)
			}

			if DEMO {
				var dt []model.Transaction
				json.Unmarshal([]byte(DEMO_TRANSACTIONS), &dt)
				c.Model().UpdateWith(dt)
			} else {
				c.Model().UpdateWith(transactions)
			}
		}
	}
}
