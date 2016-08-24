package main

//#include <stdint.h>
//#include <stdlib.h>
//#include "moc.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"runtime"
	"unsafe"
)

type Delegate_ITF interface {
	widgets.QStyledItemDelegate_ITF
	Delegate_PTR() *Delegate
}

func (p *Delegate) Delegate_PTR() *Delegate {
	return p
}

func (p *Delegate) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QStyledItemDelegate_PTR().Pointer()
	}
	return nil
}

func (p *Delegate) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QStyledItemDelegate_PTR().SetPointer(ptr)
	}
}

func PointerFromDelegate(ptr Delegate_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Delegate_PTR().Pointer()
	}
	return nil
}

func NewDelegateFromPointer(ptr unsafe.Pointer) *Delegate {
	var n = new(Delegate)
	n.SetPointer(ptr)
	return n
}
func NewDelegate(parent core.QObject_ITF) *Delegate {
	defer qt.Recovering("Delegate::Delegate")

	var tmpValue = NewDelegateFromPointer(C.Delegate_NewDelegate(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) {})
	}
	return tmpValue
}

func (ptr *Delegate) DestroyDelegate() {
	defer qt.Recovering("Delegate::~Delegate")

	if ptr.Pointer() != nil {
		C.Delegate_DestroyDelegate(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackDelegate_CreateEditor
func callbackDelegate_CreateEditor(ptr unsafe.Pointer, parent unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback Delegate::createEditor")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "Delegate::createEditor"); signal != nil {
		return widgets.PointerFromQWidget(signal.(func(*widgets.QWidget, *widgets.QStyleOptionViewItem, *core.QModelIndex) *widgets.QWidget)(widgets.NewQWidgetFromPointer(parent), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index)))
	}

	return widgets.PointerFromQWidget(NewDelegateFromPointer(ptr).CreateEditorDefault(widgets.NewQWidgetFromPointer(parent), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index)))
}

func (ptr *Delegate) ConnectCreateEditor(f func(parent *widgets.QWidget, option *widgets.QStyleOptionViewItem, index *core.QModelIndex) *widgets.QWidget) {
	defer qt.Recovering("connect Delegate::createEditor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::createEditor", f)
	}
}

func (ptr *Delegate) DisconnectCreateEditor() {
	defer qt.Recovering("disconnect Delegate::createEditor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::createEditor")
	}
}

func (ptr *Delegate) CreateEditor(parent widgets.QWidget_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) *widgets.QWidget {
	defer qt.Recovering("Delegate::createEditor")

	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.Delegate_CreateEditor(ptr.Pointer(), widgets.PointerFromQWidget(parent), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

func (ptr *Delegate) CreateEditorDefault(parent widgets.QWidget_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) *widgets.QWidget {
	defer qt.Recovering("Delegate::createEditor")

	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.Delegate_CreateEditorDefault(ptr.Pointer(), widgets.PointerFromQWidget(parent), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

//export callbackDelegate_DisplayText
func callbackDelegate_DisplayText(ptr unsafe.Pointer, value unsafe.Pointer, locale unsafe.Pointer) *C.char {
	defer qt.Recovering("callback Delegate::displayText")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "Delegate::displayText"); signal != nil {
		return C.CString(signal.(func(*core.QVariant, *core.QLocale) string)(core.NewQVariantFromPointer(value), core.NewQLocaleFromPointer(locale)))
	}

	return C.CString(NewDelegateFromPointer(ptr).DisplayTextDefault(core.NewQVariantFromPointer(value), core.NewQLocaleFromPointer(locale)))
}

func (ptr *Delegate) ConnectDisplayText(f func(value *core.QVariant, locale *core.QLocale) string) {
	defer qt.Recovering("connect Delegate::displayText")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::displayText", f)
	}
}

func (ptr *Delegate) DisconnectDisplayText() {
	defer qt.Recovering("disconnect Delegate::displayText")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::displayText")
	}
}

func (ptr *Delegate) DisplayText(value core.QVariant_ITF, locale core.QLocale_ITF) string {
	defer qt.Recovering("Delegate::displayText")

	if ptr.Pointer() != nil {
		return C.GoString(C.Delegate_DisplayText(ptr.Pointer(), core.PointerFromQVariant(value), core.PointerFromQLocale(locale)))
	}
	return ""
}

func (ptr *Delegate) DisplayTextDefault(value core.QVariant_ITF, locale core.QLocale_ITF) string {
	defer qt.Recovering("Delegate::displayText")

	if ptr.Pointer() != nil {
		return C.GoString(C.Delegate_DisplayTextDefault(ptr.Pointer(), core.PointerFromQVariant(value), core.PointerFromQLocale(locale)))
	}
	return ""
}

//export callbackDelegate_EditorEvent
func callbackDelegate_EditorEvent(ptr unsafe.Pointer, event unsafe.Pointer, model unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) C.char {
	defer qt.Recovering("callback Delegate::editorEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "Delegate::editorEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent, *core.QAbstractItemModel, *widgets.QStyleOptionViewItem, *core.QModelIndex) bool)(core.NewQEventFromPointer(event), core.NewQAbstractItemModelFromPointer(model), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewDelegateFromPointer(ptr).EditorEventDefault(core.NewQEventFromPointer(event), core.NewQAbstractItemModelFromPointer(model), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index)))))
}

func (ptr *Delegate) ConnectEditorEvent(f func(event *core.QEvent, model *core.QAbstractItemModel, option *widgets.QStyleOptionViewItem, index *core.QModelIndex) bool) {
	defer qt.Recovering("connect Delegate::editorEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::editorEvent", f)
	}
}

func (ptr *Delegate) DisconnectEditorEvent() {
	defer qt.Recovering("disconnect Delegate::editorEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::editorEvent")
	}
}

func (ptr *Delegate) EditorEvent(event core.QEvent_ITF, model core.QAbstractItemModel_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) bool {
	defer qt.Recovering("Delegate::editorEvent")

	if ptr.Pointer() != nil {
		return C.Delegate_EditorEvent(ptr.Pointer(), core.PointerFromQEvent(event), core.PointerFromQAbstractItemModel(model), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *Delegate) EditorEventDefault(event core.QEvent_ITF, model core.QAbstractItemModel_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) bool {
	defer qt.Recovering("Delegate::editorEvent")

	if ptr.Pointer() != nil {
		return C.Delegate_EditorEventDefault(ptr.Pointer(), core.PointerFromQEvent(event), core.PointerFromQAbstractItemModel(model), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

//export callbackDelegate_InitStyleOption
func callbackDelegate_InitStyleOption(ptr unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback Delegate::initStyleOption")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "Delegate::initStyleOption"); signal != nil {
		signal.(func(*widgets.QStyleOptionViewItem, *core.QModelIndex))(widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index))
	} else {
		NewDelegateFromPointer(ptr).InitStyleOptionDefault(widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *Delegate) ConnectInitStyleOption(f func(option *widgets.QStyleOptionViewItem, index *core.QModelIndex)) {
	defer qt.Recovering("connect Delegate::initStyleOption")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::initStyleOption", f)
	}
}

func (ptr *Delegate) DisconnectInitStyleOption() {
	defer qt.Recovering("disconnect Delegate::initStyleOption")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::initStyleOption")
	}
}

func (ptr *Delegate) InitStyleOption(option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("Delegate::initStyleOption")

	if ptr.Pointer() != nil {
		C.Delegate_InitStyleOption(ptr.Pointer(), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))
	}
}

func (ptr *Delegate) InitStyleOptionDefault(option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("Delegate::initStyleOption")

	if ptr.Pointer() != nil {
		C.Delegate_InitStyleOptionDefault(ptr.Pointer(), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))
	}
}

//export callbackDelegate_Paint
func callbackDelegate_Paint(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback Delegate::paint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "Delegate::paint"); signal != nil {
		signal.(func(*gui.QPainter, *widgets.QStyleOptionViewItem, *core.QModelIndex))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index))
	} else {
		NewDelegateFromPointer(ptr).PaintDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *Delegate) ConnectPaint(f func(painter *gui.QPainter, option *widgets.QStyleOptionViewItem, index *core.QModelIndex)) {
	defer qt.Recovering("connect Delegate::paint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::paint", f)
	}
}

func (ptr *Delegate) DisconnectPaint() {
	defer qt.Recovering("disconnect Delegate::paint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::paint")
	}
}

func (ptr *Delegate) Paint(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("Delegate::paint")

	if ptr.Pointer() != nil {
		C.Delegate_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))
	}
}

func (ptr *Delegate) PaintDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("Delegate::paint")

	if ptr.Pointer() != nil {
		C.Delegate_PaintDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))
	}
}

//export callbackDelegate_SetEditorData
func callbackDelegate_SetEditorData(ptr unsafe.Pointer, editor unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback Delegate::setEditorData")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "Delegate::setEditorData"); signal != nil {
		signal.(func(*widgets.QWidget, *core.QModelIndex))(widgets.NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
	} else {
		NewDelegateFromPointer(ptr).SetEditorDataDefault(widgets.NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *Delegate) ConnectSetEditorData(f func(editor *widgets.QWidget, index *core.QModelIndex)) {
	defer qt.Recovering("connect Delegate::setEditorData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::setEditorData", f)
	}
}

func (ptr *Delegate) DisconnectSetEditorData() {
	defer qt.Recovering("disconnect Delegate::setEditorData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::setEditorData")
	}
}

func (ptr *Delegate) SetEditorData(editor widgets.QWidget_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("Delegate::setEditorData")

	if ptr.Pointer() != nil {
		C.Delegate_SetEditorData(ptr.Pointer(), widgets.PointerFromQWidget(editor), core.PointerFromQModelIndex(index))
	}
}

func (ptr *Delegate) SetEditorDataDefault(editor widgets.QWidget_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("Delegate::setEditorData")

	if ptr.Pointer() != nil {
		C.Delegate_SetEditorDataDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor), core.PointerFromQModelIndex(index))
	}
}

//export callbackDelegate_SetModelData
func callbackDelegate_SetModelData(ptr unsafe.Pointer, editor unsafe.Pointer, model unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback Delegate::setModelData")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "Delegate::setModelData"); signal != nil {
		signal.(func(*widgets.QWidget, *core.QAbstractItemModel, *core.QModelIndex))(widgets.NewQWidgetFromPointer(editor), core.NewQAbstractItemModelFromPointer(model), core.NewQModelIndexFromPointer(index))
	} else {
		NewDelegateFromPointer(ptr).SetModelDataDefault(widgets.NewQWidgetFromPointer(editor), core.NewQAbstractItemModelFromPointer(model), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *Delegate) ConnectSetModelData(f func(editor *widgets.QWidget, model *core.QAbstractItemModel, index *core.QModelIndex)) {
	defer qt.Recovering("connect Delegate::setModelData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::setModelData", f)
	}
}

func (ptr *Delegate) DisconnectSetModelData() {
	defer qt.Recovering("disconnect Delegate::setModelData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::setModelData")
	}
}

func (ptr *Delegate) SetModelData(editor widgets.QWidget_ITF, model core.QAbstractItemModel_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("Delegate::setModelData")

	if ptr.Pointer() != nil {
		C.Delegate_SetModelData(ptr.Pointer(), widgets.PointerFromQWidget(editor), core.PointerFromQAbstractItemModel(model), core.PointerFromQModelIndex(index))
	}
}

func (ptr *Delegate) SetModelDataDefault(editor widgets.QWidget_ITF, model core.QAbstractItemModel_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("Delegate::setModelData")

	if ptr.Pointer() != nil {
		C.Delegate_SetModelDataDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor), core.PointerFromQAbstractItemModel(model), core.PointerFromQModelIndex(index))
	}
}

//export callbackDelegate_SizeHint
func callbackDelegate_SizeHint(ptr unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback Delegate::sizeHint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "Delegate::sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func(*widgets.QStyleOptionViewItem, *core.QModelIndex) *core.QSize)(widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQSize(NewDelegateFromPointer(ptr).SizeHintDefault(widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index)))
}

func (ptr *Delegate) ConnectSizeHint(f func(option *widgets.QStyleOptionViewItem, index *core.QModelIndex) *core.QSize) {
	defer qt.Recovering("connect Delegate::sizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::sizeHint", f)
	}
}

func (ptr *Delegate) DisconnectSizeHint() {
	defer qt.Recovering("disconnect Delegate::sizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::sizeHint")
	}
}

func (ptr *Delegate) SizeHint(option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) *core.QSize {
	defer qt.Recovering("Delegate::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.Delegate_SizeHint(ptr.Pointer(), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *Delegate) SizeHintDefault(option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) *core.QSize {
	defer qt.Recovering("Delegate::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.Delegate_SizeHintDefault(ptr.Pointer(), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackDelegate_UpdateEditorGeometry
func callbackDelegate_UpdateEditorGeometry(ptr unsafe.Pointer, editor unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback Delegate::updateEditorGeometry")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "Delegate::updateEditorGeometry"); signal != nil {
		signal.(func(*widgets.QWidget, *widgets.QStyleOptionViewItem, *core.QModelIndex))(widgets.NewQWidgetFromPointer(editor), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index))
	} else {
		NewDelegateFromPointer(ptr).UpdateEditorGeometryDefault(widgets.NewQWidgetFromPointer(editor), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *Delegate) ConnectUpdateEditorGeometry(f func(editor *widgets.QWidget, option *widgets.QStyleOptionViewItem, index *core.QModelIndex)) {
	defer qt.Recovering("connect Delegate::updateEditorGeometry")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::updateEditorGeometry", f)
	}
}

func (ptr *Delegate) DisconnectUpdateEditorGeometry() {
	defer qt.Recovering("disconnect Delegate::updateEditorGeometry")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::updateEditorGeometry")
	}
}

func (ptr *Delegate) UpdateEditorGeometry(editor widgets.QWidget_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("Delegate::updateEditorGeometry")

	if ptr.Pointer() != nil {
		C.Delegate_UpdateEditorGeometry(ptr.Pointer(), widgets.PointerFromQWidget(editor), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))
	}
}

func (ptr *Delegate) UpdateEditorGeometryDefault(editor widgets.QWidget_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("Delegate::updateEditorGeometry")

	if ptr.Pointer() != nil {
		C.Delegate_UpdateEditorGeometryDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))
	}
}

//export callbackDelegate_DestroyEditor
func callbackDelegate_DestroyEditor(ptr unsafe.Pointer, editor unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback Delegate::destroyEditor")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "Delegate::destroyEditor"); signal != nil {
		signal.(func(*widgets.QWidget, *core.QModelIndex))(widgets.NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
	} else {
		NewDelegateFromPointer(ptr).DestroyEditorDefault(widgets.NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *Delegate) ConnectDestroyEditor(f func(editor *widgets.QWidget, index *core.QModelIndex)) {
	defer qt.Recovering("connect Delegate::destroyEditor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::destroyEditor", f)
	}
}

func (ptr *Delegate) DisconnectDestroyEditor() {
	defer qt.Recovering("disconnect Delegate::destroyEditor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::destroyEditor")
	}
}

func (ptr *Delegate) DestroyEditor(editor widgets.QWidget_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("Delegate::destroyEditor")

	if ptr.Pointer() != nil {
		C.Delegate_DestroyEditor(ptr.Pointer(), widgets.PointerFromQWidget(editor), core.PointerFromQModelIndex(index))
	}
}

func (ptr *Delegate) DestroyEditorDefault(editor widgets.QWidget_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("Delegate::destroyEditor")

	if ptr.Pointer() != nil {
		C.Delegate_DestroyEditorDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor), core.PointerFromQModelIndex(index))
	}
}

//export callbackDelegate_HelpEvent
func callbackDelegate_HelpEvent(ptr unsafe.Pointer, event unsafe.Pointer, view unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) C.char {
	defer qt.Recovering("callback Delegate::helpEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "Delegate::helpEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*gui.QHelpEvent, *widgets.QAbstractItemView, *widgets.QStyleOptionViewItem, *core.QModelIndex) bool)(gui.NewQHelpEventFromPointer(event), widgets.NewQAbstractItemViewFromPointer(view), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewDelegateFromPointer(ptr).HelpEventDefault(gui.NewQHelpEventFromPointer(event), widgets.NewQAbstractItemViewFromPointer(view), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index)))))
}

func (ptr *Delegate) ConnectHelpEvent(f func(event *gui.QHelpEvent, view *widgets.QAbstractItemView, option *widgets.QStyleOptionViewItem, index *core.QModelIndex) bool) {
	defer qt.Recovering("connect Delegate::helpEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::helpEvent", f)
	}
}

func (ptr *Delegate) DisconnectHelpEvent() {
	defer qt.Recovering("disconnect Delegate::helpEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::helpEvent")
	}
}

func (ptr *Delegate) HelpEvent(event gui.QHelpEvent_ITF, view widgets.QAbstractItemView_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) bool {
	defer qt.Recovering("Delegate::helpEvent")

	if ptr.Pointer() != nil {
		return C.Delegate_HelpEvent(ptr.Pointer(), gui.PointerFromQHelpEvent(event), widgets.PointerFromQAbstractItemView(view), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *Delegate) HelpEventDefault(event gui.QHelpEvent_ITF, view widgets.QAbstractItemView_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) bool {
	defer qt.Recovering("Delegate::helpEvent")

	if ptr.Pointer() != nil {
		return C.Delegate_HelpEventDefault(ptr.Pointer(), gui.PointerFromQHelpEvent(event), widgets.PointerFromQAbstractItemView(view), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

//export callbackDelegate_TimerEvent
func callbackDelegate_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback Delegate::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "Delegate::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewDelegateFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *Delegate) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect Delegate::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::timerEvent", f)
	}
}

func (ptr *Delegate) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect Delegate::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::timerEvent")
	}
}

func (ptr *Delegate) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("Delegate::timerEvent")

	if ptr.Pointer() != nil {
		C.Delegate_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *Delegate) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("Delegate::timerEvent")

	if ptr.Pointer() != nil {
		C.Delegate_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackDelegate_ChildEvent
func callbackDelegate_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback Delegate::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "Delegate::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewDelegateFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *Delegate) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect Delegate::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::childEvent", f)
	}
}

func (ptr *Delegate) DisconnectChildEvent() {
	defer qt.Recovering("disconnect Delegate::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::childEvent")
	}
}

func (ptr *Delegate) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("Delegate::childEvent")

	if ptr.Pointer() != nil {
		C.Delegate_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *Delegate) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("Delegate::childEvent")

	if ptr.Pointer() != nil {
		C.Delegate_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackDelegate_ConnectNotify
func callbackDelegate_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback Delegate::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "Delegate::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewDelegateFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Delegate) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect Delegate::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::connectNotify", f)
	}
}

func (ptr *Delegate) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect Delegate::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::connectNotify")
	}
}

func (ptr *Delegate) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("Delegate::connectNotify")

	if ptr.Pointer() != nil {
		C.Delegate_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *Delegate) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("Delegate::connectNotify")

	if ptr.Pointer() != nil {
		C.Delegate_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackDelegate_CustomEvent
func callbackDelegate_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback Delegate::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "Delegate::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewDelegateFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *Delegate) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect Delegate::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::customEvent", f)
	}
}

func (ptr *Delegate) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect Delegate::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::customEvent")
	}
}

func (ptr *Delegate) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("Delegate::customEvent")

	if ptr.Pointer() != nil {
		C.Delegate_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *Delegate) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("Delegate::customEvent")

	if ptr.Pointer() != nil {
		C.Delegate_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackDelegate_DeleteLater
func callbackDelegate_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback Delegate::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "Delegate::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewDelegateFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *Delegate) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect Delegate::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::deleteLater", f)
	}
}

func (ptr *Delegate) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect Delegate::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::deleteLater")
	}
}

func (ptr *Delegate) DeleteLater() {
	defer qt.Recovering("Delegate::deleteLater")

	if ptr.Pointer() != nil {
		C.Delegate_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *Delegate) DeleteLaterDefault() {
	defer qt.Recovering("Delegate::deleteLater")

	if ptr.Pointer() != nil {
		C.Delegate_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackDelegate_DisconnectNotify
func callbackDelegate_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback Delegate::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "Delegate::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewDelegateFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Delegate) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect Delegate::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::disconnectNotify", f)
	}
}

func (ptr *Delegate) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect Delegate::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::disconnectNotify")
	}
}

func (ptr *Delegate) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("Delegate::disconnectNotify")

	if ptr.Pointer() != nil {
		C.Delegate_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *Delegate) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("Delegate::disconnectNotify")

	if ptr.Pointer() != nil {
		C.Delegate_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackDelegate_Event
func callbackDelegate_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback Delegate::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "Delegate::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewDelegateFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *Delegate) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect Delegate::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::event", f)
	}
}

func (ptr *Delegate) DisconnectEvent() {
	defer qt.Recovering("disconnect Delegate::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "Delegate::event")
	}
}

func (ptr *Delegate) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("Delegate::event")

	if ptr.Pointer() != nil {
		return C.Delegate_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *Delegate) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("Delegate::event")

	if ptr.Pointer() != nil {
		return C.Delegate_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackDelegate_MetaObject
