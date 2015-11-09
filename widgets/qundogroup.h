#ifdef __cplusplus
extern "C" {
#endif

void* QUndoGroup_NewQUndoGroup(void* parent);
void* QUndoGroup_ActiveStack(void* ptr);
void QUndoGroup_ConnectActiveStackChanged(void* ptr);
void QUndoGroup_DisconnectActiveStackChanged(void* ptr);
void QUndoGroup_AddStack(void* ptr, void* stack);
int QUndoGroup_CanRedo(void* ptr);
void QUndoGroup_ConnectCanRedoChanged(void* ptr);
void QUndoGroup_DisconnectCanRedoChanged(void* ptr);
int QUndoGroup_CanUndo(void* ptr);
void QUndoGroup_ConnectCanUndoChanged(void* ptr);
void QUndoGroup_DisconnectCanUndoChanged(void* ptr);
void QUndoGroup_ConnectCleanChanged(void* ptr);
void QUndoGroup_DisconnectCleanChanged(void* ptr);
void* QUndoGroup_CreateRedoAction(void* ptr, void* parent, char* prefix);
void* QUndoGroup_CreateUndoAction(void* ptr, void* parent, char* prefix);
void QUndoGroup_ConnectIndexChanged(void* ptr);
void QUndoGroup_DisconnectIndexChanged(void* ptr);
int QUndoGroup_IsClean(void* ptr);
void QUndoGroup_Redo(void* ptr);
char* QUndoGroup_RedoText(void* ptr);
void QUndoGroup_ConnectRedoTextChanged(void* ptr);
void QUndoGroup_DisconnectRedoTextChanged(void* ptr);
void QUndoGroup_RemoveStack(void* ptr, void* stack);
void QUndoGroup_SetActiveStack(void* ptr, void* stack);
void QUndoGroup_Undo(void* ptr);
char* QUndoGroup_UndoText(void* ptr);
void QUndoGroup_ConnectUndoTextChanged(void* ptr);
void QUndoGroup_DisconnectUndoTextChanged(void* ptr);
void QUndoGroup_DestroyQUndoGroup(void* ptr);

#ifdef __cplusplus
}
#endif