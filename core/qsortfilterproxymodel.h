#ifdef __cplusplus
extern "C" {
#endif

int QSortFilterProxyModel_DynamicSortFilter(void* ptr);
int QSortFilterProxyModel_FilterCaseSensitivity(void* ptr);
int QSortFilterProxyModel_FilterKeyColumn(void* ptr);
void* QSortFilterProxyModel_FilterRegExp(void* ptr);
int QSortFilterProxyModel_FilterRole(void* ptr);
int QSortFilterProxyModel_IsSortLocaleAware(void* ptr);
void QSortFilterProxyModel_SetDynamicSortFilter(void* ptr, int enable);
void QSortFilterProxyModel_SetFilterCaseSensitivity(void* ptr, int cs);
void QSortFilterProxyModel_SetFilterKeyColumn(void* ptr, int column);
void QSortFilterProxyModel_SetFilterRegExp(void* ptr, void* regExp);
void QSortFilterProxyModel_SetFilterRole(void* ptr, int role);
void QSortFilterProxyModel_SetSortCaseSensitivity(void* ptr, int cs);
void QSortFilterProxyModel_SetSortLocaleAware(void* ptr, int on);
void QSortFilterProxyModel_SetSortRole(void* ptr, int role);
int QSortFilterProxyModel_SortCaseSensitivity(void* ptr);
int QSortFilterProxyModel_SortRole(void* ptr);
void* QSortFilterProxyModel_NewQSortFilterProxyModel(void* parent);
void* QSortFilterProxyModel_Buddy(void* ptr, void* index);
int QSortFilterProxyModel_CanFetchMore(void* ptr, void* parent);
int QSortFilterProxyModel_ColumnCount(void* ptr, void* parent);
void* QSortFilterProxyModel_Data(void* ptr, void* index, int role);
int QSortFilterProxyModel_DropMimeData(void* ptr, void* data, int action, int row, int column, void* parent);
void QSortFilterProxyModel_FetchMore(void* ptr, void* parent);
int QSortFilterProxyModel_Flags(void* ptr, void* index);
int QSortFilterProxyModel_HasChildren(void* ptr, void* parent);
void* QSortFilterProxyModel_HeaderData(void* ptr, int section, int orientation, int role);
void* QSortFilterProxyModel_Index(void* ptr, int row, int column, void* parent);
int QSortFilterProxyModel_InsertColumns(void* ptr, int column, int count, void* parent);
int QSortFilterProxyModel_InsertRows(void* ptr, int row, int count, void* parent);
void QSortFilterProxyModel_Invalidate(void* ptr);
void* QSortFilterProxyModel_MapFromSource(void* ptr, void* sourceIndex);
void* QSortFilterProxyModel_MapToSource(void* ptr, void* proxyIndex);
char* QSortFilterProxyModel_MimeTypes(void* ptr);
void* QSortFilterProxyModel_Parent(void* ptr, void* child);
int QSortFilterProxyModel_RemoveColumns(void* ptr, int column, int count, void* parent);
int QSortFilterProxyModel_RemoveRows(void* ptr, int row, int count, void* parent);
int QSortFilterProxyModel_RowCount(void* ptr, void* parent);
int QSortFilterProxyModel_SetData(void* ptr, void* index, void* value, int role);
void QSortFilterProxyModel_SetFilterFixedString(void* ptr, char* pattern);
void QSortFilterProxyModel_SetFilterRegExp2(void* ptr, char* pattern);
void QSortFilterProxyModel_SetFilterWildcard(void* ptr, char* pattern);
int QSortFilterProxyModel_SetHeaderData(void* ptr, int section, int orientation, void* value, int role);
void QSortFilterProxyModel_SetSourceModel(void* ptr, void* sourceModel);
void* QSortFilterProxyModel_Sibling(void* ptr, int row, int column, void* idx);
void QSortFilterProxyModel_Sort(void* ptr, int column, int order);
int QSortFilterProxyModel_SortColumn(void* ptr);
int QSortFilterProxyModel_SortOrder(void* ptr);
int QSortFilterProxyModel_SupportedDropActions(void* ptr);
void QSortFilterProxyModel_DestroyQSortFilterProxyModel(void* ptr);

#ifdef __cplusplus
}
#endif