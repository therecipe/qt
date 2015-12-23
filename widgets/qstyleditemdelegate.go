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
func callbackQStyledItemDelegateInitStyleOption(ptrName *C.char, option unsafe.Pointer, index unsafe.Pointer) bool {
	defer qt.Recovering("callback QStyledItemDelegate::initStyleOption")

	if signal := qt.GetSignal(C.GoString(ptrName), "initStyleOption"); signal != nil {
		signal.(func(*QStyleOptionViewItem, *core.QModelIndex))(NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index))
		return true
	}
	return false

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
func callbackQStyledItemDelegateSetEditorData(ptrName *C.char, editor unsafe.Pointer, index unsafe.Pointer) bool {
	defer qt.Recovering("callback QStyledItemDelegate::setEditorData")

	if signal := qt.GetSignal(C.GoString(ptrName), "setEditorData"); signal != nil {
		signal.(func(*QWidget, *core.QModelIndex))(NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
		return true
	}
	return false

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
func callbackQStyledItemDelegateSetModelData(ptrName *C.char, editor unsafe.Pointer, model unsafe.Pointer, index unsafe.Pointer) bool {
	defer qt.Recovering("callback QStyledItemDelegate::setModelData")

	if signal := qt.GetSignal(C.GoString(ptrName), "setModelData"); signal != nil {
		signal.(func(*QWidget, *core.QAbstractItemModel, *core.QModelIndex))(NewQWidgetFromPointer(editor), core.NewQAbstractItemModelFromPointer(model), core.NewQModelIndexFromPointer(index))
		return true
	}
	return false

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
func callbackQStyledItemDelegateDestroyEditor(ptrName *C.char, editor unsafe.Pointer, index unsafe.Pointer) bool {
	defer qt.Recovering("callback QStyledItemDelegate::destroyEditor")

	if signal := qt.GetSignal(C.GoString(ptrName), "destroyEditor"); signal != nil {
		signal.(func(*QWidget, *core.QModelIndex))(NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
		return true
	}
	return false

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
func callbackQStyledItemDelegateTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStyledItemDelegate::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQStyledItemDelegateChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStyledItemDelegate::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQStyledItemDelegateCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStyledItemDelegate::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
