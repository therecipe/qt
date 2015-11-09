#ifdef __cplusplus
extern "C" {
#endif

int QPlainTextEdit_BackgroundVisible(void* ptr);
int QPlainTextEdit_BlockCount(void* ptr);
int QPlainTextEdit_CenterOnScroll(void* ptr);
int QPlainTextEdit_CursorWidth(void* ptr);
int QPlainTextEdit_IsReadOnly(void* ptr);
int QPlainTextEdit_LineWrapMode(void* ptr);
int QPlainTextEdit_OverwriteMode(void* ptr);
char* QPlainTextEdit_PlaceholderText(void* ptr);
void QPlainTextEdit_Redo(void* ptr);
void QPlainTextEdit_SetBackgroundVisible(void* ptr, int visible);
void QPlainTextEdit_SetCenterOnScroll(void* ptr, int enabled);
void QPlainTextEdit_SetCursorWidth(void* ptr, int width);
void QPlainTextEdit_SetLineWrapMode(void* ptr, int mode);
void QPlainTextEdit_SetOverwriteMode(void* ptr, int overwrite);
void QPlainTextEdit_SetPlaceholderText(void* ptr, char* placeholderText);
void QPlainTextEdit_SetReadOnly(void* ptr, int ro);
void QPlainTextEdit_SetTabChangesFocus(void* ptr, int b);
void QPlainTextEdit_SetTabStopWidth(void* ptr, int width);
void QPlainTextEdit_SetTextInteractionFlags(void* ptr, int flags);
void QPlainTextEdit_SetWordWrapMode(void* ptr, int policy);
int QPlainTextEdit_TabChangesFocus(void* ptr);
int QPlainTextEdit_TabStopWidth(void* ptr);
int QPlainTextEdit_TextInteractionFlags(void* ptr);
int QPlainTextEdit_WordWrapMode(void* ptr);
void QPlainTextEdit_ZoomIn(void* ptr, int ran);
void QPlainTextEdit_ZoomOut(void* ptr, int ran);
void* QPlainTextEdit_NewQPlainTextEdit(void* parent);
void* QPlainTextEdit_NewQPlainTextEdit2(char* text, void* parent);
char* QPlainTextEdit_AnchorAt(void* ptr, void* pos);
void QPlainTextEdit_AppendPlainText(void* ptr, char* text);
void QPlainTextEdit_CenterCursor(void* ptr);
void QPlainTextEdit_AppendHtml(void* ptr, char* html);
void QPlainTextEdit_ConnectBlockCountChanged(void* ptr);
void QPlainTextEdit_DisconnectBlockCountChanged(void* ptr);
int QPlainTextEdit_CanPaste(void* ptr);
void QPlainTextEdit_Clear(void* ptr);
void QPlainTextEdit_Copy(void* ptr);
void QPlainTextEdit_ConnectCopyAvailable(void* ptr);
void QPlainTextEdit_DisconnectCopyAvailable(void* ptr);
void* QPlainTextEdit_CreateStandardContextMenu(void* ptr);
void* QPlainTextEdit_CreateStandardContextMenu2(void* ptr, void* position);
void QPlainTextEdit_ConnectCursorPositionChanged(void* ptr);
void QPlainTextEdit_DisconnectCursorPositionChanged(void* ptr);
void QPlainTextEdit_Cut(void* ptr);
void* QPlainTextEdit_Document(void* ptr);
char* QPlainTextEdit_DocumentTitle(void* ptr);
void QPlainTextEdit_EnsureCursorVisible(void* ptr);
void* QPlainTextEdit_InputMethodQuery(void* ptr, int property);
void QPlainTextEdit_InsertPlainText(void* ptr, char* text);
int QPlainTextEdit_IsUndoRedoEnabled(void* ptr);
void* QPlainTextEdit_LoadResource(void* ptr, int ty, void* name);
int QPlainTextEdit_MaximumBlockCount(void* ptr);
void QPlainTextEdit_MergeCurrentCharFormat(void* ptr, void* modifier);
void QPlainTextEdit_ConnectModificationChanged(void* ptr);
void QPlainTextEdit_DisconnectModificationChanged(void* ptr);
void QPlainTextEdit_MoveCursor(void* ptr, int operation, int mode);
void QPlainTextEdit_Paste(void* ptr);
void QPlainTextEdit_Print(void* ptr, void* printer);
void QPlainTextEdit_ConnectRedoAvailable(void* ptr);
void QPlainTextEdit_DisconnectRedoAvailable(void* ptr);
void QPlainTextEdit_SelectAll(void* ptr);
void QPlainTextEdit_ConnectSelectionChanged(void* ptr);
void QPlainTextEdit_DisconnectSelectionChanged(void* ptr);
void QPlainTextEdit_SetCurrentCharFormat(void* ptr, void* format);
void QPlainTextEdit_SetDocument(void* ptr, void* document);
void QPlainTextEdit_SetDocumentTitle(void* ptr, char* title);
void QPlainTextEdit_SetMaximumBlockCount(void* ptr, int maximum);
void QPlainTextEdit_SetPlainText(void* ptr, char* text);
void QPlainTextEdit_SetTextCursor(void* ptr, void* cursor);
void QPlainTextEdit_SetUndoRedoEnabled(void* ptr, int enable);
void QPlainTextEdit_ConnectTextChanged(void* ptr);
void QPlainTextEdit_DisconnectTextChanged(void* ptr);
char* QPlainTextEdit_ToPlainText(void* ptr);
void QPlainTextEdit_Undo(void* ptr);
void QPlainTextEdit_ConnectUndoAvailable(void* ptr);
void QPlainTextEdit_DisconnectUndoAvailable(void* ptr);
void QPlainTextEdit_DestroyQPlainTextEdit(void* ptr);

#ifdef __cplusplus
}
#endif