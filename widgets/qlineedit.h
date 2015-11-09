#ifdef __cplusplus
extern "C" {
#endif

int QLineEdit_Alignment(void* ptr);
int QLineEdit_CursorMoveStyle(void* ptr);
int QLineEdit_CursorPosition(void* ptr);
char* QLineEdit_DisplayText(void* ptr);
int QLineEdit_DragEnabled(void* ptr);
int QLineEdit_EchoMode(void* ptr);
int QLineEdit_HasAcceptableInput(void* ptr);
int QLineEdit_HasFrame(void* ptr);
int QLineEdit_HasSelectedText(void* ptr);
char* QLineEdit_InputMask(void* ptr);
int QLineEdit_IsClearButtonEnabled(void* ptr);
int QLineEdit_IsModified(void* ptr);
int QLineEdit_IsReadOnly(void* ptr);
int QLineEdit_IsRedoAvailable(void* ptr);
int QLineEdit_IsUndoAvailable(void* ptr);
int QLineEdit_MaxLength(void* ptr);
char* QLineEdit_PlaceholderText(void* ptr);
char* QLineEdit_SelectedText(void* ptr);
void QLineEdit_SetAlignment(void* ptr, int flag);
void QLineEdit_SetClearButtonEnabled(void* ptr, int enable);
void QLineEdit_SetCursorMoveStyle(void* ptr, int style);
void QLineEdit_SetCursorPosition(void* ptr, int v);
void QLineEdit_SetDragEnabled(void* ptr, int b);
void QLineEdit_SetEchoMode(void* ptr, int v);
void QLineEdit_SetFrame(void* ptr, int v);
void QLineEdit_SetInputMask(void* ptr, char* inputMask);
void QLineEdit_SetMaxLength(void* ptr, int v);
void QLineEdit_SetModified(void* ptr, int v);
void QLineEdit_SetPlaceholderText(void* ptr, char* v);
void QLineEdit_SetReadOnly(void* ptr, int v);
void QLineEdit_SetText(void* ptr, char* v);
char* QLineEdit_Text(void* ptr);
void* QLineEdit_NewQLineEdit(void* parent);
void* QLineEdit_NewQLineEdit2(char* contents, void* parent);
void* QLineEdit_AddAction2(void* ptr, void* icon, int position);
void QLineEdit_AddAction(void* ptr, void* action, int position);
void QLineEdit_Backspace(void* ptr);
void QLineEdit_Clear(void* ptr);
void* QLineEdit_Completer(void* ptr);
void QLineEdit_Copy(void* ptr);
void* QLineEdit_CreateStandardContextMenu(void* ptr);
void QLineEdit_CursorBackward(void* ptr, int mark, int steps);
void QLineEdit_CursorForward(void* ptr, int mark, int steps);
int QLineEdit_CursorPositionAt(void* ptr, void* pos);
void QLineEdit_ConnectCursorPositionChanged(void* ptr);
void QLineEdit_DisconnectCursorPositionChanged(void* ptr);
void QLineEdit_CursorWordBackward(void* ptr, int mark);
void QLineEdit_CursorWordForward(void* ptr, int mark);
void QLineEdit_Cut(void* ptr);
void QLineEdit_Del(void* ptr);
void QLineEdit_Deselect(void* ptr);
void QLineEdit_ConnectEditingFinished(void* ptr);
void QLineEdit_DisconnectEditingFinished(void* ptr);
void QLineEdit_End(void* ptr, int mark);
int QLineEdit_Event(void* ptr, void* e);
void QLineEdit_GetTextMargins(void* ptr, int left, int top, int right, int bottom);
void QLineEdit_Home(void* ptr, int mark);
void* QLineEdit_InputMethodQuery(void* ptr, int property);
void QLineEdit_Paste(void* ptr);
void QLineEdit_Redo(void* ptr);
void QLineEdit_ConnectReturnPressed(void* ptr);
void QLineEdit_DisconnectReturnPressed(void* ptr);
void QLineEdit_SelectAll(void* ptr);
void QLineEdit_ConnectSelectionChanged(void* ptr);
void QLineEdit_DisconnectSelectionChanged(void* ptr);
int QLineEdit_SelectionStart(void* ptr);
void QLineEdit_SetCompleter(void* ptr, void* c);
void QLineEdit_SetSelection(void* ptr, int start, int length);
void QLineEdit_SetTextMargins2(void* ptr, void* margins);
void QLineEdit_SetTextMargins(void* ptr, int left, int top, int right, int bottom);
void QLineEdit_SetValidator(void* ptr, void* v);
void QLineEdit_ConnectTextChanged(void* ptr);
void QLineEdit_DisconnectTextChanged(void* ptr);
void QLineEdit_ConnectTextEdited(void* ptr);
void QLineEdit_DisconnectTextEdited(void* ptr);
void QLineEdit_Undo(void* ptr);
void* QLineEdit_Validator(void* ptr);
void QLineEdit_DestroyQLineEdit(void* ptr);

#ifdef __cplusplus
}
#endif