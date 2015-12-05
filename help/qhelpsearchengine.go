package help

//#include "help.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
		n.SetObjectName("QHelpSearchEngine_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHelpSearchEngine) QHelpSearchEngine_PTR() *QHelpSearchEngine {
	return ptr
}

func NewQHelpSearchEngine(helpEngine QHelpEngineCore_ITF, parent core.QObject_ITF) *QHelpSearchEngine {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchEngine::QHelpSearchEngine")
		}
	}()

	return NewQHelpSearchEngineFromPointer(C.QHelpSearchEngine_NewQHelpSearchEngine(PointerFromQHelpEngineCore(helpEngine), core.PointerFromQObject(parent)))
}

func (ptr *QHelpSearchEngine) CancelIndexing() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchEngine::cancelIndexing")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_CancelIndexing(ptr.Pointer())
	}
}

func (ptr *QHelpSearchEngine) CancelSearching() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchEngine::cancelSearching")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_CancelSearching(ptr.Pointer())
	}
}

func (ptr *QHelpSearchEngine) HitCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchEngine::hitCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHelpSearchEngine_HitCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHelpSearchEngine) ConnectIndexingFinished(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchEngine::indexingFinished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectIndexingFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "indexingFinished", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectIndexingFinished() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchEngine::indexingFinished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectIndexingFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "indexingFinished")
	}
}

//export callbackQHelpSearchEngineIndexingFinished
func callbackQHelpSearchEngineIndexingFinished(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchEngine::indexingFinished")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "indexingFinished").(func())()
}

func (ptr *QHelpSearchEngine) ConnectIndexingStarted(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchEngine::indexingStarted")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectIndexingStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "indexingStarted", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectIndexingStarted() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchEngine::indexingStarted")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectIndexingStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "indexingStarted")
	}
}

//export callbackQHelpSearchEngineIndexingStarted
func callbackQHelpSearchEngineIndexingStarted(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchEngine::indexingStarted")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "indexingStarted").(func())()
}

func (ptr *QHelpSearchEngine) QueryWidget() *QHelpSearchQueryWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchEngine::queryWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQHelpSearchQueryWidgetFromPointer(C.QHelpSearchEngine_QueryWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpSearchEngine) ReindexDocumentation() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchEngine::reindexDocumentation")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ReindexDocumentation(ptr.Pointer())
	}
}

func (ptr *QHelpSearchEngine) ResultWidget() *QHelpSearchResultWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchEngine::resultWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQHelpSearchResultWidgetFromPointer(C.QHelpSearchEngine_ResultWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpSearchEngine) ConnectSearchingFinished(f func(hits int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchEngine::searchingFinished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectSearchingFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "searchingFinished", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectSearchingFinished() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchEngine::searchingFinished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectSearchingFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "searchingFinished")
	}
}

//export callbackQHelpSearchEngineSearchingFinished
func callbackQHelpSearchEngineSearchingFinished(ptrName *C.char, hits C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchEngine::searchingFinished")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "searchingFinished").(func(int))(int(hits))
}

func (ptr *QHelpSearchEngine) ConnectSearchingStarted(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchEngine::searchingStarted")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectSearchingStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "searchingStarted", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectSearchingStarted() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchEngine::searchingStarted")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectSearchingStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "searchingStarted")
	}
}

//export callbackQHelpSearchEngineSearchingStarted
func callbackQHelpSearchEngineSearchingStarted(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchEngine::searchingStarted")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "searchingStarted").(func())()
}

func (ptr *QHelpSearchEngine) DestroyQHelpSearchEngine() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchEngine::~QHelpSearchEngine")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DestroyQHelpSearchEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
