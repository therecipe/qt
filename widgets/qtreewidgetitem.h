#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTreeWidgetItem_NewQTreeWidgetItem5(QtObjectPtr parent, QtObjectPtr preceding, int ty);
QtObjectPtr QTreeWidgetItem_NewQTreeWidgetItem4(QtObjectPtr parent, char* strin, int ty);
QtObjectPtr QTreeWidgetItem_NewQTreeWidgetItem3(QtObjectPtr parent, int ty);
QtObjectPtr QTreeWidgetItem_NewQTreeWidgetItem8(QtObjectPtr parent, QtObjectPtr preceding, int ty);
int QTreeWidgetItem_Flags(QtObjectPtr ptr);
void QTreeWidgetItem_SetFlags(QtObjectPtr ptr, int flags);
QtObjectPtr QTreeWidgetItem_NewQTreeWidgetItem7(QtObjectPtr parent, char* strin, int ty);
QtObjectPtr QTreeWidgetItem_NewQTreeWidgetItem6(QtObjectPtr parent, int ty);
QtObjectPtr QTreeWidgetItem_NewQTreeWidgetItem2(char* strin, int ty);
QtObjectPtr QTreeWidgetItem_NewQTreeWidgetItem9(QtObjectPtr other);
QtObjectPtr QTreeWidgetItem_NewQTreeWidgetItem(int ty);
void QTreeWidgetItem_AddChild(QtObjectPtr ptr, QtObjectPtr child);
int QTreeWidgetItem_CheckState(QtObjectPtr ptr, int column);
QtObjectPtr QTreeWidgetItem_Child(QtObjectPtr ptr, int index);
int QTreeWidgetItem_ChildCount(QtObjectPtr ptr);
int QTreeWidgetItem_ChildIndicatorPolicy(QtObjectPtr ptr);
int QTreeWidgetItem_ColumnCount(QtObjectPtr ptr);
char* QTreeWidgetItem_Data(QtObjectPtr ptr, int column, int role);
QtObjectPtr QTreeWidgetItem_Clone(QtObjectPtr ptr);
int QTreeWidgetItem_IndexOfChild(QtObjectPtr ptr, QtObjectPtr child);
void QTreeWidgetItem_InsertChild(QtObjectPtr ptr, int index, QtObjectPtr child);
int QTreeWidgetItem_IsDisabled(QtObjectPtr ptr);
int QTreeWidgetItem_IsExpanded(QtObjectPtr ptr);
int QTreeWidgetItem_IsFirstColumnSpanned(QtObjectPtr ptr);
int QTreeWidgetItem_IsHidden(QtObjectPtr ptr);
int QTreeWidgetItem_IsSelected(QtObjectPtr ptr);
QtObjectPtr QTreeWidgetItem_Parent(QtObjectPtr ptr);
void QTreeWidgetItem_Read(QtObjectPtr ptr, QtObjectPtr in);
void QTreeWidgetItem_RemoveChild(QtObjectPtr ptr, QtObjectPtr child);
void QTreeWidgetItem_SetBackground(QtObjectPtr ptr, int column, QtObjectPtr brush);
void QTreeWidgetItem_SetCheckState(QtObjectPtr ptr, int column, int state);
void QTreeWidgetItem_SetChildIndicatorPolicy(QtObjectPtr ptr, int policy);
void QTreeWidgetItem_SetData(QtObjectPtr ptr, int column, int role, char* value);
void QTreeWidgetItem_SetDisabled(QtObjectPtr ptr, int disabled);
void QTreeWidgetItem_SetExpanded(QtObjectPtr ptr, int expand);
void QTreeWidgetItem_SetFirstColumnSpanned(QtObjectPtr ptr, int span);
void QTreeWidgetItem_SetFont(QtObjectPtr ptr, int column, QtObjectPtr font);
void QTreeWidgetItem_SetForeground(QtObjectPtr ptr, int column, QtObjectPtr brush);
void QTreeWidgetItem_SetHidden(QtObjectPtr ptr, int hide);
void QTreeWidgetItem_SetIcon(QtObjectPtr ptr, int column, QtObjectPtr icon);
void QTreeWidgetItem_SetSelected(QtObjectPtr ptr, int sele);
void QTreeWidgetItem_SetSizeHint(QtObjectPtr ptr, int column, QtObjectPtr size);
void QTreeWidgetItem_SetStatusTip(QtObjectPtr ptr, int column, char* statusTip);
void QTreeWidgetItem_SetText(QtObjectPtr ptr, int column, char* text);
void QTreeWidgetItem_SetTextAlignment(QtObjectPtr ptr, int column, int alignment);
void QTreeWidgetItem_SetToolTip(QtObjectPtr ptr, int column, char* toolTip);
void QTreeWidgetItem_SetWhatsThis(QtObjectPtr ptr, int column, char* whatsThis);
void QTreeWidgetItem_SortChildren(QtObjectPtr ptr, int column, int order);
char* QTreeWidgetItem_StatusTip(QtObjectPtr ptr, int column);
QtObjectPtr QTreeWidgetItem_TakeChild(QtObjectPtr ptr, int index);
char* QTreeWidgetItem_Text(QtObjectPtr ptr, int column);
int QTreeWidgetItem_TextAlignment(QtObjectPtr ptr, int column);
char* QTreeWidgetItem_ToolTip(QtObjectPtr ptr, int column);
QtObjectPtr QTreeWidgetItem_TreeWidget(QtObjectPtr ptr);
int QTreeWidgetItem_Type(QtObjectPtr ptr);
char* QTreeWidgetItem_WhatsThis(QtObjectPtr ptr, int column);
void QTreeWidgetItem_Write(QtObjectPtr ptr, QtObjectPtr out);
void QTreeWidgetItem_DestroyQTreeWidgetItem(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif