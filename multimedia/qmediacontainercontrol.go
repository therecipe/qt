package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
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
	for len(n.ObjectName()) < len("QMediaContainerControl_") {
		n.SetObjectName("QMediaContainerControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaContainerControl) QMediaContainerControl_PTR() *QMediaContainerControl {
	return ptr
}

func (ptr *QMediaContainerControl) ContainerDescription(format string) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaContainerControl::containerDescription")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaContainerControl_ContainerDescription(ptr.Pointer(), C.CString(format)))
	}
	return ""
}

func (ptr *QMediaContainerControl) ContainerFormat() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaContainerControl::containerFormat")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaContainerControl_ContainerFormat(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaContainerControl) SetContainerFormat(format string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaContainerControl::setContainerFormat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaContainerControl_SetContainerFormat(ptr.Pointer(), C.CString(format))
	}
}

func (ptr *QMediaContainerControl) SupportedContainers() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaContainerControl::supportedContainers")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaContainerControl_SupportedContainers(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QMediaContainerControl) DestroyQMediaContainerControl() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaContainerControl::~QMediaContainerControl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaContainerControl_DestroyQMediaContainerControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
