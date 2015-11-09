#ifdef __cplusplus
extern "C" {
#endif

int QMetaMethod_Access(void* ptr);
int QMetaMethod_Invoke4(void* ptr, void* object, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9);
int QMetaMethod_Invoke2(void* ptr, void* object, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9);
int QMetaMethod_Invoke3(void* ptr, void* object, int connectionType, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9);
int QMetaMethod_Invoke(void* ptr, void* object, int connectionType, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9);
int QMetaMethod_IsValid(void* ptr);
int QMetaMethod_MethodIndex(void* ptr);
void* QMetaMethod_MethodSignature(void* ptr);
int QMetaMethod_MethodType(void* ptr);
void* QMetaMethod_Name(void* ptr);
int QMetaMethod_ParameterCount(void* ptr);
int QMetaMethod_ParameterType(void* ptr, int index);
int QMetaMethod_ReturnType(void* ptr);
int QMetaMethod_Revision(void* ptr);

#ifdef __cplusplus
}
#endif