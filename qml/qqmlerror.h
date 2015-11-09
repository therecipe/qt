#ifdef __cplusplus
extern "C" {
#endif

void* QQmlError_NewQQmlError();
void* QQmlError_NewQQmlError2(void* other);
int QQmlError_Column(void* ptr);
char* QQmlError_Description(void* ptr);
int QQmlError_IsValid(void* ptr);
int QQmlError_Line(void* ptr);
void* QQmlError_Object(void* ptr);
void QQmlError_SetColumn(void* ptr, int column);
void QQmlError_SetDescription(void* ptr, char* description);
void QQmlError_SetLine(void* ptr, int line);
void QQmlError_SetObject(void* ptr, void* object);
void QQmlError_SetUrl(void* ptr, void* url);
char* QQmlError_ToString(void* ptr);

#ifdef __cplusplus
}
#endif