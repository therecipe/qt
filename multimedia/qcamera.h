#ifdef __cplusplus
extern "C" {
#endif

int QCamera_CaptureMode(void* ptr);
void QCamera_SearchAndLock2(void* ptr, int locks);
void QCamera_SetCaptureMode(void* ptr, int mode);
int QCamera_Status(void* ptr);
void* QCamera_NewQCamera4(int position, void* parent);
void* QCamera_NewQCamera(void* parent);
void* QCamera_NewQCamera2(void* deviceName, void* parent);
void* QCamera_NewQCamera3(void* cameraInfo, void* parent);
void QCamera_ConnectCaptureModeChanged(void* ptr);
void QCamera_DisconnectCaptureModeChanged(void* ptr);
int QCamera_Error(void* ptr);
char* QCamera_ErrorString(void* ptr);
void* QCamera_Exposure(void* ptr);
void* QCamera_Focus(void* ptr);
void* QCamera_ImageProcessing(void* ptr);
int QCamera_IsCaptureModeSupported(void* ptr, int mode);
void QCamera_Load(void* ptr);
void QCamera_ConnectLockFailed(void* ptr);
void QCamera_DisconnectLockFailed(void* ptr);
int QCamera_LockStatus(void* ptr);
int QCamera_LockStatus2(void* ptr, int lockType);
void QCamera_ConnectLockStatusChanged(void* ptr);
void QCamera_DisconnectLockStatusChanged(void* ptr);
void QCamera_ConnectLocked(void* ptr);
void QCamera_DisconnectLocked(void* ptr);
int QCamera_RequestedLocks(void* ptr);
void QCamera_SearchAndLock(void* ptr);
void QCamera_SetViewfinder3(void* ptr, void* surface);
void QCamera_SetViewfinderSettings(void* ptr, void* settings);
void QCamera_Start(void* ptr);
void QCamera_ConnectStateChanged(void* ptr);
void QCamera_DisconnectStateChanged(void* ptr);
void QCamera_ConnectStatusChanged(void* ptr);
void QCamera_DisconnectStatusChanged(void* ptr);
void QCamera_Stop(void* ptr);
int QCamera_SupportedLocks(void* ptr);
void QCamera_Unload(void* ptr);
void QCamera_Unlock(void* ptr);
void QCamera_Unlock2(void* ptr, int locks);
void QCamera_DestroyQCamera(void* ptr);

#ifdef __cplusplus
}
#endif