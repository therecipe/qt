#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QCameraExposureControl_ActualValue(QtObjectPtr ptr, int parameter);
void QCameraExposureControl_ConnectActualValueChanged(QtObjectPtr ptr);
void QCameraExposureControl_DisconnectActualValueChanged(QtObjectPtr ptr);
int QCameraExposureControl_IsParameterSupported(QtObjectPtr ptr, int parameter);
void QCameraExposureControl_ConnectParameterRangeChanged(QtObjectPtr ptr);
void QCameraExposureControl_DisconnectParameterRangeChanged(QtObjectPtr ptr);
char* QCameraExposureControl_RequestedValue(QtObjectPtr ptr, int parameter);
void QCameraExposureControl_ConnectRequestedValueChanged(QtObjectPtr ptr);
void QCameraExposureControl_DisconnectRequestedValueChanged(QtObjectPtr ptr);
int QCameraExposureControl_SetValue(QtObjectPtr ptr, int parameter, char* value);
void QCameraExposureControl_DestroyQCameraExposureControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif