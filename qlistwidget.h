#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QListWidget_New_QWidget(QtObjectPtr parent);
void QListWidget_Destroy(QtObjectPtr ptr);
void QListWidget_AddItem_String(QtObjectPtr ptr, char* label);
void QListWidget_AddItem_QListWidgetItem(QtObjectPtr ptr, QtObjectPtr item);
void QListWidget_ClosePersistentEditor_QListWidgetItem(QtObjectPtr ptr, QtObjectPtr item);
int QListWidget_Count(QtObjectPtr ptr);
QtObjectPtr QListWidget_CurrentItem(QtObjectPtr ptr);
int QListWidget_CurrentRow(QtObjectPtr ptr);
void QListWidget_EditItem_QListWidgetItem(QtObjectPtr ptr, QtObjectPtr item);
void QListWidget_InsertItem_Int_QListWidgetItem(QtObjectPtr ptr, int row, QtObjectPtr item);
void QListWidget_InsertItem_Int_String(QtObjectPtr ptr, int row, char* label);
int QListWidget_IsSortingEnabled(QtObjectPtr ptr);
QtObjectPtr QListWidget_Item_Int(QtObjectPtr ptr, int row);
QtObjectPtr QListWidget_ItemAt_Int_Int(QtObjectPtr ptr, int x, int y);
QtObjectPtr QListWidget_ItemWidget_QListWidgetItem(QtObjectPtr ptr, QtObjectPtr item);
void QListWidget_OpenPersistentEditor_QListWidgetItem(QtObjectPtr ptr, QtObjectPtr item);
void QListWidget_RemoveItemWidget_QListWidgetItem(QtObjectPtr ptr, QtObjectPtr item);
int QListWidget_Row_QListWidgetItem(QtObjectPtr ptr, QtObjectPtr item);
void QListWidget_SetCurrentItem_QListWidgetItem(QtObjectPtr ptr, QtObjectPtr item);
void QListWidget_SetCurrentRow_Int(QtObjectPtr ptr, int row);
void QListWidget_SetItemWidget_QListWidgetItem_QWidget(QtObjectPtr ptr, QtObjectPtr item, QtObjectPtr widget);
void QListWidget_SetSortingEnabled_Bool(QtObjectPtr ptr, int enable);
void QListWidget_SortItems_SortOrder(QtObjectPtr ptr, int order);
QtObjectPtr QListWidget_TakeItem_Int(QtObjectPtr ptr, int row);
//Public Slots
void QListWidget_ConnectSlotClear(QtObjectPtr ptr);
void QListWidget_DisconnectSlotClear(QtObjectPtr ptr);
void QListWidget_Clear(QtObjectPtr ptr);
//Signals
void QListWidget_ConnectSignalCurrentItemChanged(QtObjectPtr ptr);
void QListWidget_DisconnectSignalCurrentItemChanged(QtObjectPtr ptr);
void QListWidget_ConnectSignalCurrentRowChanged(QtObjectPtr ptr);
void QListWidget_DisconnectSignalCurrentRowChanged(QtObjectPtr ptr);
void QListWidget_ConnectSignalCurrentTextChanged(QtObjectPtr ptr);
void QListWidget_DisconnectSignalCurrentTextChanged(QtObjectPtr ptr);
void QListWidget_ConnectSignalItemActivated(QtObjectPtr ptr);
void QListWidget_DisconnectSignalItemActivated(QtObjectPtr ptr);
void QListWidget_ConnectSignalItemChanged(QtObjectPtr ptr);
void QListWidget_DisconnectSignalItemChanged(QtObjectPtr ptr);
void QListWidget_ConnectSignalItemClicked(QtObjectPtr ptr);
void QListWidget_DisconnectSignalItemClicked(QtObjectPtr ptr);
void QListWidget_ConnectSignalItemDoubleClicked(QtObjectPtr ptr);
void QListWidget_DisconnectSignalItemDoubleClicked(QtObjectPtr ptr);
void QListWidget_ConnectSignalItemEntered(QtObjectPtr ptr);
void QListWidget_DisconnectSignalItemEntered(QtObjectPtr ptr);
void QListWidget_ConnectSignalItemPressed(QtObjectPtr ptr);
void QListWidget_DisconnectSignalItemPressed(QtObjectPtr ptr);
void QListWidget_ConnectSignalItemSelectionChanged(QtObjectPtr ptr);
void QListWidget_DisconnectSignalItemSelectionChanged(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
