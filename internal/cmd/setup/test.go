package setup

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/cmd/deploy"
	"github.com/therecipe/qt/internal/cmd/minimal"
	"github.com/therecipe/qt/internal/cmd/moc"

	"github.com/therecipe/qt/internal/utils"
)

func Test(target string, docker, vagrant bool, vagrantsystem string) {
	if docker && target == "darwin" {
		utils.Log.Warn("darwin is currently not supported as a deploy target with docker; testing the linux deployment instead")
		target = "linux"
	}

	utils.Log.Infof("running: 'qtsetup test %v' [docker=%v] [vagrant=%v]", target, docker, vagrant)

	if utils.CI() && target == runtime.GOOS && runtime.GOOS != "windows" { //TODO: split test for windows ?
		utils.Log.Infof("running setup/test %v CI", target)

		path := utils.GoQtPkgPath("internal", "cmd", "moc", "test")

		moc.Moc(path, target, "", false, false)
		minimal.Minimal(path, target, "")

		cmd := exec.Command("go", "test", "-v", "-tags=minimal", "-ldflags=\"-s\"")
		cmd.Dir = path
		if runtime.GOOS == "windows" {
			for key, value := range map[string]string{
				"PATH":   os.Getenv("PATH"),
				"GOPATH": utils.MustGoPath(),
				"GOROOT": runtime.GOROOT(),

				"TMP":  os.Getenv("TMP"),
				"TEMP": os.Getenv("TEMP"),

				"GOOS":   "windows",
				"GOARCH": "386",

				"CGO_ENABLED": "1",
			} {
				cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", key, value))
			}
		}
		utils.RunCmd(cmd, "run \"go test\"")
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

			//"charts": []string{"audio"}, //TODO: ios, ios-simulator

			"common": []string{"qml_demo", "widgets_demo"},

			//"grpc": []string{"hello_world","hello_world2"},

			//"gui": []string{"analogclock", "openglwindow", "rasterwindow"},

			//opengl: []string{"2dpainting"},

			"qml": []string{"adding", "application", "drawer_nav_x",
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
				//filepath.Join("printslides", "cmd", "printslides"),
				"prop", "prop2" /*"quickflux", "webview"*/},

			"qt3d": []string{"audio-visualizer-qml"},

			"quick": []string{"bridge", "bridge2", "calc", "dialog", "dynamic",
				"hotreload", "listview", "sailfish", "tableview", "translate", "view"},

			"sailfish": []string{"listview", "listview_variant"},

			"sql": []string{"masterdetail", "masterdetail_qml", "querymodel"},

			"uitools": []string{"calculator"},

			"webchannel": []string{"chatserver-go" /*"standalone" "webview"*/},

			"widgets": []string{"bridge2" /*"dropsite"*/, "graphicsscene", "line_edits", "pixel_editor",
				/*"renderer"*/ "share", "systray" /*"table"*/, "textedit", filepath.Join("treeview", "treeview_dual"),
				filepath.Join("treeview", "treeview_filelist"), "video_player" /*"webengine"*/},
		}
	} else {
		if strings.HasPrefix(target, "sailfish") {
			examples = map[string][]string{
				"quick": []string{"sailfish"},

				"sailfish": []string{"listview", "listview_variant"},
			}
		} else {
			examples = map[string][]string{
				"qml": []string{"application", "drawer_nav_x", "gallery"},

				"quick": []string{"calc"},

				"widgets": []string{"line_edits", "pixel_editor", "textedit"},
			}
		}
	}

	if utils.QT_VAGRANT() && target == "ios-simulator" {
		mode = "build"
	}

	for cat, list := range examples {
		for _, example := range list {
			if target != runtime.GOOS && example == "textedit" {
				continue
			}
			example := filepath.Join(cat, example)
			path := utils.GoQtPkgPath("internal", "examples", example)
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
			)
		}
	}
}
