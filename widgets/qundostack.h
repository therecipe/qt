#ifdef __cplusplus
extern "C" {
#endif

int QUndoStack_IsActive(void* ptr);
void QUndoStack_SetActive(void* ptr, int active);
void QUndoStack_SetUndoLimit(void* ptr, int limit);
int QUndoStack_UndoLimit(void* ptr);
void* QUndoStack_NewQUndoStack(void* parent);
void QUndoStack_BeginMacro(void* ptr, char* text);
int QUndoStack_CanRedo(void* ptr);
void QUndoStack_ConnectCanRedoChanged(void* ptr);
void QUndoStack_DisconnectCanRedoChanged(void* ptr);
int QUndoStack_CanUndo(void* ptr);
void QUndoStack_ConnectCanUndoChanged(void* ptr);
void QUndoStack_DisconnectCanUndoChanged(void* ptr);
void QUndoStack_ConnectCleanChanged(void* ptr);
void QUndoStack_DisconnectCleanChanged(void* ptr);
int QUndoStack_CleanIndex(void* ptr);
void QUndoStack_Clear(void* ptr);
void* QUndoStack_Command(void* ptr, int index);
int QUndoStack_Count(void* ptr);
void* QUndoStack_CreateRedoAction(void* ptr, void* parent, char* prefix);
void* QUndoStack_CreateUndoAction(void* ptr, void* parent, char* prefix);
void QUndoStack_EndMacro(void* ptr);
int QUndoStack_Index(void* ptr);
void QUndoStack_ConnectIndexChanged(void* ptr);
void QUndoStack_DisconnectIndexChanged(void* ptr);
int QUndoStack_IsClean(void* ptr);
void QUndoStack_Push(void* ptr, void* cmd);
void QUndoStack_Redo(void* ptr);
char* QUndoStack_RedoText(void* ptr);
void QUndoStack_ConnectRedoTextChanged(void* ptr);
void QUndoStack_DisconnectRedoTextChanged(void* ptr);
void QUndoStack_SetClean(void* ptr);
void QUndoStack_SetIndex(void* ptr, int idx);
char* QUndoStack_Text(void* ptr, int idx);
void QUndoStack_Undo(void* ptr);
char* QUndoStack_UndoText(void* ptr);
void QUndoStack_ConnectUndoTextChanged(void* ptr);
void QUndoStack_DisconnectUndoTextChanged(void* ptr);
void QUndoStack_DestroyQUndoStack(void* ptr);

#ifdef __cplusplus
}
#endif