#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QUndoGroup_NewQUndoGroup(QtObjectPtr parent);
QtObjectPtr QUndoGroup_ActiveStack(QtObjectPtr ptr);
void QUndoGroup_ConnectActiveStackChanged(QtObjectPtr ptr);
void QUndoGroup_DisconnectActiveStackChanged(QtObjectPtr ptr);
void QUndoGroup_AddStack(QtObjectPtr ptr, QtObjectPtr stack);
int QUndoGroup_CanRedo(QtObjectPtr ptr);
void QUndoGroup_ConnectCanRedoChanged(QtObjectPtr ptr);
void QUndoGroup_DisconnectCanRedoChanged(QtObjectPtr ptr);
int QUndoGroup_CanUndo(QtObjectPtr ptr);
void QUndoGroup_ConnectCanUndoChanged(QtObjectPtr ptr);
void QUndoGroup_DisconnectCanUndoChanged(QtObjectPtr ptr);
void QUndoGroup_ConnectCleanChanged(QtObjectPtr ptr);
void QUndoGroup_DisconnectCleanChanged(QtObjectPtr ptr);
QtObjectPtr QUndoGroup_CreateRedoAction(QtObjectPtr ptr, QtObjectPtr parent, char* prefix);
QtObjectPtr QUndoGroup_CreateUndoAction(QtObjectPtr ptr, QtObjectPtr parent, char* prefix);
void QUndoGroup_ConnectIndexChanged(QtObjectPtr ptr);
void QUndoGroup_DisconnectIndexChanged(QtObjectPtr ptr);
int QUndoGroup_IsClean(QtObjectPtr ptr);
void QUndoGroup_Redo(QtObjectPtr ptr);
char* QUndoGroup_RedoText(QtObjectPtr ptr);
void QUndoGroup_ConnectRedoTextChanged(QtObjectPtr ptr);
void QUndoGroup_DisconnectRedoTextChanged(QtObjectPtr ptr);
void QUndoGroup_RemoveStack(QtObjectPtr ptr, QtObjectPtr stack);
void QUndoGroup_SetActiveStack(QtObjectPtr ptr, QtObjectPtr stack);
void QUndoGroup_Undo(QtObjectPtr ptr);
char* QUndoGroup_UndoText(QtObjectPtr ptr);
void QUndoGroup_ConnectUndoTextChanged(QtObjectPtr ptr);
void QUndoGroup_DisconnectUndoTextChanged(QtObjectPtr ptr);
void QUndoGroup_DestroyQUndoGroup(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif