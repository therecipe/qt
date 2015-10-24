#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QAccessibleEditableTextInterface_DeleteText(QtObjectPtr ptr, int startOffset, int endOffset);
void QAccessibleEditableTextInterface_InsertText(QtObjectPtr ptr, int offset, char* text);
void QAccessibleEditableTextInterface_ReplaceText(QtObjectPtr ptr, int startOffset, int endOffset, char* text);
void QAccessibleEditableTextInterface_DestroyQAccessibleEditableTextInterface(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif