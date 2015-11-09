#ifdef __cplusplus
extern "C" {
#endif

void* QTextDocumentFragment_NewQTextDocumentFragment4(void* other);
void* QTextDocumentFragment_NewQTextDocumentFragment();
void* QTextDocumentFragment_NewQTextDocumentFragment3(void* cursor);
void* QTextDocumentFragment_NewQTextDocumentFragment2(void* document);
int QTextDocumentFragment_IsEmpty(void* ptr);
char* QTextDocumentFragment_ToHtml(void* ptr, void* encoding);
char* QTextDocumentFragment_ToPlainText(void* ptr);
void QTextDocumentFragment_DestroyQTextDocumentFragment(void* ptr);

#ifdef __cplusplus
}
#endif