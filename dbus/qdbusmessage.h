#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDBusMessage_NewQDBusMessage();
QtObjectPtr QDBusMessage_NewQDBusMessage2(QtObjectPtr other);
int QDBusMessage_AutoStartService(QtObjectPtr ptr);
char* QDBusMessage_ErrorMessage(QtObjectPtr ptr);
char* QDBusMessage_ErrorName(QtObjectPtr ptr);
char* QDBusMessage_Interface(QtObjectPtr ptr);
int QDBusMessage_IsDelayedReply(QtObjectPtr ptr);
int QDBusMessage_IsReplyRequired(QtObjectPtr ptr);
char* QDBusMessage_Member(QtObjectPtr ptr);
char* QDBusMessage_Path(QtObjectPtr ptr);
char* QDBusMessage_Service(QtObjectPtr ptr);
void QDBusMessage_SetAutoStartService(QtObjectPtr ptr, int enable);
void QDBusMessage_SetDelayedReply(QtObjectPtr ptr, int enable);
char* QDBusMessage_Signature(QtObjectPtr ptr);
int QDBusMessage_Type(QtObjectPtr ptr);
void QDBusMessage_DestroyQDBusMessage(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif