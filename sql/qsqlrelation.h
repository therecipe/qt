#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSqlRelation_NewQSqlRelation();
QtObjectPtr QSqlRelation_NewQSqlRelation2(char* tableName, char* indexColumn, char* displayColumn);
char* QSqlRelation_DisplayColumn(QtObjectPtr ptr);
char* QSqlRelation_IndexColumn(QtObjectPtr ptr);
int QSqlRelation_IsValid(QtObjectPtr ptr);
char* QSqlRelation_TableName(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif