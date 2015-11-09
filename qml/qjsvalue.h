#ifdef __cplusplus
extern "C" {
#endif

void* QJSValue_NewQJSValue3(void* other);
void* QJSValue_NewQJSValue(int value);
void* QJSValue_NewQJSValue4(int value);
void* QJSValue_NewQJSValue2(void* other);
void* QJSValue_NewQJSValue9(void* value);
void* QJSValue_NewQJSValue8(char* value);
void* QJSValue_NewQJSValue10(char* value);
void* QJSValue_NewQJSValue5(int value);
int QJSValue_DeleteProperty(void* ptr, char* name);
int QJSValue_Equals(void* ptr, void* other);
int QJSValue_HasOwnProperty(void* ptr, char* name);
int QJSValue_HasProperty(void* ptr, char* name);
int QJSValue_IsArray(void* ptr);
int QJSValue_IsBool(void* ptr);
int QJSValue_IsCallable(void* ptr);
int QJSValue_IsDate(void* ptr);
int QJSValue_IsError(void* ptr);
int QJSValue_IsNull(void* ptr);
int QJSValue_IsNumber(void* ptr);
int QJSValue_IsObject(void* ptr);
int QJSValue_IsQObject(void* ptr);
int QJSValue_IsRegExp(void* ptr);
int QJSValue_IsString(void* ptr);
int QJSValue_IsUndefined(void* ptr);
int QJSValue_IsVariant(void* ptr);
void* QJSValue_Property(void* ptr, char* name);
void* QJSValue_Prototype(void* ptr);
void QJSValue_SetProperty(void* ptr, char* name, void* value);
void QJSValue_SetPrototype(void* ptr, void* prototype);
int QJSValue_StrictlyEquals(void* ptr, void* other);
int QJSValue_ToBool(void* ptr);
void* QJSValue_ToDateTime(void* ptr);
void* QJSValue_ToQObject(void* ptr);
char* QJSValue_ToString(void* ptr);
void* QJSValue_ToVariant(void* ptr);
void QJSValue_DestroyQJSValue(void* ptr);

#ifdef __cplusplus
}
#endif