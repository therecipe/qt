#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QAbstractItemModel_Sibling(QtObjectPtr ptr, int row, int column, QtObjectPtr index);
QtObjectPtr QAbstractItemModel_Buddy(QtObjectPtr ptr, QtObjectPtr index);
int QAbstractItemModel_CanDropMimeData(QtObjectPtr ptr, QtObjectPtr data, int action, int row, int column, QtObjectPtr parent);
int QAbstractItemModel_CanFetchMore(QtObjectPtr ptr, QtObjectPtr parent);
int QAbstractItemModel_ColumnCount(QtObjectPtr ptr, QtObjectPtr parent);
void QAbstractItemModel_ConnectColumnsAboutToBeInserted(QtObjectPtr ptr);
void QAbstractItemModel_DisconnectColumnsAboutToBeInserted(QtObjectPtr ptr);
void QAbstractItemModel_ConnectColumnsAboutToBeMoved(QtObjectPtr ptr);
void QAbstractItemModel_DisconnectColumnsAboutToBeMoved(QtObjectPtr ptr);
void QAbstractItemModel_ConnectColumnsAboutToBeRemoved(QtObjectPtr ptr);
void QAbstractItemModel_DisconnectColumnsAboutToBeRemoved(QtObjectPtr ptr);
void QAbstractItemModel_ConnectColumnsInserted(QtObjectPtr ptr);
void QAbstractItemModel_DisconnectColumnsInserted(QtObjectPtr ptr);
void QAbstractItemModel_ConnectColumnsMoved(QtObjectPtr ptr);
void QAbstractItemModel_DisconnectColumnsMoved(QtObjectPtr ptr);
void QAbstractItemModel_ConnectColumnsRemoved(QtObjectPtr ptr);
void QAbstractItemModel_DisconnectColumnsRemoved(QtObjectPtr ptr);
char* QAbstractItemModel_Data(QtObjectPtr ptr, QtObjectPtr index, int role);
int QAbstractItemModel_DropMimeData(QtObjectPtr ptr, QtObjectPtr data, int action, int row, int column, QtObjectPtr parent);
void QAbstractItemModel_FetchMore(QtObjectPtr ptr, QtObjectPtr parent);
int QAbstractItemModel_Flags(QtObjectPtr ptr, QtObjectPtr index);
int QAbstractItemModel_HasChildren(QtObjectPtr ptr, QtObjectPtr parent);
int QAbstractItemModel_HasIndex(QtObjectPtr ptr, int row, int column, QtObjectPtr parent);
char* QAbstractItemModel_HeaderData(QtObjectPtr ptr, int section, int orientation, int role);
void QAbstractItemModel_ConnectHeaderDataChanged(QtObjectPtr ptr);
void QAbstractItemModel_DisconnectHeaderDataChanged(QtObjectPtr ptr);
QtObjectPtr QAbstractItemModel_Index(QtObjectPtr ptr, int row, int column, QtObjectPtr parent);
int QAbstractItemModel_InsertColumn(QtObjectPtr ptr, int column, QtObjectPtr parent);
int QAbstractItemModel_InsertColumns(QtObjectPtr ptr, int column, int count, QtObjectPtr parent);
int QAbstractItemModel_InsertRow(QtObjectPtr ptr, int row, QtObjectPtr parent);
int QAbstractItemModel_InsertRows(QtObjectPtr ptr, int row, int count, QtObjectPtr parent);
char* QAbstractItemModel_MimeTypes(QtObjectPtr ptr);
void QAbstractItemModel_ConnectModelAboutToBeReset(QtObjectPtr ptr);
void QAbstractItemModel_DisconnectModelAboutToBeReset(QtObjectPtr ptr);
void QAbstractItemModel_ConnectModelReset(QtObjectPtr ptr);
void QAbstractItemModel_DisconnectModelReset(QtObjectPtr ptr);
int QAbstractItemModel_MoveColumn(QtObjectPtr ptr, QtObjectPtr sourceParent, int sourceColumn, QtObjectPtr destinationParent, int destinationChild);
int QAbstractItemModel_MoveColumns(QtObjectPtr ptr, QtObjectPtr sourceParent, int sourceColumn, int count, QtObjectPtr destinationParent, int destinationChild);
int QAbstractItemModel_MoveRow(QtObjectPtr ptr, QtObjectPtr sourceParent, int sourceRow, QtObjectPtr destinationParent, int destinationChild);
int QAbstractItemModel_MoveRows(QtObjectPtr ptr, QtObjectPtr sourceParent, int sourceRow, int count, QtObjectPtr destinationParent, int destinationChild);
QtObjectPtr QAbstractItemModel_Parent(QtObjectPtr ptr, QtObjectPtr index);
int QAbstractItemModel_RemoveColumn(QtObjectPtr ptr, int column, QtObjectPtr parent);
int QAbstractItemModel_RemoveColumns(QtObjectPtr ptr, int column, int count, QtObjectPtr parent);
int QAbstractItemModel_RemoveRow(QtObjectPtr ptr, int row, QtObjectPtr parent);
int QAbstractItemModel_RemoveRows(QtObjectPtr ptr, int row, int count, QtObjectPtr parent);
void QAbstractItemModel_Revert(QtObjectPtr ptr);
int QAbstractItemModel_RowCount(QtObjectPtr ptr, QtObjectPtr parent);
void QAbstractItemModel_ConnectRowsAboutToBeInserted(QtObjectPtr ptr);
void QAbstractItemModel_DisconnectRowsAboutToBeInserted(QtObjectPtr ptr);
void QAbstractItemModel_ConnectRowsAboutToBeMoved(QtObjectPtr ptr);
void QAbstractItemModel_DisconnectRowsAboutToBeMoved(QtObjectPtr ptr);
void QAbstractItemModel_ConnectRowsAboutToBeRemoved(QtObjectPtr ptr);
void QAbstractItemModel_DisconnectRowsAboutToBeRemoved(QtObjectPtr ptr);
void QAbstractItemModel_ConnectRowsInserted(QtObjectPtr ptr);
void QAbstractItemModel_DisconnectRowsInserted(QtObjectPtr ptr);
void QAbstractItemModel_ConnectRowsMoved(QtObjectPtr ptr);
void QAbstractItemModel_DisconnectRowsMoved(QtObjectPtr ptr);
void QAbstractItemModel_ConnectRowsRemoved(QtObjectPtr ptr);
void QAbstractItemModel_DisconnectRowsRemoved(QtObjectPtr ptr);
int QAbstractItemModel_SetData(QtObjectPtr ptr, QtObjectPtr index, char* value, int role);
int QAbstractItemModel_SetHeaderData(QtObjectPtr ptr, int section, int orientation, char* value, int role);
void QAbstractItemModel_Sort(QtObjectPtr ptr, int column, int order);
int QAbstractItemModel_Submit(QtObjectPtr ptr);
int QAbstractItemModel_SupportedDragActions(QtObjectPtr ptr);
int QAbstractItemModel_SupportedDropActions(QtObjectPtr ptr);
void QAbstractItemModel_DestroyQAbstractItemModel(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif