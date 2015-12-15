package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QStylePlugin struct {
	core.QObject
}

type QStylePlugin_ITF interface {
	core.QObject_ITF
	QStylePlugin_PTR() *QStylePlugin
}

func PointerFromQStylePlugin(ptr QStylePlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStylePlugin_PTR().Pointer()
	}
	return nil
}

func NewQStylePluginFromPointer(ptr unsafe.Pointer) *QStylePlugin {
	var n = new(QStylePlugin)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QStylePlugin_") {
		n.SetObjectName("QStylePlugin_" + qt.Identifier())
	}
	return n
}

func (ptr *QStylePlugin) QStylePlugin_PTR() *QStylePlugin {
	return ptr
}

func (ptr *QStylePlugin) Create(key string) *QStyle {
	defer qt.Recovering("QStylePlugin::create")

	if ptr.Pointer() != nil {
		return NewQStyleFromPointer(C.QStylePlugin_Create(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QStylePlugin) DestroyQStylePlugin() {
	defer qt.Recovering("QStylePlugin::~QStylePlugin")

	if ptr.Pointer() != nil {
		C.QStylePlugin_DestroyQStylePlugin(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
