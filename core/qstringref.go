package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QStringRef struct {
	ptr unsafe.Pointer
}

type QStringRef_ITF interface {
	QStringRef_PTR() *QStringRef
}

func (p *QStringRef) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QStringRef) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQStringRef(ptr QStringRef_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStringRef_PTR().Pointer()
	}
	return nil
}

func NewQStringRefFromPointer(ptr unsafe.Pointer) *QStringRef {
	var n = new(QStringRef)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStringRef) QStringRef_PTR() *QStringRef {
	return ptr
}

func (ptr *QStringRef) Left(n int) *QStringRef {
	defer qt.Recovering("QStringRef::left")

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QStringRef_Left(ptr.Pointer(), C.int(n)))
	}
	return nil
}

func (ptr *QStringRef) Mid(position int, n int) *QStringRef {
	defer qt.Recovering("QStringRef::mid")

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QStringRef_Mid(ptr.Pointer(), C.int(position), C.int(n)))
	}
	return nil
}

func (ptr *QStringRef) Right(n int) *QStringRef {
	defer qt.Recovering("QStringRef::right")

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QStringRef_Right(ptr.Pointer(), C.int(n)))
	}
	return nil
}

func (ptr *QStringRef) AppendTo(stri string) *QStringRef {
	defer qt.Recovering("QStringRef::appendTo")

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QStringRef_AppendTo(ptr.Pointer(), C.CString(stri)))
	}
	return nil
}

func (ptr *QStringRef) Begin() *QChar {
	defer qt.Recovering("QStringRef::begin")

	if ptr.Pointer() != nil {
		return NewQCharFromPointer(C.QStringRef_Begin(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStringRef) Cbegin() *QChar {
	defer qt.Recovering("QStringRef::cbegin")

	if ptr.Pointer() != nil {
		return NewQCharFromPointer(C.QStringRef_Cbegin(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStringRef) Cend() *QChar {
	defer qt.Recovering("QStringRef::cend")

	if ptr.Pointer() != nil {
		return NewQCharFromPointer(C.QStringRef_Cend(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStringRef) Clear() {
	defer qt.Recovering("QStringRef::clear")

	if ptr.Pointer() != nil {
		C.QStringRef_Clear(ptr.Pointer())
	}
}

func QStringRef_Compare3(s1 QStringRef_ITF, s2 QLatin1String_ITF, cs Qt__CaseSensitivity) int {
	defer qt.Recovering("QStringRef::compare")

	return int(C.QStringRef_QStringRef_Compare3(PointerFromQStringRef(s1), PointerFromQLatin1String(s2), C.int(cs)))
}

func QStringRef_Compare(s1 QStringRef_ITF, s2 string, cs Qt__CaseSensitivity) int {
	defer qt.Recovering("QStringRef::compare")

	return int(C.QStringRef_QStringRef_Compare(PointerFromQStringRef(s1), C.CString(s2), C.int(cs)))
}

func QStringRef_Compare2(s1 QStringRef_ITF, s2 QStringRef_ITF, cs Qt__CaseSensitivity) int {
	defer qt.Recovering("QStringRef::compare")

	return int(C.QStringRef_QStringRef_Compare2(PointerFromQStringRef(s1), PointerFromQStringRef(s2), C.int(cs)))
}

func (ptr *QStringRef) Compare6(other QLatin1String_ITF, cs Qt__CaseSensitivity) int {
	defer qt.Recovering("QStringRef::compare")

	if ptr.Pointer() != nil {
		return int(C.QStringRef_Compare6(ptr.Pointer(), PointerFromQLatin1String(other), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) Compare4(other string, cs Qt__CaseSensitivity) int {
	defer qt.Recovering("QStringRef::compare")

	if ptr.Pointer() != nil {
		return int(C.QStringRef_Compare4(ptr.Pointer(), C.CString(other), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) Compare5(other QStringRef_ITF, cs Qt__CaseSensitivity) int {
	defer qt.Recovering("QStringRef::compare")

	if ptr.Pointer() != nil {
		return int(C.QStringRef_Compare5(ptr.Pointer(), PointerFromQStringRef(other), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) ConstData() *QChar {
	defer qt.Recovering("QStringRef::constData")

	if ptr.Pointer() != nil {
		return NewQCharFromPointer(C.QStringRef_ConstData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStringRef) Contains2(ch QChar_ITF, cs Qt__CaseSensitivity) bool {
	defer qt.Recovering("QStringRef::contains")

	if ptr.Pointer() != nil {
		return C.QStringRef_Contains2(ptr.Pointer(), PointerFromQChar(ch), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) Contains4(str QLatin1String_ITF, cs Qt__CaseSensitivity) bool {
	defer qt.Recovering("QStringRef::contains")

	if ptr.Pointer() != nil {
		return C.QStringRef_Contains4(ptr.Pointer(), PointerFromQLatin1String(str), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) Contains(str string, cs Qt__CaseSensitivity) bool {
	defer qt.Recovering("QStringRef::contains")

	if ptr.Pointer() != nil {
		return C.QStringRef_Contains(ptr.Pointer(), C.CString(str), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) Contains3(str QStringRef_ITF, cs Qt__CaseSensitivity) bool {
	defer qt.Recovering("QStringRef::contains")

	if ptr.Pointer() != nil {
		return C.QStringRef_Contains3(ptr.Pointer(), PointerFromQStringRef(str), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) Count() int {
	defer qt.Recovering("QStringRef::count")

	if ptr.Pointer() != nil {
		return int(C.QStringRef_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStringRef) Count3(ch QChar_ITF, cs Qt__CaseSensitivity) int {
	defer qt.Recovering("QStringRef::count")

	if ptr.Pointer() != nil {
		return int(C.QStringRef_Count3(ptr.Pointer(), PointerFromQChar(ch), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) Count2(str string, cs Qt__CaseSensitivity) int {
	defer qt.Recovering("QStringRef::count")

	if ptr.Pointer() != nil {
		return int(C.QStringRef_Count2(ptr.Pointer(), C.CString(str), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) Count4(str QStringRef_ITF, cs Qt__CaseSensitivity) int {
	defer qt.Recovering("QStringRef::count")

	if ptr.Pointer() != nil {
		return int(C.QStringRef_Count4(ptr.Pointer(), PointerFromQStringRef(str), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) Data() *QChar {
	defer qt.Recovering("QStringRef::data")

	if ptr.Pointer() != nil {
		return NewQCharFromPointer(C.QStringRef_Data(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStringRef) End() *QChar {
	defer qt.Recovering("QStringRef::end")

	if ptr.Pointer() != nil {
		return NewQCharFromPointer(C.QStringRef_End(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStringRef) EndsWith2(ch QChar_ITF, cs Qt__CaseSensitivity) bool {
	defer qt.Recovering("QStringRef::endsWith")

	if ptr.Pointer() != nil {
		return C.QStringRef_EndsWith2(ptr.Pointer(), PointerFromQChar(ch), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) EndsWith3(str QLatin1String_ITF, cs Qt__CaseSensitivity) bool {
	defer qt.Recovering("QStringRef::endsWith")

	if ptr.Pointer() != nil {
		return C.QStringRef_EndsWith3(ptr.Pointer(), PointerFromQLatin1String(str), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) EndsWith(str string, cs Qt__CaseSensitivity) bool {
	defer qt.Recovering("QStringRef::endsWith")

	if ptr.Pointer() != nil {
		return C.QStringRef_EndsWith(ptr.Pointer(), C.CString(str), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) EndsWith4(str QStringRef_ITF, cs Qt__CaseSensitivity) bool {
	defer qt.Recovering("QStringRef::endsWith")

	if ptr.Pointer() != nil {
		return C.QStringRef_EndsWith4(ptr.Pointer(), PointerFromQStringRef(str), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) IndexOf3(ch QChar_ITF, from int, cs Qt__CaseSensitivity) int {
	defer qt.Recovering("QStringRef::indexOf")

	if ptr.Pointer() != nil {
		return int(C.QStringRef_IndexOf3(ptr.Pointer(), PointerFromQChar(ch), C.int(from), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) IndexOf2(str QLatin1String_ITF, from int, cs Qt__CaseSensitivity) int {
	defer qt.Recovering("QStringRef::indexOf")

	if ptr.Pointer() != nil {
		return int(C.QStringRef_IndexOf2(ptr.Pointer(), PointerFromQLatin1String(str), C.int(from), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) IndexOf(str string, from int, cs Qt__CaseSensitivity) int {
	defer qt.Recovering("QStringRef::indexOf")

	if ptr.Pointer() != nil {
		return int(C.QStringRef_IndexOf(ptr.Pointer(), C.CString(str), C.int(from), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) IndexOf4(str QStringRef_ITF, from int, cs Qt__CaseSensitivity) int {
	defer qt.Recovering("QStringRef::indexOf")

	if ptr.Pointer() != nil {
		return int(C.QStringRef_IndexOf4(ptr.Pointer(), PointerFromQStringRef(str), C.int(from), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) IsEmpty() bool {
	defer qt.Recovering("QStringRef::isEmpty")

	if ptr.Pointer() != nil {
		return C.QStringRef_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStringRef) IsNull() bool {
	defer qt.Recovering("QStringRef::isNull")

	if ptr.Pointer() != nil {
		return C.QStringRef_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStringRef) LastIndexOf2(ch QChar_ITF, from int, cs Qt__CaseSensitivity) int {
	defer qt.Recovering("QStringRef::lastIndexOf")

	if ptr.Pointer() != nil {
		return int(C.QStringRef_LastIndexOf2(ptr.Pointer(), PointerFromQChar(ch), C.int(from), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) LastIndexOf3(str QLatin1String_ITF, from int, cs Qt__CaseSensitivity) int {
	defer qt.Recovering("QStringRef::lastIndexOf")

	if ptr.Pointer() != nil {
		return int(C.QStringRef_LastIndexOf3(ptr.Pointer(), PointerFromQLatin1String(str), C.int(from), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) LastIndexOf(str string, from int, cs Qt__CaseSensitivity) int {
	defer qt.Recovering("QStringRef::lastIndexOf")

	if ptr.Pointer() != nil {
		return int(C.QStringRef_LastIndexOf(ptr.Pointer(), C.CString(str), C.int(from), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) LastIndexOf4(str QStringRef_ITF, from int, cs Qt__CaseSensitivity) int {
	defer qt.Recovering("QStringRef::lastIndexOf")

	if ptr.Pointer() != nil {
		return int(C.QStringRef_LastIndexOf4(ptr.Pointer(), PointerFromQStringRef(str), C.int(from), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) Length() int {
	defer qt.Recovering("QStringRef::length")

	if ptr.Pointer() != nil {
		return int(C.QStringRef_Length(ptr.Pointer()))
	}
	return 0
}

func QStringRef_LocaleAwareCompare(s1 QStringRef_ITF, s2 string) int {
	defer qt.Recovering("QStringRef::localeAwareCompare")

	return int(C.QStringRef_QStringRef_LocaleAwareCompare(PointerFromQStringRef(s1), C.CString(s2)))
}

func QStringRef_LocaleAwareCompare2(s1 QStringRef_ITF, s2 QStringRef_ITF) int {
	defer qt.Recovering("QStringRef::localeAwareCompare")

	return int(C.QStringRef_QStringRef_LocaleAwareCompare2(PointerFromQStringRef(s1), PointerFromQStringRef(s2)))
}

func (ptr *QStringRef) LocaleAwareCompare3(other string) int {
	defer qt.Recovering("QStringRef::localeAwareCompare")

	if ptr.Pointer() != nil {
		return int(C.QStringRef_LocaleAwareCompare3(ptr.Pointer(), C.CString(other)))
	}
	return 0
}

func (ptr *QStringRef) LocaleAwareCompare4(other QStringRef_ITF) int {
	defer qt.Recovering("QStringRef::localeAwareCompare")

	if ptr.Pointer() != nil {
		return int(C.QStringRef_LocaleAwareCompare4(ptr.Pointer(), PointerFromQStringRef(other)))
	}
	return 0
}

func (ptr *QStringRef) Position() int {
	defer qt.Recovering("QStringRef::position")

	if ptr.Pointer() != nil {
		return int(C.QStringRef_Position(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStringRef) Size() int {
	defer qt.Recovering("QStringRef::size")

	if ptr.Pointer() != nil {
		return int(C.QStringRef_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStringRef) StartsWith4(ch QChar_ITF, cs Qt__CaseSensitivity) bool {
	defer qt.Recovering("QStringRef::startsWith")

	if ptr.Pointer() != nil {
		return C.QStringRef_StartsWith4(ptr.Pointer(), PointerFromQChar(ch), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) StartsWith2(str QLatin1String_ITF, cs Qt__CaseSensitivity) bool {
	defer qt.Recovering("QStringRef::startsWith")

	if ptr.Pointer() != nil {
		return C.QStringRef_StartsWith2(ptr.Pointer(), PointerFromQLatin1String(str), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) StartsWith(str string, cs Qt__CaseSensitivity) bool {
	defer qt.Recovering("QStringRef::startsWith")

	if ptr.Pointer() != nil {
		return C.QStringRef_StartsWith(ptr.Pointer(), C.CString(str), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) StartsWith3(str QStringRef_ITF, cs Qt__CaseSensitivity) bool {
	defer qt.Recovering("QStringRef::startsWith")

	if ptr.Pointer() != nil {
		return C.QStringRef_StartsWith3(ptr.Pointer(), PointerFromQStringRef(str), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) String() string {
	defer qt.Recovering("QStringRef::string")

	if ptr.Pointer() != nil {
		return C.GoString(C.QStringRef_String(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStringRef) ToInt(ok bool, base int) int {
	defer qt.Recovering("QStringRef::toInt")

	if ptr.Pointer() != nil {
		return int(C.QStringRef_ToInt(ptr.Pointer(), C.int(qt.GoBoolToInt(ok)), C.int(base)))
	}
	return 0
}

func (ptr *QStringRef) ToLatin1() *QByteArray {
	defer qt.Recovering("QStringRef::toLatin1")

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QStringRef_ToLatin1(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStringRef) ToLocal8Bit() *QByteArray {
	defer qt.Recovering("QStringRef::toLocal8Bit")

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QStringRef_ToLocal8Bit(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStringRef) ToString() string {
	defer qt.Recovering("QStringRef::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QStringRef_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStringRef) ToUtf8() *QByteArray {
	defer qt.Recovering("QStringRef::toUtf8")

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QStringRef_ToUtf8(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStringRef) Trimmed() *QStringRef {
	defer qt.Recovering("QStringRef::trimmed")

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QStringRef_Trimmed(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStringRef) Unicode() *QChar {
	defer qt.Recovering("QStringRef::unicode")

	if ptr.Pointer() != nil {
		return NewQCharFromPointer(C.QStringRef_Unicode(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStringRef) DestroyQStringRef() {
	defer qt.Recovering("QStringRef::~QStringRef")

	if ptr.Pointer() != nil {
		C.QStringRef_DestroyQStringRef(ptr.Pointer())
	}
}
