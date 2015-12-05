package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QSGSimpleTextureNode struct {
	QSGGeometryNode
}

type QSGSimpleTextureNode_ITF interface {
	QSGGeometryNode_ITF
	QSGSimpleTextureNode_PTR() *QSGSimpleTextureNode
}

func PointerFromQSGSimpleTextureNode(ptr QSGSimpleTextureNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGSimpleTextureNode_PTR().Pointer()
	}
	return nil
}

func NewQSGSimpleTextureNodeFromPointer(ptr unsafe.Pointer) *QSGSimpleTextureNode {
	var n = new(QSGSimpleTextureNode)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGSimpleTextureNode) QSGSimpleTextureNode_PTR() *QSGSimpleTextureNode {
	return ptr
}

//QSGSimpleTextureNode::TextureCoordinatesTransformFlag
type QSGSimpleTextureNode__TextureCoordinatesTransformFlag int64

const (
	QSGSimpleTextureNode__NoTransform        = QSGSimpleTextureNode__TextureCoordinatesTransformFlag(0x00)
	QSGSimpleTextureNode__MirrorHorizontally = QSGSimpleTextureNode__TextureCoordinatesTransformFlag(0x01)
	QSGSimpleTextureNode__MirrorVertically   = QSGSimpleTextureNode__TextureCoordinatesTransformFlag(0x02)
)

func NewQSGSimpleTextureNode() *QSGSimpleTextureNode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGSimpleTextureNode::QSGSimpleTextureNode")
		}
	}()

	return NewQSGSimpleTextureNodeFromPointer(C.QSGSimpleTextureNode_NewQSGSimpleTextureNode())
}

func (ptr *QSGSimpleTextureNode) Filtering() QSGTexture__Filtering {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGSimpleTextureNode::filtering")
		}
	}()

	if ptr.Pointer() != nil {
		return QSGTexture__Filtering(C.QSGSimpleTextureNode_Filtering(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGSimpleTextureNode) OwnsTexture() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGSimpleTextureNode::ownsTexture")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSGSimpleTextureNode_OwnsTexture(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGSimpleTextureNode) SetFiltering(filtering QSGTexture__Filtering) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGSimpleTextureNode::setFiltering")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetFiltering(ptr.Pointer(), C.int(filtering))
	}
}

func (ptr *QSGSimpleTextureNode) SetOwnsTexture(owns bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGSimpleTextureNode::setOwnsTexture")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetOwnsTexture(ptr.Pointer(), C.int(qt.GoBoolToInt(owns)))
	}
}

func (ptr *QSGSimpleTextureNode) SetRect(r core.QRectF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGSimpleTextureNode::setRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetRect(ptr.Pointer(), core.PointerFromQRectF(r))
	}
}

func (ptr *QSGSimpleTextureNode) SetRect2(x float64, y float64, w float64, h float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGSimpleTextureNode::setRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetRect2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h))
	}
}

func (ptr *QSGSimpleTextureNode) SetSourceRect(r core.QRectF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGSimpleTextureNode::setSourceRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetSourceRect(ptr.Pointer(), core.PointerFromQRectF(r))
	}
}

func (ptr *QSGSimpleTextureNode) SetSourceRect2(x float64, y float64, w float64, h float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGSimpleTextureNode::setSourceRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetSourceRect2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h))
	}
}

func (ptr *QSGSimpleTextureNode) SetTexture(texture QSGTexture_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGSimpleTextureNode::setTexture")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetTexture(ptr.Pointer(), PointerFromQSGTexture(texture))
	}
}

func (ptr *QSGSimpleTextureNode) SetTextureCoordinatesTransform(mode QSGSimpleTextureNode__TextureCoordinatesTransformFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGSimpleTextureNode::setTextureCoordinatesTransform")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetTextureCoordinatesTransform(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QSGSimpleTextureNode) Texture() *QSGTexture {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGSimpleTextureNode::texture")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QSGSimpleTextureNode_Texture(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGSimpleTextureNode) TextureCoordinatesTransform() QSGSimpleTextureNode__TextureCoordinatesTransformFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGSimpleTextureNode::textureCoordinatesTransform")
		}
	}()

	if ptr.Pointer() != nil {
		return QSGSimpleTextureNode__TextureCoordinatesTransformFlag(C.QSGSimpleTextureNode_TextureCoordinatesTransform(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGSimpleTextureNode) DestroyQSGSimpleTextureNode() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGSimpleTextureNode::~QSGSimpleTextureNode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_DestroyQSGSimpleTextureNode(ptr.Pointer())
	}
}
