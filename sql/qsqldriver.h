#ifdef __cplusplus
extern "C" {
#endif

int QSqlDriver_BeginTransaction(void* ptr);
void QSqlDriver_Close(void* ptr);
int QSqlDriver_CommitTransaction(void* ptr);
void* QSqlDriver_CreateResult(void* ptr);
int QSqlDriver_DbmsType(void* ptr);
char* QSqlDriver_EscapeIdentifier(void* ptr, char* identifier, int ty);
char* QSqlDriver_FormatValue(void* ptr, void* field, int trimStrings);
void* QSqlDriver_Handle(void* ptr);
int QSqlDriver_HasFeature(void* ptr, int feature);
int QSqlDriver_IsIdentifierEscaped(void* ptr, char* identifier, int ty);
int QSqlDriver_IsOpen(void* ptr);
int QSqlDriver_IsOpenError(void* ptr);
void QSqlDriver_ConnectNotification(void* ptr);
void QSqlDriver_DisconnectNotification(void* ptr);
int QSqlDriver_Open(void* ptr, char* db, char* user, char* password, char* host, int port, char* options);
int QSqlDriver_RollbackTransaction(void* ptr);
char* QSqlDriver_SqlStatement(void* ptr, int ty, char* tableName, void* rec, int preparedStatement);
char* QSqlDriver_StripDelimiters(void* ptr, char* identifier, int ty);
int QSqlDriver_SubscribeToNotification(void* ptr, char* name);
char* QSqlDriver_SubscribedToNotifications(void* ptr);
int QSqlDriver_UnsubscribeFromNotification(void* ptr, char* name);
void QSqlDriver_DestroyQSqlDriver(void* ptr);

#ifdef __cplusplus
}
#endif