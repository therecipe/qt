#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QSqlDriver_BeginTransaction(QtObjectPtr ptr);
void QSqlDriver_Close(QtObjectPtr ptr);
int QSqlDriver_CommitTransaction(QtObjectPtr ptr);
QtObjectPtr QSqlDriver_CreateResult(QtObjectPtr ptr);
int QSqlDriver_DbmsType(QtObjectPtr ptr);
char* QSqlDriver_EscapeIdentifier(QtObjectPtr ptr, char* identifier, int ty);
char* QSqlDriver_FormatValue(QtObjectPtr ptr, QtObjectPtr field, int trimStrings);
char* QSqlDriver_Handle(QtObjectPtr ptr);
int QSqlDriver_HasFeature(QtObjectPtr ptr, int feature);
int QSqlDriver_IsIdentifierEscaped(QtObjectPtr ptr, char* identifier, int ty);
int QSqlDriver_IsOpen(QtObjectPtr ptr);
int QSqlDriver_IsOpenError(QtObjectPtr ptr);
void QSqlDriver_ConnectNotification(QtObjectPtr ptr);
void QSqlDriver_DisconnectNotification(QtObjectPtr ptr);
int QSqlDriver_Open(QtObjectPtr ptr, char* db, char* user, char* password, char* host, int port, char* options);
int QSqlDriver_RollbackTransaction(QtObjectPtr ptr);
char* QSqlDriver_SqlStatement(QtObjectPtr ptr, int ty, char* tableName, QtObjectPtr rec, int preparedStatement);
char* QSqlDriver_StripDelimiters(QtObjectPtr ptr, char* identifier, int ty);
int QSqlDriver_SubscribeToNotification(QtObjectPtr ptr, char* name);
char* QSqlDriver_SubscribedToNotifications(QtObjectPtr ptr);
int QSqlDriver_UnsubscribeFromNotification(QtObjectPtr ptr, char* name);
void QSqlDriver_DestroyQSqlDriver(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif