package qt

import (
	"log"
	"runtime"
	"strconv"
	"strings"

	_ "github.com/therecipe/qt/internal/android/cgo"
)

var (
	signals = make(map[string]interface{})
	ids     int
)

func init() { runtime.LockOSThread() }

func ConnectSignal(name string, signal string, function interface{}) {
	signals[name+":"+signal] = function
}

func GetSignal(name string, signal string) interface{} {
	if signal == "destroyed" {
		defer DisconnectAllSignals(name)
	}
	return signals[name+":"+signal]
}

func DisconnectSignal(name string, signal string) {
	delete(signals, name+":"+signal)
}

func DisconnectAllSignals(name string) {
	for entry := range signals {
		if strings.Contains(entry, name) {
			delete(signals, entry)
		}
	}
}

func Identifier() string {
	ids++
	return strconv.Itoa(ids)
}

func DumpSignals() {
	println("##############################\tDUMP_SIGNALTABLE_START\t##############################")
	for entry := range signals {
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

func Recovering(fn string) {
	if recover() != nil {
		log.Printf("recovered in: %v", fn)
	}
}
