#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSqlQuery_NewQSqlQuery3(QtObjectPtr db);
QtObjectPtr QSqlQuery_NewQSqlQuery(QtObjectPtr result);
QtObjectPtr QSqlQuery_NewQSqlQuery4(QtObjectPtr other);
QtObjectPtr QSqlQuery_NewQSqlQuery2(char* query, QtObjectPtr db);
int QSqlQuery_At(QtObjectPtr ptr);
char* QSqlQuery_BoundValue(QtObjectPtr ptr, char* placeholder);
char* QSqlQuery_BoundValue2(QtObjectPtr ptr, int pos);
void QSqlQuery_Clear(QtObjectPtr ptr);
QtObjectPtr QSqlQuery_Driver(QtObjectPtr ptr);
int QSqlQuery_Exec2(QtObjectPtr ptr);
int QSqlQuery_Exec(QtObjectPtr ptr, char* query);
int QSqlQuery_ExecBatch(QtObjectPtr ptr, int mode);
char* QSqlQuery_ExecutedQuery(QtObjectPtr ptr);
void QSqlQuery_Finish(QtObjectPtr ptr);
int QSqlQuery_First(QtObjectPtr ptr);
int QSqlQuery_IsActive(QtObjectPtr ptr);
int QSqlQuery_IsForwardOnly(QtObjectPtr ptr);
int QSqlQuery_IsNull2(QtObjectPtr ptr, char* name);
int QSqlQuery_IsNull(QtObjectPtr ptr, int field);
int QSqlQuery_IsSelect(QtObjectPtr ptr);
int QSqlQuery_IsValid(QtObjectPtr ptr);
int QSqlQuery_Last(QtObjectPtr ptr);
char* QSqlQuery_LastInsertId(QtObjectPtr ptr);
char* QSqlQuery_LastQuery(QtObjectPtr ptr);
int QSqlQuery_Next(QtObjectPtr ptr);
int QSqlQuery_NextResult(QtObjectPtr ptr);
int QSqlQuery_NumRowsAffected(QtObjectPtr ptr);
int QSqlQuery_Prepare(QtObjectPtr ptr, char* query);
int QSqlQuery_Previous(QtObjectPtr ptr);
QtObjectPtr QSqlQuery_Result(QtObjectPtr ptr);
int QSqlQuery_Seek(QtObjectPtr ptr, int index, int relative);
void QSqlQuery_SetForwardOnly(QtObjectPtr ptr, int forward);
int QSqlQuery_Size(QtObjectPtr ptr);
char* QSqlQuery_Value2(QtObjectPtr ptr, char* name);
char* QSqlQuery_Value(QtObjectPtr ptr, int index);
void QSqlQuery_DestroyQSqlQuery(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif