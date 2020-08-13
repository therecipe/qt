package deploy

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"

	"github.com/therecipe/qt/internal/cmd"
	"github.com/therecipe/qt/internal/cmd/minimal"
	"github.com/therecipe/qt/internal/cmd/moc"
	"github.com/therecipe/qt/internal/cmd/rcc"

	"github.com/therecipe/qt/internal/utils"
)

func Deploy(mode, target, path string, docker bool, ldFlags, tags string, fast bool, device string, vagrant bool, vagrantsystem string, comply bool, useuic bool, quickcompiler bool) {
	defer func() { parser.State.ClassMap = make(map[string]*parser.Class) }()

	utils.Log.WithField("mode", mode).WithField("target", target).WithField("path", path).WithField("docker", docker).
		WithField("ldFlags", ldFlags).WithField("fast", fast).WithField("comply", comply).
		WithField("useuic", useuic).WithField("quickcompiler", quickcompiler).Debug("running Deploy")

	name := filepath.Base(path)
	switch name {
	case "lib", "plugins", "qml",
		"audio", "bearer", "iconengines", "imageformats", "mediaservice",
		"platforminputcontexts", "platforms", "playlistformats", "qmltooling",
		"qt", "Qt", "QT", "styles", "translations":
		name += "_project"
	}
	depPath := filepath.Join(path, "deploy", target)

	switch mode {
	case "build", "test":

		if !(fast && (strings.HasPrefix(target, "js") || strings.HasPrefix(target, "wasm"))) {
			err := os.RemoveAll(depPath)
			if err != nil {
				utils.Log.WithError(err).Panic("failed to remove deploy folder")
			}
		}

		if utils.UseGOMOD(path) {
			if !utils.ExistsDir(filepath.Join(filepath.Dir(utils.GOMOD(path)), "vendor")) {
				cmd := exec.Command("go", "mod", "vendor")
				cmd.Dir = path
				utils.RunCmd(cmd, "go mod vendor")
			}
			if utils.QT_DOCKER() {
				cmd := exec.Command("go", "get", "-v", "-d", "github.com/therecipe/qt/internal/binding/files/docs/"+utils.QT_API(utils.QT_VERSION())+"@"+cmd.QtModVersion(filepath.Dir(utils.GOMOD(path)))) //TODO: needs to pull 5.8.0 if QT_WEBKIT
				cmd.Dir = path
				if !utils.QT_PKG_CONFIG() {
					utils.RunCmdOptional(cmd, "go get docs") //TODO: this can fail if QT_PKG_CONFIG
				}

				if strings.HasPrefix(target, "sailfish") || strings.HasPrefix(target, "android") { //TODO: generate android and sailfish minimal instead
					cmd := exec.Command(filepath.Join(utils.GOBIN(), "qtsetup"), "generate", target)
					cmd.Dir = path
					utils.RunCmd(cmd, "run setup")
				}
			}
		}

		if docker || vagrant {
			args := []string{"qtdeploy", "-debug"}
			if fast {
				args = append(args, "-fast")
			}
			if comply {
				args = append(args, "-comply")
			}
			if !useuic {
				args = append(args, "-uic=false")
			}
			if quickcompiler {
				args = append(args, "-quickcompiler")
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

		if (utils.QT_FELGO() && fast) || utils.QT_FELGO_LIVE() {
			utils.Save(filepath.Join(path, "live.pro"), "")
		}

		if utils.ExistsDir(depPath + "_obj") {
			utils.RemoveAll(depPath + "_obj")
		}

		if !(utils.QT_GEN_GO_WRAPPER() && !utils.QT_FAT()) {
			rcc.Rcc(path, target, tags, os.Getenv("QTRCC_OUTPUT_DIR"), useuic, quickcompiler, true, true)
		}
		if cmd.ImportsFlutter() {
			flutter(target, path)
		}
		if !fast && !(utils.QT_GEN_GO_WRAPPER() && !utils.QT_FAT()) {
			moc.Moc(path, target, tags, false, false, true, true)
		}

		if ((!fast || utils.QT_STUB()) || ((target == "js" || target == "wasm") && (utils.QT_DOCKER() || utils.QT_VAGRANT()))) && !utils.QT_FAT() {
			minimal.Minimal(path, target, tags, true)
		}

		build(mode, target, path, ldFlags, tags, name, depPath, fast, comply)

		if !(fast || ((utils.QT_DEBUG_QML() || utils.QT_FELGO_LIVE()) && target == runtime.GOOS)) && !(utils.QT_GEN_GO_WRAPPER() && !utils.QT_FAT()) || (target == "js" || target == "wasm") {
			bundle(mode, target, path, name, depPath, tags, fast)
		} else if fast || (utils.QT_DEBUG_QML() || utils.QT_FELGO_LIVE()) || utils.QT_GEN_GO_WRAPPER() {
			switch target {
			case "darwin":
				if fn := filepath.Join(depPath, name+".app", "Contents", "Info.plist"); !utils.ExistsFile(fn) {
					utils.Save(fn, darwin_plist(name))
				}
			case "windows":
				if utils.QT_MSVC() {
					_, _, _, out := cmd.BuildEnv(target, name, depPath)
					cmd.PatchBinary(out + ".exe")
				}
			}
		}
	}

	if (mode == "run" || mode == "test") && !(fast && (strings.HasPrefix(target, "js") || strings.HasPrefix(target, "wasm"))) {
		run(target, name, depPath, device)
	}
}
