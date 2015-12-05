package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
	for len(n.ObjectName()) < len("QAbstractSpinBox_") {
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::alignment")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QAbstractSpinBox_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSpinBox) ButtonSymbols() QAbstractSpinBox__ButtonSymbols {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::buttonSymbols")
		}
	}()

	if ptr.Pointer() != nil {
		return QAbstractSpinBox__ButtonSymbols(C.QAbstractSpinBox_ButtonSymbols(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSpinBox) CorrectionMode() QAbstractSpinBox__CorrectionMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::correctionMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QAbstractSpinBox__CorrectionMode(C.QAbstractSpinBox_CorrectionMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSpinBox) HasAcceptableInput() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::hasAcceptableInput")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_HasAcceptableInput(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) HasFrame() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::hasFrame")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_HasFrame(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) IsAccelerated() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::isAccelerated")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_IsAccelerated(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) IsGroupSeparatorShown() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::isGroupSeparatorShown")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_IsGroupSeparatorShown(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) IsReadOnly() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::isReadOnly")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_IsReadOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) KeyboardTracking() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::keyboardTracking")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_KeyboardTracking(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) SetAccelerated(on bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::setAccelerated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetAccelerated(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QAbstractSpinBox) SetAlignment(flag core.Qt__AlignmentFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::setAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetAlignment(ptr.Pointer(), C.int(flag))
	}
}

func (ptr *QAbstractSpinBox) SetButtonSymbols(bs QAbstractSpinBox__ButtonSymbols) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::setButtonSymbols")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetButtonSymbols(ptr.Pointer(), C.int(bs))
	}
}

func (ptr *QAbstractSpinBox) SetCorrectionMode(cm QAbstractSpinBox__CorrectionMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::setCorrectionMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetCorrectionMode(ptr.Pointer(), C.int(cm))
	}
}

func (ptr *QAbstractSpinBox) SetFrame(v bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::setFrame")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetFrame(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractSpinBox) SetGroupSeparatorShown(shown bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::setGroupSeparatorShown")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetGroupSeparatorShown(ptr.Pointer(), C.int(qt.GoBoolToInt(shown)))
	}
}

func (ptr *QAbstractSpinBox) SetKeyboardTracking(kt bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::setKeyboardTracking")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetKeyboardTracking(ptr.Pointer(), C.int(qt.GoBoolToInt(kt)))
	}
}

func (ptr *QAbstractSpinBox) SetReadOnly(r bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::setReadOnly")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetReadOnly(ptr.Pointer(), C.int(qt.GoBoolToInt(r)))
	}
}

func (ptr *QAbstractSpinBox) SetSpecialValueText(txt string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::setSpecialValueText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetSpecialValueText(ptr.Pointer(), C.CString(txt))
	}
}

func (ptr *QAbstractSpinBox) SetWrapping(w bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::setWrapping")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SetWrapping(ptr.Pointer(), C.int(qt.GoBoolToInt(w)))
	}
}

func (ptr *QAbstractSpinBox) SpecialValueText() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::specialValueText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractSpinBox_SpecialValueText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractSpinBox) Text() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::text")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractSpinBox_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractSpinBox) Wrapping() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::wrapping")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_Wrapping(ptr.Pointer()) != 0
	}
	return false
}

func NewQAbstractSpinBox(parent QWidget_ITF) *QAbstractSpinBox {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::QAbstractSpinBox")
		}
	}()

	return NewQAbstractSpinBoxFromPointer(C.QAbstractSpinBox_NewQAbstractSpinBox(PointerFromQWidget(parent)))
}

func (ptr *QAbstractSpinBox) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_Clear(ptr.Pointer())
	}
}

func (ptr *QAbstractSpinBox) ConnectEditingFinished(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::editingFinished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_ConnectEditingFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "editingFinished", f)
	}
}

func (ptr *QAbstractSpinBox) DisconnectEditingFinished() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::editingFinished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_DisconnectEditingFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "editingFinished")
	}
}

//export callbackQAbstractSpinBoxEditingFinished
func callbackQAbstractSpinBoxEditingFinished(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::editingFinished")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "editingFinished").(func())()
}

func (ptr *QAbstractSpinBox) Event(event core.QEvent_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::event")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractSpinBox_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QAbstractSpinBox) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::inputMethodQuery")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QAbstractSpinBox_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QAbstractSpinBox) InterpretText() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::interpretText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_InterpretText(ptr.Pointer())
	}
}

func (ptr *QAbstractSpinBox) SelectAll() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::selectAll")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_SelectAll(ptr.Pointer())
	}
}

func (ptr *QAbstractSpinBox) StepBy(steps int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::stepBy")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_StepBy(ptr.Pointer(), C.int(steps))
	}
}

func (ptr *QAbstractSpinBox) StepDown() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::stepDown")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_StepDown(ptr.Pointer())
	}
}

func (ptr *QAbstractSpinBox) StepUp() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::stepUp")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_StepUp(ptr.Pointer())
	}
}

func (ptr *QAbstractSpinBox) DestroyQAbstractSpinBox() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractSpinBox::~QAbstractSpinBox")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractSpinBox_DestroyQAbstractSpinBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
