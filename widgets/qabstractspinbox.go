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

type QAbstractSpinBoxITF interface {
	QWidgetITF
	QAbstractSpinBoxPTR() *QAbstractSpinBox
}

func PointerFromQAbstractSpinBox(ptr QAbstractSpinBoxITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractSpinBoxPTR().Pointer()
	}
	return nil
}

func QAbstractSpinBoxFromPointer(ptr unsafe.Pointer) *QAbstractSpinBox {
	var n = new(QAbstractSpinBox)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAbstractSpinBox_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractSpinBox) QAbstractSpinBoxPTR() *QAbstractSpinBox {
	return ptr
}

//QAbstractSpinBox::ButtonSymbols
type QAbstractSpinBox__ButtonSymbols int

var (
	QAbstractSpinBox__UpDownArrows = QAbstractSpinBox__ButtonSymbols(0)
	QAbstractSpinBox__PlusMinus    = QAbstractSpinBox__ButtonSymbols(1)
	QAbstractSpinBox__NoButtons    = QAbstractSpinBox__ButtonSymbols(2)
)

//QAbstractSpinBox::CorrectionMode
type QAbstractSpinBox__CorrectionMode int

var (
	QAbstractSpinBox__CorrectToPreviousValue = QAbstractSpinBox__CorrectionMode(0)
	QAbstractSpinBox__CorrectToNearestValue  = QAbstractSpinBox__CorrectionMode(1)
)

//QAbstractSpinBox::StepEnabledFlag
type QAbstractSpinBox__StepEnabledFlag int

var (
	QAbstractSpinBox__StepNone        = QAbstractSpinBox__StepEnabledFlag(0x00)
	QAbstractSpinBox__StepUpEnabled   = QAbstractSpinBox__StepEnabledFlag(0x01)
	QAbstractSpinBox__StepDownEnabled = QAbstractSpinBox__StepEnabledFlag(0x02)
)

func (ptr *QAbstractSpinBox) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QAbstractSpinBox_Alignment(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractSpinBox) ButtonSymbols() QAbstractSpinBox__ButtonSymbols {
	if ptr.Pointer() != nil {
		return QAbstractSpinBox__ButtonSymbols(C.QAbstractSpinBox_ButtonSymbols(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractSpinBox) CorrectionMode() QAbstractSpinBox__CorrectionMode {
	if ptr.Pointer() != nil {
		return QAbstractSpinBox__CorrectionMode(C.QAbstractSpinBox_CorrectionMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractSpinBox) HasAcceptableInput() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_HasAcceptableInput(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) HasFrame() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_HasFrame(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) IsAccelerated() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_IsAccelerated(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) IsGroupSeparatorShown() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_IsGroupSeparatorShown(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) IsReadOnly() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_IsReadOnly(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) KeyboardTracking() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_KeyboardTracking(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) SetAccelerated(on bool) {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetAccelerated(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QAbstractSpinBox) SetAlignment(flag core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetAlignment(C.QtObjectPtr(ptr.Pointer()), C.int(flag))
	}
}

func (ptr *QAbstractSpinBox) SetButtonSymbols(bs QAbstractSpinBox__ButtonSymbols) {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetButtonSymbols(C.QtObjectPtr(ptr.Pointer()), C.int(bs))
	}
}

func (ptr *QAbstractSpinBox) SetCorrectionMode(cm QAbstractSpinBox__CorrectionMode) {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetCorrectionMode(C.QtObjectPtr(ptr.Pointer()), C.int(cm))
	}
}

func (ptr *QAbstractSpinBox) SetFrame(v bool) {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetFrame(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractSpinBox) SetGroupSeparatorShown(shown bool) {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetGroupSeparatorShown(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(shown)))
	}
}

func (ptr *QAbstractSpinBox) SetKeyboardTracking(kt bool) {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetKeyboardTracking(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(kt)))
	}
}

func (ptr *QAbstractSpinBox) SetReadOnly(r bool) {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetReadOnly(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(r)))
	}
}

func (ptr *QAbstractSpinBox) SetSpecialValueText(txt string) {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetSpecialValueText(C.QtObjectPtr(ptr.Pointer()), C.CString(txt))
	}
}

func (ptr *QAbstractSpinBox) SetWrapping(w bool) {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetWrapping(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(w)))
	}
}

func (ptr *QAbstractSpinBox) SpecialValueText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractSpinBox_SpecialValueText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QAbstractSpinBox) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractSpinBox_Text(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QAbstractSpinBox) Wrapping() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_Wrapping(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func NewQAbstractSpinBox(parent QWidgetITF) *QAbstractSpinBox {
	return QAbstractSpinBoxFromPointer(unsafe.Pointer(C.QAbstractSpinBox_NewQAbstractSpinBox(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QAbstractSpinBox) Clear() {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractSpinBox) ConnectEditingFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_ConnectEditingFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "editingFinished", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectEditingFinished() {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_DisconnectEditingFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "editingFinished")
	}
}

//export callbackQAbstractSpinBoxEditingFinished
func callbackQAbstractSpinBoxEditingFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "editingFinished").(func())()
}

func (ptr *QAbstractSpinBox) Event(event core.QEventITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_Event(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQEvent(event))) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) InputMethodQuery(query core.Qt__InputMethodQuery) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractSpinBox_InputMethodQuery(C.QtObjectPtr(ptr.Pointer()), C.int(query)))
	}
	return ""
}

func (ptr *QAbstractSpinBox) InterpretText() {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_InterpretText(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractSpinBox) SelectAll() {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SelectAll(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractSpinBox) StepBy(steps int) {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_StepBy(C.QtObjectPtr(ptr.Pointer()), C.int(steps))
	}
}

func (ptr *QAbstractSpinBox) StepDown() {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_StepDown(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractSpinBox) StepUp() {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_StepUp(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractSpinBox) DestroyQAbstractSpinBox() {
	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_DestroyQAbstractSpinBox(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
