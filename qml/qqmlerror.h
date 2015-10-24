#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QQmlError_NewQQmlError();
QtObjectPtr QQmlError_NewQQmlError2(QtObjectPtr other);
int QQmlError_Column(QtObjectPtr ptr);
char* QQmlError_Description(QtObjectPtr ptr);
int QQmlError_IsValid(QtObjectPtr ptr);
int QQmlError_Line(QtObjectPtr ptr);
QtObjectPtr QQmlError_Object(QtObjectPtr ptr);
void QQmlError_SetColumn(QtObjectPtr ptr, int column);
void QQmlError_SetDescription(QtObjectPtr ptr, char* description);
void QQmlError_SetLine(QtObjectPtr ptr, int line);
void QQmlError_SetObject(QtObjectPtr ptr, QtObjectPtr object);
void QQmlError_SetUrl(QtObjectPtr ptr, char* url);
char* QQmlError_ToString(QtObjectPtr ptr);
char* QQmlError_Url(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif