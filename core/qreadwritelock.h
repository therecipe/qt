#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QReadWriteLock_NewQReadWriteLock(int recursionMode);
void QReadWriteLock_LockForRead(QtObjectPtr ptr);
void QReadWriteLock_LockForWrite(QtObjectPtr ptr);
int QReadWriteLock_TryLockForRead(QtObjectPtr ptr);
int QReadWriteLock_TryLockForRead2(QtObjectPtr ptr, int timeout);
int QReadWriteLock_TryLockForWrite(QtObjectPtr ptr);
int QReadWriteLock_TryLockForWrite2(QtObjectPtr ptr, int timeout);
void QReadWriteLock_Unlock(QtObjectPtr ptr);
void QReadWriteLock_DestroyQReadWriteLock(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif