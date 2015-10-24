#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTableWidget_ItemAt(QtObjectPtr ptr, QtObjectPtr point);
void QTableWidget_ConnectCellActivated(QtObjectPtr ptr);
void QTableWidget_DisconnectCellActivated(QtObjectPtr ptr);
void QTableWidget_ConnectCellChanged(QtObjectPtr ptr);
void QTableWidget_DisconnectCellChanged(QtObjectPtr ptr);
void QTableWidget_ConnectCellClicked(QtObjectPtr ptr);
void QTableWidget_DisconnectCellClicked(QtObjectPtr ptr);
void QTableWidget_ConnectCellDoubleClicked(QtObjectPtr ptr);
void QTableWidget_DisconnectCellDoubleClicked(QtObjectPtr ptr);
void QTableWidget_ConnectCellEntered(QtObjectPtr ptr);
void QTableWidget_DisconnectCellEntered(QtObjectPtr ptr);
void QTableWidget_ConnectCellPressed(QtObjectPtr ptr);
void QTableWidget_DisconnectCellPressed(QtObjectPtr ptr);
QtObjectPtr QTableWidget_CellWidget(QtObjectPtr ptr, int row, int column);
void QTableWidget_Clear(QtObjectPtr ptr);
void QTableWidget_ClearContents(QtObjectPtr ptr);
void QTableWidget_ClosePersistentEditor(QtObjectPtr ptr, QtObjectPtr item);
int QTableWidget_Column(QtObjectPtr ptr, QtObjectPtr item);
int QTableWidget_ColumnCount(QtObjectPtr ptr);
void QTableWidget_ConnectCurrentCellChanged(QtObjectPtr ptr);
void QTableWidget_DisconnectCurrentCellChanged(QtObjectPtr ptr);
int QTableWidget_CurrentColumn(QtObjectPtr ptr);
QtObjectPtr QTableWidget_CurrentItem(QtObjectPtr ptr);
void QTableWidget_ConnectCurrentItemChanged(QtObjectPtr ptr);
void QTableWidget_DisconnectCurrentItemChanged(QtObjectPtr ptr);
int QTableWidget_CurrentRow(QtObjectPtr ptr);
void QTableWidget_EditItem(QtObjectPtr ptr, QtObjectPtr item);
QtObjectPtr QTableWidget_HorizontalHeaderItem(QtObjectPtr ptr, int column);
void QTableWidget_InsertColumn(QtObjectPtr ptr, int column);
void QTableWidget_InsertRow(QtObjectPtr ptr, int row);
QtObjectPtr QTableWidget_Item(QtObjectPtr ptr, int row, int column);
void QTableWidget_ConnectItemActivated(QtObjectPtr ptr);
void QTableWidget_DisconnectItemActivated(QtObjectPtr ptr);
QtObjectPtr QTableWidget_ItemAt2(QtObjectPtr ptr, int ax, int ay);
void QTableWidget_ConnectItemChanged(QtObjectPtr ptr);
void QTableWidget_DisconnectItemChanged(QtObjectPtr ptr);
void QTableWidget_ConnectItemClicked(QtObjectPtr ptr);
void QTableWidget_DisconnectItemClicked(QtObjectPtr ptr);
void QTableWidget_ConnectItemDoubleClicked(QtObjectPtr ptr);
void QTableWidget_DisconnectItemDoubleClicked(QtObjectPtr ptr);
void QTableWidget_ConnectItemEntered(QtObjectPtr ptr);
void QTableWidget_DisconnectItemEntered(QtObjectPtr ptr);
void QTableWidget_ConnectItemPressed(QtObjectPtr ptr);
void QTableWidget_DisconnectItemPressed(QtObjectPtr ptr);
QtObjectPtr QTableWidget_ItemPrototype(QtObjectPtr ptr);
void QTableWidget_ConnectItemSelectionChanged(QtObjectPtr ptr);
void QTableWidget_DisconnectItemSelectionChanged(QtObjectPtr ptr);
void QTableWidget_OpenPersistentEditor(QtObjectPtr ptr, QtObjectPtr item);
void QTableWidget_RemoveCellWidget(QtObjectPtr ptr, int row, int column);
void QTableWidget_RemoveColumn(QtObjectPtr ptr, int column);
void QTableWidget_RemoveRow(QtObjectPtr ptr, int row);
int QTableWidget_Row(QtObjectPtr ptr, QtObjectPtr item);
int QTableWidget_RowCount(QtObjectPtr ptr);
void QTableWidget_ScrollToItem(QtObjectPtr ptr, QtObjectPtr item, int hint);
void QTableWidget_SetCellWidget(QtObjectPtr ptr, int row, int column, QtObjectPtr widget);
void QTableWidget_SetColumnCount(QtObjectPtr ptr, int columns);
void QTableWidget_SetCurrentCell(QtObjectPtr ptr, int row, int column);
void QTableWidget_SetCurrentCell2(QtObjectPtr ptr, int row, int column, int command);
void QTableWidget_SetCurrentItem(QtObjectPtr ptr, QtObjectPtr item);
void QTableWidget_SetCurrentItem2(QtObjectPtr ptr, QtObjectPtr item, int command);
void QTableWidget_SetHorizontalHeaderItem(QtObjectPtr ptr, int column, QtObjectPtr item);
void QTableWidget_SetHorizontalHeaderLabels(QtObjectPtr ptr, char* labels);
void QTableWidget_SetItem(QtObjectPtr ptr, int row, int column, QtObjectPtr item);
void QTableWidget_SetItemPrototype(QtObjectPtr ptr, QtObjectPtr item);
void QTableWidget_SetRangeSelected(QtObjectPtr ptr, QtObjectPtr ran, int sele);
void QTableWidget_SetRowCount(QtObjectPtr ptr, int rows);
void QTableWidget_SetVerticalHeaderItem(QtObjectPtr ptr, int row, QtObjectPtr item);
void QTableWidget_SetVerticalHeaderLabels(QtObjectPtr ptr, char* labels);
void QTableWidget_SortItems(QtObjectPtr ptr, int column, int order);
QtObjectPtr QTableWidget_TakeHorizontalHeaderItem(QtObjectPtr ptr, int column);
QtObjectPtr QTableWidget_TakeItem(QtObjectPtr ptr, int row, int column);
QtObjectPtr QTableWidget_TakeVerticalHeaderItem(QtObjectPtr ptr, int row);
QtObjectPtr QTableWidget_VerticalHeaderItem(QtObjectPtr ptr, int row);
int QTableWidget_VisualColumn(QtObjectPtr ptr, int logicalColumn);
int QTableWidget_VisualRow(QtObjectPtr ptr, int logicalRow);
void QTableWidget_DestroyQTableWidget(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif