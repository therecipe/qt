package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QIconEngine struct {
	ptr unsafe.Pointer
}

type QIconEngine_ITF interface {
	QIconEngine_PTR() *QIconEngine
}

func (p *QIconEngine) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QIconEngine) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQIconEngine(ptr QIconEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIconEngine_PTR().Pointer()
	}
	return nil
}

func NewQIconEngineFromPointer(ptr unsafe.Pointer) *QIconEngine {
	var n = new(QIconEngine)
	n.SetPointer(ptr)
	return n
}

func (ptr *QIconEngine) QIconEngine_PTR() *QIconEngine {
	return ptr
}

//QIconEngine::IconEngineHook
type QIconEngine__IconEngineHook int64

const (
	QIconEngine__AvailableSizesHook = QIconEngine__IconEngineHook(1)
	QIconEngine__IconNameHook       = QIconEngine__IconEngineHook(2)
)

func (ptr *QIconEngine) AddFile(fileName string, size core.QSize_ITF, mode QIcon__Mode, state QIcon__State) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QIconEngine::addFile")
		}
	}()

	if ptr.Pointer() != nil {
		C.QIconEngine_AddFile(ptr.Pointer(), C.CString(fileName), core.PointerFromQSize(size), C.int(mode), C.int(state))
	}
}

func (ptr *QIconEngine) AddPixmap(pixmap QPixmap_ITF, mode QIcon__Mode, state QIcon__State) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QIconEngine::addPixmap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QIconEngine_AddPixmap(ptr.Pointer(), PointerFromQPixmap(pixmap), C.int(mode), C.int(state))
	}
}

func (ptr *QIconEngine) Clone() *QIconEngine {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QIconEngine::clone")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQIconEngineFromPointer(C.QIconEngine_Clone(ptr.Pointer()))
	}
	return nil
}

func (ptr *QIconEngine) IconName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QIconEngine::iconName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QIconEngine_IconName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QIconEngine) Key() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QIconEngine::key")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QIconEngine_Key(ptr.Pointer()))
	}
	return ""
}

func (ptr *QIconEngine) Paint(painter QPainter_ITF, rect core.QRect_ITF, mode QIcon__Mode, state QIcon__State) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QIconEngine::paint")
		}
	}()

	if ptr.Pointer() != nil {
		C.QIconEngine_Paint(ptr.Pointer(), PointerFromQPainter(painter), core.PointerFromQRect(rect), C.int(mode), C.int(state))
	}
}

func (ptr *QIconEngine) Read(in core.QDataStream_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QIconEngine::read")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QIconEngine_Read(ptr.Pointer(), core.PointerFromQDataStream(in)) != 0
	}
	return false
}

func (ptr *QIconEngine) Write(out core.QDataStream_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QIconEngine::write")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QIconEngine_Write(ptr.Pointer(), core.PointerFromQDataStream(out)) != 0
	}
	return false
}

func (ptr *QIconEngine) DestroyQIconEngine() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QIconEngine::~QIconEngine")
		}
	}()

	if ptr.Pointer() != nil {
		C.QIconEngine_DestroyQIconEngine(ptr.Pointer())
	}
}
