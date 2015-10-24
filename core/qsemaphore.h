#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSemaphore_NewQSemaphore(int n);
void QSemaphore_Acquire(QtObjectPtr ptr, int n);
int QSemaphore_Available(QtObjectPtr ptr);
void QSemaphore_Release(QtObjectPtr ptr, int n);
int QSemaphore_TryAcquire(QtObjectPtr ptr, int n);
int QSemaphore_TryAcquire2(QtObjectPtr ptr, int n, int timeout);
void QSemaphore_DestroyQSemaphore(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif