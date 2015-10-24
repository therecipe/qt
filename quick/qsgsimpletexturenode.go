package quick

//#include "qsgsimpletexturenode.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSGSimpleTextureNode struct {
	QSGGeometryNode
}

type QSGSimpleTextureNodeITF interface {
	QSGGeometryNodeITF
	QSGSimpleTextureNodePTR() *QSGSimpleTextureNode
}

func PointerFromQSGSimpleTextureNode(ptr QSGSimpleTextureNodeITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGSimpleTextureNodePTR().Pointer()
	}
	return nil
}

func QSGSimpleTextureNodeFromPointer(ptr unsafe.Pointer) *QSGSimpleTextureNode {
	var n = new(QSGSimpleTextureNode)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGSimpleTextureNode) QSGSimpleTextureNodePTR() *QSGSimpleTextureNode {
	return ptr
}

//QSGSimpleTextureNode::TextureCoordinatesTransformFlag
type QSGSimpleTextureNode__TextureCoordinatesTransformFlag int

var (
	QSGSimpleTextureNode__NoTransform        = QSGSimpleTextureNode__TextureCoordinatesTransformFlag(0x00)
	QSGSimpleTextureNode__MirrorHorizontally = QSGSimpleTextureNode__TextureCoordinatesTransformFlag(0x01)
	QSGSimpleTextureNode__MirrorVertically   = QSGSimpleTextureNode__TextureCoordinatesTransformFlag(0x02)
)

func NewQSGSimpleTextureNode() *QSGSimpleTextureNode {
	return QSGSimpleTextureNodeFromPointer(unsafe.Pointer(C.QSGSimpleTextureNode_NewQSGSimpleTextureNode()))
}

func (ptr *QSGSimpleTextureNode) Filtering() QSGTexture__Filtering {
	if ptr.Pointer() != nil {
		return QSGTexture__Filtering(C.QSGSimpleTextureNode_Filtering(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGSimpleTextureNode) OwnsTexture() bool {
	if ptr.Pointer() != nil {
		return C.QSGSimpleTextureNode_OwnsTexture(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSGSimpleTextureNode) SetFiltering(filtering QSGTexture__Filtering) {
	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetFiltering(C.QtObjectPtr(ptr.Pointer()), C.int(filtering))
	}
}

func (ptr *QSGSimpleTextureNode) SetOwnsTexture(owns bool) {
	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetOwnsTexture(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(owns)))
	}
}

func (ptr *QSGSimpleTextureNode) SetRect(r core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetRect(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(r)))
	}
}

func (ptr *QSGSimpleTextureNode) SetSourceRect(r core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetSourceRect(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(r)))
	}
}

func (ptr *QSGSimpleTextureNode) SetTexture(texture QSGTextureITF) {
	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetTexture(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSGTexture(texture)))
	}
}

func (ptr *QSGSimpleTextureNode) SetTextureCoordinatesTransform(mode QSGSimpleTextureNode__TextureCoordinatesTransformFlag) {
	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetTextureCoordinatesTransform(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QSGSimpleTextureNode) Texture() *QSGTexture {
	if ptr.Pointer() != nil {
		return QSGTextureFromPointer(unsafe.Pointer(C.QSGSimpleTextureNode_Texture(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSGSimpleTextureNode) TextureCoordinatesTransform() QSGSimpleTextureNode__TextureCoordinatesTransformFlag {
	if ptr.Pointer() != nil {
		return QSGSimpleTextureNode__TextureCoordinatesTransformFlag(C.QSGSimpleTextureNode_TextureCoordinatesTransform(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGSimpleTextureNode) DestroyQSGSimpleTextureNode() {
	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_DestroyQSGSimpleTextureNode(C.QtObjectPtr(ptr.Pointer()))
	}
}
