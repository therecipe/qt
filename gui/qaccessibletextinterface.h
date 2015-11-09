#ifdef __cplusplus
extern "C" {
#endif

void QAccessibleTextInterface_AddSelection(void* ptr, int startOffset, int endOffset);
char* QAccessibleTextInterface_Attributes(void* ptr, int offset, int startOffset, int endOffset);
int QAccessibleTextInterface_CharacterCount(void* ptr);
int QAccessibleTextInterface_CursorPosition(void* ptr);
int QAccessibleTextInterface_OffsetAtPoint(void* ptr, void* point);
void QAccessibleTextInterface_RemoveSelection(void* ptr, int selectionIndex);
void QAccessibleTextInterface_ScrollToSubstring(void* ptr, int startIndex, int endIndex);
void QAccessibleTextInterface_Selection(void* ptr, int selectionIndex, int startOffset, int endOffset);
int QAccessibleTextInterface_SelectionCount(void* ptr);
void QAccessibleTextInterface_SetCursorPosition(void* ptr, int position);
void QAccessibleTextInterface_SetSelection(void* ptr, int selectionIndex, int startOffset, int endOffset);
char* QAccessibleTextInterface_Text(void* ptr, int startOffset, int endOffset);
char* QAccessibleTextInterface_TextAfterOffset(void* ptr, int offset, int boundaryType, int startOffset, int endOffset);
char* QAccessibleTextInterface_TextAtOffset(void* ptr, int offset, int boundaryType, int startOffset, int endOffset);
char* QAccessibleTextInterface_TextBeforeOffset(void* ptr, int offset, int boundaryType, int startOffset, int endOffset);
void QAccessibleTextInterface_DestroyQAccessibleTextInterface(void* ptr);

#ifdef __cplusplus
}
#endif