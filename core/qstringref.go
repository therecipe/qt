package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::left")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QStringRef_Left(ptr.Pointer(), C.int(n)))
	}
	return nil
}

func (ptr *QStringRef) Mid(position int, n int) *QStringRef {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::mid")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QStringRef_Mid(ptr.Pointer(), C.int(position), C.int(n)))
	}
	return nil
}

func (ptr *QStringRef) Right(n int) *QStringRef {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::right")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QStringRef_Right(ptr.Pointer(), C.int(n)))
	}
	return nil
}

func (ptr *QStringRef) AppendTo(stri string) *QStringRef {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::appendTo")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QStringRef_AppendTo(ptr.Pointer(), C.CString(stri)))
	}
	return nil
}

func (ptr *QStringRef) Begin() *QChar {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::begin")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQCharFromPointer(C.QStringRef_Begin(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStringRef) Cbegin() *QChar {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::cbegin")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQCharFromPointer(C.QStringRef_Cbegin(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStringRef) Cend() *QChar {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::cend")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQCharFromPointer(C.QStringRef_Cend(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStringRef) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStringRef_Clear(ptr.Pointer())
	}
}

func QStringRef_Compare3(s1 QStringRef_ITF, s2 QLatin1String_ITF, cs Qt__CaseSensitivity) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::compare")
		}
	}()

	return int(C.QStringRef_QStringRef_Compare3(PointerFromQStringRef(s1), PointerFromQLatin1String(s2), C.int(cs)))
}

func QStringRef_Compare(s1 QStringRef_ITF, s2 string, cs Qt__CaseSensitivity) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::compare")
		}
	}()

	return int(C.QStringRef_QStringRef_Compare(PointerFromQStringRef(s1), C.CString(s2), C.int(cs)))
}

func QStringRef_Compare2(s1 QStringRef_ITF, s2 QStringRef_ITF, cs Qt__CaseSensitivity) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::compare")
		}
	}()

	return int(C.QStringRef_QStringRef_Compare2(PointerFromQStringRef(s1), PointerFromQStringRef(s2), C.int(cs)))
}

func (ptr *QStringRef) Compare6(other QLatin1String_ITF, cs Qt__CaseSensitivity) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::compare")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStringRef_Compare6(ptr.Pointer(), PointerFromQLatin1String(other), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) Compare4(other string, cs Qt__CaseSensitivity) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::compare")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStringRef_Compare4(ptr.Pointer(), C.CString(other), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) Compare5(other QStringRef_ITF, cs Qt__CaseSensitivity) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::compare")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStringRef_Compare5(ptr.Pointer(), PointerFromQStringRef(other), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) ConstData() *QChar {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::constData")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQCharFromPointer(C.QStringRef_ConstData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStringRef) Contains2(ch QChar_ITF, cs Qt__CaseSensitivity) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::contains")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStringRef_Contains2(ptr.Pointer(), PointerFromQChar(ch), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) Contains4(str QLatin1String_ITF, cs Qt__CaseSensitivity) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::contains")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStringRef_Contains4(ptr.Pointer(), PointerFromQLatin1String(str), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) Contains(str string, cs Qt__CaseSensitivity) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::contains")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStringRef_Contains(ptr.Pointer(), C.CString(str), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) Contains3(str QStringRef_ITF, cs Qt__CaseSensitivity) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::contains")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStringRef_Contains3(ptr.Pointer(), PointerFromQStringRef(str), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) Count() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::count")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStringRef_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStringRef) Count3(ch QChar_ITF, cs Qt__CaseSensitivity) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::count")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStringRef_Count3(ptr.Pointer(), PointerFromQChar(ch), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) Count2(str string, cs Qt__CaseSensitivity) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::count")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStringRef_Count2(ptr.Pointer(), C.CString(str), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) Count4(str QStringRef_ITF, cs Qt__CaseSensitivity) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::count")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStringRef_Count4(ptr.Pointer(), PointerFromQStringRef(str), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) Data() *QChar {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::data")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQCharFromPointer(C.QStringRef_Data(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStringRef) End() *QChar {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::end")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQCharFromPointer(C.QStringRef_End(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStringRef) EndsWith2(ch QChar_ITF, cs Qt__CaseSensitivity) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::endsWith")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStringRef_EndsWith2(ptr.Pointer(), PointerFromQChar(ch), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) EndsWith3(str QLatin1String_ITF, cs Qt__CaseSensitivity) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::endsWith")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStringRef_EndsWith3(ptr.Pointer(), PointerFromQLatin1String(str), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) EndsWith(str string, cs Qt__CaseSensitivity) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::endsWith")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStringRef_EndsWith(ptr.Pointer(), C.CString(str), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) EndsWith4(str QStringRef_ITF, cs Qt__CaseSensitivity) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::endsWith")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStringRef_EndsWith4(ptr.Pointer(), PointerFromQStringRef(str), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) IndexOf3(ch QChar_ITF, from int, cs Qt__CaseSensitivity) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::indexOf")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStringRef_IndexOf3(ptr.Pointer(), PointerFromQChar(ch), C.int(from), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) IndexOf2(str QLatin1String_ITF, from int, cs Qt__CaseSensitivity) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::indexOf")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStringRef_IndexOf2(ptr.Pointer(), PointerFromQLatin1String(str), C.int(from), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) IndexOf(str string, from int, cs Qt__CaseSensitivity) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::indexOf")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStringRef_IndexOf(ptr.Pointer(), C.CString(str), C.int(from), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) IndexOf4(str QStringRef_ITF, from int, cs Qt__CaseSensitivity) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::indexOf")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStringRef_IndexOf4(ptr.Pointer(), PointerFromQStringRef(str), C.int(from), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) IsEmpty() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::isEmpty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStringRef_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStringRef) IsNull() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::isNull")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStringRef_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStringRef) LastIndexOf2(ch QChar_ITF, from int, cs Qt__CaseSensitivity) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::lastIndexOf")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStringRef_LastIndexOf2(ptr.Pointer(), PointerFromQChar(ch), C.int(from), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) LastIndexOf3(str QLatin1String_ITF, from int, cs Qt__CaseSensitivity) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::lastIndexOf")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStringRef_LastIndexOf3(ptr.Pointer(), PointerFromQLatin1String(str), C.int(from), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) LastIndexOf(str string, from int, cs Qt__CaseSensitivity) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::lastIndexOf")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStringRef_LastIndexOf(ptr.Pointer(), C.CString(str), C.int(from), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) LastIndexOf4(str QStringRef_ITF, from int, cs Qt__CaseSensitivity) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::lastIndexOf")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStringRef_LastIndexOf4(ptr.Pointer(), PointerFromQStringRef(str), C.int(from), C.int(cs)))
	}
	return 0
}

func (ptr *QStringRef) Length() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::length")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStringRef_Length(ptr.Pointer()))
	}
	return 0
}

func QStringRef_LocaleAwareCompare(s1 QStringRef_ITF, s2 string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::localeAwareCompare")
		}
	}()

	return int(C.QStringRef_QStringRef_LocaleAwareCompare(PointerFromQStringRef(s1), C.CString(s2)))
}

func QStringRef_LocaleAwareCompare2(s1 QStringRef_ITF, s2 QStringRef_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::localeAwareCompare")
		}
	}()

	return int(C.QStringRef_QStringRef_LocaleAwareCompare2(PointerFromQStringRef(s1), PointerFromQStringRef(s2)))
}

func (ptr *QStringRef) LocaleAwareCompare3(other string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::localeAwareCompare")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStringRef_LocaleAwareCompare3(ptr.Pointer(), C.CString(other)))
	}
	return 0
}

func (ptr *QStringRef) LocaleAwareCompare4(other QStringRef_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::localeAwareCompare")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStringRef_LocaleAwareCompare4(ptr.Pointer(), PointerFromQStringRef(other)))
	}
	return 0
}

func (ptr *QStringRef) Position() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::position")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStringRef_Position(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStringRef) Size() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::size")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStringRef_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStringRef) StartsWith4(ch QChar_ITF, cs Qt__CaseSensitivity) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::startsWith")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStringRef_StartsWith4(ptr.Pointer(), PointerFromQChar(ch), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) StartsWith2(str QLatin1String_ITF, cs Qt__CaseSensitivity) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::startsWith")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStringRef_StartsWith2(ptr.Pointer(), PointerFromQLatin1String(str), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) StartsWith(str string, cs Qt__CaseSensitivity) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::startsWith")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStringRef_StartsWith(ptr.Pointer(), C.CString(str), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) StartsWith3(str QStringRef_ITF, cs Qt__CaseSensitivity) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::startsWith")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStringRef_StartsWith3(ptr.Pointer(), PointerFromQStringRef(str), C.int(cs)) != 0
	}
	return false
}

func (ptr *QStringRef) String() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::string")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QStringRef_String(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStringRef) ToInt(ok bool, base int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::toInt")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStringRef_ToInt(ptr.Pointer(), C.int(qt.GoBoolToInt(ok)), C.int(base)))
	}
	return 0
}

func (ptr *QStringRef) ToLatin1() *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::toLatin1")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QStringRef_ToLatin1(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStringRef) ToLocal8Bit() *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::toLocal8Bit")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QStringRef_ToLocal8Bit(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStringRef) ToString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::toString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QStringRef_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStringRef) ToUtf8() *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::toUtf8")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QStringRef_ToUtf8(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStringRef) Trimmed() *QStringRef {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::trimmed")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QStringRef_Trimmed(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStringRef) Unicode() *QChar {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::unicode")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQCharFromPointer(C.QStringRef_Unicode(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStringRef) DestroyQStringRef() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStringRef::~QStringRef")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStringRef_DestroyQStringRef(ptr.Pointer())
	}
}
