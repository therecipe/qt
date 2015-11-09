#ifdef __cplusplus
extern "C" {
#endif

void* QScriptContextInfo_NewQScriptContextInfo3();
void* QScriptContextInfo_NewQScriptContextInfo(void* context);
void* QScriptContextInfo_NewQScriptContextInfo2(void* other);
char* QScriptContextInfo_FileName(void* ptr);
int QScriptContextInfo_FunctionEndLineNumber(void* ptr);
int QScriptContextInfo_FunctionMetaIndex(void* ptr);
char* QScriptContextInfo_FunctionName(void* ptr);
char* QScriptContextInfo_FunctionParameterNames(void* ptr);
int QScriptContextInfo_FunctionStartLineNumber(void* ptr);
int QScriptContextInfo_FunctionType(void* ptr);
int QScriptContextInfo_IsNull(void* ptr);
int QScriptContextInfo_LineNumber(void* ptr);
void QScriptContextInfo_DestroyQScriptContextInfo(void* ptr);

#ifdef __cplusplus
}
#endif