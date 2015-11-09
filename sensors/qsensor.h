#ifdef __cplusplus
extern "C" {
#endif

int QSensor_AxesOrientationMode(void* ptr);
int QSensor_BufferSize(void* ptr);
int QSensor_CurrentOrientation(void* ptr);
int QSensor_DataRate(void* ptr);
char* QSensor_Description(void* ptr);
int QSensor_EfficientBufferSize(void* ptr);
int QSensor_Error(void* ptr);
void* QSensor_Identifier(void* ptr);
int QSensor_IsActive(void* ptr);
int QSensor_IsAlwaysOn(void* ptr);
int QSensor_IsBusy(void* ptr);
int QSensor_IsConnectedToBackend(void* ptr);
int QSensor_MaxBufferSize(void* ptr);
int QSensor_OutputRange(void* ptr);
void* QSensor_Reading(void* ptr);
void QSensor_SetActive(void* ptr, int active);
void QSensor_SetAlwaysOn(void* ptr, int alwaysOn);
void QSensor_SetAxesOrientationMode(void* ptr, int axesOrientationMode);
void QSensor_SetBufferSize(void* ptr, int bufferSize);
void QSensor_SetDataRate(void* ptr, int rate);
void QSensor_SetIdentifier(void* ptr, void* identifier);
void QSensor_SetOutputRange(void* ptr, int index);
void QSensor_SetUserOrientation(void* ptr, int userOrientation);
int QSensor_SkipDuplicates(void* ptr);
void* QSensor_Type(void* ptr);
int QSensor_UserOrientation(void* ptr);
void* QSensor_NewQSensor(void* ty, void* parent);
void QSensor_ConnectActiveChanged(void* ptr);
void QSensor_DisconnectActiveChanged(void* ptr);
void QSensor_AddFilter(void* ptr, void* filter);
void QSensor_ConnectAlwaysOnChanged(void* ptr);
void QSensor_DisconnectAlwaysOnChanged(void* ptr);
void QSensor_ConnectAvailableSensorsChanged(void* ptr);
void QSensor_DisconnectAvailableSensorsChanged(void* ptr);
void QSensor_ConnectAxesOrientationModeChanged(void* ptr);
void QSensor_DisconnectAxesOrientationModeChanged(void* ptr);
void QSensor_ConnectBufferSizeChanged(void* ptr);
void QSensor_DisconnectBufferSizeChanged(void* ptr);
void QSensor_ConnectBusyChanged(void* ptr);
void QSensor_DisconnectBusyChanged(void* ptr);
int QSensor_ConnectToBackend(void* ptr);
void QSensor_ConnectCurrentOrientationChanged(void* ptr);
void QSensor_DisconnectCurrentOrientationChanged(void* ptr);
void QSensor_ConnectDataRateChanged(void* ptr);
void QSensor_DisconnectDataRateChanged(void* ptr);
void* QSensor_QSensor_DefaultSensorForType(void* ty);
void QSensor_ConnectEfficientBufferSizeChanged(void* ptr);
void QSensor_DisconnectEfficientBufferSizeChanged(void* ptr);
int QSensor_IsFeatureSupported(void* ptr, int feature);
void QSensor_ConnectMaxBufferSizeChanged(void* ptr);
void QSensor_DisconnectMaxBufferSizeChanged(void* ptr);
void QSensor_ConnectReadingChanged(void* ptr);
void QSensor_DisconnectReadingChanged(void* ptr);
void QSensor_RemoveFilter(void* ptr, void* filter);
void QSensor_ConnectSensorError(void* ptr);
void QSensor_DisconnectSensorError(void* ptr);
void QSensor_SetCurrentOrientation(void* ptr, int currentOrientation);
void QSensor_SetEfficientBufferSize(void* ptr, int efficientBufferSize);
void QSensor_SetMaxBufferSize(void* ptr, int maxBufferSize);
void QSensor_SetSkipDuplicates(void* ptr, int skipDuplicates);
void QSensor_ConnectSkipDuplicatesChanged(void* ptr);
void QSensor_DisconnectSkipDuplicatesChanged(void* ptr);
int QSensor_Start(void* ptr);
void QSensor_Stop(void* ptr);
void QSensor_ConnectUserOrientationChanged(void* ptr);
void QSensor_DisconnectUserOrientationChanged(void* ptr);
void QSensor_DestroyQSensor(void* ptr);

#ifdef __cplusplus
}
#endif