#ifdef __cplusplus
extern "C" {
#endif

int QVideoDeviceSelectorControl_DefaultDevice(void* ptr);
int QVideoDeviceSelectorControl_DeviceCount(void* ptr);
char* QVideoDeviceSelectorControl_DeviceDescription(void* ptr, int index);
char* QVideoDeviceSelectorControl_DeviceName(void* ptr, int index);
void QVideoDeviceSelectorControl_ConnectDevicesChanged(void* ptr);
void QVideoDeviceSelectorControl_DisconnectDevicesChanged(void* ptr);
int QVideoDeviceSelectorControl_SelectedDevice(void* ptr);
void QVideoDeviceSelectorControl_ConnectSelectedDeviceChanged(void* ptr);
void QVideoDeviceSelectorControl_DisconnectSelectedDeviceChanged(void* ptr);
void QVideoDeviceSelectorControl_SetSelectedDevice(void* ptr, int index);
void QVideoDeviceSelectorControl_DestroyQVideoDeviceSelectorControl(void* ptr);

#ifdef __cplusplus
}
#endif