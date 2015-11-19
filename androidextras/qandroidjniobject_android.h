#ifdef __cplusplus
extern "C" {
#endif

void* QAndroidJniObject_NewQAndroidJniObject();
void* QAndroidJniObject_NewQAndroidJniObject2(char* className);
void* QAndroidJniObject_NewQAndroidJniObject3(char* className, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void* QAndroidJniObject_NewQAndroidJniObject4(void* clazz);
void* QAndroidJniObject_NewQAndroidJniObject5(void* clazz, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void* QAndroidJniObject_NewQAndroidJniObject6(void* object);
int QAndroidJniObject_CallMethodInt(void* ptr, char* methodName);
int QAndroidJniObject_CallMethodBoolean(void* ptr, char* methodName);
void QAndroidJniObject_CallMethodVoid(void* ptr, char* methodName);
int QAndroidJniObject_CallMethodInt2(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
int QAndroidJniObject_CallMethodBoolean2(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void QAndroidJniObject_CallMethodVoid2(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void* QAndroidJniObject_CallObjectMethod(void* ptr, char* methodName);
void* QAndroidJniObject_CallObjectMethod2(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
int QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt(char* className, char* methodName);
int QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean(char* className, char* methodName);
void QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid(char* className, char* methodName);
int QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt2(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
int QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean2(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid2(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
int QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt3(void* clazz, char* methodName);
int QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean3(void* clazz, char* methodName);
void QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid3(void* clazz, char* methodName);
int QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt4(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
int QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean4(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid4(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void* QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod(char* className, char* methodName);
void* QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod2(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void* QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod3(void* clazz, char* methodName);
void* QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod4(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v);
void* QAndroidJniObject_QAndroidJniObject_FromString(char* stri);
int QAndroidJniObject_GetFieldInt(void* ptr, char* fieldName);
int QAndroidJniObject_GetFieldBoolean(void* ptr, char* fieldName);
void* QAndroidJniObject_GetObjectField(void* ptr, char* fieldName);
void* QAndroidJniObject_GetObjectField2(void* ptr, char* fieldName, char* signature);
int QAndroidJniObject_QAndroidJniObject_GetStaticFieldInt(char* className, char* fieldName);
int QAndroidJniObject_QAndroidJniObject_GetStaticFieldBoolean(char* className, char* fieldName);
int QAndroidJniObject_QAndroidJniObject_GetStaticFieldInt2(void* clazz, char* fieldName);
int QAndroidJniObject_QAndroidJniObject_GetStaticFieldBoolean2(void* clazz, char* fieldName);
void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectField(char* className, char* fieldName);
void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectField2(char* className, char* fieldName, char* signature);
void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectField3(void* clazz, char* fieldName);
void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectField4(void* clazz, char* fieldName, char* signature);
int QAndroidJniObject_QAndroidJniObject_IsClassAvailable(char* className);
int QAndroidJniObject_IsValid(void* ptr);
char* QAndroidJniObject_ToString(void* ptr);
void* QAndroidJniObject_Object(void* ptr);
void QAndroidJniObject_DestroyQAndroidJniObject(void* ptr);

#ifdef __cplusplus
}
#endif