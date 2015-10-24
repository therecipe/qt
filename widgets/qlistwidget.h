#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QListWidget_Count(QtObjectPtr ptr);
int QListWidget_CurrentRow(QtObjectPtr ptr);
int QListWidget_IsSortingEnabled(QtObjectPtr ptr);
void QListWidget_SetCurrentRow(QtObjectPtr ptr, int row);
void QListWidget_SetSortingEnabled(QtObjectPtr ptr, int enable);
QtObjectPtr QListWidget_NewQListWidget(QtObjectPtr parent);
void QListWidget_AddItem2(QtObjectPtr ptr, QtObjectPtr item);
void QListWidget_AddItem(QtObjectPtr ptr, char* label);
void QListWidget_AddItems(QtObjectPtr ptr, char* labels);
void QListWidget_Clear(QtObjectPtr ptr);
void QListWidget_ClosePersistentEditor(QtObjectPtr ptr, QtObjectPtr item);
QtObjectPtr QListWidget_CurrentItem(QtObjectPtr ptr);
void QListWidget_ConnectCurrentItemChanged(QtObjectPtr ptr);
void QListWidget_DisconnectCurrentItemChanged(QtObjectPtr ptr);
void QListWidget_ConnectCurrentRowChanged(QtObjectPtr ptr);
void QListWidget_DisconnectCurrentRowChanged(QtObjectPtr ptr);
void QListWidget_ConnectCurrentTextChanged(QtObjectPtr ptr);
void QListWidget_DisconnectCurrentTextChanged(QtObjectPtr ptr);
void QListWidget_DropEvent(QtObjectPtr ptr, QtObjectPtr event);
void QListWidget_EditItem(QtObjectPtr ptr, QtObjectPtr item);
void QListWidget_InsertItem(QtObjectPtr ptr, int row, QtObjectPtr item);
void QListWidget_InsertItem2(QtObjectPtr ptr, int row, char* label);
void QListWidget_InsertItems(QtObjectPtr ptr, int row, char* labels);
QtObjectPtr QListWidget_Item(QtObjectPtr ptr, int row);
void QListWidget_ConnectItemActivated(QtObjectPtr ptr);
void QListWidget_DisconnectItemActivated(QtObjectPtr ptr);
QtObjectPtr QListWidget_ItemAt(QtObjectPtr ptr, QtObjectPtr p);
QtObjectPtr QListWidget_ItemAt2(QtObjectPtr ptr, int x, int y);
void QListWidget_ConnectItemChanged(QtObjectPtr ptr);
void QListWidget_DisconnectItemChanged(QtObjectPtr ptr);
void QListWidget_ConnectItemClicked(QtObjectPtr ptr);
void QListWidget_DisconnectItemClicked(QtObjectPtr ptr);
void QListWidget_ConnectItemDoubleClicked(QtObjectPtr ptr);
void QListWidget_DisconnectItemDoubleClicked(QtObjectPtr ptr);
void QListWidget_ConnectItemEntered(QtObjectPtr ptr);
void QListWidget_DisconnectItemEntered(QtObjectPtr ptr);
void QListWidget_ConnectItemPressed(QtObjectPtr ptr);
void QListWidget_DisconnectItemPressed(QtObjectPtr ptr);
void QListWidget_ConnectItemSelectionChanged(QtObjectPtr ptr);
void QListWidget_DisconnectItemSelectionChanged(QtObjectPtr ptr);
QtObjectPtr QListWidget_ItemWidget(QtObjectPtr ptr, QtObjectPtr item);
void QListWidget_OpenPersistentEditor(QtObjectPtr ptr, QtObjectPtr item);
void QListWidget_RemoveItemWidget(QtObjectPtr ptr, QtObjectPtr item);
int QListWidget_Row(QtObjectPtr ptr, QtObjectPtr item);
void QListWidget_ScrollToItem(QtObjectPtr ptr, QtObjectPtr item, int hint);
void QListWidget_SetCurrentItem(QtObjectPtr ptr, QtObjectPtr item);
void QListWidget_SetCurrentItem2(QtObjectPtr ptr, QtObjectPtr item, int command);
void QListWidget_SetCurrentRow2(QtObjectPtr ptr, int row, int command);
void QListWidget_SetItemWidget(QtObjectPtr ptr, QtObjectPtr item, QtObjectPtr widget);
void QListWidget_SortItems(QtObjectPtr ptr, int order);
QtObjectPtr QListWidget_TakeItem(QtObjectPtr ptr, int row);
void QListWidget_DestroyQListWidget(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif