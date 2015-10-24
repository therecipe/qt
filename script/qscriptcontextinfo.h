#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QScriptContextInfo_NewQScriptContextInfo3();
QtObjectPtr QScriptContextInfo_NewQScriptContextInfo(QtObjectPtr context);
QtObjectPtr QScriptContextInfo_NewQScriptContextInfo2(QtObjectPtr other);
char* QScriptContextInfo_FileName(QtObjectPtr ptr);
int QScriptContextInfo_FunctionEndLineNumber(QtObjectPtr ptr);
int QScriptContextInfo_FunctionMetaIndex(QtObjectPtr ptr);
char* QScriptContextInfo_FunctionName(QtObjectPtr ptr);
char* QScriptContextInfo_FunctionParameterNames(QtObjectPtr ptr);
int QScriptContextInfo_FunctionStartLineNumber(QtObjectPtr ptr);
int QScriptContextInfo_FunctionType(QtObjectPtr ptr);
int QScriptContextInfo_IsNull(QtObjectPtr ptr);
int QScriptContextInfo_LineNumber(QtObjectPtr ptr);
void QScriptContextInfo_DestroyQScriptContextInfo(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif