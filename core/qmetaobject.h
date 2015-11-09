#ifdef __cplusplus
extern "C" {
#endif

void QMetaObject_QMetaObject_ConnectSlotsByName(void* object);
int QMetaObject_QMetaObject_CheckConnectArgs2(void* signal, void* method);
int QMetaObject_QMetaObject_CheckConnectArgs(char* signal, char* method);
int QMetaObject_ClassInfoCount(void* ptr);
int QMetaObject_ClassInfoOffset(void* ptr);
int QMetaObject_ConstructorCount(void* ptr);
int QMetaObject_EnumeratorCount(void* ptr);
int QMetaObject_EnumeratorOffset(void* ptr);
int QMetaObject_IndexOfClassInfo(void* ptr, char* name);
int QMetaObject_IndexOfConstructor(void* ptr, char* constructor);
int QMetaObject_IndexOfEnumerator(void* ptr, char* name);
int QMetaObject_IndexOfMethod(void* ptr, char* method);
int QMetaObject_IndexOfProperty(void* ptr, char* name);
int QMetaObject_IndexOfSignal(void* ptr, char* signal);
int QMetaObject_IndexOfSlot(void* ptr, char* slot);
int QMetaObject_QMetaObject_InvokeMethod4(void* obj, char* member, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9);
int QMetaObject_QMetaObject_InvokeMethod2(void* obj, char* member, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9);
int QMetaObject_QMetaObject_InvokeMethod3(void* obj, char* member, int ty, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9);
int QMetaObject_QMetaObject_InvokeMethod(void* obj, char* member, int ty, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9);
int QMetaObject_MethodCount(void* ptr);
int QMetaObject_MethodOffset(void* ptr);
void* QMetaObject_NewInstance(void* ptr, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9);
void* QMetaObject_QMetaObject_NormalizedSignature(char* method);
void* QMetaObject_QMetaObject_NormalizedType(char* ty);
int QMetaObject_PropertyCount(void* ptr);
int QMetaObject_PropertyOffset(void* ptr);
void* QMetaObject_SuperClass(void* ptr);

#ifdef __cplusplus
}
#endif