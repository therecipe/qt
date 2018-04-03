package controller

import (
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/therecipe/qt/core"

	"github.com/NebulousLabs/Sia/modules"

	"github.com/therecipe/qt/internal/examples/showcases/sia/controller"
	"github.com/therecipe/qt/internal/examples/showcases/sia/files/model"
)

var FilesController *filesController

type filesController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ *model.FilesModel `property:"model"`
}

func (c *filesController) init() {
	FilesController = c

	c.SetModel(model.NewFilesModel(nil))

	go c.loop()
}

func (c *filesController) loop() {
	for range time.NewTicker(1 * time.Second).C {

		rf, errFiles := controller.Client.RenterFilesGet()

		rdq, errDowload := controller.Client.RenterDownloadsGet()

		if (errFiles != nil || errDowload != nil) && !DEMO {
			println(errFiles.Error(), errDowload.Error())
		} else {
			files := make([]model.File, len(rf.Files))

			sort.Sort(bySiaPath(rf.Files))
			for i, f := range rf.Files {
				files[i] = model.File{
					Name:          f.SiaPath,
					Size:          filesizeUnits(int64(f.Filesize)),
					Redundancy:    fmt.Sprintf("%1.1fx", f.Redundancy),
					Available:     f.Available,
					ProgressText:  fmt.Sprintf("%1.2f%%", f.UploadProgress),
					ProgressValue: f.UploadProgress,
				}
			}

			for _, df := range rdq.Downloads {
				if df.Received == df.Filesize {
					continue
				}
				for i, f := range files {
					if df.SiaPath != f.Name {
						continue
					}
					pct := 100 * float64(df.Received) / float64(df.Filesize)
					files[i].Available = false
					files[i].ProgressText = fmt.Sprintf("%1.2f%%", pct)
					files[i].ProgressValue = pct
					files[i].Error = df.Error
					if df.Error != "" {
						println("Couldn't recover", df.SiaPath, "->", df.Error)
					}
				}
			}

			if DEMO {
				var df []model.File
				json.Unmarshal([]byte(DEMO_FILES), &df)
				c.Model().UpdateWith(df)
			} else {
				c.Model().UpdateWith(files)
			}
		}

	}
}

// bySiaPath implements sort.Interface for [] modules.FileInfo based on the
// SiaPath field.
type bySiaPath []modules.FileInfo

func (s bySiaPath) Len() int           { return len(s) }
func (s bySiaPath) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s bySiaPath) Less(i, j int) bool { return s[i].SiaPath < s[j].SiaPath }

// filesize returns a string that displays a filesize in human-readable units.
func filesizeUnits(size int64) string {
	if size == 0 {
		return "0 B"
	}
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	i := int(math.Log10(float64(size)) / 3)
	return fmt.Sprintf("%.*f %s", i, float64(size)/math.Pow10(3*i), sizes[i])
}
