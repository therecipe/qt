package gui

//#include "qgenericpluginfactory.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QGenericPluginFactory struct {
	ptr unsafe.Pointer
}

type QGenericPluginFactoryITF interface {
	QGenericPluginFactoryPTR() *QGenericPluginFactory
}

func (p *QGenericPluginFactory) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGenericPluginFactory) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGenericPluginFactory(ptr QGenericPluginFactoryITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGenericPluginFactoryPTR().Pointer()
	}
	return nil
}

func QGenericPluginFactoryFromPointer(ptr unsafe.Pointer) *QGenericPluginFactory {
	var n = new(QGenericPluginFactory)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGenericPluginFactory) QGenericPluginFactoryPTR() *QGenericPluginFactory {
	return ptr
}

func QGenericPluginFactory_Create(key string, specification string) *core.QObject {
	return core.QObjectFromPointer(unsafe.Pointer(C.QGenericPluginFactory_QGenericPluginFactory_Create(C.CString(key), C.CString(specification))))
}

func QGenericPluginFactory_Keys() []string {
	return strings.Split(C.GoString(C.QGenericPluginFactory_QGenericPluginFactory_Keys()), "|")
}
