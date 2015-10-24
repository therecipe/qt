#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QTreeView_AllColumnsShowFocus(QtObjectPtr ptr);
int QTreeView_AutoExpandDelay(QtObjectPtr ptr);
void QTreeView_Collapse(QtObjectPtr ptr, QtObjectPtr index);
void QTreeView_Expand(QtObjectPtr ptr, QtObjectPtr index);
int QTreeView_ExpandsOnDoubleClick(QtObjectPtr ptr);
int QTreeView_Indentation(QtObjectPtr ptr);
int QTreeView_IsAnimated(QtObjectPtr ptr);
int QTreeView_IsExpanded(QtObjectPtr ptr, QtObjectPtr index);
int QTreeView_IsHeaderHidden(QtObjectPtr ptr);
int QTreeView_IsSortingEnabled(QtObjectPtr ptr);
int QTreeView_ItemsExpandable(QtObjectPtr ptr);
void QTreeView_ResetIndentation(QtObjectPtr ptr);
int QTreeView_RootIsDecorated(QtObjectPtr ptr);
void QTreeView_SetAllColumnsShowFocus(QtObjectPtr ptr, int enable);
void QTreeView_SetAnimated(QtObjectPtr ptr, int enable);
void QTreeView_SetAutoExpandDelay(QtObjectPtr ptr, int delay);
void QTreeView_SetExpandsOnDoubleClick(QtObjectPtr ptr, int enable);
void QTreeView_SetHeaderHidden(QtObjectPtr ptr, int hide);
void QTreeView_SetIndentation(QtObjectPtr ptr, int i);
void QTreeView_SetItemsExpandable(QtObjectPtr ptr, int enable);
void QTreeView_SetRootIsDecorated(QtObjectPtr ptr, int show);
void QTreeView_SetSortingEnabled(QtObjectPtr ptr, int enable);
void QTreeView_SetUniformRowHeights(QtObjectPtr ptr, int uniform);
void QTreeView_SetWordWrap(QtObjectPtr ptr, int on);
int QTreeView_UniformRowHeights(QtObjectPtr ptr);
int QTreeView_WordWrap(QtObjectPtr ptr);
QtObjectPtr QTreeView_NewQTreeView(QtObjectPtr parent);
void QTreeView_CollapseAll(QtObjectPtr ptr);
void QTreeView_ConnectCollapsed(QtObjectPtr ptr);
void QTreeView_DisconnectCollapsed(QtObjectPtr ptr);
int QTreeView_ColumnAt(QtObjectPtr ptr, int x);
int QTreeView_ColumnViewportPosition(QtObjectPtr ptr, int column);
int QTreeView_ColumnWidth(QtObjectPtr ptr, int column);
void QTreeView_ExpandAll(QtObjectPtr ptr);
void QTreeView_ExpandToDepth(QtObjectPtr ptr, int depth);
void QTreeView_ConnectExpanded(QtObjectPtr ptr);
void QTreeView_DisconnectExpanded(QtObjectPtr ptr);
QtObjectPtr QTreeView_Header(QtObjectPtr ptr);
void QTreeView_HideColumn(QtObjectPtr ptr, int column);
QtObjectPtr QTreeView_IndexAbove(QtObjectPtr ptr, QtObjectPtr index);
QtObjectPtr QTreeView_IndexAt(QtObjectPtr ptr, QtObjectPtr point);
QtObjectPtr QTreeView_IndexBelow(QtObjectPtr ptr, QtObjectPtr index);
int QTreeView_IsColumnHidden(QtObjectPtr ptr, int column);
int QTreeView_IsFirstColumnSpanned(QtObjectPtr ptr, int row, QtObjectPtr parent);
int QTreeView_IsRowHidden(QtObjectPtr ptr, int row, QtObjectPtr parent);
void QTreeView_KeyboardSearch(QtObjectPtr ptr, char* search);
void QTreeView_Reset(QtObjectPtr ptr);
void QTreeView_ResizeColumnToContents(QtObjectPtr ptr, int column);
void QTreeView_ScrollTo(QtObjectPtr ptr, QtObjectPtr index, int hint);
void QTreeView_SelectAll(QtObjectPtr ptr);
void QTreeView_SetColumnHidden(QtObjectPtr ptr, int column, int hide);
void QTreeView_SetColumnWidth(QtObjectPtr ptr, int column, int width);
void QTreeView_SetExpanded(QtObjectPtr ptr, QtObjectPtr index, int expanded);
void QTreeView_SetFirstColumnSpanned(QtObjectPtr ptr, int row, QtObjectPtr parent, int span);
void QTreeView_SetHeader(QtObjectPtr ptr, QtObjectPtr header);
void QTreeView_SetModel(QtObjectPtr ptr, QtObjectPtr model);
void QTreeView_SetRootIndex(QtObjectPtr ptr, QtObjectPtr index);
void QTreeView_SetRowHidden(QtObjectPtr ptr, int row, QtObjectPtr parent, int hide);
void QTreeView_SetSelectionModel(QtObjectPtr ptr, QtObjectPtr selectionModel);
void QTreeView_SetTreePosition(QtObjectPtr ptr, int index);
void QTreeView_ShowColumn(QtObjectPtr ptr, int column);
void QTreeView_SortByColumn(QtObjectPtr ptr, int column, int order);
int QTreeView_TreePosition(QtObjectPtr ptr);
void QTreeView_DestroyQTreeView(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif