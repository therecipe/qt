#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QTreeWidget_New_QWidget(QtObjectPtr parent);
void QTreeWidget_Destroy(QtObjectPtr ptr);
void QTreeWidget_AddTopLevelItem_QTreeWidgetItem(QtObjectPtr ptr, QtObjectPtr item);
void QTreeWidget_ClosePersistentEditor_QTreeWidgetItem_Int(QtObjectPtr ptr, QtObjectPtr item, int column);
int QTreeWidget_ColumnCount(QtObjectPtr ptr);
int QTreeWidget_CurrentColumn(QtObjectPtr ptr);
QtObjectPtr QTreeWidget_CurrentItem(QtObjectPtr ptr);
void QTreeWidget_EditItem_QTreeWidgetItem_Int(QtObjectPtr ptr, QtObjectPtr item, int column);
QtObjectPtr QTreeWidget_HeaderItem(QtObjectPtr ptr);
int QTreeWidget_IndexOfTopLevelItem_QTreeWidgetItem(QtObjectPtr ptr, QtObjectPtr item);
void QTreeWidget_InsertTopLevelItem_Int_QTreeWidgetItem(QtObjectPtr ptr, int index, QtObjectPtr item);
QtObjectPtr QTreeWidget_InvisibleRootItem(QtObjectPtr ptr);
int QTreeWidget_IsFirstItemColumnSpanned_QTreeWidgetItem(QtObjectPtr ptr, QtObjectPtr item);
QtObjectPtr QTreeWidget_ItemAbove_QTreeWidgetItem(QtObjectPtr ptr, QtObjectPtr item);
QtObjectPtr QTreeWidget_ItemAt_Int_Int(QtObjectPtr ptr, int x, int y);
QtObjectPtr QTreeWidget_ItemBelow_QTreeWidgetItem(QtObjectPtr ptr, QtObjectPtr item);
QtObjectPtr QTreeWidget_ItemWidget_QTreeWidgetItem_Int(QtObjectPtr ptr, QtObjectPtr item, int column);
void QTreeWidget_OpenPersistentEditor_QTreeWidgetItem_Int(QtObjectPtr ptr, QtObjectPtr item, int column);
void QTreeWidget_RemoveItemWidget_QTreeWidgetItem_Int(QtObjectPtr ptr, QtObjectPtr item, int column);
void QTreeWidget_SetColumnCount_Int(QtObjectPtr ptr, int columns);
void QTreeWidget_SetCurrentItem_QTreeWidgetItem(QtObjectPtr ptr, QtObjectPtr item);
void QTreeWidget_SetCurrentItem_QTreeWidgetItem_Int(QtObjectPtr ptr, QtObjectPtr item, int column);
void QTreeWidget_SetFirstItemColumnSpanned_QTreeWidgetItem_Bool(QtObjectPtr ptr, QtObjectPtr item, int span);
void QTreeWidget_SetHeaderItem_QTreeWidgetItem(QtObjectPtr ptr, QtObjectPtr item);
void QTreeWidget_SetHeaderLabel_String(QtObjectPtr ptr, char* label);
void QTreeWidget_SetHeaderLabels_QStringList(QtObjectPtr ptr, char* labels);
void QTreeWidget_SetItemWidget_QTreeWidgetItem_Int_QWidget(QtObjectPtr ptr, QtObjectPtr item, int column, QtObjectPtr widget);
int QTreeWidget_SortColumn(QtObjectPtr ptr);
void QTreeWidget_SortItems_Int_SortOrder(QtObjectPtr ptr, int column, int order);
QtObjectPtr QTreeWidget_TakeTopLevelItem_Int(QtObjectPtr ptr, int index);
QtObjectPtr QTreeWidget_TopLevelItem_Int(QtObjectPtr ptr, int index);
int QTreeWidget_TopLevelItemCount(QtObjectPtr ptr);
//Public Slots
void QTreeWidget_ConnectSlotClear(QtObjectPtr ptr);
void QTreeWidget_DisconnectSlotClear(QtObjectPtr ptr);
void QTreeWidget_Clear(QtObjectPtr ptr);
//Signals
void QTreeWidget_ConnectSignalCurrentItemChanged(QtObjectPtr ptr);
void QTreeWidget_DisconnectSignalCurrentItemChanged(QtObjectPtr ptr);
void QTreeWidget_ConnectSignalItemActivated(QtObjectPtr ptr);
void QTreeWidget_DisconnectSignalItemActivated(QtObjectPtr ptr);
void QTreeWidget_ConnectSignalItemChanged(QtObjectPtr ptr);
void QTreeWidget_DisconnectSignalItemChanged(QtObjectPtr ptr);
void QTreeWidget_ConnectSignalItemClicked(QtObjectPtr ptr);
void QTreeWidget_DisconnectSignalItemClicked(QtObjectPtr ptr);
void QTreeWidget_ConnectSignalItemCollapsed(QtObjectPtr ptr);
void QTreeWidget_DisconnectSignalItemCollapsed(QtObjectPtr ptr);
void QTreeWidget_ConnectSignalItemDoubleClicked(QtObjectPtr ptr);
void QTreeWidget_DisconnectSignalItemDoubleClicked(QtObjectPtr ptr);
void QTreeWidget_ConnectSignalItemEntered(QtObjectPtr ptr);
void QTreeWidget_DisconnectSignalItemEntered(QtObjectPtr ptr);
void QTreeWidget_ConnectSignalItemExpanded(QtObjectPtr ptr);
void QTreeWidget_DisconnectSignalItemExpanded(QtObjectPtr ptr);
void QTreeWidget_ConnectSignalItemPressed(QtObjectPtr ptr);
void QTreeWidget_DisconnectSignalItemPressed(QtObjectPtr ptr);
void QTreeWidget_ConnectSignalItemSelectionChanged(QtObjectPtr ptr);
void QTreeWidget_DisconnectSignalItemSelectionChanged(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
