#ifdef __cplusplus
extern "C" {
#endif

int QListView_BatchSize(void* ptr);
int QListView_Flow(void* ptr);
int QListView_IsSelectionRectVisible(void* ptr);
int QListView_IsWrapping(void* ptr);
int QListView_LayoutMode(void* ptr);
int QListView_ModelColumn(void* ptr);
int QListView_Movement(void* ptr);
int QListView_ResizeMode(void* ptr);
void QListView_SetBatchSize(void* ptr, int batchSize);
void QListView_SetFlow(void* ptr, int flow);
void QListView_SetGridSize(void* ptr, void* size);
void QListView_SetLayoutMode(void* ptr, int mode);
void QListView_SetModelColumn(void* ptr, int column);
void QListView_SetMovement(void* ptr, int movement);
void QListView_SetResizeMode(void* ptr, int mode);
void QListView_SetSelectionRectVisible(void* ptr, int show);
void QListView_SetSpacing(void* ptr, int space);
void QListView_SetUniformItemSizes(void* ptr, int enable);
void QListView_SetViewMode(void* ptr, int mode);
void QListView_SetWordWrap(void* ptr, int on);
void QListView_SetWrapping(void* ptr, int enable);
int QListView_Spacing(void* ptr);
int QListView_UniformItemSizes(void* ptr);
int QListView_ViewMode(void* ptr);
int QListView_WordWrap(void* ptr);
void* QListView_NewQListView(void* parent);
void QListView_ClearPropertyFlags(void* ptr);
void* QListView_IndexAt(void* ptr, void* p);
int QListView_IsRowHidden(void* ptr, int row);
void QListView_ScrollTo(void* ptr, void* index, int hint);
void QListView_SetRowHidden(void* ptr, int row, int hide);
void QListView_DestroyQListView(void* ptr);

#ifdef __cplusplus
}
#endif