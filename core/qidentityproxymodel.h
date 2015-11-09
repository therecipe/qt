#ifdef __cplusplus
extern "C" {
#endif

void* QIdentityProxyModel_NewQIdentityProxyModel(void* parent);
int QIdentityProxyModel_ColumnCount(void* ptr, void* parent);
int QIdentityProxyModel_DropMimeData(void* ptr, void* data, int action, int row, int column, void* parent);
void* QIdentityProxyModel_HeaderData(void* ptr, int section, int orientation, int role);
void* QIdentityProxyModel_Index(void* ptr, int row, int column, void* parent);
int QIdentityProxyModel_InsertColumns(void* ptr, int column, int count, void* parent);
int QIdentityProxyModel_InsertRows(void* ptr, int row, int count, void* parent);
void* QIdentityProxyModel_MapFromSource(void* ptr, void* sourceIndex);
void* QIdentityProxyModel_MapToSource(void* ptr, void* proxyIndex);
void* QIdentityProxyModel_Parent(void* ptr, void* child);
int QIdentityProxyModel_RemoveColumns(void* ptr, int column, int count, void* parent);
int QIdentityProxyModel_RemoveRows(void* ptr, int row, int count, void* parent);
int QIdentityProxyModel_RowCount(void* ptr, void* parent);
void QIdentityProxyModel_SetSourceModel(void* ptr, void* newSourceModel);
void* QIdentityProxyModel_Sibling(void* ptr, int row, int column, void* idx);
void QIdentityProxyModel_DestroyQIdentityProxyModel(void* ptr);

#ifdef __cplusplus
}
#endif