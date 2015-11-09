package xml

//#include "qdomnode.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDomNode struct {
	ptr unsafe.Pointer
}

type QDomNode_ITF interface {
	QDomNode_PTR() *QDomNode
}

func (p *QDomNode) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDomNode) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDomNode(ptr QDomNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNode_PTR().Pointer()
	}
	return nil
}

func NewQDomNodeFromPointer(ptr unsafe.Pointer) *QDomNode {
	var n = new(QDomNode)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomNode) QDomNode_PTR() *QDomNode {
	return ptr
}

//QDomNode::EncodingPolicy
type QDomNode__EncodingPolicy int64

const (
	QDomNode__EncodingFromDocument   = QDomNode__EncodingPolicy(1)
	QDomNode__EncodingFromTextStream = QDomNode__EncodingPolicy(2)
)

//QDomNode::NodeType
type QDomNode__NodeType int64

const (
	QDomNode__ElementNode               = QDomNode__NodeType(1)
	QDomNode__AttributeNode             = QDomNode__NodeType(2)
	QDomNode__TextNode                  = QDomNode__NodeType(3)
	QDomNode__CDATASectionNode          = QDomNode__NodeType(4)
	QDomNode__EntityReferenceNode       = QDomNode__NodeType(5)
	QDomNode__EntityNode                = QDomNode__NodeType(6)
	QDomNode__ProcessingInstructionNode = QDomNode__NodeType(7)
	QDomNode__CommentNode               = QDomNode__NodeType(8)
	QDomNode__DocumentNode              = QDomNode__NodeType(9)
	QDomNode__DocumentTypeNode          = QDomNode__NodeType(10)
	QDomNode__DocumentFragmentNode      = QDomNode__NodeType(11)
	QDomNode__NotationNode              = QDomNode__NodeType(12)
	QDomNode__BaseNode                  = QDomNode__NodeType(21)
	QDomNode__CharacterDataNode         = QDomNode__NodeType(22)
)

func NewQDomNode() *QDomNode {
	return NewQDomNodeFromPointer(C.QDomNode_NewQDomNode())
}

func NewQDomNode2(n QDomNode_ITF) *QDomNode {
	return NewQDomNodeFromPointer(C.QDomNode_NewQDomNode2(PointerFromQDomNode(n)))
}

func (ptr *QDomNode) Clear() {
	if ptr.Pointer() != nil {
		C.QDomNode_Clear(ptr.Pointer())
	}
}

func (ptr *QDomNode) ColumnNumber() int {
	if ptr.Pointer() != nil {
		return int(C.QDomNode_ColumnNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNode) HasAttributes() bool {
	if ptr.Pointer() != nil {
		return C.QDomNode_HasAttributes(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) HasChildNodes() bool {
	if ptr.Pointer() != nil {
		return C.QDomNode_HasChildNodes(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsAttr() bool {
	if ptr.Pointer() != nil {
		return C.QDomNode_IsAttr(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsCDATASection() bool {
	if ptr.Pointer() != nil {
		return C.QDomNode_IsCDATASection(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsCharacterData() bool {
	if ptr.Pointer() != nil {
		return C.QDomNode_IsCharacterData(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsComment() bool {
	if ptr.Pointer() != nil {
		return C.QDomNode_IsComment(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsDocument() bool {
	if ptr.Pointer() != nil {
		return C.QDomNode_IsDocument(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsDocumentFragment() bool {
	if ptr.Pointer() != nil {
		return C.QDomNode_IsDocumentFragment(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsDocumentType() bool {
	if ptr.Pointer() != nil {
		return C.QDomNode_IsDocumentType(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsElement() bool {
	if ptr.Pointer() != nil {
		return C.QDomNode_IsElement(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsEntity() bool {
	if ptr.Pointer() != nil {
		return C.QDomNode_IsEntity(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsEntityReference() bool {
	if ptr.Pointer() != nil {
		return C.QDomNode_IsEntityReference(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsNotation() bool {
	if ptr.Pointer() != nil {
		return C.QDomNode_IsNotation(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QDomNode_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsProcessingInstruction() bool {
	if ptr.Pointer() != nil {
		return C.QDomNode_IsProcessingInstruction(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsSupported(feature string, version string) bool {
	if ptr.Pointer() != nil {
		return C.QDomNode_IsSupported(ptr.Pointer(), C.CString(feature), C.CString(version)) != 0
	}
	return false
}

func (ptr *QDomNode) IsText() bool {
	if ptr.Pointer() != nil {
		return C.QDomNode_IsText(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) LineNumber() int {
	if ptr.Pointer() != nil {
		return int(C.QDomNode_LineNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNode) LocalName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomNode_LocalName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomNode) NamespaceURI() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomNode_NamespaceURI(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomNode) NodeName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomNode_NodeName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomNode) NodeType() QDomNode__NodeType {
	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomNode_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNode) NodeValue() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomNode_NodeValue(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomNode) Normalize() {
	if ptr.Pointer() != nil {
		C.QDomNode_Normalize(ptr.Pointer())
	}
}

func (ptr *QDomNode) Prefix() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomNode_Prefix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomNode) Save(stream core.QTextStream_ITF, indent int, encodingPolicy QDomNode__EncodingPolicy) {
	if ptr.Pointer() != nil {
		C.QDomNode_Save(ptr.Pointer(), core.PointerFromQTextStream(stream), C.int(indent), C.int(encodingPolicy))
	}
}

func (ptr *QDomNode) SetNodeValue(v string) {
	if ptr.Pointer() != nil {
		C.QDomNode_SetNodeValue(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QDomNode) SetPrefix(pre string) {
	if ptr.Pointer() != nil {
		C.QDomNode_SetPrefix(ptr.Pointer(), C.CString(pre))
	}
}

func (ptr *QDomNode) DestroyQDomNode() {
	if ptr.Pointer() != nil {
		C.QDomNode_DestroyQDomNode(ptr.Pointer())
	}
}
