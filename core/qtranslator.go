package core

//#include "qtranslator.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	if n.ObjectName() == "" {
		n.SetObjectName("QTranslator_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTranslator) QTranslator_PTR() *QTranslator {
	return ptr
}

func NewQTranslator(parent QObject_ITF) *QTranslator {
	return NewQTranslatorFromPointer(C.QTranslator_NewQTranslator(PointerFromQObject(parent)))
}

func (ptr *QTranslator) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QTranslator_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTranslator) Load2(locale QLocale_ITF, filename string, prefix string, directory string, suffix string) bool {
	if ptr.Pointer() != nil {
		return C.QTranslator_Load2(ptr.Pointer(), PointerFromQLocale(locale), C.CString(filename), C.CString(prefix), C.CString(directory), C.CString(suffix)) != 0
	}
	return false
}

func (ptr *QTranslator) Load(filename string, directory string, search_delimiters string, suffix string) bool {
	if ptr.Pointer() != nil {
		return C.QTranslator_Load(ptr.Pointer(), C.CString(filename), C.CString(directory), C.CString(search_delimiters), C.CString(suffix)) != 0
	}
	return false
}

func (ptr *QTranslator) Translate(context string, sourceText string, disambiguation string, n int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTranslator_Translate(ptr.Pointer(), C.CString(context), C.CString(sourceText), C.CString(disambiguation), C.int(n)))
	}
	return ""
}

func (ptr *QTranslator) DestroyQTranslator() {
	if ptr.Pointer() != nil {
		C.QTranslator_DestroyQTranslator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
