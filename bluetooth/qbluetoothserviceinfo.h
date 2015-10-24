#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QBluetoothServiceInfo_ServiceName_Type();
int QBluetoothServiceInfo_ServiceDescription_Type();
int QBluetoothServiceInfo_ServiceProvider_Type();
char* QBluetoothServiceInfo_ServiceDescription(QtObjectPtr ptr);
char* QBluetoothServiceInfo_ServiceName(QtObjectPtr ptr);
char* QBluetoothServiceInfo_ServiceProvider(QtObjectPtr ptr);
void QBluetoothServiceInfo_SetServiceDescription(QtObjectPtr ptr, char* description);
void QBluetoothServiceInfo_SetServiceName(QtObjectPtr ptr, char* name);
void QBluetoothServiceInfo_SetServiceProvider(QtObjectPtr ptr, char* provider);
void QBluetoothServiceInfo_SetServiceUuid(QtObjectPtr ptr, QtObjectPtr uuid);
QtObjectPtr QBluetoothServiceInfo_NewQBluetoothServiceInfo();
QtObjectPtr QBluetoothServiceInfo_NewQBluetoothServiceInfo2(QtObjectPtr other);
int QBluetoothServiceInfo_IsComplete(QtObjectPtr ptr);
int QBluetoothServiceInfo_IsRegistered(QtObjectPtr ptr);
int QBluetoothServiceInfo_IsValid(QtObjectPtr ptr);
int QBluetoothServiceInfo_ProtocolServiceMultiplexer(QtObjectPtr ptr);
int QBluetoothServiceInfo_RegisterService(QtObjectPtr ptr, QtObjectPtr localAdapter);
int QBluetoothServiceInfo_ServerChannel(QtObjectPtr ptr);
void QBluetoothServiceInfo_SetDevice(QtObjectPtr ptr, QtObjectPtr device);
int QBluetoothServiceInfo_SocketProtocol(QtObjectPtr ptr);
int QBluetoothServiceInfo_UnregisterService(QtObjectPtr ptr);
void QBluetoothServiceInfo_DestroyQBluetoothServiceInfo(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif