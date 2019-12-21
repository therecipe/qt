// +build !minimal

#pragma once

#ifndef GO_QTTESTLIB_H
#define GO_QTTESTLIB_H

#include <stdint.h>

#ifdef __cplusplus
int QSignalSpy_QSignalSpy_QRegisterMetaType();
extern "C" {
#endif

struct QtTestLib_PackedString { char* data; long long len; void* ptr; };
struct QtTestLib_PackedList { void* data; long long len; };
void* QAbstractItemModelTester_NewQAbstractItemModelTester(void* model, void* parent);
void* QAbstractItemModelTester_NewQAbstractItemModelTester2(void* model, long long mode, void* parent);
void* QAbstractItemModelTester_Model(void* ptr);
void* QSignalSpy_NewQSignalSpy(void* object, char* sign);
char QSignalSpy_IsValid(void* ptr);
void* QSignalSpy_Signal(void* ptr);
char QSignalSpy_Wait(void* ptr, int timeout);
int QSignalSpy___args_atList(void* ptr, int i);
void QSignalSpy___args_setList(void* ptr, int i);
void* QSignalSpy___args_newList(void* ptr);
int QSignalSpy___setArgs__atList(void* ptr, int i);
void QSignalSpy___setArgs__setList(void* ptr, int i);
void* QSignalSpy___setArgs__newList(void* ptr);
void* QSignalSpy___children_atList(void* ptr, int i);
void QSignalSpy___children_setList(void* ptr, void* i);
void* QSignalSpy___children_newList(void* ptr);
void* QSignalSpy___dynamicPropertyNames_atList(void* ptr, int i);
void QSignalSpy___dynamicPropertyNames_setList(void* ptr, void* i);
void* QSignalSpy___dynamicPropertyNames_newList(void* ptr);
void* QSignalSpy___findChildren_atList(void* ptr, int i);
void QSignalSpy___findChildren_setList(void* ptr, void* i);
void* QSignalSpy___findChildren_newList(void* ptr);
void* QSignalSpy___findChildren_atList3(void* ptr, int i);
void QSignalSpy___findChildren_setList3(void* ptr, void* i);
void* QSignalSpy___findChildren_newList3(void* ptr);
void QSignalSpy_ChildEventDefault(void* ptr, void* event);
void QSignalSpy_ConnectNotifyDefault(void* ptr, void* sign);
void QSignalSpy_CustomEventDefault(void* ptr, void* event);
void QSignalSpy_DeleteLaterDefault(void* ptr);
void QSignalSpy_DisconnectNotifyDefault(void* ptr, void* sign);
char QSignalSpy_EventDefault(void* ptr, void* e);
char QSignalSpy_EventFilterDefault(void* ptr, void* watched, void* event);
void* QSignalSpy_MetaObjectDefault(void* ptr);
void QSignalSpy_TimerEventDefault(void* ptr, void* event);
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
void QTestEventList_Clear(void* ptr);
void QTestEventList_Simulate(void* ptr, void* w);
void QTestEventList_DestroyQTestEventList(void* ptr);

#ifdef __cplusplus
}
#endif

#endif