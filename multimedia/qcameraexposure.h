#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QCameraExposure_ExposureMode(QtObjectPtr ptr);
int QCameraExposure_FlashMode(QtObjectPtr ptr);
int QCameraExposure_IsoSensitivity(QtObjectPtr ptr);
int QCameraExposure_MeteringMode(QtObjectPtr ptr);
void QCameraExposure_SetAutoAperture(QtObjectPtr ptr);
void QCameraExposure_SetAutoIsoSensitivity(QtObjectPtr ptr);
void QCameraExposure_SetExposureMode(QtObjectPtr ptr, int mode);
void QCameraExposure_SetFlashMode(QtObjectPtr ptr, int mode);
void QCameraExposure_SetManualIsoSensitivity(QtObjectPtr ptr, int iso);
void QCameraExposure_SetMeteringMode(QtObjectPtr ptr, int mode);
void QCameraExposure_SetSpotMeteringPoint(QtObjectPtr ptr, QtObjectPtr point);
void QCameraExposure_ConnectApertureRangeChanged(QtObjectPtr ptr);
void QCameraExposure_DisconnectApertureRangeChanged(QtObjectPtr ptr);
void QCameraExposure_ConnectFlashReady(QtObjectPtr ptr);
void QCameraExposure_DisconnectFlashReady(QtObjectPtr ptr);
int QCameraExposure_IsAvailable(QtObjectPtr ptr);
int QCameraExposure_IsExposureModeSupported(QtObjectPtr ptr, int mode);
int QCameraExposure_IsFlashModeSupported(QtObjectPtr ptr, int mode);
int QCameraExposure_IsFlashReady(QtObjectPtr ptr);
int QCameraExposure_IsMeteringModeSupported(QtObjectPtr ptr, int mode);
void QCameraExposure_ConnectIsoSensitivityChanged(QtObjectPtr ptr);
void QCameraExposure_DisconnectIsoSensitivityChanged(QtObjectPtr ptr);
int QCameraExposure_RequestedIsoSensitivity(QtObjectPtr ptr);
void QCameraExposure_SetAutoShutterSpeed(QtObjectPtr ptr);
void QCameraExposure_ConnectShutterSpeedRangeChanged(QtObjectPtr ptr);
void QCameraExposure_DisconnectShutterSpeedRangeChanged(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif