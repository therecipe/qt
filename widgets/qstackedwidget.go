package widgets

//#include "qstackedwidget.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QStackedWidget struct {
	QFrame
}

type QStackedWidgetITF interface {
	QFrameITF
	QStackedWidgetPTR() *QStackedWidget
}

func PointerFromQStackedWidget(ptr QStackedWidgetITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStackedWidgetPTR().Pointer()
	}
	return nil
}

func QStackedWidgetFromPointer(ptr unsafe.Pointer) *QStackedWidget {
	var n = new(QStackedWidget)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QStackedWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QStackedWidget) QStackedWidgetPTR() *QStackedWidget {
	return ptr
}

func (ptr *QStackedWidget) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QStackedWidget_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStackedWidget) CurrentIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QStackedWidget_CurrentIndex(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStackedWidget) SetCurrentIndex(index int) {
	if ptr.Pointer() != nil {
		C.QStackedWidget_SetCurrentIndex(C.QtObjectPtr(ptr.Pointer()), C.int(index))
	}
}

func (ptr *QStackedWidget) SetCurrentWidget(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QStackedWidget_SetCurrentWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func NewQStackedWidget(parent QWidgetITF) *QStackedWidget {
	return QStackedWidgetFromPointer(unsafe.Pointer(C.QStackedWidget_NewQStackedWidget(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QStackedWidget) AddWidget(widget QWidgetITF) int {
	if ptr.Pointer() != nil {
		return int(C.QStackedWidget_AddWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget))))
	}
	return 0
}

func (ptr *QStackedWidget) ConnectCurrentChanged(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QStackedWidget_ConnectCurrentChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QStackedWidget) DisconnectCurrentChanged() {
	if ptr.Pointer() != nil {
		C.QStackedWidget_DisconnectCurrentChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQStackedWidgetCurrentChanged
func callbackQStackedWidgetCurrentChanged(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "currentChanged").(func(int))(int(index))
}

func (ptr *QStackedWidget) CurrentWidget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QStackedWidget_CurrentWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QStackedWidget) IndexOf(widget QWidgetITF) int {
	if ptr.Pointer() != nil {
		return int(C.QStackedWidget_IndexOf(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget))))
	}
	return 0
}

func (ptr *QStackedWidget) InsertWidget(index int, widget QWidgetITF) int {
	if ptr.Pointer() != nil {
		return int(C.QStackedWidget_InsertWidget(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.QtObjectPtr(PointerFromQWidget(widget))))
	}
	return 0
}

func (ptr *QStackedWidget) RemoveWidget(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QStackedWidget_RemoveWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QStackedWidget) Widget(index int) *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QStackedWidget_Widget(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QStackedWidget) ConnectWidgetRemoved(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QStackedWidget_ConnectWidgetRemoved(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "widgetRemoved", f)
	}
}

func (ptr *QStackedWidget) DisconnectWidgetRemoved() {
	if ptr.Pointer() != nil {
		C.QStackedWidget_DisconnectWidgetRemoved(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "widgetRemoved")
	}
}

//export callbackQStackedWidgetWidgetRemoved
func callbackQStackedWidgetWidgetRemoved(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "widgetRemoved").(func(int))(int(index))
}

func (ptr *QStackedWidget) DestroyQStackedWidget() {
	if ptr.Pointer() != nil {
		C.QStackedWidget_DestroyQStackedWidget(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
