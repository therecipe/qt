#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QAccessibleTableInterface_Caption(QtObjectPtr ptr);
QtObjectPtr QAccessibleTableInterface_CellAt(QtObjectPtr ptr, int row, int column);
int QAccessibleTableInterface_ColumnCount(QtObjectPtr ptr);
char* QAccessibleTableInterface_ColumnDescription(QtObjectPtr ptr, int column);
int QAccessibleTableInterface_IsColumnSelected(QtObjectPtr ptr, int column);
int QAccessibleTableInterface_IsRowSelected(QtObjectPtr ptr, int row);
void QAccessibleTableInterface_ModelChange(QtObjectPtr ptr, QtObjectPtr event);
int QAccessibleTableInterface_RowCount(QtObjectPtr ptr);
char* QAccessibleTableInterface_RowDescription(QtObjectPtr ptr, int row);
int QAccessibleTableInterface_SelectColumn(QtObjectPtr ptr, int column);
int QAccessibleTableInterface_SelectRow(QtObjectPtr ptr, int row);
int QAccessibleTableInterface_SelectedCellCount(QtObjectPtr ptr);
int QAccessibleTableInterface_SelectedColumnCount(QtObjectPtr ptr);
int QAccessibleTableInterface_SelectedRowCount(QtObjectPtr ptr);
QtObjectPtr QAccessibleTableInterface_Summary(QtObjectPtr ptr);
int QAccessibleTableInterface_UnselectColumn(QtObjectPtr ptr, int column);
int QAccessibleTableInterface_UnselectRow(QtObjectPtr ptr, int row);
void QAccessibleTableInterface_DestroyQAccessibleTableInterface(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif