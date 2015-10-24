#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QStandardItem_NewQStandardItem();
QtObjectPtr QStandardItem_NewQStandardItem3(QtObjectPtr icon, char* text);
QtObjectPtr QStandardItem_NewQStandardItem2(char* text);
QtObjectPtr QStandardItem_NewQStandardItem4(int rows, int columns);
char* QStandardItem_AccessibleDescription(QtObjectPtr ptr);
char* QStandardItem_AccessibleText(QtObjectPtr ptr);
void QStandardItem_AppendRow2(QtObjectPtr ptr, QtObjectPtr item);
int QStandardItem_CheckState(QtObjectPtr ptr);
QtObjectPtr QStandardItem_Child(QtObjectPtr ptr, int row, int column);
QtObjectPtr QStandardItem_Clone(QtObjectPtr ptr);
int QStandardItem_Column(QtObjectPtr ptr);
int QStandardItem_ColumnCount(QtObjectPtr ptr);
char* QStandardItem_Data(QtObjectPtr ptr, int role);
int QStandardItem_Flags(QtObjectPtr ptr);
int QStandardItem_HasChildren(QtObjectPtr ptr);
QtObjectPtr QStandardItem_Index(QtObjectPtr ptr);
void QStandardItem_InsertColumns(QtObjectPtr ptr, int column, int count);
void QStandardItem_InsertRow2(QtObjectPtr ptr, int row, QtObjectPtr item);
void QStandardItem_InsertRows2(QtObjectPtr ptr, int row, int count);
int QStandardItem_IsCheckable(QtObjectPtr ptr);
int QStandardItem_IsDragEnabled(QtObjectPtr ptr);
int QStandardItem_IsDropEnabled(QtObjectPtr ptr);
int QStandardItem_IsEditable(QtObjectPtr ptr);
int QStandardItem_IsEnabled(QtObjectPtr ptr);
int QStandardItem_IsSelectable(QtObjectPtr ptr);
int QStandardItem_IsTristate(QtObjectPtr ptr);
QtObjectPtr QStandardItem_Model(QtObjectPtr ptr);
QtObjectPtr QStandardItem_Parent(QtObjectPtr ptr);
void QStandardItem_Read(QtObjectPtr ptr, QtObjectPtr in);
void QStandardItem_RemoveColumn(QtObjectPtr ptr, int column);
void QStandardItem_RemoveColumns(QtObjectPtr ptr, int column, int count);
void QStandardItem_RemoveRow(QtObjectPtr ptr, int row);
void QStandardItem_RemoveRows(QtObjectPtr ptr, int row, int count);
int QStandardItem_Row(QtObjectPtr ptr);
int QStandardItem_RowCount(QtObjectPtr ptr);
void QStandardItem_SetAccessibleDescription(QtObjectPtr ptr, char* accessibleDescription);
void QStandardItem_SetAccessibleText(QtObjectPtr ptr, char* accessibleText);
void QStandardItem_SetBackground(QtObjectPtr ptr, QtObjectPtr brush);
void QStandardItem_SetCheckState(QtObjectPtr ptr, int state);
void QStandardItem_SetCheckable(QtObjectPtr ptr, int checkable);
void QStandardItem_SetChild2(QtObjectPtr ptr, int row, QtObjectPtr item);
void QStandardItem_SetChild(QtObjectPtr ptr, int row, int column, QtObjectPtr item);
void QStandardItem_SetColumnCount(QtObjectPtr ptr, int columns);
void QStandardItem_SetData(QtObjectPtr ptr, char* value, int role);
void QStandardItem_SetDragEnabled(QtObjectPtr ptr, int dragEnabled);
void QStandardItem_SetDropEnabled(QtObjectPtr ptr, int dropEnabled);
void QStandardItem_SetEditable(QtObjectPtr ptr, int editable);
void QStandardItem_SetEnabled(QtObjectPtr ptr, int enabled);
void QStandardItem_SetFlags(QtObjectPtr ptr, int flags);
void QStandardItem_SetFont(QtObjectPtr ptr, QtObjectPtr font);
void QStandardItem_SetForeground(QtObjectPtr ptr, QtObjectPtr brush);
void QStandardItem_SetIcon(QtObjectPtr ptr, QtObjectPtr icon);
void QStandardItem_SetRowCount(QtObjectPtr ptr, int rows);
void QStandardItem_SetSelectable(QtObjectPtr ptr, int selectable);
void QStandardItem_SetSizeHint(QtObjectPtr ptr, QtObjectPtr size);
void QStandardItem_SetStatusTip(QtObjectPtr ptr, char* statusTip);
void QStandardItem_SetText(QtObjectPtr ptr, char* text);
void QStandardItem_SetTextAlignment(QtObjectPtr ptr, int alignment);
void QStandardItem_SetToolTip(QtObjectPtr ptr, char* toolTip);
void QStandardItem_SetTristate(QtObjectPtr ptr, int tristate);
void QStandardItem_SetWhatsThis(QtObjectPtr ptr, char* whatsThis);
void QStandardItem_SortChildren(QtObjectPtr ptr, int column, int order);
char* QStandardItem_StatusTip(QtObjectPtr ptr);
QtObjectPtr QStandardItem_TakeChild(QtObjectPtr ptr, int row, int column);
char* QStandardItem_Text(QtObjectPtr ptr);
int QStandardItem_TextAlignment(QtObjectPtr ptr);
char* QStandardItem_ToolTip(QtObjectPtr ptr);
int QStandardItem_Type(QtObjectPtr ptr);
char* QStandardItem_WhatsThis(QtObjectPtr ptr);
void QStandardItem_Write(QtObjectPtr ptr, QtObjectPtr out);
void QStandardItem_DestroyQStandardItem(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif