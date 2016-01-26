package xml

//#include "xml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QDomAttr struct {
	QDomNode
}

type QDomAttr_ITF interface {
	QDomNode_ITF
	QDomAttr_PTR() *QDomAttr
}

func PointerFromQDomAttr(ptr QDomAttr_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomAttr_PTR().Pointer()
	}
	return nil
}

func NewQDomAttrFromPointer(ptr unsafe.Pointer) *QDomAttr {
	var n = new(QDomAttr)
	n.SetPointer(ptr)
	return n
}

func newQDomAttrFromPointer(ptr unsafe.Pointer) *QDomAttr {
	var n = NewQDomAttrFromPointer(ptr)
	return n
}

func (ptr *QDomAttr) QDomAttr_PTR() *QDomAttr {
	return ptr
}

func NewQDomAttr() *QDomAttr {
	defer qt.Recovering("QDomAttr::QDomAttr")

	return newQDomAttrFromPointer(C.QDomAttr_NewQDomAttr())
}

func NewQDomAttr2(x QDomAttr_ITF) *QDomAttr {
	defer qt.Recovering("QDomAttr::QDomAttr")

	return newQDomAttrFromPointer(C.QDomAttr_NewQDomAttr2(PointerFromQDomAttr(x)))
}

func (ptr *QDomAttr) Name() string {
	defer qt.Recovering("QDomAttr::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomAttr_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomAttr) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomAttr::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomAttr_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomAttr) SetValue(v string) {
	defer qt.Recovering("QDomAttr::setValue")

	if ptr.Pointer() != nil {
		C.QDomAttr_SetValue(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QDomAttr) Specified() bool {
	defer qt.Recovering("QDomAttr::specified")

	if ptr.Pointer() != nil {
		return C.QDomAttr_Specified(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomAttr) Value() string {
	defer qt.Recovering("QDomAttr::value")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomAttr_Value(ptr.Pointer()))
	}
	return ""
}

type QDomCDATASection struct {
	QDomText
}

type QDomCDATASection_ITF interface {
	QDomText_ITF
	QDomCDATASection_PTR() *QDomCDATASection
}

func PointerFromQDomCDATASection(ptr QDomCDATASection_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomCDATASection_PTR().Pointer()
	}
	return nil
}

func NewQDomCDATASectionFromPointer(ptr unsafe.Pointer) *QDomCDATASection {
	var n = new(QDomCDATASection)
	n.SetPointer(ptr)
	return n
}

func newQDomCDATASectionFromPointer(ptr unsafe.Pointer) *QDomCDATASection {
	var n = NewQDomCDATASectionFromPointer(ptr)
	return n
}

func (ptr *QDomCDATASection) QDomCDATASection_PTR() *QDomCDATASection {
	return ptr
}

func NewQDomCDATASection() *QDomCDATASection {
	defer qt.Recovering("QDomCDATASection::QDomCDATASection")

	return newQDomCDATASectionFromPointer(C.QDomCDATASection_NewQDomCDATASection())
}

func NewQDomCDATASection2(x QDomCDATASection_ITF) *QDomCDATASection {
	defer qt.Recovering("QDomCDATASection::QDomCDATASection")

	return newQDomCDATASectionFromPointer(C.QDomCDATASection_NewQDomCDATASection2(PointerFromQDomCDATASection(x)))
}

func (ptr *QDomCDATASection) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomCDATASection::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomCDATASection_NodeType(ptr.Pointer()))
	}
	return 0
}

type QDomCharacterData struct {
	QDomNode
}

type QDomCharacterData_ITF interface {
	QDomNode_ITF
	QDomCharacterData_PTR() *QDomCharacterData
}

func PointerFromQDomCharacterData(ptr QDomCharacterData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomCharacterData_PTR().Pointer()
	}
	return nil
}

func NewQDomCharacterDataFromPointer(ptr unsafe.Pointer) *QDomCharacterData {
	var n = new(QDomCharacterData)
	n.SetPointer(ptr)
	return n
}

func newQDomCharacterDataFromPointer(ptr unsafe.Pointer) *QDomCharacterData {
	var n = NewQDomCharacterDataFromPointer(ptr)
	return n
}

func (ptr *QDomCharacterData) QDomCharacterData_PTR() *QDomCharacterData {
	return ptr
}

func NewQDomCharacterData() *QDomCharacterData {
	defer qt.Recovering("QDomCharacterData::QDomCharacterData")

	return newQDomCharacterDataFromPointer(C.QDomCharacterData_NewQDomCharacterData())
}

func NewQDomCharacterData2(x QDomCharacterData_ITF) *QDomCharacterData {
	defer qt.Recovering("QDomCharacterData::QDomCharacterData")

	return newQDomCharacterDataFromPointer(C.QDomCharacterData_NewQDomCharacterData2(PointerFromQDomCharacterData(x)))
}

func (ptr *QDomCharacterData) AppendData(arg string) {
	defer qt.Recovering("QDomCharacterData::appendData")

	if ptr.Pointer() != nil {
		C.QDomCharacterData_AppendData(ptr.Pointer(), C.CString(arg))
	}
}

func (ptr *QDomCharacterData) Data() string {
	defer qt.Recovering("QDomCharacterData::data")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomCharacterData_Data(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomCharacterData) Length() int {
	defer qt.Recovering("QDomCharacterData::length")

	if ptr.Pointer() != nil {
		return int(C.QDomCharacterData_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomCharacterData) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomCharacterData::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomCharacterData_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomCharacterData) SetData(v string) {
	defer qt.Recovering("QDomCharacterData::setData")

	if ptr.Pointer() != nil {
		C.QDomCharacterData_SetData(ptr.Pointer(), C.CString(v))
	}
}

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

func newQDomCommentFromPointer(ptr unsafe.Pointer) *QDomComment {
	var n = NewQDomCommentFromPointer(ptr)
	return n
}

func (ptr *QDomComment) QDomComment_PTR() *QDomComment {
	return ptr
}

func NewQDomComment() *QDomComment {
	defer qt.Recovering("QDomComment::QDomComment")

	return newQDomCommentFromPointer(C.QDomComment_NewQDomComment())
}

func NewQDomComment2(x QDomComment_ITF) *QDomComment {
	defer qt.Recovering("QDomComment::QDomComment")

	return newQDomCommentFromPointer(C.QDomComment_NewQDomComment2(PointerFromQDomComment(x)))
}

func (ptr *QDomComment) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomComment::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomComment_NodeType(ptr.Pointer()))
	}
	return 0
}

type QDomDocument struct {
	QDomNode
}

type QDomDocument_ITF interface {
	QDomNode_ITF
	QDomDocument_PTR() *QDomDocument
}

func PointerFromQDomDocument(ptr QDomDocument_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomDocument_PTR().Pointer()
	}
	return nil
}

func NewQDomDocumentFromPointer(ptr unsafe.Pointer) *QDomDocument {
	var n = new(QDomDocument)
	n.SetPointer(ptr)
	return n
}

func newQDomDocumentFromPointer(ptr unsafe.Pointer) *QDomDocument {
	var n = NewQDomDocumentFromPointer(ptr)
	return n
}

func (ptr *QDomDocument) QDomDocument_PTR() *QDomDocument {
	return ptr
}

func NewQDomDocument() *QDomDocument {
	defer qt.Recovering("QDomDocument::QDomDocument")

	return newQDomDocumentFromPointer(C.QDomDocument_NewQDomDocument())
}

func NewQDomDocument4(x QDomDocument_ITF) *QDomDocument {
	defer qt.Recovering("QDomDocument::QDomDocument")

	return newQDomDocumentFromPointer(C.QDomDocument_NewQDomDocument4(PointerFromQDomDocument(x)))
}

func NewQDomDocument3(doctype QDomDocumentType_ITF) *QDomDocument {
	defer qt.Recovering("QDomDocument::QDomDocument")

	return newQDomDocumentFromPointer(C.QDomDocument_NewQDomDocument3(PointerFromQDomDocumentType(doctype)))
}

func NewQDomDocument2(name string) *QDomDocument {
	defer qt.Recovering("QDomDocument::QDomDocument")

	return newQDomDocumentFromPointer(C.QDomDocument_NewQDomDocument2(C.CString(name)))
}

func (ptr *QDomDocument) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomDocument::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomDocument_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomDocument) SetContent7(dev core.QIODevice_ITF, errorMsg string, errorLine int, errorColumn int) bool {
	defer qt.Recovering("QDomDocument::setContent")

	if ptr.Pointer() != nil {
		return C.QDomDocument_SetContent7(ptr.Pointer(), core.PointerFromQIODevice(dev), C.CString(errorMsg), C.int(errorLine), C.int(errorColumn)) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent3(dev core.QIODevice_ITF, namespaceProcessing bool, errorMsg string, errorLine int, errorColumn int) bool {
	defer qt.Recovering("QDomDocument::setContent")

	if ptr.Pointer() != nil {
		return C.QDomDocument_SetContent3(ptr.Pointer(), core.PointerFromQIODevice(dev), C.int(qt.GoBoolToInt(namespaceProcessing)), C.CString(errorMsg), C.int(errorLine), C.int(errorColumn)) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent8(source QXmlInputSource_ITF, reader QXmlReader_ITF, errorMsg string, errorLine int, errorColumn int) bool {
	defer qt.Recovering("QDomDocument::setContent")

	if ptr.Pointer() != nil {
		return C.QDomDocument_SetContent8(ptr.Pointer(), PointerFromQXmlInputSource(source), PointerFromQXmlReader(reader), C.CString(errorMsg), C.int(errorLine), C.int(errorColumn)) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent4(source QXmlInputSource_ITF, namespaceProcessing bool, errorMsg string, errorLine int, errorColumn int) bool {
	defer qt.Recovering("QDomDocument::setContent")

	if ptr.Pointer() != nil {
		return C.QDomDocument_SetContent4(ptr.Pointer(), PointerFromQXmlInputSource(source), C.int(qt.GoBoolToInt(namespaceProcessing)), C.CString(errorMsg), C.int(errorLine), C.int(errorColumn)) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent6(buffer core.QByteArray_ITF, errorMsg string, errorLine int, errorColumn int) bool {
	defer qt.Recovering("QDomDocument::setContent")

	if ptr.Pointer() != nil {
		return C.QDomDocument_SetContent6(ptr.Pointer(), core.PointerFromQByteArray(buffer), C.CString(errorMsg), C.int(errorLine), C.int(errorColumn)) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent(data core.QByteArray_ITF, namespaceProcessing bool, errorMsg string, errorLine int, errorColumn int) bool {
	defer qt.Recovering("QDomDocument::setContent")

	if ptr.Pointer() != nil {
		return C.QDomDocument_SetContent(ptr.Pointer(), core.PointerFromQByteArray(data), C.int(qt.GoBoolToInt(namespaceProcessing)), C.CString(errorMsg), C.int(errorLine), C.int(errorColumn)) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent5(text string, errorMsg string, errorLine int, errorColumn int) bool {
	defer qt.Recovering("QDomDocument::setContent")

	if ptr.Pointer() != nil {
		return C.QDomDocument_SetContent5(ptr.Pointer(), C.CString(text), C.CString(errorMsg), C.int(errorLine), C.int(errorColumn)) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent2(text string, namespaceProcessing bool, errorMsg string, errorLine int, errorColumn int) bool {
	defer qt.Recovering("QDomDocument::setContent")

	if ptr.Pointer() != nil {
		return C.QDomDocument_SetContent2(ptr.Pointer(), C.CString(text), C.int(qt.GoBoolToInt(namespaceProcessing)), C.CString(errorMsg), C.int(errorLine), C.int(errorColumn)) != 0
	}
	return false
}

func (ptr *QDomDocument) ToByteArray(indent int) *core.QByteArray {
	defer qt.Recovering("QDomDocument::toByteArray")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QDomDocument_ToByteArray(ptr.Pointer(), C.int(indent)))
	}
	return nil
}

func (ptr *QDomDocument) ToString(indent int) string {
	defer qt.Recovering("QDomDocument::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomDocument_ToString(ptr.Pointer(), C.int(indent)))
	}
	return ""
}

func (ptr *QDomDocument) DestroyQDomDocument() {
	defer qt.Recovering("QDomDocument::~QDomDocument")

	if ptr.Pointer() != nil {
		C.QDomDocument_DestroyQDomDocument(ptr.Pointer())
	}
}

type QDomDocumentFragment struct {
	QDomNode
}

type QDomDocumentFragment_ITF interface {
	QDomNode_ITF
	QDomDocumentFragment_PTR() *QDomDocumentFragment
}

func PointerFromQDomDocumentFragment(ptr QDomDocumentFragment_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomDocumentFragment_PTR().Pointer()
	}
	return nil
}

func NewQDomDocumentFragmentFromPointer(ptr unsafe.Pointer) *QDomDocumentFragment {
	var n = new(QDomDocumentFragment)
	n.SetPointer(ptr)
	return n
}

func newQDomDocumentFragmentFromPointer(ptr unsafe.Pointer) *QDomDocumentFragment {
	var n = NewQDomDocumentFragmentFromPointer(ptr)
	return n
}

func (ptr *QDomDocumentFragment) QDomDocumentFragment_PTR() *QDomDocumentFragment {
	return ptr
}

func NewQDomDocumentFragment() *QDomDocumentFragment {
	defer qt.Recovering("QDomDocumentFragment::QDomDocumentFragment")

	return newQDomDocumentFragmentFromPointer(C.QDomDocumentFragment_NewQDomDocumentFragment())
}

func NewQDomDocumentFragment2(x QDomDocumentFragment_ITF) *QDomDocumentFragment {
	defer qt.Recovering("QDomDocumentFragment::QDomDocumentFragment")

	return newQDomDocumentFragmentFromPointer(C.QDomDocumentFragment_NewQDomDocumentFragment2(PointerFromQDomDocumentFragment(x)))
}

func (ptr *QDomDocumentFragment) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomDocumentFragment::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomDocumentFragment_NodeType(ptr.Pointer()))
	}
	return 0
}

type QDomDocumentType struct {
	QDomNode
}

type QDomDocumentType_ITF interface {
	QDomNode_ITF
	QDomDocumentType_PTR() *QDomDocumentType
}

func PointerFromQDomDocumentType(ptr QDomDocumentType_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomDocumentType_PTR().Pointer()
	}
	return nil
}

func NewQDomDocumentTypeFromPointer(ptr unsafe.Pointer) *QDomDocumentType {
	var n = new(QDomDocumentType)
	n.SetPointer(ptr)
	return n
}

func newQDomDocumentTypeFromPointer(ptr unsafe.Pointer) *QDomDocumentType {
	var n = NewQDomDocumentTypeFromPointer(ptr)
	return n
}

func (ptr *QDomDocumentType) QDomDocumentType_PTR() *QDomDocumentType {
	return ptr
}

func NewQDomDocumentType() *QDomDocumentType {
	defer qt.Recovering("QDomDocumentType::QDomDocumentType")

	return newQDomDocumentTypeFromPointer(C.QDomDocumentType_NewQDomDocumentType())
}

func NewQDomDocumentType2(n QDomDocumentType_ITF) *QDomDocumentType {
	defer qt.Recovering("QDomDocumentType::QDomDocumentType")

	return newQDomDocumentTypeFromPointer(C.QDomDocumentType_NewQDomDocumentType2(PointerFromQDomDocumentType(n)))
}

func (ptr *QDomDocumentType) InternalSubset() string {
	defer qt.Recovering("QDomDocumentType::internalSubset")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomDocumentType_InternalSubset(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomDocumentType) Name() string {
	defer qt.Recovering("QDomDocumentType::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomDocumentType_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomDocumentType) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomDocumentType::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomDocumentType_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomDocumentType) PublicId() string {
	defer qt.Recovering("QDomDocumentType::publicId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomDocumentType_PublicId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomDocumentType) SystemId() string {
	defer qt.Recovering("QDomDocumentType::systemId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomDocumentType_SystemId(ptr.Pointer()))
	}
	return ""
}

type QDomElement struct {
	QDomNode
}

type QDomElement_ITF interface {
	QDomNode_ITF
	QDomElement_PTR() *QDomElement
}

func PointerFromQDomElement(ptr QDomElement_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomElement_PTR().Pointer()
	}
	return nil
}

func NewQDomElementFromPointer(ptr unsafe.Pointer) *QDomElement {
	var n = new(QDomElement)
	n.SetPointer(ptr)
	return n
}

func newQDomElementFromPointer(ptr unsafe.Pointer) *QDomElement {
	var n = NewQDomElementFromPointer(ptr)
	return n
}

func (ptr *QDomElement) QDomElement_PTR() *QDomElement {
	return ptr
}

func NewQDomElement() *QDomElement {
	defer qt.Recovering("QDomElement::QDomElement")

	return newQDomElementFromPointer(C.QDomElement_NewQDomElement())
}

func NewQDomElement2(x QDomElement_ITF) *QDomElement {
	defer qt.Recovering("QDomElement::QDomElement")

	return newQDomElementFromPointer(C.QDomElement_NewQDomElement2(PointerFromQDomElement(x)))
}

func (ptr *QDomElement) Attribute(name string, defValue string) string {
	defer qt.Recovering("QDomElement::attribute")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomElement_Attribute(ptr.Pointer(), C.CString(name), C.CString(defValue)))
	}
	return ""
}

func (ptr *QDomElement) AttributeNS(nsURI string, localName string, defValue string) string {
	defer qt.Recovering("QDomElement::attributeNS")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomElement_AttributeNS(ptr.Pointer(), C.CString(nsURI), C.CString(localName), C.CString(defValue)))
	}
	return ""
}

func (ptr *QDomElement) HasAttribute(name string) bool {
	defer qt.Recovering("QDomElement::hasAttribute")

	if ptr.Pointer() != nil {
		return C.QDomElement_HasAttribute(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QDomElement) HasAttributeNS(nsURI string, localName string) bool {
	defer qt.Recovering("QDomElement::hasAttributeNS")

	if ptr.Pointer() != nil {
		return C.QDomElement_HasAttributeNS(ptr.Pointer(), C.CString(nsURI), C.CString(localName)) != 0
	}
	return false
}

func (ptr *QDomElement) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomElement::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomElement_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomElement) RemoveAttribute(name string) {
	defer qt.Recovering("QDomElement::removeAttribute")

	if ptr.Pointer() != nil {
		C.QDomElement_RemoveAttribute(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QDomElement) RemoveAttributeNS(nsURI string, localName string) {
	defer qt.Recovering("QDomElement::removeAttributeNS")

	if ptr.Pointer() != nil {
		C.QDomElement_RemoveAttributeNS(ptr.Pointer(), C.CString(nsURI), C.CString(localName))
	}
}

func (ptr *QDomElement) SetAttribute(name string, value string) {
	defer qt.Recovering("QDomElement::setAttribute")

	if ptr.Pointer() != nil {
		C.QDomElement_SetAttribute(ptr.Pointer(), C.CString(name), C.CString(value))
	}
}

func (ptr *QDomElement) SetAttribute2(name string, value int) {
	defer qt.Recovering("QDomElement::setAttribute")

	if ptr.Pointer() != nil {
		C.QDomElement_SetAttribute2(ptr.Pointer(), C.CString(name), C.int(value))
	}
}

func (ptr *QDomElement) SetAttributeNS(nsURI string, qName string, value string) {
	defer qt.Recovering("QDomElement::setAttributeNS")

	if ptr.Pointer() != nil {
		C.QDomElement_SetAttributeNS(ptr.Pointer(), C.CString(nsURI), C.CString(qName), C.CString(value))
	}
}

func (ptr *QDomElement) SetAttributeNS2(nsURI string, qName string, value int) {
	defer qt.Recovering("QDomElement::setAttributeNS")

	if ptr.Pointer() != nil {
		C.QDomElement_SetAttributeNS2(ptr.Pointer(), C.CString(nsURI), C.CString(qName), C.int(value))
	}
}

func (ptr *QDomElement) SetTagName(name string) {
	defer qt.Recovering("QDomElement::setTagName")

	if ptr.Pointer() != nil {
		C.QDomElement_SetTagName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QDomElement) TagName() string {
	defer qt.Recovering("QDomElement::tagName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomElement_TagName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomElement) Text() string {
	defer qt.Recovering("QDomElement::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomElement_Text(ptr.Pointer()))
	}
	return ""
}

type QDomEntity struct {
	QDomNode
}

type QDomEntity_ITF interface {
	QDomNode_ITF
	QDomEntity_PTR() *QDomEntity
}

func PointerFromQDomEntity(ptr QDomEntity_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomEntity_PTR().Pointer()
	}
	return nil
}

func NewQDomEntityFromPointer(ptr unsafe.Pointer) *QDomEntity {
	var n = new(QDomEntity)
	n.SetPointer(ptr)
	return n
}

func newQDomEntityFromPointer(ptr unsafe.Pointer) *QDomEntity {
	var n = NewQDomEntityFromPointer(ptr)
	return n
}

func (ptr *QDomEntity) QDomEntity_PTR() *QDomEntity {
	return ptr
}

func NewQDomEntity() *QDomEntity {
	defer qt.Recovering("QDomEntity::QDomEntity")

	return newQDomEntityFromPointer(C.QDomEntity_NewQDomEntity())
}

func NewQDomEntity2(x QDomEntity_ITF) *QDomEntity {
	defer qt.Recovering("QDomEntity::QDomEntity")

	return newQDomEntityFromPointer(C.QDomEntity_NewQDomEntity2(PointerFromQDomEntity(x)))
}

func (ptr *QDomEntity) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomEntity::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomEntity_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomEntity) NotationName() string {
	defer qt.Recovering("QDomEntity::notationName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomEntity_NotationName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomEntity) PublicId() string {
	defer qt.Recovering("QDomEntity::publicId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomEntity_PublicId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomEntity) SystemId() string {
	defer qt.Recovering("QDomEntity::systemId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomEntity_SystemId(ptr.Pointer()))
	}
	return ""
}

type QDomEntityReference struct {
	QDomNode
}

type QDomEntityReference_ITF interface {
	QDomNode_ITF
	QDomEntityReference_PTR() *QDomEntityReference
}

func PointerFromQDomEntityReference(ptr QDomEntityReference_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomEntityReference_PTR().Pointer()
	}
	return nil
}

func NewQDomEntityReferenceFromPointer(ptr unsafe.Pointer) *QDomEntityReference {
	var n = new(QDomEntityReference)
	n.SetPointer(ptr)
	return n
}

func newQDomEntityReferenceFromPointer(ptr unsafe.Pointer) *QDomEntityReference {
	var n = NewQDomEntityReferenceFromPointer(ptr)
	return n
}

func (ptr *QDomEntityReference) QDomEntityReference_PTR() *QDomEntityReference {
	return ptr
}

func NewQDomEntityReference() *QDomEntityReference {
	defer qt.Recovering("QDomEntityReference::QDomEntityReference")

	return newQDomEntityReferenceFromPointer(C.QDomEntityReference_NewQDomEntityReference())
}

func NewQDomEntityReference2(x QDomEntityReference_ITF) *QDomEntityReference {
	defer qt.Recovering("QDomEntityReference::QDomEntityReference")

	return newQDomEntityReferenceFromPointer(C.QDomEntityReference_NewQDomEntityReference2(PointerFromQDomEntityReference(x)))
}

func (ptr *QDomEntityReference) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomEntityReference::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomEntityReference_NodeType(ptr.Pointer()))
	}
	return 0
}

type QDomImplementation struct {
	ptr unsafe.Pointer
}

type QDomImplementation_ITF interface {
	QDomImplementation_PTR() *QDomImplementation
}

func (p *QDomImplementation) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDomImplementation) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDomImplementation(ptr QDomImplementation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomImplementation_PTR().Pointer()
	}
	return nil
}

func NewQDomImplementationFromPointer(ptr unsafe.Pointer) *QDomImplementation {
	var n = new(QDomImplementation)
	n.SetPointer(ptr)
	return n
}

func newQDomImplementationFromPointer(ptr unsafe.Pointer) *QDomImplementation {
	var n = NewQDomImplementationFromPointer(ptr)
	return n
}

func (ptr *QDomImplementation) QDomImplementation_PTR() *QDomImplementation {
	return ptr
}

//QDomImplementation::InvalidDataPolicy
type QDomImplementation__InvalidDataPolicy int64

const (
	QDomImplementation__AcceptInvalidChars = QDomImplementation__InvalidDataPolicy(0)
	QDomImplementation__DropInvalidChars   = QDomImplementation__InvalidDataPolicy(1)
	QDomImplementation__ReturnNullNode     = QDomImplementation__InvalidDataPolicy(2)
)

func NewQDomImplementation() *QDomImplementation {
	defer qt.Recovering("QDomImplementation::QDomImplementation")

	return newQDomImplementationFromPointer(C.QDomImplementation_NewQDomImplementation())
}

func NewQDomImplementation2(x QDomImplementation_ITF) *QDomImplementation {
	defer qt.Recovering("QDomImplementation::QDomImplementation")

	return newQDomImplementationFromPointer(C.QDomImplementation_NewQDomImplementation2(PointerFromQDomImplementation(x)))
}

func (ptr *QDomImplementation) HasFeature(feature string, version string) bool {
	defer qt.Recovering("QDomImplementation::hasFeature")

	if ptr.Pointer() != nil {
		return C.QDomImplementation_HasFeature(ptr.Pointer(), C.CString(feature), C.CString(version)) != 0
	}
	return false
}

func QDomImplementation_InvalidDataPolicy() QDomImplementation__InvalidDataPolicy {
	defer qt.Recovering("QDomImplementation::invalidDataPolicy")

	return QDomImplementation__InvalidDataPolicy(C.QDomImplementation_QDomImplementation_InvalidDataPolicy())
}

func (ptr *QDomImplementation) IsNull() bool {
	defer qt.Recovering("QDomImplementation::isNull")

	if ptr.Pointer() != nil {
		return C.QDomImplementation_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func QDomImplementation_SetInvalidDataPolicy(policy QDomImplementation__InvalidDataPolicy) {
	defer qt.Recovering("QDomImplementation::setInvalidDataPolicy")

	C.QDomImplementation_QDomImplementation_SetInvalidDataPolicy(C.int(policy))
}

func (ptr *QDomImplementation) DestroyQDomImplementation() {
	defer qt.Recovering("QDomImplementation::~QDomImplementation")

	if ptr.Pointer() != nil {
		C.QDomImplementation_DestroyQDomImplementation(ptr.Pointer())
	}
}

type QDomNamedNodeMap struct {
	ptr unsafe.Pointer
}

type QDomNamedNodeMap_ITF interface {
	QDomNamedNodeMap_PTR() *QDomNamedNodeMap
}

func (p *QDomNamedNodeMap) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDomNamedNodeMap) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDomNamedNodeMap(ptr QDomNamedNodeMap_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNamedNodeMap_PTR().Pointer()
	}
	return nil
}

func NewQDomNamedNodeMapFromPointer(ptr unsafe.Pointer) *QDomNamedNodeMap {
	var n = new(QDomNamedNodeMap)
	n.SetPointer(ptr)
	return n
}

func newQDomNamedNodeMapFromPointer(ptr unsafe.Pointer) *QDomNamedNodeMap {
	var n = NewQDomNamedNodeMapFromPointer(ptr)
	return n
}

func (ptr *QDomNamedNodeMap) QDomNamedNodeMap_PTR() *QDomNamedNodeMap {
	return ptr
}

func NewQDomNamedNodeMap() *QDomNamedNodeMap {
	defer qt.Recovering("QDomNamedNodeMap::QDomNamedNodeMap")

	return newQDomNamedNodeMapFromPointer(C.QDomNamedNodeMap_NewQDomNamedNodeMap())
}

func NewQDomNamedNodeMap2(n QDomNamedNodeMap_ITF) *QDomNamedNodeMap {
	defer qt.Recovering("QDomNamedNodeMap::QDomNamedNodeMap")

	return newQDomNamedNodeMapFromPointer(C.QDomNamedNodeMap_NewQDomNamedNodeMap2(PointerFromQDomNamedNodeMap(n)))
}

func (ptr *QDomNamedNodeMap) Contains(name string) bool {
	defer qt.Recovering("QDomNamedNodeMap::contains")

	if ptr.Pointer() != nil {
		return C.QDomNamedNodeMap_Contains(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QDomNamedNodeMap) Count() int {
	defer qt.Recovering("QDomNamedNodeMap::count")

	if ptr.Pointer() != nil {
		return int(C.QDomNamedNodeMap_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNamedNodeMap) IsEmpty() bool {
	defer qt.Recovering("QDomNamedNodeMap::isEmpty")

	if ptr.Pointer() != nil {
		return C.QDomNamedNodeMap_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNamedNodeMap) Length() int {
	defer qt.Recovering("QDomNamedNodeMap::length")

	if ptr.Pointer() != nil {
		return int(C.QDomNamedNodeMap_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNamedNodeMap) Size() int {
	defer qt.Recovering("QDomNamedNodeMap::size")

	if ptr.Pointer() != nil {
		return int(C.QDomNamedNodeMap_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNamedNodeMap) DestroyQDomNamedNodeMap() {
	defer qt.Recovering("QDomNamedNodeMap::~QDomNamedNodeMap")

	if ptr.Pointer() != nil {
		C.QDomNamedNodeMap_DestroyQDomNamedNodeMap(ptr.Pointer())
	}
}

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

func newQDomNodeFromPointer(ptr unsafe.Pointer) *QDomNode {
	var n = NewQDomNodeFromPointer(ptr)
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

	return newQDomNodeFromPointer(C.QDomNode_NewQDomNode())
}

func NewQDomNode2(n QDomNode_ITF) *QDomNode {
	defer qt.Recovering("QDomNode::QDomNode")

	return newQDomNodeFromPointer(C.QDomNode_NewQDomNode2(PointerFromQDomNode(n)))
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

type QDomNodeList struct {
	ptr unsafe.Pointer
}

type QDomNodeList_ITF interface {
	QDomNodeList_PTR() *QDomNodeList
}

func (p *QDomNodeList) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDomNodeList) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDomNodeList(ptr QDomNodeList_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNodeList_PTR().Pointer()
	}
	return nil
}

func NewQDomNodeListFromPointer(ptr unsafe.Pointer) *QDomNodeList {
	var n = new(QDomNodeList)
	n.SetPointer(ptr)
	return n
}

func newQDomNodeListFromPointer(ptr unsafe.Pointer) *QDomNodeList {
	var n = NewQDomNodeListFromPointer(ptr)
	return n
}

func (ptr *QDomNodeList) QDomNodeList_PTR() *QDomNodeList {
	return ptr
}

func NewQDomNodeList() *QDomNodeList {
	defer qt.Recovering("QDomNodeList::QDomNodeList")

	return newQDomNodeListFromPointer(C.QDomNodeList_NewQDomNodeList())
}

func NewQDomNodeList2(n QDomNodeList_ITF) *QDomNodeList {
	defer qt.Recovering("QDomNodeList::QDomNodeList")

	return newQDomNodeListFromPointer(C.QDomNodeList_NewQDomNodeList2(PointerFromQDomNodeList(n)))
}

func (ptr *QDomNodeList) Count() int {
	defer qt.Recovering("QDomNodeList::count")

	if ptr.Pointer() != nil {
		return int(C.QDomNodeList_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNodeList) IsEmpty() bool {
	defer qt.Recovering("QDomNodeList::isEmpty")

	if ptr.Pointer() != nil {
		return C.QDomNodeList_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDomNodeList) Length() int {
	defer qt.Recovering("QDomNodeList::length")

	if ptr.Pointer() != nil {
		return int(C.QDomNodeList_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNodeList) Size() int {
	defer qt.Recovering("QDomNodeList::size")

	if ptr.Pointer() != nil {
		return int(C.QDomNodeList_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNodeList) DestroyQDomNodeList() {
	defer qt.Recovering("QDomNodeList::~QDomNodeList")

	if ptr.Pointer() != nil {
		C.QDomNodeList_DestroyQDomNodeList(ptr.Pointer())
	}
}

type QDomNotation struct {
	QDomNode
}

type QDomNotation_ITF interface {
	QDomNode_ITF
	QDomNotation_PTR() *QDomNotation
}

func PointerFromQDomNotation(ptr QDomNotation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNotation_PTR().Pointer()
	}
	return nil
}

func NewQDomNotationFromPointer(ptr unsafe.Pointer) *QDomNotation {
	var n = new(QDomNotation)
	n.SetPointer(ptr)
	return n
}

func newQDomNotationFromPointer(ptr unsafe.Pointer) *QDomNotation {
	var n = NewQDomNotationFromPointer(ptr)
	return n
}

func (ptr *QDomNotation) QDomNotation_PTR() *QDomNotation {
	return ptr
}

func NewQDomNotation() *QDomNotation {
	defer qt.Recovering("QDomNotation::QDomNotation")

	return newQDomNotationFromPointer(C.QDomNotation_NewQDomNotation())
}

func NewQDomNotation2(x QDomNotation_ITF) *QDomNotation {
	defer qt.Recovering("QDomNotation::QDomNotation")

	return newQDomNotationFromPointer(C.QDomNotation_NewQDomNotation2(PointerFromQDomNotation(x)))
}

func (ptr *QDomNotation) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomNotation::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomNotation_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNotation) PublicId() string {
	defer qt.Recovering("QDomNotation::publicId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomNotation_PublicId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomNotation) SystemId() string {
	defer qt.Recovering("QDomNotation::systemId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomNotation_SystemId(ptr.Pointer()))
	}
	return ""
}

type QDomProcessingInstruction struct {
	QDomNode
}

type QDomProcessingInstruction_ITF interface {
	QDomNode_ITF
	QDomProcessingInstruction_PTR() *QDomProcessingInstruction
}

func PointerFromQDomProcessingInstruction(ptr QDomProcessingInstruction_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomProcessingInstruction_PTR().Pointer()
	}
	return nil
}

func NewQDomProcessingInstructionFromPointer(ptr unsafe.Pointer) *QDomProcessingInstruction {
	var n = new(QDomProcessingInstruction)
	n.SetPointer(ptr)
	return n
}

func newQDomProcessingInstructionFromPointer(ptr unsafe.Pointer) *QDomProcessingInstruction {
	var n = NewQDomProcessingInstructionFromPointer(ptr)
	return n
}

func (ptr *QDomProcessingInstruction) QDomProcessingInstruction_PTR() *QDomProcessingInstruction {
	return ptr
}

func NewQDomProcessingInstruction() *QDomProcessingInstruction {
	defer qt.Recovering("QDomProcessingInstruction::QDomProcessingInstruction")

	return newQDomProcessingInstructionFromPointer(C.QDomProcessingInstruction_NewQDomProcessingInstruction())
}

func NewQDomProcessingInstruction2(x QDomProcessingInstruction_ITF) *QDomProcessingInstruction {
	defer qt.Recovering("QDomProcessingInstruction::QDomProcessingInstruction")

	return newQDomProcessingInstructionFromPointer(C.QDomProcessingInstruction_NewQDomProcessingInstruction2(PointerFromQDomProcessingInstruction(x)))
}

func (ptr *QDomProcessingInstruction) Data() string {
	defer qt.Recovering("QDomProcessingInstruction::data")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomProcessingInstruction_Data(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomProcessingInstruction) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomProcessingInstruction::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomProcessingInstruction_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomProcessingInstruction) SetData(d string) {
	defer qt.Recovering("QDomProcessingInstruction::setData")

	if ptr.Pointer() != nil {
		C.QDomProcessingInstruction_SetData(ptr.Pointer(), C.CString(d))
	}
}

func (ptr *QDomProcessingInstruction) Target() string {
	defer qt.Recovering("QDomProcessingInstruction::target")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomProcessingInstruction_Target(ptr.Pointer()))
	}
	return ""
}

type QDomText struct {
	QDomCharacterData
}

type QDomText_ITF interface {
	QDomCharacterData_ITF
	QDomText_PTR() *QDomText
}

func PointerFromQDomText(ptr QDomText_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomText_PTR().Pointer()
	}
	return nil
}

func NewQDomTextFromPointer(ptr unsafe.Pointer) *QDomText {
	var n = new(QDomText)
	n.SetPointer(ptr)
	return n
}

func newQDomTextFromPointer(ptr unsafe.Pointer) *QDomText {
	var n = NewQDomTextFromPointer(ptr)
	return n
}

func (ptr *QDomText) QDomText_PTR() *QDomText {
	return ptr
}

func NewQDomText() *QDomText {
	defer qt.Recovering("QDomText::QDomText")

	return newQDomTextFromPointer(C.QDomText_NewQDomText())
}

func NewQDomText2(x QDomText_ITF) *QDomText {
	defer qt.Recovering("QDomText::QDomText")

	return newQDomTextFromPointer(C.QDomText_NewQDomText2(PointerFromQDomText(x)))
}

func (ptr *QDomText) NodeType() QDomNode__NodeType {
	defer qt.Recovering("QDomText::nodeType")

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomText_NodeType(ptr.Pointer()))
	}
	return 0
}

type QXmlAttributes struct {
	ptr unsafe.Pointer
}

type QXmlAttributes_ITF interface {
	QXmlAttributes_PTR() *QXmlAttributes
}

func (p *QXmlAttributes) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlAttributes) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlAttributes(ptr QXmlAttributes_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlAttributes_PTR().Pointer()
	}
	return nil
}

func NewQXmlAttributesFromPointer(ptr unsafe.Pointer) *QXmlAttributes {
	var n = new(QXmlAttributes)
	n.SetPointer(ptr)
	return n
}

func newQXmlAttributesFromPointer(ptr unsafe.Pointer) *QXmlAttributes {
	var n = NewQXmlAttributesFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlAttributes_") {
		n.SetObjectNameAbs("QXmlAttributes_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlAttributes) QXmlAttributes_PTR() *QXmlAttributes {
	return ptr
}

func NewQXmlAttributes() *QXmlAttributes {
	defer qt.Recovering("QXmlAttributes::QXmlAttributes")

	return newQXmlAttributesFromPointer(C.QXmlAttributes_NewQXmlAttributes())
}

func (ptr *QXmlAttributes) DestroyQXmlAttributes() {
	defer qt.Recovering("QXmlAttributes::~QXmlAttributes")

	if ptr.Pointer() != nil {
		C.QXmlAttributes_DestroyQXmlAttributes(ptr.Pointer())
	}
}

func (ptr *QXmlAttributes) Append(qName string, uri string, localPart string, value string) {
	defer qt.Recovering("QXmlAttributes::append")

	if ptr.Pointer() != nil {
		C.QXmlAttributes_Append(ptr.Pointer(), C.CString(qName), C.CString(uri), C.CString(localPart), C.CString(value))
	}
}

func (ptr *QXmlAttributes) Clear() {
	defer qt.Recovering("QXmlAttributes::clear")

	if ptr.Pointer() != nil {
		C.QXmlAttributes_Clear(ptr.Pointer())
	}
}

func (ptr *QXmlAttributes) Count() int {
	defer qt.Recovering("QXmlAttributes::count")

	if ptr.Pointer() != nil {
		return int(C.QXmlAttributes_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlAttributes) Index2(qName core.QLatin1String_ITF) int {
	defer qt.Recovering("QXmlAttributes::index")

	if ptr.Pointer() != nil {
		return int(C.QXmlAttributes_Index2(ptr.Pointer(), core.PointerFromQLatin1String(qName)))
	}
	return 0
}

func (ptr *QXmlAttributes) Index(qName string) int {
	defer qt.Recovering("QXmlAttributes::index")

	if ptr.Pointer() != nil {
		return int(C.QXmlAttributes_Index(ptr.Pointer(), C.CString(qName)))
	}
	return 0
}

func (ptr *QXmlAttributes) Index3(uri string, localPart string) int {
	defer qt.Recovering("QXmlAttributes::index")

	if ptr.Pointer() != nil {
		return int(C.QXmlAttributes_Index3(ptr.Pointer(), C.CString(uri), C.CString(localPart)))
	}
	return 0
}

func (ptr *QXmlAttributes) Length() int {
	defer qt.Recovering("QXmlAttributes::length")

	if ptr.Pointer() != nil {
		return int(C.QXmlAttributes_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlAttributes) LocalName(index int) string {
	defer qt.Recovering("QXmlAttributes::localName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_LocalName(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QXmlAttributes) QName(index int) string {
	defer qt.Recovering("QXmlAttributes::qName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_QName(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QXmlAttributes) Type2(qName string) string {
	defer qt.Recovering("QXmlAttributes::type")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_Type2(ptr.Pointer(), C.CString(qName)))
	}
	return ""
}

func (ptr *QXmlAttributes) Type3(uri string, localName string) string {
	defer qt.Recovering("QXmlAttributes::type")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_Type3(ptr.Pointer(), C.CString(uri), C.CString(localName)))
	}
	return ""
}

func (ptr *QXmlAttributes) Type(index int) string {
	defer qt.Recovering("QXmlAttributes::type")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_Type(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QXmlAttributes) Uri(index int) string {
	defer qt.Recovering("QXmlAttributes::uri")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_Uri(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QXmlAttributes) Value3(qName core.QLatin1String_ITF) string {
	defer qt.Recovering("QXmlAttributes::value")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_Value3(ptr.Pointer(), core.PointerFromQLatin1String(qName)))
	}
	return ""
}

func (ptr *QXmlAttributes) Value2(qName string) string {
	defer qt.Recovering("QXmlAttributes::value")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_Value2(ptr.Pointer(), C.CString(qName)))
	}
	return ""
}

func (ptr *QXmlAttributes) Value4(uri string, localName string) string {
	defer qt.Recovering("QXmlAttributes::value")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_Value4(ptr.Pointer(), C.CString(uri), C.CString(localName)))
	}
	return ""
}

func (ptr *QXmlAttributes) Value(index int) string {
	defer qt.Recovering("QXmlAttributes::value")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_Value(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QXmlAttributes) ObjectNameAbs() string {
	defer qt.Recovering("QXmlAttributes::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlAttributes) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlAttributes::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlAttributes_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QXmlContentHandler struct {
	ptr unsafe.Pointer
}

type QXmlContentHandler_ITF interface {
	QXmlContentHandler_PTR() *QXmlContentHandler
}

func (p *QXmlContentHandler) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlContentHandler) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlContentHandler(ptr QXmlContentHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlContentHandler_PTR().Pointer()
	}
	return nil
}

func NewQXmlContentHandlerFromPointer(ptr unsafe.Pointer) *QXmlContentHandler {
	var n = new(QXmlContentHandler)
	n.SetPointer(ptr)
	return n
}

func newQXmlContentHandlerFromPointer(ptr unsafe.Pointer) *QXmlContentHandler {
	var n = NewQXmlContentHandlerFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlContentHandler_") {
		n.SetObjectNameAbs("QXmlContentHandler_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlContentHandler) QXmlContentHandler_PTR() *QXmlContentHandler {
	return ptr
}

func (ptr *QXmlContentHandler) Characters(ch string) bool {
	defer qt.Recovering("QXmlContentHandler::characters")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_Characters(ptr.Pointer(), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) EndDocument() bool {
	defer qt.Recovering("QXmlContentHandler::endDocument")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_EndDocument(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) EndElement(namespaceURI string, localName string, qName string) bool {
	defer qt.Recovering("QXmlContentHandler::endElement")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_EndElement(ptr.Pointer(), C.CString(namespaceURI), C.CString(localName), C.CString(qName)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) EndPrefixMapping(prefix string) bool {
	defer qt.Recovering("QXmlContentHandler::endPrefixMapping")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_EndPrefixMapping(ptr.Pointer(), C.CString(prefix)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) ErrorString() string {
	defer qt.Recovering("QXmlContentHandler::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlContentHandler_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlContentHandler) IgnorableWhitespace(ch string) bool {
	defer qt.Recovering("QXmlContentHandler::ignorableWhitespace")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_IgnorableWhitespace(ptr.Pointer(), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) ProcessingInstruction(target string, data string) bool {
	defer qt.Recovering("QXmlContentHandler::processingInstruction")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_ProcessingInstruction(ptr.Pointer(), C.CString(target), C.CString(data)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) SetDocumentLocator(locator QXmlLocator_ITF) {
	defer qt.Recovering("QXmlContentHandler::setDocumentLocator")

	if ptr.Pointer() != nil {
		C.QXmlContentHandler_SetDocumentLocator(ptr.Pointer(), PointerFromQXmlLocator(locator))
	}
}

func (ptr *QXmlContentHandler) SkippedEntity(name string) bool {
	defer qt.Recovering("QXmlContentHandler::skippedEntity")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_SkippedEntity(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) StartDocument() bool {
	defer qt.Recovering("QXmlContentHandler::startDocument")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_StartDocument(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) StartElement(namespaceURI string, localName string, qName string, atts QXmlAttributes_ITF) bool {
	defer qt.Recovering("QXmlContentHandler::startElement")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_StartElement(ptr.Pointer(), C.CString(namespaceURI), C.CString(localName), C.CString(qName), PointerFromQXmlAttributes(atts)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) StartPrefixMapping(prefix string, uri string) bool {
	defer qt.Recovering("QXmlContentHandler::startPrefixMapping")

	if ptr.Pointer() != nil {
		return C.QXmlContentHandler_StartPrefixMapping(ptr.Pointer(), C.CString(prefix), C.CString(uri)) != 0
	}
	return false
}

func (ptr *QXmlContentHandler) DestroyQXmlContentHandler() {
	defer qt.Recovering("QXmlContentHandler::~QXmlContentHandler")

	if ptr.Pointer() != nil {
		C.QXmlContentHandler_DestroyQXmlContentHandler(ptr.Pointer())
	}
}

func (ptr *QXmlContentHandler) ObjectNameAbs() string {
	defer qt.Recovering("QXmlContentHandler::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlContentHandler_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlContentHandler) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlContentHandler::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlContentHandler_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QXmlDTDHandler struct {
	ptr unsafe.Pointer
}

type QXmlDTDHandler_ITF interface {
	QXmlDTDHandler_PTR() *QXmlDTDHandler
}

func (p *QXmlDTDHandler) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlDTDHandler) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlDTDHandler(ptr QXmlDTDHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlDTDHandler_PTR().Pointer()
	}
	return nil
}

func NewQXmlDTDHandlerFromPointer(ptr unsafe.Pointer) *QXmlDTDHandler {
	var n = new(QXmlDTDHandler)
	n.SetPointer(ptr)
	return n
}

func newQXmlDTDHandlerFromPointer(ptr unsafe.Pointer) *QXmlDTDHandler {
	var n = NewQXmlDTDHandlerFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlDTDHandler_") {
		n.SetObjectNameAbs("QXmlDTDHandler_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlDTDHandler) QXmlDTDHandler_PTR() *QXmlDTDHandler {
	return ptr
}

func (ptr *QXmlDTDHandler) ErrorString() string {
	defer qt.Recovering("QXmlDTDHandler::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlDTDHandler_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlDTDHandler) NotationDecl(name string, publicId string, systemId string) bool {
	defer qt.Recovering("QXmlDTDHandler::notationDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDTDHandler_NotationDecl(ptr.Pointer(), C.CString(name), C.CString(publicId), C.CString(systemId)) != 0
	}
	return false
}

func (ptr *QXmlDTDHandler) UnparsedEntityDecl(name string, publicId string, systemId string, notationName string) bool {
	defer qt.Recovering("QXmlDTDHandler::unparsedEntityDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDTDHandler_UnparsedEntityDecl(ptr.Pointer(), C.CString(name), C.CString(publicId), C.CString(systemId), C.CString(notationName)) != 0
	}
	return false
}

func (ptr *QXmlDTDHandler) DestroyQXmlDTDHandler() {
	defer qt.Recovering("QXmlDTDHandler::~QXmlDTDHandler")

	if ptr.Pointer() != nil {
		C.QXmlDTDHandler_DestroyQXmlDTDHandler(ptr.Pointer())
	}
}

func (ptr *QXmlDTDHandler) ObjectNameAbs() string {
	defer qt.Recovering("QXmlDTDHandler::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlDTDHandler_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlDTDHandler) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlDTDHandler::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlDTDHandler_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QXmlDeclHandler struct {
	ptr unsafe.Pointer
}

type QXmlDeclHandler_ITF interface {
	QXmlDeclHandler_PTR() *QXmlDeclHandler
}

func (p *QXmlDeclHandler) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlDeclHandler) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlDeclHandler(ptr QXmlDeclHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlDeclHandler_PTR().Pointer()
	}
	return nil
}

func NewQXmlDeclHandlerFromPointer(ptr unsafe.Pointer) *QXmlDeclHandler {
	var n = new(QXmlDeclHandler)
	n.SetPointer(ptr)
	return n
}

func newQXmlDeclHandlerFromPointer(ptr unsafe.Pointer) *QXmlDeclHandler {
	var n = NewQXmlDeclHandlerFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlDeclHandler_") {
		n.SetObjectNameAbs("QXmlDeclHandler_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlDeclHandler) QXmlDeclHandler_PTR() *QXmlDeclHandler {
	return ptr
}

func (ptr *QXmlDeclHandler) AttributeDecl(eName string, aName string, ty string, valueDefault string, value string) bool {
	defer qt.Recovering("QXmlDeclHandler::attributeDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDeclHandler_AttributeDecl(ptr.Pointer(), C.CString(eName), C.CString(aName), C.CString(ty), C.CString(valueDefault), C.CString(value)) != 0
	}
	return false
}

func (ptr *QXmlDeclHandler) ErrorString() string {
	defer qt.Recovering("QXmlDeclHandler::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlDeclHandler_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlDeclHandler) ExternalEntityDecl(name string, publicId string, systemId string) bool {
	defer qt.Recovering("QXmlDeclHandler::externalEntityDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDeclHandler_ExternalEntityDecl(ptr.Pointer(), C.CString(name), C.CString(publicId), C.CString(systemId)) != 0
	}
	return false
}

func (ptr *QXmlDeclHandler) InternalEntityDecl(name string, value string) bool {
	defer qt.Recovering("QXmlDeclHandler::internalEntityDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDeclHandler_InternalEntityDecl(ptr.Pointer(), C.CString(name), C.CString(value)) != 0
	}
	return false
}

func (ptr *QXmlDeclHandler) DestroyQXmlDeclHandler() {
	defer qt.Recovering("QXmlDeclHandler::~QXmlDeclHandler")

	if ptr.Pointer() != nil {
		C.QXmlDeclHandler_DestroyQXmlDeclHandler(ptr.Pointer())
	}
}

func (ptr *QXmlDeclHandler) ObjectNameAbs() string {
	defer qt.Recovering("QXmlDeclHandler::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlDeclHandler_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlDeclHandler) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlDeclHandler::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlDeclHandler_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QXmlDefaultHandler struct {
	QXmlDTDHandler
	QXmlErrorHandler
	QXmlLexicalHandler
	QXmlEntityResolver
	QXmlContentHandler
	QXmlDeclHandler
}

type QXmlDefaultHandler_ITF interface {
	QXmlDTDHandler_ITF
	QXmlErrorHandler_ITF
	QXmlLexicalHandler_ITF
	QXmlEntityResolver_ITF
	QXmlContentHandler_ITF
	QXmlDeclHandler_ITF
	QXmlDefaultHandler_PTR() *QXmlDefaultHandler
}

func (p *QXmlDefaultHandler) Pointer() unsafe.Pointer {
	return p.QXmlDTDHandler_PTR().Pointer()
}

func (p *QXmlDefaultHandler) SetPointer(ptr unsafe.Pointer) {
	p.QXmlDTDHandler_PTR().SetPointer(ptr)
	p.QXmlErrorHandler_PTR().SetPointer(ptr)
	p.QXmlLexicalHandler_PTR().SetPointer(ptr)
	p.QXmlEntityResolver_PTR().SetPointer(ptr)
	p.QXmlContentHandler_PTR().SetPointer(ptr)
	p.QXmlDeclHandler_PTR().SetPointer(ptr)
}

func PointerFromQXmlDefaultHandler(ptr QXmlDefaultHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlDefaultHandler_PTR().Pointer()
	}
	return nil
}

func NewQXmlDefaultHandlerFromPointer(ptr unsafe.Pointer) *QXmlDefaultHandler {
	var n = new(QXmlDefaultHandler)
	n.SetPointer(ptr)
	return n
}

func newQXmlDefaultHandlerFromPointer(ptr unsafe.Pointer) *QXmlDefaultHandler {
	var n = NewQXmlDefaultHandlerFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlDefaultHandler_") {
		n.SetObjectNameAbs("QXmlDefaultHandler_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlDefaultHandler) QXmlDefaultHandler_PTR() *QXmlDefaultHandler {
	return ptr
}

func NewQXmlDefaultHandler() *QXmlDefaultHandler {
	defer qt.Recovering("QXmlDefaultHandler::QXmlDefaultHandler")

	return newQXmlDefaultHandlerFromPointer(C.QXmlDefaultHandler_NewQXmlDefaultHandler())
}

func (ptr *QXmlDefaultHandler) DestroyQXmlDefaultHandler() {
	defer qt.Recovering("QXmlDefaultHandler::~QXmlDefaultHandler")

	if ptr.Pointer() != nil {
		C.QXmlDefaultHandler_DestroyQXmlDefaultHandler(ptr.Pointer())
	}
}

func (ptr *QXmlDefaultHandler) AttributeDecl(eName string, aName string, ty string, valueDefault string, value string) bool {
	defer qt.Recovering("QXmlDefaultHandler::attributeDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_AttributeDecl(ptr.Pointer(), C.CString(eName), C.CString(aName), C.CString(ty), C.CString(valueDefault), C.CString(value)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) Characters(ch string) bool {
	defer qt.Recovering("QXmlDefaultHandler::characters")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_Characters(ptr.Pointer(), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) Comment(ch string) bool {
	defer qt.Recovering("QXmlDefaultHandler::comment")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_Comment(ptr.Pointer(), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) EndCDATA() bool {
	defer qt.Recovering("QXmlDefaultHandler::endCDATA")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_EndCDATA(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) EndDTD() bool {
	defer qt.Recovering("QXmlDefaultHandler::endDTD")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_EndDTD(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) EndDocument() bool {
	defer qt.Recovering("QXmlDefaultHandler::endDocument")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_EndDocument(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) EndElement(namespaceURI string, localName string, qName string) bool {
	defer qt.Recovering("QXmlDefaultHandler::endElement")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_EndElement(ptr.Pointer(), C.CString(namespaceURI), C.CString(localName), C.CString(qName)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) EndEntity(name string) bool {
	defer qt.Recovering("QXmlDefaultHandler::endEntity")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_EndEntity(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) EndPrefixMapping(prefix string) bool {
	defer qt.Recovering("QXmlDefaultHandler::endPrefixMapping")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_EndPrefixMapping(ptr.Pointer(), C.CString(prefix)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) Error(exception QXmlParseException_ITF) bool {
	defer qt.Recovering("QXmlDefaultHandler::error")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_Error(ptr.Pointer(), PointerFromQXmlParseException(exception)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) ErrorString() string {
	defer qt.Recovering("QXmlDefaultHandler::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlDefaultHandler_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlDefaultHandler) ExternalEntityDecl(name string, publicId string, systemId string) bool {
	defer qt.Recovering("QXmlDefaultHandler::externalEntityDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_ExternalEntityDecl(ptr.Pointer(), C.CString(name), C.CString(publicId), C.CString(systemId)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) FatalError(exception QXmlParseException_ITF) bool {
	defer qt.Recovering("QXmlDefaultHandler::fatalError")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_FatalError(ptr.Pointer(), PointerFromQXmlParseException(exception)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) IgnorableWhitespace(ch string) bool {
	defer qt.Recovering("QXmlDefaultHandler::ignorableWhitespace")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_IgnorableWhitespace(ptr.Pointer(), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) InternalEntityDecl(name string, value string) bool {
	defer qt.Recovering("QXmlDefaultHandler::internalEntityDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_InternalEntityDecl(ptr.Pointer(), C.CString(name), C.CString(value)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) NotationDecl(name string, publicId string, systemId string) bool {
	defer qt.Recovering("QXmlDefaultHandler::notationDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_NotationDecl(ptr.Pointer(), C.CString(name), C.CString(publicId), C.CString(systemId)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) ProcessingInstruction(target string, data string) bool {
	defer qt.Recovering("QXmlDefaultHandler::processingInstruction")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_ProcessingInstruction(ptr.Pointer(), C.CString(target), C.CString(data)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) ConnectSetDocumentLocator(f func(locator *QXmlLocator)) {
	defer qt.Recovering("connect QXmlDefaultHandler::setDocumentLocator")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setDocumentLocator", f)
	}
}

func (ptr *QXmlDefaultHandler) DisconnectSetDocumentLocator() {
	defer qt.Recovering("disconnect QXmlDefaultHandler::setDocumentLocator")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setDocumentLocator")
	}
}

//export callbackQXmlDefaultHandlerSetDocumentLocator
func callbackQXmlDefaultHandlerSetDocumentLocator(ptr unsafe.Pointer, ptrName *C.char, locator unsafe.Pointer) {
	defer qt.Recovering("callback QXmlDefaultHandler::setDocumentLocator")

	if signal := qt.GetSignal(C.GoString(ptrName), "setDocumentLocator"); signal != nil {
		signal.(func(*QXmlLocator))(NewQXmlLocatorFromPointer(locator))
	} else {
		NewQXmlDefaultHandlerFromPointer(ptr).SetDocumentLocatorDefault(NewQXmlLocatorFromPointer(locator))
	}
}

func (ptr *QXmlDefaultHandler) SetDocumentLocator(locator QXmlLocator_ITF) {
	defer qt.Recovering("QXmlDefaultHandler::setDocumentLocator")

	if ptr.Pointer() != nil {
		C.QXmlDefaultHandler_SetDocumentLocator(ptr.Pointer(), PointerFromQXmlLocator(locator))
	}
}

func (ptr *QXmlDefaultHandler) SetDocumentLocatorDefault(locator QXmlLocator_ITF) {
	defer qt.Recovering("QXmlDefaultHandler::setDocumentLocator")

	if ptr.Pointer() != nil {
		C.QXmlDefaultHandler_SetDocumentLocatorDefault(ptr.Pointer(), PointerFromQXmlLocator(locator))
	}
}

func (ptr *QXmlDefaultHandler) SkippedEntity(name string) bool {
	defer qt.Recovering("QXmlDefaultHandler::skippedEntity")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_SkippedEntity(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) StartCDATA() bool {
	defer qt.Recovering("QXmlDefaultHandler::startCDATA")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_StartCDATA(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) StartDTD(name string, publicId string, systemId string) bool {
	defer qt.Recovering("QXmlDefaultHandler::startDTD")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_StartDTD(ptr.Pointer(), C.CString(name), C.CString(publicId), C.CString(systemId)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) StartDocument() bool {
	defer qt.Recovering("QXmlDefaultHandler::startDocument")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_StartDocument(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) StartElement(namespaceURI string, localName string, qName string, atts QXmlAttributes_ITF) bool {
	defer qt.Recovering("QXmlDefaultHandler::startElement")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_StartElement(ptr.Pointer(), C.CString(namespaceURI), C.CString(localName), C.CString(qName), PointerFromQXmlAttributes(atts)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) StartEntity(name string) bool {
	defer qt.Recovering("QXmlDefaultHandler::startEntity")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_StartEntity(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) StartPrefixMapping(prefix string, uri string) bool {
	defer qt.Recovering("QXmlDefaultHandler::startPrefixMapping")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_StartPrefixMapping(ptr.Pointer(), C.CString(prefix), C.CString(uri)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) UnparsedEntityDecl(name string, publicId string, systemId string, notationName string) bool {
	defer qt.Recovering("QXmlDefaultHandler::unparsedEntityDecl")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_UnparsedEntityDecl(ptr.Pointer(), C.CString(name), C.CString(publicId), C.CString(systemId), C.CString(notationName)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) Warning(exception QXmlParseException_ITF) bool {
	defer qt.Recovering("QXmlDefaultHandler::warning")

	if ptr.Pointer() != nil {
		return C.QXmlDefaultHandler_Warning(ptr.Pointer(), PointerFromQXmlParseException(exception)) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) ObjectNameAbs() string {
	defer qt.Recovering("QXmlDefaultHandler::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlDefaultHandler_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlDefaultHandler) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlDefaultHandler::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlDefaultHandler_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QXmlEntityResolver struct {
	ptr unsafe.Pointer
}

type QXmlEntityResolver_ITF interface {
	QXmlEntityResolver_PTR() *QXmlEntityResolver
}

func (p *QXmlEntityResolver) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlEntityResolver) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlEntityResolver(ptr QXmlEntityResolver_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlEntityResolver_PTR().Pointer()
	}
	return nil
}

func NewQXmlEntityResolverFromPointer(ptr unsafe.Pointer) *QXmlEntityResolver {
	var n = new(QXmlEntityResolver)
	n.SetPointer(ptr)
	return n
}

func newQXmlEntityResolverFromPointer(ptr unsafe.Pointer) *QXmlEntityResolver {
	var n = NewQXmlEntityResolverFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlEntityResolver_") {
		n.SetObjectNameAbs("QXmlEntityResolver_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlEntityResolver) QXmlEntityResolver_PTR() *QXmlEntityResolver {
	return ptr
}

func (ptr *QXmlEntityResolver) ErrorString() string {
	defer qt.Recovering("QXmlEntityResolver::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlEntityResolver_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlEntityResolver) DestroyQXmlEntityResolver() {
	defer qt.Recovering("QXmlEntityResolver::~QXmlEntityResolver")

	if ptr.Pointer() != nil {
		C.QXmlEntityResolver_DestroyQXmlEntityResolver(ptr.Pointer())
	}
}

func (ptr *QXmlEntityResolver) ObjectNameAbs() string {
	defer qt.Recovering("QXmlEntityResolver::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlEntityResolver_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlEntityResolver) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlEntityResolver::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlEntityResolver_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QXmlErrorHandler struct {
	ptr unsafe.Pointer
}

type QXmlErrorHandler_ITF interface {
	QXmlErrorHandler_PTR() *QXmlErrorHandler
}

func (p *QXmlErrorHandler) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlErrorHandler) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlErrorHandler(ptr QXmlErrorHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlErrorHandler_PTR().Pointer()
	}
	return nil
}

func NewQXmlErrorHandlerFromPointer(ptr unsafe.Pointer) *QXmlErrorHandler {
	var n = new(QXmlErrorHandler)
	n.SetPointer(ptr)
	return n
}

func newQXmlErrorHandlerFromPointer(ptr unsafe.Pointer) *QXmlErrorHandler {
	var n = NewQXmlErrorHandlerFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlErrorHandler_") {
		n.SetObjectNameAbs("QXmlErrorHandler_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlErrorHandler) QXmlErrorHandler_PTR() *QXmlErrorHandler {
	return ptr
}

func (ptr *QXmlErrorHandler) Error(exception QXmlParseException_ITF) bool {
	defer qt.Recovering("QXmlErrorHandler::error")

	if ptr.Pointer() != nil {
		return C.QXmlErrorHandler_Error(ptr.Pointer(), PointerFromQXmlParseException(exception)) != 0
	}
	return false
}

func (ptr *QXmlErrorHandler) ErrorString() string {
	defer qt.Recovering("QXmlErrorHandler::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlErrorHandler_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlErrorHandler) FatalError(exception QXmlParseException_ITF) bool {
	defer qt.Recovering("QXmlErrorHandler::fatalError")

	if ptr.Pointer() != nil {
		return C.QXmlErrorHandler_FatalError(ptr.Pointer(), PointerFromQXmlParseException(exception)) != 0
	}
	return false
}

func (ptr *QXmlErrorHandler) Warning(exception QXmlParseException_ITF) bool {
	defer qt.Recovering("QXmlErrorHandler::warning")

	if ptr.Pointer() != nil {
		return C.QXmlErrorHandler_Warning(ptr.Pointer(), PointerFromQXmlParseException(exception)) != 0
	}
	return false
}

func (ptr *QXmlErrorHandler) DestroyQXmlErrorHandler() {
	defer qt.Recovering("QXmlErrorHandler::~QXmlErrorHandler")

	if ptr.Pointer() != nil {
		C.QXmlErrorHandler_DestroyQXmlErrorHandler(ptr.Pointer())
	}
}

func (ptr *QXmlErrorHandler) ObjectNameAbs() string {
	defer qt.Recovering("QXmlErrorHandler::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlErrorHandler_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlErrorHandler) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlErrorHandler::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlErrorHandler_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QXmlInputSource struct {
	ptr unsafe.Pointer
}

type QXmlInputSource_ITF interface {
	QXmlInputSource_PTR() *QXmlInputSource
}

func (p *QXmlInputSource) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlInputSource) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlInputSource(ptr QXmlInputSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlInputSource_PTR().Pointer()
	}
	return nil
}

func NewQXmlInputSourceFromPointer(ptr unsafe.Pointer) *QXmlInputSource {
	var n = new(QXmlInputSource)
	n.SetPointer(ptr)
	return n
}

func newQXmlInputSourceFromPointer(ptr unsafe.Pointer) *QXmlInputSource {
	var n = NewQXmlInputSourceFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlInputSource_") {
		n.SetObjectNameAbs("QXmlInputSource_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlInputSource) QXmlInputSource_PTR() *QXmlInputSource {
	return ptr
}

func NewQXmlInputSource() *QXmlInputSource {
	defer qt.Recovering("QXmlInputSource::QXmlInputSource")

	return newQXmlInputSourceFromPointer(C.QXmlInputSource_NewQXmlInputSource())
}

func NewQXmlInputSource2(dev core.QIODevice_ITF) *QXmlInputSource {
	defer qt.Recovering("QXmlInputSource::QXmlInputSource")

	return newQXmlInputSourceFromPointer(C.QXmlInputSource_NewQXmlInputSource2(core.PointerFromQIODevice(dev)))
}

func (ptr *QXmlInputSource) Data() string {
	defer qt.Recovering("QXmlInputSource::data")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlInputSource_Data(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlInputSource) ConnectFetchData(f func()) {
	defer qt.Recovering("connect QXmlInputSource::fetchData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "fetchData", f)
	}
}

func (ptr *QXmlInputSource) DisconnectFetchData() {
	defer qt.Recovering("disconnect QXmlInputSource::fetchData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "fetchData")
	}
}

//export callbackQXmlInputSourceFetchData
func callbackQXmlInputSourceFetchData(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlInputSource::fetchData")

	if signal := qt.GetSignal(C.GoString(ptrName), "fetchData"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlInputSourceFromPointer(ptr).FetchDataDefault()
	}
}

func (ptr *QXmlInputSource) FetchData() {
	defer qt.Recovering("QXmlInputSource::fetchData")

	if ptr.Pointer() != nil {
		C.QXmlInputSource_FetchData(ptr.Pointer())
	}
}

func (ptr *QXmlInputSource) FetchDataDefault() {
	defer qt.Recovering("QXmlInputSource::fetchData")

	if ptr.Pointer() != nil {
		C.QXmlInputSource_FetchDataDefault(ptr.Pointer())
	}
}

func (ptr *QXmlInputSource) FromRawData(data core.QByteArray_ITF, beginning bool) string {
	defer qt.Recovering("QXmlInputSource::fromRawData")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlInputSource_FromRawData(ptr.Pointer(), core.PointerFromQByteArray(data), C.int(qt.GoBoolToInt(beginning))))
	}
	return ""
}

func (ptr *QXmlInputSource) ConnectReset(f func()) {
	defer qt.Recovering("connect QXmlInputSource::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "reset", f)
	}
}

func (ptr *QXmlInputSource) DisconnectReset() {
	defer qt.Recovering("disconnect QXmlInputSource::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "reset")
	}
}

//export callbackQXmlInputSourceReset
func callbackQXmlInputSourceReset(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlInputSource::reset")

	if signal := qt.GetSignal(C.GoString(ptrName), "reset"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlInputSourceFromPointer(ptr).ResetDefault()
	}
}

func (ptr *QXmlInputSource) Reset() {
	defer qt.Recovering("QXmlInputSource::reset")

	if ptr.Pointer() != nil {
		C.QXmlInputSource_Reset(ptr.Pointer())
	}
}

func (ptr *QXmlInputSource) ResetDefault() {
	defer qt.Recovering("QXmlInputSource::reset")

	if ptr.Pointer() != nil {
		C.QXmlInputSource_ResetDefault(ptr.Pointer())
	}
}

func (ptr *QXmlInputSource) ConnectSetData(f func(dat string)) {
	defer qt.Recovering("connect QXmlInputSource::setData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setData", f)
	}
}

func (ptr *QXmlInputSource) DisconnectSetData() {
	defer qt.Recovering("disconnect QXmlInputSource::setData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setData")
	}
}

//export callbackQXmlInputSourceSetData
func callbackQXmlInputSourceSetData(ptr unsafe.Pointer, ptrName *C.char, dat *C.char) {
	defer qt.Recovering("callback QXmlInputSource::setData")

	if signal := qt.GetSignal(C.GoString(ptrName), "setData"); signal != nil {
		signal.(func(string))(C.GoString(dat))
	} else {
		NewQXmlInputSourceFromPointer(ptr).SetDataDefault(C.GoString(dat))
	}
}

func (ptr *QXmlInputSource) SetData(dat string) {
	defer qt.Recovering("QXmlInputSource::setData")

	if ptr.Pointer() != nil {
		C.QXmlInputSource_SetData(ptr.Pointer(), C.CString(dat))
	}
}

func (ptr *QXmlInputSource) SetDataDefault(dat string) {
	defer qt.Recovering("QXmlInputSource::setData")

	if ptr.Pointer() != nil {
		C.QXmlInputSource_SetDataDefault(ptr.Pointer(), C.CString(dat))
	}
}

func (ptr *QXmlInputSource) DestroyQXmlInputSource() {
	defer qt.Recovering("QXmlInputSource::~QXmlInputSource")

	if ptr.Pointer() != nil {
		C.QXmlInputSource_DestroyQXmlInputSource(ptr.Pointer())
	}
}

func (ptr *QXmlInputSource) ObjectNameAbs() string {
	defer qt.Recovering("QXmlInputSource::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlInputSource_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlInputSource) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlInputSource::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlInputSource_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QXmlLexicalHandler struct {
	ptr unsafe.Pointer
}

type QXmlLexicalHandler_ITF interface {
	QXmlLexicalHandler_PTR() *QXmlLexicalHandler
}

func (p *QXmlLexicalHandler) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlLexicalHandler) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlLexicalHandler(ptr QXmlLexicalHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlLexicalHandler_PTR().Pointer()
	}
	return nil
}

func NewQXmlLexicalHandlerFromPointer(ptr unsafe.Pointer) *QXmlLexicalHandler {
	var n = new(QXmlLexicalHandler)
	n.SetPointer(ptr)
	return n
}

func newQXmlLexicalHandlerFromPointer(ptr unsafe.Pointer) *QXmlLexicalHandler {
	var n = NewQXmlLexicalHandlerFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlLexicalHandler_") {
		n.SetObjectNameAbs("QXmlLexicalHandler_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlLexicalHandler) QXmlLexicalHandler_PTR() *QXmlLexicalHandler {
	return ptr
}

func (ptr *QXmlLexicalHandler) Comment(ch string) bool {
	defer qt.Recovering("QXmlLexicalHandler::comment")

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_Comment(ptr.Pointer(), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) EndCDATA() bool {
	defer qt.Recovering("QXmlLexicalHandler::endCDATA")

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_EndCDATA(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) EndDTD() bool {
	defer qt.Recovering("QXmlLexicalHandler::endDTD")

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_EndDTD(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) EndEntity(name string) bool {
	defer qt.Recovering("QXmlLexicalHandler::endEntity")

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_EndEntity(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) ErrorString() string {
	defer qt.Recovering("QXmlLexicalHandler::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlLexicalHandler_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlLexicalHandler) StartCDATA() bool {
	defer qt.Recovering("QXmlLexicalHandler::startCDATA")

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_StartCDATA(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) StartDTD(name string, publicId string, systemId string) bool {
	defer qt.Recovering("QXmlLexicalHandler::startDTD")

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_StartDTD(ptr.Pointer(), C.CString(name), C.CString(publicId), C.CString(systemId)) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) StartEntity(name string) bool {
	defer qt.Recovering("QXmlLexicalHandler::startEntity")

	if ptr.Pointer() != nil {
		return C.QXmlLexicalHandler_StartEntity(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlLexicalHandler) DestroyQXmlLexicalHandler() {
	defer qt.Recovering("QXmlLexicalHandler::~QXmlLexicalHandler")

	if ptr.Pointer() != nil {
		C.QXmlLexicalHandler_DestroyQXmlLexicalHandler(ptr.Pointer())
	}
}

func (ptr *QXmlLexicalHandler) ObjectNameAbs() string {
	defer qt.Recovering("QXmlLexicalHandler::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlLexicalHandler_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlLexicalHandler) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlLexicalHandler::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlLexicalHandler_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QXmlLocator struct {
	ptr unsafe.Pointer
}

type QXmlLocator_ITF interface {
	QXmlLocator_PTR() *QXmlLocator
}

func (p *QXmlLocator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlLocator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlLocator(ptr QXmlLocator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlLocator_PTR().Pointer()
	}
	return nil
}

func NewQXmlLocatorFromPointer(ptr unsafe.Pointer) *QXmlLocator {
	var n = new(QXmlLocator)
	n.SetPointer(ptr)
	return n
}

func newQXmlLocatorFromPointer(ptr unsafe.Pointer) *QXmlLocator {
	var n = NewQXmlLocatorFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlLocator_") {
		n.SetObjectNameAbs("QXmlLocator_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlLocator) QXmlLocator_PTR() *QXmlLocator {
	return ptr
}

func (ptr *QXmlLocator) ColumnNumber() int {
	defer qt.Recovering("QXmlLocator::columnNumber")

	if ptr.Pointer() != nil {
		return int(C.QXmlLocator_ColumnNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlLocator) LineNumber() int {
	defer qt.Recovering("QXmlLocator::lineNumber")

	if ptr.Pointer() != nil {
		return int(C.QXmlLocator_LineNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlLocator) DestroyQXmlLocator() {
	defer qt.Recovering("QXmlLocator::~QXmlLocator")

	if ptr.Pointer() != nil {
		C.QXmlLocator_DestroyQXmlLocator(ptr.Pointer())
	}
}

func (ptr *QXmlLocator) ObjectNameAbs() string {
	defer qt.Recovering("QXmlLocator::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlLocator_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlLocator) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlLocator::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlLocator_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QXmlNamespaceSupport struct {
	ptr unsafe.Pointer
}

type QXmlNamespaceSupport_ITF interface {
	QXmlNamespaceSupport_PTR() *QXmlNamespaceSupport
}

func (p *QXmlNamespaceSupport) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlNamespaceSupport) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlNamespaceSupport(ptr QXmlNamespaceSupport_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlNamespaceSupport_PTR().Pointer()
	}
	return nil
}

func NewQXmlNamespaceSupportFromPointer(ptr unsafe.Pointer) *QXmlNamespaceSupport {
	var n = new(QXmlNamespaceSupport)
	n.SetPointer(ptr)
	return n
}

func newQXmlNamespaceSupportFromPointer(ptr unsafe.Pointer) *QXmlNamespaceSupport {
	var n = NewQXmlNamespaceSupportFromPointer(ptr)
	return n
}

func (ptr *QXmlNamespaceSupport) QXmlNamespaceSupport_PTR() *QXmlNamespaceSupport {
	return ptr
}

func NewQXmlNamespaceSupport() *QXmlNamespaceSupport {
	defer qt.Recovering("QXmlNamespaceSupport::QXmlNamespaceSupport")

	return newQXmlNamespaceSupportFromPointer(C.QXmlNamespaceSupport_NewQXmlNamespaceSupport())
}

func (ptr *QXmlNamespaceSupport) PopContext() {
	defer qt.Recovering("QXmlNamespaceSupport::popContext")

	if ptr.Pointer() != nil {
		C.QXmlNamespaceSupport_PopContext(ptr.Pointer())
	}
}

func (ptr *QXmlNamespaceSupport) Prefix(uri string) string {
	defer qt.Recovering("QXmlNamespaceSupport::prefix")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlNamespaceSupport_Prefix(ptr.Pointer(), C.CString(uri)))
	}
	return ""
}

func (ptr *QXmlNamespaceSupport) Prefixes() []string {
	defer qt.Recovering("QXmlNamespaceSupport::prefixes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QXmlNamespaceSupport_Prefixes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QXmlNamespaceSupport) Prefixes2(uri string) []string {
	defer qt.Recovering("QXmlNamespaceSupport::prefixes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QXmlNamespaceSupport_Prefixes2(ptr.Pointer(), C.CString(uri))), "|")
	}
	return make([]string, 0)
}

func (ptr *QXmlNamespaceSupport) PushContext() {
	defer qt.Recovering("QXmlNamespaceSupport::pushContext")

	if ptr.Pointer() != nil {
		C.QXmlNamespaceSupport_PushContext(ptr.Pointer())
	}
}

func (ptr *QXmlNamespaceSupport) Reset() {
	defer qt.Recovering("QXmlNamespaceSupport::reset")

	if ptr.Pointer() != nil {
		C.QXmlNamespaceSupport_Reset(ptr.Pointer())
	}
}

func (ptr *QXmlNamespaceSupport) SetPrefix(pre string, uri string) {
	defer qt.Recovering("QXmlNamespaceSupport::setPrefix")

	if ptr.Pointer() != nil {
		C.QXmlNamespaceSupport_SetPrefix(ptr.Pointer(), C.CString(pre), C.CString(uri))
	}
}

func (ptr *QXmlNamespaceSupport) Uri(prefix string) string {
	defer qt.Recovering("QXmlNamespaceSupport::uri")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlNamespaceSupport_Uri(ptr.Pointer(), C.CString(prefix)))
	}
	return ""
}

func (ptr *QXmlNamespaceSupport) DestroyQXmlNamespaceSupport() {
	defer qt.Recovering("QXmlNamespaceSupport::~QXmlNamespaceSupport")

	if ptr.Pointer() != nil {
		C.QXmlNamespaceSupport_DestroyQXmlNamespaceSupport(ptr.Pointer())
	}
}

type QXmlParseException struct {
	ptr unsafe.Pointer
}

type QXmlParseException_ITF interface {
	QXmlParseException_PTR() *QXmlParseException
}

func (p *QXmlParseException) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlParseException) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlParseException(ptr QXmlParseException_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlParseException_PTR().Pointer()
	}
	return nil
}

func NewQXmlParseExceptionFromPointer(ptr unsafe.Pointer) *QXmlParseException {
	var n = new(QXmlParseException)
	n.SetPointer(ptr)
	return n
}

func newQXmlParseExceptionFromPointer(ptr unsafe.Pointer) *QXmlParseException {
	var n = NewQXmlParseExceptionFromPointer(ptr)
	return n
}

func (ptr *QXmlParseException) QXmlParseException_PTR() *QXmlParseException {
	return ptr
}

func NewQXmlParseException(name string, c int, l int, p string, s string) *QXmlParseException {
	defer qt.Recovering("QXmlParseException::QXmlParseException")

	return newQXmlParseExceptionFromPointer(C.QXmlParseException_NewQXmlParseException(C.CString(name), C.int(c), C.int(l), C.CString(p), C.CString(s)))
}

func NewQXmlParseException2(other QXmlParseException_ITF) *QXmlParseException {
	defer qt.Recovering("QXmlParseException::QXmlParseException")

	return newQXmlParseExceptionFromPointer(C.QXmlParseException_NewQXmlParseException2(PointerFromQXmlParseException(other)))
}

func (ptr *QXmlParseException) ColumnNumber() int {
	defer qt.Recovering("QXmlParseException::columnNumber")

	if ptr.Pointer() != nil {
		return int(C.QXmlParseException_ColumnNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlParseException) LineNumber() int {
	defer qt.Recovering("QXmlParseException::lineNumber")

	if ptr.Pointer() != nil {
		return int(C.QXmlParseException_LineNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlParseException) Message() string {
	defer qt.Recovering("QXmlParseException::message")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlParseException_Message(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlParseException) PublicId() string {
	defer qt.Recovering("QXmlParseException::publicId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlParseException_PublicId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlParseException) SystemId() string {
	defer qt.Recovering("QXmlParseException::systemId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlParseException_SystemId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlParseException) DestroyQXmlParseException() {
	defer qt.Recovering("QXmlParseException::~QXmlParseException")

	if ptr.Pointer() != nil {
		C.QXmlParseException_DestroyQXmlParseException(ptr.Pointer())
	}
}

type QXmlReader struct {
	ptr unsafe.Pointer
}

type QXmlReader_ITF interface {
	QXmlReader_PTR() *QXmlReader
}

func (p *QXmlReader) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlReader) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlReader(ptr QXmlReader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlReader_PTR().Pointer()
	}
	return nil
}

func NewQXmlReaderFromPointer(ptr unsafe.Pointer) *QXmlReader {
	var n = new(QXmlReader)
	n.SetPointer(ptr)
	return n
}

func newQXmlReaderFromPointer(ptr unsafe.Pointer) *QXmlReader {
	var n = NewQXmlReaderFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlReader_") {
		n.SetObjectNameAbs("QXmlReader_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlReader) QXmlReader_PTR() *QXmlReader {
	return ptr
}

func (ptr *QXmlReader) DTDHandler() *QXmlDTDHandler {
	defer qt.Recovering("QXmlReader::DTDHandler")

	if ptr.Pointer() != nil {
		return NewQXmlDTDHandlerFromPointer(C.QXmlReader_DTDHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlReader) ContentHandler() *QXmlContentHandler {
	defer qt.Recovering("QXmlReader::contentHandler")

	if ptr.Pointer() != nil {
		return NewQXmlContentHandlerFromPointer(C.QXmlReader_ContentHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlReader) DeclHandler() *QXmlDeclHandler {
	defer qt.Recovering("QXmlReader::declHandler")

	if ptr.Pointer() != nil {
		return NewQXmlDeclHandlerFromPointer(C.QXmlReader_DeclHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlReader) EntityResolver() *QXmlEntityResolver {
	defer qt.Recovering("QXmlReader::entityResolver")

	if ptr.Pointer() != nil {
		return NewQXmlEntityResolverFromPointer(C.QXmlReader_EntityResolver(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlReader) ErrorHandler() *QXmlErrorHandler {
	defer qt.Recovering("QXmlReader::errorHandler")

	if ptr.Pointer() != nil {
		return NewQXmlErrorHandlerFromPointer(C.QXmlReader_ErrorHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlReader) Feature(name string, ok bool) bool {
	defer qt.Recovering("QXmlReader::feature")

	if ptr.Pointer() != nil {
		return C.QXmlReader_Feature(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(ok))) != 0
	}
	return false
}

func (ptr *QXmlReader) HasFeature(name string) bool {
	defer qt.Recovering("QXmlReader::hasFeature")

	if ptr.Pointer() != nil {
		return C.QXmlReader_HasFeature(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlReader) HasProperty(name string) bool {
	defer qt.Recovering("QXmlReader::hasProperty")

	if ptr.Pointer() != nil {
		return C.QXmlReader_HasProperty(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlReader) LexicalHandler() *QXmlLexicalHandler {
	defer qt.Recovering("QXmlReader::lexicalHandler")

	if ptr.Pointer() != nil {
		return NewQXmlLexicalHandlerFromPointer(C.QXmlReader_LexicalHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlReader) Parse(input QXmlInputSource_ITF) bool {
	defer qt.Recovering("QXmlReader::parse")

	if ptr.Pointer() != nil {
		return C.QXmlReader_Parse(ptr.Pointer(), PointerFromQXmlInputSource(input)) != 0
	}
	return false
}

func (ptr *QXmlReader) Property(name string, ok bool) unsafe.Pointer {
	defer qt.Recovering("QXmlReader::property")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QXmlReader_Property(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(ok))))
	}
	return nil
}

func (ptr *QXmlReader) SetContentHandler(handler QXmlContentHandler_ITF) {
	defer qt.Recovering("QXmlReader::setContentHandler")

	if ptr.Pointer() != nil {
		C.QXmlReader_SetContentHandler(ptr.Pointer(), PointerFromQXmlContentHandler(handler))
	}
}

func (ptr *QXmlReader) SetDTDHandler(handler QXmlDTDHandler_ITF) {
	defer qt.Recovering("QXmlReader::setDTDHandler")

	if ptr.Pointer() != nil {
		C.QXmlReader_SetDTDHandler(ptr.Pointer(), PointerFromQXmlDTDHandler(handler))
	}
}

func (ptr *QXmlReader) SetDeclHandler(handler QXmlDeclHandler_ITF) {
	defer qt.Recovering("QXmlReader::setDeclHandler")

	if ptr.Pointer() != nil {
		C.QXmlReader_SetDeclHandler(ptr.Pointer(), PointerFromQXmlDeclHandler(handler))
	}
}

func (ptr *QXmlReader) SetEntityResolver(handler QXmlEntityResolver_ITF) {
	defer qt.Recovering("QXmlReader::setEntityResolver")

	if ptr.Pointer() != nil {
		C.QXmlReader_SetEntityResolver(ptr.Pointer(), PointerFromQXmlEntityResolver(handler))
	}
}

func (ptr *QXmlReader) SetErrorHandler(handler QXmlErrorHandler_ITF) {
	defer qt.Recovering("QXmlReader::setErrorHandler")

	if ptr.Pointer() != nil {
		C.QXmlReader_SetErrorHandler(ptr.Pointer(), PointerFromQXmlErrorHandler(handler))
	}
}

func (ptr *QXmlReader) SetFeature(name string, value bool) {
	defer qt.Recovering("QXmlReader::setFeature")

	if ptr.Pointer() != nil {
		C.QXmlReader_SetFeature(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(value)))
	}
}

func (ptr *QXmlReader) SetLexicalHandler(handler QXmlLexicalHandler_ITF) {
	defer qt.Recovering("QXmlReader::setLexicalHandler")

	if ptr.Pointer() != nil {
		C.QXmlReader_SetLexicalHandler(ptr.Pointer(), PointerFromQXmlLexicalHandler(handler))
	}
}

func (ptr *QXmlReader) DestroyQXmlReader() {
	defer qt.Recovering("QXmlReader::~QXmlReader")

	if ptr.Pointer() != nil {
		C.QXmlReader_DestroyQXmlReader(ptr.Pointer())
	}
}

func (ptr *QXmlReader) ObjectNameAbs() string {
	defer qt.Recovering("QXmlReader::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlReader_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlReader) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlReader::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlReader_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QXmlSimpleReader struct {
	QXmlReader
}

type QXmlSimpleReader_ITF interface {
	QXmlReader_ITF
	QXmlSimpleReader_PTR() *QXmlSimpleReader
}

func PointerFromQXmlSimpleReader(ptr QXmlSimpleReader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlSimpleReader_PTR().Pointer()
	}
	return nil
}

func NewQXmlSimpleReaderFromPointer(ptr unsafe.Pointer) *QXmlSimpleReader {
	var n = new(QXmlSimpleReader)
	n.SetPointer(ptr)
	return n
}

func newQXmlSimpleReaderFromPointer(ptr unsafe.Pointer) *QXmlSimpleReader {
	var n = NewQXmlSimpleReaderFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlSimpleReader_") {
		n.SetObjectNameAbs("QXmlSimpleReader_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlSimpleReader) QXmlSimpleReader_PTR() *QXmlSimpleReader {
	return ptr
}

func (ptr *QXmlSimpleReader) DTDHandler() *QXmlDTDHandler {
	defer qt.Recovering("QXmlSimpleReader::DTDHandler")

	if ptr.Pointer() != nil {
		return NewQXmlDTDHandlerFromPointer(C.QXmlSimpleReader_DTDHandler(ptr.Pointer()))
	}
	return nil
}

func NewQXmlSimpleReader() *QXmlSimpleReader {
	defer qt.Recovering("QXmlSimpleReader::QXmlSimpleReader")

	return newQXmlSimpleReaderFromPointer(C.QXmlSimpleReader_NewQXmlSimpleReader())
}

func (ptr *QXmlSimpleReader) ContentHandler() *QXmlContentHandler {
	defer qt.Recovering("QXmlSimpleReader::contentHandler")

	if ptr.Pointer() != nil {
		return NewQXmlContentHandlerFromPointer(C.QXmlSimpleReader_ContentHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSimpleReader) DeclHandler() *QXmlDeclHandler {
	defer qt.Recovering("QXmlSimpleReader::declHandler")

	if ptr.Pointer() != nil {
		return NewQXmlDeclHandlerFromPointer(C.QXmlSimpleReader_DeclHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSimpleReader) EntityResolver() *QXmlEntityResolver {
	defer qt.Recovering("QXmlSimpleReader::entityResolver")

	if ptr.Pointer() != nil {
		return NewQXmlEntityResolverFromPointer(C.QXmlSimpleReader_EntityResolver(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSimpleReader) ErrorHandler() *QXmlErrorHandler {
	defer qt.Recovering("QXmlSimpleReader::errorHandler")

	if ptr.Pointer() != nil {
		return NewQXmlErrorHandlerFromPointer(C.QXmlSimpleReader_ErrorHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSimpleReader) Feature(name string, ok bool) bool {
	defer qt.Recovering("QXmlSimpleReader::feature")

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_Feature(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(ok))) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) HasFeature(name string) bool {
	defer qt.Recovering("QXmlSimpleReader::hasFeature")

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_HasFeature(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) HasProperty(name string) bool {
	defer qt.Recovering("QXmlSimpleReader::hasProperty")

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_HasProperty(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) LexicalHandler() *QXmlLexicalHandler {
	defer qt.Recovering("QXmlSimpleReader::lexicalHandler")

	if ptr.Pointer() != nil {
		return NewQXmlLexicalHandlerFromPointer(C.QXmlSimpleReader_LexicalHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSimpleReader) Parse(input QXmlInputSource_ITF) bool {
	defer qt.Recovering("QXmlSimpleReader::parse")

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_Parse(ptr.Pointer(), PointerFromQXmlInputSource(input)) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) ParseContinue() bool {
	defer qt.Recovering("QXmlSimpleReader::parseContinue")

	if ptr.Pointer() != nil {
		return C.QXmlSimpleReader_ParseContinue(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) Property(name string, ok bool) unsafe.Pointer {
	defer qt.Recovering("QXmlSimpleReader::property")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QXmlSimpleReader_Property(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(ok))))
	}
	return nil
}

func (ptr *QXmlSimpleReader) ConnectSetContentHandler(f func(handler *QXmlContentHandler)) {
	defer qt.Recovering("connect QXmlSimpleReader::setContentHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setContentHandler", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetContentHandler() {
	defer qt.Recovering("disconnect QXmlSimpleReader::setContentHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setContentHandler")
	}
}

//export callbackQXmlSimpleReaderSetContentHandler
func callbackQXmlSimpleReaderSetContentHandler(ptr unsafe.Pointer, ptrName *C.char, handler unsafe.Pointer) {
	defer qt.Recovering("callback QXmlSimpleReader::setContentHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "setContentHandler"); signal != nil {
		signal.(func(*QXmlContentHandler))(NewQXmlContentHandlerFromPointer(handler))
	}

}

func (ptr *QXmlSimpleReader) SetContentHandler(handler QXmlContentHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setContentHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetContentHandler(ptr.Pointer(), PointerFromQXmlContentHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) SetContentHandlerDefault(handler QXmlContentHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setContentHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetContentHandlerDefault(ptr.Pointer(), PointerFromQXmlContentHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) ConnectSetDTDHandler(f func(handler *QXmlDTDHandler)) {
	defer qt.Recovering("connect QXmlSimpleReader::setDTDHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setDTDHandler", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetDTDHandler() {
	defer qt.Recovering("disconnect QXmlSimpleReader::setDTDHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setDTDHandler")
	}
}

//export callbackQXmlSimpleReaderSetDTDHandler
func callbackQXmlSimpleReaderSetDTDHandler(ptr unsafe.Pointer, ptrName *C.char, handler unsafe.Pointer) {
	defer qt.Recovering("callback QXmlSimpleReader::setDTDHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "setDTDHandler"); signal != nil {
		signal.(func(*QXmlDTDHandler))(NewQXmlDTDHandlerFromPointer(handler))
	}

}

func (ptr *QXmlSimpleReader) SetDTDHandler(handler QXmlDTDHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setDTDHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetDTDHandler(ptr.Pointer(), PointerFromQXmlDTDHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) SetDTDHandlerDefault(handler QXmlDTDHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setDTDHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetDTDHandlerDefault(ptr.Pointer(), PointerFromQXmlDTDHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) ConnectSetDeclHandler(f func(handler *QXmlDeclHandler)) {
	defer qt.Recovering("connect QXmlSimpleReader::setDeclHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setDeclHandler", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetDeclHandler() {
	defer qt.Recovering("disconnect QXmlSimpleReader::setDeclHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setDeclHandler")
	}
}

//export callbackQXmlSimpleReaderSetDeclHandler
func callbackQXmlSimpleReaderSetDeclHandler(ptr unsafe.Pointer, ptrName *C.char, handler unsafe.Pointer) {
	defer qt.Recovering("callback QXmlSimpleReader::setDeclHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "setDeclHandler"); signal != nil {
		signal.(func(*QXmlDeclHandler))(NewQXmlDeclHandlerFromPointer(handler))
	}

}

func (ptr *QXmlSimpleReader) SetDeclHandler(handler QXmlDeclHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setDeclHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetDeclHandler(ptr.Pointer(), PointerFromQXmlDeclHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) SetDeclHandlerDefault(handler QXmlDeclHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setDeclHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetDeclHandlerDefault(ptr.Pointer(), PointerFromQXmlDeclHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) ConnectSetEntityResolver(f func(handler *QXmlEntityResolver)) {
	defer qt.Recovering("connect QXmlSimpleReader::setEntityResolver")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setEntityResolver", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetEntityResolver() {
	defer qt.Recovering("disconnect QXmlSimpleReader::setEntityResolver")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setEntityResolver")
	}
}

//export callbackQXmlSimpleReaderSetEntityResolver
func callbackQXmlSimpleReaderSetEntityResolver(ptr unsafe.Pointer, ptrName *C.char, handler unsafe.Pointer) {
	defer qt.Recovering("callback QXmlSimpleReader::setEntityResolver")

	if signal := qt.GetSignal(C.GoString(ptrName), "setEntityResolver"); signal != nil {
		signal.(func(*QXmlEntityResolver))(NewQXmlEntityResolverFromPointer(handler))
	}

}

func (ptr *QXmlSimpleReader) SetEntityResolver(handler QXmlEntityResolver_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setEntityResolver")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetEntityResolver(ptr.Pointer(), PointerFromQXmlEntityResolver(handler))
	}
}

func (ptr *QXmlSimpleReader) SetEntityResolverDefault(handler QXmlEntityResolver_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setEntityResolver")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetEntityResolverDefault(ptr.Pointer(), PointerFromQXmlEntityResolver(handler))
	}
}

func (ptr *QXmlSimpleReader) ConnectSetErrorHandler(f func(handler *QXmlErrorHandler)) {
	defer qt.Recovering("connect QXmlSimpleReader::setErrorHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setErrorHandler", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetErrorHandler() {
	defer qt.Recovering("disconnect QXmlSimpleReader::setErrorHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setErrorHandler")
	}
}

//export callbackQXmlSimpleReaderSetErrorHandler
func callbackQXmlSimpleReaderSetErrorHandler(ptr unsafe.Pointer, ptrName *C.char, handler unsafe.Pointer) {
	defer qt.Recovering("callback QXmlSimpleReader::setErrorHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "setErrorHandler"); signal != nil {
		signal.(func(*QXmlErrorHandler))(NewQXmlErrorHandlerFromPointer(handler))
	}

}

func (ptr *QXmlSimpleReader) SetErrorHandler(handler QXmlErrorHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setErrorHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetErrorHandler(ptr.Pointer(), PointerFromQXmlErrorHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) SetErrorHandlerDefault(handler QXmlErrorHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setErrorHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetErrorHandlerDefault(ptr.Pointer(), PointerFromQXmlErrorHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) ConnectSetFeature(f func(name string, enable bool)) {
	defer qt.Recovering("connect QXmlSimpleReader::setFeature")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setFeature", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetFeature() {
	defer qt.Recovering("disconnect QXmlSimpleReader::setFeature")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setFeature")
	}
}

//export callbackQXmlSimpleReaderSetFeature
func callbackQXmlSimpleReaderSetFeature(ptr unsafe.Pointer, ptrName *C.char, name *C.char, enable C.int) {
	defer qt.Recovering("callback QXmlSimpleReader::setFeature")

	if signal := qt.GetSignal(C.GoString(ptrName), "setFeature"); signal != nil {
		signal.(func(string, bool))(C.GoString(name), int(enable) != 0)
	}

}

func (ptr *QXmlSimpleReader) SetFeature(name string, enable bool) {
	defer qt.Recovering("QXmlSimpleReader::setFeature")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetFeature(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QXmlSimpleReader) SetFeatureDefault(name string, enable bool) {
	defer qt.Recovering("QXmlSimpleReader::setFeature")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetFeatureDefault(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QXmlSimpleReader) ConnectSetLexicalHandler(f func(handler *QXmlLexicalHandler)) {
	defer qt.Recovering("connect QXmlSimpleReader::setLexicalHandler")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setLexicalHandler", f)
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetLexicalHandler() {
	defer qt.Recovering("disconnect QXmlSimpleReader::setLexicalHandler")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setLexicalHandler")
	}
}

//export callbackQXmlSimpleReaderSetLexicalHandler
func callbackQXmlSimpleReaderSetLexicalHandler(ptr unsafe.Pointer, ptrName *C.char, handler unsafe.Pointer) {
	defer qt.Recovering("callback QXmlSimpleReader::setLexicalHandler")

	if signal := qt.GetSignal(C.GoString(ptrName), "setLexicalHandler"); signal != nil {
		signal.(func(*QXmlLexicalHandler))(NewQXmlLexicalHandlerFromPointer(handler))
	}

}

func (ptr *QXmlSimpleReader) SetLexicalHandler(handler QXmlLexicalHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setLexicalHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetLexicalHandler(ptr.Pointer(), PointerFromQXmlLexicalHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) SetLexicalHandlerDefault(handler QXmlLexicalHandler_ITF) {
	defer qt.Recovering("QXmlSimpleReader::setLexicalHandler")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetLexicalHandlerDefault(ptr.Pointer(), PointerFromQXmlLexicalHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) DestroyQXmlSimpleReader() {
	defer qt.Recovering("QXmlSimpleReader::~QXmlSimpleReader")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_DestroyQXmlSimpleReader(ptr.Pointer())
	}
}

func (ptr *QXmlSimpleReader) ObjectNameAbs() string {
	defer qt.Recovering("QXmlSimpleReader::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlSimpleReader_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlSimpleReader) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlSimpleReader::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
