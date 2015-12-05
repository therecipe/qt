package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QTranslator struct {
	QObject
}

type QTranslator_ITF interface {
	QObject_ITF
	QTranslator_PTR() *QTranslator
}

func PointerFromQTranslator(ptr QTranslator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTranslator_PTR().Pointer()
	}
	return nil
}

func NewQTranslatorFromPointer(ptr unsafe.Pointer) *QTranslator {
	var n = new(QTranslator)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTranslator_") {
		n.SetObjectName("QTranslator_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTranslator) QTranslator_PTR() *QTranslator {
	return ptr
}

func NewQTranslator(parent QObject_ITF) *QTranslator {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTranslator::QTranslator")
		}
	}()

	return NewQTranslatorFromPointer(C.QTranslator_NewQTranslator(PointerFromQObject(parent)))
}

func (ptr *QTranslator) IsEmpty() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTranslator::isEmpty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTranslator_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTranslator) Load2(locale QLocale_ITF, filename string, prefix string, directory string, suffix string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTranslator::load")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTranslator_Load2(ptr.Pointer(), PointerFromQLocale(locale), C.CString(filename), C.CString(prefix), C.CString(directory), C.CString(suffix)) != 0
	}
	return false
}

func (ptr *QTranslator) Load(filename string, directory string, search_delimiters string, suffix string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTranslator::load")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTranslator_Load(ptr.Pointer(), C.CString(filename), C.CString(directory), C.CString(search_delimiters), C.CString(suffix)) != 0
	}
	return false
}

func (ptr *QTranslator) Translate(context string, sourceText string, disambiguation string, n int) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTranslator::translate")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTranslator_Translate(ptr.Pointer(), C.CString(context), C.CString(sourceText), C.CString(disambiguation), C.int(n)))
	}
	return ""
}

func (ptr *QTranslator) DestroyQTranslator() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTranslator::~QTranslator")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTranslator_DestroyQTranslator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
