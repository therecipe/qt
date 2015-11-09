#ifdef __cplusplus
extern "C" {
#endif

int QSplitter_ChildrenCollapsible(void* ptr);
int QSplitter_HandleWidth(void* ptr);
int QSplitter_IndexOf(void* ptr, void* widget);
int QSplitter_OpaqueResize(void* ptr);
int QSplitter_Orientation(void* ptr);
void QSplitter_SetChildrenCollapsible(void* ptr, int v);
void QSplitter_SetHandleWidth(void* ptr, int v);
void QSplitter_SetOpaqueResize(void* ptr, int opaque);
void QSplitter_SetOrientation(void* ptr, int v);
void* QSplitter_NewQSplitter(void* parent);
void* QSplitter_NewQSplitter2(int orientation, void* parent);
void QSplitter_AddWidget(void* ptr, void* widget);
int QSplitter_Count(void* ptr);
void QSplitter_GetRange(void* ptr, int index, int min, int max);
void* QSplitter_Handle(void* ptr, int index);
void QSplitter_InsertWidget(void* ptr, int index, void* widget);
int QSplitter_IsCollapsible(void* ptr, int index);
void QSplitter_Refresh(void* ptr);
int QSplitter_RestoreState(void* ptr, void* state);
void* QSplitter_SaveState(void* ptr);
void QSplitter_SetCollapsible(void* ptr, int index, int collapse);
void QSplitter_SetStretchFactor(void* ptr, int index, int stretch);
void QSplitter_ConnectSplitterMoved(void* ptr);
void QSplitter_DisconnectSplitterMoved(void* ptr);
void* QSplitter_Widget(void* ptr, int index);
void QSplitter_DestroyQSplitter(void* ptr);

#ifdef __cplusplus
}
#endif