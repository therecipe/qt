#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QCameraFocusControl_FocusMode(QtObjectPtr ptr);
void QCameraFocusControl_ConnectFocusModeChanged(QtObjectPtr ptr);
void QCameraFocusControl_DisconnectFocusModeChanged(QtObjectPtr ptr);
int QCameraFocusControl_FocusPointMode(QtObjectPtr ptr);
void QCameraFocusControl_ConnectFocusPointModeChanged(QtObjectPtr ptr);
void QCameraFocusControl_DisconnectFocusPointModeChanged(QtObjectPtr ptr);
void QCameraFocusControl_ConnectFocusZonesChanged(QtObjectPtr ptr);
void QCameraFocusControl_DisconnectFocusZonesChanged(QtObjectPtr ptr);
int QCameraFocusControl_IsFocusModeSupported(QtObjectPtr ptr, int mode);
int QCameraFocusControl_IsFocusPointModeSupported(QtObjectPtr ptr, int mode);
void QCameraFocusControl_SetCustomFocusPoint(QtObjectPtr ptr, QtObjectPtr point);
void QCameraFocusControl_SetFocusMode(QtObjectPtr ptr, int mode);
void QCameraFocusControl_SetFocusPointMode(QtObjectPtr ptr, int mode);
void QCameraFocusControl_DestroyQCameraFocusControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif