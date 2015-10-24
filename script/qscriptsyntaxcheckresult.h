#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QScriptSyntaxCheckResult_NewQScriptSyntaxCheckResult(QtObjectPtr other);
int QScriptSyntaxCheckResult_ErrorColumnNumber(QtObjectPtr ptr);
int QScriptSyntaxCheckResult_ErrorLineNumber(QtObjectPtr ptr);
char* QScriptSyntaxCheckResult_ErrorMessage(QtObjectPtr ptr);
void QScriptSyntaxCheckResult_DestroyQScriptSyntaxCheckResult(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif