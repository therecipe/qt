#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QTreeWidgetItem_New_Int(int typ);
QtObjectPtr QTreeWidgetItem_New_QTreeWidget_Int(QtObjectPtr parent, int typ);
QtObjectPtr QTreeWidgetItem_New_QTreeWidget_QTreeWidgetItem_Int(QtObjectPtr parent, QtObjectPtr preceding, int typ);
QtObjectPtr QTreeWidgetItem_New_QTreeWidgetItem_Int(QtObjectPtr parent, int typ);
QtObjectPtr QTreeWidgetItem_New_QTreeWidgetItem_QTreeWidgetItem_Int(QtObjectPtr parent, QtObjectPtr preceding, int typ);
QtObjectPtr QTreeWidgetItem_New_QTreeWidgetItem(QtObjectPtr other);
void QTreeWidgetItem_Destroy(QtObjectPtr ptr);
void QTreeWidgetItem_AddChild_QTreeWidgetItem(QtObjectPtr ptr, QtObjectPtr child);
int QTreeWidgetItem_CheckState_Int(QtObjectPtr ptr, int column);
QtObjectPtr QTreeWidgetItem_Child_Int(QtObjectPtr ptr, int index);
int QTreeWidgetItem_ChildCount(QtObjectPtr ptr);
QtObjectPtr QTreeWidgetItem_Clone(QtObjectPtr ptr);
int QTreeWidgetItem_ColumnCount(QtObjectPtr ptr);
int QTreeWidgetItem_Flags(QtObjectPtr ptr);
int QTreeWidgetItem_IndexOfChild_QTreeWidgetItem(QtObjectPtr ptr, QtObjectPtr child);
void QTreeWidgetItem_InsertChild_Int_QTreeWidgetItem(QtObjectPtr ptr, int index, QtObjectPtr child);
int QTreeWidgetItem_IsDisabled(QtObjectPtr ptr);
int QTreeWidgetItem_IsExpanded(QtObjectPtr ptr);
int QTreeWidgetItem_IsFirstColumnSpanned(QtObjectPtr ptr);
int QTreeWidgetItem_IsHidden(QtObjectPtr ptr);
int QTreeWidgetItem_IsSelected(QtObjectPtr ptr);
QtObjectPtr QTreeWidgetItem_Parent(QtObjectPtr ptr);
void QTreeWidgetItem_RemoveChild_QTreeWidgetItem(QtObjectPtr ptr, QtObjectPtr child);
void QTreeWidgetItem_SetBackground_Int_QBrush(QtObjectPtr ptr, int column, QtObjectPtr brush);
void QTreeWidgetItem_SetCheckState_Int_CheckState(QtObjectPtr ptr, int column, int state);
void QTreeWidgetItem_SetDisabled_Bool(QtObjectPtr ptr, int disabled);
void QTreeWidgetItem_SetExpanded_Bool(QtObjectPtr ptr, int expand);
void QTreeWidgetItem_SetFirstColumnSpanned_Bool(QtObjectPtr ptr, int span);
void QTreeWidgetItem_SetFlags_ItemFlag(QtObjectPtr ptr, int flags);
void QTreeWidgetItem_SetForeground_Int_QBrush(QtObjectPtr ptr, int column, QtObjectPtr brush);
void QTreeWidgetItem_SetHidden_Bool(QtObjectPtr ptr, int hide);
void QTreeWidgetItem_SetSelected_Bool(QtObjectPtr ptr, int selected);
void QTreeWidgetItem_SetStatusTip_Int_String(QtObjectPtr ptr, int column, char* statusTip);
void QTreeWidgetItem_SetText_Int_String(QtObjectPtr ptr, int column, char* text);
void QTreeWidgetItem_SetTextAlignment_Int_Int(QtObjectPtr ptr, int column, int alignment);
void QTreeWidgetItem_SetToolTip_Int_String(QtObjectPtr ptr, int column, char* toolTip);
void QTreeWidgetItem_SetWhatsThis_Int_String(QtObjectPtr ptr, int column, char* whatsThis);
void QTreeWidgetItem_SortChildren_Int_SortOrder(QtObjectPtr ptr, int column, int order);
char* QTreeWidgetItem_StatusTip_Int(QtObjectPtr ptr, int column);
QtObjectPtr QTreeWidgetItem_TakeChild_Int(QtObjectPtr ptr, int index);
char* QTreeWidgetItem_Text_Int(QtObjectPtr ptr, int column);
int QTreeWidgetItem_TextAlignment_Int(QtObjectPtr ptr, int column);
char* QTreeWidgetItem_ToolTip_Int(QtObjectPtr ptr, int column);
QtObjectPtr QTreeWidgetItem_TreeWidget(QtObjectPtr ptr);
int QTreeWidgetItem_Type(QtObjectPtr ptr);
char* QTreeWidgetItem_WhatsThis_Int(QtObjectPtr ptr, int column);

#ifdef __cplusplus
}
#endif
