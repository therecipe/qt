package widgets

//#include "qtreewidgetitemiterator.h"
import "C"
import (
	"unsafe"
)

type QTreeWidgetItemIterator struct {
	ptr unsafe.Pointer
}

type QTreeWidgetItemIteratorITF interface {
	QTreeWidgetItemIteratorPTR() *QTreeWidgetItemIterator
}

func (p *QTreeWidgetItemIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTreeWidgetItemIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTreeWidgetItemIterator(ptr QTreeWidgetItemIteratorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTreeWidgetItemIteratorPTR().Pointer()
	}
	return nil
}

func QTreeWidgetItemIteratorFromPointer(ptr unsafe.Pointer) *QTreeWidgetItemIterator {
	var n = new(QTreeWidgetItemIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTreeWidgetItemIterator) QTreeWidgetItemIteratorPTR() *QTreeWidgetItemIterator {
	return ptr
}

//QTreeWidgetItemIterator::IteratorFlag
type QTreeWidgetItemIterator__IteratorFlag int

var (
	QTreeWidgetItemIterator__All           = QTreeWidgetItemIterator__IteratorFlag(0x00000000)
	QTreeWidgetItemIterator__Hidden        = QTreeWidgetItemIterator__IteratorFlag(0x00000001)
	QTreeWidgetItemIterator__NotHidden     = QTreeWidgetItemIterator__IteratorFlag(0x00000002)
	QTreeWidgetItemIterator__Selected      = QTreeWidgetItemIterator__IteratorFlag(0x00000004)
	QTreeWidgetItemIterator__Unselected    = QTreeWidgetItemIterator__IteratorFlag(0x00000008)
	QTreeWidgetItemIterator__Selectable    = QTreeWidgetItemIterator__IteratorFlag(0x00000010)
	QTreeWidgetItemIterator__NotSelectable = QTreeWidgetItemIterator__IteratorFlag(0x00000020)
	QTreeWidgetItemIterator__DragEnabled   = QTreeWidgetItemIterator__IteratorFlag(0x00000040)
	QTreeWidgetItemIterator__DragDisabled  = QTreeWidgetItemIterator__IteratorFlag(0x00000080)
	QTreeWidgetItemIterator__DropEnabled   = QTreeWidgetItemIterator__IteratorFlag(0x00000100)
	QTreeWidgetItemIterator__DropDisabled  = QTreeWidgetItemIterator__IteratorFlag(0x00000200)
	QTreeWidgetItemIterator__HasChildren   = QTreeWidgetItemIterator__IteratorFlag(0x00000400)
	QTreeWidgetItemIterator__NoChildren    = QTreeWidgetItemIterator__IteratorFlag(0x00000800)
	QTreeWidgetItemIterator__Checked       = QTreeWidgetItemIterator__IteratorFlag(0x00001000)
	QTreeWidgetItemIterator__NotChecked    = QTreeWidgetItemIterator__IteratorFlag(0x00002000)
	QTreeWidgetItemIterator__Enabled       = QTreeWidgetItemIterator__IteratorFlag(0x00004000)
	QTreeWidgetItemIterator__Disabled      = QTreeWidgetItemIterator__IteratorFlag(0x00008000)
	QTreeWidgetItemIterator__Editable      = QTreeWidgetItemIterator__IteratorFlag(0x00010000)
	QTreeWidgetItemIterator__NotEditable   = QTreeWidgetItemIterator__IteratorFlag(0x00020000)
	QTreeWidgetItemIterator__UserFlag      = QTreeWidgetItemIterator__IteratorFlag(0x01000000)
)
