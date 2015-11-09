#ifdef __cplusplus
extern "C" {
#endif

void* QReadWriteLock_NewQReadWriteLock(int recursionMode);
void QReadWriteLock_LockForRead(void* ptr);
void QReadWriteLock_LockForWrite(void* ptr);
int QReadWriteLock_TryLockForRead(void* ptr);
int QReadWriteLock_TryLockForRead2(void* ptr, int timeout);
int QReadWriteLock_TryLockForWrite(void* ptr);
int QReadWriteLock_TryLockForWrite2(void* ptr, int timeout);
void QReadWriteLock_Unlock(void* ptr);
void QReadWriteLock_DestroyQReadWriteLock(void* ptr);

#ifdef __cplusplus
}
#endif