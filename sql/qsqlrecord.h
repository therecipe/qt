#ifdef __cplusplus
extern "C" {
#endif

void* QSqlRecord_NewQSqlRecord();
void* QSqlRecord_NewQSqlRecord2(void* other);
void QSqlRecord_Append(void* ptr, void* field);
void QSqlRecord_Clear(void* ptr);
void QSqlRecord_ClearValues(void* ptr);
int QSqlRecord_Contains(void* ptr, char* name);
int QSqlRecord_Count(void* ptr);
char* QSqlRecord_FieldName(void* ptr, int index);
int QSqlRecord_IndexOf(void* ptr, char* name);
int QSqlRecord_IsEmpty(void* ptr);
int QSqlRecord_IsGenerated(void* ptr, char* name);
int QSqlRecord_IsGenerated2(void* ptr, int index);
int QSqlRecord_IsNull(void* ptr, char* name);
int QSqlRecord_IsNull2(void* ptr, int index);
void QSqlRecord_SetGenerated(void* ptr, char* name, int generated);
void QSqlRecord_SetGenerated2(void* ptr, int index, int generated);
void QSqlRecord_SetNull2(void* ptr, char* name);
void QSqlRecord_SetNull(void* ptr, int index);
void QSqlRecord_SetValue2(void* ptr, char* name, void* val);
void QSqlRecord_SetValue(void* ptr, int index, void* val);
void* QSqlRecord_Value2(void* ptr, char* name);
void* QSqlRecord_Value(void* ptr, int index);
void QSqlRecord_DestroyQSqlRecord(void* ptr);

#ifdef __cplusplus
}
#endif