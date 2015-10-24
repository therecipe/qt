#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QBluetoothDeviceDiscoveryAgent_ConnectCanceled(QtObjectPtr ptr);
void QBluetoothDeviceDiscoveryAgent_DisconnectCanceled(QtObjectPtr ptr);
void QBluetoothDeviceDiscoveryAgent_ConnectFinished(QtObjectPtr ptr);
void QBluetoothDeviceDiscoveryAgent_DisconnectFinished(QtObjectPtr ptr);
QtObjectPtr QBluetoothDeviceDiscoveryAgent_NewQBluetoothDeviceDiscoveryAgent(QtObjectPtr parent);
QtObjectPtr QBluetoothDeviceDiscoveryAgent_NewQBluetoothDeviceDiscoveryAgent2(QtObjectPtr deviceAdapter, QtObjectPtr parent);
int QBluetoothDeviceDiscoveryAgent_Error(QtObjectPtr ptr);
char* QBluetoothDeviceDiscoveryAgent_ErrorString(QtObjectPtr ptr);
int QBluetoothDeviceDiscoveryAgent_InquiryType(QtObjectPtr ptr);
int QBluetoothDeviceDiscoveryAgent_IsActive(QtObjectPtr ptr);
void QBluetoothDeviceDiscoveryAgent_SetInquiryType(QtObjectPtr ptr, int ty);
void QBluetoothDeviceDiscoveryAgent_Start(QtObjectPtr ptr);
void QBluetoothDeviceDiscoveryAgent_Stop(QtObjectPtr ptr);
void QBluetoothDeviceDiscoveryAgent_DestroyQBluetoothDeviceDiscoveryAgent(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif