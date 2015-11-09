#ifdef __cplusplus
extern "C" {
#endif

void* QAbstractItemModel_Sibling(void* ptr, int row, int column, void* index);
void* QAbstractItemModel_Buddy(void* ptr, void* index);
int QAbstractItemModel_CanDropMimeData(void* ptr, void* data, int action, int row, int column, void* parent);
int QAbstractItemModel_CanFetchMore(void* ptr, void* parent);
int QAbstractItemModel_ColumnCount(void* ptr, void* parent);
void QAbstractItemModel_ConnectColumnsAboutToBeInserted(void* ptr);
void QAbstractItemModel_DisconnectColumnsAboutToBeInserted(void* ptr);
void QAbstractItemModel_ConnectColumnsAboutToBeMoved(void* ptr);
void QAbstractItemModel_DisconnectColumnsAboutToBeMoved(void* ptr);
void QAbstractItemModel_ConnectColumnsAboutToBeRemoved(void* ptr);
void QAbstractItemModel_DisconnectColumnsAboutToBeRemoved(void* ptr);
void QAbstractItemModel_ConnectColumnsInserted(void* ptr);
void QAbstractItemModel_DisconnectColumnsInserted(void* ptr);
void QAbstractItemModel_ConnectColumnsMoved(void* ptr);
void QAbstractItemModel_DisconnectColumnsMoved(void* ptr);
void QAbstractItemModel_ConnectColumnsRemoved(void* ptr);
void QAbstractItemModel_DisconnectColumnsRemoved(void* ptr);
void* QAbstractItemModel_Data(void* ptr, void* index, int role);
int QAbstractItemModel_DropMimeData(void* ptr, void* data, int action, int row, int column, void* parent);
void QAbstractItemModel_FetchMore(void* ptr, void* parent);
int QAbstractItemModel_Flags(void* ptr, void* index);
int QAbstractItemModel_HasChildren(void* ptr, void* parent);
int QAbstractItemModel_HasIndex(void* ptr, int row, int column, void* parent);
void* QAbstractItemModel_HeaderData(void* ptr, int section, int orientation, int role);
void QAbstractItemModel_ConnectHeaderDataChanged(void* ptr);
void QAbstractItemModel_DisconnectHeaderDataChanged(void* ptr);
void* QAbstractItemModel_Index(void* ptr, int row, int column, void* parent);
int QAbstractItemModel_InsertColumn(void* ptr, int column, void* parent);
int QAbstractItemModel_InsertColumns(void* ptr, int column, int count, void* parent);
int QAbstractItemModel_InsertRow(void* ptr, int row, void* parent);
int QAbstractItemModel_InsertRows(void* ptr, int row, int count, void* parent);
char* QAbstractItemModel_MimeTypes(void* ptr);
void QAbstractItemModel_ConnectModelAboutToBeReset(void* ptr);
void QAbstractItemModel_DisconnectModelAboutToBeReset(void* ptr);
void QAbstractItemModel_ConnectModelReset(void* ptr);
void QAbstractItemModel_DisconnectModelReset(void* ptr);
int QAbstractItemModel_MoveColumn(void* ptr, void* sourceParent, int sourceColumn, void* destinationParent, int destinationChild);
int QAbstractItemModel_MoveColumns(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild);
int QAbstractItemModel_MoveRow(void* ptr, void* sourceParent, int sourceRow, void* destinationParent, int destinationChild);
int QAbstractItemModel_MoveRows(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild);
void* QAbstractItemModel_Parent(void* ptr, void* index);
int QAbstractItemModel_RemoveColumn(void* ptr, int column, void* parent);
int QAbstractItemModel_RemoveColumns(void* ptr, int column, int count, void* parent);
int QAbstractItemModel_RemoveRow(void* ptr, int row, void* parent);
int QAbstractItemModel_RemoveRows(void* ptr, int row, int count, void* parent);
void QAbstractItemModel_Revert(void* ptr);
int QAbstractItemModel_RowCount(void* ptr, void* parent);
void QAbstractItemModel_ConnectRowsAboutToBeInserted(void* ptr);
void QAbstractItemModel_DisconnectRowsAboutToBeInserted(void* ptr);
void QAbstractItemModel_ConnectRowsAboutToBeMoved(void* ptr);
void QAbstractItemModel_DisconnectRowsAboutToBeMoved(void* ptr);
void QAbstractItemModel_ConnectRowsAboutToBeRemoved(void* ptr);
void QAbstractItemModel_DisconnectRowsAboutToBeRemoved(void* ptr);
void QAbstractItemModel_ConnectRowsInserted(void* ptr);
void QAbstractItemModel_DisconnectRowsInserted(void* ptr);
void QAbstractItemModel_ConnectRowsMoved(void* ptr);
void QAbstractItemModel_DisconnectRowsMoved(void* ptr);
void QAbstractItemModel_ConnectRowsRemoved(void* ptr);
void QAbstractItemModel_DisconnectRowsRemoved(void* ptr);
int QAbstractItemModel_SetData(void* ptr, void* index, void* value, int role);
int QAbstractItemModel_SetHeaderData(void* ptr, int section, int orientation, void* value, int role);
void QAbstractItemModel_Sort(void* ptr, int column, int order);
int QAbstractItemModel_Submit(void* ptr);
int QAbstractItemModel_SupportedDragActions(void* ptr);
int QAbstractItemModel_SupportedDropActions(void* ptr);
void QAbstractItemModel_DestroyQAbstractItemModel(void* ptr);

#ifdef __cplusplus
}
#endif