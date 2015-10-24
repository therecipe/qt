#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QScriptValue_NewQScriptValue();
QtObjectPtr QScriptValue_NewQScriptValue10(int value);
QtObjectPtr QScriptValue_NewQScriptValue11(int value);
QtObjectPtr QScriptValue_NewQScriptValue16(QtObjectPtr value);
QtObjectPtr QScriptValue_NewQScriptValue2(QtObjectPtr other);
QtObjectPtr QScriptValue_NewQScriptValue15(char* value);
QtObjectPtr QScriptValue_NewQScriptValue17(char* value);
QtObjectPtr QScriptValue_NewQScriptValue12(int value);
QtObjectPtr QScriptValue_Engine(QtObjectPtr ptr);
int QScriptValue_Equals(QtObjectPtr ptr, QtObjectPtr other);
int QScriptValue_InstanceOf(QtObjectPtr ptr, QtObjectPtr other);
int QScriptValue_IsArray(QtObjectPtr ptr);
int QScriptValue_IsBool(QtObjectPtr ptr);
int QScriptValue_IsDate(QtObjectPtr ptr);
int QScriptValue_IsError(QtObjectPtr ptr);
int QScriptValue_IsFunction(QtObjectPtr ptr);
int QScriptValue_IsNull(QtObjectPtr ptr);
int QScriptValue_IsNumber(QtObjectPtr ptr);
int QScriptValue_IsObject(QtObjectPtr ptr);
int QScriptValue_IsQMetaObject(QtObjectPtr ptr);
int QScriptValue_IsQObject(QtObjectPtr ptr);
int QScriptValue_IsRegExp(QtObjectPtr ptr);
int QScriptValue_IsString(QtObjectPtr ptr);
int QScriptValue_IsUndefined(QtObjectPtr ptr);
int QScriptValue_IsValid(QtObjectPtr ptr);
int QScriptValue_IsVariant(QtObjectPtr ptr);
int QScriptValue_LessThan(QtObjectPtr ptr, QtObjectPtr other);
int QScriptValue_PropertyFlags2(QtObjectPtr ptr, QtObjectPtr name, int mode);
int QScriptValue_PropertyFlags(QtObjectPtr ptr, char* name, int mode);
QtObjectPtr QScriptValue_ScriptClass(QtObjectPtr ptr);
void QScriptValue_SetData(QtObjectPtr ptr, QtObjectPtr data);
void QScriptValue_SetProperty2(QtObjectPtr ptr, QtObjectPtr name, QtObjectPtr value, int flags);
void QScriptValue_SetProperty(QtObjectPtr ptr, char* name, QtObjectPtr value, int flags);
void QScriptValue_SetPrototype(QtObjectPtr ptr, QtObjectPtr prototype);
void QScriptValue_SetScriptClass(QtObjectPtr ptr, QtObjectPtr scriptClass);
int QScriptValue_StrictlyEquals(QtObjectPtr ptr, QtObjectPtr other);
int QScriptValue_ToBool(QtObjectPtr ptr);
QtObjectPtr QScriptValue_ToQMetaObject(QtObjectPtr ptr);
QtObjectPtr QScriptValue_ToQObject(QtObjectPtr ptr);
char* QScriptValue_ToString(QtObjectPtr ptr);
char* QScriptValue_ToVariant(QtObjectPtr ptr);
void QScriptValue_DestroyQScriptValue(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif