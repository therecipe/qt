package setup

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/therecipe/qt/internal/cmd/deploy"
	"github.com/therecipe/qt/internal/utils"
)

var exampleApps = []struct {
	path        []string
	desktopOnly bool
}{
	{path: []string{"widgets", "line_edits"}},
	{path: []string{"widgets", "video_player"}},
	{path: []string{"widgets", "graphicsscene"}},
	//{path: []string{"widgets", "dropsite"}},
	//{path: []string{"widgets", "table"}},
	{path: []string{"widgets", "treeview", "treeview_dual"}},
	{path: []string{"widgets", "treeview", "treeview_filelist"}},
	{path: []string{"widgets", "bridge2"}},
	{path: []string{"widgets", "systray"}},
	//{path: []string{"widgets", "renderer"}},
	{path: []string{"widgets", "textedit"}},
	{path: []string{"quick", "bridge"}},
	{path: []string{"quick", "bridge2"}},
	{path: []string{"quick", "calc"}},
	{path: []string{"quick", "dialog"}},
	{path: []string{"quick", "sailfish"}},
	{path: []string{"quick", "translate"}},
	{path: []string{"quick", "view"}},
	{path: []string{"quick", "tableview"}},
	{path: []string{"quick", "dynamic"}},
	{path: []string{"qml", "application"}},
	{path: []string{"qml", "material"}},
	{path: []string{"qml", "prop"}},
	{path: []string{"uitools", "calculator"}},
}

func test(buildTarget string) {
	utils.Log.Infof("running setup/test %v (~10min)", buildTarget)

	skipExample := func(example string) bool {
		if (buildTarget == "sailfish" || buildTarget == "sailfish-emulator") && (!strings.Contains(example, "quick") || (example == filepath.Join("quick", "dynamic") || example == filepath.Join("quick", "tableview") || example == filepath.Join("quick", "bridge") || example == filepath.Join("quick", "dialog"))) {
			return true
		} else if !(buildTarget == "sailfish" || buildTarget == "sailfish-emulator") && example == filepath.Join("quick", "sailfish") {
			return true
		} else if buildTarget != "desktop" && example == filepath.Join("widgets", "textedit") {
			return true
		}
		return false
	}

	//TODO: cleanup
	for _, app := range exampleApps {
		var example = filepath.Join(app.path...)

		if skipExample(example) {
			utils.Log.Infoln("skipping example", example)
			continue
		}
		utils.Log.Infoln("testing", example)

		deploy.Deploy(&deploy.State{
			BuildMode: func() string {
				if strings.ToLower(os.Getenv("CI")) == "true" {
					return "build"
				}
				return "test"
			}(),
			BuildTarget: strings.TrimSuffix(buildTarget, "-docker"),
			AppPath:     filepath.Join(utils.GoQtPkgPath("internal", "examples"), example),
			BuildDocker: strings.HasSuffix(buildTarget, "-docker"),
		})
	}
}
