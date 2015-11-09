package qt

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	_ "github.com/therecipe/qt/internal/android/cgo"
)

var signalTable = make(map[string]interface{})

func init() { runtime.LockOSThread() }

func ConnectSignal(name string, signal string, function interface{}) {
	signalTable[fmt.Sprintf("%v:%v", name, signal)] = function
}

func GetSignal(name string, signal string) interface{} {
	if signal == "destroyed" {
		defer DisconnectAllSignals(name)
	}
	return signalTable[fmt.Sprintf("%v:%v", name, signal)]
}

func DisconnectSignal(name string, signal string) {
	delete(signalTable, fmt.Sprintf("%v:%v", name, signal))
}

func DisconnectAllSignals(name string) {
	for entry := range signalTable {
		if strings.Contains(entry, name) {
			delete(signalTable, entry)
		}
	}
}

func RandomIdentifier() string {
	return fmt.Sprint(time.Now().UnixNano())
}

func DumpSignalTable() {
	println("##############################\tDUMP_SIGNALTABLE_START\t##############################")
	for entry := range signalTable {
		println(entry)
	}
	println("##############################\tDUMP_SIGNALTABLE_END\t##############################")
}

func GoBoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
