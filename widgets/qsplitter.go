package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QSplitter struct {
	QFrame
}

type QSplitter_ITF interface {
	QFrame_ITF
	QSplitter_PTR() *QSplitter
}

func PointerFromQSplitter(ptr QSplitter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSplitter_PTR().Pointer()
	}
	return nil
}

func NewQSplitterFromPointer(ptr unsafe.Pointer) *QSplitter {
	var n = new(QSplitter)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSplitter_") {
		n.SetObjectName("QSplitter_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSplitter) QSplitter_PTR() *QSplitter {
	return ptr
}

func (ptr *QSplitter) ChildrenCollapsible() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitter::childrenCollapsible")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSplitter_ChildrenCollapsible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSplitter) HandleWidth() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitter::handleWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSplitter_HandleWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSplitter) IndexOf(widget QWidget_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitter::indexOf")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSplitter_IndexOf(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QSplitter) OpaqueResize() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitter::opaqueResize")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSplitter_OpaqueResize(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSplitter) Orientation() core.Qt__Orientation {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitter::orientation")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QSplitter_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSplitter) SetChildrenCollapsible(v bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitter::setChildrenCollapsible")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSplitter_SetChildrenCollapsible(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QSplitter) SetHandleWidth(v int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitter::setHandleWidth")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSplitter_SetHandleWidth(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QSplitter) SetOpaqueResize(opaque bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitter::setOpaqueResize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSplitter_SetOpaqueResize(ptr.Pointer(), C.int(qt.GoBoolToInt(opaque)))
	}
}

func (ptr *QSplitter) SetOrientation(v core.Qt__Orientation) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitter::setOrientation")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSplitter_SetOrientation(ptr.Pointer(), C.int(v))
	}
}

func NewQSplitter(parent QWidget_ITF) *QSplitter {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitter::QSplitter")
		}
	}()

	return NewQSplitterFromPointer(C.QSplitter_NewQSplitter(PointerFromQWidget(parent)))
}

func NewQSplitter2(orientation core.Qt__Orientation, parent QWidget_ITF) *QSplitter {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitter::QSplitter")
		}
	}()

	return NewQSplitterFromPointer(C.QSplitter_NewQSplitter2(C.int(orientation), PointerFromQWidget(parent)))
}

func (ptr *QSplitter) AddWidget(widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitter::addWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSplitter_AddWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QSplitter) Count() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitter::count")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSplitter_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSplitter) GetRange(index int, min int, max int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitter::getRange")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSplitter_GetRange(ptr.Pointer(), C.int(index), C.int(min), C.int(max))
	}
}

func (ptr *QSplitter) Handle(index int) *QSplitterHandle {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitter::handle")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQSplitterHandleFromPointer(C.QSplitter_Handle(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QSplitter) InsertWidget(index int, widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitter::insertWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSplitter_InsertWidget(ptr.Pointer(), C.int(index), PointerFromQWidget(widget))
	}
}

func (ptr *QSplitter) IsCollapsible(index int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitter::isCollapsible")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSplitter_IsCollapsible(ptr.Pointer(), C.int(index)) != 0
	}
	return false
}

func (ptr *QSplitter) Refresh() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitter::refresh")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSplitter_Refresh(ptr.Pointer())
	}
}

func (ptr *QSplitter) RestoreState(state core.QByteArray_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitter::restoreState")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSplitter_RestoreState(ptr.Pointer(), core.PointerFromQByteArray(state)) != 0
	}
	return false
}

func (ptr *QSplitter) SaveState() *core.QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitter::saveState")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSplitter_SaveState(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSplitter) SetCollapsible(index int, collapse bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitter::setCollapsible")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSplitter_SetCollapsible(ptr.Pointer(), C.int(index), C.int(qt.GoBoolToInt(collapse)))
	}
}

func (ptr *QSplitter) SetStretchFactor(index int, stretch int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitter::setStretchFactor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSplitter_SetStretchFactor(ptr.Pointer(), C.int(index), C.int(stretch))
	}
}

func (ptr *QSplitter) ConnectSplitterMoved(f func(pos int, index int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitter::splitterMoved")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSplitter_ConnectSplitterMoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "splitterMoved", f)
	}
}

func (ptr *QSplitter) DisconnectSplitterMoved() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitter::splitterMoved")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSplitter_DisconnectSplitterMoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "splitterMoved")
	}
}

//export callbackQSplitterSplitterMoved
func callbackQSplitterSplitterMoved(ptrName *C.char, pos C.int, index C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitter::splitterMoved")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "splitterMoved").(func(int, int))(int(pos), int(index))
}

func (ptr *QSplitter) Widget(index int) *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitter::widget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QSplitter_Widget(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QSplitter) DestroyQSplitter() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitter::~QSplitter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSplitter_DestroyQSplitter(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
