#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QTreeView_New_QWidget(QtObjectPtr parent);
void QTreeView_Destroy(QtObjectPtr ptr);
int QTreeView_AllColumnsShowFocus(QtObjectPtr ptr);
int QTreeView_AutoExpandDelay(QtObjectPtr ptr);
int QTreeView_ColumnAt_Int(QtObjectPtr ptr, int x);
int QTreeView_ColumnViewportPosition_Int(QtObjectPtr ptr, int column);
int QTreeView_ColumnWidth_Int(QtObjectPtr ptr, int column);
int QTreeView_ExpandsOnDoubleClick(QtObjectPtr ptr);
int QTreeView_Indentation(QtObjectPtr ptr);
int QTreeView_IsAnimated(QtObjectPtr ptr);
int QTreeView_IsColumnHidden_Int(QtObjectPtr ptr, int column);
int QTreeView_IsHeaderHidden(QtObjectPtr ptr);
int QTreeView_IsSortingEnabled(QtObjectPtr ptr);
int QTreeView_ItemsExpandable(QtObjectPtr ptr);
int QTreeView_RootIsDecorated(QtObjectPtr ptr);
void QTreeView_SetAllColumnsShowFocus_Bool(QtObjectPtr ptr, int enable);
void QTreeView_SetAnimated_Bool(QtObjectPtr ptr, int enable);
void QTreeView_SetAutoExpandDelay_Int(QtObjectPtr ptr, int delay);
void QTreeView_SetColumnHidden_Int_Bool(QtObjectPtr ptr, int column, int hide);
void QTreeView_SetColumnWidth_Int_Int(QtObjectPtr ptr, int column, int width);
void QTreeView_SetExpandsOnDoubleClick_Bool(QtObjectPtr ptr, int enable);
void QTreeView_SetHeaderHidden_Bool(QtObjectPtr ptr, int hide);
void QTreeView_SetIndentation_Int(QtObjectPtr ptr, int i);
void QTreeView_SetItemsExpandable_Bool(QtObjectPtr ptr, int enable);
void QTreeView_SetRootIsDecorated_Bool(QtObjectPtr ptr, int show);
void QTreeView_SetSortingEnabled_Bool(QtObjectPtr ptr, int enable);
void QTreeView_SetTreePosition_Int(QtObjectPtr ptr, int index);
void QTreeView_SetUniformRowHeights_Bool(QtObjectPtr ptr, int uniform);
void QTreeView_SetWordWrap_Bool(QtObjectPtr ptr, int on);
void QTreeView_SortByColumn_Int_SortOrder(QtObjectPtr ptr, int column, int order);
int QTreeView_TreePosition(QtObjectPtr ptr);
int QTreeView_UniformRowHeights(QtObjectPtr ptr);
int QTreeView_WordWrap(QtObjectPtr ptr);
//Public Slots
void QTreeView_ConnectSlotCollapseAll(QtObjectPtr ptr);
void QTreeView_DisconnectSlotCollapseAll(QtObjectPtr ptr);
void QTreeView_CollapseAll(QtObjectPtr ptr);
void QTreeView_ConnectSlotExpandAll(QtObjectPtr ptr);
void QTreeView_DisconnectSlotExpandAll(QtObjectPtr ptr);
void QTreeView_ExpandAll(QtObjectPtr ptr);
void QTreeView_ConnectSlotExpandToDepth(QtObjectPtr ptr);
void QTreeView_DisconnectSlotExpandToDepth(QtObjectPtr ptr);
void QTreeView_ExpandToDepth_Int(QtObjectPtr ptr, int depth);
void QTreeView_ConnectSlotHideColumn(QtObjectPtr ptr);
void QTreeView_DisconnectSlotHideColumn(QtObjectPtr ptr);
void QTreeView_HideColumn_Int(QtObjectPtr ptr, int column);
void QTreeView_ConnectSlotResizeColumnToContents(QtObjectPtr ptr);
void QTreeView_DisconnectSlotResizeColumnToContents(QtObjectPtr ptr);
void QTreeView_ResizeColumnToContents_Int(QtObjectPtr ptr, int column);
void QTreeView_ConnectSlotShowColumn(QtObjectPtr ptr);
void QTreeView_DisconnectSlotShowColumn(QtObjectPtr ptr);
void QTreeView_ShowColumn_Int(QtObjectPtr ptr, int column);

#ifdef __cplusplus
}
#endif
