#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QTreeWidget_ColumnCount(QtObjectPtr ptr);
void QTreeWidget_SetColumnCount(QtObjectPtr ptr, int columns);
int QTreeWidget_TopLevelItemCount(QtObjectPtr ptr);
QtObjectPtr QTreeWidget_NewQTreeWidget(QtObjectPtr parent);
void QTreeWidget_AddTopLevelItem(QtObjectPtr ptr, QtObjectPtr item);
void QTreeWidget_Clear(QtObjectPtr ptr);
void QTreeWidget_ClosePersistentEditor(QtObjectPtr ptr, QtObjectPtr item, int column);
void QTreeWidget_CollapseItem(QtObjectPtr ptr, QtObjectPtr item);
int QTreeWidget_CurrentColumn(QtObjectPtr ptr);
QtObjectPtr QTreeWidget_CurrentItem(QtObjectPtr ptr);
void QTreeWidget_ConnectCurrentItemChanged(QtObjectPtr ptr);
void QTreeWidget_DisconnectCurrentItemChanged(QtObjectPtr ptr);
void QTreeWidget_EditItem(QtObjectPtr ptr, QtObjectPtr item, int column);
void QTreeWidget_ExpandItem(QtObjectPtr ptr, QtObjectPtr item);
QtObjectPtr QTreeWidget_HeaderItem(QtObjectPtr ptr);
int QTreeWidget_IndexOfTopLevelItem(QtObjectPtr ptr, QtObjectPtr item);
void QTreeWidget_InsertTopLevelItem(QtObjectPtr ptr, int index, QtObjectPtr item);
QtObjectPtr QTreeWidget_InvisibleRootItem(QtObjectPtr ptr);
int QTreeWidget_IsFirstItemColumnSpanned(QtObjectPtr ptr, QtObjectPtr item);
QtObjectPtr QTreeWidget_ItemAbove(QtObjectPtr ptr, QtObjectPtr item);
void QTreeWidget_ConnectItemActivated(QtObjectPtr ptr);
void QTreeWidget_DisconnectItemActivated(QtObjectPtr ptr);
QtObjectPtr QTreeWidget_ItemAt(QtObjectPtr ptr, QtObjectPtr p);
QtObjectPtr QTreeWidget_ItemAt2(QtObjectPtr ptr, int x, int y);
QtObjectPtr QTreeWidget_ItemBelow(QtObjectPtr ptr, QtObjectPtr item);
void QTreeWidget_ConnectItemChanged(QtObjectPtr ptr);
void QTreeWidget_DisconnectItemChanged(QtObjectPtr ptr);
void QTreeWidget_ConnectItemClicked(QtObjectPtr ptr);
void QTreeWidget_DisconnectItemClicked(QtObjectPtr ptr);
void QTreeWidget_ConnectItemCollapsed(QtObjectPtr ptr);
void QTreeWidget_DisconnectItemCollapsed(QtObjectPtr ptr);
void QTreeWidget_ConnectItemDoubleClicked(QtObjectPtr ptr);
void QTreeWidget_DisconnectItemDoubleClicked(QtObjectPtr ptr);
void QTreeWidget_ConnectItemEntered(QtObjectPtr ptr);
void QTreeWidget_DisconnectItemEntered(QtObjectPtr ptr);
void QTreeWidget_ConnectItemExpanded(QtObjectPtr ptr);
void QTreeWidget_DisconnectItemExpanded(QtObjectPtr ptr);
void QTreeWidget_ConnectItemPressed(QtObjectPtr ptr);
void QTreeWidget_DisconnectItemPressed(QtObjectPtr ptr);
void QTreeWidget_ConnectItemSelectionChanged(QtObjectPtr ptr);
void QTreeWidget_DisconnectItemSelectionChanged(QtObjectPtr ptr);
QtObjectPtr QTreeWidget_ItemWidget(QtObjectPtr ptr, QtObjectPtr item, int column);
void QTreeWidget_OpenPersistentEditor(QtObjectPtr ptr, QtObjectPtr item, int column);
void QTreeWidget_RemoveItemWidget(QtObjectPtr ptr, QtObjectPtr item, int column);
void QTreeWidget_ScrollToItem(QtObjectPtr ptr, QtObjectPtr item, int hint);
void QTreeWidget_SetCurrentItem(QtObjectPtr ptr, QtObjectPtr item);
void QTreeWidget_SetCurrentItem2(QtObjectPtr ptr, QtObjectPtr item, int column);
void QTreeWidget_SetCurrentItem3(QtObjectPtr ptr, QtObjectPtr item, int column, int command);
void QTreeWidget_SetFirstItemColumnSpanned(QtObjectPtr ptr, QtObjectPtr item, int span);
void QTreeWidget_SetHeaderItem(QtObjectPtr ptr, QtObjectPtr item);
void QTreeWidget_SetHeaderLabel(QtObjectPtr ptr, char* label);
void QTreeWidget_SetHeaderLabels(QtObjectPtr ptr, char* labels);
void QTreeWidget_SetItemWidget(QtObjectPtr ptr, QtObjectPtr item, int column, QtObjectPtr widget);
void QTreeWidget_SetSelectionModel(QtObjectPtr ptr, QtObjectPtr selectionModel);
int QTreeWidget_SortColumn(QtObjectPtr ptr);
void QTreeWidget_SortItems(QtObjectPtr ptr, int column, int order);
QtObjectPtr QTreeWidget_TakeTopLevelItem(QtObjectPtr ptr, int index);
QtObjectPtr QTreeWidget_TopLevelItem(QtObjectPtr ptr, int index);
void QTreeWidget_DestroyQTreeWidget(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif