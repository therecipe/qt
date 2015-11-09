#ifdef __cplusplus
extern "C" {
#endif

void* QQmlIncubator_NewQQmlIncubator(int mode);
void QQmlIncubator_Clear(void* ptr);
void QQmlIncubator_ForceCompletion(void* ptr);
int QQmlIncubator_IncubationMode(void* ptr);
int QQmlIncubator_IsError(void* ptr);
int QQmlIncubator_IsLoading(void* ptr);
int QQmlIncubator_IsNull(void* ptr);
int QQmlIncubator_IsReady(void* ptr);
void* QQmlIncubator_Object(void* ptr);
int QQmlIncubator_Status(void* ptr);

#ifdef __cplusplus
}
#endif