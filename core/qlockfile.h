#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QLockFile_NewQLockFile(char* fileName);
int QLockFile_Error(QtObjectPtr ptr);
int QLockFile_IsLocked(QtObjectPtr ptr);
int QLockFile_Lock(QtObjectPtr ptr);
int QLockFile_RemoveStaleLockFile(QtObjectPtr ptr);
void QLockFile_SetStaleLockTime(QtObjectPtr ptr, int staleLockTime);
int QLockFile_StaleLockTime(QtObjectPtr ptr);
int QLockFile_TryLock(QtObjectPtr ptr, int timeout);
void QLockFile_DestroyQLockFile(QtObjectPtr ptr);
void QLockFile_Unlock(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif