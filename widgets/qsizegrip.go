package widgets

//#include "qsizegrip.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSizeGrip struct {
	QWidget
}

type QSizeGripITF interface {
	QWidgetITF
	QSizeGripPTR() *QSizeGrip
}

func PointerFromQSizeGrip(ptr QSizeGripITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSizeGripPTR().Pointer()
	}
	return nil
}

func QSizeGripFromPointer(ptr unsafe.Pointer) *QSizeGrip {
	var n = new(QSizeGrip)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSizeGrip_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSizeGrip) QSizeGripPTR() *QSizeGrip {
	return ptr
}

func NewQSizeGrip(parent QWidgetITF) *QSizeGrip {
	return QSizeGripFromPointer(unsafe.Pointer(C.QSizeGrip_NewQSizeGrip(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QSizeGrip) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QSizeGrip_SetVisible(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QSizeGrip) DestroyQSizeGrip() {
	if ptr.Pointer() != nil {
		C.QSizeGrip_DestroyQSizeGrip(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
