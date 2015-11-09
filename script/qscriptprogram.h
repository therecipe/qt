#ifdef __cplusplus
extern "C" {
#endif

void* QScriptProgram_NewQScriptProgram();
void* QScriptProgram_NewQScriptProgram3(void* other);
void* QScriptProgram_NewQScriptProgram2(char* sourceCode, char* fileName, int firstLineNumber);
char* QScriptProgram_FileName(void* ptr);
int QScriptProgram_FirstLineNumber(void* ptr);
int QScriptProgram_IsNull(void* ptr);
char* QScriptProgram_SourceCode(void* ptr);
void QScriptProgram_DestroyQScriptProgram(void* ptr);

#ifdef __cplusplus
}
#endif