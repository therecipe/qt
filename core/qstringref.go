package core

//#include "qstringref.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QStringRef struct {
	ptr unsafe.Pointer
}

type QStringRefITF interface {
	QStringRefPTR() *QStringRef
}

func (p *QStringRef) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QStringRef) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQStringRef(ptr QStringRefITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStringRefPTR().Pointer()
	}
	return nil
}

func QStringRefFromPointer(ptr unsafe.Pointer) *QStringRef {
	var n = new(QStringRef)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStringRef) QStringRefPTR() *QStringRef {
	return ptr
}

func (ptr *QStringRef) Begin() *QChar {
	if ptr.Pointer() != nil {
		return QCharFromPointer(unsafe.Pointer(C.QStringRef_Begin(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QStringRef) Cbegin() *QChar {
	if ptr.Pointer() != nil {
		return QCharFromPointer(unsafe.Pointer(C.QStringRef_Cbegin(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QStringRef) Cend() *QChar {
	if ptr.Pointer() != nil {
		return QCharFromPointer(unsafe.Pointer(C.QStringRef_Cend(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QStringRef) Clear() {
	if ptr.Pointer() != nil {
		C.QStringRef_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func QStringRef_Compare3(s1 QStringRefITF, s2 QLatin1StringITF, cs Qt__CaseSensitivity) int {
	return int(C.QStringRef_QStringRef_Compare3(C.QtObjectPtr(PointerFromQStringRef(s1)), C.QtObjectPtr(PointerFromQLatin1String(s2)), C.int(cs)))
}

func QStringRef_Compare(s1 QStringRefITF, s2 string, cs Qt__CaseSensitivity) int {
	return int(C.QStringRef_QStringRef_Compare(C.QtObjectPtr(PointerFromQStringRef(s1)), C.CString(s2), C.int(cs)))
}

func QStringRef_Compare2(s1 QStringRefITF, s2 QStringRefITF, cs Qt__CaseSensitivity) int {
	return int(C.QStringRef_QStringRef_Compare2(C.QtObjectPtr(PointerFromQStringRef(s1)), C.QtObjectPtr(PointerFromQStringRef(s2)), C.int(cs)))
}

func (ptr *QStringRef) Compare6(other QLatin1StringITF, cs Qt__CaseSensitivity) int {
	if ptr.Pointer() != nil {
		return int(C.QStringRef_Compare6(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLatin1String(other)), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) Compare4(other string, cs Qt__CaseSensitivity) int {
	if ptr.Pointer() != nil {
		return int(C.QStringRef_Compare4(C.QtObjectPtr(ptr.Pointer()), C.CString(other), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) Compare5(other QStringRefITF, cs Qt__CaseSensitivity) int {
	if ptr.Pointer() != nil {
		return int(C.QStringRef_Compare5(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQStringRef(other)), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) ConstData() *QChar {
	if ptr.Pointer() != nil {
		return QCharFromPointer(unsafe.Pointer(C.QStringRef_ConstData(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QStringRef) Contains2(ch QCharITF, cs Qt__CaseSensitivity) bool {
	if ptr.Pointer() != nil {
		return C.QStringRef_Contains2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQChar(ch)), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) Contains4(str QLatin1StringITF, cs Qt__CaseSensitivity) bool {
	if ptr.Pointer() != nil {
		return C.QStringRef_Contains4(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLatin1String(str)), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) Contains(str string, cs Qt__CaseSensitivity) bool {
	if ptr.Pointer() != nil {
		return C.QStringRef_Contains(C.QtObjectPtr(ptr.Pointer()), C.CString(str), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) Contains3(str QStringRefITF, cs Qt__CaseSensitivity) bool {
	if ptr.Pointer() != nil {
		return C.QStringRef_Contains3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQStringRef(str)), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QStringRef_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStringRef) Count3(ch QCharITF, cs Qt__CaseSensitivity) int {
	if ptr.Pointer() != nil {
		return int(C.QStringRef_Count3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQChar(ch)), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) Count2(str string, cs Qt__CaseSensitivity) int {
	if ptr.Pointer() != nil {
		return int(C.QStringRef_Count2(C.QtObjectPtr(ptr.Pointer()), C.CString(str), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) Count4(str QStringRefITF, cs Qt__CaseSensitivity) int {
	if ptr.Pointer() != nil {
		return int(C.QStringRef_Count4(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQStringRef(str)), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) Data() *QChar {
	if ptr.Pointer() != nil {
		return QCharFromPointer(unsafe.Pointer(C.QStringRef_Data(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QStringRef) End() *QChar {
	if ptr.Pointer() != nil {
		return QCharFromPointer(unsafe.Pointer(C.QStringRef_End(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QStringRef) EndsWith2(ch QCharITF, cs Qt__CaseSensitivity) bool {
	if ptr.Pointer() != nil {
		return C.QStringRef_EndsWith2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQChar(ch)), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) EndsWith3(str QLatin1StringITF, cs Qt__CaseSensitivity) bool {
	if ptr.Pointer() != nil {
		return C.QStringRef_EndsWith3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLatin1String(str)), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) EndsWith(str string, cs Qt__CaseSensitivity) bool {
	if ptr.Pointer() != nil {
		return C.QStringRef_EndsWith(C.QtObjectPtr(ptr.Pointer()), C.CString(str), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) EndsWith4(str QStringRefITF, cs Qt__CaseSensitivity) bool {
	if ptr.Pointer() != nil {
		return C.QStringRef_EndsWith4(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQStringRef(str)), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) IndexOf3(ch QCharITF, from int, cs Qt__CaseSensitivity) int {
	if ptr.Pointer() != nil {
		return int(C.QStringRef_IndexOf3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQChar(ch)), C.int(from), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) IndexOf2(str QLatin1StringITF, from int, cs Qt__CaseSensitivity) int {
	if ptr.Pointer() != nil {
		return int(C.QStringRef_IndexOf2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLatin1String(str)), C.int(from), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) IndexOf(str string, from int, cs Qt__CaseSensitivity) int {
	if ptr.Pointer() != nil {
		return int(C.QStringRef_IndexOf(C.QtObjectPtr(ptr.Pointer()), C.CString(str), C.int(from), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) IndexOf4(str QStringRefITF, from int, cs Qt__CaseSensitivity) int {
	if ptr.Pointer() != nil {
		return int(C.QStringRef_IndexOf4(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQStringRef(str)), C.int(from), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QStringRef_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QStringRef) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QStringRef_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QStringRef) LastIndexOf2(ch QCharITF, from int, cs Qt__CaseSensitivity) int {
	if ptr.Pointer() != nil {
		return int(C.QStringRef_LastIndexOf2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQChar(ch)), C.int(from), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) LastIndexOf3(str QLatin1StringITF, from int, cs Qt__CaseSensitivity) int {
	if ptr.Pointer() != nil {
		return int(C.QStringRef_LastIndexOf3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLatin1String(str)), C.int(from), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) LastIndexOf(str string, from int, cs Qt__CaseSensitivity) int {
	if ptr.Pointer() != nil {
		return int(C.QStringRef_LastIndexOf(C.QtObjectPtr(ptr.Pointer()), C.CString(str), C.int(from), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) LastIndexOf4(str QStringRefITF, from int, cs Qt__CaseSensitivity) int {
	if ptr.Pointer() != nil {
		return int(C.QStringRef_LastIndexOf4(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQStringRef(str)), C.int(from), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) Length() int {
	if ptr.Pointer() != nil {
		return int(C.QStringRef_Length(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func QStringRef_LocaleAwareCompare(s1 QStringRefITF, s2 string) int {
	return int(C.QStringRef_QStringRef_LocaleAwareCompare(C.QtObjectPtr(PointerFromQStringRef(s1)), C.CString(s2)))
}

func QStringRef_LocaleAwareCompare2(s1 QStringRefITF, s2 QStringRefITF) int {
	return int(C.QStringRef_QStringRef_LocaleAwareCompare2(C.QtObjectPtr(PointerFromQStringRef(s1)), C.QtObjectPtr(PointerFromQStringRef(s2))))
}

func (ptr *QStringRef) LocaleAwareCompare3(other string) int {
	if ptr.Pointer() != nil {
		return int(C.QStringRef_LocaleAwareCompare3(C.QtObjectPtr(ptr.Pointer()), C.CString(other)))
	}
	return 0
}

func (ptr *QStringRef) LocaleAwareCompare4(other QStringRefITF) int {
	if ptr.Pointer() != nil {
		return int(C.QStringRef_LocaleAwareCompare4(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQStringRef(other))))
	}
	return 0
}

func (ptr *QStringRef) Position() int {
	if ptr.Pointer() != nil {
		return int(C.QStringRef_Position(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStringRef) Size() int {
	if ptr.Pointer() != nil {
		return int(C.QStringRef_Size(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStringRef) StartsWith4(ch QCharITF, cs Qt__CaseSensitivity) bool {
	if ptr.Pointer() != nil {
		return C.QStringRef_StartsWith4(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQChar(ch)), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) StartsWith2(str QLatin1StringITF, cs Qt__CaseSensitivity) bool {
	if ptr.Pointer() != nil {
		return C.QStringRef_StartsWith2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLatin1String(str)), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) StartsWith(str string, cs Qt__CaseSensitivity) bool {
	if ptr.Pointer() != nil {
		return C.QStringRef_StartsWith(C.QtObjectPtr(ptr.Pointer()), C.CString(str), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) StartsWith3(str QStringRefITF, cs Qt__CaseSensitivity) bool {
	if ptr.Pointer() != nil {
		return C.QStringRef_StartsWith3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQStringRef(str)), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) String() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStringRef_String(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QStringRef) ToInt(ok bool, base int) int {
	if ptr.Pointer() != nil {
		return int(C.QStringRef_ToInt(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(ok)), C.int(base)))
	}
	return 0
}

func (ptr *QStringRef) ToString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStringRef_ToString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QStringRef) Unicode() *QChar {
	if ptr.Pointer() != nil {
		return QCharFromPointer(unsafe.Pointer(C.QStringRef_Unicode(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QStringRef) DestroyQStringRef() {
	if ptr.Pointer() != nil {
		C.QStringRef_DestroyQStringRef(C.QtObjectPtr(ptr.Pointer()))
	}
}
