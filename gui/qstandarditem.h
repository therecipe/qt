#ifdef __cplusplus
extern "C" {
#endif

void* QStandardItem_NewQStandardItem();
void* QStandardItem_NewQStandardItem3(void* icon, char* text);
void* QStandardItem_NewQStandardItem2(char* text);
void* QStandardItem_NewQStandardItem4(int rows, int columns);
char* QStandardItem_AccessibleDescription(void* ptr);
char* QStandardItem_AccessibleText(void* ptr);
void QStandardItem_AppendRow2(void* ptr, void* item);
void* QStandardItem_Background(void* ptr);
int QStandardItem_CheckState(void* ptr);
void* QStandardItem_Child(void* ptr, int row, int column);
void* QStandardItem_Clone(void* ptr);
int QStandardItem_Column(void* ptr);
int QStandardItem_ColumnCount(void* ptr);
void* QStandardItem_Data(void* ptr, int role);
int QStandardItem_Flags(void* ptr);
void* QStandardItem_Foreground(void* ptr);
int QStandardItem_HasChildren(void* ptr);
void* QStandardItem_Index(void* ptr);
void QStandardItem_InsertColumns(void* ptr, int column, int count);
void QStandardItem_InsertRow2(void* ptr, int row, void* item);
void QStandardItem_InsertRows2(void* ptr, int row, int count);
int QStandardItem_IsCheckable(void* ptr);
int QStandardItem_IsDragEnabled(void* ptr);
int QStandardItem_IsDropEnabled(void* ptr);
int QStandardItem_IsEditable(void* ptr);
int QStandardItem_IsEnabled(void* ptr);
int QStandardItem_IsSelectable(void* ptr);
int QStandardItem_IsTristate(void* ptr);
void* QStandardItem_Model(void* ptr);
void* QStandardItem_Parent(void* ptr);
void QStandardItem_Read(void* ptr, void* in);
void QStandardItem_RemoveColumn(void* ptr, int column);
void QStandardItem_RemoveColumns(void* ptr, int column, int count);
void QStandardItem_RemoveRow(void* ptr, int row);
void QStandardItem_RemoveRows(void* ptr, int row, int count);
int QStandardItem_Row(void* ptr);
int QStandardItem_RowCount(void* ptr);
void QStandardItem_SetAccessibleDescription(void* ptr, char* accessibleDescription);
void QStandardItem_SetAccessibleText(void* ptr, char* accessibleText);
void QStandardItem_SetBackground(void* ptr, void* brush);
void QStandardItem_SetCheckState(void* ptr, int state);
void QStandardItem_SetCheckable(void* ptr, int checkable);
void QStandardItem_SetChild2(void* ptr, int row, void* item);
void QStandardItem_SetChild(void* ptr, int row, int column, void* item);
void QStandardItem_SetColumnCount(void* ptr, int columns);
void QStandardItem_SetData(void* ptr, void* value, int role);
void QStandardItem_SetDragEnabled(void* ptr, int dragEnabled);
void QStandardItem_SetDropEnabled(void* ptr, int dropEnabled);
void QStandardItem_SetEditable(void* ptr, int editable);
void QStandardItem_SetEnabled(void* ptr, int enabled);
void QStandardItem_SetFlags(void* ptr, int flags);
void QStandardItem_SetFont(void* ptr, void* font);
void QStandardItem_SetForeground(void* ptr, void* brush);
void QStandardItem_SetIcon(void* ptr, void* icon);
void QStandardItem_SetRowCount(void* ptr, int rows);
void QStandardItem_SetSelectable(void* ptr, int selectable);
void QStandardItem_SetSizeHint(void* ptr, void* size);
void QStandardItem_SetStatusTip(void* ptr, char* statusTip);
void QStandardItem_SetText(void* ptr, char* text);
void QStandardItem_SetTextAlignment(void* ptr, int alignment);
void QStandardItem_SetToolTip(void* ptr, char* toolTip);
void QStandardItem_SetTristate(void* ptr, int tristate);
void QStandardItem_SetWhatsThis(void* ptr, char* whatsThis);
void QStandardItem_SortChildren(void* ptr, int column, int order);
char* QStandardItem_StatusTip(void* ptr);
void* QStandardItem_TakeChild(void* ptr, int row, int column);
char* QStandardItem_Text(void* ptr);
int QStandardItem_TextAlignment(void* ptr);
char* QStandardItem_ToolTip(void* ptr);
int QStandardItem_Type(void* ptr);
char* QStandardItem_WhatsThis(void* ptr);
void QStandardItem_Write(void* ptr, void* out);
void QStandardItem_DestroyQStandardItem(void* ptr);

#ifdef __cplusplus
}
#endif