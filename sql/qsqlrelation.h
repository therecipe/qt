#ifdef __cplusplus
extern "C" {
#endif

void* QSqlRelation_NewQSqlRelation();
void* QSqlRelation_NewQSqlRelation2(char* tableName, char* indexColumn, char* displayColumn);
char* QSqlRelation_DisplayColumn(void* ptr);
char* QSqlRelation_IndexColumn(void* ptr);
int QSqlRelation_IsValid(void* ptr);
char* QSqlRelation_TableName(void* ptr);

#ifdef __cplusplus
}
#endif