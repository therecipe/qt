#ifdef __cplusplus
extern "C" {
#endif

void* QTreeWidgetItem_NewQTreeWidgetItem5(void* parent, void* preceding, int ty);
void* QTreeWidgetItem_NewQTreeWidgetItem4(void* parent, char* strin, int ty);
void* QTreeWidgetItem_NewQTreeWidgetItem3(void* parent, int ty);
void* QTreeWidgetItem_NewQTreeWidgetItem8(void* parent, void* preceding, int ty);
int QTreeWidgetItem_Flags(void* ptr);
void QTreeWidgetItem_SetFlags(void* ptr, int flags);
void* QTreeWidgetItem_NewQTreeWidgetItem7(void* parent, char* strin, int ty);
void* QTreeWidgetItem_NewQTreeWidgetItem6(void* parent, int ty);
void* QTreeWidgetItem_NewQTreeWidgetItem2(char* strin, int ty);
void* QTreeWidgetItem_NewQTreeWidgetItem9(void* other);
void* QTreeWidgetItem_NewQTreeWidgetItem(int ty);
void QTreeWidgetItem_AddChild(void* ptr, void* child);
void* QTreeWidgetItem_Background(void* ptr, int column);
int QTreeWidgetItem_CheckState(void* ptr, int column);
void* QTreeWidgetItem_Child(void* ptr, int index);
int QTreeWidgetItem_ChildCount(void* ptr);
int QTreeWidgetItem_ChildIndicatorPolicy(void* ptr);
int QTreeWidgetItem_ColumnCount(void* ptr);
void* QTreeWidgetItem_Data(void* ptr, int column, int role);
void* QTreeWidgetItem_Clone(void* ptr);
void* QTreeWidgetItem_Foreground(void* ptr, int column);
int QTreeWidgetItem_IndexOfChild(void* ptr, void* child);
void QTreeWidgetItem_InsertChild(void* ptr, int index, void* child);
int QTreeWidgetItem_IsDisabled(void* ptr);
int QTreeWidgetItem_IsExpanded(void* ptr);
int QTreeWidgetItem_IsFirstColumnSpanned(void* ptr);
int QTreeWidgetItem_IsHidden(void* ptr);
int QTreeWidgetItem_IsSelected(void* ptr);
void* QTreeWidgetItem_Parent(void* ptr);
void QTreeWidgetItem_Read(void* ptr, void* in);
void QTreeWidgetItem_RemoveChild(void* ptr, void* child);
void QTreeWidgetItem_SetBackground(void* ptr, int column, void* brush);
void QTreeWidgetItem_SetCheckState(void* ptr, int column, int state);
void QTreeWidgetItem_SetChildIndicatorPolicy(void* ptr, int policy);
void QTreeWidgetItem_SetData(void* ptr, int column, int role, void* value);
void QTreeWidgetItem_SetDisabled(void* ptr, int disabled);
void QTreeWidgetItem_SetExpanded(void* ptr, int expand);
void QTreeWidgetItem_SetFirstColumnSpanned(void* ptr, int span);
void QTreeWidgetItem_SetFont(void* ptr, int column, void* font);
void QTreeWidgetItem_SetForeground(void* ptr, int column, void* brush);
void QTreeWidgetItem_SetHidden(void* ptr, int hide);
void QTreeWidgetItem_SetIcon(void* ptr, int column, void* icon);
void QTreeWidgetItem_SetSelected(void* ptr, int sele);
void QTreeWidgetItem_SetSizeHint(void* ptr, int column, void* size);
void QTreeWidgetItem_SetStatusTip(void* ptr, int column, char* statusTip);
void QTreeWidgetItem_SetText(void* ptr, int column, char* text);
void QTreeWidgetItem_SetTextAlignment(void* ptr, int column, int alignment);
void QTreeWidgetItem_SetToolTip(void* ptr, int column, char* toolTip);
void QTreeWidgetItem_SetWhatsThis(void* ptr, int column, char* whatsThis);
void QTreeWidgetItem_SortChildren(void* ptr, int column, int order);
char* QTreeWidgetItem_StatusTip(void* ptr, int column);
void* QTreeWidgetItem_TakeChild(void* ptr, int index);
char* QTreeWidgetItem_Text(void* ptr, int column);
int QTreeWidgetItem_TextAlignment(void* ptr, int column);
char* QTreeWidgetItem_ToolTip(void* ptr, int column);
void* QTreeWidgetItem_TreeWidget(void* ptr);
int QTreeWidgetItem_Type(void* ptr);
char* QTreeWidgetItem_WhatsThis(void* ptr, int column);
void QTreeWidgetItem_Write(void* ptr, void* out);
void QTreeWidgetItem_DestroyQTreeWidgetItem(void* ptr);

#ifdef __cplusplus
}
#endif