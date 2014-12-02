#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Types
int QAbstractItemView_NoDragDrop();
int QAbstractItemView_DragOnly();
int QAbstractItemView_DropOnly();
int QAbstractItemView_DragDrop();
int QAbstractItemView_InternalMove();
int QAbstractItemView_NoEditTriggers();
int QAbstractItemView_CurrentChanged();
int QAbstractItemView_DoubleClicked();
int QAbstractItemView_SelectedClicked();
int QAbstractItemView_EditKeyPressed();
int QAbstractItemView_AnyKeyPressed();
int QAbstractItemView_AllEditTriggers();
int QAbstractItemView_EnsureVisible();
int QAbstractItemView_PositionAtTop();
int QAbstractItemView_PositionAtBottom();
int QAbstractItemView_PositionAtCenter();
int QAbstractItemView_ScrollPerItem();
int QAbstractItemView_ScrollPerPixel();
int QAbstractItemView_SelectItems();
int QAbstractItemView_SelectRows();
int QAbstractItemView_SelectColumns();
int QAbstractItemView_SingleSelection();
int QAbstractItemView_ContiguousSelection();
int QAbstractItemView_ExtendedSelection();
int QAbstractItemView_MultiSelection();
int QAbstractItemView_NoSelection();
//Public Functions
void QAbstractItemView_Destroy(QtObjectPtr ptr);
int QAbstractItemView_AlternatingRowColors(QtObjectPtr ptr);
int QAbstractItemView_AutoScrollMargin(QtObjectPtr ptr);
int QAbstractItemView_DefaultDropAction(QtObjectPtr ptr);
int QAbstractItemView_DragDropOverwriteMode(QtObjectPtr ptr);
int QAbstractItemView_DragEnabled(QtObjectPtr ptr);
int QAbstractItemView_HasAutoScroll(QtObjectPtr ptr);
void QAbstractItemView_KeyboardSearch_String(QtObjectPtr ptr, char* search);
int QAbstractItemView_SelectionMode(QtObjectPtr ptr);
void QAbstractItemView_SetAlternatingRowColors_Bool(QtObjectPtr ptr, int enable);
void QAbstractItemView_SetAutoScroll_Bool(QtObjectPtr ptr, int enable);
void QAbstractItemView_SetAutoScrollMargin_Int(QtObjectPtr ptr, int margin);
void QAbstractItemView_SetDefaultDropAction_DropAction(QtObjectPtr ptr, int dropAction);
void QAbstractItemView_SetDragDropOverwriteMode_Bool(QtObjectPtr ptr, int overwrite);
void QAbstractItemView_SetDragEnabled_Bool(QtObjectPtr ptr, int enable);
void QAbstractItemView_SetDropIndicatorShown_Bool(QtObjectPtr ptr, int enable);
void QAbstractItemView_SetSelectionMode_SelectionMode(QtObjectPtr ptr, int mode);
void QAbstractItemView_SetTabKeyNavigation_Bool(QtObjectPtr ptr, int enable);
void QAbstractItemView_SetTextElideMode_TextElideMode(QtObjectPtr ptr, int mode);
int QAbstractItemView_ShowDropIndicator(QtObjectPtr ptr);
int QAbstractItemView_SizeHintForColumn_Int(QtObjectPtr ptr, int column);
int QAbstractItemView_SizeHintForRow_Int(QtObjectPtr ptr, int row);
int QAbstractItemView_TabKeyNavigation(QtObjectPtr ptr);
int QAbstractItemView_TextElideMode(QtObjectPtr ptr);
//Public Slots
void QAbstractItemView_ConnectSlotClearSelection(QtObjectPtr ptr);
void QAbstractItemView_DisconnectSlotClearSelection(QtObjectPtr ptr);
void QAbstractItemView_ClearSelection(QtObjectPtr ptr);
void QAbstractItemView_ConnectSlotReset(QtObjectPtr ptr);
void QAbstractItemView_DisconnectSlotReset(QtObjectPtr ptr);
void QAbstractItemView_Reset(QtObjectPtr ptr);
void QAbstractItemView_ConnectSlotScrollToBottom(QtObjectPtr ptr);
void QAbstractItemView_DisconnectSlotScrollToBottom(QtObjectPtr ptr);
void QAbstractItemView_ScrollToBottom(QtObjectPtr ptr);
void QAbstractItemView_ConnectSlotScrollToTop(QtObjectPtr ptr);
void QAbstractItemView_DisconnectSlotScrollToTop(QtObjectPtr ptr);
void QAbstractItemView_ScrollToTop(QtObjectPtr ptr);
void QAbstractItemView_ConnectSlotSelectAll(QtObjectPtr ptr);
void QAbstractItemView_DisconnectSlotSelectAll(QtObjectPtr ptr);
void QAbstractItemView_SelectAll(QtObjectPtr ptr);
//Signals
void QAbstractItemView_ConnectSignalActivated(QtObjectPtr ptr);
void QAbstractItemView_DisconnectSignalActivated(QtObjectPtr ptr);
void QAbstractItemView_ConnectSignalClicked(QtObjectPtr ptr);
void QAbstractItemView_DisconnectSignalClicked(QtObjectPtr ptr);
void QAbstractItemView_ConnectSignalDoubleClicked(QtObjectPtr ptr);
void QAbstractItemView_DisconnectSignalDoubleClicked(QtObjectPtr ptr);
void QAbstractItemView_ConnectSignalEntered(QtObjectPtr ptr);
void QAbstractItemView_DisconnectSignalEntered(QtObjectPtr ptr);
void QAbstractItemView_ConnectSignalPressed(QtObjectPtr ptr);
void QAbstractItemView_DisconnectSignalPressed(QtObjectPtr ptr);
void QAbstractItemView_ConnectSignalViewportEntered(QtObjectPtr ptr);
void QAbstractItemView_DisconnectSignalViewportEntered(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
