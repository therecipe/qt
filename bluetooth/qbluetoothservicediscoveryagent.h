#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QBluetoothServiceDiscoveryAgent_ConnectCanceled(QtObjectPtr ptr);
void QBluetoothServiceDiscoveryAgent_DisconnectCanceled(QtObjectPtr ptr);
void QBluetoothServiceDiscoveryAgent_ConnectFinished(QtObjectPtr ptr);
void QBluetoothServiceDiscoveryAgent_DisconnectFinished(QtObjectPtr ptr);
QtObjectPtr QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent(QtObjectPtr parent);
QtObjectPtr QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent2(QtObjectPtr deviceAdapter, QtObjectPtr parent);
void QBluetoothServiceDiscoveryAgent_Clear(QtObjectPtr ptr);
int QBluetoothServiceDiscoveryAgent_Error(QtObjectPtr ptr);
char* QBluetoothServiceDiscoveryAgent_ErrorString(QtObjectPtr ptr);
int QBluetoothServiceDiscoveryAgent_IsActive(QtObjectPtr ptr);
int QBluetoothServiceDiscoveryAgent_SetRemoteAddress(QtObjectPtr ptr, QtObjectPtr address);
void QBluetoothServiceDiscoveryAgent_SetUuidFilter2(QtObjectPtr ptr, QtObjectPtr uuid);
void QBluetoothServiceDiscoveryAgent_Start(QtObjectPtr ptr, int mode);
void QBluetoothServiceDiscoveryAgent_Stop(QtObjectPtr ptr);
void QBluetoothServiceDiscoveryAgent_DestroyQBluetoothServiceDiscoveryAgent(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif