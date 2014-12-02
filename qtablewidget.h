#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QTableWidget_New_Int_Int_QWidget(int rows, int columns, QtObjectPtr parent);
void QTableWidget_Destroy(QtObjectPtr ptr);
QtObjectPtr QTableWidget_CellWidget_Int_Int(QtObjectPtr ptr, int row, int column);
int QTableWidget_ColumnCount(QtObjectPtr ptr);
int QTableWidget_CurrentColumn(QtObjectPtr ptr);
int QTableWidget_CurrentRow(QtObjectPtr ptr);
void QTableWidget_RemoveCellWidget_Int_Int(QtObjectPtr ptr, int row, int column);
int QTableWidget_RowCount(QtObjectPtr ptr);
void QTableWidget_SetCellWidget_Int_Int_QWidget(QtObjectPtr ptr, int row, int column, QtObjectPtr widget);
void QTableWidget_SetColumnCount_Int(QtObjectPtr ptr, int columns);
void QTableWidget_SetCurrentCell_Int_Int(QtObjectPtr ptr, int row, int column);
void QTableWidget_SetRowCount_Int(QtObjectPtr ptr, int rows);
void QTableWidget_SortItems_Int_SortOrder(QtObjectPtr ptr, int column, int order);
int QTableWidget_VisualColumn_Int(QtObjectPtr ptr, int logicalColumn);
int QTableWidget_VisualRow_Int(QtObjectPtr ptr, int logicalRow);
//Public Slots
void QTableWidget_ConnectSlotClear(QtObjectPtr ptr);
void QTableWidget_DisconnectSlotClear(QtObjectPtr ptr);
void QTableWidget_Clear(QtObjectPtr ptr);
void QTableWidget_ConnectSlotClearContents(QtObjectPtr ptr);
void QTableWidget_DisconnectSlotClearContents(QtObjectPtr ptr);
void QTableWidget_ClearContents(QtObjectPtr ptr);
void QTableWidget_ConnectSlotInsertColumn(QtObjectPtr ptr);
void QTableWidget_DisconnectSlotInsertColumn(QtObjectPtr ptr);
void QTableWidget_InsertColumn_Int(QtObjectPtr ptr, int column);
void QTableWidget_ConnectSlotInsertRow(QtObjectPtr ptr);
void QTableWidget_DisconnectSlotInsertRow(QtObjectPtr ptr);
void QTableWidget_InsertRow_Int(QtObjectPtr ptr, int row);
void QTableWidget_ConnectSlotRemoveColumn(QtObjectPtr ptr);
void QTableWidget_DisconnectSlotRemoveColumn(QtObjectPtr ptr);
void QTableWidget_RemoveColumn_Int(QtObjectPtr ptr, int column);
void QTableWidget_ConnectSlotRemoveRow(QtObjectPtr ptr);
void QTableWidget_DisconnectSlotRemoveRow(QtObjectPtr ptr);
void QTableWidget_RemoveRow_Int(QtObjectPtr ptr, int row);
//Signals
void QTableWidget_ConnectSignalCellActivated(QtObjectPtr ptr);
void QTableWidget_DisconnectSignalCellActivated(QtObjectPtr ptr);
void QTableWidget_ConnectSignalCellChanged(QtObjectPtr ptr);
void QTableWidget_DisconnectSignalCellChanged(QtObjectPtr ptr);
void QTableWidget_ConnectSignalCellClicked(QtObjectPtr ptr);
void QTableWidget_DisconnectSignalCellClicked(QtObjectPtr ptr);
void QTableWidget_ConnectSignalCellDoubleClicked(QtObjectPtr ptr);
void QTableWidget_DisconnectSignalCellDoubleClicked(QtObjectPtr ptr);
void QTableWidget_ConnectSignalCellEntered(QtObjectPtr ptr);
void QTableWidget_DisconnectSignalCellEntered(QtObjectPtr ptr);
void QTableWidget_ConnectSignalCellPressed(QtObjectPtr ptr);
void QTableWidget_DisconnectSignalCellPressed(QtObjectPtr ptr);
void QTableWidget_ConnectSignalCurrentCellChanged(QtObjectPtr ptr);
void QTableWidget_DisconnectSignalCurrentCellChanged(QtObjectPtr ptr);
void QTableWidget_ConnectSignalItemSelectionChanged(QtObjectPtr ptr);
void QTableWidget_DisconnectSignalItemSelectionChanged(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
