package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
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
	for len(n.ObjectNameAbs()) < len("QIconEngine_") {
		n.SetObjectNameAbs("QIconEngine_" + qt.Identifier())
	}
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

func (ptr *QIconEngine) Clone() *QIconEngine {
	defer qt.Recovering("QIconEngine::clone")

	if ptr.Pointer() != nil {
		return NewQIconEngineFromPointer(C.QIconEngine_Clone(ptr.Pointer()))
	}
	return nil
}

func (ptr *QIconEngine) IconName() string {
	defer qt.Recovering("QIconEngine::iconName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QIconEngine_IconName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QIconEngine) Key() string {
	defer qt.Recovering("QIconEngine::key")

	if ptr.Pointer() != nil {
		return C.GoString(C.QIconEngine_Key(ptr.Pointer()))
	}
	return ""
}

func (ptr *QIconEngine) Paint(painter QPainter_ITF, rect core.QRect_ITF, mode QIcon__Mode, state QIcon__State) {
	defer qt.Recovering("QIconEngine::paint")

	if ptr.Pointer() != nil {
		C.QIconEngine_Paint(ptr.Pointer(), PointerFromQPainter(painter), core.PointerFromQRect(rect), C.int(mode), C.int(state))
	}
}

func (ptr *QIconEngine) Read(in core.QDataStream_ITF) bool {
	defer qt.Recovering("QIconEngine::read")

	if ptr.Pointer() != nil {
		return C.QIconEngine_Read(ptr.Pointer(), core.PointerFromQDataStream(in)) != 0
	}
	return false
}

func (ptr *QIconEngine) Write(out core.QDataStream_ITF) bool {
	defer qt.Recovering("QIconEngine::write")

	if ptr.Pointer() != nil {
		return C.QIconEngine_Write(ptr.Pointer(), core.PointerFromQDataStream(out)) != 0
	}
	return false
}

func (ptr *QIconEngine) DestroyQIconEngine() {
	defer qt.Recovering("QIconEngine::~QIconEngine")

	if ptr.Pointer() != nil {
		C.QIconEngine_DestroyQIconEngine(ptr.Pointer())
	}
}

func (ptr *QIconEngine) ObjectNameAbs() string {
	defer qt.Recovering("QIconEngine::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QIconEngine_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QIconEngine) SetObjectNameAbs(name string) {
	defer qt.Recovering("QIconEngine::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QIconEngine_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
