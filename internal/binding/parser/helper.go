package parser

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"sort"
	"strings"
	"sync"

	"github.com/therecipe/qt/internal/utils"
)

const (
	SIGNAL = "signal"
	SLOT   = "slot"
	PROP   = "prop"

	IMPURE = "impure"
	PURE   = "pure"

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

	MOC = "moc"
)

func IsPackedList(v string) bool {
	return (strings.HasPrefix(v, "QList<") ||
		//TODO: QLinkedList
		strings.HasPrefix(v, "QVector<") ||
		strings.HasPrefix(v, "QStack<") ||
		strings.HasPrefix(v, "QQueue<")) &&
		//TODO: QSet

		strings.Count(v, "<") == 1 //TODO:
}

func UnpackedList(v string) string {
	return CleanValue(UnpackedListDirty(v))
}

func UnpackedListDirty(v string) string {
	return strings.Split(strings.Split(v, "<")[1], ">")[0]
}

func IsPackedMap(v string) bool {
	return (strings.HasPrefix(v, "QMap<") ||
		strings.HasPrefix(v, "QMultiMap<") ||
		strings.HasPrefix(v, "QHash<") ||
		strings.HasPrefix(v, "QMultiHash<")) &&

		strings.Count(v, "<") == 1 //TODO:
}

func UnpackedMap(v string) (string, string) {
	var splitted = strings.Split(UnpackedList(v), ",")
	return splitted[0], splitted[1]
}

func UnpackedMapDirty(v string) (string, string) {
	var splitted = strings.Split(UnpackedListDirty(v), ",")
	return splitted[0], splitted[1]
}

func CleanValue(v string) string {
	if IsPackedList(cleanValueUnsafe(v)) || IsPackedMap(cleanValueUnsafe(v)) {
		var inside = strings.Split(strings.Split(v, "<")[1], ">")[0]
		return strings.Replace(cleanValueUnsafe(v), strings.Split(strings.Split(cleanValueUnsafe(v), "<")[1], ">")[0], inside, -1)
	}
	return cleanValueUnsafe(v)
}

func cleanValueUnsafe(v string) string {
	for _, b := range []string{"*", "const", "&amp", "&", ";"} {
		v = strings.Replace(v, b, "", -1)
	}
	return strings.TrimSpace(v)
}

func CleanName(name, value string) string {
	switch name {
	case
		"type",
		"func",
		"range",
		"string",
		"int",
		"map",
		"const",
		"interface",
		"select",
		"strings",
		"new",
		"signal",
		"ptr",
		"register",
		"forever",
		"len":
		{
			return name[:len(name)-2]
		}

	case "":
		{
			var v = strings.Replace(CleanValue(value), ".", "", -1)
			if len(v) >= 3 {
				return fmt.Sprintf("v%v", strings.ToLower(v[:2]))
			} else {
				return fmt.Sprintf("v%v", strings.ToLower(v))
			}
		}

	case "f", "fmt":
		{
			return "fo"
		}
	}

	return name
}

//TODO: remove global
var LibDeps = map[string][]string{
	"Core":          {"Widgets", "Gui", "Svg"}, //Widgets, Gui //Svg (needed because it's more convenient)
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
	"Help":          {"Sql", "Network", "Widgets", "Gui", "Core"},                 //Sql + CLucene + Network (needed for static linking ios)
	"Location":      {"Positioning", "Quick", "Gui", "Core"},
	"ScriptTools":   {"Script", "Widgets", "Core"}, //Script, Widgets
	"UiTools":       {"Widgets", "Gui", "Core"},
	"X11Extras":     {"Gui", "Core"},
	"WinExtras":     {"Widgets", "Gui", "Core"},
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
	//"Quick2DRenderer":   {}, //TODO: uncomment

	"Speech":         {"Core"},
	"QuickControls2": {"Quick", "QuickWidgets", "Widgets", "Network", "Qml", "Gui", "Core"}, //Quick, QuickWidgets, Widgets, Network, Qml, Gui (needed for static linking ios)

	"Sailfish": {"Core"},
	"WebView":  {"Core"},

	"NetworkAuth":   {"Network", "Gui", "Core"},
	"RemoteObjects": {"Network", "Core"},

	"WebKit": {"WebKitWidgets", "Multimedia", "Positioning", "Widgets", "Sql", "Network", "Gui", "Sensors", "Core"},

	MOC:         make([]string, 0),
	"build_ios": {"Core", "Gui", "Network", "Sql", "Xml", "Nfc", "Script", "Sensors", "Positioning", "Widgets", "Qml", "WebSockets", "XmlPatterns", "Bluetooth", "WebChannel", "Svg", "Multimedia", "Quick", "Help", "Location", "ScriptTools", "MultimediaWidgets", "UiTools", "PrintSupport", "WebView"},
}

func ShouldBuildForTarget(module, target string) bool {

	switch target {
	case "windows":
		if runtime.GOOS == "windows" {
			return true
		}
		switch module {
		case "WebEngine", "Designer", "Speech", "WebView":
			return false
		}
		if strings.HasSuffix(module, "Extras") && module != "WinExtras" {
			return false
		}

	case "android", "android-emulator":
		switch module {
		case "DBus", "WebEngine", "Designer", "PrintSupport": //TODO: PrintSupport
			return false
		}
		if strings.HasSuffix(module, "Extras") && module != "AndroidExtras" {
			return false
		}

	case "ios", "ios-simulator":
		switch module {
		case "DBus", "WebEngine", "SerialPort", "SerialBus", "Designer", "PrintSupport": //TODO: PrintSupport
			return false
		}
		if strings.HasSuffix(module, "Extras") {
			return false
		}

	case "sailfish", "sailfish-emulator", "asteroid":
		if !IsWhiteListedSailfishLib(module) {
			return false
		}

	case "rpi1", "rpi2", "rpi3":
		switch module {
		case "WebEngine", "Designer":
			return false
		}
		if strings.HasSuffix(module, "Extras") {
			return false
		}
	}

	return true
}

func IsWhiteListedSailfishLib(name string) bool {
	switch name {
	case "Sailfish", "Core", "Quick", "Qml", "Network", "Gui", "Concurrent", "Multimedia", "Sql", "Svg", "XmlPatterns", "Xml", "DBus", "WebKit", "Sensors", "Positioning":
		return true

	default:
		return false
	}
}

func GetLibs() []string {
	libs := []string{
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
		//"X11Extras", //TODO:
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

		"Speech",
		"QuickControls2",

		"Sailfish",
		"WebView",

		//"NetworkAuth", //TODO:
		"RemoteObjects",

		"WebKit",
	}

	for i := len(libs) - 1; i >= 0; i-- {
		switch {
		case !(runtime.GOOS == "darwin" || runtime.GOOS == "linux") && (libs[i] == "WebEngine" || libs[i] == "WebView"),
			runtime.GOOS != "windows" && libs[i] == "WinExtras",
			runtime.GOOS != "darwin" && libs[i] == "MacExtras",
			runtime.GOOS != "linux" && libs[i] == "X11Extras":
			libs = append(libs[:i], libs[i+1:]...)

		case utils.QT_VERSION_NUM() < 5080 && libs[i] == "Speech":
			libs = append(libs[:i], libs[i+1:]...)

		case (utils.QT_VERSION_NUM() < 5090 || utils.QT_MXE()) && (libs[i] == "NetworkAuth" || libs[i] == "RemoteObjects"):
			libs = append(libs[:i], libs[i+1:]...)

		case !utils.QT_WEBKIT() && libs[i] == "WebKit":
			libs = append(libs[:i], libs[i+1:]...)

		case (utils.QT_MSYS2() || utils.QT_PKG_CONFIG()) && libs[i] == "Purchasing":
			libs = append(libs[:i], libs[i+1:]...)
		}
	}
	return libs
}

var (
	getCustomLibsCache      = make(map[string]string)
	getCustomLibsCacheMutex = new(sync.Mutex)
)

func GetCustomLibs(target, tags string) map[string]string {

	/*TODO: cycle dep of cmd.BuildEnv
	env, tags, _, _ := cmd.BuildEnv(target, "", "")
	if tagsCustom != "" {
		tags = append(tags, strings.Split(tagsCustom, " ")...)
	}
	*/

	wg := new(sync.WaitGroup)
	out := make(map[string]string)
	outMutex := new(sync.Mutex)

	for _, c := range State.ClassMap {
		if c.Pkg == "" || !c.IsSubClassOfQObject() {
			continue
		}

		wg.Add(1)
		go func(c *Class) {
			getCustomLibsCacheMutex.Lock()
			path, ok := getCustomLibsCache[c.Pkg]
			getCustomLibsCacheMutex.Unlock()

			if !ok {
				cmd := exec.Command("go", "list", "-e", "-f", "{{.ImportPath}}", fmt.Sprintf("-tags=\"%v\"", tags))
				cmd.Dir = c.Pkg

				/*TODO: cycle dep of cmd.BuildEnv
				for k, v := range env {
					cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", k, v))
				}
				*/

				path = strings.TrimSpace(utils.RunCmd(cmd, "get import path"))
				getCustomLibsCacheMutex.Lock()
				getCustomLibsCache[c.Pkg] = path
				getCustomLibsCacheMutex.Unlock()
			}

			outMutex.Lock()
			out[c.Module] = path
			outMutex.Unlock()

			wg.Done()
		}(c)
	}

	wg.Wait()

	return out
}

func Dump() {
	for _, c := range State.ClassMap {
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

func SortedClassNamesForModule(module string, template bool) []string {
	var output = make([]string, 0)
	for _, class := range State.ClassMap {
		for _, pm := range strings.Split(module, ",") {
			if class.Module == pm {
				output = append(output, class.Name)
			}
		}
	}
	sort.Stable(sort.StringSlice(output))

	if (module == MOC || strings.HasPrefix(module, "custom_")) && template {
		items := make(map[string]string)
		for _, cn := range output {
			if class, ok := State.ClassMap[cn]; ok {
				items[cn] = class.Bases
			}
		}

		tmpOutput := make([]string, 0)

		for item, dep := range items {
			depClass, ok := State.ClassMap[dep]
			if !ok {
				delete(items, item)
				continue
			}

			//filter out everything that has no moc dep
			if !(depClass.Module == MOC || strings.HasPrefix(depClass.Module, "custom_")) {
				tmpOutput = append(tmpOutput, item)
				delete(items, item)
				continue
			}
		}

		for len(items) > 0 {
			for item, dep := range items {

				depClass, ok := State.ClassMap[dep]
				if !ok {
					delete(items, item)
					continue
				}

				//filter out everything that has the fewest dependencies
				if hasFewestDeps(items, depClass.Name) {
					tmpOutput = append(tmpOutput, item)
					delete(items, item)
					continue
				}

				//filter out everything that has resolved dep
				for _, key := range tmpOutput {
					if key == dep {
						tmpOutput = append(tmpOutput, item)
						delete(items, item)
						break
					}
				}

			}
		}
		output = tmpOutput
	}

	return output
}

func hasFewestDeps(i map[string]string, check string) bool {
	dif := 100
	var base string
	for _, v := range i {
		c, ok := State.ClassMap[v]
		if !ok {
			continue
		}
		if ndif := len(c.GetAllBases()); ndif < dif {
			dif = ndif
			base = v
		}
	}
	return base == check
}

func SortedClassesForModule(module string, template bool) []*Class {
	var (
		classNames = SortedClassNamesForModule(module, template)
		output     = make([]*Class, len(classNames))
	)
	for i, name := range classNames {
		output[i] = State.ClassMap[name]
	}
	return output
}
