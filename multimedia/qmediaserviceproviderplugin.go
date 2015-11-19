package multimedia

//#include "qmediaserviceproviderplugin.h"
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
	if n.ObjectName() == "" {
		n.SetObjectName("QMediaServiceProviderPlugin_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaServiceProviderPlugin) QMediaServiceProviderPlugin_PTR() *QMediaServiceProviderPlugin {
	return ptr
}

func (ptr *QMediaServiceProviderPlugin) Create(key string) *QMediaService {
	if ptr.Pointer() != nil {
		return NewQMediaServiceFromPointer(C.QMediaServiceProviderPlugin_Create(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QMediaServiceProviderPlugin) Release(service QMediaService_ITF) {
	if ptr.Pointer() != nil {
		C.QMediaServiceProviderPlugin_Release(ptr.Pointer(), PointerFromQMediaService(service))
	}
}
