package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

var data_struct struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	widget := widgets.NewQWidget(nil, 0)
	window.SetCentralWidget(widget)

	layout := widgets.NewQFormLayout(widget)
	layout.SetFieldGrowthPolicy(widgets.QFormLayout__AllNonFixedFieldsGrow)

	widgetmap := make(map[string]*widgets.QWidget)
	for i := 0; i < reflect.TypeOf(data_struct).NumField(); i++ {
		name := reflect.TypeOf(data_struct).Field(i).Tag.Get("json")

		if name != "img" {
			widgetmap[name] = widgets.NewQLineEdit(nil).QWidget_PTR()

			layout.AddRow3(name, widgetmap[name])
		} else {
			widgetmap[name] = widgets.NewQLineEdit(nil).QWidget_PTR()
			widgetmap[name+"_label"] = widgets.NewQLabel(nil, 0).QWidget_PTR()

			layout.AddRow3(name, widgetmap[name])
			layout.AddRow3(name+"_label", widgetmap[name+"_label"])
		}
	}

	button := widgets.NewQPushButton2("random xkcd", nil)
	layout.AddWidget(button)
	button.ConnectClicked(func(bool) {
		rand.Seed(time.Now().UnixNano())

		resp, err := http.Get(fmt.Sprintf("https://xkcd.com/%v/info.0.json", rand.Intn(614)))
		if err != nil {
			return
		}
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)

		json.Unmarshal(data, &data_struct)

		for i := 0; i < reflect.TypeOf(data_struct).NumField(); i++ {
			name := reflect.TypeOf(data_struct).Field(i).Tag.Get("json")

			if name != "img" {
				switch reflect.ValueOf(data_struct).Field(i).Kind() {
				case reflect.String:
					widgets.NewQLineEditFromPointer(widgetmap[name].Pointer()).SetText(reflect.ValueOf(data_struct).Field(i).String())
				case reflect.Int:
					widgets.NewQLineEditFromPointer(widgetmap[name].Pointer()).SetText(strconv.Itoa(int(reflect.ValueOf(data_struct).Field(i).Int())))
				}
			} else {
				url := reflect.ValueOf(data_struct).Field(i).String()

				widgets.NewQLineEditFromPointer(widgetmap[name].Pointer()).SetText(url)

				resp, err := http.Get(url)
				if err != nil {
					return
				}
				defer resp.Body.Close()
				data, _ := ioutil.ReadAll(resp.Body)

				pix := gui.NewQPixmap()
				pix.LoadFromData(data, uint(len(data)), "", 0)
				widgets.NewQLabelFromPointer(widgetmap[name+"_label"].Pointer()).SetPixmap(pix.Scaled2(400, 400, core.Qt__KeepAspectRatio, core.Qt__SmoothTransformation))
			}
		}
	})

	window.Show()
	widgets.QApplication_Exec()
}
