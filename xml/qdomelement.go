package xml

//#include "xml.h"
import "C"
import (
	"log"
	"unsafe"
)

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

func (ptr *QDomElement) QDomElement_PTR() *QDomElement {
	return ptr
}

func NewQDomElement() *QDomElement {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomElement::QDomElement")
		}
	}()

	return NewQDomElementFromPointer(C.QDomElement_NewQDomElement())
}

func NewQDomElement2(x QDomElement_ITF) *QDomElement {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomElement::QDomElement")
		}
	}()

	return NewQDomElementFromPointer(C.QDomElement_NewQDomElement2(PointerFromQDomElement(x)))
}

func (ptr *QDomElement) Attribute(name string, defValue string) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomElement::attribute")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomElement_Attribute(ptr.Pointer(), C.CString(name), C.CString(defValue)))
	}
	return ""
}

func (ptr *QDomElement) AttributeNS(nsURI string, localName string, defValue string) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomElement::attributeNS")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomElement_AttributeNS(ptr.Pointer(), C.CString(nsURI), C.CString(localName), C.CString(defValue)))
	}
	return ""
}

func (ptr *QDomElement) HasAttribute(name string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomElement::hasAttribute")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QDomElement_HasAttribute(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QDomElement) HasAttributeNS(nsURI string, localName string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomElement::hasAttributeNS")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QDomElement_HasAttributeNS(ptr.Pointer(), C.CString(nsURI), C.CString(localName)) != 0
	}
	return false
}

func (ptr *QDomElement) NodeType() QDomNode__NodeType {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomElement::nodeType")
		}
	}()

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomElement_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomElement) RemoveAttribute(name string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomElement::removeAttribute")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDomElement_RemoveAttribute(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QDomElement) RemoveAttributeNS(nsURI string, localName string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomElement::removeAttributeNS")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDomElement_RemoveAttributeNS(ptr.Pointer(), C.CString(nsURI), C.CString(localName))
	}
}

func (ptr *QDomElement) SetAttribute(name string, value string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomElement::setAttribute")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDomElement_SetAttribute(ptr.Pointer(), C.CString(name), C.CString(value))
	}
}

func (ptr *QDomElement) SetAttribute2(name string, value int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomElement::setAttribute")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDomElement_SetAttribute2(ptr.Pointer(), C.CString(name), C.int(value))
	}
}

func (ptr *QDomElement) SetAttributeNS(nsURI string, qName string, value string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomElement::setAttributeNS")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDomElement_SetAttributeNS(ptr.Pointer(), C.CString(nsURI), C.CString(qName), C.CString(value))
	}
}

func (ptr *QDomElement) SetAttributeNS2(nsURI string, qName string, value int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomElement::setAttributeNS")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDomElement_SetAttributeNS2(ptr.Pointer(), C.CString(nsURI), C.CString(qName), C.int(value))
	}
}

func (ptr *QDomElement) SetTagName(name string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomElement::setTagName")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDomElement_SetTagName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QDomElement) TagName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomElement::tagName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomElement_TagName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomElement) Text() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomElement::text")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomElement_Text(ptr.Pointer()))
	}
	return ""
}
