#ifdef __cplusplus
extern "C" {
#endif

void* QWriteLocker_NewQWriteLocker(void* lock);
void* QWriteLocker_ReadWriteLock(void* ptr);
void QWriteLocker_Relock(void* ptr);
void QWriteLocker_Unlock(void* ptr);
void QWriteLocker_DestroyQWriteLocker(void* ptr);

#ifdef __cplusplus
}
#endif