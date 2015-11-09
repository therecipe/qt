package multimedia

//#include "qmediacontainercontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QMediaContainerControl struct {
	QMediaControl
}

type QMediaContainerControl_ITF interface {
	QMediaControl_ITF
	QMediaContainerControl_PTR() *QMediaContainerControl
}

func PointerFromQMediaContainerControl(ptr QMediaContainerControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaContainerControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaContainerControlFromPointer(ptr unsafe.Pointer) *QMediaContainerControl {
	var n = new(QMediaContainerControl)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QMediaContainerControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaContainerControl) QMediaContainerControl_PTR() *QMediaContainerControl {
	return ptr
}

func (ptr *QMediaContainerControl) ContainerDescription(format string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaContainerControl_ContainerDescription(ptr.Pointer(), C.CString(format)))
	}
	return ""
}

func (ptr *QMediaContainerControl) ContainerFormat() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaContainerControl_ContainerFormat(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaContainerControl) SetContainerFormat(format string) {
	if ptr.Pointer() != nil {
		C.QMediaContainerControl_SetContainerFormat(ptr.Pointer(), C.CString(format))
	}
}

func (ptr *QMediaContainerControl) SupportedContainers() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaContainerControl_SupportedContainers(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QMediaContainerControl) DestroyQMediaContainerControl() {
	if ptr.Pointer() != nil {
		C.QMediaContainerControl_DestroyQMediaContainerControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
