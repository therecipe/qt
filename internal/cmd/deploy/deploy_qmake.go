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

func Deploy(mode, target, path string, docker bool, ldFlags string, fast bool) {
	utils.Log.WithField("mode", mode).WithField("target", target).WithField("path", path).WithField("docker", docker).WithField("ldFlags", ldFlags).WithField("fast", fast).Debug("running Deploy")
	name := filepath.Base(path)
	depPath := filepath.Join(path, "deploy", target)

	switch mode {
	case "build", "test":

		if !fast {
			err := os.RemoveAll(depPath)
			if err != nil {
				utils.Log.WithError(err).Panic("failed to remove deploy folder")
			}
		}

		if docker {
			args := []string{"qtdeploy", "-debug"}
			if fast {
				args = append(args, "-fast")
			}
			args = append(args, []string{"-ldflags=" + ldFlags, "build"}...)
			cmd.Docker(args, target, path)
			break
		}

		rcc.Rcc(path, target, os.Getenv("QTRCC_OUTPUT_DIR"))
		if !fast {
			moc.Moc(path, target)
			minimal.Minimal(path, target)
		}

		build(mode, target, path, ldFlags, name, depPath)

		if !fast {
			bundle(mode, target, path, name, depPath)
		}

	default:
		utils.Log.Errorf("wrong mode: %v", mode) //TODO: check mode in cmd
	}

	if mode == "run" || mode == "test" {
		run(target, name, depPath)
	}
}
