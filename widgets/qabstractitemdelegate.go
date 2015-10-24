package widgets

//#include "qabstractitemdelegate.h"
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

type QAbstractItemDelegateITF interface {
	core.QObjectITF
	QAbstractItemDelegatePTR() *QAbstractItemDelegate
}

func PointerFromQAbstractItemDelegate(ptr QAbstractItemDelegateITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractItemDelegatePTR().Pointer()
	}
	return nil
}

func QAbstractItemDelegateFromPointer(ptr unsafe.Pointer) *QAbstractItemDelegate {
	var n = new(QAbstractItemDelegate)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAbstractItemDelegate_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractItemDelegate) QAbstractItemDelegatePTR() *QAbstractItemDelegate {
	return ptr
}

//QAbstractItemDelegate::EndEditHint
type QAbstractItemDelegate__EndEditHint int

var (
	QAbstractItemDelegate__NoHint           = QAbstractItemDelegate__EndEditHint(0)
	QAbstractItemDelegate__EditNextItem     = QAbstractItemDelegate__EndEditHint(1)
	QAbstractItemDelegate__EditPreviousItem = QAbstractItemDelegate__EndEditHint(2)
	QAbstractItemDelegate__SubmitModelCache = QAbstractItemDelegate__EndEditHint(3)
	QAbstractItemDelegate__RevertModelCache = QAbstractItemDelegate__EndEditHint(4)
)

func (ptr *QAbstractItemDelegate) ConnectCloseEditor(f func(editor QWidgetITF, hint QAbstractItemDelegate__EndEditHint)) {
	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_ConnectCloseEditor(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "closeEditor", f)
	}
}

func (ptr *QAbstractItemDelegate) DisconnectCloseEditor() {
	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_DisconnectCloseEditor(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "closeEditor")
	}
}

//export callbackQAbstractItemDelegateCloseEditor
func callbackQAbstractItemDelegateCloseEditor(ptrName *C.char, editor unsafe.Pointer, hint C.int) {
	qt.GetSignal(C.GoString(ptrName), "closeEditor").(func(*QWidget, QAbstractItemDelegate__EndEditHint))(QWidgetFromPointer(editor), QAbstractItemDelegate__EndEditHint(hint))
}

func (ptr *QAbstractItemDelegate) ConnectCommitData(f func(editor QWidgetITF)) {
	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_ConnectCommitData(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "commitData", f)
	}
}

func (ptr *QAbstractItemDelegate) DisconnectCommitData() {
	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_DisconnectCommitData(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "commitData")
	}
}

//export callbackQAbstractItemDelegateCommitData
func callbackQAbstractItemDelegateCommitData(ptrName *C.char, editor unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "commitData").(func(*QWidget))(QWidgetFromPointer(editor))
}

func (ptr *QAbstractItemDelegate) CreateEditor(parent QWidgetITF, option QStyleOptionViewItemITF, index core.QModelIndexITF) *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QAbstractItemDelegate_CreateEditor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(parent)), C.QtObjectPtr(PointerFromQStyleOptionViewItem(option)), C.QtObjectPtr(core.PointerFromQModelIndex(index)))))
	}
	return nil
}

func (ptr *QAbstractItemDelegate) DestroyEditor(editor QWidgetITF, index core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_DestroyEditor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(editor)), C.QtObjectPtr(core.PointerFromQModelIndex(index)))
	}
}

func (ptr *QAbstractItemDelegate) EditorEvent(event core.QEventITF, model core.QAbstractItemModelITF, option QStyleOptionViewItemITF, index core.QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemDelegate_EditorEvent(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQEvent(event)), C.QtObjectPtr(core.PointerFromQAbstractItemModel(model)), C.QtObjectPtr(PointerFromQStyleOptionViewItem(option)), C.QtObjectPtr(core.PointerFromQModelIndex(index))) != 0
	}
	return false
}

func (ptr *QAbstractItemDelegate) HelpEvent(event gui.QHelpEventITF, view QAbstractItemViewITF, option QStyleOptionViewItemITF, index core.QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemDelegate_HelpEvent(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQHelpEvent(event)), C.QtObjectPtr(PointerFromQAbstractItemView(view)), C.QtObjectPtr(PointerFromQStyleOptionViewItem(option)), C.QtObjectPtr(core.PointerFromQModelIndex(index))) != 0
	}
	return false
}

func (ptr *QAbstractItemDelegate) Paint(painter gui.QPainterITF, option QStyleOptionViewItemITF, index core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_Paint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainter(painter)), C.QtObjectPtr(PointerFromQStyleOptionViewItem(option)), C.QtObjectPtr(core.PointerFromQModelIndex(index)))
	}
}

func (ptr *QAbstractItemDelegate) SetEditorData(editor QWidgetITF, index core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_SetEditorData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(editor)), C.QtObjectPtr(core.PointerFromQModelIndex(index)))
	}
}

func (ptr *QAbstractItemDelegate) SetModelData(editor QWidgetITF, model core.QAbstractItemModelITF, index core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_SetModelData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(editor)), C.QtObjectPtr(core.PointerFromQAbstractItemModel(model)), C.QtObjectPtr(core.PointerFromQModelIndex(index)))
	}
}

func (ptr *QAbstractItemDelegate) ConnectSizeHintChanged(f func(index core.QModelIndexITF)) {
	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_ConnectSizeHintChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "sizeHintChanged", f)
	}
}

func (ptr *QAbstractItemDelegate) DisconnectSizeHintChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_DisconnectSizeHintChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "sizeHintChanged")
	}
}

//export callbackQAbstractItemDelegateSizeHintChanged
func callbackQAbstractItemDelegateSizeHintChanged(ptrName *C.char, index unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "sizeHintChanged").(func(*core.QModelIndex))(core.QModelIndexFromPointer(index))
}

func (ptr *QAbstractItemDelegate) UpdateEditorGeometry(editor QWidgetITF, option QStyleOptionViewItemITF, index core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_UpdateEditorGeometry(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(editor)), C.QtObjectPtr(PointerFromQStyleOptionViewItem(option)), C.QtObjectPtr(core.PointerFromQModelIndex(index)))
	}
}

func (ptr *QAbstractItemDelegate) DestroyQAbstractItemDelegate() {
	if ptr.Pointer() != nil {
		C.QAbstractItemDelegate_DestroyQAbstractItemDelegate(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
