package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QInputMethod struct {
	core.QObject
}

type QInputMethod_ITF interface {
	core.QObject_ITF
	QInputMethod_PTR() *QInputMethod
}

func PointerFromQInputMethod(ptr QInputMethod_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QInputMethod_PTR().Pointer()
	}
	return nil
}

func NewQInputMethodFromPointer(ptr unsafe.Pointer) *QInputMethod {
	var n = new(QInputMethod)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QInputMethod_") {
		n.SetObjectName("QInputMethod_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QInputMethod) QInputMethod_PTR() *QInputMethod {
	return ptr
}

//QInputMethod::Action
type QInputMethod__Action int64

const (
	QInputMethod__Click       = QInputMethod__Action(0)
	QInputMethod__ContextMenu = QInputMethod__Action(1)
)

func (ptr *QInputMethod) InputDirection() core.Qt__LayoutDirection {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::inputDirection")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__LayoutDirection(C.QInputMethod_InputDirection(ptr.Pointer()))
	}
	return 0
}

func (ptr *QInputMethod) IsAnimating() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::isAnimating")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QInputMethod_IsAnimating(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QInputMethod) IsVisible() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::isVisible")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QInputMethod_IsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QInputMethod) ConnectAnimatingChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::animatingChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QInputMethod_ConnectAnimatingChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "animatingChanged", f)
	}
}

func (ptr *QInputMethod) DisconnectAnimatingChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::animatingChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QInputMethod_DisconnectAnimatingChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "animatingChanged")
	}
}

//export callbackQInputMethodAnimatingChanged
func callbackQInputMethodAnimatingChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::animatingChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "animatingChanged").(func())()
}

func (ptr *QInputMethod) Commit() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::commit")
		}
	}()

	if ptr.Pointer() != nil {
		C.QInputMethod_Commit(ptr.Pointer())
	}
}

func (ptr *QInputMethod) ConnectCursorRectangleChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::cursorRectangleChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QInputMethod_ConnectCursorRectangleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cursorRectangleChanged", f)
	}
}

func (ptr *QInputMethod) DisconnectCursorRectangleChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::cursorRectangleChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QInputMethod_DisconnectCursorRectangleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cursorRectangleChanged")
	}
}

//export callbackQInputMethodCursorRectangleChanged
func callbackQInputMethodCursorRectangleChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::cursorRectangleChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "cursorRectangleChanged").(func())()
}

func (ptr *QInputMethod) Hide() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::hide")
		}
	}()

	if ptr.Pointer() != nil {
		C.QInputMethod_Hide(ptr.Pointer())
	}
}

func (ptr *QInputMethod) ConnectInputDirectionChanged(f func(newDirection core.Qt__LayoutDirection)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::inputDirectionChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QInputMethod_ConnectInputDirectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "inputDirectionChanged", f)
	}
}

func (ptr *QInputMethod) DisconnectInputDirectionChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::inputDirectionChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QInputMethod_DisconnectInputDirectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "inputDirectionChanged")
	}
}

//export callbackQInputMethodInputDirectionChanged
func callbackQInputMethodInputDirectionChanged(ptrName *C.char, newDirection C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::inputDirectionChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "inputDirectionChanged").(func(core.Qt__LayoutDirection))(core.Qt__LayoutDirection(newDirection))
}

func (ptr *QInputMethod) InvokeAction(a QInputMethod__Action, cursorPosition int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::invokeAction")
		}
	}()

	if ptr.Pointer() != nil {
		C.QInputMethod_InvokeAction(ptr.Pointer(), C.int(a), C.int(cursorPosition))
	}
}

func (ptr *QInputMethod) ConnectKeyboardRectangleChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::keyboardRectangleChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QInputMethod_ConnectKeyboardRectangleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "keyboardRectangleChanged", f)
	}
}

func (ptr *QInputMethod) DisconnectKeyboardRectangleChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::keyboardRectangleChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QInputMethod_DisconnectKeyboardRectangleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "keyboardRectangleChanged")
	}
}

//export callbackQInputMethodKeyboardRectangleChanged
func callbackQInputMethodKeyboardRectangleChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::keyboardRectangleChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "keyboardRectangleChanged").(func())()
}

func (ptr *QInputMethod) ConnectLocaleChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::localeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QInputMethod_ConnectLocaleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "localeChanged", f)
	}
}

func (ptr *QInputMethod) DisconnectLocaleChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::localeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QInputMethod_DisconnectLocaleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "localeChanged")
	}
}

//export callbackQInputMethodLocaleChanged
func callbackQInputMethodLocaleChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::localeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "localeChanged").(func())()
}

func QInputMethod_QueryFocusObject(query core.Qt__InputMethodQuery, argument core.QVariant_ITF) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::queryFocusObject")
		}
	}()

	return core.NewQVariantFromPointer(C.QInputMethod_QInputMethod_QueryFocusObject(C.int(query), core.PointerFromQVariant(argument)))
}

func (ptr *QInputMethod) Reset() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::reset")
		}
	}()

	if ptr.Pointer() != nil {
		C.QInputMethod_Reset(ptr.Pointer())
	}
}

func (ptr *QInputMethod) SetInputItemRectangle(rect core.QRectF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::setInputItemRectangle")
		}
	}()

	if ptr.Pointer() != nil {
		C.QInputMethod_SetInputItemRectangle(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QInputMethod) SetInputItemTransform(transform QTransform_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::setInputItemTransform")
		}
	}()

	if ptr.Pointer() != nil {
		C.QInputMethod_SetInputItemTransform(ptr.Pointer(), PointerFromQTransform(transform))
	}
}

func (ptr *QInputMethod) SetVisible(visible bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::setVisible")
		}
	}()

	if ptr.Pointer() != nil {
		C.QInputMethod_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QInputMethod) Show() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::show")
		}
	}()

	if ptr.Pointer() != nil {
		C.QInputMethod_Show(ptr.Pointer())
	}
}

func (ptr *QInputMethod) Update(queries core.Qt__InputMethodQuery) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::update")
		}
	}()

	if ptr.Pointer() != nil {
		C.QInputMethod_Update(ptr.Pointer(), C.int(queries))
	}
}

func (ptr *QInputMethod) ConnectVisibleChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::visibleChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QInputMethod_ConnectVisibleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "visibleChanged", f)
	}
}

func (ptr *QInputMethod) DisconnectVisibleChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::visibleChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QInputMethod_DisconnectVisibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "visibleChanged")
	}
}

//export callbackQInputMethodVisibleChanged
func callbackQInputMethodVisibleChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QInputMethod::visibleChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "visibleChanged").(func())()
}
