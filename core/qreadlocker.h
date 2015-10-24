#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QReadLocker_NewQReadLocker(QtObjectPtr lock);
QtObjectPtr QReadLocker_ReadWriteLock(QtObjectPtr ptr);
void QReadLocker_Relock(QtObjectPtr ptr);
void QReadLocker_Unlock(QtObjectPtr ptr);
void QReadLocker_DestroyQReadLocker(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif