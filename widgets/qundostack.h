#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QUndoStack_IsActive(QtObjectPtr ptr);
void QUndoStack_SetActive(QtObjectPtr ptr, int active);
void QUndoStack_SetUndoLimit(QtObjectPtr ptr, int limit);
int QUndoStack_UndoLimit(QtObjectPtr ptr);
QtObjectPtr QUndoStack_NewQUndoStack(QtObjectPtr parent);
void QUndoStack_BeginMacro(QtObjectPtr ptr, char* text);
int QUndoStack_CanRedo(QtObjectPtr ptr);
void QUndoStack_ConnectCanRedoChanged(QtObjectPtr ptr);
void QUndoStack_DisconnectCanRedoChanged(QtObjectPtr ptr);
int QUndoStack_CanUndo(QtObjectPtr ptr);
void QUndoStack_ConnectCanUndoChanged(QtObjectPtr ptr);
void QUndoStack_DisconnectCanUndoChanged(QtObjectPtr ptr);
void QUndoStack_ConnectCleanChanged(QtObjectPtr ptr);
void QUndoStack_DisconnectCleanChanged(QtObjectPtr ptr);
int QUndoStack_CleanIndex(QtObjectPtr ptr);
void QUndoStack_Clear(QtObjectPtr ptr);
QtObjectPtr QUndoStack_Command(QtObjectPtr ptr, int index);
int QUndoStack_Count(QtObjectPtr ptr);
QtObjectPtr QUndoStack_CreateRedoAction(QtObjectPtr ptr, QtObjectPtr parent, char* prefix);
QtObjectPtr QUndoStack_CreateUndoAction(QtObjectPtr ptr, QtObjectPtr parent, char* prefix);
void QUndoStack_EndMacro(QtObjectPtr ptr);
int QUndoStack_Index(QtObjectPtr ptr);
void QUndoStack_ConnectIndexChanged(QtObjectPtr ptr);
void QUndoStack_DisconnectIndexChanged(QtObjectPtr ptr);
int QUndoStack_IsClean(QtObjectPtr ptr);
void QUndoStack_Push(QtObjectPtr ptr, QtObjectPtr cmd);
void QUndoStack_Redo(QtObjectPtr ptr);
char* QUndoStack_RedoText(QtObjectPtr ptr);
void QUndoStack_ConnectRedoTextChanged(QtObjectPtr ptr);
void QUndoStack_DisconnectRedoTextChanged(QtObjectPtr ptr);
void QUndoStack_SetClean(QtObjectPtr ptr);
void QUndoStack_SetIndex(QtObjectPtr ptr, int idx);
char* QUndoStack_Text(QtObjectPtr ptr, int idx);
void QUndoStack_Undo(QtObjectPtr ptr);
char* QUndoStack_UndoText(QtObjectPtr ptr);
void QUndoStack_ConnectUndoTextChanged(QtObjectPtr ptr);
void QUndoStack_DisconnectUndoTextChanged(QtObjectPtr ptr);
void QUndoStack_DestroyQUndoStack(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif