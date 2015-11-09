#ifdef __cplusplus
extern "C" {
#endif

double QCameraExposure_Aperture(void* ptr);
double QCameraExposure_ExposureCompensation(void* ptr);
int QCameraExposure_ExposureMode(void* ptr);
int QCameraExposure_FlashMode(void* ptr);
int QCameraExposure_IsoSensitivity(void* ptr);
int QCameraExposure_MeteringMode(void* ptr);
void QCameraExposure_SetAutoAperture(void* ptr);
void QCameraExposure_SetAutoIsoSensitivity(void* ptr);
void QCameraExposure_SetExposureCompensation(void* ptr, double ev);
void QCameraExposure_SetExposureMode(void* ptr, int mode);
void QCameraExposure_SetFlashMode(void* ptr, int mode);
void QCameraExposure_SetManualAperture(void* ptr, double aperture);
void QCameraExposure_SetManualIsoSensitivity(void* ptr, int iso);
void QCameraExposure_SetMeteringMode(void* ptr, int mode);
void QCameraExposure_SetSpotMeteringPoint(void* ptr, void* point);
void QCameraExposure_ConnectApertureRangeChanged(void* ptr);
void QCameraExposure_DisconnectApertureRangeChanged(void* ptr);
void QCameraExposure_ConnectFlashReady(void* ptr);
void QCameraExposure_DisconnectFlashReady(void* ptr);
int QCameraExposure_IsAvailable(void* ptr);
int QCameraExposure_IsExposureModeSupported(void* ptr, int mode);
int QCameraExposure_IsFlashModeSupported(void* ptr, int mode);
int QCameraExposure_IsFlashReady(void* ptr);
int QCameraExposure_IsMeteringModeSupported(void* ptr, int mode);
void QCameraExposure_ConnectIsoSensitivityChanged(void* ptr);
void QCameraExposure_DisconnectIsoSensitivityChanged(void* ptr);
double QCameraExposure_RequestedAperture(void* ptr);
int QCameraExposure_RequestedIsoSensitivity(void* ptr);
double QCameraExposure_RequestedShutterSpeed(void* ptr);
void QCameraExposure_SetAutoShutterSpeed(void* ptr);
void QCameraExposure_SetManualShutterSpeed(void* ptr, double seconds);
double QCameraExposure_ShutterSpeed(void* ptr);
void QCameraExposure_ConnectShutterSpeedRangeChanged(void* ptr);
void QCameraExposure_DisconnectShutterSpeedRangeChanged(void* ptr);

#ifdef __cplusplus
}
#endif