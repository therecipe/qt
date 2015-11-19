package gui

//#include "qtextblockgroup.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTextBlockGroup struct {
	QTextObject
}

type QTextBlockGroup_ITF interface {
	QTextObject_ITF
	QTextBlockGroup_PTR() *QTextBlockGroup
}

func PointerFromQTextBlockGroup(ptr QTextBlockGroup_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextBlockGroup_PTR().Pointer()
	}
	return nil
}

func NewQTextBlockGroupFromPointer(ptr unsafe.Pointer) *QTextBlockGroup {
	var n = new(QTextBlockGroup)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTextBlockGroup_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTextBlockGroup) QTextBlockGroup_PTR() *QTextBlockGroup {
	return ptr
}
