#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QSortFilterProxyModel_DynamicSortFilter(QtObjectPtr ptr);
int QSortFilterProxyModel_FilterCaseSensitivity(QtObjectPtr ptr);
int QSortFilterProxyModel_FilterKeyColumn(QtObjectPtr ptr);
int QSortFilterProxyModel_FilterRole(QtObjectPtr ptr);
int QSortFilterProxyModel_IsSortLocaleAware(QtObjectPtr ptr);
void QSortFilterProxyModel_SetDynamicSortFilter(QtObjectPtr ptr, int enable);
void QSortFilterProxyModel_SetFilterCaseSensitivity(QtObjectPtr ptr, int cs);
void QSortFilterProxyModel_SetFilterKeyColumn(QtObjectPtr ptr, int column);
void QSortFilterProxyModel_SetFilterRegExp(QtObjectPtr ptr, QtObjectPtr regExp);
void QSortFilterProxyModel_SetFilterRole(QtObjectPtr ptr, int role);
void QSortFilterProxyModel_SetSortCaseSensitivity(QtObjectPtr ptr, int cs);
void QSortFilterProxyModel_SetSortLocaleAware(QtObjectPtr ptr, int on);
void QSortFilterProxyModel_SetSortRole(QtObjectPtr ptr, int role);
int QSortFilterProxyModel_SortCaseSensitivity(QtObjectPtr ptr);
int QSortFilterProxyModel_SortRole(QtObjectPtr ptr);
QtObjectPtr QSortFilterProxyModel_NewQSortFilterProxyModel(QtObjectPtr parent);
QtObjectPtr QSortFilterProxyModel_Buddy(QtObjectPtr ptr, QtObjectPtr index);
int QSortFilterProxyModel_CanFetchMore(QtObjectPtr ptr, QtObjectPtr parent);
int QSortFilterProxyModel_ColumnCount(QtObjectPtr ptr, QtObjectPtr parent);
char* QSortFilterProxyModel_Data(QtObjectPtr ptr, QtObjectPtr index, int role);
int QSortFilterProxyModel_DropMimeData(QtObjectPtr ptr, QtObjectPtr data, int action, int row, int column, QtObjectPtr parent);
void QSortFilterProxyModel_FetchMore(QtObjectPtr ptr, QtObjectPtr parent);
int QSortFilterProxyModel_Flags(QtObjectPtr ptr, QtObjectPtr index);
int QSortFilterProxyModel_HasChildren(QtObjectPtr ptr, QtObjectPtr parent);
char* QSortFilterProxyModel_HeaderData(QtObjectPtr ptr, int section, int orientation, int role);
QtObjectPtr QSortFilterProxyModel_Index(QtObjectPtr ptr, int row, int column, QtObjectPtr parent);
int QSortFilterProxyModel_InsertColumns(QtObjectPtr ptr, int column, int count, QtObjectPtr parent);
int QSortFilterProxyModel_InsertRows(QtObjectPtr ptr, int row, int count, QtObjectPtr parent);
void QSortFilterProxyModel_Invalidate(QtObjectPtr ptr);
QtObjectPtr QSortFilterProxyModel_MapFromSource(QtObjectPtr ptr, QtObjectPtr sourceIndex);
QtObjectPtr QSortFilterProxyModel_MapToSource(QtObjectPtr ptr, QtObjectPtr proxyIndex);
char* QSortFilterProxyModel_MimeTypes(QtObjectPtr ptr);
QtObjectPtr QSortFilterProxyModel_Parent(QtObjectPtr ptr, QtObjectPtr child);
int QSortFilterProxyModel_RemoveColumns(QtObjectPtr ptr, int column, int count, QtObjectPtr parent);
int QSortFilterProxyModel_RemoveRows(QtObjectPtr ptr, int row, int count, QtObjectPtr parent);
int QSortFilterProxyModel_RowCount(QtObjectPtr ptr, QtObjectPtr parent);
int QSortFilterProxyModel_SetData(QtObjectPtr ptr, QtObjectPtr index, char* value, int role);
void QSortFilterProxyModel_SetFilterFixedString(QtObjectPtr ptr, char* pattern);
void QSortFilterProxyModel_SetFilterRegExp2(QtObjectPtr ptr, char* pattern);
void QSortFilterProxyModel_SetFilterWildcard(QtObjectPtr ptr, char* pattern);
int QSortFilterProxyModel_SetHeaderData(QtObjectPtr ptr, int section, int orientation, char* value, int role);
void QSortFilterProxyModel_SetSourceModel(QtObjectPtr ptr, QtObjectPtr sourceModel);
QtObjectPtr QSortFilterProxyModel_Sibling(QtObjectPtr ptr, int row, int column, QtObjectPtr idx);
void QSortFilterProxyModel_Sort(QtObjectPtr ptr, int column, int order);
int QSortFilterProxyModel_SortColumn(QtObjectPtr ptr);
int QSortFilterProxyModel_SortOrder(QtObjectPtr ptr);
int QSortFilterProxyModel_SupportedDropActions(QtObjectPtr ptr);
void QSortFilterProxyModel_DestroyQSortFilterProxyModel(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif