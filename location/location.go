// +build !minimal

package location

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "location.h"
import "C"
import (
	"github.com/StarAurryon/qt"
	"github.com/StarAurryon/qt/core"
	"github.com/StarAurryon/qt/positioning"
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtLocation_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtLocation_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return []byte(gs)
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type QGeoCodeReply struct {
	core.QObject
}

type QGeoCodeReply_ITF interface {
	core.QObject_ITF
	QGeoCodeReply_PTR() *QGeoCodeReply
}

func (ptr *QGeoCodeReply) QGeoCodeReply_PTR() *QGeoCodeReply {
	return ptr
}

func (ptr *QGeoCodeReply) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QGeoCodeReply) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQGeoCodeReply(ptr QGeoCodeReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoCodeReply_PTR().Pointer()
	}
	return nil
}

func NewQGeoCodeReplyFromPointer(ptr unsafe.Pointer) (n *QGeoCodeReply) {
	n = new(QGeoCodeReply)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QGeoCodeReply__Error
//QGeoCodeReply::Error
type QGeoCodeReply__Error int64

const (
	QGeoCodeReply__NoError                QGeoCodeReply__Error = QGeoCodeReply__Error(0)
	QGeoCodeReply__EngineNotSetError      QGeoCodeReply__Error = QGeoCodeReply__Error(1)
	QGeoCodeReply__CommunicationError     QGeoCodeReply__Error = QGeoCodeReply__Error(2)
	QGeoCodeReply__ParseError             QGeoCodeReply__Error = QGeoCodeReply__Error(3)
	QGeoCodeReply__UnsupportedOptionError QGeoCodeReply__Error = QGeoCodeReply__Error(4)
	QGeoCodeReply__CombinationError       QGeoCodeReply__Error = QGeoCodeReply__Error(5)
	QGeoCodeReply__UnknownError           QGeoCodeReply__Error = QGeoCodeReply__Error(6)
)

type QGeoCodingManager struct {
	core.QObject
}

type QGeoCodingManager_ITF interface {
	core.QObject_ITF
	QGeoCodingManager_PTR() *QGeoCodingManager
}

func (ptr *QGeoCodingManager) QGeoCodingManager_PTR() *QGeoCodingManager {
	return ptr
}

func (ptr *QGeoCodingManager) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QGeoCodingManager) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQGeoCodingManager(ptr QGeoCodingManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoCodingManager_PTR().Pointer()
	}
	return nil
}

func NewQGeoCodingManagerFromPointer(ptr unsafe.Pointer) (n *QGeoCodingManager) {
	n = new(QGeoCodingManager)
	n.SetPointer(ptr)
	return
}

type QGeoCodingManagerEngine struct {
	core.QObject
}

type QGeoCodingManagerEngine_ITF interface {
	core.QObject_ITF
	QGeoCodingManagerEngine_PTR() *QGeoCodingManagerEngine
}

func (ptr *QGeoCodingManagerEngine) QGeoCodingManagerEngine_PTR() *QGeoCodingManagerEngine {
	return ptr
}

func (ptr *QGeoCodingManagerEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QGeoCodingManagerEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQGeoCodingManagerEngine(ptr QGeoCodingManagerEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoCodingManagerEngine_PTR().Pointer()
	}
	return nil
}

func NewQGeoCodingManagerEngineFromPointer(ptr unsafe.Pointer) (n *QGeoCodingManagerEngine) {
	n = new(QGeoCodingManagerEngine)
	n.SetPointer(ptr)
	return
}

type QGeoJson struct {
	ptr unsafe.Pointer
}

type QGeoJson_ITF interface {
	QGeoJson_PTR() *QGeoJson
}

func (ptr *QGeoJson) QGeoJson_PTR() *QGeoJson {
	return ptr
}

func (ptr *QGeoJson) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QGeoJson) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQGeoJson(ptr QGeoJson_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoJson_PTR().Pointer()
	}
	return nil
}

func NewQGeoJsonFromPointer(ptr unsafe.Pointer) (n *QGeoJson) {
	n = new(QGeoJson)
	n.SetPointer(ptr)
	return
}
func (ptr *QGeoJson) DestroyQGeoJson() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QGeoManeuver struct {
	ptr unsafe.Pointer
}

type QGeoManeuver_ITF interface {
	QGeoManeuver_PTR() *QGeoManeuver
}

func (ptr *QGeoManeuver) QGeoManeuver_PTR() *QGeoManeuver {
	return ptr
}

func (ptr *QGeoManeuver) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QGeoManeuver) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQGeoManeuver(ptr QGeoManeuver_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoManeuver_PTR().Pointer()
	}
	return nil
}

func NewQGeoManeuverFromPointer(ptr unsafe.Pointer) (n *QGeoManeuver) {
	n = new(QGeoManeuver)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QGeoManeuver__InstructionDirection
//QGeoManeuver::InstructionDirection
type QGeoManeuver__InstructionDirection int64

const (
	QGeoManeuver__NoDirection         QGeoManeuver__InstructionDirection = QGeoManeuver__InstructionDirection(0)
	QGeoManeuver__DirectionForward    QGeoManeuver__InstructionDirection = QGeoManeuver__InstructionDirection(1)
	QGeoManeuver__DirectionBearRight  QGeoManeuver__InstructionDirection = QGeoManeuver__InstructionDirection(2)
	QGeoManeuver__DirectionLightRight QGeoManeuver__InstructionDirection = QGeoManeuver__InstructionDirection(3)
	QGeoManeuver__DirectionRight      QGeoManeuver__InstructionDirection = QGeoManeuver__InstructionDirection(4)
	QGeoManeuver__DirectionHardRight  QGeoManeuver__InstructionDirection = QGeoManeuver__InstructionDirection(5)
	QGeoManeuver__DirectionUTurnRight QGeoManeuver__InstructionDirection = QGeoManeuver__InstructionDirection(6)
	QGeoManeuver__DirectionUTurnLeft  QGeoManeuver__InstructionDirection = QGeoManeuver__InstructionDirection(7)
	QGeoManeuver__DirectionHardLeft   QGeoManeuver__InstructionDirection = QGeoManeuver__InstructionDirection(8)
	QGeoManeuver__DirectionLeft       QGeoManeuver__InstructionDirection = QGeoManeuver__InstructionDirection(9)
	QGeoManeuver__DirectionLightLeft  QGeoManeuver__InstructionDirection = QGeoManeuver__InstructionDirection(10)
	QGeoManeuver__DirectionBearLeft   QGeoManeuver__InstructionDirection = QGeoManeuver__InstructionDirection(11)
)

func NewQGeoManeuver() *QGeoManeuver {
	tmpValue := NewQGeoManeuverFromPointer(C.QGeoManeuver_NewQGeoManeuver())
	qt.SetFinalizer(tmpValue, (*QGeoManeuver).DestroyQGeoManeuver)
	return tmpValue
}

func NewQGeoManeuver2(other QGeoManeuver_ITF) *QGeoManeuver {
	tmpValue := NewQGeoManeuverFromPointer(C.QGeoManeuver_NewQGeoManeuver2(PointerFromQGeoManeuver(other)))
	qt.SetFinalizer(tmpValue, (*QGeoManeuver).DestroyQGeoManeuver)
	return tmpValue
}

func (ptr *QGeoManeuver) Direction() QGeoManeuver__InstructionDirection {
	if ptr.Pointer() != nil {
		return QGeoManeuver__InstructionDirection(C.QGeoManeuver_Direction(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoManeuver) DistanceToNextInstruction() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoManeuver_DistanceToNextInstruction(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoManeuver) ExtendedAttributes() map[string]*core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) map[string]*core.QVariant {
			out := make(map[string]*core.QVariant, int(l.len))
			tmpList := NewQGeoManeuverFromPointer(l.data)
			for i, v := range tmpList.__extendedAttributes_keyList() {
				out[v] = tmpList.__extendedAttributes_atList(v, i)
			}
			return out
		}(C.QGeoManeuver_ExtendedAttributes(ptr.Pointer()))
	}
	return make(map[string]*core.QVariant, 0)
}

func (ptr *QGeoManeuver) InstructionText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoManeuver_InstructionText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoManeuver) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoManeuver_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGeoManeuver) Position() *positioning.QGeoCoordinate {
	if ptr.Pointer() != nil {
		tmpValue := positioning.NewQGeoCoordinateFromPointer(C.QGeoManeuver_Position(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*positioning.QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoManeuver) SetDirection(direction QGeoManeuver__InstructionDirection) {
	if ptr.Pointer() != nil {
		C.QGeoManeuver_SetDirection(ptr.Pointer(), C.longlong(direction))
	}
}

func (ptr *QGeoManeuver) SetDistanceToNextInstruction(distance float64) {
	if ptr.Pointer() != nil {
		C.QGeoManeuver_SetDistanceToNextInstruction(ptr.Pointer(), C.double(distance))
	}
}

func (ptr *QGeoManeuver) SetExtendedAttributes(extendedAttributes map[string]*core.QVariant) {
	if ptr.Pointer() != nil {
		C.QGeoManeuver_SetExtendedAttributes(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQGeoManeuverFromPointer(NewQGeoManeuverFromPointer(nil).__setExtendedAttributes_extendedAttributes_newList())
			for k, v := range extendedAttributes {
				tmpList.__setExtendedAttributes_extendedAttributes_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QGeoManeuver) SetInstructionText(instructionText string) {
	if ptr.Pointer() != nil {
		var instructionTextC *C.char
		if instructionText != "" {
			instructionTextC = C.CString(instructionText)
			defer C.free(unsafe.Pointer(instructionTextC))
		}
		C.QGeoManeuver_SetInstructionText(ptr.Pointer(), C.struct_QtLocation_PackedString{data: instructionTextC, len: C.longlong(len(instructionText))})
	}
}

func (ptr *QGeoManeuver) SetPosition(position positioning.QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoManeuver_SetPosition(ptr.Pointer(), positioning.PointerFromQGeoCoordinate(position))
	}
}

func (ptr *QGeoManeuver) SetTimeToNextInstruction(secs int) {
	if ptr.Pointer() != nil {
		C.QGeoManeuver_SetTimeToNextInstruction(ptr.Pointer(), C.int(int32(secs)))
	}
}

func (ptr *QGeoManeuver) SetWaypoint(coordinate positioning.QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoManeuver_SetWaypoint(ptr.Pointer(), positioning.PointerFromQGeoCoordinate(coordinate))
	}
}

func (ptr *QGeoManeuver) TimeToNextInstruction() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGeoManeuver_TimeToNextInstruction(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoManeuver) Waypoint() *positioning.QGeoCoordinate {
	if ptr.Pointer() != nil {
		tmpValue := positioning.NewQGeoCoordinateFromPointer(C.QGeoManeuver_Waypoint(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*positioning.QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoManeuver) DestroyQGeoManeuver() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoManeuver_DestroyQGeoManeuver(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoManeuver) __extendedAttributes_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QGeoManeuver___extendedAttributes_atList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoManeuver) __extendedAttributes_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QGeoManeuver___extendedAttributes_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QGeoManeuver) __extendedAttributes_newList() unsafe.Pointer {
	return C.QGeoManeuver___extendedAttributes_newList(ptr.Pointer())
}

func (ptr *QGeoManeuver) __extendedAttributes_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQGeoManeuverFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____extendedAttributes_keyList_atList(i)
			}
			return out
		}(C.QGeoManeuver___extendedAttributes_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QGeoManeuver) __setExtendedAttributes_extendedAttributes_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QGeoManeuver___setExtendedAttributes_extendedAttributes_atList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoManeuver) __setExtendedAttributes_extendedAttributes_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QGeoManeuver___setExtendedAttributes_extendedAttributes_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QGeoManeuver) __setExtendedAttributes_extendedAttributes_newList() unsafe.Pointer {
	return C.QGeoManeuver___setExtendedAttributes_extendedAttributes_newList(ptr.Pointer())
}

func (ptr *QGeoManeuver) __setExtendedAttributes_extendedAttributes_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQGeoManeuverFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setExtendedAttributes_extendedAttributes_keyList_atList(i)
			}
			return out
		}(C.QGeoManeuver___setExtendedAttributes_extendedAttributes_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QGeoManeuver) ____extendedAttributes_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoManeuver_____extendedAttributes_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QGeoManeuver) ____extendedAttributes_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QGeoManeuver_____extendedAttributes_keyList_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QGeoManeuver) ____extendedAttributes_keyList_newList() unsafe.Pointer {
	return C.QGeoManeuver_____extendedAttributes_keyList_newList(ptr.Pointer())
}

func (ptr *QGeoManeuver) ____setExtendedAttributes_extendedAttributes_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoManeuver_____setExtendedAttributes_extendedAttributes_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QGeoManeuver) ____setExtendedAttributes_extendedAttributes_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QGeoManeuver_____setExtendedAttributes_extendedAttributes_keyList_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QGeoManeuver) ____setExtendedAttributes_extendedAttributes_keyList_newList() unsafe.Pointer {
	return C.QGeoManeuver_____setExtendedAttributes_extendedAttributes_keyList_newList(ptr.Pointer())
}

type QGeoRoute struct {
	ptr unsafe.Pointer
}

type QGeoRoute_ITF interface {
	QGeoRoute_PTR() *QGeoRoute
}

func (ptr *QGeoRoute) QGeoRoute_PTR() *QGeoRoute {
	return ptr
}

func (ptr *QGeoRoute) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QGeoRoute) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQGeoRoute(ptr QGeoRoute_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoRoute_PTR().Pointer()
	}
	return nil
}

func NewQGeoRouteFromPointer(ptr unsafe.Pointer) (n *QGeoRoute) {
	n = new(QGeoRoute)
	n.SetPointer(ptr)
	return
}
func NewQGeoRoute() *QGeoRoute {
	tmpValue := NewQGeoRouteFromPointer(C.QGeoRoute_NewQGeoRoute())
	qt.SetFinalizer(tmpValue, (*QGeoRoute).DestroyQGeoRoute)
	return tmpValue
}

func NewQGeoRoute2(other QGeoRoute_ITF) *QGeoRoute {
	tmpValue := NewQGeoRouteFromPointer(C.QGeoRoute_NewQGeoRoute2(PointerFromQGeoRoute(other)))
	qt.SetFinalizer(tmpValue, (*QGeoRoute).DestroyQGeoRoute)
	return tmpValue
}

func (ptr *QGeoRoute) Bounds() *positioning.QGeoRectangle {
	if ptr.Pointer() != nil {
		tmpValue := positioning.NewQGeoRectangleFromPointer(C.QGeoRoute_Bounds(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*positioning.QGeoRectangle).DestroyQGeoRectangle)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRoute) Distance() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoRoute_Distance(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoute) ExtendedAttributes() map[string]*core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) map[string]*core.QVariant {
			out := make(map[string]*core.QVariant, int(l.len))
			tmpList := NewQGeoRouteFromPointer(l.data)
			for i, v := range tmpList.__extendedAttributes_keyList() {
				out[v] = tmpList.__extendedAttributes_atList(v, i)
			}
			return out
		}(C.QGeoRoute_ExtendedAttributes(ptr.Pointer()))
	}
	return make(map[string]*core.QVariant, 0)
}

func (ptr *QGeoRoute) FirstRouteSegment() *QGeoRouteSegment {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoRouteSegmentFromPointer(C.QGeoRoute_FirstRouteSegment(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QGeoRouteSegment).DestroyQGeoRouteSegment)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRoute) Path() []*positioning.QGeoCoordinate {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) []*positioning.QGeoCoordinate {
			out := make([]*positioning.QGeoCoordinate, int(l.len))
			tmpList := NewQGeoRouteFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__path_atList(i)
			}
			return out
		}(C.QGeoRoute_Path(ptr.Pointer()))
	}
	return make([]*positioning.QGeoCoordinate, 0)
}

func (ptr *QGeoRoute) Request() *QGeoRouteRequest {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoRouteRequestFromPointer(C.QGeoRoute_Request(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QGeoRouteRequest).DestroyQGeoRouteRequest)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRoute) RouteId() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoRoute_RouteId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoRoute) RouteLegs() []*QGeoRouteLeg {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) []*QGeoRouteLeg {
			out := make([]*QGeoRouteLeg, int(l.len))
			tmpList := NewQGeoRouteFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__routeLegs_atList(i)
			}
			return out
		}(C.QGeoRoute_RouteLegs(ptr.Pointer()))
	}
	return make([]*QGeoRouteLeg, 0)
}

func (ptr *QGeoRoute) SetBounds(bounds positioning.QGeoRectangle_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoute_SetBounds(ptr.Pointer(), positioning.PointerFromQGeoRectangle(bounds))
	}
}

func (ptr *QGeoRoute) SetDistance(distance float64) {
	if ptr.Pointer() != nil {
		C.QGeoRoute_SetDistance(ptr.Pointer(), C.double(distance))
	}
}

func (ptr *QGeoRoute) SetExtendedAttributes(extendedAttributes map[string]*core.QVariant) {
	if ptr.Pointer() != nil {
		C.QGeoRoute_SetExtendedAttributes(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQGeoRouteFromPointer(NewQGeoRouteFromPointer(nil).__setExtendedAttributes_extendedAttributes_newList())
			for k, v := range extendedAttributes {
				tmpList.__setExtendedAttributes_extendedAttributes_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QGeoRoute) SetFirstRouteSegment(routeSegment QGeoRouteSegment_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoute_SetFirstRouteSegment(ptr.Pointer(), PointerFromQGeoRouteSegment(routeSegment))
	}
}

func (ptr *QGeoRoute) SetPath(path []*positioning.QGeoCoordinate) {
	if ptr.Pointer() != nil {
		C.QGeoRoute_SetPath(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQGeoRouteFromPointer(NewQGeoRouteFromPointer(nil).__setPath_path_newList())
			for _, v := range path {
				tmpList.__setPath_path_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QGeoRoute) SetRequest(request QGeoRouteRequest_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoute_SetRequest(ptr.Pointer(), PointerFromQGeoRouteRequest(request))
	}
}

func (ptr *QGeoRoute) SetRouteId(id string) {
	if ptr.Pointer() != nil {
		var idC *C.char
		if id != "" {
			idC = C.CString(id)
			defer C.free(unsafe.Pointer(idC))
		}
		C.QGeoRoute_SetRouteId(ptr.Pointer(), C.struct_QtLocation_PackedString{data: idC, len: C.longlong(len(id))})
	}
}

func (ptr *QGeoRoute) SetRouteLegs(legs []*QGeoRouteLeg) {
	if ptr.Pointer() != nil {
		C.QGeoRoute_SetRouteLegs(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQGeoRouteFromPointer(NewQGeoRouteFromPointer(nil).__setRouteLegs_legs_newList())
			for _, v := range legs {
				tmpList.__setRouteLegs_legs_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QGeoRoute) SetTravelMode(mode QGeoRouteRequest__TravelMode) {
	if ptr.Pointer() != nil {
		C.QGeoRoute_SetTravelMode(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QGeoRoute) SetTravelTime(secs int) {
	if ptr.Pointer() != nil {
		C.QGeoRoute_SetTravelTime(ptr.Pointer(), C.int(int32(secs)))
	}
}

func (ptr *QGeoRoute) TravelMode() QGeoRouteRequest__TravelMode {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__TravelMode(C.QGeoRoute_TravelMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoute) TravelTime() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGeoRoute_TravelTime(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoRoute) DestroyQGeoRoute() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoRoute_DestroyQGeoRoute(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoRoute) __extendedAttributes_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QGeoRoute___extendedAttributes_atList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRoute) __extendedAttributes_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QGeoRoute___extendedAttributes_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QGeoRoute) __extendedAttributes_newList() unsafe.Pointer {
	return C.QGeoRoute___extendedAttributes_newList(ptr.Pointer())
}

func (ptr *QGeoRoute) __extendedAttributes_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQGeoRouteFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____extendedAttributes_keyList_atList(i)
			}
			return out
		}(C.QGeoRoute___extendedAttributes_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QGeoRoute) __path_atList(i int) *positioning.QGeoCoordinate {
	if ptr.Pointer() != nil {
		tmpValue := positioning.NewQGeoCoordinateFromPointer(C.QGeoRoute___path_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*positioning.QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRoute) __path_setList(i positioning.QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoute___path_setList(ptr.Pointer(), positioning.PointerFromQGeoCoordinate(i))
	}
}

func (ptr *QGeoRoute) __path_newList() unsafe.Pointer {
	return C.QGeoRoute___path_newList(ptr.Pointer())
}

func (ptr *QGeoRoute) __routeLegs_atList(i int) *QGeoRouteLeg {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoRouteLegFromPointer(C.QGeoRoute___routeLegs_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QGeoRouteLeg).DestroyQGeoRouteLeg)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRoute) __routeLegs_setList(i QGeoRouteLeg_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoute___routeLegs_setList(ptr.Pointer(), PointerFromQGeoRouteLeg(i))
	}
}

func (ptr *QGeoRoute) __routeLegs_newList() unsafe.Pointer {
	return C.QGeoRoute___routeLegs_newList(ptr.Pointer())
}

func (ptr *QGeoRoute) __setExtendedAttributes_extendedAttributes_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QGeoRoute___setExtendedAttributes_extendedAttributes_atList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRoute) __setExtendedAttributes_extendedAttributes_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QGeoRoute___setExtendedAttributes_extendedAttributes_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QGeoRoute) __setExtendedAttributes_extendedAttributes_newList() unsafe.Pointer {
	return C.QGeoRoute___setExtendedAttributes_extendedAttributes_newList(ptr.Pointer())
}

func (ptr *QGeoRoute) __setExtendedAttributes_extendedAttributes_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQGeoRouteFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setExtendedAttributes_extendedAttributes_keyList_atList(i)
			}
			return out
		}(C.QGeoRoute___setExtendedAttributes_extendedAttributes_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QGeoRoute) __setPath_path_atList(i int) *positioning.QGeoCoordinate {
	if ptr.Pointer() != nil {
		tmpValue := positioning.NewQGeoCoordinateFromPointer(C.QGeoRoute___setPath_path_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*positioning.QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRoute) __setPath_path_setList(i positioning.QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoute___setPath_path_setList(ptr.Pointer(), positioning.PointerFromQGeoCoordinate(i))
	}
}

func (ptr *QGeoRoute) __setPath_path_newList() unsafe.Pointer {
	return C.QGeoRoute___setPath_path_newList(ptr.Pointer())
}

func (ptr *QGeoRoute) __setRouteLegs_legs_atList(i int) *QGeoRouteLeg {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoRouteLegFromPointer(C.QGeoRoute___setRouteLegs_legs_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QGeoRouteLeg).DestroyQGeoRouteLeg)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRoute) __setRouteLegs_legs_setList(i QGeoRouteLeg_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoute___setRouteLegs_legs_setList(ptr.Pointer(), PointerFromQGeoRouteLeg(i))
	}
}

func (ptr *QGeoRoute) __setRouteLegs_legs_newList() unsafe.Pointer {
	return C.QGeoRoute___setRouteLegs_legs_newList(ptr.Pointer())
}

func (ptr *QGeoRoute) ____extendedAttributes_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoRoute_____extendedAttributes_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QGeoRoute) ____extendedAttributes_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QGeoRoute_____extendedAttributes_keyList_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QGeoRoute) ____extendedAttributes_keyList_newList() unsafe.Pointer {
	return C.QGeoRoute_____extendedAttributes_keyList_newList(ptr.Pointer())
}

func (ptr *QGeoRoute) ____setExtendedAttributes_extendedAttributes_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoRoute_____setExtendedAttributes_extendedAttributes_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QGeoRoute) ____setExtendedAttributes_extendedAttributes_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QGeoRoute_____setExtendedAttributes_extendedAttributes_keyList_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QGeoRoute) ____setExtendedAttributes_extendedAttributes_keyList_newList() unsafe.Pointer {
	return C.QGeoRoute_____setExtendedAttributes_extendedAttributes_keyList_newList(ptr.Pointer())
}

type QGeoRouteLeg struct {
	QGeoRoute
}

type QGeoRouteLeg_ITF interface {
	QGeoRoute_ITF
	QGeoRouteLeg_PTR() *QGeoRouteLeg
}

func (ptr *QGeoRouteLeg) QGeoRouteLeg_PTR() *QGeoRouteLeg {
	return ptr
}

func (ptr *QGeoRouteLeg) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoRoute_PTR().Pointer()
	}
	return nil
}

func (ptr *QGeoRouteLeg) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QGeoRoute_PTR().SetPointer(p)
	}
}

func PointerFromQGeoRouteLeg(ptr QGeoRouteLeg_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoRouteLeg_PTR().Pointer()
	}
	return nil
}

func NewQGeoRouteLegFromPointer(ptr unsafe.Pointer) (n *QGeoRouteLeg) {
	n = new(QGeoRouteLeg)
	n.SetPointer(ptr)
	return
}
func NewQGeoRouteLeg() *QGeoRouteLeg {
	tmpValue := NewQGeoRouteLegFromPointer(C.QGeoRouteLeg_NewQGeoRouteLeg())
	qt.SetFinalizer(tmpValue, (*QGeoRouteLeg).DestroyQGeoRouteLeg)
	return tmpValue
}

func NewQGeoRouteLeg2(other QGeoRouteLeg_ITF) *QGeoRouteLeg {
	tmpValue := NewQGeoRouteLegFromPointer(C.QGeoRouteLeg_NewQGeoRouteLeg2(PointerFromQGeoRouteLeg(other)))
	qt.SetFinalizer(tmpValue, (*QGeoRouteLeg).DestroyQGeoRouteLeg)
	return tmpValue
}

func (ptr *QGeoRouteLeg) LegIndex() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGeoRouteLeg_LegIndex(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoRouteLeg) OverallRoute() *QGeoRoute {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoRouteFromPointer(C.QGeoRouteLeg_OverallRoute(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QGeoRoute).DestroyQGeoRoute)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRouteLeg) SetLegIndex(idx int) {
	if ptr.Pointer() != nil {
		C.QGeoRouteLeg_SetLegIndex(ptr.Pointer(), C.int(int32(idx)))
	}
}

func (ptr *QGeoRouteLeg) SetOverallRoute(route QGeoRoute_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRouteLeg_SetOverallRoute(ptr.Pointer(), PointerFromQGeoRoute(route))
	}
}

func (ptr *QGeoRouteLeg) DestroyQGeoRouteLeg() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoRouteLeg_DestroyQGeoRouteLeg(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QGeoRouteReply struct {
	core.QObject
}

type QGeoRouteReply_ITF interface {
	core.QObject_ITF
	QGeoRouteReply_PTR() *QGeoRouteReply
}

func (ptr *QGeoRouteReply) QGeoRouteReply_PTR() *QGeoRouteReply {
	return ptr
}

func (ptr *QGeoRouteReply) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QGeoRouteReply) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQGeoRouteReply(ptr QGeoRouteReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoRouteReply_PTR().Pointer()
	}
	return nil
}

func NewQGeoRouteReplyFromPointer(ptr unsafe.Pointer) (n *QGeoRouteReply) {
	n = new(QGeoRouteReply)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QGeoRouteReply__Error
//QGeoRouteReply::Error
type QGeoRouteReply__Error int64

const (
	QGeoRouteReply__NoError                QGeoRouteReply__Error = QGeoRouteReply__Error(0)
	QGeoRouteReply__EngineNotSetError      QGeoRouteReply__Error = QGeoRouteReply__Error(1)
	QGeoRouteReply__CommunicationError     QGeoRouteReply__Error = QGeoRouteReply__Error(2)
	QGeoRouteReply__ParseError             QGeoRouteReply__Error = QGeoRouteReply__Error(3)
	QGeoRouteReply__UnsupportedOptionError QGeoRouteReply__Error = QGeoRouteReply__Error(4)
	QGeoRouteReply__UnknownError           QGeoRouteReply__Error = QGeoRouteReply__Error(5)
)

func NewQGeoRouteReply(error QGeoRouteReply__Error, errorString string, parent core.QObject_ITF) *QGeoRouteReply {
	var errorStringC *C.char
	if errorString != "" {
		errorStringC = C.CString(errorString)
		defer C.free(unsafe.Pointer(errorStringC))
	}
	tmpValue := NewQGeoRouteReplyFromPointer(C.QGeoRouteReply_NewQGeoRouteReply(C.longlong(error), C.struct_QtLocation_PackedString{data: errorStringC, len: C.longlong(len(errorString))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQGeoRouteReply2(request QGeoRouteRequest_ITF, parent core.QObject_ITF) *QGeoRouteReply {
	tmpValue := NewQGeoRouteReplyFromPointer(C.QGeoRouteReply_NewQGeoRouteReply2(PointerFromQGeoRouteRequest(request), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQGeoRouteReply_Abort
func callbackQGeoRouteReply_Abort(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "abort"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGeoRouteReplyFromPointer(ptr).AbortDefault()
	}
}

func (ptr *QGeoRouteReply) ConnectAbort(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "abort"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "abort", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "abort", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoRouteReply) DisconnectAbort() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "abort")
	}
}

func (ptr *QGeoRouteReply) Abort() {
	if ptr.Pointer() != nil {
		C.QGeoRouteReply_Abort(ptr.Pointer())
	}
}

func (ptr *QGeoRouteReply) AbortDefault() {
	if ptr.Pointer() != nil {
		C.QGeoRouteReply_AbortDefault(ptr.Pointer())
	}
}

//export callbackQGeoRouteReply_Aborted
func callbackQGeoRouteReply_Aborted(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "aborted"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QGeoRouteReply) ConnectAborted(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "aborted") {
			C.QGeoRouteReply_ConnectAborted(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "aborted")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "aborted"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "aborted", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "aborted", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoRouteReply) DisconnectAborted() {
	if ptr.Pointer() != nil {
		C.QGeoRouteReply_DisconnectAborted(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "aborted")
	}
}

func (ptr *QGeoRouteReply) Aborted() {
	if ptr.Pointer() != nil {
		C.QGeoRouteReply_Aborted(ptr.Pointer())
	}
}

func (ptr *QGeoRouteReply) AddRoutes(routes []*QGeoRoute) {
	if ptr.Pointer() != nil {
		C.QGeoRouteReply_AddRoutes(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQGeoRouteReplyFromPointer(NewQGeoRouteReplyFromPointer(nil).__addRoutes_routes_newList())
			for _, v := range routes {
				tmpList.__addRoutes_routes_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QGeoRouteReply) Error() QGeoRouteReply__Error {
	if ptr.Pointer() != nil {
		return QGeoRouteReply__Error(C.QGeoRouteReply_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQGeoRouteReply_Error2
func callbackQGeoRouteReply_Error2(ptr unsafe.Pointer, error C.longlong, errorString C.struct_QtLocation_PackedString) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		(*(*func(QGeoRouteReply__Error, string))(signal))(QGeoRouteReply__Error(error), cGoUnpackString(errorString))
	}

}

func (ptr *QGeoRouteReply) ConnectError2(f func(error QGeoRouteReply__Error, errorString string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error2") {
			C.QGeoRouteReply_ConnectError2(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "error")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error2"); signal != nil {
			f := func(error QGeoRouteReply__Error, errorString string) {
				(*(*func(QGeoRouteReply__Error, string))(signal))(error, errorString)
				f(error, errorString)
			}
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoRouteReply) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QGeoRouteReply_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "error2")
	}
}

func (ptr *QGeoRouteReply) Error2(error QGeoRouteReply__Error, errorString string) {
	if ptr.Pointer() != nil {
		var errorStringC *C.char
		if errorString != "" {
			errorStringC = C.CString(errorString)
			defer C.free(unsafe.Pointer(errorStringC))
		}
		C.QGeoRouteReply_Error2(ptr.Pointer(), C.longlong(error), C.struct_QtLocation_PackedString{data: errorStringC, len: C.longlong(len(errorString))})
	}
}

func (ptr *QGeoRouteReply) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoRouteReply_ErrorString(ptr.Pointer()))
	}
	return ""
}

//export callbackQGeoRouteReply_Finished
func callbackQGeoRouteReply_Finished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "finished"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QGeoRouteReply) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finished") {
			C.QGeoRouteReply_ConnectFinished(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "finished")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "finished"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "finished", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finished", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoRouteReply) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QGeoRouteReply_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "finished")
	}
}

func (ptr *QGeoRouteReply) Finished() {
	if ptr.Pointer() != nil {
		C.QGeoRouteReply_Finished(ptr.Pointer())
	}
}

func (ptr *QGeoRouteReply) IsFinished() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoRouteReply_IsFinished(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGeoRouteReply) Request() *QGeoRouteRequest {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoRouteRequestFromPointer(C.QGeoRouteReply_Request(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QGeoRouteRequest).DestroyQGeoRouteRequest)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRouteReply) Routes() []*QGeoRoute {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) []*QGeoRoute {
			out := make([]*QGeoRoute, int(l.len))
			tmpList := NewQGeoRouteReplyFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__routes_atList(i)
			}
			return out
		}(C.QGeoRouteReply_Routes(ptr.Pointer()))
	}
	return make([]*QGeoRoute, 0)
}

func (ptr *QGeoRouteReply) SetError(error QGeoRouteReply__Error, errorString string) {
	if ptr.Pointer() != nil {
		var errorStringC *C.char
		if errorString != "" {
			errorStringC = C.CString(errorString)
			defer C.free(unsafe.Pointer(errorStringC))
		}
		C.QGeoRouteReply_SetError(ptr.Pointer(), C.longlong(error), C.struct_QtLocation_PackedString{data: errorStringC, len: C.longlong(len(errorString))})
	}
}

func (ptr *QGeoRouteReply) SetFinished(finished bool) {
	if ptr.Pointer() != nil {
		C.QGeoRouteReply_SetFinished(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(finished))))
	}
}

func (ptr *QGeoRouteReply) SetRoutes(routes []*QGeoRoute) {
	if ptr.Pointer() != nil {
		C.QGeoRouteReply_SetRoutes(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQGeoRouteReplyFromPointer(NewQGeoRouteReplyFromPointer(nil).__setRoutes_routes_newList())
			for _, v := range routes {
				tmpList.__setRoutes_routes_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQGeoRouteReply_DestroyQGeoRouteReply
func callbackQGeoRouteReply_DestroyQGeoRouteReply(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGeoRouteReply"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGeoRouteReplyFromPointer(ptr).DestroyQGeoRouteReplyDefault()
	}
}

func (ptr *QGeoRouteReply) ConnectDestroyQGeoRouteReply(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGeoRouteReply"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QGeoRouteReply", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGeoRouteReply", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoRouteReply) DisconnectDestroyQGeoRouteReply() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QGeoRouteReply")
	}
}

func (ptr *QGeoRouteReply) DestroyQGeoRouteReply() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoRouteReply_DestroyQGeoRouteReply(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoRouteReply) DestroyQGeoRouteReplyDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoRouteReply_DestroyQGeoRouteReplyDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoRouteReply) __addRoutes_routes_atList(i int) *QGeoRoute {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoRouteFromPointer(C.QGeoRouteReply___addRoutes_routes_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QGeoRoute).DestroyQGeoRoute)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRouteReply) __addRoutes_routes_setList(i QGeoRoute_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRouteReply___addRoutes_routes_setList(ptr.Pointer(), PointerFromQGeoRoute(i))
	}
}

func (ptr *QGeoRouteReply) __addRoutes_routes_newList() unsafe.Pointer {
	return C.QGeoRouteReply___addRoutes_routes_newList(ptr.Pointer())
}

func (ptr *QGeoRouteReply) __routes_atList(i int) *QGeoRoute {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoRouteFromPointer(C.QGeoRouteReply___routes_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QGeoRoute).DestroyQGeoRoute)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRouteReply) __routes_setList(i QGeoRoute_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRouteReply___routes_setList(ptr.Pointer(), PointerFromQGeoRoute(i))
	}
}

func (ptr *QGeoRouteReply) __routes_newList() unsafe.Pointer {
	return C.QGeoRouteReply___routes_newList(ptr.Pointer())
}

func (ptr *QGeoRouteReply) __setRoutes_routes_atList(i int) *QGeoRoute {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoRouteFromPointer(C.QGeoRouteReply___setRoutes_routes_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QGeoRoute).DestroyQGeoRoute)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRouteReply) __setRoutes_routes_setList(i QGeoRoute_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRouteReply___setRoutes_routes_setList(ptr.Pointer(), PointerFromQGeoRoute(i))
	}
}

func (ptr *QGeoRouteReply) __setRoutes_routes_newList() unsafe.Pointer {
	return C.QGeoRouteReply___setRoutes_routes_newList(ptr.Pointer())
}

func (ptr *QGeoRouteReply) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGeoRouteReply___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRouteReply) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRouteReply___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGeoRouteReply) __children_newList() unsafe.Pointer {
	return C.QGeoRouteReply___children_newList(ptr.Pointer())
}

func (ptr *QGeoRouteReply) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QGeoRouteReply___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRouteReply) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRouteReply___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QGeoRouteReply) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QGeoRouteReply___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QGeoRouteReply) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGeoRouteReply___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRouteReply) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRouteReply___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGeoRouteReply) __findChildren_newList() unsafe.Pointer {
	return C.QGeoRouteReply___findChildren_newList(ptr.Pointer())
}

func (ptr *QGeoRouteReply) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGeoRouteReply___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRouteReply) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRouteReply___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGeoRouteReply) __findChildren_newList3() unsafe.Pointer {
	return C.QGeoRouteReply___findChildren_newList3(ptr.Pointer())
}

//export callbackQGeoRouteReply_ChildEvent
func callbackQGeoRouteReply_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGeoRouteReplyFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGeoRouteReply) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRouteReply_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQGeoRouteReply_ConnectNotify
func callbackQGeoRouteReply_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGeoRouteReplyFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGeoRouteReply) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRouteReply_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGeoRouteReply_CustomEvent
func callbackQGeoRouteReply_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQGeoRouteReplyFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGeoRouteReply) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRouteReply_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGeoRouteReply_DeleteLater
func callbackQGeoRouteReply_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGeoRouteReplyFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGeoRouteReply) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoRouteReply_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQGeoRouteReply_Destroyed
func callbackQGeoRouteReply_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQGeoRouteReply_DisconnectNotify
func callbackQGeoRouteReply_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGeoRouteReplyFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGeoRouteReply) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRouteReply_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGeoRouteReply_Event
func callbackQGeoRouteReply_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGeoRouteReplyFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QGeoRouteReply) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoRouteReply_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQGeoRouteReply_EventFilter
func callbackQGeoRouteReply_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGeoRouteReplyFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGeoRouteReply) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoRouteReply_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQGeoRouteReply_MetaObject
func callbackQGeoRouteReply_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQGeoRouteReplyFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGeoRouteReply) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGeoRouteReply_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQGeoRouteReply_ObjectNameChanged
func callbackQGeoRouteReply_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtLocation_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQGeoRouteReply_TimerEvent
func callbackQGeoRouteReply_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGeoRouteReplyFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGeoRouteReply) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRouteReply_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QGeoRouteRequest struct {
	ptr unsafe.Pointer
}

type QGeoRouteRequest_ITF interface {
	QGeoRouteRequest_PTR() *QGeoRouteRequest
}

func (ptr *QGeoRouteRequest) QGeoRouteRequest_PTR() *QGeoRouteRequest {
	return ptr
}

func (ptr *QGeoRouteRequest) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QGeoRouteRequest) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQGeoRouteRequest(ptr QGeoRouteRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoRouteRequest_PTR().Pointer()
	}
	return nil
}

func NewQGeoRouteRequestFromPointer(ptr unsafe.Pointer) (n *QGeoRouteRequest) {
	n = new(QGeoRouteRequest)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QGeoRouteRequest__TravelMode
//QGeoRouteRequest::TravelMode
type QGeoRouteRequest__TravelMode int64

const (
	QGeoRouteRequest__CarTravel           QGeoRouteRequest__TravelMode = QGeoRouteRequest__TravelMode(0x0001)
	QGeoRouteRequest__PedestrianTravel    QGeoRouteRequest__TravelMode = QGeoRouteRequest__TravelMode(0x0002)
	QGeoRouteRequest__BicycleTravel       QGeoRouteRequest__TravelMode = QGeoRouteRequest__TravelMode(0x0004)
	QGeoRouteRequest__PublicTransitTravel QGeoRouteRequest__TravelMode = QGeoRouteRequest__TravelMode(0x0008)
	QGeoRouteRequest__TruckTravel         QGeoRouteRequest__TravelMode = QGeoRouteRequest__TravelMode(0x0010)
)

//go:generate stringer -type=QGeoRouteRequest__FeatureType
//QGeoRouteRequest::FeatureType
type QGeoRouteRequest__FeatureType int64

const (
	QGeoRouteRequest__NoFeature            QGeoRouteRequest__FeatureType = QGeoRouteRequest__FeatureType(0x00000000)
	QGeoRouteRequest__TollFeature          QGeoRouteRequest__FeatureType = QGeoRouteRequest__FeatureType(0x00000001)
	QGeoRouteRequest__HighwayFeature       QGeoRouteRequest__FeatureType = QGeoRouteRequest__FeatureType(0x00000002)
	QGeoRouteRequest__PublicTransitFeature QGeoRouteRequest__FeatureType = QGeoRouteRequest__FeatureType(0x00000004)
	QGeoRouteRequest__FerryFeature         QGeoRouteRequest__FeatureType = QGeoRouteRequest__FeatureType(0x00000008)
	QGeoRouteRequest__TunnelFeature        QGeoRouteRequest__FeatureType = QGeoRouteRequest__FeatureType(0x00000010)
	QGeoRouteRequest__DirtRoadFeature      QGeoRouteRequest__FeatureType = QGeoRouteRequest__FeatureType(0x00000020)
	QGeoRouteRequest__ParksFeature         QGeoRouteRequest__FeatureType = QGeoRouteRequest__FeatureType(0x00000040)
	QGeoRouteRequest__MotorPoolLaneFeature QGeoRouteRequest__FeatureType = QGeoRouteRequest__FeatureType(0x00000080)
	QGeoRouteRequest__TrafficFeature       QGeoRouteRequest__FeatureType = QGeoRouteRequest__FeatureType(0x00000100)
)

//go:generate stringer -type=QGeoRouteRequest__FeatureWeight
//QGeoRouteRequest::FeatureWeight
type QGeoRouteRequest__FeatureWeight int64

const (
	QGeoRouteRequest__NeutralFeatureWeight  QGeoRouteRequest__FeatureWeight = QGeoRouteRequest__FeatureWeight(0x00000000)
	QGeoRouteRequest__PreferFeatureWeight   QGeoRouteRequest__FeatureWeight = QGeoRouteRequest__FeatureWeight(0x00000001)
	QGeoRouteRequest__RequireFeatureWeight  QGeoRouteRequest__FeatureWeight = QGeoRouteRequest__FeatureWeight(0x00000002)
	QGeoRouteRequest__AvoidFeatureWeight    QGeoRouteRequest__FeatureWeight = QGeoRouteRequest__FeatureWeight(0x00000004)
	QGeoRouteRequest__DisallowFeatureWeight QGeoRouteRequest__FeatureWeight = QGeoRouteRequest__FeatureWeight(0x00000008)
)

//go:generate stringer -type=QGeoRouteRequest__RouteOptimization
//QGeoRouteRequest::RouteOptimization
type QGeoRouteRequest__RouteOptimization int64

const (
	QGeoRouteRequest__ShortestRoute     QGeoRouteRequest__RouteOptimization = QGeoRouteRequest__RouteOptimization(0x0001)
	QGeoRouteRequest__FastestRoute      QGeoRouteRequest__RouteOptimization = QGeoRouteRequest__RouteOptimization(0x0002)
	QGeoRouteRequest__MostEconomicRoute QGeoRouteRequest__RouteOptimization = QGeoRouteRequest__RouteOptimization(0x0004)
	QGeoRouteRequest__MostScenicRoute   QGeoRouteRequest__RouteOptimization = QGeoRouteRequest__RouteOptimization(0x0008)
)

//go:generate stringer -type=QGeoRouteRequest__SegmentDetail
//QGeoRouteRequest::SegmentDetail
type QGeoRouteRequest__SegmentDetail int64

const (
	QGeoRouteRequest__NoSegmentData    QGeoRouteRequest__SegmentDetail = QGeoRouteRequest__SegmentDetail(0x0000)
	QGeoRouteRequest__BasicSegmentData QGeoRouteRequest__SegmentDetail = QGeoRouteRequest__SegmentDetail(0x0001)
)

//go:generate stringer -type=QGeoRouteRequest__ManeuverDetail
//QGeoRouteRequest::ManeuverDetail
type QGeoRouteRequest__ManeuverDetail int64

const (
	QGeoRouteRequest__NoManeuvers    QGeoRouteRequest__ManeuverDetail = QGeoRouteRequest__ManeuverDetail(0x0000)
	QGeoRouteRequest__BasicManeuvers QGeoRouteRequest__ManeuverDetail = QGeoRouteRequest__ManeuverDetail(0x0001)
)

func NewQGeoRouteRequest(waypoints []*positioning.QGeoCoordinate) *QGeoRouteRequest {
	tmpValue := NewQGeoRouteRequestFromPointer(C.QGeoRouteRequest_NewQGeoRouteRequest(func() unsafe.Pointer {
		tmpList := NewQGeoRouteRequestFromPointer(NewQGeoRouteRequestFromPointer(nil).__QGeoRouteRequest_waypoints_newList())
		for _, v := range waypoints {
			tmpList.__QGeoRouteRequest_waypoints_setList(v)
		}
		return tmpList.Pointer()
	}()))
	qt.SetFinalizer(tmpValue, (*QGeoRouteRequest).DestroyQGeoRouteRequest)
	return tmpValue
}

func NewQGeoRouteRequest2(origin positioning.QGeoCoordinate_ITF, destination positioning.QGeoCoordinate_ITF) *QGeoRouteRequest {
	tmpValue := NewQGeoRouteRequestFromPointer(C.QGeoRouteRequest_NewQGeoRouteRequest2(positioning.PointerFromQGeoCoordinate(origin), positioning.PointerFromQGeoCoordinate(destination)))
	qt.SetFinalizer(tmpValue, (*QGeoRouteRequest).DestroyQGeoRouteRequest)
	return tmpValue
}

func NewQGeoRouteRequest3(other QGeoRouteRequest_ITF) *QGeoRouteRequest {
	tmpValue := NewQGeoRouteRequestFromPointer(C.QGeoRouteRequest_NewQGeoRouteRequest3(PointerFromQGeoRouteRequest(other)))
	qt.SetFinalizer(tmpValue, (*QGeoRouteRequest).DestroyQGeoRouteRequest)
	return tmpValue
}

func (ptr *QGeoRouteRequest) DepartureTime() *core.QDateTime {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQDateTimeFromPointer(C.QGeoRouteRequest_DepartureTime(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRouteRequest) ExcludeAreas() []*positioning.QGeoRectangle {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) []*positioning.QGeoRectangle {
			out := make([]*positioning.QGeoRectangle, int(l.len))
			tmpList := NewQGeoRouteRequestFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__excludeAreas_atList(i)
			}
			return out
		}(C.QGeoRouteRequest_ExcludeAreas(ptr.Pointer()))
	}
	return make([]*positioning.QGeoRectangle, 0)
}

func (ptr *QGeoRouteRequest) ExtraParameters() map[string]*core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) map[string]*core.QVariant {
			out := make(map[string]*core.QVariant, int(l.len))
			tmpList := NewQGeoRouteRequestFromPointer(l.data)
			for i, v := range tmpList.__extraParameters_keyList() {
				out[v] = tmpList.__extraParameters_atList(v, i)
			}
			return out
		}(C.QGeoRouteRequest_ExtraParameters(ptr.Pointer()))
	}
	return make(map[string]*core.QVariant, 0)
}

func (ptr *QGeoRouteRequest) FeatureWeight(featureType QGeoRouteRequest__FeatureType) QGeoRouteRequest__FeatureWeight {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__FeatureWeight(C.QGeoRouteRequest_FeatureWeight(ptr.Pointer(), C.longlong(featureType)))
	}
	return 0
}

func (ptr *QGeoRouteRequest) ManeuverDetail() QGeoRouteRequest__ManeuverDetail {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__ManeuverDetail(C.QGeoRouteRequest_ManeuverDetail(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRouteRequest) NumberAlternativeRoutes() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGeoRouteRequest_NumberAlternativeRoutes(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoRouteRequest) RouteOptimization() QGeoRouteRequest__RouteOptimization {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__RouteOptimization(C.QGeoRouteRequest_RouteOptimization(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRouteRequest) SegmentDetail() QGeoRouteRequest__SegmentDetail {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__SegmentDetail(C.QGeoRouteRequest_SegmentDetail(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRouteRequest) SetDepartureTime(departureTime core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRouteRequest_SetDepartureTime(ptr.Pointer(), core.PointerFromQDateTime(departureTime))
	}
}

func (ptr *QGeoRouteRequest) SetExcludeAreas(areas []*positioning.QGeoRectangle) {
	if ptr.Pointer() != nil {
		C.QGeoRouteRequest_SetExcludeAreas(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQGeoRouteRequestFromPointer(NewQGeoRouteRequestFromPointer(nil).__setExcludeAreas_areas_newList())
			for _, v := range areas {
				tmpList.__setExcludeAreas_areas_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QGeoRouteRequest) SetExtraParameters(extraParameters map[string]*core.QVariant) {
	if ptr.Pointer() != nil {
		C.QGeoRouteRequest_SetExtraParameters(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQGeoRouteRequestFromPointer(NewQGeoRouteRequestFromPointer(nil).__setExtraParameters_extraParameters_newList())
			for k, v := range extraParameters {
				tmpList.__setExtraParameters_extraParameters_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QGeoRouteRequest) SetFeatureWeight(featureType QGeoRouteRequest__FeatureType, featureWeight QGeoRouteRequest__FeatureWeight) {
	if ptr.Pointer() != nil {
		C.QGeoRouteRequest_SetFeatureWeight(ptr.Pointer(), C.longlong(featureType), C.longlong(featureWeight))
	}
}

func (ptr *QGeoRouteRequest) SetManeuverDetail(maneuverDetail QGeoRouteRequest__ManeuverDetail) {
	if ptr.Pointer() != nil {
		C.QGeoRouteRequest_SetManeuverDetail(ptr.Pointer(), C.longlong(maneuverDetail))
	}
}

func (ptr *QGeoRouteRequest) SetNumberAlternativeRoutes(alternatives int) {
	if ptr.Pointer() != nil {
		C.QGeoRouteRequest_SetNumberAlternativeRoutes(ptr.Pointer(), C.int(int32(alternatives)))
	}
}

func (ptr *QGeoRouteRequest) SetRouteOptimization(optimization QGeoRouteRequest__RouteOptimization) {
	if ptr.Pointer() != nil {
		C.QGeoRouteRequest_SetRouteOptimization(ptr.Pointer(), C.longlong(optimization))
	}
}

func (ptr *QGeoRouteRequest) SetSegmentDetail(segmentDetail QGeoRouteRequest__SegmentDetail) {
	if ptr.Pointer() != nil {
		C.QGeoRouteRequest_SetSegmentDetail(ptr.Pointer(), C.longlong(segmentDetail))
	}
}

func (ptr *QGeoRouteRequest) SetTravelModes(travelModes QGeoRouteRequest__TravelMode) {
	if ptr.Pointer() != nil {
		C.QGeoRouteRequest_SetTravelModes(ptr.Pointer(), C.longlong(travelModes))
	}
}

func (ptr *QGeoRouteRequest) SetWaypoints(waypoints []*positioning.QGeoCoordinate) {
	if ptr.Pointer() != nil {
		C.QGeoRouteRequest_SetWaypoints(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQGeoRouteRequestFromPointer(NewQGeoRouteRequestFromPointer(nil).__setWaypoints_waypoints_newList())
			for _, v := range waypoints {
				tmpList.__setWaypoints_waypoints_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QGeoRouteRequest) TravelModes() QGeoRouteRequest__TravelMode {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__TravelMode(C.QGeoRouteRequest_TravelModes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRouteRequest) Waypoints() []*positioning.QGeoCoordinate {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) []*positioning.QGeoCoordinate {
			out := make([]*positioning.QGeoCoordinate, int(l.len))
			tmpList := NewQGeoRouteRequestFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__waypoints_atList(i)
			}
			return out
		}(C.QGeoRouteRequest_Waypoints(ptr.Pointer()))
	}
	return make([]*positioning.QGeoCoordinate, 0)
}

func (ptr *QGeoRouteRequest) DestroyQGeoRouteRequest() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoRouteRequest_DestroyQGeoRouteRequest(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoRouteRequest) __QGeoRouteRequest_waypoints_atList(i int) *positioning.QGeoCoordinate {
	if ptr.Pointer() != nil {
		tmpValue := positioning.NewQGeoCoordinateFromPointer(C.QGeoRouteRequest___QGeoRouteRequest_waypoints_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*positioning.QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRouteRequest) __QGeoRouteRequest_waypoints_setList(i positioning.QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRouteRequest___QGeoRouteRequest_waypoints_setList(ptr.Pointer(), positioning.PointerFromQGeoCoordinate(i))
	}
}

func (ptr *QGeoRouteRequest) __QGeoRouteRequest_waypoints_newList() unsafe.Pointer {
	return C.QGeoRouteRequest___QGeoRouteRequest_waypoints_newList(ptr.Pointer())
}

func (ptr *QGeoRouteRequest) __excludeAreas_atList(i int) *positioning.QGeoRectangle {
	if ptr.Pointer() != nil {
		tmpValue := positioning.NewQGeoRectangleFromPointer(C.QGeoRouteRequest___excludeAreas_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*positioning.QGeoRectangle).DestroyQGeoRectangle)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRouteRequest) __excludeAreas_setList(i positioning.QGeoRectangle_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRouteRequest___excludeAreas_setList(ptr.Pointer(), positioning.PointerFromQGeoRectangle(i))
	}
}

func (ptr *QGeoRouteRequest) __excludeAreas_newList() unsafe.Pointer {
	return C.QGeoRouteRequest___excludeAreas_newList(ptr.Pointer())
}

func (ptr *QGeoRouteRequest) __extraParameters_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QGeoRouteRequest___extraParameters_atList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRouteRequest) __extraParameters_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QGeoRouteRequest___extraParameters_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QGeoRouteRequest) __extraParameters_newList() unsafe.Pointer {
	return C.QGeoRouteRequest___extraParameters_newList(ptr.Pointer())
}

func (ptr *QGeoRouteRequest) __extraParameters_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQGeoRouteRequestFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____extraParameters_keyList_atList(i)
			}
			return out
		}(C.QGeoRouteRequest___extraParameters_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QGeoRouteRequest) __setExcludeAreas_areas_atList(i int) *positioning.QGeoRectangle {
	if ptr.Pointer() != nil {
		tmpValue := positioning.NewQGeoRectangleFromPointer(C.QGeoRouteRequest___setExcludeAreas_areas_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*positioning.QGeoRectangle).DestroyQGeoRectangle)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRouteRequest) __setExcludeAreas_areas_setList(i positioning.QGeoRectangle_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRouteRequest___setExcludeAreas_areas_setList(ptr.Pointer(), positioning.PointerFromQGeoRectangle(i))
	}
}

func (ptr *QGeoRouteRequest) __setExcludeAreas_areas_newList() unsafe.Pointer {
	return C.QGeoRouteRequest___setExcludeAreas_areas_newList(ptr.Pointer())
}

func (ptr *QGeoRouteRequest) __setExtraParameters_extraParameters_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QGeoRouteRequest___setExtraParameters_extraParameters_atList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRouteRequest) __setExtraParameters_extraParameters_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QGeoRouteRequest___setExtraParameters_extraParameters_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QGeoRouteRequest) __setExtraParameters_extraParameters_newList() unsafe.Pointer {
	return C.QGeoRouteRequest___setExtraParameters_extraParameters_newList(ptr.Pointer())
}

func (ptr *QGeoRouteRequest) __setExtraParameters_extraParameters_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQGeoRouteRequestFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setExtraParameters_extraParameters_keyList_atList(i)
			}
			return out
		}(C.QGeoRouteRequest___setExtraParameters_extraParameters_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QGeoRouteRequest) __setWaypoints_waypoints_atList(i int) *positioning.QGeoCoordinate {
	if ptr.Pointer() != nil {
		tmpValue := positioning.NewQGeoCoordinateFromPointer(C.QGeoRouteRequest___setWaypoints_waypoints_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*positioning.QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRouteRequest) __setWaypoints_waypoints_setList(i positioning.QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRouteRequest___setWaypoints_waypoints_setList(ptr.Pointer(), positioning.PointerFromQGeoCoordinate(i))
	}
}

func (ptr *QGeoRouteRequest) __setWaypoints_waypoints_newList() unsafe.Pointer {
	return C.QGeoRouteRequest___setWaypoints_waypoints_newList(ptr.Pointer())
}

func (ptr *QGeoRouteRequest) __setWaypointsMetadata_waypointMetadata_atList(i int) map[string]*core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) map[string]*core.QVariant {
			out := make(map[string]*core.QVariant, int(l.len))
			tmpList := NewQGeoRouteRequestFromPointer(l.data)
			for i, v := range tmpList.____setWaypointsMetadata_waypointMetadata_atList_keyList() {
				out[v] = tmpList.____setWaypointsMetadata_waypointMetadata_atList_atList(v, i)
			}
			return out
		}(C.QGeoRouteRequest___setWaypointsMetadata_waypointMetadata_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return make(map[string]*core.QVariant, 0)
}

func (ptr *QGeoRouteRequest) __setWaypointsMetadata_waypointMetadata_setList(i map[string]*core.QVariant) {
	if ptr.Pointer() != nil {
		C.QGeoRouteRequest___setWaypointsMetadata_waypointMetadata_setList(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQGeoRouteRequestFromPointer(NewQGeoRouteRequestFromPointer(nil).____setWaypointsMetadata_waypointMetadata_setList_i_newList())
			for k, v := range i {
				tmpList.____setWaypointsMetadata_waypointMetadata_setList_i_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QGeoRouteRequest) __setWaypointsMetadata_waypointMetadata_newList() unsafe.Pointer {
	return C.QGeoRouteRequest___setWaypointsMetadata_waypointMetadata_newList(ptr.Pointer())
}

func (ptr *QGeoRouteRequest) __waypoints_atList(i int) *positioning.QGeoCoordinate {
	if ptr.Pointer() != nil {
		tmpValue := positioning.NewQGeoCoordinateFromPointer(C.QGeoRouteRequest___waypoints_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*positioning.QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRouteRequest) __waypoints_setList(i positioning.QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRouteRequest___waypoints_setList(ptr.Pointer(), positioning.PointerFromQGeoCoordinate(i))
	}
}

func (ptr *QGeoRouteRequest) __waypoints_newList() unsafe.Pointer {
	return C.QGeoRouteRequest___waypoints_newList(ptr.Pointer())
}

func (ptr *QGeoRouteRequest) __waypointsMetadata_atList(i int) map[string]*core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) map[string]*core.QVariant {
			out := make(map[string]*core.QVariant, int(l.len))
			tmpList := NewQGeoRouteRequestFromPointer(l.data)
			for i, v := range tmpList.____waypointsMetadata_atList_keyList() {
				out[v] = tmpList.____waypointsMetadata_atList_atList(v, i)
			}
			return out
		}(C.QGeoRouteRequest___waypointsMetadata_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return make(map[string]*core.QVariant, 0)
}

func (ptr *QGeoRouteRequest) __waypointsMetadata_setList(i map[string]*core.QVariant) {
	if ptr.Pointer() != nil {
		C.QGeoRouteRequest___waypointsMetadata_setList(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQGeoRouteRequestFromPointer(NewQGeoRouteRequestFromPointer(nil).____waypointsMetadata_setList_i_newList())
			for k, v := range i {
				tmpList.____waypointsMetadata_setList_i_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QGeoRouteRequest) __waypointsMetadata_newList() unsafe.Pointer {
	return C.QGeoRouteRequest___waypointsMetadata_newList(ptr.Pointer())
}

func (ptr *QGeoRouteRequest) ____extraParameters_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoRouteRequest_____extraParameters_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QGeoRouteRequest) ____extraParameters_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QGeoRouteRequest_____extraParameters_keyList_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QGeoRouteRequest) ____extraParameters_keyList_newList() unsafe.Pointer {
	return C.QGeoRouteRequest_____extraParameters_keyList_newList(ptr.Pointer())
}

func (ptr *QGeoRouteRequest) ____setExtraParameters_extraParameters_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoRouteRequest_____setExtraParameters_extraParameters_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QGeoRouteRequest) ____setExtraParameters_extraParameters_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QGeoRouteRequest_____setExtraParameters_extraParameters_keyList_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QGeoRouteRequest) ____setExtraParameters_extraParameters_keyList_newList() unsafe.Pointer {
	return C.QGeoRouteRequest_____setExtraParameters_extraParameters_keyList_newList(ptr.Pointer())
}

func (ptr *QGeoRouteRequest) ____setWaypointsMetadata_waypointMetadata_atList_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QGeoRouteRequest_____setWaypointsMetadata_waypointMetadata_atList_atList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRouteRequest) ____setWaypointsMetadata_waypointMetadata_atList_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QGeoRouteRequest_____setWaypointsMetadata_waypointMetadata_atList_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QGeoRouteRequest) ____setWaypointsMetadata_waypointMetadata_atList_newList() unsafe.Pointer {
	return C.QGeoRouteRequest_____setWaypointsMetadata_waypointMetadata_atList_newList(ptr.Pointer())
}

func (ptr *QGeoRouteRequest) ____setWaypointsMetadata_waypointMetadata_atList_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQGeoRouteRequestFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.______setWaypointsMetadata_waypointMetadata_atList_keyList_atList(i)
			}
			return out
		}(C.QGeoRouteRequest_____setWaypointsMetadata_waypointMetadata_atList_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QGeoRouteRequest) ____setWaypointsMetadata_waypointMetadata_setList_i_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QGeoRouteRequest_____setWaypointsMetadata_waypointMetadata_setList_i_atList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRouteRequest) ____setWaypointsMetadata_waypointMetadata_setList_i_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QGeoRouteRequest_____setWaypointsMetadata_waypointMetadata_setList_i_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QGeoRouteRequest) ____setWaypointsMetadata_waypointMetadata_setList_i_newList() unsafe.Pointer {
	return C.QGeoRouteRequest_____setWaypointsMetadata_waypointMetadata_setList_i_newList(ptr.Pointer())
}

func (ptr *QGeoRouteRequest) ____setWaypointsMetadata_waypointMetadata_setList_i_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQGeoRouteRequestFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.______setWaypointsMetadata_waypointMetadata_setList_i_keyList_atList(i)
			}
			return out
		}(C.QGeoRouteRequest_____setWaypointsMetadata_waypointMetadata_setList_i_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QGeoRouteRequest) ____waypointsMetadata_atList_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QGeoRouteRequest_____waypointsMetadata_atList_atList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRouteRequest) ____waypointsMetadata_atList_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QGeoRouteRequest_____waypointsMetadata_atList_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QGeoRouteRequest) ____waypointsMetadata_atList_newList() unsafe.Pointer {
	return C.QGeoRouteRequest_____waypointsMetadata_atList_newList(ptr.Pointer())
}

func (ptr *QGeoRouteRequest) ____waypointsMetadata_atList_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQGeoRouteRequestFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.______waypointsMetadata_atList_keyList_atList(i)
			}
			return out
		}(C.QGeoRouteRequest_____waypointsMetadata_atList_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QGeoRouteRequest) ____waypointsMetadata_setList_i_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QGeoRouteRequest_____waypointsMetadata_setList_i_atList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRouteRequest) ____waypointsMetadata_setList_i_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QGeoRouteRequest_____waypointsMetadata_setList_i_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QGeoRouteRequest) ____waypointsMetadata_setList_i_newList() unsafe.Pointer {
	return C.QGeoRouteRequest_____waypointsMetadata_setList_i_newList(ptr.Pointer())
}

func (ptr *QGeoRouteRequest) ____waypointsMetadata_setList_i_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQGeoRouteRequestFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.______waypointsMetadata_setList_i_keyList_atList(i)
			}
			return out
		}(C.QGeoRouteRequest_____waypointsMetadata_setList_i_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QGeoRouteRequest) ______setWaypointsMetadata_waypointMetadata_atList_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoRouteRequest_______setWaypointsMetadata_waypointMetadata_atList_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QGeoRouteRequest) ______setWaypointsMetadata_waypointMetadata_atList_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QGeoRouteRequest_______setWaypointsMetadata_waypointMetadata_atList_keyList_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QGeoRouteRequest) ______setWaypointsMetadata_waypointMetadata_atList_keyList_newList() unsafe.Pointer {
	return C.QGeoRouteRequest_______setWaypointsMetadata_waypointMetadata_atList_keyList_newList(ptr.Pointer())
}

func (ptr *QGeoRouteRequest) ______setWaypointsMetadata_waypointMetadata_setList_i_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoRouteRequest_______setWaypointsMetadata_waypointMetadata_setList_i_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QGeoRouteRequest) ______setWaypointsMetadata_waypointMetadata_setList_i_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QGeoRouteRequest_______setWaypointsMetadata_waypointMetadata_setList_i_keyList_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QGeoRouteRequest) ______setWaypointsMetadata_waypointMetadata_setList_i_keyList_newList() unsafe.Pointer {
	return C.QGeoRouteRequest_______setWaypointsMetadata_waypointMetadata_setList_i_keyList_newList(ptr.Pointer())
}

func (ptr *QGeoRouteRequest) ______waypointsMetadata_atList_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoRouteRequest_______waypointsMetadata_atList_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QGeoRouteRequest) ______waypointsMetadata_atList_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QGeoRouteRequest_______waypointsMetadata_atList_keyList_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QGeoRouteRequest) ______waypointsMetadata_atList_keyList_newList() unsafe.Pointer {
	return C.QGeoRouteRequest_______waypointsMetadata_atList_keyList_newList(ptr.Pointer())
}

func (ptr *QGeoRouteRequest) ______waypointsMetadata_setList_i_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoRouteRequest_______waypointsMetadata_setList_i_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QGeoRouteRequest) ______waypointsMetadata_setList_i_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QGeoRouteRequest_______waypointsMetadata_setList_i_keyList_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QGeoRouteRequest) ______waypointsMetadata_setList_i_keyList_newList() unsafe.Pointer {
	return C.QGeoRouteRequest_______waypointsMetadata_setList_i_keyList_newList(ptr.Pointer())
}

type QGeoRouteSegment struct {
	ptr unsafe.Pointer
}

type QGeoRouteSegment_ITF interface {
	QGeoRouteSegment_PTR() *QGeoRouteSegment
}

func (ptr *QGeoRouteSegment) QGeoRouteSegment_PTR() *QGeoRouteSegment {
	return ptr
}

func (ptr *QGeoRouteSegment) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QGeoRouteSegment) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQGeoRouteSegment(ptr QGeoRouteSegment_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoRouteSegment_PTR().Pointer()
	}
	return nil
}

func NewQGeoRouteSegmentFromPointer(ptr unsafe.Pointer) (n *QGeoRouteSegment) {
	n = new(QGeoRouteSegment)
	n.SetPointer(ptr)
	return
}
func NewQGeoRouteSegment() *QGeoRouteSegment {
	tmpValue := NewQGeoRouteSegmentFromPointer(C.QGeoRouteSegment_NewQGeoRouteSegment())
	qt.SetFinalizer(tmpValue, (*QGeoRouteSegment).DestroyQGeoRouteSegment)
	return tmpValue
}

func NewQGeoRouteSegment2(other QGeoRouteSegment_ITF) *QGeoRouteSegment {
	tmpValue := NewQGeoRouteSegmentFromPointer(C.QGeoRouteSegment_NewQGeoRouteSegment2(PointerFromQGeoRouteSegment(other)))
	qt.SetFinalizer(tmpValue, (*QGeoRouteSegment).DestroyQGeoRouteSegment)
	return tmpValue
}

func (ptr *QGeoRouteSegment) Distance() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoRouteSegment_Distance(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRouteSegment) IsLegLastSegment() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoRouteSegment_IsLegLastSegment(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGeoRouteSegment) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoRouteSegment_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGeoRouteSegment) Maneuver() *QGeoManeuver {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoManeuverFromPointer(C.QGeoRouteSegment_Maneuver(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QGeoManeuver).DestroyQGeoManeuver)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRouteSegment) NextRouteSegment() *QGeoRouteSegment {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoRouteSegmentFromPointer(C.QGeoRouteSegment_NextRouteSegment(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QGeoRouteSegment).DestroyQGeoRouteSegment)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRouteSegment) Path() []*positioning.QGeoCoordinate {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) []*positioning.QGeoCoordinate {
			out := make([]*positioning.QGeoCoordinate, int(l.len))
			tmpList := NewQGeoRouteSegmentFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__path_atList(i)
			}
			return out
		}(C.QGeoRouteSegment_Path(ptr.Pointer()))
	}
	return make([]*positioning.QGeoCoordinate, 0)
}

func (ptr *QGeoRouteSegment) SetDistance(distance float64) {
	if ptr.Pointer() != nil {
		C.QGeoRouteSegment_SetDistance(ptr.Pointer(), C.double(distance))
	}
}

func (ptr *QGeoRouteSegment) SetManeuver(maneuver QGeoManeuver_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRouteSegment_SetManeuver(ptr.Pointer(), PointerFromQGeoManeuver(maneuver))
	}
}

func (ptr *QGeoRouteSegment) SetNextRouteSegment(routeSegment QGeoRouteSegment_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRouteSegment_SetNextRouteSegment(ptr.Pointer(), PointerFromQGeoRouteSegment(routeSegment))
	}
}

func (ptr *QGeoRouteSegment) SetPath(path []*positioning.QGeoCoordinate) {
	if ptr.Pointer() != nil {
		C.QGeoRouteSegment_SetPath(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQGeoRouteSegmentFromPointer(NewQGeoRouteSegmentFromPointer(nil).__setPath_path_newList())
			for _, v := range path {
				tmpList.__setPath_path_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QGeoRouteSegment) SetTravelTime(secs int) {
	if ptr.Pointer() != nil {
		C.QGeoRouteSegment_SetTravelTime(ptr.Pointer(), C.int(int32(secs)))
	}
}

func (ptr *QGeoRouteSegment) TravelTime() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGeoRouteSegment_TravelTime(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoRouteSegment) DestroyQGeoRouteSegment() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoRouteSegment_DestroyQGeoRouteSegment(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoRouteSegment) __path_atList(i int) *positioning.QGeoCoordinate {
	if ptr.Pointer() != nil {
		tmpValue := positioning.NewQGeoCoordinateFromPointer(C.QGeoRouteSegment___path_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*positioning.QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRouteSegment) __path_setList(i positioning.QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRouteSegment___path_setList(ptr.Pointer(), positioning.PointerFromQGeoCoordinate(i))
	}
}

func (ptr *QGeoRouteSegment) __path_newList() unsafe.Pointer {
	return C.QGeoRouteSegment___path_newList(ptr.Pointer())
}

func (ptr *QGeoRouteSegment) __setPath_path_atList(i int) *positioning.QGeoCoordinate {
	if ptr.Pointer() != nil {
		tmpValue := positioning.NewQGeoCoordinateFromPointer(C.QGeoRouteSegment___setPath_path_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*positioning.QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRouteSegment) __setPath_path_setList(i positioning.QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRouteSegment___setPath_path_setList(ptr.Pointer(), positioning.PointerFromQGeoCoordinate(i))
	}
}

func (ptr *QGeoRouteSegment) __setPath_path_newList() unsafe.Pointer {
	return C.QGeoRouteSegment___setPath_path_newList(ptr.Pointer())
}

type QGeoRoutingManager struct {
	core.QObject
}

type QGeoRoutingManager_ITF interface {
	core.QObject_ITF
	QGeoRoutingManager_PTR() *QGeoRoutingManager
}

func (ptr *QGeoRoutingManager) QGeoRoutingManager_PTR() *QGeoRoutingManager {
	return ptr
}

func (ptr *QGeoRoutingManager) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QGeoRoutingManager) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQGeoRoutingManager(ptr QGeoRoutingManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoRoutingManager_PTR().Pointer()
	}
	return nil
}

func NewQGeoRoutingManagerFromPointer(ptr unsafe.Pointer) (n *QGeoRoutingManager) {
	n = new(QGeoRoutingManager)
	n.SetPointer(ptr)
	return
}
func (ptr *QGeoRoutingManager) CalculateRoute(request QGeoRouteRequest_ITF) *QGeoRouteReply {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoRouteReplyFromPointer(C.QGeoRoutingManager_CalculateRoute(ptr.Pointer(), PointerFromQGeoRouteRequest(request)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQGeoRoutingManager_Error
func callbackQGeoRoutingManager_Error(ptr unsafe.Pointer, reply unsafe.Pointer, error C.longlong, errorString C.struct_QtLocation_PackedString) {
	if signal := qt.GetSignal(ptr, "error"); signal != nil {
		(*(*func(*QGeoRouteReply, QGeoRouteReply__Error, string))(signal))(NewQGeoRouteReplyFromPointer(reply), QGeoRouteReply__Error(error), cGoUnpackString(errorString))
	}

}

func (ptr *QGeoRoutingManager) ConnectError(f func(reply *QGeoRouteReply, error QGeoRouteReply__Error, errorString string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error") {
			C.QGeoRoutingManager_ConnectError(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "error")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error"); signal != nil {
			f := func(reply *QGeoRouteReply, error QGeoRouteReply__Error, errorString string) {
				(*(*func(*QGeoRouteReply, QGeoRouteReply__Error, string))(signal))(reply, error, errorString)
				f(reply, error, errorString)
			}
			qt.ConnectSignal(ptr.Pointer(), "error", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoRoutingManager) DisconnectError() {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManager_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "error")
	}
}

func (ptr *QGeoRoutingManager) Error(reply QGeoRouteReply_ITF, error QGeoRouteReply__Error, errorString string) {
	if ptr.Pointer() != nil {
		var errorStringC *C.char
		if errorString != "" {
			errorStringC = C.CString(errorString)
			defer C.free(unsafe.Pointer(errorStringC))
		}
		C.QGeoRoutingManager_Error(ptr.Pointer(), PointerFromQGeoRouteReply(reply), C.longlong(error), C.struct_QtLocation_PackedString{data: errorStringC, len: C.longlong(len(errorString))})
	}
}

//export callbackQGeoRoutingManager_Finished
func callbackQGeoRoutingManager_Finished(ptr unsafe.Pointer, reply unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "finished"); signal != nil {
		(*(*func(*QGeoRouteReply))(signal))(NewQGeoRouteReplyFromPointer(reply))
	}

}

func (ptr *QGeoRoutingManager) ConnectFinished(f func(reply *QGeoRouteReply)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finished") {
			C.QGeoRoutingManager_ConnectFinished(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "finished")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "finished"); signal != nil {
			f := func(reply *QGeoRouteReply) {
				(*(*func(*QGeoRouteReply))(signal))(reply)
				f(reply)
			}
			qt.ConnectSignal(ptr.Pointer(), "finished", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finished", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoRoutingManager) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManager_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "finished")
	}
}

func (ptr *QGeoRoutingManager) Finished(reply QGeoRouteReply_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManager_Finished(ptr.Pointer(), PointerFromQGeoRouteReply(reply))
	}
}

func (ptr *QGeoRoutingManager) Locale() *core.QLocale {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQLocaleFromPointer(C.QGeoRoutingManager_Locale(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QLocale).DestroyQLocale)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRoutingManager) ManagerName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoRoutingManager_ManagerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoRoutingManager) ManagerVersion() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGeoRoutingManager_ManagerVersion(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoRoutingManager) MeasurementSystem() core.QLocale__MeasurementSystem {
	if ptr.Pointer() != nil {
		return core.QLocale__MeasurementSystem(C.QGeoRoutingManager_MeasurementSystem(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManager) SetLocale(locale core.QLocale_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManager_SetLocale(ptr.Pointer(), core.PointerFromQLocale(locale))
	}
}

func (ptr *QGeoRoutingManager) SetMeasurementSystem(system core.QLocale__MeasurementSystem) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManager_SetMeasurementSystem(ptr.Pointer(), C.longlong(system))
	}
}

func (ptr *QGeoRoutingManager) SupportedFeatureTypes() QGeoRouteRequest__FeatureType {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__FeatureType(C.QGeoRoutingManager_SupportedFeatureTypes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManager) SupportedFeatureWeights() QGeoRouteRequest__FeatureWeight {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__FeatureWeight(C.QGeoRoutingManager_SupportedFeatureWeights(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManager) SupportedManeuverDetails() QGeoRouteRequest__ManeuverDetail {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__ManeuverDetail(C.QGeoRoutingManager_SupportedManeuverDetails(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManager) SupportedRouteOptimizations() QGeoRouteRequest__RouteOptimization {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__RouteOptimization(C.QGeoRoutingManager_SupportedRouteOptimizations(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManager) SupportedSegmentDetails() QGeoRouteRequest__SegmentDetail {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__SegmentDetail(C.QGeoRoutingManager_SupportedSegmentDetails(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManager) SupportedTravelModes() QGeoRouteRequest__TravelMode {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__TravelMode(C.QGeoRoutingManager_SupportedTravelModes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManager) UpdateRoute(route QGeoRoute_ITF, position positioning.QGeoCoordinate_ITF) *QGeoRouteReply {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoRouteReplyFromPointer(C.QGeoRoutingManager_UpdateRoute(ptr.Pointer(), PointerFromQGeoRoute(route), positioning.PointerFromQGeoCoordinate(position)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQGeoRoutingManager_DestroyQGeoRoutingManager
func callbackQGeoRoutingManager_DestroyQGeoRoutingManager(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGeoRoutingManager"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGeoRoutingManagerFromPointer(ptr).DestroyQGeoRoutingManagerDefault()
	}
}

func (ptr *QGeoRoutingManager) ConnectDestroyQGeoRoutingManager(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGeoRoutingManager"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QGeoRoutingManager", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGeoRoutingManager", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoRoutingManager) DisconnectDestroyQGeoRoutingManager() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QGeoRoutingManager")
	}
}

func (ptr *QGeoRoutingManager) DestroyQGeoRoutingManager() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoRoutingManager_DestroyQGeoRoutingManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoRoutingManager) DestroyQGeoRoutingManagerDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoRoutingManager_DestroyQGeoRoutingManagerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoRoutingManager) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGeoRoutingManager___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRoutingManager) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManager___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGeoRoutingManager) __children_newList() unsafe.Pointer {
	return C.QGeoRoutingManager___children_newList(ptr.Pointer())
}

func (ptr *QGeoRoutingManager) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QGeoRoutingManager___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRoutingManager) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManager___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QGeoRoutingManager) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QGeoRoutingManager___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QGeoRoutingManager) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGeoRoutingManager___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRoutingManager) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManager___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGeoRoutingManager) __findChildren_newList() unsafe.Pointer {
	return C.QGeoRoutingManager___findChildren_newList(ptr.Pointer())
}

func (ptr *QGeoRoutingManager) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGeoRoutingManager___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRoutingManager) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManager___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGeoRoutingManager) __findChildren_newList3() unsafe.Pointer {
	return C.QGeoRoutingManager___findChildren_newList3(ptr.Pointer())
}

//export callbackQGeoRoutingManager_ChildEvent
func callbackQGeoRoutingManager_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGeoRoutingManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGeoRoutingManager) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQGeoRoutingManager_ConnectNotify
func callbackQGeoRoutingManager_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGeoRoutingManagerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGeoRoutingManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManager_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGeoRoutingManager_CustomEvent
func callbackQGeoRoutingManager_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQGeoRoutingManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGeoRoutingManager) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGeoRoutingManager_DeleteLater
func callbackQGeoRoutingManager_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGeoRoutingManagerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGeoRoutingManager) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoRoutingManager_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQGeoRoutingManager_Destroyed
func callbackQGeoRoutingManager_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQGeoRoutingManager_DisconnectNotify
func callbackQGeoRoutingManager_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGeoRoutingManagerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGeoRoutingManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManager_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGeoRoutingManager_Event
func callbackQGeoRoutingManager_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGeoRoutingManagerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QGeoRoutingManager) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoRoutingManager_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQGeoRoutingManager_EventFilter
func callbackQGeoRoutingManager_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGeoRoutingManagerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGeoRoutingManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoRoutingManager_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQGeoRoutingManager_MetaObject
func callbackQGeoRoutingManager_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQGeoRoutingManagerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGeoRoutingManager) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGeoRoutingManager_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQGeoRoutingManager_ObjectNameChanged
func callbackQGeoRoutingManager_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtLocation_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQGeoRoutingManager_TimerEvent
func callbackQGeoRoutingManager_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGeoRoutingManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGeoRoutingManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QGeoRoutingManagerEngine struct {
	core.QObject
}

type QGeoRoutingManagerEngine_ITF interface {
	core.QObject_ITF
	QGeoRoutingManagerEngine_PTR() *QGeoRoutingManagerEngine
}

func (ptr *QGeoRoutingManagerEngine) QGeoRoutingManagerEngine_PTR() *QGeoRoutingManagerEngine {
	return ptr
}

func (ptr *QGeoRoutingManagerEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QGeoRoutingManagerEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQGeoRoutingManagerEngine(ptr QGeoRoutingManagerEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoRoutingManagerEngine_PTR().Pointer()
	}
	return nil
}

func NewQGeoRoutingManagerEngineFromPointer(ptr unsafe.Pointer) (n *QGeoRoutingManagerEngine) {
	n = new(QGeoRoutingManagerEngine)
	n.SetPointer(ptr)
	return
}
func NewQGeoRoutingManagerEngine(parameters map[string]*core.QVariant, parent core.QObject_ITF) *QGeoRoutingManagerEngine {
	tmpValue := NewQGeoRoutingManagerEngineFromPointer(C.QGeoRoutingManagerEngine_NewQGeoRoutingManagerEngine(func() unsafe.Pointer {
		tmpList := NewQGeoRoutingManagerEngineFromPointer(NewQGeoRoutingManagerEngineFromPointer(nil).__QGeoRoutingManagerEngine_parameters_newList())
		for k, v := range parameters {
			tmpList.__QGeoRoutingManagerEngine_parameters_setList(k, v)
		}
		return tmpList.Pointer()
	}(), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQGeoRoutingManagerEngine_CalculateRoute
func callbackQGeoRoutingManagerEngine_CalculateRoute(ptr unsafe.Pointer, request unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "calculateRoute"); signal != nil {
		return PointerFromQGeoRouteReply((*(*func(*QGeoRouteRequest) *QGeoRouteReply)(signal))(NewQGeoRouteRequestFromPointer(request)))
	}

	return PointerFromQGeoRouteReply(nil)
}

func (ptr *QGeoRoutingManagerEngine) ConnectCalculateRoute(f func(request *QGeoRouteRequest) *QGeoRouteReply) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "calculateRoute"); signal != nil {
			f := func(request *QGeoRouteRequest) *QGeoRouteReply {
				(*(*func(*QGeoRouteRequest) *QGeoRouteReply)(signal))(request)
				return f(request)
			}
			qt.ConnectSignal(ptr.Pointer(), "calculateRoute", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "calculateRoute", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoRoutingManagerEngine) DisconnectCalculateRoute() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "calculateRoute")
	}
}

func (ptr *QGeoRoutingManagerEngine) CalculateRoute(request QGeoRouteRequest_ITF) *QGeoRouteReply {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoRouteReplyFromPointer(C.QGeoRoutingManagerEngine_CalculateRoute(ptr.Pointer(), PointerFromQGeoRouteRequest(request)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQGeoRoutingManagerEngine_Error
func callbackQGeoRoutingManagerEngine_Error(ptr unsafe.Pointer, reply unsafe.Pointer, error C.longlong, errorString C.struct_QtLocation_PackedString) {
	if signal := qt.GetSignal(ptr, "error"); signal != nil {
		(*(*func(*QGeoRouteReply, QGeoRouteReply__Error, string))(signal))(NewQGeoRouteReplyFromPointer(reply), QGeoRouteReply__Error(error), cGoUnpackString(errorString))
	}

}

func (ptr *QGeoRoutingManagerEngine) ConnectError(f func(reply *QGeoRouteReply, error QGeoRouteReply__Error, errorString string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error") {
			C.QGeoRoutingManagerEngine_ConnectError(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "error")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error"); signal != nil {
			f := func(reply *QGeoRouteReply, error QGeoRouteReply__Error, errorString string) {
				(*(*func(*QGeoRouteReply, QGeoRouteReply__Error, string))(signal))(reply, error, errorString)
				f(reply, error, errorString)
			}
			qt.ConnectSignal(ptr.Pointer(), "error", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoRoutingManagerEngine) DisconnectError() {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "error")
	}
}

func (ptr *QGeoRoutingManagerEngine) Error(reply QGeoRouteReply_ITF, error QGeoRouteReply__Error, errorString string) {
	if ptr.Pointer() != nil {
		var errorStringC *C.char
		if errorString != "" {
			errorStringC = C.CString(errorString)
			defer C.free(unsafe.Pointer(errorStringC))
		}
		C.QGeoRoutingManagerEngine_Error(ptr.Pointer(), PointerFromQGeoRouteReply(reply), C.longlong(error), C.struct_QtLocation_PackedString{data: errorStringC, len: C.longlong(len(errorString))})
	}
}

//export callbackQGeoRoutingManagerEngine_Finished
func callbackQGeoRoutingManagerEngine_Finished(ptr unsafe.Pointer, reply unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "finished"); signal != nil {
		(*(*func(*QGeoRouteReply))(signal))(NewQGeoRouteReplyFromPointer(reply))
	}

}

func (ptr *QGeoRoutingManagerEngine) ConnectFinished(f func(reply *QGeoRouteReply)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finished") {
			C.QGeoRoutingManagerEngine_ConnectFinished(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "finished")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "finished"); signal != nil {
			f := func(reply *QGeoRouteReply) {
				(*(*func(*QGeoRouteReply))(signal))(reply)
				f(reply)
			}
			qt.ConnectSignal(ptr.Pointer(), "finished", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finished", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoRoutingManagerEngine) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "finished")
	}
}

func (ptr *QGeoRoutingManagerEngine) Finished(reply QGeoRouteReply_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_Finished(ptr.Pointer(), PointerFromQGeoRouteReply(reply))
	}
}

func (ptr *QGeoRoutingManagerEngine) Locale() *core.QLocale {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQLocaleFromPointer(C.QGeoRoutingManagerEngine_Locale(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QLocale).DestroyQLocale)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRoutingManagerEngine) ManagerName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoRoutingManagerEngine_ManagerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoRoutingManagerEngine) ManagerVersion() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGeoRoutingManagerEngine_ManagerVersion(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoRoutingManagerEngine) MeasurementSystem() core.QLocale__MeasurementSystem {
	if ptr.Pointer() != nil {
		return core.QLocale__MeasurementSystem(C.QGeoRoutingManagerEngine_MeasurementSystem(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManagerEngine) SetLocale(locale core.QLocale_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_SetLocale(ptr.Pointer(), core.PointerFromQLocale(locale))
	}
}

func (ptr *QGeoRoutingManagerEngine) SetMeasurementSystem(system core.QLocale__MeasurementSystem) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_SetMeasurementSystem(ptr.Pointer(), C.longlong(system))
	}
}

func (ptr *QGeoRoutingManagerEngine) SetSupportedFeatureTypes(featureTypes QGeoRouteRequest__FeatureType) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_SetSupportedFeatureTypes(ptr.Pointer(), C.longlong(featureTypes))
	}
}

func (ptr *QGeoRoutingManagerEngine) SetSupportedFeatureWeights(featureWeights QGeoRouteRequest__FeatureWeight) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_SetSupportedFeatureWeights(ptr.Pointer(), C.longlong(featureWeights))
	}
}

func (ptr *QGeoRoutingManagerEngine) SetSupportedManeuverDetails(maneuverDetails QGeoRouteRequest__ManeuverDetail) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_SetSupportedManeuverDetails(ptr.Pointer(), C.longlong(maneuverDetails))
	}
}

func (ptr *QGeoRoutingManagerEngine) SetSupportedRouteOptimizations(optimizations QGeoRouteRequest__RouteOptimization) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_SetSupportedRouteOptimizations(ptr.Pointer(), C.longlong(optimizations))
	}
}

func (ptr *QGeoRoutingManagerEngine) SetSupportedSegmentDetails(segmentDetails QGeoRouteRequest__SegmentDetail) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_SetSupportedSegmentDetails(ptr.Pointer(), C.longlong(segmentDetails))
	}
}

func (ptr *QGeoRoutingManagerEngine) SetSupportedTravelModes(travelModes QGeoRouteRequest__TravelMode) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_SetSupportedTravelModes(ptr.Pointer(), C.longlong(travelModes))
	}
}

func (ptr *QGeoRoutingManagerEngine) SupportedFeatureTypes() QGeoRouteRequest__FeatureType {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__FeatureType(C.QGeoRoutingManagerEngine_SupportedFeatureTypes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManagerEngine) SupportedFeatureWeights() QGeoRouteRequest__FeatureWeight {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__FeatureWeight(C.QGeoRoutingManagerEngine_SupportedFeatureWeights(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManagerEngine) SupportedManeuverDetails() QGeoRouteRequest__ManeuverDetail {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__ManeuverDetail(C.QGeoRoutingManagerEngine_SupportedManeuverDetails(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManagerEngine) SupportedRouteOptimizations() QGeoRouteRequest__RouteOptimization {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__RouteOptimization(C.QGeoRoutingManagerEngine_SupportedRouteOptimizations(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManagerEngine) SupportedSegmentDetails() QGeoRouteRequest__SegmentDetail {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__SegmentDetail(C.QGeoRoutingManagerEngine_SupportedSegmentDetails(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoutingManagerEngine) SupportedTravelModes() QGeoRouteRequest__TravelMode {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__TravelMode(C.QGeoRoutingManagerEngine_SupportedTravelModes(ptr.Pointer()))
	}
	return 0
}

//export callbackQGeoRoutingManagerEngine_UpdateRoute
func callbackQGeoRoutingManagerEngine_UpdateRoute(ptr unsafe.Pointer, route unsafe.Pointer, position unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "updateRoute"); signal != nil {
		return PointerFromQGeoRouteReply((*(*func(*QGeoRoute, *positioning.QGeoCoordinate) *QGeoRouteReply)(signal))(NewQGeoRouteFromPointer(route), positioning.NewQGeoCoordinateFromPointer(position)))
	}

	return PointerFromQGeoRouteReply(NewQGeoRoutingManagerEngineFromPointer(ptr).UpdateRouteDefault(NewQGeoRouteFromPointer(route), positioning.NewQGeoCoordinateFromPointer(position)))
}

func (ptr *QGeoRoutingManagerEngine) ConnectUpdateRoute(f func(route *QGeoRoute, position *positioning.QGeoCoordinate) *QGeoRouteReply) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "updateRoute"); signal != nil {
			f := func(route *QGeoRoute, position *positioning.QGeoCoordinate) *QGeoRouteReply {
				(*(*func(*QGeoRoute, *positioning.QGeoCoordinate) *QGeoRouteReply)(signal))(route, position)
				return f(route, position)
			}
			qt.ConnectSignal(ptr.Pointer(), "updateRoute", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updateRoute", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoRoutingManagerEngine) DisconnectUpdateRoute() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "updateRoute")
	}
}

func (ptr *QGeoRoutingManagerEngine) UpdateRoute(route QGeoRoute_ITF, position positioning.QGeoCoordinate_ITF) *QGeoRouteReply {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoRouteReplyFromPointer(C.QGeoRoutingManagerEngine_UpdateRoute(ptr.Pointer(), PointerFromQGeoRoute(route), positioning.PointerFromQGeoCoordinate(position)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRoutingManagerEngine) UpdateRouteDefault(route QGeoRoute_ITF, position positioning.QGeoCoordinate_ITF) *QGeoRouteReply {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoRouteReplyFromPointer(C.QGeoRoutingManagerEngine_UpdateRouteDefault(ptr.Pointer(), PointerFromQGeoRoute(route), positioning.PointerFromQGeoCoordinate(position)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQGeoRoutingManagerEngine_DestroyQGeoRoutingManagerEngine
func callbackQGeoRoutingManagerEngine_DestroyQGeoRoutingManagerEngine(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGeoRoutingManagerEngine"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGeoRoutingManagerEngineFromPointer(ptr).DestroyQGeoRoutingManagerEngineDefault()
	}
}

func (ptr *QGeoRoutingManagerEngine) ConnectDestroyQGeoRoutingManagerEngine(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGeoRoutingManagerEngine"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QGeoRoutingManagerEngine", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGeoRoutingManagerEngine", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoRoutingManagerEngine) DisconnectDestroyQGeoRoutingManagerEngine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QGeoRoutingManagerEngine")
	}
}

func (ptr *QGeoRoutingManagerEngine) DestroyQGeoRoutingManagerEngine() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoRoutingManagerEngine_DestroyQGeoRoutingManagerEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoRoutingManagerEngine) DestroyQGeoRoutingManagerEngineDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoRoutingManagerEngine_DestroyQGeoRoutingManagerEngineDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoRoutingManagerEngine) __QGeoRoutingManagerEngine_parameters_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QGeoRoutingManagerEngine___QGeoRoutingManagerEngine_parameters_atList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRoutingManagerEngine) __QGeoRoutingManagerEngine_parameters_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QGeoRoutingManagerEngine___QGeoRoutingManagerEngine_parameters_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QGeoRoutingManagerEngine) __QGeoRoutingManagerEngine_parameters_newList() unsafe.Pointer {
	return C.QGeoRoutingManagerEngine___QGeoRoutingManagerEngine_parameters_newList(ptr.Pointer())
}

func (ptr *QGeoRoutingManagerEngine) __QGeoRoutingManagerEngine_parameters_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQGeoRoutingManagerEngineFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____QGeoRoutingManagerEngine_parameters_keyList_atList(i)
			}
			return out
		}(C.QGeoRoutingManagerEngine___QGeoRoutingManagerEngine_parameters_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QGeoRoutingManagerEngine) ____QGeoRoutingManagerEngine_parameters_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoRoutingManagerEngine_____QGeoRoutingManagerEngine_parameters_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QGeoRoutingManagerEngine) ____QGeoRoutingManagerEngine_parameters_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QGeoRoutingManagerEngine_____QGeoRoutingManagerEngine_parameters_keyList_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QGeoRoutingManagerEngine) ____QGeoRoutingManagerEngine_parameters_keyList_newList() unsafe.Pointer {
	return C.QGeoRoutingManagerEngine_____QGeoRoutingManagerEngine_parameters_keyList_newList(ptr.Pointer())
}

func (ptr *QGeoRoutingManagerEngine) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGeoRoutingManagerEngine___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRoutingManagerEngine) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGeoRoutingManagerEngine) __children_newList() unsafe.Pointer {
	return C.QGeoRoutingManagerEngine___children_newList(ptr.Pointer())
}

func (ptr *QGeoRoutingManagerEngine) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QGeoRoutingManagerEngine___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRoutingManagerEngine) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QGeoRoutingManagerEngine) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QGeoRoutingManagerEngine___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QGeoRoutingManagerEngine) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGeoRoutingManagerEngine___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRoutingManagerEngine) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGeoRoutingManagerEngine) __findChildren_newList() unsafe.Pointer {
	return C.QGeoRoutingManagerEngine___findChildren_newList(ptr.Pointer())
}

func (ptr *QGeoRoutingManagerEngine) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGeoRoutingManagerEngine___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRoutingManagerEngine) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGeoRoutingManagerEngine) __findChildren_newList3() unsafe.Pointer {
	return C.QGeoRoutingManagerEngine___findChildren_newList3(ptr.Pointer())
}

//export callbackQGeoRoutingManagerEngine_ChildEvent
func callbackQGeoRoutingManagerEngine_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGeoRoutingManagerEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGeoRoutingManagerEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQGeoRoutingManagerEngine_ConnectNotify
func callbackQGeoRoutingManagerEngine_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGeoRoutingManagerEngineFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGeoRoutingManagerEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGeoRoutingManagerEngine_CustomEvent
func callbackQGeoRoutingManagerEngine_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQGeoRoutingManagerEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGeoRoutingManagerEngine) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGeoRoutingManagerEngine_DeleteLater
func callbackQGeoRoutingManagerEngine_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGeoRoutingManagerEngineFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGeoRoutingManagerEngine) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoRoutingManagerEngine_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQGeoRoutingManagerEngine_Destroyed
func callbackQGeoRoutingManagerEngine_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQGeoRoutingManagerEngine_DisconnectNotify
func callbackQGeoRoutingManagerEngine_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGeoRoutingManagerEngineFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGeoRoutingManagerEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGeoRoutingManagerEngine_Event
func callbackQGeoRoutingManagerEngine_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGeoRoutingManagerEngineFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QGeoRoutingManagerEngine) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoRoutingManagerEngine_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQGeoRoutingManagerEngine_EventFilter
func callbackQGeoRoutingManagerEngine_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGeoRoutingManagerEngineFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGeoRoutingManagerEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoRoutingManagerEngine_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQGeoRoutingManagerEngine_MetaObject
func callbackQGeoRoutingManagerEngine_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQGeoRoutingManagerEngineFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGeoRoutingManagerEngine) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGeoRoutingManagerEngine_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQGeoRoutingManagerEngine_ObjectNameChanged
func callbackQGeoRoutingManagerEngine_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtLocation_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQGeoRoutingManagerEngine_TimerEvent
func callbackQGeoRoutingManagerEngine_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGeoRoutingManagerEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGeoRoutingManagerEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoutingManagerEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QGeoServiceProvider struct {
	core.QObject
}

type QGeoServiceProvider_ITF interface {
	core.QObject_ITF
	QGeoServiceProvider_PTR() *QGeoServiceProvider
}

func (ptr *QGeoServiceProvider) QGeoServiceProvider_PTR() *QGeoServiceProvider {
	return ptr
}

func (ptr *QGeoServiceProvider) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QGeoServiceProvider) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQGeoServiceProvider(ptr QGeoServiceProvider_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoServiceProvider_PTR().Pointer()
	}
	return nil
}

func NewQGeoServiceProviderFromPointer(ptr unsafe.Pointer) (n *QGeoServiceProvider) {
	n = new(QGeoServiceProvider)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QGeoServiceProvider__Error
//QGeoServiceProvider::Error
type QGeoServiceProvider__Error int64

const (
	QGeoServiceProvider__NoError                       QGeoServiceProvider__Error = QGeoServiceProvider__Error(0)
	QGeoServiceProvider__NotSupportedError             QGeoServiceProvider__Error = QGeoServiceProvider__Error(1)
	QGeoServiceProvider__UnknownParameterError         QGeoServiceProvider__Error = QGeoServiceProvider__Error(2)
	QGeoServiceProvider__MissingRequiredParameterError QGeoServiceProvider__Error = QGeoServiceProvider__Error(3)
	QGeoServiceProvider__ConnectionError               QGeoServiceProvider__Error = QGeoServiceProvider__Error(4)
	QGeoServiceProvider__LoaderError                   QGeoServiceProvider__Error = QGeoServiceProvider__Error(5)
)

//go:generate stringer -type=QGeoServiceProvider__RoutingFeature
//QGeoServiceProvider::RoutingFeature
type QGeoServiceProvider__RoutingFeature int64

var (
	QGeoServiceProvider__NoRoutingFeatures          QGeoServiceProvider__RoutingFeature = QGeoServiceProvider__RoutingFeature(0)
	QGeoServiceProvider__OnlineRoutingFeature       QGeoServiceProvider__RoutingFeature = QGeoServiceProvider__RoutingFeature(C.QGeoServiceProvider_OnlineRoutingFeature_Type())
	QGeoServiceProvider__OfflineRoutingFeature      QGeoServiceProvider__RoutingFeature = QGeoServiceProvider__RoutingFeature(C.QGeoServiceProvider_OfflineRoutingFeature_Type())
	QGeoServiceProvider__LocalizedRoutingFeature    QGeoServiceProvider__RoutingFeature = QGeoServiceProvider__RoutingFeature(C.QGeoServiceProvider_LocalizedRoutingFeature_Type())
	QGeoServiceProvider__RouteUpdatesFeature        QGeoServiceProvider__RoutingFeature = QGeoServiceProvider__RoutingFeature(C.QGeoServiceProvider_RouteUpdatesFeature_Type())
	QGeoServiceProvider__AlternativeRoutesFeature   QGeoServiceProvider__RoutingFeature = QGeoServiceProvider__RoutingFeature(C.QGeoServiceProvider_AlternativeRoutesFeature_Type())
	QGeoServiceProvider__ExcludeAreasRoutingFeature QGeoServiceProvider__RoutingFeature = QGeoServiceProvider__RoutingFeature(C.QGeoServiceProvider_ExcludeAreasRoutingFeature_Type())
	QGeoServiceProvider__AnyRoutingFeatures         QGeoServiceProvider__RoutingFeature = QGeoServiceProvider__RoutingFeature(C.QGeoServiceProvider_AnyRoutingFeatures_Type())
)

//go:generate stringer -type=QGeoServiceProvider__GeocodingFeature
//QGeoServiceProvider::GeocodingFeature
type QGeoServiceProvider__GeocodingFeature int64

var (
	QGeoServiceProvider__NoGeocodingFeatures       QGeoServiceProvider__GeocodingFeature = QGeoServiceProvider__GeocodingFeature(0)
	QGeoServiceProvider__OnlineGeocodingFeature    QGeoServiceProvider__GeocodingFeature = QGeoServiceProvider__GeocodingFeature(C.QGeoServiceProvider_OnlineGeocodingFeature_Type())
	QGeoServiceProvider__OfflineGeocodingFeature   QGeoServiceProvider__GeocodingFeature = QGeoServiceProvider__GeocodingFeature(C.QGeoServiceProvider_OfflineGeocodingFeature_Type())
	QGeoServiceProvider__ReverseGeocodingFeature   QGeoServiceProvider__GeocodingFeature = QGeoServiceProvider__GeocodingFeature(C.QGeoServiceProvider_ReverseGeocodingFeature_Type())
	QGeoServiceProvider__LocalizedGeocodingFeature QGeoServiceProvider__GeocodingFeature = QGeoServiceProvider__GeocodingFeature(C.QGeoServiceProvider_LocalizedGeocodingFeature_Type())
	QGeoServiceProvider__AnyGeocodingFeatures      QGeoServiceProvider__GeocodingFeature = QGeoServiceProvider__GeocodingFeature(C.QGeoServiceProvider_AnyGeocodingFeatures_Type())
)

//go:generate stringer -type=QGeoServiceProvider__MappingFeature
//QGeoServiceProvider::MappingFeature
type QGeoServiceProvider__MappingFeature int64

var (
	QGeoServiceProvider__NoMappingFeatures       QGeoServiceProvider__MappingFeature = QGeoServiceProvider__MappingFeature(0)
	QGeoServiceProvider__OnlineMappingFeature    QGeoServiceProvider__MappingFeature = QGeoServiceProvider__MappingFeature(C.QGeoServiceProvider_OnlineMappingFeature_Type())
	QGeoServiceProvider__OfflineMappingFeature   QGeoServiceProvider__MappingFeature = QGeoServiceProvider__MappingFeature(C.QGeoServiceProvider_OfflineMappingFeature_Type())
	QGeoServiceProvider__LocalizedMappingFeature QGeoServiceProvider__MappingFeature = QGeoServiceProvider__MappingFeature(C.QGeoServiceProvider_LocalizedMappingFeature_Type())
	QGeoServiceProvider__AnyMappingFeatures      QGeoServiceProvider__MappingFeature = QGeoServiceProvider__MappingFeature(C.QGeoServiceProvider_AnyMappingFeatures_Type())
)

//go:generate stringer -type=QGeoServiceProvider__PlacesFeature
//QGeoServiceProvider::PlacesFeature
type QGeoServiceProvider__PlacesFeature int64

var (
	QGeoServiceProvider__NoPlacesFeatures            QGeoServiceProvider__PlacesFeature = QGeoServiceProvider__PlacesFeature(0)
	QGeoServiceProvider__OnlinePlacesFeature         QGeoServiceProvider__PlacesFeature = QGeoServiceProvider__PlacesFeature(C.QGeoServiceProvider_OnlinePlacesFeature_Type())
	QGeoServiceProvider__OfflinePlacesFeature        QGeoServiceProvider__PlacesFeature = QGeoServiceProvider__PlacesFeature(C.QGeoServiceProvider_OfflinePlacesFeature_Type())
	QGeoServiceProvider__SavePlaceFeature            QGeoServiceProvider__PlacesFeature = QGeoServiceProvider__PlacesFeature(C.QGeoServiceProvider_SavePlaceFeature_Type())
	QGeoServiceProvider__RemovePlaceFeature          QGeoServiceProvider__PlacesFeature = QGeoServiceProvider__PlacesFeature(C.QGeoServiceProvider_RemovePlaceFeature_Type())
	QGeoServiceProvider__SaveCategoryFeature         QGeoServiceProvider__PlacesFeature = QGeoServiceProvider__PlacesFeature(C.QGeoServiceProvider_SaveCategoryFeature_Type())
	QGeoServiceProvider__RemoveCategoryFeature       QGeoServiceProvider__PlacesFeature = QGeoServiceProvider__PlacesFeature(C.QGeoServiceProvider_RemoveCategoryFeature_Type())
	QGeoServiceProvider__PlaceRecommendationsFeature QGeoServiceProvider__PlacesFeature = QGeoServiceProvider__PlacesFeature(C.QGeoServiceProvider_PlaceRecommendationsFeature_Type())
	QGeoServiceProvider__SearchSuggestionsFeature    QGeoServiceProvider__PlacesFeature = QGeoServiceProvider__PlacesFeature(C.QGeoServiceProvider_SearchSuggestionsFeature_Type())
	QGeoServiceProvider__LocalizedPlacesFeature      QGeoServiceProvider__PlacesFeature = QGeoServiceProvider__PlacesFeature(C.QGeoServiceProvider_LocalizedPlacesFeature_Type())
	QGeoServiceProvider__NotificationsFeature        QGeoServiceProvider__PlacesFeature = QGeoServiceProvider__PlacesFeature(C.QGeoServiceProvider_NotificationsFeature_Type())
	QGeoServiceProvider__PlaceMatchingFeature        QGeoServiceProvider__PlacesFeature = QGeoServiceProvider__PlacesFeature(C.QGeoServiceProvider_PlaceMatchingFeature_Type())
	QGeoServiceProvider__AnyPlacesFeatures           QGeoServiceProvider__PlacesFeature = QGeoServiceProvider__PlacesFeature(C.QGeoServiceProvider_AnyPlacesFeatures_Type())
)

//go:generate stringer -type=QGeoServiceProvider__NavigationFeature
//QGeoServiceProvider::NavigationFeature
type QGeoServiceProvider__NavigationFeature int64

var (
	QGeoServiceProvider__NoNavigationFeatures     QGeoServiceProvider__NavigationFeature = QGeoServiceProvider__NavigationFeature(0)
	QGeoServiceProvider__OnlineNavigationFeature  QGeoServiceProvider__NavigationFeature = QGeoServiceProvider__NavigationFeature(C.QGeoServiceProvider_OnlineNavigationFeature_Type())
	QGeoServiceProvider__OfflineNavigationFeature QGeoServiceProvider__NavigationFeature = QGeoServiceProvider__NavigationFeature(C.QGeoServiceProvider_OfflineNavigationFeature_Type())
	QGeoServiceProvider__AnyNavigationFeatures    QGeoServiceProvider__NavigationFeature = QGeoServiceProvider__NavigationFeature(C.QGeoServiceProvider_AnyNavigationFeatures_Type())
)

func NewQGeoServiceProvider(providerName string, parameters map[string]*core.QVariant, allowExperimental bool) *QGeoServiceProvider {
	var providerNameC *C.char
	if providerName != "" {
		providerNameC = C.CString(providerName)
		defer C.free(unsafe.Pointer(providerNameC))
	}
	tmpValue := NewQGeoServiceProviderFromPointer(C.QGeoServiceProvider_NewQGeoServiceProvider(C.struct_QtLocation_PackedString{data: providerNameC, len: C.longlong(len(providerName))}, func() unsafe.Pointer {
		tmpList := NewQGeoServiceProviderFromPointer(NewQGeoServiceProviderFromPointer(nil).__QGeoServiceProvider_parameters_newList())
		for k, v := range parameters {
			tmpList.__QGeoServiceProvider_parameters_setList(k, v)
		}
		return tmpList.Pointer()
	}(), C.char(int8(qt.GoBoolToInt(allowExperimental)))))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QGeoServiceProvider_AvailableServiceProviders() []string {
	return unpackStringList(cGoUnpackString(C.QGeoServiceProvider_QGeoServiceProvider_AvailableServiceProviders()))
}

func (ptr *QGeoServiceProvider) AvailableServiceProviders() []string {
	return unpackStringList(cGoUnpackString(C.QGeoServiceProvider_QGeoServiceProvider_AvailableServiceProviders()))
}

func (ptr *QGeoServiceProvider) Error() QGeoServiceProvider__Error {
	if ptr.Pointer() != nil {
		return QGeoServiceProvider__Error(C.QGeoServiceProvider_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoServiceProvider) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoServiceProvider_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoServiceProvider) GeocodingError() QGeoServiceProvider__Error {
	if ptr.Pointer() != nil {
		return QGeoServiceProvider__Error(C.QGeoServiceProvider_GeocodingError(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoServiceProvider) GeocodingErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoServiceProvider_GeocodingErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoServiceProvider) GeocodingFeatures() QGeoServiceProvider__GeocodingFeature {
	if ptr.Pointer() != nil {
		return QGeoServiceProvider__GeocodingFeature(C.QGeoServiceProvider_GeocodingFeatures(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoServiceProvider) GeocodingManager() *QGeoCodingManager {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoCodingManagerFromPointer(C.QGeoServiceProvider_GeocodingManager(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoServiceProvider) MappingError() QGeoServiceProvider__Error {
	if ptr.Pointer() != nil {
		return QGeoServiceProvider__Error(C.QGeoServiceProvider_MappingError(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoServiceProvider) MappingErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoServiceProvider_MappingErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoServiceProvider) MappingFeatures() QGeoServiceProvider__MappingFeature {
	if ptr.Pointer() != nil {
		return QGeoServiceProvider__MappingFeature(C.QGeoServiceProvider_MappingFeatures(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoServiceProvider) NavigationError() QGeoServiceProvider__Error {
	if ptr.Pointer() != nil {
		return QGeoServiceProvider__Error(C.QGeoServiceProvider_NavigationError(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoServiceProvider) NavigationErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoServiceProvider_NavigationErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoServiceProvider) NavigationFeatures() QGeoServiceProvider__NavigationFeature {
	if ptr.Pointer() != nil {
		return QGeoServiceProvider__NavigationFeature(C.QGeoServiceProvider_NavigationFeatures(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoServiceProvider) PlaceManager() *QPlaceManager {
	if ptr.Pointer() != nil {
		tmpValue := NewQPlaceManagerFromPointer(C.QGeoServiceProvider_PlaceManager(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoServiceProvider) PlacesError() QGeoServiceProvider__Error {
	if ptr.Pointer() != nil {
		return QGeoServiceProvider__Error(C.QGeoServiceProvider_PlacesError(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoServiceProvider) PlacesErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoServiceProvider_PlacesErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoServiceProvider) PlacesFeatures() QGeoServiceProvider__PlacesFeature {
	if ptr.Pointer() != nil {
		return QGeoServiceProvider__PlacesFeature(C.QGeoServiceProvider_PlacesFeatures(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoServiceProvider) RoutingError() QGeoServiceProvider__Error {
	if ptr.Pointer() != nil {
		return QGeoServiceProvider__Error(C.QGeoServiceProvider_RoutingError(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoServiceProvider) RoutingErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoServiceProvider_RoutingErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoServiceProvider) RoutingFeatures() QGeoServiceProvider__RoutingFeature {
	if ptr.Pointer() != nil {
		return QGeoServiceProvider__RoutingFeature(C.QGeoServiceProvider_RoutingFeatures(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoServiceProvider) RoutingManager() *QGeoRoutingManager {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoRoutingManagerFromPointer(C.QGeoServiceProvider_RoutingManager(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoServiceProvider) SetAllowExperimental(allow bool) {
	if ptr.Pointer() != nil {
		C.QGeoServiceProvider_SetAllowExperimental(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(allow))))
	}
}

func (ptr *QGeoServiceProvider) SetLocale(locale core.QLocale_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoServiceProvider_SetLocale(ptr.Pointer(), core.PointerFromQLocale(locale))
	}
}

func (ptr *QGeoServiceProvider) SetParameters(parameters map[string]*core.QVariant) {
	if ptr.Pointer() != nil {
		C.QGeoServiceProvider_SetParameters(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQGeoServiceProviderFromPointer(NewQGeoServiceProviderFromPointer(nil).__setParameters_parameters_newList())
			for k, v := range parameters {
				tmpList.__setParameters_parameters_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQGeoServiceProvider_DestroyQGeoServiceProvider
func callbackQGeoServiceProvider_DestroyQGeoServiceProvider(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGeoServiceProvider"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGeoServiceProviderFromPointer(ptr).DestroyQGeoServiceProviderDefault()
	}
}

func (ptr *QGeoServiceProvider) ConnectDestroyQGeoServiceProvider(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGeoServiceProvider"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QGeoServiceProvider", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGeoServiceProvider", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoServiceProvider) DisconnectDestroyQGeoServiceProvider() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QGeoServiceProvider")
	}
}

func (ptr *QGeoServiceProvider) DestroyQGeoServiceProvider() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoServiceProvider_DestroyQGeoServiceProvider(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoServiceProvider) DestroyQGeoServiceProviderDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoServiceProvider_DestroyQGeoServiceProviderDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoServiceProvider) __QGeoServiceProvider_parameters_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QGeoServiceProvider___QGeoServiceProvider_parameters_atList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoServiceProvider) __QGeoServiceProvider_parameters_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QGeoServiceProvider___QGeoServiceProvider_parameters_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QGeoServiceProvider) __QGeoServiceProvider_parameters_newList() unsafe.Pointer {
	return C.QGeoServiceProvider___QGeoServiceProvider_parameters_newList(ptr.Pointer())
}

func (ptr *QGeoServiceProvider) __QGeoServiceProvider_parameters_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQGeoServiceProviderFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____QGeoServiceProvider_parameters_keyList_atList(i)
			}
			return out
		}(C.QGeoServiceProvider___QGeoServiceProvider_parameters_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QGeoServiceProvider) __setParameters_parameters_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QGeoServiceProvider___setParameters_parameters_atList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoServiceProvider) __setParameters_parameters_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QGeoServiceProvider___setParameters_parameters_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QGeoServiceProvider) __setParameters_parameters_newList() unsafe.Pointer {
	return C.QGeoServiceProvider___setParameters_parameters_newList(ptr.Pointer())
}

func (ptr *QGeoServiceProvider) __setParameters_parameters_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQGeoServiceProviderFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setParameters_parameters_keyList_atList(i)
			}
			return out
		}(C.QGeoServiceProvider___setParameters_parameters_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QGeoServiceProvider) ____QGeoServiceProvider_parameters_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoServiceProvider_____QGeoServiceProvider_parameters_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QGeoServiceProvider) ____QGeoServiceProvider_parameters_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QGeoServiceProvider_____QGeoServiceProvider_parameters_keyList_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QGeoServiceProvider) ____QGeoServiceProvider_parameters_keyList_newList() unsafe.Pointer {
	return C.QGeoServiceProvider_____QGeoServiceProvider_parameters_keyList_newList(ptr.Pointer())
}

func (ptr *QGeoServiceProvider) ____setParameters_parameters_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoServiceProvider_____setParameters_parameters_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QGeoServiceProvider) ____setParameters_parameters_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QGeoServiceProvider_____setParameters_parameters_keyList_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QGeoServiceProvider) ____setParameters_parameters_keyList_newList() unsafe.Pointer {
	return C.QGeoServiceProvider_____setParameters_parameters_keyList_newList(ptr.Pointer())
}

func (ptr *QGeoServiceProvider) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGeoServiceProvider___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoServiceProvider) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoServiceProvider___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGeoServiceProvider) __children_newList() unsafe.Pointer {
	return C.QGeoServiceProvider___children_newList(ptr.Pointer())
}

func (ptr *QGeoServiceProvider) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QGeoServiceProvider___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoServiceProvider) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoServiceProvider___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QGeoServiceProvider) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QGeoServiceProvider___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QGeoServiceProvider) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGeoServiceProvider___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoServiceProvider) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoServiceProvider___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGeoServiceProvider) __findChildren_newList() unsafe.Pointer {
	return C.QGeoServiceProvider___findChildren_newList(ptr.Pointer())
}

func (ptr *QGeoServiceProvider) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGeoServiceProvider___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoServiceProvider) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoServiceProvider___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGeoServiceProvider) __findChildren_newList3() unsafe.Pointer {
	return C.QGeoServiceProvider___findChildren_newList3(ptr.Pointer())
}

//export callbackQGeoServiceProvider_ChildEvent
func callbackQGeoServiceProvider_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGeoServiceProviderFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGeoServiceProvider) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoServiceProvider_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQGeoServiceProvider_ConnectNotify
func callbackQGeoServiceProvider_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGeoServiceProviderFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGeoServiceProvider) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoServiceProvider_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGeoServiceProvider_CustomEvent
func callbackQGeoServiceProvider_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQGeoServiceProviderFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGeoServiceProvider) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoServiceProvider_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGeoServiceProvider_DeleteLater
func callbackQGeoServiceProvider_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGeoServiceProviderFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGeoServiceProvider) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoServiceProvider_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQGeoServiceProvider_Destroyed
func callbackQGeoServiceProvider_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQGeoServiceProvider_DisconnectNotify
func callbackQGeoServiceProvider_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGeoServiceProviderFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGeoServiceProvider) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoServiceProvider_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGeoServiceProvider_Event
func callbackQGeoServiceProvider_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGeoServiceProviderFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QGeoServiceProvider) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoServiceProvider_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQGeoServiceProvider_EventFilter
func callbackQGeoServiceProvider_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGeoServiceProviderFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGeoServiceProvider) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoServiceProvider_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQGeoServiceProvider_MetaObject
func callbackQGeoServiceProvider_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQGeoServiceProviderFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGeoServiceProvider) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGeoServiceProvider_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQGeoServiceProvider_ObjectNameChanged
func callbackQGeoServiceProvider_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtLocation_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQGeoServiceProvider_TimerEvent
func callbackQGeoServiceProvider_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGeoServiceProviderFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGeoServiceProvider) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoServiceProvider_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QGeoServiceProviderFactory struct {
	ptr unsafe.Pointer
}

type QGeoServiceProviderFactory_ITF interface {
	QGeoServiceProviderFactory_PTR() *QGeoServiceProviderFactory
}

func (ptr *QGeoServiceProviderFactory) QGeoServiceProviderFactory_PTR() *QGeoServiceProviderFactory {
	return ptr
}

func (ptr *QGeoServiceProviderFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QGeoServiceProviderFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQGeoServiceProviderFactory(ptr QGeoServiceProviderFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoServiceProviderFactory_PTR().Pointer()
	}
	return nil
}

func NewQGeoServiceProviderFactoryFromPointer(ptr unsafe.Pointer) (n *QGeoServiceProviderFactory) {
	n = new(QGeoServiceProviderFactory)
	n.SetPointer(ptr)
	return
}

//export callbackQGeoServiceProviderFactory_DestroyQGeoServiceProviderFactory
func callbackQGeoServiceProviderFactory_DestroyQGeoServiceProviderFactory(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGeoServiceProviderFactory"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGeoServiceProviderFactoryFromPointer(ptr).DestroyQGeoServiceProviderFactoryDefault()
	}
}

func (ptr *QGeoServiceProviderFactory) ConnectDestroyQGeoServiceProviderFactory(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGeoServiceProviderFactory"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QGeoServiceProviderFactory", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGeoServiceProviderFactory", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoServiceProviderFactory) DisconnectDestroyQGeoServiceProviderFactory() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QGeoServiceProviderFactory")
	}
}

func (ptr *QGeoServiceProviderFactory) DestroyQGeoServiceProviderFactory() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoServiceProviderFactory_DestroyQGeoServiceProviderFactory(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoServiceProviderFactory) DestroyQGeoServiceProviderFactoryDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoServiceProviderFactory_DestroyQGeoServiceProviderFactoryDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoServiceProviderFactory) __createGeocodingManagerEngine_parameters_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QGeoServiceProviderFactory___createGeocodingManagerEngine_parameters_atList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoServiceProviderFactory) __createGeocodingManagerEngine_parameters_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QGeoServiceProviderFactory___createGeocodingManagerEngine_parameters_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QGeoServiceProviderFactory) __createGeocodingManagerEngine_parameters_newList() unsafe.Pointer {
	return C.QGeoServiceProviderFactory___createGeocodingManagerEngine_parameters_newList(ptr.Pointer())
}

func (ptr *QGeoServiceProviderFactory) __createGeocodingManagerEngine_parameters_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQGeoServiceProviderFactoryFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____createGeocodingManagerEngine_parameters_keyList_atList(i)
			}
			return out
		}(C.QGeoServiceProviderFactory___createGeocodingManagerEngine_parameters_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QGeoServiceProviderFactory) __createMappingManagerEngine_parameters_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QGeoServiceProviderFactory___createMappingManagerEngine_parameters_atList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoServiceProviderFactory) __createMappingManagerEngine_parameters_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QGeoServiceProviderFactory___createMappingManagerEngine_parameters_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QGeoServiceProviderFactory) __createMappingManagerEngine_parameters_newList() unsafe.Pointer {
	return C.QGeoServiceProviderFactory___createMappingManagerEngine_parameters_newList(ptr.Pointer())
}

func (ptr *QGeoServiceProviderFactory) __createMappingManagerEngine_parameters_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQGeoServiceProviderFactoryFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____createMappingManagerEngine_parameters_keyList_atList(i)
			}
			return out
		}(C.QGeoServiceProviderFactory___createMappingManagerEngine_parameters_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QGeoServiceProviderFactory) __createPlaceManagerEngine_parameters_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QGeoServiceProviderFactory___createPlaceManagerEngine_parameters_atList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoServiceProviderFactory) __createPlaceManagerEngine_parameters_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QGeoServiceProviderFactory___createPlaceManagerEngine_parameters_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QGeoServiceProviderFactory) __createPlaceManagerEngine_parameters_newList() unsafe.Pointer {
	return C.QGeoServiceProviderFactory___createPlaceManagerEngine_parameters_newList(ptr.Pointer())
}

func (ptr *QGeoServiceProviderFactory) __createPlaceManagerEngine_parameters_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQGeoServiceProviderFactoryFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____createPlaceManagerEngine_parameters_keyList_atList(i)
			}
			return out
		}(C.QGeoServiceProviderFactory___createPlaceManagerEngine_parameters_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QGeoServiceProviderFactory) __createRoutingManagerEngine_parameters_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QGeoServiceProviderFactory___createRoutingManagerEngine_parameters_atList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoServiceProviderFactory) __createRoutingManagerEngine_parameters_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QGeoServiceProviderFactory___createRoutingManagerEngine_parameters_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QGeoServiceProviderFactory) __createRoutingManagerEngine_parameters_newList() unsafe.Pointer {
	return C.QGeoServiceProviderFactory___createRoutingManagerEngine_parameters_newList(ptr.Pointer())
}

func (ptr *QGeoServiceProviderFactory) __createRoutingManagerEngine_parameters_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtLocation_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQGeoServiceProviderFactoryFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____createRoutingManagerEngine_parameters_keyList_atList(i)
			}
			return out
		}(C.QGeoServiceProviderFactory___createRoutingManagerEngine_parameters_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QGeoServiceProviderFactory) ____createGeocodingManagerEngine_parameters_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoServiceProviderFactory_____createGeocodingManagerEngine_parameters_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QGeoServiceProviderFactory) ____createGeocodingManagerEngine_parameters_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QGeoServiceProviderFactory_____createGeocodingManagerEngine_parameters_keyList_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QGeoServiceProviderFactory) ____createGeocodingManagerEngine_parameters_keyList_newList() unsafe.Pointer {
	return C.QGeoServiceProviderFactory_____createGeocodingManagerEngine_parameters_keyList_newList(ptr.Pointer())
}

func (ptr *QGeoServiceProviderFactory) ____createMappingManagerEngine_parameters_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoServiceProviderFactory_____createMappingManagerEngine_parameters_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QGeoServiceProviderFactory) ____createMappingManagerEngine_parameters_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QGeoServiceProviderFactory_____createMappingManagerEngine_parameters_keyList_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QGeoServiceProviderFactory) ____createMappingManagerEngine_parameters_keyList_newList() unsafe.Pointer {
	return C.QGeoServiceProviderFactory_____createMappingManagerEngine_parameters_keyList_newList(ptr.Pointer())
}

func (ptr *QGeoServiceProviderFactory) ____createPlaceManagerEngine_parameters_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoServiceProviderFactory_____createPlaceManagerEngine_parameters_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QGeoServiceProviderFactory) ____createPlaceManagerEngine_parameters_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QGeoServiceProviderFactory_____createPlaceManagerEngine_parameters_keyList_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QGeoServiceProviderFactory) ____createPlaceManagerEngine_parameters_keyList_newList() unsafe.Pointer {
	return C.QGeoServiceProviderFactory_____createPlaceManagerEngine_parameters_keyList_newList(ptr.Pointer())
}

func (ptr *QGeoServiceProviderFactory) ____createRoutingManagerEngine_parameters_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoServiceProviderFactory_____createRoutingManagerEngine_parameters_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QGeoServiceProviderFactory) ____createRoutingManagerEngine_parameters_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QGeoServiceProviderFactory_____createRoutingManagerEngine_parameters_keyList_setList(ptr.Pointer(), C.struct_QtLocation_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QGeoServiceProviderFactory) ____createRoutingManagerEngine_parameters_keyList_newList() unsafe.Pointer {
	return C.QGeoServiceProviderFactory_____createRoutingManagerEngine_parameters_keyList_newList(ptr.Pointer())
}

type QGeoServiceProviderFactoryV2 struct {
	QGeoServiceProviderFactory
}

type QGeoServiceProviderFactoryV2_ITF interface {
	QGeoServiceProviderFactory_ITF
	QGeoServiceProviderFactoryV2_PTR() *QGeoServiceProviderFactoryV2
}

func (ptr *QGeoServiceProviderFactoryV2) QGeoServiceProviderFactoryV2_PTR() *QGeoServiceProviderFactoryV2 {
	return ptr
}

func (ptr *QGeoServiceProviderFactoryV2) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoServiceProviderFactory_PTR().Pointer()
	}
	return nil
}

func (ptr *QGeoServiceProviderFactoryV2) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QGeoServiceProviderFactory_PTR().SetPointer(p)
	}
}

func PointerFromQGeoServiceProviderFactoryV2(ptr QGeoServiceProviderFactoryV2_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoServiceProviderFactoryV2_PTR().Pointer()
	}
	return nil
}

func NewQGeoServiceProviderFactoryV2FromPointer(ptr unsafe.Pointer) (n *QGeoServiceProviderFactoryV2) {
	n = new(QGeoServiceProviderFactoryV2)
	n.SetPointer(ptr)
	return
}

type QGeoServiceProviderFactoryV3 struct {
	QGeoServiceProviderFactoryV2
}

type QGeoServiceProviderFactoryV3_ITF interface {
	QGeoServiceProviderFactoryV2_ITF
	QGeoServiceProviderFactoryV3_PTR() *QGeoServiceProviderFactoryV3
}

func (ptr *QGeoServiceProviderFactoryV3) QGeoServiceProviderFactoryV3_PTR() *QGeoServiceProviderFactoryV3 {
	return ptr
}

func (ptr *QGeoServiceProviderFactoryV3) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoServiceProviderFactoryV2_PTR().Pointer()
	}
	return nil
}

func (ptr *QGeoServiceProviderFactoryV3) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QGeoServiceProviderFactoryV2_PTR().SetPointer(p)
	}
}

func PointerFromQGeoServiceProviderFactoryV3(ptr QGeoServiceProviderFactoryV3_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoServiceProviderFactoryV3_PTR().Pointer()
	}
	return nil
}

func NewQGeoServiceProviderFactoryV3FromPointer(ptr unsafe.Pointer) (n *QGeoServiceProviderFactoryV3) {
	n = new(QGeoServiceProviderFactoryV3)
	n.SetPointer(ptr)
	return
}
func (ptr *QGeoServiceProviderFactoryV3) DestroyQGeoServiceProviderFactoryV3() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QLocation struct {
	ptr unsafe.Pointer
}

type QLocation_ITF interface {
	QLocation_PTR() *QLocation
}

func (ptr *QLocation) QLocation_PTR() *QLocation {
	return ptr
}

func (ptr *QLocation) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QLocation) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQLocation(ptr QLocation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLocation_PTR().Pointer()
	}
	return nil
}

func NewQLocationFromPointer(ptr unsafe.Pointer) (n *QLocation) {
	n = new(QLocation)
	n.SetPointer(ptr)
	return
}
func (ptr *QLocation) DestroyQLocation() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QLocation__Visibility
//QLocation::Visibility
type QLocation__Visibility int64

const (
	QLocation__UnspecifiedVisibility QLocation__Visibility = QLocation__Visibility(0x00)
	QLocation__DeviceVisibility      QLocation__Visibility = QLocation__Visibility(0x01)
	QLocation__PrivateVisibility     QLocation__Visibility = QLocation__Visibility(0x02)
	QLocation__PublicVisibility      QLocation__Visibility = QLocation__Visibility(0x04)
)

type QPlace struct {
	ptr unsafe.Pointer
}

type QPlace_ITF interface {
	QPlace_PTR() *QPlace
}

func (ptr *QPlace) QPlace_PTR() *QPlace {
	return ptr
}

func (ptr *QPlace) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPlace) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPlace(ptr QPlace_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlace_PTR().Pointer()
	}
	return nil
}

func NewQPlaceFromPointer(ptr unsafe.Pointer) (n *QPlace) {
	n = new(QPlace)
	n.SetPointer(ptr)
	return
}

type QPlaceAttribute struct {
	ptr unsafe.Pointer
}

type QPlaceAttribute_ITF interface {
	QPlaceAttribute_PTR() *QPlaceAttribute
}

func (ptr *QPlaceAttribute) QPlaceAttribute_PTR() *QPlaceAttribute {
	return ptr
}

func (ptr *QPlaceAttribute) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPlaceAttribute) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPlaceAttribute(ptr QPlaceAttribute_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceAttribute_PTR().Pointer()
	}
	return nil
}

func NewQPlaceAttributeFromPointer(ptr unsafe.Pointer) (n *QPlaceAttribute) {
	n = new(QPlaceAttribute)
	n.SetPointer(ptr)
	return
}

type QPlaceCategory struct {
	ptr unsafe.Pointer
}

type QPlaceCategory_ITF interface {
	QPlaceCategory_PTR() *QPlaceCategory
}

func (ptr *QPlaceCategory) QPlaceCategory_PTR() *QPlaceCategory {
	return ptr
}

func (ptr *QPlaceCategory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPlaceCategory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPlaceCategory(ptr QPlaceCategory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceCategory_PTR().Pointer()
	}
	return nil
}

func NewQPlaceCategoryFromPointer(ptr unsafe.Pointer) (n *QPlaceCategory) {
	n = new(QPlaceCategory)
	n.SetPointer(ptr)
	return
}

type QPlaceContactDetail struct {
	ptr unsafe.Pointer
}

type QPlaceContactDetail_ITF interface {
	QPlaceContactDetail_PTR() *QPlaceContactDetail
}

func (ptr *QPlaceContactDetail) QPlaceContactDetail_PTR() *QPlaceContactDetail {
	return ptr
}

func (ptr *QPlaceContactDetail) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPlaceContactDetail) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPlaceContactDetail(ptr QPlaceContactDetail_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceContactDetail_PTR().Pointer()
	}
	return nil
}

func NewQPlaceContactDetailFromPointer(ptr unsafe.Pointer) (n *QPlaceContactDetail) {
	n = new(QPlaceContactDetail)
	n.SetPointer(ptr)
	return
}

type QPlaceContent struct {
	ptr unsafe.Pointer
}

type QPlaceContent_ITF interface {
	QPlaceContent_PTR() *QPlaceContent
}

func (ptr *QPlaceContent) QPlaceContent_PTR() *QPlaceContent {
	return ptr
}

func (ptr *QPlaceContent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPlaceContent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPlaceContent(ptr QPlaceContent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceContent_PTR().Pointer()
	}
	return nil
}

func NewQPlaceContentFromPointer(ptr unsafe.Pointer) (n *QPlaceContent) {
	n = new(QPlaceContent)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QPlaceContent__Type
//QPlaceContent::Type
type QPlaceContent__Type int64

const (
	QPlaceContent__NoType        QPlaceContent__Type = QPlaceContent__Type(0)
	QPlaceContent__ImageType     QPlaceContent__Type = QPlaceContent__Type(1)
	QPlaceContent__ReviewType    QPlaceContent__Type = QPlaceContent__Type(2)
	QPlaceContent__EditorialType QPlaceContent__Type = QPlaceContent__Type(3)
	QPlaceContent__CustomType    QPlaceContent__Type = QPlaceContent__Type(0x0100)
)

type QPlaceContentReply struct {
	QPlaceReply
}

type QPlaceContentReply_ITF interface {
	QPlaceReply_ITF
	QPlaceContentReply_PTR() *QPlaceContentReply
}

func (ptr *QPlaceContentReply) QPlaceContentReply_PTR() *QPlaceContentReply {
	return ptr
}

func (ptr *QPlaceContentReply) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceReply_PTR().Pointer()
	}
	return nil
}

func (ptr *QPlaceContentReply) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QPlaceReply_PTR().SetPointer(p)
	}
}

func PointerFromQPlaceContentReply(ptr QPlaceContentReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceContentReply_PTR().Pointer()
	}
	return nil
}

func NewQPlaceContentReplyFromPointer(ptr unsafe.Pointer) (n *QPlaceContentReply) {
	n = new(QPlaceContentReply)
	n.SetPointer(ptr)
	return
}

type QPlaceContentRequest struct {
	ptr unsafe.Pointer
}

type QPlaceContentRequest_ITF interface {
	QPlaceContentRequest_PTR() *QPlaceContentRequest
}

func (ptr *QPlaceContentRequest) QPlaceContentRequest_PTR() *QPlaceContentRequest {
	return ptr
}

func (ptr *QPlaceContentRequest) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPlaceContentRequest) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPlaceContentRequest(ptr QPlaceContentRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceContentRequest_PTR().Pointer()
	}
	return nil
}

func NewQPlaceContentRequestFromPointer(ptr unsafe.Pointer) (n *QPlaceContentRequest) {
	n = new(QPlaceContentRequest)
	n.SetPointer(ptr)
	return
}

type QPlaceDetailsReply struct {
	QPlaceReply
}

type QPlaceDetailsReply_ITF interface {
	QPlaceReply_ITF
	QPlaceDetailsReply_PTR() *QPlaceDetailsReply
}

func (ptr *QPlaceDetailsReply) QPlaceDetailsReply_PTR() *QPlaceDetailsReply {
	return ptr
}

func (ptr *QPlaceDetailsReply) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceReply_PTR().Pointer()
	}
	return nil
}

func (ptr *QPlaceDetailsReply) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QPlaceReply_PTR().SetPointer(p)
	}
}

func PointerFromQPlaceDetailsReply(ptr QPlaceDetailsReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceDetailsReply_PTR().Pointer()
	}
	return nil
}

func NewQPlaceDetailsReplyFromPointer(ptr unsafe.Pointer) (n *QPlaceDetailsReply) {
	n = new(QPlaceDetailsReply)
	n.SetPointer(ptr)
	return
}

type QPlaceEditorial struct {
	QPlaceContent
}

type QPlaceEditorial_ITF interface {
	QPlaceContent_ITF
	QPlaceEditorial_PTR() *QPlaceEditorial
}

func (ptr *QPlaceEditorial) QPlaceEditorial_PTR() *QPlaceEditorial {
	return ptr
}

func (ptr *QPlaceEditorial) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceContent_PTR().Pointer()
	}
	return nil
}

func (ptr *QPlaceEditorial) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QPlaceContent_PTR().SetPointer(p)
	}
}

func PointerFromQPlaceEditorial(ptr QPlaceEditorial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceEditorial_PTR().Pointer()
	}
	return nil
}

func NewQPlaceEditorialFromPointer(ptr unsafe.Pointer) (n *QPlaceEditorial) {
	n = new(QPlaceEditorial)
	n.SetPointer(ptr)
	return
}

type QPlaceIcon struct {
	ptr unsafe.Pointer
}

type QPlaceIcon_ITF interface {
	QPlaceIcon_PTR() *QPlaceIcon
}

func (ptr *QPlaceIcon) QPlaceIcon_PTR() *QPlaceIcon {
	return ptr
}

func (ptr *QPlaceIcon) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPlaceIcon) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPlaceIcon(ptr QPlaceIcon_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceIcon_PTR().Pointer()
	}
	return nil
}

func NewQPlaceIconFromPointer(ptr unsafe.Pointer) (n *QPlaceIcon) {
	n = new(QPlaceIcon)
	n.SetPointer(ptr)
	return
}

type QPlaceIdReply struct {
	QPlaceReply
}

type QPlaceIdReply_ITF interface {
	QPlaceReply_ITF
	QPlaceIdReply_PTR() *QPlaceIdReply
}

func (ptr *QPlaceIdReply) QPlaceIdReply_PTR() *QPlaceIdReply {
	return ptr
}

func (ptr *QPlaceIdReply) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceReply_PTR().Pointer()
	}
	return nil
}

func (ptr *QPlaceIdReply) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QPlaceReply_PTR().SetPointer(p)
	}
}

func PointerFromQPlaceIdReply(ptr QPlaceIdReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceIdReply_PTR().Pointer()
	}
	return nil
}

func NewQPlaceIdReplyFromPointer(ptr unsafe.Pointer) (n *QPlaceIdReply) {
	n = new(QPlaceIdReply)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QPlaceIdReply__OperationType
//QPlaceIdReply::OperationType
type QPlaceIdReply__OperationType int64

const (
	QPlaceIdReply__SavePlace      QPlaceIdReply__OperationType = QPlaceIdReply__OperationType(0)
	QPlaceIdReply__SaveCategory   QPlaceIdReply__OperationType = QPlaceIdReply__OperationType(1)
	QPlaceIdReply__RemovePlace    QPlaceIdReply__OperationType = QPlaceIdReply__OperationType(2)
	QPlaceIdReply__RemoveCategory QPlaceIdReply__OperationType = QPlaceIdReply__OperationType(3)
)

type QPlaceImage struct {
	QPlaceContent
}

type QPlaceImage_ITF interface {
	QPlaceContent_ITF
	QPlaceImage_PTR() *QPlaceImage
}

func (ptr *QPlaceImage) QPlaceImage_PTR() *QPlaceImage {
	return ptr
}

func (ptr *QPlaceImage) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceContent_PTR().Pointer()
	}
	return nil
}

func (ptr *QPlaceImage) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QPlaceContent_PTR().SetPointer(p)
	}
}

func PointerFromQPlaceImage(ptr QPlaceImage_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceImage_PTR().Pointer()
	}
	return nil
}

func NewQPlaceImageFromPointer(ptr unsafe.Pointer) (n *QPlaceImage) {
	n = new(QPlaceImage)
	n.SetPointer(ptr)
	return
}

type QPlaceManager struct {
	core.QObject
}

type QPlaceManager_ITF interface {
	core.QObject_ITF
	QPlaceManager_PTR() *QPlaceManager
}

func (ptr *QPlaceManager) QPlaceManager_PTR() *QPlaceManager {
	return ptr
}

func (ptr *QPlaceManager) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QPlaceManager) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQPlaceManager(ptr QPlaceManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceManager_PTR().Pointer()
	}
	return nil
}

func NewQPlaceManagerFromPointer(ptr unsafe.Pointer) (n *QPlaceManager) {
	n = new(QPlaceManager)
	n.SetPointer(ptr)
	return
}

type QPlaceManagerEngine struct {
	core.QObject
}

type QPlaceManagerEngine_ITF interface {
	core.QObject_ITF
	QPlaceManagerEngine_PTR() *QPlaceManagerEngine
}

func (ptr *QPlaceManagerEngine) QPlaceManagerEngine_PTR() *QPlaceManagerEngine {
	return ptr
}

func (ptr *QPlaceManagerEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QPlaceManagerEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQPlaceManagerEngine(ptr QPlaceManagerEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceManagerEngine_PTR().Pointer()
	}
	return nil
}

func NewQPlaceManagerEngineFromPointer(ptr unsafe.Pointer) (n *QPlaceManagerEngine) {
	n = new(QPlaceManagerEngine)
	n.SetPointer(ptr)
	return
}

type QPlaceMatchReply struct {
	QPlaceReply
}

type QPlaceMatchReply_ITF interface {
	QPlaceReply_ITF
	QPlaceMatchReply_PTR() *QPlaceMatchReply
}

func (ptr *QPlaceMatchReply) QPlaceMatchReply_PTR() *QPlaceMatchReply {
	return ptr
}

func (ptr *QPlaceMatchReply) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceReply_PTR().Pointer()
	}
	return nil
}

func (ptr *QPlaceMatchReply) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QPlaceReply_PTR().SetPointer(p)
	}
}

func PointerFromQPlaceMatchReply(ptr QPlaceMatchReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceMatchReply_PTR().Pointer()
	}
	return nil
}

func NewQPlaceMatchReplyFromPointer(ptr unsafe.Pointer) (n *QPlaceMatchReply) {
	n = new(QPlaceMatchReply)
	n.SetPointer(ptr)
	return
}

type QPlaceMatchRequest struct {
	ptr unsafe.Pointer
}

type QPlaceMatchRequest_ITF interface {
	QPlaceMatchRequest_PTR() *QPlaceMatchRequest
}

func (ptr *QPlaceMatchRequest) QPlaceMatchRequest_PTR() *QPlaceMatchRequest {
	return ptr
}

func (ptr *QPlaceMatchRequest) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPlaceMatchRequest) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPlaceMatchRequest(ptr QPlaceMatchRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceMatchRequest_PTR().Pointer()
	}
	return nil
}

func NewQPlaceMatchRequestFromPointer(ptr unsafe.Pointer) (n *QPlaceMatchRequest) {
	n = new(QPlaceMatchRequest)
	n.SetPointer(ptr)
	return
}

type QPlaceProposedSearchResult struct {
	QPlaceSearchResult
}

type QPlaceProposedSearchResult_ITF interface {
	QPlaceSearchResult_ITF
	QPlaceProposedSearchResult_PTR() *QPlaceProposedSearchResult
}

func (ptr *QPlaceProposedSearchResult) QPlaceProposedSearchResult_PTR() *QPlaceProposedSearchResult {
	return ptr
}

func (ptr *QPlaceProposedSearchResult) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceSearchResult_PTR().Pointer()
	}
	return nil
}

func (ptr *QPlaceProposedSearchResult) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QPlaceSearchResult_PTR().SetPointer(p)
	}
}

func PointerFromQPlaceProposedSearchResult(ptr QPlaceProposedSearchResult_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceProposedSearchResult_PTR().Pointer()
	}
	return nil
}

func NewQPlaceProposedSearchResultFromPointer(ptr unsafe.Pointer) (n *QPlaceProposedSearchResult) {
	n = new(QPlaceProposedSearchResult)
	n.SetPointer(ptr)
	return
}

type QPlaceRatings struct {
	ptr unsafe.Pointer
}

type QPlaceRatings_ITF interface {
	QPlaceRatings_PTR() *QPlaceRatings
}

func (ptr *QPlaceRatings) QPlaceRatings_PTR() *QPlaceRatings {
	return ptr
}

func (ptr *QPlaceRatings) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPlaceRatings) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPlaceRatings(ptr QPlaceRatings_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceRatings_PTR().Pointer()
	}
	return nil
}

func NewQPlaceRatingsFromPointer(ptr unsafe.Pointer) (n *QPlaceRatings) {
	n = new(QPlaceRatings)
	n.SetPointer(ptr)
	return
}

type QPlaceReply struct {
	core.QObject
}

type QPlaceReply_ITF interface {
	core.QObject_ITF
	QPlaceReply_PTR() *QPlaceReply
}

func (ptr *QPlaceReply) QPlaceReply_PTR() *QPlaceReply {
	return ptr
}

func (ptr *QPlaceReply) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QPlaceReply) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQPlaceReply(ptr QPlaceReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceReply_PTR().Pointer()
	}
	return nil
}

func NewQPlaceReplyFromPointer(ptr unsafe.Pointer) (n *QPlaceReply) {
	n = new(QPlaceReply)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QPlaceReply__Error
//QPlaceReply::Error
type QPlaceReply__Error int64

const (
	QPlaceReply__NoError                   QPlaceReply__Error = QPlaceReply__Error(0)
	QPlaceReply__PlaceDoesNotExistError    QPlaceReply__Error = QPlaceReply__Error(1)
	QPlaceReply__CategoryDoesNotExistError QPlaceReply__Error = QPlaceReply__Error(2)
	QPlaceReply__CommunicationError        QPlaceReply__Error = QPlaceReply__Error(3)
	QPlaceReply__ParseError                QPlaceReply__Error = QPlaceReply__Error(4)
	QPlaceReply__PermissionsError          QPlaceReply__Error = QPlaceReply__Error(5)
	QPlaceReply__UnsupportedError          QPlaceReply__Error = QPlaceReply__Error(6)
	QPlaceReply__BadArgumentError          QPlaceReply__Error = QPlaceReply__Error(7)
	QPlaceReply__CancelError               QPlaceReply__Error = QPlaceReply__Error(8)
	QPlaceReply__UnknownError              QPlaceReply__Error = QPlaceReply__Error(9)
)

//go:generate stringer -type=QPlaceReply__Type
//QPlaceReply::Type
type QPlaceReply__Type int64

const (
	QPlaceReply__Reply                 QPlaceReply__Type = QPlaceReply__Type(0)
	QPlaceReply__DetailsReply          QPlaceReply__Type = QPlaceReply__Type(1)
	QPlaceReply__SearchReply           QPlaceReply__Type = QPlaceReply__Type(2)
	QPlaceReply__SearchSuggestionReply QPlaceReply__Type = QPlaceReply__Type(3)
	QPlaceReply__ContentReply          QPlaceReply__Type = QPlaceReply__Type(4)
	QPlaceReply__IdReply               QPlaceReply__Type = QPlaceReply__Type(5)
	QPlaceReply__MatchReply            QPlaceReply__Type = QPlaceReply__Type(6)
)

type QPlaceResult struct {
	QPlaceSearchResult
}

type QPlaceResult_ITF interface {
	QPlaceSearchResult_ITF
	QPlaceResult_PTR() *QPlaceResult
}

func (ptr *QPlaceResult) QPlaceResult_PTR() *QPlaceResult {
	return ptr
}

func (ptr *QPlaceResult) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceSearchResult_PTR().Pointer()
	}
	return nil
}

func (ptr *QPlaceResult) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QPlaceSearchResult_PTR().SetPointer(p)
	}
}

func PointerFromQPlaceResult(ptr QPlaceResult_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceResult_PTR().Pointer()
	}
	return nil
}

func NewQPlaceResultFromPointer(ptr unsafe.Pointer) (n *QPlaceResult) {
	n = new(QPlaceResult)
	n.SetPointer(ptr)
	return
}

type QPlaceReview struct {
	QPlaceContent
}

type QPlaceReview_ITF interface {
	QPlaceContent_ITF
	QPlaceReview_PTR() *QPlaceReview
}

func (ptr *QPlaceReview) QPlaceReview_PTR() *QPlaceReview {
	return ptr
}

func (ptr *QPlaceReview) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceContent_PTR().Pointer()
	}
	return nil
}

func (ptr *QPlaceReview) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QPlaceContent_PTR().SetPointer(p)
	}
}

func PointerFromQPlaceReview(ptr QPlaceReview_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceReview_PTR().Pointer()
	}
	return nil
}

func NewQPlaceReviewFromPointer(ptr unsafe.Pointer) (n *QPlaceReview) {
	n = new(QPlaceReview)
	n.SetPointer(ptr)
	return
}

type QPlaceSearchReply struct {
	QPlaceReply
}

type QPlaceSearchReply_ITF interface {
	QPlaceReply_ITF
	QPlaceSearchReply_PTR() *QPlaceSearchReply
}

func (ptr *QPlaceSearchReply) QPlaceSearchReply_PTR() *QPlaceSearchReply {
	return ptr
}

func (ptr *QPlaceSearchReply) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceReply_PTR().Pointer()
	}
	return nil
}

func (ptr *QPlaceSearchReply) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QPlaceReply_PTR().SetPointer(p)
	}
}

func PointerFromQPlaceSearchReply(ptr QPlaceSearchReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceSearchReply_PTR().Pointer()
	}
	return nil
}

func NewQPlaceSearchReplyFromPointer(ptr unsafe.Pointer) (n *QPlaceSearchReply) {
	n = new(QPlaceSearchReply)
	n.SetPointer(ptr)
	return
}

type QPlaceSearchRequest struct {
	ptr unsafe.Pointer
}

type QPlaceSearchRequest_ITF interface {
	QPlaceSearchRequest_PTR() *QPlaceSearchRequest
}

func (ptr *QPlaceSearchRequest) QPlaceSearchRequest_PTR() *QPlaceSearchRequest {
	return ptr
}

func (ptr *QPlaceSearchRequest) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPlaceSearchRequest) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPlaceSearchRequest(ptr QPlaceSearchRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceSearchRequest_PTR().Pointer()
	}
	return nil
}

func NewQPlaceSearchRequestFromPointer(ptr unsafe.Pointer) (n *QPlaceSearchRequest) {
	n = new(QPlaceSearchRequest)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QPlaceSearchRequest__RelevanceHint
//QPlaceSearchRequest::RelevanceHint
type QPlaceSearchRequest__RelevanceHint int64

const (
	QPlaceSearchRequest__UnspecifiedHint      QPlaceSearchRequest__RelevanceHint = QPlaceSearchRequest__RelevanceHint(0)
	QPlaceSearchRequest__DistanceHint         QPlaceSearchRequest__RelevanceHint = QPlaceSearchRequest__RelevanceHint(1)
	QPlaceSearchRequest__LexicalPlaceNameHint QPlaceSearchRequest__RelevanceHint = QPlaceSearchRequest__RelevanceHint(2)
)

type QPlaceSearchResult struct {
	ptr unsafe.Pointer
}

type QPlaceSearchResult_ITF interface {
	QPlaceSearchResult_PTR() *QPlaceSearchResult
}

func (ptr *QPlaceSearchResult) QPlaceSearchResult_PTR() *QPlaceSearchResult {
	return ptr
}

func (ptr *QPlaceSearchResult) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPlaceSearchResult) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPlaceSearchResult(ptr QPlaceSearchResult_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceSearchResult_PTR().Pointer()
	}
	return nil
}

func NewQPlaceSearchResultFromPointer(ptr unsafe.Pointer) (n *QPlaceSearchResult) {
	n = new(QPlaceSearchResult)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QPlaceSearchResult__SearchResultType
//QPlaceSearchResult::SearchResultType
type QPlaceSearchResult__SearchResultType int64

const (
	QPlaceSearchResult__UnknownSearchResult  QPlaceSearchResult__SearchResultType = QPlaceSearchResult__SearchResultType(0)
	QPlaceSearchResult__PlaceResult          QPlaceSearchResult__SearchResultType = QPlaceSearchResult__SearchResultType(1)
	QPlaceSearchResult__ProposedSearchResult QPlaceSearchResult__SearchResultType = QPlaceSearchResult__SearchResultType(2)
)

type QPlaceSearchSuggestionReply struct {
	QPlaceReply
}

type QPlaceSearchSuggestionReply_ITF interface {
	QPlaceReply_ITF
	QPlaceSearchSuggestionReply_PTR() *QPlaceSearchSuggestionReply
}

func (ptr *QPlaceSearchSuggestionReply) QPlaceSearchSuggestionReply_PTR() *QPlaceSearchSuggestionReply {
	return ptr
}

func (ptr *QPlaceSearchSuggestionReply) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceReply_PTR().Pointer()
	}
	return nil
}

func (ptr *QPlaceSearchSuggestionReply) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QPlaceReply_PTR().SetPointer(p)
	}
}

func PointerFromQPlaceSearchSuggestionReply(ptr QPlaceSearchSuggestionReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceSearchSuggestionReply_PTR().Pointer()
	}
	return nil
}

func NewQPlaceSearchSuggestionReplyFromPointer(ptr unsafe.Pointer) (n *QPlaceSearchSuggestionReply) {
	n = new(QPlaceSearchSuggestionReply)
	n.SetPointer(ptr)
	return
}

type QPlaceSupplier struct {
	ptr unsafe.Pointer
}

type QPlaceSupplier_ITF interface {
	QPlaceSupplier_PTR() *QPlaceSupplier
}

func (ptr *QPlaceSupplier) QPlaceSupplier_PTR() *QPlaceSupplier {
	return ptr
}

func (ptr *QPlaceSupplier) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPlaceSupplier) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPlaceSupplier(ptr QPlaceSupplier_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceSupplier_PTR().Pointer()
	}
	return nil
}

func NewQPlaceSupplierFromPointer(ptr unsafe.Pointer) (n *QPlaceSupplier) {
	n = new(QPlaceSupplier)
	n.SetPointer(ptr)
	return
}

type QPlaceUser struct {
	ptr unsafe.Pointer
}

type QPlaceUser_ITF interface {
	QPlaceUser_PTR() *QPlaceUser
}

func (ptr *QPlaceUser) QPlaceUser_PTR() *QPlaceUser {
	return ptr
}

func (ptr *QPlaceUser) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPlaceUser) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPlaceUser(ptr QPlaceUser_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceUser_PTR().Pointer()
	}
	return nil
}

func NewQPlaceUserFromPointer(ptr unsafe.Pointer) (n *QPlaceUser) {
	n = new(QPlaceUser)
	n.SetPointer(ptr)
	return
}
func init() {
	qt.ItfMap["location.QGeoCodeReply_ITF"] = QGeoCodeReply{}
	qt.EnumMap["location.QGeoCodeReply__NoError"] = int64(QGeoCodeReply__NoError)
	qt.EnumMap["location.QGeoCodeReply__EngineNotSetError"] = int64(QGeoCodeReply__EngineNotSetError)
	qt.EnumMap["location.QGeoCodeReply__CommunicationError"] = int64(QGeoCodeReply__CommunicationError)
	qt.EnumMap["location.QGeoCodeReply__ParseError"] = int64(QGeoCodeReply__ParseError)
	qt.EnumMap["location.QGeoCodeReply__UnsupportedOptionError"] = int64(QGeoCodeReply__UnsupportedOptionError)
	qt.EnumMap["location.QGeoCodeReply__CombinationError"] = int64(QGeoCodeReply__CombinationError)
	qt.EnumMap["location.QGeoCodeReply__UnknownError"] = int64(QGeoCodeReply__UnknownError)
	qt.ItfMap["location.QGeoCodingManager_ITF"] = QGeoCodingManager{}
	qt.ItfMap["location.QGeoCodingManagerEngine_ITF"] = QGeoCodingManagerEngine{}
	qt.ItfMap["location.QGeoJson_ITF"] = QGeoJson{}
	qt.ItfMap["location.QGeoManeuver_ITF"] = QGeoManeuver{}
	qt.FuncMap["location.NewQGeoManeuver"] = NewQGeoManeuver
	qt.FuncMap["location.NewQGeoManeuver2"] = NewQGeoManeuver2
	qt.EnumMap["location.QGeoManeuver__NoDirection"] = int64(QGeoManeuver__NoDirection)
	qt.EnumMap["location.QGeoManeuver__DirectionForward"] = int64(QGeoManeuver__DirectionForward)
	qt.EnumMap["location.QGeoManeuver__DirectionBearRight"] = int64(QGeoManeuver__DirectionBearRight)
	qt.EnumMap["location.QGeoManeuver__DirectionLightRight"] = int64(QGeoManeuver__DirectionLightRight)
	qt.EnumMap["location.QGeoManeuver__DirectionRight"] = int64(QGeoManeuver__DirectionRight)
	qt.EnumMap["location.QGeoManeuver__DirectionHardRight"] = int64(QGeoManeuver__DirectionHardRight)
	qt.EnumMap["location.QGeoManeuver__DirectionUTurnRight"] = int64(QGeoManeuver__DirectionUTurnRight)
	qt.EnumMap["location.QGeoManeuver__DirectionUTurnLeft"] = int64(QGeoManeuver__DirectionUTurnLeft)
	qt.EnumMap["location.QGeoManeuver__DirectionHardLeft"] = int64(QGeoManeuver__DirectionHardLeft)
	qt.EnumMap["location.QGeoManeuver__DirectionLeft"] = int64(QGeoManeuver__DirectionLeft)
	qt.EnumMap["location.QGeoManeuver__DirectionLightLeft"] = int64(QGeoManeuver__DirectionLightLeft)
	qt.EnumMap["location.QGeoManeuver__DirectionBearLeft"] = int64(QGeoManeuver__DirectionBearLeft)
	qt.ItfMap["location.QGeoRoute_ITF"] = QGeoRoute{}
	qt.FuncMap["location.NewQGeoRoute"] = NewQGeoRoute
	qt.FuncMap["location.NewQGeoRoute2"] = NewQGeoRoute2
	qt.ItfMap["location.QGeoRouteLeg_ITF"] = QGeoRouteLeg{}
	qt.FuncMap["location.NewQGeoRouteLeg"] = NewQGeoRouteLeg
	qt.FuncMap["location.NewQGeoRouteLeg2"] = NewQGeoRouteLeg2
	qt.ItfMap["location.QGeoRouteReply_ITF"] = QGeoRouteReply{}
	qt.FuncMap["location.NewQGeoRouteReply"] = NewQGeoRouteReply
	qt.FuncMap["location.NewQGeoRouteReply2"] = NewQGeoRouteReply2
	qt.EnumMap["location.QGeoRouteReply__NoError"] = int64(QGeoRouteReply__NoError)
	qt.EnumMap["location.QGeoRouteReply__EngineNotSetError"] = int64(QGeoRouteReply__EngineNotSetError)
	qt.EnumMap["location.QGeoRouteReply__CommunicationError"] = int64(QGeoRouteReply__CommunicationError)
	qt.EnumMap["location.QGeoRouteReply__ParseError"] = int64(QGeoRouteReply__ParseError)
	qt.EnumMap["location.QGeoRouteReply__UnsupportedOptionError"] = int64(QGeoRouteReply__UnsupportedOptionError)
	qt.EnumMap["location.QGeoRouteReply__UnknownError"] = int64(QGeoRouteReply__UnknownError)
	qt.ItfMap["location.QGeoRouteRequest_ITF"] = QGeoRouteRequest{}
	qt.FuncMap["location.NewQGeoRouteRequest"] = NewQGeoRouteRequest
	qt.FuncMap["location.NewQGeoRouteRequest2"] = NewQGeoRouteRequest2
	qt.FuncMap["location.NewQGeoRouteRequest3"] = NewQGeoRouteRequest3
	qt.EnumMap["location.QGeoRouteRequest__CarTravel"] = int64(QGeoRouteRequest__CarTravel)
	qt.EnumMap["location.QGeoRouteRequest__PedestrianTravel"] = int64(QGeoRouteRequest__PedestrianTravel)
	qt.EnumMap["location.QGeoRouteRequest__BicycleTravel"] = int64(QGeoRouteRequest__BicycleTravel)
	qt.EnumMap["location.QGeoRouteRequest__PublicTransitTravel"] = int64(QGeoRouteRequest__PublicTransitTravel)
	qt.EnumMap["location.QGeoRouteRequest__TruckTravel"] = int64(QGeoRouteRequest__TruckTravel)
	qt.EnumMap["location.QGeoRouteRequest__NoFeature"] = int64(QGeoRouteRequest__NoFeature)
	qt.EnumMap["location.QGeoRouteRequest__TollFeature"] = int64(QGeoRouteRequest__TollFeature)
	qt.EnumMap["location.QGeoRouteRequest__HighwayFeature"] = int64(QGeoRouteRequest__HighwayFeature)
	qt.EnumMap["location.QGeoRouteRequest__PublicTransitFeature"] = int64(QGeoRouteRequest__PublicTransitFeature)
	qt.EnumMap["location.QGeoRouteRequest__FerryFeature"] = int64(QGeoRouteRequest__FerryFeature)
	qt.EnumMap["location.QGeoRouteRequest__TunnelFeature"] = int64(QGeoRouteRequest__TunnelFeature)
	qt.EnumMap["location.QGeoRouteRequest__DirtRoadFeature"] = int64(QGeoRouteRequest__DirtRoadFeature)
	qt.EnumMap["location.QGeoRouteRequest__ParksFeature"] = int64(QGeoRouteRequest__ParksFeature)
	qt.EnumMap["location.QGeoRouteRequest__MotorPoolLaneFeature"] = int64(QGeoRouteRequest__MotorPoolLaneFeature)
	qt.EnumMap["location.QGeoRouteRequest__TrafficFeature"] = int64(QGeoRouteRequest__TrafficFeature)
	qt.EnumMap["location.QGeoRouteRequest__NeutralFeatureWeight"] = int64(QGeoRouteRequest__NeutralFeatureWeight)
	qt.EnumMap["location.QGeoRouteRequest__PreferFeatureWeight"] = int64(QGeoRouteRequest__PreferFeatureWeight)
	qt.EnumMap["location.QGeoRouteRequest__RequireFeatureWeight"] = int64(QGeoRouteRequest__RequireFeatureWeight)
	qt.EnumMap["location.QGeoRouteRequest__AvoidFeatureWeight"] = int64(QGeoRouteRequest__AvoidFeatureWeight)
	qt.EnumMap["location.QGeoRouteRequest__DisallowFeatureWeight"] = int64(QGeoRouteRequest__DisallowFeatureWeight)
	qt.EnumMap["location.QGeoRouteRequest__ShortestRoute"] = int64(QGeoRouteRequest__ShortestRoute)
	qt.EnumMap["location.QGeoRouteRequest__FastestRoute"] = int64(QGeoRouteRequest__FastestRoute)
	qt.EnumMap["location.QGeoRouteRequest__MostEconomicRoute"] = int64(QGeoRouteRequest__MostEconomicRoute)
	qt.EnumMap["location.QGeoRouteRequest__MostScenicRoute"] = int64(QGeoRouteRequest__MostScenicRoute)
	qt.EnumMap["location.QGeoRouteRequest__NoSegmentData"] = int64(QGeoRouteRequest__NoSegmentData)
	qt.EnumMap["location.QGeoRouteRequest__BasicSegmentData"] = int64(QGeoRouteRequest__BasicSegmentData)
	qt.EnumMap["location.QGeoRouteRequest__NoManeuvers"] = int64(QGeoRouteRequest__NoManeuvers)
	qt.EnumMap["location.QGeoRouteRequest__BasicManeuvers"] = int64(QGeoRouteRequest__BasicManeuvers)
	qt.ItfMap["location.QGeoRouteSegment_ITF"] = QGeoRouteSegment{}
	qt.FuncMap["location.NewQGeoRouteSegment"] = NewQGeoRouteSegment
	qt.FuncMap["location.NewQGeoRouteSegment2"] = NewQGeoRouteSegment2
	qt.ItfMap["location.QGeoRoutingManager_ITF"] = QGeoRoutingManager{}
	qt.ItfMap["location.QGeoRoutingManagerEngine_ITF"] = QGeoRoutingManagerEngine{}
	qt.FuncMap["location.NewQGeoRoutingManagerEngine"] = NewQGeoRoutingManagerEngine
	qt.ItfMap["location.QGeoServiceProvider_ITF"] = QGeoServiceProvider{}
	qt.FuncMap["location.NewQGeoServiceProvider"] = NewQGeoServiceProvider
	qt.FuncMap["location.QGeoServiceProvider_AvailableServiceProviders"] = QGeoServiceProvider_AvailableServiceProviders
	qt.EnumMap["location.QGeoServiceProvider__NoError"] = int64(QGeoServiceProvider__NoError)
	qt.EnumMap["location.QGeoServiceProvider__NotSupportedError"] = int64(QGeoServiceProvider__NotSupportedError)
	qt.EnumMap["location.QGeoServiceProvider__UnknownParameterError"] = int64(QGeoServiceProvider__UnknownParameterError)
	qt.EnumMap["location.QGeoServiceProvider__MissingRequiredParameterError"] = int64(QGeoServiceProvider__MissingRequiredParameterError)
	qt.EnumMap["location.QGeoServiceProvider__ConnectionError"] = int64(QGeoServiceProvider__ConnectionError)
	qt.EnumMap["location.QGeoServiceProvider__LoaderError"] = int64(QGeoServiceProvider__LoaderError)
	qt.EnumMap["location.QGeoServiceProvider__NoRoutingFeatures"] = int64(QGeoServiceProvider__NoRoutingFeatures)
	qt.EnumMap["location.QGeoServiceProvider__OnlineRoutingFeature"] = int64(QGeoServiceProvider__OnlineRoutingFeature)
	qt.EnumMap["location.QGeoServiceProvider__OfflineRoutingFeature"] = int64(QGeoServiceProvider__OfflineRoutingFeature)
	qt.EnumMap["location.QGeoServiceProvider__LocalizedRoutingFeature"] = int64(QGeoServiceProvider__LocalizedRoutingFeature)
	qt.EnumMap["location.QGeoServiceProvider__RouteUpdatesFeature"] = int64(QGeoServiceProvider__RouteUpdatesFeature)
	qt.EnumMap["location.QGeoServiceProvider__AlternativeRoutesFeature"] = int64(QGeoServiceProvider__AlternativeRoutesFeature)
	qt.EnumMap["location.QGeoServiceProvider__ExcludeAreasRoutingFeature"] = int64(QGeoServiceProvider__ExcludeAreasRoutingFeature)
	qt.EnumMap["location.QGeoServiceProvider__AnyRoutingFeatures"] = int64(QGeoServiceProvider__AnyRoutingFeatures)
	qt.EnumMap["location.QGeoServiceProvider__NoGeocodingFeatures"] = int64(QGeoServiceProvider__NoGeocodingFeatures)
	qt.EnumMap["location.QGeoServiceProvider__OnlineGeocodingFeature"] = int64(QGeoServiceProvider__OnlineGeocodingFeature)
	qt.EnumMap["location.QGeoServiceProvider__OfflineGeocodingFeature"] = int64(QGeoServiceProvider__OfflineGeocodingFeature)
	qt.EnumMap["location.QGeoServiceProvider__ReverseGeocodingFeature"] = int64(QGeoServiceProvider__ReverseGeocodingFeature)
	qt.EnumMap["location.QGeoServiceProvider__LocalizedGeocodingFeature"] = int64(QGeoServiceProvider__LocalizedGeocodingFeature)
	qt.EnumMap["location.QGeoServiceProvider__AnyGeocodingFeatures"] = int64(QGeoServiceProvider__AnyGeocodingFeatures)
	qt.EnumMap["location.QGeoServiceProvider__NoMappingFeatures"] = int64(QGeoServiceProvider__NoMappingFeatures)
	qt.EnumMap["location.QGeoServiceProvider__OnlineMappingFeature"] = int64(QGeoServiceProvider__OnlineMappingFeature)
	qt.EnumMap["location.QGeoServiceProvider__OfflineMappingFeature"] = int64(QGeoServiceProvider__OfflineMappingFeature)
	qt.EnumMap["location.QGeoServiceProvider__LocalizedMappingFeature"] = int64(QGeoServiceProvider__LocalizedMappingFeature)
	qt.EnumMap["location.QGeoServiceProvider__AnyMappingFeatures"] = int64(QGeoServiceProvider__AnyMappingFeatures)
	qt.EnumMap["location.QGeoServiceProvider__NoPlacesFeatures"] = int64(QGeoServiceProvider__NoPlacesFeatures)
	qt.EnumMap["location.QGeoServiceProvider__OnlinePlacesFeature"] = int64(QGeoServiceProvider__OnlinePlacesFeature)
	qt.EnumMap["location.QGeoServiceProvider__OfflinePlacesFeature"] = int64(QGeoServiceProvider__OfflinePlacesFeature)
	qt.EnumMap["location.QGeoServiceProvider__SavePlaceFeature"] = int64(QGeoServiceProvider__SavePlaceFeature)
	qt.EnumMap["location.QGeoServiceProvider__RemovePlaceFeature"] = int64(QGeoServiceProvider__RemovePlaceFeature)
	qt.EnumMap["location.QGeoServiceProvider__SaveCategoryFeature"] = int64(QGeoServiceProvider__SaveCategoryFeature)
	qt.EnumMap["location.QGeoServiceProvider__RemoveCategoryFeature"] = int64(QGeoServiceProvider__RemoveCategoryFeature)
	qt.EnumMap["location.QGeoServiceProvider__PlaceRecommendationsFeature"] = int64(QGeoServiceProvider__PlaceRecommendationsFeature)
	qt.EnumMap["location.QGeoServiceProvider__SearchSuggestionsFeature"] = int64(QGeoServiceProvider__SearchSuggestionsFeature)
	qt.EnumMap["location.QGeoServiceProvider__LocalizedPlacesFeature"] = int64(QGeoServiceProvider__LocalizedPlacesFeature)
	qt.EnumMap["location.QGeoServiceProvider__NotificationsFeature"] = int64(QGeoServiceProvider__NotificationsFeature)
	qt.EnumMap["location.QGeoServiceProvider__PlaceMatchingFeature"] = int64(QGeoServiceProvider__PlaceMatchingFeature)
	qt.EnumMap["location.QGeoServiceProvider__AnyPlacesFeatures"] = int64(QGeoServiceProvider__AnyPlacesFeatures)
	qt.EnumMap["location.QGeoServiceProvider__NoNavigationFeatures"] = int64(QGeoServiceProvider__NoNavigationFeatures)
	qt.EnumMap["location.QGeoServiceProvider__OnlineNavigationFeature"] = int64(QGeoServiceProvider__OnlineNavigationFeature)
	qt.EnumMap["location.QGeoServiceProvider__OfflineNavigationFeature"] = int64(QGeoServiceProvider__OfflineNavigationFeature)
	qt.EnumMap["location.QGeoServiceProvider__AnyNavigationFeatures"] = int64(QGeoServiceProvider__AnyNavigationFeatures)
	qt.ItfMap["location.QGeoServiceProviderFactory_ITF"] = QGeoServiceProviderFactory{}
	qt.ItfMap["location.QGeoServiceProviderFactoryV2_ITF"] = QGeoServiceProviderFactoryV2{}
	qt.ItfMap["location.QGeoServiceProviderFactoryV3_ITF"] = QGeoServiceProviderFactoryV3{}
	qt.ItfMap["location.QLocation_ITF"] = QLocation{}
	qt.EnumMap["location.QLocation__UnspecifiedVisibility"] = int64(QLocation__UnspecifiedVisibility)
	qt.EnumMap["location.QLocation__DeviceVisibility"] = int64(QLocation__DeviceVisibility)
	qt.EnumMap["location.QLocation__PrivateVisibility"] = int64(QLocation__PrivateVisibility)
	qt.EnumMap["location.QLocation__PublicVisibility"] = int64(QLocation__PublicVisibility)
	qt.ItfMap["location.QPlace_ITF"] = QPlace{}
	qt.ItfMap["location.QPlaceAttribute_ITF"] = QPlaceAttribute{}
	qt.ItfMap["location.QPlaceCategory_ITF"] = QPlaceCategory{}
	qt.ItfMap["location.QPlaceContactDetail_ITF"] = QPlaceContactDetail{}
	qt.ItfMap["location.QPlaceContent_ITF"] = QPlaceContent{}
	qt.EnumMap["location.QPlaceContent__NoType"] = int64(QPlaceContent__NoType)
	qt.EnumMap["location.QPlaceContent__ImageType"] = int64(QPlaceContent__ImageType)
	qt.EnumMap["location.QPlaceContent__ReviewType"] = int64(QPlaceContent__ReviewType)
	qt.EnumMap["location.QPlaceContent__EditorialType"] = int64(QPlaceContent__EditorialType)
	qt.EnumMap["location.QPlaceContent__CustomType"] = int64(QPlaceContent__CustomType)
	qt.ItfMap["location.QPlaceContentReply_ITF"] = QPlaceContentReply{}
	qt.ItfMap["location.QPlaceContentRequest_ITF"] = QPlaceContentRequest{}
	qt.ItfMap["location.QPlaceDetailsReply_ITF"] = QPlaceDetailsReply{}
	qt.ItfMap["location.QPlaceEditorial_ITF"] = QPlaceEditorial{}
	qt.ItfMap["location.QPlaceIcon_ITF"] = QPlaceIcon{}
	qt.ItfMap["location.QPlaceIdReply_ITF"] = QPlaceIdReply{}
	qt.EnumMap["location.QPlaceIdReply__SavePlace"] = int64(QPlaceIdReply__SavePlace)
	qt.EnumMap["location.QPlaceIdReply__SaveCategory"] = int64(QPlaceIdReply__SaveCategory)
	qt.EnumMap["location.QPlaceIdReply__RemovePlace"] = int64(QPlaceIdReply__RemovePlace)
	qt.EnumMap["location.QPlaceIdReply__RemoveCategory"] = int64(QPlaceIdReply__RemoveCategory)
	qt.ItfMap["location.QPlaceImage_ITF"] = QPlaceImage{}
	qt.ItfMap["location.QPlaceManager_ITF"] = QPlaceManager{}
	qt.ItfMap["location.QPlaceManagerEngine_ITF"] = QPlaceManagerEngine{}
	qt.ItfMap["location.QPlaceMatchReply_ITF"] = QPlaceMatchReply{}
	qt.ItfMap["location.QPlaceMatchRequest_ITF"] = QPlaceMatchRequest{}
	qt.ItfMap["location.QPlaceProposedSearchResult_ITF"] = QPlaceProposedSearchResult{}
	qt.ItfMap["location.QPlaceRatings_ITF"] = QPlaceRatings{}
	qt.ItfMap["location.QPlaceReply_ITF"] = QPlaceReply{}
	qt.EnumMap["location.QPlaceReply__NoError"] = int64(QPlaceReply__NoError)
	qt.EnumMap["location.QPlaceReply__PlaceDoesNotExistError"] = int64(QPlaceReply__PlaceDoesNotExistError)
	qt.EnumMap["location.QPlaceReply__CategoryDoesNotExistError"] = int64(QPlaceReply__CategoryDoesNotExistError)
	qt.EnumMap["location.QPlaceReply__CommunicationError"] = int64(QPlaceReply__CommunicationError)
	qt.EnumMap["location.QPlaceReply__ParseError"] = int64(QPlaceReply__ParseError)
	qt.EnumMap["location.QPlaceReply__PermissionsError"] = int64(QPlaceReply__PermissionsError)
	qt.EnumMap["location.QPlaceReply__UnsupportedError"] = int64(QPlaceReply__UnsupportedError)
	qt.EnumMap["location.QPlaceReply__BadArgumentError"] = int64(QPlaceReply__BadArgumentError)
	qt.EnumMap["location.QPlaceReply__CancelError"] = int64(QPlaceReply__CancelError)
	qt.EnumMap["location.QPlaceReply__UnknownError"] = int64(QPlaceReply__UnknownError)
	qt.EnumMap["location.QPlaceReply__Reply"] = int64(QPlaceReply__Reply)
	qt.EnumMap["location.QPlaceReply__DetailsReply"] = int64(QPlaceReply__DetailsReply)
	qt.EnumMap["location.QPlaceReply__SearchReply"] = int64(QPlaceReply__SearchReply)
	qt.EnumMap["location.QPlaceReply__SearchSuggestionReply"] = int64(QPlaceReply__SearchSuggestionReply)
	qt.EnumMap["location.QPlaceReply__ContentReply"] = int64(QPlaceReply__ContentReply)
	qt.EnumMap["location.QPlaceReply__IdReply"] = int64(QPlaceReply__IdReply)
	qt.EnumMap["location.QPlaceReply__MatchReply"] = int64(QPlaceReply__MatchReply)
	qt.ItfMap["location.QPlaceResult_ITF"] = QPlaceResult{}
	qt.ItfMap["location.QPlaceReview_ITF"] = QPlaceReview{}
	qt.ItfMap["location.QPlaceSearchReply_ITF"] = QPlaceSearchReply{}
	qt.ItfMap["location.QPlaceSearchRequest_ITF"] = QPlaceSearchRequest{}
	qt.EnumMap["location.QPlaceSearchRequest__UnspecifiedHint"] = int64(QPlaceSearchRequest__UnspecifiedHint)
	qt.EnumMap["location.QPlaceSearchRequest__DistanceHint"] = int64(QPlaceSearchRequest__DistanceHint)
	qt.EnumMap["location.QPlaceSearchRequest__LexicalPlaceNameHint"] = int64(QPlaceSearchRequest__LexicalPlaceNameHint)
	qt.ItfMap["location.QPlaceSearchResult_ITF"] = QPlaceSearchResult{}
	qt.EnumMap["location.QPlaceSearchResult__UnknownSearchResult"] = int64(QPlaceSearchResult__UnknownSearchResult)
	qt.EnumMap["location.QPlaceSearchResult__PlaceResult"] = int64(QPlaceSearchResult__PlaceResult)
	qt.EnumMap["location.QPlaceSearchResult__ProposedSearchResult"] = int64(QPlaceSearchResult__ProposedSearchResult)
	qt.ItfMap["location.QPlaceSearchSuggestionReply_ITF"] = QPlaceSearchSuggestionReply{}
	qt.ItfMap["location.QPlaceSupplier_ITF"] = QPlaceSupplier{}
	qt.ItfMap["location.QPlaceUser_ITF"] = QPlaceUser{}
}
