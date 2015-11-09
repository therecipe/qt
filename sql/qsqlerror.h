#ifdef __cplusplus
extern "C" {
#endif

void* QSqlError_NewQSqlError3(void* other);
void* QSqlError_NewQSqlError(char* driverText, char* databaseText, int ty, char* code);
char* QSqlError_DatabaseText(void* ptr);
char* QSqlError_DriverText(void* ptr);
int QSqlError_IsValid(void* ptr);
char* QSqlError_NativeErrorCode(void* ptr);
char* QSqlError_Text(void* ptr);
int QSqlError_Type(void* ptr);
void QSqlError_DestroyQSqlError(void* ptr);

#ifdef __cplusplus
}
#endif