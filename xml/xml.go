// +build !minimal

package xml

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/internal"
	"strings"
	"unsafe"
)

func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type QDomAttr struct {
	QDomNode
}

type QDomAttr_ITF interface {
	QDomNode_ITF
	QDomAttr_PTR() *QDomAttr
}

func (ptr *QDomAttr) QDomAttr_PTR() *QDomAttr {
	return ptr
}

func (ptr *QDomAttr) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QDomAttr) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDomNode_PTR().SetPointer(p)
	}
}

func PointerFromQDomAttr(ptr QDomAttr_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomAttr_PTR().Pointer()
	}
	return nil
}

func (n *QDomAttr) InitFromInternal(ptr uintptr, name string) {
	n.QDomNode_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDomAttr) ClassNameInternalF() string {
	return n.QDomNode_PTR().ClassNameInternalF()
}

func NewQDomAttrFromPointer(ptr unsafe.Pointer) (n *QDomAttr) {
	n = new(QDomAttr)
	n.InitFromInternal(uintptr(ptr), "xml.QDomAttr")
	return
}

func (ptr *QDomAttr) DestroyQDomAttr() {
}

func NewQDomAttr() *QDomAttr {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomAttr", ""}).(*QDomAttr)
}

func NewQDomAttr2(x QDomAttr_ITF) *QDomAttr {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomAttr2", "", x}).(*QDomAttr)
}

func (ptr *QDomAttr) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QDomAttr) OwnerElement() *QDomElement {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OwnerElement"}).(*QDomElement)
}

func (ptr *QDomAttr) SetValue(v string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetValue", v})
}

func (ptr *QDomAttr) Specified() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Specified"}).(bool)
}

func (ptr *QDomAttr) Value() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Value"}).(string)
}

type QDomCDATASection struct {
	QDomText
}

type QDomCDATASection_ITF interface {
	QDomText_ITF
	QDomCDATASection_PTR() *QDomCDATASection
}

func (ptr *QDomCDATASection) QDomCDATASection_PTR() *QDomCDATASection {
	return ptr
}

func (ptr *QDomCDATASection) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomText_PTR().Pointer()
	}
	return nil
}

func (ptr *QDomCDATASection) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDomText_PTR().SetPointer(p)
	}
}

func PointerFromQDomCDATASection(ptr QDomCDATASection_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomCDATASection_PTR().Pointer()
	}
	return nil
}

func (n *QDomCDATASection) InitFromInternal(ptr uintptr, name string) {
	n.QDomText_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDomCDATASection) ClassNameInternalF() string {
	return n.QDomText_PTR().ClassNameInternalF()
}

func NewQDomCDATASectionFromPointer(ptr unsafe.Pointer) (n *QDomCDATASection) {
	n = new(QDomCDATASection)
	n.InitFromInternal(uintptr(ptr), "xml.QDomCDATASection")
	return
}

func (ptr *QDomCDATASection) DestroyQDomCDATASection() {
}

func NewQDomCDATASection() *QDomCDATASection {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomCDATASection", ""}).(*QDomCDATASection)
}

func NewQDomCDATASection2(x QDomCDATASection_ITF) *QDomCDATASection {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomCDATASection2", "", x}).(*QDomCDATASection)
}

type QDomCharacterData struct {
	QDomNode
}

type QDomCharacterData_ITF interface {
	QDomNode_ITF
	QDomCharacterData_PTR() *QDomCharacterData
}

func (ptr *QDomCharacterData) QDomCharacterData_PTR() *QDomCharacterData {
	return ptr
}

func (ptr *QDomCharacterData) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QDomCharacterData) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDomNode_PTR().SetPointer(p)
	}
}

func PointerFromQDomCharacterData(ptr QDomCharacterData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomCharacterData_PTR().Pointer()
	}
	return nil
}

func (n *QDomCharacterData) InitFromInternal(ptr uintptr, name string) {
	n.QDomNode_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDomCharacterData) ClassNameInternalF() string {
	return n.QDomNode_PTR().ClassNameInternalF()
}

func NewQDomCharacterDataFromPointer(ptr unsafe.Pointer) (n *QDomCharacterData) {
	n = new(QDomCharacterData)
	n.InitFromInternal(uintptr(ptr), "xml.QDomCharacterData")
	return
}

func (ptr *QDomCharacterData) DestroyQDomCharacterData() {
}

func NewQDomCharacterData() *QDomCharacterData {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomCharacterData", ""}).(*QDomCharacterData)
}

func NewQDomCharacterData2(x QDomCharacterData_ITF) *QDomCharacterData {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomCharacterData2", "", x}).(*QDomCharacterData)
}

func (ptr *QDomCharacterData) AppendData(arg string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AppendData", arg})
}

func (ptr *QDomCharacterData) Data() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Data"}).(string)
}

func (ptr *QDomCharacterData) DeleteData(offset uint, count uint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteData", offset, count})
}

func (ptr *QDomCharacterData) InsertData(offset uint, arg string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InsertData", offset, arg})
}

func (ptr *QDomCharacterData) Length() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Length"}).(float64))
}

func (ptr *QDomCharacterData) ReplaceData(offset uint, count uint, arg string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReplaceData", offset, count, arg})
}

func (ptr *QDomCharacterData) SetData(v string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetData", v})
}

func (ptr *QDomCharacterData) SubstringData(offset uint, count uint) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SubstringData", offset, count}).(string)
}

type QDomComment struct {
	QDomCharacterData
}

type QDomComment_ITF interface {
	QDomCharacterData_ITF
	QDomComment_PTR() *QDomComment
}

func (ptr *QDomComment) QDomComment_PTR() *QDomComment {
	return ptr
}

func (ptr *QDomComment) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomCharacterData_PTR().Pointer()
	}
	return nil
}

func (ptr *QDomComment) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDomCharacterData_PTR().SetPointer(p)
	}
}

func PointerFromQDomComment(ptr QDomComment_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomComment_PTR().Pointer()
	}
	return nil
}

func (n *QDomComment) InitFromInternal(ptr uintptr, name string) {
	n.QDomCharacterData_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDomComment) ClassNameInternalF() string {
	return n.QDomCharacterData_PTR().ClassNameInternalF()
}

func NewQDomCommentFromPointer(ptr unsafe.Pointer) (n *QDomComment) {
	n = new(QDomComment)
	n.InitFromInternal(uintptr(ptr), "xml.QDomComment")
	return
}

func (ptr *QDomComment) DestroyQDomComment() {
}

func NewQDomComment() *QDomComment {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomComment", ""}).(*QDomComment)
}

func NewQDomComment2(x QDomComment_ITF) *QDomComment {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomComment2", "", x}).(*QDomComment)
}

type QDomDocument struct {
	QDomNode
}

type QDomDocument_ITF interface {
	QDomNode_ITF
	QDomDocument_PTR() *QDomDocument
}

func (ptr *QDomDocument) QDomDocument_PTR() *QDomDocument {
	return ptr
}

func (ptr *QDomDocument) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QDomDocument) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDomNode_PTR().SetPointer(p)
	}
}

func PointerFromQDomDocument(ptr QDomDocument_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomDocument_PTR().Pointer()
	}
	return nil
}

func (n *QDomDocument) InitFromInternal(ptr uintptr, name string) {
	n.QDomNode_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDomDocument) ClassNameInternalF() string {
	return n.QDomNode_PTR().ClassNameInternalF()
}

func NewQDomDocumentFromPointer(ptr unsafe.Pointer) (n *QDomDocument) {
	n = new(QDomDocument)
	n.InitFromInternal(uintptr(ptr), "xml.QDomDocument")
	return
}
func NewQDomDocument() *QDomDocument {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomDocument", ""}).(*QDomDocument)
}

func NewQDomDocument2(name string) *QDomDocument {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomDocument2", "", name}).(*QDomDocument)
}

func NewQDomDocument3(doctype QDomDocumentType_ITF) *QDomDocument {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomDocument3", "", doctype}).(*QDomDocument)
}

func NewQDomDocument4(x QDomDocument_ITF) *QDomDocument {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomDocument4", "", x}).(*QDomDocument)
}

func (ptr *QDomDocument) CreateAttribute(name string) *QDomAttr {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateAttribute", name}).(*QDomAttr)
}

func (ptr *QDomDocument) CreateAttributeNS(nsURI string, qName string) *QDomAttr {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateAttributeNS", nsURI, qName}).(*QDomAttr)
}

func (ptr *QDomDocument) CreateCDATASection(value string) *QDomCDATASection {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateCDATASection", value}).(*QDomCDATASection)
}

func (ptr *QDomDocument) CreateComment(value string) *QDomComment {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateComment", value}).(*QDomComment)
}

func (ptr *QDomDocument) CreateDocumentFragment() *QDomDocumentFragment {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateDocumentFragment"}).(*QDomDocumentFragment)
}

func (ptr *QDomDocument) CreateElement(tagName string) *QDomElement {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateElement", tagName}).(*QDomElement)
}

func (ptr *QDomDocument) CreateElementNS(nsURI string, qName string) *QDomElement {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateElementNS", nsURI, qName}).(*QDomElement)
}

func (ptr *QDomDocument) CreateEntityReference(name string) *QDomEntityReference {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateEntityReference", name}).(*QDomEntityReference)
}

func (ptr *QDomDocument) CreateProcessingInstruction(target string, data string) *QDomProcessingInstruction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateProcessingInstruction", target, data}).(*QDomProcessingInstruction)
}

func (ptr *QDomDocument) CreateTextNode(value string) *QDomText {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateTextNode", value}).(*QDomText)
}

func (ptr *QDomDocument) Doctype() *QDomDocumentType {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Doctype"}).(*QDomDocumentType)
}

func (ptr *QDomDocument) DocumentElement() *QDomElement {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DocumentElement"}).(*QDomElement)
}

func (ptr *QDomDocument) ElementById(elementId string) *QDomElement {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ElementById", elementId}).(*QDomElement)
}

func (ptr *QDomDocument) ElementsByTagName(tagname string) *QDomNodeList {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ElementsByTagName", tagname}).(*QDomNodeList)
}

func (ptr *QDomDocument) ElementsByTagNameNS(nsURI string, localName string) *QDomNodeList {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ElementsByTagNameNS", nsURI, localName}).(*QDomNodeList)
}

func (ptr *QDomDocument) Implementation() *QDomImplementation {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Implementation"}).(*QDomImplementation)
}

func (ptr *QDomDocument) ImportNode(importedNode QDomNode_ITF, deep bool) *QDomNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ImportNode", importedNode, deep}).(*QDomNode)
}

func (ptr *QDomDocument) SetContent(data core.QByteArray_ITF, namespaceProcessing bool, errorMsg string, errorLine int, errorColumn int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetContent", data, namespaceProcessing, errorMsg, errorLine, errorColumn}).(bool)
}

func (ptr *QDomDocument) SetContent2(text string, namespaceProcessing bool, errorMsg string, errorLine int, errorColumn int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetContent2", text, namespaceProcessing, errorMsg, errorLine, errorColumn}).(bool)
}

func (ptr *QDomDocument) SetContent3(dev core.QIODevice_ITF, namespaceProcessing bool, errorMsg string, errorLine int, errorColumn int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetContent3", dev, namespaceProcessing, errorMsg, errorLine, errorColumn}).(bool)
}

func (ptr *QDomDocument) SetContent4(source QXmlInputSource_ITF, namespaceProcessing bool, errorMsg string, errorLine int, errorColumn int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetContent4", source, namespaceProcessing, errorMsg, errorLine, errorColumn}).(bool)
}

func (ptr *QDomDocument) SetContent5(buffer core.QByteArray_ITF, errorMsg string, errorLine int, errorColumn int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetContent5", buffer, errorMsg, errorLine, errorColumn}).(bool)
}

func (ptr *QDomDocument) SetContent6(text string, errorMsg string, errorLine int, errorColumn int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetContent6", text, errorMsg, errorLine, errorColumn}).(bool)
}

func (ptr *QDomDocument) SetContent7(dev core.QIODevice_ITF, errorMsg string, errorLine int, errorColumn int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetContent7", dev, errorMsg, errorLine, errorColumn}).(bool)
}

func (ptr *QDomDocument) SetContent8(source QXmlInputSource_ITF, reader QXmlReader_ITF, errorMsg string, errorLine int, errorColumn int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetContent8", source, reader, errorMsg, errorLine, errorColumn}).(bool)
}

func (ptr *QDomDocument) ToByteArray(indent int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToByteArray", indent}).(*core.QByteArray)
}

func (ptr *QDomDocument) ToString(indent int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToString", indent}).(string)
}

func (ptr *QDomDocument) DestroyQDomDocument() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDomDocument"})
}

type QDomDocumentFragment struct {
	QDomNode
}

type QDomDocumentFragment_ITF interface {
	QDomNode_ITF
	QDomDocumentFragment_PTR() *QDomDocumentFragment
}

func (ptr *QDomDocumentFragment) QDomDocumentFragment_PTR() *QDomDocumentFragment {
	return ptr
}

func (ptr *QDomDocumentFragment) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QDomDocumentFragment) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDomNode_PTR().SetPointer(p)
	}
}

func PointerFromQDomDocumentFragment(ptr QDomDocumentFragment_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomDocumentFragment_PTR().Pointer()
	}
	return nil
}

func (n *QDomDocumentFragment) InitFromInternal(ptr uintptr, name string) {
	n.QDomNode_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDomDocumentFragment) ClassNameInternalF() string {
	return n.QDomNode_PTR().ClassNameInternalF()
}

func NewQDomDocumentFragmentFromPointer(ptr unsafe.Pointer) (n *QDomDocumentFragment) {
	n = new(QDomDocumentFragment)
	n.InitFromInternal(uintptr(ptr), "xml.QDomDocumentFragment")
	return
}

func (ptr *QDomDocumentFragment) DestroyQDomDocumentFragment() {
}

func NewQDomDocumentFragment() *QDomDocumentFragment {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomDocumentFragment", ""}).(*QDomDocumentFragment)
}

func NewQDomDocumentFragment2(x QDomDocumentFragment_ITF) *QDomDocumentFragment {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomDocumentFragment2", "", x}).(*QDomDocumentFragment)
}

type QDomDocumentType struct {
	QDomNode
}

type QDomDocumentType_ITF interface {
	QDomNode_ITF
	QDomDocumentType_PTR() *QDomDocumentType
}

func (ptr *QDomDocumentType) QDomDocumentType_PTR() *QDomDocumentType {
	return ptr
}

func (ptr *QDomDocumentType) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QDomDocumentType) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDomNode_PTR().SetPointer(p)
	}
}

func PointerFromQDomDocumentType(ptr QDomDocumentType_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomDocumentType_PTR().Pointer()
	}
	return nil
}

func (n *QDomDocumentType) InitFromInternal(ptr uintptr, name string) {
	n.QDomNode_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDomDocumentType) ClassNameInternalF() string {
	return n.QDomNode_PTR().ClassNameInternalF()
}

func NewQDomDocumentTypeFromPointer(ptr unsafe.Pointer) (n *QDomDocumentType) {
	n = new(QDomDocumentType)
	n.InitFromInternal(uintptr(ptr), "xml.QDomDocumentType")
	return
}

func (ptr *QDomDocumentType) DestroyQDomDocumentType() {
}

func NewQDomDocumentType() *QDomDocumentType {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomDocumentType", ""}).(*QDomDocumentType)
}

func NewQDomDocumentType2(n QDomDocumentType_ITF) *QDomDocumentType {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomDocumentType2", "", n}).(*QDomDocumentType)
}

func (ptr *QDomDocumentType) Entities() *QDomNamedNodeMap {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Entities"}).(*QDomNamedNodeMap)
}

func (ptr *QDomDocumentType) InternalSubset() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InternalSubset"}).(string)
}

func (ptr *QDomDocumentType) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QDomDocumentType) Notations() *QDomNamedNodeMap {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Notations"}).(*QDomNamedNodeMap)
}

func (ptr *QDomDocumentType) PublicId() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PublicId"}).(string)
}

func (ptr *QDomDocumentType) SystemId() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SystemId"}).(string)
}

type QDomElement struct {
	QDomNode
}

type QDomElement_ITF interface {
	QDomNode_ITF
	QDomElement_PTR() *QDomElement
}

func (ptr *QDomElement) QDomElement_PTR() *QDomElement {
	return ptr
}

func (ptr *QDomElement) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QDomElement) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDomNode_PTR().SetPointer(p)
	}
}

func PointerFromQDomElement(ptr QDomElement_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomElement_PTR().Pointer()
	}
	return nil
}

func (n *QDomElement) InitFromInternal(ptr uintptr, name string) {
	n.QDomNode_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDomElement) ClassNameInternalF() string {
	return n.QDomNode_PTR().ClassNameInternalF()
}

func NewQDomElementFromPointer(ptr unsafe.Pointer) (n *QDomElement) {
	n = new(QDomElement)
	n.InitFromInternal(uintptr(ptr), "xml.QDomElement")
	return
}

func (ptr *QDomElement) DestroyQDomElement() {
}

func NewQDomElement() *QDomElement {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomElement", ""}).(*QDomElement)
}

func NewQDomElement2(x QDomElement_ITF) *QDomElement {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomElement2", "", x}).(*QDomElement)
}

func (ptr *QDomElement) Attribute(name string, defValue string) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Attribute", name, defValue}).(string)
}

func (ptr *QDomElement) AttributeNS(nsURI string, localName string, defValue string) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AttributeNS", nsURI, localName, defValue}).(string)
}

func (ptr *QDomElement) AttributeNode(name string) *QDomAttr {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AttributeNode", name}).(*QDomAttr)
}

func (ptr *QDomElement) AttributeNodeNS(nsURI string, localName string) *QDomAttr {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AttributeNodeNS", nsURI, localName}).(*QDomAttr)
}

func (ptr *QDomElement) ElementsByTagName(tagname string) *QDomNodeList {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ElementsByTagName", tagname}).(*QDomNodeList)
}

func (ptr *QDomElement) ElementsByTagNameNS(nsURI string, localName string) *QDomNodeList {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ElementsByTagNameNS", nsURI, localName}).(*QDomNodeList)
}

func (ptr *QDomElement) HasAttribute(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasAttribute", name}).(bool)
}

func (ptr *QDomElement) HasAttributeNS(nsURI string, localName string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasAttributeNS", nsURI, localName}).(bool)
}

func (ptr *QDomElement) RemoveAttribute(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveAttribute", name})
}

func (ptr *QDomElement) RemoveAttributeNS(nsURI string, localName string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveAttributeNS", nsURI, localName})
}

func (ptr *QDomElement) RemoveAttributeNode(oldAttr QDomAttr_ITF) *QDomAttr {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveAttributeNode", oldAttr}).(*QDomAttr)
}

func (ptr *QDomElement) SetAttribute(name string, value string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAttribute", name, value})
}

func (ptr *QDomElement) SetAttribute2(name string, value int64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAttribute2", name, value})
}

func (ptr *QDomElement) SetAttribute3(name string, value uint64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAttribute3", name, value})
}

func (ptr *QDomElement) SetAttribute4(name string, value int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAttribute4", name, value})
}

func (ptr *QDomElement) SetAttribute5(name string, value uint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAttribute5", name, value})
}

func (ptr *QDomElement) SetAttribute6(name string, value float32) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAttribute6", name, value})
}

func (ptr *QDomElement) SetAttribute7(name string, value float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAttribute7", name, value})
}

func (ptr *QDomElement) SetAttributeNS(nsURI string, qName string, value string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAttributeNS", nsURI, qName, value})
}

func (ptr *QDomElement) SetAttributeNS2(nsURI string, qName string, value int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAttributeNS2", nsURI, qName, value})
}

func (ptr *QDomElement) SetAttributeNS3(nsURI string, qName string, value uint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAttributeNS3", nsURI, qName, value})
}

func (ptr *QDomElement) SetAttributeNS4(nsURI string, qName string, value int64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAttributeNS4", nsURI, qName, value})
}

func (ptr *QDomElement) SetAttributeNS5(nsURI string, qName string, value uint64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAttributeNS5", nsURI, qName, value})
}

func (ptr *QDomElement) SetAttributeNS6(nsURI string, qName string, value float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAttributeNS6", nsURI, qName, value})
}

func (ptr *QDomElement) SetAttributeNode(newAttr QDomAttr_ITF) *QDomAttr {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAttributeNode", newAttr}).(*QDomAttr)
}

func (ptr *QDomElement) SetAttributeNodeNS(newAttr QDomAttr_ITF) *QDomAttr {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAttributeNodeNS", newAttr}).(*QDomAttr)
}

func (ptr *QDomElement) SetTagName(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTagName", name})
}

func (ptr *QDomElement) TagName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TagName"}).(string)
}

func (ptr *QDomElement) Text() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Text"}).(string)
}

type QDomEntity struct {
	QDomNode
}

type QDomEntity_ITF interface {
	QDomNode_ITF
	QDomEntity_PTR() *QDomEntity
}

func (ptr *QDomEntity) QDomEntity_PTR() *QDomEntity {
	return ptr
}

func (ptr *QDomEntity) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QDomEntity) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDomNode_PTR().SetPointer(p)
	}
}

func PointerFromQDomEntity(ptr QDomEntity_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomEntity_PTR().Pointer()
	}
	return nil
}

func (n *QDomEntity) InitFromInternal(ptr uintptr, name string) {
	n.QDomNode_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDomEntity) ClassNameInternalF() string {
	return n.QDomNode_PTR().ClassNameInternalF()
}

func NewQDomEntityFromPointer(ptr unsafe.Pointer) (n *QDomEntity) {
	n = new(QDomEntity)
	n.InitFromInternal(uintptr(ptr), "xml.QDomEntity")
	return
}

func (ptr *QDomEntity) DestroyQDomEntity() {
}

func NewQDomEntity() *QDomEntity {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomEntity", ""}).(*QDomEntity)
}

func NewQDomEntity2(x QDomEntity_ITF) *QDomEntity {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomEntity2", "", x}).(*QDomEntity)
}

func (ptr *QDomEntity) NotationName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NotationName"}).(string)
}

func (ptr *QDomEntity) PublicId() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PublicId"}).(string)
}

func (ptr *QDomEntity) SystemId() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SystemId"}).(string)
}

type QDomEntityReference struct {
	QDomNode
}

type QDomEntityReference_ITF interface {
	QDomNode_ITF
	QDomEntityReference_PTR() *QDomEntityReference
}

func (ptr *QDomEntityReference) QDomEntityReference_PTR() *QDomEntityReference {
	return ptr
}

func (ptr *QDomEntityReference) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QDomEntityReference) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDomNode_PTR().SetPointer(p)
	}
}

func PointerFromQDomEntityReference(ptr QDomEntityReference_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomEntityReference_PTR().Pointer()
	}
	return nil
}

func (n *QDomEntityReference) InitFromInternal(ptr uintptr, name string) {
	n.QDomNode_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDomEntityReference) ClassNameInternalF() string {
	return n.QDomNode_PTR().ClassNameInternalF()
}

func NewQDomEntityReferenceFromPointer(ptr unsafe.Pointer) (n *QDomEntityReference) {
	n = new(QDomEntityReference)
	n.InitFromInternal(uintptr(ptr), "xml.QDomEntityReference")
	return
}

func (ptr *QDomEntityReference) DestroyQDomEntityReference() {
}

func NewQDomEntityReference() *QDomEntityReference {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomEntityReference", ""}).(*QDomEntityReference)
}

func NewQDomEntityReference2(x QDomEntityReference_ITF) *QDomEntityReference {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomEntityReference2", "", x}).(*QDomEntityReference)
}

type QDomImplementation struct {
	internal.Internal
}

type QDomImplementation_ITF interface {
	QDomImplementation_PTR() *QDomImplementation
}

func (ptr *QDomImplementation) QDomImplementation_PTR() *QDomImplementation {
	return ptr
}

func (ptr *QDomImplementation) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDomImplementation) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDomImplementation(ptr QDomImplementation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomImplementation_PTR().Pointer()
	}
	return nil
}

func (n *QDomImplementation) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDomImplementationFromPointer(ptr unsafe.Pointer) (n *QDomImplementation) {
	n = new(QDomImplementation)
	n.InitFromInternal(uintptr(ptr), "xml.QDomImplementation")
	return
}

//go:generate stringer -type=QDomImplementation__InvalidDataPolicy
//QDomImplementation::InvalidDataPolicy
type QDomImplementation__InvalidDataPolicy int64

const (
	QDomImplementation__AcceptInvalidChars QDomImplementation__InvalidDataPolicy = QDomImplementation__InvalidDataPolicy(0)
	QDomImplementation__DropInvalidChars   QDomImplementation__InvalidDataPolicy = QDomImplementation__InvalidDataPolicy(1)
	QDomImplementation__ReturnNullNode     QDomImplementation__InvalidDataPolicy = QDomImplementation__InvalidDataPolicy(2)
)

func NewQDomImplementation() *QDomImplementation {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomImplementation", ""}).(*QDomImplementation)
}

func NewQDomImplementation2(x QDomImplementation_ITF) *QDomImplementation {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomImplementation2", "", x}).(*QDomImplementation)
}

func (ptr *QDomImplementation) CreateDocument(nsURI string, qName string, doctype QDomDocumentType_ITF) *QDomDocument {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateDocument", nsURI, qName, doctype}).(*QDomDocument)
}

func (ptr *QDomImplementation) CreateDocumentType(qName string, publicId string, systemId string) *QDomDocumentType {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateDocumentType", qName, publicId, systemId}).(*QDomDocumentType)
}

func (ptr *QDomImplementation) HasFeature(feature string, version string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasFeature", feature, version}).(bool)
}

func QDomImplementation_InvalidDataPolicy() QDomImplementation__InvalidDataPolicy {

	return QDomImplementation__InvalidDataPolicy(internal.CallLocalFunction([]interface{}{"", "", "xml.QDomImplementation_InvalidDataPolicy", ""}).(float64))
}

func (ptr *QDomImplementation) InvalidDataPolicy() QDomImplementation__InvalidDataPolicy {

	return QDomImplementation__InvalidDataPolicy(internal.CallLocalFunction([]interface{}{"", "", "xml.QDomImplementation_InvalidDataPolicy", ""}).(float64))
}

func (ptr *QDomImplementation) IsNull() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull"}).(bool)
}

func QDomImplementation_SetInvalidDataPolicy(policy QDomImplementation__InvalidDataPolicy) {

	internal.CallLocalFunction([]interface{}{"", "", "xml.QDomImplementation_SetInvalidDataPolicy", "", policy})
}

func (ptr *QDomImplementation) SetInvalidDataPolicy(policy QDomImplementation__InvalidDataPolicy) {

	internal.CallLocalFunction([]interface{}{"", "", "xml.QDomImplementation_SetInvalidDataPolicy", "", policy})
}

func (ptr *QDomImplementation) DestroyQDomImplementation() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDomImplementation"})
}

type QDomNamedNodeMap struct {
	internal.Internal
}

type QDomNamedNodeMap_ITF interface {
	QDomNamedNodeMap_PTR() *QDomNamedNodeMap
}

func (ptr *QDomNamedNodeMap) QDomNamedNodeMap_PTR() *QDomNamedNodeMap {
	return ptr
}

func (ptr *QDomNamedNodeMap) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDomNamedNodeMap) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDomNamedNodeMap(ptr QDomNamedNodeMap_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNamedNodeMap_PTR().Pointer()
	}
	return nil
}

func (n *QDomNamedNodeMap) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDomNamedNodeMapFromPointer(ptr unsafe.Pointer) (n *QDomNamedNodeMap) {
	n = new(QDomNamedNodeMap)
	n.InitFromInternal(uintptr(ptr), "xml.QDomNamedNodeMap")
	return
}
func NewQDomNamedNodeMap() *QDomNamedNodeMap {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomNamedNodeMap", ""}).(*QDomNamedNodeMap)
}

func NewQDomNamedNodeMap2(n QDomNamedNodeMap_ITF) *QDomNamedNodeMap {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomNamedNodeMap2", "", n}).(*QDomNamedNodeMap)
}

func (ptr *QDomNamedNodeMap) Contains(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Contains", name}).(bool)
}

func (ptr *QDomNamedNodeMap) Count() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Count"}).(float64))
}

func (ptr *QDomNamedNodeMap) IsEmpty() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsEmpty"}).(bool)
}

func (ptr *QDomNamedNodeMap) Item(index int) *QDomNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Item", index}).(*QDomNode)
}

func (ptr *QDomNamedNodeMap) Length() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Length"}).(float64))
}

func (ptr *QDomNamedNodeMap) NamedItem(name string) *QDomNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NamedItem", name}).(*QDomNode)
}

func (ptr *QDomNamedNodeMap) NamedItemNS(nsURI string, localName string) *QDomNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NamedItemNS", nsURI, localName}).(*QDomNode)
}

func (ptr *QDomNamedNodeMap) RemoveNamedItem(name string) *QDomNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveNamedItem", name}).(*QDomNode)
}

func (ptr *QDomNamedNodeMap) RemoveNamedItemNS(nsURI string, localName string) *QDomNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveNamedItemNS", nsURI, localName}).(*QDomNode)
}

func (ptr *QDomNamedNodeMap) SetNamedItem(newNode QDomNode_ITF) *QDomNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetNamedItem", newNode}).(*QDomNode)
}

func (ptr *QDomNamedNodeMap) SetNamedItemNS(newNode QDomNode_ITF) *QDomNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetNamedItemNS", newNode}).(*QDomNode)
}

func (ptr *QDomNamedNodeMap) Size() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Size"}).(float64))
}

func (ptr *QDomNamedNodeMap) DestroyQDomNamedNodeMap() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDomNamedNodeMap"})
}

type QDomNode struct {
	internal.Internal
}

type QDomNode_ITF interface {
	QDomNode_PTR() *QDomNode
}

func (ptr *QDomNode) QDomNode_PTR() *QDomNode {
	return ptr
}

func (ptr *QDomNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDomNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDomNode(ptr QDomNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNode_PTR().Pointer()
	}
	return nil
}

func (n *QDomNode) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDomNodeFromPointer(ptr unsafe.Pointer) (n *QDomNode) {
	n = new(QDomNode)
	n.InitFromInternal(uintptr(ptr), "xml.QDomNode")
	return
}

//go:generate stringer -type=QDomNode__NodeType
//QDomNode::NodeType
type QDomNode__NodeType int64

const (
	QDomNode__ElementNode               QDomNode__NodeType = QDomNode__NodeType(1)
	QDomNode__AttributeNode             QDomNode__NodeType = QDomNode__NodeType(2)
	QDomNode__TextNode                  QDomNode__NodeType = QDomNode__NodeType(3)
	QDomNode__CDATASectionNode          QDomNode__NodeType = QDomNode__NodeType(4)
	QDomNode__EntityReferenceNode       QDomNode__NodeType = QDomNode__NodeType(5)
	QDomNode__EntityNode                QDomNode__NodeType = QDomNode__NodeType(6)
	QDomNode__ProcessingInstructionNode QDomNode__NodeType = QDomNode__NodeType(7)
	QDomNode__CommentNode               QDomNode__NodeType = QDomNode__NodeType(8)
	QDomNode__DocumentNode              QDomNode__NodeType = QDomNode__NodeType(9)
	QDomNode__DocumentTypeNode          QDomNode__NodeType = QDomNode__NodeType(10)
	QDomNode__DocumentFragmentNode      QDomNode__NodeType = QDomNode__NodeType(11)
	QDomNode__NotationNode              QDomNode__NodeType = QDomNode__NodeType(12)
	QDomNode__BaseNode                  QDomNode__NodeType = QDomNode__NodeType(21)
	QDomNode__CharacterDataNode         QDomNode__NodeType = QDomNode__NodeType(22)
)

//go:generate stringer -type=QDomNode__EncodingPolicy
//QDomNode::EncodingPolicy
type QDomNode__EncodingPolicy int64

const (
	QDomNode__EncodingFromDocument   QDomNode__EncodingPolicy = QDomNode__EncodingPolicy(1)
	QDomNode__EncodingFromTextStream QDomNode__EncodingPolicy = QDomNode__EncodingPolicy(2)
)

func NewQDomNode() *QDomNode {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomNode", ""}).(*QDomNode)
}

func NewQDomNode2(n QDomNode_ITF) *QDomNode {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomNode2", "", n}).(*QDomNode)
}

func (ptr *QDomNode) AppendChild(newChild QDomNode_ITF) *QDomNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AppendChild", newChild}).(*QDomNode)
}

func (ptr *QDomNode) ChildNodes() *QDomNodeList {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildNodes"}).(*QDomNodeList)
}

func (ptr *QDomNode) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QDomNode) CloneNode(deep bool) *QDomNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloneNode", deep}).(*QDomNode)
}

func (ptr *QDomNode) ColumnNumber() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnNumber"}).(float64))
}

func (ptr *QDomNode) FirstChild() *QDomNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstChild"}).(*QDomNode)
}

func (ptr *QDomNode) FirstChildElement(tagName string) *QDomElement {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstChildElement", tagName}).(*QDomElement)
}

func (ptr *QDomNode) HasAttributes() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasAttributes"}).(bool)
}

func (ptr *QDomNode) HasChildNodes() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasChildNodes"}).(bool)
}

func (ptr *QDomNode) InsertAfter(newChild QDomNode_ITF, refChild QDomNode_ITF) *QDomNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InsertAfter", newChild, refChild}).(*QDomNode)
}

func (ptr *QDomNode) InsertBefore(newChild QDomNode_ITF, refChild QDomNode_ITF) *QDomNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InsertBefore", newChild, refChild}).(*QDomNode)
}

func (ptr *QDomNode) IsAttr() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsAttr"}).(bool)
}

func (ptr *QDomNode) IsCDATASection() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsCDATASection"}).(bool)
}

func (ptr *QDomNode) IsCharacterData() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsCharacterData"}).(bool)
}

func (ptr *QDomNode) IsComment() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsComment"}).(bool)
}

func (ptr *QDomNode) IsDocument() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsDocument"}).(bool)
}

func (ptr *QDomNode) IsDocumentFragment() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsDocumentFragment"}).(bool)
}

func (ptr *QDomNode) IsDocumentType() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsDocumentType"}).(bool)
}

func (ptr *QDomNode) IsElement() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsElement"}).(bool)
}

func (ptr *QDomNode) IsEntity() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsEntity"}).(bool)
}

func (ptr *QDomNode) IsEntityReference() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsEntityReference"}).(bool)
}

func (ptr *QDomNode) IsNotation() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNotation"}).(bool)
}

func (ptr *QDomNode) IsNull() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull"}).(bool)
}

func (ptr *QDomNode) IsProcessingInstruction() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsProcessingInstruction"}).(bool)
}

func (ptr *QDomNode) IsSupported(feature string, version string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsSupported", feature, version}).(bool)
}

func (ptr *QDomNode) IsText() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsText"}).(bool)
}

func (ptr *QDomNode) LastChild() *QDomNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastChild"}).(*QDomNode)
}

func (ptr *QDomNode) LastChildElement(tagName string) *QDomElement {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastChildElement", tagName}).(*QDomElement)
}

func (ptr *QDomNode) LineNumber() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LineNumber"}).(float64))
}

func (ptr *QDomNode) LocalName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LocalName"}).(string)
}

func (ptr *QDomNode) NamedItem(name string) *QDomNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NamedItem", name}).(*QDomNode)
}

func (ptr *QDomNode) NamespaceURI() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NamespaceURI"}).(string)
}

func (ptr *QDomNode) NextSibling() *QDomNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NextSibling"}).(*QDomNode)
}

func (ptr *QDomNode) NextSiblingElement(tagName string) *QDomElement {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NextSiblingElement", tagName}).(*QDomElement)
}

func (ptr *QDomNode) NodeName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NodeName"}).(string)
}

func (ptr *QDomNode) NodeType() QDomNode__NodeType {

	return QDomNode__NodeType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NodeType"}).(float64))
}

func (ptr *QDomNode) NodeValue() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NodeValue"}).(string)
}

func (ptr *QDomNode) Normalize() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Normalize"})
}

func (ptr *QDomNode) OwnerDocument() *QDomDocument {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OwnerDocument"}).(*QDomDocument)
}

func (ptr *QDomNode) ParentNode() *QDomNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ParentNode"}).(*QDomNode)
}

func (ptr *QDomNode) Prefix() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Prefix"}).(string)
}

func (ptr *QDomNode) PreviousSibling() *QDomNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PreviousSibling"}).(*QDomNode)
}

func (ptr *QDomNode) PreviousSiblingElement(tagName string) *QDomElement {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PreviousSiblingElement", tagName}).(*QDomElement)
}

func (ptr *QDomNode) RemoveChild(oldChild QDomNode_ITF) *QDomNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveChild", oldChild}).(*QDomNode)
}

func (ptr *QDomNode) ReplaceChild(newChild QDomNode_ITF, oldChild QDomNode_ITF) *QDomNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReplaceChild", newChild, oldChild}).(*QDomNode)
}

func (ptr *QDomNode) Save(stream core.QTextStream_ITF, indent int, encodingPolicy QDomNode__EncodingPolicy) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Save", stream, indent, encodingPolicy})
}

func (ptr *QDomNode) SetNodeValue(v string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetNodeValue", v})
}

func (ptr *QDomNode) SetPrefix(pre string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPrefix", pre})
}

func (ptr *QDomNode) ToAttr() *QDomAttr {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToAttr"}).(*QDomAttr)
}

func (ptr *QDomNode) ToCDATASection() *QDomCDATASection {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToCDATASection"}).(*QDomCDATASection)
}

func (ptr *QDomNode) ToCharacterData() *QDomCharacterData {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToCharacterData"}).(*QDomCharacterData)
}

func (ptr *QDomNode) ToComment() *QDomComment {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToComment"}).(*QDomComment)
}

func (ptr *QDomNode) ToDocument() *QDomDocument {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToDocument"}).(*QDomDocument)
}

func (ptr *QDomNode) ToDocumentFragment() *QDomDocumentFragment {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToDocumentFragment"}).(*QDomDocumentFragment)
}

func (ptr *QDomNode) ToDocumentType() *QDomDocumentType {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToDocumentType"}).(*QDomDocumentType)
}

func (ptr *QDomNode) ToElement() *QDomElement {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToElement"}).(*QDomElement)
}

func (ptr *QDomNode) ToEntity() *QDomEntity {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToEntity"}).(*QDomEntity)
}

func (ptr *QDomNode) ToEntityReference() *QDomEntityReference {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToEntityReference"}).(*QDomEntityReference)
}

func (ptr *QDomNode) ToNotation() *QDomNotation {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToNotation"}).(*QDomNotation)
}

func (ptr *QDomNode) ToProcessingInstruction() *QDomProcessingInstruction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToProcessingInstruction"}).(*QDomProcessingInstruction)
}

func (ptr *QDomNode) ToText() *QDomText {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToText"}).(*QDomText)
}

func (ptr *QDomNode) DestroyQDomNode() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDomNode"})
}

type QDomNodeList struct {
	internal.Internal
}

type QDomNodeList_ITF interface {
	QDomNodeList_PTR() *QDomNodeList
}

func (ptr *QDomNodeList) QDomNodeList_PTR() *QDomNodeList {
	return ptr
}

func (ptr *QDomNodeList) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDomNodeList) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDomNodeList(ptr QDomNodeList_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNodeList_PTR().Pointer()
	}
	return nil
}

func (n *QDomNodeList) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDomNodeListFromPointer(ptr unsafe.Pointer) (n *QDomNodeList) {
	n = new(QDomNodeList)
	n.InitFromInternal(uintptr(ptr), "xml.QDomNodeList")
	return
}
func NewQDomNodeList() *QDomNodeList {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomNodeList", ""}).(*QDomNodeList)
}

func NewQDomNodeList2(n QDomNodeList_ITF) *QDomNodeList {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomNodeList2", "", n}).(*QDomNodeList)
}

func (ptr *QDomNodeList) At(index int) *QDomNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "At", index}).(*QDomNode)
}

func (ptr *QDomNodeList) Count() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Count"}).(float64))
}

func (ptr *QDomNodeList) IsEmpty() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsEmpty"}).(bool)
}

func (ptr *QDomNodeList) Item(index int) *QDomNode {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Item", index}).(*QDomNode)
}

func (ptr *QDomNodeList) Length() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Length"}).(float64))
}

func (ptr *QDomNodeList) Size() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Size"}).(float64))
}

func (ptr *QDomNodeList) DestroyQDomNodeList() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDomNodeList"})
}

type QDomNotation struct {
	QDomNode
}

type QDomNotation_ITF interface {
	QDomNode_ITF
	QDomNotation_PTR() *QDomNotation
}

func (ptr *QDomNotation) QDomNotation_PTR() *QDomNotation {
	return ptr
}

func (ptr *QDomNotation) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QDomNotation) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDomNode_PTR().SetPointer(p)
	}
}

func PointerFromQDomNotation(ptr QDomNotation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNotation_PTR().Pointer()
	}
	return nil
}

func (n *QDomNotation) InitFromInternal(ptr uintptr, name string) {
	n.QDomNode_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDomNotation) ClassNameInternalF() string {
	return n.QDomNode_PTR().ClassNameInternalF()
}

func NewQDomNotationFromPointer(ptr unsafe.Pointer) (n *QDomNotation) {
	n = new(QDomNotation)
	n.InitFromInternal(uintptr(ptr), "xml.QDomNotation")
	return
}

func (ptr *QDomNotation) DestroyQDomNotation() {
}

func NewQDomNotation() *QDomNotation {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomNotation", ""}).(*QDomNotation)
}

func NewQDomNotation2(x QDomNotation_ITF) *QDomNotation {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomNotation2", "", x}).(*QDomNotation)
}

func (ptr *QDomNotation) PublicId() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PublicId"}).(string)
}

func (ptr *QDomNotation) SystemId() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SystemId"}).(string)
}

type QDomProcessingInstruction struct {
	QDomNode
}

type QDomProcessingInstruction_ITF interface {
	QDomNode_ITF
	QDomProcessingInstruction_PTR() *QDomProcessingInstruction
}

func (ptr *QDomProcessingInstruction) QDomProcessingInstruction_PTR() *QDomProcessingInstruction {
	return ptr
}

func (ptr *QDomProcessingInstruction) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QDomProcessingInstruction) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDomNode_PTR().SetPointer(p)
	}
}

func PointerFromQDomProcessingInstruction(ptr QDomProcessingInstruction_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomProcessingInstruction_PTR().Pointer()
	}
	return nil
}

func (n *QDomProcessingInstruction) InitFromInternal(ptr uintptr, name string) {
	n.QDomNode_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDomProcessingInstruction) ClassNameInternalF() string {
	return n.QDomNode_PTR().ClassNameInternalF()
}

func NewQDomProcessingInstructionFromPointer(ptr unsafe.Pointer) (n *QDomProcessingInstruction) {
	n = new(QDomProcessingInstruction)
	n.InitFromInternal(uintptr(ptr), "xml.QDomProcessingInstruction")
	return
}

func (ptr *QDomProcessingInstruction) DestroyQDomProcessingInstruction() {
}

func NewQDomProcessingInstruction() *QDomProcessingInstruction {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomProcessingInstruction", ""}).(*QDomProcessingInstruction)
}

func NewQDomProcessingInstruction2(x QDomProcessingInstruction_ITF) *QDomProcessingInstruction {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomProcessingInstruction2", "", x}).(*QDomProcessingInstruction)
}

func (ptr *QDomProcessingInstruction) Data() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Data"}).(string)
}

func (ptr *QDomProcessingInstruction) SetData(d string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetData", d})
}

func (ptr *QDomProcessingInstruction) Target() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Target"}).(string)
}

type QDomText struct {
	QDomCharacterData
}

type QDomText_ITF interface {
	QDomCharacterData_ITF
	QDomText_PTR() *QDomText
}

func (ptr *QDomText) QDomText_PTR() *QDomText {
	return ptr
}

func (ptr *QDomText) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomCharacterData_PTR().Pointer()
	}
	return nil
}

func (ptr *QDomText) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDomCharacterData_PTR().SetPointer(p)
	}
}

func PointerFromQDomText(ptr QDomText_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomText_PTR().Pointer()
	}
	return nil
}

func (n *QDomText) InitFromInternal(ptr uintptr, name string) {
	n.QDomCharacterData_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDomText) ClassNameInternalF() string {
	return n.QDomCharacterData_PTR().ClassNameInternalF()
}

func NewQDomTextFromPointer(ptr unsafe.Pointer) (n *QDomText) {
	n = new(QDomText)
	n.InitFromInternal(uintptr(ptr), "xml.QDomText")
	return
}

func (ptr *QDomText) DestroyQDomText() {
}

func NewQDomText() *QDomText {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomText", ""}).(*QDomText)
}

func NewQDomText2(x QDomText_ITF) *QDomText {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQDomText2", "", x}).(*QDomText)
}

func (ptr *QDomText) SplitText(offset int) *QDomText {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SplitText", offset}).(*QDomText)
}

type QXmlAttributes struct {
	internal.Internal
}

type QXmlAttributes_ITF interface {
	QXmlAttributes_PTR() *QXmlAttributes
}

func (ptr *QXmlAttributes) QXmlAttributes_PTR() *QXmlAttributes {
	return ptr
}

func (ptr *QXmlAttributes) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QXmlAttributes) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQXmlAttributes(ptr QXmlAttributes_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlAttributes_PTR().Pointer()
	}
	return nil
}

func (n *QXmlAttributes) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQXmlAttributesFromPointer(ptr unsafe.Pointer) (n *QXmlAttributes) {
	n = new(QXmlAttributes)
	n.InitFromInternal(uintptr(ptr), "xml.QXmlAttributes")
	return
}
func NewQXmlAttributes() *QXmlAttributes {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQXmlAttributes", ""}).(*QXmlAttributes)
}

func (ptr *QXmlAttributes) Append(qName string, uri string, localPart string, value string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Append", qName, uri, localPart, value})
}

func (ptr *QXmlAttributes) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QXmlAttributes) Count() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Count"}).(float64))
}

func (ptr *QXmlAttributes) Index(qName string) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Index", qName}).(float64))
}

func (ptr *QXmlAttributes) Index2(qName core.QLatin1String_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Index2", qName}).(float64))
}

func (ptr *QXmlAttributes) Index3(uri string, localPart string) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Index3", uri, localPart}).(float64))
}

func (ptr *QXmlAttributes) Length() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Length"}).(float64))
}

func (ptr *QXmlAttributes) LocalName(index int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LocalName", index}).(string)
}

func (ptr *QXmlAttributes) QName(index int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "QName", index}).(string)
}

func (ptr *QXmlAttributes) Swap(other QXmlAttributes_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QXmlAttributes) Type(index int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type", index}).(string)
}

func (ptr *QXmlAttributes) Type2(qName string) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type2", qName}).(string)
}

func (ptr *QXmlAttributes) Type3(uri string, localName string) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type3", uri, localName}).(string)
}

func (ptr *QXmlAttributes) Uri(index int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Uri", index}).(string)
}

func (ptr *QXmlAttributes) Value(index int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Value", index}).(string)
}

func (ptr *QXmlAttributes) Value2(qName string) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Value2", qName}).(string)
}

func (ptr *QXmlAttributes) Value3(qName core.QLatin1String_ITF) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Value3", qName}).(string)
}

func (ptr *QXmlAttributes) Value4(uri string, localName string) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Value4", uri, localName}).(string)
}

func (ptr *QXmlAttributes) ConnectDestroyQXmlAttributes(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQXmlAttributes", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlAttributes) DisconnectDestroyQXmlAttributes() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQXmlAttributes"})
}

func (ptr *QXmlAttributes) DestroyQXmlAttributes() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlAttributes"})
}

func (ptr *QXmlAttributes) DestroyQXmlAttributesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlAttributesDefault"})
}

type QXmlContentHandler struct {
	internal.Internal
}

type QXmlContentHandler_ITF interface {
	QXmlContentHandler_PTR() *QXmlContentHandler
}

func (ptr *QXmlContentHandler) QXmlContentHandler_PTR() *QXmlContentHandler {
	return ptr
}

func (ptr *QXmlContentHandler) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QXmlContentHandler) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQXmlContentHandler(ptr QXmlContentHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlContentHandler_PTR().Pointer()
	}
	return nil
}

func (n *QXmlContentHandler) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQXmlContentHandlerFromPointer(ptr unsafe.Pointer) (n *QXmlContentHandler) {
	n = new(QXmlContentHandler)
	n.InitFromInternal(uintptr(ptr), "xml.QXmlContentHandler")
	return
}
func (ptr *QXmlContentHandler) ConnectCharacters(f func(ch string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCharacters", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlContentHandler) DisconnectCharacters() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCharacters"})
}

func (ptr *QXmlContentHandler) Characters(ch string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Characters", ch}).(bool)
}

func (ptr *QXmlContentHandler) ConnectEndDocument(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectEndDocument", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlContentHandler) DisconnectEndDocument() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectEndDocument"})
}

func (ptr *QXmlContentHandler) EndDocument() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EndDocument"}).(bool)
}

func (ptr *QXmlContentHandler) ConnectEndElement(f func(namespaceURI string, localName string, qName string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectEndElement", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlContentHandler) DisconnectEndElement() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectEndElement"})
}

func (ptr *QXmlContentHandler) EndElement(namespaceURI string, localName string, qName string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EndElement", namespaceURI, localName, qName}).(bool)
}

func (ptr *QXmlContentHandler) ConnectEndPrefixMapping(f func(prefix string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectEndPrefixMapping", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlContentHandler) DisconnectEndPrefixMapping() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectEndPrefixMapping"})
}

func (ptr *QXmlContentHandler) EndPrefixMapping(prefix string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EndPrefixMapping", prefix}).(bool)
}

func (ptr *QXmlContentHandler) ConnectErrorString(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectErrorString", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlContentHandler) DisconnectErrorString() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectErrorString"})
}

func (ptr *QXmlContentHandler) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func (ptr *QXmlContentHandler) ConnectIgnorableWhitespace(f func(ch string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIgnorableWhitespace", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlContentHandler) DisconnectIgnorableWhitespace() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIgnorableWhitespace"})
}

func (ptr *QXmlContentHandler) IgnorableWhitespace(ch string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IgnorableWhitespace", ch}).(bool)
}

func (ptr *QXmlContentHandler) ConnectProcessingInstruction(f func(target string, data string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectProcessingInstruction", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlContentHandler) DisconnectProcessingInstruction() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectProcessingInstruction"})
}

func (ptr *QXmlContentHandler) ProcessingInstruction(target string, data string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProcessingInstruction", target, data}).(bool)
}

func (ptr *QXmlContentHandler) ConnectSetDocumentLocator(f func(locator *QXmlLocator)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetDocumentLocator", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlContentHandler) DisconnectSetDocumentLocator() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetDocumentLocator"})
}

func (ptr *QXmlContentHandler) SetDocumentLocator(locator QXmlLocator_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDocumentLocator", locator})
}

func (ptr *QXmlContentHandler) ConnectSkippedEntity(f func(name string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSkippedEntity", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlContentHandler) DisconnectSkippedEntity() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSkippedEntity"})
}

func (ptr *QXmlContentHandler) SkippedEntity(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SkippedEntity", name}).(bool)
}

func (ptr *QXmlContentHandler) ConnectStartDocument(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStartDocument", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlContentHandler) DisconnectStartDocument() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStartDocument"})
}

func (ptr *QXmlContentHandler) StartDocument() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartDocument"}).(bool)
}

func (ptr *QXmlContentHandler) ConnectStartElement(f func(namespaceURI string, localName string, qName string, atts *QXmlAttributes) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStartElement", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlContentHandler) DisconnectStartElement() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStartElement"})
}

func (ptr *QXmlContentHandler) StartElement(namespaceURI string, localName string, qName string, atts QXmlAttributes_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartElement", namespaceURI, localName, qName, atts}).(bool)
}

func (ptr *QXmlContentHandler) ConnectStartPrefixMapping(f func(prefix string, uri string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStartPrefixMapping", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlContentHandler) DisconnectStartPrefixMapping() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStartPrefixMapping"})
}

func (ptr *QXmlContentHandler) StartPrefixMapping(prefix string, uri string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartPrefixMapping", prefix, uri}).(bool)
}

func (ptr *QXmlContentHandler) ConnectDestroyQXmlContentHandler(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQXmlContentHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlContentHandler) DisconnectDestroyQXmlContentHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQXmlContentHandler"})
}

func (ptr *QXmlContentHandler) DestroyQXmlContentHandler() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlContentHandler"})
}

func (ptr *QXmlContentHandler) DestroyQXmlContentHandlerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlContentHandlerDefault"})
}

type QXmlDTDHandler struct {
	internal.Internal
}

type QXmlDTDHandler_ITF interface {
	QXmlDTDHandler_PTR() *QXmlDTDHandler
}

func (ptr *QXmlDTDHandler) QXmlDTDHandler_PTR() *QXmlDTDHandler {
	return ptr
}

func (ptr *QXmlDTDHandler) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QXmlDTDHandler) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQXmlDTDHandler(ptr QXmlDTDHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlDTDHandler_PTR().Pointer()
	}
	return nil
}

func (n *QXmlDTDHandler) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQXmlDTDHandlerFromPointer(ptr unsafe.Pointer) (n *QXmlDTDHandler) {
	n = new(QXmlDTDHandler)
	n.InitFromInternal(uintptr(ptr), "xml.QXmlDTDHandler")
	return
}
func (ptr *QXmlDTDHandler) ConnectErrorString(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectErrorString", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDTDHandler) DisconnectErrorString() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectErrorString"})
}

func (ptr *QXmlDTDHandler) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func (ptr *QXmlDTDHandler) ConnectNotationDecl(f func(name string, publicId string, systemId string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotationDecl", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDTDHandler) DisconnectNotationDecl() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotationDecl"})
}

func (ptr *QXmlDTDHandler) NotationDecl(name string, publicId string, systemId string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NotationDecl", name, publicId, systemId}).(bool)
}

func (ptr *QXmlDTDHandler) ConnectUnparsedEntityDecl(f func(name string, publicId string, systemId string, notationName string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUnparsedEntityDecl", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDTDHandler) DisconnectUnparsedEntityDecl() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUnparsedEntityDecl"})
}

func (ptr *QXmlDTDHandler) UnparsedEntityDecl(name string, publicId string, systemId string, notationName string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UnparsedEntityDecl", name, publicId, systemId, notationName}).(bool)
}

func (ptr *QXmlDTDHandler) ConnectDestroyQXmlDTDHandler(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQXmlDTDHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDTDHandler) DisconnectDestroyQXmlDTDHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQXmlDTDHandler"})
}

func (ptr *QXmlDTDHandler) DestroyQXmlDTDHandler() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlDTDHandler"})
}

func (ptr *QXmlDTDHandler) DestroyQXmlDTDHandlerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlDTDHandlerDefault"})
}

type QXmlDeclHandler struct {
	internal.Internal
}

type QXmlDeclHandler_ITF interface {
	QXmlDeclHandler_PTR() *QXmlDeclHandler
}

func (ptr *QXmlDeclHandler) QXmlDeclHandler_PTR() *QXmlDeclHandler {
	return ptr
}

func (ptr *QXmlDeclHandler) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QXmlDeclHandler) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQXmlDeclHandler(ptr QXmlDeclHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlDeclHandler_PTR().Pointer()
	}
	return nil
}

func (n *QXmlDeclHandler) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQXmlDeclHandlerFromPointer(ptr unsafe.Pointer) (n *QXmlDeclHandler) {
	n = new(QXmlDeclHandler)
	n.InitFromInternal(uintptr(ptr), "xml.QXmlDeclHandler")
	return
}
func (ptr *QXmlDeclHandler) ConnectAttributeDecl(f func(eName string, aName string, ty string, valueDefault string, value string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAttributeDecl", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDeclHandler) DisconnectAttributeDecl() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAttributeDecl"})
}

func (ptr *QXmlDeclHandler) AttributeDecl(eName string, aName string, ty string, valueDefault string, value string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AttributeDecl", eName, aName, ty, valueDefault, value}).(bool)
}

func (ptr *QXmlDeclHandler) ConnectErrorString(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectErrorString", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDeclHandler) DisconnectErrorString() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectErrorString"})
}

func (ptr *QXmlDeclHandler) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func (ptr *QXmlDeclHandler) ConnectExternalEntityDecl(f func(name string, publicId string, systemId string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectExternalEntityDecl", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDeclHandler) DisconnectExternalEntityDecl() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectExternalEntityDecl"})
}

func (ptr *QXmlDeclHandler) ExternalEntityDecl(name string, publicId string, systemId string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExternalEntityDecl", name, publicId, systemId}).(bool)
}

func (ptr *QXmlDeclHandler) ConnectInternalEntityDecl(f func(name string, value string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInternalEntityDecl", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDeclHandler) DisconnectInternalEntityDecl() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInternalEntityDecl"})
}

func (ptr *QXmlDeclHandler) InternalEntityDecl(name string, value string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InternalEntityDecl", name, value}).(bool)
}

func (ptr *QXmlDeclHandler) ConnectDestroyQXmlDeclHandler(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQXmlDeclHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDeclHandler) DisconnectDestroyQXmlDeclHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQXmlDeclHandler"})
}

func (ptr *QXmlDeclHandler) DestroyQXmlDeclHandler() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlDeclHandler"})
}

func (ptr *QXmlDeclHandler) DestroyQXmlDeclHandlerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlDeclHandlerDefault"})
}

type QXmlDefaultHandler struct {
	QXmlContentHandler
	QXmlErrorHandler
	QXmlDTDHandler
	QXmlEntityResolver
	QXmlLexicalHandler
	QXmlDeclHandler
}

type QXmlDefaultHandler_ITF interface {
	QXmlContentHandler_ITF
	QXmlErrorHandler_ITF
	QXmlDTDHandler_ITF
	QXmlEntityResolver_ITF
	QXmlLexicalHandler_ITF
	QXmlDeclHandler_ITF
	QXmlDefaultHandler_PTR() *QXmlDefaultHandler
}

func (ptr *QXmlDefaultHandler) QXmlDefaultHandler_PTR() *QXmlDefaultHandler {
	return ptr
}

func (ptr *QXmlDefaultHandler) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlContentHandler_PTR().Pointer()
	}
	return nil
}

func (ptr *QXmlDefaultHandler) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QXmlContentHandler_PTR().SetPointer(p)
		ptr.QXmlErrorHandler_PTR().SetPointer(p)
		ptr.QXmlDTDHandler_PTR().SetPointer(p)
		ptr.QXmlEntityResolver_PTR().SetPointer(p)
		ptr.QXmlLexicalHandler_PTR().SetPointer(p)
		ptr.QXmlDeclHandler_PTR().SetPointer(p)
	}
}

func PointerFromQXmlDefaultHandler(ptr QXmlDefaultHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlDefaultHandler_PTR().Pointer()
	}
	return nil
}

func (n *QXmlDefaultHandler) InitFromInternal(ptr uintptr, name string) {
	n.QXmlContentHandler_PTR().InitFromInternal(uintptr(ptr), name)
	n.QXmlErrorHandler_PTR().InitFromInternal(uintptr(ptr), name)
	n.QXmlDTDHandler_PTR().InitFromInternal(uintptr(ptr), name)
	n.QXmlEntityResolver_PTR().InitFromInternal(uintptr(ptr), name)
	n.QXmlLexicalHandler_PTR().InitFromInternal(uintptr(ptr), name)
	n.QXmlDeclHandler_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QXmlDefaultHandler) ClassNameInternalF() string {
	return n.QXmlContentHandler_PTR().ClassNameInternalF()
}

func NewQXmlDefaultHandlerFromPointer(ptr unsafe.Pointer) (n *QXmlDefaultHandler) {
	n = new(QXmlDefaultHandler)
	n.InitFromInternal(uintptr(ptr), "xml.QXmlDefaultHandler")
	return
}
func NewQXmlDefaultHandler() *QXmlDefaultHandler {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQXmlDefaultHandler", ""}).(*QXmlDefaultHandler)
}

func (ptr *QXmlDefaultHandler) ConnectAttributeDecl(f func(eName string, aName string, ty string, valueDefault string, value string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectAttributeDecl", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectAttributeDecl() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectAttributeDecl"})
}

func (ptr *QXmlDefaultHandler) AttributeDecl(eName string, aName string, ty string, valueDefault string, value string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "AttributeDecl", eName, aName, ty, valueDefault, value}).(bool)
}

func (ptr *QXmlDefaultHandler) AttributeDeclDefault(eName string, aName string, ty string, valueDefault string, value string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "AttributeDeclDefault", eName, aName, ty, valueDefault, value}).(bool)
}

func (ptr *QXmlDefaultHandler) ConnectCharacters(f func(ch string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectCharacters", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectCharacters() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectCharacters"})
}

func (ptr *QXmlDefaultHandler) Characters(ch string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "Characters", ch}).(bool)
}

func (ptr *QXmlDefaultHandler) CharactersDefault(ch string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "CharactersDefault", ch}).(bool)
}

func (ptr *QXmlDefaultHandler) ConnectComment(f func(ch string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectComment", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectComment() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectComment"})
}

func (ptr *QXmlDefaultHandler) Comment(ch string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "Comment", ch}).(bool)
}

func (ptr *QXmlDefaultHandler) CommentDefault(ch string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "CommentDefault", ch}).(bool)
}

func (ptr *QXmlDefaultHandler) ConnectEndCDATA(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectEndCDATA", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectEndCDATA() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectEndCDATA"})
}

func (ptr *QXmlDefaultHandler) EndCDATA() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "EndCDATA"}).(bool)
}

func (ptr *QXmlDefaultHandler) EndCDATADefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "EndCDATADefault"}).(bool)
}

func (ptr *QXmlDefaultHandler) ConnectEndDTD(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectEndDTD", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectEndDTD() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectEndDTD"})
}

func (ptr *QXmlDefaultHandler) EndDTD() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "EndDTD"}).(bool)
}

func (ptr *QXmlDefaultHandler) EndDTDDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "EndDTDDefault"}).(bool)
}

func (ptr *QXmlDefaultHandler) ConnectEndDocument(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectEndDocument", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectEndDocument() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectEndDocument"})
}

func (ptr *QXmlDefaultHandler) EndDocument() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "EndDocument"}).(bool)
}

func (ptr *QXmlDefaultHandler) EndDocumentDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "EndDocumentDefault"}).(bool)
}

func (ptr *QXmlDefaultHandler) ConnectEndElement(f func(namespaceURI string, localName string, qName string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectEndElement", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectEndElement() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectEndElement"})
}

func (ptr *QXmlDefaultHandler) EndElement(namespaceURI string, localName string, qName string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "EndElement", namespaceURI, localName, qName}).(bool)
}

func (ptr *QXmlDefaultHandler) EndElementDefault(namespaceURI string, localName string, qName string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "EndElementDefault", namespaceURI, localName, qName}).(bool)
}

func (ptr *QXmlDefaultHandler) ConnectEndEntity(f func(name string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectEndEntity", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectEndEntity() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectEndEntity"})
}

func (ptr *QXmlDefaultHandler) EndEntity(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "EndEntity", name}).(bool)
}

func (ptr *QXmlDefaultHandler) EndEntityDefault(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "EndEntityDefault", name}).(bool)
}

func (ptr *QXmlDefaultHandler) ConnectEndPrefixMapping(f func(prefix string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectEndPrefixMapping", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectEndPrefixMapping() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectEndPrefixMapping"})
}

func (ptr *QXmlDefaultHandler) EndPrefixMapping(prefix string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "EndPrefixMapping", prefix}).(bool)
}

func (ptr *QXmlDefaultHandler) EndPrefixMappingDefault(prefix string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "EndPrefixMappingDefault", prefix}).(bool)
}

func (ptr *QXmlDefaultHandler) ConnectError(f func(exception *QXmlParseException) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectError", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectError() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectError"})
}

func (ptr *QXmlDefaultHandler) Error(exception QXmlParseException_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "Error", exception}).(bool)
}

func (ptr *QXmlDefaultHandler) ErrorDefault(exception QXmlParseException_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ErrorDefault", exception}).(bool)
}

func (ptr *QXmlDefaultHandler) ConnectErrorString(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectErrorString", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectErrorString() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectErrorString"})
}

func (ptr *QXmlDefaultHandler) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ErrorString"}).(string)
}

func (ptr *QXmlDefaultHandler) ErrorStringDefault() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ErrorStringDefault"}).(string)
}

func (ptr *QXmlDefaultHandler) ConnectExternalEntityDecl(f func(name string, publicId string, systemId string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectExternalEntityDecl", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectExternalEntityDecl() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectExternalEntityDecl"})
}

func (ptr *QXmlDefaultHandler) ExternalEntityDecl(name string, publicId string, systemId string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ExternalEntityDecl", name, publicId, systemId}).(bool)
}

func (ptr *QXmlDefaultHandler) ExternalEntityDeclDefault(name string, publicId string, systemId string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ExternalEntityDeclDefault", name, publicId, systemId}).(bool)
}

func (ptr *QXmlDefaultHandler) ConnectFatalError(f func(exception *QXmlParseException) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectFatalError", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectFatalError() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectFatalError"})
}

func (ptr *QXmlDefaultHandler) FatalError(exception QXmlParseException_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "FatalError", exception}).(bool)
}

func (ptr *QXmlDefaultHandler) FatalErrorDefault(exception QXmlParseException_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "FatalErrorDefault", exception}).(bool)
}

func (ptr *QXmlDefaultHandler) ConnectIgnorableWhitespace(f func(ch string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectIgnorableWhitespace", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectIgnorableWhitespace() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectIgnorableWhitespace"})
}

func (ptr *QXmlDefaultHandler) IgnorableWhitespace(ch string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "IgnorableWhitespace", ch}).(bool)
}

func (ptr *QXmlDefaultHandler) IgnorableWhitespaceDefault(ch string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "IgnorableWhitespaceDefault", ch}).(bool)
}

func (ptr *QXmlDefaultHandler) ConnectInternalEntityDecl(f func(name string, value string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectInternalEntityDecl", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectInternalEntityDecl() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectInternalEntityDecl"})
}

func (ptr *QXmlDefaultHandler) InternalEntityDecl(name string, value string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "InternalEntityDecl", name, value}).(bool)
}

func (ptr *QXmlDefaultHandler) InternalEntityDeclDefault(name string, value string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "InternalEntityDeclDefault", name, value}).(bool)
}

func (ptr *QXmlDefaultHandler) ConnectNotationDecl(f func(name string, publicId string, systemId string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectNotationDecl", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectNotationDecl() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectNotationDecl"})
}

func (ptr *QXmlDefaultHandler) NotationDecl(name string, publicId string, systemId string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "NotationDecl", name, publicId, systemId}).(bool)
}

func (ptr *QXmlDefaultHandler) NotationDeclDefault(name string, publicId string, systemId string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "NotationDeclDefault", name, publicId, systemId}).(bool)
}

func (ptr *QXmlDefaultHandler) ConnectProcessingInstruction(f func(target string, data string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectProcessingInstruction", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectProcessingInstruction() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectProcessingInstruction"})
}

func (ptr *QXmlDefaultHandler) ProcessingInstruction(target string, data string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ProcessingInstruction", target, data}).(bool)
}

func (ptr *QXmlDefaultHandler) ProcessingInstructionDefault(target string, data string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ProcessingInstructionDefault", target, data}).(bool)
}

func (ptr *QXmlDefaultHandler) ConnectSetDocumentLocator(f func(locator *QXmlLocator)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectSetDocumentLocator", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectSetDocumentLocator() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectSetDocumentLocator"})
}

func (ptr *QXmlDefaultHandler) SetDocumentLocator(locator QXmlLocator_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "SetDocumentLocator", locator})
}

func (ptr *QXmlDefaultHandler) SetDocumentLocatorDefault(locator QXmlLocator_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "SetDocumentLocatorDefault", locator})
}

func (ptr *QXmlDefaultHandler) ConnectSkippedEntity(f func(name string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectSkippedEntity", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectSkippedEntity() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectSkippedEntity"})
}

func (ptr *QXmlDefaultHandler) SkippedEntity(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "SkippedEntity", name}).(bool)
}

func (ptr *QXmlDefaultHandler) SkippedEntityDefault(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "SkippedEntityDefault", name}).(bool)
}

func (ptr *QXmlDefaultHandler) ConnectStartCDATA(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectStartCDATA", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectStartCDATA() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectStartCDATA"})
}

func (ptr *QXmlDefaultHandler) StartCDATA() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "StartCDATA"}).(bool)
}

func (ptr *QXmlDefaultHandler) StartCDATADefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "StartCDATADefault"}).(bool)
}

func (ptr *QXmlDefaultHandler) ConnectStartDTD(f func(name string, publicId string, systemId string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectStartDTD", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectStartDTD() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectStartDTD"})
}

func (ptr *QXmlDefaultHandler) StartDTD(name string, publicId string, systemId string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "StartDTD", name, publicId, systemId}).(bool)
}

func (ptr *QXmlDefaultHandler) StartDTDDefault(name string, publicId string, systemId string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "StartDTDDefault", name, publicId, systemId}).(bool)
}

func (ptr *QXmlDefaultHandler) ConnectStartDocument(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectStartDocument", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectStartDocument() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectStartDocument"})
}

func (ptr *QXmlDefaultHandler) StartDocument() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "StartDocument"}).(bool)
}

func (ptr *QXmlDefaultHandler) StartDocumentDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "StartDocumentDefault"}).(bool)
}

func (ptr *QXmlDefaultHandler) ConnectStartElement(f func(namespaceURI string, localName string, qName string, atts *QXmlAttributes) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectStartElement", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectStartElement() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectStartElement"})
}

func (ptr *QXmlDefaultHandler) StartElement(namespaceURI string, localName string, qName string, atts QXmlAttributes_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "StartElement", namespaceURI, localName, qName, atts}).(bool)
}

func (ptr *QXmlDefaultHandler) StartElementDefault(namespaceURI string, localName string, qName string, atts QXmlAttributes_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "StartElementDefault", namespaceURI, localName, qName, atts}).(bool)
}

func (ptr *QXmlDefaultHandler) ConnectStartEntity(f func(name string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectStartEntity", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectStartEntity() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectStartEntity"})
}

func (ptr *QXmlDefaultHandler) StartEntity(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "StartEntity", name}).(bool)
}

func (ptr *QXmlDefaultHandler) StartEntityDefault(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "StartEntityDefault", name}).(bool)
}

func (ptr *QXmlDefaultHandler) ConnectStartPrefixMapping(f func(prefix string, uri string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectStartPrefixMapping", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectStartPrefixMapping() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectStartPrefixMapping"})
}

func (ptr *QXmlDefaultHandler) StartPrefixMapping(prefix string, uri string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "StartPrefixMapping", prefix, uri}).(bool)
}

func (ptr *QXmlDefaultHandler) StartPrefixMappingDefault(prefix string, uri string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "StartPrefixMappingDefault", prefix, uri}).(bool)
}

func (ptr *QXmlDefaultHandler) ConnectUnparsedEntityDecl(f func(name string, publicId string, systemId string, notationName string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectUnparsedEntityDecl", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectUnparsedEntityDecl() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectUnparsedEntityDecl"})
}

func (ptr *QXmlDefaultHandler) UnparsedEntityDecl(name string, publicId string, systemId string, notationName string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "UnparsedEntityDecl", name, publicId, systemId, notationName}).(bool)
}

func (ptr *QXmlDefaultHandler) UnparsedEntityDeclDefault(name string, publicId string, systemId string, notationName string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "UnparsedEntityDeclDefault", name, publicId, systemId, notationName}).(bool)
}

func (ptr *QXmlDefaultHandler) ConnectWarning(f func(exception *QXmlParseException) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectWarning", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectWarning() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectWarning"})
}

func (ptr *QXmlDefaultHandler) Warning(exception QXmlParseException_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "Warning", exception}).(bool)
}

func (ptr *QXmlDefaultHandler) WarningDefault(exception QXmlParseException_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "WarningDefault", exception}).(bool)
}

func (ptr *QXmlDefaultHandler) ConnectDestroyQXmlDefaultHandler(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "ConnectDestroyQXmlDefaultHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlDefaultHandler) DisconnectDestroyQXmlDefaultHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DisconnectDestroyQXmlDefaultHandler"})
}

func (ptr *QXmlDefaultHandler) DestroyQXmlDefaultHandler() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DestroyQXmlDefaultHandler"})
}

func (ptr *QXmlDefaultHandler) DestroyQXmlDefaultHandlerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QXmlContentHandler_PTR().ClassNameInternalF(), "DestroyQXmlDefaultHandlerDefault"})
}

type QXmlEntityResolver struct {
	internal.Internal
}

type QXmlEntityResolver_ITF interface {
	QXmlEntityResolver_PTR() *QXmlEntityResolver
}

func (ptr *QXmlEntityResolver) QXmlEntityResolver_PTR() *QXmlEntityResolver {
	return ptr
}

func (ptr *QXmlEntityResolver) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QXmlEntityResolver) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQXmlEntityResolver(ptr QXmlEntityResolver_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlEntityResolver_PTR().Pointer()
	}
	return nil
}

func (n *QXmlEntityResolver) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQXmlEntityResolverFromPointer(ptr unsafe.Pointer) (n *QXmlEntityResolver) {
	n = new(QXmlEntityResolver)
	n.InitFromInternal(uintptr(ptr), "xml.QXmlEntityResolver")
	return
}
func (ptr *QXmlEntityResolver) ConnectErrorString(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectErrorString", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlEntityResolver) DisconnectErrorString() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectErrorString"})
}

func (ptr *QXmlEntityResolver) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func (ptr *QXmlEntityResolver) ConnectDestroyQXmlEntityResolver(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQXmlEntityResolver", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlEntityResolver) DisconnectDestroyQXmlEntityResolver() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQXmlEntityResolver"})
}

func (ptr *QXmlEntityResolver) DestroyQXmlEntityResolver() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlEntityResolver"})
}

func (ptr *QXmlEntityResolver) DestroyQXmlEntityResolverDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlEntityResolverDefault"})
}

type QXmlErrorHandler struct {
	internal.Internal
}

type QXmlErrorHandler_ITF interface {
	QXmlErrorHandler_PTR() *QXmlErrorHandler
}

func (ptr *QXmlErrorHandler) QXmlErrorHandler_PTR() *QXmlErrorHandler {
	return ptr
}

func (ptr *QXmlErrorHandler) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QXmlErrorHandler) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQXmlErrorHandler(ptr QXmlErrorHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlErrorHandler_PTR().Pointer()
	}
	return nil
}

func (n *QXmlErrorHandler) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQXmlErrorHandlerFromPointer(ptr unsafe.Pointer) (n *QXmlErrorHandler) {
	n = new(QXmlErrorHandler)
	n.InitFromInternal(uintptr(ptr), "xml.QXmlErrorHandler")
	return
}
func (ptr *QXmlErrorHandler) ConnectError(f func(exception *QXmlParseException) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectError", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlErrorHandler) DisconnectError() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError"})
}

func (ptr *QXmlErrorHandler) Error(exception QXmlParseException_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error", exception}).(bool)
}

func (ptr *QXmlErrorHandler) ConnectErrorString(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectErrorString", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlErrorHandler) DisconnectErrorString() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectErrorString"})
}

func (ptr *QXmlErrorHandler) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func (ptr *QXmlErrorHandler) ConnectFatalError(f func(exception *QXmlParseException) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFatalError", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlErrorHandler) DisconnectFatalError() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFatalError"})
}

func (ptr *QXmlErrorHandler) FatalError(exception QXmlParseException_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FatalError", exception}).(bool)
}

func (ptr *QXmlErrorHandler) ConnectWarning(f func(exception *QXmlParseException) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWarning", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlErrorHandler) DisconnectWarning() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWarning"})
}

func (ptr *QXmlErrorHandler) Warning(exception QXmlParseException_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Warning", exception}).(bool)
}

func (ptr *QXmlErrorHandler) ConnectDestroyQXmlErrorHandler(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQXmlErrorHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlErrorHandler) DisconnectDestroyQXmlErrorHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQXmlErrorHandler"})
}

func (ptr *QXmlErrorHandler) DestroyQXmlErrorHandler() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlErrorHandler"})
}

func (ptr *QXmlErrorHandler) DestroyQXmlErrorHandlerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlErrorHandlerDefault"})
}

type QXmlInputSource struct {
	internal.Internal
}

type QXmlInputSource_ITF interface {
	QXmlInputSource_PTR() *QXmlInputSource
}

func (ptr *QXmlInputSource) QXmlInputSource_PTR() *QXmlInputSource {
	return ptr
}

func (ptr *QXmlInputSource) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QXmlInputSource) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQXmlInputSource(ptr QXmlInputSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlInputSource_PTR().Pointer()
	}
	return nil
}

func (n *QXmlInputSource) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQXmlInputSourceFromPointer(ptr unsafe.Pointer) (n *QXmlInputSource) {
	n = new(QXmlInputSource)
	n.InitFromInternal(uintptr(ptr), "xml.QXmlInputSource")
	return
}
func NewQXmlInputSource() *QXmlInputSource {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQXmlInputSource", ""}).(*QXmlInputSource)
}

func NewQXmlInputSource2(dev core.QIODevice_ITF) *QXmlInputSource {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQXmlInputSource2", "", dev}).(*QXmlInputSource)
}

func (ptr *QXmlInputSource) ConnectData(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlInputSource) DisconnectData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectData"})
}

func (ptr *QXmlInputSource) Data() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Data"}).(string)
}

func (ptr *QXmlInputSource) DataDefault() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DataDefault"}).(string)
}

func (ptr *QXmlInputSource) ConnectFetchData(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFetchData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlInputSource) DisconnectFetchData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFetchData"})
}

func (ptr *QXmlInputSource) FetchData() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FetchData"})
}

func (ptr *QXmlInputSource) FetchDataDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FetchDataDefault"})
}

func (ptr *QXmlInputSource) ConnectFromRawData(f func(data *core.QByteArray, beginning bool) string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFromRawData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlInputSource) DisconnectFromRawData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFromRawData"})
}

func (ptr *QXmlInputSource) FromRawData(data core.QByteArray_ITF, beginning bool) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FromRawData", data, beginning}).(string)
}

func (ptr *QXmlInputSource) FromRawDataDefault(data core.QByteArray_ITF, beginning bool) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FromRawDataDefault", data, beginning}).(string)
}

func (ptr *QXmlInputSource) ConnectNext(f func() *core.QChar) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNext", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlInputSource) DisconnectNext() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNext"})
}

func (ptr *QXmlInputSource) Next() *core.QChar {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Next"}).(*core.QChar)
}

func (ptr *QXmlInputSource) NextDefault() *core.QChar {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NextDefault"}).(*core.QChar)
}

func (ptr *QXmlInputSource) ConnectReset(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReset", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlInputSource) DisconnectReset() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReset"})
}

func (ptr *QXmlInputSource) Reset() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reset"})
}

func (ptr *QXmlInputSource) ResetDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResetDefault"})
}

func (ptr *QXmlInputSource) ConnectSetData(f func(dat string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlInputSource) DisconnectSetData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetData"})
}

func (ptr *QXmlInputSource) SetData(dat string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetData", dat})
}

func (ptr *QXmlInputSource) SetDataDefault(dat string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDataDefault", dat})
}

func (ptr *QXmlInputSource) ConnectSetData2(f func(dat *core.QByteArray)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetData2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlInputSource) DisconnectSetData2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetData2"})
}

func (ptr *QXmlInputSource) SetData2(dat core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetData2", dat})
}

func (ptr *QXmlInputSource) SetData2Default(dat core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetData2Default", dat})
}

func (ptr *QXmlInputSource) ConnectDestroyQXmlInputSource(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQXmlInputSource", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlInputSource) DisconnectDestroyQXmlInputSource() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQXmlInputSource"})
}

func (ptr *QXmlInputSource) DestroyQXmlInputSource() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlInputSource"})
}

func (ptr *QXmlInputSource) DestroyQXmlInputSourceDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlInputSourceDefault"})
}

type QXmlLexicalHandler struct {
	internal.Internal
}

type QXmlLexicalHandler_ITF interface {
	QXmlLexicalHandler_PTR() *QXmlLexicalHandler
}

func (ptr *QXmlLexicalHandler) QXmlLexicalHandler_PTR() *QXmlLexicalHandler {
	return ptr
}

func (ptr *QXmlLexicalHandler) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QXmlLexicalHandler) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQXmlLexicalHandler(ptr QXmlLexicalHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlLexicalHandler_PTR().Pointer()
	}
	return nil
}

func (n *QXmlLexicalHandler) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQXmlLexicalHandlerFromPointer(ptr unsafe.Pointer) (n *QXmlLexicalHandler) {
	n = new(QXmlLexicalHandler)
	n.InitFromInternal(uintptr(ptr), "xml.QXmlLexicalHandler")
	return
}
func (ptr *QXmlLexicalHandler) ConnectComment(f func(ch string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectComment", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlLexicalHandler) DisconnectComment() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectComment"})
}

func (ptr *QXmlLexicalHandler) Comment(ch string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Comment", ch}).(bool)
}

func (ptr *QXmlLexicalHandler) ConnectEndCDATA(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectEndCDATA", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlLexicalHandler) DisconnectEndCDATA() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectEndCDATA"})
}

func (ptr *QXmlLexicalHandler) EndCDATA() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EndCDATA"}).(bool)
}

func (ptr *QXmlLexicalHandler) ConnectEndDTD(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectEndDTD", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlLexicalHandler) DisconnectEndDTD() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectEndDTD"})
}

func (ptr *QXmlLexicalHandler) EndDTD() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EndDTD"}).(bool)
}

func (ptr *QXmlLexicalHandler) ConnectEndEntity(f func(name string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectEndEntity", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlLexicalHandler) DisconnectEndEntity() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectEndEntity"})
}

func (ptr *QXmlLexicalHandler) EndEntity(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EndEntity", name}).(bool)
}

func (ptr *QXmlLexicalHandler) ConnectErrorString(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectErrorString", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlLexicalHandler) DisconnectErrorString() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectErrorString"})
}

func (ptr *QXmlLexicalHandler) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func (ptr *QXmlLexicalHandler) ConnectStartCDATA(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStartCDATA", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlLexicalHandler) DisconnectStartCDATA() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStartCDATA"})
}

func (ptr *QXmlLexicalHandler) StartCDATA() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartCDATA"}).(bool)
}

func (ptr *QXmlLexicalHandler) ConnectStartDTD(f func(name string, publicId string, systemId string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStartDTD", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlLexicalHandler) DisconnectStartDTD() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStartDTD"})
}

func (ptr *QXmlLexicalHandler) StartDTD(name string, publicId string, systemId string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartDTD", name, publicId, systemId}).(bool)
}

func (ptr *QXmlLexicalHandler) ConnectStartEntity(f func(name string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStartEntity", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlLexicalHandler) DisconnectStartEntity() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStartEntity"})
}

func (ptr *QXmlLexicalHandler) StartEntity(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartEntity", name}).(bool)
}

func (ptr *QXmlLexicalHandler) ConnectDestroyQXmlLexicalHandler(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQXmlLexicalHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlLexicalHandler) DisconnectDestroyQXmlLexicalHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQXmlLexicalHandler"})
}

func (ptr *QXmlLexicalHandler) DestroyQXmlLexicalHandler() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlLexicalHandler"})
}

func (ptr *QXmlLexicalHandler) DestroyQXmlLexicalHandlerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlLexicalHandlerDefault"})
}

type QXmlLocator struct {
	internal.Internal
}

type QXmlLocator_ITF interface {
	QXmlLocator_PTR() *QXmlLocator
}

func (ptr *QXmlLocator) QXmlLocator_PTR() *QXmlLocator {
	return ptr
}

func (ptr *QXmlLocator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QXmlLocator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQXmlLocator(ptr QXmlLocator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlLocator_PTR().Pointer()
	}
	return nil
}

func (n *QXmlLocator) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQXmlLocatorFromPointer(ptr unsafe.Pointer) (n *QXmlLocator) {
	n = new(QXmlLocator)
	n.InitFromInternal(uintptr(ptr), "xml.QXmlLocator")
	return
}
func NewQXmlLocator() *QXmlLocator {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQXmlLocator", ""}).(*QXmlLocator)
}

func (ptr *QXmlLocator) ConnectColumnNumber(f func() int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColumnNumber", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlLocator) DisconnectColumnNumber() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColumnNumber"})
}

func (ptr *QXmlLocator) ColumnNumber() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnNumber"}).(float64))
}

func (ptr *QXmlLocator) ConnectLineNumber(f func() int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLineNumber", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlLocator) DisconnectLineNumber() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLineNumber"})
}

func (ptr *QXmlLocator) LineNumber() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LineNumber"}).(float64))
}

func (ptr *QXmlLocator) ConnectDestroyQXmlLocator(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQXmlLocator", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlLocator) DisconnectDestroyQXmlLocator() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQXmlLocator"})
}

func (ptr *QXmlLocator) DestroyQXmlLocator() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlLocator"})
}

func (ptr *QXmlLocator) DestroyQXmlLocatorDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlLocatorDefault"})
}

type QXmlNamespaceSupport struct {
	internal.Internal
}

type QXmlNamespaceSupport_ITF interface {
	QXmlNamespaceSupport_PTR() *QXmlNamespaceSupport
}

func (ptr *QXmlNamespaceSupport) QXmlNamespaceSupport_PTR() *QXmlNamespaceSupport {
	return ptr
}

func (ptr *QXmlNamespaceSupport) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QXmlNamespaceSupport) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQXmlNamespaceSupport(ptr QXmlNamespaceSupport_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlNamespaceSupport_PTR().Pointer()
	}
	return nil
}

func (n *QXmlNamespaceSupport) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQXmlNamespaceSupportFromPointer(ptr unsafe.Pointer) (n *QXmlNamespaceSupport) {
	n = new(QXmlNamespaceSupport)
	n.InitFromInternal(uintptr(ptr), "xml.QXmlNamespaceSupport")
	return
}
func NewQXmlNamespaceSupport() *QXmlNamespaceSupport {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQXmlNamespaceSupport", ""}).(*QXmlNamespaceSupport)
}

func (ptr *QXmlNamespaceSupport) PopContext() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PopContext"})
}

func (ptr *QXmlNamespaceSupport) Prefix(uri string) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Prefix", uri}).(string)
}

func (ptr *QXmlNamespaceSupport) Prefixes() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Prefixes"}).([]string)
}

func (ptr *QXmlNamespaceSupport) Prefixes2(uri string) []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Prefixes2", uri}).([]string)
}

func (ptr *QXmlNamespaceSupport) ProcessName(qname string, isAttribute bool, nsuri string, localname string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProcessName", qname, isAttribute, nsuri, localname})
}

func (ptr *QXmlNamespaceSupport) PushContext() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PushContext"})
}

func (ptr *QXmlNamespaceSupport) Reset() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reset"})
}

func (ptr *QXmlNamespaceSupport) SetPrefix(pre string, uri string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPrefix", pre, uri})
}

func (ptr *QXmlNamespaceSupport) SplitName(qname string, prefix string, localname string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SplitName", qname, prefix, localname})
}

func (ptr *QXmlNamespaceSupport) Uri(prefix string) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Uri", prefix}).(string)
}

func (ptr *QXmlNamespaceSupport) DestroyQXmlNamespaceSupport() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlNamespaceSupport"})
}

type QXmlParseException struct {
	internal.Internal
}

type QXmlParseException_ITF interface {
	QXmlParseException_PTR() *QXmlParseException
}

func (ptr *QXmlParseException) QXmlParseException_PTR() *QXmlParseException {
	return ptr
}

func (ptr *QXmlParseException) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QXmlParseException) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQXmlParseException(ptr QXmlParseException_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlParseException_PTR().Pointer()
	}
	return nil
}

func (n *QXmlParseException) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQXmlParseExceptionFromPointer(ptr unsafe.Pointer) (n *QXmlParseException) {
	n = new(QXmlParseException)
	n.InitFromInternal(uintptr(ptr), "xml.QXmlParseException")
	return
}
func NewQXmlParseException(name string, c int, l int, p string, s string) *QXmlParseException {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQXmlParseException", "", name, c, l, p, s}).(*QXmlParseException)
}

func NewQXmlParseException2(other QXmlParseException_ITF) *QXmlParseException {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQXmlParseException2", "", other}).(*QXmlParseException)
}

func (ptr *QXmlParseException) ColumnNumber() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnNumber"}).(float64))
}

func (ptr *QXmlParseException) LineNumber() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LineNumber"}).(float64))
}

func (ptr *QXmlParseException) Message() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Message"}).(string)
}

func (ptr *QXmlParseException) PublicId() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PublicId"}).(string)
}

func (ptr *QXmlParseException) SystemId() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SystemId"}).(string)
}

func (ptr *QXmlParseException) DestroyQXmlParseException() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlParseException"})
}

type QXmlReader struct {
	internal.Internal
}

type QXmlReader_ITF interface {
	QXmlReader_PTR() *QXmlReader
}

func (ptr *QXmlReader) QXmlReader_PTR() *QXmlReader {
	return ptr
}

func (ptr *QXmlReader) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QXmlReader) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQXmlReader(ptr QXmlReader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlReader_PTR().Pointer()
	}
	return nil
}

func (n *QXmlReader) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQXmlReaderFromPointer(ptr unsafe.Pointer) (n *QXmlReader) {
	n = new(QXmlReader)
	n.InitFromInternal(uintptr(ptr), "xml.QXmlReader")
	return
}
func (ptr *QXmlReader) ConnectDTDHandler(f func() *QXmlDTDHandler) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDTDHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlReader) DisconnectDTDHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDTDHandler"})
}

func (ptr *QXmlReader) DTDHandler() *QXmlDTDHandler {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DTDHandler"}).(*QXmlDTDHandler)
}

func (ptr *QXmlReader) ConnectContentHandler(f func() *QXmlContentHandler) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectContentHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlReader) DisconnectContentHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectContentHandler"})
}

func (ptr *QXmlReader) ContentHandler() *QXmlContentHandler {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContentHandler"}).(*QXmlContentHandler)
}

func (ptr *QXmlReader) ConnectDeclHandler(f func() *QXmlDeclHandler) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDeclHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlReader) DisconnectDeclHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDeclHandler"})
}

func (ptr *QXmlReader) DeclHandler() *QXmlDeclHandler {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeclHandler"}).(*QXmlDeclHandler)
}

func (ptr *QXmlReader) ConnectEntityResolver(f func() *QXmlEntityResolver) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectEntityResolver", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlReader) DisconnectEntityResolver() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectEntityResolver"})
}

func (ptr *QXmlReader) EntityResolver() *QXmlEntityResolver {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EntityResolver"}).(*QXmlEntityResolver)
}

func (ptr *QXmlReader) ConnectErrorHandler(f func() *QXmlErrorHandler) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectErrorHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlReader) DisconnectErrorHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectErrorHandler"})
}

func (ptr *QXmlReader) ErrorHandler() *QXmlErrorHandler {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorHandler"}).(*QXmlErrorHandler)
}

func (ptr *QXmlReader) ConnectFeature(f func(name string, ok *bool) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFeature", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlReader) DisconnectFeature() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFeature"})
}

func (ptr *QXmlReader) Feature(name string, ok *bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Feature", name, ok}).(bool)
}

func (ptr *QXmlReader) ConnectHasFeature(f func(name string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHasFeature", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlReader) DisconnectHasFeature() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHasFeature"})
}

func (ptr *QXmlReader) HasFeature(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasFeature", name}).(bool)
}

func (ptr *QXmlReader) ConnectHasProperty(f func(name string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHasProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlReader) DisconnectHasProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHasProperty"})
}

func (ptr *QXmlReader) HasProperty(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasProperty", name}).(bool)
}

func (ptr *QXmlReader) ConnectLexicalHandler(f func() *QXmlLexicalHandler) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLexicalHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlReader) DisconnectLexicalHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLexicalHandler"})
}

func (ptr *QXmlReader) LexicalHandler() *QXmlLexicalHandler {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LexicalHandler"}).(*QXmlLexicalHandler)
}

func (ptr *QXmlReader) ConnectProperty(f func(name string, ok *bool) unsafe.Pointer) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlReader) DisconnectProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectProperty"})
}

func (ptr *QXmlReader) Property(name string, ok *bool) unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Property", name, ok}).(unsafe.Pointer)
}

func (ptr *QXmlReader) ConnectSetContentHandler(f func(handler *QXmlContentHandler)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetContentHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlReader) DisconnectSetContentHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetContentHandler"})
}

func (ptr *QXmlReader) SetContentHandler(handler QXmlContentHandler_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetContentHandler", handler})
}

func (ptr *QXmlReader) ConnectSetDTDHandler(f func(handler *QXmlDTDHandler)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetDTDHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlReader) DisconnectSetDTDHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetDTDHandler"})
}

func (ptr *QXmlReader) SetDTDHandler(handler QXmlDTDHandler_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDTDHandler", handler})
}

func (ptr *QXmlReader) ConnectSetDeclHandler(f func(handler *QXmlDeclHandler)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetDeclHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlReader) DisconnectSetDeclHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetDeclHandler"})
}

func (ptr *QXmlReader) SetDeclHandler(handler QXmlDeclHandler_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDeclHandler", handler})
}

func (ptr *QXmlReader) ConnectSetEntityResolver(f func(handler *QXmlEntityResolver)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetEntityResolver", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlReader) DisconnectSetEntityResolver() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetEntityResolver"})
}

func (ptr *QXmlReader) SetEntityResolver(handler QXmlEntityResolver_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEntityResolver", handler})
}

func (ptr *QXmlReader) ConnectSetErrorHandler(f func(handler *QXmlErrorHandler)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetErrorHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlReader) DisconnectSetErrorHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetErrorHandler"})
}

func (ptr *QXmlReader) SetErrorHandler(handler QXmlErrorHandler_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetErrorHandler", handler})
}

func (ptr *QXmlReader) ConnectSetFeature(f func(name string, value bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetFeature", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlReader) DisconnectSetFeature() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetFeature"})
}

func (ptr *QXmlReader) SetFeature(name string, value bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFeature", name, value})
}

func (ptr *QXmlReader) ConnectSetLexicalHandler(f func(handler *QXmlLexicalHandler)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetLexicalHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlReader) DisconnectSetLexicalHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetLexicalHandler"})
}

func (ptr *QXmlReader) SetLexicalHandler(handler QXmlLexicalHandler_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLexicalHandler", handler})
}

func (ptr *QXmlReader) ConnectSetProperty(f func(name string, value unsafe.Pointer)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlReader) DisconnectSetProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetProperty"})
}

func (ptr *QXmlReader) SetProperty(name string, value unsafe.Pointer) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetProperty", name, value})
}

func (ptr *QXmlReader) ConnectDestroyQXmlReader(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQXmlReader", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlReader) DisconnectDestroyQXmlReader() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQXmlReader"})
}

func (ptr *QXmlReader) DestroyQXmlReader() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlReader"})
}

func (ptr *QXmlReader) DestroyQXmlReaderDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlReaderDefault"})
}

type QXmlSimpleReader struct {
	QXmlReader
}

type QXmlSimpleReader_ITF interface {
	QXmlReader_ITF
	QXmlSimpleReader_PTR() *QXmlSimpleReader
}

func (ptr *QXmlSimpleReader) QXmlSimpleReader_PTR() *QXmlSimpleReader {
	return ptr
}

func (ptr *QXmlSimpleReader) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlReader_PTR().Pointer()
	}
	return nil
}

func (ptr *QXmlSimpleReader) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QXmlReader_PTR().SetPointer(p)
	}
}

func PointerFromQXmlSimpleReader(ptr QXmlSimpleReader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlSimpleReader_PTR().Pointer()
	}
	return nil
}

func (n *QXmlSimpleReader) InitFromInternal(ptr uintptr, name string) {
	n.QXmlReader_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QXmlSimpleReader) ClassNameInternalF() string {
	return n.QXmlReader_PTR().ClassNameInternalF()
}

func NewQXmlSimpleReaderFromPointer(ptr unsafe.Pointer) (n *QXmlSimpleReader) {
	n = new(QXmlSimpleReader)
	n.InitFromInternal(uintptr(ptr), "xml.QXmlSimpleReader")
	return
}
func (ptr *QXmlSimpleReader) ConnectDTDHandler(f func() *QXmlDTDHandler) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDTDHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSimpleReader) DisconnectDTDHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDTDHandler"})
}

func (ptr *QXmlSimpleReader) DTDHandler() *QXmlDTDHandler {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DTDHandler"}).(*QXmlDTDHandler)
}

func (ptr *QXmlSimpleReader) DTDHandlerDefault() *QXmlDTDHandler {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DTDHandlerDefault"}).(*QXmlDTDHandler)
}

func NewQXmlSimpleReader() *QXmlSimpleReader {

	return internal.CallLocalFunction([]interface{}{"", "", "xml.NewQXmlSimpleReader", ""}).(*QXmlSimpleReader)
}

func (ptr *QXmlSimpleReader) ConnectContentHandler(f func() *QXmlContentHandler) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectContentHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSimpleReader) DisconnectContentHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectContentHandler"})
}

func (ptr *QXmlSimpleReader) ContentHandler() *QXmlContentHandler {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContentHandler"}).(*QXmlContentHandler)
}

func (ptr *QXmlSimpleReader) ContentHandlerDefault() *QXmlContentHandler {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContentHandlerDefault"}).(*QXmlContentHandler)
}

func (ptr *QXmlSimpleReader) ConnectDeclHandler(f func() *QXmlDeclHandler) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDeclHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSimpleReader) DisconnectDeclHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDeclHandler"})
}

func (ptr *QXmlSimpleReader) DeclHandler() *QXmlDeclHandler {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeclHandler"}).(*QXmlDeclHandler)
}

func (ptr *QXmlSimpleReader) DeclHandlerDefault() *QXmlDeclHandler {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeclHandlerDefault"}).(*QXmlDeclHandler)
}

func (ptr *QXmlSimpleReader) ConnectEntityResolver(f func() *QXmlEntityResolver) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectEntityResolver", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSimpleReader) DisconnectEntityResolver() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectEntityResolver"})
}

func (ptr *QXmlSimpleReader) EntityResolver() *QXmlEntityResolver {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EntityResolver"}).(*QXmlEntityResolver)
}

func (ptr *QXmlSimpleReader) EntityResolverDefault() *QXmlEntityResolver {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EntityResolverDefault"}).(*QXmlEntityResolver)
}

func (ptr *QXmlSimpleReader) ConnectErrorHandler(f func() *QXmlErrorHandler) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectErrorHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSimpleReader) DisconnectErrorHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectErrorHandler"})
}

func (ptr *QXmlSimpleReader) ErrorHandler() *QXmlErrorHandler {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorHandler"}).(*QXmlErrorHandler)
}

func (ptr *QXmlSimpleReader) ErrorHandlerDefault() *QXmlErrorHandler {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorHandlerDefault"}).(*QXmlErrorHandler)
}

func (ptr *QXmlSimpleReader) ConnectFeature(f func(name string, ok *bool) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFeature", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSimpleReader) DisconnectFeature() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFeature"})
}

func (ptr *QXmlSimpleReader) Feature(name string, ok *bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Feature", name, ok}).(bool)
}

func (ptr *QXmlSimpleReader) FeatureDefault(name string, ok *bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FeatureDefault", name, ok}).(bool)
}

func (ptr *QXmlSimpleReader) ConnectHasFeature(f func(name string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHasFeature", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSimpleReader) DisconnectHasFeature() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHasFeature"})
}

func (ptr *QXmlSimpleReader) HasFeature(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasFeature", name}).(bool)
}

func (ptr *QXmlSimpleReader) HasFeatureDefault(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasFeatureDefault", name}).(bool)
}

func (ptr *QXmlSimpleReader) ConnectHasProperty(f func(name string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHasProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSimpleReader) DisconnectHasProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHasProperty"})
}

func (ptr *QXmlSimpleReader) HasProperty(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasProperty", name}).(bool)
}

func (ptr *QXmlSimpleReader) HasPropertyDefault(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasPropertyDefault", name}).(bool)
}

func (ptr *QXmlSimpleReader) ConnectLexicalHandler(f func() *QXmlLexicalHandler) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLexicalHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSimpleReader) DisconnectLexicalHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLexicalHandler"})
}

func (ptr *QXmlSimpleReader) LexicalHandler() *QXmlLexicalHandler {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LexicalHandler"}).(*QXmlLexicalHandler)
}

func (ptr *QXmlSimpleReader) LexicalHandlerDefault() *QXmlLexicalHandler {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LexicalHandlerDefault"}).(*QXmlLexicalHandler)
}

func (ptr *QXmlSimpleReader) ConnectParse(f func(input *QXmlInputSource) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectParse", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSimpleReader) DisconnectParse() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectParse"})
}

func (ptr *QXmlSimpleReader) Parse(input QXmlInputSource_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Parse", input}).(bool)
}

func (ptr *QXmlSimpleReader) ParseDefault(input QXmlInputSource_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ParseDefault", input}).(bool)
}

func (ptr *QXmlSimpleReader) ConnectParse2(f func(input *QXmlInputSource) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectParse2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSimpleReader) DisconnectParse2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectParse2"})
}

func (ptr *QXmlSimpleReader) Parse2(input QXmlInputSource_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Parse2", input}).(bool)
}

func (ptr *QXmlSimpleReader) Parse2Default(input QXmlInputSource_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Parse2Default", input}).(bool)
}

func (ptr *QXmlSimpleReader) ConnectParse3(f func(input *QXmlInputSource, incremental bool) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectParse3", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSimpleReader) DisconnectParse3() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectParse3"})
}

func (ptr *QXmlSimpleReader) Parse3(input QXmlInputSource_ITF, incremental bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Parse3", input, incremental}).(bool)
}

func (ptr *QXmlSimpleReader) Parse3Default(input QXmlInputSource_ITF, incremental bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Parse3Default", input, incremental}).(bool)
}

func (ptr *QXmlSimpleReader) ConnectParseContinue(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectParseContinue", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSimpleReader) DisconnectParseContinue() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectParseContinue"})
}

func (ptr *QXmlSimpleReader) ParseContinue() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ParseContinue"}).(bool)
}

func (ptr *QXmlSimpleReader) ParseContinueDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ParseContinueDefault"}).(bool)
}

func (ptr *QXmlSimpleReader) ConnectProperty(f func(name string, ok *bool) unsafe.Pointer) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSimpleReader) DisconnectProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectProperty"})
}

func (ptr *QXmlSimpleReader) Property(name string, ok *bool) unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Property", name, ok}).(unsafe.Pointer)
}

func (ptr *QXmlSimpleReader) PropertyDefault(name string, ok *bool) unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PropertyDefault", name, ok}).(unsafe.Pointer)
}

func (ptr *QXmlSimpleReader) ConnectSetContentHandler(f func(handler *QXmlContentHandler)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetContentHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSimpleReader) DisconnectSetContentHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetContentHandler"})
}

func (ptr *QXmlSimpleReader) SetContentHandler(handler QXmlContentHandler_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetContentHandler", handler})
}

func (ptr *QXmlSimpleReader) SetContentHandlerDefault(handler QXmlContentHandler_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetContentHandlerDefault", handler})
}

func (ptr *QXmlSimpleReader) ConnectSetDTDHandler(f func(handler *QXmlDTDHandler)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetDTDHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSimpleReader) DisconnectSetDTDHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetDTDHandler"})
}

func (ptr *QXmlSimpleReader) SetDTDHandler(handler QXmlDTDHandler_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDTDHandler", handler})
}

func (ptr *QXmlSimpleReader) SetDTDHandlerDefault(handler QXmlDTDHandler_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDTDHandlerDefault", handler})
}

func (ptr *QXmlSimpleReader) ConnectSetDeclHandler(f func(handler *QXmlDeclHandler)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetDeclHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSimpleReader) DisconnectSetDeclHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetDeclHandler"})
}

func (ptr *QXmlSimpleReader) SetDeclHandler(handler QXmlDeclHandler_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDeclHandler", handler})
}

func (ptr *QXmlSimpleReader) SetDeclHandlerDefault(handler QXmlDeclHandler_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDeclHandlerDefault", handler})
}

func (ptr *QXmlSimpleReader) ConnectSetEntityResolver(f func(handler *QXmlEntityResolver)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetEntityResolver", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSimpleReader) DisconnectSetEntityResolver() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetEntityResolver"})
}

func (ptr *QXmlSimpleReader) SetEntityResolver(handler QXmlEntityResolver_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEntityResolver", handler})
}

func (ptr *QXmlSimpleReader) SetEntityResolverDefault(handler QXmlEntityResolver_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEntityResolverDefault", handler})
}

func (ptr *QXmlSimpleReader) ConnectSetErrorHandler(f func(handler *QXmlErrorHandler)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetErrorHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSimpleReader) DisconnectSetErrorHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetErrorHandler"})
}

func (ptr *QXmlSimpleReader) SetErrorHandler(handler QXmlErrorHandler_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetErrorHandler", handler})
}

func (ptr *QXmlSimpleReader) SetErrorHandlerDefault(handler QXmlErrorHandler_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetErrorHandlerDefault", handler})
}

func (ptr *QXmlSimpleReader) ConnectSetFeature(f func(name string, enable bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetFeature", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSimpleReader) DisconnectSetFeature() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetFeature"})
}

func (ptr *QXmlSimpleReader) SetFeature(name string, enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFeature", name, enable})
}

func (ptr *QXmlSimpleReader) SetFeatureDefault(name string, enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFeatureDefault", name, enable})
}

func (ptr *QXmlSimpleReader) ConnectSetLexicalHandler(f func(handler *QXmlLexicalHandler)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetLexicalHandler", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSimpleReader) DisconnectSetLexicalHandler() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetLexicalHandler"})
}

func (ptr *QXmlSimpleReader) SetLexicalHandler(handler QXmlLexicalHandler_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLexicalHandler", handler})
}

func (ptr *QXmlSimpleReader) SetLexicalHandlerDefault(handler QXmlLexicalHandler_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLexicalHandlerDefault", handler})
}

func (ptr *QXmlSimpleReader) ConnectSetProperty(f func(name string, value unsafe.Pointer)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSimpleReader) DisconnectSetProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetProperty"})
}

func (ptr *QXmlSimpleReader) SetProperty(name string, value unsafe.Pointer) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetProperty", name, value})
}

func (ptr *QXmlSimpleReader) SetPropertyDefault(name string, value unsafe.Pointer) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPropertyDefault", name, value})
}

func (ptr *QXmlSimpleReader) ConnectDestroyQXmlSimpleReader(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQXmlSimpleReader", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QXmlSimpleReader) DisconnectDestroyQXmlSimpleReader() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQXmlSimpleReader"})
}

func (ptr *QXmlSimpleReader) DestroyQXmlSimpleReader() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlSimpleReader"})
}

func (ptr *QXmlSimpleReader) DestroyQXmlSimpleReaderDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQXmlSimpleReaderDefault"})
}

func init() {
	internal.ConstructorTable["xml.QDomAttr"] = NewQDomAttrFromPointer
	internal.ConstructorTable["xml.QDomCDATASection"] = NewQDomCDATASectionFromPointer
	internal.ConstructorTable["xml.QDomCharacterData"] = NewQDomCharacterDataFromPointer
	internal.ConstructorTable["xml.QDomComment"] = NewQDomCommentFromPointer
	internal.ConstructorTable["xml.QDomDocument"] = NewQDomDocumentFromPointer
	internal.ConstructorTable["xml.QDomDocumentFragment"] = NewQDomDocumentFragmentFromPointer
	internal.ConstructorTable["xml.QDomDocumentType"] = NewQDomDocumentTypeFromPointer
	internal.ConstructorTable["xml.QDomElement"] = NewQDomElementFromPointer
	internal.ConstructorTable["xml.QDomEntity"] = NewQDomEntityFromPointer
	internal.ConstructorTable["xml.QDomEntityReference"] = NewQDomEntityReferenceFromPointer
	internal.ConstructorTable["xml.QDomImplementation"] = NewQDomImplementationFromPointer
	internal.ConstructorTable["xml.QDomNamedNodeMap"] = NewQDomNamedNodeMapFromPointer
	internal.ConstructorTable["xml.QDomNode"] = NewQDomNodeFromPointer
	internal.ConstructorTable["xml.QDomNodeList"] = NewQDomNodeListFromPointer
	internal.ConstructorTable["xml.QDomNotation"] = NewQDomNotationFromPointer
	internal.ConstructorTable["xml.QDomProcessingInstruction"] = NewQDomProcessingInstructionFromPointer
	internal.ConstructorTable["xml.QDomText"] = NewQDomTextFromPointer
	internal.ConstructorTable["xml.QXmlAttributes"] = NewQXmlAttributesFromPointer
	internal.ConstructorTable["xml.QXmlContentHandler"] = NewQXmlContentHandlerFromPointer
	internal.ConstructorTable["xml.QXmlDTDHandler"] = NewQXmlDTDHandlerFromPointer
	internal.ConstructorTable["xml.QXmlDeclHandler"] = NewQXmlDeclHandlerFromPointer
	internal.ConstructorTable["xml.QXmlDefaultHandler"] = NewQXmlDefaultHandlerFromPointer
	internal.ConstructorTable["xml.QXmlEntityResolver"] = NewQXmlEntityResolverFromPointer
	internal.ConstructorTable["xml.QXmlErrorHandler"] = NewQXmlErrorHandlerFromPointer
	internal.ConstructorTable["xml.QXmlInputSource"] = NewQXmlInputSourceFromPointer
	internal.ConstructorTable["xml.QXmlLexicalHandler"] = NewQXmlLexicalHandlerFromPointer
	internal.ConstructorTable["xml.QXmlLocator"] = NewQXmlLocatorFromPointer
	internal.ConstructorTable["xml.QXmlNamespaceSupport"] = NewQXmlNamespaceSupportFromPointer
	internal.ConstructorTable["xml.QXmlParseException"] = NewQXmlParseExceptionFromPointer
	internal.ConstructorTable["xml.QXmlReader"] = NewQXmlReaderFromPointer
	internal.ConstructorTable["xml.QXmlSimpleReader"] = NewQXmlSimpleReaderFromPointer
}
