#ifdef __cplusplus
extern "C" {
#endif

void* QAccessibleTableInterface_Caption(void* ptr);
void* QAccessibleTableInterface_CellAt(void* ptr, int row, int column);
int QAccessibleTableInterface_ColumnCount(void* ptr);
char* QAccessibleTableInterface_ColumnDescription(void* ptr, int column);
int QAccessibleTableInterface_IsColumnSelected(void* ptr, int column);
int QAccessibleTableInterface_IsRowSelected(void* ptr, int row);
void QAccessibleTableInterface_ModelChange(void* ptr, void* event);
int QAccessibleTableInterface_RowCount(void* ptr);
char* QAccessibleTableInterface_RowDescription(void* ptr, int row);
int QAccessibleTableInterface_SelectColumn(void* ptr, int column);
int QAccessibleTableInterface_SelectRow(void* ptr, int row);
int QAccessibleTableInterface_SelectedCellCount(void* ptr);
int QAccessibleTableInterface_SelectedColumnCount(void* ptr);
int QAccessibleTableInterface_SelectedRowCount(void* ptr);
void* QAccessibleTableInterface_Summary(void* ptr);
int QAccessibleTableInterface_UnselectColumn(void* ptr, int column);
int QAccessibleTableInterface_UnselectRow(void* ptr, int row);
void QAccessibleTableInterface_DestroyQAccessibleTableInterface(void* ptr);

#ifdef __cplusplus
}
#endif