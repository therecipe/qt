package core

//#include "qsignalmapper.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSignalMapper struct {
	QObject
}

type QSignalMapper_ITF interface {
	QObject_ITF
	QSignalMapper_PTR() *QSignalMapper
}

func PointerFromQSignalMapper(ptr QSignalMapper_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSignalMapper_PTR().Pointer()
	}
	return nil
}

func NewQSignalMapperFromPointer(ptr unsafe.Pointer) *QSignalMapper {
	var n = new(QSignalMapper)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSignalMapper_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSignalMapper) QSignalMapper_PTR() *QSignalMapper {
	return ptr
}

func NewQSignalMapper(parent QObject_ITF) *QSignalMapper {
	return NewQSignalMapperFromPointer(C.QSignalMapper_NewQSignalMapper(PointerFromQObject(parent)))
}

func (ptr *QSignalMapper) Map() {
	if ptr.Pointer() != nil {
		C.QSignalMapper_Map(ptr.Pointer())
	}
}

func (ptr *QSignalMapper) Map2(sender QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalMapper_Map2(ptr.Pointer(), PointerFromQObject(sender))
	}
}

func (ptr *QSignalMapper) ConnectMapped(f func(i int)) {
	if ptr.Pointer() != nil {
		C.QSignalMapper_ConnectMapped(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mapped", f)
	}
}

func (ptr *QSignalMapper) DisconnectMapped() {
	if ptr.Pointer() != nil {
		C.QSignalMapper_DisconnectMapped(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mapped")
	}
}

//export callbackQSignalMapperMapped
func callbackQSignalMapperMapped(ptrName *C.char, i C.int) {
	qt.GetSignal(C.GoString(ptrName), "mapped").(func(int))(int(i))
}

func (ptr *QSignalMapper) Mapping4(object QObject_ITF) *QObject {
	if ptr.Pointer() != nil {
		return NewQObjectFromPointer(C.QSignalMapper_Mapping4(ptr.Pointer(), PointerFromQObject(object)))
	}
	return nil
}

func (ptr *QSignalMapper) Mapping2(id string) *QObject {
	if ptr.Pointer() != nil {
		return NewQObjectFromPointer(C.QSignalMapper_Mapping2(ptr.Pointer(), C.CString(id)))
	}
	return nil
}

func (ptr *QSignalMapper) Mapping(id int) *QObject {
	if ptr.Pointer() != nil {
		return NewQObjectFromPointer(C.QSignalMapper_Mapping(ptr.Pointer(), C.int(id)))
	}
	return nil
}

func (ptr *QSignalMapper) RemoveMappings(sender QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalMapper_RemoveMappings(ptr.Pointer(), PointerFromQObject(sender))
	}
}

func (ptr *QSignalMapper) SetMapping4(sender QObject_ITF, object QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalMapper_SetMapping4(ptr.Pointer(), PointerFromQObject(sender), PointerFromQObject(object))
	}
}

func (ptr *QSignalMapper) SetMapping2(sender QObject_ITF, text string) {
	if ptr.Pointer() != nil {
		C.QSignalMapper_SetMapping2(ptr.Pointer(), PointerFromQObject(sender), C.CString(text))
	}
}

func (ptr *QSignalMapper) SetMapping(sender QObject_ITF, id int) {
	if ptr.Pointer() != nil {
		C.QSignalMapper_SetMapping(ptr.Pointer(), PointerFromQObject(sender), C.int(id))
	}
}

func (ptr *QSignalMapper) DestroyQSignalMapper() {
	if ptr.Pointer() != nil {
		C.QSignalMapper_DestroyQSignalMapper(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
