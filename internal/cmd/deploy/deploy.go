package deploy

import (
	"os"
	"path/filepath"

	"github.com/therecipe/qt/internal/cmd"
	"github.com/therecipe/qt/internal/cmd/minimal"
	"github.com/therecipe/qt/internal/cmd/moc"
	"github.com/therecipe/qt/internal/cmd/rcc"

	"github.com/therecipe/qt/internal/utils"
)

func Deploy(mode, target, path string, docker bool, ldFlags, tags string, fast bool, device string, vagrant bool, vagrantsystem string) {
	utils.Log.WithField("mode", mode).WithField("target", target).WithField("path", path).WithField("docker", docker).WithField("ldFlags", ldFlags).WithField("fast", fast).Debug("running Deploy")
	name := filepath.Base(path)
	depPath := filepath.Join(path, "deploy", target)

	switch mode {
	case "build", "test":

		if docker || vagrant {
			args := []string{"qtdeploy", "-debug"}
			if fast {
				args = append(args, "-fast")
			}
			if vagrantsystem == "docker" {
				args = append(args, "-docker")
			}
			args = append(args, []string{"-ldflags=" + ldFlags, "-tags=" + tags, "build"}...)

			if docker {
				cmd.Docker(args, target, path, false)
			} else {
				cmd.Vagrant(args, target, path, false, vagrantsystem)
			}
			break
		}

		if !fast {
			err := os.RemoveAll(depPath)
			if err != nil {
				utils.Log.WithError(err).Panic("failed to remove deploy folder")
			}
		}

		rcc.Rcc(path, target, tags, os.Getenv("QTRCC_OUTPUT_DIR"))
		if !fast {
			moc.Moc(path, target, tags, false, false)
		}

		if (!fast || utils.QT_STUB()) && !utils.QT_FAT() {
			minimal.Minimal(path, target, tags)
		}

		build(mode, target, path, ldFlags, tags, name, depPath, fast)

		if !(fast || utils.QT_DEBUG_QML()) {
			bundle(mode, target, path, name, depPath)
		}
	}

	if mode == "run" || mode == "test" {
		run(target, name, depPath, device)
	}
}
