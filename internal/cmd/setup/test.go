package setup

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/bluszcz/cutego/internal/binding/parser"
	"github.com/bluszcz/cutego/internal/binding/templater"

	"github.com/bluszcz/cutego/internal/cmd"
	"github.com/bluszcz/cutego/internal/cmd/deploy"
	"github.com/bluszcz/cutego/internal/cmd/minimal"
	"github.com/bluszcz/cutego/internal/cmd/moc"
	"github.com/bluszcz/cutego/internal/cmd/rcc"

	"github.com/bluszcz/cutego/internal/utils"
)

func Test(target string, docker, vagrant bool, vagrantsystem string) {
	utils.Log.Infof("running: 'qtsetup test %v' [docker=%v] [vagrant=%v]", target, docker, vagrant)

	if utils.CI() && target == runtime.GOOS && !(runtime.GOOS == "windows" || utils.QT_STATIC()) && !utils.QT_GEN_GO_WRAPPER() { //TODO: test on windows
		utils.Log.Infof("running setup/test %v CI", target)

		path := utils.GoQtPkgPath("internal", "cmd", "moc", "test")

		moc.Moc(path, target, "", false, false, false, true)
		minimal.Minimal(path, target, "", true)

		var pattern string
		if utils.GOVERSION_NUM() >= 110 {
			pattern = "all="
		}

		cmd := exec.Command("go", "test", "-v", "-tags=minimal", fmt.Sprintf("-ldflags=%v\"-s\"", pattern), "-race")
		cmd.Env = append(os.Environ(), []string{"GODEBUG=cgocheck=2", "GO111MODULE=off"}...)
		if runtime.GOOS == "linux" {
			cmd.Env = append(cmd.Env, "QT_QPA_PLATFORM=offscreen")
		}
		cmd.Dir = path
		utils.RunCmd(cmd, "run \"go test\"")

		parser.State.ClassMap = make(map[string]*parser.Class)
		rcc.ResourceNames = make(map[string]string)
		moc.ResourceNames = make(map[string]string)
	}

	mode := "test"
	var examples map[string][]string
	if utils.CI() {
		mode = "build"
		examples = map[string][]string{
			"androidextras": []string{"jni", "notification"},

			"canvas3d": []string{"framebuffer", "interaction", "jsonmodels",
				"quickitemtexture", "textureandlight",
				filepath.Join("threejs", "cellphone"),
				filepath.Join("threejs", "oneqt"),
				filepath.Join("threejs", "planets"),
			},

			"charts": []string{"audio", "dynamicspline", "modeldata"},

			"common": []string{"qml_demo", "widgets_demo"},

			//"grpc": []string{"hello_world","hello_world2"}, //TODO: re-gen protobuf go files

			//"gui": []string{"analogclock", "openglwindow", "rasterwindow"},

			//"opengl": []string{"2dpainting"},

			"qml": []string{"adding", "application",
				"custom_scheme", "drawer_nav_x",

				filepath.Join("extending", "chapter1-basics"),
				filepath.Join("extending", "chapter2-methods"),
				filepath.Join("extending", "chapter3-bindings"),
				filepath.Join("extending", "chapter4-customPropertyTypes"),
				filepath.Join("extending", "components", "test_dir"),
				filepath.Join("extending", "components", "test_dir_2"),
				filepath.Join("extending", "components", "test_go"),
				filepath.Join("extending", "components", "test_module"),
				filepath.Join("extending", "components", "test_module_2"),
				filepath.Join("extending", "components", "test_qml"),
				filepath.Join("extending", "components", "test_qml_go"),
				"gallery", "material",
				"prop", "prop2" /*"webview"*/},

			"qt3d": []string{"audio-visualizer-qml"},

			"quick": []string{"bridge", "bridge2", "calc" /*"cookies"*/, "dialog", "dynamic",
				"hotreload", "listview", "photoviewer", "sailfish", "tableview", "translate", "view"},

			"sailfish": []string{"listview", "listview_variant"},

			"sensors": []string{"accelbubble"},

			"showcases": []string{"wallet"},

			"sql": []string{"masterdetail", "masterdetail_qml", "querymodel"},

			"uitools": []string{"calculator", "calculator_manual"},

			"virtualkeyboard": []string{"qml", "widgets", "widgets_embedded"},

			"webchannel": []string{"chatserver-go" /*"standalone" "webview"*/},

			//"webkit": []string{"browser"},

			"widgets": []string{"bridge2" /*"dropsite"*/, "graphicsscene", "line_edits", "pixel_editor",
				/*"renderer"*/ "share", "systray" /*"table"*/, "textedit",
				filepath.Join("treeview", "treeview_dual"), filepath.Join("treeview", "treeview_filelist"),
				"video_player" /*"webengine"*/, "xkcd"},
		}
	} else {
		if strings.HasPrefix(target, "sailfish") {
			examples = map[string][]string{
				"quick": []string{"sailfish"},

				"sailfish": []string{"listview", "listview_variant"},
			}
		} else {
			examples = map[string][]string{
				"showcases": []string{"wallet"},

				"qml": []string{"gallery"},

				"quick": []string{"calc"},

				"widgets": []string{"textedit"},
			}
		}
	}

	if utils.QT_VAGRANT() && target == "ios-simulator" {
		mode = "build"
	}

	for cat, list := range examples {
		for _, example := range list {
			if target != runtime.GOOS && (example == "textedit" || example == "wallet") {
				continue
			}
			if cat == "virtualkeyboard" && (strings.Contains(target, "ios") || strings.Contains(target, "android")) {
				continue
			}
			if (target == "js" || target == "wasm") &&
				(cat == "charts" || cat == "uitools" || cat == "sql" ||
					cat == "androidextras" || cat == "qt3d" || cat == "webchannel" ||
					strings.HasPrefix(example, "treeview") ||
					example == "video_player" || example == "custom_scheme") {
				continue
			}
			if example == "wallet" && target == runtime.GOOS && runtime.GOOS == "windows" {
				continue
			}

			example := filepath.Join(cat, example)

			path := filepath.Join(strings.TrimSpace(utils.GoListOptional("{{.Dir}}", "github.com/bluszcz/cutego/internal/examples", "-find", "get doc dir")), example)
			utils.Log.Infof("testing %v", example)
			deploy.Deploy(
				mode,
				target,
				path,
				docker,
				"",
				"",
				false,
				"",
				vagrant,
				vagrantsystem,
				false,
				true,
				false, //TODO: re-enable once quickcompiler files can be properly cleaned up automatically !strings.HasPrefix(target, "sailfish"),
			)
			templater.CleanupDepsForCI()
			templater.CleanupDepsForCI = func() {}
			cmd.CleanupRegisteredImportsForCI()
		}
	}
}
