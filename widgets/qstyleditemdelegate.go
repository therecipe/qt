package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"log"
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
		n.SetObjectName("QStyledItemDelegate_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QStyledItemDelegate) QStyledItemDelegate_PTR() *QStyledItemDelegate {
	return ptr
}

func NewQStyledItemDelegate(parent core.QObject_ITF) *QStyledItemDelegate {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyledItemDelegate::QStyledItemDelegate")
		}
	}()

	return NewQStyledItemDelegateFromPointer(C.QStyledItemDelegate_NewQStyledItemDelegate(core.PointerFromQObject(parent)))
}

func (ptr *QStyledItemDelegate) CreateEditor(parent QWidget_ITF, option QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyledItemDelegate::createEditor")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QStyledItemDelegate_CreateEditor(ptr.Pointer(), PointerFromQWidget(parent), PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QStyledItemDelegate) DisplayText(value core.QVariant_ITF, locale core.QLocale_ITF) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyledItemDelegate::displayText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QStyledItemDelegate_DisplayText(ptr.Pointer(), core.PointerFromQVariant(value), core.PointerFromQLocale(locale)))
	}
	return ""
}

func (ptr *QStyledItemDelegate) ItemEditorFactory() *QItemEditorFactory {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyledItemDelegate::itemEditorFactory")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQItemEditorFactoryFromPointer(C.QStyledItemDelegate_ItemEditorFactory(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStyledItemDelegate) Paint(painter gui.QPainter_ITF, option QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyledItemDelegate::paint")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QStyledItemDelegate) SetEditorData(editor QWidget_ITF, index core.QModelIndex_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyledItemDelegate::setEditorData")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_SetEditorData(ptr.Pointer(), PointerFromQWidget(editor), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QStyledItemDelegate) SetItemEditorFactory(factory QItemEditorFactory_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyledItemDelegate::setItemEditorFactory")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_SetItemEditorFactory(ptr.Pointer(), PointerFromQItemEditorFactory(factory))
	}
}

func (ptr *QStyledItemDelegate) SetModelData(editor QWidget_ITF, model core.QAbstractItemModel_ITF, index core.QModelIndex_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyledItemDelegate::setModelData")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_SetModelData(ptr.Pointer(), PointerFromQWidget(editor), core.PointerFromQAbstractItemModel(model), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QStyledItemDelegate) UpdateEditorGeometry(editor QWidget_ITF, option QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyledItemDelegate::updateEditorGeometry")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_UpdateEditorGeometry(ptr.Pointer(), PointerFromQWidget(editor), PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QStyledItemDelegate) DestroyQStyledItemDelegate() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyledItemDelegate::~QStyledItemDelegate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStyledItemDelegate_DestroyQStyledItemDelegate(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
