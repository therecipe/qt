#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QCamera_CaptureMode(QtObjectPtr ptr);
void QCamera_SearchAndLock2(QtObjectPtr ptr, int locks);
void QCamera_SetCaptureMode(QtObjectPtr ptr, int mode);
int QCamera_Status(QtObjectPtr ptr);
QtObjectPtr QCamera_NewQCamera4(int position, QtObjectPtr parent);
QtObjectPtr QCamera_NewQCamera(QtObjectPtr parent);
QtObjectPtr QCamera_NewQCamera2(QtObjectPtr deviceName, QtObjectPtr parent);
QtObjectPtr QCamera_NewQCamera3(QtObjectPtr cameraInfo, QtObjectPtr parent);
void QCamera_ConnectCaptureModeChanged(QtObjectPtr ptr);
void QCamera_DisconnectCaptureModeChanged(QtObjectPtr ptr);
int QCamera_Error(QtObjectPtr ptr);
char* QCamera_ErrorString(QtObjectPtr ptr);
QtObjectPtr QCamera_Exposure(QtObjectPtr ptr);
QtObjectPtr QCamera_Focus(QtObjectPtr ptr);
QtObjectPtr QCamera_ImageProcessing(QtObjectPtr ptr);
int QCamera_IsCaptureModeSupported(QtObjectPtr ptr, int mode);
void QCamera_Load(QtObjectPtr ptr);
void QCamera_ConnectLockFailed(QtObjectPtr ptr);
void QCamera_DisconnectLockFailed(QtObjectPtr ptr);
int QCamera_LockStatus(QtObjectPtr ptr);
int QCamera_LockStatus2(QtObjectPtr ptr, int lockType);
void QCamera_ConnectLockStatusChanged(QtObjectPtr ptr);
void QCamera_DisconnectLockStatusChanged(QtObjectPtr ptr);
void QCamera_ConnectLocked(QtObjectPtr ptr);
void QCamera_DisconnectLocked(QtObjectPtr ptr);
int QCamera_RequestedLocks(QtObjectPtr ptr);
void QCamera_SearchAndLock(QtObjectPtr ptr);
void QCamera_SetViewfinder3(QtObjectPtr ptr, QtObjectPtr surface);
void QCamera_SetViewfinderSettings(QtObjectPtr ptr, QtObjectPtr settings);
void QCamera_Start(QtObjectPtr ptr);
void QCamera_ConnectStateChanged(QtObjectPtr ptr);
void QCamera_DisconnectStateChanged(QtObjectPtr ptr);
void QCamera_ConnectStatusChanged(QtObjectPtr ptr);
void QCamera_DisconnectStatusChanged(QtObjectPtr ptr);
void QCamera_Stop(QtObjectPtr ptr);
int QCamera_SupportedLocks(QtObjectPtr ptr);
void QCamera_Unload(QtObjectPtr ptr);
void QCamera_Unlock(QtObjectPtr ptr);
void QCamera_Unlock2(QtObjectPtr ptr, int locks);
void QCamera_DestroyQCamera(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif