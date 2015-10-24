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

type QSignalMapperITF interface {
	QObjectITF
	QSignalMapperPTR() *QSignalMapper
}

func PointerFromQSignalMapper(ptr QSignalMapperITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSignalMapperPTR().Pointer()
	}
	return nil
}

func QSignalMapperFromPointer(ptr unsafe.Pointer) *QSignalMapper {
	var n = new(QSignalMapper)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSignalMapper_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSignalMapper) QSignalMapperPTR() *QSignalMapper {
	return ptr
}

func NewQSignalMapper(parent QObjectITF) *QSignalMapper {
	return QSignalMapperFromPointer(unsafe.Pointer(C.QSignalMapper_NewQSignalMapper(C.QtObjectPtr(PointerFromQObject(parent)))))
}

func (ptr *QSignalMapper) Map() {
	if ptr.Pointer() != nil {
		C.QSignalMapper_Map(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSignalMapper) Map2(sender QObjectITF) {
	if ptr.Pointer() != nil {
		C.QSignalMapper_Map2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(sender)))
	}
}

func (ptr *QSignalMapper) ConnectMapped(f func(i int)) {
	if ptr.Pointer() != nil {
		C.QSignalMapper_ConnectMapped(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "mapped", f)
	}
}

func (ptr *QSignalMapper) DisconnectMapped() {
	if ptr.Pointer() != nil {
		C.QSignalMapper_DisconnectMapped(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "mapped")
	}
}

//export callbackQSignalMapperMapped
func callbackQSignalMapperMapped(ptrName *C.char, i C.int) {
	qt.GetSignal(C.GoString(ptrName), "mapped").(func(int))(int(i))
}

func (ptr *QSignalMapper) Mapping4(object QObjectITF) *QObject {
	if ptr.Pointer() != nil {
		return QObjectFromPointer(unsafe.Pointer(C.QSignalMapper_Mapping4(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(object)))))
	}
	return nil
}

func (ptr *QSignalMapper) Mapping2(id string) *QObject {
	if ptr.Pointer() != nil {
		return QObjectFromPointer(unsafe.Pointer(C.QSignalMapper_Mapping2(C.QtObjectPtr(ptr.Pointer()), C.CString(id))))
	}
	return nil
}

func (ptr *QSignalMapper) Mapping(id int) *QObject {
	if ptr.Pointer() != nil {
		return QObjectFromPointer(unsafe.Pointer(C.QSignalMapper_Mapping(C.QtObjectPtr(ptr.Pointer()), C.int(id))))
	}
	return nil
}

func (ptr *QSignalMapper) RemoveMappings(sender QObjectITF) {
	if ptr.Pointer() != nil {
		C.QSignalMapper_RemoveMappings(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(sender)))
	}
}

func (ptr *QSignalMapper) SetMapping4(sender QObjectITF, object QObjectITF) {
	if ptr.Pointer() != nil {
		C.QSignalMapper_SetMapping4(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(sender)), C.QtObjectPtr(PointerFromQObject(object)))
	}
}

func (ptr *QSignalMapper) SetMapping2(sender QObjectITF, text string) {
	if ptr.Pointer() != nil {
		C.QSignalMapper_SetMapping2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(sender)), C.CString(text))
	}
}

func (ptr *QSignalMapper) SetMapping(sender QObjectITF, id int) {
	if ptr.Pointer() != nil {
		C.QSignalMapper_SetMapping(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(sender)), C.int(id))
	}
}

func (ptr *QSignalMapper) DestroyQSignalMapper() {
	if ptr.Pointer() != nil {
		C.QSignalMapper_DestroyQSignalMapper(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
