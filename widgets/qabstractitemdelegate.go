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
func callbackQAbstractItemDelegateCloseEditor(ptrName *C.char, editor unsafe.Pointer, hint C.int) {
	defer qt.Recovering("callback QAbstractItemDelegate::closeEditor")

	var signal = qt.GetSignal(C.GoString(ptrName), "closeEditor")
	if signal != nil {
		signal.(func(*QWidget, QAbstractItemDelegate__EndEditHint))(NewQWidgetFromPointer(editor), QAbstractItemDelegate__EndEditHint(hint))
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
func callbackQAbstractItemDelegateCommitData(ptrName *C.char, editor unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractItemDelegate::commitData")

	var signal = qt.GetSignal(C.GoString(ptrName), "commitData")
	if signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(editor))
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
func callbackQAbstractItemDelegateDestroyEditor(ptrName *C.char, editor unsafe.Pointer, index unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemDelegate::destroyEditor")

	var signal = qt.GetSignal(C.GoString(ptrName), "destroyEditor")
	if signal != nil {
		defer signal.(func(*QWidget, *core.QModelIndex))(NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
		return true
	}
	return false

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
func callbackQAbstractItemDelegateSetEditorData(ptrName *C.char, editor unsafe.Pointer, index unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemDelegate::setEditorData")

	var signal = qt.GetSignal(C.GoString(ptrName), "setEditorData")
	if signal != nil {
		defer signal.(func(*QWidget, *core.QModelIndex))(NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
		return true
	}
	return false

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
func callbackQAbstractItemDelegateSetModelData(ptrName *C.char, editor unsafe.Pointer, model unsafe.Pointer, index unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractItemDelegate::setModelData")

	var signal = qt.GetSignal(C.GoString(ptrName), "setModelData")
	if signal != nil {
		defer signal.(func(*QWidget, *core.QAbstractItemModel, *core.QModelIndex))(NewQWidgetFromPointer(editor), core.NewQAbstractItemModelFromPointer(model), core.NewQModelIndexFromPointer(index))
		return true
	}
	return false

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
func callbackQAbstractItemDelegateSizeHintChanged(ptrName *C.char, index unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractItemDelegate::sizeHintChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "sizeHintChanged")
	if signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	}

}

func (ptr *QAbstractItemDelegate) DestroyQAbstractItemDelegate() {
	defer qt.Recovering("QAbstractItemDelegate::~QAbstractItemDelegate")

	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_DestroyQAbstractItemDelegate(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
