package widgets

//#include "qcolumnview.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QColumnView struct {
	QAbstractItemView
}

type QColumnView_ITF interface {
	QAbstractItemView_ITF
	QColumnView_PTR() *QColumnView
}

func PointerFromQColumnView(ptr QColumnView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QColumnView_PTR().Pointer()
	}
	return nil
}

func NewQColumnViewFromPointer(ptr unsafe.Pointer) *QColumnView {
	var n = new(QColumnView)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QColumnView_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QColumnView) QColumnView_PTR() *QColumnView {
	return ptr
}

func (ptr *QColumnView) ResizeGripsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QColumnView_ResizeGripsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QColumnView) SetResizeGripsVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QColumnView_SetResizeGripsVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func NewQColumnView(parent QWidget_ITF) *QColumnView {
	return NewQColumnViewFromPointer(C.QColumnView_NewQColumnView(PointerFromQWidget(parent)))
}

func (ptr *QColumnView) IndexAt(point core.QPoint_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QColumnView_IndexAt(ptr.Pointer(), core.PointerFromQPoint(point)))
	}
	return nil
}

func (ptr *QColumnView) PreviewWidget() *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QColumnView_PreviewWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QColumnView) ScrollTo(index core.QModelIndex_ITF, hint QAbstractItemView__ScrollHint) {
	if ptr.Pointer() != nil {
		C.QColumnView_ScrollTo(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(hint))
	}
}

func (ptr *QColumnView) SelectAll() {
	if ptr.Pointer() != nil {
		C.QColumnView_SelectAll(ptr.Pointer())
	}
}

func (ptr *QColumnView) SetModel(model core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QColumnView_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QColumnView) SetPreviewWidget(widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QColumnView_SetPreviewWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QColumnView) SetRootIndex(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QColumnView_SetRootIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QColumnView) SetSelectionModel(newSelectionModel core.QItemSelectionModel_ITF) {
	if ptr.Pointer() != nil {
		C.QColumnView_SetSelectionModel(ptr.Pointer(), core.PointerFromQItemSelectionModel(newSelectionModel))
	}
}

func (ptr *QColumnView) ConnectUpdatePreviewWidget(f func(index *core.QModelIndex)) {
	if ptr.Pointer() != nil {
		C.QColumnView_ConnectUpdatePreviewWidget(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "updatePreviewWidget", f)
	}
}

func (ptr *QColumnView) DisconnectUpdatePreviewWidget() {
	if ptr.Pointer() != nil {
		C.QColumnView_DisconnectUpdatePreviewWidget(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "updatePreviewWidget")
	}
}

//export callbackQColumnViewUpdatePreviewWidget
func callbackQColumnViewUpdatePreviewWidget(ptrName *C.char, index unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "updatePreviewWidget").(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
}

func (ptr *QColumnView) DestroyQColumnView() {
	if ptr.Pointer() != nil {
		C.QColumnView_DestroyQColumnView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
