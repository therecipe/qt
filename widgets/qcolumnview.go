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

type QColumnViewITF interface {
	QAbstractItemViewITF
	QColumnViewPTR() *QColumnView
}

func PointerFromQColumnView(ptr QColumnViewITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QColumnViewPTR().Pointer()
	}
	return nil
}

func QColumnViewFromPointer(ptr unsafe.Pointer) *QColumnView {
	var n = new(QColumnView)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QColumnView_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QColumnView) QColumnViewPTR() *QColumnView {
	return ptr
}

func (ptr *QColumnView) ResizeGripsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QColumnView_ResizeGripsVisible(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QColumnView) SetResizeGripsVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QColumnView_SetResizeGripsVisible(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(visible)))
	}
}

func NewQColumnView(parent QWidgetITF) *QColumnView {
	return QColumnViewFromPointer(unsafe.Pointer(C.QColumnView_NewQColumnView(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QColumnView) IndexAt(point core.QPointITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.QModelIndexFromPointer(unsafe.Pointer(C.QColumnView_IndexAt(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(point)))))
	}
	return nil
}

func (ptr *QColumnView) PreviewWidget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QColumnView_PreviewWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QColumnView) ScrollTo(index core.QModelIndexITF, hint QAbstractItemView__ScrollHint) {
	if ptr.Pointer() != nil {
		C.QColumnView_ScrollTo(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)), C.int(hint))
	}
}

func (ptr *QColumnView) SelectAll() {
	if ptr.Pointer() != nil {
		C.QColumnView_SelectAll(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QColumnView) SetModel(model core.QAbstractItemModelITF) {
	if ptr.Pointer() != nil {
		C.QColumnView_SetModel(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQAbstractItemModel(model)))
	}
}

func (ptr *QColumnView) SetPreviewWidget(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QColumnView_SetPreviewWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QColumnView) SetRootIndex(index core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QColumnView_SetRootIndex(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)))
	}
}

func (ptr *QColumnView) SetSelectionModel(newSelectionModel core.QItemSelectionModelITF) {
	if ptr.Pointer() != nil {
		C.QColumnView_SetSelectionModel(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQItemSelectionModel(newSelectionModel)))
	}
}

func (ptr *QColumnView) ConnectUpdatePreviewWidget(f func(index core.QModelIndexITF)) {
	if ptr.Pointer() != nil {
		C.QColumnView_ConnectUpdatePreviewWidget(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "updatePreviewWidget", f)
	}
}

func (ptr *QColumnView) DisconnectUpdatePreviewWidget() {
	if ptr.Pointer() != nil {
		C.QColumnView_DisconnectUpdatePreviewWidget(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "updatePreviewWidget")
	}
}

//export callbackQColumnViewUpdatePreviewWidget
func callbackQColumnViewUpdatePreviewWidget(ptrName *C.char, index unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "updatePreviewWidget").(func(*core.QModelIndex))(core.QModelIndexFromPointer(index))
}

func (ptr *QColumnView) DestroyQColumnView() {
	if ptr.Pointer() != nil {
		C.QColumnView_DestroyQColumnView(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
