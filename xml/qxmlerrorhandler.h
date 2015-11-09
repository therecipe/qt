#ifdef __cplusplus
extern "C" {
#endif

int QXmlErrorHandler_Error(void* ptr, void* exception);
char* QXmlErrorHandler_ErrorString(void* ptr);
int QXmlErrorHandler_FatalError(void* ptr, void* exception);
int QXmlErrorHandler_Warning(void* ptr, void* exception);
void QXmlErrorHandler_DestroyQXmlErrorHandler(void* ptr);

#ifdef __cplusplus
}
#endif