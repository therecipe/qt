package multimedia

//#include "qmediacontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMediaControl struct {
	core.QObject
}

type QMediaControl_ITF interface {
	core.QObject_ITF
	QMediaControl_PTR() *QMediaControl
}

func PointerFromQMediaControl(ptr QMediaControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaControlFromPointer(ptr unsafe.Pointer) *QMediaControl {
	var n = new(QMediaControl)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QMediaControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaControl) QMediaControl_PTR() *QMediaControl {
	return ptr
}

func (ptr *QMediaControl) DestroyQMediaControl() {
	if ptr.Pointer() != nil {
		C.QMediaControl_DestroyQMediaControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
