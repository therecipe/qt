package qt

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
	"unsafe"
)

var (
	Logger = log.New(os.Stderr, "", log.Ltime)

	signals      = make(map[string]interface{}) //TODO: unsafe.Pointer
	signalsMutex = new(sync.Mutex)

	objects      = make(map[string]interface{}) //TODO: unsafe.Pointer
	objectsMutex = new(sync.Mutex)

	objectsTemp      = make(map[unsafe.Pointer]interface{})
	objectsTempMutex = new(sync.Mutex)
)

func init() { runtime.LockOSThread() }

func ExistsSignal(name, signal string) bool {
	signalsMutex.Lock()
	_, exists := signals[fmt.Sprintf("%v:%v", name, signal)]
	signalsMutex.Unlock()
	return exists
}

func LendSignal(name, signal string) interface{} {
	var s interface{}
	signalsMutex.Lock()
	s = signals[fmt.Sprintf("%v:%v", name, signal)]
	signalsMutex.Unlock()
	return s
}

func GetSignal(name, signal string) interface{} {
	if signal == "destroyed" || strings.HasPrefix(signal, "~") {
		defer DisconnectAllSignals(name, signal)
	}
	return LendSignal(name, signal)
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

func DisconnectAllSignals(name, signal string) {
	signalsMutex.Lock()
	for entry := range signals {
		if (signal == "destroyed" || !strings.HasSuffix(entry, fmt.Sprintf(":%v", "destroyed"))) && strings.HasPrefix(entry, fmt.Sprintf("%v:", name)) {
			delete(signals, entry)
		}
	}
	signalsMutex.Unlock()
	if signal == "destroyed" {
		Unregister(name)
	}
}

func DumpSignals() {
	Debug("##############################\tSIGNALSTABLE_START\t##############################")
	signalsMutex.Lock()
	for entry := range signals {
		Debug(entry)
	}
	signalsMutex.Unlock()
	Debug("##############################\tSIGNALSTABLE_END\t##############################")
}

func CountSignals() int {
	var c int
	signalsMutex.Lock()
	c = len(signals)
	signalsMutex.Unlock()
	return c
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

func ClearSignals() {
	signalsMutex.Lock()
	signals = make(map[string]interface{})
	signalsMutex.Unlock()
}

func Register(cPtr unsafe.Pointer, gPtr interface{}) {
	objectsMutex.Lock()
	objects[fmt.Sprint(cPtr)] = gPtr
	objectsMutex.Unlock()
}

func Receive(cPtr unsafe.Pointer) (o interface{}, ok bool) {
	objectsMutex.Lock()
	o, ok = objects[fmt.Sprint(cPtr)]
	objectsMutex.Unlock()
	return
}

func Unregister(cPtr string) {
	objectsMutex.Lock()
	delete(objects, cPtr)
	objectsMutex.Unlock()
}

func RegisterTemp(cPtr unsafe.Pointer, gPtr interface{}) {
	objectsTempMutex.Lock()
	objectsTemp[cPtr] = gPtr
	objectsTempMutex.Unlock()
}

func ReceiveTemp(cPtr unsafe.Pointer) (o interface{}, ok bool) {
	objectsTempMutex.Lock()
	o, ok = objectsTemp[cPtr]
	objectsTempMutex.Unlock()
	return
}

func UnregisterTemp(cPtr unsafe.Pointer) {
	objectsTempMutex.Lock()
	delete(objectsTemp, cPtr)
	objectsTempMutex.Unlock()
}
