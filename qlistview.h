#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QListView_New_QWidget(QtObjectPtr parent);
void QListView_Destroy(QtObjectPtr ptr);
int QListView_BatchSize(QtObjectPtr ptr);
void QListView_ClearPropertyFlags(QtObjectPtr ptr);
int QListView_IsRowHidden_Int(QtObjectPtr ptr, int row);
int QListView_IsSelectionRectVisible(QtObjectPtr ptr);
int QListView_IsWrapping(QtObjectPtr ptr);
int QListView_ModelColumn(QtObjectPtr ptr);
void QListView_SetBatchSize_Int(QtObjectPtr ptr, int batchSize);
void QListView_SetModelColumn_Int(QtObjectPtr ptr, int column);
void QListView_SetRowHidden_Int_Bool(QtObjectPtr ptr, int row, int hide);
void QListView_SetSelectionRectVisible_Bool(QtObjectPtr ptr, int show);
void QListView_SetSpacing_Int(QtObjectPtr ptr, int space);
void QListView_SetUniformItemSizes_Bool(QtObjectPtr ptr, int enable);
void QListView_SetWordWrap_Bool(QtObjectPtr ptr, int on);
void QListView_SetWrapping_Bool(QtObjectPtr ptr, int enable);
int QListView_Spacing(QtObjectPtr ptr);
int QListView_UniformItemSizes(QtObjectPtr ptr);
int QListView_WordWrap(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
