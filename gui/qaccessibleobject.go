package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAccessibleObject struct {
	QAccessibleInterface
}

type QAccessibleObject_ITF interface {
	QAccessibleInterface_ITF
	QAccessibleObject_PTR() *QAccessibleObject
}

func PointerFromQAccessibleObject(ptr QAccessibleObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleObject_PTR().Pointer()
	}
	return nil
}

func NewQAccessibleObjectFromPointer(ptr unsafe.Pointer) *QAccessibleObject {
	var n = new(QAccessibleObject)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QAccessibleObject_") {
		n.SetObjectNameAbs("QAccessibleObject_" + qt.Identifier())
	}
	return n
}

func (ptr *QAccessibleObject) QAccessibleObject_PTR() *QAccessibleObject {
	return ptr
}

func (ptr *QAccessibleObject) ChildAt(x int, y int) *QAccessibleInterface {
	defer qt.Recovering("QAccessibleObject::childAt")

	if ptr.Pointer() != nil {
		return NewQAccessibleInterfaceFromPointer(C.QAccessibleObject_ChildAt(ptr.Pointer(), C.int(x), C.int(y)))
	}
	return nil
}

func (ptr *QAccessibleObject) IsValid() bool {
	defer qt.Recovering("QAccessibleObject::isValid")

	if ptr.Pointer() != nil {
		return C.QAccessibleObject_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAccessibleObject) Object() *core.QObject {
	defer qt.Recovering("QAccessibleObject::object")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QAccessibleObject_Object(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleObject) Rect() *core.QRect {
	defer qt.Recovering("QAccessibleObject::rect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QAccessibleObject_Rect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleObject) ConnectSetText(f func(t QAccessible__Text, text string)) {
	defer qt.Recovering("connect QAccessibleObject::setText")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setText", f)
	}
}

func (ptr *QAccessibleObject) DisconnectSetText() {
	defer qt.Recovering("disconnect QAccessibleObject::setText")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setText")
	}
}

//export callbackQAccessibleObjectSetText
func callbackQAccessibleObjectSetText(ptr unsafe.Pointer, ptrName *C.char, t C.int, text *C.char) {
	defer qt.Recovering("callback QAccessibleObject::setText")

	if signal := qt.GetSignal(C.GoString(ptrName), "setText"); signal != nil {
		signal.(func(QAccessible__Text, string))(QAccessible__Text(t), C.GoString(text))
	} else {
		NewQAccessibleObjectFromPointer(ptr).SetTextDefault(QAccessible__Text(t), C.GoString(text))
	}
}

func (ptr *QAccessibleObject) SetText(t QAccessible__Text, text string) {
	defer qt.Recovering("QAccessibleObject::setText")

	if ptr.Pointer() != nil {
		C.QAccessibleObject_SetText(ptr.Pointer(), C.int(t), C.CString(text))
	}
}

func (ptr *QAccessibleObject) SetTextDefault(t QAccessible__Text, text string) {
	defer qt.Recovering("QAccessibleObject::setText")

	if ptr.Pointer() != nil {
		C.QAccessibleObject_SetTextDefault(ptr.Pointer(), C.int(t), C.CString(text))
	}
}

func (ptr *QAccessibleObject) DestroyQAccessibleObject() {
	defer qt.Recovering("QAccessibleObject::~QAccessibleObject")

	if ptr.Pointer() != nil {
		C.QAccessibleObject_DestroyQAccessibleObject(ptr.Pointer())
	}
}

func (ptr *QAccessibleObject) ObjectNameAbs() string {
	defer qt.Recovering("QAccessibleObject::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleObject_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAccessibleObject) SetObjectNameAbs(name string) {
	defer qt.Recovering("QAccessibleObject::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QAccessibleObject_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
