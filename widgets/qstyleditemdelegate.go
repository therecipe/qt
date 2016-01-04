package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QStyledItemDelegate struct {
	QAbstractItemDelegate
}

type QStyledItemDelegate_ITF interface {
	QAbstractItemDelegate_ITF
	QStyledItemDelegate_PTR() *QStyledItemDelegate
}

func PointerFromQStyledItemDelegate(ptr QStyledItemDelegate_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyledItemDelegate_PTR().Pointer()
	}
	return nil
}

func NewQStyledItemDelegateFromPointer(ptr unsafe.Pointer) *QStyledItemDelegate {
	var n = new(QStyledItemDelegate)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QStyledItemDelegate_") {
		n.SetObjectName("QStyledItemDelegate_" + qt.Identifier())
	}
	return n
}

func (ptr *QStyledItemDelegate) QStyledItemDelegate_PTR() *QStyledItemDelegate {
	return ptr
}

func (ptr *QStyledItemDelegate) EventFilter(editor core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QStyledItemDelegate::eventFilter")

	if ptr.Pointer() != nil {
		return C.QStyledItemDelegate_EventFilter(ptr.Pointer(), core.PointerFromQObject(editor), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func NewQStyledItemDelegate(parent core.QObject_ITF) *QStyledItemDelegate {
	defer qt.Recovering("QStyledItemDelegate::QStyledItemDelegate")

	return NewQStyledItemDelegateFromPointer(C.QStyledItemDelegate_NewQStyledItemDelegate(core.PointerFromQObject(parent)))
}

func (ptr *QStyledItemDelegate) CreateEditor(parent QWidget_ITF, option QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) *QWidget {
	defer qt.Recovering("QStyledItemDelegate::createEditor")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QStyledItemDelegate_CreateEditor(ptr.Pointer(), PointerFromQWidget(parent), PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QStyledItemDelegate) DisplayText(value core.QVariant_ITF, locale core.QLocale_ITF) string {
	defer qt.Recovering("QStyledItemDelegate::displayText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QStyledItemDelegate_DisplayText(ptr.Pointer(), core.PointerFromQVariant(value), core.PointerFromQLocale(locale)))
	}
	return ""
}

func (ptr *QStyledItemDelegate) EditorEvent(event core.QEvent_ITF, model core.QAbstractItemModel_ITF, option QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) bool {
	defer qt.Recovering("QStyledItemDelegate::editorEvent")

	if ptr.Pointer() != nil {
		return C.QStyledItemDelegate_EditorEvent(ptr.Pointer(), core.PointerFromQEvent(event), core.PointerFromQAbstractItemModel(model), PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QStyledItemDelegate) ConnectInitStyleOption(f func(option *QStyleOptionViewItem, index *core.QModelIndex)) {
	defer qt.Recovering("connect QStyledItemDelegate::initStyleOption")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initStyleOption", f)
	}
}

func (ptr *QStyledItemDelegate) DisconnectInitStyleOption() {
	defer qt.Recovering("disconnect QStyledItemDelegate::initStyleOption")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initStyleOption")
	}
}

//export callbackQStyledItemDelegateInitStyleOption
func callbackQStyledItemDelegateInitStyleOption(ptr unsafe.Pointer, ptrName *C.char, option unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QStyledItemDelegate::initStyleOption")

	if signal := qt.GetSignal(C.GoString(ptrName), "initStyleOption"); signal != nil {
		signal.(func(*QStyleOptionViewItem, *core.QModelIndex))(NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index))
	} else {
		NewQStyledItemDelegateFromPointer(ptr).InitStyleOptionDefault(NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QStyledItemDelegate) InitStyleOption(option QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QStyledItemDelegate::initStyleOption")

	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_InitStyleOption(ptr.Pointer(), PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QStyledItemDelegate) InitStyleOptionDefault(option QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QStyledItemDelegate::initStyleOption")

	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_InitStyleOptionDefault(ptr.Pointer(), PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QStyledItemDelegate) ItemEditorFactory() *QItemEditorFactory {
	defer qt.Recovering("QStyledItemDelegate::itemEditorFactory")

	if ptr.Pointer() != nil {
		return NewQItemEditorFactoryFromPointer(C.QStyledItemDelegate_ItemEditorFactory(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStyledItemDelegate) ConnectSetEditorData(f func(editor *QWidget, index *core.QModelIndex)) {
	defer qt.Recovering("connect QStyledItemDelegate::setEditorData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setEditorData", f)
	}
}

func (ptr *QStyledItemDelegate) DisconnectSetEditorData() {
	defer qt.Recovering("disconnect QStyledItemDelegate::setEditorData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setEditorData")
	}
}

//export callbackQStyledItemDelegateSetEditorData
func callbackQStyledItemDelegateSetEditorData(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QStyledItemDelegate::setEditorData")

	if signal := qt.GetSignal(C.GoString(ptrName), "setEditorData"); signal != nil {
		signal.(func(*QWidget, *core.QModelIndex))(NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
	} else {
		NewQStyledItemDelegateFromPointer(ptr).SetEditorDataDefault(NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QStyledItemDelegate) SetEditorData(editor QWidget_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QStyledItemDelegate::setEditorData")

	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_SetEditorData(ptr.Pointer(), PointerFromQWidget(editor), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QStyledItemDelegate) SetEditorDataDefault(editor QWidget_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QStyledItemDelegate::setEditorData")

	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_SetEditorDataDefault(ptr.Pointer(), PointerFromQWidget(editor), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QStyledItemDelegate) SetItemEditorFactory(factory QItemEditorFactory_ITF) {
	defer qt.Recovering("QStyledItemDelegate::setItemEditorFactory")

	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_SetItemEditorFactory(ptr.Pointer(), PointerFromQItemEditorFactory(factory))
	}
}

func (ptr *QStyledItemDelegate) ConnectSetModelData(f func(editor *QWidget, model *core.QAbstractItemModel, index *core.QModelIndex)) {
	defer qt.Recovering("connect QStyledItemDelegate::setModelData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setModelData", f)
	}
}

func (ptr *QStyledItemDelegate) DisconnectSetModelData() {
	defer qt.Recovering("disconnect QStyledItemDelegate::setModelData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setModelData")
	}
}

//export callbackQStyledItemDelegateSetModelData
func callbackQStyledItemDelegateSetModelData(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer, model unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QStyledItemDelegate::setModelData")

	if signal := qt.GetSignal(C.GoString(ptrName), "setModelData"); signal != nil {
		signal.(func(*QWidget, *core.QAbstractItemModel, *core.QModelIndex))(NewQWidgetFromPointer(editor), core.NewQAbstractItemModelFromPointer(model), core.NewQModelIndexFromPointer(index))
	} else {
		NewQStyledItemDelegateFromPointer(ptr).SetModelDataDefault(NewQWidgetFromPointer(editor), core.NewQAbstractItemModelFromPointer(model), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QStyledItemDelegate) SetModelData(editor QWidget_ITF, model core.QAbstractItemModel_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QStyledItemDelegate::setModelData")

	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_SetModelData(ptr.Pointer(), PointerFromQWidget(editor), core.PointerFromQAbstractItemModel(model), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QStyledItemDelegate) SetModelDataDefault(editor QWidget_ITF, model core.QAbstractItemModel_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QStyledItemDelegate::setModelData")

	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_SetModelDataDefault(ptr.Pointer(), PointerFromQWidget(editor), core.PointerFromQAbstractItemModel(model), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QStyledItemDelegate) SizeHint(option QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) *core.QSize {
	defer qt.Recovering("QStyledItemDelegate::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QStyledItemDelegate_SizeHint(ptr.Pointer(), PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QStyledItemDelegate) DestroyQStyledItemDelegate() {
	defer qt.Recovering("QStyledItemDelegate::~QStyledItemDelegate")

	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_DestroyQStyledItemDelegate(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QStyledItemDelegate) ConnectDestroyEditor(f func(editor *QWidget, index *core.QModelIndex)) {
	defer qt.Recovering("connect QStyledItemDelegate::destroyEditor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "destroyEditor", f)
	}
}

func (ptr *QStyledItemDelegate) DisconnectDestroyEditor() {
	defer qt.Recovering("disconnect QStyledItemDelegate::destroyEditor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "destroyEditor")
	}
}

//export callbackQStyledItemDelegateDestroyEditor
func callbackQStyledItemDelegateDestroyEditor(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QStyledItemDelegate::destroyEditor")

	if signal := qt.GetSignal(C.GoString(ptrName), "destroyEditor"); signal != nil {
		signal.(func(*QWidget, *core.QModelIndex))(NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
	} else {
		NewQStyledItemDelegateFromPointer(ptr).DestroyEditorDefault(NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QStyledItemDelegate) DestroyEditor(editor QWidget_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QStyledItemDelegate::destroyEditor")

	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_DestroyEditor(ptr.Pointer(), PointerFromQWidget(editor), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QStyledItemDelegate) DestroyEditorDefault(editor QWidget_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QStyledItemDelegate::destroyEditor")

	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_DestroyEditorDefault(ptr.Pointer(), PointerFromQWidget(editor), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QStyledItemDelegate) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QStyledItemDelegate::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QStyledItemDelegate) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QStyledItemDelegate::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQStyledItemDelegateTimerEvent
func callbackQStyledItemDelegateTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStyledItemDelegate::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQStyledItemDelegateFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QStyledItemDelegate) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QStyledItemDelegate::timerEvent")

	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QStyledItemDelegate) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QStyledItemDelegate::timerEvent")

	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QStyledItemDelegate) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QStyledItemDelegate::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QStyledItemDelegate) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QStyledItemDelegate::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQStyledItemDelegateChildEvent
func callbackQStyledItemDelegateChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStyledItemDelegate::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQStyledItemDelegateFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QStyledItemDelegate) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QStyledItemDelegate::childEvent")

	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QStyledItemDelegate) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QStyledItemDelegate::childEvent")

	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QStyledItemDelegate) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QStyledItemDelegate::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QStyledItemDelegate) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QStyledItemDelegate::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQStyledItemDelegateCustomEvent
func callbackQStyledItemDelegateCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStyledItemDelegate::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQStyledItemDelegateFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QStyledItemDelegate) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QStyledItemDelegate::customEvent")

	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QStyledItemDelegate) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QStyledItemDelegate::customEvent")

	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
