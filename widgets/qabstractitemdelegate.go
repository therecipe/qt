package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QAbstractItemDelegate struct {
	core.QObject
}

type QAbstractItemDelegate_ITF interface {
	core.QObject_ITF
	QAbstractItemDelegate_PTR() *QAbstractItemDelegate
}

func PointerFromQAbstractItemDelegate(ptr QAbstractItemDelegate_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractItemDelegate_PTR().Pointer()
	}
	return nil
}

func NewQAbstractItemDelegateFromPointer(ptr unsafe.Pointer) *QAbstractItemDelegate {
	var n = new(QAbstractItemDelegate)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractItemDelegate_") {
		n.SetObjectName("QAbstractItemDelegate_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractItemDelegate) QAbstractItemDelegate_PTR() *QAbstractItemDelegate {
	return ptr
}

//QAbstractItemDelegate::EndEditHint
type QAbstractItemDelegate__EndEditHint int64

const (
	QAbstractItemDelegate__NoHint           = QAbstractItemDelegate__EndEditHint(0)
	QAbstractItemDelegate__EditNextItem     = QAbstractItemDelegate__EndEditHint(1)
	QAbstractItemDelegate__EditPreviousItem = QAbstractItemDelegate__EndEditHint(2)
	QAbstractItemDelegate__SubmitModelCache = QAbstractItemDelegate__EndEditHint(3)
	QAbstractItemDelegate__RevertModelCache = QAbstractItemDelegate__EndEditHint(4)
)

func (ptr *QAbstractItemDelegate) ConnectCloseEditor(f func(editor *QWidget, hint QAbstractItemDelegate__EndEditHint)) {
	defer qt.Recovering("connect QAbstractItemDelegate::closeEditor")

	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_ConnectCloseEditor(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "closeEditor", f)
	}
}

func (ptr *QAbstractItemDelegate) DisconnectCloseEditor() {
	defer qt.Recovering("disconnect QAbstractItemDelegate::closeEditor")

	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_DisconnectCloseEditor(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "closeEditor")
	}
}

//export callbackQAbstractItemDelegateCloseEditor
func callbackQAbstractItemDelegateCloseEditor(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer, hint C.int) {
	defer qt.Recovering("callback QAbstractItemDelegate::closeEditor")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEditor"); signal != nil {
		signal.(func(*QWidget, QAbstractItemDelegate__EndEditHint))(NewQWidgetFromPointer(editor), QAbstractItemDelegate__EndEditHint(hint))
	}

}

func (ptr *QAbstractItemDelegate) CloseEditor(editor QWidget_ITF, hint QAbstractItemDelegate__EndEditHint) {
	defer qt.Recovering("QAbstractItemDelegate::closeEditor")

	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_CloseEditor(ptr.Pointer(), PointerFromQWidget(editor), C.int(hint))
	}
}

func (ptr *QAbstractItemDelegate) ConnectCommitData(f func(editor *QWidget)) {
	defer qt.Recovering("connect QAbstractItemDelegate::commitData")

	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_ConnectCommitData(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "commitData", f)
	}
}

func (ptr *QAbstractItemDelegate) DisconnectCommitData() {
	defer qt.Recovering("disconnect QAbstractItemDelegate::commitData")

	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_DisconnectCommitData(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "commitData")
	}
}

//export callbackQAbstractItemDelegateCommitData
func callbackQAbstractItemDelegateCommitData(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractItemDelegate::commitData")

	if signal := qt.GetSignal(C.GoString(ptrName), "commitData"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(editor))
	}

}

func (ptr *QAbstractItemDelegate) CommitData(editor QWidget_ITF) {
	defer qt.Recovering("QAbstractItemDelegate::commitData")

	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_CommitData(ptr.Pointer(), PointerFromQWidget(editor))
	}
}

func (ptr *QAbstractItemDelegate) CreateEditor(parent QWidget_ITF, option QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) *QWidget {
	defer qt.Recovering("QAbstractItemDelegate::createEditor")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QAbstractItemDelegate_CreateEditor(ptr.Pointer(), PointerFromQWidget(parent), PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QAbstractItemDelegate) ConnectDestroyEditor(f func(editor *QWidget, index *core.QModelIndex)) {
	defer qt.Recovering("connect QAbstractItemDelegate::destroyEditor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "destroyEditor", f)
	}
}

func (ptr *QAbstractItemDelegate) DisconnectDestroyEditor() {
	defer qt.Recovering("disconnect QAbstractItemDelegate::destroyEditor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "destroyEditor")
	}
}

//export callbackQAbstractItemDelegateDestroyEditor
func callbackQAbstractItemDelegateDestroyEditor(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractItemDelegate::destroyEditor")

	if signal := qt.GetSignal(C.GoString(ptrName), "destroyEditor"); signal != nil {
		signal.(func(*QWidget, *core.QModelIndex))(NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
	} else {
		NewQAbstractItemDelegateFromPointer(ptr).DestroyEditorDefault(NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QAbstractItemDelegate) DestroyEditor(editor QWidget_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QAbstractItemDelegate::destroyEditor")

	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_DestroyEditor(ptr.Pointer(), PointerFromQWidget(editor), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QAbstractItemDelegate) DestroyEditorDefault(editor QWidget_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QAbstractItemDelegate::destroyEditor")

	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_DestroyEditorDefault(ptr.Pointer(), PointerFromQWidget(editor), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QAbstractItemDelegate) EditorEvent(event core.QEvent_ITF, model core.QAbstractItemModel_ITF, option QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) bool {
	defer qt.Recovering("QAbstractItemDelegate::editorEvent")

	if ptr.Pointer() != nil {
		return C.QAbstractItemDelegate_EditorEvent(ptr.Pointer(), core.PointerFromQEvent(event), core.PointerFromQAbstractItemModel(model), PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QAbstractItemDelegate) HelpEvent(event gui.QHelpEvent_ITF, view QAbstractItemView_ITF, option QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) bool {
	defer qt.Recovering("QAbstractItemDelegate::helpEvent")

	if ptr.Pointer() != nil {
		return C.QAbstractItemDelegate_HelpEvent(ptr.Pointer(), gui.PointerFromQHelpEvent(event), PointerFromQAbstractItemView(view), PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QAbstractItemDelegate) Paint(painter gui.QPainter_ITF, option QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QAbstractItemDelegate::paint")

	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QAbstractItemDelegate) ConnectSetEditorData(f func(editor *QWidget, index *core.QModelIndex)) {
	defer qt.Recovering("connect QAbstractItemDelegate::setEditorData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setEditorData", f)
	}
}

func (ptr *QAbstractItemDelegate) DisconnectSetEditorData() {
	defer qt.Recovering("disconnect QAbstractItemDelegate::setEditorData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setEditorData")
	}
}

//export callbackQAbstractItemDelegateSetEditorData
func callbackQAbstractItemDelegateSetEditorData(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractItemDelegate::setEditorData")

	if signal := qt.GetSignal(C.GoString(ptrName), "setEditorData"); signal != nil {
		signal.(func(*QWidget, *core.QModelIndex))(NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
	} else {
		NewQAbstractItemDelegateFromPointer(ptr).SetEditorDataDefault(NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QAbstractItemDelegate) SetEditorData(editor QWidget_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QAbstractItemDelegate::setEditorData")

	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_SetEditorData(ptr.Pointer(), PointerFromQWidget(editor), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QAbstractItemDelegate) SetEditorDataDefault(editor QWidget_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QAbstractItemDelegate::setEditorData")

	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_SetEditorDataDefault(ptr.Pointer(), PointerFromQWidget(editor), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QAbstractItemDelegate) ConnectSetModelData(f func(editor *QWidget, model *core.QAbstractItemModel, index *core.QModelIndex)) {
	defer qt.Recovering("connect QAbstractItemDelegate::setModelData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setModelData", f)
	}
}

func (ptr *QAbstractItemDelegate) DisconnectSetModelData() {
	defer qt.Recovering("disconnect QAbstractItemDelegate::setModelData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setModelData")
	}
}

//export callbackQAbstractItemDelegateSetModelData
func callbackQAbstractItemDelegateSetModelData(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer, model unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractItemDelegate::setModelData")

	if signal := qt.GetSignal(C.GoString(ptrName), "setModelData"); signal != nil {
		signal.(func(*QWidget, *core.QAbstractItemModel, *core.QModelIndex))(NewQWidgetFromPointer(editor), core.NewQAbstractItemModelFromPointer(model), core.NewQModelIndexFromPointer(index))
	} else {
		NewQAbstractItemDelegateFromPointer(ptr).SetModelDataDefault(NewQWidgetFromPointer(editor), core.NewQAbstractItemModelFromPointer(model), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QAbstractItemDelegate) SetModelData(editor QWidget_ITF, model core.QAbstractItemModel_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QAbstractItemDelegate::setModelData")

	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_SetModelData(ptr.Pointer(), PointerFromQWidget(editor), core.PointerFromQAbstractItemModel(model), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QAbstractItemDelegate) SetModelDataDefault(editor QWidget_ITF, model core.QAbstractItemModel_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QAbstractItemDelegate::setModelData")

	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_SetModelDataDefault(ptr.Pointer(), PointerFromQWidget(editor), core.PointerFromQAbstractItemModel(model), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QAbstractItemDelegate) SizeHint(option QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) *core.QSize {
	defer qt.Recovering("QAbstractItemDelegate::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QAbstractItemDelegate_SizeHint(ptr.Pointer(), PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QAbstractItemDelegate) ConnectSizeHintChanged(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QAbstractItemDelegate::sizeHintChanged")

	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_ConnectSizeHintChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sizeHintChanged", f)
	}
}

func (ptr *QAbstractItemDelegate) DisconnectSizeHintChanged() {
	defer qt.Recovering("disconnect QAbstractItemDelegate::sizeHintChanged")

	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_DisconnectSizeHintChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sizeHintChanged")
	}
}

//export callbackQAbstractItemDelegateSizeHintChanged
func callbackQAbstractItemDelegateSizeHintChanged(ptr unsafe.Pointer, ptrName *C.char, index unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractItemDelegate::sizeHintChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "sizeHintChanged"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	}

}

func (ptr *QAbstractItemDelegate) SizeHintChanged(index core.QModelIndex_ITF) {
	defer qt.Recovering("QAbstractItemDelegate::sizeHintChanged")

	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_SizeHintChanged(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QAbstractItemDelegate) DestroyQAbstractItemDelegate() {
	defer qt.Recovering("QAbstractItemDelegate::~QAbstractItemDelegate")

	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_DestroyQAbstractItemDelegate(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractItemDelegate) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAbstractItemDelegate::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAbstractItemDelegate) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAbstractItemDelegate::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAbstractItemDelegateTimerEvent
func callbackQAbstractItemDelegateTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractItemDelegate::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractItemDelegateFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractItemDelegate) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractItemDelegate::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractItemDelegate) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractItemDelegate::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractItemDelegate) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAbstractItemDelegate::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAbstractItemDelegate) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAbstractItemDelegate::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAbstractItemDelegateChildEvent
func callbackQAbstractItemDelegateChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractItemDelegate::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractItemDelegateFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractItemDelegate) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractItemDelegate::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractItemDelegate) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractItemDelegate::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractItemDelegate) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractItemDelegate::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAbstractItemDelegate) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAbstractItemDelegate::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAbstractItemDelegateCustomEvent
func callbackQAbstractItemDelegateCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractItemDelegate::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractItemDelegateFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractItemDelegate) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractItemDelegate::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractItemDelegate) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractItemDelegate::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
