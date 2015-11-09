#ifdef __cplusplus
extern "C" {
#endif

int QTreeView_AllColumnsShowFocus(void* ptr);
int QTreeView_AutoExpandDelay(void* ptr);
void QTreeView_Collapse(void* ptr, void* index);
void QTreeView_Expand(void* ptr, void* index);
int QTreeView_ExpandsOnDoubleClick(void* ptr);
int QTreeView_Indentation(void* ptr);
int QTreeView_IsAnimated(void* ptr);
int QTreeView_IsExpanded(void* ptr, void* index);
int QTreeView_IsHeaderHidden(void* ptr);
int QTreeView_IsSortingEnabled(void* ptr);
int QTreeView_ItemsExpandable(void* ptr);
void QTreeView_ResetIndentation(void* ptr);
int QTreeView_RootIsDecorated(void* ptr);
void QTreeView_SetAllColumnsShowFocus(void* ptr, int enable);
void QTreeView_SetAnimated(void* ptr, int enable);
void QTreeView_SetAutoExpandDelay(void* ptr, int delay);
void QTreeView_SetExpandsOnDoubleClick(void* ptr, int enable);
void QTreeView_SetHeaderHidden(void* ptr, int hide);
void QTreeView_SetIndentation(void* ptr, int i);
void QTreeView_SetItemsExpandable(void* ptr, int enable);
void QTreeView_SetRootIsDecorated(void* ptr, int show);
void QTreeView_SetSortingEnabled(void* ptr, int enable);
void QTreeView_SetUniformRowHeights(void* ptr, int uniform);
void QTreeView_SetWordWrap(void* ptr, int on);
int QTreeView_UniformRowHeights(void* ptr);
int QTreeView_WordWrap(void* ptr);
void* QTreeView_NewQTreeView(void* parent);
void QTreeView_CollapseAll(void* ptr);
void QTreeView_ConnectCollapsed(void* ptr);
void QTreeView_DisconnectCollapsed(void* ptr);
int QTreeView_ColumnAt(void* ptr, int x);
int QTreeView_ColumnViewportPosition(void* ptr, int column);
int QTreeView_ColumnWidth(void* ptr, int column);
void QTreeView_ExpandAll(void* ptr);
void QTreeView_ExpandToDepth(void* ptr, int depth);
void QTreeView_ConnectExpanded(void* ptr);
void QTreeView_DisconnectExpanded(void* ptr);
void* QTreeView_Header(void* ptr);
void QTreeView_HideColumn(void* ptr, int column);
void* QTreeView_IndexAbove(void* ptr, void* index);
void* QTreeView_IndexAt(void* ptr, void* point);
void* QTreeView_IndexBelow(void* ptr, void* index);
int QTreeView_IsColumnHidden(void* ptr, int column);
int QTreeView_IsFirstColumnSpanned(void* ptr, int row, void* parent);
int QTreeView_IsRowHidden(void* ptr, int row, void* parent);
void QTreeView_KeyboardSearch(void* ptr, char* search);
void QTreeView_Reset(void* ptr);
void QTreeView_ResizeColumnToContents(void* ptr, int column);
void QTreeView_ScrollTo(void* ptr, void* index, int hint);
void QTreeView_SelectAll(void* ptr);
void QTreeView_SetColumnHidden(void* ptr, int column, int hide);
void QTreeView_SetColumnWidth(void* ptr, int column, int width);
void QTreeView_SetExpanded(void* ptr, void* index, int expanded);
void QTreeView_SetFirstColumnSpanned(void* ptr, int row, void* parent, int span);
void QTreeView_SetHeader(void* ptr, void* header);
void QTreeView_SetModel(void* ptr, void* model);
void QTreeView_SetRootIndex(void* ptr, void* index);
void QTreeView_SetRowHidden(void* ptr, int row, void* parent, int hide);
void QTreeView_SetSelectionModel(void* ptr, void* selectionModel);
void QTreeView_SetTreePosition(void* ptr, int index);
void QTreeView_ShowColumn(void* ptr, int column);
void QTreeView_SortByColumn(void* ptr, int column, int order);
int QTreeView_TreePosition(void* ptr);
void QTreeView_DestroyQTreeView(void* ptr);

#ifdef __cplusplus
}
#endif