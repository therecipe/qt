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

type QTranslatorITF interface {
	QObjectITF
	QTranslatorPTR() *QTranslator
}

func PointerFromQTranslator(ptr QTranslatorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTranslatorPTR().Pointer()
	}
	return nil
}

func QTranslatorFromPointer(ptr unsafe.Pointer) *QTranslator {
	var n = new(QTranslator)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTranslator_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTranslator) QTranslatorPTR() *QTranslator {
	return ptr
}

func NewQTranslator(parent QObjectITF) *QTranslator {
	return QTranslatorFromPointer(unsafe.Pointer(C.QTranslator_NewQTranslator(C.QtObjectPtr(PointerFromQObject(parent)))))
}

func (ptr *QTranslator) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QTranslator_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTranslator) Load2(locale QLocaleITF, filename string, prefix string, directory string, suffix string) bool {
	if ptr.Pointer() != nil {
		return C.QTranslator_Load2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLocale(locale)), C.CString(filename), C.CString(prefix), C.CString(directory), C.CString(suffix)) != 0
	}
	return false
}

func (ptr *QTranslator) Load(filename string, directory string, search_delimiters string, suffix string) bool {
	if ptr.Pointer() != nil {
		return C.QTranslator_Load(C.QtObjectPtr(ptr.Pointer()), C.CString(filename), C.CString(directory), C.CString(search_delimiters), C.CString(suffix)) != 0
	}
	return false
}

func (ptr *QTranslator) Translate(context string, sourceText string, disambiguation string, n int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTranslator_Translate(C.QtObjectPtr(ptr.Pointer()), C.CString(context), C.CString(sourceText), C.CString(disambiguation), C.int(n)))
	}
	return ""
}

func (ptr *QTranslator) DestroyQTranslator() {
	if ptr.Pointer() != nil {
		C.QTranslator_DestroyQTranslator(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
