#pragma once

#ifndef GO_QTTESTLIB_H
#define GO_QTTESTLIB_H

#ifdef __cplusplus
extern "C" {
#endif

void* QSignalSpy_NewQSignalSpy(void* object, char* sign);
int QSignalSpy_IsValid(void* ptr);
char* QSignalSpy_Signal(void* ptr);
int QSignalSpy_Wait(void* ptr, int timeout);
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
int QSignalSpy_Event(void* ptr, void* e);
int QSignalSpy_EventDefault(void* ptr, void* e);
int QSignalSpy_EventFilter(void* ptr, void* watched, void* event);
int QSignalSpy_EventFilterDefault(void* ptr, void* watched, void* event);
void* QSignalSpy_MetaObject(void* ptr);
void* QSignalSpy_MetaObjectDefault(void* ptr);
void* QTestEventList_NewQTestEventList();
void* QTestEventList_NewQTestEventList2(void* other);
void QTestEventList_AddDelay(void* ptr, int msecs);
void QTestEventList_AddKeyClick(void* ptr, int qtKey, int modifiers, int msecs);
void QTestEventList_AddKeyClick2(void* ptr, char* ascii, int modifiers, int msecs);
void QTestEventList_AddKeyClicks(void* ptr, char* keys, int modifiers, int msecs);
void QTestEventList_AddKeyPress(void* ptr, int qtKey, int modifiers, int msecs);
void QTestEventList_AddKeyPress2(void* ptr, char* ascii, int modifiers, int msecs);
void QTestEventList_AddKeyRelease(void* ptr, int qtKey, int modifiers, int msecs);
void QTestEventList_AddKeyRelease2(void* ptr, char* ascii, int modifiers, int msecs);
void QTestEventList_AddMouseClick(void* ptr, int button, int modifiers, void* pos, int delay);
void QTestEventList_AddMouseDClick(void* ptr, int button, int modifiers, void* pos, int delay);
void QTestEventList_AddMouseMove(void* ptr, void* pos, int delay);
void QTestEventList_AddMousePress(void* ptr, int button, int modifiers, void* pos, int delay);
void QTestEventList_AddMouseRelease(void* ptr, int button, int modifiers, void* pos, int delay);
void QTestEventList_Clear(void* ptr);
void QTestEventList_Simulate(void* ptr, void* w);
void QTestEventList_DestroyQTestEventList(void* ptr);

#ifdef __cplusplus
}
#endif

#endif