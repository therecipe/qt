#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QTableView_GridStyle(QtObjectPtr ptr);
int QTableView_IsCornerButtonEnabled(QtObjectPtr ptr);
int QTableView_IsSortingEnabled(QtObjectPtr ptr);
void QTableView_SetCornerButtonEnabled(QtObjectPtr ptr, int enable);
void QTableView_SetGridStyle(QtObjectPtr ptr, int style);
void QTableView_SetShowGrid(QtObjectPtr ptr, int show);
void QTableView_SetSpan(QtObjectPtr ptr, int row, int column, int rowSpanCount, int columnSpanCount);
void QTableView_SetWordWrap(QtObjectPtr ptr, int on);
int QTableView_ShowGrid(QtObjectPtr ptr);
int QTableView_WordWrap(QtObjectPtr ptr);
void QTableView_ClearSpans(QtObjectPtr ptr);
int QTableView_ColumnAt(QtObjectPtr ptr, int x);
int QTableView_ColumnSpan(QtObjectPtr ptr, int row, int column);
int QTableView_ColumnViewportPosition(QtObjectPtr ptr, int column);
int QTableView_ColumnWidth(QtObjectPtr ptr, int column);
void QTableView_HideColumn(QtObjectPtr ptr, int column);
void QTableView_HideRow(QtObjectPtr ptr, int row);
QtObjectPtr QTableView_HorizontalHeader(QtObjectPtr ptr);
QtObjectPtr QTableView_IndexAt(QtObjectPtr ptr, QtObjectPtr pos);
int QTableView_IsColumnHidden(QtObjectPtr ptr, int column);
int QTableView_IsRowHidden(QtObjectPtr ptr, int row);
void QTableView_ResizeColumnToContents(QtObjectPtr ptr, int column);
void QTableView_ResizeColumnsToContents(QtObjectPtr ptr);
void QTableView_ResizeRowToContents(QtObjectPtr ptr, int row);
void QTableView_ResizeRowsToContents(QtObjectPtr ptr);
int QTableView_RowAt(QtObjectPtr ptr, int y);
int QTableView_RowHeight(QtObjectPtr ptr, int row);
int QTableView_RowSpan(QtObjectPtr ptr, int row, int column);
int QTableView_RowViewportPosition(QtObjectPtr ptr, int row);
void QTableView_SelectColumn(QtObjectPtr ptr, int column);
void QTableView_SelectRow(QtObjectPtr ptr, int row);
void QTableView_SetColumnHidden(QtObjectPtr ptr, int column, int hide);
void QTableView_SetColumnWidth(QtObjectPtr ptr, int column, int width);
void QTableView_SetHorizontalHeader(QtObjectPtr ptr, QtObjectPtr header);
void QTableView_SetModel(QtObjectPtr ptr, QtObjectPtr model);
void QTableView_SetRootIndex(QtObjectPtr ptr, QtObjectPtr index);
void QTableView_SetRowHeight(QtObjectPtr ptr, int row, int height);
void QTableView_SetRowHidden(QtObjectPtr ptr, int row, int hide);
void QTableView_SetSelectionModel(QtObjectPtr ptr, QtObjectPtr selectionModel);
void QTableView_SetSortingEnabled(QtObjectPtr ptr, int enable);
void QTableView_SetVerticalHeader(QtObjectPtr ptr, QtObjectPtr header);
void QTableView_ShowColumn(QtObjectPtr ptr, int column);
void QTableView_ShowRow(QtObjectPtr ptr, int row);
void QTableView_SortByColumn(QtObjectPtr ptr, int column, int order);
QtObjectPtr QTableView_VerticalHeader(QtObjectPtr ptr);
void QTableView_DestroyQTableView(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif