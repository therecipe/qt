#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QPlainTextEdit_BackgroundVisible(QtObjectPtr ptr);
int QPlainTextEdit_BlockCount(QtObjectPtr ptr);
int QPlainTextEdit_CenterOnScroll(QtObjectPtr ptr);
int QPlainTextEdit_CursorWidth(QtObjectPtr ptr);
int QPlainTextEdit_IsReadOnly(QtObjectPtr ptr);
int QPlainTextEdit_LineWrapMode(QtObjectPtr ptr);
int QPlainTextEdit_OverwriteMode(QtObjectPtr ptr);
char* QPlainTextEdit_PlaceholderText(QtObjectPtr ptr);
void QPlainTextEdit_Redo(QtObjectPtr ptr);
void QPlainTextEdit_SetBackgroundVisible(QtObjectPtr ptr, int visible);
void QPlainTextEdit_SetCenterOnScroll(QtObjectPtr ptr, int enabled);
void QPlainTextEdit_SetCursorWidth(QtObjectPtr ptr, int width);
void QPlainTextEdit_SetLineWrapMode(QtObjectPtr ptr, int mode);
void QPlainTextEdit_SetOverwriteMode(QtObjectPtr ptr, int overwrite);
void QPlainTextEdit_SetPlaceholderText(QtObjectPtr ptr, char* placeholderText);
void QPlainTextEdit_SetReadOnly(QtObjectPtr ptr, int ro);
void QPlainTextEdit_SetTabChangesFocus(QtObjectPtr ptr, int b);
void QPlainTextEdit_SetTabStopWidth(QtObjectPtr ptr, int width);
void QPlainTextEdit_SetTextInteractionFlags(QtObjectPtr ptr, int flags);
void QPlainTextEdit_SetWordWrapMode(QtObjectPtr ptr, int policy);
int QPlainTextEdit_TabChangesFocus(QtObjectPtr ptr);
int QPlainTextEdit_TabStopWidth(QtObjectPtr ptr);
int QPlainTextEdit_TextInteractionFlags(QtObjectPtr ptr);
int QPlainTextEdit_WordWrapMode(QtObjectPtr ptr);
void QPlainTextEdit_ZoomIn(QtObjectPtr ptr, int ran);
void QPlainTextEdit_ZoomOut(QtObjectPtr ptr, int ran);
QtObjectPtr QPlainTextEdit_NewQPlainTextEdit(QtObjectPtr parent);
QtObjectPtr QPlainTextEdit_NewQPlainTextEdit2(char* text, QtObjectPtr parent);
char* QPlainTextEdit_AnchorAt(QtObjectPtr ptr, QtObjectPtr pos);
void QPlainTextEdit_AppendPlainText(QtObjectPtr ptr, char* text);
void QPlainTextEdit_CenterCursor(QtObjectPtr ptr);
void QPlainTextEdit_AppendHtml(QtObjectPtr ptr, char* html);
void QPlainTextEdit_ConnectBlockCountChanged(QtObjectPtr ptr);
void QPlainTextEdit_DisconnectBlockCountChanged(QtObjectPtr ptr);
int QPlainTextEdit_CanPaste(QtObjectPtr ptr);
void QPlainTextEdit_Clear(QtObjectPtr ptr);
void QPlainTextEdit_Copy(QtObjectPtr ptr);
void QPlainTextEdit_ConnectCopyAvailable(QtObjectPtr ptr);
void QPlainTextEdit_DisconnectCopyAvailable(QtObjectPtr ptr);
QtObjectPtr QPlainTextEdit_CreateStandardContextMenu(QtObjectPtr ptr);
QtObjectPtr QPlainTextEdit_CreateStandardContextMenu2(QtObjectPtr ptr, QtObjectPtr position);
void QPlainTextEdit_ConnectCursorPositionChanged(QtObjectPtr ptr);
void QPlainTextEdit_DisconnectCursorPositionChanged(QtObjectPtr ptr);
void QPlainTextEdit_Cut(QtObjectPtr ptr);
QtObjectPtr QPlainTextEdit_Document(QtObjectPtr ptr);
char* QPlainTextEdit_DocumentTitle(QtObjectPtr ptr);
void QPlainTextEdit_EnsureCursorVisible(QtObjectPtr ptr);
char* QPlainTextEdit_InputMethodQuery(QtObjectPtr ptr, int property);
void QPlainTextEdit_InsertPlainText(QtObjectPtr ptr, char* text);
int QPlainTextEdit_IsUndoRedoEnabled(QtObjectPtr ptr);
char* QPlainTextEdit_LoadResource(QtObjectPtr ptr, int ty, char* name);
int QPlainTextEdit_MaximumBlockCount(QtObjectPtr ptr);
void QPlainTextEdit_MergeCurrentCharFormat(QtObjectPtr ptr, QtObjectPtr modifier);
void QPlainTextEdit_ConnectModificationChanged(QtObjectPtr ptr);
void QPlainTextEdit_DisconnectModificationChanged(QtObjectPtr ptr);
void QPlainTextEdit_MoveCursor(QtObjectPtr ptr, int operation, int mode);
void QPlainTextEdit_Paste(QtObjectPtr ptr);
void QPlainTextEdit_Print(QtObjectPtr ptr, QtObjectPtr printer);
void QPlainTextEdit_ConnectRedoAvailable(QtObjectPtr ptr);
void QPlainTextEdit_DisconnectRedoAvailable(QtObjectPtr ptr);
void QPlainTextEdit_SelectAll(QtObjectPtr ptr);
void QPlainTextEdit_ConnectSelectionChanged(QtObjectPtr ptr);
void QPlainTextEdit_DisconnectSelectionChanged(QtObjectPtr ptr);
void QPlainTextEdit_SetCurrentCharFormat(QtObjectPtr ptr, QtObjectPtr format);
void QPlainTextEdit_SetDocument(QtObjectPtr ptr, QtObjectPtr document);
void QPlainTextEdit_SetDocumentTitle(QtObjectPtr ptr, char* title);
void QPlainTextEdit_SetMaximumBlockCount(QtObjectPtr ptr, int maximum);
void QPlainTextEdit_SetPlainText(QtObjectPtr ptr, char* text);
void QPlainTextEdit_SetTextCursor(QtObjectPtr ptr, QtObjectPtr cursor);
void QPlainTextEdit_SetUndoRedoEnabled(QtObjectPtr ptr, int enable);
void QPlainTextEdit_ConnectTextChanged(QtObjectPtr ptr);
void QPlainTextEdit_DisconnectTextChanged(QtObjectPtr ptr);
char* QPlainTextEdit_ToPlainText(QtObjectPtr ptr);
void QPlainTextEdit_Undo(QtObjectPtr ptr);
void QPlainTextEdit_ConnectUndoAvailable(QtObjectPtr ptr);
void QPlainTextEdit_DisconnectUndoAvailable(QtObjectPtr ptr);
void QPlainTextEdit_DestroyQPlainTextEdit(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif