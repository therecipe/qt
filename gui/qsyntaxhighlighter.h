#ifdef __cplusplus
extern "C" {
#endif

void* QSyntaxHighlighter_Document(void* ptr);
void QSyntaxHighlighter_Rehighlight(void* ptr);
void QSyntaxHighlighter_RehighlightBlock(void* ptr, void* block);
void QSyntaxHighlighter_SetDocument(void* ptr, void* doc);
void QSyntaxHighlighter_DestroyQSyntaxHighlighter(void* ptr);

#ifdef __cplusplus
}
#endif