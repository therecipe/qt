#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QBluetoothServer_NewQBluetoothServer(int serverType, QtObjectPtr parent);
void QBluetoothServer_ConnectNewConnection(QtObjectPtr ptr);
void QBluetoothServer_DisconnectNewConnection(QtObjectPtr ptr);
int QBluetoothServer_Error(QtObjectPtr ptr);
int QBluetoothServer_IsListening(QtObjectPtr ptr);
int QBluetoothServer_MaxPendingConnections(QtObjectPtr ptr);
int QBluetoothServer_ServerType(QtObjectPtr ptr);
void QBluetoothServer_DestroyQBluetoothServer(QtObjectPtr ptr);
void QBluetoothServer_Close(QtObjectPtr ptr);
int QBluetoothServer_HasPendingConnections(QtObjectPtr ptr);
QtObjectPtr QBluetoothServer_NextPendingConnection(QtObjectPtr ptr);
void QBluetoothServer_SetMaxPendingConnections(QtObjectPtr ptr, int numConnections);

#ifdef __cplusplus
}
#endif