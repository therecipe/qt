#ifdef __cplusplus
extern "C" {
#endif

void* QScriptSyntaxCheckResult_NewQScriptSyntaxCheckResult(void* other);
int QScriptSyntaxCheckResult_ErrorColumnNumber(void* ptr);
int QScriptSyntaxCheckResult_ErrorLineNumber(void* ptr);
char* QScriptSyntaxCheckResult_ErrorMessage(void* ptr);
void QScriptSyntaxCheckResult_DestroyQScriptSyntaxCheckResult(void* ptr);

#ifdef __cplusplus
}
#endif