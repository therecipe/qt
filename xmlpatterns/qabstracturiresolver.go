package xmlpatterns

//#include "xmlpatterns.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAbstractUriResolver struct {
	core.QObject
}

type QAbstractUriResolver_ITF interface {
	core.QObject_ITF
	QAbstractUriResolver_PTR() *QAbstractUriResolver
}

func PointerFromQAbstractUriResolver(ptr QAbstractUriResolver_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractUriResolver_PTR().Pointer()
	}
	return nil
}

func NewQAbstractUriResolverFromPointer(ptr unsafe.Pointer) *QAbstractUriResolver {
	var n = new(QAbstractUriResolver)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractUriResolver_") {
		n.SetObjectName("QAbstractUriResolver_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractUriResolver) QAbstractUriResolver_PTR() *QAbstractUriResolver {
	return ptr
}

func (ptr *QAbstractUriResolver) DestroyQAbstractUriResolver() {
	defer qt.Recovering("QAbstractUriResolver::~QAbstractUriResolver")

	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_DestroyQAbstractUriResolver(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
