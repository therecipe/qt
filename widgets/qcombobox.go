package widgets

//#include "qcombobox.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"strings"
	"unsafe"
)

type QComboBox struct {
	QWidget
}

type QComboBoxITF interface {
	QWidgetITF
	QComboBoxPTR() *QComboBox
}

func PointerFromQComboBox(ptr QComboBoxITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QComboBoxPTR().Pointer()
	}
	return nil
}

func QComboBoxFromPointer(ptr unsafe.Pointer) *QComboBox {
	var n = new(QComboBox)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QComboBox_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QComboBox) QComboBoxPTR() *QComboBox {
	return ptr
}

//QComboBox::InsertPolicy
type QComboBox__InsertPolicy int

var (
	QComboBox__NoInsert             = QComboBox__InsertPolicy(0)
	QComboBox__InsertAtTop          = QComboBox__InsertPolicy(1)
	QComboBox__InsertAtCurrent      = QComboBox__InsertPolicy(2)
	QComboBox__InsertAtBottom       = QComboBox__InsertPolicy(3)
	QComboBox__InsertAfterCurrent   = QComboBox__InsertPolicy(4)
	QComboBox__InsertBeforeCurrent  = QComboBox__InsertPolicy(5)
	QComboBox__InsertAlphabetically = QComboBox__InsertPolicy(6)
)

//QComboBox::SizeAdjustPolicy
type QComboBox__SizeAdjustPolicy int

var (
	QComboBox__AdjustToContents                      = QComboBox__SizeAdjustPolicy(0)
	QComboBox__AdjustToContentsOnFirstShow           = QComboBox__SizeAdjustPolicy(1)
	QComboBox__AdjustToMinimumContentsLength         = QComboBox__SizeAdjustPolicy(2)
	QComboBox__AdjustToMinimumContentsLengthWithIcon = QComboBox__SizeAdjustPolicy(3)
)

func (ptr *QComboBox) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QComboBox_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QComboBox) CurrentData(role int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QComboBox_CurrentData(C.QtObjectPtr(ptr.Pointer()), C.int(role)))
	}
	return ""
}

func (ptr *QComboBox) CurrentIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QComboBox_CurrentIndex(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QComboBox) CurrentText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QComboBox_CurrentText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QComboBox) DuplicatesEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QComboBox_DuplicatesEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QComboBox) HasFrame() bool {
	if ptr.Pointer() != nil {
		return C.QComboBox_HasFrame(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QComboBox) InsertPolicy() QComboBox__InsertPolicy {
	if ptr.Pointer() != nil {
		return QComboBox__InsertPolicy(C.QComboBox_InsertPolicy(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QComboBox) IsEditable() bool {
	if ptr.Pointer() != nil {
		return C.QComboBox_IsEditable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QComboBox) MaxCount() int {
	if ptr.Pointer() != nil {
		return int(C.QComboBox_MaxCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QComboBox) MaxVisibleItems() int {
	if ptr.Pointer() != nil {
		return int(C.QComboBox_MaxVisibleItems(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QComboBox) MinimumContentsLength() int {
	if ptr.Pointer() != nil {
		return int(C.QComboBox_MinimumContentsLength(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QComboBox) ModelColumn() int {
	if ptr.Pointer() != nil {
		return int(C.QComboBox_ModelColumn(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QComboBox) SetCompleter(completer QCompleterITF) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetCompleter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQCompleter(completer)))
	}
}

func (ptr *QComboBox) SetCurrentIndex(index int) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetCurrentIndex(C.QtObjectPtr(ptr.Pointer()), C.int(index))
	}
}

func (ptr *QComboBox) SetCurrentText(text string) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetCurrentText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QComboBox) SetDuplicatesEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetDuplicatesEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QComboBox) SetEditable(editable bool) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetEditable(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(editable)))
	}
}

func (ptr *QComboBox) SetFrame(v bool) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetFrame(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QComboBox) SetIconSize(size core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetIconSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(size)))
	}
}

func (ptr *QComboBox) SetInsertPolicy(policy QComboBox__InsertPolicy) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetInsertPolicy(C.QtObjectPtr(ptr.Pointer()), C.int(policy))
	}
}

func (ptr *QComboBox) SetMaxCount(max int) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetMaxCount(C.QtObjectPtr(ptr.Pointer()), C.int(max))
	}
}

func (ptr *QComboBox) SetMaxVisibleItems(maxItems int) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetMaxVisibleItems(C.QtObjectPtr(ptr.Pointer()), C.int(maxItems))
	}
}

func (ptr *QComboBox) SetMinimumContentsLength(characters int) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetMinimumContentsLength(C.QtObjectPtr(ptr.Pointer()), C.int(characters))
	}
}

func (ptr *QComboBox) SetModelColumn(visibleColumn int) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetModelColumn(C.QtObjectPtr(ptr.Pointer()), C.int(visibleColumn))
	}
}

func (ptr *QComboBox) SetSizeAdjustPolicy(policy QComboBox__SizeAdjustPolicy) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetSizeAdjustPolicy(C.QtObjectPtr(ptr.Pointer()), C.int(policy))
	}
}

func (ptr *QComboBox) SetValidator(validator gui.QValidatorITF) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetValidator(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQValidator(validator)))
	}
}

func (ptr *QComboBox) SizeAdjustPolicy() QComboBox__SizeAdjustPolicy {
	if ptr.Pointer() != nil {
		return QComboBox__SizeAdjustPolicy(C.QComboBox_SizeAdjustPolicy(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQComboBox(parent QWidgetITF) *QComboBox {
	return QComboBoxFromPointer(unsafe.Pointer(C.QComboBox_NewQComboBox(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QComboBox) ConnectActivated(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QComboBox_ConnectActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "activated", f)
	}
}

func (ptr *QComboBox) DisconnectActivated() {
	if ptr.Pointer() != nil {
		C.QComboBox_DisconnectActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "activated")
	}
}

//export callbackQComboBoxActivated
func callbackQComboBoxActivated(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "activated").(func(int))(int(index))
}

func (ptr *QComboBox) AddItem2(icon gui.QIconITF, text string, userData string) {
	if ptr.Pointer() != nil {
		C.QComboBox_AddItem2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQIcon(icon)), C.CString(text), C.CString(userData))
	}
}

func (ptr *QComboBox) AddItem(text string, userData string) {
	if ptr.Pointer() != nil {
		C.QComboBox_AddItem(C.QtObjectPtr(ptr.Pointer()), C.CString(text), C.CString(userData))
	}
}

func (ptr *QComboBox) AddItems(texts []string) {
	if ptr.Pointer() != nil {
		C.QComboBox_AddItems(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(texts, "|")))
	}
}

func (ptr *QComboBox) Clear() {
	if ptr.Pointer() != nil {
		C.QComboBox_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QComboBox) ClearEditText() {
	if ptr.Pointer() != nil {
		C.QComboBox_ClearEditText(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QComboBox) Completer() *QCompleter {
	if ptr.Pointer() != nil {
		return QCompleterFromPointer(unsafe.Pointer(C.QComboBox_Completer(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QComboBox) ConnectCurrentIndexChanged(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QComboBox_ConnectCurrentIndexChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "currentIndexChanged", f)
	}
}

func (ptr *QComboBox) DisconnectCurrentIndexChanged() {
	if ptr.Pointer() != nil {
		C.QComboBox_DisconnectCurrentIndexChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "currentIndexChanged")
	}
}

//export callbackQComboBoxCurrentIndexChanged
func callbackQComboBoxCurrentIndexChanged(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "currentIndexChanged").(func(int))(int(index))
}

func (ptr *QComboBox) ConnectCurrentTextChanged(f func(text string)) {
	if ptr.Pointer() != nil {
		C.QComboBox_ConnectCurrentTextChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "currentTextChanged", f)
	}
}

func (ptr *QComboBox) DisconnectCurrentTextChanged() {
	if ptr.Pointer() != nil {
		C.QComboBox_DisconnectCurrentTextChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "currentTextChanged")
	}
}

//export callbackQComboBoxCurrentTextChanged
func callbackQComboBoxCurrentTextChanged(ptrName *C.char, text *C.char) {
	qt.GetSignal(C.GoString(ptrName), "currentTextChanged").(func(string))(C.GoString(text))
}

func (ptr *QComboBox) ConnectEditTextChanged(f func(text string)) {
	if ptr.Pointer() != nil {
		C.QComboBox_ConnectEditTextChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "editTextChanged", f)
	}
}

func (ptr *QComboBox) DisconnectEditTextChanged() {
	if ptr.Pointer() != nil {
		C.QComboBox_DisconnectEditTextChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "editTextChanged")
	}
}

//export callbackQComboBoxEditTextChanged
func callbackQComboBoxEditTextChanged(ptrName *C.char, text *C.char) {
	qt.GetSignal(C.GoString(ptrName), "editTextChanged").(func(string))(C.GoString(text))
}

func (ptr *QComboBox) Event(event core.QEventITF) bool {
	if ptr.Pointer() != nil {
		return C.QComboBox_Event(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQEvent(event))) != 0
	}
	return false
}

func (ptr *QComboBox) FindData(data string, role int, flags core.Qt__MatchFlag) int {
	if ptr.Pointer() != nil {
		return int(C.QComboBox_FindData(C.QtObjectPtr(ptr.Pointer()), C.CString(data), C.int(role), C.int(flags)))
	}
	return 0
}

func (ptr *QComboBox) FindText(text string, flags core.Qt__MatchFlag) int {
	if ptr.Pointer() != nil {
		return int(C.QComboBox_FindText(C.QtObjectPtr(ptr.Pointer()), C.CString(text), C.int(flags)))
	}
	return 0
}

func (ptr *QComboBox) HidePopup() {
	if ptr.Pointer() != nil {
		C.QComboBox_HidePopup(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QComboBox) ConnectHighlighted(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QComboBox_ConnectHighlighted(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "highlighted", f)
	}
}

func (ptr *QComboBox) DisconnectHighlighted() {
	if ptr.Pointer() != nil {
		C.QComboBox_DisconnectHighlighted(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "highlighted")
	}
}

//export callbackQComboBoxHighlighted
func callbackQComboBoxHighlighted(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "highlighted").(func(int))(int(index))
}

func (ptr *QComboBox) InputMethodQuery(query core.Qt__InputMethodQuery) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QComboBox_InputMethodQuery(C.QtObjectPtr(ptr.Pointer()), C.int(query)))
	}
	return ""
}

func (ptr *QComboBox) InsertItem2(index int, icon gui.QIconITF, text string, userData string) {
	if ptr.Pointer() != nil {
		C.QComboBox_InsertItem2(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.QtObjectPtr(gui.PointerFromQIcon(icon)), C.CString(text), C.CString(userData))
	}
}

func (ptr *QComboBox) InsertItem(index int, text string, userData string) {
	if ptr.Pointer() != nil {
		C.QComboBox_InsertItem(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.CString(text), C.CString(userData))
	}
}

func (ptr *QComboBox) InsertItems(index int, list []string) {
	if ptr.Pointer() != nil {
		C.QComboBox_InsertItems(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.CString(strings.Join(list, "|")))
	}
}

func (ptr *QComboBox) InsertSeparator(index int) {
	if ptr.Pointer() != nil {
		C.QComboBox_InsertSeparator(C.QtObjectPtr(ptr.Pointer()), C.int(index))
	}
}

func (ptr *QComboBox) ItemData(index int, role int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QComboBox_ItemData(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.int(role)))
	}
	return ""
}

func (ptr *QComboBox) ItemDelegate() *QAbstractItemDelegate {
	if ptr.Pointer() != nil {
		return QAbstractItemDelegateFromPointer(unsafe.Pointer(C.QComboBox_ItemDelegate(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QComboBox) ItemText(index int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QComboBox_ItemText(C.QtObjectPtr(ptr.Pointer()), C.int(index)))
	}
	return ""
}

func (ptr *QComboBox) LineEdit() *QLineEdit {
	if ptr.Pointer() != nil {
		return QLineEditFromPointer(unsafe.Pointer(C.QComboBox_LineEdit(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QComboBox) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		return core.QAbstractItemModelFromPointer(unsafe.Pointer(C.QComboBox_Model(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QComboBox) RemoveItem(index int) {
	if ptr.Pointer() != nil {
		C.QComboBox_RemoveItem(C.QtObjectPtr(ptr.Pointer()), C.int(index))
	}
}

func (ptr *QComboBox) RootModelIndex() *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.QModelIndexFromPointer(unsafe.Pointer(C.QComboBox_RootModelIndex(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QComboBox) SetEditText(text string) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetEditText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QComboBox) SetItemData(index int, value string, role int) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetItemData(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.CString(value), C.int(role))
	}
}

func (ptr *QComboBox) SetItemDelegate(delegate QAbstractItemDelegateITF) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetItemDelegate(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractItemDelegate(delegate)))
	}
}

func (ptr *QComboBox) SetItemIcon(index int, icon gui.QIconITF) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetItemIcon(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.QtObjectPtr(gui.PointerFromQIcon(icon)))
	}
}

func (ptr *QComboBox) SetItemText(index int, text string) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetItemText(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.CString(text))
	}
}

func (ptr *QComboBox) SetLineEdit(edit QLineEditITF) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetLineEdit(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLineEdit(edit)))
	}
}

func (ptr *QComboBox) SetModel(model core.QAbstractItemModelITF) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetModel(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQAbstractItemModel(model)))
	}
}

func (ptr *QComboBox) SetRootModelIndex(index core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetRootModelIndex(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)))
	}
}

func (ptr *QComboBox) SetView(itemView QAbstractItemViewITF) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetView(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractItemView(itemView)))
	}
}

func (ptr *QComboBox) ShowPopup() {
	if ptr.Pointer() != nil {
		C.QComboBox_ShowPopup(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QComboBox) Validator() *gui.QValidator {
	if ptr.Pointer() != nil {
		return gui.QValidatorFromPointer(unsafe.Pointer(C.QComboBox_Validator(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QComboBox) View() *QAbstractItemView {
	if ptr.Pointer() != nil {
		return QAbstractItemViewFromPointer(unsafe.Pointer(C.QComboBox_View(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QComboBox) DestroyQComboBox() {
	if ptr.Pointer() != nil {
		C.QComboBox_DestroyQComboBox(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
