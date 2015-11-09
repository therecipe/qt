#ifdef __cplusplus
extern "C" {
#endif

int QTableView_GridStyle(void* ptr);
int QTableView_IsCornerButtonEnabled(void* ptr);
int QTableView_IsSortingEnabled(void* ptr);
void QTableView_SetCornerButtonEnabled(void* ptr, int enable);
void QTableView_SetGridStyle(void* ptr, int style);
void QTableView_SetShowGrid(void* ptr, int show);
void QTableView_SetSpan(void* ptr, int row, int column, int rowSpanCount, int columnSpanCount);
void QTableView_SetWordWrap(void* ptr, int on);
int QTableView_ShowGrid(void* ptr);
int QTableView_WordWrap(void* ptr);
void QTableView_ClearSpans(void* ptr);
int QTableView_ColumnAt(void* ptr, int x);
int QTableView_ColumnSpan(void* ptr, int row, int column);
int QTableView_ColumnViewportPosition(void* ptr, int column);
int QTableView_ColumnWidth(void* ptr, int column);
void QTableView_HideColumn(void* ptr, int column);
void QTableView_HideRow(void* ptr, int row);
void* QTableView_HorizontalHeader(void* ptr);
void* QTableView_IndexAt(void* ptr, void* pos);
int QTableView_IsColumnHidden(void* ptr, int column);
int QTableView_IsRowHidden(void* ptr, int row);
void QTableView_ResizeColumnToContents(void* ptr, int column);
void QTableView_ResizeColumnsToContents(void* ptr);
void QTableView_ResizeRowToContents(void* ptr, int row);
void QTableView_ResizeRowsToContents(void* ptr);
int QTableView_RowAt(void* ptr, int y);
int QTableView_RowHeight(void* ptr, int row);
int QTableView_RowSpan(void* ptr, int row, int column);
int QTableView_RowViewportPosition(void* ptr, int row);
void QTableView_SelectColumn(void* ptr, int column);
void QTableView_SelectRow(void* ptr, int row);
void QTableView_SetColumnHidden(void* ptr, int column, int hide);
void QTableView_SetColumnWidth(void* ptr, int column, int width);
void QTableView_SetHorizontalHeader(void* ptr, void* header);
void QTableView_SetModel(void* ptr, void* model);
void QTableView_SetRootIndex(void* ptr, void* index);
void QTableView_SetRowHeight(void* ptr, int row, int height);
void QTableView_SetRowHidden(void* ptr, int row, int hide);
void QTableView_SetSelectionModel(void* ptr, void* selectionModel);
void QTableView_SetSortingEnabled(void* ptr, int enable);
void QTableView_SetVerticalHeader(void* ptr, void* header);
void QTableView_ShowColumn(void* ptr, int column);
void QTableView_ShowRow(void* ptr, int row);
void QTableView_SortByColumn(void* ptr, int column, int order);
void* QTableView_VerticalHeader(void* ptr);
void QTableView_DestroyQTableView(void* ptr);

#ifdef __cplusplus
}
#endif