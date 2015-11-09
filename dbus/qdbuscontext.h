#ifdef __cplusplus
extern "C" {
#endif

void* QDBusContext_NewQDBusContext();
int QDBusContext_CalledFromDBus(void* ptr);
int QDBusContext_IsDelayedReply(void* ptr);
void QDBusContext_SendErrorReply2(void* ptr, int ty, char* msg);
void QDBusContext_SendErrorReply(void* ptr, char* name, char* msg);
void QDBusContext_SetDelayedReply(void* ptr, int enable);
void QDBusContext_DestroyQDBusContext(void* ptr);

#ifdef __cplusplus
}
#endif