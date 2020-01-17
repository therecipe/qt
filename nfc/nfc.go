// +build !minimal

package nfc

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "nfc.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtNfc_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtNfc_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type QNdefFilter struct {
	ptr unsafe.Pointer
}

type QNdefFilter_ITF interface {
	QNdefFilter_PTR() *QNdefFilter
}

func (ptr *QNdefFilter) QNdefFilter_PTR() *QNdefFilter {
	return ptr
}

func (ptr *QNdefFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QNdefFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQNdefFilter(ptr QNdefFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefFilter_PTR().Pointer()
	}
	return nil
}

func NewQNdefFilterFromPointer(ptr unsafe.Pointer) (n *QNdefFilter) {
	n = new(QNdefFilter)
	n.SetPointer(ptr)
	return
}
func NewQNdefFilter() *QNdefFilter {
	tmpValue := NewQNdefFilterFromPointer(C.QNdefFilter_NewQNdefFilter())
	qt.SetFinalizer(tmpValue, (*QNdefFilter).DestroyQNdefFilter)
	return tmpValue
}

func NewQNdefFilter2(other QNdefFilter_ITF) *QNdefFilter {
	tmpValue := NewQNdefFilterFromPointer(C.QNdefFilter_NewQNdefFilter2(PointerFromQNdefFilter(other)))
	qt.SetFinalizer(tmpValue, (*QNdefFilter).DestroyQNdefFilter)
	return tmpValue
}

func (ptr *QNdefFilter) AppendRecord2(typeNameFormat QNdefRecord__TypeNameFormat, ty core.QByteArray_ITF, min uint, max uint) {
	if ptr.Pointer() != nil {
		C.QNdefFilter_AppendRecord2(ptr.Pointer(), C.longlong(typeNameFormat), core.PointerFromQByteArray(ty), C.uint(uint32(min)), C.uint(uint32(max)))
	}
}

func (ptr *QNdefFilter) Clear() {
	if ptr.Pointer() != nil {
		C.QNdefFilter_Clear(ptr.Pointer())
	}
}

func (ptr *QNdefFilter) OrderMatch() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNdefFilter_OrderMatch(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNdefFilter) RecordCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QNdefFilter_RecordCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNdefFilter) SetOrderMatch(on bool) {
	if ptr.Pointer() != nil {
		C.QNdefFilter_SetOrderMatch(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(on))))
	}
}

func (ptr *QNdefFilter) DestroyQNdefFilter() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNdefFilter_DestroyQNdefFilter(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QNdefMessage struct {
	ptr unsafe.Pointer
}

type QNdefMessage_ITF interface {
	QNdefMessage_PTR() *QNdefMessage
}

func (ptr *QNdefMessage) QNdefMessage_PTR() *QNdefMessage {
	return ptr
}

func (ptr *QNdefMessage) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QNdefMessage) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQNdefMessage(ptr QNdefMessage_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefMessage_PTR().Pointer()
	}
	return nil
}

func NewQNdefMessageFromPointer(ptr unsafe.Pointer) (n *QNdefMessage) {
	n = new(QNdefMessage)
	n.SetPointer(ptr)
	return
}
func (ptr *QNdefMessage) DestroyQNdefMessage() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQNdefMessage() *QNdefMessage {
	tmpValue := NewQNdefMessageFromPointer(C.QNdefMessage_NewQNdefMessage())
	qt.SetFinalizer(tmpValue, (*QNdefMessage).DestroyQNdefMessage)
	return tmpValue
}

func NewQNdefMessage2(record QNdefRecord_ITF) *QNdefMessage {
	tmpValue := NewQNdefMessageFromPointer(C.QNdefMessage_NewQNdefMessage2(PointerFromQNdefRecord(record)))
	qt.SetFinalizer(tmpValue, (*QNdefMessage).DestroyQNdefMessage)
	return tmpValue
}

func NewQNdefMessage3(message QNdefMessage_ITF) *QNdefMessage {
	tmpValue := NewQNdefMessageFromPointer(C.QNdefMessage_NewQNdefMessage3(PointerFromQNdefMessage(message)))
	qt.SetFinalizer(tmpValue, (*QNdefMessage).DestroyQNdefMessage)
	return tmpValue
}

func NewQNdefMessage4(records []*QNdefRecord) *QNdefMessage {
	tmpValue := NewQNdefMessageFromPointer(C.QNdefMessage_NewQNdefMessage4(func() unsafe.Pointer {
		tmpList := NewQNdefMessageFromPointer(NewQNdefMessageFromPointer(nil).__QNdefMessage_records_newList4())
		for _, v := range records {
			tmpList.__QNdefMessage_records_setList4(v)
		}
		return tmpList.Pointer()
	}()))
	qt.SetFinalizer(tmpValue, (*QNdefMessage).DestroyQNdefMessage)
	return tmpValue
}

func QNdefMessage_FromByteArray(message core.QByteArray_ITF) *QNdefMessage {
	tmpValue := NewQNdefMessageFromPointer(C.QNdefMessage_QNdefMessage_FromByteArray(core.PointerFromQByteArray(message)))
	qt.SetFinalizer(tmpValue, (*QNdefMessage).DestroyQNdefMessage)
	return tmpValue
}

func (ptr *QNdefMessage) FromByteArray(message core.QByteArray_ITF) *QNdefMessage {
	tmpValue := NewQNdefMessageFromPointer(C.QNdefMessage_QNdefMessage_FromByteArray(core.PointerFromQByteArray(message)))
	qt.SetFinalizer(tmpValue, (*QNdefMessage).DestroyQNdefMessage)
	return tmpValue
}

func (ptr *QNdefMessage) ToByteArray() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QNdefMessage_ToByteArray(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNdefMessage) __QNdefMessage_records_atList4(i int) *QNdefRecord {
	if ptr.Pointer() != nil {
		tmpValue := NewQNdefRecordFromPointer(C.QNdefMessage___QNdefMessage_records_atList4(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QNdefRecord).DestroyQNdefRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QNdefMessage) __QNdefMessage_records_setList4(i QNdefRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QNdefMessage___QNdefMessage_records_setList4(ptr.Pointer(), PointerFromQNdefRecord(i))
	}
}

func (ptr *QNdefMessage) __QNdefMessage_records_newList4() unsafe.Pointer {
	return C.QNdefMessage___QNdefMessage_records_newList4(ptr.Pointer())
}

type QNdefNfcIconRecord struct {
	QNdefRecord
}

type QNdefNfcIconRecord_ITF interface {
	QNdefRecord_ITF
	QNdefNfcIconRecord_PTR() *QNdefNfcIconRecord
}

func (ptr *QNdefNfcIconRecord) QNdefNfcIconRecord_PTR() *QNdefNfcIconRecord {
	return ptr
}

func (ptr *QNdefNfcIconRecord) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefRecord_PTR().Pointer()
	}
	return nil
}

func (ptr *QNdefNfcIconRecord) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QNdefRecord_PTR().SetPointer(p)
	}
}

func PointerFromQNdefNfcIconRecord(ptr QNdefNfcIconRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefNfcIconRecord_PTR().Pointer()
	}
	return nil
}

func NewQNdefNfcIconRecordFromPointer(ptr unsafe.Pointer) (n *QNdefNfcIconRecord) {
	n = new(QNdefNfcIconRecord)
	n.SetPointer(ptr)
	return
}
func (ptr *QNdefNfcIconRecord) DestroyQNdefNfcIconRecord() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QNdefNfcSmartPosterRecord struct {
	QNdefRecord
}

type QNdefNfcSmartPosterRecord_ITF interface {
	QNdefRecord_ITF
	QNdefNfcSmartPosterRecord_PTR() *QNdefNfcSmartPosterRecord
}

func (ptr *QNdefNfcSmartPosterRecord) QNdefNfcSmartPosterRecord_PTR() *QNdefNfcSmartPosterRecord {
	return ptr
}

func (ptr *QNdefNfcSmartPosterRecord) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefRecord_PTR().Pointer()
	}
	return nil
}

func (ptr *QNdefNfcSmartPosterRecord) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QNdefRecord_PTR().SetPointer(p)
	}
}

func PointerFromQNdefNfcSmartPosterRecord(ptr QNdefNfcSmartPosterRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefNfcSmartPosterRecord_PTR().Pointer()
	}
	return nil
}

func NewQNdefNfcSmartPosterRecordFromPointer(ptr unsafe.Pointer) (n *QNdefNfcSmartPosterRecord) {
	n = new(QNdefNfcSmartPosterRecord)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QNdefNfcSmartPosterRecord__Action
//QNdefNfcSmartPosterRecord::Action
type QNdefNfcSmartPosterRecord__Action int64

const (
	QNdefNfcSmartPosterRecord__UnspecifiedAction QNdefNfcSmartPosterRecord__Action = QNdefNfcSmartPosterRecord__Action(-1)
	QNdefNfcSmartPosterRecord__DoAction          QNdefNfcSmartPosterRecord__Action = QNdefNfcSmartPosterRecord__Action(0)
	QNdefNfcSmartPosterRecord__SaveAction        QNdefNfcSmartPosterRecord__Action = QNdefNfcSmartPosterRecord__Action(1)
	QNdefNfcSmartPosterRecord__EditAction        QNdefNfcSmartPosterRecord__Action = QNdefNfcSmartPosterRecord__Action(2)
)

func NewQNdefNfcSmartPosterRecord() *QNdefNfcSmartPosterRecord {
	tmpValue := NewQNdefNfcSmartPosterRecordFromPointer(C.QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord())
	qt.SetFinalizer(tmpValue, (*QNdefNfcSmartPosterRecord).DestroyQNdefNfcSmartPosterRecord)
	return tmpValue
}

func NewQNdefNfcSmartPosterRecord2(other QNdefRecord_ITF) *QNdefNfcSmartPosterRecord {
	tmpValue := NewQNdefNfcSmartPosterRecordFromPointer(C.QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord2(PointerFromQNdefRecord(other)))
	qt.SetFinalizer(tmpValue, (*QNdefNfcSmartPosterRecord).DestroyQNdefNfcSmartPosterRecord)
	return tmpValue
}

func NewQNdefNfcSmartPosterRecord3(other QNdefNfcSmartPosterRecord_ITF) *QNdefNfcSmartPosterRecord {
	tmpValue := NewQNdefNfcSmartPosterRecordFromPointer(C.QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord3(PointerFromQNdefNfcSmartPosterRecord(other)))
	qt.SetFinalizer(tmpValue, (*QNdefNfcSmartPosterRecord).DestroyQNdefNfcSmartPosterRecord)
	return tmpValue
}

func (ptr *QNdefNfcSmartPosterRecord) Action() QNdefNfcSmartPosterRecord__Action {
	if ptr.Pointer() != nil {
		return QNdefNfcSmartPosterRecord__Action(C.QNdefNfcSmartPosterRecord_Action(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNdefNfcSmartPosterRecord) AddIcon(icon QNdefNfcIconRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_AddIcon(ptr.Pointer(), PointerFromQNdefNfcIconRecord(icon))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) AddIcon2(ty core.QByteArray_ITF, data core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_AddIcon2(ptr.Pointer(), core.PointerFromQByteArray(ty), core.PointerFromQByteArray(data))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) AddTitle(text QNdefNfcTextRecord_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNdefNfcSmartPosterRecord_AddTitle(ptr.Pointer(), PointerFromQNdefNfcTextRecord(text))) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) AddTitle2(text string, locale string, encoding QNdefNfcTextRecord__Encoding) bool {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		var localeC *C.char
		if locale != "" {
			localeC = C.CString(locale)
			defer C.free(unsafe.Pointer(localeC))
		}
		return int8(C.QNdefNfcSmartPosterRecord_AddTitle2(ptr.Pointer(), C.struct_QtNfc_PackedString{data: textC, len: C.longlong(len(text))}, C.struct_QtNfc_PackedString{data: localeC, len: C.longlong(len(locale))}, C.longlong(encoding))) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) HasAction() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNdefNfcSmartPosterRecord_HasAction(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) HasIcon(mimetype core.QByteArray_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNdefNfcSmartPosterRecord_HasIcon(ptr.Pointer(), core.PointerFromQByteArray(mimetype))) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) HasSize() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNdefNfcSmartPosterRecord_HasSize(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) HasTitle(locale string) bool {
	if ptr.Pointer() != nil {
		var localeC *C.char
		if locale != "" {
			localeC = C.CString(locale)
			defer C.free(unsafe.Pointer(localeC))
		}
		return int8(C.QNdefNfcSmartPosterRecord_HasTitle(ptr.Pointer(), C.struct_QtNfc_PackedString{data: localeC, len: C.longlong(len(locale))})) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) HasTypeInfo() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNdefNfcSmartPosterRecord_HasTypeInfo(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) Icon(mimetype core.QByteArray_ITF) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QNdefNfcSmartPosterRecord_Icon(ptr.Pointer(), core.PointerFromQByteArray(mimetype)))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNdefNfcSmartPosterRecord) IconCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QNdefNfcSmartPosterRecord_IconCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNdefNfcSmartPosterRecord) IconRecord(index int) *QNdefNfcIconRecord {
	if ptr.Pointer() != nil {
		tmpValue := NewQNdefNfcIconRecordFromPointer(C.QNdefNfcSmartPosterRecord_IconRecord(ptr.Pointer(), C.int(int32(index))))
		qt.SetFinalizer(tmpValue, (*QNdefNfcIconRecord).DestroyQNdefNfcIconRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QNdefNfcSmartPosterRecord) IconRecords() []*QNdefNfcIconRecord {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtNfc_PackedList) []*QNdefNfcIconRecord {
			out := make([]*QNdefNfcIconRecord, int(l.len))
			tmpList := NewQNdefNfcSmartPosterRecordFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__iconRecords_atList(i)
			}
			return out
		}(C.QNdefNfcSmartPosterRecord_IconRecords(ptr.Pointer()))
	}
	return make([]*QNdefNfcIconRecord, 0)
}

func (ptr *QNdefNfcSmartPosterRecord) RemoveIcon(icon QNdefNfcIconRecord_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNdefNfcSmartPosterRecord_RemoveIcon(ptr.Pointer(), PointerFromQNdefNfcIconRecord(icon))) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) RemoveIcon2(ty core.QByteArray_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNdefNfcSmartPosterRecord_RemoveIcon2(ptr.Pointer(), core.PointerFromQByteArray(ty))) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) RemoveTitle(text QNdefNfcTextRecord_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNdefNfcSmartPosterRecord_RemoveTitle(ptr.Pointer(), PointerFromQNdefNfcTextRecord(text))) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) RemoveTitle2(locale string) bool {
	if ptr.Pointer() != nil {
		var localeC *C.char
		if locale != "" {
			localeC = C.CString(locale)
			defer C.free(unsafe.Pointer(localeC))
		}
		return int8(C.QNdefNfcSmartPosterRecord_RemoveTitle2(ptr.Pointer(), C.struct_QtNfc_PackedString{data: localeC, len: C.longlong(len(locale))})) != 0
	}
	return false
}

func (ptr *QNdefNfcSmartPosterRecord) SetAction(act QNdefNfcSmartPosterRecord__Action) {
	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_SetAction(ptr.Pointer(), C.longlong(act))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) SetIcons(icons []*QNdefNfcIconRecord) {
	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_SetIcons(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQNdefNfcSmartPosterRecordFromPointer(NewQNdefNfcSmartPosterRecordFromPointer(nil).__setIcons_icons_newList())
			for _, v := range icons {
				tmpList.__setIcons_icons_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QNdefNfcSmartPosterRecord) SetSize(size uint) {
	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_SetSize(ptr.Pointer(), C.uint(uint32(size)))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) SetTitles(titles []*QNdefNfcTextRecord) {
	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_SetTitles(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQNdefNfcSmartPosterRecordFromPointer(NewQNdefNfcSmartPosterRecordFromPointer(nil).__setTitles_titles_newList())
			for _, v := range titles {
				tmpList.__setTitles_titles_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QNdefNfcSmartPosterRecord) SetTypeInfo(ty core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_SetTypeInfo(ptr.Pointer(), core.PointerFromQByteArray(ty))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) SetUri(url QNdefNfcUriRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_SetUri(ptr.Pointer(), PointerFromQNdefNfcUriRecord(url))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) SetUri2(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord_SetUri2(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) Size() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QNdefNfcSmartPosterRecord_Size(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNdefNfcSmartPosterRecord) Title(locale string) string {
	if ptr.Pointer() != nil {
		var localeC *C.char
		if locale != "" {
			localeC = C.CString(locale)
			defer C.free(unsafe.Pointer(localeC))
		}
		return cGoUnpackString(C.QNdefNfcSmartPosterRecord_Title(ptr.Pointer(), C.struct_QtNfc_PackedString{data: localeC, len: C.longlong(len(locale))}))
	}
	return ""
}

func (ptr *QNdefNfcSmartPosterRecord) TitleCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QNdefNfcSmartPosterRecord_TitleCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNdefNfcSmartPosterRecord) TitleRecord(index int) *QNdefNfcTextRecord {
	if ptr.Pointer() != nil {
		tmpValue := NewQNdefNfcTextRecordFromPointer(C.QNdefNfcSmartPosterRecord_TitleRecord(ptr.Pointer(), C.int(int32(index))))
		qt.SetFinalizer(tmpValue, (*QNdefNfcTextRecord).DestroyQNdefNfcTextRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QNdefNfcSmartPosterRecord) TypeInfo() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QNdefNfcSmartPosterRecord_TypeInfo(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNdefNfcSmartPosterRecord) Uri() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QNdefNfcSmartPosterRecord_Uri(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QNdefNfcSmartPosterRecord) UriRecord() *QNdefNfcUriRecord {
	if ptr.Pointer() != nil {
		tmpValue := NewQNdefNfcUriRecordFromPointer(C.QNdefNfcSmartPosterRecord_UriRecord(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QNdefNfcUriRecord).DestroyQNdefNfcUriRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QNdefNfcSmartPosterRecord) DestroyQNdefNfcSmartPosterRecord() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNdefNfcSmartPosterRecord_DestroyQNdefNfcSmartPosterRecord(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNdefNfcSmartPosterRecord) __iconRecords_atList(i int) *QNdefNfcIconRecord {
	if ptr.Pointer() != nil {
		tmpValue := NewQNdefNfcIconRecordFromPointer(C.QNdefNfcSmartPosterRecord___iconRecords_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QNdefNfcIconRecord).DestroyQNdefNfcIconRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QNdefNfcSmartPosterRecord) __iconRecords_setList(i QNdefNfcIconRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord___iconRecords_setList(ptr.Pointer(), PointerFromQNdefNfcIconRecord(i))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) __iconRecords_newList() unsafe.Pointer {
	return C.QNdefNfcSmartPosterRecord___iconRecords_newList(ptr.Pointer())
}

func (ptr *QNdefNfcSmartPosterRecord) __setIcons_icons_atList(i int) *QNdefNfcIconRecord {
	if ptr.Pointer() != nil {
		tmpValue := NewQNdefNfcIconRecordFromPointer(C.QNdefNfcSmartPosterRecord___setIcons_icons_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QNdefNfcIconRecord).DestroyQNdefNfcIconRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QNdefNfcSmartPosterRecord) __setIcons_icons_setList(i QNdefNfcIconRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord___setIcons_icons_setList(ptr.Pointer(), PointerFromQNdefNfcIconRecord(i))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) __setIcons_icons_newList() unsafe.Pointer {
	return C.QNdefNfcSmartPosterRecord___setIcons_icons_newList(ptr.Pointer())
}

func (ptr *QNdefNfcSmartPosterRecord) __setTitles_titles_atList(i int) *QNdefNfcTextRecord {
	if ptr.Pointer() != nil {
		tmpValue := NewQNdefNfcTextRecordFromPointer(C.QNdefNfcSmartPosterRecord___setTitles_titles_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QNdefNfcTextRecord).DestroyQNdefNfcTextRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QNdefNfcSmartPosterRecord) __setTitles_titles_setList(i QNdefNfcTextRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord___setTitles_titles_setList(ptr.Pointer(), PointerFromQNdefNfcTextRecord(i))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) __setTitles_titles_newList() unsafe.Pointer {
	return C.QNdefNfcSmartPosterRecord___setTitles_titles_newList(ptr.Pointer())
}

func (ptr *QNdefNfcSmartPosterRecord) __titleRecords_atList(i int) *QNdefNfcTextRecord {
	if ptr.Pointer() != nil {
		tmpValue := NewQNdefNfcTextRecordFromPointer(C.QNdefNfcSmartPosterRecord___titleRecords_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QNdefNfcTextRecord).DestroyQNdefNfcTextRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QNdefNfcSmartPosterRecord) __titleRecords_setList(i QNdefNfcTextRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QNdefNfcSmartPosterRecord___titleRecords_setList(ptr.Pointer(), PointerFromQNdefNfcTextRecord(i))
	}
}

func (ptr *QNdefNfcSmartPosterRecord) __titleRecords_newList() unsafe.Pointer {
	return C.QNdefNfcSmartPosterRecord___titleRecords_newList(ptr.Pointer())
}

type QNdefNfcTextRecord struct {
	QNdefRecord
}

type QNdefNfcTextRecord_ITF interface {
	QNdefRecord_ITF
	QNdefNfcTextRecord_PTR() *QNdefNfcTextRecord
}

func (ptr *QNdefNfcTextRecord) QNdefNfcTextRecord_PTR() *QNdefNfcTextRecord {
	return ptr
}

func (ptr *QNdefNfcTextRecord) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefRecord_PTR().Pointer()
	}
	return nil
}

func (ptr *QNdefNfcTextRecord) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QNdefRecord_PTR().SetPointer(p)
	}
}

func PointerFromQNdefNfcTextRecord(ptr QNdefNfcTextRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefNfcTextRecord_PTR().Pointer()
	}
	return nil
}

func NewQNdefNfcTextRecordFromPointer(ptr unsafe.Pointer) (n *QNdefNfcTextRecord) {
	n = new(QNdefNfcTextRecord)
	n.SetPointer(ptr)
	return
}
func (ptr *QNdefNfcTextRecord) DestroyQNdefNfcTextRecord() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QNdefNfcTextRecord__Encoding
//QNdefNfcTextRecord::Encoding
type QNdefNfcTextRecord__Encoding int64

const (
	QNdefNfcTextRecord__Utf8  QNdefNfcTextRecord__Encoding = QNdefNfcTextRecord__Encoding(0)
	QNdefNfcTextRecord__Utf16 QNdefNfcTextRecord__Encoding = QNdefNfcTextRecord__Encoding(1)
)

func NewQNdefNfcTextRecord() *QNdefNfcTextRecord {
	tmpValue := NewQNdefNfcTextRecordFromPointer(C.QNdefNfcTextRecord_NewQNdefNfcTextRecord())
	qt.SetFinalizer(tmpValue, (*QNdefNfcTextRecord).DestroyQNdefNfcTextRecord)
	return tmpValue
}

func NewQNdefNfcTextRecord2(other QNdefRecord_ITF) *QNdefNfcTextRecord {
	tmpValue := NewQNdefNfcTextRecordFromPointer(C.QNdefNfcTextRecord_NewQNdefNfcTextRecord2(PointerFromQNdefRecord(other)))
	qt.SetFinalizer(tmpValue, (*QNdefNfcTextRecord).DestroyQNdefNfcTextRecord)
	return tmpValue
}

func (ptr *QNdefNfcTextRecord) Encoding() QNdefNfcTextRecord__Encoding {
	if ptr.Pointer() != nil {
		return QNdefNfcTextRecord__Encoding(C.QNdefNfcTextRecord_Encoding(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNdefNfcTextRecord) Locale() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QNdefNfcTextRecord_Locale(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNdefNfcTextRecord) SetEncoding(encoding QNdefNfcTextRecord__Encoding) {
	if ptr.Pointer() != nil {
		C.QNdefNfcTextRecord_SetEncoding(ptr.Pointer(), C.longlong(encoding))
	}
}

func (ptr *QNdefNfcTextRecord) SetLocale(locale string) {
	if ptr.Pointer() != nil {
		var localeC *C.char
		if locale != "" {
			localeC = C.CString(locale)
			defer C.free(unsafe.Pointer(localeC))
		}
		C.QNdefNfcTextRecord_SetLocale(ptr.Pointer(), C.struct_QtNfc_PackedString{data: localeC, len: C.longlong(len(locale))})
	}
}

func (ptr *QNdefNfcTextRecord) SetText(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QNdefNfcTextRecord_SetText(ptr.Pointer(), C.struct_QtNfc_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

func (ptr *QNdefNfcTextRecord) Text() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QNdefNfcTextRecord_Text(ptr.Pointer()))
	}
	return ""
}

type QNdefNfcUriRecord struct {
	QNdefRecord
}

type QNdefNfcUriRecord_ITF interface {
	QNdefRecord_ITF
	QNdefNfcUriRecord_PTR() *QNdefNfcUriRecord
}

func (ptr *QNdefNfcUriRecord) QNdefNfcUriRecord_PTR() *QNdefNfcUriRecord {
	return ptr
}

func (ptr *QNdefNfcUriRecord) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefRecord_PTR().Pointer()
	}
	return nil
}

func (ptr *QNdefNfcUriRecord) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QNdefRecord_PTR().SetPointer(p)
	}
}

func PointerFromQNdefNfcUriRecord(ptr QNdefNfcUriRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefNfcUriRecord_PTR().Pointer()
	}
	return nil
}

func NewQNdefNfcUriRecordFromPointer(ptr unsafe.Pointer) (n *QNdefNfcUriRecord) {
	n = new(QNdefNfcUriRecord)
	n.SetPointer(ptr)
	return
}
func (ptr *QNdefNfcUriRecord) DestroyQNdefNfcUriRecord() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQNdefNfcUriRecord() *QNdefNfcUriRecord {
	tmpValue := NewQNdefNfcUriRecordFromPointer(C.QNdefNfcUriRecord_NewQNdefNfcUriRecord())
	qt.SetFinalizer(tmpValue, (*QNdefNfcUriRecord).DestroyQNdefNfcUriRecord)
	return tmpValue
}

func NewQNdefNfcUriRecord2(other QNdefRecord_ITF) *QNdefNfcUriRecord {
	tmpValue := NewQNdefNfcUriRecordFromPointer(C.QNdefNfcUriRecord_NewQNdefNfcUriRecord2(PointerFromQNdefRecord(other)))
	qt.SetFinalizer(tmpValue, (*QNdefNfcUriRecord).DestroyQNdefNfcUriRecord)
	return tmpValue
}

func (ptr *QNdefNfcUriRecord) SetUri(uri core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QNdefNfcUriRecord_SetUri(ptr.Pointer(), core.PointerFromQUrl(uri))
	}
}

func (ptr *QNdefNfcUriRecord) Uri() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QNdefNfcUriRecord_Uri(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

type QNdefRecord struct {
	ptr unsafe.Pointer
}

type QNdefRecord_ITF interface {
	QNdefRecord_PTR() *QNdefRecord
}

func (ptr *QNdefRecord) QNdefRecord_PTR() *QNdefRecord {
	return ptr
}

func (ptr *QNdefRecord) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QNdefRecord) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQNdefRecord(ptr QNdefRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNdefRecord_PTR().Pointer()
	}
	return nil
}

func NewQNdefRecordFromPointer(ptr unsafe.Pointer) (n *QNdefRecord) {
	n = new(QNdefRecord)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QNdefRecord__TypeNameFormat
//QNdefRecord::TypeNameFormat
type QNdefRecord__TypeNameFormat int64

const (
	QNdefRecord__Empty       QNdefRecord__TypeNameFormat = QNdefRecord__TypeNameFormat(0x00)
	QNdefRecord__NfcRtd      QNdefRecord__TypeNameFormat = QNdefRecord__TypeNameFormat(0x01)
	QNdefRecord__Mime        QNdefRecord__TypeNameFormat = QNdefRecord__TypeNameFormat(0x02)
	QNdefRecord__Uri         QNdefRecord__TypeNameFormat = QNdefRecord__TypeNameFormat(0x03)
	QNdefRecord__ExternalRtd QNdefRecord__TypeNameFormat = QNdefRecord__TypeNameFormat(0x04)
	QNdefRecord__Unknown     QNdefRecord__TypeNameFormat = QNdefRecord__TypeNameFormat(0x05)
)

func NewQNdefRecord() *QNdefRecord {
	tmpValue := NewQNdefRecordFromPointer(C.QNdefRecord_NewQNdefRecord())
	qt.SetFinalizer(tmpValue, (*QNdefRecord).DestroyQNdefRecord)
	return tmpValue
}

func NewQNdefRecord2(other QNdefRecord_ITF) *QNdefRecord {
	tmpValue := NewQNdefRecordFromPointer(C.QNdefRecord_NewQNdefRecord2(PointerFromQNdefRecord(other)))
	qt.SetFinalizer(tmpValue, (*QNdefRecord).DestroyQNdefRecord)
	return tmpValue
}

func (ptr *QNdefRecord) Id() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QNdefRecord_Id(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNdefRecord) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNdefRecord_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNdefRecord) Payload() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QNdefRecord_Payload(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNdefRecord) SetId(id core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNdefRecord_SetId(ptr.Pointer(), core.PointerFromQByteArray(id))
	}
}

func (ptr *QNdefRecord) SetPayload(payload core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNdefRecord_SetPayload(ptr.Pointer(), core.PointerFromQByteArray(payload))
	}
}

func (ptr *QNdefRecord) SetType(ty core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNdefRecord_SetType(ptr.Pointer(), core.PointerFromQByteArray(ty))
	}
}

func (ptr *QNdefRecord) SetTypeNameFormat(typeNameFormat QNdefRecord__TypeNameFormat) {
	if ptr.Pointer() != nil {
		C.QNdefRecord_SetTypeNameFormat(ptr.Pointer(), C.longlong(typeNameFormat))
	}
}

func (ptr *QNdefRecord) Type() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QNdefRecord_Type(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNdefRecord) TypeNameFormat() QNdefRecord__TypeNameFormat {
	if ptr.Pointer() != nil {
		return QNdefRecord__TypeNameFormat(C.QNdefRecord_TypeNameFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNdefRecord) DestroyQNdefRecord() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNdefRecord_DestroyQNdefRecord(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QNearFieldManager struct {
	core.QObject
}

type QNearFieldManager_ITF interface {
	core.QObject_ITF
	QNearFieldManager_PTR() *QNearFieldManager
}

func (ptr *QNearFieldManager) QNearFieldManager_PTR() *QNearFieldManager {
	return ptr
}

func (ptr *QNearFieldManager) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QNearFieldManager) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQNearFieldManager(ptr QNearFieldManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNearFieldManager_PTR().Pointer()
	}
	return nil
}

func NewQNearFieldManagerFromPointer(ptr unsafe.Pointer) (n *QNearFieldManager) {
	n = new(QNearFieldManager)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QNearFieldManager__AdapterState
//QNearFieldManager::AdapterState
type QNearFieldManager__AdapterState int64

const (
	QNearFieldManager__Offline    QNearFieldManager__AdapterState = QNearFieldManager__AdapterState(1)
	QNearFieldManager__TurningOn  QNearFieldManager__AdapterState = QNearFieldManager__AdapterState(2)
	QNearFieldManager__Online     QNearFieldManager__AdapterState = QNearFieldManager__AdapterState(3)
	QNearFieldManager__TurningOff QNearFieldManager__AdapterState = QNearFieldManager__AdapterState(4)
)

//go:generate stringer -type=QNearFieldManager__TargetAccessMode
//QNearFieldManager::TargetAccessMode
type QNearFieldManager__TargetAccessMode int64

const (
	QNearFieldManager__NoTargetAccess              QNearFieldManager__TargetAccessMode = QNearFieldManager__TargetAccessMode(0x00)
	QNearFieldManager__NdefReadTargetAccess        QNearFieldManager__TargetAccessMode = QNearFieldManager__TargetAccessMode(0x01)
	QNearFieldManager__NdefWriteTargetAccess       QNearFieldManager__TargetAccessMode = QNearFieldManager__TargetAccessMode(0x02)
	QNearFieldManager__TagTypeSpecificTargetAccess QNearFieldManager__TargetAccessMode = QNearFieldManager__TargetAccessMode(0x04)
)

func NewQNearFieldManager(parent core.QObject_ITF) *QNearFieldManager {
	tmpValue := NewQNearFieldManagerFromPointer(C.QNearFieldManager_NewQNearFieldManager(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QNearFieldManager) IsAvailable() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNearFieldManager_IsAvailable(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNearFieldManager) IsSupported() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNearFieldManager_IsSupported(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNearFieldManager) RegisterNdefMessageHandler(object core.QObject_ITF, method string) int {
	if ptr.Pointer() != nil {
		var methodC *C.char
		if method != "" {
			methodC = C.CString(method)
			defer C.free(unsafe.Pointer(methodC))
		}
		return int(int32(C.QNearFieldManager_RegisterNdefMessageHandler(ptr.Pointer(), core.PointerFromQObject(object), methodC)))
	}
	return 0
}

func (ptr *QNearFieldManager) RegisterNdefMessageHandler2(typeNameFormat QNdefRecord__TypeNameFormat, ty core.QByteArray_ITF, object core.QObject_ITF, method string) int {
	if ptr.Pointer() != nil {
		var methodC *C.char
		if method != "" {
			methodC = C.CString(method)
			defer C.free(unsafe.Pointer(methodC))
		}
		return int(int32(C.QNearFieldManager_RegisterNdefMessageHandler2(ptr.Pointer(), C.longlong(typeNameFormat), core.PointerFromQByteArray(ty), core.PointerFromQObject(object), methodC)))
	}
	return 0
}

func (ptr *QNearFieldManager) RegisterNdefMessageHandler3(filter QNdefFilter_ITF, object core.QObject_ITF, method string) int {
	if ptr.Pointer() != nil {
		var methodC *C.char
		if method != "" {
			methodC = C.CString(method)
			defer C.free(unsafe.Pointer(methodC))
		}
		return int(int32(C.QNearFieldManager_RegisterNdefMessageHandler3(ptr.Pointer(), PointerFromQNdefFilter(filter), core.PointerFromQObject(object), methodC)))
	}
	return 0
}

func (ptr *QNearFieldManager) SetTargetAccessModes(accessModes QNearFieldManager__TargetAccessMode) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_SetTargetAccessModes(ptr.Pointer(), C.longlong(accessModes))
	}
}

func (ptr *QNearFieldManager) StartTargetDetection() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNearFieldManager_StartTargetDetection(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNearFieldManager) StopTargetDetection() {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_StopTargetDetection(ptr.Pointer())
	}
}

func (ptr *QNearFieldManager) TargetAccessModes() QNearFieldManager__TargetAccessMode {
	if ptr.Pointer() != nil {
		return QNearFieldManager__TargetAccessMode(C.QNearFieldManager_TargetAccessModes(ptr.Pointer()))
	}
	return 0
}

//export callbackQNearFieldManager_TargetDetected
func callbackQNearFieldManager_TargetDetected(ptr unsafe.Pointer, target unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "targetDetected"); signal != nil {
		(*(*func(*QNearFieldTarget))(signal))(NewQNearFieldTargetFromPointer(target))
	}

}

func (ptr *QNearFieldManager) ConnectTargetDetected(f func(target *QNearFieldTarget)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "targetDetected") {
			C.QNearFieldManager_ConnectTargetDetected(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "targetDetected")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "targetDetected"); signal != nil {
			f := func(target *QNearFieldTarget) {
				(*(*func(*QNearFieldTarget))(signal))(target)
				f(target)
			}
			qt.ConnectSignal(ptr.Pointer(), "targetDetected", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "targetDetected", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNearFieldManager) DisconnectTargetDetected() {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_DisconnectTargetDetected(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "targetDetected")
	}
}

func (ptr *QNearFieldManager) TargetDetected(target QNearFieldTarget_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_TargetDetected(ptr.Pointer(), PointerFromQNearFieldTarget(target))
	}
}

//export callbackQNearFieldManager_TargetLost
func callbackQNearFieldManager_TargetLost(ptr unsafe.Pointer, target unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "targetLost"); signal != nil {
		(*(*func(*QNearFieldTarget))(signal))(NewQNearFieldTargetFromPointer(target))
	}

}

func (ptr *QNearFieldManager) ConnectTargetLost(f func(target *QNearFieldTarget)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "targetLost") {
			C.QNearFieldManager_ConnectTargetLost(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "targetLost")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "targetLost"); signal != nil {
			f := func(target *QNearFieldTarget) {
				(*(*func(*QNearFieldTarget))(signal))(target)
				f(target)
			}
			qt.ConnectSignal(ptr.Pointer(), "targetLost", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "targetLost", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNearFieldManager) DisconnectTargetLost() {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_DisconnectTargetLost(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "targetLost")
	}
}

func (ptr *QNearFieldManager) TargetLost(target QNearFieldTarget_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_TargetLost(ptr.Pointer(), PointerFromQNearFieldTarget(target))
	}
}

func (ptr *QNearFieldManager) UnregisterNdefMessageHandler(handlerId int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNearFieldManager_UnregisterNdefMessageHandler(ptr.Pointer(), C.int(int32(handlerId)))) != 0
	}
	return false
}

//export callbackQNearFieldManager_DestroyQNearFieldManager
func callbackQNearFieldManager_DestroyQNearFieldManager(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QNearFieldManager"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNearFieldManagerFromPointer(ptr).DestroyQNearFieldManagerDefault()
	}
}

func (ptr *QNearFieldManager) ConnectDestroyQNearFieldManager(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QNearFieldManager"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QNearFieldManager", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QNearFieldManager", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNearFieldManager) DisconnectDestroyQNearFieldManager() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QNearFieldManager")
	}
}

func (ptr *QNearFieldManager) DestroyQNearFieldManager() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNearFieldManager_DestroyQNearFieldManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNearFieldManager) DestroyQNearFieldManagerDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNearFieldManager_DestroyQNearFieldManagerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNearFieldManager) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QNearFieldManager___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNearFieldManager) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QNearFieldManager) __children_newList() unsafe.Pointer {
	return C.QNearFieldManager___children_newList(ptr.Pointer())
}

func (ptr *QNearFieldManager) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QNearFieldManager___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNearFieldManager) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QNearFieldManager) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QNearFieldManager___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QNearFieldManager) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QNearFieldManager___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNearFieldManager) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QNearFieldManager) __findChildren_newList() unsafe.Pointer {
	return C.QNearFieldManager___findChildren_newList(ptr.Pointer())
}

func (ptr *QNearFieldManager) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QNearFieldManager___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNearFieldManager) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QNearFieldManager) __findChildren_newList3() unsafe.Pointer {
	return C.QNearFieldManager___findChildren_newList3(ptr.Pointer())
}

//export callbackQNearFieldManager_ChildEvent
func callbackQNearFieldManager_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNearFieldManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNearFieldManager) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNearFieldManager_ConnectNotify
func callbackQNearFieldManager_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNearFieldManagerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNearFieldManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNearFieldManager_CustomEvent
func callbackQNearFieldManager_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQNearFieldManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNearFieldManager) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNearFieldManager_DeleteLater
func callbackQNearFieldManager_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNearFieldManagerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNearFieldManager) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNearFieldManager_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQNearFieldManager_Destroyed
func callbackQNearFieldManager_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQNearFieldManager_DisconnectNotify
func callbackQNearFieldManager_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNearFieldManagerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNearFieldManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNearFieldManager_Event
func callbackQNearFieldManager_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNearFieldManagerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNearFieldManager) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNearFieldManager_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQNearFieldManager_EventFilter
func callbackQNearFieldManager_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNearFieldManagerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNearFieldManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNearFieldManager_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQNearFieldManager_MetaObject
func callbackQNearFieldManager_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQNearFieldManagerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNearFieldManager) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNearFieldManager_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQNearFieldManager_ObjectNameChanged
func callbackQNearFieldManager_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtNfc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQNearFieldManager_TimerEvent
func callbackQNearFieldManager_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNearFieldManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNearFieldManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QNearFieldShareManager struct {
	core.QObject
}

type QNearFieldShareManager_ITF interface {
	core.QObject_ITF
	QNearFieldShareManager_PTR() *QNearFieldShareManager
}

func (ptr *QNearFieldShareManager) QNearFieldShareManager_PTR() *QNearFieldShareManager {
	return ptr
}

func (ptr *QNearFieldShareManager) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QNearFieldShareManager) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQNearFieldShareManager(ptr QNearFieldShareManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNearFieldShareManager_PTR().Pointer()
	}
	return nil
}

func NewQNearFieldShareManagerFromPointer(ptr unsafe.Pointer) (n *QNearFieldShareManager) {
	n = new(QNearFieldShareManager)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QNearFieldShareManager__ShareMode
//QNearFieldShareManager::ShareMode
type QNearFieldShareManager__ShareMode int64

const (
	QNearFieldShareManager__NoShare   QNearFieldShareManager__ShareMode = QNearFieldShareManager__ShareMode(0x00)
	QNearFieldShareManager__NdefShare QNearFieldShareManager__ShareMode = QNearFieldShareManager__ShareMode(0x01)
	QNearFieldShareManager__FileShare QNearFieldShareManager__ShareMode = QNearFieldShareManager__ShareMode(0x02)
)

//go:generate stringer -type=QNearFieldShareManager__ShareError
//QNearFieldShareManager::ShareError
type QNearFieldShareManager__ShareError int64

const (
	QNearFieldShareManager__NoError                     QNearFieldShareManager__ShareError = QNearFieldShareManager__ShareError(0)
	QNearFieldShareManager__UnknownError                QNearFieldShareManager__ShareError = QNearFieldShareManager__ShareError(1)
	QNearFieldShareManager__InvalidShareContentError    QNearFieldShareManager__ShareError = QNearFieldShareManager__ShareError(2)
	QNearFieldShareManager__ShareCanceledError          QNearFieldShareManager__ShareError = QNearFieldShareManager__ShareError(3)
	QNearFieldShareManager__ShareInterruptedError       QNearFieldShareManager__ShareError = QNearFieldShareManager__ShareError(4)
	QNearFieldShareManager__ShareRejectedError          QNearFieldShareManager__ShareError = QNearFieldShareManager__ShareError(5)
	QNearFieldShareManager__UnsupportedShareModeError   QNearFieldShareManager__ShareError = QNearFieldShareManager__ShareError(6)
	QNearFieldShareManager__ShareAlreadyInProgressError QNearFieldShareManager__ShareError = QNearFieldShareManager__ShareError(7)
	QNearFieldShareManager__SharePermissionDeniedError  QNearFieldShareManager__ShareError = QNearFieldShareManager__ShareError(8)
)

func NewQNearFieldShareManager(parent core.QObject_ITF) *QNearFieldShareManager {
	tmpValue := NewQNearFieldShareManagerFromPointer(C.QNearFieldShareManager_NewQNearFieldShareManager(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQNearFieldShareManager_Error
func callbackQNearFieldShareManager_Error(ptr unsafe.Pointer, error C.longlong) {
	if signal := qt.GetSignal(ptr, "error"); signal != nil {
		(*(*func(QNearFieldShareManager__ShareError))(signal))(QNearFieldShareManager__ShareError(error))
	}

}

func (ptr *QNearFieldShareManager) ConnectError(f func(error QNearFieldShareManager__ShareError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error") {
			C.QNearFieldShareManager_ConnectError(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "error")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error"); signal != nil {
			f := func(error QNearFieldShareManager__ShareError) {
				(*(*func(QNearFieldShareManager__ShareError))(signal))(error)
				f(error)
			}
			qt.ConnectSignal(ptr.Pointer(), "error", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNearFieldShareManager) DisconnectError() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "error")
	}
}

func (ptr *QNearFieldShareManager) Error(error QNearFieldShareManager__ShareError) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_Error(ptr.Pointer(), C.longlong(error))
	}
}

func (ptr *QNearFieldShareManager) SetShareModes(mode QNearFieldShareManager__ShareMode) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_SetShareModes(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QNearFieldShareManager) ShareError() QNearFieldShareManager__ShareError {
	if ptr.Pointer() != nil {
		return QNearFieldShareManager__ShareError(C.QNearFieldShareManager_ShareError(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNearFieldShareManager) ShareModes() QNearFieldShareManager__ShareMode {
	if ptr.Pointer() != nil {
		return QNearFieldShareManager__ShareMode(C.QNearFieldShareManager_ShareModes(ptr.Pointer()))
	}
	return 0
}

//export callbackQNearFieldShareManager_ShareModesChanged
func callbackQNearFieldShareManager_ShareModesChanged(ptr unsafe.Pointer, modes C.longlong) {
	if signal := qt.GetSignal(ptr, "shareModesChanged"); signal != nil {
		(*(*func(QNearFieldShareManager__ShareMode))(signal))(QNearFieldShareManager__ShareMode(modes))
	}

}

func (ptr *QNearFieldShareManager) ConnectShareModesChanged(f func(modes QNearFieldShareManager__ShareMode)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "shareModesChanged") {
			C.QNearFieldShareManager_ConnectShareModesChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "shareModesChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "shareModesChanged"); signal != nil {
			f := func(modes QNearFieldShareManager__ShareMode) {
				(*(*func(QNearFieldShareManager__ShareMode))(signal))(modes)
				f(modes)
			}
			qt.ConnectSignal(ptr.Pointer(), "shareModesChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "shareModesChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNearFieldShareManager) DisconnectShareModesChanged() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectShareModesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "shareModesChanged")
	}
}

func (ptr *QNearFieldShareManager) ShareModesChanged(modes QNearFieldShareManager__ShareMode) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ShareModesChanged(ptr.Pointer(), C.longlong(modes))
	}
}

func QNearFieldShareManager_SupportedShareModes() QNearFieldShareManager__ShareMode {
	return QNearFieldShareManager__ShareMode(C.QNearFieldShareManager_QNearFieldShareManager_SupportedShareModes())
}

func (ptr *QNearFieldShareManager) SupportedShareModes() QNearFieldShareManager__ShareMode {
	return QNearFieldShareManager__ShareMode(C.QNearFieldShareManager_QNearFieldShareManager_SupportedShareModes())
}

//export callbackQNearFieldShareManager_TargetDetected
func callbackQNearFieldShareManager_TargetDetected(ptr unsafe.Pointer, shareTarget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "targetDetected"); signal != nil {
		(*(*func(*QNearFieldShareTarget))(signal))(NewQNearFieldShareTargetFromPointer(shareTarget))
	}

}

func (ptr *QNearFieldShareManager) ConnectTargetDetected(f func(shareTarget *QNearFieldShareTarget)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "targetDetected") {
			C.QNearFieldShareManager_ConnectTargetDetected(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "targetDetected")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "targetDetected"); signal != nil {
			f := func(shareTarget *QNearFieldShareTarget) {
				(*(*func(*QNearFieldShareTarget))(signal))(shareTarget)
				f(shareTarget)
			}
			qt.ConnectSignal(ptr.Pointer(), "targetDetected", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "targetDetected", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNearFieldShareManager) DisconnectTargetDetected() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectTargetDetected(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "targetDetected")
	}
}

func (ptr *QNearFieldShareManager) TargetDetected(shareTarget QNearFieldShareTarget_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_TargetDetected(ptr.Pointer(), PointerFromQNearFieldShareTarget(shareTarget))
	}
}

//export callbackQNearFieldShareManager_DestroyQNearFieldShareManager
func callbackQNearFieldShareManager_DestroyQNearFieldShareManager(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QNearFieldShareManager"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNearFieldShareManagerFromPointer(ptr).DestroyQNearFieldShareManagerDefault()
	}
}

func (ptr *QNearFieldShareManager) ConnectDestroyQNearFieldShareManager(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QNearFieldShareManager"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QNearFieldShareManager", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QNearFieldShareManager", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNearFieldShareManager) DisconnectDestroyQNearFieldShareManager() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QNearFieldShareManager")
	}
}

func (ptr *QNearFieldShareManager) DestroyQNearFieldShareManager() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNearFieldShareManager_DestroyQNearFieldShareManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNearFieldShareManager) DestroyQNearFieldShareManagerDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNearFieldShareManager_DestroyQNearFieldShareManagerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNearFieldShareManager) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QNearFieldShareManager___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNearFieldShareManager) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QNearFieldShareManager) __children_newList() unsafe.Pointer {
	return C.QNearFieldShareManager___children_newList(ptr.Pointer())
}

func (ptr *QNearFieldShareManager) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QNearFieldShareManager___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNearFieldShareManager) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QNearFieldShareManager) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QNearFieldShareManager___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QNearFieldShareManager) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QNearFieldShareManager___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNearFieldShareManager) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QNearFieldShareManager) __findChildren_newList() unsafe.Pointer {
	return C.QNearFieldShareManager___findChildren_newList(ptr.Pointer())
}

func (ptr *QNearFieldShareManager) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QNearFieldShareManager___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNearFieldShareManager) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QNearFieldShareManager) __findChildren_newList3() unsafe.Pointer {
	return C.QNearFieldShareManager___findChildren_newList3(ptr.Pointer())
}

//export callbackQNearFieldShareManager_ChildEvent
func callbackQNearFieldShareManager_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNearFieldShareManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNearFieldShareManager) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNearFieldShareManager_ConnectNotify
func callbackQNearFieldShareManager_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNearFieldShareManagerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNearFieldShareManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNearFieldShareManager_CustomEvent
func callbackQNearFieldShareManager_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQNearFieldShareManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNearFieldShareManager) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNearFieldShareManager_DeleteLater
func callbackQNearFieldShareManager_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNearFieldShareManagerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNearFieldShareManager) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNearFieldShareManager_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQNearFieldShareManager_Destroyed
func callbackQNearFieldShareManager_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQNearFieldShareManager_DisconnectNotify
func callbackQNearFieldShareManager_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNearFieldShareManagerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNearFieldShareManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNearFieldShareManager_Event
func callbackQNearFieldShareManager_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNearFieldShareManagerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNearFieldShareManager) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNearFieldShareManager_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQNearFieldShareManager_EventFilter
func callbackQNearFieldShareManager_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNearFieldShareManagerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNearFieldShareManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNearFieldShareManager_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQNearFieldShareManager_MetaObject
func callbackQNearFieldShareManager_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQNearFieldShareManagerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNearFieldShareManager) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNearFieldShareManager_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQNearFieldShareManager_ObjectNameChanged
func callbackQNearFieldShareManager_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtNfc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQNearFieldShareManager_TimerEvent
func callbackQNearFieldShareManager_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNearFieldShareManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNearFieldShareManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QNearFieldShareTarget struct {
	core.QObject
}

type QNearFieldShareTarget_ITF interface {
	core.QObject_ITF
	QNearFieldShareTarget_PTR() *QNearFieldShareTarget
}

func (ptr *QNearFieldShareTarget) QNearFieldShareTarget_PTR() *QNearFieldShareTarget {
	return ptr
}

func (ptr *QNearFieldShareTarget) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QNearFieldShareTarget) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQNearFieldShareTarget(ptr QNearFieldShareTarget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNearFieldShareTarget_PTR().Pointer()
	}
	return nil
}

func NewQNearFieldShareTargetFromPointer(ptr unsafe.Pointer) (n *QNearFieldShareTarget) {
	n = new(QNearFieldShareTarget)
	n.SetPointer(ptr)
	return
}
func (ptr *QNearFieldShareTarget) Cancel() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_Cancel(ptr.Pointer())
	}
}

//export callbackQNearFieldShareTarget_Error
func callbackQNearFieldShareTarget_Error(ptr unsafe.Pointer, error C.longlong) {
	if signal := qt.GetSignal(ptr, "error"); signal != nil {
		(*(*func(QNearFieldShareManager__ShareError))(signal))(QNearFieldShareManager__ShareError(error))
	}

}

func (ptr *QNearFieldShareTarget) ConnectError(f func(error QNearFieldShareManager__ShareError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error") {
			C.QNearFieldShareTarget_ConnectError(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "error")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error"); signal != nil {
			f := func(error QNearFieldShareManager__ShareError) {
				(*(*func(QNearFieldShareManager__ShareError))(signal))(error)
				f(error)
			}
			qt.ConnectSignal(ptr.Pointer(), "error", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNearFieldShareTarget) DisconnectError() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "error")
	}
}

func (ptr *QNearFieldShareTarget) Error(error QNearFieldShareManager__ShareError) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_Error(ptr.Pointer(), C.longlong(error))
	}
}

func (ptr *QNearFieldShareTarget) IsShareInProgress() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNearFieldShareTarget_IsShareInProgress(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNearFieldShareTarget) Share(message QNdefMessage_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNearFieldShareTarget_Share(ptr.Pointer(), PointerFromQNdefMessage(message))) != 0
	}
	return false
}

func (ptr *QNearFieldShareTarget) Share2(files []*core.QFileInfo) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNearFieldShareTarget_Share2(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQNearFieldShareTargetFromPointer(NewQNearFieldShareTargetFromPointer(nil).__share_files_newList2())
			for _, v := range files {
				tmpList.__share_files_setList2(v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

func (ptr *QNearFieldShareTarget) ShareError() QNearFieldShareManager__ShareError {
	if ptr.Pointer() != nil {
		return QNearFieldShareManager__ShareError(C.QNearFieldShareTarget_ShareError(ptr.Pointer()))
	}
	return 0
}

//export callbackQNearFieldShareTarget_ShareFinished
func callbackQNearFieldShareTarget_ShareFinished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "shareFinished"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QNearFieldShareTarget) ConnectShareFinished(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "shareFinished") {
			C.QNearFieldShareTarget_ConnectShareFinished(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "shareFinished")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "shareFinished"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "shareFinished", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "shareFinished", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNearFieldShareTarget) DisconnectShareFinished() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_DisconnectShareFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "shareFinished")
	}
}

func (ptr *QNearFieldShareTarget) ShareFinished() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ShareFinished(ptr.Pointer())
	}
}

func (ptr *QNearFieldShareTarget) ShareModes() QNearFieldShareManager__ShareMode {
	if ptr.Pointer() != nil {
		return QNearFieldShareManager__ShareMode(C.QNearFieldShareTarget_ShareModes(ptr.Pointer()))
	}
	return 0
}

//export callbackQNearFieldShareTarget_DestroyQNearFieldShareTarget
func callbackQNearFieldShareTarget_DestroyQNearFieldShareTarget(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QNearFieldShareTarget"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNearFieldShareTargetFromPointer(ptr).DestroyQNearFieldShareTargetDefault()
	}
}

func (ptr *QNearFieldShareTarget) ConnectDestroyQNearFieldShareTarget(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QNearFieldShareTarget"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QNearFieldShareTarget", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QNearFieldShareTarget", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNearFieldShareTarget) DisconnectDestroyQNearFieldShareTarget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QNearFieldShareTarget")
	}
}

func (ptr *QNearFieldShareTarget) DestroyQNearFieldShareTarget() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNearFieldShareTarget_DestroyQNearFieldShareTarget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNearFieldShareTarget) DestroyQNearFieldShareTargetDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNearFieldShareTarget_DestroyQNearFieldShareTargetDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNearFieldShareTarget) __share_files_atList2(i int) *core.QFileInfo {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQFileInfoFromPointer(C.QNearFieldShareTarget___share_files_atList2(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QFileInfo).DestroyQFileInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QNearFieldShareTarget) __share_files_setList2(i core.QFileInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget___share_files_setList2(ptr.Pointer(), core.PointerFromQFileInfo(i))
	}
}

func (ptr *QNearFieldShareTarget) __share_files_newList2() unsafe.Pointer {
	return C.QNearFieldShareTarget___share_files_newList2(ptr.Pointer())
}

func (ptr *QNearFieldShareTarget) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QNearFieldShareTarget___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNearFieldShareTarget) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QNearFieldShareTarget) __children_newList() unsafe.Pointer {
	return C.QNearFieldShareTarget___children_newList(ptr.Pointer())
}

func (ptr *QNearFieldShareTarget) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QNearFieldShareTarget___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNearFieldShareTarget) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QNearFieldShareTarget) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QNearFieldShareTarget___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QNearFieldShareTarget) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QNearFieldShareTarget___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNearFieldShareTarget) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QNearFieldShareTarget) __findChildren_newList() unsafe.Pointer {
	return C.QNearFieldShareTarget___findChildren_newList(ptr.Pointer())
}

func (ptr *QNearFieldShareTarget) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QNearFieldShareTarget___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNearFieldShareTarget) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QNearFieldShareTarget) __findChildren_newList3() unsafe.Pointer {
	return C.QNearFieldShareTarget___findChildren_newList3(ptr.Pointer())
}

//export callbackQNearFieldShareTarget_ChildEvent
func callbackQNearFieldShareTarget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNearFieldShareTargetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNearFieldShareTarget) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNearFieldShareTarget_ConnectNotify
func callbackQNearFieldShareTarget_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNearFieldShareTargetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNearFieldShareTarget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNearFieldShareTarget_CustomEvent
func callbackQNearFieldShareTarget_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQNearFieldShareTargetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNearFieldShareTarget) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNearFieldShareTarget_DeleteLater
func callbackQNearFieldShareTarget_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNearFieldShareTargetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNearFieldShareTarget) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNearFieldShareTarget_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQNearFieldShareTarget_Destroyed
func callbackQNearFieldShareTarget_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQNearFieldShareTarget_DisconnectNotify
func callbackQNearFieldShareTarget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNearFieldShareTargetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNearFieldShareTarget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNearFieldShareTarget_Event
func callbackQNearFieldShareTarget_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNearFieldShareTargetFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNearFieldShareTarget) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNearFieldShareTarget_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQNearFieldShareTarget_EventFilter
func callbackQNearFieldShareTarget_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNearFieldShareTargetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNearFieldShareTarget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNearFieldShareTarget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQNearFieldShareTarget_MetaObject
func callbackQNearFieldShareTarget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQNearFieldShareTargetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNearFieldShareTarget) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNearFieldShareTarget_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQNearFieldShareTarget_ObjectNameChanged
func callbackQNearFieldShareTarget_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtNfc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQNearFieldShareTarget_TimerEvent
func callbackQNearFieldShareTarget_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNearFieldShareTargetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNearFieldShareTarget) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QNearFieldTarget struct {
	core.QObject
}

type QNearFieldTarget_ITF interface {
	core.QObject_ITF
	QNearFieldTarget_PTR() *QNearFieldTarget
}

func (ptr *QNearFieldTarget) QNearFieldTarget_PTR() *QNearFieldTarget {
	return ptr
}

func (ptr *QNearFieldTarget) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QNearFieldTarget) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQNearFieldTarget(ptr QNearFieldTarget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNearFieldTarget_PTR().Pointer()
	}
	return nil
}

func NewQNearFieldTargetFromPointer(ptr unsafe.Pointer) (n *QNearFieldTarget) {
	n = new(QNearFieldTarget)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QNearFieldTarget__Type
//QNearFieldTarget::Type
type QNearFieldTarget__Type int64

const (
	QNearFieldTarget__ProprietaryTag QNearFieldTarget__Type = QNearFieldTarget__Type(0)
	QNearFieldTarget__NfcTagType1    QNearFieldTarget__Type = QNearFieldTarget__Type(1)
	QNearFieldTarget__NfcTagType2    QNearFieldTarget__Type = QNearFieldTarget__Type(2)
	QNearFieldTarget__NfcTagType3    QNearFieldTarget__Type = QNearFieldTarget__Type(3)
	QNearFieldTarget__NfcTagType4    QNearFieldTarget__Type = QNearFieldTarget__Type(4)
	QNearFieldTarget__MifareTag      QNearFieldTarget__Type = QNearFieldTarget__Type(5)
)

//go:generate stringer -type=QNearFieldTarget__AccessMethod
//QNearFieldTarget::AccessMethod
type QNearFieldTarget__AccessMethod int64

const (
	QNearFieldTarget__UnknownAccess         QNearFieldTarget__AccessMethod = QNearFieldTarget__AccessMethod(0x00)
	QNearFieldTarget__NdefAccess            QNearFieldTarget__AccessMethod = QNearFieldTarget__AccessMethod(0x01)
	QNearFieldTarget__TagTypeSpecificAccess QNearFieldTarget__AccessMethod = QNearFieldTarget__AccessMethod(0x02)
	QNearFieldTarget__LlcpAccess            QNearFieldTarget__AccessMethod = QNearFieldTarget__AccessMethod(0x04)
)

//go:generate stringer -type=QNearFieldTarget__Error
//QNearFieldTarget::Error
type QNearFieldTarget__Error int64

const (
	QNearFieldTarget__NoError                QNearFieldTarget__Error = QNearFieldTarget__Error(0)
	QNearFieldTarget__UnknownError           QNearFieldTarget__Error = QNearFieldTarget__Error(1)
	QNearFieldTarget__UnsupportedError       QNearFieldTarget__Error = QNearFieldTarget__Error(2)
	QNearFieldTarget__TargetOutOfRangeError  QNearFieldTarget__Error = QNearFieldTarget__Error(3)
	QNearFieldTarget__NoResponseError        QNearFieldTarget__Error = QNearFieldTarget__Error(4)
	QNearFieldTarget__ChecksumMismatchError  QNearFieldTarget__Error = QNearFieldTarget__Error(5)
	QNearFieldTarget__InvalidParametersError QNearFieldTarget__Error = QNearFieldTarget__Error(6)
	QNearFieldTarget__NdefReadError          QNearFieldTarget__Error = QNearFieldTarget__Error(7)
	QNearFieldTarget__NdefWriteError         QNearFieldTarget__Error = QNearFieldTarget__Error(8)
	QNearFieldTarget__CommandError           QNearFieldTarget__Error = QNearFieldTarget__Error(9)
)

func NewQNearFieldTarget(parent core.QObject_ITF) *QNearFieldTarget {
	tmpValue := NewQNearFieldTargetFromPointer(C.QNearFieldTarget_NewQNearFieldTarget(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQNearFieldTarget_AccessMethods
func callbackQNearFieldTarget_AccessMethods(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "accessMethods"); signal != nil {
		return C.longlong((*(*func() QNearFieldTarget__AccessMethod)(signal))())
	}

	return C.longlong(0)
}

func (ptr *QNearFieldTarget) ConnectAccessMethods(f func() QNearFieldTarget__AccessMethod) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "accessMethods"); signal != nil {
			f := func() QNearFieldTarget__AccessMethod {
				(*(*func() QNearFieldTarget__AccessMethod)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "accessMethods", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "accessMethods", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNearFieldTarget) DisconnectAccessMethods() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "accessMethods")
	}
}

func (ptr *QNearFieldTarget) AccessMethods() QNearFieldTarget__AccessMethod {
	if ptr.Pointer() != nil {
		return QNearFieldTarget__AccessMethod(C.QNearFieldTarget_AccessMethods(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNearFieldTarget) Disconnect() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNearFieldTarget_Disconnect(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQNearFieldTarget_Disconnected
func callbackQNearFieldTarget_Disconnected(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnected"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QNearFieldTarget) ConnectDisconnected(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "disconnected") {
			C.QNearFieldTarget_ConnectDisconnected(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "disconnected")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "disconnected"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "disconnected", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "disconnected", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNearFieldTarget) DisconnectDisconnected() {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "disconnected")
	}
}

func (ptr *QNearFieldTarget) Disconnected() {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_Disconnected(ptr.Pointer())
	}
}

//export callbackQNearFieldTarget_HasNdefMessage
func callbackQNearFieldTarget_HasNdefMessage(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasNdefMessage"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNearFieldTargetFromPointer(ptr).HasNdefMessageDefault())))
}

func (ptr *QNearFieldTarget) ConnectHasNdefMessage(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hasNdefMessage"); signal != nil {
			f := func() bool {
				(*(*func() bool)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "hasNdefMessage", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hasNdefMessage", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNearFieldTarget) DisconnectHasNdefMessage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hasNdefMessage")
	}
}

func (ptr *QNearFieldTarget) HasNdefMessage() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNearFieldTarget_HasNdefMessage(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNearFieldTarget) HasNdefMessageDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNearFieldTarget_HasNdefMessageDefault(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNearFieldTarget) IsProcessingCommand() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNearFieldTarget_IsProcessingCommand(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNearFieldTarget) KeepConnection() bool {
	if ptr.Pointer() != nil {
		return int8(C.QNearFieldTarget_KeepConnection(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNearFieldTarget) MaxCommandLength() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QNearFieldTarget_MaxCommandLength(ptr.Pointer())))
	}
	return 0
}

//export callbackQNearFieldTarget_NdefMessageRead
func callbackQNearFieldTarget_NdefMessageRead(ptr unsafe.Pointer, message unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "ndefMessageRead"); signal != nil {
		(*(*func(*QNdefMessage))(signal))(NewQNdefMessageFromPointer(message))
	}

}

func (ptr *QNearFieldTarget) ConnectNdefMessageRead(f func(message *QNdefMessage)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "ndefMessageRead") {
			C.QNearFieldTarget_ConnectNdefMessageRead(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "ndefMessageRead")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "ndefMessageRead"); signal != nil {
			f := func(message *QNdefMessage) {
				(*(*func(*QNdefMessage))(signal))(message)
				f(message)
			}
			qt.ConnectSignal(ptr.Pointer(), "ndefMessageRead", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "ndefMessageRead", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNearFieldTarget) DisconnectNdefMessageRead() {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_DisconnectNdefMessageRead(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "ndefMessageRead")
	}
}

func (ptr *QNearFieldTarget) NdefMessageRead(message QNdefMessage_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_NdefMessageRead(ptr.Pointer(), PointerFromQNdefMessage(message))
	}
}

//export callbackQNearFieldTarget_NdefMessagesWritten
func callbackQNearFieldTarget_NdefMessagesWritten(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "ndefMessagesWritten"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QNearFieldTarget) ConnectNdefMessagesWritten(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "ndefMessagesWritten") {
			C.QNearFieldTarget_ConnectNdefMessagesWritten(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "ndefMessagesWritten")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "ndefMessagesWritten"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "ndefMessagesWritten", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "ndefMessagesWritten", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNearFieldTarget) DisconnectNdefMessagesWritten() {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_DisconnectNdefMessagesWritten(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "ndefMessagesWritten")
	}
}

func (ptr *QNearFieldTarget) NdefMessagesWritten() {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_NdefMessagesWritten(ptr.Pointer())
	}
}

func (ptr *QNearFieldTarget) SetKeepConnection(isPersistent bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNearFieldTarget_SetKeepConnection(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isPersistent))))) != 0
	}
	return false
}

//export callbackQNearFieldTarget_Type
func callbackQNearFieldTarget_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QNearFieldTarget__Type)(signal))())
	}

	return C.longlong(0)
}

func (ptr *QNearFieldTarget) ConnectType(f func() QNearFieldTarget__Type) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QNearFieldTarget__Type {
				(*(*func() QNearFieldTarget__Type)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNearFieldTarget) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QNearFieldTarget) Type() QNearFieldTarget__Type {
	if ptr.Pointer() != nil {
		return QNearFieldTarget__Type(C.QNearFieldTarget_Type(ptr.Pointer()))
	}
	return 0
}

//export callbackQNearFieldTarget_Uid
func callbackQNearFieldTarget_Uid(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "uid"); signal != nil {
		return core.PointerFromQByteArray((*(*func() *core.QByteArray)(signal))())
	}

	return core.PointerFromQByteArray(core.NewQByteArray())
}

func (ptr *QNearFieldTarget) ConnectUid(f func() *core.QByteArray) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "uid"); signal != nil {
			f := func() *core.QByteArray {
				(*(*func() *core.QByteArray)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "uid", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "uid", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNearFieldTarget) DisconnectUid() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "uid")
	}
}

func (ptr *QNearFieldTarget) Uid() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QNearFieldTarget_Uid(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

//export callbackQNearFieldTarget_Url
func callbackQNearFieldTarget_Url(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "url"); signal != nil {
		return core.PointerFromQUrl((*(*func() *core.QUrl)(signal))())
	}

	return core.PointerFromQUrl(NewQNearFieldTargetFromPointer(ptr).UrlDefault())
}

func (ptr *QNearFieldTarget) ConnectUrl(f func() *core.QUrl) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "url"); signal != nil {
			f := func() *core.QUrl {
				(*(*func() *core.QUrl)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "url", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "url", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNearFieldTarget) DisconnectUrl() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "url")
	}
}

func (ptr *QNearFieldTarget) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QNearFieldTarget_Url(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QNearFieldTarget) UrlDefault() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QNearFieldTarget_UrlDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQNearFieldTarget_DestroyQNearFieldTarget
func callbackQNearFieldTarget_DestroyQNearFieldTarget(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QNearFieldTarget"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNearFieldTargetFromPointer(ptr).DestroyQNearFieldTargetDefault()
	}
}

func (ptr *QNearFieldTarget) ConnectDestroyQNearFieldTarget(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QNearFieldTarget"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QNearFieldTarget", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QNearFieldTarget", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNearFieldTarget) DisconnectDestroyQNearFieldTarget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QNearFieldTarget")
	}
}

func (ptr *QNearFieldTarget) DestroyQNearFieldTarget() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNearFieldTarget_DestroyQNearFieldTarget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNearFieldTarget) DestroyQNearFieldTargetDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNearFieldTarget_DestroyQNearFieldTargetDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNearFieldTarget) __sendCommands_commands_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QNearFieldTarget___sendCommands_commands_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNearFieldTarget) __sendCommands_commands_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget___sendCommands_commands_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QNearFieldTarget) __sendCommands_commands_newList() unsafe.Pointer {
	return C.QNearFieldTarget___sendCommands_commands_newList(ptr.Pointer())
}

func (ptr *QNearFieldTarget) __writeNdefMessages_messages_atList(i int) *QNdefMessage {
	if ptr.Pointer() != nil {
		tmpValue := NewQNdefMessageFromPointer(C.QNearFieldTarget___writeNdefMessages_messages_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QNdefMessage).DestroyQNdefMessage)
		return tmpValue
	}
	return nil
}

func (ptr *QNearFieldTarget) __writeNdefMessages_messages_setList(i QNdefMessage_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget___writeNdefMessages_messages_setList(ptr.Pointer(), PointerFromQNdefMessage(i))
	}
}

func (ptr *QNearFieldTarget) __writeNdefMessages_messages_newList() unsafe.Pointer {
	return C.QNearFieldTarget___writeNdefMessages_messages_newList(ptr.Pointer())
}

func (ptr *QNearFieldTarget) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QNearFieldTarget___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNearFieldTarget) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QNearFieldTarget) __children_newList() unsafe.Pointer {
	return C.QNearFieldTarget___children_newList(ptr.Pointer())
}

func (ptr *QNearFieldTarget) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QNearFieldTarget___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QNearFieldTarget) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QNearFieldTarget) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QNearFieldTarget___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QNearFieldTarget) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QNearFieldTarget___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNearFieldTarget) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QNearFieldTarget) __findChildren_newList() unsafe.Pointer {
	return C.QNearFieldTarget___findChildren_newList(ptr.Pointer())
}

func (ptr *QNearFieldTarget) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QNearFieldTarget___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNearFieldTarget) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QNearFieldTarget) __findChildren_newList3() unsafe.Pointer {
	return C.QNearFieldTarget___findChildren_newList3(ptr.Pointer())
}

//export callbackQNearFieldTarget_ChildEvent
func callbackQNearFieldTarget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNearFieldTargetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNearFieldTarget) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNearFieldTarget_ConnectNotify
func callbackQNearFieldTarget_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNearFieldTargetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNearFieldTarget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNearFieldTarget_CustomEvent
func callbackQNearFieldTarget_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQNearFieldTargetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNearFieldTarget) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNearFieldTarget_DeleteLater
func callbackQNearFieldTarget_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNearFieldTargetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNearFieldTarget) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNearFieldTarget_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQNearFieldTarget_Destroyed
func callbackQNearFieldTarget_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQNearFieldTarget_DisconnectNotify
func callbackQNearFieldTarget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNearFieldTargetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNearFieldTarget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNearFieldTarget_Event
func callbackQNearFieldTarget_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNearFieldTargetFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNearFieldTarget) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNearFieldTarget_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQNearFieldTarget_EventFilter
func callbackQNearFieldTarget_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNearFieldTargetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNearFieldTarget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QNearFieldTarget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQNearFieldTarget_MetaObject
func callbackQNearFieldTarget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQNearFieldTargetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNearFieldTarget) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNearFieldTarget_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQNearFieldTarget_ObjectNameChanged
func callbackQNearFieldTarget_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtNfc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQNearFieldTarget_TimerEvent
func callbackQNearFieldTarget_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNearFieldTargetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNearFieldTarget) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QQmlNdefRecord struct {
	core.QObject
}

type QQmlNdefRecord_ITF interface {
	core.QObject_ITF
	QQmlNdefRecord_PTR() *QQmlNdefRecord
}

func (ptr *QQmlNdefRecord) QQmlNdefRecord_PTR() *QQmlNdefRecord {
	return ptr
}

func (ptr *QQmlNdefRecord) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlNdefRecord) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQmlNdefRecord(ptr QQmlNdefRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlNdefRecord_PTR().Pointer()
	}
	return nil
}

func NewQQmlNdefRecordFromPointer(ptr unsafe.Pointer) (n *QQmlNdefRecord) {
	n = new(QQmlNdefRecord)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QQmlNdefRecord__TypeNameFormat
//QQmlNdefRecord::TypeNameFormat
type QQmlNdefRecord__TypeNameFormat int64

const (
	QQmlNdefRecord__Empty       QQmlNdefRecord__TypeNameFormat = QQmlNdefRecord__TypeNameFormat(QNdefRecord__Empty)
	QQmlNdefRecord__NfcRtd      QQmlNdefRecord__TypeNameFormat = QQmlNdefRecord__TypeNameFormat(QNdefRecord__NfcRtd)
	QQmlNdefRecord__Mime        QQmlNdefRecord__TypeNameFormat = QQmlNdefRecord__TypeNameFormat(QNdefRecord__Mime)
	QQmlNdefRecord__Uri         QQmlNdefRecord__TypeNameFormat = QQmlNdefRecord__TypeNameFormat(QNdefRecord__Uri)
	QQmlNdefRecord__ExternalRtd QQmlNdefRecord__TypeNameFormat = QQmlNdefRecord__TypeNameFormat(QNdefRecord__ExternalRtd)
	QQmlNdefRecord__Unknown     QQmlNdefRecord__TypeNameFormat = QQmlNdefRecord__TypeNameFormat(QNdefRecord__Unknown)
)

func NewQQmlNdefRecord(parent core.QObject_ITF) *QQmlNdefRecord {
	tmpValue := NewQQmlNdefRecordFromPointer(C.QQmlNdefRecord_NewQQmlNdefRecord(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlNdefRecord2(record QNdefRecord_ITF, parent core.QObject_ITF) *QQmlNdefRecord {
	tmpValue := NewQQmlNdefRecordFromPointer(C.QQmlNdefRecord_NewQQmlNdefRecord2(PointerFromQNdefRecord(record), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQmlNdefRecord) Record() *QNdefRecord {
	if ptr.Pointer() != nil {
		tmpValue := NewQNdefRecordFromPointer(C.QQmlNdefRecord_Record(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QNdefRecord).DestroyQNdefRecord)
		return tmpValue
	}
	return nil
}

//export callbackQQmlNdefRecord_RecordChanged
func callbackQQmlNdefRecord_RecordChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "recordChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QQmlNdefRecord) ConnectRecordChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "recordChanged") {
			C.QQmlNdefRecord_ConnectRecordChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "recordChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "recordChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "recordChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "recordChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QQmlNdefRecord) DisconnectRecordChanged() {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DisconnectRecordChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "recordChanged")
	}
}

func (ptr *QQmlNdefRecord) RecordChanged() {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_RecordChanged(ptr.Pointer())
	}
}

func (ptr *QQmlNdefRecord) SetRecord(record QNdefRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_SetRecord(ptr.Pointer(), PointerFromQNdefRecord(record))
	}
}

func (ptr *QQmlNdefRecord) SetType(newtype string) {
	if ptr.Pointer() != nil {
		var newtypeC *C.char
		if newtype != "" {
			newtypeC = C.CString(newtype)
			defer C.free(unsafe.Pointer(newtypeC))
		}
		C.QQmlNdefRecord_SetType(ptr.Pointer(), C.struct_QtNfc_PackedString{data: newtypeC, len: C.longlong(len(newtype))})
	}
}

func (ptr *QQmlNdefRecord) SetTypeNameFormat(newTypeNameFormat QQmlNdefRecord__TypeNameFormat) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_SetTypeNameFormat(ptr.Pointer(), C.longlong(newTypeNameFormat))
	}
}

func (ptr *QQmlNdefRecord) Type() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQmlNdefRecord_Type(ptr.Pointer()))
	}
	return ""
}

//export callbackQQmlNdefRecord_TypeChanged
func callbackQQmlNdefRecord_TypeChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "typeChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QQmlNdefRecord) ConnectTypeChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "typeChanged") {
			C.QQmlNdefRecord_ConnectTypeChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "typeChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "typeChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "typeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "typeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QQmlNdefRecord) DisconnectTypeChanged() {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DisconnectTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "typeChanged")
	}
}

func (ptr *QQmlNdefRecord) TypeChanged() {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_TypeChanged(ptr.Pointer())
	}
}

func (ptr *QQmlNdefRecord) TypeNameFormat() QQmlNdefRecord__TypeNameFormat {
	if ptr.Pointer() != nil {
		return QQmlNdefRecord__TypeNameFormat(C.QQmlNdefRecord_TypeNameFormat(ptr.Pointer()))
	}
	return 0
}

//export callbackQQmlNdefRecord_TypeNameFormatChanged
func callbackQQmlNdefRecord_TypeNameFormatChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "typeNameFormatChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QQmlNdefRecord) ConnectTypeNameFormatChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "typeNameFormatChanged") {
			C.QQmlNdefRecord_ConnectTypeNameFormatChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "typeNameFormatChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "typeNameFormatChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "typeNameFormatChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "typeNameFormatChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QQmlNdefRecord) DisconnectTypeNameFormatChanged() {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DisconnectTypeNameFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "typeNameFormatChanged")
	}
}

func (ptr *QQmlNdefRecord) TypeNameFormatChanged() {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_TypeNameFormatChanged(ptr.Pointer())
	}
}

//export callbackQQmlNdefRecord_DestroyQQmlNdefRecord
func callbackQQmlNdefRecord_DestroyQQmlNdefRecord(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QQmlNdefRecord"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQQmlNdefRecordFromPointer(ptr).DestroyQQmlNdefRecordDefault()
	}
}

func (ptr *QQmlNdefRecord) ConnectDestroyQQmlNdefRecord(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QQmlNdefRecord"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QQmlNdefRecord", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QQmlNdefRecord", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QQmlNdefRecord) DisconnectDestroyQQmlNdefRecord() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QQmlNdefRecord")
	}
}

func (ptr *QQmlNdefRecord) DestroyQQmlNdefRecord() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QQmlNdefRecord_DestroyQQmlNdefRecord(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlNdefRecord) DestroyQQmlNdefRecordDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QQmlNdefRecord_DestroyQQmlNdefRecordDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlNdefRecord) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlNdefRecord___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlNdefRecord) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlNdefRecord) __children_newList() unsafe.Pointer {
	return C.QQmlNdefRecord___children_newList(ptr.Pointer())
}

func (ptr *QQmlNdefRecord) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QQmlNdefRecord___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlNdefRecord) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QQmlNdefRecord) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QQmlNdefRecord___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QQmlNdefRecord) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlNdefRecord___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlNdefRecord) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlNdefRecord) __findChildren_newList() unsafe.Pointer {
	return C.QQmlNdefRecord___findChildren_newList(ptr.Pointer())
}

func (ptr *QQmlNdefRecord) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlNdefRecord___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlNdefRecord) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlNdefRecord) __findChildren_newList3() unsafe.Pointer {
	return C.QQmlNdefRecord___findChildren_newList3(ptr.Pointer())
}

//export callbackQQmlNdefRecord_ChildEvent
func callbackQQmlNdefRecord_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlNdefRecordFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlNdefRecord) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQmlNdefRecord_ConnectNotify
func callbackQQmlNdefRecord_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlNdefRecordFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlNdefRecord) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlNdefRecord_CustomEvent
func callbackQQmlNdefRecord_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlNdefRecordFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlNdefRecord) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQmlNdefRecord_DeleteLater
func callbackQQmlNdefRecord_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQQmlNdefRecordFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQmlNdefRecord) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QQmlNdefRecord_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQQmlNdefRecord_Destroyed
func callbackQQmlNdefRecord_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQQmlNdefRecord_DisconnectNotify
func callbackQQmlNdefRecord_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlNdefRecordFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlNdefRecord) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlNdefRecord_Event
func callbackQQmlNdefRecord_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlNdefRecordFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQmlNdefRecord) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlNdefRecord_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQQmlNdefRecord_EventFilter
func callbackQQmlNdefRecord_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlNdefRecordFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQmlNdefRecord) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlNdefRecord_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQQmlNdefRecord_MetaObject
func callbackQQmlNdefRecord_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQQmlNdefRecordFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQmlNdefRecord) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlNdefRecord_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQQmlNdefRecord_ObjectNameChanged
func callbackQQmlNdefRecord_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtNfc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQQmlNdefRecord_TimerEvent
func callbackQQmlNdefRecord_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlNdefRecordFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlNdefRecord) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlNdefRecord_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func init() {
	qt.ItfMap["nfc.QNdefFilter_ITF"] = QNdefFilter{}
	qt.FuncMap["nfc.NewQNdefFilter"] = NewQNdefFilter
	qt.FuncMap["nfc.NewQNdefFilter2"] = NewQNdefFilter2
	qt.ItfMap["nfc.QNdefMessage_ITF"] = QNdefMessage{}
	qt.FuncMap["nfc.NewQNdefMessage"] = NewQNdefMessage
	qt.FuncMap["nfc.NewQNdefMessage2"] = NewQNdefMessage2
	qt.FuncMap["nfc.NewQNdefMessage3"] = NewQNdefMessage3
	qt.FuncMap["nfc.NewQNdefMessage4"] = NewQNdefMessage4
	qt.FuncMap["nfc.QNdefMessage_FromByteArray"] = QNdefMessage_FromByteArray
	qt.ItfMap["nfc.QNdefNfcIconRecord_ITF"] = QNdefNfcIconRecord{}
	qt.ItfMap["nfc.QNdefNfcSmartPosterRecord_ITF"] = QNdefNfcSmartPosterRecord{}
	qt.FuncMap["nfc.NewQNdefNfcSmartPosterRecord"] = NewQNdefNfcSmartPosterRecord
	qt.FuncMap["nfc.NewQNdefNfcSmartPosterRecord2"] = NewQNdefNfcSmartPosterRecord2
	qt.FuncMap["nfc.NewQNdefNfcSmartPosterRecord3"] = NewQNdefNfcSmartPosterRecord3
	qt.EnumMap["nfc.QNdefNfcSmartPosterRecord__UnspecifiedAction"] = int64(QNdefNfcSmartPosterRecord__UnspecifiedAction)
	qt.EnumMap["nfc.QNdefNfcSmartPosterRecord__DoAction"] = int64(QNdefNfcSmartPosterRecord__DoAction)
	qt.EnumMap["nfc.QNdefNfcSmartPosterRecord__SaveAction"] = int64(QNdefNfcSmartPosterRecord__SaveAction)
	qt.EnumMap["nfc.QNdefNfcSmartPosterRecord__EditAction"] = int64(QNdefNfcSmartPosterRecord__EditAction)
	qt.ItfMap["nfc.QNdefNfcTextRecord_ITF"] = QNdefNfcTextRecord{}
	qt.FuncMap["nfc.NewQNdefNfcTextRecord"] = NewQNdefNfcTextRecord
	qt.FuncMap["nfc.NewQNdefNfcTextRecord2"] = NewQNdefNfcTextRecord2
	qt.EnumMap["nfc.QNdefNfcTextRecord__Utf8"] = int64(QNdefNfcTextRecord__Utf8)
	qt.EnumMap["nfc.QNdefNfcTextRecord__Utf16"] = int64(QNdefNfcTextRecord__Utf16)
	qt.ItfMap["nfc.QNdefNfcUriRecord_ITF"] = QNdefNfcUriRecord{}
	qt.FuncMap["nfc.NewQNdefNfcUriRecord"] = NewQNdefNfcUriRecord
	qt.FuncMap["nfc.NewQNdefNfcUriRecord2"] = NewQNdefNfcUriRecord2
	qt.ItfMap["nfc.QNdefRecord_ITF"] = QNdefRecord{}
	qt.FuncMap["nfc.NewQNdefRecord"] = NewQNdefRecord
	qt.FuncMap["nfc.NewQNdefRecord2"] = NewQNdefRecord2
	qt.EnumMap["nfc.QNdefRecord__Empty"] = int64(QNdefRecord__Empty)
	qt.EnumMap["nfc.QNdefRecord__NfcRtd"] = int64(QNdefRecord__NfcRtd)
	qt.EnumMap["nfc.QNdefRecord__Mime"] = int64(QNdefRecord__Mime)
	qt.EnumMap["nfc.QNdefRecord__Uri"] = int64(QNdefRecord__Uri)
	qt.EnumMap["nfc.QNdefRecord__ExternalRtd"] = int64(QNdefRecord__ExternalRtd)
	qt.EnumMap["nfc.QNdefRecord__Unknown"] = int64(QNdefRecord__Unknown)
	qt.ItfMap["nfc.QNearFieldManager_ITF"] = QNearFieldManager{}
	qt.FuncMap["nfc.NewQNearFieldManager"] = NewQNearFieldManager
	qt.EnumMap["nfc.QNearFieldManager__Offline"] = int64(QNearFieldManager__Offline)
	qt.EnumMap["nfc.QNearFieldManager__TurningOn"] = int64(QNearFieldManager__TurningOn)
	qt.EnumMap["nfc.QNearFieldManager__Online"] = int64(QNearFieldManager__Online)
	qt.EnumMap["nfc.QNearFieldManager__TurningOff"] = int64(QNearFieldManager__TurningOff)
	qt.EnumMap["nfc.QNearFieldManager__NoTargetAccess"] = int64(QNearFieldManager__NoTargetAccess)
	qt.EnumMap["nfc.QNearFieldManager__NdefReadTargetAccess"] = int64(QNearFieldManager__NdefReadTargetAccess)
	qt.EnumMap["nfc.QNearFieldManager__NdefWriteTargetAccess"] = int64(QNearFieldManager__NdefWriteTargetAccess)
	qt.EnumMap["nfc.QNearFieldManager__TagTypeSpecificTargetAccess"] = int64(QNearFieldManager__TagTypeSpecificTargetAccess)
	qt.ItfMap["nfc.QNearFieldShareManager_ITF"] = QNearFieldShareManager{}
	qt.FuncMap["nfc.NewQNearFieldShareManager"] = NewQNearFieldShareManager
	qt.FuncMap["nfc.QNearFieldShareManager_SupportedShareModes"] = QNearFieldShareManager_SupportedShareModes
	qt.EnumMap["nfc.QNearFieldShareManager__NoShare"] = int64(QNearFieldShareManager__NoShare)
	qt.EnumMap["nfc.QNearFieldShareManager__NdefShare"] = int64(QNearFieldShareManager__NdefShare)
	qt.EnumMap["nfc.QNearFieldShareManager__FileShare"] = int64(QNearFieldShareManager__FileShare)
	qt.EnumMap["nfc.QNearFieldShareManager__NoError"] = int64(QNearFieldShareManager__NoError)
	qt.EnumMap["nfc.QNearFieldShareManager__UnknownError"] = int64(QNearFieldShareManager__UnknownError)
	qt.EnumMap["nfc.QNearFieldShareManager__InvalidShareContentError"] = int64(QNearFieldShareManager__InvalidShareContentError)
	qt.EnumMap["nfc.QNearFieldShareManager__ShareCanceledError"] = int64(QNearFieldShareManager__ShareCanceledError)
	qt.EnumMap["nfc.QNearFieldShareManager__ShareInterruptedError"] = int64(QNearFieldShareManager__ShareInterruptedError)
	qt.EnumMap["nfc.QNearFieldShareManager__ShareRejectedError"] = int64(QNearFieldShareManager__ShareRejectedError)
	qt.EnumMap["nfc.QNearFieldShareManager__UnsupportedShareModeError"] = int64(QNearFieldShareManager__UnsupportedShareModeError)
	qt.EnumMap["nfc.QNearFieldShareManager__ShareAlreadyInProgressError"] = int64(QNearFieldShareManager__ShareAlreadyInProgressError)
	qt.EnumMap["nfc.QNearFieldShareManager__SharePermissionDeniedError"] = int64(QNearFieldShareManager__SharePermissionDeniedError)
	qt.ItfMap["nfc.QNearFieldShareTarget_ITF"] = QNearFieldShareTarget{}
	qt.ItfMap["nfc.QNearFieldTarget_ITF"] = QNearFieldTarget{}
	qt.FuncMap["nfc.NewQNearFieldTarget"] = NewQNearFieldTarget
	qt.EnumMap["nfc.QNearFieldTarget__ProprietaryTag"] = int64(QNearFieldTarget__ProprietaryTag)
	qt.EnumMap["nfc.QNearFieldTarget__NfcTagType1"] = int64(QNearFieldTarget__NfcTagType1)
	qt.EnumMap["nfc.QNearFieldTarget__NfcTagType2"] = int64(QNearFieldTarget__NfcTagType2)
	qt.EnumMap["nfc.QNearFieldTarget__NfcTagType3"] = int64(QNearFieldTarget__NfcTagType3)
	qt.EnumMap["nfc.QNearFieldTarget__NfcTagType4"] = int64(QNearFieldTarget__NfcTagType4)
	qt.EnumMap["nfc.QNearFieldTarget__MifareTag"] = int64(QNearFieldTarget__MifareTag)
	qt.EnumMap["nfc.QNearFieldTarget__UnknownAccess"] = int64(QNearFieldTarget__UnknownAccess)
	qt.EnumMap["nfc.QNearFieldTarget__NdefAccess"] = int64(QNearFieldTarget__NdefAccess)
	qt.EnumMap["nfc.QNearFieldTarget__TagTypeSpecificAccess"] = int64(QNearFieldTarget__TagTypeSpecificAccess)
	qt.EnumMap["nfc.QNearFieldTarget__LlcpAccess"] = int64(QNearFieldTarget__LlcpAccess)
	qt.EnumMap["nfc.QNearFieldTarget__NoError"] = int64(QNearFieldTarget__NoError)
	qt.EnumMap["nfc.QNearFieldTarget__UnknownError"] = int64(QNearFieldTarget__UnknownError)
	qt.EnumMap["nfc.QNearFieldTarget__UnsupportedError"] = int64(QNearFieldTarget__UnsupportedError)
	qt.EnumMap["nfc.QNearFieldTarget__TargetOutOfRangeError"] = int64(QNearFieldTarget__TargetOutOfRangeError)
	qt.EnumMap["nfc.QNearFieldTarget__NoResponseError"] = int64(QNearFieldTarget__NoResponseError)
	qt.EnumMap["nfc.QNearFieldTarget__ChecksumMismatchError"] = int64(QNearFieldTarget__ChecksumMismatchError)
	qt.EnumMap["nfc.QNearFieldTarget__InvalidParametersError"] = int64(QNearFieldTarget__InvalidParametersError)
	qt.EnumMap["nfc.QNearFieldTarget__NdefReadError"] = int64(QNearFieldTarget__NdefReadError)
	qt.EnumMap["nfc.QNearFieldTarget__NdefWriteError"] = int64(QNearFieldTarget__NdefWriteError)
	qt.EnumMap["nfc.QNearFieldTarget__CommandError"] = int64(QNearFieldTarget__CommandError)
	qt.ItfMap["nfc.QQmlNdefRecord_ITF"] = QQmlNdefRecord{}
	qt.FuncMap["nfc.NewQQmlNdefRecord"] = NewQQmlNdefRecord
	qt.FuncMap["nfc.NewQQmlNdefRecord2"] = NewQQmlNdefRecord2
	qt.EnumMap["nfc.QQmlNdefRecord__Empty"] = int64(QQmlNdefRecord__Empty)
	qt.EnumMap["nfc.QQmlNdefRecord__NfcRtd"] = int64(QQmlNdefRecord__NfcRtd)
	qt.EnumMap["nfc.QQmlNdefRecord__Mime"] = int64(QQmlNdefRecord__Mime)
	qt.EnumMap["nfc.QQmlNdefRecord__Uri"] = int64(QQmlNdefRecord__Uri)
	qt.EnumMap["nfc.QQmlNdefRecord__ExternalRtd"] = int64(QQmlNdefRecord__ExternalRtd)
	qt.EnumMap["nfc.QQmlNdefRecord__Unknown"] = int64(QQmlNdefRecord__Unknown)
}
