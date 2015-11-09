#ifdef __cplusplus
extern "C" {
#endif

int QTreeWidget_ColumnCount(void* ptr);
void QTreeWidget_SetColumnCount(void* ptr, int columns);
int QTreeWidget_TopLevelItemCount(void* ptr);
void* QTreeWidget_NewQTreeWidget(void* parent);
void QTreeWidget_AddTopLevelItem(void* ptr, void* item);
void QTreeWidget_Clear(void* ptr);
void QTreeWidget_ClosePersistentEditor(void* ptr, void* item, int column);
void QTreeWidget_CollapseItem(void* ptr, void* item);
int QTreeWidget_CurrentColumn(void* ptr);
void* QTreeWidget_CurrentItem(void* ptr);
void QTreeWidget_ConnectCurrentItemChanged(void* ptr);
void QTreeWidget_DisconnectCurrentItemChanged(void* ptr);
void QTreeWidget_EditItem(void* ptr, void* item, int column);
void QTreeWidget_ExpandItem(void* ptr, void* item);
void* QTreeWidget_HeaderItem(void* ptr);
int QTreeWidget_IndexOfTopLevelItem(void* ptr, void* item);
void QTreeWidget_InsertTopLevelItem(void* ptr, int index, void* item);
void* QTreeWidget_InvisibleRootItem(void* ptr);
int QTreeWidget_IsFirstItemColumnSpanned(void* ptr, void* item);
void* QTreeWidget_ItemAbove(void* ptr, void* item);
void QTreeWidget_ConnectItemActivated(void* ptr);
void QTreeWidget_DisconnectItemActivated(void* ptr);
void* QTreeWidget_ItemAt(void* ptr, void* p);
void* QTreeWidget_ItemAt2(void* ptr, int x, int y);
void* QTreeWidget_ItemBelow(void* ptr, void* item);
void QTreeWidget_ConnectItemChanged(void* ptr);
void QTreeWidget_DisconnectItemChanged(void* ptr);
void QTreeWidget_ConnectItemClicked(void* ptr);
void QTreeWidget_DisconnectItemClicked(void* ptr);
void QTreeWidget_ConnectItemCollapsed(void* ptr);
void QTreeWidget_DisconnectItemCollapsed(void* ptr);
void QTreeWidget_ConnectItemDoubleClicked(void* ptr);
void QTreeWidget_DisconnectItemDoubleClicked(void* ptr);
void QTreeWidget_ConnectItemEntered(void* ptr);
void QTreeWidget_DisconnectItemEntered(void* ptr);
void QTreeWidget_ConnectItemExpanded(void* ptr);
void QTreeWidget_DisconnectItemExpanded(void* ptr);
void QTreeWidget_ConnectItemPressed(void* ptr);
void QTreeWidget_DisconnectItemPressed(void* ptr);
void QTreeWidget_ConnectItemSelectionChanged(void* ptr);
void QTreeWidget_DisconnectItemSelectionChanged(void* ptr);
void* QTreeWidget_ItemWidget(void* ptr, void* item, int column);
void QTreeWidget_OpenPersistentEditor(void* ptr, void* item, int column);
void QTreeWidget_RemoveItemWidget(void* ptr, void* item, int column);
void QTreeWidget_ScrollToItem(void* ptr, void* item, int hint);
void QTreeWidget_SetCurrentItem(void* ptr, void* item);
void QTreeWidget_SetCurrentItem2(void* ptr, void* item, int column);
void QTreeWidget_SetCurrentItem3(void* ptr, void* item, int column, int command);
void QTreeWidget_SetFirstItemColumnSpanned(void* ptr, void* item, int span);
void QTreeWidget_SetHeaderItem(void* ptr, void* item);
void QTreeWidget_SetHeaderLabel(void* ptr, char* label);
void QTreeWidget_SetHeaderLabels(void* ptr, char* labels);
void QTreeWidget_SetItemWidget(void* ptr, void* item, int column, void* widget);
void QTreeWidget_SetSelectionModel(void* ptr, void* selectionModel);
int QTreeWidget_SortColumn(void* ptr);
void QTreeWidget_SortItems(void* ptr, int column, int order);
void* QTreeWidget_TakeTopLevelItem(void* ptr, int index);
void* QTreeWidget_TopLevelItem(void* ptr, int index);
void QTreeWidget_DestroyQTreeWidget(void* ptr);

#ifdef __cplusplus
}
#endif