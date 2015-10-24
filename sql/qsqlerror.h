#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSqlError_NewQSqlError3(QtObjectPtr other);
QtObjectPtr QSqlError_NewQSqlError(char* driverText, char* databaseText, int ty, char* code);
char* QSqlError_DatabaseText(QtObjectPtr ptr);
char* QSqlError_DriverText(QtObjectPtr ptr);
int QSqlError_IsValid(QtObjectPtr ptr);
char* QSqlError_NativeErrorCode(QtObjectPtr ptr);
char* QSqlError_Text(QtObjectPtr ptr);
int QSqlError_Type(QtObjectPtr ptr);
void QSqlError_DestroyQSqlError(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif