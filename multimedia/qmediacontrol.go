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

type QMediaControlITF interface {
	core.QObjectITF
	QMediaControlPTR() *QMediaControl
}

func PointerFromQMediaControl(ptr QMediaControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaControlPTR().Pointer()
	}
	return nil
}

func QMediaControlFromPointer(ptr unsafe.Pointer) *QMediaControl {
	var n = new(QMediaControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMediaControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaControl) QMediaControlPTR() *QMediaControl {
	return ptr
}

func (ptr *QMediaControl) DestroyQMediaControl() {
	if ptr.Pointer() != nil {
		C.QMediaControl_DestroyQMediaControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
