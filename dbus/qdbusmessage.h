#ifdef __cplusplus
extern "C" {
#endif

void* QDBusMessage_NewQDBusMessage();
void* QDBusMessage_NewQDBusMessage2(void* other);
int QDBusMessage_AutoStartService(void* ptr);
char* QDBusMessage_ErrorMessage(void* ptr);
char* QDBusMessage_ErrorName(void* ptr);
char* QDBusMessage_Interface(void* ptr);
int QDBusMessage_IsDelayedReply(void* ptr);
int QDBusMessage_IsReplyRequired(void* ptr);
char* QDBusMessage_Member(void* ptr);
char* QDBusMessage_Path(void* ptr);
char* QDBusMessage_Service(void* ptr);
void QDBusMessage_SetAutoStartService(void* ptr, int enable);
void QDBusMessage_SetDelayedReply(void* ptr, int enable);
char* QDBusMessage_Signature(void* ptr);
int QDBusMessage_Type(void* ptr);
void QDBusMessage_DestroyQDBusMessage(void* ptr);

#ifdef __cplusplus
}
#endif