package gui

//#include "qfontdatabase.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QFontDatabase struct {
	ptr unsafe.Pointer
}

type QFontDatabaseITF interface {
	QFontDatabasePTR() *QFontDatabase
}

func (p *QFontDatabase) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QFontDatabase) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQFontDatabase(ptr QFontDatabaseITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFontDatabasePTR().Pointer()
	}
	return nil
}

func QFontDatabaseFromPointer(ptr unsafe.Pointer) *QFontDatabase {
	var n = new(QFontDatabase)
	n.SetPointer(ptr)
	return n
}

func (ptr *QFontDatabase) QFontDatabasePTR() *QFontDatabase {
	return ptr
}

//QFontDatabase::SystemFont
type QFontDatabase__SystemFont int

var (
	QFontDatabase__GeneralFont          = QFontDatabase__SystemFont(0)
	QFontDatabase__FixedFont            = QFontDatabase__SystemFont(1)
	QFontDatabase__TitleFont            = QFontDatabase__SystemFont(2)
	QFontDatabase__SmallestReadableFont = QFontDatabase__SystemFont(3)
)

//QFontDatabase::WritingSystem
type QFontDatabase__WritingSystem int

var (
	QFontDatabase__Any                 = QFontDatabase__WritingSystem(0)
	QFontDatabase__Latin               = QFontDatabase__WritingSystem(1)
	QFontDatabase__Greek               = QFontDatabase__WritingSystem(2)
	QFontDatabase__Cyrillic            = QFontDatabase__WritingSystem(3)
	QFontDatabase__Armenian            = QFontDatabase__WritingSystem(4)
	QFontDatabase__Hebrew              = QFontDatabase__WritingSystem(5)
	QFontDatabase__Arabic              = QFontDatabase__WritingSystem(6)
	QFontDatabase__Syriac              = QFontDatabase__WritingSystem(7)
	QFontDatabase__Thaana              = QFontDatabase__WritingSystem(8)
	QFontDatabase__Devanagari          = QFontDatabase__WritingSystem(9)
	QFontDatabase__Bengali             = QFontDatabase__WritingSystem(10)
	QFontDatabase__Gurmukhi            = QFontDatabase__WritingSystem(11)
	QFontDatabase__Gujarati            = QFontDatabase__WritingSystem(12)
	QFontDatabase__Oriya               = QFontDatabase__WritingSystem(13)
	QFontDatabase__Tamil               = QFontDatabase__WritingSystem(14)
	QFontDatabase__Telugu              = QFontDatabase__WritingSystem(15)
	QFontDatabase__Kannada             = QFontDatabase__WritingSystem(16)
	QFontDatabase__Malayalam           = QFontDatabase__WritingSystem(17)
	QFontDatabase__Sinhala             = QFontDatabase__WritingSystem(18)
	QFontDatabase__Thai                = QFontDatabase__WritingSystem(19)
	QFontDatabase__Lao                 = QFontDatabase__WritingSystem(20)
	QFontDatabase__Tibetan             = QFontDatabase__WritingSystem(21)
	QFontDatabase__Myanmar             = QFontDatabase__WritingSystem(22)
	QFontDatabase__Georgian            = QFontDatabase__WritingSystem(23)
	QFontDatabase__Khmer               = QFontDatabase__WritingSystem(24)
	QFontDatabase__SimplifiedChinese   = QFontDatabase__WritingSystem(25)
	QFontDatabase__TraditionalChinese  = QFontDatabase__WritingSystem(26)
	QFontDatabase__Japanese            = QFontDatabase__WritingSystem(27)
	QFontDatabase__Korean              = QFontDatabase__WritingSystem(28)
	QFontDatabase__Vietnamese          = QFontDatabase__WritingSystem(29)
	QFontDatabase__Symbol              = QFontDatabase__WritingSystem(30)
	QFontDatabase__Other               = QFontDatabase__WritingSystem(QFontDatabase__Symbol)
	QFontDatabase__Ogham               = QFontDatabase__WritingSystem(C.QFontDatabase_Ogham_Type())
	QFontDatabase__Runic               = QFontDatabase__WritingSystem(C.QFontDatabase_Runic_Type())
	QFontDatabase__Nko                 = QFontDatabase__WritingSystem(C.QFontDatabase_Nko_Type())
	QFontDatabase__WritingSystemsCount = QFontDatabase__WritingSystem(C.QFontDatabase_WritingSystemsCount_Type())
)

func QFontDatabase_RemoveAllApplicationFonts() bool {
	return C.QFontDatabase_QFontDatabase_RemoveAllApplicationFonts() != 0
}

func QFontDatabase_RemoveApplicationFont(id int) bool {
	return C.QFontDatabase_QFontDatabase_RemoveApplicationFont(C.int(id)) != 0
}

func NewQFontDatabase() *QFontDatabase {
	return QFontDatabaseFromPointer(unsafe.Pointer(C.QFontDatabase_NewQFontDatabase()))
}

func QFontDatabase_AddApplicationFont(fileName string) int {
	return int(C.QFontDatabase_QFontDatabase_AddApplicationFont(C.CString(fileName)))
}

func QFontDatabase_AddApplicationFontFromData(fontData core.QByteArrayITF) int {
	return int(C.QFontDatabase_QFontDatabase_AddApplicationFontFromData(C.QtObjectPtr(core.PointerFromQByteArray(fontData))))
}

func QFontDatabase_ApplicationFontFamilies(id int) []string {
	return strings.Split(C.GoString(C.QFontDatabase_QFontDatabase_ApplicationFontFamilies(C.int(id))), "|")
}

func (ptr *QFontDatabase) Bold(family string, style string) bool {
	if ptr.Pointer() != nil {
		return C.QFontDatabase_Bold(C.QtObjectPtr(ptr.Pointer()), C.CString(family), C.CString(style)) != 0
	}
	return false
}

func (ptr *QFontDatabase) Families(writingSystem QFontDatabase__WritingSystem) []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFontDatabase_Families(C.QtObjectPtr(ptr.Pointer()), C.int(writingSystem))), "|")
	}
	return make([]string, 0)
}

func (ptr *QFontDatabase) IsBitmapScalable(family string, style string) bool {
	if ptr.Pointer() != nil {
		return C.QFontDatabase_IsBitmapScalable(C.QtObjectPtr(ptr.Pointer()), C.CString(family), C.CString(style)) != 0
	}
	return false
}

func (ptr *QFontDatabase) IsFixedPitch(family string, style string) bool {
	if ptr.Pointer() != nil {
		return C.QFontDatabase_IsFixedPitch(C.QtObjectPtr(ptr.Pointer()), C.CString(family), C.CString(style)) != 0
	}
	return false
}

func (ptr *QFontDatabase) IsPrivateFamily(family string) bool {
	if ptr.Pointer() != nil {
		return C.QFontDatabase_IsPrivateFamily(C.QtObjectPtr(ptr.Pointer()), C.CString(family)) != 0
	}
	return false
}

func (ptr *QFontDatabase) IsScalable(family string, style string) bool {
	if ptr.Pointer() != nil {
		return C.QFontDatabase_IsScalable(C.QtObjectPtr(ptr.Pointer()), C.CString(family), C.CString(style)) != 0
	}
	return false
}

func (ptr *QFontDatabase) IsSmoothlyScalable(family string, style string) bool {
	if ptr.Pointer() != nil {
		return C.QFontDatabase_IsSmoothlyScalable(C.QtObjectPtr(ptr.Pointer()), C.CString(family), C.CString(style)) != 0
	}
	return false
}

func (ptr *QFontDatabase) Italic(family string, style string) bool {
	if ptr.Pointer() != nil {
		return C.QFontDatabase_Italic(C.QtObjectPtr(ptr.Pointer()), C.CString(family), C.CString(style)) != 0
	}
	return false
}

func (ptr *QFontDatabase) StyleString(font QFontITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFontDatabase_StyleString(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQFont(font))))
	}
	return ""
}

func (ptr *QFontDatabase) StyleString2(fontInfo QFontInfoITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFontDatabase_StyleString2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQFontInfo(fontInfo))))
	}
	return ""
}

func (ptr *QFontDatabase) Styles(family string) []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFontDatabase_Styles(C.QtObjectPtr(ptr.Pointer()), C.CString(family))), "|")
	}
	return make([]string, 0)
}

func (ptr *QFontDatabase) Weight(family string, style string) int {
	if ptr.Pointer() != nil {
		return int(C.QFontDatabase_Weight(C.QtObjectPtr(ptr.Pointer()), C.CString(family), C.CString(style)))
	}
	return 0
}

func QFontDatabase_WritingSystemName(writingSystem QFontDatabase__WritingSystem) string {
	return C.GoString(C.QFontDatabase_QFontDatabase_WritingSystemName(C.int(writingSystem)))
}

func QFontDatabase_WritingSystemSample(writingSystem QFontDatabase__WritingSystem) string {
	return C.GoString(C.QFontDatabase_QFontDatabase_WritingSystemSample(C.int(writingSystem)))
}
