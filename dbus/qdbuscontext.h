#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDBusContext_NewQDBusContext();
int QDBusContext_CalledFromDBus(QtObjectPtr ptr);
int QDBusContext_IsDelayedReply(QtObjectPtr ptr);
void QDBusContext_SendErrorReply2(QtObjectPtr ptr, int ty, char* msg);
void QDBusContext_SendErrorReply(QtObjectPtr ptr, char* name, char* msg);
void QDBusContext_SetDelayedReply(QtObjectPtr ptr, int enable);
void QDBusContext_DestroyQDBusContext(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif