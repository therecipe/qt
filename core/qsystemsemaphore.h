#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSystemSemaphore_NewQSystemSemaphore(char* key, int initialValue, int mode);
int QSystemSemaphore_Acquire(QtObjectPtr ptr);
int QSystemSemaphore_Error(QtObjectPtr ptr);
char* QSystemSemaphore_ErrorString(QtObjectPtr ptr);
char* QSystemSemaphore_Key(QtObjectPtr ptr);
int QSystemSemaphore_Release(QtObjectPtr ptr, int n);
void QSystemSemaphore_SetKey(QtObjectPtr ptr, char* key, int initialValue, int mode);
void QSystemSemaphore_DestroyQSystemSemaphore(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif