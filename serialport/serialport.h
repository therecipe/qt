#pragma once

#ifndef GO_QTSERIALPORT_H
#define GO_QTSERIALPORT_H

#ifdef __cplusplus
extern "C" {
#endif

void QSerialPort_ClearError(void* ptr);
int QSerialPort_DataBits(void* ptr);
int QSerialPort_Error(void* ptr);
int QSerialPort_FlowControl(void* ptr);
int QSerialPort_IsBreakEnabled(void* ptr);
int QSerialPort_IsDataTerminalReady(void* ptr);
int QSerialPort_IsRequestToSend(void* ptr);
int QSerialPort_Parity(void* ptr);
int QSerialPort_SetBreakEnabled(void* ptr, int set);
int QSerialPort_SetDataBits(void* ptr, int dataBits);
int QSerialPort_SetDataTerminalReady(void* ptr, int set);
int QSerialPort_SetFlowControl(void* ptr, int flowControl);
int QSerialPort_SetParity(void* ptr, int parity);
int QSerialPort_SetRequestToSend(void* ptr, int set);
int QSerialPort_SetStopBits(void* ptr, int stopBits);
int QSerialPort_StopBits(void* ptr);
void* QSerialPort_NewQSerialPort(void* parent);
void* QSerialPort_NewQSerialPort3(void* serialPortInfo, void* parent);
void* QSerialPort_NewQSerialPort2(char* name, void* parent);
int QSerialPort_AtEnd(void* ptr);
void QSerialPort_ConnectBreakEnabledChanged(void* ptr);
void QSerialPort_DisconnectBreakEnabledChanged(void* ptr);
void QSerialPort_BreakEnabledChanged(void* ptr, int set);
long long QSerialPort_BytesAvailable(void* ptr);
long long QSerialPort_BytesToWrite(void* ptr);
int QSerialPort_CanReadLine(void* ptr);
int QSerialPort_Clear(void* ptr, int directions);
void QSerialPort_Close(void* ptr);
void QSerialPort_ConnectDataBitsChanged(void* ptr);
void QSerialPort_DisconnectDataBitsChanged(void* ptr);
void QSerialPort_DataBitsChanged(void* ptr, int dataBits);
void QSerialPort_ConnectDataTerminalReadyChanged(void* ptr);
void QSerialPort_DisconnectDataTerminalReadyChanged(void* ptr);
void QSerialPort_DataTerminalReadyChanged(void* ptr, int set);
void QSerialPort_ConnectError2(void* ptr);
void QSerialPort_DisconnectError2(void* ptr);
void QSerialPort_Error2(void* ptr, int error);
void QSerialPort_ConnectFlowControlChanged(void* ptr);
void QSerialPort_DisconnectFlowControlChanged(void* ptr);
void QSerialPort_FlowControlChanged(void* ptr, int flow);
int QSerialPort_Flush(void* ptr);
int QSerialPort_IsSequential(void* ptr);
int QSerialPort_Open(void* ptr, int mode);
void QSerialPort_ConnectParityChanged(void* ptr);
void QSerialPort_DisconnectParityChanged(void* ptr);
void QSerialPort_ParityChanged(void* ptr, int parity);
int QSerialPort_PinoutSignals(void* ptr);
char* QSerialPort_PortName(void* ptr);
long long QSerialPort_ReadBufferSize(void* ptr);
long long QSerialPort_ReadLineData(void* ptr, char* data, long long maxSize);
void QSerialPort_ConnectRequestToSendChanged(void* ptr);
void QSerialPort_DisconnectRequestToSendChanged(void* ptr);
void QSerialPort_RequestToSendChanged(void* ptr, int set);
void QSerialPort_SetPort(void* ptr, void* serialPortInfo);
void QSerialPort_SetPortName(void* ptr, char* name);
void QSerialPort_SetReadBufferSize(void* ptr, long long size);
void QSerialPort_ConnectStopBitsChanged(void* ptr);
void QSerialPort_DisconnectStopBitsChanged(void* ptr);
void QSerialPort_StopBitsChanged(void* ptr, int stopBits);
int QSerialPort_WaitForBytesWritten(void* ptr, int msecs);
int QSerialPort_WaitForReadyRead(void* ptr, int msecs);
long long QSerialPort_WriteData(void* ptr, char* data, long long maxSize);
void QSerialPort_DestroyQSerialPort(void* ptr);
long long QSerialPort_Pos(void* ptr);
long long QSerialPort_PosDefault(void* ptr);
int QSerialPort_Reset(void* ptr);
int QSerialPort_ResetDefault(void* ptr);
int QSerialPort_Seek(void* ptr, long long pos);
int QSerialPort_SeekDefault(void* ptr, long long pos);
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
int QSerialPort_Event(void* ptr, void* e);
int QSerialPort_EventDefault(void* ptr, void* e);
int QSerialPort_EventFilter(void* ptr, void* watched, void* event);
int QSerialPort_EventFilterDefault(void* ptr, void* watched, void* event);
void* QSerialPort_MetaObject(void* ptr);
void* QSerialPort_MetaObjectDefault(void* ptr);
void QSerialPortInfo_Swap(void* ptr, void* other);
void* QSerialPortInfo_NewQSerialPortInfo();
void* QSerialPortInfo_NewQSerialPortInfo2(void* port);
void* QSerialPortInfo_NewQSerialPortInfo4(void* other);
void* QSerialPortInfo_NewQSerialPortInfo3(char* name);
char* QSerialPortInfo_Description(void* ptr);
int QSerialPortInfo_HasProductIdentifier(void* ptr);
int QSerialPortInfo_HasVendorIdentifier(void* ptr);
int QSerialPortInfo_IsNull(void* ptr);
char* QSerialPortInfo_Manufacturer(void* ptr);
char* QSerialPortInfo_PortName(void* ptr);
char* QSerialPortInfo_SerialNumber(void* ptr);
char* QSerialPortInfo_SystemLocation(void* ptr);
void QSerialPortInfo_DestroyQSerialPortInfo(void* ptr);
int QSerialPortInfo_IsBusy(void* ptr);

#ifdef __cplusplus
}
#endif

#endif