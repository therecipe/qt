#ifdef __cplusplus
extern "C" {
#endif

void* QLockFile_NewQLockFile(char* fileName);
int QLockFile_Error(void* ptr);
int QLockFile_IsLocked(void* ptr);
int QLockFile_Lock(void* ptr);
int QLockFile_RemoveStaleLockFile(void* ptr);
void QLockFile_SetStaleLockTime(void* ptr, int staleLockTime);
int QLockFile_StaleLockTime(void* ptr);
int QLockFile_TryLock(void* ptr, int timeout);
void QLockFile_DestroyQLockFile(void* ptr);
void QLockFile_Unlock(void* ptr);

#ifdef __cplusplus
}
#endif