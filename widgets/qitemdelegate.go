package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QItemDelegate struct {
	QAbstractItemDelegate
}

type QItemDelegate_ITF interface {
	QAbstractItemDelegate_ITF
	QItemDelegate_PTR() *QItemDelegate
}

func PointerFromQItemDelegate(ptr QItemDelegate_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QItemDelegate_PTR().Pointer()
	}
	return nil
}

func NewQItemDelegateFromPointer(ptr unsafe.Pointer) *QItemDelegate {
	var n = new(QItemDelegate)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QItemDelegate_") {
		n.SetObjectName("QItemDelegate_" + qt.Identifier())
	}
	return n
}

func (ptr *QItemDelegate) QItemDelegate_PTR() *QItemDelegate {
	return ptr
}

func (ptr *QItemDelegate) EventFilter(editor core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QItemDelegate::eventFilter")

	if ptr.Pointer() != nil {
		return C.QItemDelegate_EventFilter(ptr.Pointer(), core.PointerFromQObject(editor), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QItemDelegate) HasClipping() bool {
	defer qt.Recovering("QItemDelegate::hasClipping")

	if ptr.Pointer() != nil {
		return C.QItemDelegate_HasClipping(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QItemDelegate) SetClipping(clip bool) {
	defer qt.Recovering("QItemDelegate::setClipping")

	if ptr.Pointer() != nil {
		C.QItemDelegate_SetClipping(ptr.Pointer(), C.int(qt.GoBoolToInt(clip)))
	}
}

func NewQItemDelegate(parent core.QObject_ITF) *QItemDelegate {
	defer qt.Recovering("QItemDelegate::QItemDelegate")

	return NewQItemDelegateFromPointer(C.QItemDelegate_NewQItemDelegate(core.PointerFromQObject(parent)))
}

func (ptr *QItemDelegate) CreateEditor(parent QWidget_ITF, option QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) *QWidget {
	defer qt.Recovering("QItemDelegate::createEditor")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QItemDelegate_CreateEditor(ptr.Pointer(), PointerFromQWidget(parent), PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QItemDelegate) EditorEvent(event core.QEvent_ITF, model core.QAbstractItemModel_ITF, option QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) bool {
	defer qt.Recovering("QItemDelegate::editorEvent")

	if ptr.Pointer() != nil {
		return C.QItemDelegate_EditorEvent(ptr.Pointer(), core.PointerFromQEvent(event), core.PointerFromQAbstractItemModel(model), PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QItemDelegate) ItemEditorFactory() *QItemEditorFactory {
	defer qt.Recovering("QItemDelegate::itemEditorFactory")

	if ptr.Pointer() != nil {
		return NewQItemEditorFactoryFromPointer(C.QItemDelegate_ItemEditorFactory(ptr.Pointer()))
	}
	return nil
}

func (ptr *QItemDelegate) ConnectSetEditorData(f func(editor *QWidget, index *core.QModelIndex)) {
	defer qt.Recovering("connect QItemDelegate::setEditorData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setEditorData", f)
	}
}

func (ptr *QItemDelegate) DisconnectSetEditorData() {
	defer qt.Recovering("disconnect QItemDelegate::setEditorData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setEditorData")
	}
}

//export callbackQItemDelegateSetEditorData
func callbackQItemDelegateSetEditorData(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QItemDelegate::setEditorData")

	if signal := qt.GetSignal(C.GoString(ptrName), "setEditorData"); signal != nil {
		signal.(func(*QWidget, *core.QModelIndex))(NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
	} else {
		NewQItemDelegateFromPointer(ptr).SetEditorDataDefault(NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QItemDelegate) SetEditorData(editor QWidget_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QItemDelegate::setEditorData")

	if ptr.Pointer() != nil {
		C.QItemDelegate_SetEditorData(ptr.Pointer(), PointerFromQWidget(editor), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QItemDelegate) SetEditorDataDefault(editor QWidget_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QItemDelegate::setEditorData")

	if ptr.Pointer() != nil {
		C.QItemDelegate_SetEditorDataDefault(ptr.Pointer(), PointerFromQWidget(editor), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QItemDelegate) SetItemEditorFactory(factory QItemEditorFactory_ITF) {
	defer qt.Recovering("QItemDelegate::setItemEditorFactory")

	if ptr.Pointer() != nil {
		C.QItemDelegate_SetItemEditorFactory(ptr.Pointer(), PointerFromQItemEditorFactory(factory))
	}
}

func (ptr *QItemDelegate) ConnectSetModelData(f func(editor *QWidget, model *core.QAbstractItemModel, index *core.QModelIndex)) {
	defer qt.Recovering("connect QItemDelegate::setModelData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setModelData", f)
	}
}

func (ptr *QItemDelegate) DisconnectSetModelData() {
	defer qt.Recovering("disconnect QItemDelegate::setModelData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setModelData")
	}
}

//export callbackQItemDelegateSetModelData
func callbackQItemDelegateSetModelData(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer, model unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QItemDelegate::setModelData")

	if signal := qt.GetSignal(C.GoString(ptrName), "setModelData"); signal != nil {
		signal.(func(*QWidget, *core.QAbstractItemModel, *core.QModelIndex))(NewQWidgetFromPointer(editor), core.NewQAbstractItemModelFromPointer(model), core.NewQModelIndexFromPointer(index))
	} else {
		NewQItemDelegateFromPointer(ptr).SetModelDataDefault(NewQWidgetFromPointer(editor), core.NewQAbstractItemModelFromPointer(model), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QItemDelegate) SetModelData(editor QWidget_ITF, model core.QAbstractItemModel_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QItemDelegate::setModelData")

	if ptr.Pointer() != nil {
		C.QItemDelegate_SetModelData(ptr.Pointer(), PointerFromQWidget(editor), core.PointerFromQAbstractItemModel(model), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QItemDelegate) SetModelDataDefault(editor QWidget_ITF, model core.QAbstractItemModel_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QItemDelegate::setModelData")

	if ptr.Pointer() != nil {
		C.QItemDelegate_SetModelDataDefault(ptr.Pointer(), PointerFromQWidget(editor), core.PointerFromQAbstractItemModel(model), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QItemDelegate) SizeHint(option QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) *core.QSize {
	defer qt.Recovering("QItemDelegate::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QItemDelegate_SizeHint(ptr.Pointer(), PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QItemDelegate) DestroyQItemDelegate() {
	defer qt.Recovering("QItemDelegate::~QItemDelegate")

	if ptr.Pointer() != nil {
		C.QItemDelegate_DestroyQItemDelegate(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QItemDelegate) ConnectDestroyEditor(f func(editor *QWidget, index *core.QModelIndex)) {
	defer qt.Recovering("connect QItemDelegate::destroyEditor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "destroyEditor", f)
	}
}

func (ptr *QItemDelegate) DisconnectDestroyEditor() {
	defer qt.Recovering("disconnect QItemDelegate::destroyEditor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "destroyEditor")
	}
}

//export callbackQItemDelegateDestroyEditor
func callbackQItemDelegateDestroyEditor(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QItemDelegate::destroyEditor")

	if signal := qt.GetSignal(C.GoString(ptrName), "destroyEditor"); signal != nil {
		signal.(func(*QWidget, *core.QModelIndex))(NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
	} else {
		NewQItemDelegateFromPointer(ptr).DestroyEditorDefault(NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QItemDelegate) DestroyEditor(editor QWidget_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QItemDelegate::destroyEditor")

	if ptr.Pointer() != nil {
		C.QItemDelegate_DestroyEditor(ptr.Pointer(), PointerFromQWidget(editor), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QItemDelegate) DestroyEditorDefault(editor QWidget_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QItemDelegate::destroyEditor")

	if ptr.Pointer() != nil {
		C.QItemDelegate_DestroyEditorDefault(ptr.Pointer(), PointerFromQWidget(editor), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QItemDelegate) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QItemDelegate::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QItemDelegate) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QItemDelegate::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQItemDelegateTimerEvent
func callbackQItemDelegateTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QItemDelegate::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQItemDelegateFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QItemDelegate) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QItemDelegate::timerEvent")

	if ptr.Pointer() != nil {
		C.QItemDelegate_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QItemDelegate) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QItemDelegate::timerEvent")

	if ptr.Pointer() != nil {
		C.QItemDelegate_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QItemDelegate) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QItemDelegate::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QItemDelegate) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QItemDelegate::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQItemDelegateChildEvent
func callbackQItemDelegateChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QItemDelegate::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQItemDelegateFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QItemDelegate) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QItemDelegate::childEvent")

	if ptr.Pointer() != nil {
		C.QItemDelegate_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QItemDelegate) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QItemDelegate::childEvent")

	if ptr.Pointer() != nil {
		C.QItemDelegate_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QItemDelegate) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QItemDelegate::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QItemDelegate) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QItemDelegate::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQItemDelegateCustomEvent
func callbackQItemDelegateCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QItemDelegate::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQItemDelegateFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QItemDelegate) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QItemDelegate::customEvent")

	if ptr.Pointer() != nil {
		C.QItemDelegate_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QItemDelegate) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QItemDelegate::customEvent")

	if ptr.Pointer() != nil {
		C.QItemDelegate_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
