#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QTextCursor_InsertBlock3(QtObjectPtr ptr, QtObjectPtr format, QtObjectPtr charFormat);
QtObjectPtr QTextCursor_InsertTable2(QtObjectPtr ptr, int rows, int columns);
QtObjectPtr QTextCursor_InsertTable(QtObjectPtr ptr, int rows, int columns, QtObjectPtr format);
void QTextCursor_InsertText2(QtObjectPtr ptr, char* text, QtObjectPtr format);
int QTextCursor_MovePosition(QtObjectPtr ptr, int operation, int mode, int n);
QtObjectPtr QTextCursor_NewQTextCursor();
QtObjectPtr QTextCursor_NewQTextCursor2(QtObjectPtr document);
QtObjectPtr QTextCursor_NewQTextCursor4(QtObjectPtr frame);
QtObjectPtr QTextCursor_NewQTextCursor5(QtObjectPtr block);
QtObjectPtr QTextCursor_NewQTextCursor7(QtObjectPtr cursor);
int QTextCursor_Anchor(QtObjectPtr ptr);
int QTextCursor_AtBlockEnd(QtObjectPtr ptr);
int QTextCursor_AtBlockStart(QtObjectPtr ptr);
int QTextCursor_AtEnd(QtObjectPtr ptr);
int QTextCursor_AtStart(QtObjectPtr ptr);
void QTextCursor_BeginEditBlock(QtObjectPtr ptr);
int QTextCursor_BlockNumber(QtObjectPtr ptr);
void QTextCursor_ClearSelection(QtObjectPtr ptr);
int QTextCursor_ColumnNumber(QtObjectPtr ptr);
QtObjectPtr QTextCursor_CreateList2(QtObjectPtr ptr, int style);
QtObjectPtr QTextCursor_CreateList(QtObjectPtr ptr, QtObjectPtr format);
QtObjectPtr QTextCursor_CurrentFrame(QtObjectPtr ptr);
QtObjectPtr QTextCursor_CurrentList(QtObjectPtr ptr);
QtObjectPtr QTextCursor_CurrentTable(QtObjectPtr ptr);
void QTextCursor_DeleteChar(QtObjectPtr ptr);
void QTextCursor_DeletePreviousChar(QtObjectPtr ptr);
QtObjectPtr QTextCursor_Document(QtObjectPtr ptr);
void QTextCursor_EndEditBlock(QtObjectPtr ptr);
int QTextCursor_HasComplexSelection(QtObjectPtr ptr);
int QTextCursor_HasSelection(QtObjectPtr ptr);
void QTextCursor_InsertBlock(QtObjectPtr ptr);
void QTextCursor_InsertBlock2(QtObjectPtr ptr, QtObjectPtr format);
void QTextCursor_InsertFragment(QtObjectPtr ptr, QtObjectPtr fragment);
QtObjectPtr QTextCursor_InsertFrame(QtObjectPtr ptr, QtObjectPtr format);
void QTextCursor_InsertHtml(QtObjectPtr ptr, char* html);
void QTextCursor_InsertImage4(QtObjectPtr ptr, QtObjectPtr image, char* name);
void QTextCursor_InsertImage3(QtObjectPtr ptr, char* name);
void QTextCursor_InsertImage(QtObjectPtr ptr, QtObjectPtr format);
void QTextCursor_InsertImage2(QtObjectPtr ptr, QtObjectPtr format, int alignment);
QtObjectPtr QTextCursor_InsertList2(QtObjectPtr ptr, int style);
QtObjectPtr QTextCursor_InsertList(QtObjectPtr ptr, QtObjectPtr format);
void QTextCursor_InsertText(QtObjectPtr ptr, char* text);
int QTextCursor_IsCopyOf(QtObjectPtr ptr, QtObjectPtr other);
int QTextCursor_IsNull(QtObjectPtr ptr);
void QTextCursor_JoinPreviousEditBlock(QtObjectPtr ptr);
int QTextCursor_KeepPositionOnInsert(QtObjectPtr ptr);
void QTextCursor_MergeBlockCharFormat(QtObjectPtr ptr, QtObjectPtr modifier);
void QTextCursor_MergeBlockFormat(QtObjectPtr ptr, QtObjectPtr modifier);
void QTextCursor_MergeCharFormat(QtObjectPtr ptr, QtObjectPtr modifier);
int QTextCursor_Position(QtObjectPtr ptr);
int QTextCursor_PositionInBlock(QtObjectPtr ptr);
void QTextCursor_RemoveSelectedText(QtObjectPtr ptr);
void QTextCursor_Select(QtObjectPtr ptr, int selection);
void QTextCursor_SelectedTableCells(QtObjectPtr ptr, int firstRow, int numRows, int firstColumn, int numColumns);
char* QTextCursor_SelectedText(QtObjectPtr ptr);
int QTextCursor_SelectionEnd(QtObjectPtr ptr);
int QTextCursor_SelectionStart(QtObjectPtr ptr);
void QTextCursor_SetBlockCharFormat(QtObjectPtr ptr, QtObjectPtr format);
void QTextCursor_SetBlockFormat(QtObjectPtr ptr, QtObjectPtr format);
void QTextCursor_SetCharFormat(QtObjectPtr ptr, QtObjectPtr format);
void QTextCursor_SetKeepPositionOnInsert(QtObjectPtr ptr, int b);
void QTextCursor_SetPosition(QtObjectPtr ptr, int pos, int m);
void QTextCursor_SetVerticalMovementX(QtObjectPtr ptr, int x);
void QTextCursor_SetVisualNavigation(QtObjectPtr ptr, int b);
void QTextCursor_Swap(QtObjectPtr ptr, QtObjectPtr other);
int QTextCursor_VerticalMovementX(QtObjectPtr ptr);
int QTextCursor_VisualNavigation(QtObjectPtr ptr);
void QTextCursor_DestroyQTextCursor(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif