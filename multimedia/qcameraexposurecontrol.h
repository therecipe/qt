#ifdef __cplusplus
extern "C" {
#endif

void* QCameraExposureControl_ActualValue(void* ptr, int parameter);
void QCameraExposureControl_ConnectActualValueChanged(void* ptr);
void QCameraExposureControl_DisconnectActualValueChanged(void* ptr);
int QCameraExposureControl_IsParameterSupported(void* ptr, int parameter);
void QCameraExposureControl_ConnectParameterRangeChanged(void* ptr);
void QCameraExposureControl_DisconnectParameterRangeChanged(void* ptr);
void* QCameraExposureControl_RequestedValue(void* ptr, int parameter);
void QCameraExposureControl_ConnectRequestedValueChanged(void* ptr);
void QCameraExposureControl_DisconnectRequestedValueChanged(void* ptr);
int QCameraExposureControl_SetValue(void* ptr, int parameter, void* value);
void QCameraExposureControl_DestroyQCameraExposureControl(void* ptr);

#ifdef __cplusplus
}
#endif