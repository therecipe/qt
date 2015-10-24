package qml

//#include "qqmlexpression.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QQmlExpression struct {
	core.QObject
}

type QQmlExpressionITF interface {
	core.QObjectITF
	QQmlExpressionPTR() *QQmlExpression
}

func PointerFromQQmlExpression(ptr QQmlExpressionITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlExpressionPTR().Pointer()
	}
	return nil
}

func QQmlExpressionFromPointer(ptr unsafe.Pointer) *QQmlExpression {
	var n = new(QQmlExpression)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QQmlExpression_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQmlExpression) QQmlExpressionPTR() *QQmlExpression {
	return ptr
}

func NewQQmlExpression() *QQmlExpression {
	return QQmlExpressionFromPointer(unsafe.Pointer(C.QQmlExpression_NewQQmlExpression()))
}

func NewQQmlExpression2(ctxt QQmlContextITF, scope core.QObjectITF, expression string, parent core.QObjectITF) *QQmlExpression {
	return QQmlExpressionFromPointer(unsafe.Pointer(C.QQmlExpression_NewQQmlExpression2(C.QtObjectPtr(PointerFromQQmlContext(ctxt)), C.QtObjectPtr(core.PointerFromQObject(scope)), C.CString(expression), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQQmlExpression3(script QQmlScriptStringITF, ctxt QQmlContextITF, scope core.QObjectITF, parent core.QObjectITF) *QQmlExpression {
	return QQmlExpressionFromPointer(unsafe.Pointer(C.QQmlExpression_NewQQmlExpression3(C.QtObjectPtr(PointerFromQQmlScriptString(script)), C.QtObjectPtr(PointerFromQQmlContext(ctxt)), C.QtObjectPtr(core.PointerFromQObject(scope)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QQmlExpression) ClearError() {
	if ptr.Pointer() != nil {
		C.QQmlExpression_ClearError(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QQmlExpression) ColumnNumber() int {
	if ptr.Pointer() != nil {
		return int(C.QQmlExpression_ColumnNumber(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlExpression) Context() *QQmlContext {
	if ptr.Pointer() != nil {
		return QQmlContextFromPointer(unsafe.Pointer(C.QQmlExpression_Context(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQmlExpression) Engine() *QQmlEngine {
	if ptr.Pointer() != nil {
		return QQmlEngineFromPointer(unsafe.Pointer(C.QQmlExpression_Engine(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQmlExpression) Evaluate(valueIsUndefined bool) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlExpression_Evaluate(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(valueIsUndefined))))
	}
	return ""
}

func (ptr *QQmlExpression) Expression() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlExpression_Expression(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QQmlExpression) HasError() bool {
	if ptr.Pointer() != nil {
		return C.QQmlExpression_HasError(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlExpression) LineNumber() int {
	if ptr.Pointer() != nil {
		return int(C.QQmlExpression_LineNumber(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlExpression) NotifyOnValueChanged() bool {
	if ptr.Pointer() != nil {
		return C.QQmlExpression_NotifyOnValueChanged(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlExpression) ScopeObject() *core.QObject {
	if ptr.Pointer() != nil {
		return core.QObjectFromPointer(unsafe.Pointer(C.QQmlExpression_ScopeObject(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQmlExpression) SetExpression(expression string) {
	if ptr.Pointer() != nil {
		C.QQmlExpression_SetExpression(C.QtObjectPtr(ptr.Pointer()), C.CString(expression))
	}
}

func (ptr *QQmlExpression) SetNotifyOnValueChanged(notifyOnChange bool) {
	if ptr.Pointer() != nil {
		C.QQmlExpression_SetNotifyOnValueChanged(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(notifyOnChange)))
	}
}

func (ptr *QQmlExpression) SetSourceLocation(url string, line int, column int) {
	if ptr.Pointer() != nil {
		C.QQmlExpression_SetSourceLocation(C.QtObjectPtr(ptr.Pointer()), C.CString(url), C.int(line), C.int(column))
	}
}

func (ptr *QQmlExpression) SourceFile() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlExpression_SourceFile(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QQmlExpression) ConnectValueChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQmlExpression_ConnectValueChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "valueChanged", f)
	}
}

func (ptr *QQmlExpression) DisconnectValueChanged() {
	if ptr.Pointer() != nil {
		C.QQmlExpression_DisconnectValueChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "valueChanged")
	}
}

//export callbackQQmlExpressionValueChanged
func callbackQQmlExpressionValueChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "valueChanged").(func())()
}

func (ptr *QQmlExpression) DestroyQQmlExpression() {
	if ptr.Pointer() != nil {
		C.QQmlExpression_DestroyQQmlExpression(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
