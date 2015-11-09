package xml

//#include "qdomcomment.h"
import "C"
import (
	"unsafe"
)

type QDomComment struct {
	QDomCharacterData
}

type QDomComment_ITF interface {
	QDomCharacterData_ITF
	QDomComment_PTR() *QDomComment
}

func PointerFromQDomComment(ptr QDomComment_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomComment_PTR().Pointer()
	}
	return nil
}

func NewQDomCommentFromPointer(ptr unsafe.Pointer) *QDomComment {
	var n = new(QDomComment)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomComment) QDomComment_PTR() *QDomComment {
	return ptr
}

func NewQDomComment() *QDomComment {
	return NewQDomCommentFromPointer(C.QDomComment_NewQDomComment())
}

func NewQDomComment2(x QDomComment_ITF) *QDomComment {
	return NewQDomCommentFromPointer(C.QDomComment_NewQDomComment2(PointerFromQDomComment(x)))
}

func (ptr *QDomComment) NodeType() QDomNode__NodeType {
	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomComment_NodeType(ptr.Pointer()))
	}
	return 0
}
