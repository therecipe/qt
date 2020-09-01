package mpv

/*
#cgo pkg-config: mpv
#cgo linux,amd64 LDFLAGS: -L ${SRCDIR}/mpv-examples/libmpv/qml -lmpvtest

void initMpv();
*/
import "C"
import (
	"github.com/therecipe/qt/core"
	_ "github.com/therecipe/qt/quick"
)

type stub struct{ core.QObject } //TODO: needed for linking at the moment

func InitMpv() {
	C.initMpv()
}
