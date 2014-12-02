#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QTableView_New_QWidget(QtObjectPtr parent);
void QTableView_Destroy(QtObjectPtr ptr);
void QTableView_ClearSpans(QtObjectPtr ptr);
int QTableView_ColumnAt_Int(QtObjectPtr ptr, int x);
int QTableView_ColumnSpan_Int_Int(QtObjectPtr ptr, int row, int column);
int QTableView_ColumnViewportPosition_Int(QtObjectPtr ptr, int column);
int QTableView_ColumnWidth_Int(QtObjectPtr ptr, int column);
int QTableView_GridStyle(QtObjectPtr ptr);
int QTableView_IsColumnHidden_Int(QtObjectPtr ptr, int column);
int QTableView_IsCornerButtonEnabled(QtObjectPtr ptr);
int QTableView_IsRowHidden_Int(QtObjectPtr ptr, int row);
int QTableView_IsSortingEnabled(QtObjectPtr ptr);
int QTableView_RowAt_Int(QtObjectPtr ptr, int y);
int QTableView_RowHeight_Int(QtObjectPtr ptr, int row);
int QTableView_RowSpan_Int_Int(QtObjectPtr ptr, int row, int column);
int QTableView_RowViewportPosition_Int(QtObjectPtr ptr, int row);
void QTableView_SetColumnHidden_Int_Bool(QtObjectPtr ptr, int column, int hide);
void QTableView_SetColumnWidth_Int_Int(QtObjectPtr ptr, int column, int width);
void QTableView_SetCornerButtonEnabled_Bool(QtObjectPtr ptr, int enable);
void QTableView_SetGridStyle_PenStyle(QtObjectPtr ptr, int style);
void QTableView_SetRowHeight_Int_Int(QtObjectPtr ptr, int row, int height);
void QTableView_SetRowHidden_Int_Bool(QtObjectPtr ptr, int row, int hide);
void QTableView_SetSortingEnabled_Bool(QtObjectPtr ptr, int enable);
void QTableView_SetSpan_Int_Int_Int_Int(QtObjectPtr ptr, int row, int column, int rowSpanCount, int columnSpanCount);
void QTableView_SetWordWrap_Bool(QtObjectPtr ptr, int on);
int QTableView_ShowGrid(QtObjectPtr ptr);
void QTableView_SortByColumn_Int_SortOrder(QtObjectPtr ptr, int column, int order);
int QTableView_WordWrap(QtObjectPtr ptr);
//Public Slots
void QTableView_ConnectSlotHideColumn(QtObjectPtr ptr);
void QTableView_DisconnectSlotHideColumn(QtObjectPtr ptr);
void QTableView_HideColumn_Int(QtObjectPtr ptr, int column);
void QTableView_ConnectSlotHideRow(QtObjectPtr ptr);
void QTableView_DisconnectSlotHideRow(QtObjectPtr ptr);
void QTableView_HideRow_Int(QtObjectPtr ptr, int row);
void QTableView_ConnectSlotResizeColumnToContents(QtObjectPtr ptr);
void QTableView_DisconnectSlotResizeColumnToContents(QtObjectPtr ptr);
void QTableView_ResizeColumnToContents_Int(QtObjectPtr ptr, int column);
void QTableView_ConnectSlotResizeColumnsToContents(QtObjectPtr ptr);
void QTableView_DisconnectSlotResizeColumnsToContents(QtObjectPtr ptr);
void QTableView_ResizeColumnsToContents(QtObjectPtr ptr);
void QTableView_ConnectSlotResizeRowToContents(QtObjectPtr ptr);
void QTableView_DisconnectSlotResizeRowToContents(QtObjectPtr ptr);
void QTableView_ResizeRowToContents_Int(QtObjectPtr ptr, int row);
void QTableView_ConnectSlotResizeRowsToContents(QtObjectPtr ptr);
void QTableView_DisconnectSlotResizeRowsToContents(QtObjectPtr ptr);
void QTableView_ResizeRowsToContents(QtObjectPtr ptr);
void QTableView_ConnectSlotSelectColumn(QtObjectPtr ptr);
void QTableView_DisconnectSlotSelectColumn(QtObjectPtr ptr);
void QTableView_SelectColumn_Int(QtObjectPtr ptr, int column);
void QTableView_ConnectSlotSelectRow(QtObjectPtr ptr);
void QTableView_DisconnectSlotSelectRow(QtObjectPtr ptr);
void QTableView_SelectRow_Int(QtObjectPtr ptr, int row);
void QTableView_ConnectSlotSetShowGrid(QtObjectPtr ptr);
void QTableView_DisconnectSlotSetShowGrid(QtObjectPtr ptr);
void QTableView_SetShowGrid_Bool(QtObjectPtr ptr, int show);
void QTableView_ConnectSlotShowColumn(QtObjectPtr ptr);
void QTableView_DisconnectSlotShowColumn(QtObjectPtr ptr);
void QTableView_ShowColumn_Int(QtObjectPtr ptr, int column);
void QTableView_ConnectSlotShowRow(QtObjectPtr ptr);
void QTableView_DisconnectSlotShowRow(QtObjectPtr ptr);
void QTableView_ShowRow_Int(QtObjectPtr ptr, int row);

#ifdef __cplusplus
}
#endif
