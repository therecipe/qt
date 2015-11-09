package cgo

import "C"
import (
	"os"
	"strings"

	"github.com/therecipe/qt/internal/android/callfn"
)

//export callMain
func callMain(mainPC uintptr) { go callfn.CallFn(mainPC) }

//export setAndroidParamsAndEnv
func setAndroidParamsAndEnv(p *C.char, e *C.char) {
	os.Args = strings.Split(C.GoString(p), "\t")

	for _, env := range strings.Split(C.GoString(e), "\t") {
		if strings.Contains(env, "=") {
			var envS = strings.Split(env, "=")
			os.Setenv(envS[0], envS[1])
		}
	}
	os.Setenv("PWD", os.Getenv("HOME"))
}
