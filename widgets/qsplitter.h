#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QSplitter_ChildrenCollapsible(QtObjectPtr ptr);
int QSplitter_HandleWidth(QtObjectPtr ptr);
int QSplitter_IndexOf(QtObjectPtr ptr, QtObjectPtr widget);
int QSplitter_OpaqueResize(QtObjectPtr ptr);
int QSplitter_Orientation(QtObjectPtr ptr);
void QSplitter_SetChildrenCollapsible(QtObjectPtr ptr, int v);
void QSplitter_SetHandleWidth(QtObjectPtr ptr, int v);
void QSplitter_SetOpaqueResize(QtObjectPtr ptr, int opaque);
void QSplitter_SetOrientation(QtObjectPtr ptr, int v);
QtObjectPtr QSplitter_NewQSplitter(QtObjectPtr parent);
QtObjectPtr QSplitter_NewQSplitter2(int orientation, QtObjectPtr parent);
void QSplitter_AddWidget(QtObjectPtr ptr, QtObjectPtr widget);
int QSplitter_Count(QtObjectPtr ptr);
void QSplitter_GetRange(QtObjectPtr ptr, int index, int min, int max);
QtObjectPtr QSplitter_Handle(QtObjectPtr ptr, int index);
void QSplitter_InsertWidget(QtObjectPtr ptr, int index, QtObjectPtr widget);
int QSplitter_IsCollapsible(QtObjectPtr ptr, int index);
void QSplitter_Refresh(QtObjectPtr ptr);
int QSplitter_RestoreState(QtObjectPtr ptr, QtObjectPtr state);
void QSplitter_SetCollapsible(QtObjectPtr ptr, int index, int collapse);
void QSplitter_SetStretchFactor(QtObjectPtr ptr, int index, int stretch);
void QSplitter_ConnectSplitterMoved(QtObjectPtr ptr);
void QSplitter_DisconnectSplitterMoved(QtObjectPtr ptr);
QtObjectPtr QSplitter_Widget(QtObjectPtr ptr, int index);
void QSplitter_DestroyQSplitter(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif