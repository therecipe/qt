#ifdef __cplusplus
extern "C" {
#endif

void* QSqlQuery_NewQSqlQuery3(void* db);
void* QSqlQuery_NewQSqlQuery(void* result);
void* QSqlQuery_NewQSqlQuery4(void* other);
void* QSqlQuery_NewQSqlQuery2(char* query, void* db);
int QSqlQuery_At(void* ptr);
void* QSqlQuery_BoundValue(void* ptr, char* placeholder);
void* QSqlQuery_BoundValue2(void* ptr, int pos);
void QSqlQuery_Clear(void* ptr);
void* QSqlQuery_Driver(void* ptr);
int QSqlQuery_Exec2(void* ptr);
int QSqlQuery_Exec(void* ptr, char* query);
int QSqlQuery_ExecBatch(void* ptr, int mode);
char* QSqlQuery_ExecutedQuery(void* ptr);
void QSqlQuery_Finish(void* ptr);
int QSqlQuery_First(void* ptr);
int QSqlQuery_IsActive(void* ptr);
int QSqlQuery_IsForwardOnly(void* ptr);
int QSqlQuery_IsNull2(void* ptr, char* name);
int QSqlQuery_IsNull(void* ptr, int field);
int QSqlQuery_IsSelect(void* ptr);
int QSqlQuery_IsValid(void* ptr);
int QSqlQuery_Last(void* ptr);
void* QSqlQuery_LastInsertId(void* ptr);
char* QSqlQuery_LastQuery(void* ptr);
int QSqlQuery_Next(void* ptr);
int QSqlQuery_NextResult(void* ptr);
int QSqlQuery_NumRowsAffected(void* ptr);
int QSqlQuery_Prepare(void* ptr, char* query);
int QSqlQuery_Previous(void* ptr);
void* QSqlQuery_Result(void* ptr);
int QSqlQuery_Seek(void* ptr, int index, int relative);
void QSqlQuery_SetForwardOnly(void* ptr, int forward);
int QSqlQuery_Size(void* ptr);
void* QSqlQuery_Value2(void* ptr, char* name);
void* QSqlQuery_Value(void* ptr, int index);
void QSqlQuery_DestroyQSqlQuery(void* ptr);

#ifdef __cplusplus
}
#endif