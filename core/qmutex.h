#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QMutex_Lock(QtObjectPtr ptr);
int QMutex_TryLock(QtObjectPtr ptr, int timeout);
void QMutex_Unlock(QtObjectPtr ptr);
QtObjectPtr QMutex_NewQMutex(int mode);
int QMutex_IsRecursive(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif