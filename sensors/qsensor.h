#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QSensor_AxesOrientationMode(QtObjectPtr ptr);
int QSensor_BufferSize(QtObjectPtr ptr);
int QSensor_CurrentOrientation(QtObjectPtr ptr);
int QSensor_DataRate(QtObjectPtr ptr);
char* QSensor_Description(QtObjectPtr ptr);
int QSensor_EfficientBufferSize(QtObjectPtr ptr);
int QSensor_Error(QtObjectPtr ptr);
int QSensor_IsActive(QtObjectPtr ptr);
int QSensor_IsAlwaysOn(QtObjectPtr ptr);
int QSensor_IsBusy(QtObjectPtr ptr);
int QSensor_IsConnectedToBackend(QtObjectPtr ptr);
int QSensor_MaxBufferSize(QtObjectPtr ptr);
int QSensor_OutputRange(QtObjectPtr ptr);
QtObjectPtr QSensor_Reading(QtObjectPtr ptr);
void QSensor_SetActive(QtObjectPtr ptr, int active);
void QSensor_SetAlwaysOn(QtObjectPtr ptr, int alwaysOn);
void QSensor_SetAxesOrientationMode(QtObjectPtr ptr, int axesOrientationMode);
void QSensor_SetBufferSize(QtObjectPtr ptr, int bufferSize);
void QSensor_SetDataRate(QtObjectPtr ptr, int rate);
void QSensor_SetIdentifier(QtObjectPtr ptr, QtObjectPtr identifier);
void QSensor_SetOutputRange(QtObjectPtr ptr, int index);
void QSensor_SetUserOrientation(QtObjectPtr ptr, int userOrientation);
int QSensor_SkipDuplicates(QtObjectPtr ptr);
int QSensor_UserOrientation(QtObjectPtr ptr);
QtObjectPtr QSensor_NewQSensor(QtObjectPtr ty, QtObjectPtr parent);
void QSensor_ConnectActiveChanged(QtObjectPtr ptr);
void QSensor_DisconnectActiveChanged(QtObjectPtr ptr);
void QSensor_AddFilter(QtObjectPtr ptr, QtObjectPtr filter);
void QSensor_ConnectAlwaysOnChanged(QtObjectPtr ptr);
void QSensor_DisconnectAlwaysOnChanged(QtObjectPtr ptr);
void QSensor_ConnectAvailableSensorsChanged(QtObjectPtr ptr);
void QSensor_DisconnectAvailableSensorsChanged(QtObjectPtr ptr);
void QSensor_ConnectAxesOrientationModeChanged(QtObjectPtr ptr);
void QSensor_DisconnectAxesOrientationModeChanged(QtObjectPtr ptr);
void QSensor_ConnectBufferSizeChanged(QtObjectPtr ptr);
void QSensor_DisconnectBufferSizeChanged(QtObjectPtr ptr);
void QSensor_ConnectBusyChanged(QtObjectPtr ptr);
void QSensor_DisconnectBusyChanged(QtObjectPtr ptr);
int QSensor_ConnectToBackend(QtObjectPtr ptr);
void QSensor_ConnectCurrentOrientationChanged(QtObjectPtr ptr);
void QSensor_DisconnectCurrentOrientationChanged(QtObjectPtr ptr);
void QSensor_ConnectDataRateChanged(QtObjectPtr ptr);
void QSensor_DisconnectDataRateChanged(QtObjectPtr ptr);
void QSensor_ConnectEfficientBufferSizeChanged(QtObjectPtr ptr);
void QSensor_DisconnectEfficientBufferSizeChanged(QtObjectPtr ptr);
int QSensor_IsFeatureSupported(QtObjectPtr ptr, int feature);
void QSensor_ConnectMaxBufferSizeChanged(QtObjectPtr ptr);
void QSensor_DisconnectMaxBufferSizeChanged(QtObjectPtr ptr);
void QSensor_ConnectReadingChanged(QtObjectPtr ptr);
void QSensor_DisconnectReadingChanged(QtObjectPtr ptr);
void QSensor_RemoveFilter(QtObjectPtr ptr, QtObjectPtr filter);
void QSensor_ConnectSensorError(QtObjectPtr ptr);
void QSensor_DisconnectSensorError(QtObjectPtr ptr);
void QSensor_SetCurrentOrientation(QtObjectPtr ptr, int currentOrientation);
void QSensor_SetEfficientBufferSize(QtObjectPtr ptr, int efficientBufferSize);
void QSensor_SetMaxBufferSize(QtObjectPtr ptr, int maxBufferSize);
void QSensor_SetSkipDuplicates(QtObjectPtr ptr, int skipDuplicates);
void QSensor_ConnectSkipDuplicatesChanged(QtObjectPtr ptr);
void QSensor_DisconnectSkipDuplicatesChanged(QtObjectPtr ptr);
int QSensor_Start(QtObjectPtr ptr);
void QSensor_Stop(QtObjectPtr ptr);
void QSensor_ConnectUserOrientationChanged(QtObjectPtr ptr);
void QSensor_DisconnectUserOrientationChanged(QtObjectPtr ptr);
void QSensor_DestroyQSensor(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif