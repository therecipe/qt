package core

//#include "qbytearray.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QByteArray struct {
	ptr unsafe.Pointer
}

type QByteArray_ITF interface {
	QByteArray_PTR() *QByteArray
}

func (p *QByteArray) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QByteArray) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQByteArray(ptr QByteArray_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QByteArray_PTR().Pointer()
	}
	return nil
}

func NewQByteArrayFromPointer(ptr unsafe.Pointer) *QByteArray {
	var n = new(QByteArray)
	n.SetPointer(ptr)
	return n
}

func (ptr *QByteArray) QByteArray_PTR() *QByteArray {
	return ptr
}

//QByteArray::Base64Option
type QByteArray__Base64Option int64

const (
	QByteArray__Base64Encoding     = QByteArray__Base64Option(0)
	QByteArray__Base64UrlEncoding  = QByteArray__Base64Option(1)
	QByteArray__KeepTrailingEquals = QByteArray__Base64Option(0)
	QByteArray__OmitTrailingEquals = QByteArray__Base64Option(2)
)

func (ptr *QByteArray) Clear() {
	if ptr.Pointer() != nil {
		C.QByteArray_Clear(ptr.Pointer())
	}
}

func (ptr *QByteArray) IndexOf3(str string, from int) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_IndexOf3(ptr.Pointer(), C.CString(str), C.int(from)))
	}
	return 0
}

func (ptr *QByteArray) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QByteArray_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QByteArray) LastIndexOf(ba QByteArray_ITF, from int) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_LastIndexOf(ptr.Pointer(), PointerFromQByteArray(ba), C.int(from)))
	}
	return 0
}

func (ptr *QByteArray) LastIndexOf3(str string, from int) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_LastIndexOf3(ptr.Pointer(), C.CString(str), C.int(from)))
	}
	return 0
}

func NewQByteArray() *QByteArray {
	return NewQByteArrayFromPointer(C.QByteArray_NewQByteArray())
}

func NewQByteArray6(other QByteArray_ITF) *QByteArray {
	return NewQByteArrayFromPointer(C.QByteArray_NewQByteArray6(PointerFromQByteArray(other)))
}

func NewQByteArray5(other QByteArray_ITF) *QByteArray {
	return NewQByteArrayFromPointer(C.QByteArray_NewQByteArray5(PointerFromQByteArray(other)))
}

func NewQByteArray2(data string, size int) *QByteArray {
	return NewQByteArrayFromPointer(C.QByteArray_NewQByteArray2(C.CString(data), C.int(size)))
}

func NewQByteArray3(size int, ch string) *QByteArray {
	return NewQByteArrayFromPointer(C.QByteArray_NewQByteArray3(C.int(size), C.CString(ch)))
}

func (ptr *QByteArray) Append5(ch string) *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArray_Append5(ptr.Pointer(), C.CString(ch)))
	}
	return nil
}

func (ptr *QByteArray) Append(ba QByteArray_ITF) *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArray_Append(ptr.Pointer(), PointerFromQByteArray(ba)))
	}
	return nil
}

func (ptr *QByteArray) Append2(str string) *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArray_Append2(ptr.Pointer(), C.CString(str)))
	}
	return nil
}

func (ptr *QByteArray) Append3(str string) *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArray_Append3(ptr.Pointer(), C.CString(str)))
	}
	return nil
}

func (ptr *QByteArray) Append4(str string, len int) *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArray_Append4(ptr.Pointer(), C.CString(str), C.int(len)))
	}
	return nil
}

func (ptr *QByteArray) Capacity() int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_Capacity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QByteArray) Chop(n int) {
	if ptr.Pointer() != nil {
		C.QByteArray_Chop(ptr.Pointer(), C.int(n))
	}
}

func (ptr *QByteArray) Contains3(ch string) bool {
	if ptr.Pointer() != nil {
		return C.QByteArray_Contains3(ptr.Pointer(), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QByteArray) Contains(ba QByteArray_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QByteArray_Contains(ptr.Pointer(), PointerFromQByteArray(ba)) != 0
	}
	return false
}

func (ptr *QByteArray) Contains2(str string) bool {
	if ptr.Pointer() != nil {
		return C.QByteArray_Contains2(ptr.Pointer(), C.CString(str)) != 0
	}
	return false
}

func (ptr *QByteArray) Count4() int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_Count4(ptr.Pointer()))
	}
	return 0
}

func (ptr *QByteArray) Count3(ch string) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_Count3(ptr.Pointer(), C.CString(ch)))
	}
	return 0
}

func (ptr *QByteArray) Count(ba QByteArray_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_Count(ptr.Pointer(), PointerFromQByteArray(ba)))
	}
	return 0
}

func (ptr *QByteArray) Count2(str string) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_Count2(ptr.Pointer(), C.CString(str)))
	}
	return 0
}

func (ptr *QByteArray) EndsWith3(ch string) bool {
	if ptr.Pointer() != nil {
		return C.QByteArray_EndsWith3(ptr.Pointer(), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QByteArray) EndsWith(ba QByteArray_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QByteArray_EndsWith(ptr.Pointer(), PointerFromQByteArray(ba)) != 0
	}
	return false
}

func (ptr *QByteArray) EndsWith2(str string) bool {
	if ptr.Pointer() != nil {
		return C.QByteArray_EndsWith2(ptr.Pointer(), C.CString(str)) != 0
	}
	return false
}

func (ptr *QByteArray) Fill(ch string, size int) *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArray_Fill(ptr.Pointer(), C.CString(ch), C.int(size)))
	}
	return nil
}

func QByteArray_FromBase64(base64 QByteArray_ITF) *QByteArray {
	return NewQByteArrayFromPointer(C.QByteArray_QByteArray_FromBase64(PointerFromQByteArray(base64)))
}

func QByteArray_FromBase642(base64 QByteArray_ITF, options QByteArray__Base64Option) *QByteArray {
	return NewQByteArrayFromPointer(C.QByteArray_QByteArray_FromBase642(PointerFromQByteArray(base64), C.int(options)))
}

func QByteArray_FromHex(hexEncoded QByteArray_ITF) *QByteArray {
	return NewQByteArrayFromPointer(C.QByteArray_QByteArray_FromHex(PointerFromQByteArray(hexEncoded)))
}

func QByteArray_FromPercentEncoding(input QByteArray_ITF, percent string) *QByteArray {
	return NewQByteArrayFromPointer(C.QByteArray_QByteArray_FromPercentEncoding(PointerFromQByteArray(input), C.CString(percent)))
}

func QByteArray_FromRawData(data string, size int) *QByteArray {
	return NewQByteArrayFromPointer(C.QByteArray_QByteArray_FromRawData(C.CString(data), C.int(size)))
}

func (ptr *QByteArray) IndexOf4(ch string, from int) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_IndexOf4(ptr.Pointer(), C.CString(ch), C.int(from)))
	}
	return 0
}

func (ptr *QByteArray) IndexOf(ba QByteArray_ITF, from int) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_IndexOf(ptr.Pointer(), PointerFromQByteArray(ba), C.int(from)))
	}
	return 0
}

func (ptr *QByteArray) IndexOf2(str string, from int) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_IndexOf2(ptr.Pointer(), C.CString(str), C.int(from)))
	}
	return 0
}

func (ptr *QByteArray) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QByteArray_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QByteArray) LastIndexOf4(ch string, from int) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_LastIndexOf4(ptr.Pointer(), C.CString(ch), C.int(from)))
	}
	return 0
}

func (ptr *QByteArray) LastIndexOf2(str string, from int) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_LastIndexOf2(ptr.Pointer(), C.CString(str), C.int(from)))
	}
	return 0
}

func (ptr *QByteArray) Left(len int) *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArray_Left(ptr.Pointer(), C.int(len)))
	}
	return nil
}

func (ptr *QByteArray) LeftJustified(width int, fill string, truncate bool) *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArray_LeftJustified(ptr.Pointer(), C.int(width), C.CString(fill), C.int(qt.GoBoolToInt(truncate))))
	}
	return nil
}

func (ptr *QByteArray) Length() int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QByteArray) Mid(pos int, len int) *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArray_Mid(ptr.Pointer(), C.int(pos), C.int(len)))
	}
	return nil
}

func QByteArray_Number(n int, base int) *QByteArray {
	return NewQByteArrayFromPointer(C.QByteArray_QByteArray_Number(C.int(n), C.int(base)))
}

func (ptr *QByteArray) Prepend4(ch string) *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArray_Prepend4(ptr.Pointer(), C.CString(ch)))
	}
	return nil
}

func (ptr *QByteArray) Prepend(ba QByteArray_ITF) *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArray_Prepend(ptr.Pointer(), PointerFromQByteArray(ba)))
	}
	return nil
}

func (ptr *QByteArray) Prepend2(str string) *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArray_Prepend2(ptr.Pointer(), C.CString(str)))
	}
	return nil
}

func (ptr *QByteArray) Prepend3(str string, len int) *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArray_Prepend3(ptr.Pointer(), C.CString(str), C.int(len)))
	}
	return nil
}

func (ptr *QByteArray) Push_back3(ch string) {
	if ptr.Pointer() != nil {
		C.QByteArray_Push_back3(ptr.Pointer(), C.CString(ch))
	}
}

func (ptr *QByteArray) Push_back(other QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QByteArray_Push_back(ptr.Pointer(), PointerFromQByteArray(other))
	}
}

func (ptr *QByteArray) Push_back2(str string) {
	if ptr.Pointer() != nil {
		C.QByteArray_Push_back2(ptr.Pointer(), C.CString(str))
	}
}

func (ptr *QByteArray) Push_front3(ch string) {
	if ptr.Pointer() != nil {
		C.QByteArray_Push_front3(ptr.Pointer(), C.CString(ch))
	}
}

func (ptr *QByteArray) Push_front(other QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QByteArray_Push_front(ptr.Pointer(), PointerFromQByteArray(other))
	}
}

func (ptr *QByteArray) Push_front2(str string) {
	if ptr.Pointer() != nil {
		C.QByteArray_Push_front2(ptr.Pointer(), C.CString(str))
	}
}

func (ptr *QByteArray) Repeated(times int) *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArray_Repeated(ptr.Pointer(), C.int(times)))
	}
	return nil
}

func (ptr *QByteArray) Reserve(size int) {
	if ptr.Pointer() != nil {
		C.QByteArray_Reserve(ptr.Pointer(), C.int(size))
	}
}

func (ptr *QByteArray) Resize(size int) {
	if ptr.Pointer() != nil {
		C.QByteArray_Resize(ptr.Pointer(), C.int(size))
	}
}

func (ptr *QByteArray) Right(len int) *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArray_Right(ptr.Pointer(), C.int(len)))
	}
	return nil
}

func (ptr *QByteArray) RightJustified(width int, fill string, truncate bool) *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArray_RightJustified(ptr.Pointer(), C.int(width), C.CString(fill), C.int(qt.GoBoolToInt(truncate))))
	}
	return nil
}

func (ptr *QByteArray) SetNum(n int, base int) *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArray_SetNum(ptr.Pointer(), C.int(n), C.int(base)))
	}
	return nil
}

func (ptr *QByteArray) Size() int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QByteArray) Squeeze() {
	if ptr.Pointer() != nil {
		C.QByteArray_Squeeze(ptr.Pointer())
	}
}

func (ptr *QByteArray) StartsWith3(ch string) bool {
	if ptr.Pointer() != nil {
		return C.QByteArray_StartsWith3(ptr.Pointer(), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QByteArray) StartsWith(ba QByteArray_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QByteArray_StartsWith(ptr.Pointer(), PointerFromQByteArray(ba)) != 0
	}
	return false
}

func (ptr *QByteArray) StartsWith2(str string) bool {
	if ptr.Pointer() != nil {
		return C.QByteArray_StartsWith2(ptr.Pointer(), C.CString(str)) != 0
	}
	return false
}

func (ptr *QByteArray) Swap(other QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QByteArray_Swap(ptr.Pointer(), PointerFromQByteArray(other))
	}
}

func (ptr *QByteArray) ToBase64() *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArray_ToBase64(ptr.Pointer()))
	}
	return nil
}

func (ptr *QByteArray) ToBase642(options QByteArray__Base64Option) *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArray_ToBase642(ptr.Pointer(), C.int(options)))
	}
	return nil
}

func (ptr *QByteArray) ToHex() *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArray_ToHex(ptr.Pointer()))
	}
	return nil
}

func (ptr *QByteArray) ToInt(ok bool, base int) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_ToInt(ptr.Pointer(), C.int(qt.GoBoolToInt(ok)), C.int(base)))
	}
	return 0
}

func (ptr *QByteArray) ToPercentEncoding(exclude QByteArray_ITF, include QByteArray_ITF, percent string) *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArray_ToPercentEncoding(ptr.Pointer(), PointerFromQByteArray(exclude), PointerFromQByteArray(include), C.CString(percent)))
	}
	return nil
}

func (ptr *QByteArray) Truncate(pos int) {
	if ptr.Pointer() != nil {
		C.QByteArray_Truncate(ptr.Pointer(), C.int(pos))
	}
}

func (ptr *QByteArray) DestroyQByteArray() {
	if ptr.Pointer() != nil {
		C.QByteArray_DestroyQByteArray(ptr.Pointer())
	}
}

func (ptr *QByteArray) Simplified() *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArray_Simplified(ptr.Pointer()))
	}
	return nil
}

func (ptr *QByteArray) ToLower() *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArray_ToLower(ptr.Pointer()))
	}
	return nil
}

func (ptr *QByteArray) ToUpper() *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArray_ToUpper(ptr.Pointer()))
	}
	return nil
}

func (ptr *QByteArray) Trimmed() *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QByteArray_Trimmed(ptr.Pointer()))
	}
	return nil
}
