#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QIdentityProxyModel_NewQIdentityProxyModel(QtObjectPtr parent);
int QIdentityProxyModel_ColumnCount(QtObjectPtr ptr, QtObjectPtr parent);
int QIdentityProxyModel_DropMimeData(QtObjectPtr ptr, QtObjectPtr data, int action, int row, int column, QtObjectPtr parent);
char* QIdentityProxyModel_HeaderData(QtObjectPtr ptr, int section, int orientation, int role);
QtObjectPtr QIdentityProxyModel_Index(QtObjectPtr ptr, int row, int column, QtObjectPtr parent);
int QIdentityProxyModel_InsertColumns(QtObjectPtr ptr, int column, int count, QtObjectPtr parent);
int QIdentityProxyModel_InsertRows(QtObjectPtr ptr, int row, int count, QtObjectPtr parent);
QtObjectPtr QIdentityProxyModel_MapFromSource(QtObjectPtr ptr, QtObjectPtr sourceIndex);
QtObjectPtr QIdentityProxyModel_MapToSource(QtObjectPtr ptr, QtObjectPtr proxyIndex);
QtObjectPtr QIdentityProxyModel_Parent(QtObjectPtr ptr, QtObjectPtr child);
int QIdentityProxyModel_RemoveColumns(QtObjectPtr ptr, int column, int count, QtObjectPtr parent);
int QIdentityProxyModel_RemoveRows(QtObjectPtr ptr, int row, int count, QtObjectPtr parent);
int QIdentityProxyModel_RowCount(QtObjectPtr ptr, QtObjectPtr parent);
void QIdentityProxyModel_SetSourceModel(QtObjectPtr ptr, QtObjectPtr newSourceModel);
QtObjectPtr QIdentityProxyModel_Sibling(QtObjectPtr ptr, int row, int column, QtObjectPtr idx);
void QIdentityProxyModel_DestroyQIdentityProxyModel(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif