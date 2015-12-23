package widgets

//#include "widgets.h"
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

type QCompleter_ITF interface {
	core.QObject_ITF
	QCompleter_PTR() *QCompleter
}

func PointerFromQCompleter(ptr QCompleter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCompleter_PTR().Pointer()
	}
	return nil
}

func NewQCompleterFromPointer(ptr unsafe.Pointer) *QCompleter {
	var n = new(QCompleter)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCompleter_") {
		n.SetObjectName("QCompleter_" + qt.Identifier())
	}
	return n
}

func (ptr *QCompleter) QCompleter_PTR() *QCompleter {
	return ptr
}

//QCompleter::CompletionMode
type QCompleter__CompletionMode int64

const (
	QCompleter__PopupCompletion           = QCompleter__CompletionMode(0)
	QCompleter__UnfilteredPopupCompletion = QCompleter__CompletionMode(1)
	QCompleter__InlineCompletion          = QCompleter__CompletionMode(2)
)

//QCompleter::ModelSorting
type QCompleter__ModelSorting int64

const (
	QCompleter__UnsortedModel                = QCompleter__ModelSorting(0)
	QCompleter__CaseSensitivelySortedModel   = QCompleter__ModelSorting(1)
	QCompleter__CaseInsensitivelySortedModel = QCompleter__ModelSorting(2)
)

func (ptr *QCompleter) CaseSensitivity() core.Qt__CaseSensitivity {
	defer qt.Recovering("QCompleter::caseSensitivity")

	if ptr.Pointer() != nil {
		return core.Qt__CaseSensitivity(C.QCompleter_CaseSensitivity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompleter) CompletionColumn() int {
	defer qt.Recovering("QCompleter::completionColumn")

	if ptr.Pointer() != nil {
		return int(C.QCompleter_CompletionColumn(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompleter) CompletionMode() QCompleter__CompletionMode {
	defer qt.Recovering("QCompleter::completionMode")

	if ptr.Pointer() != nil {
		return QCompleter__CompletionMode(C.QCompleter_CompletionMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompleter) CompletionPrefix() string {
	defer qt.Recovering("QCompleter::completionPrefix")

	if ptr.Pointer() != nil {
		return C.GoString(C.QCompleter_CompletionPrefix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCompleter) CompletionRole() int {
	defer qt.Recovering("QCompleter::completionRole")

	if ptr.Pointer() != nil {
		return int(C.QCompleter_CompletionRole(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompleter) FilterMode() core.Qt__MatchFlag {
	defer qt.Recovering("QCompleter::filterMode")

	if ptr.Pointer() != nil {
		return core.Qt__MatchFlag(C.QCompleter_FilterMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompleter) MaxVisibleItems() int {
	defer qt.Recovering("QCompleter::maxVisibleItems")

	if ptr.Pointer() != nil {
		return int(C.QCompleter_MaxVisibleItems(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompleter) ModelSorting() QCompleter__ModelSorting {
	defer qt.Recovering("QCompleter::modelSorting")

	if ptr.Pointer() != nil {
		return QCompleter__ModelSorting(C.QCompleter_ModelSorting(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompleter) SetCaseSensitivity(caseSensitivity core.Qt__CaseSensitivity) {
	defer qt.Recovering("QCompleter::setCaseSensitivity")

	if ptr.Pointer() != nil {
		C.QCompleter_SetCaseSensitivity(ptr.Pointer(), C.int(caseSensitivity))
	}
}

func (ptr *QCompleter) SetCompletionColumn(column int) {
	defer qt.Recovering("QCompleter::setCompletionColumn")

	if ptr.Pointer() != nil {
		C.QCompleter_SetCompletionColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QCompleter) SetCompletionMode(mode QCompleter__CompletionMode) {
	defer qt.Recovering("QCompleter::setCompletionMode")

	if ptr.Pointer() != nil {
		C.QCompleter_SetCompletionMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCompleter) SetCompletionPrefix(prefix string) {
	defer qt.Recovering("QCompleter::setCompletionPrefix")

	if ptr.Pointer() != nil {
		C.QCompleter_SetCompletionPrefix(ptr.Pointer(), C.CString(prefix))
	}
}

func (ptr *QCompleter) SetCompletionRole(role int) {
	defer qt.Recovering("QCompleter::setCompletionRole")

	if ptr.Pointer() != nil {
		C.QCompleter_SetCompletionRole(ptr.Pointer(), C.int(role))
	}
}

func (ptr *QCompleter) SetFilterMode(filterMode core.Qt__MatchFlag) {
	defer qt.Recovering("QCompleter::setFilterMode")

	if ptr.Pointer() != nil {
		C.QCompleter_SetFilterMode(ptr.Pointer(), C.int(filterMode))
	}
}

func (ptr *QCompleter) SetMaxVisibleItems(maxItems int) {
	defer qt.Recovering("QCompleter::setMaxVisibleItems")

	if ptr.Pointer() != nil {
		C.QCompleter_SetMaxVisibleItems(ptr.Pointer(), C.int(maxItems))
	}
}

func (ptr *QCompleter) SetModelSorting(sorting QCompleter__ModelSorting) {
	defer qt.Recovering("QCompleter::setModelSorting")

	if ptr.Pointer() != nil {
		C.QCompleter_SetModelSorting(ptr.Pointer(), C.int(sorting))
	}
}

func (ptr *QCompleter) SetWrapAround(wrap bool) {
	defer qt.Recovering("QCompleter::setWrapAround")

	if ptr.Pointer() != nil {
		C.QCompleter_SetWrapAround(ptr.Pointer(), C.int(qt.GoBoolToInt(wrap)))
	}
}

func (ptr *QCompleter) WrapAround() bool {
	defer qt.Recovering("QCompleter::wrapAround")

	if ptr.Pointer() != nil {
		return C.QCompleter_WrapAround(ptr.Pointer()) != 0
	}
	return false
}

func NewQCompleter2(model core.QAbstractItemModel_ITF, parent core.QObject_ITF) *QCompleter {
	defer qt.Recovering("QCompleter::QCompleter")

	return NewQCompleterFromPointer(C.QCompleter_NewQCompleter2(core.PointerFromQAbstractItemModel(model), core.PointerFromQObject(parent)))
}

func NewQCompleter(parent core.QObject_ITF) *QCompleter {
	defer qt.Recovering("QCompleter::QCompleter")

	return NewQCompleterFromPointer(C.QCompleter_NewQCompleter(core.PointerFromQObject(parent)))
}

func NewQCompleter3(list []string, parent core.QObject_ITF) *QCompleter {
	defer qt.Recovering("QCompleter::QCompleter")

	return NewQCompleterFromPointer(C.QCompleter_NewQCompleter3(C.CString(strings.Join(list, ",,,")), core.PointerFromQObject(parent)))
}

func (ptr *QCompleter) ConnectActivated2(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QCompleter::activated")

	if ptr.Pointer() != nil {
		C.QCompleter_ConnectActivated2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activated2", f)
	}
}

func (ptr *QCompleter) DisconnectActivated2() {
	defer qt.Recovering("disconnect QCompleter::activated")

	if ptr.Pointer() != nil {
		C.QCompleter_DisconnectActivated2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activated2")
	}
}

//export callbackQCompleterActivated2
func callbackQCompleterActivated2(ptrName *C.char, index unsafe.Pointer) {
	defer qt.Recovering("callback QCompleter::activated")

	if signal := qt.GetSignal(C.GoString(ptrName), "activated2"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	}

}

func (ptr *QCompleter) ConnectActivated(f func(text string)) {
	defer qt.Recovering("connect QCompleter::activated")

	if ptr.Pointer() != nil {
		C.QCompleter_ConnectActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activated", f)
	}
}

func (ptr *QCompleter) DisconnectActivated() {
	defer qt.Recovering("disconnect QCompleter::activated")

	if ptr.Pointer() != nil {
		C.QCompleter_DisconnectActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activated")
	}
}

//export callbackQCompleterActivated
func callbackQCompleterActivated(ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QCompleter::activated")

	if signal := qt.GetSignal(C.GoString(ptrName), "activated"); signal != nil {
		signal.(func(string))(C.GoString(text))
	}

}

func (ptr *QCompleter) Complete(rect core.QRect_ITF) {
	defer qt.Recovering("QCompleter::complete")

	if ptr.Pointer() != nil {
		C.QCompleter_Complete(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QCompleter) CompletionCount() int {
	defer qt.Recovering("QCompleter::completionCount")

	if ptr.Pointer() != nil {
		return int(C.QCompleter_CompletionCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompleter) CompletionModel() *core.QAbstractItemModel {
	defer qt.Recovering("QCompleter::completionModel")

	if ptr.Pointer() != nil {
		return core.NewQAbstractItemModelFromPointer(C.QCompleter_CompletionModel(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCompleter) CurrentCompletion() string {
	defer qt.Recovering("QCompleter::currentCompletion")

	if ptr.Pointer() != nil {
		return C.GoString(C.QCompleter_CurrentCompletion(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCompleter) CurrentIndex() *core.QModelIndex {
	defer qt.Recovering("QCompleter::currentIndex")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QCompleter_CurrentIndex(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCompleter) CurrentRow() int {
	defer qt.Recovering("QCompleter::currentRow")

	if ptr.Pointer() != nil {
		return int(C.QCompleter_CurrentRow(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompleter) ConnectHighlighted2(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QCompleter::highlighted")

	if ptr.Pointer() != nil {
		C.QCompleter_ConnectHighlighted2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "highlighted2", f)
	}
}

func (ptr *QCompleter) DisconnectHighlighted2() {
	defer qt.Recovering("disconnect QCompleter::highlighted")

	if ptr.Pointer() != nil {
		C.QCompleter_DisconnectHighlighted2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "highlighted2")
	}
}

//export callbackQCompleterHighlighted2
func callbackQCompleterHighlighted2(ptrName *C.char, index unsafe.Pointer) {
	defer qt.Recovering("callback QCompleter::highlighted")

	if signal := qt.GetSignal(C.GoString(ptrName), "highlighted2"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	}

}

func (ptr *QCompleter) ConnectHighlighted(f func(text string)) {
	defer qt.Recovering("connect QCompleter::highlighted")

	if ptr.Pointer() != nil {
		C.QCompleter_ConnectHighlighted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "highlighted", f)
	}
}

func (ptr *QCompleter) DisconnectHighlighted() {
	defer qt.Recovering("disconnect QCompleter::highlighted")

	if ptr.Pointer() != nil {
		C.QCompleter_DisconnectHighlighted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "highlighted")
	}
}

//export callbackQCompleterHighlighted
func callbackQCompleterHighlighted(ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QCompleter::highlighted")

	if signal := qt.GetSignal(C.GoString(ptrName), "highlighted"); signal != nil {
		signal.(func(string))(C.GoString(text))
	}

}

func (ptr *QCompleter) Model() *core.QAbstractItemModel {
	defer qt.Recovering("QCompleter::model")

	if ptr.Pointer() != nil {
		return core.NewQAbstractItemModelFromPointer(C.QCompleter_Model(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCompleter) PathFromIndex(index core.QModelIndex_ITF) string {
	defer qt.Recovering("QCompleter::pathFromIndex")

	if ptr.Pointer() != nil {
		return C.GoString(C.QCompleter_PathFromIndex(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return ""
}

func (ptr *QCompleter) Popup() *QAbstractItemView {
	defer qt.Recovering("QCompleter::popup")

	if ptr.Pointer() != nil {
		return NewQAbstractItemViewFromPointer(C.QCompleter_Popup(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCompleter) SetCurrentRow(row int) bool {
	defer qt.Recovering("QCompleter::setCurrentRow")

	if ptr.Pointer() != nil {
		return C.QCompleter_SetCurrentRow(ptr.Pointer(), C.int(row)) != 0
	}
	return false
}

func (ptr *QCompleter) SetModel(model core.QAbstractItemModel_ITF) {
	defer qt.Recovering("QCompleter::setModel")

	if ptr.Pointer() != nil {
		C.QCompleter_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QCompleter) SetPopup(popup QAbstractItemView_ITF) {
	defer qt.Recovering("QCompleter::setPopup")

	if ptr.Pointer() != nil {
		C.QCompleter_SetPopup(ptr.Pointer(), PointerFromQAbstractItemView(popup))
	}
}

func (ptr *QCompleter) SetWidget(widget QWidget_ITF) {
	defer qt.Recovering("QCompleter::setWidget")

	if ptr.Pointer() != nil {
		C.QCompleter_SetWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QCompleter) SplitPath(path string) []string {
	defer qt.Recovering("QCompleter::splitPath")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QCompleter_SplitPath(ptr.Pointer(), C.CString(path))), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QCompleter) Widget() *QWidget {
	defer qt.Recovering("QCompleter::widget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QCompleter_Widget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCompleter) DestroyQCompleter() {
	defer qt.Recovering("QCompleter::~QCompleter")

	if ptr.Pointer() != nil {
		C.QCompleter_DestroyQCompleter(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCompleter) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCompleter::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCompleter) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCompleter::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCompleterTimerEvent
func callbackQCompleterTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCompleter::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCompleter) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCompleter::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCompleter) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCompleter::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCompleterChildEvent
func callbackQCompleterChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCompleter::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCompleter) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCompleter::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCompleter) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCompleter::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCompleterCustomEvent
func callbackQCompleterCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCompleter::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
