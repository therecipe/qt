#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QCameraControl_CanChangeProperty(QtObjectPtr ptr, int changeType, int status);
int QCameraControl_CaptureMode(QtObjectPtr ptr);
void QCameraControl_ConnectCaptureModeChanged(QtObjectPtr ptr);
void QCameraControl_DisconnectCaptureModeChanged(QtObjectPtr ptr);
void QCameraControl_ConnectError(QtObjectPtr ptr);
void QCameraControl_DisconnectError(QtObjectPtr ptr);
int QCameraControl_IsCaptureModeSupported(QtObjectPtr ptr, int mode);
void QCameraControl_SetCaptureMode(QtObjectPtr ptr, int mode);
void QCameraControl_SetState(QtObjectPtr ptr, int state);
void QCameraControl_ConnectStateChanged(QtObjectPtr ptr);
void QCameraControl_DisconnectStateChanged(QtObjectPtr ptr);
int QCameraControl_Status(QtObjectPtr ptr);
void QCameraControl_ConnectStatusChanged(QtObjectPtr ptr);
void QCameraControl_DisconnectStatusChanged(QtObjectPtr ptr);
void QCameraControl_DestroyQCameraControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif