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

type QMediaContainerControlITF interface {
	QMediaControlITF
	QMediaContainerControlPTR() *QMediaContainerControl
}

func PointerFromQMediaContainerControl(ptr QMediaContainerControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaContainerControlPTR().Pointer()
	}
	return nil
}

func QMediaContainerControlFromPointer(ptr unsafe.Pointer) *QMediaContainerControl {
	var n = new(QMediaContainerControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMediaContainerControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaContainerControl) QMediaContainerControlPTR() *QMediaContainerControl {
	return ptr
}

func (ptr *QMediaContainerControl) ContainerDescription(format string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaContainerControl_ContainerDescription(C.QtObjectPtr(ptr.Pointer()), C.CString(format)))
	}
	return ""
}

func (ptr *QMediaContainerControl) ContainerFormat() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaContainerControl_ContainerFormat(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QMediaContainerControl) SetContainerFormat(format string) {
	if ptr.Pointer() != nil {
		C.QMediaContainerControl_SetContainerFormat(C.QtObjectPtr(ptr.Pointer()), C.CString(format))
	}
}

func (ptr *QMediaContainerControl) SupportedContainers() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaContainerControl_SupportedContainers(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QMediaContainerControl) DestroyQMediaContainerControl() {
	if ptr.Pointer() != nil {
		C.QMediaContainerControl_DestroyQMediaContainerControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
