package parser

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/utils"
)

const (
	SIGNAL = "signal"
	SLOT   = "slot"

	IMPURE = "impure"
	PURE   = "pure"

	MOC              = "main"
	PLAIN            = "plain"
	CONSTRUCTOR      = "constructor"
	COPY_CONSTRUCTOR = "copy-constructor"
	MOVE_CONSTRUCTOR = "move-constructor"
	DESTRUCTOR       = "destructor"

	CONNECT    = "Connect"
	DISCONNECT = "Disconnect"
	CALLBACK   = "callback"

	GETTER = "getter"
	SETTER = "setter"

	VOID = "void"

	TILDE = "~"
)

func IsPackedList(v string) bool {
	return (strings.HasPrefix(v, "QList<") ||
		strings.HasPrefix(v, "QVector<") ||
		strings.HasPrefix(v, "QStack<") ||
		strings.HasPrefix(v, "QQueue<")) &&

		strings.Count(v, "<") == 1 &&
		!strings.Contains(v, ":") &&
		CurrentState.ClassMap[UnpackedList(v)] != nil
}

func UnpackedList(v string) string {
	var CleanValue = func(v string) string {
		for _, b := range []string{"*", "const", "&amp", "&", ";"} {
			v = strings.Replace(v, b, "", -1)
		}
		return strings.TrimSpace(v)
	}
	return CleanValue(strings.Split(strings.Split(v, "<")[1], ">")[0])
}

var LibDeps = map[string][]string{
	"Core":          {"Widgets", "Gui"}, //Widgets, Gui
	"AndroidExtras": {"Core"},
	"Gui":           {"Widgets", "Core"}, //Widgets
	"Network":       {"Core"},
	"Xml":           {"XmlPatterns", "Core"}, //XmlPatterns
	"DBus":          {"Core"},
	"Nfc":           {"Core"},
	"Script":        {"Core"},
	"Sensors":       {"Core"},
	"Positioning":   {"Core"},
	"Widgets":       {"Gui", "Core"},
	"Sql":           {"Widgets", "Gui", "Core"}, //Widgets, Gui
	"MacExtras":     {"Gui", "Core"},
	"Qml":           {"Network", "Core"},
	"WebSockets":    {"Network", "Core"},
	"XmlPatterns":   {"Network", "Core"},
	"Bluetooth":     {"Core"},
	"WebChannel":    {"Network", "Qml", "Core"}, //Network (needed for static linking ios)
	"Svg":           {"Widgets", "Gui", "Core"},
	"Multimedia":    {"MultimediaWidgets", "Widgets", "Network", "Gui", "Core"},   //MultimediaWidgets, Widgets
	"Quick":         {"QuickWidgets", "Widgets", "Network", "Qml", "Gui", "Core"}, //QuickWidgets, Widgets, Network (needed for static linking ios)
	"Help":          {"Sql", "CLucene", "Network", "Widgets", "Gui", "Core"},      //Sql + CLucene + Network (needed for static linking ios)
	"Location":      {"Positioning", "Quick", "Gui", "Core"},
	"ScriptTools":   {"Script", "Widgets", "Core"}, //Script, Widgets
	"UiTools":       {"Widgets", "Gui", "Core"},
	"X11Extras":     {"Gui", "Core"},
	"WinExtras":     {"Gui", "Core"},
	"WebEngine":     {"Widgets", "WebEngineWidgets", "WebChannel", "Network", "WebEngineCore", "Quick", "Gui", "Qml", "Core"}, //Widgets, WebEngineWidgets, WebChannel, Network
	"TestLib":       {"Widgets", "Gui", "Core"},                                                                               //Widgets, Gui
	"SerialPort":    {"Core"},
	"SerialBus":     {"Core"},
	"PrintSupport":  {"Widgets", "Gui", "Core"},
	//"PlatformHeaders": []string{}, //TODO: uncomment
	"Designer": {"UiPlugin", "Widgets", "Gui", "Xml", "Core"},
	"Scxml":    {"Network", "Qml", "Core"}, //Network (needed for static linking ios)
	"Gamepad":  {"Gui", "Core"},

	"Purchasing":        {"Core"},
	"DataVisualization": {"Gui", "Core"},
	"Charts":            {"Widgets", "Gui", "Core"},
	//"Quick2DRenderer":   {},                         //TODO: uncomment

	"Sailfish": {"Core"},

	MOC:         make([]string, 0),
	"build_ios": {"Core", "Gui", "Network", "Sql", "Xml", "DBus", "Nfc", "Script", "Sensors", "Positioning", "Widgets", "Qml", "WebSockets", "XmlPatterns", "Bluetooth", "WebChannel", "Svg", "Multimedia", "Quick", "Help", "Location", "ScriptTools", "MultimediaWidgets", "UiTools", "PrintSupport"},
}

var Libs = []string{
	"Core",
	"AndroidExtras",
	"Gui",
	"Network",
	"Xml",
	"DBus",
	"Nfc",
	"Script", //depreached (planned) in 5.6
	"Sensors",
	"Positioning",
	"Widgets",
	"Sql",
	"MacExtras",
	"Qml",
	"WebSockets",
	"XmlPatterns",
	"Bluetooth",
	"WebChannel",
	"Svg",
	"Multimedia",
	"Quick",
	"Help",
	"Location",
	"ScriptTools", //depreached (planned) in 5.6
	"UiTools",
	"X11Extras",
	"WinExtras",
	"WebEngine",
	"TestLib",
	"SerialPort",
	"SerialBus",
	"PrintSupport",
	//"PlatformHeaders", //missing imports/guards
	"Designer",
	"Scxml",
	"Gamepad",

	"Purchasing",        //GPLv3 & LGPLv3
	"DataVisualization", //GPLv3
	"Charts",            //GPLv3
	//"Quick2DRenderer",   //GPLv3

	"Sailfish",
}

func GetLibs() []string {
	for i := len(Libs) - 1; i >= 0; i-- {
		switch {
		case
			!(runtime.GOOS == "darwin" || runtime.GOOS == "linux") && Libs[i] == "WebEngine",
			runtime.GOOS != "windows" && Libs[i] == "WinExtras",
			runtime.GOOS != "darwin" && Libs[i] == "MacExtras",
			runtime.GOOS != "linux" && Libs[i] == "X11Extras":
			{
				Libs = append(Libs[:i], Libs[i+1:]...)
			}
		}
	}

	return Libs
}

func Dump() {
	for _, c := range CurrentState.ClassMap {
		var bb = new(bytes.Buffer)
		defer bb.Reset()

		fmt.Fprint(bb, "funcs\n\n")
		for _, f := range c.Functions {
			fmt.Fprintln(bb, f)
		}

		fmt.Fprint(bb, "\n\nenums\n\n")
		for _, e := range c.Enums {
			fmt.Fprintln(bb, e)
		}

		utils.MkdirAll(utils.GoQtPkgPath("internal", "binding", "dump", c.Module))
		utils.SaveBytes(utils.GoQtPkgPath("internal", "binding", "dump", c.Module, fmt.Sprintf("%v.txt", c.Name)), bb.Bytes())
	}
}
