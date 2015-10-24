package gui

//#include "qinputmethod.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QInputMethod struct {
	core.QObject
}

type QInputMethodITF interface {
	core.QObjectITF
	QInputMethodPTR() *QInputMethod
}

func PointerFromQInputMethod(ptr QInputMethodITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QInputMethodPTR().Pointer()
	}
	return nil
}

func QInputMethodFromPointer(ptr unsafe.Pointer) *QInputMethod {
	var n = new(QInputMethod)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QInputMethod_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QInputMethod) QInputMethodPTR() *QInputMethod {
	return ptr
}

//QInputMethod::Action
type QInputMethod__Action int

var (
	QInputMethod__Click       = QInputMethod__Action(0)
	QInputMethod__ContextMenu = QInputMethod__Action(1)
)

func (ptr *QInputMethod) InputDirection() core.Qt__LayoutDirection {
	if ptr.Pointer() != nil {
		return core.Qt__LayoutDirection(C.QInputMethod_InputDirection(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QInputMethod) IsAnimating() bool {
	if ptr.Pointer() != nil {
		return C.QInputMethod_IsAnimating(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QInputMethod) IsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QInputMethod_IsVisible(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QInputMethod) ConnectAnimatingChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QInputMethod_ConnectAnimatingChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "animatingChanged", f)
	}
}

func (ptr *QInputMethod) DisconnectAnimatingChanged() {
	if ptr.Pointer() != nil {
		C.QInputMethod_DisconnectAnimatingChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "animatingChanged")
	}
}

//export callbackQInputMethodAnimatingChanged
func callbackQInputMethodAnimatingChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "animatingChanged").(func())()
}

func (ptr *QInputMethod) Commit() {
	if ptr.Pointer() != nil {
		C.QInputMethod_Commit(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QInputMethod) ConnectCursorRectangleChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QInputMethod_ConnectCursorRectangleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "cursorRectangleChanged", f)
	}
}

func (ptr *QInputMethod) DisconnectCursorRectangleChanged() {
	if ptr.Pointer() != nil {
		C.QInputMethod_DisconnectCursorRectangleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "cursorRectangleChanged")
	}
}

//export callbackQInputMethodCursorRectangleChanged
func callbackQInputMethodCursorRectangleChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "cursorRectangleChanged").(func())()
}

func (ptr *QInputMethod) Hide() {
	if ptr.Pointer() != nil {
		C.QInputMethod_Hide(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QInputMethod) ConnectInputDirectionChanged(f func(newDirection core.Qt__LayoutDirection)) {
	if ptr.Pointer() != nil {
		C.QInputMethod_ConnectInputDirectionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "inputDirectionChanged", f)
	}
}

func (ptr *QInputMethod) DisconnectInputDirectionChanged() {
	if ptr.Pointer() != nil {
		C.QInputMethod_DisconnectInputDirectionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "inputDirectionChanged")
	}
}

//export callbackQInputMethodInputDirectionChanged
func callbackQInputMethodInputDirectionChanged(ptrName *C.char, newDirection C.int) {
	qt.GetSignal(C.GoString(ptrName), "inputDirectionChanged").(func(core.Qt__LayoutDirection))(core.Qt__LayoutDirection(newDirection))
}

func (ptr *QInputMethod) InvokeAction(a QInputMethod__Action, cursorPosition int) {
	if ptr.Pointer() != nil {
		C.QInputMethod_InvokeAction(C.QtObjectPtr(ptr.Pointer()), C.int(a), C.int(cursorPosition))
	}
}

func (ptr *QInputMethod) ConnectKeyboardRectangleChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QInputMethod_ConnectKeyboardRectangleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "keyboardRectangleChanged", f)
	}
}

func (ptr *QInputMethod) DisconnectKeyboardRectangleChanged() {
	if ptr.Pointer() != nil {
		C.QInputMethod_DisconnectKeyboardRectangleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "keyboardRectangleChanged")
	}
}

//export callbackQInputMethodKeyboardRectangleChanged
func callbackQInputMethodKeyboardRectangleChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "keyboardRectangleChanged").(func())()
}

func (ptr *QInputMethod) ConnectLocaleChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QInputMethod_ConnectLocaleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "localeChanged", f)
	}
}

func (ptr *QInputMethod) DisconnectLocaleChanged() {
	if ptr.Pointer() != nil {
		C.QInputMethod_DisconnectLocaleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "localeChanged")
	}
}

//export callbackQInputMethodLocaleChanged
func callbackQInputMethodLocaleChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "localeChanged").(func())()
}

func QInputMethod_QueryFocusObject(query core.Qt__InputMethodQuery, argument string) string {
	return C.GoString(C.QInputMethod_QInputMethod_QueryFocusObject(C.int(query), C.CString(argument)))
}

func (ptr *QInputMethod) Reset() {
	if ptr.Pointer() != nil {
		C.QInputMethod_Reset(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QInputMethod) SetInputItemRectangle(rect core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QInputMethod_SetInputItemRectangle(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rect)))
	}
}

func (ptr *QInputMethod) SetInputItemTransform(transform QTransformITF) {
	if ptr.Pointer() != nil {
		C.QInputMethod_SetInputItemTransform(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTransform(transform)))
	}
}

func (ptr *QInputMethod) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QInputMethod_SetVisible(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QInputMethod) Show() {
	if ptr.Pointer() != nil {
		C.QInputMethod_Show(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QInputMethod) Update(queries core.Qt__InputMethodQuery) {
	if ptr.Pointer() != nil {
		C.QInputMethod_Update(C.QtObjectPtr(ptr.Pointer()), C.int(queries))
	}
}

func (ptr *QInputMethod) ConnectVisibleChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QInputMethod_ConnectVisibleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "visibleChanged", f)
	}
}

func (ptr *QInputMethod) DisconnectVisibleChanged() {
	if ptr.Pointer() != nil {
		C.QInputMethod_DisconnectVisibleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "visibleChanged")
	}
}

//export callbackQInputMethodVisibleChanged
func callbackQInputMethodVisibleChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "visibleChanged").(func())()
}
