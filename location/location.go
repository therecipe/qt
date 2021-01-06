// +build !minimal

package location

import (
	"github.com/dev-drprasad/qt/core"
	"github.com/dev-drprasad/qt/internal"
	"github.com/dev-drprasad/qt/positioning"
	"unsafe"
)

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

func (n *QGeoCodeReply) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QGeoCodeReply) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQGeoCodeReplyFromPointer(ptr unsafe.Pointer) (n *QGeoCodeReply) {
	n = new(QGeoCodeReply)
	n.InitFromInternal(uintptr(ptr), "location.QGeoCodeReply")
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

func (n *QGeoCodingManager) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QGeoCodingManager) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQGeoCodingManagerFromPointer(ptr unsafe.Pointer) (n *QGeoCodingManager) {
	n = new(QGeoCodingManager)
	n.InitFromInternal(uintptr(ptr), "location.QGeoCodingManager")
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

func (n *QGeoCodingManagerEngine) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QGeoCodingManagerEngine) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQGeoCodingManagerEngineFromPointer(ptr unsafe.Pointer) (n *QGeoCodingManagerEngine) {
	n = new(QGeoCodingManagerEngine)
	n.InitFromInternal(uintptr(ptr), "location.QGeoCodingManagerEngine")
	return
}

type QGeoJson struct {
	internal.Internal
}

type QGeoJson_ITF interface {
	QGeoJson_PTR() *QGeoJson
}

func (ptr *QGeoJson) QGeoJson_PTR() *QGeoJson {
	return ptr
}

func (ptr *QGeoJson) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QGeoJson) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQGeoJson(ptr QGeoJson_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoJson_PTR().Pointer()
	}
	return nil
}

func (n *QGeoJson) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQGeoJsonFromPointer(ptr unsafe.Pointer) (n *QGeoJson) {
	n = new(QGeoJson)
	n.InitFromInternal(uintptr(ptr), "location.QGeoJson")
	return
}

func (ptr *QGeoJson) DestroyQGeoJson() {
}

type QGeoManeuver struct {
	internal.Internal
}

type QGeoManeuver_ITF interface {
	QGeoManeuver_PTR() *QGeoManeuver
}

func (ptr *QGeoManeuver) QGeoManeuver_PTR() *QGeoManeuver {
	return ptr
}

func (ptr *QGeoManeuver) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QGeoManeuver) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQGeoManeuver(ptr QGeoManeuver_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoManeuver_PTR().Pointer()
	}
	return nil
}

func (n *QGeoManeuver) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQGeoManeuverFromPointer(ptr unsafe.Pointer) (n *QGeoManeuver) {
	n = new(QGeoManeuver)
	n.InitFromInternal(uintptr(ptr), "location.QGeoManeuver")
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

	return internal.CallLocalFunction([]interface{}{"", "", "location.NewQGeoManeuver", ""}).(*QGeoManeuver)
}

func NewQGeoManeuver2(other QGeoManeuver_ITF) *QGeoManeuver {

	return internal.CallLocalFunction([]interface{}{"", "", "location.NewQGeoManeuver2", "", other}).(*QGeoManeuver)
}

func (ptr *QGeoManeuver) Direction() QGeoManeuver__InstructionDirection {

	return QGeoManeuver__InstructionDirection(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Direction"}).(float64))
}

func (ptr *QGeoManeuver) DistanceToNextInstruction() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DistanceToNextInstruction"}).(float64)
}

func (ptr *QGeoManeuver) ExtendedAttributes() map[string]*core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExtendedAttributes"}).(map[string]*core.QVariant)
}

func (ptr *QGeoManeuver) InstructionText() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InstructionText"}).(string)
}

func (ptr *QGeoManeuver) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QGeoManeuver) Position() *positioning.QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Position"}).(*positioning.QGeoCoordinate)
}

func (ptr *QGeoManeuver) SetDirection(direction QGeoManeuver__InstructionDirection) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDirection", direction})
}

func (ptr *QGeoManeuver) SetDistanceToNextInstruction(distance float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDistanceToNextInstruction", distance})
}

func (ptr *QGeoManeuver) SetExtendedAttributes(extendedAttributes map[string]*core.QVariant) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetExtendedAttributes", extendedAttributes})
}

func (ptr *QGeoManeuver) SetInstructionText(instructionText string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetInstructionText", instructionText})
}

func (ptr *QGeoManeuver) SetPosition(position positioning.QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPosition", position})
}

func (ptr *QGeoManeuver) SetTimeToNextInstruction(secs int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTimeToNextInstruction", secs})
}

func (ptr *QGeoManeuver) SetWaypoint(coordinate positioning.QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWaypoint", coordinate})
}

func (ptr *QGeoManeuver) TimeToNextInstruction() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimeToNextInstruction"}).(float64))
}

func (ptr *QGeoManeuver) Waypoint() *positioning.QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Waypoint"}).(*positioning.QGeoCoordinate)
}

func (ptr *QGeoManeuver) DestroyQGeoManeuver() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoManeuver"})
}

func (ptr *QGeoManeuver) __extendedAttributes_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__extendedAttributes_atList", v, i}).(*core.QVariant)
}

func (ptr *QGeoManeuver) __extendedAttributes_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__extendedAttributes_setList", key, i})
}

func (ptr *QGeoManeuver) __extendedAttributes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__extendedAttributes_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoManeuver) __extendedAttributes_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__extendedAttributes_keyList"}).([]string)
}

func (ptr *QGeoManeuver) __setExtendedAttributes_extendedAttributes_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setExtendedAttributes_extendedAttributes_atList", v, i}).(*core.QVariant)
}

func (ptr *QGeoManeuver) __setExtendedAttributes_extendedAttributes_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setExtendedAttributes_extendedAttributes_setList", key, i})
}

func (ptr *QGeoManeuver) __setExtendedAttributes_extendedAttributes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setExtendedAttributes_extendedAttributes_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoManeuver) __setExtendedAttributes_extendedAttributes_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setExtendedAttributes_extendedAttributes_keyList"}).([]string)
}

func (ptr *QGeoManeuver) ____extendedAttributes_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____extendedAttributes_keyList_atList", i}).(string)
}

func (ptr *QGeoManeuver) ____extendedAttributes_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____extendedAttributes_keyList_setList", i})
}

func (ptr *QGeoManeuver) ____extendedAttributes_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____extendedAttributes_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoManeuver) ____setExtendedAttributes_extendedAttributes_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setExtendedAttributes_extendedAttributes_keyList_atList", i}).(string)
}

func (ptr *QGeoManeuver) ____setExtendedAttributes_extendedAttributes_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setExtendedAttributes_extendedAttributes_keyList_setList", i})
}

func (ptr *QGeoManeuver) ____setExtendedAttributes_extendedAttributes_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setExtendedAttributes_extendedAttributes_keyList_newList"}).(unsafe.Pointer)
}

type QGeoRoute struct {
	internal.Internal
}

type QGeoRoute_ITF interface {
	QGeoRoute_PTR() *QGeoRoute
}

func (ptr *QGeoRoute) QGeoRoute_PTR() *QGeoRoute {
	return ptr
}

func (ptr *QGeoRoute) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QGeoRoute) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQGeoRoute(ptr QGeoRoute_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoRoute_PTR().Pointer()
	}
	return nil
}

func (n *QGeoRoute) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQGeoRouteFromPointer(ptr unsafe.Pointer) (n *QGeoRoute) {
	n = new(QGeoRoute)
	n.InitFromInternal(uintptr(ptr), "location.QGeoRoute")
	return
}
func NewQGeoRoute() *QGeoRoute {

	return internal.CallLocalFunction([]interface{}{"", "", "location.NewQGeoRoute", ""}).(*QGeoRoute)
}

func NewQGeoRoute2(other QGeoRoute_ITF) *QGeoRoute {

	return internal.CallLocalFunction([]interface{}{"", "", "location.NewQGeoRoute2", "", other}).(*QGeoRoute)
}

func (ptr *QGeoRoute) Bounds() *positioning.QGeoRectangle {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Bounds"}).(*positioning.QGeoRectangle)
}

func (ptr *QGeoRoute) Distance() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Distance"}).(float64)
}

func (ptr *QGeoRoute) ExtendedAttributes() map[string]*core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExtendedAttributes"}).(map[string]*core.QVariant)
}

func (ptr *QGeoRoute) FirstRouteSegment() *QGeoRouteSegment {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FirstRouteSegment"}).(*QGeoRouteSegment)
}

func (ptr *QGeoRoute) Path() []*positioning.QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Path"}).([]*positioning.QGeoCoordinate)
}

func (ptr *QGeoRoute) Request() *QGeoRouteRequest {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Request"}).(*QGeoRouteRequest)
}

func (ptr *QGeoRoute) RouteId() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RouteId"}).(string)
}

func (ptr *QGeoRoute) RouteLegs() []*QGeoRouteLeg {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RouteLegs"}).([]*QGeoRouteLeg)
}

func (ptr *QGeoRoute) SetBounds(bounds positioning.QGeoRectangle_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBounds", bounds})
}

func (ptr *QGeoRoute) SetDistance(distance float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDistance", distance})
}

func (ptr *QGeoRoute) SetExtendedAttributes(extendedAttributes map[string]*core.QVariant) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetExtendedAttributes", extendedAttributes})
}

func (ptr *QGeoRoute) SetFirstRouteSegment(routeSegment QGeoRouteSegment_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFirstRouteSegment", routeSegment})
}

func (ptr *QGeoRoute) SetPath(path []*positioning.QGeoCoordinate) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPath", path})
}

func (ptr *QGeoRoute) SetRequest(request QGeoRouteRequest_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRequest", request})
}

func (ptr *QGeoRoute) SetRouteId(id string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRouteId", id})
}

func (ptr *QGeoRoute) SetRouteLegs(legs []*QGeoRouteLeg) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRouteLegs", legs})
}

func (ptr *QGeoRoute) SetTravelMode(mode QGeoRouteRequest__TravelMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTravelMode", mode})
}

func (ptr *QGeoRoute) SetTravelTime(secs int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTravelTime", secs})
}

func (ptr *QGeoRoute) TravelMode() QGeoRouteRequest__TravelMode {

	return QGeoRouteRequest__TravelMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TravelMode"}).(float64))
}

func (ptr *QGeoRoute) TravelTime() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TravelTime"}).(float64))
}

func (ptr *QGeoRoute) DestroyQGeoRoute() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoRoute"})
}

func (ptr *QGeoRoute) __extendedAttributes_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__extendedAttributes_atList", v, i}).(*core.QVariant)
}

func (ptr *QGeoRoute) __extendedAttributes_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__extendedAttributes_setList", key, i})
}

func (ptr *QGeoRoute) __extendedAttributes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__extendedAttributes_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRoute) __extendedAttributes_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__extendedAttributes_keyList"}).([]string)
}

func (ptr *QGeoRoute) __path_atList(i int) *positioning.QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__path_atList", i}).(*positioning.QGeoCoordinate)
}

func (ptr *QGeoRoute) __path_setList(i positioning.QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__path_setList", i})
}

func (ptr *QGeoRoute) __path_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__path_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRoute) __routeLegs_atList(i int) *QGeoRouteLeg {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__routeLegs_atList", i}).(*QGeoRouteLeg)
}

func (ptr *QGeoRoute) __routeLegs_setList(i QGeoRouteLeg_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__routeLegs_setList", i})
}

func (ptr *QGeoRoute) __routeLegs_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__routeLegs_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRoute) __setExtendedAttributes_extendedAttributes_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setExtendedAttributes_extendedAttributes_atList", v, i}).(*core.QVariant)
}

func (ptr *QGeoRoute) __setExtendedAttributes_extendedAttributes_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setExtendedAttributes_extendedAttributes_setList", key, i})
}

func (ptr *QGeoRoute) __setExtendedAttributes_extendedAttributes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setExtendedAttributes_extendedAttributes_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRoute) __setExtendedAttributes_extendedAttributes_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setExtendedAttributes_extendedAttributes_keyList"}).([]string)
}

func (ptr *QGeoRoute) __setPath_path_atList(i int) *positioning.QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setPath_path_atList", i}).(*positioning.QGeoCoordinate)
}

func (ptr *QGeoRoute) __setPath_path_setList(i positioning.QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setPath_path_setList", i})
}

func (ptr *QGeoRoute) __setPath_path_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setPath_path_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRoute) __setRouteLegs_legs_atList(i int) *QGeoRouteLeg {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setRouteLegs_legs_atList", i}).(*QGeoRouteLeg)
}

func (ptr *QGeoRoute) __setRouteLegs_legs_setList(i QGeoRouteLeg_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setRouteLegs_legs_setList", i})
}

func (ptr *QGeoRoute) __setRouteLegs_legs_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setRouteLegs_legs_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRoute) ____extendedAttributes_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____extendedAttributes_keyList_atList", i}).(string)
}

func (ptr *QGeoRoute) ____extendedAttributes_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____extendedAttributes_keyList_setList", i})
}

func (ptr *QGeoRoute) ____extendedAttributes_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____extendedAttributes_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRoute) ____setExtendedAttributes_extendedAttributes_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setExtendedAttributes_extendedAttributes_keyList_atList", i}).(string)
}

func (ptr *QGeoRoute) ____setExtendedAttributes_extendedAttributes_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setExtendedAttributes_extendedAttributes_keyList_setList", i})
}

func (ptr *QGeoRoute) ____setExtendedAttributes_extendedAttributes_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setExtendedAttributes_extendedAttributes_keyList_newList"}).(unsafe.Pointer)
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

func (n *QGeoRouteLeg) InitFromInternal(ptr uintptr, name string) {
	n.QGeoRoute_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QGeoRouteLeg) ClassNameInternalF() string {
	return n.QGeoRoute_PTR().ClassNameInternalF()
}

func NewQGeoRouteLegFromPointer(ptr unsafe.Pointer) (n *QGeoRouteLeg) {
	n = new(QGeoRouteLeg)
	n.InitFromInternal(uintptr(ptr), "location.QGeoRouteLeg")
	return
}
func NewQGeoRouteLeg() *QGeoRouteLeg {

	return internal.CallLocalFunction([]interface{}{"", "", "location.NewQGeoRouteLeg", ""}).(*QGeoRouteLeg)
}

func NewQGeoRouteLeg2(other QGeoRouteLeg_ITF) *QGeoRouteLeg {

	return internal.CallLocalFunction([]interface{}{"", "", "location.NewQGeoRouteLeg2", "", other}).(*QGeoRouteLeg)
}

func (ptr *QGeoRouteLeg) LegIndex() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LegIndex"}).(float64))
}

func (ptr *QGeoRouteLeg) OverallRoute() *QGeoRoute {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OverallRoute"}).(*QGeoRoute)
}

func (ptr *QGeoRouteLeg) SetLegIndex(idx int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLegIndex", idx})
}

func (ptr *QGeoRouteLeg) SetOverallRoute(route QGeoRoute_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOverallRoute", route})
}

func (ptr *QGeoRouteLeg) DestroyQGeoRouteLeg() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoRouteLeg"})
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

func (n *QGeoRouteReply) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QGeoRouteReply) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQGeoRouteReplyFromPointer(ptr unsafe.Pointer) (n *QGeoRouteReply) {
	n = new(QGeoRouteReply)
	n.InitFromInternal(uintptr(ptr), "location.QGeoRouteReply")
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

	return internal.CallLocalFunction([]interface{}{"", "", "location.NewQGeoRouteReply", "", error, errorString, parent}).(*QGeoRouteReply)
}

func NewQGeoRouteReply2(request QGeoRouteRequest_ITF, parent core.QObject_ITF) *QGeoRouteReply {

	return internal.CallLocalFunction([]interface{}{"", "", "location.NewQGeoRouteReply2", "", request, parent}).(*QGeoRouteReply)
}

func (ptr *QGeoRouteReply) ConnectAbort(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAbort", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoRouteReply) DisconnectAbort() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAbort"})
}

func (ptr *QGeoRouteReply) Abort() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Abort"})
}

func (ptr *QGeoRouteReply) AbortDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AbortDefault"})
}

func (ptr *QGeoRouteReply) ConnectAborted(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAborted", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoRouteReply) DisconnectAborted() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAborted"})
}

func (ptr *QGeoRouteReply) Aborted() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Aborted"})
}

func (ptr *QGeoRouteReply) AddRoutes(routes []*QGeoRoute) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddRoutes", routes})
}

func (ptr *QGeoRouteReply) Error() QGeoRouteReply__Error {

	return QGeoRouteReply__Error(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QGeoRouteReply) ConnectError2(f func(error QGeoRouteReply__Error, errorString string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectError2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoRouteReply) DisconnectError2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError2"})
}

func (ptr *QGeoRouteReply) Error2(error QGeoRouteReply__Error, errorString string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error2", error, errorString})
}

func (ptr *QGeoRouteReply) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func (ptr *QGeoRouteReply) ConnectFinished(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFinished", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoRouteReply) DisconnectFinished() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFinished"})
}

func (ptr *QGeoRouteReply) Finished() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Finished"})
}

func (ptr *QGeoRouteReply) IsFinished() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsFinished"}).(bool)
}

func (ptr *QGeoRouteReply) Request() *QGeoRouteRequest {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Request"}).(*QGeoRouteRequest)
}

func (ptr *QGeoRouteReply) Routes() []*QGeoRoute {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Routes"}).([]*QGeoRoute)
}

func (ptr *QGeoRouteReply) SetError(error QGeoRouteReply__Error, errorString string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetError", error, errorString})
}

func (ptr *QGeoRouteReply) SetFinished(finished bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFinished", finished})
}

func (ptr *QGeoRouteReply) SetRoutes(routes []*QGeoRoute) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRoutes", routes})
}

func (ptr *QGeoRouteReply) ConnectDestroyQGeoRouteReply(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQGeoRouteReply", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoRouteReply) DisconnectDestroyQGeoRouteReply() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQGeoRouteReply"})
}

func (ptr *QGeoRouteReply) DestroyQGeoRouteReply() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoRouteReply"})
}

func (ptr *QGeoRouteReply) DestroyQGeoRouteReplyDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoRouteReplyDefault"})
}

func (ptr *QGeoRouteReply) __addRoutes_routes_atList(i int) *QGeoRoute {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addRoutes_routes_atList", i}).(*QGeoRoute)
}

func (ptr *QGeoRouteReply) __addRoutes_routes_setList(i QGeoRoute_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addRoutes_routes_setList", i})
}

func (ptr *QGeoRouteReply) __addRoutes_routes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addRoutes_routes_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRouteReply) __routes_atList(i int) *QGeoRoute {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__routes_atList", i}).(*QGeoRoute)
}

func (ptr *QGeoRouteReply) __routes_setList(i QGeoRoute_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__routes_setList", i})
}

func (ptr *QGeoRouteReply) __routes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__routes_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRouteReply) __setRoutes_routes_atList(i int) *QGeoRoute {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setRoutes_routes_atList", i}).(*QGeoRoute)
}

func (ptr *QGeoRouteReply) __setRoutes_routes_setList(i QGeoRoute_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setRoutes_routes_setList", i})
}

func (ptr *QGeoRouteReply) __setRoutes_routes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setRoutes_routes_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRouteReply) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QGeoRouteReply) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QGeoRouteReply) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRouteReply) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QGeoRouteReply) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QGeoRouteReply) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRouteReply) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QGeoRouteReply) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QGeoRouteReply) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRouteReply) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QGeoRouteReply) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QGeoRouteReply) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QGeoRouteReply) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QGeoRouteReply) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QGeoRouteReply) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QGeoRouteReply) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QGeoRouteReply) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QGeoRouteReply) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QGeoRouteReply) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QGeoRouteReply) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QGeoRouteReply) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QGeoRouteRequest struct {
	internal.Internal
}

type QGeoRouteRequest_ITF interface {
	QGeoRouteRequest_PTR() *QGeoRouteRequest
}

func (ptr *QGeoRouteRequest) QGeoRouteRequest_PTR() *QGeoRouteRequest {
	return ptr
}

func (ptr *QGeoRouteRequest) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QGeoRouteRequest) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQGeoRouteRequest(ptr QGeoRouteRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoRouteRequest_PTR().Pointer()
	}
	return nil
}

func (n *QGeoRouteRequest) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQGeoRouteRequestFromPointer(ptr unsafe.Pointer) (n *QGeoRouteRequest) {
	n = new(QGeoRouteRequest)
	n.InitFromInternal(uintptr(ptr), "location.QGeoRouteRequest")
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

	return internal.CallLocalFunction([]interface{}{"", "", "location.NewQGeoRouteRequest", "", waypoints}).(*QGeoRouteRequest)
}

func NewQGeoRouteRequest2(origin positioning.QGeoCoordinate_ITF, destination positioning.QGeoCoordinate_ITF) *QGeoRouteRequest {

	return internal.CallLocalFunction([]interface{}{"", "", "location.NewQGeoRouteRequest2", "", origin, destination}).(*QGeoRouteRequest)
}

func NewQGeoRouteRequest3(other QGeoRouteRequest_ITF) *QGeoRouteRequest {

	return internal.CallLocalFunction([]interface{}{"", "", "location.NewQGeoRouteRequest3", "", other}).(*QGeoRouteRequest)
}

func (ptr *QGeoRouteRequest) DepartureTime() *core.QDateTime {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DepartureTime"}).(*core.QDateTime)
}

func (ptr *QGeoRouteRequest) ExcludeAreas() []*positioning.QGeoRectangle {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExcludeAreas"}).([]*positioning.QGeoRectangle)
}

func (ptr *QGeoRouteRequest) ExtraParameters() map[string]*core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExtraParameters"}).(map[string]*core.QVariant)
}

func (ptr *QGeoRouteRequest) FeatureWeight(featureType QGeoRouteRequest__FeatureType) QGeoRouteRequest__FeatureWeight {

	return QGeoRouteRequest__FeatureWeight(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FeatureWeight", featureType}).(float64))
}

func (ptr *QGeoRouteRequest) ManeuverDetail() QGeoRouteRequest__ManeuverDetail {

	return QGeoRouteRequest__ManeuverDetail(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ManeuverDetail"}).(float64))
}

func (ptr *QGeoRouteRequest) NumberAlternativeRoutes() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NumberAlternativeRoutes"}).(float64))
}

func (ptr *QGeoRouteRequest) RouteOptimization() QGeoRouteRequest__RouteOptimization {

	return QGeoRouteRequest__RouteOptimization(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RouteOptimization"}).(float64))
}

func (ptr *QGeoRouteRequest) SegmentDetail() QGeoRouteRequest__SegmentDetail {

	return QGeoRouteRequest__SegmentDetail(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SegmentDetail"}).(float64))
}

func (ptr *QGeoRouteRequest) SetDepartureTime(departureTime core.QDateTime_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDepartureTime", departureTime})
}

func (ptr *QGeoRouteRequest) SetExcludeAreas(areas []*positioning.QGeoRectangle) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetExcludeAreas", areas})
}

func (ptr *QGeoRouteRequest) SetExtraParameters(extraParameters map[string]*core.QVariant) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetExtraParameters", extraParameters})
}

func (ptr *QGeoRouteRequest) SetFeatureWeight(featureType QGeoRouteRequest__FeatureType, featureWeight QGeoRouteRequest__FeatureWeight) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFeatureWeight", featureType, featureWeight})
}

func (ptr *QGeoRouteRequest) SetManeuverDetail(maneuverDetail QGeoRouteRequest__ManeuverDetail) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetManeuverDetail", maneuverDetail})
}

func (ptr *QGeoRouteRequest) SetNumberAlternativeRoutes(alternatives int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetNumberAlternativeRoutes", alternatives})
}

func (ptr *QGeoRouteRequest) SetRouteOptimization(optimization QGeoRouteRequest__RouteOptimization) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRouteOptimization", optimization})
}

func (ptr *QGeoRouteRequest) SetSegmentDetail(segmentDetail QGeoRouteRequest__SegmentDetail) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSegmentDetail", segmentDetail})
}

func (ptr *QGeoRouteRequest) SetTravelModes(travelModes QGeoRouteRequest__TravelMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTravelModes", travelModes})
}

func (ptr *QGeoRouteRequest) SetWaypoints(waypoints []*positioning.QGeoCoordinate) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWaypoints", waypoints})
}

func (ptr *QGeoRouteRequest) TravelModes() QGeoRouteRequest__TravelMode {

	return QGeoRouteRequest__TravelMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TravelModes"}).(float64))
}

func (ptr *QGeoRouteRequest) Waypoints() []*positioning.QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Waypoints"}).([]*positioning.QGeoCoordinate)
}

func (ptr *QGeoRouteRequest) DestroyQGeoRouteRequest() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoRouteRequest"})
}

func (ptr *QGeoRouteRequest) __QGeoRouteRequest_waypoints_atList(i int) *positioning.QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QGeoRouteRequest_waypoints_atList", i}).(*positioning.QGeoCoordinate)
}

func (ptr *QGeoRouteRequest) __QGeoRouteRequest_waypoints_setList(i positioning.QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QGeoRouteRequest_waypoints_setList", i})
}

func (ptr *QGeoRouteRequest) __QGeoRouteRequest_waypoints_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QGeoRouteRequest_waypoints_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRouteRequest) __excludeAreas_atList(i int) *positioning.QGeoRectangle {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__excludeAreas_atList", i}).(*positioning.QGeoRectangle)
}

func (ptr *QGeoRouteRequest) __excludeAreas_setList(i positioning.QGeoRectangle_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__excludeAreas_setList", i})
}

func (ptr *QGeoRouteRequest) __excludeAreas_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__excludeAreas_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRouteRequest) __extraParameters_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__extraParameters_atList", v, i}).(*core.QVariant)
}

func (ptr *QGeoRouteRequest) __extraParameters_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__extraParameters_setList", key, i})
}

func (ptr *QGeoRouteRequest) __extraParameters_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__extraParameters_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRouteRequest) __extraParameters_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__extraParameters_keyList"}).([]string)
}

func (ptr *QGeoRouteRequest) __setExcludeAreas_areas_atList(i int) *positioning.QGeoRectangle {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setExcludeAreas_areas_atList", i}).(*positioning.QGeoRectangle)
}

func (ptr *QGeoRouteRequest) __setExcludeAreas_areas_setList(i positioning.QGeoRectangle_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setExcludeAreas_areas_setList", i})
}

func (ptr *QGeoRouteRequest) __setExcludeAreas_areas_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setExcludeAreas_areas_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRouteRequest) __setExtraParameters_extraParameters_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setExtraParameters_extraParameters_atList", v, i}).(*core.QVariant)
}

func (ptr *QGeoRouteRequest) __setExtraParameters_extraParameters_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setExtraParameters_extraParameters_setList", key, i})
}

func (ptr *QGeoRouteRequest) __setExtraParameters_extraParameters_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setExtraParameters_extraParameters_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRouteRequest) __setExtraParameters_extraParameters_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setExtraParameters_extraParameters_keyList"}).([]string)
}

func (ptr *QGeoRouteRequest) __setWaypoints_waypoints_atList(i int) *positioning.QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setWaypoints_waypoints_atList", i}).(*positioning.QGeoCoordinate)
}

func (ptr *QGeoRouteRequest) __setWaypoints_waypoints_setList(i positioning.QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setWaypoints_waypoints_setList", i})
}

func (ptr *QGeoRouteRequest) __setWaypoints_waypoints_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setWaypoints_waypoints_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRouteRequest) __setWaypointsMetadata_waypointMetadata_atList(i int) map[string]*core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setWaypointsMetadata_waypointMetadata_atList", i}).(map[string]*core.QVariant)
}

func (ptr *QGeoRouteRequest) __setWaypointsMetadata_waypointMetadata_setList(i map[string]*core.QVariant) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setWaypointsMetadata_waypointMetadata_setList", i})
}

func (ptr *QGeoRouteRequest) __setWaypointsMetadata_waypointMetadata_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setWaypointsMetadata_waypointMetadata_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRouteRequest) __waypoints_atList(i int) *positioning.QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__waypoints_atList", i}).(*positioning.QGeoCoordinate)
}

func (ptr *QGeoRouteRequest) __waypoints_setList(i positioning.QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__waypoints_setList", i})
}

func (ptr *QGeoRouteRequest) __waypoints_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__waypoints_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRouteRequest) __waypointsMetadata_atList(i int) map[string]*core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__waypointsMetadata_atList", i}).(map[string]*core.QVariant)
}

func (ptr *QGeoRouteRequest) __waypointsMetadata_setList(i map[string]*core.QVariant) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__waypointsMetadata_setList", i})
}

func (ptr *QGeoRouteRequest) __waypointsMetadata_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__waypointsMetadata_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRouteRequest) ____extraParameters_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____extraParameters_keyList_atList", i}).(string)
}

func (ptr *QGeoRouteRequest) ____extraParameters_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____extraParameters_keyList_setList", i})
}

func (ptr *QGeoRouteRequest) ____extraParameters_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____extraParameters_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRouteRequest) ____setExtraParameters_extraParameters_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setExtraParameters_extraParameters_keyList_atList", i}).(string)
}

func (ptr *QGeoRouteRequest) ____setExtraParameters_extraParameters_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setExtraParameters_extraParameters_keyList_setList", i})
}

func (ptr *QGeoRouteRequest) ____setExtraParameters_extraParameters_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setExtraParameters_extraParameters_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRouteRequest) ____setWaypointsMetadata_waypointMetadata_atList_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setWaypointsMetadata_waypointMetadata_atList_atList", v, i}).(*core.QVariant)
}

func (ptr *QGeoRouteRequest) ____setWaypointsMetadata_waypointMetadata_atList_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setWaypointsMetadata_waypointMetadata_atList_setList", key, i})
}

func (ptr *QGeoRouteRequest) ____setWaypointsMetadata_waypointMetadata_atList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setWaypointsMetadata_waypointMetadata_atList_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRouteRequest) ____setWaypointsMetadata_waypointMetadata_atList_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setWaypointsMetadata_waypointMetadata_atList_keyList"}).([]string)
}

func (ptr *QGeoRouteRequest) ____setWaypointsMetadata_waypointMetadata_setList_i_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setWaypointsMetadata_waypointMetadata_setList_i_atList", v, i}).(*core.QVariant)
}

func (ptr *QGeoRouteRequest) ____setWaypointsMetadata_waypointMetadata_setList_i_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setWaypointsMetadata_waypointMetadata_setList_i_setList", key, i})
}

func (ptr *QGeoRouteRequest) ____setWaypointsMetadata_waypointMetadata_setList_i_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setWaypointsMetadata_waypointMetadata_setList_i_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRouteRequest) ____setWaypointsMetadata_waypointMetadata_setList_i_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setWaypointsMetadata_waypointMetadata_setList_i_keyList"}).([]string)
}

func (ptr *QGeoRouteRequest) ____waypointsMetadata_atList_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____waypointsMetadata_atList_atList", v, i}).(*core.QVariant)
}

func (ptr *QGeoRouteRequest) ____waypointsMetadata_atList_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____waypointsMetadata_atList_setList", key, i})
}

func (ptr *QGeoRouteRequest) ____waypointsMetadata_atList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____waypointsMetadata_atList_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRouteRequest) ____waypointsMetadata_atList_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____waypointsMetadata_atList_keyList"}).([]string)
}

func (ptr *QGeoRouteRequest) ____waypointsMetadata_setList_i_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____waypointsMetadata_setList_i_atList", v, i}).(*core.QVariant)
}

func (ptr *QGeoRouteRequest) ____waypointsMetadata_setList_i_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____waypointsMetadata_setList_i_setList", key, i})
}

func (ptr *QGeoRouteRequest) ____waypointsMetadata_setList_i_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____waypointsMetadata_setList_i_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRouteRequest) ____waypointsMetadata_setList_i_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____waypointsMetadata_setList_i_keyList"}).([]string)
}

func (ptr *QGeoRouteRequest) ______setWaypointsMetadata_waypointMetadata_atList_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "______setWaypointsMetadata_waypointMetadata_atList_keyList_atList", i}).(string)
}

func (ptr *QGeoRouteRequest) ______setWaypointsMetadata_waypointMetadata_atList_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "______setWaypointsMetadata_waypointMetadata_atList_keyList_setList", i})
}

func (ptr *QGeoRouteRequest) ______setWaypointsMetadata_waypointMetadata_atList_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "______setWaypointsMetadata_waypointMetadata_atList_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRouteRequest) ______setWaypointsMetadata_waypointMetadata_setList_i_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "______setWaypointsMetadata_waypointMetadata_setList_i_keyList_atList", i}).(string)
}

func (ptr *QGeoRouteRequest) ______setWaypointsMetadata_waypointMetadata_setList_i_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "______setWaypointsMetadata_waypointMetadata_setList_i_keyList_setList", i})
}

func (ptr *QGeoRouteRequest) ______setWaypointsMetadata_waypointMetadata_setList_i_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "______setWaypointsMetadata_waypointMetadata_setList_i_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRouteRequest) ______waypointsMetadata_atList_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "______waypointsMetadata_atList_keyList_atList", i}).(string)
}

func (ptr *QGeoRouteRequest) ______waypointsMetadata_atList_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "______waypointsMetadata_atList_keyList_setList", i})
}

func (ptr *QGeoRouteRequest) ______waypointsMetadata_atList_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "______waypointsMetadata_atList_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRouteRequest) ______waypointsMetadata_setList_i_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "______waypointsMetadata_setList_i_keyList_atList", i}).(string)
}

func (ptr *QGeoRouteRequest) ______waypointsMetadata_setList_i_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "______waypointsMetadata_setList_i_keyList_setList", i})
}

func (ptr *QGeoRouteRequest) ______waypointsMetadata_setList_i_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "______waypointsMetadata_setList_i_keyList_newList"}).(unsafe.Pointer)
}

type QGeoRouteSegment struct {
	internal.Internal
}

type QGeoRouteSegment_ITF interface {
	QGeoRouteSegment_PTR() *QGeoRouteSegment
}

func (ptr *QGeoRouteSegment) QGeoRouteSegment_PTR() *QGeoRouteSegment {
	return ptr
}

func (ptr *QGeoRouteSegment) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QGeoRouteSegment) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQGeoRouteSegment(ptr QGeoRouteSegment_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoRouteSegment_PTR().Pointer()
	}
	return nil
}

func (n *QGeoRouteSegment) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQGeoRouteSegmentFromPointer(ptr unsafe.Pointer) (n *QGeoRouteSegment) {
	n = new(QGeoRouteSegment)
	n.InitFromInternal(uintptr(ptr), "location.QGeoRouteSegment")
	return
}
func NewQGeoRouteSegment() *QGeoRouteSegment {

	return internal.CallLocalFunction([]interface{}{"", "", "location.NewQGeoRouteSegment", ""}).(*QGeoRouteSegment)
}

func NewQGeoRouteSegment2(other QGeoRouteSegment_ITF) *QGeoRouteSegment {

	return internal.CallLocalFunction([]interface{}{"", "", "location.NewQGeoRouteSegment2", "", other}).(*QGeoRouteSegment)
}

func (ptr *QGeoRouteSegment) Distance() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Distance"}).(float64)
}

func (ptr *QGeoRouteSegment) IsLegLastSegment() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsLegLastSegment"}).(bool)
}

func (ptr *QGeoRouteSegment) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QGeoRouteSegment) Maneuver() *QGeoManeuver {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Maneuver"}).(*QGeoManeuver)
}

func (ptr *QGeoRouteSegment) NextRouteSegment() *QGeoRouteSegment {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NextRouteSegment"}).(*QGeoRouteSegment)
}

func (ptr *QGeoRouteSegment) Path() []*positioning.QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Path"}).([]*positioning.QGeoCoordinate)
}

func (ptr *QGeoRouteSegment) SetDistance(distance float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDistance", distance})
}

func (ptr *QGeoRouteSegment) SetManeuver(maneuver QGeoManeuver_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetManeuver", maneuver})
}

func (ptr *QGeoRouteSegment) SetNextRouteSegment(routeSegment QGeoRouteSegment_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetNextRouteSegment", routeSegment})
}

func (ptr *QGeoRouteSegment) SetPath(path []*positioning.QGeoCoordinate) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPath", path})
}

func (ptr *QGeoRouteSegment) SetTravelTime(secs int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTravelTime", secs})
}

func (ptr *QGeoRouteSegment) TravelTime() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TravelTime"}).(float64))
}

func (ptr *QGeoRouteSegment) DestroyQGeoRouteSegment() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoRouteSegment"})
}

func (ptr *QGeoRouteSegment) __path_atList(i int) *positioning.QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__path_atList", i}).(*positioning.QGeoCoordinate)
}

func (ptr *QGeoRouteSegment) __path_setList(i positioning.QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__path_setList", i})
}

func (ptr *QGeoRouteSegment) __path_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__path_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRouteSegment) __setPath_path_atList(i int) *positioning.QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setPath_path_atList", i}).(*positioning.QGeoCoordinate)
}

func (ptr *QGeoRouteSegment) __setPath_path_setList(i positioning.QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setPath_path_setList", i})
}

func (ptr *QGeoRouteSegment) __setPath_path_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setPath_path_newList"}).(unsafe.Pointer)
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

func (n *QGeoRoutingManager) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QGeoRoutingManager) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQGeoRoutingManagerFromPointer(ptr unsafe.Pointer) (n *QGeoRoutingManager) {
	n = new(QGeoRoutingManager)
	n.InitFromInternal(uintptr(ptr), "location.QGeoRoutingManager")
	return
}
func (ptr *QGeoRoutingManager) CalculateRoute(request QGeoRouteRequest_ITF) *QGeoRouteReply {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CalculateRoute", request}).(*QGeoRouteReply)
}

func (ptr *QGeoRoutingManager) ConnectError(f func(reply *QGeoRouteReply, error QGeoRouteReply__Error, errorString string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectError", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoRoutingManager) DisconnectError() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError"})
}

func (ptr *QGeoRoutingManager) Error(reply QGeoRouteReply_ITF, error QGeoRouteReply__Error, errorString string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error", reply, error, errorString})
}

func (ptr *QGeoRoutingManager) ConnectFinished(f func(reply *QGeoRouteReply)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFinished", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoRoutingManager) DisconnectFinished() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFinished"})
}

func (ptr *QGeoRoutingManager) Finished(reply QGeoRouteReply_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Finished", reply})
}

func (ptr *QGeoRoutingManager) Locale() *core.QLocale {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Locale"}).(*core.QLocale)
}

func (ptr *QGeoRoutingManager) ManagerName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ManagerName"}).(string)
}

func (ptr *QGeoRoutingManager) ManagerVersion() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ManagerVersion"}).(float64))
}

func (ptr *QGeoRoutingManager) MeasurementSystem() core.QLocale__MeasurementSystem {

	return core.QLocale__MeasurementSystem(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MeasurementSystem"}).(float64))
}

func (ptr *QGeoRoutingManager) SetLocale(locale core.QLocale_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLocale", locale})
}

func (ptr *QGeoRoutingManager) SetMeasurementSystem(system core.QLocale__MeasurementSystem) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMeasurementSystem", system})
}

func (ptr *QGeoRoutingManager) SupportedFeatureTypes() QGeoRouteRequest__FeatureType {

	return QGeoRouteRequest__FeatureType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedFeatureTypes"}).(float64))
}

func (ptr *QGeoRoutingManager) SupportedFeatureWeights() QGeoRouteRequest__FeatureWeight {

	return QGeoRouteRequest__FeatureWeight(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedFeatureWeights"}).(float64))
}

func (ptr *QGeoRoutingManager) SupportedManeuverDetails() QGeoRouteRequest__ManeuverDetail {

	return QGeoRouteRequest__ManeuverDetail(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedManeuverDetails"}).(float64))
}

func (ptr *QGeoRoutingManager) SupportedRouteOptimizations() QGeoRouteRequest__RouteOptimization {

	return QGeoRouteRequest__RouteOptimization(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedRouteOptimizations"}).(float64))
}

func (ptr *QGeoRoutingManager) SupportedSegmentDetails() QGeoRouteRequest__SegmentDetail {

	return QGeoRouteRequest__SegmentDetail(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedSegmentDetails"}).(float64))
}

func (ptr *QGeoRoutingManager) SupportedTravelModes() QGeoRouteRequest__TravelMode {

	return QGeoRouteRequest__TravelMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedTravelModes"}).(float64))
}

func (ptr *QGeoRoutingManager) UpdateRoute(route QGeoRoute_ITF, position positioning.QGeoCoordinate_ITF) *QGeoRouteReply {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateRoute", route, position}).(*QGeoRouteReply)
}

func (ptr *QGeoRoutingManager) ConnectDestroyQGeoRoutingManager(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQGeoRoutingManager", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoRoutingManager) DisconnectDestroyQGeoRoutingManager() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQGeoRoutingManager"})
}

func (ptr *QGeoRoutingManager) DestroyQGeoRoutingManager() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoRoutingManager"})
}

func (ptr *QGeoRoutingManager) DestroyQGeoRoutingManagerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoRoutingManagerDefault"})
}

func (ptr *QGeoRoutingManager) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QGeoRoutingManager) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QGeoRoutingManager) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRoutingManager) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QGeoRoutingManager) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QGeoRoutingManager) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRoutingManager) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QGeoRoutingManager) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QGeoRoutingManager) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRoutingManager) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QGeoRoutingManager) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QGeoRoutingManager) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QGeoRoutingManager) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QGeoRoutingManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QGeoRoutingManager) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QGeoRoutingManager) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QGeoRoutingManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QGeoRoutingManager) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QGeoRoutingManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QGeoRoutingManager) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QGeoRoutingManager) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
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

func (n *QGeoRoutingManagerEngine) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QGeoRoutingManagerEngine) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQGeoRoutingManagerEngineFromPointer(ptr unsafe.Pointer) (n *QGeoRoutingManagerEngine) {
	n = new(QGeoRoutingManagerEngine)
	n.InitFromInternal(uintptr(ptr), "location.QGeoRoutingManagerEngine")
	return
}
func NewQGeoRoutingManagerEngine(parameters map[string]*core.QVariant, parent core.QObject_ITF) *QGeoRoutingManagerEngine {

	return internal.CallLocalFunction([]interface{}{"", "", "location.NewQGeoRoutingManagerEngine", "", parameters, parent}).(*QGeoRoutingManagerEngine)
}

func (ptr *QGeoRoutingManagerEngine) ConnectCalculateRoute(f func(request *QGeoRouteRequest) *QGeoRouteReply) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCalculateRoute", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoRoutingManagerEngine) DisconnectCalculateRoute() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCalculateRoute"})
}

func (ptr *QGeoRoutingManagerEngine) CalculateRoute(request QGeoRouteRequest_ITF) *QGeoRouteReply {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CalculateRoute", request}).(*QGeoRouteReply)
}

func (ptr *QGeoRoutingManagerEngine) ConnectError(f func(reply *QGeoRouteReply, error QGeoRouteReply__Error, errorString string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectError", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoRoutingManagerEngine) DisconnectError() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError"})
}

func (ptr *QGeoRoutingManagerEngine) Error(reply QGeoRouteReply_ITF, error QGeoRouteReply__Error, errorString string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error", reply, error, errorString})
}

func (ptr *QGeoRoutingManagerEngine) ConnectFinished(f func(reply *QGeoRouteReply)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFinished", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoRoutingManagerEngine) DisconnectFinished() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFinished"})
}

func (ptr *QGeoRoutingManagerEngine) Finished(reply QGeoRouteReply_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Finished", reply})
}

func (ptr *QGeoRoutingManagerEngine) Locale() *core.QLocale {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Locale"}).(*core.QLocale)
}

func (ptr *QGeoRoutingManagerEngine) ManagerName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ManagerName"}).(string)
}

func (ptr *QGeoRoutingManagerEngine) ManagerVersion() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ManagerVersion"}).(float64))
}

func (ptr *QGeoRoutingManagerEngine) MeasurementSystem() core.QLocale__MeasurementSystem {

	return core.QLocale__MeasurementSystem(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MeasurementSystem"}).(float64))
}

func (ptr *QGeoRoutingManagerEngine) SetLocale(locale core.QLocale_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLocale", locale})
}

func (ptr *QGeoRoutingManagerEngine) SetMeasurementSystem(system core.QLocale__MeasurementSystem) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMeasurementSystem", system})
}

func (ptr *QGeoRoutingManagerEngine) SetSupportedFeatureTypes(featureTypes QGeoRouteRequest__FeatureType) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSupportedFeatureTypes", featureTypes})
}

func (ptr *QGeoRoutingManagerEngine) SetSupportedFeatureWeights(featureWeights QGeoRouteRequest__FeatureWeight) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSupportedFeatureWeights", featureWeights})
}

func (ptr *QGeoRoutingManagerEngine) SetSupportedManeuverDetails(maneuverDetails QGeoRouteRequest__ManeuverDetail) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSupportedManeuverDetails", maneuverDetails})
}

func (ptr *QGeoRoutingManagerEngine) SetSupportedRouteOptimizations(optimizations QGeoRouteRequest__RouteOptimization) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSupportedRouteOptimizations", optimizations})
}

func (ptr *QGeoRoutingManagerEngine) SetSupportedSegmentDetails(segmentDetails QGeoRouteRequest__SegmentDetail) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSupportedSegmentDetails", segmentDetails})
}

func (ptr *QGeoRoutingManagerEngine) SetSupportedTravelModes(travelModes QGeoRouteRequest__TravelMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSupportedTravelModes", travelModes})
}

func (ptr *QGeoRoutingManagerEngine) SupportedFeatureTypes() QGeoRouteRequest__FeatureType {

	return QGeoRouteRequest__FeatureType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedFeatureTypes"}).(float64))
}

func (ptr *QGeoRoutingManagerEngine) SupportedFeatureWeights() QGeoRouteRequest__FeatureWeight {

	return QGeoRouteRequest__FeatureWeight(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedFeatureWeights"}).(float64))
}

func (ptr *QGeoRoutingManagerEngine) SupportedManeuverDetails() QGeoRouteRequest__ManeuverDetail {

	return QGeoRouteRequest__ManeuverDetail(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedManeuverDetails"}).(float64))
}

func (ptr *QGeoRoutingManagerEngine) SupportedRouteOptimizations() QGeoRouteRequest__RouteOptimization {

	return QGeoRouteRequest__RouteOptimization(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedRouteOptimizations"}).(float64))
}

func (ptr *QGeoRoutingManagerEngine) SupportedSegmentDetails() QGeoRouteRequest__SegmentDetail {

	return QGeoRouteRequest__SegmentDetail(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedSegmentDetails"}).(float64))
}

func (ptr *QGeoRoutingManagerEngine) SupportedTravelModes() QGeoRouteRequest__TravelMode {

	return QGeoRouteRequest__TravelMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedTravelModes"}).(float64))
}

func (ptr *QGeoRoutingManagerEngine) ConnectUpdateRoute(f func(route *QGeoRoute, position *positioning.QGeoCoordinate) *QGeoRouteReply) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUpdateRoute", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoRoutingManagerEngine) DisconnectUpdateRoute() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUpdateRoute"})
}

func (ptr *QGeoRoutingManagerEngine) UpdateRoute(route QGeoRoute_ITF, position positioning.QGeoCoordinate_ITF) *QGeoRouteReply {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateRoute", route, position}).(*QGeoRouteReply)
}

func (ptr *QGeoRoutingManagerEngine) UpdateRouteDefault(route QGeoRoute_ITF, position positioning.QGeoCoordinate_ITF) *QGeoRouteReply {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateRouteDefault", route, position}).(*QGeoRouteReply)
}

func (ptr *QGeoRoutingManagerEngine) ConnectDestroyQGeoRoutingManagerEngine(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQGeoRoutingManagerEngine", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoRoutingManagerEngine) DisconnectDestroyQGeoRoutingManagerEngine() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQGeoRoutingManagerEngine"})
}

func (ptr *QGeoRoutingManagerEngine) DestroyQGeoRoutingManagerEngine() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoRoutingManagerEngine"})
}

func (ptr *QGeoRoutingManagerEngine) DestroyQGeoRoutingManagerEngineDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoRoutingManagerEngineDefault"})
}

func (ptr *QGeoRoutingManagerEngine) __QGeoRoutingManagerEngine_parameters_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QGeoRoutingManagerEngine_parameters_atList", v, i}).(*core.QVariant)
}

func (ptr *QGeoRoutingManagerEngine) __QGeoRoutingManagerEngine_parameters_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QGeoRoutingManagerEngine_parameters_setList", key, i})
}

func (ptr *QGeoRoutingManagerEngine) __QGeoRoutingManagerEngine_parameters_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QGeoRoutingManagerEngine_parameters_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRoutingManagerEngine) __QGeoRoutingManagerEngine_parameters_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QGeoRoutingManagerEngine_parameters_keyList"}).([]string)
}

func (ptr *QGeoRoutingManagerEngine) ____QGeoRoutingManagerEngine_parameters_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____QGeoRoutingManagerEngine_parameters_keyList_atList", i}).(string)
}

func (ptr *QGeoRoutingManagerEngine) ____QGeoRoutingManagerEngine_parameters_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____QGeoRoutingManagerEngine_parameters_keyList_setList", i})
}

func (ptr *QGeoRoutingManagerEngine) ____QGeoRoutingManagerEngine_parameters_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____QGeoRoutingManagerEngine_parameters_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRoutingManagerEngine) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QGeoRoutingManagerEngine) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QGeoRoutingManagerEngine) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRoutingManagerEngine) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QGeoRoutingManagerEngine) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QGeoRoutingManagerEngine) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRoutingManagerEngine) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QGeoRoutingManagerEngine) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QGeoRoutingManagerEngine) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoRoutingManagerEngine) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QGeoRoutingManagerEngine) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QGeoRoutingManagerEngine) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QGeoRoutingManagerEngine) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QGeoRoutingManagerEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QGeoRoutingManagerEngine) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QGeoRoutingManagerEngine) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QGeoRoutingManagerEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QGeoRoutingManagerEngine) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QGeoRoutingManagerEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QGeoRoutingManagerEngine) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QGeoRoutingManagerEngine) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
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

func (n *QGeoServiceProvider) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QGeoServiceProvider) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQGeoServiceProviderFromPointer(ptr unsafe.Pointer) (n *QGeoServiceProvider) {
	n = new(QGeoServiceProvider)
	n.InitFromInternal(uintptr(ptr), "location.QGeoServiceProvider")
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
	QGeoServiceProvider__OnlineRoutingFeature       QGeoServiceProvider__RoutingFeature = QGeoServiceProvider__RoutingFeature(0)
	QGeoServiceProvider__OfflineRoutingFeature      QGeoServiceProvider__RoutingFeature = QGeoServiceProvider__RoutingFeature(0)
	QGeoServiceProvider__LocalizedRoutingFeature    QGeoServiceProvider__RoutingFeature = QGeoServiceProvider__RoutingFeature(0)
	QGeoServiceProvider__RouteUpdatesFeature        QGeoServiceProvider__RoutingFeature = QGeoServiceProvider__RoutingFeature(0)
	QGeoServiceProvider__AlternativeRoutesFeature   QGeoServiceProvider__RoutingFeature = QGeoServiceProvider__RoutingFeature(0)
	QGeoServiceProvider__ExcludeAreasRoutingFeature QGeoServiceProvider__RoutingFeature = QGeoServiceProvider__RoutingFeature(0)
	QGeoServiceProvider__AnyRoutingFeatures         QGeoServiceProvider__RoutingFeature = QGeoServiceProvider__RoutingFeature(0)
)

//go:generate stringer -type=QGeoServiceProvider__GeocodingFeature
//QGeoServiceProvider::GeocodingFeature
type QGeoServiceProvider__GeocodingFeature int64

var (
	QGeoServiceProvider__NoGeocodingFeatures       QGeoServiceProvider__GeocodingFeature = QGeoServiceProvider__GeocodingFeature(0)
	QGeoServiceProvider__OnlineGeocodingFeature    QGeoServiceProvider__GeocodingFeature = QGeoServiceProvider__GeocodingFeature(0)
	QGeoServiceProvider__OfflineGeocodingFeature   QGeoServiceProvider__GeocodingFeature = QGeoServiceProvider__GeocodingFeature(0)
	QGeoServiceProvider__ReverseGeocodingFeature   QGeoServiceProvider__GeocodingFeature = QGeoServiceProvider__GeocodingFeature(0)
	QGeoServiceProvider__LocalizedGeocodingFeature QGeoServiceProvider__GeocodingFeature = QGeoServiceProvider__GeocodingFeature(0)
	QGeoServiceProvider__AnyGeocodingFeatures      QGeoServiceProvider__GeocodingFeature = QGeoServiceProvider__GeocodingFeature(0)
)

//go:generate stringer -type=QGeoServiceProvider__MappingFeature
//QGeoServiceProvider::MappingFeature
type QGeoServiceProvider__MappingFeature int64

var (
	QGeoServiceProvider__NoMappingFeatures       QGeoServiceProvider__MappingFeature = QGeoServiceProvider__MappingFeature(0)
	QGeoServiceProvider__OnlineMappingFeature    QGeoServiceProvider__MappingFeature = QGeoServiceProvider__MappingFeature(0)
	QGeoServiceProvider__OfflineMappingFeature   QGeoServiceProvider__MappingFeature = QGeoServiceProvider__MappingFeature(0)
	QGeoServiceProvider__LocalizedMappingFeature QGeoServiceProvider__MappingFeature = QGeoServiceProvider__MappingFeature(0)
	QGeoServiceProvider__AnyMappingFeatures      QGeoServiceProvider__MappingFeature = QGeoServiceProvider__MappingFeature(0)
)

//go:generate stringer -type=QGeoServiceProvider__PlacesFeature
//QGeoServiceProvider::PlacesFeature
type QGeoServiceProvider__PlacesFeature int64

var (
	QGeoServiceProvider__NoPlacesFeatures            QGeoServiceProvider__PlacesFeature = QGeoServiceProvider__PlacesFeature(0)
	QGeoServiceProvider__OnlinePlacesFeature         QGeoServiceProvider__PlacesFeature = QGeoServiceProvider__PlacesFeature(0)
	QGeoServiceProvider__OfflinePlacesFeature        QGeoServiceProvider__PlacesFeature = QGeoServiceProvider__PlacesFeature(0)
	QGeoServiceProvider__SavePlaceFeature            QGeoServiceProvider__PlacesFeature = QGeoServiceProvider__PlacesFeature(0)
	QGeoServiceProvider__RemovePlaceFeature          QGeoServiceProvider__PlacesFeature = QGeoServiceProvider__PlacesFeature(0)
	QGeoServiceProvider__SaveCategoryFeature         QGeoServiceProvider__PlacesFeature = QGeoServiceProvider__PlacesFeature(0)
	QGeoServiceProvider__RemoveCategoryFeature       QGeoServiceProvider__PlacesFeature = QGeoServiceProvider__PlacesFeature(0)
	QGeoServiceProvider__PlaceRecommendationsFeature QGeoServiceProvider__PlacesFeature = QGeoServiceProvider__PlacesFeature(0)
	QGeoServiceProvider__SearchSuggestionsFeature    QGeoServiceProvider__PlacesFeature = QGeoServiceProvider__PlacesFeature(0)
	QGeoServiceProvider__LocalizedPlacesFeature      QGeoServiceProvider__PlacesFeature = QGeoServiceProvider__PlacesFeature(0)
	QGeoServiceProvider__NotificationsFeature        QGeoServiceProvider__PlacesFeature = QGeoServiceProvider__PlacesFeature(0)
	QGeoServiceProvider__PlaceMatchingFeature        QGeoServiceProvider__PlacesFeature = QGeoServiceProvider__PlacesFeature(0)
	QGeoServiceProvider__AnyPlacesFeatures           QGeoServiceProvider__PlacesFeature = QGeoServiceProvider__PlacesFeature(0)
)

//go:generate stringer -type=QGeoServiceProvider__NavigationFeature
//QGeoServiceProvider::NavigationFeature
type QGeoServiceProvider__NavigationFeature int64

var (
	QGeoServiceProvider__NoNavigationFeatures     QGeoServiceProvider__NavigationFeature = QGeoServiceProvider__NavigationFeature(0)
	QGeoServiceProvider__OnlineNavigationFeature  QGeoServiceProvider__NavigationFeature = QGeoServiceProvider__NavigationFeature(0)
	QGeoServiceProvider__OfflineNavigationFeature QGeoServiceProvider__NavigationFeature = QGeoServiceProvider__NavigationFeature(0)
	QGeoServiceProvider__AnyNavigationFeatures    QGeoServiceProvider__NavigationFeature = QGeoServiceProvider__NavigationFeature(0)
)

func NewQGeoServiceProvider(providerName string, parameters map[string]*core.QVariant, allowExperimental bool) *QGeoServiceProvider {

	return internal.CallLocalFunction([]interface{}{"", "", "location.NewQGeoServiceProvider", "", providerName, parameters, allowExperimental}).(*QGeoServiceProvider)
}

func QGeoServiceProvider_AvailableServiceProviders() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "location.QGeoServiceProvider_AvailableServiceProviders", ""}).([]string)
}

func (ptr *QGeoServiceProvider) AvailableServiceProviders() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "location.QGeoServiceProvider_AvailableServiceProviders", ""}).([]string)
}

func (ptr *QGeoServiceProvider) Error() QGeoServiceProvider__Error {

	return QGeoServiceProvider__Error(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QGeoServiceProvider) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func (ptr *QGeoServiceProvider) GeocodingError() QGeoServiceProvider__Error {

	return QGeoServiceProvider__Error(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GeocodingError"}).(float64))
}

func (ptr *QGeoServiceProvider) GeocodingErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GeocodingErrorString"}).(string)
}

func (ptr *QGeoServiceProvider) GeocodingFeatures() QGeoServiceProvider__GeocodingFeature {

	return QGeoServiceProvider__GeocodingFeature(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GeocodingFeatures"}).(float64))
}

func (ptr *QGeoServiceProvider) GeocodingManager() *QGeoCodingManager {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GeocodingManager"}).(*QGeoCodingManager)
}

func (ptr *QGeoServiceProvider) MappingError() QGeoServiceProvider__Error {

	return QGeoServiceProvider__Error(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MappingError"}).(float64))
}

func (ptr *QGeoServiceProvider) MappingErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MappingErrorString"}).(string)
}

func (ptr *QGeoServiceProvider) MappingFeatures() QGeoServiceProvider__MappingFeature {

	return QGeoServiceProvider__MappingFeature(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MappingFeatures"}).(float64))
}

func (ptr *QGeoServiceProvider) NavigationError() QGeoServiceProvider__Error {

	return QGeoServiceProvider__Error(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NavigationError"}).(float64))
}

func (ptr *QGeoServiceProvider) NavigationErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NavigationErrorString"}).(string)
}

func (ptr *QGeoServiceProvider) NavigationFeatures() QGeoServiceProvider__NavigationFeature {

	return QGeoServiceProvider__NavigationFeature(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NavigationFeatures"}).(float64))
}

func (ptr *QGeoServiceProvider) PlaceManager() *QPlaceManager {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PlaceManager"}).(*QPlaceManager)
}

func (ptr *QGeoServiceProvider) PlacesError() QGeoServiceProvider__Error {

	return QGeoServiceProvider__Error(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PlacesError"}).(float64))
}

func (ptr *QGeoServiceProvider) PlacesErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PlacesErrorString"}).(string)
}

func (ptr *QGeoServiceProvider) PlacesFeatures() QGeoServiceProvider__PlacesFeature {

	return QGeoServiceProvider__PlacesFeature(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PlacesFeatures"}).(float64))
}

func (ptr *QGeoServiceProvider) RoutingError() QGeoServiceProvider__Error {

	return QGeoServiceProvider__Error(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RoutingError"}).(float64))
}

func (ptr *QGeoServiceProvider) RoutingErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RoutingErrorString"}).(string)
}

func (ptr *QGeoServiceProvider) RoutingFeatures() QGeoServiceProvider__RoutingFeature {

	return QGeoServiceProvider__RoutingFeature(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RoutingFeatures"}).(float64))
}

func (ptr *QGeoServiceProvider) RoutingManager() *QGeoRoutingManager {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RoutingManager"}).(*QGeoRoutingManager)
}

func (ptr *QGeoServiceProvider) SetAllowExperimental(allow bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAllowExperimental", allow})
}

func (ptr *QGeoServiceProvider) SetLocale(locale core.QLocale_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLocale", locale})
}

func (ptr *QGeoServiceProvider) SetParameters(parameters map[string]*core.QVariant) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetParameters", parameters})
}

func (ptr *QGeoServiceProvider) ConnectDestroyQGeoServiceProvider(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQGeoServiceProvider", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoServiceProvider) DisconnectDestroyQGeoServiceProvider() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQGeoServiceProvider"})
}

func (ptr *QGeoServiceProvider) DestroyQGeoServiceProvider() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoServiceProvider"})
}

func (ptr *QGeoServiceProvider) DestroyQGeoServiceProviderDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoServiceProviderDefault"})
}

func (ptr *QGeoServiceProvider) __QGeoServiceProvider_parameters_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QGeoServiceProvider_parameters_atList", v, i}).(*core.QVariant)
}

func (ptr *QGeoServiceProvider) __QGeoServiceProvider_parameters_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QGeoServiceProvider_parameters_setList", key, i})
}

func (ptr *QGeoServiceProvider) __QGeoServiceProvider_parameters_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QGeoServiceProvider_parameters_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoServiceProvider) __QGeoServiceProvider_parameters_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QGeoServiceProvider_parameters_keyList"}).([]string)
}

func (ptr *QGeoServiceProvider) __setParameters_parameters_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setParameters_parameters_atList", v, i}).(*core.QVariant)
}

func (ptr *QGeoServiceProvider) __setParameters_parameters_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setParameters_parameters_setList", key, i})
}

func (ptr *QGeoServiceProvider) __setParameters_parameters_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setParameters_parameters_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoServiceProvider) __setParameters_parameters_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setParameters_parameters_keyList"}).([]string)
}

func (ptr *QGeoServiceProvider) ____QGeoServiceProvider_parameters_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____QGeoServiceProvider_parameters_keyList_atList", i}).(string)
}

func (ptr *QGeoServiceProvider) ____QGeoServiceProvider_parameters_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____QGeoServiceProvider_parameters_keyList_setList", i})
}

func (ptr *QGeoServiceProvider) ____QGeoServiceProvider_parameters_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____QGeoServiceProvider_parameters_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoServiceProvider) ____setParameters_parameters_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setParameters_parameters_keyList_atList", i}).(string)
}

func (ptr *QGeoServiceProvider) ____setParameters_parameters_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setParameters_parameters_keyList_setList", i})
}

func (ptr *QGeoServiceProvider) ____setParameters_parameters_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setParameters_parameters_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoServiceProvider) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QGeoServiceProvider) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QGeoServiceProvider) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoServiceProvider) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QGeoServiceProvider) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QGeoServiceProvider) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoServiceProvider) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QGeoServiceProvider) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QGeoServiceProvider) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoServiceProvider) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QGeoServiceProvider) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QGeoServiceProvider) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QGeoServiceProvider) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QGeoServiceProvider) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QGeoServiceProvider) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QGeoServiceProvider) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QGeoServiceProvider) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QGeoServiceProvider) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QGeoServiceProvider) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QGeoServiceProvider) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QGeoServiceProvider) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QGeoServiceProviderFactory struct {
	internal.Internal
}

type QGeoServiceProviderFactory_ITF interface {
	QGeoServiceProviderFactory_PTR() *QGeoServiceProviderFactory
}

func (ptr *QGeoServiceProviderFactory) QGeoServiceProviderFactory_PTR() *QGeoServiceProviderFactory {
	return ptr
}

func (ptr *QGeoServiceProviderFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QGeoServiceProviderFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQGeoServiceProviderFactory(ptr QGeoServiceProviderFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoServiceProviderFactory_PTR().Pointer()
	}
	return nil
}

func (n *QGeoServiceProviderFactory) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQGeoServiceProviderFactoryFromPointer(ptr unsafe.Pointer) (n *QGeoServiceProviderFactory) {
	n = new(QGeoServiceProviderFactory)
	n.InitFromInternal(uintptr(ptr), "location.QGeoServiceProviderFactory")
	return
}
func (ptr *QGeoServiceProviderFactory) ConnectDestroyQGeoServiceProviderFactory(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQGeoServiceProviderFactory", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoServiceProviderFactory) DisconnectDestroyQGeoServiceProviderFactory() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQGeoServiceProviderFactory"})
}

func (ptr *QGeoServiceProviderFactory) DestroyQGeoServiceProviderFactory() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoServiceProviderFactory"})
}

func (ptr *QGeoServiceProviderFactory) DestroyQGeoServiceProviderFactoryDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoServiceProviderFactoryDefault"})
}

func (ptr *QGeoServiceProviderFactory) __createGeocodingManagerEngine_parameters_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__createGeocodingManagerEngine_parameters_atList", v, i}).(*core.QVariant)
}

func (ptr *QGeoServiceProviderFactory) __createGeocodingManagerEngine_parameters_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__createGeocodingManagerEngine_parameters_setList", key, i})
}

func (ptr *QGeoServiceProviderFactory) __createGeocodingManagerEngine_parameters_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__createGeocodingManagerEngine_parameters_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoServiceProviderFactory) __createGeocodingManagerEngine_parameters_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__createGeocodingManagerEngine_parameters_keyList"}).([]string)
}

func (ptr *QGeoServiceProviderFactory) __createMappingManagerEngine_parameters_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__createMappingManagerEngine_parameters_atList", v, i}).(*core.QVariant)
}

func (ptr *QGeoServiceProviderFactory) __createMappingManagerEngine_parameters_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__createMappingManagerEngine_parameters_setList", key, i})
}

func (ptr *QGeoServiceProviderFactory) __createMappingManagerEngine_parameters_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__createMappingManagerEngine_parameters_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoServiceProviderFactory) __createMappingManagerEngine_parameters_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__createMappingManagerEngine_parameters_keyList"}).([]string)
}

func (ptr *QGeoServiceProviderFactory) __createPlaceManagerEngine_parameters_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__createPlaceManagerEngine_parameters_atList", v, i}).(*core.QVariant)
}

func (ptr *QGeoServiceProviderFactory) __createPlaceManagerEngine_parameters_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__createPlaceManagerEngine_parameters_setList", key, i})
}

func (ptr *QGeoServiceProviderFactory) __createPlaceManagerEngine_parameters_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__createPlaceManagerEngine_parameters_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoServiceProviderFactory) __createPlaceManagerEngine_parameters_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__createPlaceManagerEngine_parameters_keyList"}).([]string)
}

func (ptr *QGeoServiceProviderFactory) __createRoutingManagerEngine_parameters_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__createRoutingManagerEngine_parameters_atList", v, i}).(*core.QVariant)
}

func (ptr *QGeoServiceProviderFactory) __createRoutingManagerEngine_parameters_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__createRoutingManagerEngine_parameters_setList", key, i})
}

func (ptr *QGeoServiceProviderFactory) __createRoutingManagerEngine_parameters_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__createRoutingManagerEngine_parameters_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoServiceProviderFactory) __createRoutingManagerEngine_parameters_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__createRoutingManagerEngine_parameters_keyList"}).([]string)
}

func (ptr *QGeoServiceProviderFactory) ____createGeocodingManagerEngine_parameters_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____createGeocodingManagerEngine_parameters_keyList_atList", i}).(string)
}

func (ptr *QGeoServiceProviderFactory) ____createGeocodingManagerEngine_parameters_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____createGeocodingManagerEngine_parameters_keyList_setList", i})
}

func (ptr *QGeoServiceProviderFactory) ____createGeocodingManagerEngine_parameters_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____createGeocodingManagerEngine_parameters_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoServiceProviderFactory) ____createMappingManagerEngine_parameters_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____createMappingManagerEngine_parameters_keyList_atList", i}).(string)
}

func (ptr *QGeoServiceProviderFactory) ____createMappingManagerEngine_parameters_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____createMappingManagerEngine_parameters_keyList_setList", i})
}

func (ptr *QGeoServiceProviderFactory) ____createMappingManagerEngine_parameters_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____createMappingManagerEngine_parameters_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoServiceProviderFactory) ____createPlaceManagerEngine_parameters_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____createPlaceManagerEngine_parameters_keyList_atList", i}).(string)
}

func (ptr *QGeoServiceProviderFactory) ____createPlaceManagerEngine_parameters_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____createPlaceManagerEngine_parameters_keyList_setList", i})
}

func (ptr *QGeoServiceProviderFactory) ____createPlaceManagerEngine_parameters_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____createPlaceManagerEngine_parameters_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoServiceProviderFactory) ____createRoutingManagerEngine_parameters_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____createRoutingManagerEngine_parameters_keyList_atList", i}).(string)
}

func (ptr *QGeoServiceProviderFactory) ____createRoutingManagerEngine_parameters_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____createRoutingManagerEngine_parameters_keyList_setList", i})
}

func (ptr *QGeoServiceProviderFactory) ____createRoutingManagerEngine_parameters_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____createRoutingManagerEngine_parameters_keyList_newList"}).(unsafe.Pointer)
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

func (n *QGeoServiceProviderFactoryV2) InitFromInternal(ptr uintptr, name string) {
	n.QGeoServiceProviderFactory_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QGeoServiceProviderFactoryV2) ClassNameInternalF() string {
	return n.QGeoServiceProviderFactory_PTR().ClassNameInternalF()
}

func NewQGeoServiceProviderFactoryV2FromPointer(ptr unsafe.Pointer) (n *QGeoServiceProviderFactoryV2) {
	n = new(QGeoServiceProviderFactoryV2)
	n.InitFromInternal(uintptr(ptr), "location.QGeoServiceProviderFactoryV2")
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

func (n *QGeoServiceProviderFactoryV3) InitFromInternal(ptr uintptr, name string) {
	n.QGeoServiceProviderFactoryV2_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QGeoServiceProviderFactoryV3) ClassNameInternalF() string {
	return n.QGeoServiceProviderFactoryV2_PTR().ClassNameInternalF()
}

func NewQGeoServiceProviderFactoryV3FromPointer(ptr unsafe.Pointer) (n *QGeoServiceProviderFactoryV3) {
	n = new(QGeoServiceProviderFactoryV3)
	n.InitFromInternal(uintptr(ptr), "location.QGeoServiceProviderFactoryV3")
	return
}

func (ptr *QGeoServiceProviderFactoryV3) DestroyQGeoServiceProviderFactoryV3() {
}

type QLocation struct {
	internal.Internal
}

type QLocation_ITF interface {
	QLocation_PTR() *QLocation
}

func (ptr *QLocation) QLocation_PTR() *QLocation {
	return ptr
}

func (ptr *QLocation) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QLocation) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQLocation(ptr QLocation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLocation_PTR().Pointer()
	}
	return nil
}

func (n *QLocation) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQLocationFromPointer(ptr unsafe.Pointer) (n *QLocation) {
	n = new(QLocation)
	n.InitFromInternal(uintptr(ptr), "location.QLocation")
	return
}

func (ptr *QLocation) DestroyQLocation() {
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
	internal.Internal
}

type QPlace_ITF interface {
	QPlace_PTR() *QPlace
}

func (ptr *QPlace) QPlace_PTR() *QPlace {
	return ptr
}

func (ptr *QPlace) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QPlace) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQPlace(ptr QPlace_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlace_PTR().Pointer()
	}
	return nil
}

func (n *QPlace) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQPlaceFromPointer(ptr unsafe.Pointer) (n *QPlace) {
	n = new(QPlace)
	n.InitFromInternal(uintptr(ptr), "location.QPlace")
	return
}

type QPlaceAttribute struct {
	internal.Internal
}

type QPlaceAttribute_ITF interface {
	QPlaceAttribute_PTR() *QPlaceAttribute
}

func (ptr *QPlaceAttribute) QPlaceAttribute_PTR() *QPlaceAttribute {
	return ptr
}

func (ptr *QPlaceAttribute) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QPlaceAttribute) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQPlaceAttribute(ptr QPlaceAttribute_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceAttribute_PTR().Pointer()
	}
	return nil
}

func (n *QPlaceAttribute) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQPlaceAttributeFromPointer(ptr unsafe.Pointer) (n *QPlaceAttribute) {
	n = new(QPlaceAttribute)
	n.InitFromInternal(uintptr(ptr), "location.QPlaceAttribute")
	return
}

type QPlaceCategory struct {
	internal.Internal
}

type QPlaceCategory_ITF interface {
	QPlaceCategory_PTR() *QPlaceCategory
}

func (ptr *QPlaceCategory) QPlaceCategory_PTR() *QPlaceCategory {
	return ptr
}

func (ptr *QPlaceCategory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QPlaceCategory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQPlaceCategory(ptr QPlaceCategory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceCategory_PTR().Pointer()
	}
	return nil
}

func (n *QPlaceCategory) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQPlaceCategoryFromPointer(ptr unsafe.Pointer) (n *QPlaceCategory) {
	n = new(QPlaceCategory)
	n.InitFromInternal(uintptr(ptr), "location.QPlaceCategory")
	return
}

type QPlaceContactDetail struct {
	internal.Internal
}

type QPlaceContactDetail_ITF interface {
	QPlaceContactDetail_PTR() *QPlaceContactDetail
}

func (ptr *QPlaceContactDetail) QPlaceContactDetail_PTR() *QPlaceContactDetail {
	return ptr
}

func (ptr *QPlaceContactDetail) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QPlaceContactDetail) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQPlaceContactDetail(ptr QPlaceContactDetail_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceContactDetail_PTR().Pointer()
	}
	return nil
}

func (n *QPlaceContactDetail) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQPlaceContactDetailFromPointer(ptr unsafe.Pointer) (n *QPlaceContactDetail) {
	n = new(QPlaceContactDetail)
	n.InitFromInternal(uintptr(ptr), "location.QPlaceContactDetail")
	return
}

type QPlaceContent struct {
	internal.Internal
}

type QPlaceContent_ITF interface {
	QPlaceContent_PTR() *QPlaceContent
}

func (ptr *QPlaceContent) QPlaceContent_PTR() *QPlaceContent {
	return ptr
}

func (ptr *QPlaceContent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QPlaceContent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQPlaceContent(ptr QPlaceContent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceContent_PTR().Pointer()
	}
	return nil
}

func (n *QPlaceContent) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQPlaceContentFromPointer(ptr unsafe.Pointer) (n *QPlaceContent) {
	n = new(QPlaceContent)
	n.InitFromInternal(uintptr(ptr), "location.QPlaceContent")
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

func (n *QPlaceContentReply) InitFromInternal(ptr uintptr, name string) {
	n.QPlaceReply_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPlaceContentReply) ClassNameInternalF() string {
	return n.QPlaceReply_PTR().ClassNameInternalF()
}

func NewQPlaceContentReplyFromPointer(ptr unsafe.Pointer) (n *QPlaceContentReply) {
	n = new(QPlaceContentReply)
	n.InitFromInternal(uintptr(ptr), "location.QPlaceContentReply")
	return
}

type QPlaceContentRequest struct {
	internal.Internal
}

type QPlaceContentRequest_ITF interface {
	QPlaceContentRequest_PTR() *QPlaceContentRequest
}

func (ptr *QPlaceContentRequest) QPlaceContentRequest_PTR() *QPlaceContentRequest {
	return ptr
}

func (ptr *QPlaceContentRequest) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QPlaceContentRequest) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQPlaceContentRequest(ptr QPlaceContentRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceContentRequest_PTR().Pointer()
	}
	return nil
}

func (n *QPlaceContentRequest) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQPlaceContentRequestFromPointer(ptr unsafe.Pointer) (n *QPlaceContentRequest) {
	n = new(QPlaceContentRequest)
	n.InitFromInternal(uintptr(ptr), "location.QPlaceContentRequest")
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

func (n *QPlaceDetailsReply) InitFromInternal(ptr uintptr, name string) {
	n.QPlaceReply_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPlaceDetailsReply) ClassNameInternalF() string {
	return n.QPlaceReply_PTR().ClassNameInternalF()
}

func NewQPlaceDetailsReplyFromPointer(ptr unsafe.Pointer) (n *QPlaceDetailsReply) {
	n = new(QPlaceDetailsReply)
	n.InitFromInternal(uintptr(ptr), "location.QPlaceDetailsReply")
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

func (n *QPlaceEditorial) InitFromInternal(ptr uintptr, name string) {
	n.QPlaceContent_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPlaceEditorial) ClassNameInternalF() string {
	return n.QPlaceContent_PTR().ClassNameInternalF()
}

func NewQPlaceEditorialFromPointer(ptr unsafe.Pointer) (n *QPlaceEditorial) {
	n = new(QPlaceEditorial)
	n.InitFromInternal(uintptr(ptr), "location.QPlaceEditorial")
	return
}

type QPlaceIcon struct {
	internal.Internal
}

type QPlaceIcon_ITF interface {
	QPlaceIcon_PTR() *QPlaceIcon
}

func (ptr *QPlaceIcon) QPlaceIcon_PTR() *QPlaceIcon {
	return ptr
}

func (ptr *QPlaceIcon) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QPlaceIcon) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQPlaceIcon(ptr QPlaceIcon_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceIcon_PTR().Pointer()
	}
	return nil
}

func (n *QPlaceIcon) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQPlaceIconFromPointer(ptr unsafe.Pointer) (n *QPlaceIcon) {
	n = new(QPlaceIcon)
	n.InitFromInternal(uintptr(ptr), "location.QPlaceIcon")
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

func (n *QPlaceIdReply) InitFromInternal(ptr uintptr, name string) {
	n.QPlaceReply_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPlaceIdReply) ClassNameInternalF() string {
	return n.QPlaceReply_PTR().ClassNameInternalF()
}

func NewQPlaceIdReplyFromPointer(ptr unsafe.Pointer) (n *QPlaceIdReply) {
	n = new(QPlaceIdReply)
	n.InitFromInternal(uintptr(ptr), "location.QPlaceIdReply")
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

func (n *QPlaceImage) InitFromInternal(ptr uintptr, name string) {
	n.QPlaceContent_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPlaceImage) ClassNameInternalF() string {
	return n.QPlaceContent_PTR().ClassNameInternalF()
}

func NewQPlaceImageFromPointer(ptr unsafe.Pointer) (n *QPlaceImage) {
	n = new(QPlaceImage)
	n.InitFromInternal(uintptr(ptr), "location.QPlaceImage")
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

func (n *QPlaceManager) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPlaceManager) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQPlaceManagerFromPointer(ptr unsafe.Pointer) (n *QPlaceManager) {
	n = new(QPlaceManager)
	n.InitFromInternal(uintptr(ptr), "location.QPlaceManager")
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

func (n *QPlaceManagerEngine) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPlaceManagerEngine) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQPlaceManagerEngineFromPointer(ptr unsafe.Pointer) (n *QPlaceManagerEngine) {
	n = new(QPlaceManagerEngine)
	n.InitFromInternal(uintptr(ptr), "location.QPlaceManagerEngine")
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

func (n *QPlaceMatchReply) InitFromInternal(ptr uintptr, name string) {
	n.QPlaceReply_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPlaceMatchReply) ClassNameInternalF() string {
	return n.QPlaceReply_PTR().ClassNameInternalF()
}

func NewQPlaceMatchReplyFromPointer(ptr unsafe.Pointer) (n *QPlaceMatchReply) {
	n = new(QPlaceMatchReply)
	n.InitFromInternal(uintptr(ptr), "location.QPlaceMatchReply")
	return
}

type QPlaceMatchRequest struct {
	internal.Internal
}

type QPlaceMatchRequest_ITF interface {
	QPlaceMatchRequest_PTR() *QPlaceMatchRequest
}

func (ptr *QPlaceMatchRequest) QPlaceMatchRequest_PTR() *QPlaceMatchRequest {
	return ptr
}

func (ptr *QPlaceMatchRequest) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QPlaceMatchRequest) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQPlaceMatchRequest(ptr QPlaceMatchRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceMatchRequest_PTR().Pointer()
	}
	return nil
}

func (n *QPlaceMatchRequest) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQPlaceMatchRequestFromPointer(ptr unsafe.Pointer) (n *QPlaceMatchRequest) {
	n = new(QPlaceMatchRequest)
	n.InitFromInternal(uintptr(ptr), "location.QPlaceMatchRequest")
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

func (n *QPlaceProposedSearchResult) InitFromInternal(ptr uintptr, name string) {
	n.QPlaceSearchResult_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPlaceProposedSearchResult) ClassNameInternalF() string {
	return n.QPlaceSearchResult_PTR().ClassNameInternalF()
}

func NewQPlaceProposedSearchResultFromPointer(ptr unsafe.Pointer) (n *QPlaceProposedSearchResult) {
	n = new(QPlaceProposedSearchResult)
	n.InitFromInternal(uintptr(ptr), "location.QPlaceProposedSearchResult")
	return
}

type QPlaceRatings struct {
	internal.Internal
}

type QPlaceRatings_ITF interface {
	QPlaceRatings_PTR() *QPlaceRatings
}

func (ptr *QPlaceRatings) QPlaceRatings_PTR() *QPlaceRatings {
	return ptr
}

func (ptr *QPlaceRatings) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QPlaceRatings) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQPlaceRatings(ptr QPlaceRatings_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceRatings_PTR().Pointer()
	}
	return nil
}

func (n *QPlaceRatings) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQPlaceRatingsFromPointer(ptr unsafe.Pointer) (n *QPlaceRatings) {
	n = new(QPlaceRatings)
	n.InitFromInternal(uintptr(ptr), "location.QPlaceRatings")
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

func (n *QPlaceReply) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPlaceReply) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQPlaceReplyFromPointer(ptr unsafe.Pointer) (n *QPlaceReply) {
	n = new(QPlaceReply)
	n.InitFromInternal(uintptr(ptr), "location.QPlaceReply")
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

func (n *QPlaceResult) InitFromInternal(ptr uintptr, name string) {
	n.QPlaceSearchResult_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPlaceResult) ClassNameInternalF() string {
	return n.QPlaceSearchResult_PTR().ClassNameInternalF()
}

func NewQPlaceResultFromPointer(ptr unsafe.Pointer) (n *QPlaceResult) {
	n = new(QPlaceResult)
	n.InitFromInternal(uintptr(ptr), "location.QPlaceResult")
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

func (n *QPlaceReview) InitFromInternal(ptr uintptr, name string) {
	n.QPlaceContent_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPlaceReview) ClassNameInternalF() string {
	return n.QPlaceContent_PTR().ClassNameInternalF()
}

func NewQPlaceReviewFromPointer(ptr unsafe.Pointer) (n *QPlaceReview) {
	n = new(QPlaceReview)
	n.InitFromInternal(uintptr(ptr), "location.QPlaceReview")
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

func (n *QPlaceSearchReply) InitFromInternal(ptr uintptr, name string) {
	n.QPlaceReply_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPlaceSearchReply) ClassNameInternalF() string {
	return n.QPlaceReply_PTR().ClassNameInternalF()
}

func NewQPlaceSearchReplyFromPointer(ptr unsafe.Pointer) (n *QPlaceSearchReply) {
	n = new(QPlaceSearchReply)
	n.InitFromInternal(uintptr(ptr), "location.QPlaceSearchReply")
	return
}

type QPlaceSearchRequest struct {
	internal.Internal
}

type QPlaceSearchRequest_ITF interface {
	QPlaceSearchRequest_PTR() *QPlaceSearchRequest
}

func (ptr *QPlaceSearchRequest) QPlaceSearchRequest_PTR() *QPlaceSearchRequest {
	return ptr
}

func (ptr *QPlaceSearchRequest) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QPlaceSearchRequest) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQPlaceSearchRequest(ptr QPlaceSearchRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceSearchRequest_PTR().Pointer()
	}
	return nil
}

func (n *QPlaceSearchRequest) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQPlaceSearchRequestFromPointer(ptr unsafe.Pointer) (n *QPlaceSearchRequest) {
	n = new(QPlaceSearchRequest)
	n.InitFromInternal(uintptr(ptr), "location.QPlaceSearchRequest")
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
	internal.Internal
}

type QPlaceSearchResult_ITF interface {
	QPlaceSearchResult_PTR() *QPlaceSearchResult
}

func (ptr *QPlaceSearchResult) QPlaceSearchResult_PTR() *QPlaceSearchResult {
	return ptr
}

func (ptr *QPlaceSearchResult) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QPlaceSearchResult) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQPlaceSearchResult(ptr QPlaceSearchResult_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceSearchResult_PTR().Pointer()
	}
	return nil
}

func (n *QPlaceSearchResult) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQPlaceSearchResultFromPointer(ptr unsafe.Pointer) (n *QPlaceSearchResult) {
	n = new(QPlaceSearchResult)
	n.InitFromInternal(uintptr(ptr), "location.QPlaceSearchResult")
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

func (n *QPlaceSearchSuggestionReply) InitFromInternal(ptr uintptr, name string) {
	n.QPlaceReply_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QPlaceSearchSuggestionReply) ClassNameInternalF() string {
	return n.QPlaceReply_PTR().ClassNameInternalF()
}

func NewQPlaceSearchSuggestionReplyFromPointer(ptr unsafe.Pointer) (n *QPlaceSearchSuggestionReply) {
	n = new(QPlaceSearchSuggestionReply)
	n.InitFromInternal(uintptr(ptr), "location.QPlaceSearchSuggestionReply")
	return
}

type QPlaceSupplier struct {
	internal.Internal
}

type QPlaceSupplier_ITF interface {
	QPlaceSupplier_PTR() *QPlaceSupplier
}

func (ptr *QPlaceSupplier) QPlaceSupplier_PTR() *QPlaceSupplier {
	return ptr
}

func (ptr *QPlaceSupplier) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QPlaceSupplier) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQPlaceSupplier(ptr QPlaceSupplier_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceSupplier_PTR().Pointer()
	}
	return nil
}

func (n *QPlaceSupplier) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQPlaceSupplierFromPointer(ptr unsafe.Pointer) (n *QPlaceSupplier) {
	n = new(QPlaceSupplier)
	n.InitFromInternal(uintptr(ptr), "location.QPlaceSupplier")
	return
}

type QPlaceUser struct {
	internal.Internal
}

type QPlaceUser_ITF interface {
	QPlaceUser_PTR() *QPlaceUser
}

func (ptr *QPlaceUser) QPlaceUser_PTR() *QPlaceUser {
	return ptr
}

func (ptr *QPlaceUser) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QPlaceUser) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQPlaceUser(ptr QPlaceUser_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceUser_PTR().Pointer()
	}
	return nil
}

func (n *QPlaceUser) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQPlaceUserFromPointer(ptr unsafe.Pointer) (n *QPlaceUser) {
	n = new(QPlaceUser)
	n.InitFromInternal(uintptr(ptr), "location.QPlaceUser")
	return
}
func init() {
	internal.ConstructorTable["location.QGeoManeuver"] = NewQGeoManeuverFromPointer
	internal.ConstructorTable["location.QGeoRoute"] = NewQGeoRouteFromPointer
	internal.ConstructorTable["location.QGeoRouteLeg"] = NewQGeoRouteLegFromPointer
	internal.ConstructorTable["location.QGeoRouteReply"] = NewQGeoRouteReplyFromPointer
	internal.ConstructorTable["location.QGeoRouteRequest"] = NewQGeoRouteRequestFromPointer
	internal.ConstructorTable["location.QGeoRouteSegment"] = NewQGeoRouteSegmentFromPointer
	internal.ConstructorTable["location.QGeoRoutingManager"] = NewQGeoRoutingManagerFromPointer
	internal.ConstructorTable["location.QGeoRoutingManagerEngine"] = NewQGeoRoutingManagerEngineFromPointer
	internal.ConstructorTable["location.QGeoServiceProvider"] = NewQGeoServiceProviderFromPointer
	internal.ConstructorTable["location.QGeoServiceProviderFactory"] = NewQGeoServiceProviderFactoryFromPointer
	internal.ConstructorTable["location.QLocation"] = NewQLocationFromPointer
}
