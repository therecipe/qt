#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QListView_BatchSize(QtObjectPtr ptr);
int QListView_Flow(QtObjectPtr ptr);
int QListView_IsSelectionRectVisible(QtObjectPtr ptr);
int QListView_IsWrapping(QtObjectPtr ptr);
int QListView_LayoutMode(QtObjectPtr ptr);
int QListView_ModelColumn(QtObjectPtr ptr);
int QListView_Movement(QtObjectPtr ptr);
int QListView_ResizeMode(QtObjectPtr ptr);
void QListView_SetBatchSize(QtObjectPtr ptr, int batchSize);
void QListView_SetFlow(QtObjectPtr ptr, int flow);
void QListView_SetGridSize(QtObjectPtr ptr, QtObjectPtr size);
void QListView_SetLayoutMode(QtObjectPtr ptr, int mode);
void QListView_SetModelColumn(QtObjectPtr ptr, int column);
void QListView_SetMovement(QtObjectPtr ptr, int movement);
void QListView_SetResizeMode(QtObjectPtr ptr, int mode);
void QListView_SetSelectionRectVisible(QtObjectPtr ptr, int show);
void QListView_SetSpacing(QtObjectPtr ptr, int space);
void QListView_SetUniformItemSizes(QtObjectPtr ptr, int enable);
void QListView_SetViewMode(QtObjectPtr ptr, int mode);
void QListView_SetWordWrap(QtObjectPtr ptr, int on);
void QListView_SetWrapping(QtObjectPtr ptr, int enable);
int QListView_Spacing(QtObjectPtr ptr);
int QListView_UniformItemSizes(QtObjectPtr ptr);
int QListView_ViewMode(QtObjectPtr ptr);
int QListView_WordWrap(QtObjectPtr ptr);
QtObjectPtr QListView_NewQListView(QtObjectPtr parent);
void QListView_ClearPropertyFlags(QtObjectPtr ptr);
QtObjectPtr QListView_IndexAt(QtObjectPtr ptr, QtObjectPtr p);
int QListView_IsRowHidden(QtObjectPtr ptr, int row);
void QListView_ScrollTo(QtObjectPtr ptr, QtObjectPtr index, int hint);
void QListView_SetRowHidden(QtObjectPtr ptr, int row, int hide);
void QListView_DestroyQListView(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif