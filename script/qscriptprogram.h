#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QScriptProgram_NewQScriptProgram();
QtObjectPtr QScriptProgram_NewQScriptProgram3(QtObjectPtr other);
QtObjectPtr QScriptProgram_NewQScriptProgram2(char* sourceCode, char* fileName, int firstLineNumber);
char* QScriptProgram_FileName(QtObjectPtr ptr);
int QScriptProgram_FirstLineNumber(QtObjectPtr ptr);
int QScriptProgram_IsNull(QtObjectPtr ptr);
char* QScriptProgram_SourceCode(QtObjectPtr ptr);
void QScriptProgram_DestroyQScriptProgram(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif