package widgets

//#include "qcompleter.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QCompleter struct {
	core.QObject
}

type QCompleterITF interface {
	core.QObjectITF
	QCompleterPTR() *QCompleter
}

func PointerFromQCompleter(ptr QCompleterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCompleterPTR().Pointer()
	}
	return nil
}

func QCompleterFromPointer(ptr unsafe.Pointer) *QCompleter {
	var n = new(QCompleter)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCompleter_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCompleter) QCompleterPTR() *QCompleter {
	return ptr
}

//QCompleter::CompletionMode
type QCompleter__CompletionMode int

var (
	QCompleter__PopupCompletion           = QCompleter__CompletionMode(0)
	QCompleter__UnfilteredPopupCompletion = QCompleter__CompletionMode(1)
	QCompleter__InlineCompletion          = QCompleter__CompletionMode(2)
)

//QCompleter::ModelSorting
type QCompleter__ModelSorting int

var (
	QCompleter__UnsortedModel                = QCompleter__ModelSorting(0)
	QCompleter__CaseSensitivelySortedModel   = QCompleter__ModelSorting(1)
	QCompleter__CaseInsensitivelySortedModel = QCompleter__ModelSorting(2)
)

func (ptr *QCompleter) CaseSensitivity() core.Qt__CaseSensitivity {
	if ptr.Pointer() != nil {
		return core.Qt__CaseSensitivity(C.QCompleter_CaseSensitivity(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCompleter) CompletionColumn() int {
	if ptr.Pointer() != nil {
		return int(C.QCompleter_CompletionColumn(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCompleter) CompletionMode() QCompleter__CompletionMode {
	if ptr.Pointer() != nil {
		return QCompleter__CompletionMode(C.QCompleter_CompletionMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCompleter) CompletionPrefix() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QCompleter_CompletionPrefix(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QCompleter) CompletionRole() int {
	if ptr.Pointer() != nil {
		return int(C.QCompleter_CompletionRole(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCompleter) FilterMode() core.Qt__MatchFlag {
	if ptr.Pointer() != nil {
		return core.Qt__MatchFlag(C.QCompleter_FilterMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCompleter) MaxVisibleItems() int {
	if ptr.Pointer() != nil {
		return int(C.QCompleter_MaxVisibleItems(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCompleter) ModelSorting() QCompleter__ModelSorting {
	if ptr.Pointer() != nil {
		return QCompleter__ModelSorting(C.QCompleter_ModelSorting(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCompleter) SetCaseSensitivity(caseSensitivity core.Qt__CaseSensitivity) {
	if ptr.Pointer() != nil {
		C.QCompleter_SetCaseSensitivity(C.QtObjectPtr(ptr.Pointer()), C.int(caseSensitivity))
	}
}

func (ptr *QCompleter) SetCompletionColumn(column int) {
	if ptr.Pointer() != nil {
		C.QCompleter_SetCompletionColumn(C.QtObjectPtr(ptr.Pointer()), C.int(column))
	}
}

func (ptr *QCompleter) SetCompletionMode(mode QCompleter__CompletionMode) {
	if ptr.Pointer() != nil {
		C.QCompleter_SetCompletionMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QCompleter) SetCompletionPrefix(prefix string) {
	if ptr.Pointer() != nil {
		C.QCompleter_SetCompletionPrefix(C.QtObjectPtr(ptr.Pointer()), C.CString(prefix))
	}
}

func (ptr *QCompleter) SetCompletionRole(role int) {
	if ptr.Pointer() != nil {
		C.QCompleter_SetCompletionRole(C.QtObjectPtr(ptr.Pointer()), C.int(role))
	}
}

func (ptr *QCompleter) SetFilterMode(filterMode core.Qt__MatchFlag) {
	if ptr.Pointer() != nil {
		C.QCompleter_SetFilterMode(C.QtObjectPtr(ptr.Pointer()), C.int(filterMode))
	}
}

func (ptr *QCompleter) SetMaxVisibleItems(maxItems int) {
	if ptr.Pointer() != nil {
		C.QCompleter_SetMaxVisibleItems(C.QtObjectPtr(ptr.Pointer()), C.int(maxItems))
	}
}

func (ptr *QCompleter) SetModelSorting(sorting QCompleter__ModelSorting) {
	if ptr.Pointer() != nil {
		C.QCompleter_SetModelSorting(C.QtObjectPtr(ptr.Pointer()), C.int(sorting))
	}
}

func (ptr *QCompleter) SetWrapAround(wrap bool) {
	if ptr.Pointer() != nil {
		C.QCompleter_SetWrapAround(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(wrap)))
	}
}

func (ptr *QCompleter) WrapAround() bool {
	if ptr.Pointer() != nil {
		return C.QCompleter_WrapAround(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func NewQCompleter2(model core.QAbstractItemModelITF, parent core.QObjectITF) *QCompleter {
	return QCompleterFromPointer(unsafe.Pointer(C.QCompleter_NewQCompleter2(C.QtObjectPtr(core.PointerFromQAbstractItemModel(model)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQCompleter(parent core.QObjectITF) *QCompleter {
	return QCompleterFromPointer(unsafe.Pointer(C.QCompleter_NewQCompleter(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQCompleter3(list []string, parent core.QObjectITF) *QCompleter {
	return QCompleterFromPointer(unsafe.Pointer(C.QCompleter_NewQCompleter3(C.CString(strings.Join(list, "|")), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QCompleter) ConnectActivated(f func(text string)) {
	if ptr.Pointer() != nil {
		C.QCompleter_ConnectActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "activated", f)
	}
}

func (ptr *QCompleter) DisconnectActivated() {
	if ptr.Pointer() != nil {
		C.QCompleter_DisconnectActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "activated")
	}
}

//export callbackQCompleterActivated
func callbackQCompleterActivated(ptrName *C.char, text *C.char) {
	qt.GetSignal(C.GoString(ptrName), "activated").(func(string))(C.GoString(text))
}

func (ptr *QCompleter) Complete(rect core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QCompleter_Complete(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rect)))
	}
}

func (ptr *QCompleter) CompletionCount() int {
	if ptr.Pointer() != nil {
		return int(C.QCompleter_CompletionCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCompleter) CompletionModel() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		return core.QAbstractItemModelFromPointer(unsafe.Pointer(C.QCompleter_CompletionModel(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QCompleter) CurrentCompletion() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QCompleter_CurrentCompletion(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QCompleter) CurrentIndex() *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.QModelIndexFromPointer(unsafe.Pointer(C.QCompleter_CurrentIndex(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QCompleter) CurrentRow() int {
	if ptr.Pointer() != nil {
		return int(C.QCompleter_CurrentRow(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCompleter) ConnectHighlighted(f func(text string)) {
	if ptr.Pointer() != nil {
		C.QCompleter_ConnectHighlighted(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "highlighted", f)
	}
}

func (ptr *QCompleter) DisconnectHighlighted() {
	if ptr.Pointer() != nil {
		C.QCompleter_DisconnectHighlighted(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "highlighted")
	}
}

//export callbackQCompleterHighlighted
func callbackQCompleterHighlighted(ptrName *C.char, text *C.char) {
	qt.GetSignal(C.GoString(ptrName), "highlighted").(func(string))(C.GoString(text))
}

func (ptr *QCompleter) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		return core.QAbstractItemModelFromPointer(unsafe.Pointer(C.QCompleter_Model(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QCompleter) PathFromIndex(index core.QModelIndexITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QCompleter_PathFromIndex(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index))))
	}
	return ""
}

func (ptr *QCompleter) Popup() *QAbstractItemView {
	if ptr.Pointer() != nil {
		return QAbstractItemViewFromPointer(unsafe.Pointer(C.QCompleter_Popup(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QCompleter) SetCurrentRow(row int) bool {
	if ptr.Pointer() != nil {
		return C.QCompleter_SetCurrentRow(C.QtObjectPtr(ptr.Pointer()), C.int(row)) != 0
	}
	return false
}

func (ptr *QCompleter) SetModel(model core.QAbstractItemModelITF) {
	if ptr.Pointer() != nil {
		C.QCompleter_SetModel(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQAbstractItemModel(model)))
	}
}

func (ptr *QCompleter) SetPopup(popup QAbstractItemViewITF) {
	if ptr.Pointer() != nil {
		C.QCompleter_SetPopup(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractItemView(popup)))
	}
}

func (ptr *QCompleter) SetWidget(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QCompleter_SetWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QCompleter) SplitPath(path string) []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QCompleter_SplitPath(C.QtObjectPtr(ptr.Pointer()), C.CString(path))), "|")
	}
	return make([]string, 0)
}

func (ptr *QCompleter) Widget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QCompleter_Widget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QCompleter) DestroyQCompleter() {
	if ptr.Pointer() != nil {
		C.QCompleter_DestroyQCompleter(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
