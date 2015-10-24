package xmlpatterns

//#include "qabstracturiresolver.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAbstractUriResolver struct {
	core.QObject
}

type QAbstractUriResolverITF interface {
	core.QObjectITF
	QAbstractUriResolverPTR() *QAbstractUriResolver
}

func PointerFromQAbstractUriResolver(ptr QAbstractUriResolverITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractUriResolverPTR().Pointer()
	}
	return nil
}

func QAbstractUriResolverFromPointer(ptr unsafe.Pointer) *QAbstractUriResolver {
	var n = new(QAbstractUriResolver)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAbstractUriResolver_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractUriResolver) QAbstractUriResolverPTR() *QAbstractUriResolver {
	return ptr
}

func (ptr *QAbstractUriResolver) Resolve(relative string, baseURI string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractUriResolver_Resolve(C.QtObjectPtr(ptr.Pointer()), C.CString(relative), C.CString(baseURI)))
	}
	return ""
}

func (ptr *QAbstractUriResolver) DestroyQAbstractUriResolver() {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_DestroyQAbstractUriResolver(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
