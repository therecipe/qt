#ifdef __cplusplus
extern "C" {
#endif

void* QUndoCommand_NewQUndoCommand(void* parent);
void* QUndoCommand_NewQUndoCommand2(char* text, void* parent);
char* QUndoCommand_ActionText(void* ptr);
void* QUndoCommand_Child(void* ptr, int index);
int QUndoCommand_ChildCount(void* ptr);
int QUndoCommand_Id(void* ptr);
int QUndoCommand_MergeWith(void* ptr, void* command);
void QUndoCommand_Redo(void* ptr);
void QUndoCommand_SetText(void* ptr, char* text);
char* QUndoCommand_Text(void* ptr);
void QUndoCommand_Undo(void* ptr);
void QUndoCommand_DestroyQUndoCommand(void* ptr);

#ifdef __cplusplus
}
#endif