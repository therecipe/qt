#ifdef __cplusplus
extern "C" {
#endif

int QAbstractItemView_AlternatingRowColors(void* ptr);
int QAbstractItemView_AutoScrollMargin(void* ptr);
int QAbstractItemView_DefaultDropAction(void* ptr);
int QAbstractItemView_DragDropMode(void* ptr);
int QAbstractItemView_DragDropOverwriteMode(void* ptr);
int QAbstractItemView_DragEnabled(void* ptr);
int QAbstractItemView_EditTriggers(void* ptr);
int QAbstractItemView_HasAutoScroll(void* ptr);
int QAbstractItemView_HorizontalScrollMode(void* ptr);
int QAbstractItemView_SelectionBehavior(void* ptr);
int QAbstractItemView_SelectionMode(void* ptr);
void QAbstractItemView_SetAlternatingRowColors(void* ptr, int enable);
void QAbstractItemView_SetAutoScroll(void* ptr, int enable);
void QAbstractItemView_SetAutoScrollMargin(void* ptr, int margin);
void QAbstractItemView_SetDefaultDropAction(void* ptr, int dropAction);
void QAbstractItemView_SetDragDropMode(void* ptr, int behavior);
void QAbstractItemView_SetDragDropOverwriteMode(void* ptr, int overwrite);
void QAbstractItemView_SetDragEnabled(void* ptr, int enable);
void QAbstractItemView_SetDropIndicatorShown(void* ptr, int enable);
void QAbstractItemView_SetEditTriggers(void* ptr, int triggers);
void QAbstractItemView_SetHorizontalScrollMode(void* ptr, int mode);
void QAbstractItemView_SetIconSize(void* ptr, void* size);
void QAbstractItemView_SetSelectionBehavior(void* ptr, int behavior);
void QAbstractItemView_SetSelectionMode(void* ptr, int mode);
void QAbstractItemView_SetTabKeyNavigation(void* ptr, int enable);
void QAbstractItemView_SetTextElideMode(void* ptr, int mode);
void QAbstractItemView_SetVerticalScrollMode(void* ptr, int mode);
int QAbstractItemView_ShowDropIndicator(void* ptr);
int QAbstractItemView_TabKeyNavigation(void* ptr);
int QAbstractItemView_TextElideMode(void* ptr);
int QAbstractItemView_VerticalScrollMode(void* ptr);
void QAbstractItemView_ConnectActivated(void* ptr);
void QAbstractItemView_DisconnectActivated(void* ptr);
void QAbstractItemView_ClearSelection(void* ptr);
void QAbstractItemView_ConnectClicked(void* ptr);
void QAbstractItemView_DisconnectClicked(void* ptr);
void QAbstractItemView_ClosePersistentEditor(void* ptr, void* index);
void* QAbstractItemView_CurrentIndex(void* ptr);
void QAbstractItemView_ConnectDoubleClicked(void* ptr);
void QAbstractItemView_DisconnectDoubleClicked(void* ptr);
void QAbstractItemView_Edit(void* ptr, void* index);
void QAbstractItemView_ConnectEntered(void* ptr);
void QAbstractItemView_DisconnectEntered(void* ptr);
void* QAbstractItemView_IndexAt(void* ptr, void* point);
void* QAbstractItemView_IndexWidget(void* ptr, void* index);
void* QAbstractItemView_InputMethodQuery(void* ptr, int query);
void* QAbstractItemView_ItemDelegate(void* ptr);
void* QAbstractItemView_ItemDelegate2(void* ptr, void* index);
void* QAbstractItemView_ItemDelegateForColumn(void* ptr, int column);
void* QAbstractItemView_ItemDelegateForRow(void* ptr, int row);
void QAbstractItemView_KeyboardSearch(void* ptr, char* search);
void* QAbstractItemView_Model(void* ptr);
void QAbstractItemView_OpenPersistentEditor(void* ptr, void* index);
void QAbstractItemView_ConnectPressed(void* ptr);
void QAbstractItemView_DisconnectPressed(void* ptr);
void QAbstractItemView_Reset(void* ptr);
void* QAbstractItemView_RootIndex(void* ptr);
void QAbstractItemView_ScrollTo(void* ptr, void* index, int hint);
void QAbstractItemView_ScrollToBottom(void* ptr);
void QAbstractItemView_ScrollToTop(void* ptr);
void QAbstractItemView_SelectAll(void* ptr);
void* QAbstractItemView_SelectionModel(void* ptr);
void QAbstractItemView_SetCurrentIndex(void* ptr, void* index);
void QAbstractItemView_SetIndexWidget(void* ptr, void* index, void* widget);
void QAbstractItemView_SetItemDelegate(void* ptr, void* delegate);
void QAbstractItemView_SetItemDelegateForColumn(void* ptr, int column, void* delegate);
void QAbstractItemView_SetItemDelegateForRow(void* ptr, int row, void* delegate);
void QAbstractItemView_SetModel(void* ptr, void* model);
void QAbstractItemView_SetRootIndex(void* ptr, void* index);
void QAbstractItemView_SetSelectionModel(void* ptr, void* selectionModel);
int QAbstractItemView_SizeHintForColumn(void* ptr, int column);
int QAbstractItemView_SizeHintForRow(void* ptr, int row);
void QAbstractItemView_Update(void* ptr, void* index);
void QAbstractItemView_ConnectViewportEntered(void* ptr);
void QAbstractItemView_DisconnectViewportEntered(void* ptr);
void QAbstractItemView_DestroyQAbstractItemView(void* ptr);

#ifdef __cplusplus
}
#endif