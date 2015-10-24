package widgets

//#include "qdial.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QDial struct {
	QAbstractSlider
}

type QDialITF interface {
	QAbstractSliderITF
	QDialPTR() *QDial
}

func PointerFromQDial(ptr QDialITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDialPTR().Pointer()
	}
	return nil
}

func QDialFromPointer(ptr unsafe.Pointer) *QDial {
	var n = new(QDial)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QDial_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDial) QDialPTR() *QDial {
	return ptr
}

func (ptr *QDial) NotchSize() int {
	if ptr.Pointer() != nil {
		return int(C.QDial_NotchSize(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDial) NotchesVisible() bool {
	if ptr.Pointer() != nil {
		return C.QDial_NotchesVisible(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDial) SetNotchesVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QDial_SetNotchesVisible(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QDial) SetWrapping(on bool) {
	if ptr.Pointer() != nil {
		C.QDial_SetWrapping(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QDial) Wrapping() bool {
	if ptr.Pointer() != nil {
		return C.QDial_Wrapping(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func NewQDial(parent QWidgetITF) *QDial {
	return QDialFromPointer(unsafe.Pointer(C.QDial_NewQDial(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QDial) DestroyQDial() {
	if ptr.Pointer() != nil {
		C.QDial_DestroyQDial(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
