package qt

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"

	_ "github.com/therecipe/qt/internal/android/cgo"
)

var (
	Logger = log.New(os.Stderr, "", log.Ltime)

	signals      = make(map[string]interface{})
	signalsMutex = new(sync.Mutex)

	ids int
)

func init() { runtime.LockOSThread() }

func ExistsSignal(name, signal string) bool {
	signalsMutex.Lock()
	var _, exists = signals[fmt.Sprintf("%v:%v", name, signal)]
	signalsMutex.Unlock()
	return exists
}

func GetSignal(name, signal string) interface{} {
	if strings.HasSuffix(signal, ":destroyed") || strings.HasSuffix(signal, ":deleteLater") || strings.Contains(signal, ":~") {
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
	Debug("##############################\tSIGNALSTABLE_START\t##############################")
	for entry := range signals {
		Debug(entry)
	}
	Debug("##############################\tSIGNALSTABLE_END\t##############################")
}

func GoBoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func Recover(fn string) {
	if recover() != nil {
		Debug("RECOVERED:", fn)
	}
}

func Debug(fn ...interface{}) {
	if strings.ToLower(os.Getenv("QT_DEBUG")) == "true" {
		Logger.Println(fn...)
	}
}
