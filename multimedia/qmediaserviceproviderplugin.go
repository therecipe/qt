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

type QMediaServiceProviderPluginITF interface {
	core.QObjectITF
	QMediaServiceProviderPluginPTR() *QMediaServiceProviderPlugin
}

func PointerFromQMediaServiceProviderPlugin(ptr QMediaServiceProviderPluginITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaServiceProviderPluginPTR().Pointer()
	}
	return nil
}

func QMediaServiceProviderPluginFromPointer(ptr unsafe.Pointer) *QMediaServiceProviderPlugin {
	var n = new(QMediaServiceProviderPlugin)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMediaServiceProviderPlugin_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaServiceProviderPlugin) QMediaServiceProviderPluginPTR() *QMediaServiceProviderPlugin {
	return ptr
}

func (ptr *QMediaServiceProviderPlugin) Create(key string) *QMediaService {
	if ptr.Pointer() != nil {
		return QMediaServiceFromPointer(unsafe.Pointer(C.QMediaServiceProviderPlugin_Create(C.QtObjectPtr(ptr.Pointer()), C.CString(key))))
	}
	return nil
}

func (ptr *QMediaServiceProviderPlugin) Release(service QMediaServiceITF) {
	if ptr.Pointer() != nil {
		C.QMediaServiceProviderPlugin_Release(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMediaService(service)))
	}
}
