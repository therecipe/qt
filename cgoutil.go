package qt

import (
	"C"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"runtime"
	"strings"
)

var signalTable = make(map[string]interface{})

func init() { runtime.LockOSThread() }

func ConnectSignal(name string, signal string, function interface{}) {
	signalTable[name+":"+signal] = function
}

func GetSignal(name string, signal string) interface{} {
	if signal == "destroyed" {
		defer disconnectAllSignals(name)
	}
	return signalTable[name+":"+signal]
}

func DisconnectSignal(name string, signal string) {
	delete(signalTable, name+":"+signal)
}

func disconnectAllSignals(name string) {
	for entry := range signalTable {
		if strings.Contains(entry, name) {
			delete(signalTable, entry)
		}
	}
}

func RandomIdentifier() string {
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

func GoBoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
