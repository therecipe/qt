#ifdef __cplusplus
extern "C" {
#endif

int QCameraFocusControl_FocusMode(void* ptr);
void QCameraFocusControl_ConnectFocusModeChanged(void* ptr);
void QCameraFocusControl_DisconnectFocusModeChanged(void* ptr);
int QCameraFocusControl_FocusPointMode(void* ptr);
void QCameraFocusControl_ConnectFocusPointModeChanged(void* ptr);
void QCameraFocusControl_DisconnectFocusPointModeChanged(void* ptr);
void QCameraFocusControl_ConnectFocusZonesChanged(void* ptr);
void QCameraFocusControl_DisconnectFocusZonesChanged(void* ptr);
int QCameraFocusControl_IsFocusModeSupported(void* ptr, int mode);
int QCameraFocusControl_IsFocusPointModeSupported(void* ptr, int mode);
void QCameraFocusControl_SetCustomFocusPoint(void* ptr, void* point);
void QCameraFocusControl_SetFocusMode(void* ptr, int mode);
void QCameraFocusControl_SetFocusPointMode(void* ptr, int mode);
void QCameraFocusControl_DestroyQCameraFocusControl(void* ptr);

#ifdef __cplusplus
}
#endif