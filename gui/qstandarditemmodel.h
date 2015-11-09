#ifdef __cplusplus
extern "C" {
#endif

void QStandardItemModel_SetSortRole(void* ptr, int role);
int QStandardItemModel_SortRole(void* ptr);
void* QStandardItemModel_NewQStandardItemModel(void* parent);
void* QStandardItemModel_NewQStandardItemModel2(int rows, int columns, void* parent);
void QStandardItemModel_AppendRow2(void* ptr, void* item);
void QStandardItemModel_Clear(void* ptr);
int QStandardItemModel_ColumnCount(void* ptr, void* parent);
void* QStandardItemModel_Data(void* ptr, void* index, int role);
int QStandardItemModel_DropMimeData(void* ptr, void* data, int action, int row, int column, void* parent);
int QStandardItemModel_Flags(void* ptr, void* index);
int QStandardItemModel_HasChildren(void* ptr, void* parent);
void* QStandardItemModel_HeaderData(void* ptr, int section, int orientation, int role);
void* QStandardItemModel_HorizontalHeaderItem(void* ptr, int column);
void* QStandardItemModel_Index(void* ptr, int row, int column, void* parent);
void* QStandardItemModel_IndexFromItem(void* ptr, void* item);
int QStandardItemModel_InsertColumn2(void* ptr, int column, void* parent);
int QStandardItemModel_InsertColumns(void* ptr, int column, int count, void* parent);
int QStandardItemModel_InsertRow2(void* ptr, int row, void* parent);
void QStandardItemModel_InsertRow3(void* ptr, int row, void* item);
int QStandardItemModel_InsertRows(void* ptr, int row, int count, void* parent);
void* QStandardItemModel_InvisibleRootItem(void* ptr);
void* QStandardItemModel_Item(void* ptr, int row, int column);
void QStandardItemModel_ConnectItemChanged(void* ptr);
void QStandardItemModel_DisconnectItemChanged(void* ptr);
void* QStandardItemModel_ItemFromIndex(void* ptr, void* index);
void* QStandardItemModel_ItemPrototype(void* ptr);
char* QStandardItemModel_MimeTypes(void* ptr);
void* QStandardItemModel_Parent(void* ptr, void* child);
int QStandardItemModel_RemoveColumns(void* ptr, int column, int count, void* parent);
int QStandardItemModel_RemoveRows(void* ptr, int row, int count, void* parent);
int QStandardItemModel_RowCount(void* ptr, void* parent);
void QStandardItemModel_SetColumnCount(void* ptr, int columns);
int QStandardItemModel_SetData(void* ptr, void* index, void* value, int role);
int QStandardItemModel_SetHeaderData(void* ptr, int section, int orientation, void* value, int role);
void QStandardItemModel_SetHorizontalHeaderItem(void* ptr, int column, void* item);
void QStandardItemModel_SetHorizontalHeaderLabels(void* ptr, char* labels);
void QStandardItemModel_SetItem2(void* ptr, int row, void* item);
void QStandardItemModel_SetItem(void* ptr, int row, int column, void* item);
void QStandardItemModel_SetItemPrototype(void* ptr, void* item);
void QStandardItemModel_SetRowCount(void* ptr, int rows);
void QStandardItemModel_SetVerticalHeaderItem(void* ptr, int row, void* item);
void QStandardItemModel_SetVerticalHeaderLabels(void* ptr, char* labels);
void* QStandardItemModel_Sibling(void* ptr, int row, int column, void* idx);
void QStandardItemModel_Sort(void* ptr, int column, int order);
int QStandardItemModel_SupportedDropActions(void* ptr);
void* QStandardItemModel_TakeHorizontalHeaderItem(void* ptr, int column);
void* QStandardItemModel_TakeItem(void* ptr, int row, int column);
void* QStandardItemModel_TakeVerticalHeaderItem(void* ptr, int row);
void* QStandardItemModel_VerticalHeaderItem(void* ptr, int row);
void QStandardItemModel_DestroyQStandardItemModel(void* ptr);

#ifdef __cplusplus
}
#endif