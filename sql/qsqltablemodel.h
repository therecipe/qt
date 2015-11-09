#ifdef __cplusplus
extern "C" {
#endif

void QSqlTableModel_ConnectBeforeDelete(void* ptr);
void QSqlTableModel_DisconnectBeforeDelete(void* ptr);
void QSqlTableModel_Clear(void* ptr);
void* QSqlTableModel_Data(void* ptr, void* index, int role);
int QSqlTableModel_EditStrategy(void* ptr);
int QSqlTableModel_FieldIndex(void* ptr, char* fieldName);
char* QSqlTableModel_Filter(void* ptr);
int QSqlTableModel_Flags(void* ptr, void* index);
void* QSqlTableModel_HeaderData(void* ptr, int section, int orientation, int role);
int QSqlTableModel_InsertRecord(void* ptr, int row, void* record);
int QSqlTableModel_InsertRows(void* ptr, int row, int count, void* parent);
int QSqlTableModel_IsDirty2(void* ptr);
int QSqlTableModel_IsDirty(void* ptr, void* index);
int QSqlTableModel_RemoveColumns(void* ptr, int column, int count, void* parent);
int QSqlTableModel_RemoveRows(void* ptr, int row, int count, void* parent);
void QSqlTableModel_Revert(void* ptr);
void QSqlTableModel_RevertAll(void* ptr);
void QSqlTableModel_RevertRow(void* ptr, int row);
int QSqlTableModel_RowCount(void* ptr, void* parent);
int QSqlTableModel_Select(void* ptr);
int QSqlTableModel_SelectRow(void* ptr, int row);
int QSqlTableModel_SetData(void* ptr, void* index, void* value, int role);
void QSqlTableModel_SetEditStrategy(void* ptr, int strategy);
void QSqlTableModel_SetFilter(void* ptr, char* filter);
int QSqlTableModel_SetRecord(void* ptr, int row, void* values);
void QSqlTableModel_SetSort(void* ptr, int column, int order);
void QSqlTableModel_SetTable(void* ptr, char* tableName);
void QSqlTableModel_Sort(void* ptr, int column, int order);
int QSqlTableModel_Submit(void* ptr);
int QSqlTableModel_SubmitAll(void* ptr);
char* QSqlTableModel_TableName(void* ptr);
void QSqlTableModel_DestroyQSqlTableModel(void* ptr);

#ifdef __cplusplus
}
#endif