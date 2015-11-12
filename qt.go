package qt

import (
	"runtime"
	"strconv"
	"strings"

	_ "github.com/therecipe/qt/internal/android/cgo"
)

var (
	signalTable = make(map[string]interface{})
	ids         int
)

func init() { runtime.LockOSThread() }

func ConnectSignal(name string, signal string, function interface{}) {
	signalTable[name+":"+signal] = function
}

func GetSignal(name string, signal string) interface{} {
	if signal == "destroyed" {
		defer DisconnectAllSignals(name)
	}
	return signalTable[name+":"+signal]
}

func DisconnectSignal(name string, signal string) {
	delete(signalTable, name+":"+signal)
}

func DisconnectAllSignals(name string) {
	for entry := range signalTable {
		if strings.Contains(entry, name) {
			delete(signalTable, entry)
		}
	}
}

func RandomIdentifier() string {
	ids++
	return strconv.Itoa(ids)
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
