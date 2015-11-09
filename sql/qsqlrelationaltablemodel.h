#ifdef __cplusplus
extern "C" {
#endif

void QSqlRelationalTableModel_Clear(void* ptr);
void* QSqlRelationalTableModel_Data(void* ptr, void* index, int role);
void* QSqlRelationalTableModel_RelationModel(void* ptr, int column);
int QSqlRelationalTableModel_RemoveColumns(void* ptr, int column, int count, void* parent);
void QSqlRelationalTableModel_RevertRow(void* ptr, int row);
int QSqlRelationalTableModel_Select(void* ptr);
int QSqlRelationalTableModel_SetData(void* ptr, void* index, void* value, int role);
void QSqlRelationalTableModel_SetJoinMode(void* ptr, int joinMode);
void QSqlRelationalTableModel_SetRelation(void* ptr, int column, void* relation);
void QSqlRelationalTableModel_SetTable(void* ptr, char* table);
void QSqlRelationalTableModel_DestroyQSqlRelationalTableModel(void* ptr);

#ifdef __cplusplus
}
#endif