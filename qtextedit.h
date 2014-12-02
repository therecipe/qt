#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QTextEdit_New_QWidget(QtObjectPtr parent);
QtObjectPtr QTextEdit_New_String_QWidget(char* text, QtObjectPtr parent);
void QTextEdit_Destroy(QtObjectPtr ptr);
int QTextEdit_AcceptRichText(QtObjectPtr ptr);
int QTextEdit_Alignment(QtObjectPtr ptr);
int QTextEdit_CanPaste(QtObjectPtr ptr);
QtObjectPtr QTextEdit_CreateStandardContextMenu(QtObjectPtr ptr);
int QTextEdit_CursorWidth(QtObjectPtr ptr);
char* QTextEdit_DocumentTitle(QtObjectPtr ptr);
void QTextEdit_EnsureCursorVisible(QtObjectPtr ptr);
char* QTextEdit_FontFamily(QtObjectPtr ptr);
int QTextEdit_FontItalic(QtObjectPtr ptr);
int QTextEdit_FontUnderline(QtObjectPtr ptr);
int QTextEdit_FontWeight(QtObjectPtr ptr);
int QTextEdit_IsReadOnly(QtObjectPtr ptr);
int QTextEdit_IsUndoRedoEnabled(QtObjectPtr ptr);
int QTextEdit_LineWrapColumnOrWidth(QtObjectPtr ptr);
int QTextEdit_OverwriteMode(QtObjectPtr ptr);
char* QTextEdit_PlaceholderText(QtObjectPtr ptr);
void QTextEdit_SetAcceptRichText_Bool(QtObjectPtr ptr, int accept);
void QTextEdit_SetCursorWidth_Int(QtObjectPtr ptr, int width);
void QTextEdit_SetDocumentTitle_String(QtObjectPtr ptr, char* title);
void QTextEdit_SetLineWrapColumnOrWidth_Int(QtObjectPtr ptr, int w);
void QTextEdit_SetOverwriteMode_Bool(QtObjectPtr ptr, int overwrite);
void QTextEdit_SetPlaceholderText_String(QtObjectPtr ptr, char* placeholderText);
void QTextEdit_SetReadOnly_Bool(QtObjectPtr ptr, int ro);
void QTextEdit_SetTabChangesFocus_Bool(QtObjectPtr ptr, int b);
void QTextEdit_SetTabStopWidth_Int(QtObjectPtr ptr, int width);
void QTextEdit_SetTextInteractionFlags_TextInteractionFlag(QtObjectPtr ptr, int flags);
void QTextEdit_SetUndoRedoEnabled_Bool(QtObjectPtr ptr, int enable);
int QTextEdit_TabChangesFocus(QtObjectPtr ptr);
int QTextEdit_TabStopWidth(QtObjectPtr ptr);
int QTextEdit_TextInteractionFlags(QtObjectPtr ptr);
char* QTextEdit_ToHtml(QtObjectPtr ptr);
char* QTextEdit_ToPlainText(QtObjectPtr ptr);
//Public Slots
void QTextEdit_ConnectSlotClear(QtObjectPtr ptr);
void QTextEdit_DisconnectSlotClear(QtObjectPtr ptr);
void QTextEdit_Clear(QtObjectPtr ptr);
void QTextEdit_ConnectSlotCopy(QtObjectPtr ptr);
void QTextEdit_DisconnectSlotCopy(QtObjectPtr ptr);
void QTextEdit_Copy(QtObjectPtr ptr);
void QTextEdit_ConnectSlotCut(QtObjectPtr ptr);
void QTextEdit_DisconnectSlotCut(QtObjectPtr ptr);
void QTextEdit_Cut(QtObjectPtr ptr);
void QTextEdit_ConnectSlotPaste(QtObjectPtr ptr);
void QTextEdit_DisconnectSlotPaste(QtObjectPtr ptr);
void QTextEdit_Paste(QtObjectPtr ptr);
void QTextEdit_ConnectSlotRedo(QtObjectPtr ptr);
void QTextEdit_DisconnectSlotRedo(QtObjectPtr ptr);
void QTextEdit_Redo(QtObjectPtr ptr);
void QTextEdit_ConnectSlotSelectAll(QtObjectPtr ptr);
void QTextEdit_DisconnectSlotSelectAll(QtObjectPtr ptr);
void QTextEdit_SelectAll(QtObjectPtr ptr);
void QTextEdit_ConnectSlotSetFontItalic(QtObjectPtr ptr);
void QTextEdit_DisconnectSlotSetFontItalic(QtObjectPtr ptr);
void QTextEdit_SetFontItalic_Bool(QtObjectPtr ptr, int italic);
void QTextEdit_ConnectSlotSetFontUnderline(QtObjectPtr ptr);
void QTextEdit_DisconnectSlotSetFontUnderline(QtObjectPtr ptr);
void QTextEdit_SetFontUnderline_Bool(QtObjectPtr ptr, int underline);
void QTextEdit_ConnectSlotSetFontWeight(QtObjectPtr ptr);
void QTextEdit_DisconnectSlotSetFontWeight(QtObjectPtr ptr);
void QTextEdit_SetFontWeight_Int(QtObjectPtr ptr, int weight);
void QTextEdit_ConnectSlotUndo(QtObjectPtr ptr);
void QTextEdit_DisconnectSlotUndo(QtObjectPtr ptr);
void QTextEdit_Undo(QtObjectPtr ptr);
void QTextEdit_ConnectSlotZoomIn(QtObjectPtr ptr);
void QTextEdit_DisconnectSlotZoomIn(QtObjectPtr ptr);
void QTextEdit_ZoomIn_Int(QtObjectPtr ptr, int rang);
void QTextEdit_ConnectSlotZoomOut(QtObjectPtr ptr);
void QTextEdit_DisconnectSlotZoomOut(QtObjectPtr ptr);
void QTextEdit_ZoomOut_Int(QtObjectPtr ptr, int rang);
//Signals
void QTextEdit_ConnectSignalCopyAvailable(QtObjectPtr ptr);
void QTextEdit_DisconnectSignalCopyAvailable(QtObjectPtr ptr);
void QTextEdit_ConnectSignalCursorPositionChanged(QtObjectPtr ptr);
void QTextEdit_DisconnectSignalCursorPositionChanged(QtObjectPtr ptr);
void QTextEdit_ConnectSignalRedoAvailable(QtObjectPtr ptr);
void QTextEdit_DisconnectSignalRedoAvailable(QtObjectPtr ptr);
void QTextEdit_ConnectSignalSelectionChanged(QtObjectPtr ptr);
void QTextEdit_DisconnectSignalSelectionChanged(QtObjectPtr ptr);
void QTextEdit_ConnectSignalTextChanged(QtObjectPtr ptr);
void QTextEdit_DisconnectSignalTextChanged(QtObjectPtr ptr);
void QTextEdit_ConnectSignalUndoAvailable(QtObjectPtr ptr);
void QTextEdit_DisconnectSignalUndoAvailable(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
