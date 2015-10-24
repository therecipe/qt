#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QQmlProperty_NewQQmlProperty();
QtObjectPtr QQmlProperty_NewQQmlProperty2(QtObjectPtr obj);
QtObjectPtr QQmlProperty_NewQQmlProperty3(QtObjectPtr obj, QtObjectPtr ctxt);
QtObjectPtr QQmlProperty_NewQQmlProperty4(QtObjectPtr obj, QtObjectPtr engine);
QtObjectPtr QQmlProperty_NewQQmlProperty5(QtObjectPtr obj, char* name);
QtObjectPtr QQmlProperty_NewQQmlProperty6(QtObjectPtr obj, char* name, QtObjectPtr ctxt);
QtObjectPtr QQmlProperty_NewQQmlProperty7(QtObjectPtr obj, char* name, QtObjectPtr engine);
QtObjectPtr QQmlProperty_NewQQmlProperty8(QtObjectPtr other);
int QQmlProperty_ConnectNotifySignal(QtObjectPtr ptr, QtObjectPtr dest, char* slot);
int QQmlProperty_ConnectNotifySignal2(QtObjectPtr ptr, QtObjectPtr dest, int method);
int QQmlProperty_HasNotifySignal(QtObjectPtr ptr);
int QQmlProperty_Index(QtObjectPtr ptr);
int QQmlProperty_IsDesignable(QtObjectPtr ptr);
int QQmlProperty_IsProperty(QtObjectPtr ptr);
int QQmlProperty_IsResettable(QtObjectPtr ptr);
int QQmlProperty_IsSignalProperty(QtObjectPtr ptr);
int QQmlProperty_IsValid(QtObjectPtr ptr);
int QQmlProperty_IsWritable(QtObjectPtr ptr);
char* QQmlProperty_Name(QtObjectPtr ptr);
int QQmlProperty_NeedsNotifySignal(QtObjectPtr ptr);
QtObjectPtr QQmlProperty_Object(QtObjectPtr ptr);
int QQmlProperty_PropertyType(QtObjectPtr ptr);
int QQmlProperty_PropertyTypeCategory(QtObjectPtr ptr);
char* QQmlProperty_QQmlProperty_Read2(QtObjectPtr object, char* name);
char* QQmlProperty_QQmlProperty_Read3(QtObjectPtr object, char* name, QtObjectPtr ctxt);
char* QQmlProperty_QQmlProperty_Read4(QtObjectPtr object, char* name, QtObjectPtr engine);
char* QQmlProperty_Read(QtObjectPtr ptr);
int QQmlProperty_Reset(QtObjectPtr ptr);
int QQmlProperty_Type(QtObjectPtr ptr);
int QQmlProperty_QQmlProperty_Write2(QtObjectPtr object, char* name, char* value);
int QQmlProperty_QQmlProperty_Write3(QtObjectPtr object, char* name, char* value, QtObjectPtr ctxt);
int QQmlProperty_QQmlProperty_Write4(QtObjectPtr object, char* name, char* value, QtObjectPtr engine);
int QQmlProperty_Write(QtObjectPtr ptr, char* value);

#ifdef __cplusplus
}
#endif