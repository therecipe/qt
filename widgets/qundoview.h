#ifdef __cplusplus
extern "C" {
#endif

char* QUndoView_EmptyLabel(void* ptr);
void QUndoView_SetCleanIcon(void* ptr, void* icon);
void QUndoView_SetEmptyLabel(void* ptr, char* label);
void* QUndoView_NewQUndoView3(void* group, void* parent);
void* QUndoView_NewQUndoView2(void* stack, void* parent);
void* QUndoView_NewQUndoView(void* parent);
void* QUndoView_Group(void* ptr);
void QUndoView_SetGroup(void* ptr, void* group);
void QUndoView_SetStack(void* ptr, void* stack);
void* QUndoView_Stack(void* ptr);
void QUndoView_DestroyQUndoView(void* ptr);

#ifdef __cplusplus
}
#endif