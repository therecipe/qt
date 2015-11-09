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

type QIconEnginePlugin_ITF interface {
	core.QObject_ITF
	QIconEnginePlugin_PTR() *QIconEnginePlugin
}

func PointerFromQIconEnginePlugin(ptr QIconEnginePlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIconEnginePlugin_PTR().Pointer()
	}
	return nil
}

func NewQIconEnginePluginFromPointer(ptr unsafe.Pointer) *QIconEnginePlugin {
	var n = new(QIconEnginePlugin)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QIconEnginePlugin_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QIconEnginePlugin) QIconEnginePlugin_PTR() *QIconEnginePlugin {
	return ptr
}
