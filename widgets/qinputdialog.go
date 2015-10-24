package widgets

//#include "qinputdialog.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QInputDialog struct {
	QDialog
}

type QInputDialogITF interface {
	QDialogITF
	QInputDialogPTR() *QInputDialog
}

func PointerFromQInputDialog(ptr QInputDialogITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QInputDialogPTR().Pointer()
	}
	return nil
}

func QInputDialogFromPointer(ptr unsafe.Pointer) *QInputDialog {
	var n = new(QInputDialog)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QInputDialog_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QInputDialog) QInputDialogPTR() *QInputDialog {
	return ptr
}

//QInputDialog::InputDialogOption
type QInputDialog__InputDialogOption int

var (
	QInputDialog__NoButtons                    = QInputDialog__InputDialogOption(0x00000001)
	QInputDialog__UseListViewForComboBoxItems  = QInputDialog__InputDialogOption(0x00000002)
	QInputDialog__UsePlainTextEditForTextInput = QInputDialog__InputDialogOption(0x00000004)
)

//QInputDialog::InputMode
type QInputDialog__InputMode int

var (
	QInputDialog__TextInput   = QInputDialog__InputMode(0)
	QInputDialog__IntInput    = QInputDialog__InputMode(1)
	QInputDialog__DoubleInput = QInputDialog__InputMode(2)
)

func (ptr *QInputDialog) CancelButtonText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QInputDialog_CancelButtonText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QInputDialog) ComboBoxItems() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QInputDialog_ComboBoxItems(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QInputDialog) DoubleDecimals() int {
	if ptr.Pointer() != nil {
		return int(C.QInputDialog_DoubleDecimals(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QInputDialog) InputMode() QInputDialog__InputMode {
	if ptr.Pointer() != nil {
		return QInputDialog__InputMode(C.QInputDialog_InputMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QInputDialog) IntMaximum() int {
	if ptr.Pointer() != nil {
		return int(C.QInputDialog_IntMaximum(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QInputDialog) IntMinimum() int {
	if ptr.Pointer() != nil {
		return int(C.QInputDialog_IntMinimum(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QInputDialog) IntStep() int {
	if ptr.Pointer() != nil {
		return int(C.QInputDialog_IntStep(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QInputDialog) IntValue() int {
	if ptr.Pointer() != nil {
		return int(C.QInputDialog_IntValue(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QInputDialog) IsComboBoxEditable() bool {
	if ptr.Pointer() != nil {
		return C.QInputDialog_IsComboBoxEditable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QInputDialog) LabelText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QInputDialog_LabelText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QInputDialog) OkButtonText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QInputDialog_OkButtonText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QInputDialog) Options() QInputDialog__InputDialogOption {
	if ptr.Pointer() != nil {
		return QInputDialog__InputDialogOption(C.QInputDialog_Options(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QInputDialog) SetCancelButtonText(text string) {
	if ptr.Pointer() != nil {
		C.QInputDialog_SetCancelButtonText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QInputDialog) SetComboBoxEditable(editable bool) {
	if ptr.Pointer() != nil {
		C.QInputDialog_SetComboBoxEditable(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(editable)))
	}
}

func (ptr *QInputDialog) SetComboBoxItems(items []string) {
	if ptr.Pointer() != nil {
		C.QInputDialog_SetComboBoxItems(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(items, "|")))
	}
}

func (ptr *QInputDialog) SetDoubleDecimals(decimals int) {
	if ptr.Pointer() != nil {
		C.QInputDialog_SetDoubleDecimals(C.QtObjectPtr(ptr.Pointer()), C.int(decimals))
	}
}

func (ptr *QInputDialog) SetInputMode(mode QInputDialog__InputMode) {
	if ptr.Pointer() != nil {
		C.QInputDialog_SetInputMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QInputDialog) SetIntMaximum(max int) {
	if ptr.Pointer() != nil {
		C.QInputDialog_SetIntMaximum(C.QtObjectPtr(ptr.Pointer()), C.int(max))
	}
}

func (ptr *QInputDialog) SetIntMinimum(min int) {
	if ptr.Pointer() != nil {
		C.QInputDialog_SetIntMinimum(C.QtObjectPtr(ptr.Pointer()), C.int(min))
	}
}

func (ptr *QInputDialog) SetIntStep(step int) {
	if ptr.Pointer() != nil {
		C.QInputDialog_SetIntStep(C.QtObjectPtr(ptr.Pointer()), C.int(step))
	}
}

func (ptr *QInputDialog) SetIntValue(value int) {
	if ptr.Pointer() != nil {
		C.QInputDialog_SetIntValue(C.QtObjectPtr(ptr.Pointer()), C.int(value))
	}
}

func (ptr *QInputDialog) SetLabelText(text string) {
	if ptr.Pointer() != nil {
		C.QInputDialog_SetLabelText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QInputDialog) SetOkButtonText(text string) {
	if ptr.Pointer() != nil {
		C.QInputDialog_SetOkButtonText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QInputDialog) SetOptions(options QInputDialog__InputDialogOption) {
	if ptr.Pointer() != nil {
		C.QInputDialog_SetOptions(C.QtObjectPtr(ptr.Pointer()), C.int(options))
	}
}

func (ptr *QInputDialog) SetTextEchoMode(mode QLineEdit__EchoMode) {
	if ptr.Pointer() != nil {
		C.QInputDialog_SetTextEchoMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QInputDialog) SetTextValue(text string) {
	if ptr.Pointer() != nil {
		C.QInputDialog_SetTextValue(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QInputDialog) TextEchoMode() QLineEdit__EchoMode {
	if ptr.Pointer() != nil {
		return QLineEdit__EchoMode(C.QInputDialog_TextEchoMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QInputDialog) TextValue() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QInputDialog_TextValue(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func NewQInputDialog(parent QWidgetITF, flags core.Qt__WindowType) *QInputDialog {
	return QInputDialogFromPointer(unsafe.Pointer(C.QInputDialog_NewQInputDialog(C.QtObjectPtr(PointerFromQWidget(parent)), C.int(flags))))
}

func (ptr *QInputDialog) Done(result int) {
	if ptr.Pointer() != nil {
		C.QInputDialog_Done(C.QtObjectPtr(ptr.Pointer()), C.int(result))
	}
}

func QInputDialog_GetInt(parent QWidgetITF, title string, label string, value int, min int, max int, step int, ok bool, flags core.Qt__WindowType) int {
	return int(C.QInputDialog_QInputDialog_GetInt(C.QtObjectPtr(PointerFromQWidget(parent)), C.CString(title), C.CString(label), C.int(value), C.int(min), C.int(max), C.int(step), C.int(qt.GoBoolToInt(ok)), C.int(flags)))
}

func QInputDialog_GetItem(parent QWidgetITF, title string, label string, items []string, current int, editable bool, ok bool, flags core.Qt__WindowType, inputMethodHints core.Qt__InputMethodHint) string {
	return C.GoString(C.QInputDialog_QInputDialog_GetItem(C.QtObjectPtr(PointerFromQWidget(parent)), C.CString(title), C.CString(label), C.CString(strings.Join(items, "|")), C.int(current), C.int(qt.GoBoolToInt(editable)), C.int(qt.GoBoolToInt(ok)), C.int(flags), C.int(inputMethodHints)))
}

func QInputDialog_GetMultiLineText(parent QWidgetITF, title string, label string, text string, ok bool, flags core.Qt__WindowType, inputMethodHints core.Qt__InputMethodHint) string {
	return C.GoString(C.QInputDialog_QInputDialog_GetMultiLineText(C.QtObjectPtr(PointerFromQWidget(parent)), C.CString(title), C.CString(label), C.CString(text), C.int(qt.GoBoolToInt(ok)), C.int(flags), C.int(inputMethodHints)))
}

func QInputDialog_GetText(parent QWidgetITF, title string, label string, mode QLineEdit__EchoMode, text string, ok bool, flags core.Qt__WindowType, inputMethodHints core.Qt__InputMethodHint) string {
	return C.GoString(C.QInputDialog_QInputDialog_GetText(C.QtObjectPtr(PointerFromQWidget(parent)), C.CString(title), C.CString(label), C.int(mode), C.CString(text), C.int(qt.GoBoolToInt(ok)), C.int(flags), C.int(inputMethodHints)))
}

func (ptr *QInputDialog) ConnectIntValueChanged(f func(value int)) {
	if ptr.Pointer() != nil {
		C.QInputDialog_ConnectIntValueChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "intValueChanged", f)
	}
}

func (ptr *QInputDialog) DisconnectIntValueChanged() {
	if ptr.Pointer() != nil {
		C.QInputDialog_DisconnectIntValueChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "intValueChanged")
	}
}

//export callbackQInputDialogIntValueChanged
func callbackQInputDialogIntValueChanged(ptrName *C.char, value C.int) {
	qt.GetSignal(C.GoString(ptrName), "intValueChanged").(func(int))(int(value))
}

func (ptr *QInputDialog) ConnectIntValueSelected(f func(value int)) {
	if ptr.Pointer() != nil {
		C.QInputDialog_ConnectIntValueSelected(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "intValueSelected", f)
	}
}

func (ptr *QInputDialog) DisconnectIntValueSelected() {
	if ptr.Pointer() != nil {
		C.QInputDialog_DisconnectIntValueSelected(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "intValueSelected")
	}
}

//export callbackQInputDialogIntValueSelected
func callbackQInputDialogIntValueSelected(ptrName *C.char, value C.int) {
	qt.GetSignal(C.GoString(ptrName), "intValueSelected").(func(int))(int(value))
}

func (ptr *QInputDialog) Open(receiver core.QObjectITF, member string) {
	if ptr.Pointer() != nil {
		C.QInputDialog_Open(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQObject(receiver)), C.CString(member))
	}
}

func (ptr *QInputDialog) SetIntRange(min int, max int) {
	if ptr.Pointer() != nil {
		C.QInputDialog_SetIntRange(C.QtObjectPtr(ptr.Pointer()), C.int(min), C.int(max))
	}
}

func (ptr *QInputDialog) SetOption(option QInputDialog__InputDialogOption, on bool) {
	if ptr.Pointer() != nil {
		C.QInputDialog_SetOption(C.QtObjectPtr(ptr.Pointer()), C.int(option), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QInputDialog) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QInputDialog_SetVisible(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QInputDialog) TestOption(option QInputDialog__InputDialogOption) bool {
	if ptr.Pointer() != nil {
		return C.QInputDialog_TestOption(C.QtObjectPtr(ptr.Pointer()), C.int(option)) != 0
	}
	return false
}

func (ptr *QInputDialog) ConnectTextValueChanged(f func(text string)) {
	if ptr.Pointer() != nil {
		C.QInputDialog_ConnectTextValueChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "textValueChanged", f)
	}
}

func (ptr *QInputDialog) DisconnectTextValueChanged() {
	if ptr.Pointer() != nil {
		C.QInputDialog_DisconnectTextValueChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "textValueChanged")
	}
}

//export callbackQInputDialogTextValueChanged
func callbackQInputDialogTextValueChanged(ptrName *C.char, text *C.char) {
	qt.GetSignal(C.GoString(ptrName), "textValueChanged").(func(string))(C.GoString(text))
}

func (ptr *QInputDialog) ConnectTextValueSelected(f func(text string)) {
	if ptr.Pointer() != nil {
		C.QInputDialog_ConnectTextValueSelected(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "textValueSelected", f)
	}
}

func (ptr *QInputDialog) DisconnectTextValueSelected() {
	if ptr.Pointer() != nil {
		C.QInputDialog_DisconnectTextValueSelected(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "textValueSelected")
	}
}

//export callbackQInputDialogTextValueSelected
func callbackQInputDialogTextValueSelected(ptrName *C.char, text *C.char) {
	qt.GetSignal(C.GoString(ptrName), "textValueSelected").(func(string))(C.GoString(text))
}

func (ptr *QInputDialog) DestroyQInputDialog() {
	if ptr.Pointer() != nil {
		C.QInputDialog_DestroyQInputDialog(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
