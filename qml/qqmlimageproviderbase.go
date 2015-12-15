package qml

//#include "qml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QQmlImageProviderBase struct {
	ptr unsafe.Pointer
}

type QQmlImageProviderBase_ITF interface {
	QQmlImageProviderBase_PTR() *QQmlImageProviderBase
}

func (p *QQmlImageProviderBase) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlImageProviderBase) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlImageProviderBase(ptr QQmlImageProviderBase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlImageProviderBase_PTR().Pointer()
	}
	return nil
}

func NewQQmlImageProviderBaseFromPointer(ptr unsafe.Pointer) *QQmlImageProviderBase {
	var n = new(QQmlImageProviderBase)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlImageProviderBase) QQmlImageProviderBase_PTR() *QQmlImageProviderBase {
	return ptr
}

//QQmlImageProviderBase::Flag
type QQmlImageProviderBase__Flag int64

const (
	QQmlImageProviderBase__ForceAsynchronousImageLoading = QQmlImageProviderBase__Flag(0x01)
)

//QQmlImageProviderBase::ImageType
type QQmlImageProviderBase__ImageType int64

const (
	QQmlImageProviderBase__Image   = QQmlImageProviderBase__ImageType(0)
	QQmlImageProviderBase__Pixmap  = QQmlImageProviderBase__ImageType(1)
	QQmlImageProviderBase__Texture = QQmlImageProviderBase__ImageType(2)
	QQmlImageProviderBase__Invalid = QQmlImageProviderBase__ImageType(3)
)

func (ptr *QQmlImageProviderBase) Flags() QQmlImageProviderBase__Flag {
	defer qt.Recovering("QQmlImageProviderBase::flags")

	if ptr.Pointer() != nil {
		return QQmlImageProviderBase__Flag(C.QQmlImageProviderBase_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlImageProviderBase) ImageType() QQmlImageProviderBase__ImageType {
	defer qt.Recovering("QQmlImageProviderBase::imageType")

	if ptr.Pointer() != nil {
		return QQmlImageProviderBase__ImageType(C.QQmlImageProviderBase_ImageType(ptr.Pointer()))
	}
	return 0
}
