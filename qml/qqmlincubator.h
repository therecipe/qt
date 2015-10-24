#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QQmlIncubator_NewQQmlIncubator(int mode);
void QQmlIncubator_Clear(QtObjectPtr ptr);
void QQmlIncubator_ForceCompletion(QtObjectPtr ptr);
int QQmlIncubator_IncubationMode(QtObjectPtr ptr);
int QQmlIncubator_IsError(QtObjectPtr ptr);
int QQmlIncubator_IsLoading(QtObjectPtr ptr);
int QQmlIncubator_IsNull(QtObjectPtr ptr);
int QQmlIncubator_IsReady(QtObjectPtr ptr);
QtObjectPtr QQmlIncubator_Object(QtObjectPtr ptr);
int QQmlIncubator_Status(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif