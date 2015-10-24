#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QCameraFocus_FocusMode(QtObjectPtr ptr);
int QCameraFocus_FocusPointMode(QtObjectPtr ptr);
void QCameraFocus_SetCustomFocusPoint(QtObjectPtr ptr, QtObjectPtr point);
void QCameraFocus_SetFocusMode(QtObjectPtr ptr, int mode);
void QCameraFocus_SetFocusPointMode(QtObjectPtr ptr, int mode);
void QCameraFocus_ConnectFocusZonesChanged(QtObjectPtr ptr);
void QCameraFocus_DisconnectFocusZonesChanged(QtObjectPtr ptr);
int QCameraFocus_IsAvailable(QtObjectPtr ptr);
int QCameraFocus_IsFocusModeSupported(QtObjectPtr ptr, int mode);
int QCameraFocus_IsFocusPointModeSupported(QtObjectPtr ptr, int mode);

#ifdef __cplusplus
}
#endif