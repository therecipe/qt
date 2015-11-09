#ifdef __cplusplus
extern "C" {
#endif

int QListWidget_Count(void* ptr);
int QListWidget_CurrentRow(void* ptr);
int QListWidget_IsSortingEnabled(void* ptr);
void QListWidget_SetCurrentRow(void* ptr, int row);
void QListWidget_SetSortingEnabled(void* ptr, int enable);
void* QListWidget_NewQListWidget(void* parent);
void QListWidget_AddItem2(void* ptr, void* item);
void QListWidget_AddItem(void* ptr, char* label);
void QListWidget_AddItems(void* ptr, char* labels);
void QListWidget_Clear(void* ptr);
void QListWidget_ClosePersistentEditor(void* ptr, void* item);
void* QListWidget_CurrentItem(void* ptr);
void QListWidget_ConnectCurrentItemChanged(void* ptr);
void QListWidget_DisconnectCurrentItemChanged(void* ptr);
void QListWidget_ConnectCurrentRowChanged(void* ptr);
void QListWidget_DisconnectCurrentRowChanged(void* ptr);
void QListWidget_ConnectCurrentTextChanged(void* ptr);
void QListWidget_DisconnectCurrentTextChanged(void* ptr);
void QListWidget_DropEvent(void* ptr, void* event);
void QListWidget_EditItem(void* ptr, void* item);
void QListWidget_InsertItem(void* ptr, int row, void* item);
void QListWidget_InsertItem2(void* ptr, int row, char* label);
void QListWidget_InsertItems(void* ptr, int row, char* labels);
void* QListWidget_Item(void* ptr, int row);
void QListWidget_ConnectItemActivated(void* ptr);
void QListWidget_DisconnectItemActivated(void* ptr);
void* QListWidget_ItemAt(void* ptr, void* p);
void* QListWidget_ItemAt2(void* ptr, int x, int y);
void QListWidget_ConnectItemChanged(void* ptr);
void QListWidget_DisconnectItemChanged(void* ptr);
void QListWidget_ConnectItemClicked(void* ptr);
void QListWidget_DisconnectItemClicked(void* ptr);
void QListWidget_ConnectItemDoubleClicked(void* ptr);
void QListWidget_DisconnectItemDoubleClicked(void* ptr);
void QListWidget_ConnectItemEntered(void* ptr);
void QListWidget_DisconnectItemEntered(void* ptr);
void QListWidget_ConnectItemPressed(void* ptr);
void QListWidget_DisconnectItemPressed(void* ptr);
void QListWidget_ConnectItemSelectionChanged(void* ptr);
void QListWidget_DisconnectItemSelectionChanged(void* ptr);
void* QListWidget_ItemWidget(void* ptr, void* item);
void QListWidget_OpenPersistentEditor(void* ptr, void* item);
void QListWidget_RemoveItemWidget(void* ptr, void* item);
int QListWidget_Row(void* ptr, void* item);
void QListWidget_ScrollToItem(void* ptr, void* item, int hint);
void QListWidget_SetCurrentItem(void* ptr, void* item);
void QListWidget_SetCurrentItem2(void* ptr, void* item, int command);
void QListWidget_SetCurrentRow2(void* ptr, int row, int command);
void QListWidget_SetItemWidget(void* ptr, void* item, void* widget);
void QListWidget_SortItems(void* ptr, int order);
void* QListWidget_TakeItem(void* ptr, int row);
void QListWidget_DestroyQListWidget(void* ptr);

#ifdef __cplusplus
}
#endif