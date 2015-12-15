package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMediaServiceProviderPlugin struct {
	core.QObject
}

type QMediaServiceProviderPlugin_ITF interface {
	core.QObject_ITF
	QMediaServiceProviderPlugin_PTR() *QMediaServiceProviderPlugin
}

func PointerFromQMediaServiceProviderPlugin(ptr QMediaServiceProviderPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaServiceProviderPlugin_PTR().Pointer()
	}
	return nil
}

func NewQMediaServiceProviderPluginFromPointer(ptr unsafe.Pointer) *QMediaServiceProviderPlugin {
	var n = new(QMediaServiceProviderPlugin)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMediaServiceProviderPlugin_") {
		n.SetObjectName("QMediaServiceProviderPlugin_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaServiceProviderPlugin) QMediaServiceProviderPlugin_PTR() *QMediaServiceProviderPlugin {
	return ptr
}

func (ptr *QMediaServiceProviderPlugin) Create(key string) *QMediaService {
	defer qt.Recovering("QMediaServiceProviderPlugin::create")

	if ptr.Pointer() != nil {
		return NewQMediaServiceFromPointer(C.QMediaServiceProviderPlugin_Create(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QMediaServiceProviderPlugin) Release(service QMediaService_ITF) {
	defer qt.Recovering("QMediaServiceProviderPlugin::release")

	if ptr.Pointer() != nil {
		C.QMediaServiceProviderPlugin_Release(ptr.Pointer(), PointerFromQMediaService(service))
	}
}
