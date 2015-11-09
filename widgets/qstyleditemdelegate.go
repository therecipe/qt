package widgets

//#include "qstyleditemdelegate.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
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
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QStyledItemDelegate_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QStyledItemDelegate) QStyledItemDelegate_PTR() *QStyledItemDelegate {
	return ptr
}

func NewQStyledItemDelegate(parent core.QObject_ITF) *QStyledItemDelegate {
	return NewQStyledItemDelegateFromPointer(C.QStyledItemDelegate_NewQStyledItemDelegate(core.PointerFromQObject(parent)))
}

func (ptr *QStyledItemDelegate) CreateEditor(parent QWidget_ITF, option QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QStyledItemDelegate_CreateEditor(ptr.Pointer(), PointerFromQWidget(parent), PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QStyledItemDelegate) DisplayText(value core.QVariant_ITF, locale core.QLocale_ITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStyledItemDelegate_DisplayText(ptr.Pointer(), core.PointerFromQVariant(value), core.PointerFromQLocale(locale)))
	}
	return ""
}

func (ptr *QStyledItemDelegate) ItemEditorFactory() *QItemEditorFactory {
	if ptr.Pointer() != nil {
		return NewQItemEditorFactoryFromPointer(C.QStyledItemDelegate_ItemEditorFactory(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStyledItemDelegate) Paint(painter gui.QPainter_ITF, option QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QStyledItemDelegate) SetEditorData(editor QWidget_ITF, index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_SetEditorData(ptr.Pointer(), PointerFromQWidget(editor), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QStyledItemDelegate) SetItemEditorFactory(factory QItemEditorFactory_ITF) {
	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_SetItemEditorFactory(ptr.Pointer(), PointerFromQItemEditorFactory(factory))
	}
}

func (ptr *QStyledItemDelegate) SetModelData(editor QWidget_ITF, model core.QAbstractItemModel_ITF, index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_SetModelData(ptr.Pointer(), PointerFromQWidget(editor), core.PointerFromQAbstractItemModel(model), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QStyledItemDelegate) UpdateEditorGeometry(editor QWidget_ITF, option QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_UpdateEditorGeometry(ptr.Pointer(), PointerFromQWidget(editor), PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QStyledItemDelegate) DestroyQStyledItemDelegate() {
	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_DestroyQStyledItemDelegate(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
