package help

//#include "help.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QHelpEngine struct {
	QHelpEngineCore
}

type QHelpEngine_ITF interface {
	QHelpEngineCore_ITF
	QHelpEngine_PTR() *QHelpEngine
}

func PointerFromQHelpEngine(ptr QHelpEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpEngine_PTR().Pointer()
	}
	return nil
}

func NewQHelpEngineFromPointer(ptr unsafe.Pointer) *QHelpEngine {
	var n = new(QHelpEngine)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QHelpEngine_") {
		n.SetObjectName("QHelpEngine_" + qt.Identifier())
	}
	return n
}

func (ptr *QHelpEngine) QHelpEngine_PTR() *QHelpEngine {
	return ptr
}

func NewQHelpEngine(collectionFile string, parent core.QObject_ITF) *QHelpEngine {
	defer qt.Recovering("QHelpEngine::QHelpEngine")

	return NewQHelpEngineFromPointer(C.QHelpEngine_NewQHelpEngine(C.CString(collectionFile), core.PointerFromQObject(parent)))
}

func (ptr *QHelpEngine) ContentModel() *QHelpContentModel {
	defer qt.Recovering("QHelpEngine::contentModel")

	if ptr.Pointer() != nil {
		return NewQHelpContentModelFromPointer(C.QHelpEngine_ContentModel(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpEngine) ContentWidget() *QHelpContentWidget {
	defer qt.Recovering("QHelpEngine::contentWidget")

	if ptr.Pointer() != nil {
		return NewQHelpContentWidgetFromPointer(C.QHelpEngine_ContentWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpEngine) IndexModel() *QHelpIndexModel {
	defer qt.Recovering("QHelpEngine::indexModel")

	if ptr.Pointer() != nil {
		return NewQHelpIndexModelFromPointer(C.QHelpEngine_IndexModel(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpEngine) IndexWidget() *QHelpIndexWidget {
	defer qt.Recovering("QHelpEngine::indexWidget")

	if ptr.Pointer() != nil {
		return NewQHelpIndexWidgetFromPointer(C.QHelpEngine_IndexWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpEngine) SearchEngine() *QHelpSearchEngine {
	defer qt.Recovering("QHelpEngine::searchEngine")

	if ptr.Pointer() != nil {
		return NewQHelpSearchEngineFromPointer(C.QHelpEngine_SearchEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpEngine) DestroyQHelpEngine() {
	defer qt.Recovering("QHelpEngine::~QHelpEngine")

	if ptr.Pointer() != nil {
		C.QHelpEngine_DestroyQHelpEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
