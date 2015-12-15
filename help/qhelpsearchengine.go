package help

//#include "help.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QHelpSearchEngine struct {
	core.QObject
}

type QHelpSearchEngine_ITF interface {
	core.QObject_ITF
	QHelpSearchEngine_PTR() *QHelpSearchEngine
}

func PointerFromQHelpSearchEngine(ptr QHelpSearchEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchEngine_PTR().Pointer()
	}
	return nil
}

func NewQHelpSearchEngineFromPointer(ptr unsafe.Pointer) *QHelpSearchEngine {
	var n = new(QHelpSearchEngine)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QHelpSearchEngine_") {
		n.SetObjectName("QHelpSearchEngine_" + qt.Identifier())
	}
	return n
}

func (ptr *QHelpSearchEngine) QHelpSearchEngine_PTR() *QHelpSearchEngine {
	return ptr
}

func NewQHelpSearchEngine(helpEngine QHelpEngineCore_ITF, parent core.QObject_ITF) *QHelpSearchEngine {
	defer qt.Recovering("QHelpSearchEngine::QHelpSearchEngine")

	return NewQHelpSearchEngineFromPointer(C.QHelpSearchEngine_NewQHelpSearchEngine(PointerFromQHelpEngineCore(helpEngine), core.PointerFromQObject(parent)))
}

func (ptr *QHelpSearchEngine) CancelIndexing() {
	defer qt.Recovering("QHelpSearchEngine::cancelIndexing")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_CancelIndexing(ptr.Pointer())
	}
}

func (ptr *QHelpSearchEngine) CancelSearching() {
	defer qt.Recovering("QHelpSearchEngine::cancelSearching")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_CancelSearching(ptr.Pointer())
	}
}

func (ptr *QHelpSearchEngine) HitCount() int {
	defer qt.Recovering("QHelpSearchEngine::hitCount")

	if ptr.Pointer() != nil {
		return int(C.QHelpSearchEngine_HitCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHelpSearchEngine) ConnectIndexingFinished(f func()) {
	defer qt.Recovering("connect QHelpSearchEngine::indexingFinished")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectIndexingFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "indexingFinished", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectIndexingFinished() {
	defer qt.Recovering("disconnect QHelpSearchEngine::indexingFinished")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectIndexingFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "indexingFinished")
	}
}

//export callbackQHelpSearchEngineIndexingFinished
func callbackQHelpSearchEngineIndexingFinished(ptrName *C.char) {
	defer qt.Recovering("callback QHelpSearchEngine::indexingFinished")

	var signal = qt.GetSignal(C.GoString(ptrName), "indexingFinished")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpSearchEngine) ConnectIndexingStarted(f func()) {
	defer qt.Recovering("connect QHelpSearchEngine::indexingStarted")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectIndexingStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "indexingStarted", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectIndexingStarted() {
	defer qt.Recovering("disconnect QHelpSearchEngine::indexingStarted")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectIndexingStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "indexingStarted")
	}
}

//export callbackQHelpSearchEngineIndexingStarted
func callbackQHelpSearchEngineIndexingStarted(ptrName *C.char) {
	defer qt.Recovering("callback QHelpSearchEngine::indexingStarted")

	var signal = qt.GetSignal(C.GoString(ptrName), "indexingStarted")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpSearchEngine) QueryWidget() *QHelpSearchQueryWidget {
	defer qt.Recovering("QHelpSearchEngine::queryWidget")

	if ptr.Pointer() != nil {
		return NewQHelpSearchQueryWidgetFromPointer(C.QHelpSearchEngine_QueryWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpSearchEngine) ReindexDocumentation() {
	defer qt.Recovering("QHelpSearchEngine::reindexDocumentation")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ReindexDocumentation(ptr.Pointer())
	}
}

func (ptr *QHelpSearchEngine) ResultWidget() *QHelpSearchResultWidget {
	defer qt.Recovering("QHelpSearchEngine::resultWidget")

	if ptr.Pointer() != nil {
		return NewQHelpSearchResultWidgetFromPointer(C.QHelpSearchEngine_ResultWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpSearchEngine) ConnectSearchingFinished(f func(hits int)) {
	defer qt.Recovering("connect QHelpSearchEngine::searchingFinished")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectSearchingFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "searchingFinished", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectSearchingFinished() {
	defer qt.Recovering("disconnect QHelpSearchEngine::searchingFinished")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectSearchingFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "searchingFinished")
	}
}

//export callbackQHelpSearchEngineSearchingFinished
func callbackQHelpSearchEngineSearchingFinished(ptrName *C.char, hits C.int) {
	defer qt.Recovering("callback QHelpSearchEngine::searchingFinished")

	var signal = qt.GetSignal(C.GoString(ptrName), "searchingFinished")
	if signal != nil {
		signal.(func(int))(int(hits))
	}

}

func (ptr *QHelpSearchEngine) ConnectSearchingStarted(f func()) {
	defer qt.Recovering("connect QHelpSearchEngine::searchingStarted")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectSearchingStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "searchingStarted", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectSearchingStarted() {
	defer qt.Recovering("disconnect QHelpSearchEngine::searchingStarted")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectSearchingStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "searchingStarted")
	}
}

//export callbackQHelpSearchEngineSearchingStarted
func callbackQHelpSearchEngineSearchingStarted(ptrName *C.char) {
	defer qt.Recovering("callback QHelpSearchEngine::searchingStarted")

	var signal = qt.GetSignal(C.GoString(ptrName), "searchingStarted")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpSearchEngine) DestroyQHelpSearchEngine() {
	defer qt.Recovering("QHelpSearchEngine::~QHelpSearchEngine")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DestroyQHelpSearchEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
