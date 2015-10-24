#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSqlRecord_NewQSqlRecord();
QtObjectPtr QSqlRecord_NewQSqlRecord2(QtObjectPtr other);
void QSqlRecord_Append(QtObjectPtr ptr, QtObjectPtr field);
void QSqlRecord_Clear(QtObjectPtr ptr);
void QSqlRecord_ClearValues(QtObjectPtr ptr);
int QSqlRecord_Contains(QtObjectPtr ptr, char* name);
int QSqlRecord_Count(QtObjectPtr ptr);
char* QSqlRecord_FieldName(QtObjectPtr ptr, int index);
int QSqlRecord_IndexOf(QtObjectPtr ptr, char* name);
int QSqlRecord_IsEmpty(QtObjectPtr ptr);
int QSqlRecord_IsGenerated(QtObjectPtr ptr, char* name);
int QSqlRecord_IsGenerated2(QtObjectPtr ptr, int index);
int QSqlRecord_IsNull(QtObjectPtr ptr, char* name);
int QSqlRecord_IsNull2(QtObjectPtr ptr, int index);
void QSqlRecord_SetGenerated(QtObjectPtr ptr, char* name, int generated);
void QSqlRecord_SetGenerated2(QtObjectPtr ptr, int index, int generated);
void QSqlRecord_SetNull2(QtObjectPtr ptr, char* name);
void QSqlRecord_SetNull(QtObjectPtr ptr, int index);
void QSqlRecord_SetValue2(QtObjectPtr ptr, char* name, char* val);
void QSqlRecord_SetValue(QtObjectPtr ptr, int index, char* val);
char* QSqlRecord_Value2(QtObjectPtr ptr, char* name);
char* QSqlRecord_Value(QtObjectPtr ptr, int index);
void QSqlRecord_DestroyQSqlRecord(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif