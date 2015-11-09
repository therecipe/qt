#ifdef __cplusplus
extern "C" {
#endif

void* QTextListFormat_NewQTextListFormat();
int QTextListFormat_Indent(void* ptr);
int QTextListFormat_IsValid(void* ptr);
char* QTextListFormat_NumberPrefix(void* ptr);
char* QTextListFormat_NumberSuffix(void* ptr);
void QTextListFormat_SetIndent(void* ptr, int indentation);
void QTextListFormat_SetNumberPrefix(void* ptr, char* numberPrefix);
void QTextListFormat_SetNumberSuffix(void* ptr, char* numberSuffix);
void QTextListFormat_SetStyle(void* ptr, int style);
int QTextListFormat_Style(void* ptr);

#ifdef __cplusplus
}
#endif