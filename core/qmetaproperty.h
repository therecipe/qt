#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QMetaProperty_HasNotifySignal(QtObjectPtr ptr);
int QMetaProperty_IsConstant(QtObjectPtr ptr);
int QMetaProperty_IsDesignable(QtObjectPtr ptr, QtObjectPtr object);
int QMetaProperty_IsEnumType(QtObjectPtr ptr);
int QMetaProperty_IsFinal(QtObjectPtr ptr);
int QMetaProperty_IsFlagType(QtObjectPtr ptr);
int QMetaProperty_IsReadable(QtObjectPtr ptr);
int QMetaProperty_IsResettable(QtObjectPtr ptr);
int QMetaProperty_IsScriptable(QtObjectPtr ptr, QtObjectPtr object);
int QMetaProperty_IsStored(QtObjectPtr ptr, QtObjectPtr object);
int QMetaProperty_IsUser(QtObjectPtr ptr, QtObjectPtr object);
int QMetaProperty_IsValid(QtObjectPtr ptr);
int QMetaProperty_IsWritable(QtObjectPtr ptr);
int QMetaProperty_NotifySignalIndex(QtObjectPtr ptr);
int QMetaProperty_PropertyIndex(QtObjectPtr ptr);
char* QMetaProperty_Read(QtObjectPtr ptr, QtObjectPtr object);
int QMetaProperty_Reset(QtObjectPtr ptr, QtObjectPtr object);
int QMetaProperty_Revision(QtObjectPtr ptr);
int QMetaProperty_UserType(QtObjectPtr ptr);
int QMetaProperty_Write(QtObjectPtr ptr, QtObjectPtr object, char* value);

#ifdef __cplusplus
}
#endif