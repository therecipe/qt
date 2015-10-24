#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QUndoCommand_NewQUndoCommand(QtObjectPtr parent);
QtObjectPtr QUndoCommand_NewQUndoCommand2(char* text, QtObjectPtr parent);
char* QUndoCommand_ActionText(QtObjectPtr ptr);
QtObjectPtr QUndoCommand_Child(QtObjectPtr ptr, int index);
int QUndoCommand_ChildCount(QtObjectPtr ptr);
int QUndoCommand_Id(QtObjectPtr ptr);
int QUndoCommand_MergeWith(QtObjectPtr ptr, QtObjectPtr command);
void QUndoCommand_Redo(QtObjectPtr ptr);
void QUndoCommand_SetText(QtObjectPtr ptr, char* text);
char* QUndoCommand_Text(QtObjectPtr ptr);
void QUndoCommand_Undo(QtObjectPtr ptr);
void QUndoCommand_DestroyQUndoCommand(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif