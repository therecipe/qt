#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QSqlTableModel_ConnectBeforeDelete(QtObjectPtr ptr);
void QSqlTableModel_DisconnectBeforeDelete(QtObjectPtr ptr);
void QSqlTableModel_Clear(QtObjectPtr ptr);
char* QSqlTableModel_Data(QtObjectPtr ptr, QtObjectPtr index, int role);
int QSqlTableModel_EditStrategy(QtObjectPtr ptr);
int QSqlTableModel_FieldIndex(QtObjectPtr ptr, char* fieldName);
char* QSqlTableModel_Filter(QtObjectPtr ptr);
int QSqlTableModel_Flags(QtObjectPtr ptr, QtObjectPtr index);
char* QSqlTableModel_HeaderData(QtObjectPtr ptr, int section, int orientation, int role);
int QSqlTableModel_InsertRecord(QtObjectPtr ptr, int row, QtObjectPtr record);
int QSqlTableModel_InsertRows(QtObjectPtr ptr, int row, int count, QtObjectPtr parent);
int QSqlTableModel_IsDirty2(QtObjectPtr ptr);
int QSqlTableModel_IsDirty(QtObjectPtr ptr, QtObjectPtr index);
int QSqlTableModel_RemoveColumns(QtObjectPtr ptr, int column, int count, QtObjectPtr parent);
int QSqlTableModel_RemoveRows(QtObjectPtr ptr, int row, int count, QtObjectPtr parent);
void QSqlTableModel_Revert(QtObjectPtr ptr);
void QSqlTableModel_RevertAll(QtObjectPtr ptr);
void QSqlTableModel_RevertRow(QtObjectPtr ptr, int row);
int QSqlTableModel_RowCount(QtObjectPtr ptr, QtObjectPtr parent);
int QSqlTableModel_Select(QtObjectPtr ptr);
int QSqlTableModel_SelectRow(QtObjectPtr ptr, int row);
int QSqlTableModel_SetData(QtObjectPtr ptr, QtObjectPtr index, char* value, int role);
void QSqlTableModel_SetEditStrategy(QtObjectPtr ptr, int strategy);
void QSqlTableModel_SetFilter(QtObjectPtr ptr, char* filter);
int QSqlTableModel_SetRecord(QtObjectPtr ptr, int row, QtObjectPtr values);
void QSqlTableModel_SetSort(QtObjectPtr ptr, int column, int order);
void QSqlTableModel_SetTable(QtObjectPtr ptr, char* tableName);
void QSqlTableModel_Sort(QtObjectPtr ptr, int column, int order);
int QSqlTableModel_Submit(QtObjectPtr ptr);
int QSqlTableModel_SubmitAll(QtObjectPtr ptr);
char* QSqlTableModel_TableName(QtObjectPtr ptr);
void QSqlTableModel_DestroyQSqlTableModel(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif