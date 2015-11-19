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
	if n.ObjectName() == "" {
		n.SetObjectName("QItemDelegate_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QItemDelegate) QItemDelegate_PTR() *QItemDelegate {
	return ptr
}

func (ptr *QItemDelegate) HasClipping() bool {
	if ptr.Pointer() != nil {
		return C.QItemDelegate_HasClipping(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QItemDelegate) SetClipping(clip bool) {
	if ptr.Pointer() != nil {
		C.QItemDelegate_SetClipping(ptr.Pointer(), C.int(qt.GoBoolToInt(clip)))
	}
}

func NewQItemDelegate(parent core.QObject_ITF) *QItemDelegate {
	return NewQItemDelegateFromPointer(C.QItemDelegate_NewQItemDelegate(core.PointerFromQObject(parent)))
}

func (ptr *QItemDelegate) CreateEditor(parent QWidget_ITF, option QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QItemDelegate_CreateEditor(ptr.Pointer(), PointerFromQWidget(parent), PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QItemDelegate) ItemEditorFactory() *QItemEditorFactory {
	if ptr.Pointer() != nil {
		return NewQItemEditorFactoryFromPointer(C.QItemDelegate_ItemEditorFactory(ptr.Pointer()))
	}
	return nil
}

func (ptr *QItemDelegate) Paint(painter gui.QPainter_ITF, option QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QItemDelegate_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QItemDelegate) SetEditorData(editor QWidget_ITF, index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QItemDelegate_SetEditorData(ptr.Pointer(), PointerFromQWidget(editor), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QItemDelegate) SetItemEditorFactory(factory QItemEditorFactory_ITF) {
	if ptr.Pointer() != nil {
		C.QItemDelegate_SetItemEditorFactory(ptr.Pointer(), PointerFromQItemEditorFactory(factory))
	}
}

func (ptr *QItemDelegate) SetModelData(editor QWidget_ITF, model core.QAbstractItemModel_ITF, index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QItemDelegate_SetModelData(ptr.Pointer(), PointerFromQWidget(editor), core.PointerFromQAbstractItemModel(model), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QItemDelegate) UpdateEditorGeometry(editor QWidget_ITF, option QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QItemDelegate_UpdateEditorGeometry(ptr.Pointer(), PointerFromQWidget(editor), PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QItemDelegate) DestroyQItemDelegate() {
	if ptr.Pointer() != nil {
		C.QItemDelegate_DestroyQItemDelegate(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
