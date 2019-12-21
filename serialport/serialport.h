// +build !minimal

#pragma once

#ifndef GO_QTSERIALPORT_H
#define GO_QTSERIALPORT_H

#include <stdint.h>

#ifdef __cplusplus
int QSerialPort_QSerialPort_QRegisterMetaType();
extern "C" {
#endif

struct QtSerialPort_PackedString { char* data; long long len; void* ptr; };
struct QtSerialPort_PackedList { void* data; long long len; };
void* QSerialPort_NewQSerialPort(void* parent);
void* QSerialPort_NewQSerialPort2(struct QtSerialPort_PackedString name, void* parent);
void* QSerialPort_NewQSerialPort3(void* serialPortInfo, void* parent);
char QSerialPort_AtEndDefault(void* ptr);
int QSerialPort_BaudRate(void* ptr, long long directions);
void QSerialPort_ConnectBaudRateChanged(void* ptr, long long t);
void QSerialPort_DisconnectBaudRateChanged(void* ptr);
void QSerialPort_BaudRateChanged(void* ptr, int baudRate, long long directions);
void QSerialPort_ConnectBreakEnabledChanged(void* ptr, long long t);
void QSerialPort_DisconnectBreakEnabledChanged(void* ptr);
void QSerialPort_BreakEnabledChanged(void* ptr, char set);
long long QSerialPort_BytesAvailableDefault(void* ptr);
long long QSerialPort_BytesToWriteDefault(void* ptr);
char QSerialPort_CanReadLineDefault(void* ptr);
char QSerialPort_Clear(void* ptr, long long directions);
void QSerialPort_ClearError(void* ptr);
void QSerialPort_CloseDefault(void* ptr);
long long QSerialPort_DataBits(void* ptr);
void QSerialPort_ConnectDataBitsChanged(void* ptr, long long t);
void QSerialPort_DisconnectDataBitsChanged(void* ptr);
void QSerialPort_DataBitsChanged(void* ptr, long long dataBits);
void QSerialPort_ConnectDataTerminalReadyChanged(void* ptr, long long t);
void QSerialPort_DisconnectDataTerminalReadyChanged(void* ptr);
void QSerialPort_DataTerminalReadyChanged(void* ptr, char set);
long long QSerialPort_Error(void* ptr);
void QSerialPort_ConnectErrorOccurred(void* ptr, long long t);
void QSerialPort_DisconnectErrorOccurred(void* ptr);
void QSerialPort_ErrorOccurred(void* ptr, long long error);
long long QSerialPort_FlowControl(void* ptr);
void QSerialPort_ConnectFlowControlChanged(void* ptr, long long t);
void QSerialPort_DisconnectFlowControlChanged(void* ptr);
void QSerialPort_FlowControlChanged(void* ptr, long long flow);
char QSerialPort_Flush(void* ptr);
char QSerialPort_IsBreakEnabled(void* ptr);
char QSerialPort_IsDataTerminalReady(void* ptr);
char QSerialPort_IsRequestToSend(void* ptr);
char QSerialPort_IsSequentialDefault(void* ptr);
char QSerialPort_OpenDefault(void* ptr, long long mode);
long long QSerialPort_Parity(void* ptr);
void QSerialPort_ConnectParityChanged(void* ptr, long long t);
void QSerialPort_DisconnectParityChanged(void* ptr);
void QSerialPort_ParityChanged(void* ptr, long long parity);
long long QSerialPort_PinoutSignals(void* ptr);
struct QtSerialPort_PackedString QSerialPort_PortName(void* ptr);
long long QSerialPort_ReadBufferSize(void* ptr);
long long QSerialPort_ReadData(void* ptr, char* data, long long maxSize);
long long QSerialPort_ReadDataDefault(void* ptr, char* data, long long maxSize);
long long QSerialPort_ReadLineDataDefault(void* ptr, char* data, long long maxSize);
void QSerialPort_ConnectRequestToSendChanged(void* ptr, long long t);
void QSerialPort_DisconnectRequestToSendChanged(void* ptr);
void QSerialPort_RequestToSendChanged(void* ptr, char set);
char QSerialPort_SendBreak(void* ptr, int duration);
char QSerialPort_SetBaudRate(void* ptr, int baudRate, long long directions);
char QSerialPort_SetBreakEnabled(void* ptr, char set);
char QSerialPort_SetDataBits(void* ptr, long long dataBits);
char QSerialPort_SetDataTerminalReady(void* ptr, char set);
char QSerialPort_SetFlowControl(void* ptr, long long flowControl);
char QSerialPort_SetParity(void* ptr, long long parity);
void QSerialPort_SetPort(void* ptr, void* serialPortInfo);
void QSerialPort_SetPortName(void* ptr, struct QtSerialPort_PackedString name);
void QSerialPort_SetReadBufferSize(void* ptr, long long size);
char QSerialPort_SetRequestToSend(void* ptr, char set);
void QSerialPort_SetSettingsRestoredOnClose(void* ptr, char restore);
char QSerialPort_SetStopBits(void* ptr, long long stopBits);
char QSerialPort_SettingsRestoredOnClose(void* ptr);
long long QSerialPort_StopBits(void* ptr);
void QSerialPort_ConnectStopBitsChanged(void* ptr, long long t);
void QSerialPort_DisconnectStopBitsChanged(void* ptr);
void QSerialPort_StopBitsChanged(void* ptr, long long stopBits);
char QSerialPort_WaitForBytesWrittenDefault(void* ptr, int msecs);
char QSerialPort_WaitForReadyReadDefault(void* ptr, int msecs);
long long QSerialPort_WriteData(void* ptr, char* data, long long maxSize);
long long QSerialPort_WriteDataDefault(void* ptr, char* data, long long maxSize);
void QSerialPort_DestroyQSerialPort(void* ptr);
void QSerialPort_DestroyQSerialPortDefault(void* ptr);
void* QSerialPort___children_atList(void* ptr, int i);
void QSerialPort___children_setList(void* ptr, void* i);
void* QSerialPort___children_newList(void* ptr);
void* QSerialPort___dynamicPropertyNames_atList(void* ptr, int i);
void QSerialPort___dynamicPropertyNames_setList(void* ptr, void* i);
void* QSerialPort___dynamicPropertyNames_newList(void* ptr);
void* QSerialPort___findChildren_atList(void* ptr, int i);
void QSerialPort___findChildren_setList(void* ptr, void* i);
void* QSerialPort___findChildren_newList(void* ptr);
void* QSerialPort___findChildren_atList3(void* ptr, int i);
void QSerialPort___findChildren_setList3(void* ptr, void* i);
void* QSerialPort___findChildren_newList3(void* ptr);
long long QSerialPort_PosDefault(void* ptr);
char QSerialPort_ResetDefault(void* ptr);
char QSerialPort_SeekDefault(void* ptr, long long pos);
long long QSerialPort_SizeDefault(void* ptr);
void QSerialPort_ChildEventDefault(void* ptr, void* event);
void QSerialPort_ConnectNotifyDefault(void* ptr, void* sign);
void QSerialPort_CustomEventDefault(void* ptr, void* event);
void QSerialPort_DeleteLaterDefault(void* ptr);
void QSerialPort_DisconnectNotifyDefault(void* ptr, void* sign);
char QSerialPort_EventDefault(void* ptr, void* e);
char QSerialPort_EventFilterDefault(void* ptr, void* watched, void* event);
void* QSerialPort_MetaObjectDefault(void* ptr);
void QSerialPort_TimerEventDefault(void* ptr, void* event);
void* QSerialPortInfo_NewQSerialPortInfo();
void* QSerialPortInfo_NewQSerialPortInfo2(void* port);
void* QSerialPortInfo_NewQSerialPortInfo3(struct QtSerialPort_PackedString name);
void* QSerialPortInfo_NewQSerialPortInfo4(void* other);
struct QtSerialPort_PackedList QSerialPortInfo_QSerialPortInfo_AvailablePorts();
struct QtSerialPort_PackedString QSerialPortInfo_Description(void* ptr);
char QSerialPortInfo_HasProductIdentifier(void* ptr);
char QSerialPortInfo_HasVendorIdentifier(void* ptr);
char QSerialPortInfo_IsBusy(void* ptr);
char QSerialPortInfo_IsNull(void* ptr);
struct QtSerialPort_PackedString QSerialPortInfo_Manufacturer(void* ptr);
struct QtSerialPort_PackedString QSerialPortInfo_PortName(void* ptr);
unsigned short QSerialPortInfo_ProductIdentifier(void* ptr);
struct QtSerialPort_PackedString QSerialPortInfo_SerialNumber(void* ptr);
struct QtSerialPort_PackedList QSerialPortInfo_QSerialPortInfo_StandardBaudRates();
void QSerialPortInfo_Swap(void* ptr, void* other);
struct QtSerialPort_PackedString QSerialPortInfo_SystemLocation(void* ptr);
unsigned short QSerialPortInfo_VendorIdentifier(void* ptr);
void QSerialPortInfo_DestroyQSerialPortInfo(void* ptr);
void* QSerialPortInfo___availablePorts_atList(void* ptr, int i);
void QSerialPortInfo___availablePorts_setList(void* ptr, void* i);
void* QSerialPortInfo___availablePorts_newList(void* ptr);
int QSerialPortInfo___standardBaudRates_atList(void* ptr, int i);
void QSerialPortInfo___standardBaudRates_setList(void* ptr, int i);
void* QSerialPortInfo___standardBaudRates_newList(void* ptr);

#ifdef __cplusplus
}
#endif

#endif