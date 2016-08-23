// +build !minimal

#pragma once

#ifndef GO_QTTESTLIB_H
#define GO_QTTESTLIB_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

void* QSignalSpy_NewQSignalSpy(void* object, char* sign);
char QSignalSpy_IsValid(void* ptr);
char* QSignalSpy_Signal(void* ptr);
char QSignalSpy_Wait(void* ptr, int timeout);
void QSignalSpy_TimerEvent(void* ptr, void* event);
void QSignalSpy_TimerEventDefault(void* ptr, void* event);
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
char QSignalSpy_Event(void* ptr, void* e);
char QSignalSpy_EventDefault(void* ptr, void* e);
char QSignalSpy_EventFilter(void* ptr, void* watched, void* event);
char QSignalSpy_EventFilterDefault(void* ptr, void* watched, void* event);
void* QSignalSpy_MetaObject(void* ptr);
void* QSignalSpy_MetaObjectDefault(void* ptr);
void* QTestEventList_NewQTestEventList();
void* QTestEventList_NewQTestEventList2(void* other);
void QTestEventList_AddDelay(void* ptr, int msecs);
void QTestEventList_AddKeyClick(void* ptr, long long qtKey, long long modifiers, int msecs);
void QTestEventList_AddKeyClick2(void* ptr, char* ascii, long long modifiers, int msecs);
void QTestEventList_AddKeyClicks(void* ptr, char* keys, long long modifiers, int msecs);
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