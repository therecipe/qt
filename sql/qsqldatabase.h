#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSqlDatabase_NewQSqlDatabase();
QtObjectPtr QSqlDatabase_NewQSqlDatabase2(QtObjectPtr other);
void QSqlDatabase_Close(QtObjectPtr ptr);
int QSqlDatabase_Commit(QtObjectPtr ptr);
char* QSqlDatabase_ConnectOptions(QtObjectPtr ptr);
char* QSqlDatabase_ConnectionName(QtObjectPtr ptr);
char* QSqlDatabase_QSqlDatabase_ConnectionNames();
int QSqlDatabase_QSqlDatabase_Contains(char* connectionName);
char* QSqlDatabase_DatabaseName(QtObjectPtr ptr);
QtObjectPtr QSqlDatabase_Driver(QtObjectPtr ptr);
char* QSqlDatabase_DriverName(QtObjectPtr ptr);
char* QSqlDatabase_QSqlDatabase_Drivers();
char* QSqlDatabase_HostName(QtObjectPtr ptr);
int QSqlDatabase_QSqlDatabase_IsDriverAvailable(char* name);
int QSqlDatabase_IsOpen(QtObjectPtr ptr);
int QSqlDatabase_IsOpenError(QtObjectPtr ptr);
int QSqlDatabase_IsValid(QtObjectPtr ptr);
int QSqlDatabase_Open(QtObjectPtr ptr);
int QSqlDatabase_Open2(QtObjectPtr ptr, char* user, char* password);
char* QSqlDatabase_Password(QtObjectPtr ptr);
int QSqlDatabase_Port(QtObjectPtr ptr);
void QSqlDatabase_QSqlDatabase_RegisterSqlDriver(char* name, QtObjectPtr creator);
void QSqlDatabase_QSqlDatabase_RemoveDatabase(char* connectionName);
int QSqlDatabase_Rollback(QtObjectPtr ptr);
void QSqlDatabase_SetConnectOptions(QtObjectPtr ptr, char* options);
void QSqlDatabase_SetDatabaseName(QtObjectPtr ptr, char* name);
void QSqlDatabase_SetHostName(QtObjectPtr ptr, char* host);
void QSqlDatabase_SetPassword(QtObjectPtr ptr, char* password);
void QSqlDatabase_SetPort(QtObjectPtr ptr, int port);
void QSqlDatabase_SetUserName(QtObjectPtr ptr, char* name);
int QSqlDatabase_Transaction(QtObjectPtr ptr);
char* QSqlDatabase_UserName(QtObjectPtr ptr);
void QSqlDatabase_DestroyQSqlDatabase(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif