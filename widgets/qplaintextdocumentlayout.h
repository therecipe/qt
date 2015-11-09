#ifdef __cplusplus
extern "C" {
#endif

int QPlainTextDocumentLayout_CursorWidth(void* ptr);
void QPlainTextDocumentLayout_SetCursorWidth(void* ptr, int width);
void* QPlainTextDocumentLayout_NewQPlainTextDocumentLayout(void* document);
void QPlainTextDocumentLayout_EnsureBlockLayout(void* ptr, void* block);
int QPlainTextDocumentLayout_PageCount(void* ptr);
void QPlainTextDocumentLayout_RequestUpdate(void* ptr);
void QPlainTextDocumentLayout_DestroyQPlainTextDocumentLayout(void* ptr);

#ifdef __cplusplus
}
#endif