package sql

//#include "qsqldriverplugin.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSqlDriverPlugin struct {
	core.QObject
}

type QSqlDriverPlugin_ITF interface {
	core.QObject_ITF
	QSqlDriverPlugin_PTR() *QSqlDriverPlugin
}

func PointerFromQSqlDriverPlugin(ptr QSqlDriverPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlDriverPlugin_PTR().Pointer()
	}
	return nil
}

func NewQSqlDriverPluginFromPointer(ptr unsafe.Pointer) *QSqlDriverPlugin {
	var n = new(QSqlDriverPlugin)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QSqlDriverPlugin_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSqlDriverPlugin) QSqlDriverPlugin_PTR() *QSqlDriverPlugin {
	return ptr
}

func (ptr *QSqlDriverPlugin) Create(key string) *QSqlDriver {
	if ptr.Pointer() != nil {
		return NewQSqlDriverFromPointer(C.QSqlDriverPlugin_Create(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QSqlDriverPlugin) DestroyQSqlDriverPlugin() {
	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_DestroyQSqlDriverPlugin(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
