package setup

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/cmd/deploy"
	"github.com/therecipe/qt/internal/utils"
)

func test(buildTarget string) {

	//TODO: split qtmoc test for windows
	if utils.IsCI() && runtime.GOOS != "windows" {
		utils.Log.Infof("running setup/test CI (~2min)")

		utils.RunCmd(exec.Command(filepath.Join(os.Getenv("GOPATH"), "bin", "qtmoc"), utils.GoQtPkgPath("internal", "cmd", "moc", "test")), "run qtmoc")
		utils.RunCmd(exec.Command(filepath.Join(os.Getenv("GOPATH"), "bin", "qtminimal"), "desktop", utils.GoQtPkgPath("internal", "cmd", "moc", "test")), "run qtminimal")

		var cmd = exec.Command("go", "test", "-v", "-tags=minimal")
		cmd.Dir = utils.GoQtPkgPath("internal", "cmd", "moc", "test")
		if runtime.GOOS == "windows" {
			for key, value := range map[string]string{
				"PATH":   os.Getenv("PATH"),
				"GOPATH": utils.MustGoPath(),
				"GOROOT": runtime.GOROOT(),

				"TMP":  os.Getenv("TMP"),
				"TEMP": os.Getenv("TEMP"),

				"GOOS":   runtime.GOOS,
				"GOARCH": "386",

				"CGO_ENABLED": "1",
			} {
				cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", key, value))
			}
		}
		utils.RunCmd(cmd, "run qtmoc")
	}

	var examples map[string][]string
	if utils.IsCI() {
		examples = map[string][]string{
			"androidextras": []string{"jni", "notification"},

			"canvas3d": []string{"framebuffer", "interaction", "jsonmodels",
				"quickitemtexture", "textureandlight",
				filepath.Join("threejs", "cellphone"),
				filepath.Join("threejs", "controls"),
				filepath.Join("threejs", "oneqt"),
				filepath.Join("threejs", "planets"),
			},

			"charts": []string{"audio"},

			//"grpc": []string{"hello_world","hello_world2"},

			"gui": []string{"analogclock", "rasterwindow"},

			"qml": []string{"application", "drawer_nav_x", "gallery", "material",
				"prop", "prop2" /*, "webview"*/},

			"qt3d": []string{"audio-visualizer-qml"},

			"quick": []string{"bridge", "bridge2", "calc", "dialog", "dynamic",
				"hotreload", "listview", "sailfish", "tableview", "translate", "view"},

			"sailfish": []string{"listview", "listview_variant"},

			"sql": []string{"querymodel"},

			"uitools": []string{"calculator"},

			"widgets": []string{"bridge2" /*"dropsite",*/, "graphicsscene", "line_edits", "pixel_editor",
				/*"renderer",*/ "systray" /*"table",*/, "textedit", filepath.Join("treeview", "treeview_dual"),
				filepath.Join("treeview", "treeview_filelist"), "video_player"},
		}
	} else {
		if strings.HasPrefix(buildTarget, "sailfish") {
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

	utils.Log.Infof("running setup/test %v (~5min)", buildTarget)
	for cat, list := range examples {
		for _, example := range list {
			if buildTarget != "desktop" && example == "textedit" {
				continue
			}
			var example = filepath.Join(cat, example)

			utils.Log.Infoln("testing", example)

			deploy.Deploy(&deploy.State{
				BuildMode:   "test",
				BuildTarget: strings.TrimSuffix(buildTarget, "-docker"),
				AppPath:     utils.GoQtPkgPath("internal", "examples", example),
				BuildDocker: strings.HasSuffix(buildTarget, "-docker"),
			})
		}
	}
}
