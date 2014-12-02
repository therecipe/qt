package qt

import (
	"C"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"runtime"
	"strings"
)

var signalTable = make(map[string]func())

func init() { runtime.LockOSThread() }

func connectSignal(name string, signal string, function func()) {
	signalTable[name+":"+signal] = function
}

func getSignal(name string, signal string) func() {
	if signal == "destroyed" {
		defer disconnectAllSignals(name)
	}

	var signalFunction = signalTable[name+":"+signal]
	if signalFunction == nil {
		return func() {}
	}
	return signalFunction
}

func disconnectSignal(name string, signal string) {
	delete(signalTable, name+":"+signal)
}

func disconnectAllSignals(name string) {
	for entry := range signalTable {
		if strings.Contains(entry, name) {
			delete(signalTable, entry)
		}
	}
}

func randomIdentifier() string {
	var (
		length = 15
		b      = make([]byte, length)
	)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:length]
}

func DumpSignalTable() {
	fmt.Println("************ DUMP-SIGNALTABLE ************")
	for entry := range signalTable {
		fmt.Println(entry)
	}
	fmt.Println("************ DUMP-SIGNALTABLE ************")
}

func goBoolToCInt(b bool) C.int {
	if b {
		return C.int(1)
	}
	return C.int(0)
}
