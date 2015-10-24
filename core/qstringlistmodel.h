#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QStringListModel_Data(QtObjectPtr ptr, QtObjectPtr index, int role);
int QStringListModel_Flags(QtObjectPtr ptr, QtObjectPtr index);
int QStringListModel_InsertRows(QtObjectPtr ptr, int row, int count, QtObjectPtr parent);
int QStringListModel_RemoveRows(QtObjectPtr ptr, int row, int count, QtObjectPtr parent);
int QStringListModel_RowCount(QtObjectPtr ptr, QtObjectPtr parent);
int QStringListModel_SetData(QtObjectPtr ptr, QtObjectPtr index, char* value, int role);
void QStringListModel_SetStringList(QtObjectPtr ptr, char* strin);
QtObjectPtr QStringListModel_Sibling(QtObjectPtr ptr, int row, int column, QtObjectPtr idx);
void QStringListModel_Sort(QtObjectPtr ptr, int column, int order);
char* QStringListModel_StringList(QtObjectPtr ptr);
int QStringListModel_SupportedDropActions(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif