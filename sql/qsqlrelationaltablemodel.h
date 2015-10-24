#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QSqlRelationalTableModel_Clear(QtObjectPtr ptr);
char* QSqlRelationalTableModel_Data(QtObjectPtr ptr, QtObjectPtr index, int role);
QtObjectPtr QSqlRelationalTableModel_RelationModel(QtObjectPtr ptr, int column);
int QSqlRelationalTableModel_RemoveColumns(QtObjectPtr ptr, int column, int count, QtObjectPtr parent);
void QSqlRelationalTableModel_RevertRow(QtObjectPtr ptr, int row);
int QSqlRelationalTableModel_Select(QtObjectPtr ptr);
int QSqlRelationalTableModel_SetData(QtObjectPtr ptr, QtObjectPtr index, char* value, int role);
void QSqlRelationalTableModel_SetJoinMode(QtObjectPtr ptr, int joinMode);
void QSqlRelationalTableModel_SetRelation(QtObjectPtr ptr, int column, QtObjectPtr relation);
void QSqlRelationalTableModel_SetTable(QtObjectPtr ptr, char* table);
void QSqlRelationalTableModel_DestroyQSqlRelationalTableModel(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif