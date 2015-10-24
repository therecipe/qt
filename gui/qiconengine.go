package gui

//#include "qiconengine.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QIconEngine struct {
	ptr unsafe.Pointer
}

type QIconEngineITF interface {
	QIconEnginePTR() *QIconEngine
}

func (p *QIconEngine) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QIconEngine) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQIconEngine(ptr QIconEngineITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIconEnginePTR().Pointer()
	}
	return nil
}

func QIconEngineFromPointer(ptr unsafe.Pointer) *QIconEngine {
	var n = new(QIconEngine)
	n.SetPointer(ptr)
	return n
}

func (ptr *QIconEngine) QIconEnginePTR() *QIconEngine {
	return ptr
}

//QIconEngine::IconEngineHook
type QIconEngine__IconEngineHook int

var (
	QIconEngine__AvailableSizesHook = QIconEngine__IconEngineHook(1)
	QIconEngine__IconNameHook       = QIconEngine__IconEngineHook(2)
)

func (ptr *QIconEngine) AddFile(fileName string, size core.QSizeITF, mode QIcon__Mode, state QIcon__State) {
	if ptr.Pointer() != nil {
		C.QIconEngine_AddFile(C.QtObjectPtr(ptr.Pointer()), C.CString(fileName), C.QtObjectPtr(core.PointerFromQSize(size)), C.int(mode), C.int(state))
	}
}

func (ptr *QIconEngine) AddPixmap(pixmap QPixmapITF, mode QIcon__Mode, state QIcon__State) {
	if ptr.Pointer() != nil {
		C.QIconEngine_AddPixmap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPixmap(pixmap)), C.int(mode), C.int(state))
	}
}

func (ptr *QIconEngine) Clone() *QIconEngine {
	if ptr.Pointer() != nil {
		return QIconEngineFromPointer(unsafe.Pointer(C.QIconEngine_Clone(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QIconEngine) IconName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QIconEngine_IconName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QIconEngine) Key() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QIconEngine_Key(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QIconEngine) Paint(painter QPainterITF, rect core.QRectITF, mode QIcon__Mode, state QIcon__State) {
	if ptr.Pointer() != nil {
		C.QIconEngine_Paint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPainter(painter)), C.QtObjectPtr(core.PointerFromQRect(rect)), C.int(mode), C.int(state))
	}
}

func (ptr *QIconEngine) Read(in core.QDataStreamITF) bool {
	if ptr.Pointer() != nil {
		return C.QIconEngine_Read(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDataStream(in))) != 0
	}
	return false
}

func (ptr *QIconEngine) Write(out core.QDataStreamITF) bool {
	if ptr.Pointer() != nil {
		return C.QIconEngine_Write(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDataStream(out))) != 0
	}
	return false
}

func (ptr *QIconEngine) DestroyQIconEngine() {
	if ptr.Pointer() != nil {
		C.QIconEngine_DestroyQIconEngine(C.QtObjectPtr(ptr.Pointer()))
	}
}
