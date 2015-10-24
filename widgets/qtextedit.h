#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QTextEdit_AcceptRichText(QtObjectPtr ptr);
int QTextEdit_AutoFormatting(QtObjectPtr ptr);
int QTextEdit_CursorWidth(QtObjectPtr ptr);
QtObjectPtr QTextEdit_Document(QtObjectPtr ptr);
int QTextEdit_IsReadOnly(QtObjectPtr ptr);
int QTextEdit_LineWrapColumnOrWidth(QtObjectPtr ptr);
int QTextEdit_LineWrapMode(QtObjectPtr ptr);
int QTextEdit_OverwriteMode(QtObjectPtr ptr);
char* QTextEdit_PlaceholderText(QtObjectPtr ptr);
void QTextEdit_Redo(QtObjectPtr ptr);
void QTextEdit_SetAcceptRichText(QtObjectPtr ptr, int accept);
void QTextEdit_SetAutoFormatting(QtObjectPtr ptr, int features);
void QTextEdit_SetCursorWidth(QtObjectPtr ptr, int width);
void QTextEdit_SetDocument(QtObjectPtr ptr, QtObjectPtr document);
void QTextEdit_SetFontWeight(QtObjectPtr ptr, int weight);
void QTextEdit_SetHtml(QtObjectPtr ptr, char* text);
void QTextEdit_SetLineWrapColumnOrWidth(QtObjectPtr ptr, int w);
void QTextEdit_SetLineWrapMode(QtObjectPtr ptr, int mode);
void QTextEdit_SetOverwriteMode(QtObjectPtr ptr, int overwrite);
void QTextEdit_SetPlaceholderText(QtObjectPtr ptr, char* placeholderText);
void QTextEdit_SetReadOnly(QtObjectPtr ptr, int ro);
void QTextEdit_SetTabChangesFocus(QtObjectPtr ptr, int b);
void QTextEdit_SetTabStopWidth(QtObjectPtr ptr, int width);
void QTextEdit_SetTextInteractionFlags(QtObjectPtr ptr, int flags);
void QTextEdit_SetWordWrapMode(QtObjectPtr ptr, int policy);
int QTextEdit_TabChangesFocus(QtObjectPtr ptr);
int QTextEdit_TabStopWidth(QtObjectPtr ptr);
int QTextEdit_TextInteractionFlags(QtObjectPtr ptr);
char* QTextEdit_ToHtml(QtObjectPtr ptr);
int QTextEdit_WordWrapMode(QtObjectPtr ptr);
void QTextEdit_ZoomIn(QtObjectPtr ptr, int ran);
void QTextEdit_ZoomOut(QtObjectPtr ptr, int ran);
QtObjectPtr QTextEdit_NewQTextEdit(QtObjectPtr parent);
QtObjectPtr QTextEdit_NewQTextEdit2(char* text, QtObjectPtr parent);
int QTextEdit_Alignment(QtObjectPtr ptr);
char* QTextEdit_AnchorAt(QtObjectPtr ptr, QtObjectPtr pos);
void QTextEdit_Append(QtObjectPtr ptr, char* text);
int QTextEdit_CanPaste(QtObjectPtr ptr);
void QTextEdit_Clear(QtObjectPtr ptr);
void QTextEdit_Copy(QtObjectPtr ptr);
void QTextEdit_ConnectCopyAvailable(QtObjectPtr ptr);
void QTextEdit_DisconnectCopyAvailable(QtObjectPtr ptr);
QtObjectPtr QTextEdit_CreateStandardContextMenu(QtObjectPtr ptr);
QtObjectPtr QTextEdit_CreateStandardContextMenu2(QtObjectPtr ptr, QtObjectPtr position);
void QTextEdit_ConnectCursorPositionChanged(QtObjectPtr ptr);
void QTextEdit_DisconnectCursorPositionChanged(QtObjectPtr ptr);
void QTextEdit_Cut(QtObjectPtr ptr);
char* QTextEdit_DocumentTitle(QtObjectPtr ptr);
void QTextEdit_EnsureCursorVisible(QtObjectPtr ptr);
char* QTextEdit_FontFamily(QtObjectPtr ptr);
int QTextEdit_FontItalic(QtObjectPtr ptr);
int QTextEdit_FontUnderline(QtObjectPtr ptr);
int QTextEdit_FontWeight(QtObjectPtr ptr);
char* QTextEdit_InputMethodQuery(QtObjectPtr ptr, int property);
void QTextEdit_InsertHtml(QtObjectPtr ptr, char* text);
void QTextEdit_InsertPlainText(QtObjectPtr ptr, char* text);
int QTextEdit_IsUndoRedoEnabled(QtObjectPtr ptr);
char* QTextEdit_LoadResource(QtObjectPtr ptr, int ty, char* name);
void QTextEdit_MergeCurrentCharFormat(QtObjectPtr ptr, QtObjectPtr modifier);
void QTextEdit_MoveCursor(QtObjectPtr ptr, int operation, int mode);
void QTextEdit_Paste(QtObjectPtr ptr);
void QTextEdit_Print(QtObjectPtr ptr, QtObjectPtr printer);
void QTextEdit_ConnectRedoAvailable(QtObjectPtr ptr);
void QTextEdit_DisconnectRedoAvailable(QtObjectPtr ptr);
void QTextEdit_ScrollToAnchor(QtObjectPtr ptr, char* name);
void QTextEdit_SelectAll(QtObjectPtr ptr);
void QTextEdit_ConnectSelectionChanged(QtObjectPtr ptr);
void QTextEdit_DisconnectSelectionChanged(QtObjectPtr ptr);
void QTextEdit_SetAlignment(QtObjectPtr ptr, int a);
void QTextEdit_SetCurrentCharFormat(QtObjectPtr ptr, QtObjectPtr format);
void QTextEdit_SetCurrentFont(QtObjectPtr ptr, QtObjectPtr f);
void QTextEdit_SetDocumentTitle(QtObjectPtr ptr, char* title);
void QTextEdit_SetFontFamily(QtObjectPtr ptr, char* fontFamily);
void QTextEdit_SetFontItalic(QtObjectPtr ptr, int italic);
void QTextEdit_SetFontUnderline(QtObjectPtr ptr, int underline);
void QTextEdit_SetPlainText(QtObjectPtr ptr, char* text);
void QTextEdit_SetText(QtObjectPtr ptr, char* text);
void QTextEdit_SetTextBackgroundColor(QtObjectPtr ptr, QtObjectPtr c);
void QTextEdit_SetTextColor(QtObjectPtr ptr, QtObjectPtr c);
void QTextEdit_SetTextCursor(QtObjectPtr ptr, QtObjectPtr cursor);
void QTextEdit_SetUndoRedoEnabled(QtObjectPtr ptr, int enable);
void QTextEdit_ConnectTextChanged(QtObjectPtr ptr);
void QTextEdit_DisconnectTextChanged(QtObjectPtr ptr);
char* QTextEdit_ToPlainText(QtObjectPtr ptr);
void QTextEdit_Undo(QtObjectPtr ptr);
void QTextEdit_ConnectUndoAvailable(QtObjectPtr ptr);
void QTextEdit_DisconnectUndoAvailable(QtObjectPtr ptr);
void QTextEdit_DestroyQTextEdit(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif