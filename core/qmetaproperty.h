#ifdef __cplusplus
extern "C" {
#endif

int QMetaProperty_HasNotifySignal(void* ptr);
int QMetaProperty_IsConstant(void* ptr);
int QMetaProperty_IsDesignable(void* ptr, void* object);
int QMetaProperty_IsEnumType(void* ptr);
int QMetaProperty_IsFinal(void* ptr);
int QMetaProperty_IsFlagType(void* ptr);
int QMetaProperty_IsReadable(void* ptr);
int QMetaProperty_IsResettable(void* ptr);
int QMetaProperty_IsScriptable(void* ptr, void* object);
int QMetaProperty_IsStored(void* ptr, void* object);
int QMetaProperty_IsUser(void* ptr, void* object);
int QMetaProperty_IsValid(void* ptr);
int QMetaProperty_IsWritable(void* ptr);
int QMetaProperty_NotifySignalIndex(void* ptr);
int QMetaProperty_PropertyIndex(void* ptr);
void* QMetaProperty_Read(void* ptr, void* object);
int QMetaProperty_Reset(void* ptr, void* object);
int QMetaProperty_Revision(void* ptr);
int QMetaProperty_UserType(void* ptr);
int QMetaProperty_Write(void* ptr, void* object, void* value);

#ifdef __cplusplus
}
#endif