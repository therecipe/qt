#ifdef __cplusplus
extern "C" {
#endif

void* QScriptValue_NewQScriptValue();
void* QScriptValue_NewQScriptValue10(int value);
void* QScriptValue_NewQScriptValue11(int value);
void* QScriptValue_NewQScriptValue16(void* value);
void* QScriptValue_NewQScriptValue2(void* other);
void* QScriptValue_NewQScriptValue15(char* value);
void* QScriptValue_NewQScriptValue17(char* value);
void* QScriptValue_NewQScriptValue12(int value);
void* QScriptValue_Call2(void* ptr, void* thisObject, void* arguments);
void* QScriptValue_Construct2(void* ptr, void* arguments);
void* QScriptValue_Data(void* ptr);
void* QScriptValue_Engine(void* ptr);
int QScriptValue_Equals(void* ptr, void* other);
int QScriptValue_InstanceOf(void* ptr, void* other);
int QScriptValue_IsArray(void* ptr);
int QScriptValue_IsBool(void* ptr);
int QScriptValue_IsDate(void* ptr);
int QScriptValue_IsError(void* ptr);
int QScriptValue_IsFunction(void* ptr);
int QScriptValue_IsNull(void* ptr);
int QScriptValue_IsNumber(void* ptr);
int QScriptValue_IsObject(void* ptr);
int QScriptValue_IsQMetaObject(void* ptr);
int QScriptValue_IsQObject(void* ptr);
int QScriptValue_IsRegExp(void* ptr);
int QScriptValue_IsString(void* ptr);
int QScriptValue_IsUndefined(void* ptr);
int QScriptValue_IsValid(void* ptr);
int QScriptValue_IsVariant(void* ptr);
int QScriptValue_LessThan(void* ptr, void* other);
void* QScriptValue_Property2(void* ptr, void* name, int mode);
void* QScriptValue_Property(void* ptr, char* name, int mode);
int QScriptValue_PropertyFlags2(void* ptr, void* name, int mode);
int QScriptValue_PropertyFlags(void* ptr, char* name, int mode);
void* QScriptValue_Prototype(void* ptr);
void* QScriptValue_ScriptClass(void* ptr);
void QScriptValue_SetData(void* ptr, void* data);
void QScriptValue_SetProperty2(void* ptr, void* name, void* value, int flags);
void QScriptValue_SetProperty(void* ptr, char* name, void* value, int flags);
void QScriptValue_SetPrototype(void* ptr, void* prototype);
void QScriptValue_SetScriptClass(void* ptr, void* scriptClass);
int QScriptValue_StrictlyEquals(void* ptr, void* other);
int QScriptValue_ToBool(void* ptr);
void* QScriptValue_ToDateTime(void* ptr);
void* QScriptValue_ToQMetaObject(void* ptr);
void* QScriptValue_ToQObject(void* ptr);
void* QScriptValue_ToRegExp(void* ptr);
char* QScriptValue_ToString(void* ptr);
void* QScriptValue_ToVariant(void* ptr);
void QScriptValue_DestroyQScriptValue(void* ptr);

#ifdef __cplusplus
}
#endif