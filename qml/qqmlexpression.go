package qml

//#include "qml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QQmlExpression struct {
	core.QObject
}

type QQmlExpression_ITF interface {
	core.QObject_ITF
	QQmlExpression_PTR() *QQmlExpression
}

func PointerFromQQmlExpression(ptr QQmlExpression_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlExpression_PTR().Pointer()
	}
	return nil
}

func NewQQmlExpressionFromPointer(ptr unsafe.Pointer) *QQmlExpression {
	var n = new(QQmlExpression)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QQmlExpression_") {
		n.SetObjectName("QQmlExpression_" + qt.Identifier())
	}
	return n
}

func (ptr *QQmlExpression) QQmlExpression_PTR() *QQmlExpression {
	return ptr
}

func NewQQmlExpression() *QQmlExpression {
	defer qt.Recovering("QQmlExpression::QQmlExpression")

	return NewQQmlExpressionFromPointer(C.QQmlExpression_NewQQmlExpression())
}

func NewQQmlExpression2(ctxt QQmlContext_ITF, scope core.QObject_ITF, expression string, parent core.QObject_ITF) *QQmlExpression {
	defer qt.Recovering("QQmlExpression::QQmlExpression")

	return NewQQmlExpressionFromPointer(C.QQmlExpression_NewQQmlExpression2(PointerFromQQmlContext(ctxt), core.PointerFromQObject(scope), C.CString(expression), core.PointerFromQObject(parent)))
}

func NewQQmlExpression3(script QQmlScriptString_ITF, ctxt QQmlContext_ITF, scope core.QObject_ITF, parent core.QObject_ITF) *QQmlExpression {
	defer qt.Recovering("QQmlExpression::QQmlExpression")

	return NewQQmlExpressionFromPointer(C.QQmlExpression_NewQQmlExpression3(PointerFromQQmlScriptString(script), PointerFromQQmlContext(ctxt), core.PointerFromQObject(scope), core.PointerFromQObject(parent)))
}

func (ptr *QQmlExpression) ClearError() {
	defer qt.Recovering("QQmlExpression::clearError")

	if ptr.Pointer() != nil {
		C.QQmlExpression_ClearError(ptr.Pointer())
	}
}

func (ptr *QQmlExpression) ColumnNumber() int {
	defer qt.Recovering("QQmlExpression::columnNumber")

	if ptr.Pointer() != nil {
		return int(C.QQmlExpression_ColumnNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlExpression) Context() *QQmlContext {
	defer qt.Recovering("QQmlExpression::context")

	if ptr.Pointer() != nil {
		return NewQQmlContextFromPointer(C.QQmlExpression_Context(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlExpression) Engine() *QQmlEngine {
	defer qt.Recovering("QQmlExpression::engine")

	if ptr.Pointer() != nil {
		return NewQQmlEngineFromPointer(C.QQmlExpression_Engine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlExpression) Evaluate(valueIsUndefined bool) *core.QVariant {
	defer qt.Recovering("QQmlExpression::evaluate")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QQmlExpression_Evaluate(ptr.Pointer(), C.int(qt.GoBoolToInt(valueIsUndefined))))
	}
	return nil
}

func (ptr *QQmlExpression) Expression() string {
	defer qt.Recovering("QQmlExpression::expression")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlExpression_Expression(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlExpression) HasError() bool {
	defer qt.Recovering("QQmlExpression::hasError")

	if ptr.Pointer() != nil {
		return C.QQmlExpression_HasError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlExpression) LineNumber() int {
	defer qt.Recovering("QQmlExpression::lineNumber")

	if ptr.Pointer() != nil {
		return int(C.QQmlExpression_LineNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlExpression) NotifyOnValueChanged() bool {
	defer qt.Recovering("QQmlExpression::notifyOnValueChanged")

	if ptr.Pointer() != nil {
		return C.QQmlExpression_NotifyOnValueChanged(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlExpression) ScopeObject() *core.QObject {
	defer qt.Recovering("QQmlExpression::scopeObject")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QQmlExpression_ScopeObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlExpression) SetExpression(expression string) {
	defer qt.Recovering("QQmlExpression::setExpression")

	if ptr.Pointer() != nil {
		C.QQmlExpression_SetExpression(ptr.Pointer(), C.CString(expression))
	}
}

func (ptr *QQmlExpression) SetNotifyOnValueChanged(notifyOnChange bool) {
	defer qt.Recovering("QQmlExpression::setNotifyOnValueChanged")

	if ptr.Pointer() != nil {
		C.QQmlExpression_SetNotifyOnValueChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(notifyOnChange)))
	}
}

func (ptr *QQmlExpression) SetSourceLocation(url string, line int, column int) {
	defer qt.Recovering("QQmlExpression::setSourceLocation")

	if ptr.Pointer() != nil {
		C.QQmlExpression_SetSourceLocation(ptr.Pointer(), C.CString(url), C.int(line), C.int(column))
	}
}

func (ptr *QQmlExpression) SourceFile() string {
	defer qt.Recovering("QQmlExpression::sourceFile")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlExpression_SourceFile(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlExpression) ConnectValueChanged(f func()) {
	defer qt.Recovering("connect QQmlExpression::valueChanged")

	if ptr.Pointer() != nil {
		C.QQmlExpression_ConnectValueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "valueChanged", f)
	}
}

func (ptr *QQmlExpression) DisconnectValueChanged() {
	defer qt.Recovering("disconnect QQmlExpression::valueChanged")

	if ptr.Pointer() != nil {
		C.QQmlExpression_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "valueChanged")
	}
}

//export callbackQQmlExpressionValueChanged
func callbackQQmlExpressionValueChanged(ptrName *C.char) {
	defer qt.Recovering("callback QQmlExpression::valueChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "valueChanged")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QQmlExpression) DestroyQQmlExpression() {
	defer qt.Recovering("QQmlExpression::~QQmlExpression")

	if ptr.Pointer() != nil {
		C.QQmlExpression_DestroyQQmlExpression(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
