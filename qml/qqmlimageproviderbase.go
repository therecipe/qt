package qml

//#include "qqmlimageproviderbase.h"
import "C"
import (
	"unsafe"
)

type QQmlImageProviderBase struct {
	ptr unsafe.Pointer
}

type QQmlImageProviderBaseITF interface {
	QQmlImageProviderBasePTR() *QQmlImageProviderBase
}

func (p *QQmlImageProviderBase) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlImageProviderBase) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlImageProviderBase(ptr QQmlImageProviderBaseITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlImageProviderBasePTR().Pointer()
	}
	return nil
}

func QQmlImageProviderBaseFromPointer(ptr unsafe.Pointer) *QQmlImageProviderBase {
	var n = new(QQmlImageProviderBase)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlImageProviderBase) QQmlImageProviderBasePTR() *QQmlImageProviderBase {
	return ptr
}

//QQmlImageProviderBase::Flag
type QQmlImageProviderBase__Flag int

var (
	QQmlImageProviderBase__ForceAsynchronousImageLoading = QQmlImageProviderBase__Flag(0x01)
)

//QQmlImageProviderBase::ImageType
type QQmlImageProviderBase__ImageType int

var (
	QQmlImageProviderBase__Image   = QQmlImageProviderBase__ImageType(0)
	QQmlImageProviderBase__Pixmap  = QQmlImageProviderBase__ImageType(1)
	QQmlImageProviderBase__Texture = QQmlImageProviderBase__ImageType(2)
	QQmlImageProviderBase__Invalid = QQmlImageProviderBase__ImageType(3)
)

func (ptr *QQmlImageProviderBase) Flags() QQmlImageProviderBase__Flag {
	if ptr.Pointer() != nil {
		return QQmlImageProviderBase__Flag(C.QQmlImageProviderBase_Flags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlImageProviderBase) ImageType() QQmlImageProviderBase__ImageType {
	if ptr.Pointer() != nil {
		return QQmlImageProviderBase__ImageType(C.QQmlImageProviderBase_ImageType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
