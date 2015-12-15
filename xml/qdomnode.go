package xml

//#include "xml.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	defer qt.Recovering("QDomNode::QDomNode")

	return NewQDomNodeFromPointer(C.QDomNode_NewQDomNode())
}

func NewQDomNode2(n QDomNode_ITF) *QDomNode {
	defer qt.Recovering("QDomNode::QDomNode")

	return NewQDomNodeFromPointer(C.QDomNode_NewQDomNode2(PointerFromQDomNode(n)))
}

func (ptr *QDomNode) Clear() {
	defer qt.Recovering("QDomNode::clear")

	if ptr.Pointer() != nil {
		C.QDomNode_Clear(ptr.Pointer())
	}
}

func (ptr *QDomNode) ColumnNumber() int {
	defer qt.Recovering("QDomNode::columnNumber")

	if ptr.Pointer() != nil {
		return int(C.QDomNode_ColumnNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNode) HasAttributes() bool {
	defer qt.Recovering("QDomNode::hasAttributes")

	if ptr.Pointer() != nil {
		return C.QDomNode_HasAttributes(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) HasChildNodes() bool {
	defer qt.Recovering("QDomNode::hasChildNodes")

	if ptr.Pointer() != nil {
		return C.QDomNode_HasChildNodes(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsAttr() bool {
	defer qt.Recovering("QDomNode::isAttr")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsAttr(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsCDATASection() bool {
	defer qt.Recovering("QDomNode::isCDATASection")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsCDATASection(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsCharacterData() bool {
	defer qt.Recovering("QDomNode::isCharacterData")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsCharacterData(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsComment() bool {
	defer qt.Recovering("QDomNode::isComment")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsComment(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsDocument() bool {
	defer qt.Recovering("QDomNode::isDocument")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsDocument(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsDocumentFragment() bool {
	defer qt.Recovering("QDomNode::isDocumentFragment")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsDocumentFragment(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsDocumentType() bool {
	defer qt.Recovering("QDomNode::isDocumentType")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsDocumentType(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsElement() bool {
	defer qt.Recovering("QDomNode::isElement")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsElement(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsEntity() bool {
	defer qt.Recovering("QDomNode::isEntity")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsEntity(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsEntityReference() bool {
	defer qt.Recovering("QDomNode::isEntityReference")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsEntityReference(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsNotation() bool {
	defer qt.Recovering("QDomNode::isNotation")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsNotation(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsNull() bool {
	defer qt.Recovering("QDomNode::isNull")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsProcessingInstruction() bool {
	defer qt.Recovering("QDomNode::isProcessingInstruction")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsProcessingInstruction(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) IsSupported(feature string, version string) bool {
	defer qt.Recovering("QDomNode::isSupported")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsSupported(ptr.Pointer(), C.CString(feature), C.CString(version)) != 0
	}
	return false
}

func (ptr *QDomNode) IsText() bool {
	defer qt.Recovering("QDomNode::isText")

	if ptr.Pointer() != nil {
		return C.QDomNode_IsText(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNode) LineNumber() int {
	defer qt.Recovering("QDomNode::lineNumber")

	if ptr.Pointer() != nil {
		return int(C.QDomNode_LineNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNode) LocalName() string {
	defer qt.Recovering("QDomNode::localName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomNode_LocalName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomNode) NamespaceURI() string {
	defer qt.Recovering("QDomNode::namespaceURI")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomNode_NamespaceURI(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomNode) NodeName() string {
	defer qt.Recovering("QDomNode::nodeName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomNode_NodeName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomNode) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomNode::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomNode_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNode) NodeValue() string {
	defer qt.Recovering("QDomNode::nodeValue")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomNode_NodeValue(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomNode) Normalize() {
	defer qt.Recovering("QDomNode::normalize")

	if ptr.Pointer() != nil {
		C.QDomNode_Normalize(ptr.Pointer())
	}
}

func (ptr *QDomNode) Prefix() string {
	defer qt.Recovering("QDomNode::prefix")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomNode_Prefix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomNode) Save(stream core.QTextStream_ITF, indent int, encodingPolicy QDomNode__EncodingPolicy) {
	defer qt.Recovering("QDomNode::save")

	if ptr.Pointer() != nil {
		C.QDomNode_Save(ptr.Pointer(), core.PointerFromQTextStream(stream), C.int(indent), C.int(encodingPolicy))
	}
}

func (ptr *QDomNode) SetNodeValue(v string) {
	defer qt.Recovering("QDomNode::setNodeValue")

	if ptr.Pointer() != nil {
		C.QDomNode_SetNodeValue(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QDomNode) SetPrefix(pre string) {
	defer qt.Recovering("QDomNode::setPrefix")

	if ptr.Pointer() != nil {
		C.QDomNode_SetPrefix(ptr.Pointer(), C.CString(pre))
	}
}

func (ptr *QDomNode) DestroyQDomNode() {
	defer qt.Recovering("QDomNode::~QDomNode")

	if ptr.Pointer() != nil {
		C.QDomNode_DestroyQDomNode(ptr.Pointer())
	}
}
