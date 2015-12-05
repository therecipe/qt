package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"log"
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
	for len(n.ObjectName()) < len("QComboBox_") {
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::count")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QComboBox_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QComboBox) CurrentData(role int) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::currentData")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QComboBox_CurrentData(ptr.Pointer(), C.int(role)))
	}
	return nil
}

func (ptr *QComboBox) CurrentIndex() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::currentIndex")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QComboBox_CurrentIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QComboBox) CurrentText() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::currentText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QComboBox_CurrentText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QComboBox) DuplicatesEnabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::duplicatesEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QComboBox_DuplicatesEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QComboBox) HasFrame() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::hasFrame")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QComboBox_HasFrame(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QComboBox) InsertPolicy() QComboBox__InsertPolicy {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::insertPolicy")
		}
	}()

	if ptr.Pointer() != nil {
		return QComboBox__InsertPolicy(C.QComboBox_InsertPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QComboBox) IsEditable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::isEditable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QComboBox_IsEditable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QComboBox) MaxCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::maxCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QComboBox_MaxCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QComboBox) MaxVisibleItems() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::maxVisibleItems")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QComboBox_MaxVisibleItems(ptr.Pointer()))
	}
	return 0
}

func (ptr *QComboBox) MinimumContentsLength() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::minimumContentsLength")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QComboBox_MinimumContentsLength(ptr.Pointer()))
	}
	return 0
}

func (ptr *QComboBox) ModelColumn() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::modelColumn")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QComboBox_ModelColumn(ptr.Pointer()))
	}
	return 0
}

func (ptr *QComboBox) SetCompleter(completer QCompleter_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::setCompleter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_SetCompleter(ptr.Pointer(), PointerFromQCompleter(completer))
	}
}

func (ptr *QComboBox) SetCurrentIndex(index int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::setCurrentIndex")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_SetCurrentIndex(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QComboBox) SetCurrentText(text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::setCurrentText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_SetCurrentText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QComboBox) SetDuplicatesEnabled(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::setDuplicatesEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_SetDuplicatesEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QComboBox) SetEditable(editable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::setEditable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_SetEditable(ptr.Pointer(), C.int(qt.GoBoolToInt(editable)))
	}
}

func (ptr *QComboBox) SetFrame(v bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::setFrame")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_SetFrame(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QComboBox) SetIconSize(size core.QSize_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::setIconSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_SetIconSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QComboBox) SetInsertPolicy(policy QComboBox__InsertPolicy) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::setInsertPolicy")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_SetInsertPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QComboBox) SetMaxCount(max int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::setMaxCount")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_SetMaxCount(ptr.Pointer(), C.int(max))
	}
}

func (ptr *QComboBox) SetMaxVisibleItems(maxItems int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::setMaxVisibleItems")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_SetMaxVisibleItems(ptr.Pointer(), C.int(maxItems))
	}
}

func (ptr *QComboBox) SetMinimumContentsLength(characters int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::setMinimumContentsLength")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_SetMinimumContentsLength(ptr.Pointer(), C.int(characters))
	}
}

func (ptr *QComboBox) SetModelColumn(visibleColumn int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::setModelColumn")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_SetModelColumn(ptr.Pointer(), C.int(visibleColumn))
	}
}

func (ptr *QComboBox) SetSizeAdjustPolicy(policy QComboBox__SizeAdjustPolicy) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::setSizeAdjustPolicy")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_SetSizeAdjustPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QComboBox) SetValidator(validator gui.QValidator_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::setValidator")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_SetValidator(ptr.Pointer(), gui.PointerFromQValidator(validator))
	}
}

func (ptr *QComboBox) SizeAdjustPolicy() QComboBox__SizeAdjustPolicy {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::sizeAdjustPolicy")
		}
	}()

	if ptr.Pointer() != nil {
		return QComboBox__SizeAdjustPolicy(C.QComboBox_SizeAdjustPolicy(ptr.Pointer()))
	}
	return 0
}

func NewQComboBox(parent QWidget_ITF) *QComboBox {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::QComboBox")
		}
	}()

	return NewQComboBoxFromPointer(C.QComboBox_NewQComboBox(PointerFromQWidget(parent)))
}

func (ptr *QComboBox) ConnectActivated(f func(index int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::activated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_ConnectActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activated", f)
	}
}

func (ptr *QComboBox) DisconnectActivated() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::activated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_DisconnectActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activated")
	}
}

//export callbackQComboBoxActivated
func callbackQComboBoxActivated(ptrName *C.char, index C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::activated")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "activated").(func(int))(int(index))
}

func (ptr *QComboBox) AddItem2(icon gui.QIcon_ITF, text string, userData core.QVariant_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::addItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_AddItem2(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(text), core.PointerFromQVariant(userData))
	}
}

func (ptr *QComboBox) AddItem(text string, userData core.QVariant_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::addItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_AddItem(ptr.Pointer(), C.CString(text), core.PointerFromQVariant(userData))
	}
}

func (ptr *QComboBox) AddItems(texts []string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::addItems")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_AddItems(ptr.Pointer(), C.CString(strings.Join(texts, ",,,")))
	}
}

func (ptr *QComboBox) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_Clear(ptr.Pointer())
	}
}

func (ptr *QComboBox) ClearEditText() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::clearEditText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_ClearEditText(ptr.Pointer())
	}
}

func (ptr *QComboBox) Completer() *QCompleter {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::completer")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQCompleterFromPointer(C.QComboBox_Completer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QComboBox) ConnectCurrentIndexChanged(f func(index int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::currentIndexChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_ConnectCurrentIndexChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentIndexChanged", f)
	}
}

func (ptr *QComboBox) DisconnectCurrentIndexChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::currentIndexChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_DisconnectCurrentIndexChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentIndexChanged")
	}
}

//export callbackQComboBoxCurrentIndexChanged
func callbackQComboBoxCurrentIndexChanged(ptrName *C.char, index C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::currentIndexChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "currentIndexChanged").(func(int))(int(index))
}

func (ptr *QComboBox) ConnectCurrentTextChanged(f func(text string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::currentTextChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_ConnectCurrentTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentTextChanged", f)
	}
}

func (ptr *QComboBox) DisconnectCurrentTextChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::currentTextChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_DisconnectCurrentTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentTextChanged")
	}
}

//export callbackQComboBoxCurrentTextChanged
func callbackQComboBoxCurrentTextChanged(ptrName *C.char, text *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::currentTextChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "currentTextChanged").(func(string))(C.GoString(text))
}

func (ptr *QComboBox) ConnectEditTextChanged(f func(text string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::editTextChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_ConnectEditTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "editTextChanged", f)
	}
}

func (ptr *QComboBox) DisconnectEditTextChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::editTextChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_DisconnectEditTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "editTextChanged")
	}
}

//export callbackQComboBoxEditTextChanged
func callbackQComboBoxEditTextChanged(ptrName *C.char, text *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::editTextChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "editTextChanged").(func(string))(C.GoString(text))
}

func (ptr *QComboBox) Event(event core.QEvent_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::event")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QComboBox_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QComboBox) FindData(data core.QVariant_ITF, role int, flags core.Qt__MatchFlag) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::findData")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QComboBox_FindData(ptr.Pointer(), core.PointerFromQVariant(data), C.int(role), C.int(flags)))
	}
	return 0
}

func (ptr *QComboBox) FindText(text string, flags core.Qt__MatchFlag) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::findText")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QComboBox_FindText(ptr.Pointer(), C.CString(text), C.int(flags)))
	}
	return 0
}

func (ptr *QComboBox) HidePopup() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::hidePopup")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_HidePopup(ptr.Pointer())
	}
}

func (ptr *QComboBox) ConnectHighlighted(f func(index int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::highlighted")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_ConnectHighlighted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "highlighted", f)
	}
}

func (ptr *QComboBox) DisconnectHighlighted() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::highlighted")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_DisconnectHighlighted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "highlighted")
	}
}

//export callbackQComboBoxHighlighted
func callbackQComboBoxHighlighted(ptrName *C.char, index C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::highlighted")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "highlighted").(func(int))(int(index))
}

func (ptr *QComboBox) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::inputMethodQuery")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QComboBox_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QComboBox) InsertItem2(index int, icon gui.QIcon_ITF, text string, userData core.QVariant_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::insertItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_InsertItem2(ptr.Pointer(), C.int(index), gui.PointerFromQIcon(icon), C.CString(text), core.PointerFromQVariant(userData))
	}
}

func (ptr *QComboBox) InsertItem(index int, text string, userData core.QVariant_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::insertItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_InsertItem(ptr.Pointer(), C.int(index), C.CString(text), core.PointerFromQVariant(userData))
	}
}

func (ptr *QComboBox) InsertItems(index int, list []string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::insertItems")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_InsertItems(ptr.Pointer(), C.int(index), C.CString(strings.Join(list, ",,,")))
	}
}

func (ptr *QComboBox) InsertSeparator(index int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::insertSeparator")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_InsertSeparator(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QComboBox) ItemData(index int, role int) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::itemData")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QComboBox_ItemData(ptr.Pointer(), C.int(index), C.int(role)))
	}
	return nil
}

func (ptr *QComboBox) ItemDelegate() *QAbstractItemDelegate {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::itemDelegate")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAbstractItemDelegateFromPointer(C.QComboBox_ItemDelegate(ptr.Pointer()))
	}
	return nil
}

func (ptr *QComboBox) ItemText(index int) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::itemText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QComboBox_ItemText(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QComboBox) LineEdit() *QLineEdit {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::lineEdit")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQLineEditFromPointer(C.QComboBox_LineEdit(ptr.Pointer()))
	}
	return nil
}

func (ptr *QComboBox) Model() *core.QAbstractItemModel {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::model")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQAbstractItemModelFromPointer(C.QComboBox_Model(ptr.Pointer()))
	}
	return nil
}

func (ptr *QComboBox) RemoveItem(index int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::removeItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_RemoveItem(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QComboBox) RootModelIndex() *core.QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::rootModelIndex")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QComboBox_RootModelIndex(ptr.Pointer()))
	}
	return nil
}

func (ptr *QComboBox) SetEditText(text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::setEditText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_SetEditText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QComboBox) SetItemData(index int, value core.QVariant_ITF, role int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::setItemData")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_SetItemData(ptr.Pointer(), C.int(index), core.PointerFromQVariant(value), C.int(role))
	}
}

func (ptr *QComboBox) SetItemDelegate(delegate QAbstractItemDelegate_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::setItemDelegate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_SetItemDelegate(ptr.Pointer(), PointerFromQAbstractItemDelegate(delegate))
	}
}

func (ptr *QComboBox) SetItemIcon(index int, icon gui.QIcon_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::setItemIcon")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_SetItemIcon(ptr.Pointer(), C.int(index), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QComboBox) SetItemText(index int, text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::setItemText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_SetItemText(ptr.Pointer(), C.int(index), C.CString(text))
	}
}

func (ptr *QComboBox) SetLineEdit(edit QLineEdit_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::setLineEdit")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_SetLineEdit(ptr.Pointer(), PointerFromQLineEdit(edit))
	}
}

func (ptr *QComboBox) SetModel(model core.QAbstractItemModel_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::setModel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QComboBox) SetRootModelIndex(index core.QModelIndex_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::setRootModelIndex")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_SetRootModelIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QComboBox) SetView(itemView QAbstractItemView_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::setView")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_SetView(ptr.Pointer(), PointerFromQAbstractItemView(itemView))
	}
}

func (ptr *QComboBox) ShowPopup() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::showPopup")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_ShowPopup(ptr.Pointer())
	}
}

func (ptr *QComboBox) Validator() *gui.QValidator {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::validator")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQValidatorFromPointer(C.QComboBox_Validator(ptr.Pointer()))
	}
	return nil
}

func (ptr *QComboBox) View() *QAbstractItemView {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::view")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAbstractItemViewFromPointer(C.QComboBox_View(ptr.Pointer()))
	}
	return nil
}

func (ptr *QComboBox) DestroyQComboBox() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QComboBox::~QComboBox")
		}
	}()

	if ptr.Pointer() != nil {
		C.QComboBox_DestroyQComboBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
