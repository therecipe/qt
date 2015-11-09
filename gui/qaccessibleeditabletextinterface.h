#ifdef __cplusplus
extern "C" {
#endif

void QAccessibleEditableTextInterface_DeleteText(void* ptr, int startOffset, int endOffset);
void QAccessibleEditableTextInterface_InsertText(void* ptr, int offset, char* text);
void QAccessibleEditableTextInterface_ReplaceText(void* ptr, int startOffset, int endOffset, char* text);
void QAccessibleEditableTextInterface_DestroyQAccessibleEditableTextInterface(void* ptr);

#ifdef __cplusplus
}
#endif