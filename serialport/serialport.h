// +build !minimal

#pragma once

#ifndef GO_QTSERIALPORT_H
#define GO_QTSERIALPORT_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtSerialPort_PackedString { char* data; long long len; };
struct QtSerialPort_PackedList { void* data; long long len; };
int QSerialPort_BaudRate(void* ptr, long long directions);
void QSerialPort_ClearError(void* ptr);
long long QSerialPort_DataBits(void* ptr);
long long QSerialPort_Error(void* ptr);
long long QSerialPort_FlowControl(void* ptr);
char QSerialPort_IsBreakEnabled(void* ptr);
char QSerialPort_IsDataTerminalReady(void* ptr);
char QSerialPort_IsRequestToSend(void* ptr);
long long QSerialPort_Parity(void* ptr);
char QSerialPort_SetBaudRate(void* ptr, int baudRate, long long directions);
char QSerialPort_SetBreakEnabled(void* ptr, char set);
char QSerialPort_SetDataBits(void* ptr, long long dataBits);
char QSerialPort_SetDataTerminalReady(void* ptr, char set);
char QSerialPort_SetFlowControl(void* ptr, long long flowControl);
char QSerialPort_SetParity(void* ptr, long long parity);
char QSerialPort_SetRequestToSend(void* ptr, char set);
char QSerialPort_SetStopBits(void* ptr, long long stopBits);
long long QSerialPort_StopBits(void* ptr);
void* QSerialPort_NewQSerialPort(void* parent);
void* QSerialPort_NewQSerialPort3(void* serialPortInfo, void* parent);
void* QSerialPort_NewQSerialPort2(char* name, void* parent);
char QSerialPort_AtEnd(void* ptr);
void QSerialPort_ConnectBaudRateChanged(void* ptr);
void QSerialPort_DisconnectBaudRateChanged(void* ptr);
void QSerialPort_BaudRateChanged(void* ptr, int baudRate, long long directions);
void QSerialPort_ConnectBreakEnabledChanged(void* ptr);
void QSerialPort_DisconnectBreakEnabledChanged(void* ptr);
void QSerialPort_BreakEnabledChanged(void* ptr, char set);
long long QSerialPort_BytesAvailable(void* ptr);
long long QSerialPort_BytesToWrite(void* ptr);
char QSerialPort_CanReadLine(void* ptr);
char QSerialPort_Clear(void* ptr, long long directions);
void QSerialPort_Close(void* ptr);
void QSerialPort_ConnectDataBitsChanged(void* ptr);
void QSerialPort_DisconnectDataBitsChanged(void* ptr);
void QSerialPort_DataBitsChanged(void* ptr, long long dataBits);
void QSerialPort_ConnectDataTerminalReadyChanged(void* ptr);
void QSerialPort_DisconnectDataTerminalReadyChanged(void* ptr);
void QSerialPort_DataTerminalReadyChanged(void* ptr, char set);
void QSerialPort_ConnectError2(void* ptr);
void QSerialPort_DisconnectError2(void* ptr);
void QSerialPort_Error2(void* ptr, long long error);
void QSerialPort_ConnectFlowControlChanged(void* ptr);
void QSerialPort_DisconnectFlowControlChanged(void* ptr);
void QSerialPort_FlowControlChanged(void* ptr, long long flow);
char QSerialPort_Flush(void* ptr);
char QSerialPort_IsSequential(void* ptr);
char QSerialPort_Open(void* ptr, long long mode);
void QSerialPort_ConnectParityChanged(void* ptr);
void QSerialPort_DisconnectParityChanged(void* ptr);
void QSerialPort_ParityChanged(void* ptr, long long parity);
long long QSerialPort_PinoutSignals(void* ptr);
struct QtSerialPort_PackedString QSerialPort_PortName(void* ptr);
long long QSerialPort_ReadBufferSize(void* ptr);
long long QSerialPort_ReadLineData(void* ptr, char* data, long long maxSize);
void QSerialPort_ConnectRequestToSendChanged(void* ptr);
void QSerialPort_DisconnectRequestToSendChanged(void* ptr);
void QSerialPort_RequestToSendChanged(void* ptr, char set);
void QSerialPort_SetPort(void* ptr, void* serialPortInfo);
void QSerialPort_SetPortName(void* ptr, char* name);
void QSerialPort_SetReadBufferSize(void* ptr, long long size);
void QSerialPort_ConnectStopBitsChanged(void* ptr);
void QSerialPort_DisconnectStopBitsChanged(void* ptr);
void QSerialPort_StopBitsChanged(void* ptr, long long stopBits);
char QSerialPort_WaitForBytesWritten(void* ptr, int msecs);
char QSerialPort_WaitForReadyRead(void* ptr, int msecs);
long long QSerialPort_WriteData(void* ptr, char* data, long long maxSize);
void QSerialPort_DestroyQSerialPort(void* ptr);
void QSerialPort_DestroyQSerialPortDefault(void* ptr);
long long QSerialPort_Pos(void* ptr);
long long QSerialPort_PosDefault(void* ptr);
char QSerialPort_Reset(void* ptr);
char QSerialPort_ResetDefault(void* ptr);
char QSerialPort_Seek(void* ptr, long long pos);
char QSerialPort_SeekDefault(void* ptr, long long pos);
long long QSerialPort_Size(void* ptr);
long long QSerialPort_SizeDefault(void* ptr);
void QSerialPort_TimerEvent(void* ptr, void* event);
void QSerialPort_TimerEventDefault(void* ptr, void* event);
void QSerialPort_ChildEvent(void* ptr, void* event);
void QSerialPort_ChildEventDefault(void* ptr, void* event);
void QSerialPort_ConnectNotify(void* ptr, void* sign);
void QSerialPort_ConnectNotifyDefault(void* ptr, void* sign);
void QSerialPort_CustomEvent(void* ptr, void* event);
void QSerialPort_CustomEventDefault(void* ptr, void* event);
void QSerialPort_DeleteLater(void* ptr);
void QSerialPort_DeleteLaterDefault(void* ptr);
void QSerialPort_DisconnectNotify(void* ptr, void* sign);
void QSerialPort_DisconnectNotifyDefault(void* ptr, void* sign);
char QSerialPort_Event(void* ptr, void* e);
char QSerialPort_EventDefault(void* ptr, void* e);
char QSerialPort_EventFilter(void* ptr, void* watched, void* event);
char QSerialPort_EventFilterDefault(void* ptr, void* watched, void* event);
void* QSerialPort_MetaObject(void* ptr);
void* QSerialPort_MetaObjectDefault(void* ptr);
void QSerialPortInfo_Swap(void* ptr, void* other);
void* QSerialPortInfo_NewQSerialPortInfo();
void* QSerialPortInfo_NewQSerialPortInfo2(void* port);
void* QSerialPortInfo_NewQSerialPortInfo4(void* other);
void* QSerialPortInfo_NewQSerialPortInfo3(char* name);
struct QtSerialPort_PackedString QSerialPortInfo_Description(void* ptr);
char QSerialPortInfo_HasProductIdentifier(void* ptr);
char QSerialPortInfo_HasVendorIdentifier(void* ptr);
char QSerialPortInfo_IsNull(void* ptr);
struct QtSerialPort_PackedString QSerialPortInfo_Manufacturer(void* ptr);
struct QtSerialPort_PackedString QSerialPortInfo_PortName(void* ptr);
unsigned short QSerialPortInfo_ProductIdentifier(void* ptr);
struct QtSerialPort_PackedString QSerialPortInfo_SerialNumber(void* ptr);
struct QtSerialPort_PackedString QSerialPortInfo_SystemLocation(void* ptr);
unsigned short QSerialPortInfo_VendorIdentifier(void* ptr);
void QSerialPortInfo_DestroyQSerialPortInfo(void* ptr);
struct QtSerialPort_PackedList QSerialPortInfo_QSerialPortInfo_AvailablePorts();
char QSerialPortInfo_IsBusy(void* ptr);
void* QSerialPortInfo_availablePorts_atList(void* ptr, int i);

#ifdef __cplusplus
}
#endif

#endif