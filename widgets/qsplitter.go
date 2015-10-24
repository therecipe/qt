package widgets

//#include "qsplitter.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSplitter struct {
	QFrame
}

type QSplitterITF interface {
	QFrameITF
	QSplitterPTR() *QSplitter
}

func PointerFromQSplitter(ptr QSplitterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSplitterPTR().Pointer()
	}
	return nil
}

func QSplitterFromPointer(ptr unsafe.Pointer) *QSplitter {
	var n = new(QSplitter)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSplitter_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSplitter) QSplitterPTR() *QSplitter {
	return ptr
}

func (ptr *QSplitter) ChildrenCollapsible() bool {
	if ptr.Pointer() != nil {
		return C.QSplitter_ChildrenCollapsible(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSplitter) HandleWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QSplitter_HandleWidth(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSplitter) IndexOf(widget QWidgetITF) int {
	if ptr.Pointer() != nil {
		return int(C.QSplitter_IndexOf(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget))))
	}
	return 0
}

func (ptr *QSplitter) OpaqueResize() bool {
	if ptr.Pointer() != nil {
		return C.QSplitter_OpaqueResize(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSplitter) Orientation() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QSplitter_Orientation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSplitter) SetChildrenCollapsible(v bool) {
	if ptr.Pointer() != nil {
		C.QSplitter_SetChildrenCollapsible(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QSplitter) SetHandleWidth(v int) {
	if ptr.Pointer() != nil {
		C.QSplitter_SetHandleWidth(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QSplitter) SetOpaqueResize(opaque bool) {
	if ptr.Pointer() != nil {
		C.QSplitter_SetOpaqueResize(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(opaque)))
	}
}

func (ptr *QSplitter) SetOrientation(v core.Qt__Orientation) {
	if ptr.Pointer() != nil {
		C.QSplitter_SetOrientation(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func NewQSplitter(parent QWidgetITF) *QSplitter {
	return QSplitterFromPointer(unsafe.Pointer(C.QSplitter_NewQSplitter(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQSplitter2(orientation core.Qt__Orientation, parent QWidgetITF) *QSplitter {
	return QSplitterFromPointer(unsafe.Pointer(C.QSplitter_NewQSplitter2(C.int(orientation), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QSplitter) AddWidget(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QSplitter_AddWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QSplitter) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QSplitter_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSplitter) GetRange(index int, min int, max int) {
	if ptr.Pointer() != nil {
		C.QSplitter_GetRange(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.int(min), C.int(max))
	}
}

func (ptr *QSplitter) Handle(index int) *QSplitterHandle {
	if ptr.Pointer() != nil {
		return QSplitterHandleFromPointer(unsafe.Pointer(C.QSplitter_Handle(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QSplitter) InsertWidget(index int, widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QSplitter_InsertWidget(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QSplitter) IsCollapsible(index int) bool {
	if ptr.Pointer() != nil {
		return C.QSplitter_IsCollapsible(C.QtObjectPtr(ptr.Pointer()), C.int(index)) != 0
	}
	return false
}

func (ptr *QSplitter) Refresh() {
	if ptr.Pointer() != nil {
		C.QSplitter_Refresh(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSplitter) RestoreState(state core.QByteArrayITF) bool {
	if ptr.Pointer() != nil {
		return C.QSplitter_RestoreState(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(state))) != 0
	}
	return false
}

func (ptr *QSplitter) SetCollapsible(index int, collapse bool) {
	if ptr.Pointer() != nil {
		C.QSplitter_SetCollapsible(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.int(qt.GoBoolToInt(collapse)))
	}
}

func (ptr *QSplitter) SetStretchFactor(index int, stretch int) {
	if ptr.Pointer() != nil {
		C.QSplitter_SetStretchFactor(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.int(stretch))
	}
}

func (ptr *QSplitter) ConnectSplitterMoved(f func(pos int, index int)) {
	if ptr.Pointer() != nil {
		C.QSplitter_ConnectSplitterMoved(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "splitterMoved", f)
	}
}

func (ptr *QSplitter) DisconnectSplitterMoved() {
	if ptr.Pointer() != nil {
		C.QSplitter_DisconnectSplitterMoved(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "splitterMoved")
	}
}

//export callbackQSplitterSplitterMoved
func callbackQSplitterSplitterMoved(ptrName *C.char, pos C.int, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "splitterMoved").(func(int, int))(int(pos), int(index))
}

func (ptr *QSplitter) Widget(index int) *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QSplitter_Widget(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QSplitter) DestroyQSplitter() {
	if ptr.Pointer() != nil {
		C.QSplitter_DestroyQSplitter(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
