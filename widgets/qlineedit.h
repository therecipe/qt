#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QLineEdit_Alignment(QtObjectPtr ptr);
int QLineEdit_CursorMoveStyle(QtObjectPtr ptr);
int QLineEdit_CursorPosition(QtObjectPtr ptr);
char* QLineEdit_DisplayText(QtObjectPtr ptr);
int QLineEdit_DragEnabled(QtObjectPtr ptr);
int QLineEdit_EchoMode(QtObjectPtr ptr);
int QLineEdit_HasAcceptableInput(QtObjectPtr ptr);
int QLineEdit_HasFrame(QtObjectPtr ptr);
int QLineEdit_HasSelectedText(QtObjectPtr ptr);
char* QLineEdit_InputMask(QtObjectPtr ptr);
int QLineEdit_IsClearButtonEnabled(QtObjectPtr ptr);
int QLineEdit_IsModified(QtObjectPtr ptr);
int QLineEdit_IsReadOnly(QtObjectPtr ptr);
int QLineEdit_IsRedoAvailable(QtObjectPtr ptr);
int QLineEdit_IsUndoAvailable(QtObjectPtr ptr);
int QLineEdit_MaxLength(QtObjectPtr ptr);
char* QLineEdit_PlaceholderText(QtObjectPtr ptr);
char* QLineEdit_SelectedText(QtObjectPtr ptr);
void QLineEdit_SetAlignment(QtObjectPtr ptr, int flag);
void QLineEdit_SetClearButtonEnabled(QtObjectPtr ptr, int enable);
void QLineEdit_SetCursorMoveStyle(QtObjectPtr ptr, int style);
void QLineEdit_SetCursorPosition(QtObjectPtr ptr, int v);
void QLineEdit_SetDragEnabled(QtObjectPtr ptr, int b);
void QLineEdit_SetEchoMode(QtObjectPtr ptr, int v);
void QLineEdit_SetFrame(QtObjectPtr ptr, int v);
void QLineEdit_SetInputMask(QtObjectPtr ptr, char* inputMask);
void QLineEdit_SetMaxLength(QtObjectPtr ptr, int v);
void QLineEdit_SetModified(QtObjectPtr ptr, int v);
void QLineEdit_SetPlaceholderText(QtObjectPtr ptr, char* v);
void QLineEdit_SetReadOnly(QtObjectPtr ptr, int v);
void QLineEdit_SetText(QtObjectPtr ptr, char* v);
char* QLineEdit_Text(QtObjectPtr ptr);
QtObjectPtr QLineEdit_NewQLineEdit(QtObjectPtr parent);
QtObjectPtr QLineEdit_NewQLineEdit2(char* contents, QtObjectPtr parent);
QtObjectPtr QLineEdit_AddAction2(QtObjectPtr ptr, QtObjectPtr icon, int position);
void QLineEdit_AddAction(QtObjectPtr ptr, QtObjectPtr action, int position);
void QLineEdit_Backspace(QtObjectPtr ptr);
void QLineEdit_Clear(QtObjectPtr ptr);
QtObjectPtr QLineEdit_Completer(QtObjectPtr ptr);
void QLineEdit_Copy(QtObjectPtr ptr);
QtObjectPtr QLineEdit_CreateStandardContextMenu(QtObjectPtr ptr);
void QLineEdit_CursorBackward(QtObjectPtr ptr, int mark, int steps);
void QLineEdit_CursorForward(QtObjectPtr ptr, int mark, int steps);
int QLineEdit_CursorPositionAt(QtObjectPtr ptr, QtObjectPtr pos);
void QLineEdit_ConnectCursorPositionChanged(QtObjectPtr ptr);
void QLineEdit_DisconnectCursorPositionChanged(QtObjectPtr ptr);
void QLineEdit_CursorWordBackward(QtObjectPtr ptr, int mark);
void QLineEdit_CursorWordForward(QtObjectPtr ptr, int mark);
void QLineEdit_Cut(QtObjectPtr ptr);
void QLineEdit_Del(QtObjectPtr ptr);
void QLineEdit_Deselect(QtObjectPtr ptr);
void QLineEdit_ConnectEditingFinished(QtObjectPtr ptr);
void QLineEdit_DisconnectEditingFinished(QtObjectPtr ptr);
void QLineEdit_End(QtObjectPtr ptr, int mark);
int QLineEdit_Event(QtObjectPtr ptr, QtObjectPtr e);
void QLineEdit_GetTextMargins(QtObjectPtr ptr, int left, int top, int right, int bottom);
void QLineEdit_Home(QtObjectPtr ptr, int mark);
char* QLineEdit_InputMethodQuery(QtObjectPtr ptr, int property);
void QLineEdit_Paste(QtObjectPtr ptr);
void QLineEdit_Redo(QtObjectPtr ptr);
void QLineEdit_ConnectReturnPressed(QtObjectPtr ptr);
void QLineEdit_DisconnectReturnPressed(QtObjectPtr ptr);
void QLineEdit_SelectAll(QtObjectPtr ptr);
void QLineEdit_ConnectSelectionChanged(QtObjectPtr ptr);
void QLineEdit_DisconnectSelectionChanged(QtObjectPtr ptr);
int QLineEdit_SelectionStart(QtObjectPtr ptr);
void QLineEdit_SetCompleter(QtObjectPtr ptr, QtObjectPtr c);
void QLineEdit_SetSelection(QtObjectPtr ptr, int start, int length);
void QLineEdit_SetTextMargins2(QtObjectPtr ptr, QtObjectPtr margins);
void QLineEdit_SetTextMargins(QtObjectPtr ptr, int left, int top, int right, int bottom);
void QLineEdit_SetValidator(QtObjectPtr ptr, QtObjectPtr v);
void QLineEdit_ConnectTextChanged(QtObjectPtr ptr);
void QLineEdit_DisconnectTextChanged(QtObjectPtr ptr);
void QLineEdit_ConnectTextEdited(QtObjectPtr ptr);
void QLineEdit_DisconnectTextEdited(QtObjectPtr ptr);
void QLineEdit_Undo(QtObjectPtr ptr);
QtObjectPtr QLineEdit_Validator(QtObjectPtr ptr);
void QLineEdit_DestroyQLineEdit(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif