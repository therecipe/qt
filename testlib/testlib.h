// +build !minimal

#pragma once

#ifndef GO_QTTESTLIB_H
#define GO_QTTESTLIB_H

#include <stdint.h>

#ifdef __cplusplus
int QSignalSpy_QSignalSpy_QRegisterMetaType();
extern "C" {
#endif

struct QtTestLib_PackedString { char* data; long long len; };
struct QtTestLib_PackedList { void* data; long long len; };
void* QSignalSpy_NewQSignalSpy(void* object, char* sign);
char QSignalSpy_Wait(void* ptr, int timeout);
void* QSignalSpy_Signal(void* ptr);
char QSignalSpy_IsValid(void* ptr);
void* QSignalSpy___dynamicPropertyNames_atList(void* ptr, int i);
void QSignalSpy___dynamicPropertyNames_setList(void* ptr, void* i);
void* QSignalSpy___dynamicPropertyNames_newList(void* ptr);
void* QSignalSpy___findChildren_atList2(void* ptr, int i);
void QSignalSpy___findChildren_setList2(void* ptr, void* i);
void* QSignalSpy___findChildren_newList2(void* ptr);
void* QSignalSpy___findChildren_atList3(void* ptr, int i);
void QSignalSpy___findChildren_setList3(void* ptr, void* i);
void* QSignalSpy___findChildren_newList3(void* ptr);
void* QSignalSpy___findChildren_atList(void* ptr, int i);
void QSignalSpy___findChildren_setList(void* ptr, void* i);
void* QSignalSpy___findChildren_newList(void* ptr);
void* QSignalSpy___children_atList(void* ptr, int i);
void QSignalSpy___children_setList(void* ptr, void* i);
void* QSignalSpy___children_newList(void* ptr);
void* QSignalSpy___QList_other_atList3(void* ptr, int i);
void QSignalSpy___QList_other_setList3(void* ptr, void* i);
void* QSignalSpy___QList_other_newList3(void* ptr);
void* QSignalSpy___QList_other_atList2(void* ptr, int i);
void QSignalSpy___QList_other_setList2(void* ptr, void* i);
void* QSignalSpy___QList_other_newList2(void* ptr);
void* QSignalSpy___fromSet_atList(void* ptr, int i);
void QSignalSpy___fromSet_setList(void* ptr, void* i);
void* QSignalSpy___fromSet_newList(void* ptr);
void* QSignalSpy___fromStdList_atList(void* ptr, int i);
void QSignalSpy___fromStdList_setList(void* ptr, void* i);
void* QSignalSpy___fromStdList_newList(void* ptr);
void* QSignalSpy___fromVector_atList(void* ptr, int i);
void QSignalSpy___fromVector_setList(void* ptr, void* i);
void* QSignalSpy___fromVector_newList(void* ptr);
void* QSignalSpy___fromVector_vector_atList(void* ptr, int i);
void QSignalSpy___fromVector_vector_setList(void* ptr, void* i);
void* QSignalSpy___fromVector_vector_newList(void* ptr);
void* QSignalSpy___append_value_atList2(void* ptr, int i);
void QSignalSpy___append_value_setList2(void* ptr, void* i);
void* QSignalSpy___append_value_newList2(void* ptr);
void* QSignalSpy___swap_other_atList(void* ptr, int i);
void QSignalSpy___swap_other_setList(void* ptr, void* i);
void* QSignalSpy___swap_other_newList(void* ptr);
void* QSignalSpy___mid_atList(void* ptr, int i);
void QSignalSpy___mid_setList(void* ptr, void* i);
void* QSignalSpy___mid_newList(void* ptr);
void* QSignalSpy___toVector_atList(void* ptr, int i);
void QSignalSpy___toVector_setList(void* ptr, void* i);
void* QSignalSpy___toVector_newList(void* ptr);
char QSignalSpy_Event(void* ptr, void* e);
char QSignalSpy_EventDefault(void* ptr, void* e);
char QSignalSpy_EventFilter(void* ptr, void* watched, void* event);
char QSignalSpy_EventFilterDefault(void* ptr, void* watched, void* event);
void QSignalSpy_ChildEvent(void* ptr, void* event);
void QSignalSpy_ChildEventDefault(void* ptr, void* event);
void QSignalSpy_ConnectNotify(void* ptr, void* sign);
void QSignalSpy_ConnectNotifyDefault(void* ptr, void* sign);
void QSignalSpy_CustomEvent(void* ptr, void* event);
void QSignalSpy_CustomEventDefault(void* ptr, void* event);
void QSignalSpy_DeleteLater(void* ptr);
void QSignalSpy_DeleteLaterDefault(void* ptr);
void QSignalSpy_DisconnectNotify(void* ptr, void* sign);
void QSignalSpy_DisconnectNotifyDefault(void* ptr, void* sign);
void QSignalSpy_TimerEvent(void* ptr, void* event);
void QSignalSpy_TimerEventDefault(void* ptr, void* event);
void* QSignalSpy_MetaObject(void* ptr);
void* QSignalSpy_MetaObjectDefault(void* ptr);
void* QTestEventList_NewQTestEventList();
void* QTestEventList_NewQTestEventList2(void* other);
void QTestEventList_AddDelay(void* ptr, int msecs);
void QTestEventList_AddKeyClick(void* ptr, long long qtKey, long long modifiers, int msecs);
void QTestEventList_AddKeyClick2(void* ptr, char* ascii, long long modifiers, int msecs);
void QTestEventList_AddKeyClicks(void* ptr, struct QtTestLib_PackedString keys, long long modifiers, int msecs);
void QTestEventList_AddKeyPress(void* ptr, long long qtKey, long long modifiers, int msecs);
void QTestEventList_AddKeyPress2(void* ptr, char* ascii, long long modifiers, int msecs);
void QTestEventList_AddKeyRelease(void* ptr, long long qtKey, long long modifiers, int msecs);
void QTestEventList_AddKeyRelease2(void* ptr, char* ascii, long long modifiers, int msecs);
void QTestEventList_AddMouseClick(void* ptr, long long button, long long modifiers, void* pos, int delay);
void QTestEventList_AddMouseDClick(void* ptr, long long button, long long modifiers, void* pos, int delay);
void QTestEventList_AddMouseMove(void* ptr, void* pos, int delay);
void QTestEventList_AddMousePress(void* ptr, long long button, long long modifiers, void* pos, int delay);
void QTestEventList_AddMouseRelease(void* ptr, long long button, long long modifiers, void* pos, int delay);
void QTestEventList_Simulate(void* ptr, void* w);
void QTestEventList_DestroyQTestEventList(void* ptr);
void* QTestEventList___QList_other_atList3(void* ptr, int i);
void QTestEventList___QList_other_setList3(void* ptr, void* i);
void* QTestEventList___QList_other_newList3(void* ptr);
void* QTestEventList___QList_other_atList2(void* ptr, int i);
void QTestEventList___QList_other_setList2(void* ptr, void* i);
void* QTestEventList___QList_other_newList2(void* ptr);
void* QTestEventList___fromSet_atList(void* ptr, int i);
void QTestEventList___fromSet_setList(void* ptr, void* i);
void* QTestEventList___fromSet_newList(void* ptr);
void* QTestEventList___fromStdList_atList(void* ptr, int i);
void QTestEventList___fromStdList_setList(void* ptr, void* i);
void* QTestEventList___fromStdList_newList(void* ptr);
void* QTestEventList___fromVector_atList(void* ptr, int i);
void QTestEventList___fromVector_setList(void* ptr, void* i);
void* QTestEventList___fromVector_newList(void* ptr);
void* QTestEventList___fromVector_vector_atList(void* ptr, int i);
void QTestEventList___fromVector_vector_setList(void* ptr, void* i);
void* QTestEventList___fromVector_vector_newList(void* ptr);
void* QTestEventList___append_value_atList2(void* ptr, int i);
void QTestEventList___append_value_setList2(void* ptr, void* i);
void* QTestEventList___append_value_newList2(void* ptr);
void* QTestEventList___swap_other_atList(void* ptr, int i);
void QTestEventList___swap_other_setList(void* ptr, void* i);
void* QTestEventList___swap_other_newList(void* ptr);
void* QTestEventList___mid_atList(void* ptr, int i);
void QTestEventList___mid_setList(void* ptr, void* i);
void* QTestEventList___mid_newList(void* ptr);
void* QTestEventList___toVector_atList(void* ptr, int i);
void QTestEventList___toVector_setList(void* ptr, void* i);
void* QTestEventList___toVector_newList(void* ptr);

#ifdef __cplusplus
}
#endif

#endif