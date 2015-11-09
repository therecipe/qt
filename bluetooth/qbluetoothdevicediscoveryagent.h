#ifdef __cplusplus
extern "C" {
#endif

void QBluetoothDeviceDiscoveryAgent_ConnectCanceled(void* ptr);
void QBluetoothDeviceDiscoveryAgent_DisconnectCanceled(void* ptr);
void QBluetoothDeviceDiscoveryAgent_ConnectFinished(void* ptr);
void QBluetoothDeviceDiscoveryAgent_DisconnectFinished(void* ptr);
void* QBluetoothDeviceDiscoveryAgent_NewQBluetoothDeviceDiscoveryAgent(void* parent);
void* QBluetoothDeviceDiscoveryAgent_NewQBluetoothDeviceDiscoveryAgent2(void* deviceAdapter, void* parent);
int QBluetoothDeviceDiscoveryAgent_Error(void* ptr);
char* QBluetoothDeviceDiscoveryAgent_ErrorString(void* ptr);
int QBluetoothDeviceDiscoveryAgent_InquiryType(void* ptr);
int QBluetoothDeviceDiscoveryAgent_IsActive(void* ptr);
void QBluetoothDeviceDiscoveryAgent_SetInquiryType(void* ptr, int ty);
void QBluetoothDeviceDiscoveryAgent_Start(void* ptr);
void QBluetoothDeviceDiscoveryAgent_Stop(void* ptr);
void QBluetoothDeviceDiscoveryAgent_DestroyQBluetoothDeviceDiscoveryAgent(void* ptr);

#ifdef __cplusplus
}
#endif