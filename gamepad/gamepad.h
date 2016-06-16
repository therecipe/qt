// +build !minimal

#pragma once

#ifndef GO_QTGAMEPAD_H
#define GO_QTGAMEPAD_H

#ifdef __cplusplus
extern "C" {
#endif

int QGamepad_ButtonA(void* ptr);
int QGamepad_ButtonB(void* ptr);
int QGamepad_ButtonCenter(void* ptr);
int QGamepad_ButtonDown(void* ptr);
int QGamepad_ButtonGuide(void* ptr);
int QGamepad_ButtonL1(void* ptr);
int QGamepad_ButtonL3(void* ptr);
int QGamepad_ButtonLeft(void* ptr);
int QGamepad_ButtonR1(void* ptr);
int QGamepad_ButtonR3(void* ptr);
int QGamepad_ButtonRight(void* ptr);
int QGamepad_ButtonSelect(void* ptr);
int QGamepad_ButtonStart(void* ptr);
int QGamepad_ButtonUp(void* ptr);
int QGamepad_ButtonX(void* ptr);
int QGamepad_ButtonY(void* ptr);
int QGamepad_DeviceId(void* ptr);
int QGamepad_IsConnected(void* ptr);
char* QGamepad_Name(void* ptr);
void QGamepad_SetDeviceId(void* ptr, int number);
void* QGamepad_NewQGamepad(int deviceId, void* parent);
void QGamepad_ConnectButtonAChanged(void* ptr);
void QGamepad_DisconnectButtonAChanged(void* ptr);
void QGamepad_ButtonAChanged(void* ptr, int value);
void QGamepad_ConnectButtonBChanged(void* ptr);
void QGamepad_DisconnectButtonBChanged(void* ptr);
void QGamepad_ButtonBChanged(void* ptr, int value);
void QGamepad_ConnectButtonCenterChanged(void* ptr);
void QGamepad_DisconnectButtonCenterChanged(void* ptr);
void QGamepad_ButtonCenterChanged(void* ptr, int value);
void QGamepad_ConnectButtonDownChanged(void* ptr);
void QGamepad_DisconnectButtonDownChanged(void* ptr);
void QGamepad_ButtonDownChanged(void* ptr, int value);
void QGamepad_ConnectButtonGuideChanged(void* ptr);
void QGamepad_DisconnectButtonGuideChanged(void* ptr);
void QGamepad_ButtonGuideChanged(void* ptr, int value);
void QGamepad_ConnectButtonL1Changed(void* ptr);
void QGamepad_DisconnectButtonL1Changed(void* ptr);
void QGamepad_ButtonL1Changed(void* ptr, int value);
void QGamepad_ConnectButtonL3Changed(void* ptr);
void QGamepad_DisconnectButtonL3Changed(void* ptr);
void QGamepad_ButtonL3Changed(void* ptr, int value);
void QGamepad_ConnectButtonLeftChanged(void* ptr);
void QGamepad_DisconnectButtonLeftChanged(void* ptr);
void QGamepad_ButtonLeftChanged(void* ptr, int value);
void QGamepad_ConnectButtonR1Changed(void* ptr);
void QGamepad_DisconnectButtonR1Changed(void* ptr);
void QGamepad_ButtonR1Changed(void* ptr, int value);
void QGamepad_ConnectButtonR3Changed(void* ptr);
void QGamepad_DisconnectButtonR3Changed(void* ptr);
void QGamepad_ButtonR3Changed(void* ptr, int value);
void QGamepad_ConnectButtonRightChanged(void* ptr);
void QGamepad_DisconnectButtonRightChanged(void* ptr);
void QGamepad_ButtonRightChanged(void* ptr, int value);
void QGamepad_ConnectButtonSelectChanged(void* ptr);
void QGamepad_DisconnectButtonSelectChanged(void* ptr);
void QGamepad_ButtonSelectChanged(void* ptr, int value);
void QGamepad_ConnectButtonStartChanged(void* ptr);
void QGamepad_DisconnectButtonStartChanged(void* ptr);
void QGamepad_ButtonStartChanged(void* ptr, int value);
void QGamepad_ConnectButtonUpChanged(void* ptr);
void QGamepad_DisconnectButtonUpChanged(void* ptr);
void QGamepad_ButtonUpChanged(void* ptr, int value);
void QGamepad_ConnectButtonXChanged(void* ptr);
void QGamepad_DisconnectButtonXChanged(void* ptr);
void QGamepad_ButtonXChanged(void* ptr, int value);
void QGamepad_ConnectButtonYChanged(void* ptr);
void QGamepad_DisconnectButtonYChanged(void* ptr);
void QGamepad_ButtonYChanged(void* ptr, int value);
void QGamepad_ConnectConnectedChanged(void* ptr);
void QGamepad_DisconnectConnectedChanged(void* ptr);
void QGamepad_ConnectedChanged(void* ptr, int value);
void QGamepad_ConnectDeviceIdChanged(void* ptr);
void QGamepad_DisconnectDeviceIdChanged(void* ptr);
void QGamepad_DeviceIdChanged(void* ptr, int value);
void QGamepad_ConnectNameChanged(void* ptr);
void QGamepad_DisconnectNameChanged(void* ptr);
void QGamepad_NameChanged(void* ptr, char* value);
void QGamepad_DestroyQGamepad(void* ptr);
void QGamepad_TimerEvent(void* ptr, void* event);
void QGamepad_TimerEventDefault(void* ptr, void* event);
void QGamepad_ChildEvent(void* ptr, void* event);
void QGamepad_ChildEventDefault(void* ptr, void* event);
void QGamepad_ConnectNotify(void* ptr, void* sign);
void QGamepad_ConnectNotifyDefault(void* ptr, void* sign);
void QGamepad_CustomEvent(void* ptr, void* event);
void QGamepad_CustomEventDefault(void* ptr, void* event);
void QGamepad_DeleteLater(void* ptr);
void QGamepad_DeleteLaterDefault(void* ptr);
void QGamepad_DisconnectNotify(void* ptr, void* sign);
void QGamepad_DisconnectNotifyDefault(void* ptr, void* sign);
int QGamepad_Event(void* ptr, void* e);
int QGamepad_EventDefault(void* ptr, void* e);
int QGamepad_EventFilter(void* ptr, void* watched, void* event);
int QGamepad_EventFilterDefault(void* ptr, void* watched, void* event);
void* QGamepad_MetaObject(void* ptr);
void* QGamepad_MetaObjectDefault(void* ptr);

#ifdef __cplusplus
}
#endif

#endif