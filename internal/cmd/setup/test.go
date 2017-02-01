package setup

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/therecipe/qt/internal/cmd/deploy"
	"github.com/therecipe/qt/internal/utils"
)

func test(buildTarget string) {

	if utils.IsCI() {
		utils.Log.Infof("running setup/test CI (~2min)")

		utils.RunCmd(exec.Command(filepath.Join(os.Getenv("GOPATH"), "qtmoc"), utils.GoQtPkgPath("internal", "cmd", "moc", "test")), "run qtmoc")
		utils.RunCmd(exec.Command(filepath.Join(os.Getenv("GOPATH"), "qtminimal"), "desktop", utils.GoQtPkgPath("internal", "cmd", "moc", "test")), "run qtminimal")

		var cmd = exec.Command("go", "test", "-v", "-tags=minimal")
		cmd.Dir = utils.GoQtPkgPath("internal", "cmd", "moc", "test")
		utils.RunCmd(cmd, "run qtmoc")
	}

	var examples map[string][]string
	if utils.IsCI() {
		examples = map[string][]string{
			"androidextras": []string{"jni"},

			//"grpc": []string{"hello_world","hello_world2"},

			"qml": []string{"application", "drawer_nav_x", "gallery", "material",
				"prop", "prop2"},

			"quick": []string{"bridge", "bridge2", "calc", "dialog", "dynamic",
				"hotreload", "listview", "sailfish", "tableview", "translate", "view"},

			"sql": []string{"querymodel"},

			"uitools": []string{"calculator"},

			"widgets": []string{"bridge2" /*"dropsite",*/, "graphicsscene", "line_edits", "pixel_editor",
				/*"renderer",*/ "systray", "table", "textedit", filepath.Join("treeview", "treeview_dual"),
				filepath.Join("treeview", "treeview_filelist"), "video_player"},
		}
	} else {
		examples = map[string][]string{
			"qml": []string{"application", "drawer_nav_x", "gallery"},

			"quick": []string{"calc"},

			"sql": []string{"querymodel"},

			"widgets": []string{"line_edits", "pixel_editor", "textedit", "video_player"},
		}
	}

	utils.Log.Infof("running setup/test %v (~5min)", buildTarget)
	for cat, list := range examples {
		for _, example := range list {
			var example = filepath.Join(cat, example)

			if strings.HasPrefix(buildTarget, "sailfish") && strings.HasPrefix(example, "widgets") {
				utils.Log.Infoln("skipping example", example)
				continue
			}
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
