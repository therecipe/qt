#ifdef __cplusplus
extern "C" {
#endif

void* QAbstractProxyModel_Buddy(void* ptr, void* index);
int QAbstractProxyModel_CanDropMimeData(void* ptr, void* data, int action, int row, int column, void* parent);
int QAbstractProxyModel_CanFetchMore(void* ptr, void* parent);
void* QAbstractProxyModel_Data(void* ptr, void* proxyIndex, int role);
int QAbstractProxyModel_DropMimeData(void* ptr, void* data, int action, int row, int column, void* parent);
void QAbstractProxyModel_FetchMore(void* ptr, void* parent);
int QAbstractProxyModel_Flags(void* ptr, void* index);
int QAbstractProxyModel_HasChildren(void* ptr, void* parent);
void* QAbstractProxyModel_HeaderData(void* ptr, int section, int orientation, int role);
void* QAbstractProxyModel_MapFromSource(void* ptr, void* sourceIndex);
void* QAbstractProxyModel_MapToSource(void* ptr, void* proxyIndex);
char* QAbstractProxyModel_MimeTypes(void* ptr);
void QAbstractProxyModel_Revert(void* ptr);
int QAbstractProxyModel_SetData(void* ptr, void* index, void* value, int role);
int QAbstractProxyModel_SetHeaderData(void* ptr, int section, int orientation, void* value, int role);
void QAbstractProxyModel_SetSourceModel(void* ptr, void* sourceModel);
void* QAbstractProxyModel_Sibling(void* ptr, int row, int column, void* idx);
void QAbstractProxyModel_Sort(void* ptr, int column, int order);
void* QAbstractProxyModel_SourceModel(void* ptr);
void QAbstractProxyModel_ConnectSourceModelChanged(void* ptr);
void QAbstractProxyModel_DisconnectSourceModelChanged(void* ptr);
int QAbstractProxyModel_Submit(void* ptr);
int QAbstractProxyModel_SupportedDragActions(void* ptr);
int QAbstractProxyModel_SupportedDropActions(void* ptr);
void QAbstractProxyModel_DestroyQAbstractProxyModel(void* ptr);

#ifdef __cplusplus
}
#endif