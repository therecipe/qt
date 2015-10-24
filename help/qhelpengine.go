package help

//#include "qhelpengine.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QHelpEngine struct {
	QHelpEngineCore
}

type QHelpEngineITF interface {
	QHelpEngineCoreITF
	QHelpEnginePTR() *QHelpEngine
}

func PointerFromQHelpEngine(ptr QHelpEngineITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpEnginePTR().Pointer()
	}
	return nil
}

func QHelpEngineFromPointer(ptr unsafe.Pointer) *QHelpEngine {
	var n = new(QHelpEngine)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QHelpEngine_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHelpEngine) QHelpEnginePTR() *QHelpEngine {
	return ptr
}

func NewQHelpEngine(collectionFile string, parent core.QObjectITF) *QHelpEngine {
	return QHelpEngineFromPointer(unsafe.Pointer(C.QHelpEngine_NewQHelpEngine(C.CString(collectionFile), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QHelpEngine) ContentModel() *QHelpContentModel {
	if ptr.Pointer() != nil {
		return QHelpContentModelFromPointer(unsafe.Pointer(C.QHelpEngine_ContentModel(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QHelpEngine) ContentWidget() *QHelpContentWidget {
	if ptr.Pointer() != nil {
		return QHelpContentWidgetFromPointer(unsafe.Pointer(C.QHelpEngine_ContentWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QHelpEngine) IndexModel() *QHelpIndexModel {
	if ptr.Pointer() != nil {
		return QHelpIndexModelFromPointer(unsafe.Pointer(C.QHelpEngine_IndexModel(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QHelpEngine) IndexWidget() *QHelpIndexWidget {
	if ptr.Pointer() != nil {
		return QHelpIndexWidgetFromPointer(unsafe.Pointer(C.QHelpEngine_IndexWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QHelpEngine) SearchEngine() *QHelpSearchEngine {
	if ptr.Pointer() != nil {
		return QHelpSearchEngineFromPointer(unsafe.Pointer(C.QHelpEngine_SearchEngine(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QHelpEngine) DestroyQHelpEngine() {
	if ptr.Pointer() != nil {
		C.QHelpEngine_DestroyQHelpEngine(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
