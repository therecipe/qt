package dbus

//#include "dbus.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDBusArgument struct {
	ptr unsafe.Pointer
}

type QDBusArgument_ITF interface {
	QDBusArgument_PTR() *QDBusArgument
}

func (p *QDBusArgument) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusArgument) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusArgument(ptr QDBusArgument_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusArgument_PTR().Pointer()
	}
	return nil
}

func NewQDBusArgumentFromPointer(ptr unsafe.Pointer) *QDBusArgument {
	var n = new(QDBusArgument)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusArgument) QDBusArgument_PTR() *QDBusArgument {
	return ptr
}

//QDBusArgument::ElementType
type QDBusArgument__ElementType int64

const (
	QDBusArgument__BasicType     = QDBusArgument__ElementType(0)
	QDBusArgument__VariantType   = QDBusArgument__ElementType(1)
	QDBusArgument__ArrayType     = QDBusArgument__ElementType(2)
	QDBusArgument__StructureType = QDBusArgument__ElementType(3)
	QDBusArgument__MapType       = QDBusArgument__ElementType(4)
	QDBusArgument__MapEntryType  = QDBusArgument__ElementType(5)
	QDBusArgument__UnknownType   = QDBusArgument__ElementType(-1)
)

func NewQDBusArgument() *QDBusArgument {
	defer qt.Recovering("QDBusArgument::QDBusArgument")

	return NewQDBusArgumentFromPointer(C.QDBusArgument_NewQDBusArgument())
}

func NewQDBusArgument2(other QDBusArgument_ITF) *QDBusArgument {
	defer qt.Recovering("QDBusArgument::QDBusArgument")

	return NewQDBusArgumentFromPointer(C.QDBusArgument_NewQDBusArgument2(PointerFromQDBusArgument(other)))
}

func (ptr *QDBusArgument) AsVariant() *core.QVariant {
	defer qt.Recovering("QDBusArgument::asVariant")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QDBusArgument_AsVariant(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusArgument) AtEnd() bool {
	defer qt.Recovering("QDBusArgument::atEnd")

	if ptr.Pointer() != nil {
		return C.QDBusArgument_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusArgument) BeginArray(id int) {
	defer qt.Recovering("QDBusArgument::beginArray")

	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginArray(ptr.Pointer(), C.int(id))
	}
}

func (ptr *QDBusArgument) BeginArray2() {
	defer qt.Recovering("QDBusArgument::beginArray")

	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginArray2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) BeginMap(kid int, vid int) {
	defer qt.Recovering("QDBusArgument::beginMap")

	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginMap(ptr.Pointer(), C.int(kid), C.int(vid))
	}
}

func (ptr *QDBusArgument) BeginMap2() {
	defer qt.Recovering("QDBusArgument::beginMap")

	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginMap2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) BeginMapEntry() {
	defer qt.Recovering("QDBusArgument::beginMapEntry")

	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginMapEntry(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) BeginMapEntry2() {
	defer qt.Recovering("QDBusArgument::beginMapEntry")

	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginMapEntry2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) BeginStructure() {
	defer qt.Recovering("QDBusArgument::beginStructure")

	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginStructure(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) BeginStructure2() {
	defer qt.Recovering("QDBusArgument::beginStructure")

	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginStructure2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) CurrentType() QDBusArgument__ElementType {
	defer qt.Recovering("QDBusArgument::currentType")

	if ptr.Pointer() != nil {
		return QDBusArgument__ElementType(C.QDBusArgument_CurrentType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDBusArgument) EndArray() {
	defer qt.Recovering("QDBusArgument::endArray")

	if ptr.Pointer() != nil {
		C.QDBusArgument_EndArray(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndArray2() {
	defer qt.Recovering("QDBusArgument::endArray")

	if ptr.Pointer() != nil {
		C.QDBusArgument_EndArray2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndMap() {
	defer qt.Recovering("QDBusArgument::endMap")

	if ptr.Pointer() != nil {
		C.QDBusArgument_EndMap(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndMap2() {
	defer qt.Recovering("QDBusArgument::endMap")

	if ptr.Pointer() != nil {
		C.QDBusArgument_EndMap2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndMapEntry() {
	defer qt.Recovering("QDBusArgument::endMapEntry")

	if ptr.Pointer() != nil {
		C.QDBusArgument_EndMapEntry(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndMapEntry2() {
	defer qt.Recovering("QDBusArgument::endMapEntry")

	if ptr.Pointer() != nil {
		C.QDBusArgument_EndMapEntry2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndStructure() {
	defer qt.Recovering("QDBusArgument::endStructure")

	if ptr.Pointer() != nil {
		C.QDBusArgument_EndStructure(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndStructure2() {
	defer qt.Recovering("QDBusArgument::endStructure")

	if ptr.Pointer() != nil {
		C.QDBusArgument_EndStructure2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) DestroyQDBusArgument() {
	defer qt.Recovering("QDBusArgument::~QDBusArgument")

	if ptr.Pointer() != nil {
		C.QDBusArgument_DestroyQDBusArgument(ptr.Pointer())
	}
}
