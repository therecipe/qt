#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QAccessibleTextInterface_AddSelection(QtObjectPtr ptr, int startOffset, int endOffset);
char* QAccessibleTextInterface_Attributes(QtObjectPtr ptr, int offset, int startOffset, int endOffset);
int QAccessibleTextInterface_CharacterCount(QtObjectPtr ptr);
int QAccessibleTextInterface_CursorPosition(QtObjectPtr ptr);
int QAccessibleTextInterface_OffsetAtPoint(QtObjectPtr ptr, QtObjectPtr point);
void QAccessibleTextInterface_RemoveSelection(QtObjectPtr ptr, int selectionIndex);
void QAccessibleTextInterface_ScrollToSubstring(QtObjectPtr ptr, int startIndex, int endIndex);
void QAccessibleTextInterface_Selection(QtObjectPtr ptr, int selectionIndex, int startOffset, int endOffset);
int QAccessibleTextInterface_SelectionCount(QtObjectPtr ptr);
void QAccessibleTextInterface_SetCursorPosition(QtObjectPtr ptr, int position);
void QAccessibleTextInterface_SetSelection(QtObjectPtr ptr, int selectionIndex, int startOffset, int endOffset);
char* QAccessibleTextInterface_Text(QtObjectPtr ptr, int startOffset, int endOffset);
char* QAccessibleTextInterface_TextAfterOffset(QtObjectPtr ptr, int offset, int boundaryType, int startOffset, int endOffset);
char* QAccessibleTextInterface_TextAtOffset(QtObjectPtr ptr, int offset, int boundaryType, int startOffset, int endOffset);
char* QAccessibleTextInterface_TextBeforeOffset(QtObjectPtr ptr, int offset, int boundaryType, int startOffset, int endOffset);
void QAccessibleTextInterface_DestroyQAccessibleTextInterface(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif