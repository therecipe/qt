#ifdef __cplusplus
extern "C" {
#endif

int QBluetoothServiceInfo_ServiceName_Type();
int QBluetoothServiceInfo_ServiceDescription_Type();
int QBluetoothServiceInfo_ServiceProvider_Type();
char* QBluetoothServiceInfo_ServiceDescription(void* ptr);
char* QBluetoothServiceInfo_ServiceName(void* ptr);
char* QBluetoothServiceInfo_ServiceProvider(void* ptr);
void QBluetoothServiceInfo_SetServiceDescription(void* ptr, char* description);
void QBluetoothServiceInfo_SetServiceName(void* ptr, char* name);
void QBluetoothServiceInfo_SetServiceProvider(void* ptr, char* provider);
void QBluetoothServiceInfo_SetServiceUuid(void* ptr, void* uuid);
void* QBluetoothServiceInfo_NewQBluetoothServiceInfo();
void* QBluetoothServiceInfo_NewQBluetoothServiceInfo2(void* other);
int QBluetoothServiceInfo_IsComplete(void* ptr);
int QBluetoothServiceInfo_IsRegistered(void* ptr);
int QBluetoothServiceInfo_IsValid(void* ptr);
int QBluetoothServiceInfo_ProtocolServiceMultiplexer(void* ptr);
int QBluetoothServiceInfo_RegisterService(void* ptr, void* localAdapter);
int QBluetoothServiceInfo_ServerChannel(void* ptr);
void QBluetoothServiceInfo_SetDevice(void* ptr, void* device);
int QBluetoothServiceInfo_SocketProtocol(void* ptr);
int QBluetoothServiceInfo_UnregisterService(void* ptr);
void QBluetoothServiceInfo_DestroyQBluetoothServiceInfo(void* ptr);

#ifdef __cplusplus
}
#endif