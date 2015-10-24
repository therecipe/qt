#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QWriteLocker_NewQWriteLocker(QtObjectPtr lock);
QtObjectPtr QWriteLocker_ReadWriteLock(QtObjectPtr ptr);
void QWriteLocker_Relock(QtObjectPtr ptr);
void QWriteLocker_Unlock(QtObjectPtr ptr);
void QWriteLocker_DestroyQWriteLocker(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif