// +build !minimal

#pragma once

#ifndef GO_QTGAMEPAD_H
#define GO_QTGAMEPAD_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtGamepad_PackedString { char* data; long long len; };
struct QtGamepad_PackedList { void* data; long long len; };
void* QGamepad_NewQGamepad(int deviceId, void* parent);
void QGamepad_ConnectAxisLeftXChanged(void* ptr);
void QGamepad_DisconnectAxisLeftXChanged(void* ptr);
void QGamepad_AxisLeftXChanged(void* ptr, double value);
void QGamepad_ConnectAxisLeftYChanged(void* ptr);
void QGamepad_DisconnectAxisLeftYChanged(void* ptr);
void QGamepad_AxisLeftYChanged(void* ptr, double value);
void QGamepad_ConnectAxisRightXChanged(void* ptr);
void QGamepad_DisconnectAxisRightXChanged(void* ptr);
void QGamepad_AxisRightXChanged(void* ptr, double value);
void QGamepad_ConnectAxisRightYChanged(void* ptr);
void QGamepad_DisconnectAxisRightYChanged(void* ptr);
void QGamepad_AxisRightYChanged(void* ptr, double value);
void QGamepad_ConnectButtonAChanged(void* ptr);
void QGamepad_DisconnectButtonAChanged(void* ptr);
void QGamepad_ButtonAChanged(void* ptr, char value);
void QGamepad_ConnectButtonBChanged(void* ptr);
void QGamepad_DisconnectButtonBChanged(void* ptr);
void QGamepad_ButtonBChanged(void* ptr, char value);
void QGamepad_ConnectButtonCenterChanged(void* ptr);
void QGamepad_DisconnectButtonCenterChanged(void* ptr);
void QGamepad_ButtonCenterChanged(void* ptr, char value);
void QGamepad_ConnectButtonDownChanged(void* ptr);
void QGamepad_DisconnectButtonDownChanged(void* ptr);
void QGamepad_ButtonDownChanged(void* ptr, char value);
void QGamepad_ConnectButtonGuideChanged(void* ptr);
void QGamepad_DisconnectButtonGuideChanged(void* ptr);
void QGamepad_ButtonGuideChanged(void* ptr, char value);
void QGamepad_ConnectButtonL1Changed(void* ptr);
void QGamepad_DisconnectButtonL1Changed(void* ptr);
void QGamepad_ButtonL1Changed(void* ptr, char value);
void QGamepad_ConnectButtonL2Changed(void* ptr);
void QGamepad_DisconnectButtonL2Changed(void* ptr);
void QGamepad_ButtonL2Changed(void* ptr, double value);
void QGamepad_ConnectButtonL3Changed(void* ptr);
void QGamepad_DisconnectButtonL3Changed(void* ptr);
void QGamepad_ButtonL3Changed(void* ptr, char value);
void QGamepad_ConnectButtonLeftChanged(void* ptr);
void QGamepad_DisconnectButtonLeftChanged(void* ptr);
void QGamepad_ButtonLeftChanged(void* ptr, char value);
void QGamepad_ConnectButtonR1Changed(void* ptr);
void QGamepad_DisconnectButtonR1Changed(void* ptr);
void QGamepad_ButtonR1Changed(void* ptr, char value);
void QGamepad_ConnectButtonR2Changed(void* ptr);
void QGamepad_DisconnectButtonR2Changed(void* ptr);
void QGamepad_ButtonR2Changed(void* ptr, double value);
void QGamepad_ConnectButtonR3Changed(void* ptr);
void QGamepad_DisconnectButtonR3Changed(void* ptr);
void QGamepad_ButtonR3Changed(void* ptr, char value);
void QGamepad_ConnectButtonRightChanged(void* ptr);
void QGamepad_DisconnectButtonRightChanged(void* ptr);
void QGamepad_ButtonRightChanged(void* ptr, char value);
void QGamepad_ConnectButtonSelectChanged(void* ptr);
void QGamepad_DisconnectButtonSelectChanged(void* ptr);
void QGamepad_ButtonSelectChanged(void* ptr, char value);
void QGamepad_ConnectButtonStartChanged(void* ptr);
void QGamepad_DisconnectButtonStartChanged(void* ptr);
void QGamepad_ButtonStartChanged(void* ptr, char value);
void QGamepad_ConnectButtonUpChanged(void* ptr);
void QGamepad_DisconnectButtonUpChanged(void* ptr);
void QGamepad_ButtonUpChanged(void* ptr, char value);
void QGamepad_ConnectButtonXChanged(void* ptr);
void QGamepad_DisconnectButtonXChanged(void* ptr);
void QGamepad_ButtonXChanged(void* ptr, char value);
void QGamepad_ConnectButtonYChanged(void* ptr);
void QGamepad_DisconnectButtonYChanged(void* ptr);
void QGamepad_ButtonYChanged(void* ptr, char value);
void QGamepad_ConnectConnectedChanged(void* ptr);
void QGamepad_DisconnectConnectedChanged(void* ptr);
void QGamepad_ConnectedChanged(void* ptr, char value);
void QGamepad_ConnectDeviceIdChanged(void* ptr);
void QGamepad_DisconnectDeviceIdChanged(void* ptr);
void QGamepad_DeviceIdChanged(void* ptr, int value);
void QGamepad_ConnectNameChanged(void* ptr);
void QGamepad_DisconnectNameChanged(void* ptr);
void QGamepad_NameChanged(void* ptr, char* value);
void QGamepad_SetDeviceId(void* ptr, int number);
void QGamepad_SetDeviceIdDefault(void* ptr, int number);
void QGamepad_DestroyQGamepad(void* ptr);
struct QtGamepad_PackedString QGamepad_Name(void* ptr);
char QGamepad_ButtonA(void* ptr);
char QGamepad_ButtonB(void* ptr);
char QGamepad_ButtonCenter(void* ptr);
char QGamepad_ButtonDown(void* ptr);
char QGamepad_ButtonGuide(void* ptr);
char QGamepad_ButtonL1(void* ptr);
char QGamepad_ButtonL3(void* ptr);
char QGamepad_ButtonLeft(void* ptr);
char QGamepad_ButtonR1(void* ptr);
char QGamepad_ButtonR3(void* ptr);
char QGamepad_ButtonRight(void* ptr);
char QGamepad_ButtonSelect(void* ptr);
char QGamepad_ButtonStart(void* ptr);
char QGamepad_ButtonUp(void* ptr);
char QGamepad_ButtonX(void* ptr);
char QGamepad_ButtonY(void* ptr);
char QGamepad_IsConnected(void* ptr);
double QGamepad_AxisLeftX(void* ptr);
double QGamepad_AxisLeftY(void* ptr);
double QGamepad_AxisRightX(void* ptr);
double QGamepad_AxisRightY(void* ptr);
double QGamepad_ButtonL2(void* ptr);
double QGamepad_ButtonR2(void* ptr);
int QGamepad_DeviceId(void* ptr);
void* QGamepad___dynamicPropertyNames_atList(void* ptr, int i);
void QGamepad___dynamicPropertyNames_setList(void* ptr, void* i);
void* QGamepad___dynamicPropertyNames_newList(void* ptr);
void* QGamepad___findChildren_atList2(void* ptr, int i);
void QGamepad___findChildren_setList2(void* ptr, void* i);
void* QGamepad___findChildren_newList2(void* ptr);
void* QGamepad___findChildren_atList3(void* ptr, int i);
void QGamepad___findChildren_setList3(void* ptr, void* i);
void* QGamepad___findChildren_newList3(void* ptr);
void* QGamepad___findChildren_atList(void* ptr, int i);
void QGamepad___findChildren_setList(void* ptr, void* i);
void* QGamepad___findChildren_newList(void* ptr);
void* QGamepad___children_atList(void* ptr, int i);
void QGamepad___children_setList(void* ptr, void* i);
void* QGamepad___children_newList(void* ptr);
char QGamepad_EventDefault(void* ptr, void* e);
char QGamepad_EventFilterDefault(void* ptr, void* watched, void* event);
void QGamepad_ChildEventDefault(void* ptr, void* event);
void QGamepad_ConnectNotifyDefault(void* ptr, void* sign);
void QGamepad_CustomEventDefault(void* ptr, void* event);
void QGamepad_DeleteLaterDefault(void* ptr);
void QGamepad_DisconnectNotifyDefault(void* ptr, void* sign);
void QGamepad_TimerEventDefault(void* ptr, void* event);
void* QGamepad_MetaObjectDefault(void* ptr);

#ifdef __cplusplus
}
#endif

#endif