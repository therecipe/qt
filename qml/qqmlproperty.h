#ifdef __cplusplus
extern "C" {
#endif

void* QQmlProperty_NewQQmlProperty();
void* QQmlProperty_NewQQmlProperty2(void* obj);
void* QQmlProperty_NewQQmlProperty3(void* obj, void* ctxt);
void* QQmlProperty_NewQQmlProperty4(void* obj, void* engine);
void* QQmlProperty_NewQQmlProperty5(void* obj, char* name);
void* QQmlProperty_NewQQmlProperty6(void* obj, char* name, void* ctxt);
void* QQmlProperty_NewQQmlProperty7(void* obj, char* name, void* engine);
void* QQmlProperty_NewQQmlProperty8(void* other);
int QQmlProperty_ConnectNotifySignal(void* ptr, void* dest, char* slot);
int QQmlProperty_ConnectNotifySignal2(void* ptr, void* dest, int method);
int QQmlProperty_HasNotifySignal(void* ptr);
int QQmlProperty_Index(void* ptr);
int QQmlProperty_IsDesignable(void* ptr);
int QQmlProperty_IsProperty(void* ptr);
int QQmlProperty_IsResettable(void* ptr);
int QQmlProperty_IsSignalProperty(void* ptr);
int QQmlProperty_IsValid(void* ptr);
int QQmlProperty_IsWritable(void* ptr);
char* QQmlProperty_Name(void* ptr);
int QQmlProperty_NeedsNotifySignal(void* ptr);
void* QQmlProperty_Object(void* ptr);
int QQmlProperty_PropertyType(void* ptr);
int QQmlProperty_PropertyTypeCategory(void* ptr);
void* QQmlProperty_QQmlProperty_Read2(void* object, char* name);
void* QQmlProperty_QQmlProperty_Read3(void* object, char* name, void* ctxt);
void* QQmlProperty_QQmlProperty_Read4(void* object, char* name, void* engine);
void* QQmlProperty_Read(void* ptr);
int QQmlProperty_Reset(void* ptr);
int QQmlProperty_Type(void* ptr);
int QQmlProperty_QQmlProperty_Write2(void* object, char* name, void* value);
int QQmlProperty_QQmlProperty_Write3(void* object, char* name, void* value, void* ctxt);
int QQmlProperty_QQmlProperty_Write4(void* object, char* name, void* value, void* engine);
int QQmlProperty_Write(void* ptr, void* value);

#ifdef __cplusplus
}
#endif