#ifdef __cplusplus
extern "C" {
#endif

void QBluetoothSocket_ConnectConnected(void* ptr);
void QBluetoothSocket_DisconnectConnected(void* ptr);
void QBluetoothSocket_ConnectDisconnected(void* ptr);
void QBluetoothSocket_DisconnectDisconnected(void* ptr);
void QBluetoothSocket_ConnectStateChanged(void* ptr);
void QBluetoothSocket_DisconnectStateChanged(void* ptr);
void* QBluetoothSocket_NewQBluetoothSocket(int socketType, void* parent);
void* QBluetoothSocket_NewQBluetoothSocket2(void* parent);
void QBluetoothSocket_Abort(void* ptr);
int QBluetoothSocket_CanReadLine(void* ptr);
void QBluetoothSocket_Close(void* ptr);
void QBluetoothSocket_ConnectToService2(void* ptr, void* address, void* uuid, int openMode);
void QBluetoothSocket_ConnectToService(void* ptr, void* service, int openMode);
void QBluetoothSocket_DisconnectFromService(void* ptr);
int QBluetoothSocket_Error(void* ptr);
char* QBluetoothSocket_ErrorString(void* ptr);
int QBluetoothSocket_IsSequential(void* ptr);
char* QBluetoothSocket_LocalName(void* ptr);
char* QBluetoothSocket_PeerName(void* ptr);
int QBluetoothSocket_SetSocketDescriptor(void* ptr, int socketDescriptor, int socketType, int socketState, int openMode);
int QBluetoothSocket_SocketDescriptor(void* ptr);
int QBluetoothSocket_SocketType(void* ptr);
void QBluetoothSocket_DestroyQBluetoothSocket(void* ptr);

#ifdef __cplusplus
}
#endif