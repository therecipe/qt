package controller

import (
	"os/exec"
	"strings"

	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/internal/examples/showcases/sia/controller"
	wcontroller "github.com/therecipe/qt/internal/examples/showcases/sia/wallet/dialog/controller"
)

var PathToSiac string

type terminalController struct {
	core.QObject

	_ func(cmd string) string `slot:"command,auto"`
}

func (c *terminalController) command(cmd string) string {
	if cmd == "wallet unlock" {
		if controller.Controller.IsLocked() {
			wcontroller.Controller.Show("unlock")
			return ""
		}
		return "Wallet already unlocked"
	} else {
		ecmd := exec.Command(PathToSiac, strings.Split(cmd, " ")...)
		out, err := ecmd.CombinedOutput()
		if err != nil {
			println(err.Error())
		}
		return string(out)
	}
}
