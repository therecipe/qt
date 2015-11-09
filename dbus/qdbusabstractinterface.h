#ifdef __cplusplus
extern "C" {
#endif

char* QDBusAbstractInterface_Interface(void* ptr);
int QDBusAbstractInterface_IsValid(void* ptr);
char* QDBusAbstractInterface_Path(void* ptr);
char* QDBusAbstractInterface_Service(void* ptr);
void QDBusAbstractInterface_SetTimeout(void* ptr, int timeout);
int QDBusAbstractInterface_Timeout(void* ptr);
void QDBusAbstractInterface_DestroyQDBusAbstractInterface(void* ptr);

#ifdef __cplusplus
}
#endif