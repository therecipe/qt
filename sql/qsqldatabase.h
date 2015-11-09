#ifdef __cplusplus
extern "C" {
#endif

void* QSqlDatabase_NewQSqlDatabase();
void* QSqlDatabase_NewQSqlDatabase2(void* other);
void QSqlDatabase_Close(void* ptr);
int QSqlDatabase_Commit(void* ptr);
char* QSqlDatabase_ConnectOptions(void* ptr);
char* QSqlDatabase_ConnectionName(void* ptr);
char* QSqlDatabase_QSqlDatabase_ConnectionNames();
int QSqlDatabase_QSqlDatabase_Contains(char* connectionName);
char* QSqlDatabase_DatabaseName(void* ptr);
void* QSqlDatabase_Driver(void* ptr);
char* QSqlDatabase_DriverName(void* ptr);
char* QSqlDatabase_QSqlDatabase_Drivers();
char* QSqlDatabase_HostName(void* ptr);
int QSqlDatabase_QSqlDatabase_IsDriverAvailable(char* name);
int QSqlDatabase_IsOpen(void* ptr);
int QSqlDatabase_IsOpenError(void* ptr);
int QSqlDatabase_IsValid(void* ptr);
int QSqlDatabase_Open(void* ptr);
int QSqlDatabase_Open2(void* ptr, char* user, char* password);
char* QSqlDatabase_Password(void* ptr);
int QSqlDatabase_Port(void* ptr);
void QSqlDatabase_QSqlDatabase_RegisterSqlDriver(char* name, void* creator);
void QSqlDatabase_QSqlDatabase_RemoveDatabase(char* connectionName);
int QSqlDatabase_Rollback(void* ptr);
void QSqlDatabase_SetConnectOptions(void* ptr, char* options);
void QSqlDatabase_SetDatabaseName(void* ptr, char* name);
void QSqlDatabase_SetHostName(void* ptr, char* host);
void QSqlDatabase_SetPassword(void* ptr, char* password);
void QSqlDatabase_SetPort(void* ptr, int port);
void QSqlDatabase_SetUserName(void* ptr, char* name);
int QSqlDatabase_Transaction(void* ptr);
char* QSqlDatabase_UserName(void* ptr);
void QSqlDatabase_DestroyQSqlDatabase(void* ptr);

#ifdef __cplusplus
}
#endif