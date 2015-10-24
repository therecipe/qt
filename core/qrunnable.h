#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QRunnable_AutoDelete(QtObjectPtr ptr);
void QRunnable_Run(QtObjectPtr ptr);
void QRunnable_SetAutoDelete(QtObjectPtr ptr, int autoDelete);
void QRunnable_DestroyQRunnable(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif