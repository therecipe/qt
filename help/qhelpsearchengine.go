package help

//#include "qhelpsearchengine.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QHelpSearchEngine struct {
	core.QObject
}

type QHelpSearchEngineITF interface {
	core.QObjectITF
	QHelpSearchEnginePTR() *QHelpSearchEngine
}

func PointerFromQHelpSearchEngine(ptr QHelpSearchEngineITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchEnginePTR().Pointer()
	}
	return nil
}

func QHelpSearchEngineFromPointer(ptr unsafe.Pointer) *QHelpSearchEngine {
	var n = new(QHelpSearchEngine)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QHelpSearchEngine_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHelpSearchEngine) QHelpSearchEnginePTR() *QHelpSearchEngine {
	return ptr
}

func NewQHelpSearchEngine(helpEngine QHelpEngineCoreITF, parent core.QObjectITF) *QHelpSearchEngine {
	return QHelpSearchEngineFromPointer(unsafe.Pointer(C.QHelpSearchEngine_NewQHelpSearchEngine(C.QtObjectPtr(PointerFromQHelpEngineCore(helpEngine)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QHelpSearchEngine) CancelIndexing() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_CancelIndexing(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QHelpSearchEngine) CancelSearching() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_CancelSearching(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QHelpSearchEngine) HitCount() int {
	if ptr.Pointer() != nil {
		return int(C.QHelpSearchEngine_HitCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHelpSearchEngine) ConnectIndexingFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectIndexingFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "indexingFinished", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectIndexingFinished() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectIndexingFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "indexingFinished")
	}
}

//export callbackQHelpSearchEngineIndexingFinished
func callbackQHelpSearchEngineIndexingFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "indexingFinished").(func())()
}

func (ptr *QHelpSearchEngine) ConnectIndexingStarted(f func()) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectIndexingStarted(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "indexingStarted", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectIndexingStarted() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectIndexingStarted(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "indexingStarted")
	}
}

//export callbackQHelpSearchEngineIndexingStarted
func callbackQHelpSearchEngineIndexingStarted(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "indexingStarted").(func())()
}

func (ptr *QHelpSearchEngine) QueryWidget() *QHelpSearchQueryWidget {
	if ptr.Pointer() != nil {
		return QHelpSearchQueryWidgetFromPointer(unsafe.Pointer(C.QHelpSearchEngine_QueryWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QHelpSearchEngine) ReindexDocumentation() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ReindexDocumentation(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QHelpSearchEngine) ResultWidget() *QHelpSearchResultWidget {
	if ptr.Pointer() != nil {
		return QHelpSearchResultWidgetFromPointer(unsafe.Pointer(C.QHelpSearchEngine_ResultWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QHelpSearchEngine) ConnectSearchingFinished(f func(hits int)) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectSearchingFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "searchingFinished", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectSearchingFinished() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectSearchingFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "searchingFinished")
	}
}

//export callbackQHelpSearchEngineSearchingFinished
func callbackQHelpSearchEngineSearchingFinished(ptrName *C.char, hits C.int) {
	qt.GetSignal(C.GoString(ptrName), "searchingFinished").(func(int))(int(hits))
}

func (ptr *QHelpSearchEngine) ConnectSearchingStarted(f func()) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectSearchingStarted(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "searchingStarted", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectSearchingStarted() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectSearchingStarted(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "searchingStarted")
	}
}

//export callbackQHelpSearchEngineSearchingStarted
func callbackQHelpSearchEngineSearchingStarted(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "searchingStarted").(func())()
}

func (ptr *QHelpSearchEngine) DestroyQHelpSearchEngine() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DestroyQHelpSearchEngine(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
