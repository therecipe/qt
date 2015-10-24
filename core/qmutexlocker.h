#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QMutexLocker_NewQMutexLocker(QtObjectPtr mutex);
QtObjectPtr QMutexLocker_Mutex(QtObjectPtr ptr);
void QMutexLocker_Relock(QtObjectPtr ptr);
void QMutexLocker_Unlock(QtObjectPtr ptr);
void QMutexLocker_DestroyQMutexLocker(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif