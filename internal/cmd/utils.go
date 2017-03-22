package cmd

import (
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/therecipe/qt/internal/utils"
)

func GetImports(path string, level int) []string {
	utils.Log.WithField("path", path).WithField("level", level).Debug("imports")

	var imports []string

	level++
	if level > 5 {
		return imports
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		utils.Log.WithError(err).Fatal("failed to read dir")
	}
	var fileList []string
	for _, f := range files {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".go" {
			fileList = append(fileList, filepath.Join(path, f.Name()))
		}
	}

	for _, f := range fileList {
		file, err := parser.ParseFile(token.NewFileSet(), f, nil, 0)
		if err != nil {
			utils.Log.WithError(err).Debugln("failed to parse", f)
			continue
		}
		for _, i := range file.Imports {
			if strings.Contains(i.Path.Value, "github.com/therecipe/qt") && !strings.Contains(i.Path.Value, "qt/internal") {
				continue
			}

			for _, gopath := range strings.Split(os.Getenv("GOPATH"), string(os.PathListSeparator)) {
				path := filepath.Join(gopath, "src", strings.Replace(i.Path.Value, "\"", "", -1))
				if utils.ExistsDir(path) {
					var has bool
					for _, i := range imports {
						if i == path {
							has = true
							break
						}
					}
					if has {
						continue
					}
					imports = append(imports, path)
					for _, path := range GetImports(path, level) {
						var has bool
						for _, i := range imports {
							if i == path {
								has = true
								break
							}
						}
						if !has {
							imports = append(imports, path)
						}
					}
				}
			}

		}
	}

	return imports
}
