#ifdef __cplusplus
extern "C" {
#endif

void* QSqlIndex_NewQSqlIndex2(void* other);
void* QSqlIndex_NewQSqlIndex(char* cursorname, char* name);
void QSqlIndex_Append(void* ptr, void* field);
void QSqlIndex_Append2(void* ptr, void* field, int desc);
char* QSqlIndex_CursorName(void* ptr);
int QSqlIndex_IsDescending(void* ptr, int i);
char* QSqlIndex_Name(void* ptr);
void QSqlIndex_SetCursorName(void* ptr, char* cursorName);
void QSqlIndex_SetDescending(void* ptr, int i, int desc);
void QSqlIndex_SetName(void* ptr, char* name);
void QSqlIndex_DestroyQSqlIndex(void* ptr);

#ifdef __cplusplus
}
#endif