#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QJSValue_NewQJSValue3(QtObjectPtr other);
QtObjectPtr QJSValue_NewQJSValue(int value);
QtObjectPtr QJSValue_NewQJSValue4(int value);
QtObjectPtr QJSValue_NewQJSValue2(QtObjectPtr other);
QtObjectPtr QJSValue_NewQJSValue9(QtObjectPtr value);
QtObjectPtr QJSValue_NewQJSValue8(char* value);
QtObjectPtr QJSValue_NewQJSValue10(char* value);
QtObjectPtr QJSValue_NewQJSValue5(int value);
int QJSValue_DeleteProperty(QtObjectPtr ptr, char* name);
int QJSValue_Equals(QtObjectPtr ptr, QtObjectPtr other);
int QJSValue_HasOwnProperty(QtObjectPtr ptr, char* name);
int QJSValue_HasProperty(QtObjectPtr ptr, char* name);
int QJSValue_IsArray(QtObjectPtr ptr);
int QJSValue_IsBool(QtObjectPtr ptr);
int QJSValue_IsCallable(QtObjectPtr ptr);
int QJSValue_IsDate(QtObjectPtr ptr);
int QJSValue_IsError(QtObjectPtr ptr);
int QJSValue_IsNull(QtObjectPtr ptr);
int QJSValue_IsNumber(QtObjectPtr ptr);
int QJSValue_IsObject(QtObjectPtr ptr);
int QJSValue_IsQObject(QtObjectPtr ptr);
int QJSValue_IsRegExp(QtObjectPtr ptr);
int QJSValue_IsString(QtObjectPtr ptr);
int QJSValue_IsUndefined(QtObjectPtr ptr);
int QJSValue_IsVariant(QtObjectPtr ptr);
void QJSValue_SetProperty(QtObjectPtr ptr, char* name, QtObjectPtr value);
void QJSValue_SetPrototype(QtObjectPtr ptr, QtObjectPtr prototype);
int QJSValue_StrictlyEquals(QtObjectPtr ptr, QtObjectPtr other);
int QJSValue_ToBool(QtObjectPtr ptr);
QtObjectPtr QJSValue_ToQObject(QtObjectPtr ptr);
char* QJSValue_ToString(QtObjectPtr ptr);
char* QJSValue_ToVariant(QtObjectPtr ptr);
void QJSValue_DestroyQJSValue(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif