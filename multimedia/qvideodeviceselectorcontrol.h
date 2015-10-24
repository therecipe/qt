#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QVideoDeviceSelectorControl_DefaultDevice(QtObjectPtr ptr);
int QVideoDeviceSelectorControl_DeviceCount(QtObjectPtr ptr);
char* QVideoDeviceSelectorControl_DeviceDescription(QtObjectPtr ptr, int index);
char* QVideoDeviceSelectorControl_DeviceName(QtObjectPtr ptr, int index);
void QVideoDeviceSelectorControl_ConnectDevicesChanged(QtObjectPtr ptr);
void QVideoDeviceSelectorControl_DisconnectDevicesChanged(QtObjectPtr ptr);
int QVideoDeviceSelectorControl_SelectedDevice(QtObjectPtr ptr);
void QVideoDeviceSelectorControl_ConnectSelectedDeviceChanged(QtObjectPtr ptr);
void QVideoDeviceSelectorControl_DisconnectSelectedDeviceChanged(QtObjectPtr ptr);
void QVideoDeviceSelectorControl_SetSelectedDevice(QtObjectPtr ptr, int index);
void QVideoDeviceSelectorControl_DestroyQVideoDeviceSelectorControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif