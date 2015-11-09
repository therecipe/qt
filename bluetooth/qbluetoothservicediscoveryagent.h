#ifdef __cplusplus
extern "C" {
#endif

void QBluetoothServiceDiscoveryAgent_ConnectCanceled(void* ptr);
void QBluetoothServiceDiscoveryAgent_DisconnectCanceled(void* ptr);
void QBluetoothServiceDiscoveryAgent_ConnectFinished(void* ptr);
void QBluetoothServiceDiscoveryAgent_DisconnectFinished(void* ptr);
void* QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent(void* parent);
void* QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent2(void* deviceAdapter, void* parent);
void QBluetoothServiceDiscoveryAgent_Clear(void* ptr);
int QBluetoothServiceDiscoveryAgent_Error(void* ptr);
char* QBluetoothServiceDiscoveryAgent_ErrorString(void* ptr);
int QBluetoothServiceDiscoveryAgent_IsActive(void* ptr);
int QBluetoothServiceDiscoveryAgent_SetRemoteAddress(void* ptr, void* address);
void QBluetoothServiceDiscoveryAgent_SetUuidFilter2(void* ptr, void* uuid);
void QBluetoothServiceDiscoveryAgent_Start(void* ptr, int mode);
void QBluetoothServiceDiscoveryAgent_Stop(void* ptr);
void QBluetoothServiceDiscoveryAgent_DestroyQBluetoothServiceDiscoveryAgent(void* ptr);

#ifdef __cplusplus
}
#endif