#ifdef __cplusplus
extern "C" {
#endif

int QSqlQueryModel_RowCount(void* ptr, void* parent);
void* QSqlQueryModel_Data(void* ptr, void* item, int role);
int QSqlQueryModel_CanFetchMore(void* ptr, void* parent);
void QSqlQueryModel_Clear(void* ptr);
int QSqlQueryModel_ColumnCount(void* ptr, void* index);
void QSqlQueryModel_FetchMore(void* ptr, void* parent);
void* QSqlQueryModel_HeaderData(void* ptr, int section, int orientation, int role);
int QSqlQueryModel_InsertColumns(void* ptr, int column, int count, void* parent);
int QSqlQueryModel_RemoveColumns(void* ptr, int column, int count, void* parent);
int QSqlQueryModel_SetHeaderData(void* ptr, int section, int orientation, void* value, int role);
void QSqlQueryModel_SetQuery(void* ptr, void* query);
void QSqlQueryModel_SetQuery2(void* ptr, char* query, void* db);
void QSqlQueryModel_DestroyQSqlQueryModel(void* ptr);

#ifdef __cplusplus
}
#endif