package qt

import (
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
	"unsafe"
)

var (
	Logger = log.New(os.Stderr, "", log.Ltime)

	signals      = make(map[unsafe.Pointer]map[string]unsafe.Pointer)
	signalsJNI   = make(map[string]map[string]unsafe.Pointer)
	signalsMutex = new(sync.Mutex)

	objects      = make(map[unsafe.Pointer]interface{})
	objectsMutex = new(sync.Mutex)

	objectsTemp      = make(map[unsafe.Pointer]unsafe.Pointer)
	objectsTempMutex = new(sync.Mutex)

	connectionTypes      = make(map[unsafe.Pointer]map[string]int64)
	connectionTypesMutex = new(sync.Mutex)
)

func init() { runtime.LockOSThread() }

func ExistsSignal(cPtr unsafe.Pointer, signal string) (exists bool) {
	signalsMutex.Lock()
	_, exists = signals[cPtr][signal]
	signalsMutex.Unlock()
	return
}

func LendSignal(cPtr unsafe.Pointer, signal string) (s unsafe.Pointer) {
	signalsMutex.Lock()
	s = signals[cPtr][signal]
	signalsMutex.Unlock()
	return
}

func lendSignalJNI(cPtr, signal string) (s unsafe.Pointer) {
	signalsMutex.Lock()
	s = signalsJNI[cPtr][signal]
	signalsMutex.Unlock()
	return
}

func GetSignal(cPtr interface{}, signal string) unsafe.Pointer {
	if dcPtr, ok := cPtr.(unsafe.Pointer); ok {
		if signal == "destroyed" || strings.HasPrefix(signal, "~") {
			defer DisconnectAllSignals(dcPtr, signal)
		}
		return LendSignal(dcPtr, signal)
	}
	return lendSignalJNI(cPtr.(string), signal)
}

func ConnectSignal(cPtr interface{}, signal string, function unsafe.Pointer) {
	if dcPtr, ok := cPtr.(unsafe.Pointer); ok {
		signalsMutex.Lock()
		if s, exists := signals[dcPtr]; !exists {
			signals[dcPtr] = map[string]unsafe.Pointer{signal: function}
		} else {
			s[signal] = function
		}
		signalsMutex.Unlock()
	} else {
		connectSignalJNI(cPtr.(string), signal, function)
	}
}

func connectSignalJNI(cPtr, signal string, function unsafe.Pointer) {
	signalsMutex.Lock()
	if s, exists := signalsJNI[cPtr]; !exists {
		signalsJNI[cPtr] = map[string]unsafe.Pointer{signal: function}
	} else {
		s[signal] = function
	}
	signalsMutex.Unlock()
}

func DisconnectSignal(cPtr interface{}, signal string) {
	if dcPtr, ok := cPtr.(unsafe.Pointer); ok {
		signalsMutex.Lock()
		delete(signals[dcPtr], signal)
		if len(signals[dcPtr]) == 0 {
			delete(signals, dcPtr)
		}
		signalsMutex.Unlock()
	} else {
		disconnectSignalJNI(cPtr.(string), signal)
	}
}

func disconnectSignalJNI(cPtr, signal string) {
	signalsMutex.Lock()
	delete(signalsJNI[cPtr], signal)
	if len(signalsJNI[cPtr]) == 0 {
		delete(signalsJNI, cPtr)
	}
	signalsMutex.Unlock()
}

func DisconnectAllSignals(cPtr unsafe.Pointer, signal string) {
	signalsMutex.Lock()
	if s, exists := signals[cPtr]["destroyed"]; signal != "destroyed" && exists {
		signals[cPtr] = map[string]unsafe.Pointer{"destroyed": s}
	} else {
		delete(signals, cPtr)
	}
	signalsMutex.Unlock()
}

func DumpSignals() {
	Debug("##############################\tSIGNALS_TABLE_START\t##############################")
	signalsMutex.Lock()
	for cPtr, entry := range signals {
		Debug(cPtr, entry)
	}
	signalsMutex.Unlock()
	Debug("##############################\tSIGNALS_TABLE_END\t##############################")
}

func DumpObjects() {
	Debug("##############################\tOBJECTS_TABLE_START\t##############################")
	objectsMutex.Lock()
	for cPtr, entry := range objects {
		Debug(cPtr, entry)
	}
	objectsMutex.Unlock()
	Debug("##############################\tOBJECTS_TABLE_END\t##############################")
}

func DumpTempObjects() {
	Debug("##############################\tTMP_OBJECTS_TABLE_START\t##############################")
	objectsTempMutex.Lock()
	for cPtr, entry := range objectsTemp {
		Debug(cPtr, entry)
	}
	objectsTempMutex.Unlock()
	Debug("##############################\tTMP_OBJECTS_TABLE_END\t##############################")
}

func DumpConnectionTypes() {
	Debug("##############################\tCON_MODES_TABLE_START\t##############################")
	connectionTypesMutex.Lock()
	for cPtr, entry := range connectionTypes {
		Debug(cPtr, entry)
	}
	connectionTypesMutex.Unlock()
	Debug("##############################\tCON_MODES_TABLE_END\t##############################")
}

func CountSignals() (c int) {
	signalsMutex.Lock()
	c = len(signals)
	signalsMutex.Unlock()
	return
}

func GoBoolToInt(b bool) int8 {
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
	if strings.ToLower(os.Getenv("QT_DEBUG")) == "true" || runtime.GOARCH == "js" || runtime.GOARCH == "wasm" {
		Logger.Println(fn...)
	}
}

func ClearSignals() {
	signalsMutex.Lock()
	signals = make(map[unsafe.Pointer]map[string]unsafe.Pointer)
	signalsMutex.Unlock()
}

func Register(cPtr unsafe.Pointer, gPtr interface{}) {
	objectsMutex.Lock()
	objects[cPtr] = gPtr
	objectsMutex.Unlock()
}

func Receive(cPtr unsafe.Pointer) (o interface{}, ok bool) {
	objectsMutex.Lock()
	o, ok = objects[cPtr]
	objectsMutex.Unlock()
	return
}

func Unregister(cPtr unsafe.Pointer) {
	objectsMutex.Lock()
	delete(objects, cPtr)
	objectsMutex.Unlock()
}

func RegisterTemp(cPtr unsafe.Pointer, gPtr unsafe.Pointer) {
	objectsTempMutex.Lock()
	objectsTemp[cPtr] = gPtr
	objectsTempMutex.Unlock()
}

func ReceiveTemp(cPtr unsafe.Pointer) (o unsafe.Pointer, ok bool) {
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

func RegisterConnectionType(cPtr unsafe.Pointer, signal string, mode int64) {
	connectionTypesMutex.Lock()
	if s, exists := connectionTypes[cPtr]; !exists {
		connectionTypes[cPtr] = map[string]int64{signal: mode}
	} else {
		s[signal] = mode
	}
	connectionTypesMutex.Unlock()
}

func ConnectionType(cPtr unsafe.Pointer, signal string) (m int64) {
	connectionTypesMutex.Lock()
	if s, exists := connectionTypes[cPtr]; exists {
		if lm, ok := s[signal]; ok {
			delete(s, signal)
			if len(s) == 0 {
				delete(connectionTypes, cPtr)
			}
			m = lm
		}
	}
	connectionTypesMutex.Unlock()
	return
}
