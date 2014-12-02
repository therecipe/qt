#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Types
int QLineEdit_LeadingPosition();
int QLineEdit_TrailingPosition();
int QLineEdit_Normal();
int QLineEdit_NoEcho();
int QLineEdit_Password();
int QLineEdit_PasswordEchoOnEdit();
//Public Functions
QtObjectPtr QLineEdit_New_QWidget(QtObjectPtr parent);
QtObjectPtr QLineEdit_New_String_QWidget(char* contents, QtObjectPtr parent);
void QLineEdit_Destroy(QtObjectPtr ptr);
int QLineEdit_Alignment(QtObjectPtr ptr);
void QLineEdit_Backspace(QtObjectPtr ptr);
QtObjectPtr QLineEdit_CreateStandardContextMenu(QtObjectPtr ptr);
void QLineEdit_CursorBackward_Bool_Int(QtObjectPtr ptr, int mark, int steps);
void QLineEdit_CursorForward_Bool_Int(QtObjectPtr ptr, int mark, int steps);
int QLineEdit_CursorMoveStyle(QtObjectPtr ptr);
int QLineEdit_CursorPosition(QtObjectPtr ptr);
void QLineEdit_CursorWordBackward_Bool(QtObjectPtr ptr, int mark);
void QLineEdit_CursorWordForward_Bool(QtObjectPtr ptr, int mark);
void QLineEdit_Del(QtObjectPtr ptr);
void QLineEdit_Deselect(QtObjectPtr ptr);
char* QLineEdit_DisplayText(QtObjectPtr ptr);
int QLineEdit_DragEnabled(QtObjectPtr ptr);
int QLineEdit_EchoMode(QtObjectPtr ptr);
void QLineEdit_End_Bool(QtObjectPtr ptr, int mark);
int QLineEdit_HasAcceptableInput(QtObjectPtr ptr);
int QLineEdit_HasFrame(QtObjectPtr ptr);
int QLineEdit_HasSelectedText(QtObjectPtr ptr);
void QLineEdit_Home_Bool(QtObjectPtr ptr, int mark);
char* QLineEdit_InputMask(QtObjectPtr ptr);
void QLineEdit_Insert_String(QtObjectPtr ptr, char* newText);
int QLineEdit_IsClearButtonEnabled(QtObjectPtr ptr);
int QLineEdit_IsModified(QtObjectPtr ptr);
int QLineEdit_IsReadOnly(QtObjectPtr ptr);
int QLineEdit_IsRedoAvailable(QtObjectPtr ptr);
int QLineEdit_IsUndoAvailable(QtObjectPtr ptr);
int QLineEdit_MaxLength(QtObjectPtr ptr);
char* QLineEdit_PlaceholderText(QtObjectPtr ptr);
char* QLineEdit_SelectedText(QtObjectPtr ptr);
int QLineEdit_SelectionStart(QtObjectPtr ptr);
void QLineEdit_SetAlignment_AlignmentFlag(QtObjectPtr ptr, int flag);
void QLineEdit_SetClearButtonEnabled_Bool(QtObjectPtr ptr, int enable);
void QLineEdit_SetCursorMoveStyle_CursorMoveStyle(QtObjectPtr ptr, int style);
void QLineEdit_SetCursorPosition_Int(QtObjectPtr ptr, int cursorPosition);
void QLineEdit_SetDragEnabled_Bool(QtObjectPtr ptr, int dragEnabled);
void QLineEdit_SetEchoMode_EchoMode(QtObjectPtr ptr, int Ec);
void QLineEdit_SetFrame_Bool(QtObjectPtr ptr, int frame);
void QLineEdit_SetInputMask_String(QtObjectPtr ptr, char* inputMask);
void QLineEdit_SetMaxLength_Int(QtObjectPtr ptr, int maxLength);
void QLineEdit_SetModified_Bool(QtObjectPtr ptr, int modified);
void QLineEdit_SetPlaceholderText_String(QtObjectPtr ptr, char* placeholderText);
void QLineEdit_SetReadOnly_Bool(QtObjectPtr ptr, int readOnly);
void QLineEdit_SetSelection_Int_Int(QtObjectPtr ptr, int start, int length);
void QLineEdit_SetTextMargins_Int_Int_Int_Int(QtObjectPtr ptr, int left, int top, int right, int bottom);
void QLineEdit_SetValidator_QValidator(QtObjectPtr ptr, QtObjectPtr v);
char* QLineEdit_Text(QtObjectPtr ptr);
//Public Slots
void QLineEdit_ConnectSlotClear(QtObjectPtr ptr);
void QLineEdit_DisconnectSlotClear(QtObjectPtr ptr);
void QLineEdit_Clear(QtObjectPtr ptr);
void QLineEdit_ConnectSlotCopy(QtObjectPtr ptr);
void QLineEdit_DisconnectSlotCopy(QtObjectPtr ptr);
void QLineEdit_Copy(QtObjectPtr ptr);
void QLineEdit_ConnectSlotCut(QtObjectPtr ptr);
void QLineEdit_DisconnectSlotCut(QtObjectPtr ptr);
void QLineEdit_Cut(QtObjectPtr ptr);
void QLineEdit_ConnectSlotPaste(QtObjectPtr ptr);
void QLineEdit_DisconnectSlotPaste(QtObjectPtr ptr);
void QLineEdit_Paste(QtObjectPtr ptr);
void QLineEdit_ConnectSlotRedo(QtObjectPtr ptr);
void QLineEdit_DisconnectSlotRedo(QtObjectPtr ptr);
void QLineEdit_Redo(QtObjectPtr ptr);
void QLineEdit_ConnectSlotSelectAll(QtObjectPtr ptr);
void QLineEdit_DisconnectSlotSelectAll(QtObjectPtr ptr);
void QLineEdit_SelectAll(QtObjectPtr ptr);
void QLineEdit_ConnectSlotSetText(QtObjectPtr ptr);
void QLineEdit_DisconnectSlotSetText(QtObjectPtr ptr);
void QLineEdit_SetText_String(QtObjectPtr ptr, char* text);
void QLineEdit_ConnectSlotUndo(QtObjectPtr ptr);
void QLineEdit_DisconnectSlotUndo(QtObjectPtr ptr);
void QLineEdit_Undo(QtObjectPtr ptr);
//Signals
void QLineEdit_ConnectSignalCursorPositionChanged(QtObjectPtr ptr);
void QLineEdit_DisconnectSignalCursorPositionChanged(QtObjectPtr ptr);
void QLineEdit_ConnectSignalEditingFinished(QtObjectPtr ptr);
void QLineEdit_DisconnectSignalEditingFinished(QtObjectPtr ptr);
void QLineEdit_ConnectSignalReturnPressed(QtObjectPtr ptr);
void QLineEdit_DisconnectSignalReturnPressed(QtObjectPtr ptr);
void QLineEdit_ConnectSignalSelectionChanged(QtObjectPtr ptr);
void QLineEdit_DisconnectSignalSelectionChanged(QtObjectPtr ptr);
void QLineEdit_ConnectSignalTextChanged(QtObjectPtr ptr);
void QLineEdit_DisconnectSignalTextChanged(QtObjectPtr ptr);
void QLineEdit_ConnectSignalTextEdited(QtObjectPtr ptr);
void QLineEdit_DisconnectSignalTextEdited(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
