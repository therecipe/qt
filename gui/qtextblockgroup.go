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

type QTextBlockGroupITF interface {
	QTextObjectITF
	QTextBlockGroupPTR() *QTextBlockGroup
}

func PointerFromQTextBlockGroup(ptr QTextBlockGroupITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextBlockGroupPTR().Pointer()
	}
	return nil
}

func QTextBlockGroupFromPointer(ptr unsafe.Pointer) *QTextBlockGroup {
	var n = new(QTextBlockGroup)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTextBlockGroup_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTextBlockGroup) QTextBlockGroupPTR() *QTextBlockGroup {
	return ptr
}
