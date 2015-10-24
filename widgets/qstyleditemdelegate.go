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

type QStyledItemDelegateITF interface {
	QAbstractItemDelegateITF
	QStyledItemDelegatePTR() *QStyledItemDelegate
}

func PointerFromQStyledItemDelegate(ptr QStyledItemDelegateITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyledItemDelegatePTR().Pointer()
	}
	return nil
}

func QStyledItemDelegateFromPointer(ptr unsafe.Pointer) *QStyledItemDelegate {
	var n = new(QStyledItemDelegate)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QStyledItemDelegate_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QStyledItemDelegate) QStyledItemDelegatePTR() *QStyledItemDelegate {
	return ptr
}

func NewQStyledItemDelegate(parent core.QObjectITF) *QStyledItemDelegate {
	return QStyledItemDelegateFromPointer(unsafe.Pointer(C.QStyledItemDelegate_NewQStyledItemDelegate(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QStyledItemDelegate) CreateEditor(parent QWidgetITF, option QStyleOptionViewItemITF, index core.QModelIndexITF) *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QStyledItemDelegate_CreateEditor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(parent)), C.QtObjectPtr(PointerFromQStyleOptionViewItem(option)), C.QtObjectPtr(core.PointerFromQModelIndex(index)))))
	}
	return nil
}

func (ptr *QStyledItemDelegate) DisplayText(value string, locale core.QLocaleITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStyledItemDelegate_DisplayText(C.QtObjectPtr(ptr.Pointer()), C.CString(value), C.QtObjectPtr(core.PointerFromQLocale(locale))))
	}
	return ""
}

func (ptr *QStyledItemDelegate) ItemEditorFactory() *QItemEditorFactory {
	if ptr.Pointer() != nil {
		return QItemEditorFactoryFromPointer(unsafe.Pointer(C.QStyledItemDelegate_ItemEditorFactory(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QStyledItemDelegate) Paint(painter gui.QPainterITF, option QStyleOptionViewItemITF, index core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_Paint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainter(painter)), C.QtObjectPtr(PointerFromQStyleOptionViewItem(option)), C.QtObjectPtr(core.PointerFromQModelIndex(index)))
	}
}

func (ptr *QStyledItemDelegate) SetEditorData(editor QWidgetITF, index core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_SetEditorData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(editor)), C.QtObjectPtr(core.PointerFromQModelIndex(index)))
	}
}

func (ptr *QStyledItemDelegate) SetItemEditorFactory(factory QItemEditorFactoryITF) {
	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_SetItemEditorFactory(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQItemEditorFactory(factory)))
	}
}

func (ptr *QStyledItemDelegate) SetModelData(editor QWidgetITF, model core.QAbstractItemModelITF, index core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_SetModelData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(editor)), C.QtObjectPtr(core.PointerFromQAbstractItemModel(model)), C.QtObjectPtr(core.PointerFromQModelIndex(index)))
	}
}

func (ptr *QStyledItemDelegate) UpdateEditorGeometry(editor QWidgetITF, option QStyleOptionViewItemITF, index core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_UpdateEditorGeometry(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(editor)), C.QtObjectPtr(PointerFromQStyleOptionViewItem(option)), C.QtObjectPtr(core.PointerFromQModelIndex(index)))
	}
}

func (ptr *QStyledItemDelegate) DestroyQStyledItemDelegate() {
	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_DestroyQStyledItemDelegate(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
