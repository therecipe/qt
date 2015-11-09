#ifdef __cplusplus
extern "C" {
#endif

void* QBluetoothDeviceInfo_NewQBluetoothDeviceInfo();
void* QBluetoothDeviceInfo_NewQBluetoothDeviceInfo4(void* other);
int QBluetoothDeviceInfo_CoreConfigurations(void* ptr);
int QBluetoothDeviceInfo_IsCached(void* ptr);
int QBluetoothDeviceInfo_IsValid(void* ptr);
int QBluetoothDeviceInfo_MajorDeviceClass(void* ptr);
char* QBluetoothDeviceInfo_Name(void* ptr);
int QBluetoothDeviceInfo_ServiceClasses(void* ptr);
int QBluetoothDeviceInfo_ServiceUuidsCompleteness(void* ptr);
void QBluetoothDeviceInfo_SetCached(void* ptr, int cached);
void QBluetoothDeviceInfo_SetCoreConfigurations(void* ptr, int coreConfigs);
void QBluetoothDeviceInfo_SetDeviceUuid(void* ptr, void* uuid);
void QBluetoothDeviceInfo_DestroyQBluetoothDeviceInfo(void* ptr);

#ifdef __cplusplus
}
#endif