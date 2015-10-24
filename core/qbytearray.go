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

type QByteArrayITF interface {
	QByteArrayPTR() *QByteArray
}

func (p *QByteArray) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QByteArray) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQByteArray(ptr QByteArrayITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QByteArrayPTR().Pointer()
	}
	return nil
}

func QByteArrayFromPointer(ptr unsafe.Pointer) *QByteArray {
	var n = new(QByteArray)
	n.SetPointer(ptr)
	return n
}

func (ptr *QByteArray) QByteArrayPTR() *QByteArray {
	return ptr
}

//QByteArray::Base64Option
type QByteArray__Base64Option int

var (
	QByteArray__Base64Encoding     = QByteArray__Base64Option(0)
	QByteArray__Base64UrlEncoding  = QByteArray__Base64Option(1)
	QByteArray__KeepTrailingEquals = QByteArray__Base64Option(0)
	QByteArray__OmitTrailingEquals = QByteArray__Base64Option(2)
)

func (ptr *QByteArray) Clear() {
	if ptr.Pointer() != nil {
		C.QByteArray_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QByteArray) IndexOf3(str string, from int) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_IndexOf3(C.QtObjectPtr(ptr.Pointer()), C.CString(str), C.int(from)))
	}
	return 0
}

func (ptr *QByteArray) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QByteArray_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QByteArray) LastIndexOf(ba QByteArrayITF, from int) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_LastIndexOf(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQByteArray(ba)), C.int(from)))
	}
	return 0
}

func (ptr *QByteArray) LastIndexOf3(str string, from int) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_LastIndexOf3(C.QtObjectPtr(ptr.Pointer()), C.CString(str), C.int(from)))
	}
	return 0
}

func NewQByteArray() *QByteArray {
	return QByteArrayFromPointer(unsafe.Pointer(C.QByteArray_NewQByteArray()))
}

func NewQByteArray6(other QByteArrayITF) *QByteArray {
	return QByteArrayFromPointer(unsafe.Pointer(C.QByteArray_NewQByteArray6(C.QtObjectPtr(PointerFromQByteArray(other)))))
}

func NewQByteArray5(other QByteArrayITF) *QByteArray {
	return QByteArrayFromPointer(unsafe.Pointer(C.QByteArray_NewQByteArray5(C.QtObjectPtr(PointerFromQByteArray(other)))))
}

func NewQByteArray2(data string, size int) *QByteArray {
	return QByteArrayFromPointer(unsafe.Pointer(C.QByteArray_NewQByteArray2(C.CString(data), C.int(size))))
}

func NewQByteArray3(size int, ch string) *QByteArray {
	return QByteArrayFromPointer(unsafe.Pointer(C.QByteArray_NewQByteArray3(C.int(size), C.CString(ch))))
}

func (ptr *QByteArray) Capacity() int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_Capacity(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QByteArray) Chop(n int) {
	if ptr.Pointer() != nil {
		C.QByteArray_Chop(C.QtObjectPtr(ptr.Pointer()), C.int(n))
	}
}

func (ptr *QByteArray) Contains3(ch string) bool {
	if ptr.Pointer() != nil {
		return C.QByteArray_Contains3(C.QtObjectPtr(ptr.Pointer()), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QByteArray) Contains(ba QByteArrayITF) bool {
	if ptr.Pointer() != nil {
		return C.QByteArray_Contains(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQByteArray(ba))) != 0
	}
	return false
}

func (ptr *QByteArray) Contains2(str string) bool {
	if ptr.Pointer() != nil {
		return C.QByteArray_Contains2(C.QtObjectPtr(ptr.Pointer()), C.CString(str)) != 0
	}
	return false
}

func (ptr *QByteArray) Count4() int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_Count4(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QByteArray) Count3(ch string) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_Count3(C.QtObjectPtr(ptr.Pointer()), C.CString(ch)))
	}
	return 0
}

func (ptr *QByteArray) Count(ba QByteArrayITF) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_Count(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQByteArray(ba))))
	}
	return 0
}

func (ptr *QByteArray) Count2(str string) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_Count2(C.QtObjectPtr(ptr.Pointer()), C.CString(str)))
	}
	return 0
}

func (ptr *QByteArray) EndsWith3(ch string) bool {
	if ptr.Pointer() != nil {
		return C.QByteArray_EndsWith3(C.QtObjectPtr(ptr.Pointer()), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QByteArray) EndsWith(ba QByteArrayITF) bool {
	if ptr.Pointer() != nil {
		return C.QByteArray_EndsWith(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQByteArray(ba))) != 0
	}
	return false
}

func (ptr *QByteArray) EndsWith2(str string) bool {
	if ptr.Pointer() != nil {
		return C.QByteArray_EndsWith2(C.QtObjectPtr(ptr.Pointer()), C.CString(str)) != 0
	}
	return false
}

func (ptr *QByteArray) IndexOf4(ch string, from int) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_IndexOf4(C.QtObjectPtr(ptr.Pointer()), C.CString(ch), C.int(from)))
	}
	return 0
}

func (ptr *QByteArray) IndexOf(ba QByteArrayITF, from int) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_IndexOf(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQByteArray(ba)), C.int(from)))
	}
	return 0
}

func (ptr *QByteArray) IndexOf2(str string, from int) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_IndexOf2(C.QtObjectPtr(ptr.Pointer()), C.CString(str), C.int(from)))
	}
	return 0
}

func (ptr *QByteArray) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QByteArray_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QByteArray) LastIndexOf4(ch string, from int) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_LastIndexOf4(C.QtObjectPtr(ptr.Pointer()), C.CString(ch), C.int(from)))
	}
	return 0
}

func (ptr *QByteArray) LastIndexOf2(str string, from int) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_LastIndexOf2(C.QtObjectPtr(ptr.Pointer()), C.CString(str), C.int(from)))
	}
	return 0
}

func (ptr *QByteArray) Length() int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_Length(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QByteArray) Push_back3(ch string) {
	if ptr.Pointer() != nil {
		C.QByteArray_Push_back3(C.QtObjectPtr(ptr.Pointer()), C.CString(ch))
	}
}

func (ptr *QByteArray) Push_back(other QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QByteArray_Push_back(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQByteArray(other)))
	}
}

func (ptr *QByteArray) Push_back2(str string) {
	if ptr.Pointer() != nil {
		C.QByteArray_Push_back2(C.QtObjectPtr(ptr.Pointer()), C.CString(str))
	}
}

func (ptr *QByteArray) Push_front3(ch string) {
	if ptr.Pointer() != nil {
		C.QByteArray_Push_front3(C.QtObjectPtr(ptr.Pointer()), C.CString(ch))
	}
}

func (ptr *QByteArray) Push_front(other QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QByteArray_Push_front(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQByteArray(other)))
	}
}

func (ptr *QByteArray) Push_front2(str string) {
	if ptr.Pointer() != nil {
		C.QByteArray_Push_front2(C.QtObjectPtr(ptr.Pointer()), C.CString(str))
	}
}

func (ptr *QByteArray) Reserve(size int) {
	if ptr.Pointer() != nil {
		C.QByteArray_Reserve(C.QtObjectPtr(ptr.Pointer()), C.int(size))
	}
}

func (ptr *QByteArray) Resize(size int) {
	if ptr.Pointer() != nil {
		C.QByteArray_Resize(C.QtObjectPtr(ptr.Pointer()), C.int(size))
	}
}

func (ptr *QByteArray) Size() int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_Size(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QByteArray) Squeeze() {
	if ptr.Pointer() != nil {
		C.QByteArray_Squeeze(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QByteArray) StartsWith3(ch string) bool {
	if ptr.Pointer() != nil {
		return C.QByteArray_StartsWith3(C.QtObjectPtr(ptr.Pointer()), C.CString(ch)) != 0
	}
	return false
}

func (ptr *QByteArray) StartsWith(ba QByteArrayITF) bool {
	if ptr.Pointer() != nil {
		return C.QByteArray_StartsWith(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQByteArray(ba))) != 0
	}
	return false
}

func (ptr *QByteArray) StartsWith2(str string) bool {
	if ptr.Pointer() != nil {
		return C.QByteArray_StartsWith2(C.QtObjectPtr(ptr.Pointer()), C.CString(str)) != 0
	}
	return false
}

func (ptr *QByteArray) Swap(other QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QByteArray_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQByteArray(other)))
	}
}

func (ptr *QByteArray) ToInt(ok bool, base int) int {
	if ptr.Pointer() != nil {
		return int(C.QByteArray_ToInt(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(ok)), C.int(base)))
	}
	return 0
}

func (ptr *QByteArray) Truncate(pos int) {
	if ptr.Pointer() != nil {
		C.QByteArray_Truncate(C.QtObjectPtr(ptr.Pointer()), C.int(pos))
	}
}

func (ptr *QByteArray) DestroyQByteArray() {
	if ptr.Pointer() != nil {
		C.QByteArray_DestroyQByteArray(C.QtObjectPtr(ptr.Pointer()))
	}
}
