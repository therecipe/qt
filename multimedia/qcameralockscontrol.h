#ifdef __cplusplus
extern "C" {
#endif

int QCameraLocksControl_LockStatus(void* ptr, int lock);
void QCameraLocksControl_ConnectLockStatusChanged(void* ptr);
void QCameraLocksControl_DisconnectLockStatusChanged(void* ptr);
void QCameraLocksControl_SearchAndLock(void* ptr, int locks);
int QCameraLocksControl_SupportedLocks(void* ptr);
void QCameraLocksControl_Unlock(void* ptr, int locks);
void QCameraLocksControl_DestroyQCameraLocksControl(void* ptr);

#ifdef __cplusplus
}
#endif