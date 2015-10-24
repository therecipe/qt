#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QAbstractItemView_AlternatingRowColors(QtObjectPtr ptr);
int QAbstractItemView_AutoScrollMargin(QtObjectPtr ptr);
int QAbstractItemView_DefaultDropAction(QtObjectPtr ptr);
int QAbstractItemView_DragDropMode(QtObjectPtr ptr);
int QAbstractItemView_DragDropOverwriteMode(QtObjectPtr ptr);
int QAbstractItemView_DragEnabled(QtObjectPtr ptr);
int QAbstractItemView_EditTriggers(QtObjectPtr ptr);
int QAbstractItemView_HasAutoScroll(QtObjectPtr ptr);
int QAbstractItemView_HorizontalScrollMode(QtObjectPtr ptr);
int QAbstractItemView_SelectionBehavior(QtObjectPtr ptr);
int QAbstractItemView_SelectionMode(QtObjectPtr ptr);
void QAbstractItemView_SetAlternatingRowColors(QtObjectPtr ptr, int enable);
void QAbstractItemView_SetAutoScroll(QtObjectPtr ptr, int enable);
void QAbstractItemView_SetAutoScrollMargin(QtObjectPtr ptr, int margin);
void QAbstractItemView_SetDefaultDropAction(QtObjectPtr ptr, int dropAction);
void QAbstractItemView_SetDragDropMode(QtObjectPtr ptr, int behavior);
void QAbstractItemView_SetDragDropOverwriteMode(QtObjectPtr ptr, int overwrite);
void QAbstractItemView_SetDragEnabled(QtObjectPtr ptr, int enable);
void QAbstractItemView_SetDropIndicatorShown(QtObjectPtr ptr, int enable);
void QAbstractItemView_SetEditTriggers(QtObjectPtr ptr, int triggers);
void QAbstractItemView_SetHorizontalScrollMode(QtObjectPtr ptr, int mode);
void QAbstractItemView_SetIconSize(QtObjectPtr ptr, QtObjectPtr size);
void QAbstractItemView_SetSelectionBehavior(QtObjectPtr ptr, int behavior);
void QAbstractItemView_SetSelectionMode(QtObjectPtr ptr, int mode);
void QAbstractItemView_SetTabKeyNavigation(QtObjectPtr ptr, int enable);
void QAbstractItemView_SetTextElideMode(QtObjectPtr ptr, int mode);
void QAbstractItemView_SetVerticalScrollMode(QtObjectPtr ptr, int mode);
int QAbstractItemView_ShowDropIndicator(QtObjectPtr ptr);
int QAbstractItemView_TabKeyNavigation(QtObjectPtr ptr);
int QAbstractItemView_TextElideMode(QtObjectPtr ptr);
int QAbstractItemView_VerticalScrollMode(QtObjectPtr ptr);
void QAbstractItemView_ConnectActivated(QtObjectPtr ptr);
void QAbstractItemView_DisconnectActivated(QtObjectPtr ptr);
void QAbstractItemView_ClearSelection(QtObjectPtr ptr);
void QAbstractItemView_ConnectClicked(QtObjectPtr ptr);
void QAbstractItemView_DisconnectClicked(QtObjectPtr ptr);
void QAbstractItemView_ClosePersistentEditor(QtObjectPtr ptr, QtObjectPtr index);
QtObjectPtr QAbstractItemView_CurrentIndex(QtObjectPtr ptr);
void QAbstractItemView_ConnectDoubleClicked(QtObjectPtr ptr);
void QAbstractItemView_DisconnectDoubleClicked(QtObjectPtr ptr);
void QAbstractItemView_Edit(QtObjectPtr ptr, QtObjectPtr index);
void QAbstractItemView_ConnectEntered(QtObjectPtr ptr);
void QAbstractItemView_DisconnectEntered(QtObjectPtr ptr);
QtObjectPtr QAbstractItemView_IndexAt(QtObjectPtr ptr, QtObjectPtr point);
QtObjectPtr QAbstractItemView_IndexWidget(QtObjectPtr ptr, QtObjectPtr index);
char* QAbstractItemView_InputMethodQuery(QtObjectPtr ptr, int query);
QtObjectPtr QAbstractItemView_ItemDelegate(QtObjectPtr ptr);
QtObjectPtr QAbstractItemView_ItemDelegate2(QtObjectPtr ptr, QtObjectPtr index);
QtObjectPtr QAbstractItemView_ItemDelegateForColumn(QtObjectPtr ptr, int column);
QtObjectPtr QAbstractItemView_ItemDelegateForRow(QtObjectPtr ptr, int row);
void QAbstractItemView_KeyboardSearch(QtObjectPtr ptr, char* search);
QtObjectPtr QAbstractItemView_Model(QtObjectPtr ptr);
void QAbstractItemView_OpenPersistentEditor(QtObjectPtr ptr, QtObjectPtr index);
void QAbstractItemView_ConnectPressed(QtObjectPtr ptr);
void QAbstractItemView_DisconnectPressed(QtObjectPtr ptr);
void QAbstractItemView_Reset(QtObjectPtr ptr);
QtObjectPtr QAbstractItemView_RootIndex(QtObjectPtr ptr);
void QAbstractItemView_ScrollTo(QtObjectPtr ptr, QtObjectPtr index, int hint);
void QAbstractItemView_ScrollToBottom(QtObjectPtr ptr);
void QAbstractItemView_ScrollToTop(QtObjectPtr ptr);
void QAbstractItemView_SelectAll(QtObjectPtr ptr);
QtObjectPtr QAbstractItemView_SelectionModel(QtObjectPtr ptr);
void QAbstractItemView_SetCurrentIndex(QtObjectPtr ptr, QtObjectPtr index);
void QAbstractItemView_SetIndexWidget(QtObjectPtr ptr, QtObjectPtr index, QtObjectPtr widget);
void QAbstractItemView_SetItemDelegate(QtObjectPtr ptr, QtObjectPtr delegate);
void QAbstractItemView_SetItemDelegateForColumn(QtObjectPtr ptr, int column, QtObjectPtr delegate);
void QAbstractItemView_SetItemDelegateForRow(QtObjectPtr ptr, int row, QtObjectPtr delegate);
void QAbstractItemView_SetModel(QtObjectPtr ptr, QtObjectPtr model);
void QAbstractItemView_SetRootIndex(QtObjectPtr ptr, QtObjectPtr index);
void QAbstractItemView_SetSelectionModel(QtObjectPtr ptr, QtObjectPtr selectionModel);
int QAbstractItemView_SizeHintForColumn(QtObjectPtr ptr, int column);
int QAbstractItemView_SizeHintForRow(QtObjectPtr ptr, int row);
void QAbstractItemView_Update(QtObjectPtr ptr, QtObjectPtr index);
void QAbstractItemView_ConnectViewportEntered(QtObjectPtr ptr);
void QAbstractItemView_DisconnectViewportEntered(QtObjectPtr ptr);
void QAbstractItemView_DestroyQAbstractItemView(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif