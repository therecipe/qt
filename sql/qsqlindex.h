#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSqlIndex_NewQSqlIndex2(QtObjectPtr other);
QtObjectPtr QSqlIndex_NewQSqlIndex(char* cursorname, char* name);
void QSqlIndex_Append(QtObjectPtr ptr, QtObjectPtr field);
void QSqlIndex_Append2(QtObjectPtr ptr, QtObjectPtr field, int desc);
char* QSqlIndex_CursorName(QtObjectPtr ptr);
int QSqlIndex_IsDescending(QtObjectPtr ptr, int i);
char* QSqlIndex_Name(QtObjectPtr ptr);
void QSqlIndex_SetCursorName(QtObjectPtr ptr, char* cursorName);
void QSqlIndex_SetDescending(QtObjectPtr ptr, int i, int desc);
void QSqlIndex_SetName(QtObjectPtr ptr, char* name);
void QSqlIndex_DestroyQSqlIndex(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif