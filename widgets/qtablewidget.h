#ifdef __cplusplus
extern "C" {
#endif

void* QTableWidget_ItemAt(void* ptr, void* point);
void QTableWidget_ConnectCellActivated(void* ptr);
void QTableWidget_DisconnectCellActivated(void* ptr);
void QTableWidget_ConnectCellChanged(void* ptr);
void QTableWidget_DisconnectCellChanged(void* ptr);
void QTableWidget_ConnectCellClicked(void* ptr);
void QTableWidget_DisconnectCellClicked(void* ptr);
void QTableWidget_ConnectCellDoubleClicked(void* ptr);
void QTableWidget_DisconnectCellDoubleClicked(void* ptr);
void QTableWidget_ConnectCellEntered(void* ptr);
void QTableWidget_DisconnectCellEntered(void* ptr);
void QTableWidget_ConnectCellPressed(void* ptr);
void QTableWidget_DisconnectCellPressed(void* ptr);
void* QTableWidget_CellWidget(void* ptr, int row, int column);
void QTableWidget_Clear(void* ptr);
void QTableWidget_ClearContents(void* ptr);
void QTableWidget_ClosePersistentEditor(void* ptr, void* item);
int QTableWidget_Column(void* ptr, void* item);
int QTableWidget_ColumnCount(void* ptr);
void QTableWidget_ConnectCurrentCellChanged(void* ptr);
void QTableWidget_DisconnectCurrentCellChanged(void* ptr);
int QTableWidget_CurrentColumn(void* ptr);
void* QTableWidget_CurrentItem(void* ptr);
void QTableWidget_ConnectCurrentItemChanged(void* ptr);
void QTableWidget_DisconnectCurrentItemChanged(void* ptr);
int QTableWidget_CurrentRow(void* ptr);
void QTableWidget_EditItem(void* ptr, void* item);
void* QTableWidget_HorizontalHeaderItem(void* ptr, int column);
void QTableWidget_InsertColumn(void* ptr, int column);
void QTableWidget_InsertRow(void* ptr, int row);
void* QTableWidget_Item(void* ptr, int row, int column);
void QTableWidget_ConnectItemActivated(void* ptr);
void QTableWidget_DisconnectItemActivated(void* ptr);
void* QTableWidget_ItemAt2(void* ptr, int ax, int ay);
void QTableWidget_ConnectItemChanged(void* ptr);
void QTableWidget_DisconnectItemChanged(void* ptr);
void QTableWidget_ConnectItemClicked(void* ptr);
void QTableWidget_DisconnectItemClicked(void* ptr);
void QTableWidget_ConnectItemDoubleClicked(void* ptr);
void QTableWidget_DisconnectItemDoubleClicked(void* ptr);
void QTableWidget_ConnectItemEntered(void* ptr);
void QTableWidget_DisconnectItemEntered(void* ptr);
void QTableWidget_ConnectItemPressed(void* ptr);
void QTableWidget_DisconnectItemPressed(void* ptr);
void* QTableWidget_ItemPrototype(void* ptr);
void QTableWidget_ConnectItemSelectionChanged(void* ptr);
void QTableWidget_DisconnectItemSelectionChanged(void* ptr);
void QTableWidget_OpenPersistentEditor(void* ptr, void* item);
void QTableWidget_RemoveCellWidget(void* ptr, int row, int column);
void QTableWidget_RemoveColumn(void* ptr, int column);
void QTableWidget_RemoveRow(void* ptr, int row);
int QTableWidget_Row(void* ptr, void* item);
int QTableWidget_RowCount(void* ptr);
void QTableWidget_ScrollToItem(void* ptr, void* item, int hint);
void QTableWidget_SetCellWidget(void* ptr, int row, int column, void* widget);
void QTableWidget_SetColumnCount(void* ptr, int columns);
void QTableWidget_SetCurrentCell(void* ptr, int row, int column);
void QTableWidget_SetCurrentCell2(void* ptr, int row, int column, int command);
void QTableWidget_SetCurrentItem(void* ptr, void* item);
void QTableWidget_SetCurrentItem2(void* ptr, void* item, int command);
void QTableWidget_SetHorizontalHeaderItem(void* ptr, int column, void* item);
void QTableWidget_SetHorizontalHeaderLabels(void* ptr, char* labels);
void QTableWidget_SetItem(void* ptr, int row, int column, void* item);
void QTableWidget_SetItemPrototype(void* ptr, void* item);
void QTableWidget_SetRangeSelected(void* ptr, void* ran, int sele);
void QTableWidget_SetRowCount(void* ptr, int rows);
void QTableWidget_SetVerticalHeaderItem(void* ptr, int row, void* item);
void QTableWidget_SetVerticalHeaderLabels(void* ptr, char* labels);
void QTableWidget_SortItems(void* ptr, int column, int order);
void* QTableWidget_TakeHorizontalHeaderItem(void* ptr, int column);
void* QTableWidget_TakeItem(void* ptr, int row, int column);
void* QTableWidget_TakeVerticalHeaderItem(void* ptr, int row);
void* QTableWidget_VerticalHeaderItem(void* ptr, int row);
int QTableWidget_VisualColumn(void* ptr, int logicalColumn);
int QTableWidget_VisualRow(void* ptr, int logicalRow);
void QTableWidget_DestroyQTableWidget(void* ptr);

#ifdef __cplusplus
}
#endif