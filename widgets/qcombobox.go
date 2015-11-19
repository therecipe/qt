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

type QComboBox_ITF interface {
	QWidget_ITF
	QComboBox_PTR() *QComboBox
}

func PointerFromQComboBox(ptr QComboBox_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QComboBox_PTR().Pointer()
	}
	return nil
}

func NewQComboBoxFromPointer(ptr unsafe.Pointer) *QComboBox {
	var n = new(QComboBox)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QComboBox_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QComboBox) QComboBox_PTR() *QComboBox {
	return ptr
}

//QComboBox::InsertPolicy
type QComboBox__InsertPolicy int64

const (
	QComboBox__NoInsert             = QComboBox__InsertPolicy(0)
	QComboBox__InsertAtTop          = QComboBox__InsertPolicy(1)
	QComboBox__InsertAtCurrent      = QComboBox__InsertPolicy(2)
	QComboBox__InsertAtBottom       = QComboBox__InsertPolicy(3)
	QComboBox__InsertAfterCurrent   = QComboBox__InsertPolicy(4)
	QComboBox__InsertBeforeCurrent  = QComboBox__InsertPolicy(5)
	QComboBox__InsertAlphabetically = QComboBox__InsertPolicy(6)
)

//QComboBox::SizeAdjustPolicy
type QComboBox__SizeAdjustPolicy int64

const (
	QComboBox__AdjustToContents                      = QComboBox__SizeAdjustPolicy(0)
	QComboBox__AdjustToContentsOnFirstShow           = QComboBox__SizeAdjustPolicy(1)
	QComboBox__AdjustToMinimumContentsLength         = QComboBox__SizeAdjustPolicy(2)
	QComboBox__AdjustToMinimumContentsLengthWithIcon = QComboBox__SizeAdjustPolicy(3)
)

func (ptr *QComboBox) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QComboBox_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QComboBox) CurrentData(role int) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QComboBox_CurrentData(ptr.Pointer(), C.int(role)))
	}
	return nil
}

func (ptr *QComboBox) CurrentIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QComboBox_CurrentIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QComboBox) CurrentText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QComboBox_CurrentText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QComboBox) DuplicatesEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QComboBox_DuplicatesEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QComboBox) HasFrame() bool {
	if ptr.Pointer() != nil {
		return C.QComboBox_HasFrame(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QComboBox) InsertPolicy() QComboBox__InsertPolicy {
	if ptr.Pointer() != nil {
		return QComboBox__InsertPolicy(C.QComboBox_InsertPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QComboBox) IsEditable() bool {
	if ptr.Pointer() != nil {
		return C.QComboBox_IsEditable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QComboBox) MaxCount() int {
	if ptr.Pointer() != nil {
		return int(C.QComboBox_MaxCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QComboBox) MaxVisibleItems() int {
	if ptr.Pointer() != nil {
		return int(C.QComboBox_MaxVisibleItems(ptr.Pointer()))
	}
	return 0
}

func (ptr *QComboBox) MinimumContentsLength() int {
	if ptr.Pointer() != nil {
		return int(C.QComboBox_MinimumContentsLength(ptr.Pointer()))
	}
	return 0
}

func (ptr *QComboBox) ModelColumn() int {
	if ptr.Pointer() != nil {
		return int(C.QComboBox_ModelColumn(ptr.Pointer()))
	}
	return 0
}

func (ptr *QComboBox) SetCompleter(completer QCompleter_ITF) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetCompleter(ptr.Pointer(), PointerFromQCompleter(completer))
	}
}

func (ptr *QComboBox) SetCurrentIndex(index int) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetCurrentIndex(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QComboBox) SetCurrentText(text string) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetCurrentText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QComboBox) SetDuplicatesEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetDuplicatesEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QComboBox) SetEditable(editable bool) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetEditable(ptr.Pointer(), C.int(qt.GoBoolToInt(editable)))
	}
}

func (ptr *QComboBox) SetFrame(v bool) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetFrame(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QComboBox) SetIconSize(size core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetIconSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QComboBox) SetInsertPolicy(policy QComboBox__InsertPolicy) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetInsertPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QComboBox) SetMaxCount(max int) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetMaxCount(ptr.Pointer(), C.int(max))
	}
}

func (ptr *QComboBox) SetMaxVisibleItems(maxItems int) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetMaxVisibleItems(ptr.Pointer(), C.int(maxItems))
	}
}

func (ptr *QComboBox) SetMinimumContentsLength(characters int) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetMinimumContentsLength(ptr.Pointer(), C.int(characters))
	}
}

func (ptr *QComboBox) SetModelColumn(visibleColumn int) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetModelColumn(ptr.Pointer(), C.int(visibleColumn))
	}
}

func (ptr *QComboBox) SetSizeAdjustPolicy(policy QComboBox__SizeAdjustPolicy) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetSizeAdjustPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QComboBox) SetValidator(validator gui.QValidator_ITF) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetValidator(ptr.Pointer(), gui.PointerFromQValidator(validator))
	}
}

func (ptr *QComboBox) SizeAdjustPolicy() QComboBox__SizeAdjustPolicy {
	if ptr.Pointer() != nil {
		return QComboBox__SizeAdjustPolicy(C.QComboBox_SizeAdjustPolicy(ptr.Pointer()))
	}
	return 0
}

func NewQComboBox(parent QWidget_ITF) *QComboBox {
	return NewQComboBoxFromPointer(C.QComboBox_NewQComboBox(PointerFromQWidget(parent)))
}

func (ptr *QComboBox) ConnectActivated(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QComboBox_ConnectActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activated", f)
	}
}

func (ptr *QComboBox) DisconnectActivated() {
	if ptr.Pointer() != nil {
		C.QComboBox_DisconnectActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activated")
	}
}

//export callbackQComboBoxActivated
func callbackQComboBoxActivated(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "activated").(func(int))(int(index))
}

func (ptr *QComboBox) AddItem2(icon gui.QIcon_ITF, text string, userData core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QComboBox_AddItem2(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(text), core.PointerFromQVariant(userData))
	}
}

func (ptr *QComboBox) AddItem(text string, userData core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QComboBox_AddItem(ptr.Pointer(), C.CString(text), core.PointerFromQVariant(userData))
	}
}

func (ptr *QComboBox) AddItems(texts []string) {
	if ptr.Pointer() != nil {
		C.QComboBox_AddItems(ptr.Pointer(), C.CString(strings.Join(texts, "|")))
	}
}

func (ptr *QComboBox) Clear() {
	if ptr.Pointer() != nil {
		C.QComboBox_Clear(ptr.Pointer())
	}
}

func (ptr *QComboBox) ClearEditText() {
	if ptr.Pointer() != nil {
		C.QComboBox_ClearEditText(ptr.Pointer())
	}
}

func (ptr *QComboBox) Completer() *QCompleter {
	if ptr.Pointer() != nil {
		return NewQCompleterFromPointer(C.QComboBox_Completer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QComboBox) ConnectCurrentIndexChanged(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QComboBox_ConnectCurrentIndexChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentIndexChanged", f)
	}
}

func (ptr *QComboBox) DisconnectCurrentIndexChanged() {
	if ptr.Pointer() != nil {
		C.QComboBox_DisconnectCurrentIndexChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentIndexChanged")
	}
}

//export callbackQComboBoxCurrentIndexChanged
func callbackQComboBoxCurrentIndexChanged(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "currentIndexChanged").(func(int))(int(index))
}

func (ptr *QComboBox) ConnectCurrentTextChanged(f func(text string)) {
	if ptr.Pointer() != nil {
		C.QComboBox_ConnectCurrentTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentTextChanged", f)
	}
}

func (ptr *QComboBox) DisconnectCurrentTextChanged() {
	if ptr.Pointer() != nil {
		C.QComboBox_DisconnectCurrentTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentTextChanged")
	}
}

//export callbackQComboBoxCurrentTextChanged
func callbackQComboBoxCurrentTextChanged(ptrName *C.char, text *C.char) {
	qt.GetSignal(C.GoString(ptrName), "currentTextChanged").(func(string))(C.GoString(text))
}

func (ptr *QComboBox) ConnectEditTextChanged(f func(text string)) {
	if ptr.Pointer() != nil {
		C.QComboBox_ConnectEditTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "editTextChanged", f)
	}
}

func (ptr *QComboBox) DisconnectEditTextChanged() {
	if ptr.Pointer() != nil {
		C.QComboBox_DisconnectEditTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "editTextChanged")
	}
}

//export callbackQComboBoxEditTextChanged
func callbackQComboBoxEditTextChanged(ptrName *C.char, text *C.char) {
	qt.GetSignal(C.GoString(ptrName), "editTextChanged").(func(string))(C.GoString(text))
}

func (ptr *QComboBox) Event(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QComboBox_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QComboBox) FindData(data core.QVariant_ITF, role int, flags core.Qt__MatchFlag) int {
	if ptr.Pointer() != nil {
		return int(C.QComboBox_FindData(ptr.Pointer(), core.PointerFromQVariant(data), C.int(role), C.int(flags)))
	}
	return 0
}

func (ptr *QComboBox) FindText(text string, flags core.Qt__MatchFlag) int {
	if ptr.Pointer() != nil {
		return int(C.QComboBox_FindText(ptr.Pointer(), C.CString(text), C.int(flags)))
	}
	return 0
}

func (ptr *QComboBox) HidePopup() {
	if ptr.Pointer() != nil {
		C.QComboBox_HidePopup(ptr.Pointer())
	}
}

func (ptr *QComboBox) ConnectHighlighted(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QComboBox_ConnectHighlighted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "highlighted", f)
	}
}

func (ptr *QComboBox) DisconnectHighlighted() {
	if ptr.Pointer() != nil {
		C.QComboBox_DisconnectHighlighted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "highlighted")
	}
}

//export callbackQComboBoxHighlighted
func callbackQComboBoxHighlighted(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "highlighted").(func(int))(int(index))
}

func (ptr *QComboBox) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QComboBox_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QComboBox) InsertItem2(index int, icon gui.QIcon_ITF, text string, userData core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QComboBox_InsertItem2(ptr.Pointer(), C.int(index), gui.PointerFromQIcon(icon), C.CString(text), core.PointerFromQVariant(userData))
	}
}

func (ptr *QComboBox) InsertItem(index int, text string, userData core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QComboBox_InsertItem(ptr.Pointer(), C.int(index), C.CString(text), core.PointerFromQVariant(userData))
	}
}

func (ptr *QComboBox) InsertItems(index int, list []string) {
	if ptr.Pointer() != nil {
		C.QComboBox_InsertItems(ptr.Pointer(), C.int(index), C.CString(strings.Join(list, "|")))
	}
}

func (ptr *QComboBox) InsertSeparator(index int) {
	if ptr.Pointer() != nil {
		C.QComboBox_InsertSeparator(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QComboBox) ItemData(index int, role int) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QComboBox_ItemData(ptr.Pointer(), C.int(index), C.int(role)))
	}
	return nil
}

func (ptr *QComboBox) ItemDelegate() *QAbstractItemDelegate {
	if ptr.Pointer() != nil {
		return NewQAbstractItemDelegateFromPointer(C.QComboBox_ItemDelegate(ptr.Pointer()))
	}
	return nil
}

func (ptr *QComboBox) ItemText(index int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QComboBox_ItemText(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QComboBox) LineEdit() *QLineEdit {
	if ptr.Pointer() != nil {
		return NewQLineEditFromPointer(C.QComboBox_LineEdit(ptr.Pointer()))
	}
	return nil
}

func (ptr *QComboBox) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		return core.NewQAbstractItemModelFromPointer(C.QComboBox_Model(ptr.Pointer()))
	}
	return nil
}

func (ptr *QComboBox) RemoveItem(index int) {
	if ptr.Pointer() != nil {
		C.QComboBox_RemoveItem(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QComboBox) RootModelIndex() *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QComboBox_RootModelIndex(ptr.Pointer()))
	}
	return nil
}

func (ptr *QComboBox) SetEditText(text string) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetEditText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QComboBox) SetItemData(index int, value core.QVariant_ITF, role int) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetItemData(ptr.Pointer(), C.int(index), core.PointerFromQVariant(value), C.int(role))
	}
}

func (ptr *QComboBox) SetItemDelegate(delegate QAbstractItemDelegate_ITF) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetItemDelegate(ptr.Pointer(), PointerFromQAbstractItemDelegate(delegate))
	}
}

func (ptr *QComboBox) SetItemIcon(index int, icon gui.QIcon_ITF) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetItemIcon(ptr.Pointer(), C.int(index), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QComboBox) SetItemText(index int, text string) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetItemText(ptr.Pointer(), C.int(index), C.CString(text))
	}
}

func (ptr *QComboBox) SetLineEdit(edit QLineEdit_ITF) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetLineEdit(ptr.Pointer(), PointerFromQLineEdit(edit))
	}
}

func (ptr *QComboBox) SetModel(model core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QComboBox) SetRootModelIndex(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetRootModelIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QComboBox) SetView(itemView QAbstractItemView_ITF) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetView(ptr.Pointer(), PointerFromQAbstractItemView(itemView))
	}
}

func (ptr *QComboBox) ShowPopup() {
	if ptr.Pointer() != nil {
		C.QComboBox_ShowPopup(ptr.Pointer())
	}
}

func (ptr *QComboBox) Validator() *gui.QValidator {
	if ptr.Pointer() != nil {
		return gui.NewQValidatorFromPointer(C.QComboBox_Validator(ptr.Pointer()))
	}
	return nil
}

func (ptr *QComboBox) View() *QAbstractItemView {
	if ptr.Pointer() != nil {
		return NewQAbstractItemViewFromPointer(C.QComboBox_View(ptr.Pointer()))
	}
	return nil
}

func (ptr *QComboBox) DestroyQComboBox() {
	if ptr.Pointer() != nil {
		C.QComboBox_DestroyQComboBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
