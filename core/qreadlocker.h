#ifdef __cplusplus
extern "C" {
#endif

void* QReadLocker_NewQReadLocker(void* lock);
void* QReadLocker_ReadWriteLock(void* ptr);
void QReadLocker_Relock(void* ptr);
void QReadLocker_Unlock(void* ptr);
void QReadLocker_DestroyQReadLocker(void* ptr);

#ifdef __cplusplus
}
#endif