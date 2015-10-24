#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QStandardItemModel_SetSortRole(QtObjectPtr ptr, int role);
int QStandardItemModel_SortRole(QtObjectPtr ptr);
QtObjectPtr QStandardItemModel_NewQStandardItemModel(QtObjectPtr parent);
QtObjectPtr QStandardItemModel_NewQStandardItemModel2(int rows, int columns, QtObjectPtr parent);
void QStandardItemModel_AppendRow2(QtObjectPtr ptr, QtObjectPtr item);
void QStandardItemModel_Clear(QtObjectPtr ptr);
int QStandardItemModel_ColumnCount(QtObjectPtr ptr, QtObjectPtr parent);
char* QStandardItemModel_Data(QtObjectPtr ptr, QtObjectPtr index, int role);
int QStandardItemModel_DropMimeData(QtObjectPtr ptr, QtObjectPtr data, int action, int row, int column, QtObjectPtr parent);
int QStandardItemModel_Flags(QtObjectPtr ptr, QtObjectPtr index);
int QStandardItemModel_HasChildren(QtObjectPtr ptr, QtObjectPtr parent);
char* QStandardItemModel_HeaderData(QtObjectPtr ptr, int section, int orientation, int role);
QtObjectPtr QStandardItemModel_HorizontalHeaderItem(QtObjectPtr ptr, int column);
QtObjectPtr QStandardItemModel_Index(QtObjectPtr ptr, int row, int column, QtObjectPtr parent);
QtObjectPtr QStandardItemModel_IndexFromItem(QtObjectPtr ptr, QtObjectPtr item);
int QStandardItemModel_InsertColumn2(QtObjectPtr ptr, int column, QtObjectPtr parent);
int QStandardItemModel_InsertColumns(QtObjectPtr ptr, int column, int count, QtObjectPtr parent);
int QStandardItemModel_InsertRow2(QtObjectPtr ptr, int row, QtObjectPtr parent);
void QStandardItemModel_InsertRow3(QtObjectPtr ptr, int row, QtObjectPtr item);
int QStandardItemModel_InsertRows(QtObjectPtr ptr, int row, int count, QtObjectPtr parent);
QtObjectPtr QStandardItemModel_InvisibleRootItem(QtObjectPtr ptr);
QtObjectPtr QStandardItemModel_Item(QtObjectPtr ptr, int row, int column);
void QStandardItemModel_ConnectItemChanged(QtObjectPtr ptr);
void QStandardItemModel_DisconnectItemChanged(QtObjectPtr ptr);
QtObjectPtr QStandardItemModel_ItemFromIndex(QtObjectPtr ptr, QtObjectPtr index);
QtObjectPtr QStandardItemModel_ItemPrototype(QtObjectPtr ptr);
char* QStandardItemModel_MimeTypes(QtObjectPtr ptr);
QtObjectPtr QStandardItemModel_Parent(QtObjectPtr ptr, QtObjectPtr child);
int QStandardItemModel_RemoveColumns(QtObjectPtr ptr, int column, int count, QtObjectPtr parent);
int QStandardItemModel_RemoveRows(QtObjectPtr ptr, int row, int count, QtObjectPtr parent);
int QStandardItemModel_RowCount(QtObjectPtr ptr, QtObjectPtr parent);
void QStandardItemModel_SetColumnCount(QtObjectPtr ptr, int columns);
int QStandardItemModel_SetData(QtObjectPtr ptr, QtObjectPtr index, char* value, int role);
int QStandardItemModel_SetHeaderData(QtObjectPtr ptr, int section, int orientation, char* value, int role);
void QStandardItemModel_SetHorizontalHeaderItem(QtObjectPtr ptr, int column, QtObjectPtr item);
void QStandardItemModel_SetHorizontalHeaderLabels(QtObjectPtr ptr, char* labels);
void QStandardItemModel_SetItem2(QtObjectPtr ptr, int row, QtObjectPtr item);
void QStandardItemModel_SetItem(QtObjectPtr ptr, int row, int column, QtObjectPtr item);
void QStandardItemModel_SetItemPrototype(QtObjectPtr ptr, QtObjectPtr item);
void QStandardItemModel_SetRowCount(QtObjectPtr ptr, int rows);
void QStandardItemModel_SetVerticalHeaderItem(QtObjectPtr ptr, int row, QtObjectPtr item);
void QStandardItemModel_SetVerticalHeaderLabels(QtObjectPtr ptr, char* labels);
QtObjectPtr QStandardItemModel_Sibling(QtObjectPtr ptr, int row, int column, QtObjectPtr idx);
void QStandardItemModel_Sort(QtObjectPtr ptr, int column, int order);
int QStandardItemModel_SupportedDropActions(QtObjectPtr ptr);
QtObjectPtr QStandardItemModel_TakeHorizontalHeaderItem(QtObjectPtr ptr, int column);
QtObjectPtr QStandardItemModel_TakeItem(QtObjectPtr ptr, int row, int column);
QtObjectPtr QStandardItemModel_TakeVerticalHeaderItem(QtObjectPtr ptr, int row);
QtObjectPtr QStandardItemModel_VerticalHeaderItem(QtObjectPtr ptr, int row);
void QStandardItemModel_DestroyQStandardItemModel(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif