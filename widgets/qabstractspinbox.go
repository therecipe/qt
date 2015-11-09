package widgets

//#include "qabstractspinbox.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAbstractSpinBox struct {
	QWidget
}

type QAbstractSpinBox_ITF interface {
	QWidget_ITF
	QAbstractSpinBox_PTR() *QAbstractSpinBox
}

func PointerFromQAbstractSpinBox(ptr QAbstractSpinBox_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractSpinBox_PTR().Pointer()
	}
	return nil
}

func NewQAbstractSpinBoxFromPointer(ptr unsafe.Pointer) *QAbstractSpinBox {
	var n = new(QAbstractSpinBox)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QAbstractSpinBox_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractSpinBox) QAbstractSpinBox_PTR() *QAbstractSpinBox {
	return ptr
}

//QAbstractSpinBox::ButtonSymbols
type QAbstractSpinBox__ButtonSymbols int64

const (
	QAbstractSpinBox__UpDownArrows = QAbstractSpinBox__ButtonSymbols(0)
	QAbstractSpinBox__PlusMinus    = QAbstractSpinBox__ButtonSymbols(1)
	QAbstractSpinBox__NoButtons    = QAbstractSpinBox__ButtonSymbols(2)
)

//QAbstractSpinBox::CorrectionMode
type QAbstractSpinBox__CorrectionMode int64

const (
	QAbstractSpinBox__CorrectToPreviousValue = QAbstractSpinBox__CorrectionMode(0)
	QAbstractSpinBox__CorrectToNearestValue  = QAbstractSpinBox__CorrectionMode(1)
)

//QAbstractSpinBox::StepEnabledFlag
type QAbstractSpinBox__StepEnabledFlag int64

const (
	QAbstractSpinBox__StepNone        = QAbstractSpinBox__StepEnabledFlag(0x00)
	QAbstractSpinBox__StepUpEnabled   = QAbstractSpinBox__StepEnabledFlag(0x01)
	QAbstractSpinBox__StepDownEnabled = QAbstractSpinBox__StepEnabledFlag(0x02)
)

func (ptr *QAbstractSpinBox) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QAbstractSpinBox_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSpinBox) ButtonSymbols() QAbstractSpinBox__ButtonSymbols {
	if ptr.Pointer() != nil {
		return QAbstractSpinBox__ButtonSymbols(C.QAbstractSpinBox_ButtonSymbols(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSpinBox) CorrectionMode() QAbstractSpinBox__CorrectionMode {
	if ptr.Pointer() != nil {
		return QAbstractSpinBox__CorrectionMode(C.QAbstractSpinBox_CorrectionMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSpinBox) HasAcceptableInput() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_HasAcceptableInput(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) HasFrame() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_HasFrame(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) IsAccelerated() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_IsAccelerated(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) IsGroupSeparatorShown() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_IsGroupSeparatorShown(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) IsReadOnly() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_IsReadOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) KeyboardTracking() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_KeyboardTracking(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) SetAccelerated(on bool) {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetAccelerated(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QAbstractSpinBox) SetAlignment(flag core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetAlignment(ptr.Pointer(), C.int(flag))
	}
}

func (ptr *QAbstractSpinBox) SetButtonSymbols(bs QAbstractSpinBox__ButtonSymbols) {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetButtonSymbols(ptr.Pointer(), C.int(bs))
	}
}

func (ptr *QAbstractSpinBox) SetCorrectionMode(cm QAbstractSpinBox__CorrectionMode) {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetCorrectionMode(ptr.Pointer(), C.int(cm))
	}
}

func (ptr *QAbstractSpinBox) SetFrame(v bool) {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetFrame(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractSpinBox) SetGroupSeparatorShown(shown bool) {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetGroupSeparatorShown(ptr.Pointer(), C.int(qt.GoBoolToInt(shown)))
	}
}

func (ptr *QAbstractSpinBox) SetKeyboardTracking(kt bool) {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetKeyboardTracking(ptr.Pointer(), C.int(qt.GoBoolToInt(kt)))
	}
}

func (ptr *QAbstractSpinBox) SetReadOnly(r bool) {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetReadOnly(ptr.Pointer(), C.int(qt.GoBoolToInt(r)))
	}
}

func (ptr *QAbstractSpinBox) SetSpecialValueText(txt string) {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetSpecialValueText(ptr.Pointer(), C.CString(txt))
	}
}

func (ptr *QAbstractSpinBox) SetWrapping(w bool) {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetWrapping(ptr.Pointer(), C.int(qt.GoBoolToInt(w)))
	}
}

func (ptr *QAbstractSpinBox) SpecialValueText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractSpinBox_SpecialValueText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractSpinBox) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractSpinBox_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractSpinBox) Wrapping() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_Wrapping(ptr.Pointer()) != 0
	}
	return false
}

func NewQAbstractSpinBox(parent QWidget_ITF) *QAbstractSpinBox {
	return NewQAbstractSpinBoxFromPointer(C.QAbstractSpinBox_NewQAbstractSpinBox(PointerFromQWidget(parent)))
}

func (ptr *QAbstractSpinBox) Clear() {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_Clear(ptr.Pointer())
	}
}

func (ptr *QAbstractSpinBox) ConnectEditingFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_ConnectEditingFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "editingFinished", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectEditingFinished() {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_DisconnectEditingFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "editingFinished")
	}
}

//export callbackQAbstractSpinBoxEditingFinished
func callbackQAbstractSpinBoxEditingFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "editingFinished").(func())()
}

func (ptr *QAbstractSpinBox) Event(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QAbstractSpinBox_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QAbstractSpinBox) InterpretText() {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_InterpretText(ptr.Pointer())
	}
}

func (ptr *QAbstractSpinBox) SelectAll() {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SelectAll(ptr.Pointer())
	}
}

func (ptr *QAbstractSpinBox) StepBy(steps int) {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_StepBy(ptr.Pointer(), C.int(steps))
	}
}

func (ptr *QAbstractSpinBox) StepDown() {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_StepDown(ptr.Pointer())
	}
}

func (ptr *QAbstractSpinBox) StepUp() {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_StepUp(ptr.Pointer())
	}
}

func (ptr *QAbstractSpinBox) DestroyQAbstractSpinBox() {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_DestroyQAbstractSpinBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
