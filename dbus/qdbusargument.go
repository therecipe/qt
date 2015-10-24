package dbus

//#include "qdbusargument.h"
import "C"
import (
	"unsafe"
)

type QDBusArgument struct {
	ptr unsafe.Pointer
}

type QDBusArgumentITF interface {
	QDBusArgumentPTR() *QDBusArgument
}

func (p *QDBusArgument) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusArgument) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusArgument(ptr QDBusArgumentITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusArgumentPTR().Pointer()
	}
	return nil
}

func QDBusArgumentFromPointer(ptr unsafe.Pointer) *QDBusArgument {
	var n = new(QDBusArgument)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusArgument) QDBusArgumentPTR() *QDBusArgument {
	return ptr
}

//QDBusArgument::ElementType
type QDBusArgument__ElementType int

var (
	QDBusArgument__BasicType     = QDBusArgument__ElementType(0)
	QDBusArgument__VariantType   = QDBusArgument__ElementType(1)
	QDBusArgument__ArrayType     = QDBusArgument__ElementType(2)
	QDBusArgument__StructureType = QDBusArgument__ElementType(3)
	QDBusArgument__MapType       = QDBusArgument__ElementType(4)
	QDBusArgument__MapEntryType  = QDBusArgument__ElementType(5)
	QDBusArgument__UnknownType   = QDBusArgument__ElementType(-1)
)

func NewQDBusArgument() *QDBusArgument {
	return QDBusArgumentFromPointer(unsafe.Pointer(C.QDBusArgument_NewQDBusArgument()))
}

func NewQDBusArgument2(other QDBusArgumentITF) *QDBusArgument {
	return QDBusArgumentFromPointer(unsafe.Pointer(C.QDBusArgument_NewQDBusArgument2(C.QtObjectPtr(PointerFromQDBusArgument(other)))))
}

func (ptr *QDBusArgument) AsVariant() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusArgument_AsVariant(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDBusArgument) AtEnd() bool {
	if ptr.Pointer() != nil {
		return C.QDBusArgument_AtEnd(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDBusArgument) BeginArray(id int) {
	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginArray(C.QtObjectPtr(ptr.Pointer()), C.int(id))
	}
}

func (ptr *QDBusArgument) BeginArray2() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginArray2(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDBusArgument) BeginMap(kid int, vid int) {
	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginMap(C.QtObjectPtr(ptr.Pointer()), C.int(kid), C.int(vid))
	}
}

func (ptr *QDBusArgument) BeginMap2() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginMap2(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDBusArgument) BeginMapEntry() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginMapEntry(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDBusArgument) BeginMapEntry2() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginMapEntry2(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDBusArgument) BeginStructure() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginStructure(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDBusArgument) BeginStructure2() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginStructure2(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDBusArgument) CurrentType() QDBusArgument__ElementType {
	if ptr.Pointer() != nil {
		return QDBusArgument__ElementType(C.QDBusArgument_CurrentType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDBusArgument) EndArray() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_EndArray(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDBusArgument) EndArray2() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_EndArray2(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDBusArgument) EndMap() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_EndMap(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDBusArgument) EndMap2() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_EndMap2(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDBusArgument) EndMapEntry() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_EndMapEntry(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDBusArgument) EndMapEntry2() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_EndMapEntry2(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDBusArgument) EndStructure() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_EndStructure(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDBusArgument) EndStructure2() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_EndStructure2(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDBusArgument) DestroyQDBusArgument() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_DestroyQDBusArgument(C.QtObjectPtr(ptr.Pointer()))
	}
}
