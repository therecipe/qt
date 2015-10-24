#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QXmlErrorHandler_Error(QtObjectPtr ptr, QtObjectPtr exception);
char* QXmlErrorHandler_ErrorString(QtObjectPtr ptr);
int QXmlErrorHandler_FatalError(QtObjectPtr ptr, QtObjectPtr exception);
int QXmlErrorHandler_Warning(QtObjectPtr ptr, QtObjectPtr exception);
void QXmlErrorHandler_DestroyQXmlErrorHandler(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif