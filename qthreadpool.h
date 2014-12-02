#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QThreadPool_New_QObject(QtObjectPtr parent);
void QThreadPool_Destroy(QtObjectPtr ptr);
int QThreadPool_ActiveThreadCount(QtObjectPtr ptr);
void QThreadPool_Clear(QtObjectPtr ptr);
int QThreadPool_ExpiryTimeout(QtObjectPtr ptr);
int QThreadPool_MaxThreadCount(QtObjectPtr ptr);
void QThreadPool_ReleaseThread(QtObjectPtr ptr);
void QThreadPool_ReserveThread(QtObjectPtr ptr);
void QThreadPool_SetExpiryTimeout_Int(QtObjectPtr ptr, int expiryTimeout);
void QThreadPool_SetMaxThreadCount_Int(QtObjectPtr ptr, int maxThreadCount);
int QThreadPool_WaitForDone_Int(QtObjectPtr ptr, int msecs);
//Static Public Members
QtObjectPtr QThreadPool_GlobalInstance();

#ifdef __cplusplus
}
#endif
