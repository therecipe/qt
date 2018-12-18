// +build !minimal

package xml

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "xml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtXml_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtXml_PackedString) []byte {
	if int(s.len) == -1 {
		return []byte(C.GoString(s.data))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
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

func NewQDomAttrFromPointer(ptr unsafe.Pointer) (n *QDomAttr) {
	n = new(QDomAttr)
	n.SetPointer(ptr)
	return
}

func (ptr *QDomAttr) DestroyQDomAttr() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQDomAttr() *QDomAttr {
	tmpValue := NewQDomAttrFromPointer(C.QDomAttr_NewQDomAttr())
	runtime.SetFinalizer(tmpValue, (*QDomAttr).DestroyQDomAttr)
	return tmpValue
}

func NewQDomAttr2(x QDomAttr_ITF) *QDomAttr {
	tmpValue := NewQDomAttrFromPointer(C.QDomAttr_NewQDomAttr2(PointerFromQDomAttr(x)))
	runtime.SetFinalizer(tmpValue, (*QDomAttr).DestroyQDomAttr)
	return tmpValue
}

func (ptr *QDomAttr) SetValue(v string) {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		C.QDomAttr_SetValue(ptr.Pointer(), C.struct_QtXml_PackedString{data: vC, len: C.longlong(len(v))})
	}
}

func (ptr *QDomAttr) OwnerElement() *QDomElement {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomElementFromPointer(C.QDomAttr_OwnerElement(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDomElement).DestroyQDomElement)
		return tmpValue
	}
	return nil
}

func (ptr *QDomAttr) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDomAttr_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomAttr) Value() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDomAttr_Value(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomAttr) Specified() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDomAttr_Specified(ptr.Pointer())) != 0
	}
	return false
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

func NewQDomCDATASectionFromPointer(ptr unsafe.Pointer) (n *QDomCDATASection) {
	n = new(QDomCDATASection)
	n.SetPointer(ptr)
	return
}

func (ptr *QDomCDATASection) DestroyQDomCDATASection() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQDomCDATASection() *QDomCDATASection {
	tmpValue := NewQDomCDATASectionFromPointer(C.QDomCDATASection_NewQDomCDATASection())
	runtime.SetFinalizer(tmpValue, (*QDomCDATASection).DestroyQDomCDATASection)
	return tmpValue
}

func NewQDomCDATASection2(x QDomCDATASection_ITF) *QDomCDATASection {
	tmpValue := NewQDomCDATASectionFromPointer(C.QDomCDATASection_NewQDomCDATASection2(PointerFromQDomCDATASection(x)))
	runtime.SetFinalizer(tmpValue, (*QDomCDATASection).DestroyQDomCDATASection)
	return tmpValue
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

func NewQDomCharacterDataFromPointer(ptr unsafe.Pointer) (n *QDomCharacterData) {
	n = new(QDomCharacterData)
	n.SetPointer(ptr)
	return
}

func (ptr *QDomCharacterData) DestroyQDomCharacterData() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQDomCharacterData() *QDomCharacterData {
	tmpValue := NewQDomCharacterDataFromPointer(C.QDomCharacterData_NewQDomCharacterData())
	runtime.SetFinalizer(tmpValue, (*QDomCharacterData).DestroyQDomCharacterData)
	return tmpValue
}

func NewQDomCharacterData2(x QDomCharacterData_ITF) *QDomCharacterData {
	tmpValue := NewQDomCharacterDataFromPointer(C.QDomCharacterData_NewQDomCharacterData2(PointerFromQDomCharacterData(x)))
	runtime.SetFinalizer(tmpValue, (*QDomCharacterData).DestroyQDomCharacterData)
	return tmpValue
}

func (ptr *QDomCharacterData) SubstringData(offset uint, count uint) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDomCharacterData_SubstringData(ptr.Pointer(), C.ulong(uint32(offset)), C.ulong(uint32(count))))
	}
	return ""
}

func (ptr *QDomCharacterData) AppendData(arg string) {
	if ptr.Pointer() != nil {
		var argC *C.char
		if arg != "" {
			argC = C.CString(arg)
			defer C.free(unsafe.Pointer(argC))
		}
		C.QDomCharacterData_AppendData(ptr.Pointer(), C.struct_QtXml_PackedString{data: argC, len: C.longlong(len(arg))})
	}
}

func (ptr *QDomCharacterData) DeleteData(offset uint, count uint) {
	if ptr.Pointer() != nil {
		C.QDomCharacterData_DeleteData(ptr.Pointer(), C.ulong(uint32(offset)), C.ulong(uint32(count)))
	}
}

func (ptr *QDomCharacterData) InsertData(offset uint, arg string) {
	if ptr.Pointer() != nil {
		var argC *C.char
		if arg != "" {
			argC = C.CString(arg)
			defer C.free(unsafe.Pointer(argC))
		}
		C.QDomCharacterData_InsertData(ptr.Pointer(), C.ulong(uint32(offset)), C.struct_QtXml_PackedString{data: argC, len: C.longlong(len(arg))})
	}
}

func (ptr *QDomCharacterData) ReplaceData(offset uint, count uint, arg string) {
	if ptr.Pointer() != nil {
		var argC *C.char
		if arg != "" {
			argC = C.CString(arg)
			defer C.free(unsafe.Pointer(argC))
		}
		C.QDomCharacterData_ReplaceData(ptr.Pointer(), C.ulong(uint32(offset)), C.ulong(uint32(count)), C.struct_QtXml_PackedString{data: argC, len: C.longlong(len(arg))})
	}
}

func (ptr *QDomCharacterData) SetData(v string) {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		C.QDomCharacterData_SetData(ptr.Pointer(), C.struct_QtXml_PackedString{data: vC, len: C.longlong(len(v))})
	}
}

func (ptr *QDomCharacterData) Data() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDomCharacterData_Data(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomCharacterData) Length() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDomCharacterData_Length(ptr.Pointer())))
	}
	return 0
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

func NewQDomCommentFromPointer(ptr unsafe.Pointer) (n *QDomComment) {
	n = new(QDomComment)
	n.SetPointer(ptr)
	return
}

func (ptr *QDomComment) DestroyQDomComment() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQDomComment() *QDomComment {
	tmpValue := NewQDomCommentFromPointer(C.QDomComment_NewQDomComment())
	runtime.SetFinalizer(tmpValue, (*QDomComment).DestroyQDomComment)
	return tmpValue
}

func NewQDomComment2(x QDomComment_ITF) *QDomComment {
	tmpValue := NewQDomCommentFromPointer(C.QDomComment_NewQDomComment2(PointerFromQDomComment(x)))
	runtime.SetFinalizer(tmpValue, (*QDomComment).DestroyQDomComment)
	return tmpValue
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

func NewQDomDocumentFromPointer(ptr unsafe.Pointer) (n *QDomDocument) {
	n = new(QDomDocument)
	n.SetPointer(ptr)
	return
}
func (ptr *QDomDocument) CreateAttribute(name string) *QDomAttr {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		tmpValue := NewQDomAttrFromPointer(C.QDomDocument_CreateAttribute(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}))
		runtime.SetFinalizer(tmpValue, (*QDomAttr).DestroyQDomAttr)
		return tmpValue
	}
	return nil
}

func (ptr *QDomDocument) CreateAttributeNS(nsURI string, qName string) *QDomAttr {
	if ptr.Pointer() != nil {
		var nsURIC *C.char
		if nsURI != "" {
			nsURIC = C.CString(nsURI)
			defer C.free(unsafe.Pointer(nsURIC))
		}
		var qNameC *C.char
		if qName != "" {
			qNameC = C.CString(qName)
			defer C.free(unsafe.Pointer(qNameC))
		}
		tmpValue := NewQDomAttrFromPointer(C.QDomDocument_CreateAttributeNS(ptr.Pointer(), C.struct_QtXml_PackedString{data: nsURIC, len: C.longlong(len(nsURI))}, C.struct_QtXml_PackedString{data: qNameC, len: C.longlong(len(qName))}))
		runtime.SetFinalizer(tmpValue, (*QDomAttr).DestroyQDomAttr)
		return tmpValue
	}
	return nil
}

func (ptr *QDomDocument) CreateCDATASection(value string) *QDomCDATASection {
	if ptr.Pointer() != nil {
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		tmpValue := NewQDomCDATASectionFromPointer(C.QDomDocument_CreateCDATASection(ptr.Pointer(), C.struct_QtXml_PackedString{data: valueC, len: C.longlong(len(value))}))
		runtime.SetFinalizer(tmpValue, (*QDomCDATASection).DestroyQDomCDATASection)
		return tmpValue
	}
	return nil
}

func (ptr *QDomDocument) CreateComment(value string) *QDomComment {
	if ptr.Pointer() != nil {
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		tmpValue := NewQDomCommentFromPointer(C.QDomDocument_CreateComment(ptr.Pointer(), C.struct_QtXml_PackedString{data: valueC, len: C.longlong(len(value))}))
		runtime.SetFinalizer(tmpValue, (*QDomComment).DestroyQDomComment)
		return tmpValue
	}
	return nil
}

func NewQDomDocument() *QDomDocument {
	tmpValue := NewQDomDocumentFromPointer(C.QDomDocument_NewQDomDocument())
	runtime.SetFinalizer(tmpValue, (*QDomDocument).DestroyQDomDocument)
	return tmpValue
}

func NewQDomDocument4(x QDomDocument_ITF) *QDomDocument {
	tmpValue := NewQDomDocumentFromPointer(C.QDomDocument_NewQDomDocument4(PointerFromQDomDocument(x)))
	runtime.SetFinalizer(tmpValue, (*QDomDocument).DestroyQDomDocument)
	return tmpValue
}

func NewQDomDocument3(doctype QDomDocumentType_ITF) *QDomDocument {
	tmpValue := NewQDomDocumentFromPointer(C.QDomDocument_NewQDomDocument3(PointerFromQDomDocumentType(doctype)))
	runtime.SetFinalizer(tmpValue, (*QDomDocument).DestroyQDomDocument)
	return tmpValue
}

func NewQDomDocument2(name string) *QDomDocument {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQDomDocumentFromPointer(C.QDomDocument_NewQDomDocument2(C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}))
	runtime.SetFinalizer(tmpValue, (*QDomDocument).DestroyQDomDocument)
	return tmpValue
}

func (ptr *QDomDocument) CreateDocumentFragment() *QDomDocumentFragment {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomDocumentFragmentFromPointer(C.QDomDocument_CreateDocumentFragment(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDomDocumentFragment).DestroyQDomDocumentFragment)
		return tmpValue
	}
	return nil
}

func (ptr *QDomDocument) CreateElement(tagName string) *QDomElement {
	if ptr.Pointer() != nil {
		var tagNameC *C.char
		if tagName != "" {
			tagNameC = C.CString(tagName)
			defer C.free(unsafe.Pointer(tagNameC))
		}
		tmpValue := NewQDomElementFromPointer(C.QDomDocument_CreateElement(ptr.Pointer(), C.struct_QtXml_PackedString{data: tagNameC, len: C.longlong(len(tagName))}))
		runtime.SetFinalizer(tmpValue, (*QDomElement).DestroyQDomElement)
		return tmpValue
	}
	return nil
}

func (ptr *QDomDocument) CreateElementNS(nsURI string, qName string) *QDomElement {
	if ptr.Pointer() != nil {
		var nsURIC *C.char
		if nsURI != "" {
			nsURIC = C.CString(nsURI)
			defer C.free(unsafe.Pointer(nsURIC))
		}
		var qNameC *C.char
		if qName != "" {
			qNameC = C.CString(qName)
			defer C.free(unsafe.Pointer(qNameC))
		}
		tmpValue := NewQDomElementFromPointer(C.QDomDocument_CreateElementNS(ptr.Pointer(), C.struct_QtXml_PackedString{data: nsURIC, len: C.longlong(len(nsURI))}, C.struct_QtXml_PackedString{data: qNameC, len: C.longlong(len(qName))}))
		runtime.SetFinalizer(tmpValue, (*QDomElement).DestroyQDomElement)
		return tmpValue
	}
	return nil
}

func (ptr *QDomDocument) ElementById(elementId string) *QDomElement {
	if ptr.Pointer() != nil {
		var elementIdC *C.char
		if elementId != "" {
			elementIdC = C.CString(elementId)
			defer C.free(unsafe.Pointer(elementIdC))
		}
		tmpValue := NewQDomElementFromPointer(C.QDomDocument_ElementById(ptr.Pointer(), C.struct_QtXml_PackedString{data: elementIdC, len: C.longlong(len(elementId))}))
		runtime.SetFinalizer(tmpValue, (*QDomElement).DestroyQDomElement)
		return tmpValue
	}
	return nil
}

func (ptr *QDomDocument) CreateEntityReference(name string) *QDomEntityReference {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		tmpValue := NewQDomEntityReferenceFromPointer(C.QDomDocument_CreateEntityReference(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}))
		runtime.SetFinalizer(tmpValue, (*QDomEntityReference).DestroyQDomEntityReference)
		return tmpValue
	}
	return nil
}

func (ptr *QDomDocument) ImportNode(importedNode QDomNode_ITF, deep bool) *QDomNode {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomNodeFromPointer(C.QDomDocument_ImportNode(ptr.Pointer(), PointerFromQDomNode(importedNode), C.char(int8(qt.GoBoolToInt(deep)))))
		runtime.SetFinalizer(tmpValue, (*QDomNode).DestroyQDomNode)
		return tmpValue
	}
	return nil
}

func (ptr *QDomDocument) ElementsByTagNameNS(nsURI string, localName string) *QDomNodeList {
	if ptr.Pointer() != nil {
		var nsURIC *C.char
		if nsURI != "" {
			nsURIC = C.CString(nsURI)
			defer C.free(unsafe.Pointer(nsURIC))
		}
		var localNameC *C.char
		if localName != "" {
			localNameC = C.CString(localName)
			defer C.free(unsafe.Pointer(localNameC))
		}
		tmpValue := NewQDomNodeListFromPointer(C.QDomDocument_ElementsByTagNameNS(ptr.Pointer(), C.struct_QtXml_PackedString{data: nsURIC, len: C.longlong(len(nsURI))}, C.struct_QtXml_PackedString{data: localNameC, len: C.longlong(len(localName))}))
		runtime.SetFinalizer(tmpValue, (*QDomNodeList).DestroyQDomNodeList)
		return tmpValue
	}
	return nil
}

func (ptr *QDomDocument) CreateProcessingInstruction(target string, data string) *QDomProcessingInstruction {
	if ptr.Pointer() != nil {
		var targetC *C.char
		if target != "" {
			targetC = C.CString(target)
			defer C.free(unsafe.Pointer(targetC))
		}
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		tmpValue := NewQDomProcessingInstructionFromPointer(C.QDomDocument_CreateProcessingInstruction(ptr.Pointer(), C.struct_QtXml_PackedString{data: targetC, len: C.longlong(len(target))}, C.struct_QtXml_PackedString{data: dataC, len: C.longlong(len(data))}))
		runtime.SetFinalizer(tmpValue, (*QDomProcessingInstruction).DestroyQDomProcessingInstruction)
		return tmpValue
	}
	return nil
}

func (ptr *QDomDocument) CreateTextNode(value string) *QDomText {
	if ptr.Pointer() != nil {
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		tmpValue := NewQDomTextFromPointer(C.QDomDocument_CreateTextNode(ptr.Pointer(), C.struct_QtXml_PackedString{data: valueC, len: C.longlong(len(value))}))
		runtime.SetFinalizer(tmpValue, (*QDomText).DestroyQDomText)
		return tmpValue
	}
	return nil
}

func (ptr *QDomDocument) SetContent7(dev core.QIODevice_ITF, errorMsg string, errorLine int, errorColumn int) bool {
	if ptr.Pointer() != nil {
		var errorMsgC *C.char
		if errorMsg != "" {
			errorMsgC = C.CString(errorMsg)
			defer C.free(unsafe.Pointer(errorMsgC))
		}
		return int8(C.QDomDocument_SetContent7(ptr.Pointer(), core.PointerFromQIODevice(dev), C.struct_QtXml_PackedString{data: errorMsgC, len: C.longlong(len(errorMsg))}, C.int(int32(errorLine)), C.int(int32(errorColumn)))) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent3(dev core.QIODevice_ITF, namespaceProcessing bool, errorMsg string, errorLine int, errorColumn int) bool {
	if ptr.Pointer() != nil {
		var errorMsgC *C.char
		if errorMsg != "" {
			errorMsgC = C.CString(errorMsg)
			defer C.free(unsafe.Pointer(errorMsgC))
		}
		return int8(C.QDomDocument_SetContent3(ptr.Pointer(), core.PointerFromQIODevice(dev), C.char(int8(qt.GoBoolToInt(namespaceProcessing))), C.struct_QtXml_PackedString{data: errorMsgC, len: C.longlong(len(errorMsg))}, C.int(int32(errorLine)), C.int(int32(errorColumn)))) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent8(source QXmlInputSource_ITF, reader QXmlReader_ITF, errorMsg string, errorLine int, errorColumn int) bool {
	if ptr.Pointer() != nil {
		var errorMsgC *C.char
		if errorMsg != "" {
			errorMsgC = C.CString(errorMsg)
			defer C.free(unsafe.Pointer(errorMsgC))
		}
		return int8(C.QDomDocument_SetContent8(ptr.Pointer(), PointerFromQXmlInputSource(source), PointerFromQXmlReader(reader), C.struct_QtXml_PackedString{data: errorMsgC, len: C.longlong(len(errorMsg))}, C.int(int32(errorLine)), C.int(int32(errorColumn)))) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent4(source QXmlInputSource_ITF, namespaceProcessing bool, errorMsg string, errorLine int, errorColumn int) bool {
	if ptr.Pointer() != nil {
		var errorMsgC *C.char
		if errorMsg != "" {
			errorMsgC = C.CString(errorMsg)
			defer C.free(unsafe.Pointer(errorMsgC))
		}
		return int8(C.QDomDocument_SetContent4(ptr.Pointer(), PointerFromQXmlInputSource(source), C.char(int8(qt.GoBoolToInt(namespaceProcessing))), C.struct_QtXml_PackedString{data: errorMsgC, len: C.longlong(len(errorMsg))}, C.int(int32(errorLine)), C.int(int32(errorColumn)))) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent5(buffer core.QByteArray_ITF, errorMsg string, errorLine int, errorColumn int) bool {
	if ptr.Pointer() != nil {
		var errorMsgC *C.char
		if errorMsg != "" {
			errorMsgC = C.CString(errorMsg)
			defer C.free(unsafe.Pointer(errorMsgC))
		}
		return int8(C.QDomDocument_SetContent5(ptr.Pointer(), core.PointerFromQByteArray(buffer), C.struct_QtXml_PackedString{data: errorMsgC, len: C.longlong(len(errorMsg))}, C.int(int32(errorLine)), C.int(int32(errorColumn)))) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent(data core.QByteArray_ITF, namespaceProcessing bool, errorMsg string, errorLine int, errorColumn int) bool {
	if ptr.Pointer() != nil {
		var errorMsgC *C.char
		if errorMsg != "" {
			errorMsgC = C.CString(errorMsg)
			defer C.free(unsafe.Pointer(errorMsgC))
		}
		return int8(C.QDomDocument_SetContent(ptr.Pointer(), core.PointerFromQByteArray(data), C.char(int8(qt.GoBoolToInt(namespaceProcessing))), C.struct_QtXml_PackedString{data: errorMsgC, len: C.longlong(len(errorMsg))}, C.int(int32(errorLine)), C.int(int32(errorColumn)))) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent6(text string, errorMsg string, errorLine int, errorColumn int) bool {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		var errorMsgC *C.char
		if errorMsg != "" {
			errorMsgC = C.CString(errorMsg)
			defer C.free(unsafe.Pointer(errorMsgC))
		}
		return int8(C.QDomDocument_SetContent6(ptr.Pointer(), C.struct_QtXml_PackedString{data: textC, len: C.longlong(len(text))}, C.struct_QtXml_PackedString{data: errorMsgC, len: C.longlong(len(errorMsg))}, C.int(int32(errorLine)), C.int(int32(errorColumn)))) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent2(text string, namespaceProcessing bool, errorMsg string, errorLine int, errorColumn int) bool {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		var errorMsgC *C.char
		if errorMsg != "" {
			errorMsgC = C.CString(errorMsg)
			defer C.free(unsafe.Pointer(errorMsgC))
		}
		return int8(C.QDomDocument_SetContent2(ptr.Pointer(), C.struct_QtXml_PackedString{data: textC, len: C.longlong(len(text))}, C.char(int8(qt.GoBoolToInt(namespaceProcessing))), C.struct_QtXml_PackedString{data: errorMsgC, len: C.longlong(len(errorMsg))}, C.int(int32(errorLine)), C.int(int32(errorColumn)))) != 0
	}
	return false
}

func (ptr *QDomDocument) DestroyQDomDocument() {
	if ptr.Pointer() != nil {
		C.QDomDocument_DestroyQDomDocument(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDomDocument) ToByteArray(indent int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QDomDocument_ToByteArray(ptr.Pointer(), C.int(int32(indent))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QDomDocument) Doctype() *QDomDocumentType {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomDocumentTypeFromPointer(C.QDomDocument_Doctype(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDomDocumentType).DestroyQDomDocumentType)
		return tmpValue
	}
	return nil
}

func (ptr *QDomDocument) DocumentElement() *QDomElement {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomElementFromPointer(C.QDomDocument_DocumentElement(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDomElement).DestroyQDomElement)
		return tmpValue
	}
	return nil
}

func (ptr *QDomDocument) Implementation() *QDomImplementation {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomImplementationFromPointer(C.QDomDocument_Implementation(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDomImplementation).DestroyQDomImplementation)
		return tmpValue
	}
	return nil
}

func (ptr *QDomDocument) ElementsByTagName(tagname string) *QDomNodeList {
	if ptr.Pointer() != nil {
		var tagnameC *C.char
		if tagname != "" {
			tagnameC = C.CString(tagname)
			defer C.free(unsafe.Pointer(tagnameC))
		}
		tmpValue := NewQDomNodeListFromPointer(C.QDomDocument_ElementsByTagName(ptr.Pointer(), C.struct_QtXml_PackedString{data: tagnameC, len: C.longlong(len(tagname))}))
		runtime.SetFinalizer(tmpValue, (*QDomNodeList).DestroyQDomNodeList)
		return tmpValue
	}
	return nil
}

func (ptr *QDomDocument) ToString(indent int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDomDocument_ToString(ptr.Pointer(), C.int(int32(indent))))
	}
	return ""
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

func NewQDomDocumentFragmentFromPointer(ptr unsafe.Pointer) (n *QDomDocumentFragment) {
	n = new(QDomDocumentFragment)
	n.SetPointer(ptr)
	return
}

func (ptr *QDomDocumentFragment) DestroyQDomDocumentFragment() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQDomDocumentFragment() *QDomDocumentFragment {
	tmpValue := NewQDomDocumentFragmentFromPointer(C.QDomDocumentFragment_NewQDomDocumentFragment())
	runtime.SetFinalizer(tmpValue, (*QDomDocumentFragment).DestroyQDomDocumentFragment)
	return tmpValue
}

func NewQDomDocumentFragment2(x QDomDocumentFragment_ITF) *QDomDocumentFragment {
	tmpValue := NewQDomDocumentFragmentFromPointer(C.QDomDocumentFragment_NewQDomDocumentFragment2(PointerFromQDomDocumentFragment(x)))
	runtime.SetFinalizer(tmpValue, (*QDomDocumentFragment).DestroyQDomDocumentFragment)
	return tmpValue
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

func NewQDomDocumentTypeFromPointer(ptr unsafe.Pointer) (n *QDomDocumentType) {
	n = new(QDomDocumentType)
	n.SetPointer(ptr)
	return
}

func (ptr *QDomDocumentType) DestroyQDomDocumentType() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQDomDocumentType() *QDomDocumentType {
	tmpValue := NewQDomDocumentTypeFromPointer(C.QDomDocumentType_NewQDomDocumentType())
	runtime.SetFinalizer(tmpValue, (*QDomDocumentType).DestroyQDomDocumentType)
	return tmpValue
}

func NewQDomDocumentType2(n QDomDocumentType_ITF) *QDomDocumentType {
	tmpValue := NewQDomDocumentTypeFromPointer(C.QDomDocumentType_NewQDomDocumentType2(PointerFromQDomDocumentType(n)))
	runtime.SetFinalizer(tmpValue, (*QDomDocumentType).DestroyQDomDocumentType)
	return tmpValue
}

func (ptr *QDomDocumentType) Entities() *QDomNamedNodeMap {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomNamedNodeMapFromPointer(C.QDomDocumentType_Entities(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDomNamedNodeMap).DestroyQDomNamedNodeMap)
		return tmpValue
	}
	return nil
}

func (ptr *QDomDocumentType) Notations() *QDomNamedNodeMap {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomNamedNodeMapFromPointer(C.QDomDocumentType_Notations(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDomNamedNodeMap).DestroyQDomNamedNodeMap)
		return tmpValue
	}
	return nil
}

func (ptr *QDomDocumentType) InternalSubset() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDomDocumentType_InternalSubset(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomDocumentType) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDomDocumentType_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomDocumentType) PublicId() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDomDocumentType_PublicId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomDocumentType) SystemId() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDomDocumentType_SystemId(ptr.Pointer()))
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

func NewQDomElementFromPointer(ptr unsafe.Pointer) (n *QDomElement) {
	n = new(QDomElement)
	n.SetPointer(ptr)
	return
}

func (ptr *QDomElement) DestroyQDomElement() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDomElement) AttributeNode(name string) *QDomAttr {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		tmpValue := NewQDomAttrFromPointer(C.QDomElement_AttributeNode(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}))
		runtime.SetFinalizer(tmpValue, (*QDomAttr).DestroyQDomAttr)
		return tmpValue
	}
	return nil
}

func (ptr *QDomElement) AttributeNodeNS(nsURI string, localName string) *QDomAttr {
	if ptr.Pointer() != nil {
		var nsURIC *C.char
		if nsURI != "" {
			nsURIC = C.CString(nsURI)
			defer C.free(unsafe.Pointer(nsURIC))
		}
		var localNameC *C.char
		if localName != "" {
			localNameC = C.CString(localName)
			defer C.free(unsafe.Pointer(localNameC))
		}
		tmpValue := NewQDomAttrFromPointer(C.QDomElement_AttributeNodeNS(ptr.Pointer(), C.struct_QtXml_PackedString{data: nsURIC, len: C.longlong(len(nsURI))}, C.struct_QtXml_PackedString{data: localNameC, len: C.longlong(len(localName))}))
		runtime.SetFinalizer(tmpValue, (*QDomAttr).DestroyQDomAttr)
		return tmpValue
	}
	return nil
}

func (ptr *QDomElement) RemoveAttributeNode(oldAttr QDomAttr_ITF) *QDomAttr {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomAttrFromPointer(C.QDomElement_RemoveAttributeNode(ptr.Pointer(), PointerFromQDomAttr(oldAttr)))
		runtime.SetFinalizer(tmpValue, (*QDomAttr).DestroyQDomAttr)
		return tmpValue
	}
	return nil
}

func (ptr *QDomElement) SetAttributeNode(newAttr QDomAttr_ITF) *QDomAttr {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomAttrFromPointer(C.QDomElement_SetAttributeNode(ptr.Pointer(), PointerFromQDomAttr(newAttr)))
		runtime.SetFinalizer(tmpValue, (*QDomAttr).DestroyQDomAttr)
		return tmpValue
	}
	return nil
}

func (ptr *QDomElement) SetAttributeNodeNS(newAttr QDomAttr_ITF) *QDomAttr {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomAttrFromPointer(C.QDomElement_SetAttributeNodeNS(ptr.Pointer(), PointerFromQDomAttr(newAttr)))
		runtime.SetFinalizer(tmpValue, (*QDomAttr).DestroyQDomAttr)
		return tmpValue
	}
	return nil
}

func NewQDomElement() *QDomElement {
	tmpValue := NewQDomElementFromPointer(C.QDomElement_NewQDomElement())
	runtime.SetFinalizer(tmpValue, (*QDomElement).DestroyQDomElement)
	return tmpValue
}

func NewQDomElement2(x QDomElement_ITF) *QDomElement {
	tmpValue := NewQDomElementFromPointer(C.QDomElement_NewQDomElement2(PointerFromQDomElement(x)))
	runtime.SetFinalizer(tmpValue, (*QDomElement).DestroyQDomElement)
	return tmpValue
}

func (ptr *QDomElement) RemoveAttribute(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QDomElement_RemoveAttribute(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QDomElement) RemoveAttributeNS(nsURI string, localName string) {
	if ptr.Pointer() != nil {
		var nsURIC *C.char
		if nsURI != "" {
			nsURIC = C.CString(nsURI)
			defer C.free(unsafe.Pointer(nsURIC))
		}
		var localNameC *C.char
		if localName != "" {
			localNameC = C.CString(localName)
			defer C.free(unsafe.Pointer(localNameC))
		}
		C.QDomElement_RemoveAttributeNS(ptr.Pointer(), C.struct_QtXml_PackedString{data: nsURIC, len: C.longlong(len(nsURI))}, C.struct_QtXml_PackedString{data: localNameC, len: C.longlong(len(localName))})
	}
}

func (ptr *QDomElement) SetAttribute(name string, value string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		C.QDomElement_SetAttribute(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtXml_PackedString{data: valueC, len: C.longlong(len(value))})
	}
}

func (ptr *QDomElement) SetAttribute7(name string, value float64) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QDomElement_SetAttribute7(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, C.double(value))
	}
}

func (ptr *QDomElement) SetAttribute6(name string, value float32) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QDomElement_SetAttribute6(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, C.float(value))
	}
}

func (ptr *QDomElement) SetAttribute4(name string, value int) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QDomElement_SetAttribute4(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, C.int(int32(value)))
	}
}

func (ptr *QDomElement) SetAttribute2(name string, value int64) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QDomElement_SetAttribute2(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, C.longlong(value))
	}
}

func (ptr *QDomElement) SetAttribute3(name string, value uint64) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QDomElement_SetAttribute3(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, C.ulonglong(value))
	}
}

func (ptr *QDomElement) SetAttribute5(name string, value uint) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QDomElement_SetAttribute5(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, C.uint(uint32(value)))
	}
}

func (ptr *QDomElement) SetAttributeNS(nsURI string, qName string, value string) {
	if ptr.Pointer() != nil {
		var nsURIC *C.char
		if nsURI != "" {
			nsURIC = C.CString(nsURI)
			defer C.free(unsafe.Pointer(nsURIC))
		}
		var qNameC *C.char
		if qName != "" {
			qNameC = C.CString(qName)
			defer C.free(unsafe.Pointer(qNameC))
		}
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		C.QDomElement_SetAttributeNS(ptr.Pointer(), C.struct_QtXml_PackedString{data: nsURIC, len: C.longlong(len(nsURI))}, C.struct_QtXml_PackedString{data: qNameC, len: C.longlong(len(qName))}, C.struct_QtXml_PackedString{data: valueC, len: C.longlong(len(value))})
	}
}

func (ptr *QDomElement) SetAttributeNS6(nsURI string, qName string, value float64) {
	if ptr.Pointer() != nil {
		var nsURIC *C.char
		if nsURI != "" {
			nsURIC = C.CString(nsURI)
			defer C.free(unsafe.Pointer(nsURIC))
		}
		var qNameC *C.char
		if qName != "" {
			qNameC = C.CString(qName)
			defer C.free(unsafe.Pointer(qNameC))
		}
		C.QDomElement_SetAttributeNS6(ptr.Pointer(), C.struct_QtXml_PackedString{data: nsURIC, len: C.longlong(len(nsURI))}, C.struct_QtXml_PackedString{data: qNameC, len: C.longlong(len(qName))}, C.double(value))
	}
}

func (ptr *QDomElement) SetAttributeNS2(nsURI string, qName string, value int) {
	if ptr.Pointer() != nil {
		var nsURIC *C.char
		if nsURI != "" {
			nsURIC = C.CString(nsURI)
			defer C.free(unsafe.Pointer(nsURIC))
		}
		var qNameC *C.char
		if qName != "" {
			qNameC = C.CString(qName)
			defer C.free(unsafe.Pointer(qNameC))
		}
		C.QDomElement_SetAttributeNS2(ptr.Pointer(), C.struct_QtXml_PackedString{data: nsURIC, len: C.longlong(len(nsURI))}, C.struct_QtXml_PackedString{data: qNameC, len: C.longlong(len(qName))}, C.int(int32(value)))
	}
}

func (ptr *QDomElement) SetAttributeNS4(nsURI string, qName string, value int64) {
	if ptr.Pointer() != nil {
		var nsURIC *C.char
		if nsURI != "" {
			nsURIC = C.CString(nsURI)
			defer C.free(unsafe.Pointer(nsURIC))
		}
		var qNameC *C.char
		if qName != "" {
			qNameC = C.CString(qName)
			defer C.free(unsafe.Pointer(qNameC))
		}
		C.QDomElement_SetAttributeNS4(ptr.Pointer(), C.struct_QtXml_PackedString{data: nsURIC, len: C.longlong(len(nsURI))}, C.struct_QtXml_PackedString{data: qNameC, len: C.longlong(len(qName))}, C.longlong(value))
	}
}

func (ptr *QDomElement) SetAttributeNS5(nsURI string, qName string, value uint64) {
	if ptr.Pointer() != nil {
		var nsURIC *C.char
		if nsURI != "" {
			nsURIC = C.CString(nsURI)
			defer C.free(unsafe.Pointer(nsURIC))
		}
		var qNameC *C.char
		if qName != "" {
			qNameC = C.CString(qName)
			defer C.free(unsafe.Pointer(qNameC))
		}
		C.QDomElement_SetAttributeNS5(ptr.Pointer(), C.struct_QtXml_PackedString{data: nsURIC, len: C.longlong(len(nsURI))}, C.struct_QtXml_PackedString{data: qNameC, len: C.longlong(len(qName))}, C.ulonglong(value))
	}
}

func (ptr *QDomElement) SetAttributeNS3(nsURI string, qName string, value uint) {
	if ptr.Pointer() != nil {
		var nsURIC *C.char
		if nsURI != "" {
			nsURIC = C.CString(nsURI)
			defer C.free(unsafe.Pointer(nsURIC))
		}
		var qNameC *C.char
		if qName != "" {
			qNameC = C.CString(qName)
			defer C.free(unsafe.Pointer(qNameC))
		}
		C.QDomElement_SetAttributeNS3(ptr.Pointer(), C.struct_QtXml_PackedString{data: nsURIC, len: C.longlong(len(nsURI))}, C.struct_QtXml_PackedString{data: qNameC, len: C.longlong(len(qName))}, C.uint(uint32(value)))
	}
}

func (ptr *QDomElement) SetTagName(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QDomElement_SetTagName(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QDomElement) ElementsByTagName(tagname string) *QDomNodeList {
	if ptr.Pointer() != nil {
		var tagnameC *C.char
		if tagname != "" {
			tagnameC = C.CString(tagname)
			defer C.free(unsafe.Pointer(tagnameC))
		}
		tmpValue := NewQDomNodeListFromPointer(C.QDomElement_ElementsByTagName(ptr.Pointer(), C.struct_QtXml_PackedString{data: tagnameC, len: C.longlong(len(tagname))}))
		runtime.SetFinalizer(tmpValue, (*QDomNodeList).DestroyQDomNodeList)
		return tmpValue
	}
	return nil
}

func (ptr *QDomElement) ElementsByTagNameNS(nsURI string, localName string) *QDomNodeList {
	if ptr.Pointer() != nil {
		var nsURIC *C.char
		if nsURI != "" {
			nsURIC = C.CString(nsURI)
			defer C.free(unsafe.Pointer(nsURIC))
		}
		var localNameC *C.char
		if localName != "" {
			localNameC = C.CString(localName)
			defer C.free(unsafe.Pointer(localNameC))
		}
		tmpValue := NewQDomNodeListFromPointer(C.QDomElement_ElementsByTagNameNS(ptr.Pointer(), C.struct_QtXml_PackedString{data: nsURIC, len: C.longlong(len(nsURI))}, C.struct_QtXml_PackedString{data: localNameC, len: C.longlong(len(localName))}))
		runtime.SetFinalizer(tmpValue, (*QDomNodeList).DestroyQDomNodeList)
		return tmpValue
	}
	return nil
}

func (ptr *QDomElement) Attribute(name string, defValue string) string {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var defValueC *C.char
		if defValue != "" {
			defValueC = C.CString(defValue)
			defer C.free(unsafe.Pointer(defValueC))
		}
		return cGoUnpackString(C.QDomElement_Attribute(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtXml_PackedString{data: defValueC, len: C.longlong(len(defValue))}))
	}
	return ""
}

func (ptr *QDomElement) AttributeNS(nsURI string, localName string, defValue string) string {
	if ptr.Pointer() != nil {
		var nsURIC *C.char
		if nsURI != "" {
			nsURIC = C.CString(nsURI)
			defer C.free(unsafe.Pointer(nsURIC))
		}
		var localNameC *C.char
		if localName != "" {
			localNameC = C.CString(localName)
			defer C.free(unsafe.Pointer(localNameC))
		}
		var defValueC *C.char
		if defValue != "" {
			defValueC = C.CString(defValue)
			defer C.free(unsafe.Pointer(defValueC))
		}
		return cGoUnpackString(C.QDomElement_AttributeNS(ptr.Pointer(), C.struct_QtXml_PackedString{data: nsURIC, len: C.longlong(len(nsURI))}, C.struct_QtXml_PackedString{data: localNameC, len: C.longlong(len(localName))}, C.struct_QtXml_PackedString{data: defValueC, len: C.longlong(len(defValue))}))
	}
	return ""
}

func (ptr *QDomElement) TagName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDomElement_TagName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomElement) Text() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDomElement_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomElement) HasAttribute(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QDomElement_HasAttribute(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

func (ptr *QDomElement) HasAttributeNS(nsURI string, localName string) bool {
	if ptr.Pointer() != nil {
		var nsURIC *C.char
		if nsURI != "" {
			nsURIC = C.CString(nsURI)
			defer C.free(unsafe.Pointer(nsURIC))
		}
		var localNameC *C.char
		if localName != "" {
			localNameC = C.CString(localName)
			defer C.free(unsafe.Pointer(localNameC))
		}
		return int8(C.QDomElement_HasAttributeNS(ptr.Pointer(), C.struct_QtXml_PackedString{data: nsURIC, len: C.longlong(len(nsURI))}, C.struct_QtXml_PackedString{data: localNameC, len: C.longlong(len(localName))})) != 0
	}
	return false
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

func NewQDomEntityFromPointer(ptr unsafe.Pointer) (n *QDomEntity) {
	n = new(QDomEntity)
	n.SetPointer(ptr)
	return
}

func (ptr *QDomEntity) DestroyQDomEntity() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQDomEntity() *QDomEntity {
	tmpValue := NewQDomEntityFromPointer(C.QDomEntity_NewQDomEntity())
	runtime.SetFinalizer(tmpValue, (*QDomEntity).DestroyQDomEntity)
	return tmpValue
}

func NewQDomEntity2(x QDomEntity_ITF) *QDomEntity {
	tmpValue := NewQDomEntityFromPointer(C.QDomEntity_NewQDomEntity2(PointerFromQDomEntity(x)))
	runtime.SetFinalizer(tmpValue, (*QDomEntity).DestroyQDomEntity)
	return tmpValue
}

func (ptr *QDomEntity) NotationName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDomEntity_NotationName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomEntity) PublicId() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDomEntity_PublicId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomEntity) SystemId() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDomEntity_SystemId(ptr.Pointer()))
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

func NewQDomEntityReferenceFromPointer(ptr unsafe.Pointer) (n *QDomEntityReference) {
	n = new(QDomEntityReference)
	n.SetPointer(ptr)
	return
}

func (ptr *QDomEntityReference) DestroyQDomEntityReference() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQDomEntityReference() *QDomEntityReference {
	tmpValue := NewQDomEntityReferenceFromPointer(C.QDomEntityReference_NewQDomEntityReference())
	runtime.SetFinalizer(tmpValue, (*QDomEntityReference).DestroyQDomEntityReference)
	return tmpValue
}

func NewQDomEntityReference2(x QDomEntityReference_ITF) *QDomEntityReference {
	tmpValue := NewQDomEntityReferenceFromPointer(C.QDomEntityReference_NewQDomEntityReference2(PointerFromQDomEntityReference(x)))
	runtime.SetFinalizer(tmpValue, (*QDomEntityReference).DestroyQDomEntityReference)
	return tmpValue
}

type QDomImplementation struct {
	ptr unsafe.Pointer
}

type QDomImplementation_ITF interface {
	QDomImplementation_PTR() *QDomImplementation
}

func (ptr *QDomImplementation) QDomImplementation_PTR() *QDomImplementation {
	return ptr
}

func (ptr *QDomImplementation) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDomImplementation) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQDomImplementation(ptr QDomImplementation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomImplementation_PTR().Pointer()
	}
	return nil
}

func NewQDomImplementationFromPointer(ptr unsafe.Pointer) (n *QDomImplementation) {
	n = new(QDomImplementation)
	n.SetPointer(ptr)
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

func (ptr *QDomImplementation) CreateDocument(nsURI string, qName string, doctype QDomDocumentType_ITF) *QDomDocument {
	if ptr.Pointer() != nil {
		var nsURIC *C.char
		if nsURI != "" {
			nsURIC = C.CString(nsURI)
			defer C.free(unsafe.Pointer(nsURIC))
		}
		var qNameC *C.char
		if qName != "" {
			qNameC = C.CString(qName)
			defer C.free(unsafe.Pointer(qNameC))
		}
		tmpValue := NewQDomDocumentFromPointer(C.QDomImplementation_CreateDocument(ptr.Pointer(), C.struct_QtXml_PackedString{data: nsURIC, len: C.longlong(len(nsURI))}, C.struct_QtXml_PackedString{data: qNameC, len: C.longlong(len(qName))}, PointerFromQDomDocumentType(doctype)))
		runtime.SetFinalizer(tmpValue, (*QDomDocument).DestroyQDomDocument)
		return tmpValue
	}
	return nil
}

func (ptr *QDomImplementation) CreateDocumentType(qName string, publicId string, systemId string) *QDomDocumentType {
	if ptr.Pointer() != nil {
		var qNameC *C.char
		if qName != "" {
			qNameC = C.CString(qName)
			defer C.free(unsafe.Pointer(qNameC))
		}
		var publicIdC *C.char
		if publicId != "" {
			publicIdC = C.CString(publicId)
			defer C.free(unsafe.Pointer(publicIdC))
		}
		var systemIdC *C.char
		if systemId != "" {
			systemIdC = C.CString(systemId)
			defer C.free(unsafe.Pointer(systemIdC))
		}
		tmpValue := NewQDomDocumentTypeFromPointer(C.QDomImplementation_CreateDocumentType(ptr.Pointer(), C.struct_QtXml_PackedString{data: qNameC, len: C.longlong(len(qName))}, C.struct_QtXml_PackedString{data: publicIdC, len: C.longlong(len(publicId))}, C.struct_QtXml_PackedString{data: systemIdC, len: C.longlong(len(systemId))}))
		runtime.SetFinalizer(tmpValue, (*QDomDocumentType).DestroyQDomDocumentType)
		return tmpValue
	}
	return nil
}

func NewQDomImplementation() *QDomImplementation {
	tmpValue := NewQDomImplementationFromPointer(C.QDomImplementation_NewQDomImplementation())
	runtime.SetFinalizer(tmpValue, (*QDomImplementation).DestroyQDomImplementation)
	return tmpValue
}

func NewQDomImplementation2(x QDomImplementation_ITF) *QDomImplementation {
	tmpValue := NewQDomImplementationFromPointer(C.QDomImplementation_NewQDomImplementation2(PointerFromQDomImplementation(x)))
	runtime.SetFinalizer(tmpValue, (*QDomImplementation).DestroyQDomImplementation)
	return tmpValue
}

func QDomImplementation_InvalidDataPolicy() QDomImplementation__InvalidDataPolicy {
	return QDomImplementation__InvalidDataPolicy(C.QDomImplementation_QDomImplementation_InvalidDataPolicy())
}

func (ptr *QDomImplementation) InvalidDataPolicy() QDomImplementation__InvalidDataPolicy {
	return QDomImplementation__InvalidDataPolicy(C.QDomImplementation_QDomImplementation_InvalidDataPolicy())
}

func (ptr *QDomImplementation) IsNull() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDomImplementation_IsNull(ptr.Pointer())) != 0
	}
	return false
}

func QDomImplementation_SetInvalidDataPolicy(policy QDomImplementation__InvalidDataPolicy) {
	C.QDomImplementation_QDomImplementation_SetInvalidDataPolicy(C.longlong(policy))
}

func (ptr *QDomImplementation) SetInvalidDataPolicy(policy QDomImplementation__InvalidDataPolicy) {
	C.QDomImplementation_QDomImplementation_SetInvalidDataPolicy(C.longlong(policy))
}

func (ptr *QDomImplementation) DestroyQDomImplementation() {
	if ptr.Pointer() != nil {
		C.QDomImplementation_DestroyQDomImplementation(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDomImplementation) HasFeature(feature string, version string) bool {
	if ptr.Pointer() != nil {
		var featureC *C.char
		if feature != "" {
			featureC = C.CString(feature)
			defer C.free(unsafe.Pointer(featureC))
		}
		var versionC *C.char
		if version != "" {
			versionC = C.CString(version)
			defer C.free(unsafe.Pointer(versionC))
		}
		return int8(C.QDomImplementation_HasFeature(ptr.Pointer(), C.struct_QtXml_PackedString{data: featureC, len: C.longlong(len(feature))}, C.struct_QtXml_PackedString{data: versionC, len: C.longlong(len(version))})) != 0
	}
	return false
}

type QDomNamedNodeMap struct {
	ptr unsafe.Pointer
}

type QDomNamedNodeMap_ITF interface {
	QDomNamedNodeMap_PTR() *QDomNamedNodeMap
}

func (ptr *QDomNamedNodeMap) QDomNamedNodeMap_PTR() *QDomNamedNodeMap {
	return ptr
}

func (ptr *QDomNamedNodeMap) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDomNamedNodeMap) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQDomNamedNodeMap(ptr QDomNamedNodeMap_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNamedNodeMap_PTR().Pointer()
	}
	return nil
}

func NewQDomNamedNodeMapFromPointer(ptr unsafe.Pointer) (n *QDomNamedNodeMap) {
	n = new(QDomNamedNodeMap)
	n.SetPointer(ptr)
	return
}
func NewQDomNamedNodeMap() *QDomNamedNodeMap {
	tmpValue := NewQDomNamedNodeMapFromPointer(C.QDomNamedNodeMap_NewQDomNamedNodeMap())
	runtime.SetFinalizer(tmpValue, (*QDomNamedNodeMap).DestroyQDomNamedNodeMap)
	return tmpValue
}

func NewQDomNamedNodeMap2(n QDomNamedNodeMap_ITF) *QDomNamedNodeMap {
	tmpValue := NewQDomNamedNodeMapFromPointer(C.QDomNamedNodeMap_NewQDomNamedNodeMap2(PointerFromQDomNamedNodeMap(n)))
	runtime.SetFinalizer(tmpValue, (*QDomNamedNodeMap).DestroyQDomNamedNodeMap)
	return tmpValue
}

func (ptr *QDomNamedNodeMap) RemoveNamedItem(name string) *QDomNode {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		tmpValue := NewQDomNodeFromPointer(C.QDomNamedNodeMap_RemoveNamedItem(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}))
		runtime.SetFinalizer(tmpValue, (*QDomNode).DestroyQDomNode)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNamedNodeMap) RemoveNamedItemNS(nsURI string, localName string) *QDomNode {
	if ptr.Pointer() != nil {
		var nsURIC *C.char
		if nsURI != "" {
			nsURIC = C.CString(nsURI)
			defer C.free(unsafe.Pointer(nsURIC))
		}
		var localNameC *C.char
		if localName != "" {
			localNameC = C.CString(localName)
			defer C.free(unsafe.Pointer(localNameC))
		}
		tmpValue := NewQDomNodeFromPointer(C.QDomNamedNodeMap_RemoveNamedItemNS(ptr.Pointer(), C.struct_QtXml_PackedString{data: nsURIC, len: C.longlong(len(nsURI))}, C.struct_QtXml_PackedString{data: localNameC, len: C.longlong(len(localName))}))
		runtime.SetFinalizer(tmpValue, (*QDomNode).DestroyQDomNode)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNamedNodeMap) SetNamedItem(newNode QDomNode_ITF) *QDomNode {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomNodeFromPointer(C.QDomNamedNodeMap_SetNamedItem(ptr.Pointer(), PointerFromQDomNode(newNode)))
		runtime.SetFinalizer(tmpValue, (*QDomNode).DestroyQDomNode)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNamedNodeMap) SetNamedItemNS(newNode QDomNode_ITF) *QDomNode {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomNodeFromPointer(C.QDomNamedNodeMap_SetNamedItemNS(ptr.Pointer(), PointerFromQDomNode(newNode)))
		runtime.SetFinalizer(tmpValue, (*QDomNode).DestroyQDomNode)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNamedNodeMap) DestroyQDomNamedNodeMap() {
	if ptr.Pointer() != nil {
		C.QDomNamedNodeMap_DestroyQDomNamedNodeMap(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDomNamedNodeMap) Item(index int) *QDomNode {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomNodeFromPointer(C.QDomNamedNodeMap_Item(ptr.Pointer(), C.int(int32(index))))
		runtime.SetFinalizer(tmpValue, (*QDomNode).DestroyQDomNode)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNamedNodeMap) NamedItem(name string) *QDomNode {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		tmpValue := NewQDomNodeFromPointer(C.QDomNamedNodeMap_NamedItem(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}))
		runtime.SetFinalizer(tmpValue, (*QDomNode).DestroyQDomNode)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNamedNodeMap) NamedItemNS(nsURI string, localName string) *QDomNode {
	if ptr.Pointer() != nil {
		var nsURIC *C.char
		if nsURI != "" {
			nsURIC = C.CString(nsURI)
			defer C.free(unsafe.Pointer(nsURIC))
		}
		var localNameC *C.char
		if localName != "" {
			localNameC = C.CString(localName)
			defer C.free(unsafe.Pointer(localNameC))
		}
		tmpValue := NewQDomNodeFromPointer(C.QDomNamedNodeMap_NamedItemNS(ptr.Pointer(), C.struct_QtXml_PackedString{data: nsURIC, len: C.longlong(len(nsURI))}, C.struct_QtXml_PackedString{data: localNameC, len: C.longlong(len(localName))}))
		runtime.SetFinalizer(tmpValue, (*QDomNode).DestroyQDomNode)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNamedNodeMap) Contains(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QDomNamedNodeMap_Contains(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

func (ptr *QDomNamedNodeMap) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDomNamedNodeMap_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDomNamedNodeMap) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDomNamedNodeMap_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDomNamedNodeMap) Length() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDomNamedNodeMap_Length(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDomNamedNodeMap) Size() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDomNamedNodeMap_Size(ptr.Pointer())))
	}
	return 0
}

type QDomNode struct {
	ptr unsafe.Pointer
}

type QDomNode_ITF interface {
	QDomNode_PTR() *QDomNode
}

func (ptr *QDomNode) QDomNode_PTR() *QDomNode {
	return ptr
}

func (ptr *QDomNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDomNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQDomNode(ptr QDomNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNode_PTR().Pointer()
	}
	return nil
}

func NewQDomNodeFromPointer(ptr unsafe.Pointer) (n *QDomNode) {
	n = new(QDomNode)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QDomNode__EncodingPolicy
//QDomNode::EncodingPolicy
type QDomNode__EncodingPolicy int64

const (
	QDomNode__EncodingFromDocument   QDomNode__EncodingPolicy = QDomNode__EncodingPolicy(1)
	QDomNode__EncodingFromTextStream QDomNode__EncodingPolicy = QDomNode__EncodingPolicy(2)
)

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

func (ptr *QDomNode) AppendChild(newChild QDomNode_ITF) *QDomNode {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomNodeFromPointer(C.QDomNode_AppendChild(ptr.Pointer(), PointerFromQDomNode(newChild)))
		runtime.SetFinalizer(tmpValue, (*QDomNode).DestroyQDomNode)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) InsertAfter(newChild QDomNode_ITF, refChild QDomNode_ITF) *QDomNode {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomNodeFromPointer(C.QDomNode_InsertAfter(ptr.Pointer(), PointerFromQDomNode(newChild), PointerFromQDomNode(refChild)))
		runtime.SetFinalizer(tmpValue, (*QDomNode).DestroyQDomNode)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) InsertBefore(newChild QDomNode_ITF, refChild QDomNode_ITF) *QDomNode {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomNodeFromPointer(C.QDomNode_InsertBefore(ptr.Pointer(), PointerFromQDomNode(newChild), PointerFromQDomNode(refChild)))
		runtime.SetFinalizer(tmpValue, (*QDomNode).DestroyQDomNode)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) RemoveChild(oldChild QDomNode_ITF) *QDomNode {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomNodeFromPointer(C.QDomNode_RemoveChild(ptr.Pointer(), PointerFromQDomNode(oldChild)))
		runtime.SetFinalizer(tmpValue, (*QDomNode).DestroyQDomNode)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) ReplaceChild(newChild QDomNode_ITF, oldChild QDomNode_ITF) *QDomNode {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomNodeFromPointer(C.QDomNode_ReplaceChild(ptr.Pointer(), PointerFromQDomNode(newChild), PointerFromQDomNode(oldChild)))
		runtime.SetFinalizer(tmpValue, (*QDomNode).DestroyQDomNode)
		return tmpValue
	}
	return nil
}

func NewQDomNode() *QDomNode {
	tmpValue := NewQDomNodeFromPointer(C.QDomNode_NewQDomNode())
	runtime.SetFinalizer(tmpValue, (*QDomNode).DestroyQDomNode)
	return tmpValue
}

func NewQDomNode2(n QDomNode_ITF) *QDomNode {
	tmpValue := NewQDomNodeFromPointer(C.QDomNode_NewQDomNode2(PointerFromQDomNode(n)))
	runtime.SetFinalizer(tmpValue, (*QDomNode).DestroyQDomNode)
	return tmpValue
}

func (ptr *QDomNode) Clear() {
	if ptr.Pointer() != nil {
		C.QDomNode_Clear(ptr.Pointer())
	}
}

func (ptr *QDomNode) Normalize() {
	if ptr.Pointer() != nil {
		C.QDomNode_Normalize(ptr.Pointer())
	}
}

func (ptr *QDomNode) SetNodeValue(v string) {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		C.QDomNode_SetNodeValue(ptr.Pointer(), C.struct_QtXml_PackedString{data: vC, len: C.longlong(len(v))})
	}
}

func (ptr *QDomNode) SetPrefix(pre string) {
	if ptr.Pointer() != nil {
		var preC *C.char
		if pre != "" {
			preC = C.CString(pre)
			defer C.free(unsafe.Pointer(preC))
		}
		C.QDomNode_SetPrefix(ptr.Pointer(), C.struct_QtXml_PackedString{data: preC, len: C.longlong(len(pre))})
	}
}

func (ptr *QDomNode) DestroyQDomNode() {
	if ptr.Pointer() != nil {
		C.QDomNode_DestroyQDomNode(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDomNode) ToAttr() *QDomAttr {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomAttrFromPointer(C.QDomNode_ToAttr(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDomAttr).DestroyQDomAttr)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) ToCDATASection() *QDomCDATASection {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomCDATASectionFromPointer(C.QDomNode_ToCDATASection(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDomCDATASection).DestroyQDomCDATASection)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) ToCharacterData() *QDomCharacterData {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomCharacterDataFromPointer(C.QDomNode_ToCharacterData(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDomCharacterData).DestroyQDomCharacterData)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) ToComment() *QDomComment {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomCommentFromPointer(C.QDomNode_ToComment(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDomComment).DestroyQDomComment)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) OwnerDocument() *QDomDocument {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomDocumentFromPointer(C.QDomNode_OwnerDocument(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDomDocument).DestroyQDomDocument)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) ToDocument() *QDomDocument {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomDocumentFromPointer(C.QDomNode_ToDocument(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDomDocument).DestroyQDomDocument)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) ToDocumentFragment() *QDomDocumentFragment {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomDocumentFragmentFromPointer(C.QDomNode_ToDocumentFragment(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDomDocumentFragment).DestroyQDomDocumentFragment)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) ToDocumentType() *QDomDocumentType {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomDocumentTypeFromPointer(C.QDomNode_ToDocumentType(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDomDocumentType).DestroyQDomDocumentType)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) FirstChildElement(tagName string) *QDomElement {
	if ptr.Pointer() != nil {
		var tagNameC *C.char
		if tagName != "" {
			tagNameC = C.CString(tagName)
			defer C.free(unsafe.Pointer(tagNameC))
		}
		tmpValue := NewQDomElementFromPointer(C.QDomNode_FirstChildElement(ptr.Pointer(), C.struct_QtXml_PackedString{data: tagNameC, len: C.longlong(len(tagName))}))
		runtime.SetFinalizer(tmpValue, (*QDomElement).DestroyQDomElement)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) LastChildElement(tagName string) *QDomElement {
	if ptr.Pointer() != nil {
		var tagNameC *C.char
		if tagName != "" {
			tagNameC = C.CString(tagName)
			defer C.free(unsafe.Pointer(tagNameC))
		}
		tmpValue := NewQDomElementFromPointer(C.QDomNode_LastChildElement(ptr.Pointer(), C.struct_QtXml_PackedString{data: tagNameC, len: C.longlong(len(tagName))}))
		runtime.SetFinalizer(tmpValue, (*QDomElement).DestroyQDomElement)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) NextSiblingElement(tagName string) *QDomElement {
	if ptr.Pointer() != nil {
		var tagNameC *C.char
		if tagName != "" {
			tagNameC = C.CString(tagName)
			defer C.free(unsafe.Pointer(tagNameC))
		}
		tmpValue := NewQDomElementFromPointer(C.QDomNode_NextSiblingElement(ptr.Pointer(), C.struct_QtXml_PackedString{data: tagNameC, len: C.longlong(len(tagName))}))
		runtime.SetFinalizer(tmpValue, (*QDomElement).DestroyQDomElement)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) PreviousSiblingElement(tagName string) *QDomElement {
	if ptr.Pointer() != nil {
		var tagNameC *C.char
		if tagName != "" {
			tagNameC = C.CString(tagName)
			defer C.free(unsafe.Pointer(tagNameC))
		}
		tmpValue := NewQDomElementFromPointer(C.QDomNode_PreviousSiblingElement(ptr.Pointer(), C.struct_QtXml_PackedString{data: tagNameC, len: C.longlong(len(tagName))}))
		runtime.SetFinalizer(tmpValue, (*QDomElement).DestroyQDomElement)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) ToElement() *QDomElement {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomElementFromPointer(C.QDomNode_ToElement(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDomElement).DestroyQDomElement)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) ToEntity() *QDomEntity {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomEntityFromPointer(C.QDomNode_ToEntity(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDomEntity).DestroyQDomEntity)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) ToEntityReference() *QDomEntityReference {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomEntityReferenceFromPointer(C.QDomNode_ToEntityReference(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDomEntityReference).DestroyQDomEntityReference)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) CloneNode(deep bool) *QDomNode {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomNodeFromPointer(C.QDomNode_CloneNode(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(deep)))))
		runtime.SetFinalizer(tmpValue, (*QDomNode).DestroyQDomNode)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) FirstChild() *QDomNode {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomNodeFromPointer(C.QDomNode_FirstChild(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDomNode).DestroyQDomNode)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) LastChild() *QDomNode {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomNodeFromPointer(C.QDomNode_LastChild(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDomNode).DestroyQDomNode)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) NamedItem(name string) *QDomNode {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		tmpValue := NewQDomNodeFromPointer(C.QDomNode_NamedItem(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}))
		runtime.SetFinalizer(tmpValue, (*QDomNode).DestroyQDomNode)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) NextSibling() *QDomNode {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomNodeFromPointer(C.QDomNode_NextSibling(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDomNode).DestroyQDomNode)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) ParentNode() *QDomNode {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomNodeFromPointer(C.QDomNode_ParentNode(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDomNode).DestroyQDomNode)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) PreviousSibling() *QDomNode {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomNodeFromPointer(C.QDomNode_PreviousSibling(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDomNode).DestroyQDomNode)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) NodeType() QDomNode__NodeType {
	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomNode_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomNode) ChildNodes() *QDomNodeList {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomNodeListFromPointer(C.QDomNode_ChildNodes(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDomNodeList).DestroyQDomNodeList)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) ToNotation() *QDomNotation {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomNotationFromPointer(C.QDomNode_ToNotation(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDomNotation).DestroyQDomNotation)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) ToProcessingInstruction() *QDomProcessingInstruction {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomProcessingInstructionFromPointer(C.QDomNode_ToProcessingInstruction(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDomProcessingInstruction).DestroyQDomProcessingInstruction)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) ToText() *QDomText {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomTextFromPointer(C.QDomNode_ToText(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDomText).DestroyQDomText)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNode) LocalName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDomNode_LocalName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomNode) NamespaceURI() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDomNode_NamespaceURI(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomNode) NodeName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDomNode_NodeName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomNode) NodeValue() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDomNode_NodeValue(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomNode) Prefix() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDomNode_Prefix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomNode) HasAttributes() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDomNode_HasAttributes(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDomNode) HasChildNodes() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDomNode_HasChildNodes(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDomNode) IsAttr() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDomNode_IsAttr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDomNode) IsCDATASection() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDomNode_IsCDATASection(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDomNode) IsCharacterData() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDomNode_IsCharacterData(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDomNode) IsComment() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDomNode_IsComment(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDomNode) IsDocument() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDomNode_IsDocument(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDomNode) IsDocumentFragment() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDomNode_IsDocumentFragment(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDomNode) IsDocumentType() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDomNode_IsDocumentType(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDomNode) IsElement() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDomNode_IsElement(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDomNode) IsEntity() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDomNode_IsEntity(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDomNode) IsEntityReference() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDomNode_IsEntityReference(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDomNode) IsNotation() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDomNode_IsNotation(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDomNode) IsNull() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDomNode_IsNull(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDomNode) IsProcessingInstruction() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDomNode_IsProcessingInstruction(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDomNode) IsSupported(feature string, version string) bool {
	if ptr.Pointer() != nil {
		var featureC *C.char
		if feature != "" {
			featureC = C.CString(feature)
			defer C.free(unsafe.Pointer(featureC))
		}
		var versionC *C.char
		if version != "" {
			versionC = C.CString(version)
			defer C.free(unsafe.Pointer(versionC))
		}
		return int8(C.QDomNode_IsSupported(ptr.Pointer(), C.struct_QtXml_PackedString{data: featureC, len: C.longlong(len(feature))}, C.struct_QtXml_PackedString{data: versionC, len: C.longlong(len(version))})) != 0
	}
	return false
}

func (ptr *QDomNode) IsText() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDomNode_IsText(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDomNode) ColumnNumber() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDomNode_ColumnNumber(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDomNode) LineNumber() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDomNode_LineNumber(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDomNode) Save(stream core.QTextStream_ITF, indent int, encodingPolicy QDomNode__EncodingPolicy) {
	if ptr.Pointer() != nil {
		C.QDomNode_Save(ptr.Pointer(), core.PointerFromQTextStream(stream), C.int(int32(indent)), C.longlong(encodingPolicy))
	}
}

type QDomNodeList struct {
	ptr unsafe.Pointer
}

type QDomNodeList_ITF interface {
	QDomNodeList_PTR() *QDomNodeList
}

func (ptr *QDomNodeList) QDomNodeList_PTR() *QDomNodeList {
	return ptr
}

func (ptr *QDomNodeList) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDomNodeList) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQDomNodeList(ptr QDomNodeList_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomNodeList_PTR().Pointer()
	}
	return nil
}

func NewQDomNodeListFromPointer(ptr unsafe.Pointer) (n *QDomNodeList) {
	n = new(QDomNodeList)
	n.SetPointer(ptr)
	return
}
func NewQDomNodeList() *QDomNodeList {
	tmpValue := NewQDomNodeListFromPointer(C.QDomNodeList_NewQDomNodeList())
	runtime.SetFinalizer(tmpValue, (*QDomNodeList).DestroyQDomNodeList)
	return tmpValue
}

func NewQDomNodeList2(n QDomNodeList_ITF) *QDomNodeList {
	tmpValue := NewQDomNodeListFromPointer(C.QDomNodeList_NewQDomNodeList2(PointerFromQDomNodeList(n)))
	runtime.SetFinalizer(tmpValue, (*QDomNodeList).DestroyQDomNodeList)
	return tmpValue
}

func (ptr *QDomNodeList) DestroyQDomNodeList() {
	if ptr.Pointer() != nil {
		C.QDomNodeList_DestroyQDomNodeList(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDomNodeList) At(index int) *QDomNode {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomNodeFromPointer(C.QDomNodeList_At(ptr.Pointer(), C.int(int32(index))))
		runtime.SetFinalizer(tmpValue, (*QDomNode).DestroyQDomNode)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNodeList) Item(index int) *QDomNode {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomNodeFromPointer(C.QDomNodeList_Item(ptr.Pointer(), C.int(int32(index))))
		runtime.SetFinalizer(tmpValue, (*QDomNode).DestroyQDomNode)
		return tmpValue
	}
	return nil
}

func (ptr *QDomNodeList) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDomNodeList_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDomNodeList) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDomNodeList_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDomNodeList) Length() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDomNodeList_Length(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDomNodeList) Size() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDomNodeList_Size(ptr.Pointer())))
	}
	return 0
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

func NewQDomNotationFromPointer(ptr unsafe.Pointer) (n *QDomNotation) {
	n = new(QDomNotation)
	n.SetPointer(ptr)
	return
}

func (ptr *QDomNotation) DestroyQDomNotation() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQDomNotation() *QDomNotation {
	tmpValue := NewQDomNotationFromPointer(C.QDomNotation_NewQDomNotation())
	runtime.SetFinalizer(tmpValue, (*QDomNotation).DestroyQDomNotation)
	return tmpValue
}

func NewQDomNotation2(x QDomNotation_ITF) *QDomNotation {
	tmpValue := NewQDomNotationFromPointer(C.QDomNotation_NewQDomNotation2(PointerFromQDomNotation(x)))
	runtime.SetFinalizer(tmpValue, (*QDomNotation).DestroyQDomNotation)
	return tmpValue
}

func (ptr *QDomNotation) PublicId() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDomNotation_PublicId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomNotation) SystemId() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDomNotation_SystemId(ptr.Pointer()))
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

func NewQDomProcessingInstructionFromPointer(ptr unsafe.Pointer) (n *QDomProcessingInstruction) {
	n = new(QDomProcessingInstruction)
	n.SetPointer(ptr)
	return
}

func (ptr *QDomProcessingInstruction) DestroyQDomProcessingInstruction() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQDomProcessingInstruction() *QDomProcessingInstruction {
	tmpValue := NewQDomProcessingInstructionFromPointer(C.QDomProcessingInstruction_NewQDomProcessingInstruction())
	runtime.SetFinalizer(tmpValue, (*QDomProcessingInstruction).DestroyQDomProcessingInstruction)
	return tmpValue
}

func NewQDomProcessingInstruction2(x QDomProcessingInstruction_ITF) *QDomProcessingInstruction {
	tmpValue := NewQDomProcessingInstructionFromPointer(C.QDomProcessingInstruction_NewQDomProcessingInstruction2(PointerFromQDomProcessingInstruction(x)))
	runtime.SetFinalizer(tmpValue, (*QDomProcessingInstruction).DestroyQDomProcessingInstruction)
	return tmpValue
}

func (ptr *QDomProcessingInstruction) SetData(d string) {
	if ptr.Pointer() != nil {
		var dC *C.char
		if d != "" {
			dC = C.CString(d)
			defer C.free(unsafe.Pointer(dC))
		}
		C.QDomProcessingInstruction_SetData(ptr.Pointer(), C.struct_QtXml_PackedString{data: dC, len: C.longlong(len(d))})
	}
}

func (ptr *QDomProcessingInstruction) Data() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDomProcessingInstruction_Data(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomProcessingInstruction) Target() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDomProcessingInstruction_Target(ptr.Pointer()))
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

func NewQDomTextFromPointer(ptr unsafe.Pointer) (n *QDomText) {
	n = new(QDomText)
	n.SetPointer(ptr)
	return
}

func (ptr *QDomText) DestroyQDomText() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDomText) SplitText(offset int) *QDomText {
	if ptr.Pointer() != nil {
		tmpValue := NewQDomTextFromPointer(C.QDomText_SplitText(ptr.Pointer(), C.int(int32(offset))))
		runtime.SetFinalizer(tmpValue, (*QDomText).DestroyQDomText)
		return tmpValue
	}
	return nil
}

func NewQDomText() *QDomText {
	tmpValue := NewQDomTextFromPointer(C.QDomText_NewQDomText())
	runtime.SetFinalizer(tmpValue, (*QDomText).DestroyQDomText)
	return tmpValue
}

func NewQDomText2(x QDomText_ITF) *QDomText {
	tmpValue := NewQDomTextFromPointer(C.QDomText_NewQDomText2(PointerFromQDomText(x)))
	runtime.SetFinalizer(tmpValue, (*QDomText).DestroyQDomText)
	return tmpValue
}

type QXmlAttributes struct {
	ptr unsafe.Pointer
}

type QXmlAttributes_ITF interface {
	QXmlAttributes_PTR() *QXmlAttributes
}

func (ptr *QXmlAttributes) QXmlAttributes_PTR() *QXmlAttributes {
	return ptr
}

func (ptr *QXmlAttributes) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlAttributes) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQXmlAttributes(ptr QXmlAttributes_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlAttributes_PTR().Pointer()
	}
	return nil
}

func NewQXmlAttributesFromPointer(ptr unsafe.Pointer) (n *QXmlAttributes) {
	n = new(QXmlAttributes)
	n.SetPointer(ptr)
	return
}
func NewQXmlAttributes() *QXmlAttributes {
	return NewQXmlAttributesFromPointer(C.QXmlAttributes_NewQXmlAttributes())
}

func NewQXmlAttributes3(vqx QXmlAttributes_ITF) *QXmlAttributes {
	return NewQXmlAttributesFromPointer(C.QXmlAttributes_NewQXmlAttributes3(PointerFromQXmlAttributes(vqx)))
}

func NewQXmlAttributes2(vqx QXmlAttributes_ITF) *QXmlAttributes {
	return NewQXmlAttributesFromPointer(C.QXmlAttributes_NewQXmlAttributes2(PointerFromQXmlAttributes(vqx)))
}

func (ptr *QXmlAttributes) Append(qName string, uri string, localPart string, value string) {
	if ptr.Pointer() != nil {
		var qNameC *C.char
		if qName != "" {
			qNameC = C.CString(qName)
			defer C.free(unsafe.Pointer(qNameC))
		}
		var uriC *C.char
		if uri != "" {
			uriC = C.CString(uri)
			defer C.free(unsafe.Pointer(uriC))
		}
		var localPartC *C.char
		if localPart != "" {
			localPartC = C.CString(localPart)
			defer C.free(unsafe.Pointer(localPartC))
		}
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		C.QXmlAttributes_Append(ptr.Pointer(), C.struct_QtXml_PackedString{data: qNameC, len: C.longlong(len(qName))}, C.struct_QtXml_PackedString{data: uriC, len: C.longlong(len(uri))}, C.struct_QtXml_PackedString{data: localPartC, len: C.longlong(len(localPart))}, C.struct_QtXml_PackedString{data: valueC, len: C.longlong(len(value))})
	}
}

func (ptr *QXmlAttributes) Clear() {
	if ptr.Pointer() != nil {
		C.QXmlAttributes_Clear(ptr.Pointer())
	}
}

func (ptr *QXmlAttributes) Swap(other QXmlAttributes_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlAttributes_Swap(ptr.Pointer(), PointerFromQXmlAttributes(other))
	}
}

//export callbackQXmlAttributes_DestroyQXmlAttributes
func callbackQXmlAttributes_DestroyQXmlAttributes(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QXmlAttributes"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlAttributesFromPointer(ptr).DestroyQXmlAttributesDefault()
	}
}

func (ptr *QXmlAttributes) ConnectDestroyQXmlAttributes(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QXmlAttributes"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QXmlAttributes", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QXmlAttributes", f)
		}
	}
}

func (ptr *QXmlAttributes) DisconnectDestroyQXmlAttributes() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QXmlAttributes")
	}
}

func (ptr *QXmlAttributes) DestroyQXmlAttributes() {
	if ptr.Pointer() != nil {
		C.QXmlAttributes_DestroyQXmlAttributes(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXmlAttributes) DestroyQXmlAttributesDefault() {
	if ptr.Pointer() != nil {
		C.QXmlAttributes_DestroyQXmlAttributesDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXmlAttributes) LocalName(index int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlAttributes_LocalName(ptr.Pointer(), C.int(int32(index))))
	}
	return ""
}

func (ptr *QXmlAttributes) QName(index int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlAttributes_QName(ptr.Pointer(), C.int(int32(index))))
	}
	return ""
}

func (ptr *QXmlAttributes) Type2(qName string) string {
	if ptr.Pointer() != nil {
		var qNameC *C.char
		if qName != "" {
			qNameC = C.CString(qName)
			defer C.free(unsafe.Pointer(qNameC))
		}
		return cGoUnpackString(C.QXmlAttributes_Type2(ptr.Pointer(), C.struct_QtXml_PackedString{data: qNameC, len: C.longlong(len(qName))}))
	}
	return ""
}

func (ptr *QXmlAttributes) Type3(uri string, localName string) string {
	if ptr.Pointer() != nil {
		var uriC *C.char
		if uri != "" {
			uriC = C.CString(uri)
			defer C.free(unsafe.Pointer(uriC))
		}
		var localNameC *C.char
		if localName != "" {
			localNameC = C.CString(localName)
			defer C.free(unsafe.Pointer(localNameC))
		}
		return cGoUnpackString(C.QXmlAttributes_Type3(ptr.Pointer(), C.struct_QtXml_PackedString{data: uriC, len: C.longlong(len(uri))}, C.struct_QtXml_PackedString{data: localNameC, len: C.longlong(len(localName))}))
	}
	return ""
}

func (ptr *QXmlAttributes) Type(index int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlAttributes_Type(ptr.Pointer(), C.int(int32(index))))
	}
	return ""
}

func (ptr *QXmlAttributes) Uri(index int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlAttributes_Uri(ptr.Pointer(), C.int(int32(index))))
	}
	return ""
}

func (ptr *QXmlAttributes) Value3(qName core.QLatin1String_ITF) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlAttributes_Value3(ptr.Pointer(), core.PointerFromQLatin1String(qName)))
	}
	return ""
}

func (ptr *QXmlAttributes) Value2(qName string) string {
	if ptr.Pointer() != nil {
		var qNameC *C.char
		if qName != "" {
			qNameC = C.CString(qName)
			defer C.free(unsafe.Pointer(qNameC))
		}
		return cGoUnpackString(C.QXmlAttributes_Value2(ptr.Pointer(), C.struct_QtXml_PackedString{data: qNameC, len: C.longlong(len(qName))}))
	}
	return ""
}

func (ptr *QXmlAttributes) Value4(uri string, localName string) string {
	if ptr.Pointer() != nil {
		var uriC *C.char
		if uri != "" {
			uriC = C.CString(uri)
			defer C.free(unsafe.Pointer(uriC))
		}
		var localNameC *C.char
		if localName != "" {
			localNameC = C.CString(localName)
			defer C.free(unsafe.Pointer(localNameC))
		}
		return cGoUnpackString(C.QXmlAttributes_Value4(ptr.Pointer(), C.struct_QtXml_PackedString{data: uriC, len: C.longlong(len(uri))}, C.struct_QtXml_PackedString{data: localNameC, len: C.longlong(len(localName))}))
	}
	return ""
}

func (ptr *QXmlAttributes) Value(index int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlAttributes_Value(ptr.Pointer(), C.int(int32(index))))
	}
	return ""
}

func (ptr *QXmlAttributes) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QXmlAttributes_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QXmlAttributes) Index2(qName core.QLatin1String_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QXmlAttributes_Index2(ptr.Pointer(), core.PointerFromQLatin1String(qName))))
	}
	return 0
}

func (ptr *QXmlAttributes) Index(qName string) int {
	if ptr.Pointer() != nil {
		var qNameC *C.char
		if qName != "" {
			qNameC = C.CString(qName)
			defer C.free(unsafe.Pointer(qNameC))
		}
		return int(int32(C.QXmlAttributes_Index(ptr.Pointer(), C.struct_QtXml_PackedString{data: qNameC, len: C.longlong(len(qName))})))
	}
	return 0
}

func (ptr *QXmlAttributes) Index3(uri string, localPart string) int {
	if ptr.Pointer() != nil {
		var uriC *C.char
		if uri != "" {
			uriC = C.CString(uri)
			defer C.free(unsafe.Pointer(uriC))
		}
		var localPartC *C.char
		if localPart != "" {
			localPartC = C.CString(localPart)
			defer C.free(unsafe.Pointer(localPartC))
		}
		return int(int32(C.QXmlAttributes_Index3(ptr.Pointer(), C.struct_QtXml_PackedString{data: uriC, len: C.longlong(len(uri))}, C.struct_QtXml_PackedString{data: localPartC, len: C.longlong(len(localPart))})))
	}
	return 0
}

func (ptr *QXmlAttributes) Length() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QXmlAttributes_Length(ptr.Pointer())))
	}
	return 0
}

type QXmlContentHandler struct {
	ptr unsafe.Pointer
}

type QXmlContentHandler_ITF interface {
	QXmlContentHandler_PTR() *QXmlContentHandler
}

func (ptr *QXmlContentHandler) QXmlContentHandler_PTR() *QXmlContentHandler {
	return ptr
}

func (ptr *QXmlContentHandler) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlContentHandler) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQXmlContentHandler(ptr QXmlContentHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlContentHandler_PTR().Pointer()
	}
	return nil
}

func NewQXmlContentHandlerFromPointer(ptr unsafe.Pointer) (n *QXmlContentHandler) {
	n = new(QXmlContentHandler)
	n.SetPointer(ptr)
	return
}

//export callbackQXmlContentHandler_Characters
func callbackQXmlContentHandler_Characters(ptr unsafe.Pointer, ch C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "characters"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(ch)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlContentHandler) ConnectCharacters(f func(ch string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "characters"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "characters", func(ch string) bool {
				signal.(func(string) bool)(ch)
				return f(ch)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "characters", f)
		}
	}
}

func (ptr *QXmlContentHandler) DisconnectCharacters() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "characters")
	}
}

func (ptr *QXmlContentHandler) Characters(ch string) bool {
	if ptr.Pointer() != nil {
		var chC *C.char
		if ch != "" {
			chC = C.CString(ch)
			defer C.free(unsafe.Pointer(chC))
		}
		return int8(C.QXmlContentHandler_Characters(ptr.Pointer(), C.struct_QtXml_PackedString{data: chC, len: C.longlong(len(ch))})) != 0
	}
	return false
}

//export callbackQXmlContentHandler_EndDocument
func callbackQXmlContentHandler_EndDocument(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "endDocument"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlContentHandler) ConnectEndDocument(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "endDocument"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "endDocument", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endDocument", f)
		}
	}
}

func (ptr *QXmlContentHandler) DisconnectEndDocument() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "endDocument")
	}
}

func (ptr *QXmlContentHandler) EndDocument() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlContentHandler_EndDocument(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQXmlContentHandler_EndElement
func callbackQXmlContentHandler_EndElement(ptr unsafe.Pointer, namespaceURI C.struct_QtXml_PackedString, localName C.struct_QtXml_PackedString, qName C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "endElement"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, string, string) bool)(cGoUnpackString(namespaceURI), cGoUnpackString(localName), cGoUnpackString(qName)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlContentHandler) ConnectEndElement(f func(namespaceURI string, localName string, qName string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "endElement"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "endElement", func(namespaceURI string, localName string, qName string) bool {
				signal.(func(string, string, string) bool)(namespaceURI, localName, qName)
				return f(namespaceURI, localName, qName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endElement", f)
		}
	}
}

func (ptr *QXmlContentHandler) DisconnectEndElement() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "endElement")
	}
}

func (ptr *QXmlContentHandler) EndElement(namespaceURI string, localName string, qName string) bool {
	if ptr.Pointer() != nil {
		var namespaceURIC *C.char
		if namespaceURI != "" {
			namespaceURIC = C.CString(namespaceURI)
			defer C.free(unsafe.Pointer(namespaceURIC))
		}
		var localNameC *C.char
		if localName != "" {
			localNameC = C.CString(localName)
			defer C.free(unsafe.Pointer(localNameC))
		}
		var qNameC *C.char
		if qName != "" {
			qNameC = C.CString(qName)
			defer C.free(unsafe.Pointer(qNameC))
		}
		return int8(C.QXmlContentHandler_EndElement(ptr.Pointer(), C.struct_QtXml_PackedString{data: namespaceURIC, len: C.longlong(len(namespaceURI))}, C.struct_QtXml_PackedString{data: localNameC, len: C.longlong(len(localName))}, C.struct_QtXml_PackedString{data: qNameC, len: C.longlong(len(qName))})) != 0
	}
	return false
}

//export callbackQXmlContentHandler_EndPrefixMapping
func callbackQXmlContentHandler_EndPrefixMapping(ptr unsafe.Pointer, prefix C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "endPrefixMapping"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(prefix)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlContentHandler) ConnectEndPrefixMapping(f func(prefix string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "endPrefixMapping"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "endPrefixMapping", func(prefix string) bool {
				signal.(func(string) bool)(prefix)
				return f(prefix)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endPrefixMapping", f)
		}
	}
}

func (ptr *QXmlContentHandler) DisconnectEndPrefixMapping() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "endPrefixMapping")
	}
}

func (ptr *QXmlContentHandler) EndPrefixMapping(prefix string) bool {
	if ptr.Pointer() != nil {
		var prefixC *C.char
		if prefix != "" {
			prefixC = C.CString(prefix)
			defer C.free(unsafe.Pointer(prefixC))
		}
		return int8(C.QXmlContentHandler_EndPrefixMapping(ptr.Pointer(), C.struct_QtXml_PackedString{data: prefixC, len: C.longlong(len(prefix))})) != 0
	}
	return false
}

//export callbackQXmlContentHandler_IgnorableWhitespace
func callbackQXmlContentHandler_IgnorableWhitespace(ptr unsafe.Pointer, ch C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "ignorableWhitespace"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(ch)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlContentHandler) ConnectIgnorableWhitespace(f func(ch string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "ignorableWhitespace"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "ignorableWhitespace", func(ch string) bool {
				signal.(func(string) bool)(ch)
				return f(ch)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "ignorableWhitespace", f)
		}
	}
}

func (ptr *QXmlContentHandler) DisconnectIgnorableWhitespace() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "ignorableWhitespace")
	}
}

func (ptr *QXmlContentHandler) IgnorableWhitespace(ch string) bool {
	if ptr.Pointer() != nil {
		var chC *C.char
		if ch != "" {
			chC = C.CString(ch)
			defer C.free(unsafe.Pointer(chC))
		}
		return int8(C.QXmlContentHandler_IgnorableWhitespace(ptr.Pointer(), C.struct_QtXml_PackedString{data: chC, len: C.longlong(len(ch))})) != 0
	}
	return false
}

//export callbackQXmlContentHandler_ProcessingInstruction
func callbackQXmlContentHandler_ProcessingInstruction(ptr unsafe.Pointer, target C.struct_QtXml_PackedString, data C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "processingInstruction"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, string) bool)(cGoUnpackString(target), cGoUnpackString(data)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlContentHandler) ConnectProcessingInstruction(f func(target string, data string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "processingInstruction"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "processingInstruction", func(target string, data string) bool {
				signal.(func(string, string) bool)(target, data)
				return f(target, data)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "processingInstruction", f)
		}
	}
}

func (ptr *QXmlContentHandler) DisconnectProcessingInstruction() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "processingInstruction")
	}
}

func (ptr *QXmlContentHandler) ProcessingInstruction(target string, data string) bool {
	if ptr.Pointer() != nil {
		var targetC *C.char
		if target != "" {
			targetC = C.CString(target)
			defer C.free(unsafe.Pointer(targetC))
		}
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		return int8(C.QXmlContentHandler_ProcessingInstruction(ptr.Pointer(), C.struct_QtXml_PackedString{data: targetC, len: C.longlong(len(target))}, C.struct_QtXml_PackedString{data: dataC, len: C.longlong(len(data))})) != 0
	}
	return false
}

//export callbackQXmlContentHandler_SkippedEntity
func callbackQXmlContentHandler_SkippedEntity(ptr unsafe.Pointer, name C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "skippedEntity"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(name)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlContentHandler) ConnectSkippedEntity(f func(name string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "skippedEntity"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "skippedEntity", func(name string) bool {
				signal.(func(string) bool)(name)
				return f(name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "skippedEntity", f)
		}
	}
}

func (ptr *QXmlContentHandler) DisconnectSkippedEntity() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "skippedEntity")
	}
}

func (ptr *QXmlContentHandler) SkippedEntity(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QXmlContentHandler_SkippedEntity(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

//export callbackQXmlContentHandler_StartDocument
func callbackQXmlContentHandler_StartDocument(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "startDocument"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlContentHandler) ConnectStartDocument(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startDocument"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "startDocument", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startDocument", f)
		}
	}
}

func (ptr *QXmlContentHandler) DisconnectStartDocument() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "startDocument")
	}
}

func (ptr *QXmlContentHandler) StartDocument() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlContentHandler_StartDocument(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQXmlContentHandler_StartElement
func callbackQXmlContentHandler_StartElement(ptr unsafe.Pointer, namespaceURI C.struct_QtXml_PackedString, localName C.struct_QtXml_PackedString, qName C.struct_QtXml_PackedString, atts unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "startElement"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, string, string, *QXmlAttributes) bool)(cGoUnpackString(namespaceURI), cGoUnpackString(localName), cGoUnpackString(qName), NewQXmlAttributesFromPointer(atts)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlContentHandler) ConnectStartElement(f func(namespaceURI string, localName string, qName string, atts *QXmlAttributes) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startElement"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "startElement", func(namespaceURI string, localName string, qName string, atts *QXmlAttributes) bool {
				signal.(func(string, string, string, *QXmlAttributes) bool)(namespaceURI, localName, qName, atts)
				return f(namespaceURI, localName, qName, atts)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startElement", f)
		}
	}
}

func (ptr *QXmlContentHandler) DisconnectStartElement() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "startElement")
	}
}

func (ptr *QXmlContentHandler) StartElement(namespaceURI string, localName string, qName string, atts QXmlAttributes_ITF) bool {
	if ptr.Pointer() != nil {
		var namespaceURIC *C.char
		if namespaceURI != "" {
			namespaceURIC = C.CString(namespaceURI)
			defer C.free(unsafe.Pointer(namespaceURIC))
		}
		var localNameC *C.char
		if localName != "" {
			localNameC = C.CString(localName)
			defer C.free(unsafe.Pointer(localNameC))
		}
		var qNameC *C.char
		if qName != "" {
			qNameC = C.CString(qName)
			defer C.free(unsafe.Pointer(qNameC))
		}
		return int8(C.QXmlContentHandler_StartElement(ptr.Pointer(), C.struct_QtXml_PackedString{data: namespaceURIC, len: C.longlong(len(namespaceURI))}, C.struct_QtXml_PackedString{data: localNameC, len: C.longlong(len(localName))}, C.struct_QtXml_PackedString{data: qNameC, len: C.longlong(len(qName))}, PointerFromQXmlAttributes(atts))) != 0
	}
	return false
}

//export callbackQXmlContentHandler_StartPrefixMapping
func callbackQXmlContentHandler_StartPrefixMapping(ptr unsafe.Pointer, prefix C.struct_QtXml_PackedString, uri C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "startPrefixMapping"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, string) bool)(cGoUnpackString(prefix), cGoUnpackString(uri)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlContentHandler) ConnectStartPrefixMapping(f func(prefix string, uri string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startPrefixMapping"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "startPrefixMapping", func(prefix string, uri string) bool {
				signal.(func(string, string) bool)(prefix, uri)
				return f(prefix, uri)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startPrefixMapping", f)
		}
	}
}

func (ptr *QXmlContentHandler) DisconnectStartPrefixMapping() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "startPrefixMapping")
	}
}

func (ptr *QXmlContentHandler) StartPrefixMapping(prefix string, uri string) bool {
	if ptr.Pointer() != nil {
		var prefixC *C.char
		if prefix != "" {
			prefixC = C.CString(prefix)
			defer C.free(unsafe.Pointer(prefixC))
		}
		var uriC *C.char
		if uri != "" {
			uriC = C.CString(uri)
			defer C.free(unsafe.Pointer(uriC))
		}
		return int8(C.QXmlContentHandler_StartPrefixMapping(ptr.Pointer(), C.struct_QtXml_PackedString{data: prefixC, len: C.longlong(len(prefix))}, C.struct_QtXml_PackedString{data: uriC, len: C.longlong(len(uri))})) != 0
	}
	return false
}

//export callbackQXmlContentHandler_SetDocumentLocator
func callbackQXmlContentHandler_SetDocumentLocator(ptr unsafe.Pointer, locator unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setDocumentLocator"); signal != nil {
		signal.(func(*QXmlLocator))(NewQXmlLocatorFromPointer(locator))
	}

}

func (ptr *QXmlContentHandler) ConnectSetDocumentLocator(f func(locator *QXmlLocator)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setDocumentLocator"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setDocumentLocator", func(locator *QXmlLocator) {
				signal.(func(*QXmlLocator))(locator)
				f(locator)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setDocumentLocator", f)
		}
	}
}

func (ptr *QXmlContentHandler) DisconnectSetDocumentLocator() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setDocumentLocator")
	}
}

func (ptr *QXmlContentHandler) SetDocumentLocator(locator QXmlLocator_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlContentHandler_SetDocumentLocator(ptr.Pointer(), PointerFromQXmlLocator(locator))
	}
}

//export callbackQXmlContentHandler_DestroyQXmlContentHandler
func callbackQXmlContentHandler_DestroyQXmlContentHandler(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QXmlContentHandler"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlContentHandlerFromPointer(ptr).DestroyQXmlContentHandlerDefault()
	}
}

func (ptr *QXmlContentHandler) ConnectDestroyQXmlContentHandler(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QXmlContentHandler"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QXmlContentHandler", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QXmlContentHandler", f)
		}
	}
}

func (ptr *QXmlContentHandler) DisconnectDestroyQXmlContentHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QXmlContentHandler")
	}
}

func (ptr *QXmlContentHandler) DestroyQXmlContentHandler() {
	if ptr.Pointer() != nil {
		C.QXmlContentHandler_DestroyQXmlContentHandler(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXmlContentHandler) DestroyQXmlContentHandlerDefault() {
	if ptr.Pointer() != nil {
		C.QXmlContentHandler_DestroyQXmlContentHandlerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQXmlContentHandler_ErrorString
func callbackQXmlContentHandler_ErrorString(ptr unsafe.Pointer) C.struct_QtXml_PackedString {
	if signal := qt.GetSignal(ptr, "errorString"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtXml_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtXml_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QXmlContentHandler) ConnectErrorString(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "errorString"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "errorString", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "errorString", f)
		}
	}
}

func (ptr *QXmlContentHandler) DisconnectErrorString() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "errorString")
	}
}

func (ptr *QXmlContentHandler) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlContentHandler_ErrorString(ptr.Pointer()))
	}
	return ""
}

type QXmlDTDHandler struct {
	ptr unsafe.Pointer
}

type QXmlDTDHandler_ITF interface {
	QXmlDTDHandler_PTR() *QXmlDTDHandler
}

func (ptr *QXmlDTDHandler) QXmlDTDHandler_PTR() *QXmlDTDHandler {
	return ptr
}

func (ptr *QXmlDTDHandler) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlDTDHandler) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQXmlDTDHandler(ptr QXmlDTDHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlDTDHandler_PTR().Pointer()
	}
	return nil
}

func NewQXmlDTDHandlerFromPointer(ptr unsafe.Pointer) (n *QXmlDTDHandler) {
	n = new(QXmlDTDHandler)
	n.SetPointer(ptr)
	return
}

//export callbackQXmlDTDHandler_NotationDecl
func callbackQXmlDTDHandler_NotationDecl(ptr unsafe.Pointer, name C.struct_QtXml_PackedString, publicId C.struct_QtXml_PackedString, systemId C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "notationDecl"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, string, string) bool)(cGoUnpackString(name), cGoUnpackString(publicId), cGoUnpackString(systemId)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlDTDHandler) ConnectNotationDecl(f func(name string, publicId string, systemId string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "notationDecl"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "notationDecl", func(name string, publicId string, systemId string) bool {
				signal.(func(string, string, string) bool)(name, publicId, systemId)
				return f(name, publicId, systemId)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "notationDecl", f)
		}
	}
}

func (ptr *QXmlDTDHandler) DisconnectNotationDecl() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "notationDecl")
	}
}

func (ptr *QXmlDTDHandler) NotationDecl(name string, publicId string, systemId string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var publicIdC *C.char
		if publicId != "" {
			publicIdC = C.CString(publicId)
			defer C.free(unsafe.Pointer(publicIdC))
		}
		var systemIdC *C.char
		if systemId != "" {
			systemIdC = C.CString(systemId)
			defer C.free(unsafe.Pointer(systemIdC))
		}
		return int8(C.QXmlDTDHandler_NotationDecl(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtXml_PackedString{data: publicIdC, len: C.longlong(len(publicId))}, C.struct_QtXml_PackedString{data: systemIdC, len: C.longlong(len(systemId))})) != 0
	}
	return false
}

//export callbackQXmlDTDHandler_UnparsedEntityDecl
func callbackQXmlDTDHandler_UnparsedEntityDecl(ptr unsafe.Pointer, name C.struct_QtXml_PackedString, publicId C.struct_QtXml_PackedString, systemId C.struct_QtXml_PackedString, notationName C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "unparsedEntityDecl"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, string, string, string) bool)(cGoUnpackString(name), cGoUnpackString(publicId), cGoUnpackString(systemId), cGoUnpackString(notationName)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlDTDHandler) ConnectUnparsedEntityDecl(f func(name string, publicId string, systemId string, notationName string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "unparsedEntityDecl"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "unparsedEntityDecl", func(name string, publicId string, systemId string, notationName string) bool {
				signal.(func(string, string, string, string) bool)(name, publicId, systemId, notationName)
				return f(name, publicId, systemId, notationName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "unparsedEntityDecl", f)
		}
	}
}

func (ptr *QXmlDTDHandler) DisconnectUnparsedEntityDecl() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "unparsedEntityDecl")
	}
}

func (ptr *QXmlDTDHandler) UnparsedEntityDecl(name string, publicId string, systemId string, notationName string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var publicIdC *C.char
		if publicId != "" {
			publicIdC = C.CString(publicId)
			defer C.free(unsafe.Pointer(publicIdC))
		}
		var systemIdC *C.char
		if systemId != "" {
			systemIdC = C.CString(systemId)
			defer C.free(unsafe.Pointer(systemIdC))
		}
		var notationNameC *C.char
		if notationName != "" {
			notationNameC = C.CString(notationName)
			defer C.free(unsafe.Pointer(notationNameC))
		}
		return int8(C.QXmlDTDHandler_UnparsedEntityDecl(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtXml_PackedString{data: publicIdC, len: C.longlong(len(publicId))}, C.struct_QtXml_PackedString{data: systemIdC, len: C.longlong(len(systemId))}, C.struct_QtXml_PackedString{data: notationNameC, len: C.longlong(len(notationName))})) != 0
	}
	return false
}

//export callbackQXmlDTDHandler_DestroyQXmlDTDHandler
func callbackQXmlDTDHandler_DestroyQXmlDTDHandler(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QXmlDTDHandler"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlDTDHandlerFromPointer(ptr).DestroyQXmlDTDHandlerDefault()
	}
}

func (ptr *QXmlDTDHandler) ConnectDestroyQXmlDTDHandler(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QXmlDTDHandler"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QXmlDTDHandler", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QXmlDTDHandler", f)
		}
	}
}

func (ptr *QXmlDTDHandler) DisconnectDestroyQXmlDTDHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QXmlDTDHandler")
	}
}

func (ptr *QXmlDTDHandler) DestroyQXmlDTDHandler() {
	if ptr.Pointer() != nil {
		C.QXmlDTDHandler_DestroyQXmlDTDHandler(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXmlDTDHandler) DestroyQXmlDTDHandlerDefault() {
	if ptr.Pointer() != nil {
		C.QXmlDTDHandler_DestroyQXmlDTDHandlerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQXmlDTDHandler_ErrorString
func callbackQXmlDTDHandler_ErrorString(ptr unsafe.Pointer) C.struct_QtXml_PackedString {
	if signal := qt.GetSignal(ptr, "errorString"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtXml_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtXml_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QXmlDTDHandler) ConnectErrorString(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "errorString"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "errorString", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "errorString", f)
		}
	}
}

func (ptr *QXmlDTDHandler) DisconnectErrorString() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "errorString")
	}
}

func (ptr *QXmlDTDHandler) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlDTDHandler_ErrorString(ptr.Pointer()))
	}
	return ""
}

type QXmlDeclHandler struct {
	ptr unsafe.Pointer
}

type QXmlDeclHandler_ITF interface {
	QXmlDeclHandler_PTR() *QXmlDeclHandler
}

func (ptr *QXmlDeclHandler) QXmlDeclHandler_PTR() *QXmlDeclHandler {
	return ptr
}

func (ptr *QXmlDeclHandler) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlDeclHandler) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQXmlDeclHandler(ptr QXmlDeclHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlDeclHandler_PTR().Pointer()
	}
	return nil
}

func NewQXmlDeclHandlerFromPointer(ptr unsafe.Pointer) (n *QXmlDeclHandler) {
	n = new(QXmlDeclHandler)
	n.SetPointer(ptr)
	return
}

//export callbackQXmlDeclHandler_AttributeDecl
func callbackQXmlDeclHandler_AttributeDecl(ptr unsafe.Pointer, eName C.struct_QtXml_PackedString, aName C.struct_QtXml_PackedString, ty C.struct_QtXml_PackedString, valueDefault C.struct_QtXml_PackedString, value C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "attributeDecl"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, string, string, string, string) bool)(cGoUnpackString(eName), cGoUnpackString(aName), cGoUnpackString(ty), cGoUnpackString(valueDefault), cGoUnpackString(value)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlDeclHandler) ConnectAttributeDecl(f func(eName string, aName string, ty string, valueDefault string, value string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "attributeDecl"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "attributeDecl", func(eName string, aName string, ty string, valueDefault string, value string) bool {
				signal.(func(string, string, string, string, string) bool)(eName, aName, ty, valueDefault, value)
				return f(eName, aName, ty, valueDefault, value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "attributeDecl", f)
		}
	}
}

func (ptr *QXmlDeclHandler) DisconnectAttributeDecl() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "attributeDecl")
	}
}

func (ptr *QXmlDeclHandler) AttributeDecl(eName string, aName string, ty string, valueDefault string, value string) bool {
	if ptr.Pointer() != nil {
		var eNameC *C.char
		if eName != "" {
			eNameC = C.CString(eName)
			defer C.free(unsafe.Pointer(eNameC))
		}
		var aNameC *C.char
		if aName != "" {
			aNameC = C.CString(aName)
			defer C.free(unsafe.Pointer(aNameC))
		}
		var tyC *C.char
		if ty != "" {
			tyC = C.CString(ty)
			defer C.free(unsafe.Pointer(tyC))
		}
		var valueDefaultC *C.char
		if valueDefault != "" {
			valueDefaultC = C.CString(valueDefault)
			defer C.free(unsafe.Pointer(valueDefaultC))
		}
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		return int8(C.QXmlDeclHandler_AttributeDecl(ptr.Pointer(), C.struct_QtXml_PackedString{data: eNameC, len: C.longlong(len(eName))}, C.struct_QtXml_PackedString{data: aNameC, len: C.longlong(len(aName))}, C.struct_QtXml_PackedString{data: tyC, len: C.longlong(len(ty))}, C.struct_QtXml_PackedString{data: valueDefaultC, len: C.longlong(len(valueDefault))}, C.struct_QtXml_PackedString{data: valueC, len: C.longlong(len(value))})) != 0
	}
	return false
}

//export callbackQXmlDeclHandler_ExternalEntityDecl
func callbackQXmlDeclHandler_ExternalEntityDecl(ptr unsafe.Pointer, name C.struct_QtXml_PackedString, publicId C.struct_QtXml_PackedString, systemId C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "externalEntityDecl"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, string, string) bool)(cGoUnpackString(name), cGoUnpackString(publicId), cGoUnpackString(systemId)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlDeclHandler) ConnectExternalEntityDecl(f func(name string, publicId string, systemId string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "externalEntityDecl"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "externalEntityDecl", func(name string, publicId string, systemId string) bool {
				signal.(func(string, string, string) bool)(name, publicId, systemId)
				return f(name, publicId, systemId)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "externalEntityDecl", f)
		}
	}
}

func (ptr *QXmlDeclHandler) DisconnectExternalEntityDecl() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "externalEntityDecl")
	}
}

func (ptr *QXmlDeclHandler) ExternalEntityDecl(name string, publicId string, systemId string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var publicIdC *C.char
		if publicId != "" {
			publicIdC = C.CString(publicId)
			defer C.free(unsafe.Pointer(publicIdC))
		}
		var systemIdC *C.char
		if systemId != "" {
			systemIdC = C.CString(systemId)
			defer C.free(unsafe.Pointer(systemIdC))
		}
		return int8(C.QXmlDeclHandler_ExternalEntityDecl(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtXml_PackedString{data: publicIdC, len: C.longlong(len(publicId))}, C.struct_QtXml_PackedString{data: systemIdC, len: C.longlong(len(systemId))})) != 0
	}
	return false
}

//export callbackQXmlDeclHandler_InternalEntityDecl
func callbackQXmlDeclHandler_InternalEntityDecl(ptr unsafe.Pointer, name C.struct_QtXml_PackedString, value C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "internalEntityDecl"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, string) bool)(cGoUnpackString(name), cGoUnpackString(value)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlDeclHandler) ConnectInternalEntityDecl(f func(name string, value string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "internalEntityDecl"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "internalEntityDecl", func(name string, value string) bool {
				signal.(func(string, string) bool)(name, value)
				return f(name, value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "internalEntityDecl", f)
		}
	}
}

func (ptr *QXmlDeclHandler) DisconnectInternalEntityDecl() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "internalEntityDecl")
	}
}

func (ptr *QXmlDeclHandler) InternalEntityDecl(name string, value string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		return int8(C.QXmlDeclHandler_InternalEntityDecl(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtXml_PackedString{data: valueC, len: C.longlong(len(value))})) != 0
	}
	return false
}

//export callbackQXmlDeclHandler_DestroyQXmlDeclHandler
func callbackQXmlDeclHandler_DestroyQXmlDeclHandler(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QXmlDeclHandler"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlDeclHandlerFromPointer(ptr).DestroyQXmlDeclHandlerDefault()
	}
}

func (ptr *QXmlDeclHandler) ConnectDestroyQXmlDeclHandler(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QXmlDeclHandler"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QXmlDeclHandler", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QXmlDeclHandler", f)
		}
	}
}

func (ptr *QXmlDeclHandler) DisconnectDestroyQXmlDeclHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QXmlDeclHandler")
	}
}

func (ptr *QXmlDeclHandler) DestroyQXmlDeclHandler() {
	if ptr.Pointer() != nil {
		C.QXmlDeclHandler_DestroyQXmlDeclHandler(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXmlDeclHandler) DestroyQXmlDeclHandlerDefault() {
	if ptr.Pointer() != nil {
		C.QXmlDeclHandler_DestroyQXmlDeclHandlerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQXmlDeclHandler_ErrorString
func callbackQXmlDeclHandler_ErrorString(ptr unsafe.Pointer) C.struct_QtXml_PackedString {
	if signal := qt.GetSignal(ptr, "errorString"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtXml_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtXml_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QXmlDeclHandler) ConnectErrorString(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "errorString"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "errorString", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "errorString", f)
		}
	}
}

func (ptr *QXmlDeclHandler) DisconnectErrorString() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "errorString")
	}
}

func (ptr *QXmlDeclHandler) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlDeclHandler_ErrorString(ptr.Pointer()))
	}
	return ""
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

func NewQXmlDefaultHandlerFromPointer(ptr unsafe.Pointer) (n *QXmlDefaultHandler) {
	n = new(QXmlDefaultHandler)
	n.SetPointer(ptr)
	return
}
func NewQXmlDefaultHandler() *QXmlDefaultHandler {
	return NewQXmlDefaultHandlerFromPointer(C.QXmlDefaultHandler_NewQXmlDefaultHandler())
}

//export callbackQXmlDefaultHandler_AttributeDecl
func callbackQXmlDefaultHandler_AttributeDecl(ptr unsafe.Pointer, eName C.struct_QtXml_PackedString, aName C.struct_QtXml_PackedString, ty C.struct_QtXml_PackedString, valueDefault C.struct_QtXml_PackedString, value C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "attributeDecl"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, string, string, string, string) bool)(cGoUnpackString(eName), cGoUnpackString(aName), cGoUnpackString(ty), cGoUnpackString(valueDefault), cGoUnpackString(value)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).AttributeDeclDefault(cGoUnpackString(eName), cGoUnpackString(aName), cGoUnpackString(ty), cGoUnpackString(valueDefault), cGoUnpackString(value)))))
}

func (ptr *QXmlDefaultHandler) ConnectAttributeDecl(f func(eName string, aName string, ty string, valueDefault string, value string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "attributeDecl"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "attributeDecl", func(eName string, aName string, ty string, valueDefault string, value string) bool {
				signal.(func(string, string, string, string, string) bool)(eName, aName, ty, valueDefault, value)
				return f(eName, aName, ty, valueDefault, value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "attributeDecl", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectAttributeDecl() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "attributeDecl")
	}
}

func (ptr *QXmlDefaultHandler) AttributeDecl(eName string, aName string, ty string, valueDefault string, value string) bool {
	if ptr.Pointer() != nil {
		var eNameC *C.char
		if eName != "" {
			eNameC = C.CString(eName)
			defer C.free(unsafe.Pointer(eNameC))
		}
		var aNameC *C.char
		if aName != "" {
			aNameC = C.CString(aName)
			defer C.free(unsafe.Pointer(aNameC))
		}
		var tyC *C.char
		if ty != "" {
			tyC = C.CString(ty)
			defer C.free(unsafe.Pointer(tyC))
		}
		var valueDefaultC *C.char
		if valueDefault != "" {
			valueDefaultC = C.CString(valueDefault)
			defer C.free(unsafe.Pointer(valueDefaultC))
		}
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		return int8(C.QXmlDefaultHandler_AttributeDecl(ptr.Pointer(), C.struct_QtXml_PackedString{data: eNameC, len: C.longlong(len(eName))}, C.struct_QtXml_PackedString{data: aNameC, len: C.longlong(len(aName))}, C.struct_QtXml_PackedString{data: tyC, len: C.longlong(len(ty))}, C.struct_QtXml_PackedString{data: valueDefaultC, len: C.longlong(len(valueDefault))}, C.struct_QtXml_PackedString{data: valueC, len: C.longlong(len(value))})) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) AttributeDeclDefault(eName string, aName string, ty string, valueDefault string, value string) bool {
	if ptr.Pointer() != nil {
		var eNameC *C.char
		if eName != "" {
			eNameC = C.CString(eName)
			defer C.free(unsafe.Pointer(eNameC))
		}
		var aNameC *C.char
		if aName != "" {
			aNameC = C.CString(aName)
			defer C.free(unsafe.Pointer(aNameC))
		}
		var tyC *C.char
		if ty != "" {
			tyC = C.CString(ty)
			defer C.free(unsafe.Pointer(tyC))
		}
		var valueDefaultC *C.char
		if valueDefault != "" {
			valueDefaultC = C.CString(valueDefault)
			defer C.free(unsafe.Pointer(valueDefaultC))
		}
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		return int8(C.QXmlDefaultHandler_AttributeDeclDefault(ptr.Pointer(), C.struct_QtXml_PackedString{data: eNameC, len: C.longlong(len(eName))}, C.struct_QtXml_PackedString{data: aNameC, len: C.longlong(len(aName))}, C.struct_QtXml_PackedString{data: tyC, len: C.longlong(len(ty))}, C.struct_QtXml_PackedString{data: valueDefaultC, len: C.longlong(len(valueDefault))}, C.struct_QtXml_PackedString{data: valueC, len: C.longlong(len(value))})) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_Characters
func callbackQXmlDefaultHandler_Characters(ptr unsafe.Pointer, ch C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "characters"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(ch)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).CharactersDefault(cGoUnpackString(ch)))))
}

func (ptr *QXmlDefaultHandler) ConnectCharacters(f func(ch string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "characters"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "characters", func(ch string) bool {
				signal.(func(string) bool)(ch)
				return f(ch)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "characters", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectCharacters() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "characters")
	}
}

func (ptr *QXmlDefaultHandler) Characters(ch string) bool {
	if ptr.Pointer() != nil {
		var chC *C.char
		if ch != "" {
			chC = C.CString(ch)
			defer C.free(unsafe.Pointer(chC))
		}
		return int8(C.QXmlDefaultHandler_Characters(ptr.Pointer(), C.struct_QtXml_PackedString{data: chC, len: C.longlong(len(ch))})) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) CharactersDefault(ch string) bool {
	if ptr.Pointer() != nil {
		var chC *C.char
		if ch != "" {
			chC = C.CString(ch)
			defer C.free(unsafe.Pointer(chC))
		}
		return int8(C.QXmlDefaultHandler_CharactersDefault(ptr.Pointer(), C.struct_QtXml_PackedString{data: chC, len: C.longlong(len(ch))})) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_Comment
func callbackQXmlDefaultHandler_Comment(ptr unsafe.Pointer, ch C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "comment"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(ch)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).CommentDefault(cGoUnpackString(ch)))))
}

func (ptr *QXmlDefaultHandler) ConnectComment(f func(ch string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "comment"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "comment", func(ch string) bool {
				signal.(func(string) bool)(ch)
				return f(ch)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "comment", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectComment() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "comment")
	}
}

func (ptr *QXmlDefaultHandler) Comment(ch string) bool {
	if ptr.Pointer() != nil {
		var chC *C.char
		if ch != "" {
			chC = C.CString(ch)
			defer C.free(unsafe.Pointer(chC))
		}
		return int8(C.QXmlDefaultHandler_Comment(ptr.Pointer(), C.struct_QtXml_PackedString{data: chC, len: C.longlong(len(ch))})) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) CommentDefault(ch string) bool {
	if ptr.Pointer() != nil {
		var chC *C.char
		if ch != "" {
			chC = C.CString(ch)
			defer C.free(unsafe.Pointer(chC))
		}
		return int8(C.QXmlDefaultHandler_CommentDefault(ptr.Pointer(), C.struct_QtXml_PackedString{data: chC, len: C.longlong(len(ch))})) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_EndCDATA
func callbackQXmlDefaultHandler_EndCDATA(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "endCDATA"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).EndCDATADefault())))
}

func (ptr *QXmlDefaultHandler) ConnectEndCDATA(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "endCDATA"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "endCDATA", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endCDATA", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectEndCDATA() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "endCDATA")
	}
}

func (ptr *QXmlDefaultHandler) EndCDATA() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlDefaultHandler_EndCDATA(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) EndCDATADefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlDefaultHandler_EndCDATADefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_EndDTD
func callbackQXmlDefaultHandler_EndDTD(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "endDTD"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).EndDTDDefault())))
}

func (ptr *QXmlDefaultHandler) ConnectEndDTD(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "endDTD"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "endDTD", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endDTD", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectEndDTD() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "endDTD")
	}
}

func (ptr *QXmlDefaultHandler) EndDTD() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlDefaultHandler_EndDTD(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) EndDTDDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlDefaultHandler_EndDTDDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_EndDocument
func callbackQXmlDefaultHandler_EndDocument(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "endDocument"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).EndDocumentDefault())))
}

func (ptr *QXmlDefaultHandler) ConnectEndDocument(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "endDocument"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "endDocument", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endDocument", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectEndDocument() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "endDocument")
	}
}

func (ptr *QXmlDefaultHandler) EndDocument() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlDefaultHandler_EndDocument(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) EndDocumentDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlDefaultHandler_EndDocumentDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_EndElement
func callbackQXmlDefaultHandler_EndElement(ptr unsafe.Pointer, namespaceURI C.struct_QtXml_PackedString, localName C.struct_QtXml_PackedString, qName C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "endElement"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, string, string) bool)(cGoUnpackString(namespaceURI), cGoUnpackString(localName), cGoUnpackString(qName)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).EndElementDefault(cGoUnpackString(namespaceURI), cGoUnpackString(localName), cGoUnpackString(qName)))))
}

func (ptr *QXmlDefaultHandler) ConnectEndElement(f func(namespaceURI string, localName string, qName string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "endElement"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "endElement", func(namespaceURI string, localName string, qName string) bool {
				signal.(func(string, string, string) bool)(namespaceURI, localName, qName)
				return f(namespaceURI, localName, qName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endElement", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectEndElement() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "endElement")
	}
}

func (ptr *QXmlDefaultHandler) EndElement(namespaceURI string, localName string, qName string) bool {
	if ptr.Pointer() != nil {
		var namespaceURIC *C.char
		if namespaceURI != "" {
			namespaceURIC = C.CString(namespaceURI)
			defer C.free(unsafe.Pointer(namespaceURIC))
		}
		var localNameC *C.char
		if localName != "" {
			localNameC = C.CString(localName)
			defer C.free(unsafe.Pointer(localNameC))
		}
		var qNameC *C.char
		if qName != "" {
			qNameC = C.CString(qName)
			defer C.free(unsafe.Pointer(qNameC))
		}
		return int8(C.QXmlDefaultHandler_EndElement(ptr.Pointer(), C.struct_QtXml_PackedString{data: namespaceURIC, len: C.longlong(len(namespaceURI))}, C.struct_QtXml_PackedString{data: localNameC, len: C.longlong(len(localName))}, C.struct_QtXml_PackedString{data: qNameC, len: C.longlong(len(qName))})) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) EndElementDefault(namespaceURI string, localName string, qName string) bool {
	if ptr.Pointer() != nil {
		var namespaceURIC *C.char
		if namespaceURI != "" {
			namespaceURIC = C.CString(namespaceURI)
			defer C.free(unsafe.Pointer(namespaceURIC))
		}
		var localNameC *C.char
		if localName != "" {
			localNameC = C.CString(localName)
			defer C.free(unsafe.Pointer(localNameC))
		}
		var qNameC *C.char
		if qName != "" {
			qNameC = C.CString(qName)
			defer C.free(unsafe.Pointer(qNameC))
		}
		return int8(C.QXmlDefaultHandler_EndElementDefault(ptr.Pointer(), C.struct_QtXml_PackedString{data: namespaceURIC, len: C.longlong(len(namespaceURI))}, C.struct_QtXml_PackedString{data: localNameC, len: C.longlong(len(localName))}, C.struct_QtXml_PackedString{data: qNameC, len: C.longlong(len(qName))})) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_EndEntity
func callbackQXmlDefaultHandler_EndEntity(ptr unsafe.Pointer, name C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "endEntity"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(name)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).EndEntityDefault(cGoUnpackString(name)))))
}

func (ptr *QXmlDefaultHandler) ConnectEndEntity(f func(name string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "endEntity"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "endEntity", func(name string) bool {
				signal.(func(string) bool)(name)
				return f(name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endEntity", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectEndEntity() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "endEntity")
	}
}

func (ptr *QXmlDefaultHandler) EndEntity(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QXmlDefaultHandler_EndEntity(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) EndEntityDefault(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QXmlDefaultHandler_EndEntityDefault(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_EndPrefixMapping
func callbackQXmlDefaultHandler_EndPrefixMapping(ptr unsafe.Pointer, prefix C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "endPrefixMapping"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(prefix)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).EndPrefixMappingDefault(cGoUnpackString(prefix)))))
}

func (ptr *QXmlDefaultHandler) ConnectEndPrefixMapping(f func(prefix string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "endPrefixMapping"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "endPrefixMapping", func(prefix string) bool {
				signal.(func(string) bool)(prefix)
				return f(prefix)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endPrefixMapping", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectEndPrefixMapping() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "endPrefixMapping")
	}
}

func (ptr *QXmlDefaultHandler) EndPrefixMapping(prefix string) bool {
	if ptr.Pointer() != nil {
		var prefixC *C.char
		if prefix != "" {
			prefixC = C.CString(prefix)
			defer C.free(unsafe.Pointer(prefixC))
		}
		return int8(C.QXmlDefaultHandler_EndPrefixMapping(ptr.Pointer(), C.struct_QtXml_PackedString{data: prefixC, len: C.longlong(len(prefix))})) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) EndPrefixMappingDefault(prefix string) bool {
	if ptr.Pointer() != nil {
		var prefixC *C.char
		if prefix != "" {
			prefixC = C.CString(prefix)
			defer C.free(unsafe.Pointer(prefixC))
		}
		return int8(C.QXmlDefaultHandler_EndPrefixMappingDefault(ptr.Pointer(), C.struct_QtXml_PackedString{data: prefixC, len: C.longlong(len(prefix))})) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_Error
func callbackQXmlDefaultHandler_Error(ptr unsafe.Pointer, exception unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "error"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QXmlParseException) bool)(NewQXmlParseExceptionFromPointer(exception)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).ErrorDefault(NewQXmlParseExceptionFromPointer(exception)))))
}

func (ptr *QXmlDefaultHandler) ConnectError(f func(exception *QXmlParseException) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "error"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "error", func(exception *QXmlParseException) bool {
				signal.(func(*QXmlParseException) bool)(exception)
				return f(exception)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectError() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "error")
	}
}

func (ptr *QXmlDefaultHandler) Error(exception QXmlParseException_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlDefaultHandler_Error(ptr.Pointer(), PointerFromQXmlParseException(exception))) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) ErrorDefault(exception QXmlParseException_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlDefaultHandler_ErrorDefault(ptr.Pointer(), PointerFromQXmlParseException(exception))) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_ExternalEntityDecl
func callbackQXmlDefaultHandler_ExternalEntityDecl(ptr unsafe.Pointer, name C.struct_QtXml_PackedString, publicId C.struct_QtXml_PackedString, systemId C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "externalEntityDecl"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, string, string) bool)(cGoUnpackString(name), cGoUnpackString(publicId), cGoUnpackString(systemId)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).ExternalEntityDeclDefault(cGoUnpackString(name), cGoUnpackString(publicId), cGoUnpackString(systemId)))))
}

func (ptr *QXmlDefaultHandler) ConnectExternalEntityDecl(f func(name string, publicId string, systemId string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "externalEntityDecl"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "externalEntityDecl", func(name string, publicId string, systemId string) bool {
				signal.(func(string, string, string) bool)(name, publicId, systemId)
				return f(name, publicId, systemId)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "externalEntityDecl", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectExternalEntityDecl() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "externalEntityDecl")
	}
}

func (ptr *QXmlDefaultHandler) ExternalEntityDecl(name string, publicId string, systemId string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var publicIdC *C.char
		if publicId != "" {
			publicIdC = C.CString(publicId)
			defer C.free(unsafe.Pointer(publicIdC))
		}
		var systemIdC *C.char
		if systemId != "" {
			systemIdC = C.CString(systemId)
			defer C.free(unsafe.Pointer(systemIdC))
		}
		return int8(C.QXmlDefaultHandler_ExternalEntityDecl(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtXml_PackedString{data: publicIdC, len: C.longlong(len(publicId))}, C.struct_QtXml_PackedString{data: systemIdC, len: C.longlong(len(systemId))})) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) ExternalEntityDeclDefault(name string, publicId string, systemId string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var publicIdC *C.char
		if publicId != "" {
			publicIdC = C.CString(publicId)
			defer C.free(unsafe.Pointer(publicIdC))
		}
		var systemIdC *C.char
		if systemId != "" {
			systemIdC = C.CString(systemId)
			defer C.free(unsafe.Pointer(systemIdC))
		}
		return int8(C.QXmlDefaultHandler_ExternalEntityDeclDefault(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtXml_PackedString{data: publicIdC, len: C.longlong(len(publicId))}, C.struct_QtXml_PackedString{data: systemIdC, len: C.longlong(len(systemId))})) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_FatalError
func callbackQXmlDefaultHandler_FatalError(ptr unsafe.Pointer, exception unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "fatalError"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QXmlParseException) bool)(NewQXmlParseExceptionFromPointer(exception)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).FatalErrorDefault(NewQXmlParseExceptionFromPointer(exception)))))
}

func (ptr *QXmlDefaultHandler) ConnectFatalError(f func(exception *QXmlParseException) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "fatalError"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fatalError", func(exception *QXmlParseException) bool {
				signal.(func(*QXmlParseException) bool)(exception)
				return f(exception)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fatalError", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectFatalError() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "fatalError")
	}
}

func (ptr *QXmlDefaultHandler) FatalError(exception QXmlParseException_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlDefaultHandler_FatalError(ptr.Pointer(), PointerFromQXmlParseException(exception))) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) FatalErrorDefault(exception QXmlParseException_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlDefaultHandler_FatalErrorDefault(ptr.Pointer(), PointerFromQXmlParseException(exception))) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_IgnorableWhitespace
func callbackQXmlDefaultHandler_IgnorableWhitespace(ptr unsafe.Pointer, ch C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "ignorableWhitespace"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(ch)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).IgnorableWhitespaceDefault(cGoUnpackString(ch)))))
}

func (ptr *QXmlDefaultHandler) ConnectIgnorableWhitespace(f func(ch string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "ignorableWhitespace"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "ignorableWhitespace", func(ch string) bool {
				signal.(func(string) bool)(ch)
				return f(ch)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "ignorableWhitespace", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectIgnorableWhitespace() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "ignorableWhitespace")
	}
}

func (ptr *QXmlDefaultHandler) IgnorableWhitespace(ch string) bool {
	if ptr.Pointer() != nil {
		var chC *C.char
		if ch != "" {
			chC = C.CString(ch)
			defer C.free(unsafe.Pointer(chC))
		}
		return int8(C.QXmlDefaultHandler_IgnorableWhitespace(ptr.Pointer(), C.struct_QtXml_PackedString{data: chC, len: C.longlong(len(ch))})) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) IgnorableWhitespaceDefault(ch string) bool {
	if ptr.Pointer() != nil {
		var chC *C.char
		if ch != "" {
			chC = C.CString(ch)
			defer C.free(unsafe.Pointer(chC))
		}
		return int8(C.QXmlDefaultHandler_IgnorableWhitespaceDefault(ptr.Pointer(), C.struct_QtXml_PackedString{data: chC, len: C.longlong(len(ch))})) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_InternalEntityDecl
func callbackQXmlDefaultHandler_InternalEntityDecl(ptr unsafe.Pointer, name C.struct_QtXml_PackedString, value C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "internalEntityDecl"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, string) bool)(cGoUnpackString(name), cGoUnpackString(value)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).InternalEntityDeclDefault(cGoUnpackString(name), cGoUnpackString(value)))))
}

func (ptr *QXmlDefaultHandler) ConnectInternalEntityDecl(f func(name string, value string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "internalEntityDecl"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "internalEntityDecl", func(name string, value string) bool {
				signal.(func(string, string) bool)(name, value)
				return f(name, value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "internalEntityDecl", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectInternalEntityDecl() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "internalEntityDecl")
	}
}

func (ptr *QXmlDefaultHandler) InternalEntityDecl(name string, value string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		return int8(C.QXmlDefaultHandler_InternalEntityDecl(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtXml_PackedString{data: valueC, len: C.longlong(len(value))})) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) InternalEntityDeclDefault(name string, value string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		return int8(C.QXmlDefaultHandler_InternalEntityDeclDefault(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtXml_PackedString{data: valueC, len: C.longlong(len(value))})) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_NotationDecl
func callbackQXmlDefaultHandler_NotationDecl(ptr unsafe.Pointer, name C.struct_QtXml_PackedString, publicId C.struct_QtXml_PackedString, systemId C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "notationDecl"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, string, string) bool)(cGoUnpackString(name), cGoUnpackString(publicId), cGoUnpackString(systemId)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).NotationDeclDefault(cGoUnpackString(name), cGoUnpackString(publicId), cGoUnpackString(systemId)))))
}

func (ptr *QXmlDefaultHandler) ConnectNotationDecl(f func(name string, publicId string, systemId string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "notationDecl"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "notationDecl", func(name string, publicId string, systemId string) bool {
				signal.(func(string, string, string) bool)(name, publicId, systemId)
				return f(name, publicId, systemId)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "notationDecl", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectNotationDecl() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "notationDecl")
	}
}

func (ptr *QXmlDefaultHandler) NotationDecl(name string, publicId string, systemId string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var publicIdC *C.char
		if publicId != "" {
			publicIdC = C.CString(publicId)
			defer C.free(unsafe.Pointer(publicIdC))
		}
		var systemIdC *C.char
		if systemId != "" {
			systemIdC = C.CString(systemId)
			defer C.free(unsafe.Pointer(systemIdC))
		}
		return int8(C.QXmlDefaultHandler_NotationDecl(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtXml_PackedString{data: publicIdC, len: C.longlong(len(publicId))}, C.struct_QtXml_PackedString{data: systemIdC, len: C.longlong(len(systemId))})) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) NotationDeclDefault(name string, publicId string, systemId string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var publicIdC *C.char
		if publicId != "" {
			publicIdC = C.CString(publicId)
			defer C.free(unsafe.Pointer(publicIdC))
		}
		var systemIdC *C.char
		if systemId != "" {
			systemIdC = C.CString(systemId)
			defer C.free(unsafe.Pointer(systemIdC))
		}
		return int8(C.QXmlDefaultHandler_NotationDeclDefault(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtXml_PackedString{data: publicIdC, len: C.longlong(len(publicId))}, C.struct_QtXml_PackedString{data: systemIdC, len: C.longlong(len(systemId))})) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_ProcessingInstruction
func callbackQXmlDefaultHandler_ProcessingInstruction(ptr unsafe.Pointer, target C.struct_QtXml_PackedString, data C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "processingInstruction"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, string) bool)(cGoUnpackString(target), cGoUnpackString(data)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).ProcessingInstructionDefault(cGoUnpackString(target), cGoUnpackString(data)))))
}

func (ptr *QXmlDefaultHandler) ConnectProcessingInstruction(f func(target string, data string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "processingInstruction"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "processingInstruction", func(target string, data string) bool {
				signal.(func(string, string) bool)(target, data)
				return f(target, data)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "processingInstruction", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectProcessingInstruction() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "processingInstruction")
	}
}

func (ptr *QXmlDefaultHandler) ProcessingInstruction(target string, data string) bool {
	if ptr.Pointer() != nil {
		var targetC *C.char
		if target != "" {
			targetC = C.CString(target)
			defer C.free(unsafe.Pointer(targetC))
		}
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		return int8(C.QXmlDefaultHandler_ProcessingInstruction(ptr.Pointer(), C.struct_QtXml_PackedString{data: targetC, len: C.longlong(len(target))}, C.struct_QtXml_PackedString{data: dataC, len: C.longlong(len(data))})) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) ProcessingInstructionDefault(target string, data string) bool {
	if ptr.Pointer() != nil {
		var targetC *C.char
		if target != "" {
			targetC = C.CString(target)
			defer C.free(unsafe.Pointer(targetC))
		}
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		return int8(C.QXmlDefaultHandler_ProcessingInstructionDefault(ptr.Pointer(), C.struct_QtXml_PackedString{data: targetC, len: C.longlong(len(target))}, C.struct_QtXml_PackedString{data: dataC, len: C.longlong(len(data))})) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_SkippedEntity
func callbackQXmlDefaultHandler_SkippedEntity(ptr unsafe.Pointer, name C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "skippedEntity"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(name)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).SkippedEntityDefault(cGoUnpackString(name)))))
}

func (ptr *QXmlDefaultHandler) ConnectSkippedEntity(f func(name string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "skippedEntity"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "skippedEntity", func(name string) bool {
				signal.(func(string) bool)(name)
				return f(name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "skippedEntity", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectSkippedEntity() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "skippedEntity")
	}
}

func (ptr *QXmlDefaultHandler) SkippedEntity(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QXmlDefaultHandler_SkippedEntity(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) SkippedEntityDefault(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QXmlDefaultHandler_SkippedEntityDefault(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_StartCDATA
func callbackQXmlDefaultHandler_StartCDATA(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "startCDATA"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).StartCDATADefault())))
}

func (ptr *QXmlDefaultHandler) ConnectStartCDATA(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startCDATA"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "startCDATA", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startCDATA", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectStartCDATA() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "startCDATA")
	}
}

func (ptr *QXmlDefaultHandler) StartCDATA() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlDefaultHandler_StartCDATA(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) StartCDATADefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlDefaultHandler_StartCDATADefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_StartDTD
func callbackQXmlDefaultHandler_StartDTD(ptr unsafe.Pointer, name C.struct_QtXml_PackedString, publicId C.struct_QtXml_PackedString, systemId C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "startDTD"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, string, string) bool)(cGoUnpackString(name), cGoUnpackString(publicId), cGoUnpackString(systemId)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).StartDTDDefault(cGoUnpackString(name), cGoUnpackString(publicId), cGoUnpackString(systemId)))))
}

func (ptr *QXmlDefaultHandler) ConnectStartDTD(f func(name string, publicId string, systemId string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startDTD"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "startDTD", func(name string, publicId string, systemId string) bool {
				signal.(func(string, string, string) bool)(name, publicId, systemId)
				return f(name, publicId, systemId)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startDTD", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectStartDTD() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "startDTD")
	}
}

func (ptr *QXmlDefaultHandler) StartDTD(name string, publicId string, systemId string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var publicIdC *C.char
		if publicId != "" {
			publicIdC = C.CString(publicId)
			defer C.free(unsafe.Pointer(publicIdC))
		}
		var systemIdC *C.char
		if systemId != "" {
			systemIdC = C.CString(systemId)
			defer C.free(unsafe.Pointer(systemIdC))
		}
		return int8(C.QXmlDefaultHandler_StartDTD(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtXml_PackedString{data: publicIdC, len: C.longlong(len(publicId))}, C.struct_QtXml_PackedString{data: systemIdC, len: C.longlong(len(systemId))})) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) StartDTDDefault(name string, publicId string, systemId string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var publicIdC *C.char
		if publicId != "" {
			publicIdC = C.CString(publicId)
			defer C.free(unsafe.Pointer(publicIdC))
		}
		var systemIdC *C.char
		if systemId != "" {
			systemIdC = C.CString(systemId)
			defer C.free(unsafe.Pointer(systemIdC))
		}
		return int8(C.QXmlDefaultHandler_StartDTDDefault(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtXml_PackedString{data: publicIdC, len: C.longlong(len(publicId))}, C.struct_QtXml_PackedString{data: systemIdC, len: C.longlong(len(systemId))})) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_StartDocument
func callbackQXmlDefaultHandler_StartDocument(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "startDocument"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).StartDocumentDefault())))
}

func (ptr *QXmlDefaultHandler) ConnectStartDocument(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startDocument"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "startDocument", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startDocument", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectStartDocument() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "startDocument")
	}
}

func (ptr *QXmlDefaultHandler) StartDocument() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlDefaultHandler_StartDocument(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) StartDocumentDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlDefaultHandler_StartDocumentDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_StartElement
func callbackQXmlDefaultHandler_StartElement(ptr unsafe.Pointer, namespaceURI C.struct_QtXml_PackedString, localName C.struct_QtXml_PackedString, qName C.struct_QtXml_PackedString, atts unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "startElement"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, string, string, *QXmlAttributes) bool)(cGoUnpackString(namespaceURI), cGoUnpackString(localName), cGoUnpackString(qName), NewQXmlAttributesFromPointer(atts)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).StartElementDefault(cGoUnpackString(namespaceURI), cGoUnpackString(localName), cGoUnpackString(qName), NewQXmlAttributesFromPointer(atts)))))
}

func (ptr *QXmlDefaultHandler) ConnectStartElement(f func(namespaceURI string, localName string, qName string, atts *QXmlAttributes) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startElement"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "startElement", func(namespaceURI string, localName string, qName string, atts *QXmlAttributes) bool {
				signal.(func(string, string, string, *QXmlAttributes) bool)(namespaceURI, localName, qName, atts)
				return f(namespaceURI, localName, qName, atts)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startElement", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectStartElement() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "startElement")
	}
}

func (ptr *QXmlDefaultHandler) StartElement(namespaceURI string, localName string, qName string, atts QXmlAttributes_ITF) bool {
	if ptr.Pointer() != nil {
		var namespaceURIC *C.char
		if namespaceURI != "" {
			namespaceURIC = C.CString(namespaceURI)
			defer C.free(unsafe.Pointer(namespaceURIC))
		}
		var localNameC *C.char
		if localName != "" {
			localNameC = C.CString(localName)
			defer C.free(unsafe.Pointer(localNameC))
		}
		var qNameC *C.char
		if qName != "" {
			qNameC = C.CString(qName)
			defer C.free(unsafe.Pointer(qNameC))
		}
		return int8(C.QXmlDefaultHandler_StartElement(ptr.Pointer(), C.struct_QtXml_PackedString{data: namespaceURIC, len: C.longlong(len(namespaceURI))}, C.struct_QtXml_PackedString{data: localNameC, len: C.longlong(len(localName))}, C.struct_QtXml_PackedString{data: qNameC, len: C.longlong(len(qName))}, PointerFromQXmlAttributes(atts))) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) StartElementDefault(namespaceURI string, localName string, qName string, atts QXmlAttributes_ITF) bool {
	if ptr.Pointer() != nil {
		var namespaceURIC *C.char
		if namespaceURI != "" {
			namespaceURIC = C.CString(namespaceURI)
			defer C.free(unsafe.Pointer(namespaceURIC))
		}
		var localNameC *C.char
		if localName != "" {
			localNameC = C.CString(localName)
			defer C.free(unsafe.Pointer(localNameC))
		}
		var qNameC *C.char
		if qName != "" {
			qNameC = C.CString(qName)
			defer C.free(unsafe.Pointer(qNameC))
		}
		return int8(C.QXmlDefaultHandler_StartElementDefault(ptr.Pointer(), C.struct_QtXml_PackedString{data: namespaceURIC, len: C.longlong(len(namespaceURI))}, C.struct_QtXml_PackedString{data: localNameC, len: C.longlong(len(localName))}, C.struct_QtXml_PackedString{data: qNameC, len: C.longlong(len(qName))}, PointerFromQXmlAttributes(atts))) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_StartEntity
func callbackQXmlDefaultHandler_StartEntity(ptr unsafe.Pointer, name C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "startEntity"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(name)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).StartEntityDefault(cGoUnpackString(name)))))
}

func (ptr *QXmlDefaultHandler) ConnectStartEntity(f func(name string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startEntity"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "startEntity", func(name string) bool {
				signal.(func(string) bool)(name)
				return f(name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startEntity", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectStartEntity() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "startEntity")
	}
}

func (ptr *QXmlDefaultHandler) StartEntity(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QXmlDefaultHandler_StartEntity(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) StartEntityDefault(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QXmlDefaultHandler_StartEntityDefault(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_StartPrefixMapping
func callbackQXmlDefaultHandler_StartPrefixMapping(ptr unsafe.Pointer, prefix C.struct_QtXml_PackedString, uri C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "startPrefixMapping"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, string) bool)(cGoUnpackString(prefix), cGoUnpackString(uri)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).StartPrefixMappingDefault(cGoUnpackString(prefix), cGoUnpackString(uri)))))
}

func (ptr *QXmlDefaultHandler) ConnectStartPrefixMapping(f func(prefix string, uri string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startPrefixMapping"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "startPrefixMapping", func(prefix string, uri string) bool {
				signal.(func(string, string) bool)(prefix, uri)
				return f(prefix, uri)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startPrefixMapping", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectStartPrefixMapping() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "startPrefixMapping")
	}
}

func (ptr *QXmlDefaultHandler) StartPrefixMapping(prefix string, uri string) bool {
	if ptr.Pointer() != nil {
		var prefixC *C.char
		if prefix != "" {
			prefixC = C.CString(prefix)
			defer C.free(unsafe.Pointer(prefixC))
		}
		var uriC *C.char
		if uri != "" {
			uriC = C.CString(uri)
			defer C.free(unsafe.Pointer(uriC))
		}
		return int8(C.QXmlDefaultHandler_StartPrefixMapping(ptr.Pointer(), C.struct_QtXml_PackedString{data: prefixC, len: C.longlong(len(prefix))}, C.struct_QtXml_PackedString{data: uriC, len: C.longlong(len(uri))})) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) StartPrefixMappingDefault(prefix string, uri string) bool {
	if ptr.Pointer() != nil {
		var prefixC *C.char
		if prefix != "" {
			prefixC = C.CString(prefix)
			defer C.free(unsafe.Pointer(prefixC))
		}
		var uriC *C.char
		if uri != "" {
			uriC = C.CString(uri)
			defer C.free(unsafe.Pointer(uriC))
		}
		return int8(C.QXmlDefaultHandler_StartPrefixMappingDefault(ptr.Pointer(), C.struct_QtXml_PackedString{data: prefixC, len: C.longlong(len(prefix))}, C.struct_QtXml_PackedString{data: uriC, len: C.longlong(len(uri))})) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_UnparsedEntityDecl
func callbackQXmlDefaultHandler_UnparsedEntityDecl(ptr unsafe.Pointer, name C.struct_QtXml_PackedString, publicId C.struct_QtXml_PackedString, systemId C.struct_QtXml_PackedString, notationName C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "unparsedEntityDecl"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, string, string, string) bool)(cGoUnpackString(name), cGoUnpackString(publicId), cGoUnpackString(systemId), cGoUnpackString(notationName)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).UnparsedEntityDeclDefault(cGoUnpackString(name), cGoUnpackString(publicId), cGoUnpackString(systemId), cGoUnpackString(notationName)))))
}

func (ptr *QXmlDefaultHandler) ConnectUnparsedEntityDecl(f func(name string, publicId string, systemId string, notationName string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "unparsedEntityDecl"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "unparsedEntityDecl", func(name string, publicId string, systemId string, notationName string) bool {
				signal.(func(string, string, string, string) bool)(name, publicId, systemId, notationName)
				return f(name, publicId, systemId, notationName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "unparsedEntityDecl", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectUnparsedEntityDecl() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "unparsedEntityDecl")
	}
}

func (ptr *QXmlDefaultHandler) UnparsedEntityDecl(name string, publicId string, systemId string, notationName string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var publicIdC *C.char
		if publicId != "" {
			publicIdC = C.CString(publicId)
			defer C.free(unsafe.Pointer(publicIdC))
		}
		var systemIdC *C.char
		if systemId != "" {
			systemIdC = C.CString(systemId)
			defer C.free(unsafe.Pointer(systemIdC))
		}
		var notationNameC *C.char
		if notationName != "" {
			notationNameC = C.CString(notationName)
			defer C.free(unsafe.Pointer(notationNameC))
		}
		return int8(C.QXmlDefaultHandler_UnparsedEntityDecl(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtXml_PackedString{data: publicIdC, len: C.longlong(len(publicId))}, C.struct_QtXml_PackedString{data: systemIdC, len: C.longlong(len(systemId))}, C.struct_QtXml_PackedString{data: notationNameC, len: C.longlong(len(notationName))})) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) UnparsedEntityDeclDefault(name string, publicId string, systemId string, notationName string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var publicIdC *C.char
		if publicId != "" {
			publicIdC = C.CString(publicId)
			defer C.free(unsafe.Pointer(publicIdC))
		}
		var systemIdC *C.char
		if systemId != "" {
			systemIdC = C.CString(systemId)
			defer C.free(unsafe.Pointer(systemIdC))
		}
		var notationNameC *C.char
		if notationName != "" {
			notationNameC = C.CString(notationName)
			defer C.free(unsafe.Pointer(notationNameC))
		}
		return int8(C.QXmlDefaultHandler_UnparsedEntityDeclDefault(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtXml_PackedString{data: publicIdC, len: C.longlong(len(publicId))}, C.struct_QtXml_PackedString{data: systemIdC, len: C.longlong(len(systemId))}, C.struct_QtXml_PackedString{data: notationNameC, len: C.longlong(len(notationName))})) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_Warning
func callbackQXmlDefaultHandler_Warning(ptr unsafe.Pointer, exception unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "warning"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QXmlParseException) bool)(NewQXmlParseExceptionFromPointer(exception)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlDefaultHandlerFromPointer(ptr).WarningDefault(NewQXmlParseExceptionFromPointer(exception)))))
}

func (ptr *QXmlDefaultHandler) ConnectWarning(f func(exception *QXmlParseException) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "warning"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "warning", func(exception *QXmlParseException) bool {
				signal.(func(*QXmlParseException) bool)(exception)
				return f(exception)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "warning", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectWarning() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "warning")
	}
}

func (ptr *QXmlDefaultHandler) Warning(exception QXmlParseException_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlDefaultHandler_Warning(ptr.Pointer(), PointerFromQXmlParseException(exception))) != 0
	}
	return false
}

func (ptr *QXmlDefaultHandler) WarningDefault(exception QXmlParseException_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlDefaultHandler_WarningDefault(ptr.Pointer(), PointerFromQXmlParseException(exception))) != 0
	}
	return false
}

//export callbackQXmlDefaultHandler_SetDocumentLocator
func callbackQXmlDefaultHandler_SetDocumentLocator(ptr unsafe.Pointer, locator unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setDocumentLocator"); signal != nil {
		signal.(func(*QXmlLocator))(NewQXmlLocatorFromPointer(locator))
	} else {
		NewQXmlDefaultHandlerFromPointer(ptr).SetDocumentLocatorDefault(NewQXmlLocatorFromPointer(locator))
	}
}

func (ptr *QXmlDefaultHandler) ConnectSetDocumentLocator(f func(locator *QXmlLocator)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setDocumentLocator"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setDocumentLocator", func(locator *QXmlLocator) {
				signal.(func(*QXmlLocator))(locator)
				f(locator)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setDocumentLocator", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectSetDocumentLocator() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setDocumentLocator")
	}
}

func (ptr *QXmlDefaultHandler) SetDocumentLocator(locator QXmlLocator_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlDefaultHandler_SetDocumentLocator(ptr.Pointer(), PointerFromQXmlLocator(locator))
	}
}

func (ptr *QXmlDefaultHandler) SetDocumentLocatorDefault(locator QXmlLocator_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlDefaultHandler_SetDocumentLocatorDefault(ptr.Pointer(), PointerFromQXmlLocator(locator))
	}
}

//export callbackQXmlDefaultHandler_DestroyQXmlDefaultHandler
func callbackQXmlDefaultHandler_DestroyQXmlDefaultHandler(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QXmlDefaultHandler"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlDefaultHandlerFromPointer(ptr).DestroyQXmlDefaultHandlerDefault()
	}
}

func (ptr *QXmlDefaultHandler) ConnectDestroyQXmlDefaultHandler(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QXmlDefaultHandler"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QXmlDefaultHandler", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QXmlDefaultHandler", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectDestroyQXmlDefaultHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QXmlDefaultHandler")
	}
}

func (ptr *QXmlDefaultHandler) DestroyQXmlDefaultHandler() {
	if ptr.Pointer() != nil {
		C.QXmlDefaultHandler_DestroyQXmlDefaultHandler(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXmlDefaultHandler) DestroyQXmlDefaultHandlerDefault() {
	if ptr.Pointer() != nil {
		C.QXmlDefaultHandler_DestroyQXmlDefaultHandlerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQXmlDefaultHandler_ErrorString
func callbackQXmlDefaultHandler_ErrorString(ptr unsafe.Pointer) C.struct_QtXml_PackedString {
	if signal := qt.GetSignal(ptr, "errorString"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtXml_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQXmlDefaultHandlerFromPointer(ptr).ErrorStringDefault()
	return C.struct_QtXml_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QXmlDefaultHandler) ConnectErrorString(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "errorString"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "errorString", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "errorString", f)
		}
	}
}

func (ptr *QXmlDefaultHandler) DisconnectErrorString() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "errorString")
	}
}

func (ptr *QXmlDefaultHandler) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlDefaultHandler_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlDefaultHandler) ErrorStringDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlDefaultHandler_ErrorStringDefault(ptr.Pointer()))
	}
	return ""
}

type QXmlEntityResolver struct {
	ptr unsafe.Pointer
}

type QXmlEntityResolver_ITF interface {
	QXmlEntityResolver_PTR() *QXmlEntityResolver
}

func (ptr *QXmlEntityResolver) QXmlEntityResolver_PTR() *QXmlEntityResolver {
	return ptr
}

func (ptr *QXmlEntityResolver) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlEntityResolver) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQXmlEntityResolver(ptr QXmlEntityResolver_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlEntityResolver_PTR().Pointer()
	}
	return nil
}

func NewQXmlEntityResolverFromPointer(ptr unsafe.Pointer) (n *QXmlEntityResolver) {
	n = new(QXmlEntityResolver)
	n.SetPointer(ptr)
	return
}

//export callbackQXmlEntityResolver_DestroyQXmlEntityResolver
func callbackQXmlEntityResolver_DestroyQXmlEntityResolver(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QXmlEntityResolver"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlEntityResolverFromPointer(ptr).DestroyQXmlEntityResolverDefault()
	}
}

func (ptr *QXmlEntityResolver) ConnectDestroyQXmlEntityResolver(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QXmlEntityResolver"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QXmlEntityResolver", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QXmlEntityResolver", f)
		}
	}
}

func (ptr *QXmlEntityResolver) DisconnectDestroyQXmlEntityResolver() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QXmlEntityResolver")
	}
}

func (ptr *QXmlEntityResolver) DestroyQXmlEntityResolver() {
	if ptr.Pointer() != nil {
		C.QXmlEntityResolver_DestroyQXmlEntityResolver(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXmlEntityResolver) DestroyQXmlEntityResolverDefault() {
	if ptr.Pointer() != nil {
		C.QXmlEntityResolver_DestroyQXmlEntityResolverDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQXmlEntityResolver_ErrorString
func callbackQXmlEntityResolver_ErrorString(ptr unsafe.Pointer) C.struct_QtXml_PackedString {
	if signal := qt.GetSignal(ptr, "errorString"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtXml_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtXml_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QXmlEntityResolver) ConnectErrorString(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "errorString"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "errorString", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "errorString", f)
		}
	}
}

func (ptr *QXmlEntityResolver) DisconnectErrorString() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "errorString")
	}
}

func (ptr *QXmlEntityResolver) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlEntityResolver_ErrorString(ptr.Pointer()))
	}
	return ""
}

type QXmlErrorHandler struct {
	ptr unsafe.Pointer
}

type QXmlErrorHandler_ITF interface {
	QXmlErrorHandler_PTR() *QXmlErrorHandler
}

func (ptr *QXmlErrorHandler) QXmlErrorHandler_PTR() *QXmlErrorHandler {
	return ptr
}

func (ptr *QXmlErrorHandler) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlErrorHandler) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQXmlErrorHandler(ptr QXmlErrorHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlErrorHandler_PTR().Pointer()
	}
	return nil
}

func NewQXmlErrorHandlerFromPointer(ptr unsafe.Pointer) (n *QXmlErrorHandler) {
	n = new(QXmlErrorHandler)
	n.SetPointer(ptr)
	return
}

//export callbackQXmlErrorHandler_Error
func callbackQXmlErrorHandler_Error(ptr unsafe.Pointer, exception unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "error"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QXmlParseException) bool)(NewQXmlParseExceptionFromPointer(exception)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlErrorHandler) ConnectError(f func(exception *QXmlParseException) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "error"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "error", func(exception *QXmlParseException) bool {
				signal.(func(*QXmlParseException) bool)(exception)
				return f(exception)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error", f)
		}
	}
}

func (ptr *QXmlErrorHandler) DisconnectError() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "error")
	}
}

func (ptr *QXmlErrorHandler) Error(exception QXmlParseException_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlErrorHandler_Error(ptr.Pointer(), PointerFromQXmlParseException(exception))) != 0
	}
	return false
}

//export callbackQXmlErrorHandler_FatalError
func callbackQXmlErrorHandler_FatalError(ptr unsafe.Pointer, exception unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "fatalError"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QXmlParseException) bool)(NewQXmlParseExceptionFromPointer(exception)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlErrorHandler) ConnectFatalError(f func(exception *QXmlParseException) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "fatalError"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fatalError", func(exception *QXmlParseException) bool {
				signal.(func(*QXmlParseException) bool)(exception)
				return f(exception)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fatalError", f)
		}
	}
}

func (ptr *QXmlErrorHandler) DisconnectFatalError() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "fatalError")
	}
}

func (ptr *QXmlErrorHandler) FatalError(exception QXmlParseException_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlErrorHandler_FatalError(ptr.Pointer(), PointerFromQXmlParseException(exception))) != 0
	}
	return false
}

//export callbackQXmlErrorHandler_Warning
func callbackQXmlErrorHandler_Warning(ptr unsafe.Pointer, exception unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "warning"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QXmlParseException) bool)(NewQXmlParseExceptionFromPointer(exception)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlErrorHandler) ConnectWarning(f func(exception *QXmlParseException) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "warning"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "warning", func(exception *QXmlParseException) bool {
				signal.(func(*QXmlParseException) bool)(exception)
				return f(exception)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "warning", f)
		}
	}
}

func (ptr *QXmlErrorHandler) DisconnectWarning() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "warning")
	}
}

func (ptr *QXmlErrorHandler) Warning(exception QXmlParseException_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlErrorHandler_Warning(ptr.Pointer(), PointerFromQXmlParseException(exception))) != 0
	}
	return false
}

//export callbackQXmlErrorHandler_DestroyQXmlErrorHandler
func callbackQXmlErrorHandler_DestroyQXmlErrorHandler(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QXmlErrorHandler"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlErrorHandlerFromPointer(ptr).DestroyQXmlErrorHandlerDefault()
	}
}

func (ptr *QXmlErrorHandler) ConnectDestroyQXmlErrorHandler(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QXmlErrorHandler"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QXmlErrorHandler", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QXmlErrorHandler", f)
		}
	}
}

func (ptr *QXmlErrorHandler) DisconnectDestroyQXmlErrorHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QXmlErrorHandler")
	}
}

func (ptr *QXmlErrorHandler) DestroyQXmlErrorHandler() {
	if ptr.Pointer() != nil {
		C.QXmlErrorHandler_DestroyQXmlErrorHandler(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXmlErrorHandler) DestroyQXmlErrorHandlerDefault() {
	if ptr.Pointer() != nil {
		C.QXmlErrorHandler_DestroyQXmlErrorHandlerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQXmlErrorHandler_ErrorString
func callbackQXmlErrorHandler_ErrorString(ptr unsafe.Pointer) C.struct_QtXml_PackedString {
	if signal := qt.GetSignal(ptr, "errorString"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtXml_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtXml_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QXmlErrorHandler) ConnectErrorString(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "errorString"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "errorString", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "errorString", f)
		}
	}
}

func (ptr *QXmlErrorHandler) DisconnectErrorString() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "errorString")
	}
}

func (ptr *QXmlErrorHandler) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlErrorHandler_ErrorString(ptr.Pointer()))
	}
	return ""
}

type QXmlInputSource struct {
	ptr unsafe.Pointer
}

type QXmlInputSource_ITF interface {
	QXmlInputSource_PTR() *QXmlInputSource
}

func (ptr *QXmlInputSource) QXmlInputSource_PTR() *QXmlInputSource {
	return ptr
}

func (ptr *QXmlInputSource) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlInputSource) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQXmlInputSource(ptr QXmlInputSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlInputSource_PTR().Pointer()
	}
	return nil
}

func NewQXmlInputSourceFromPointer(ptr unsafe.Pointer) (n *QXmlInputSource) {
	n = new(QXmlInputSource)
	n.SetPointer(ptr)
	return
}

//export callbackQXmlInputSource_Next
func callbackQXmlInputSource_Next(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "next"); signal != nil {
		return core.PointerFromQChar(signal.(func() *core.QChar)())
	}

	return core.PointerFromQChar(NewQXmlInputSourceFromPointer(ptr).NextDefault())
}

func (ptr *QXmlInputSource) ConnectNext(f func() *core.QChar) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "next"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "next", func() *core.QChar {
				signal.(func() *core.QChar)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "next", f)
		}
	}
}

func (ptr *QXmlInputSource) DisconnectNext() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "next")
	}
}

func (ptr *QXmlInputSource) Next() *core.QChar {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQCharFromPointer(C.QXmlInputSource_Next(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QChar).DestroyQChar)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlInputSource) NextDefault() *core.QChar {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQCharFromPointer(C.QXmlInputSource_NextDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QChar).DestroyQChar)
		return tmpValue
	}
	return nil
}

//export callbackQXmlInputSource_FromRawData
func callbackQXmlInputSource_FromRawData(ptr unsafe.Pointer, data unsafe.Pointer, beginning C.char) C.struct_QtXml_PackedString {
	if signal := qt.GetSignal(ptr, "fromRawData"); signal != nil {
		tempVal := signal.(func(*core.QByteArray, bool) string)(core.NewQByteArrayFromPointer(data), int8(beginning) != 0)
		return C.struct_QtXml_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQXmlInputSourceFromPointer(ptr).FromRawDataDefault(core.NewQByteArrayFromPointer(data), int8(beginning) != 0)
	return C.struct_QtXml_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QXmlInputSource) ConnectFromRawData(f func(data *core.QByteArray, beginning bool) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "fromRawData"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fromRawData", func(data *core.QByteArray, beginning bool) string {
				signal.(func(*core.QByteArray, bool) string)(data, beginning)
				return f(data, beginning)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fromRawData", f)
		}
	}
}

func (ptr *QXmlInputSource) DisconnectFromRawData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "fromRawData")
	}
}

func (ptr *QXmlInputSource) FromRawData(data core.QByteArray_ITF, beginning bool) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlInputSource_FromRawData(ptr.Pointer(), core.PointerFromQByteArray(data), C.char(int8(qt.GoBoolToInt(beginning)))))
	}
	return ""
}

func (ptr *QXmlInputSource) FromRawDataDefault(data core.QByteArray_ITF, beginning bool) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlInputSource_FromRawDataDefault(ptr.Pointer(), core.PointerFromQByteArray(data), C.char(int8(qt.GoBoolToInt(beginning)))))
	}
	return ""
}

func NewQXmlInputSource() *QXmlInputSource {
	return NewQXmlInputSourceFromPointer(C.QXmlInputSource_NewQXmlInputSource())
}

func NewQXmlInputSource2(dev core.QIODevice_ITF) *QXmlInputSource {
	return NewQXmlInputSourceFromPointer(C.QXmlInputSource_NewQXmlInputSource2(core.PointerFromQIODevice(dev)))
}

//export callbackQXmlInputSource_FetchData
func callbackQXmlInputSource_FetchData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchData"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlInputSourceFromPointer(ptr).FetchDataDefault()
	}
}

func (ptr *QXmlInputSource) ConnectFetchData(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "fetchData"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fetchData", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fetchData", f)
		}
	}
}

func (ptr *QXmlInputSource) DisconnectFetchData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "fetchData")
	}
}

func (ptr *QXmlInputSource) FetchData() {
	if ptr.Pointer() != nil {
		C.QXmlInputSource_FetchData(ptr.Pointer())
	}
}

func (ptr *QXmlInputSource) FetchDataDefault() {
	if ptr.Pointer() != nil {
		C.QXmlInputSource_FetchDataDefault(ptr.Pointer())
	}
}

//export callbackQXmlInputSource_Reset
func callbackQXmlInputSource_Reset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "reset"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlInputSourceFromPointer(ptr).ResetDefault()
	}
}

func (ptr *QXmlInputSource) ConnectReset(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "reset"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "reset", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "reset", f)
		}
	}
}

func (ptr *QXmlInputSource) DisconnectReset() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "reset")
	}
}

func (ptr *QXmlInputSource) Reset() {
	if ptr.Pointer() != nil {
		C.QXmlInputSource_Reset(ptr.Pointer())
	}
}

func (ptr *QXmlInputSource) ResetDefault() {
	if ptr.Pointer() != nil {
		C.QXmlInputSource_ResetDefault(ptr.Pointer())
	}
}

//export callbackQXmlInputSource_SetData2
func callbackQXmlInputSource_SetData2(ptr unsafe.Pointer, dat unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setData2"); signal != nil {
		signal.(func(*core.QByteArray))(core.NewQByteArrayFromPointer(dat))
	} else {
		NewQXmlInputSourceFromPointer(ptr).SetData2Default(core.NewQByteArrayFromPointer(dat))
	}
}

func (ptr *QXmlInputSource) ConnectSetData2(f func(dat *core.QByteArray)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setData2"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setData2", func(dat *core.QByteArray) {
				signal.(func(*core.QByteArray))(dat)
				f(dat)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setData2", f)
		}
	}
}

func (ptr *QXmlInputSource) DisconnectSetData2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setData2")
	}
}

func (ptr *QXmlInputSource) SetData2(dat core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlInputSource_SetData2(ptr.Pointer(), core.PointerFromQByteArray(dat))
	}
}

func (ptr *QXmlInputSource) SetData2Default(dat core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlInputSource_SetData2Default(ptr.Pointer(), core.PointerFromQByteArray(dat))
	}
}

//export callbackQXmlInputSource_SetData
func callbackQXmlInputSource_SetData(ptr unsafe.Pointer, dat C.struct_QtXml_PackedString) {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		signal.(func(string))(cGoUnpackString(dat))
	} else {
		NewQXmlInputSourceFromPointer(ptr).SetDataDefault(cGoUnpackString(dat))
	}
}

func (ptr *QXmlInputSource) ConnectSetData(f func(dat string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setData"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setData", func(dat string) {
				signal.(func(string))(dat)
				f(dat)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setData", f)
		}
	}
}

func (ptr *QXmlInputSource) DisconnectSetData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setData")
	}
}

func (ptr *QXmlInputSource) SetData(dat string) {
	if ptr.Pointer() != nil {
		var datC *C.char
		if dat != "" {
			datC = C.CString(dat)
			defer C.free(unsafe.Pointer(datC))
		}
		C.QXmlInputSource_SetData(ptr.Pointer(), C.struct_QtXml_PackedString{data: datC, len: C.longlong(len(dat))})
	}
}

func (ptr *QXmlInputSource) SetDataDefault(dat string) {
	if ptr.Pointer() != nil {
		var datC *C.char
		if dat != "" {
			datC = C.CString(dat)
			defer C.free(unsafe.Pointer(datC))
		}
		C.QXmlInputSource_SetDataDefault(ptr.Pointer(), C.struct_QtXml_PackedString{data: datC, len: C.longlong(len(dat))})
	}
}

//export callbackQXmlInputSource_DestroyQXmlInputSource
func callbackQXmlInputSource_DestroyQXmlInputSource(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QXmlInputSource"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlInputSourceFromPointer(ptr).DestroyQXmlInputSourceDefault()
	}
}

func (ptr *QXmlInputSource) ConnectDestroyQXmlInputSource(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QXmlInputSource"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QXmlInputSource", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QXmlInputSource", f)
		}
	}
}

func (ptr *QXmlInputSource) DisconnectDestroyQXmlInputSource() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QXmlInputSource")
	}
}

func (ptr *QXmlInputSource) DestroyQXmlInputSource() {
	if ptr.Pointer() != nil {
		C.QXmlInputSource_DestroyQXmlInputSource(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QXmlInputSource) DestroyQXmlInputSourceDefault() {
	if ptr.Pointer() != nil {
		C.QXmlInputSource_DestroyQXmlInputSourceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQXmlInputSource_Data
func callbackQXmlInputSource_Data(ptr unsafe.Pointer) C.struct_QtXml_PackedString {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtXml_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQXmlInputSourceFromPointer(ptr).DataDefault()
	return C.struct_QtXml_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QXmlInputSource) ConnectData(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "data"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "data", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "data", f)
		}
	}
}

func (ptr *QXmlInputSource) DisconnectData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "data")
	}
}

func (ptr *QXmlInputSource) Data() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlInputSource_Data(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlInputSource) DataDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlInputSource_DataDefault(ptr.Pointer()))
	}
	return ""
}

func QXmlInputSource_EndOfData() uint16 {
	return uint16(C.QXmlInputSource_QXmlInputSource_EndOfData())
}

func (ptr *QXmlInputSource) EndOfData() uint16 {
	return uint16(C.QXmlInputSource_QXmlInputSource_EndOfData())
}

func QXmlInputSource_EndOfDocument() uint16 {
	return uint16(C.QXmlInputSource_QXmlInputSource_EndOfDocument())
}

func (ptr *QXmlInputSource) EndOfDocument() uint16 {
	return uint16(C.QXmlInputSource_QXmlInputSource_EndOfDocument())
}

type QXmlLexicalHandler struct {
	ptr unsafe.Pointer
}

type QXmlLexicalHandler_ITF interface {
	QXmlLexicalHandler_PTR() *QXmlLexicalHandler
}

func (ptr *QXmlLexicalHandler) QXmlLexicalHandler_PTR() *QXmlLexicalHandler {
	return ptr
}

func (ptr *QXmlLexicalHandler) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlLexicalHandler) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQXmlLexicalHandler(ptr QXmlLexicalHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlLexicalHandler_PTR().Pointer()
	}
	return nil
}

func NewQXmlLexicalHandlerFromPointer(ptr unsafe.Pointer) (n *QXmlLexicalHandler) {
	n = new(QXmlLexicalHandler)
	n.SetPointer(ptr)
	return
}

//export callbackQXmlLexicalHandler_Comment
func callbackQXmlLexicalHandler_Comment(ptr unsafe.Pointer, ch C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "comment"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(ch)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlLexicalHandler) ConnectComment(f func(ch string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "comment"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "comment", func(ch string) bool {
				signal.(func(string) bool)(ch)
				return f(ch)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "comment", f)
		}
	}
}

func (ptr *QXmlLexicalHandler) DisconnectComment() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "comment")
	}
}

func (ptr *QXmlLexicalHandler) Comment(ch string) bool {
	if ptr.Pointer() != nil {
		var chC *C.char
		if ch != "" {
			chC = C.CString(ch)
			defer C.free(unsafe.Pointer(chC))
		}
		return int8(C.QXmlLexicalHandler_Comment(ptr.Pointer(), C.struct_QtXml_PackedString{data: chC, len: C.longlong(len(ch))})) != 0
	}
	return false
}

//export callbackQXmlLexicalHandler_EndCDATA
func callbackQXmlLexicalHandler_EndCDATA(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "endCDATA"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlLexicalHandler) ConnectEndCDATA(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "endCDATA"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "endCDATA", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endCDATA", f)
		}
	}
}

func (ptr *QXmlLexicalHandler) DisconnectEndCDATA() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "endCDATA")
	}
}

func (ptr *QXmlLexicalHandler) EndCDATA() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlLexicalHandler_EndCDATA(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQXmlLexicalHandler_EndDTD
func callbackQXmlLexicalHandler_EndDTD(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "endDTD"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlLexicalHandler) ConnectEndDTD(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "endDTD"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "endDTD", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endDTD", f)
		}
	}
}

func (ptr *QXmlLexicalHandler) DisconnectEndDTD() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "endDTD")
	}
}

func (ptr *QXmlLexicalHandler) EndDTD() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlLexicalHandler_EndDTD(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQXmlLexicalHandler_EndEntity
func callbackQXmlLexicalHandler_EndEntity(ptr unsafe.Pointer, name C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "endEntity"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(name)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlLexicalHandler) ConnectEndEntity(f func(name string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "endEntity"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "endEntity", func(name string) bool {
				signal.(func(string) bool)(name)
				return f(name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endEntity", f)
		}
	}
}

func (ptr *QXmlLexicalHandler) DisconnectEndEntity() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "endEntity")
	}
}

func (ptr *QXmlLexicalHandler) EndEntity(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QXmlLexicalHandler_EndEntity(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

//export callbackQXmlLexicalHandler_StartCDATA
func callbackQXmlLexicalHandler_StartCDATA(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "startCDATA"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlLexicalHandler) ConnectStartCDATA(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startCDATA"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "startCDATA", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startCDATA", f)
		}
	}
}

func (ptr *QXmlLexicalHandler) DisconnectStartCDATA() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "startCDATA")
	}
}

func (ptr *QXmlLexicalHandler) StartCDATA() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlLexicalHandler_StartCDATA(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQXmlLexicalHandler_StartDTD
func callbackQXmlLexicalHandler_StartDTD(ptr unsafe.Pointer, name C.struct_QtXml_PackedString, publicId C.struct_QtXml_PackedString, systemId C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "startDTD"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, string, string) bool)(cGoUnpackString(name), cGoUnpackString(publicId), cGoUnpackString(systemId)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlLexicalHandler) ConnectStartDTD(f func(name string, publicId string, systemId string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startDTD"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "startDTD", func(name string, publicId string, systemId string) bool {
				signal.(func(string, string, string) bool)(name, publicId, systemId)
				return f(name, publicId, systemId)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startDTD", f)
		}
	}
}

func (ptr *QXmlLexicalHandler) DisconnectStartDTD() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "startDTD")
	}
}

func (ptr *QXmlLexicalHandler) StartDTD(name string, publicId string, systemId string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var publicIdC *C.char
		if publicId != "" {
			publicIdC = C.CString(publicId)
			defer C.free(unsafe.Pointer(publicIdC))
		}
		var systemIdC *C.char
		if systemId != "" {
			systemIdC = C.CString(systemId)
			defer C.free(unsafe.Pointer(systemIdC))
		}
		return int8(C.QXmlLexicalHandler_StartDTD(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtXml_PackedString{data: publicIdC, len: C.longlong(len(publicId))}, C.struct_QtXml_PackedString{data: systemIdC, len: C.longlong(len(systemId))})) != 0
	}
	return false
}

//export callbackQXmlLexicalHandler_StartEntity
func callbackQXmlLexicalHandler_StartEntity(ptr unsafe.Pointer, name C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "startEntity"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(name)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlLexicalHandler) ConnectStartEntity(f func(name string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startEntity"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "startEntity", func(name string) bool {
				signal.(func(string) bool)(name)
				return f(name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startEntity", f)
		}
	}
}

func (ptr *QXmlLexicalHandler) DisconnectStartEntity() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "startEntity")
	}
}

func (ptr *QXmlLexicalHandler) StartEntity(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QXmlLexicalHandler_StartEntity(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

//export callbackQXmlLexicalHandler_DestroyQXmlLexicalHandler
func callbackQXmlLexicalHandler_DestroyQXmlLexicalHandler(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QXmlLexicalHandler"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlLexicalHandlerFromPointer(ptr).DestroyQXmlLexicalHandlerDefault()
	}
}

func (ptr *QXmlLexicalHandler) ConnectDestroyQXmlLexicalHandler(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QXmlLexicalHandler"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QXmlLexicalHandler", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QXmlLexicalHandler", f)
		}
	}
}

func (ptr *QXmlLexicalHandler) DisconnectDestroyQXmlLexicalHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QXmlLexicalHandler")
	}
}

func (ptr *QXmlLexicalHandler) DestroyQXmlLexicalHandler() {
	if ptr.Pointer() != nil {
		C.QXmlLexicalHandler_DestroyQXmlLexicalHandler(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXmlLexicalHandler) DestroyQXmlLexicalHandlerDefault() {
	if ptr.Pointer() != nil {
		C.QXmlLexicalHandler_DestroyQXmlLexicalHandlerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQXmlLexicalHandler_ErrorString
func callbackQXmlLexicalHandler_ErrorString(ptr unsafe.Pointer) C.struct_QtXml_PackedString {
	if signal := qt.GetSignal(ptr, "errorString"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtXml_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtXml_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QXmlLexicalHandler) ConnectErrorString(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "errorString"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "errorString", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "errorString", f)
		}
	}
}

func (ptr *QXmlLexicalHandler) DisconnectErrorString() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "errorString")
	}
}

func (ptr *QXmlLexicalHandler) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlLexicalHandler_ErrorString(ptr.Pointer()))
	}
	return ""
}

type QXmlLocator struct {
	ptr unsafe.Pointer
}

type QXmlLocator_ITF interface {
	QXmlLocator_PTR() *QXmlLocator
}

func (ptr *QXmlLocator) QXmlLocator_PTR() *QXmlLocator {
	return ptr
}

func (ptr *QXmlLocator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlLocator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQXmlLocator(ptr QXmlLocator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlLocator_PTR().Pointer()
	}
	return nil
}

func NewQXmlLocatorFromPointer(ptr unsafe.Pointer) (n *QXmlLocator) {
	n = new(QXmlLocator)
	n.SetPointer(ptr)
	return
}
func NewQXmlLocator() *QXmlLocator {
	return NewQXmlLocatorFromPointer(C.QXmlLocator_NewQXmlLocator())
}

//export callbackQXmlLocator_DestroyQXmlLocator
func callbackQXmlLocator_DestroyQXmlLocator(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QXmlLocator"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlLocatorFromPointer(ptr).DestroyQXmlLocatorDefault()
	}
}

func (ptr *QXmlLocator) ConnectDestroyQXmlLocator(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QXmlLocator"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QXmlLocator", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QXmlLocator", f)
		}
	}
}

func (ptr *QXmlLocator) DisconnectDestroyQXmlLocator() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QXmlLocator")
	}
}

func (ptr *QXmlLocator) DestroyQXmlLocator() {
	if ptr.Pointer() != nil {
		C.QXmlLocator_DestroyQXmlLocator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXmlLocator) DestroyQXmlLocatorDefault() {
	if ptr.Pointer() != nil {
		C.QXmlLocator_DestroyQXmlLocatorDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQXmlLocator_ColumnNumber
func callbackQXmlLocator_ColumnNumber(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnNumber"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QXmlLocator) ConnectColumnNumber(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "columnNumber"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "columnNumber", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "columnNumber", f)
		}
	}
}

func (ptr *QXmlLocator) DisconnectColumnNumber() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "columnNumber")
	}
}

func (ptr *QXmlLocator) ColumnNumber() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QXmlLocator_ColumnNumber(ptr.Pointer())))
	}
	return 0
}

//export callbackQXmlLocator_LineNumber
func callbackQXmlLocator_LineNumber(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "lineNumber"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QXmlLocator) ConnectLineNumber(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "lineNumber"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "lineNumber", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lineNumber", f)
		}
	}
}

func (ptr *QXmlLocator) DisconnectLineNumber() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "lineNumber")
	}
}

func (ptr *QXmlLocator) LineNumber() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QXmlLocator_LineNumber(ptr.Pointer())))
	}
	return 0
}

type QXmlNamespaceSupport struct {
	ptr unsafe.Pointer
}

type QXmlNamespaceSupport_ITF interface {
	QXmlNamespaceSupport_PTR() *QXmlNamespaceSupport
}

func (ptr *QXmlNamespaceSupport) QXmlNamespaceSupport_PTR() *QXmlNamespaceSupport {
	return ptr
}

func (ptr *QXmlNamespaceSupport) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlNamespaceSupport) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQXmlNamespaceSupport(ptr QXmlNamespaceSupport_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlNamespaceSupport_PTR().Pointer()
	}
	return nil
}

func NewQXmlNamespaceSupportFromPointer(ptr unsafe.Pointer) (n *QXmlNamespaceSupport) {
	n = new(QXmlNamespaceSupport)
	n.SetPointer(ptr)
	return
}
func NewQXmlNamespaceSupport() *QXmlNamespaceSupport {
	tmpValue := NewQXmlNamespaceSupportFromPointer(C.QXmlNamespaceSupport_NewQXmlNamespaceSupport())
	runtime.SetFinalizer(tmpValue, (*QXmlNamespaceSupport).DestroyQXmlNamespaceSupport)
	return tmpValue
}

func (ptr *QXmlNamespaceSupport) PopContext() {
	if ptr.Pointer() != nil {
		C.QXmlNamespaceSupport_PopContext(ptr.Pointer())
	}
}

func (ptr *QXmlNamespaceSupport) PushContext() {
	if ptr.Pointer() != nil {
		C.QXmlNamespaceSupport_PushContext(ptr.Pointer())
	}
}

func (ptr *QXmlNamespaceSupport) Reset() {
	if ptr.Pointer() != nil {
		C.QXmlNamespaceSupport_Reset(ptr.Pointer())
	}
}

func (ptr *QXmlNamespaceSupport) SetPrefix(pre string, uri string) {
	if ptr.Pointer() != nil {
		var preC *C.char
		if pre != "" {
			preC = C.CString(pre)
			defer C.free(unsafe.Pointer(preC))
		}
		var uriC *C.char
		if uri != "" {
			uriC = C.CString(uri)
			defer C.free(unsafe.Pointer(uriC))
		}
		C.QXmlNamespaceSupport_SetPrefix(ptr.Pointer(), C.struct_QtXml_PackedString{data: preC, len: C.longlong(len(pre))}, C.struct_QtXml_PackedString{data: uriC, len: C.longlong(len(uri))})
	}
}

func (ptr *QXmlNamespaceSupport) DestroyQXmlNamespaceSupport() {
	if ptr.Pointer() != nil {
		C.QXmlNamespaceSupport_DestroyQXmlNamespaceSupport(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QXmlNamespaceSupport) Prefix(uri string) string {
	if ptr.Pointer() != nil {
		var uriC *C.char
		if uri != "" {
			uriC = C.CString(uri)
			defer C.free(unsafe.Pointer(uriC))
		}
		return cGoUnpackString(C.QXmlNamespaceSupport_Prefix(ptr.Pointer(), C.struct_QtXml_PackedString{data: uriC, len: C.longlong(len(uri))}))
	}
	return ""
}

func (ptr *QXmlNamespaceSupport) Uri(prefix string) string {
	if ptr.Pointer() != nil {
		var prefixC *C.char
		if prefix != "" {
			prefixC = C.CString(prefix)
			defer C.free(unsafe.Pointer(prefixC))
		}
		return cGoUnpackString(C.QXmlNamespaceSupport_Uri(ptr.Pointer(), C.struct_QtXml_PackedString{data: prefixC, len: C.longlong(len(prefix))}))
	}
	return ""
}

func (ptr *QXmlNamespaceSupport) Prefixes() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QXmlNamespaceSupport_Prefixes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QXmlNamespaceSupport) Prefixes2(uri string) []string {
	if ptr.Pointer() != nil {
		var uriC *C.char
		if uri != "" {
			uriC = C.CString(uri)
			defer C.free(unsafe.Pointer(uriC))
		}
		return strings.Split(cGoUnpackString(C.QXmlNamespaceSupport_Prefixes2(ptr.Pointer(), C.struct_QtXml_PackedString{data: uriC, len: C.longlong(len(uri))})), "|")
	}
	return make([]string, 0)
}

func (ptr *QXmlNamespaceSupport) ProcessName(qname string, isAttribute bool, nsuri string, localname string) {
	if ptr.Pointer() != nil {
		var qnameC *C.char
		if qname != "" {
			qnameC = C.CString(qname)
			defer C.free(unsafe.Pointer(qnameC))
		}
		var nsuriC *C.char
		if nsuri != "" {
			nsuriC = C.CString(nsuri)
			defer C.free(unsafe.Pointer(nsuriC))
		}
		var localnameC *C.char
		if localname != "" {
			localnameC = C.CString(localname)
			defer C.free(unsafe.Pointer(localnameC))
		}
		C.QXmlNamespaceSupport_ProcessName(ptr.Pointer(), C.struct_QtXml_PackedString{data: qnameC, len: C.longlong(len(qname))}, C.char(int8(qt.GoBoolToInt(isAttribute))), C.struct_QtXml_PackedString{data: nsuriC, len: C.longlong(len(nsuri))}, C.struct_QtXml_PackedString{data: localnameC, len: C.longlong(len(localname))})
	}
}

func (ptr *QXmlNamespaceSupport) SplitName(qname string, prefix string, localname string) {
	if ptr.Pointer() != nil {
		var qnameC *C.char
		if qname != "" {
			qnameC = C.CString(qname)
			defer C.free(unsafe.Pointer(qnameC))
		}
		var prefixC *C.char
		if prefix != "" {
			prefixC = C.CString(prefix)
			defer C.free(unsafe.Pointer(prefixC))
		}
		var localnameC *C.char
		if localname != "" {
			localnameC = C.CString(localname)
			defer C.free(unsafe.Pointer(localnameC))
		}
		C.QXmlNamespaceSupport_SplitName(ptr.Pointer(), C.struct_QtXml_PackedString{data: qnameC, len: C.longlong(len(qname))}, C.struct_QtXml_PackedString{data: prefixC, len: C.longlong(len(prefix))}, C.struct_QtXml_PackedString{data: localnameC, len: C.longlong(len(localname))})
	}
}

type QXmlParseException struct {
	ptr unsafe.Pointer
}

type QXmlParseException_ITF interface {
	QXmlParseException_PTR() *QXmlParseException
}

func (ptr *QXmlParseException) QXmlParseException_PTR() *QXmlParseException {
	return ptr
}

func (ptr *QXmlParseException) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlParseException) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQXmlParseException(ptr QXmlParseException_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlParseException_PTR().Pointer()
	}
	return nil
}

func NewQXmlParseExceptionFromPointer(ptr unsafe.Pointer) (n *QXmlParseException) {
	n = new(QXmlParseException)
	n.SetPointer(ptr)
	return
}
func NewQXmlParseException(name string, c int, l int, p string, s string) *QXmlParseException {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	var pC *C.char
	if p != "" {
		pC = C.CString(p)
		defer C.free(unsafe.Pointer(pC))
	}
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	tmpValue := NewQXmlParseExceptionFromPointer(C.QXmlParseException_NewQXmlParseException(C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, C.int(int32(c)), C.int(int32(l)), C.struct_QtXml_PackedString{data: pC, len: C.longlong(len(p))}, C.struct_QtXml_PackedString{data: sC, len: C.longlong(len(s))}))
	runtime.SetFinalizer(tmpValue, (*QXmlParseException).DestroyQXmlParseException)
	return tmpValue
}

func NewQXmlParseException2(other QXmlParseException_ITF) *QXmlParseException {
	tmpValue := NewQXmlParseExceptionFromPointer(C.QXmlParseException_NewQXmlParseException2(PointerFromQXmlParseException(other)))
	runtime.SetFinalizer(tmpValue, (*QXmlParseException).DestroyQXmlParseException)
	return tmpValue
}

func (ptr *QXmlParseException) DestroyQXmlParseException() {
	if ptr.Pointer() != nil {
		C.QXmlParseException_DestroyQXmlParseException(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QXmlParseException) Message() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlParseException_Message(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlParseException) PublicId() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlParseException_PublicId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlParseException) SystemId() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlParseException_SystemId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlParseException) ColumnNumber() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QXmlParseException_ColumnNumber(ptr.Pointer())))
	}
	return 0
}

func (ptr *QXmlParseException) LineNumber() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QXmlParseException_LineNumber(ptr.Pointer())))
	}
	return 0
}

type QXmlReader struct {
	ptr unsafe.Pointer
}

type QXmlReader_ITF interface {
	QXmlReader_PTR() *QXmlReader
}

func (ptr *QXmlReader) QXmlReader_PTR() *QXmlReader {
	return ptr
}

func (ptr *QXmlReader) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlReader) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQXmlReader(ptr QXmlReader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlReader_PTR().Pointer()
	}
	return nil
}

func NewQXmlReaderFromPointer(ptr unsafe.Pointer) (n *QXmlReader) {
	n = new(QXmlReader)
	n.SetPointer(ptr)
	return
}

//export callbackQXmlReader_Parse
func callbackQXmlReader_Parse(ptr unsafe.Pointer, input unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "parse"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QXmlInputSource) bool)(NewQXmlInputSourceFromPointer(input)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlReader) ConnectParse(f func(input *QXmlInputSource) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "parse"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "parse", func(input *QXmlInputSource) bool {
				signal.(func(*QXmlInputSource) bool)(input)
				return f(input)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "parse", f)
		}
	}
}

func (ptr *QXmlReader) DisconnectParse() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "parse")
	}
}

func (ptr *QXmlReader) Parse(input QXmlInputSource_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlReader_Parse(ptr.Pointer(), PointerFromQXmlInputSource(input))) != 0
	}
	return false
}

//export callbackQXmlReader_SetContentHandler
func callbackQXmlReader_SetContentHandler(ptr unsafe.Pointer, handler unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setContentHandler"); signal != nil {
		signal.(func(*QXmlContentHandler))(NewQXmlContentHandlerFromPointer(handler))
	}

}

func (ptr *QXmlReader) ConnectSetContentHandler(f func(handler *QXmlContentHandler)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setContentHandler"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setContentHandler", func(handler *QXmlContentHandler) {
				signal.(func(*QXmlContentHandler))(handler)
				f(handler)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setContentHandler", f)
		}
	}
}

func (ptr *QXmlReader) DisconnectSetContentHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setContentHandler")
	}
}

func (ptr *QXmlReader) SetContentHandler(handler QXmlContentHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlReader_SetContentHandler(ptr.Pointer(), PointerFromQXmlContentHandler(handler))
	}
}

//export callbackQXmlReader_SetDTDHandler
func callbackQXmlReader_SetDTDHandler(ptr unsafe.Pointer, handler unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setDTDHandler"); signal != nil {
		signal.(func(*QXmlDTDHandler))(NewQXmlDTDHandlerFromPointer(handler))
	}

}

func (ptr *QXmlReader) ConnectSetDTDHandler(f func(handler *QXmlDTDHandler)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setDTDHandler"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setDTDHandler", func(handler *QXmlDTDHandler) {
				signal.(func(*QXmlDTDHandler))(handler)
				f(handler)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setDTDHandler", f)
		}
	}
}

func (ptr *QXmlReader) DisconnectSetDTDHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setDTDHandler")
	}
}

func (ptr *QXmlReader) SetDTDHandler(handler QXmlDTDHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlReader_SetDTDHandler(ptr.Pointer(), PointerFromQXmlDTDHandler(handler))
	}
}

//export callbackQXmlReader_SetDeclHandler
func callbackQXmlReader_SetDeclHandler(ptr unsafe.Pointer, handler unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setDeclHandler"); signal != nil {
		signal.(func(*QXmlDeclHandler))(NewQXmlDeclHandlerFromPointer(handler))
	}

}

func (ptr *QXmlReader) ConnectSetDeclHandler(f func(handler *QXmlDeclHandler)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setDeclHandler"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setDeclHandler", func(handler *QXmlDeclHandler) {
				signal.(func(*QXmlDeclHandler))(handler)
				f(handler)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setDeclHandler", f)
		}
	}
}

func (ptr *QXmlReader) DisconnectSetDeclHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setDeclHandler")
	}
}

func (ptr *QXmlReader) SetDeclHandler(handler QXmlDeclHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlReader_SetDeclHandler(ptr.Pointer(), PointerFromQXmlDeclHandler(handler))
	}
}

//export callbackQXmlReader_SetEntityResolver
func callbackQXmlReader_SetEntityResolver(ptr unsafe.Pointer, handler unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setEntityResolver"); signal != nil {
		signal.(func(*QXmlEntityResolver))(NewQXmlEntityResolverFromPointer(handler))
	}

}

func (ptr *QXmlReader) ConnectSetEntityResolver(f func(handler *QXmlEntityResolver)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setEntityResolver"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setEntityResolver", func(handler *QXmlEntityResolver) {
				signal.(func(*QXmlEntityResolver))(handler)
				f(handler)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setEntityResolver", f)
		}
	}
}

func (ptr *QXmlReader) DisconnectSetEntityResolver() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setEntityResolver")
	}
}

func (ptr *QXmlReader) SetEntityResolver(handler QXmlEntityResolver_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlReader_SetEntityResolver(ptr.Pointer(), PointerFromQXmlEntityResolver(handler))
	}
}

//export callbackQXmlReader_SetErrorHandler
func callbackQXmlReader_SetErrorHandler(ptr unsafe.Pointer, handler unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setErrorHandler"); signal != nil {
		signal.(func(*QXmlErrorHandler))(NewQXmlErrorHandlerFromPointer(handler))
	}

}

func (ptr *QXmlReader) ConnectSetErrorHandler(f func(handler *QXmlErrorHandler)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setErrorHandler"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setErrorHandler", func(handler *QXmlErrorHandler) {
				signal.(func(*QXmlErrorHandler))(handler)
				f(handler)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setErrorHandler", f)
		}
	}
}

func (ptr *QXmlReader) DisconnectSetErrorHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setErrorHandler")
	}
}

func (ptr *QXmlReader) SetErrorHandler(handler QXmlErrorHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlReader_SetErrorHandler(ptr.Pointer(), PointerFromQXmlErrorHandler(handler))
	}
}

//export callbackQXmlReader_SetFeature
func callbackQXmlReader_SetFeature(ptr unsafe.Pointer, name C.struct_QtXml_PackedString, value C.char) {
	if signal := qt.GetSignal(ptr, "setFeature"); signal != nil {
		signal.(func(string, bool))(cGoUnpackString(name), int8(value) != 0)
	}

}

func (ptr *QXmlReader) ConnectSetFeature(f func(name string, value bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFeature"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setFeature", func(name string, value bool) {
				signal.(func(string, bool))(name, value)
				f(name, value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFeature", f)
		}
	}
}

func (ptr *QXmlReader) DisconnectSetFeature() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFeature")
	}
}

func (ptr *QXmlReader) SetFeature(name string, value bool) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QXmlReader_SetFeature(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQXmlReader_SetLexicalHandler
func callbackQXmlReader_SetLexicalHandler(ptr unsafe.Pointer, handler unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setLexicalHandler"); signal != nil {
		signal.(func(*QXmlLexicalHandler))(NewQXmlLexicalHandlerFromPointer(handler))
	}

}

func (ptr *QXmlReader) ConnectSetLexicalHandler(f func(handler *QXmlLexicalHandler)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setLexicalHandler"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setLexicalHandler", func(handler *QXmlLexicalHandler) {
				signal.(func(*QXmlLexicalHandler))(handler)
				f(handler)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setLexicalHandler", f)
		}
	}
}

func (ptr *QXmlReader) DisconnectSetLexicalHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setLexicalHandler")
	}
}

func (ptr *QXmlReader) SetLexicalHandler(handler QXmlLexicalHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlReader_SetLexicalHandler(ptr.Pointer(), PointerFromQXmlLexicalHandler(handler))
	}
}

//export callbackQXmlReader_SetProperty
func callbackQXmlReader_SetProperty(ptr unsafe.Pointer, name C.struct_QtXml_PackedString, value unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setProperty"); signal != nil {
		signal.(func(string, unsafe.Pointer))(cGoUnpackString(name), value)
	}

}

func (ptr *QXmlReader) ConnectSetProperty(f func(name string, value unsafe.Pointer)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setProperty", func(name string, value unsafe.Pointer) {
				signal.(func(string, unsafe.Pointer))(name, value)
				f(name, value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setProperty", f)
		}
	}
}

func (ptr *QXmlReader) DisconnectSetProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setProperty")
	}
}

func (ptr *QXmlReader) SetProperty(name string, value unsafe.Pointer) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QXmlReader_SetProperty(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, value)
	}
}

//export callbackQXmlReader_DestroyQXmlReader
func callbackQXmlReader_DestroyQXmlReader(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QXmlReader"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlReaderFromPointer(ptr).DestroyQXmlReaderDefault()
	}
}

func (ptr *QXmlReader) ConnectDestroyQXmlReader(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QXmlReader"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QXmlReader", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QXmlReader", f)
		}
	}
}

func (ptr *QXmlReader) DisconnectDestroyQXmlReader() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QXmlReader")
	}
}

func (ptr *QXmlReader) DestroyQXmlReader() {
	if ptr.Pointer() != nil {
		C.QXmlReader_DestroyQXmlReader(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXmlReader) DestroyQXmlReaderDefault() {
	if ptr.Pointer() != nil {
		C.QXmlReader_DestroyQXmlReaderDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQXmlReader_ContentHandler
func callbackQXmlReader_ContentHandler(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "contentHandler"); signal != nil {
		return PointerFromQXmlContentHandler(signal.(func() *QXmlContentHandler)())
	}

	return PointerFromQXmlContentHandler(nil)
}

func (ptr *QXmlReader) ConnectContentHandler(f func() *QXmlContentHandler) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "contentHandler"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "contentHandler", func() *QXmlContentHandler {
				signal.(func() *QXmlContentHandler)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "contentHandler", f)
		}
	}
}

func (ptr *QXmlReader) DisconnectContentHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "contentHandler")
	}
}

func (ptr *QXmlReader) ContentHandler() *QXmlContentHandler {
	if ptr.Pointer() != nil {
		return NewQXmlContentHandlerFromPointer(C.QXmlReader_ContentHandler(ptr.Pointer()))
	}
	return nil
}

//export callbackQXmlReader_DTDHandler
func callbackQXmlReader_DTDHandler(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "DTDHandler"); signal != nil {
		return PointerFromQXmlDTDHandler(signal.(func() *QXmlDTDHandler)())
	}

	return PointerFromQXmlDTDHandler(nil)
}

func (ptr *QXmlReader) ConnectDTDHandler(f func() *QXmlDTDHandler) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "DTDHandler"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "DTDHandler", func() *QXmlDTDHandler {
				signal.(func() *QXmlDTDHandler)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "DTDHandler", f)
		}
	}
}

func (ptr *QXmlReader) DisconnectDTDHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "DTDHandler")
	}
}

func (ptr *QXmlReader) DTDHandler() *QXmlDTDHandler {
	if ptr.Pointer() != nil {
		return NewQXmlDTDHandlerFromPointer(C.QXmlReader_DTDHandler(ptr.Pointer()))
	}
	return nil
}

//export callbackQXmlReader_DeclHandler
func callbackQXmlReader_DeclHandler(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "declHandler"); signal != nil {
		return PointerFromQXmlDeclHandler(signal.(func() *QXmlDeclHandler)())
	}

	return PointerFromQXmlDeclHandler(nil)
}

func (ptr *QXmlReader) ConnectDeclHandler(f func() *QXmlDeclHandler) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "declHandler"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "declHandler", func() *QXmlDeclHandler {
				signal.(func() *QXmlDeclHandler)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "declHandler", f)
		}
	}
}

func (ptr *QXmlReader) DisconnectDeclHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "declHandler")
	}
}

func (ptr *QXmlReader) DeclHandler() *QXmlDeclHandler {
	if ptr.Pointer() != nil {
		return NewQXmlDeclHandlerFromPointer(C.QXmlReader_DeclHandler(ptr.Pointer()))
	}
	return nil
}

//export callbackQXmlReader_EntityResolver
func callbackQXmlReader_EntityResolver(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "entityResolver"); signal != nil {
		return PointerFromQXmlEntityResolver(signal.(func() *QXmlEntityResolver)())
	}

	return PointerFromQXmlEntityResolver(nil)
}

func (ptr *QXmlReader) ConnectEntityResolver(f func() *QXmlEntityResolver) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "entityResolver"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "entityResolver", func() *QXmlEntityResolver {
				signal.(func() *QXmlEntityResolver)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "entityResolver", f)
		}
	}
}

func (ptr *QXmlReader) DisconnectEntityResolver() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "entityResolver")
	}
}

func (ptr *QXmlReader) EntityResolver() *QXmlEntityResolver {
	if ptr.Pointer() != nil {
		return NewQXmlEntityResolverFromPointer(C.QXmlReader_EntityResolver(ptr.Pointer()))
	}
	return nil
}

//export callbackQXmlReader_ErrorHandler
func callbackQXmlReader_ErrorHandler(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "errorHandler"); signal != nil {
		return PointerFromQXmlErrorHandler(signal.(func() *QXmlErrorHandler)())
	}

	return PointerFromQXmlErrorHandler(nil)
}

func (ptr *QXmlReader) ConnectErrorHandler(f func() *QXmlErrorHandler) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "errorHandler"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "errorHandler", func() *QXmlErrorHandler {
				signal.(func() *QXmlErrorHandler)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "errorHandler", f)
		}
	}
}

func (ptr *QXmlReader) DisconnectErrorHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "errorHandler")
	}
}

func (ptr *QXmlReader) ErrorHandler() *QXmlErrorHandler {
	if ptr.Pointer() != nil {
		return NewQXmlErrorHandlerFromPointer(C.QXmlReader_ErrorHandler(ptr.Pointer()))
	}
	return nil
}

//export callbackQXmlReader_LexicalHandler
func callbackQXmlReader_LexicalHandler(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "lexicalHandler"); signal != nil {
		return PointerFromQXmlLexicalHandler(signal.(func() *QXmlLexicalHandler)())
	}

	return PointerFromQXmlLexicalHandler(nil)
}

func (ptr *QXmlReader) ConnectLexicalHandler(f func() *QXmlLexicalHandler) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "lexicalHandler"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "lexicalHandler", func() *QXmlLexicalHandler {
				signal.(func() *QXmlLexicalHandler)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lexicalHandler", f)
		}
	}
}

func (ptr *QXmlReader) DisconnectLexicalHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "lexicalHandler")
	}
}

func (ptr *QXmlReader) LexicalHandler() *QXmlLexicalHandler {
	if ptr.Pointer() != nil {
		return NewQXmlLexicalHandlerFromPointer(C.QXmlReader_LexicalHandler(ptr.Pointer()))
	}
	return nil
}

//export callbackQXmlReader_Feature
func callbackQXmlReader_Feature(ptr unsafe.Pointer, name C.struct_QtXml_PackedString, ok *C.char) C.char {
	okR := int8(*ok) != 0
	defer func() { *ok = C.char(int8(qt.GoBoolToInt(okR))) }()
	if signal := qt.GetSignal(ptr, "feature"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, *bool) bool)(cGoUnpackString(name), &okR))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlReader) ConnectFeature(f func(name string, ok *bool) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "feature"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "feature", func(name string, ok *bool) bool {
				signal.(func(string, *bool) bool)(name, ok)
				return f(name, ok)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "feature", f)
		}
	}
}

func (ptr *QXmlReader) DisconnectFeature() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "feature")
	}
}

func (ptr *QXmlReader) Feature(name string, ok *bool) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		okC := C.char(int8(qt.GoBoolToInt(*ok)))
		defer func() { *ok = int8(okC) != 0 }()
		return int8(C.QXmlReader_Feature(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, &okC)) != 0
	}
	return false
}

//export callbackQXmlReader_HasFeature
func callbackQXmlReader_HasFeature(ptr unsafe.Pointer, name C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "hasFeature"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(name)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlReader) ConnectHasFeature(f func(name string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hasFeature"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "hasFeature", func(name string) bool {
				signal.(func(string) bool)(name)
				return f(name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hasFeature", f)
		}
	}
}

func (ptr *QXmlReader) DisconnectHasFeature() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hasFeature")
	}
}

func (ptr *QXmlReader) HasFeature(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QXmlReader_HasFeature(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

//export callbackQXmlReader_HasProperty
func callbackQXmlReader_HasProperty(ptr unsafe.Pointer, name C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "hasProperty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(name)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QXmlReader) ConnectHasProperty(f func(name string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hasProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "hasProperty", func(name string) bool {
				signal.(func(string) bool)(name)
				return f(name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hasProperty", f)
		}
	}
}

func (ptr *QXmlReader) DisconnectHasProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hasProperty")
	}
}

func (ptr *QXmlReader) HasProperty(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QXmlReader_HasProperty(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

//export callbackQXmlReader_Property
func callbackQXmlReader_Property(ptr unsafe.Pointer, name C.struct_QtXml_PackedString, ok *C.char) unsafe.Pointer {
	okR := int8(*ok) != 0
	defer func() { *ok = C.char(int8(qt.GoBoolToInt(okR))) }()
	if signal := qt.GetSignal(ptr, "property"); signal != nil {
		return signal.(func(string, *bool) unsafe.Pointer)(cGoUnpackString(name), &okR)
	}

	return nil
}

func (ptr *QXmlReader) ConnectProperty(f func(name string, ok *bool) unsafe.Pointer) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "property"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "property", func(name string, ok *bool) unsafe.Pointer {
				signal.(func(string, *bool) unsafe.Pointer)(name, ok)
				return f(name, ok)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "property", f)
		}
	}
}

func (ptr *QXmlReader) DisconnectProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "property")
	}
}

func (ptr *QXmlReader) Property(name string, ok *bool) unsafe.Pointer {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		okC := C.char(int8(qt.GoBoolToInt(*ok)))
		defer func() { *ok = int8(okC) != 0 }()
		return C.QXmlReader_Property(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, &okC)
	}
	return nil
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

func NewQXmlSimpleReaderFromPointer(ptr unsafe.Pointer) (n *QXmlSimpleReader) {
	n = new(QXmlSimpleReader)
	n.SetPointer(ptr)
	return
}
func NewQXmlSimpleReader() *QXmlSimpleReader {
	return NewQXmlSimpleReaderFromPointer(C.QXmlSimpleReader_NewQXmlSimpleReader())
}

//export callbackQXmlSimpleReader_Parse
func callbackQXmlSimpleReader_Parse(ptr unsafe.Pointer, input unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "parse"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QXmlInputSource) bool)(NewQXmlInputSourceFromPointer(input)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlSimpleReaderFromPointer(ptr).ParseDefault(NewQXmlInputSourceFromPointer(input)))))
}

func (ptr *QXmlSimpleReader) ConnectParse(f func(input *QXmlInputSource) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "parse"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "parse", func(input *QXmlInputSource) bool {
				signal.(func(*QXmlInputSource) bool)(input)
				return f(input)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "parse", f)
		}
	}
}

func (ptr *QXmlSimpleReader) DisconnectParse() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "parse")
	}
}

func (ptr *QXmlSimpleReader) Parse(input QXmlInputSource_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlSimpleReader_Parse(ptr.Pointer(), PointerFromQXmlInputSource(input))) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) ParseDefault(input QXmlInputSource_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlSimpleReader_ParseDefault(ptr.Pointer(), PointerFromQXmlInputSource(input))) != 0
	}
	return false
}

//export callbackQXmlSimpleReader_Parse2
func callbackQXmlSimpleReader_Parse2(ptr unsafe.Pointer, input unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "parse2"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QXmlInputSource) bool)(NewQXmlInputSourceFromPointer(input)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlSimpleReaderFromPointer(ptr).Parse2Default(NewQXmlInputSourceFromPointer(input)))))
}

func (ptr *QXmlSimpleReader) ConnectParse2(f func(input *QXmlInputSource) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "parse2"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "parse2", func(input *QXmlInputSource) bool {
				signal.(func(*QXmlInputSource) bool)(input)
				return f(input)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "parse2", f)
		}
	}
}

func (ptr *QXmlSimpleReader) DisconnectParse2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "parse2")
	}
}

func (ptr *QXmlSimpleReader) Parse2(input QXmlInputSource_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlSimpleReader_Parse2(ptr.Pointer(), PointerFromQXmlInputSource(input))) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) Parse2Default(input QXmlInputSource_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlSimpleReader_Parse2Default(ptr.Pointer(), PointerFromQXmlInputSource(input))) != 0
	}
	return false
}

//export callbackQXmlSimpleReader_Parse3
func callbackQXmlSimpleReader_Parse3(ptr unsafe.Pointer, input unsafe.Pointer, incremental C.char) C.char {
	if signal := qt.GetSignal(ptr, "parse3"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QXmlInputSource, bool) bool)(NewQXmlInputSourceFromPointer(input), int8(incremental) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlSimpleReaderFromPointer(ptr).Parse3Default(NewQXmlInputSourceFromPointer(input), int8(incremental) != 0))))
}

func (ptr *QXmlSimpleReader) ConnectParse3(f func(input *QXmlInputSource, incremental bool) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "parse3"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "parse3", func(input *QXmlInputSource, incremental bool) bool {
				signal.(func(*QXmlInputSource, bool) bool)(input, incremental)
				return f(input, incremental)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "parse3", f)
		}
	}
}

func (ptr *QXmlSimpleReader) DisconnectParse3() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "parse3")
	}
}

func (ptr *QXmlSimpleReader) Parse3(input QXmlInputSource_ITF, incremental bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlSimpleReader_Parse3(ptr.Pointer(), PointerFromQXmlInputSource(input), C.char(int8(qt.GoBoolToInt(incremental))))) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) Parse3Default(input QXmlInputSource_ITF, incremental bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlSimpleReader_Parse3Default(ptr.Pointer(), PointerFromQXmlInputSource(input), C.char(int8(qt.GoBoolToInt(incremental))))) != 0
	}
	return false
}

//export callbackQXmlSimpleReader_ParseContinue
func callbackQXmlSimpleReader_ParseContinue(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "parseContinue"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlSimpleReaderFromPointer(ptr).ParseContinueDefault())))
}

func (ptr *QXmlSimpleReader) ConnectParseContinue(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "parseContinue"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "parseContinue", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "parseContinue", f)
		}
	}
}

func (ptr *QXmlSimpleReader) DisconnectParseContinue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "parseContinue")
	}
}

func (ptr *QXmlSimpleReader) ParseContinue() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlSimpleReader_ParseContinue(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) ParseContinueDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlSimpleReader_ParseContinueDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQXmlSimpleReader_SetContentHandler
func callbackQXmlSimpleReader_SetContentHandler(ptr unsafe.Pointer, handler unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setContentHandler"); signal != nil {
		signal.(func(*QXmlContentHandler))(NewQXmlContentHandlerFromPointer(handler))
	} else {
		NewQXmlSimpleReaderFromPointer(ptr).SetContentHandlerDefault(NewQXmlContentHandlerFromPointer(handler))
	}
}

func (ptr *QXmlSimpleReader) ConnectSetContentHandler(f func(handler *QXmlContentHandler)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setContentHandler"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setContentHandler", func(handler *QXmlContentHandler) {
				signal.(func(*QXmlContentHandler))(handler)
				f(handler)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setContentHandler", f)
		}
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetContentHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setContentHandler")
	}
}

func (ptr *QXmlSimpleReader) SetContentHandler(handler QXmlContentHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetContentHandler(ptr.Pointer(), PointerFromQXmlContentHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) SetContentHandlerDefault(handler QXmlContentHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetContentHandlerDefault(ptr.Pointer(), PointerFromQXmlContentHandler(handler))
	}
}

//export callbackQXmlSimpleReader_SetDTDHandler
func callbackQXmlSimpleReader_SetDTDHandler(ptr unsafe.Pointer, handler unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setDTDHandler"); signal != nil {
		signal.(func(*QXmlDTDHandler))(NewQXmlDTDHandlerFromPointer(handler))
	} else {
		NewQXmlSimpleReaderFromPointer(ptr).SetDTDHandlerDefault(NewQXmlDTDHandlerFromPointer(handler))
	}
}

func (ptr *QXmlSimpleReader) ConnectSetDTDHandler(f func(handler *QXmlDTDHandler)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setDTDHandler"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setDTDHandler", func(handler *QXmlDTDHandler) {
				signal.(func(*QXmlDTDHandler))(handler)
				f(handler)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setDTDHandler", f)
		}
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetDTDHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setDTDHandler")
	}
}

func (ptr *QXmlSimpleReader) SetDTDHandler(handler QXmlDTDHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetDTDHandler(ptr.Pointer(), PointerFromQXmlDTDHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) SetDTDHandlerDefault(handler QXmlDTDHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetDTDHandlerDefault(ptr.Pointer(), PointerFromQXmlDTDHandler(handler))
	}
}

//export callbackQXmlSimpleReader_SetDeclHandler
func callbackQXmlSimpleReader_SetDeclHandler(ptr unsafe.Pointer, handler unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setDeclHandler"); signal != nil {
		signal.(func(*QXmlDeclHandler))(NewQXmlDeclHandlerFromPointer(handler))
	} else {
		NewQXmlSimpleReaderFromPointer(ptr).SetDeclHandlerDefault(NewQXmlDeclHandlerFromPointer(handler))
	}
}

func (ptr *QXmlSimpleReader) ConnectSetDeclHandler(f func(handler *QXmlDeclHandler)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setDeclHandler"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setDeclHandler", func(handler *QXmlDeclHandler) {
				signal.(func(*QXmlDeclHandler))(handler)
				f(handler)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setDeclHandler", f)
		}
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetDeclHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setDeclHandler")
	}
}

func (ptr *QXmlSimpleReader) SetDeclHandler(handler QXmlDeclHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetDeclHandler(ptr.Pointer(), PointerFromQXmlDeclHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) SetDeclHandlerDefault(handler QXmlDeclHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetDeclHandlerDefault(ptr.Pointer(), PointerFromQXmlDeclHandler(handler))
	}
}

//export callbackQXmlSimpleReader_SetEntityResolver
func callbackQXmlSimpleReader_SetEntityResolver(ptr unsafe.Pointer, handler unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setEntityResolver"); signal != nil {
		signal.(func(*QXmlEntityResolver))(NewQXmlEntityResolverFromPointer(handler))
	} else {
		NewQXmlSimpleReaderFromPointer(ptr).SetEntityResolverDefault(NewQXmlEntityResolverFromPointer(handler))
	}
}

func (ptr *QXmlSimpleReader) ConnectSetEntityResolver(f func(handler *QXmlEntityResolver)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setEntityResolver"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setEntityResolver", func(handler *QXmlEntityResolver) {
				signal.(func(*QXmlEntityResolver))(handler)
				f(handler)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setEntityResolver", f)
		}
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetEntityResolver() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setEntityResolver")
	}
}

func (ptr *QXmlSimpleReader) SetEntityResolver(handler QXmlEntityResolver_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetEntityResolver(ptr.Pointer(), PointerFromQXmlEntityResolver(handler))
	}
}

func (ptr *QXmlSimpleReader) SetEntityResolverDefault(handler QXmlEntityResolver_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetEntityResolverDefault(ptr.Pointer(), PointerFromQXmlEntityResolver(handler))
	}
}

//export callbackQXmlSimpleReader_SetErrorHandler
func callbackQXmlSimpleReader_SetErrorHandler(ptr unsafe.Pointer, handler unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setErrorHandler"); signal != nil {
		signal.(func(*QXmlErrorHandler))(NewQXmlErrorHandlerFromPointer(handler))
	} else {
		NewQXmlSimpleReaderFromPointer(ptr).SetErrorHandlerDefault(NewQXmlErrorHandlerFromPointer(handler))
	}
}

func (ptr *QXmlSimpleReader) ConnectSetErrorHandler(f func(handler *QXmlErrorHandler)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setErrorHandler"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setErrorHandler", func(handler *QXmlErrorHandler) {
				signal.(func(*QXmlErrorHandler))(handler)
				f(handler)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setErrorHandler", f)
		}
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetErrorHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setErrorHandler")
	}
}

func (ptr *QXmlSimpleReader) SetErrorHandler(handler QXmlErrorHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetErrorHandler(ptr.Pointer(), PointerFromQXmlErrorHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) SetErrorHandlerDefault(handler QXmlErrorHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetErrorHandlerDefault(ptr.Pointer(), PointerFromQXmlErrorHandler(handler))
	}
}

//export callbackQXmlSimpleReader_SetFeature
func callbackQXmlSimpleReader_SetFeature(ptr unsafe.Pointer, name C.struct_QtXml_PackedString, enable C.char) {
	if signal := qt.GetSignal(ptr, "setFeature"); signal != nil {
		signal.(func(string, bool))(cGoUnpackString(name), int8(enable) != 0)
	} else {
		NewQXmlSimpleReaderFromPointer(ptr).SetFeatureDefault(cGoUnpackString(name), int8(enable) != 0)
	}
}

func (ptr *QXmlSimpleReader) ConnectSetFeature(f func(name string, enable bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFeature"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setFeature", func(name string, enable bool) {
				signal.(func(string, bool))(name, enable)
				f(name, enable)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFeature", f)
		}
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetFeature() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFeature")
	}
}

func (ptr *QXmlSimpleReader) SetFeature(name string, enable bool) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QXmlSimpleReader_SetFeature(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QXmlSimpleReader) SetFeatureDefault(name string, enable bool) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QXmlSimpleReader_SetFeatureDefault(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, C.char(int8(qt.GoBoolToInt(enable))))
	}
}

//export callbackQXmlSimpleReader_SetLexicalHandler
func callbackQXmlSimpleReader_SetLexicalHandler(ptr unsafe.Pointer, handler unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setLexicalHandler"); signal != nil {
		signal.(func(*QXmlLexicalHandler))(NewQXmlLexicalHandlerFromPointer(handler))
	} else {
		NewQXmlSimpleReaderFromPointer(ptr).SetLexicalHandlerDefault(NewQXmlLexicalHandlerFromPointer(handler))
	}
}

func (ptr *QXmlSimpleReader) ConnectSetLexicalHandler(f func(handler *QXmlLexicalHandler)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setLexicalHandler"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setLexicalHandler", func(handler *QXmlLexicalHandler) {
				signal.(func(*QXmlLexicalHandler))(handler)
				f(handler)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setLexicalHandler", f)
		}
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetLexicalHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setLexicalHandler")
	}
}

func (ptr *QXmlSimpleReader) SetLexicalHandler(handler QXmlLexicalHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetLexicalHandler(ptr.Pointer(), PointerFromQXmlLexicalHandler(handler))
	}
}

func (ptr *QXmlSimpleReader) SetLexicalHandlerDefault(handler QXmlLexicalHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_SetLexicalHandlerDefault(ptr.Pointer(), PointerFromQXmlLexicalHandler(handler))
	}
}

//export callbackQXmlSimpleReader_SetProperty
func callbackQXmlSimpleReader_SetProperty(ptr unsafe.Pointer, name C.struct_QtXml_PackedString, value unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setProperty"); signal != nil {
		signal.(func(string, unsafe.Pointer))(cGoUnpackString(name), value)
	} else {
		NewQXmlSimpleReaderFromPointer(ptr).SetPropertyDefault(cGoUnpackString(name), value)
	}
}

func (ptr *QXmlSimpleReader) ConnectSetProperty(f func(name string, value unsafe.Pointer)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setProperty", func(name string, value unsafe.Pointer) {
				signal.(func(string, unsafe.Pointer))(name, value)
				f(name, value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setProperty", f)
		}
	}
}

func (ptr *QXmlSimpleReader) DisconnectSetProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setProperty")
	}
}

func (ptr *QXmlSimpleReader) SetProperty(name string, value unsafe.Pointer) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QXmlSimpleReader_SetProperty(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, value)
	}
}

func (ptr *QXmlSimpleReader) SetPropertyDefault(name string, value unsafe.Pointer) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QXmlSimpleReader_SetPropertyDefault(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, value)
	}
}

//export callbackQXmlSimpleReader_DestroyQXmlSimpleReader
func callbackQXmlSimpleReader_DestroyQXmlSimpleReader(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QXmlSimpleReader"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlSimpleReaderFromPointer(ptr).DestroyQXmlSimpleReaderDefault()
	}
}

func (ptr *QXmlSimpleReader) ConnectDestroyQXmlSimpleReader(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QXmlSimpleReader"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QXmlSimpleReader", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QXmlSimpleReader", f)
		}
	}
}

func (ptr *QXmlSimpleReader) DisconnectDestroyQXmlSimpleReader() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QXmlSimpleReader")
	}
}

func (ptr *QXmlSimpleReader) DestroyQXmlSimpleReader() {
	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_DestroyQXmlSimpleReader(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXmlSimpleReader) DestroyQXmlSimpleReaderDefault() {
	if ptr.Pointer() != nil {
		C.QXmlSimpleReader_DestroyQXmlSimpleReaderDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQXmlSimpleReader_ContentHandler
func callbackQXmlSimpleReader_ContentHandler(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "contentHandler"); signal != nil {
		return PointerFromQXmlContentHandler(signal.(func() *QXmlContentHandler)())
	}

	return PointerFromQXmlContentHandler(NewQXmlSimpleReaderFromPointer(ptr).ContentHandlerDefault())
}

func (ptr *QXmlSimpleReader) ConnectContentHandler(f func() *QXmlContentHandler) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "contentHandler"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "contentHandler", func() *QXmlContentHandler {
				signal.(func() *QXmlContentHandler)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "contentHandler", f)
		}
	}
}

func (ptr *QXmlSimpleReader) DisconnectContentHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "contentHandler")
	}
}

func (ptr *QXmlSimpleReader) ContentHandler() *QXmlContentHandler {
	if ptr.Pointer() != nil {
		return NewQXmlContentHandlerFromPointer(C.QXmlSimpleReader_ContentHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSimpleReader) ContentHandlerDefault() *QXmlContentHandler {
	if ptr.Pointer() != nil {
		return NewQXmlContentHandlerFromPointer(C.QXmlSimpleReader_ContentHandlerDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQXmlSimpleReader_DTDHandler
func callbackQXmlSimpleReader_DTDHandler(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "DTDHandler"); signal != nil {
		return PointerFromQXmlDTDHandler(signal.(func() *QXmlDTDHandler)())
	}

	return PointerFromQXmlDTDHandler(NewQXmlSimpleReaderFromPointer(ptr).DTDHandlerDefault())
}

func (ptr *QXmlSimpleReader) ConnectDTDHandler(f func() *QXmlDTDHandler) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "DTDHandler"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "DTDHandler", func() *QXmlDTDHandler {
				signal.(func() *QXmlDTDHandler)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "DTDHandler", f)
		}
	}
}

func (ptr *QXmlSimpleReader) DisconnectDTDHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "DTDHandler")
	}
}

func (ptr *QXmlSimpleReader) DTDHandler() *QXmlDTDHandler {
	if ptr.Pointer() != nil {
		return NewQXmlDTDHandlerFromPointer(C.QXmlSimpleReader_DTDHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSimpleReader) DTDHandlerDefault() *QXmlDTDHandler {
	if ptr.Pointer() != nil {
		return NewQXmlDTDHandlerFromPointer(C.QXmlSimpleReader_DTDHandlerDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQXmlSimpleReader_DeclHandler
func callbackQXmlSimpleReader_DeclHandler(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "declHandler"); signal != nil {
		return PointerFromQXmlDeclHandler(signal.(func() *QXmlDeclHandler)())
	}

	return PointerFromQXmlDeclHandler(NewQXmlSimpleReaderFromPointer(ptr).DeclHandlerDefault())
}

func (ptr *QXmlSimpleReader) ConnectDeclHandler(f func() *QXmlDeclHandler) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "declHandler"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "declHandler", func() *QXmlDeclHandler {
				signal.(func() *QXmlDeclHandler)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "declHandler", f)
		}
	}
}

func (ptr *QXmlSimpleReader) DisconnectDeclHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "declHandler")
	}
}

func (ptr *QXmlSimpleReader) DeclHandler() *QXmlDeclHandler {
	if ptr.Pointer() != nil {
		return NewQXmlDeclHandlerFromPointer(C.QXmlSimpleReader_DeclHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSimpleReader) DeclHandlerDefault() *QXmlDeclHandler {
	if ptr.Pointer() != nil {
		return NewQXmlDeclHandlerFromPointer(C.QXmlSimpleReader_DeclHandlerDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQXmlSimpleReader_EntityResolver
func callbackQXmlSimpleReader_EntityResolver(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "entityResolver"); signal != nil {
		return PointerFromQXmlEntityResolver(signal.(func() *QXmlEntityResolver)())
	}

	return PointerFromQXmlEntityResolver(NewQXmlSimpleReaderFromPointer(ptr).EntityResolverDefault())
}

func (ptr *QXmlSimpleReader) ConnectEntityResolver(f func() *QXmlEntityResolver) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "entityResolver"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "entityResolver", func() *QXmlEntityResolver {
				signal.(func() *QXmlEntityResolver)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "entityResolver", f)
		}
	}
}

func (ptr *QXmlSimpleReader) DisconnectEntityResolver() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "entityResolver")
	}
}

func (ptr *QXmlSimpleReader) EntityResolver() *QXmlEntityResolver {
	if ptr.Pointer() != nil {
		return NewQXmlEntityResolverFromPointer(C.QXmlSimpleReader_EntityResolver(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSimpleReader) EntityResolverDefault() *QXmlEntityResolver {
	if ptr.Pointer() != nil {
		return NewQXmlEntityResolverFromPointer(C.QXmlSimpleReader_EntityResolverDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQXmlSimpleReader_ErrorHandler
func callbackQXmlSimpleReader_ErrorHandler(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "errorHandler"); signal != nil {
		return PointerFromQXmlErrorHandler(signal.(func() *QXmlErrorHandler)())
	}

	return PointerFromQXmlErrorHandler(NewQXmlSimpleReaderFromPointer(ptr).ErrorHandlerDefault())
}

func (ptr *QXmlSimpleReader) ConnectErrorHandler(f func() *QXmlErrorHandler) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "errorHandler"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "errorHandler", func() *QXmlErrorHandler {
				signal.(func() *QXmlErrorHandler)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "errorHandler", f)
		}
	}
}

func (ptr *QXmlSimpleReader) DisconnectErrorHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "errorHandler")
	}
}

func (ptr *QXmlSimpleReader) ErrorHandler() *QXmlErrorHandler {
	if ptr.Pointer() != nil {
		return NewQXmlErrorHandlerFromPointer(C.QXmlSimpleReader_ErrorHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSimpleReader) ErrorHandlerDefault() *QXmlErrorHandler {
	if ptr.Pointer() != nil {
		return NewQXmlErrorHandlerFromPointer(C.QXmlSimpleReader_ErrorHandlerDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQXmlSimpleReader_LexicalHandler
func callbackQXmlSimpleReader_LexicalHandler(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "lexicalHandler"); signal != nil {
		return PointerFromQXmlLexicalHandler(signal.(func() *QXmlLexicalHandler)())
	}

	return PointerFromQXmlLexicalHandler(NewQXmlSimpleReaderFromPointer(ptr).LexicalHandlerDefault())
}

func (ptr *QXmlSimpleReader) ConnectLexicalHandler(f func() *QXmlLexicalHandler) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "lexicalHandler"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "lexicalHandler", func() *QXmlLexicalHandler {
				signal.(func() *QXmlLexicalHandler)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lexicalHandler", f)
		}
	}
}

func (ptr *QXmlSimpleReader) DisconnectLexicalHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "lexicalHandler")
	}
}

func (ptr *QXmlSimpleReader) LexicalHandler() *QXmlLexicalHandler {
	if ptr.Pointer() != nil {
		return NewQXmlLexicalHandlerFromPointer(C.QXmlSimpleReader_LexicalHandler(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlSimpleReader) LexicalHandlerDefault() *QXmlLexicalHandler {
	if ptr.Pointer() != nil {
		return NewQXmlLexicalHandlerFromPointer(C.QXmlSimpleReader_LexicalHandlerDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQXmlSimpleReader_Feature
func callbackQXmlSimpleReader_Feature(ptr unsafe.Pointer, name C.struct_QtXml_PackedString, ok *C.char) C.char {
	okR := int8(*ok) != 0
	defer func() { *ok = C.char(int8(qt.GoBoolToInt(okR))) }()
	if signal := qt.GetSignal(ptr, "feature"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, *bool) bool)(cGoUnpackString(name), &okR))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlSimpleReaderFromPointer(ptr).FeatureDefault(cGoUnpackString(name), &okR))))
}

func (ptr *QXmlSimpleReader) ConnectFeature(f func(name string, ok *bool) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "feature"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "feature", func(name string, ok *bool) bool {
				signal.(func(string, *bool) bool)(name, ok)
				return f(name, ok)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "feature", f)
		}
	}
}

func (ptr *QXmlSimpleReader) DisconnectFeature() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "feature")
	}
}

func (ptr *QXmlSimpleReader) Feature(name string, ok *bool) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		okC := C.char(int8(qt.GoBoolToInt(*ok)))
		defer func() { *ok = int8(okC) != 0 }()
		return int8(C.QXmlSimpleReader_Feature(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, &okC)) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) FeatureDefault(name string, ok *bool) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		okC := C.char(int8(qt.GoBoolToInt(*ok)))
		defer func() { *ok = int8(okC) != 0 }()
		return int8(C.QXmlSimpleReader_FeatureDefault(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, &okC)) != 0
	}
	return false
}

//export callbackQXmlSimpleReader_HasFeature
func callbackQXmlSimpleReader_HasFeature(ptr unsafe.Pointer, name C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "hasFeature"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(name)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlSimpleReaderFromPointer(ptr).HasFeatureDefault(cGoUnpackString(name)))))
}

func (ptr *QXmlSimpleReader) ConnectHasFeature(f func(name string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hasFeature"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "hasFeature", func(name string) bool {
				signal.(func(string) bool)(name)
				return f(name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hasFeature", f)
		}
	}
}

func (ptr *QXmlSimpleReader) DisconnectHasFeature() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hasFeature")
	}
}

func (ptr *QXmlSimpleReader) HasFeature(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QXmlSimpleReader_HasFeature(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) HasFeatureDefault(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QXmlSimpleReader_HasFeatureDefault(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

//export callbackQXmlSimpleReader_HasProperty
func callbackQXmlSimpleReader_HasProperty(ptr unsafe.Pointer, name C.struct_QtXml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "hasProperty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(name)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQXmlSimpleReaderFromPointer(ptr).HasPropertyDefault(cGoUnpackString(name)))))
}

func (ptr *QXmlSimpleReader) ConnectHasProperty(f func(name string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hasProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "hasProperty", func(name string) bool {
				signal.(func(string) bool)(name)
				return f(name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hasProperty", f)
		}
	}
}

func (ptr *QXmlSimpleReader) DisconnectHasProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hasProperty")
	}
}

func (ptr *QXmlSimpleReader) HasProperty(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QXmlSimpleReader_HasProperty(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

func (ptr *QXmlSimpleReader) HasPropertyDefault(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QXmlSimpleReader_HasPropertyDefault(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

//export callbackQXmlSimpleReader_Property
func callbackQXmlSimpleReader_Property(ptr unsafe.Pointer, name C.struct_QtXml_PackedString, ok *C.char) unsafe.Pointer {
	okR := int8(*ok) != 0
	defer func() { *ok = C.char(int8(qt.GoBoolToInt(okR))) }()
	if signal := qt.GetSignal(ptr, "property"); signal != nil {
		return signal.(func(string, *bool) unsafe.Pointer)(cGoUnpackString(name), &okR)
	}

	return NewQXmlSimpleReaderFromPointer(ptr).PropertyDefault(cGoUnpackString(name), &okR)
}

func (ptr *QXmlSimpleReader) ConnectProperty(f func(name string, ok *bool) unsafe.Pointer) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "property"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "property", func(name string, ok *bool) unsafe.Pointer {
				signal.(func(string, *bool) unsafe.Pointer)(name, ok)
				return f(name, ok)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "property", f)
		}
	}
}

func (ptr *QXmlSimpleReader) DisconnectProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "property")
	}
}

func (ptr *QXmlSimpleReader) Property(name string, ok *bool) unsafe.Pointer {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		okC := C.char(int8(qt.GoBoolToInt(*ok)))
		defer func() { *ok = int8(okC) != 0 }()
		return C.QXmlSimpleReader_Property(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, &okC)
	}
	return nil
}

func (ptr *QXmlSimpleReader) PropertyDefault(name string, ok *bool) unsafe.Pointer {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		okC := C.char(int8(qt.GoBoolToInt(*ok)))
		defer func() { *ok = int8(okC) != 0 }()
		return C.QXmlSimpleReader_PropertyDefault(ptr.Pointer(), C.struct_QtXml_PackedString{data: nameC, len: C.longlong(len(name))}, &okC)
	}
	return nil
}
