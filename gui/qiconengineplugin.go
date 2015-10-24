package gui

//#include "qiconengineplugin.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QIconEnginePlugin struct {
	core.QObject
}

type QIconEnginePluginITF interface {
	core.QObjectITF
	QIconEnginePluginPTR() *QIconEnginePlugin
}

func PointerFromQIconEnginePlugin(ptr QIconEnginePluginITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIconEnginePluginPTR().Pointer()
	}
	return nil
}

func QIconEnginePluginFromPointer(ptr unsafe.Pointer) *QIconEnginePlugin {
	var n = new(QIconEnginePlugin)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QIconEnginePlugin_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QIconEnginePlugin) QIconEnginePluginPTR() *QIconEnginePlugin {
	return ptr
}
