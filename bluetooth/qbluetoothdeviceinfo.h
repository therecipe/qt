#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QBluetoothDeviceInfo_NewQBluetoothDeviceInfo();
QtObjectPtr QBluetoothDeviceInfo_NewQBluetoothDeviceInfo4(QtObjectPtr other);
int QBluetoothDeviceInfo_CoreConfigurations(QtObjectPtr ptr);
int QBluetoothDeviceInfo_IsCached(QtObjectPtr ptr);
int QBluetoothDeviceInfo_IsValid(QtObjectPtr ptr);
int QBluetoothDeviceInfo_MajorDeviceClass(QtObjectPtr ptr);
char* QBluetoothDeviceInfo_Name(QtObjectPtr ptr);
int QBluetoothDeviceInfo_ServiceClasses(QtObjectPtr ptr);
int QBluetoothDeviceInfo_ServiceUuidsCompleteness(QtObjectPtr ptr);
void QBluetoothDeviceInfo_SetCached(QtObjectPtr ptr, int cached);
void QBluetoothDeviceInfo_SetCoreConfigurations(QtObjectPtr ptr, int coreConfigs);
void QBluetoothDeviceInfo_SetDeviceUuid(QtObjectPtr ptr, QtObjectPtr uuid);
void QBluetoothDeviceInfo_DestroyQBluetoothDeviceInfo(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif