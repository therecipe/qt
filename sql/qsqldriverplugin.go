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

type QSqlDriverPluginITF interface {
	core.QObjectITF
	QSqlDriverPluginPTR() *QSqlDriverPlugin
}

func PointerFromQSqlDriverPlugin(ptr QSqlDriverPluginITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlDriverPluginPTR().Pointer()
	}
	return nil
}

func QSqlDriverPluginFromPointer(ptr unsafe.Pointer) *QSqlDriverPlugin {
	var n = new(QSqlDriverPlugin)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSqlDriverPlugin_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSqlDriverPlugin) QSqlDriverPluginPTR() *QSqlDriverPlugin {
	return ptr
}

func (ptr *QSqlDriverPlugin) Create(key string) *QSqlDriver {
	if ptr.Pointer() != nil {
		return QSqlDriverFromPointer(unsafe.Pointer(C.QSqlDriverPlugin_Create(C.QtObjectPtr(ptr.Pointer()), C.CString(key))))
	}
	return nil
}

func (ptr *QSqlDriverPlugin) DestroyQSqlDriverPlugin() {
	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_DestroyQSqlDriverPlugin(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
