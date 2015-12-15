package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QStaticPlugin struct {
	ptr unsafe.Pointer
}

type QStaticPlugin_ITF interface {
	QStaticPlugin_PTR() *QStaticPlugin
}

func (p *QStaticPlugin) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QStaticPlugin) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQStaticPlugin(ptr QStaticPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStaticPlugin_PTR().Pointer()
	}
	return nil
}

func NewQStaticPluginFromPointer(ptr unsafe.Pointer) *QStaticPlugin {
	var n = new(QStaticPlugin)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStaticPlugin) QStaticPlugin_PTR() *QStaticPlugin {
	return ptr
}

func (ptr *QStaticPlugin) Instance() *QObject {
	defer qt.Recovering("QStaticPlugin::instance")

	if ptr.Pointer() != nil {
		return NewQObjectFromPointer(C.QStaticPlugin_Instance(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStaticPlugin) MetaData() *QJsonObject {
	defer qt.Recovering("QStaticPlugin::metaData")

	if ptr.Pointer() != nil {
		return NewQJsonObjectFromPointer(C.QStaticPlugin_MetaData(ptr.Pointer()))
	}
	return nil
}
