#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QSqlQueryModel_RowCount(QtObjectPtr ptr, QtObjectPtr parent);
char* QSqlQueryModel_Data(QtObjectPtr ptr, QtObjectPtr item, int role);
int QSqlQueryModel_CanFetchMore(QtObjectPtr ptr, QtObjectPtr parent);
void QSqlQueryModel_Clear(QtObjectPtr ptr);
int QSqlQueryModel_ColumnCount(QtObjectPtr ptr, QtObjectPtr index);
void QSqlQueryModel_FetchMore(QtObjectPtr ptr, QtObjectPtr parent);
char* QSqlQueryModel_HeaderData(QtObjectPtr ptr, int section, int orientation, int role);
int QSqlQueryModel_InsertColumns(QtObjectPtr ptr, int column, int count, QtObjectPtr parent);
int QSqlQueryModel_RemoveColumns(QtObjectPtr ptr, int column, int count, QtObjectPtr parent);
int QSqlQueryModel_SetHeaderData(QtObjectPtr ptr, int section, int orientation, char* value, int role);
void QSqlQueryModel_SetQuery(QtObjectPtr ptr, QtObjectPtr query);
void QSqlQueryModel_SetQuery2(QtObjectPtr ptr, char* query, QtObjectPtr db);
void QSqlQueryModel_DestroyQSqlQueryModel(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif