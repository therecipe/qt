#ifdef __cplusplus
extern "C" {
#endif

void QTextCursor_InsertBlock3(void* ptr, void* format, void* charFormat);
void* QTextCursor_InsertTable2(void* ptr, int rows, int columns);
void* QTextCursor_InsertTable(void* ptr, int rows, int columns, void* format);
void QTextCursor_InsertText2(void* ptr, char* text, void* format);
int QTextCursor_MovePosition(void* ptr, int operation, int mode, int n);
void* QTextCursor_NewQTextCursor();
void* QTextCursor_NewQTextCursor2(void* document);
void* QTextCursor_NewQTextCursor4(void* frame);
void* QTextCursor_NewQTextCursor5(void* block);
void* QTextCursor_NewQTextCursor7(void* cursor);
int QTextCursor_Anchor(void* ptr);
int QTextCursor_AtBlockEnd(void* ptr);
int QTextCursor_AtBlockStart(void* ptr);
int QTextCursor_AtEnd(void* ptr);
int QTextCursor_AtStart(void* ptr);
void QTextCursor_BeginEditBlock(void* ptr);
int QTextCursor_BlockNumber(void* ptr);
void QTextCursor_ClearSelection(void* ptr);
int QTextCursor_ColumnNumber(void* ptr);
void* QTextCursor_CreateList2(void* ptr, int style);
void* QTextCursor_CreateList(void* ptr, void* format);
void* QTextCursor_CurrentFrame(void* ptr);
void* QTextCursor_CurrentList(void* ptr);
void* QTextCursor_CurrentTable(void* ptr);
void QTextCursor_DeleteChar(void* ptr);
void QTextCursor_DeletePreviousChar(void* ptr);
void* QTextCursor_Document(void* ptr);
void QTextCursor_EndEditBlock(void* ptr);
int QTextCursor_HasComplexSelection(void* ptr);
int QTextCursor_HasSelection(void* ptr);
void QTextCursor_InsertBlock(void* ptr);
void QTextCursor_InsertBlock2(void* ptr, void* format);
void QTextCursor_InsertFragment(void* ptr, void* fragment);
void* QTextCursor_InsertFrame(void* ptr, void* format);
void QTextCursor_InsertHtml(void* ptr, char* html);
void QTextCursor_InsertImage4(void* ptr, void* image, char* name);
void QTextCursor_InsertImage3(void* ptr, char* name);
void QTextCursor_InsertImage(void* ptr, void* format);
void QTextCursor_InsertImage2(void* ptr, void* format, int alignment);
void* QTextCursor_InsertList2(void* ptr, int style);
void* QTextCursor_InsertList(void* ptr, void* format);
void QTextCursor_InsertText(void* ptr, char* text);
int QTextCursor_IsCopyOf(void* ptr, void* other);
int QTextCursor_IsNull(void* ptr);
void QTextCursor_JoinPreviousEditBlock(void* ptr);
int QTextCursor_KeepPositionOnInsert(void* ptr);
void QTextCursor_MergeBlockCharFormat(void* ptr, void* modifier);
void QTextCursor_MergeBlockFormat(void* ptr, void* modifier);
void QTextCursor_MergeCharFormat(void* ptr, void* modifier);
int QTextCursor_Position(void* ptr);
int QTextCursor_PositionInBlock(void* ptr);
void QTextCursor_RemoveSelectedText(void* ptr);
void QTextCursor_Select(void* ptr, int selection);
void QTextCursor_SelectedTableCells(void* ptr, int firstRow, int numRows, int firstColumn, int numColumns);
char* QTextCursor_SelectedText(void* ptr);
int QTextCursor_SelectionEnd(void* ptr);
int QTextCursor_SelectionStart(void* ptr);
void QTextCursor_SetBlockCharFormat(void* ptr, void* format);
void QTextCursor_SetBlockFormat(void* ptr, void* format);
void QTextCursor_SetCharFormat(void* ptr, void* format);
void QTextCursor_SetKeepPositionOnInsert(void* ptr, int b);
void QTextCursor_SetPosition(void* ptr, int pos, int m);
void QTextCursor_SetVerticalMovementX(void* ptr, int x);
void QTextCursor_SetVisualNavigation(void* ptr, int b);
void QTextCursor_Swap(void* ptr, void* other);
int QTextCursor_VerticalMovementX(void* ptr);
int QTextCursor_VisualNavigation(void* ptr);
void QTextCursor_DestroyQTextCursor(void* ptr);

#ifdef __cplusplus
}
#endif