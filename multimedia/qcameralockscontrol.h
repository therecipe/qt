#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QCameraLocksControl_LockStatus(QtObjectPtr ptr, int lock);
void QCameraLocksControl_ConnectLockStatusChanged(QtObjectPtr ptr);
void QCameraLocksControl_DisconnectLockStatusChanged(QtObjectPtr ptr);
void QCameraLocksControl_SearchAndLock(QtObjectPtr ptr, int locks);
int QCameraLocksControl_SupportedLocks(QtObjectPtr ptr);
void QCameraLocksControl_Unlock(QtObjectPtr ptr, int locks);
void QCameraLocksControl_DestroyQCameraLocksControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif