package xml

//#include "xml.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

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

func (ptr *QXmlAttributes) QXmlAttributes_PTR() *QXmlAttributes {
	return ptr
}

func NewQXmlAttributes() *QXmlAttributes {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlAttributes::QXmlAttributes")
		}
	}()

	return NewQXmlAttributesFromPointer(C.QXmlAttributes_NewQXmlAttributes())
}

func (ptr *QXmlAttributes) DestroyQXmlAttributes() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlAttributes::~QXmlAttributes")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlAttributes_DestroyQXmlAttributes(ptr.Pointer())
	}
}

func (ptr *QXmlAttributes) Append(qName string, uri string, localPart string, value string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlAttributes::append")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlAttributes_Append(ptr.Pointer(), C.CString(qName), C.CString(uri), C.CString(localPart), C.CString(value))
	}
}

func (ptr *QXmlAttributes) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlAttributes::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlAttributes_Clear(ptr.Pointer())
	}
}

func (ptr *QXmlAttributes) Count() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlAttributes::count")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QXmlAttributes_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlAttributes) Index2(qName core.QLatin1String_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlAttributes::index")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QXmlAttributes_Index2(ptr.Pointer(), core.PointerFromQLatin1String(qName)))
	}
	return 0
}

func (ptr *QXmlAttributes) Index(qName string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlAttributes::index")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QXmlAttributes_Index(ptr.Pointer(), C.CString(qName)))
	}
	return 0
}

func (ptr *QXmlAttributes) Index3(uri string, localPart string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlAttributes::index")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QXmlAttributes_Index3(ptr.Pointer(), C.CString(uri), C.CString(localPart)))
	}
	return 0
}

func (ptr *QXmlAttributes) Length() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlAttributes::length")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QXmlAttributes_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlAttributes) LocalName(index int) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlAttributes::localName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_LocalName(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QXmlAttributes) QName(index int) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlAttributes::qName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_QName(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QXmlAttributes) Type2(qName string) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlAttributes::type")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_Type2(ptr.Pointer(), C.CString(qName)))
	}
	return ""
}

func (ptr *QXmlAttributes) Type3(uri string, localName string) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlAttributes::type")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_Type3(ptr.Pointer(), C.CString(uri), C.CString(localName)))
	}
	return ""
}

func (ptr *QXmlAttributes) Type(index int) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlAttributes::type")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_Type(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QXmlAttributes) Uri(index int) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlAttributes::uri")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_Uri(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QXmlAttributes) Value3(qName core.QLatin1String_ITF) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlAttributes::value")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_Value3(ptr.Pointer(), core.PointerFromQLatin1String(qName)))
	}
	return ""
}

func (ptr *QXmlAttributes) Value2(qName string) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlAttributes::value")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_Value2(ptr.Pointer(), C.CString(qName)))
	}
	return ""
}

func (ptr *QXmlAttributes) Value4(uri string, localName string) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlAttributes::value")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_Value4(ptr.Pointer(), C.CString(uri), C.CString(localName)))
	}
	return ""
}

func (ptr *QXmlAttributes) Value(index int) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlAttributes::value")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlAttributes_Value(ptr.Pointer(), C.int(index)))
	}
	return ""
}
