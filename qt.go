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
	signalsMutex sync.Mutex

	objects      = make(map[unsafe.Pointer]interface{})
	objectsMutex sync.Mutex

	objectsTemp      = make(map[unsafe.Pointer]unsafe.Pointer)
	objectsTempMutex sync.Mutex

	connectionTypes      = make(map[unsafe.Pointer]map[string]int64)
	connectionTypesMutex sync.Mutex

	finalizerMap      = make(map[unsafe.Pointer]struct{})
	finalizerMapMutex sync.Mutex

	//

	FuncMap      = make(map[string]interface{})
	FuncMapMutex sync.Mutex
	ItfMap       = make(map[string]interface{})
	itfMapMutex  sync.Mutex
	EnumMap      = make(map[string]int64)
	EnumMapMutex sync.Mutex
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

func GetFuncMap(n string) (o interface{}, ok bool) {
	FuncMapMutex.Lock()
	o, ok = FuncMap[n]
	FuncMapMutex.Unlock()
	return
}

func SetFuncMap(n string, v interface{}) {
	FuncMapMutex.Lock()
	FuncMap[n] = v
	FuncMapMutex.Unlock()
}

func GetItfMap(n string) (o interface{}, ok bool) {
	itfMapMutex.Lock()
	o, ok = ItfMap[n]
	itfMapMutex.Unlock()
	return
}

func SetItfMap(n string, v interface{}) {
	itfMapMutex.Lock()
	ItfMap[n] = v
	itfMapMutex.Unlock()
}

func GetEnumMap(n string) (o int64, ok bool) {
	EnumMapMutex.Lock()
	o, ok = EnumMap[n]
	EnumMapMutex.Unlock()
	return
}

func SetEnumMap(n string, v int64) {
	EnumMapMutex.Lock()
	EnumMap[n] = v
	EnumMapMutex.Unlock()
}

type ptr_itf interface {
	Pointer() unsafe.Pointer
	SetPointer(p unsafe.Pointer)
}

func SetFinalizer(ptr interface{}, f interface{}) {
	finalizerMapMutex.Lock()
	cPtr := ptr.(ptr_itf).Pointer()
	if _, ok := finalizerMap[cPtr]; !ok || f == nil || cPtr == nil {
		if cPtr != nil || (cPtr == nil && f == nil) {
			runtime.SetFinalizer(ptr, f)
		}
		if f == nil {
			delete(finalizerMap, cPtr)
		} else if cPtr != nil {
			finalizerMap[cPtr] = struct{}{}
		}
	} else {
		runtime.SetFinalizer(ptr, func(p ptr_itf) { p.SetPointer(nil) })
	}
	finalizerMapMutex.Unlock()
}

func HasFinalizer(ptr interface{}) (ok bool) {
	finalizerMapMutex.Lock()
	_, ok = finalizerMap[ptr.(ptr_itf).Pointer()]
	finalizerMapMutex.Unlock()
	return
}
