#ifdef __cplusplus
extern "C" {
#endif

int QCameraControl_CanChangeProperty(void* ptr, int changeType, int status);
int QCameraControl_CaptureMode(void* ptr);
void QCameraControl_ConnectCaptureModeChanged(void* ptr);
void QCameraControl_DisconnectCaptureModeChanged(void* ptr);
void QCameraControl_ConnectError(void* ptr);
void QCameraControl_DisconnectError(void* ptr);
int QCameraControl_IsCaptureModeSupported(void* ptr, int mode);
void QCameraControl_SetCaptureMode(void* ptr, int mode);
void QCameraControl_SetState(void* ptr, int state);
void QCameraControl_ConnectStateChanged(void* ptr);
void QCameraControl_DisconnectStateChanged(void* ptr);
int QCameraControl_Status(void* ptr);
void QCameraControl_ConnectStatusChanged(void* ptr);
void QCameraControl_DisconnectStatusChanged(void* ptr);
void QCameraControl_DestroyQCameraControl(void* ptr);

#ifdef __cplusplus
}
#endif