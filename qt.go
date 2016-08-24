package qt

import (
	"encoding/hex"
	"fmt"
	"runtime"
	"strings"
	"sync"

	_ "github.com/therecipe/qt/internal/android"
)

var (
	signals      = make(map[string]interface{})
	signalsMutex = new(sync.Mutex)
	ids          int
)

func init() { runtime.LockOSThread() }

func ExistsSignal(name, signal string) bool {
	signalsMutex.Lock()
	var _, exists = signals[fmt.Sprintf("%v:%v", name, signal)]
	signalsMutex.Unlock()
	return exists
}

func GetSignal(name, signal string) interface{} {
	if strings.HasSuffix(signal, ":destroyed") || strings.HasSuffix(signal, ":deleteLater") {
		defer DisconnectAllSignals(name)
	}
	var s interface{}
	signalsMutex.Lock()
	s = signals[fmt.Sprintf("%v:%v", name, signal)]
	signalsMutex.Unlock()
	return s
}

func ConnectSignal(name, signal string, function interface{}) {
	signalsMutex.Lock()
	signals[fmt.Sprintf("%v:%v", name, signal)] = function
	signalsMutex.Unlock()
}

func DisconnectSignal(name, signal string) {
	signalsMutex.Lock()
	delete(signals, fmt.Sprintf("%v:%v", name, signal))
	signalsMutex.Unlock()
}

func DisconnectAllSignals(name string) {
	for entry := range signals {
		if strings.HasPrefix(entry, fmt.Sprintf("%v:", name)) {
			signalsMutex.Lock()
			delete(signals, entry)
			signalsMutex.Unlock()
		}
	}
}

func Identifier() string {
	ids++
	return fmt.Sprint(ids)
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
		println("recovered in:", fn)
	}
}

func HexDecodeToString(in string) string {
	var out, err = hex.DecodeString(in)
	if err != nil {
		return ""
	}
	return string(out)
}
