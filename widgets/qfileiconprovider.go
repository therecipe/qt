package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QFileIconProvider struct {
	ptr unsafe.Pointer
}

type QFileIconProvider_ITF interface {
	QFileIconProvider_PTR() *QFileIconProvider
}

func (p *QFileIconProvider) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QFileIconProvider) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQFileIconProvider(ptr QFileIconProvider_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFileIconProvider_PTR().Pointer()
	}
	return nil
}

func NewQFileIconProviderFromPointer(ptr unsafe.Pointer) *QFileIconProvider {
	var n = new(QFileIconProvider)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QFileIconProvider_") {
		n.SetObjectNameAbs("QFileIconProvider_" + qt.Identifier())
	}
	return n
}

func (ptr *QFileIconProvider) QFileIconProvider_PTR() *QFileIconProvider {
	return ptr
}

//QFileIconProvider::IconType
type QFileIconProvider__IconType int64

const (
	QFileIconProvider__Computer = QFileIconProvider__IconType(0)
	QFileIconProvider__Desktop  = QFileIconProvider__IconType(1)
	QFileIconProvider__Trashcan = QFileIconProvider__IconType(2)
	QFileIconProvider__Network  = QFileIconProvider__IconType(3)
	QFileIconProvider__Drive    = QFileIconProvider__IconType(4)
	QFileIconProvider__Folder   = QFileIconProvider__IconType(5)
	QFileIconProvider__File     = QFileIconProvider__IconType(6)
)

//QFileIconProvider::Option
type QFileIconProvider__Option int64

const (
	QFileIconProvider__DontUseCustomDirectoryIcons = QFileIconProvider__Option(0x00000001)
)

func NewQFileIconProvider() *QFileIconProvider {
	defer qt.Recovering("QFileIconProvider::QFileIconProvider")

	return NewQFileIconProviderFromPointer(C.QFileIconProvider_NewQFileIconProvider())
}

func (ptr *QFileIconProvider) Options() QFileIconProvider__Option {
	defer qt.Recovering("QFileIconProvider::options")

	if ptr.Pointer() != nil {
		return QFileIconProvider__Option(C.QFileIconProvider_Options(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFileIconProvider) SetOptions(options QFileIconProvider__Option) {
	defer qt.Recovering("QFileIconProvider::setOptions")

	if ptr.Pointer() != nil {
		C.QFileIconProvider_SetOptions(ptr.Pointer(), C.int(options))
	}
}

func (ptr *QFileIconProvider) Type(info core.QFileInfo_ITF) string {
	defer qt.Recovering("QFileIconProvider::type")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileIconProvider_Type(ptr.Pointer(), core.PointerFromQFileInfo(info)))
	}
	return ""
}

func (ptr *QFileIconProvider) DestroyQFileIconProvider() {
	defer qt.Recovering("QFileIconProvider::~QFileIconProvider")

	if ptr.Pointer() != nil {
		C.QFileIconProvider_DestroyQFileIconProvider(ptr.Pointer())
	}
}

func (ptr *QFileIconProvider) ObjectNameAbs() string {
	defer qt.Recovering("QFileIconProvider::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileIconProvider_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileIconProvider) SetObjectNameAbs(name string) {
	defer qt.Recovering("QFileIconProvider::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QFileIconProvider_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
