package widgets

//#include "qitemdelegate.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QItemDelegate struct {
	QAbstractItemDelegate
}

type QItemDelegateITF interface {
	QAbstractItemDelegateITF
	QItemDelegatePTR() *QItemDelegate
}

func PointerFromQItemDelegate(ptr QItemDelegateITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QItemDelegatePTR().Pointer()
	}
	return nil
}

func QItemDelegateFromPointer(ptr unsafe.Pointer) *QItemDelegate {
	var n = new(QItemDelegate)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QItemDelegate_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QItemDelegate) QItemDelegatePTR() *QItemDelegate {
	return ptr
}

func (ptr *QItemDelegate) HasClipping() bool {
	if ptr.Pointer() != nil {
		return C.QItemDelegate_HasClipping(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QItemDelegate) SetClipping(clip bool) {
	if ptr.Pointer() != nil {
		C.QItemDelegate_SetClipping(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(clip)))
	}
}

func NewQItemDelegate(parent core.QObjectITF) *QItemDelegate {
	return QItemDelegateFromPointer(unsafe.Pointer(C.QItemDelegate_NewQItemDelegate(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QItemDelegate) CreateEditor(parent QWidgetITF, option QStyleOptionViewItemITF, index core.QModelIndexITF) *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QItemDelegate_CreateEditor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(parent)), C.QtObjectPtr(PointerFromQStyleOptionViewItem(option)), C.QtObjectPtr(core.PointerFromQModelIndex(index)))))
	}
	return nil
}

func (ptr *QItemDelegate) ItemEditorFactory() *QItemEditorFactory {
	if ptr.Pointer() != nil {
		return QItemEditorFactoryFromPointer(unsafe.Pointer(C.QItemDelegate_ItemEditorFactory(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QItemDelegate) Paint(painter gui.QPainterITF, option QStyleOptionViewItemITF, index core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QItemDelegate_Paint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainter(painter)), C.QtObjectPtr(PointerFromQStyleOptionViewItem(option)), C.QtObjectPtr(core.PointerFromQModelIndex(index)))
	}
}

func (ptr *QItemDelegate) SetEditorData(editor QWidgetITF, index core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QItemDelegate_SetEditorData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(editor)), C.QtObjectPtr(core.PointerFromQModelIndex(index)))
	}
}

func (ptr *QItemDelegate) SetItemEditorFactory(factory QItemEditorFactoryITF) {
	if ptr.Pointer() != nil {
		C.QItemDelegate_SetItemEditorFactory(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQItemEditorFactory(factory)))
	}
}

func (ptr *QItemDelegate) SetModelData(editor QWidgetITF, model core.QAbstractItemModelITF, index core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QItemDelegate_SetModelData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(editor)), C.QtObjectPtr(core.PointerFromQAbstractItemModel(model)), C.QtObjectPtr(core.PointerFromQModelIndex(index)))
	}
}

func (ptr *QItemDelegate) UpdateEditorGeometry(editor QWidgetITF, option QStyleOptionViewItemITF, index core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QItemDelegate_UpdateEditorGeometry(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(editor)), C.QtObjectPtr(PointerFromQStyleOptionViewItem(option)), C.QtObjectPtr(core.PointerFromQModelIndex(index)))
	}
}

func (ptr *QItemDelegate) DestroyQItemDelegate() {
	if ptr.Pointer() != nil {
		C.QItemDelegate_DestroyQItemDelegate(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
