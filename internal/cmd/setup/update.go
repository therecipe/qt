package setup

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/therecipe/qt/internal/utils"
)

func fetch() {
	var fetch = exec.Command("git", "fetch", "--all")
	fetch.Dir = filepath.Join(utils.MustGoPath(), "src", "github.com", "therecipe", "qt")
	utils.RunCmd(fetch, "run \"git fetch\"")
}

func tooling() {
	utils.RunCmd(exec.Command("go", "install", "-v", fmt.Sprintf("github.com/therecipe/qt/cmd/...")), "run \"go install\"")
}

func update() {
	utils.RunCmd(exec.Command("go", "clean", "-i", "github.com/therecipe/qt/cmd/..."), "run \"go clean cmd\"")
	utils.RunCmd(exec.Command("go", "clean", "-i", "github.com/therecipe/qt/internal/..."), "run \"go clean internal\"")

	fetch()

	var checkoutCmd = exec.Command("git", "checkout", "--", utils.GoQtPkgPath("cmd"))
	checkoutCmd.Dir = filepath.Join(utils.MustGoPath(), "src", "github.com", "therecipe", "qt")
	utils.RunCmd(checkoutCmd, "run \"git checkout cmd\"")

	var checkoutInternal = exec.Command("git", "checkout", "--", utils.GoQtPkgPath("internal"))
	checkoutInternal.Dir = filepath.Join(utils.MustGoPath(), "src", "github.com", "therecipe", "qt")
	utils.RunCmd(checkoutInternal, "run \"git checkout internal\"")

	tooling()
}

func upgrade() {
	utils.RunCmd(exec.Command("go", "clean", "-i", "github.com/therecipe/qt/..."), "run \"go clean\"")

	fetch()

	var clean = exec.Command("git", "clean", "--force")
	clean.Dir = filepath.Join(utils.MustGoPath(), "src", "github.com", "therecipe", "qt")
	utils.RunCmd(clean, "run \"git clean\"")

	var reset = exec.Command("git", "reset", "--hard")
	reset.Dir = filepath.Join(utils.MustGoPath(), "src", "github.com", "therecipe", "qt")
	utils.RunCmd(reset, "run \"git reset\"")

	tooling()
}
