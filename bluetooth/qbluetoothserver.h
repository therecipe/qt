#ifdef __cplusplus
extern "C" {
#endif

void* QBluetoothServer_NewQBluetoothServer(int serverType, void* parent);
void QBluetoothServer_ConnectNewConnection(void* ptr);
void QBluetoothServer_DisconnectNewConnection(void* ptr);
int QBluetoothServer_Error(void* ptr);
int QBluetoothServer_IsListening(void* ptr);
int QBluetoothServer_MaxPendingConnections(void* ptr);
int QBluetoothServer_ServerType(void* ptr);
void QBluetoothServer_DestroyQBluetoothServer(void* ptr);
void QBluetoothServer_Close(void* ptr);
int QBluetoothServer_HasPendingConnections(void* ptr);
void* QBluetoothServer_NextPendingConnection(void* ptr);
void QBluetoothServer_SetMaxPendingConnections(void* ptr, int numConnections);

#ifdef __cplusplus
}
#endif