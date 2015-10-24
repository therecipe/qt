#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QBluetoothSocket_ConnectConnected(QtObjectPtr ptr);
void QBluetoothSocket_DisconnectConnected(QtObjectPtr ptr);
void QBluetoothSocket_ConnectDisconnected(QtObjectPtr ptr);
void QBluetoothSocket_DisconnectDisconnected(QtObjectPtr ptr);
void QBluetoothSocket_ConnectStateChanged(QtObjectPtr ptr);
void QBluetoothSocket_DisconnectStateChanged(QtObjectPtr ptr);
QtObjectPtr QBluetoothSocket_NewQBluetoothSocket(int socketType, QtObjectPtr parent);
QtObjectPtr QBluetoothSocket_NewQBluetoothSocket2(QtObjectPtr parent);
void QBluetoothSocket_Abort(QtObjectPtr ptr);
int QBluetoothSocket_CanReadLine(QtObjectPtr ptr);
void QBluetoothSocket_Close(QtObjectPtr ptr);
void QBluetoothSocket_ConnectToService2(QtObjectPtr ptr, QtObjectPtr address, QtObjectPtr uuid, int openMode);
void QBluetoothSocket_ConnectToService(QtObjectPtr ptr, QtObjectPtr service, int openMode);
void QBluetoothSocket_DisconnectFromService(QtObjectPtr ptr);
int QBluetoothSocket_Error(QtObjectPtr ptr);
char* QBluetoothSocket_ErrorString(QtObjectPtr ptr);
int QBluetoothSocket_IsSequential(QtObjectPtr ptr);
char* QBluetoothSocket_LocalName(QtObjectPtr ptr);
char* QBluetoothSocket_PeerName(QtObjectPtr ptr);
int QBluetoothSocket_SetSocketDescriptor(QtObjectPtr ptr, int socketDescriptor, int socketType, int socketState, int openMode);
int QBluetoothSocket_SocketDescriptor(QtObjectPtr ptr);
int QBluetoothSocket_SocketType(QtObjectPtr ptr);
void QBluetoothSocket_DestroyQBluetoothSocket(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif