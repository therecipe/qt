#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QMetaObject_QMetaObject_ConnectSlotsByName(QtObjectPtr object);
int QMetaObject_QMetaObject_CheckConnectArgs2(QtObjectPtr signal, QtObjectPtr method);
int QMetaObject_QMetaObject_CheckConnectArgs(char* signal, char* method);
int QMetaObject_ClassInfoCount(QtObjectPtr ptr);
int QMetaObject_ClassInfoOffset(QtObjectPtr ptr);
int QMetaObject_ConstructorCount(QtObjectPtr ptr);
int QMetaObject_EnumeratorCount(QtObjectPtr ptr);
int QMetaObject_EnumeratorOffset(QtObjectPtr ptr);
int QMetaObject_IndexOfClassInfo(QtObjectPtr ptr, char* name);
int QMetaObject_IndexOfConstructor(QtObjectPtr ptr, char* constructor);
int QMetaObject_IndexOfEnumerator(QtObjectPtr ptr, char* name);
int QMetaObject_IndexOfMethod(QtObjectPtr ptr, char* method);
int QMetaObject_IndexOfProperty(QtObjectPtr ptr, char* name);
int QMetaObject_IndexOfSignal(QtObjectPtr ptr, char* signal);
int QMetaObject_IndexOfSlot(QtObjectPtr ptr, char* slot);
int QMetaObject_QMetaObject_InvokeMethod4(QtObjectPtr obj, char* member, QtObjectPtr val0, QtObjectPtr val1, QtObjectPtr val2, QtObjectPtr val3, QtObjectPtr val4, QtObjectPtr val5, QtObjectPtr val6, QtObjectPtr val7, QtObjectPtr val8, QtObjectPtr val9);
int QMetaObject_QMetaObject_InvokeMethod2(QtObjectPtr obj, char* member, QtObjectPtr ret, QtObjectPtr val0, QtObjectPtr val1, QtObjectPtr val2, QtObjectPtr val3, QtObjectPtr val4, QtObjectPtr val5, QtObjectPtr val6, QtObjectPtr val7, QtObjectPtr val8, QtObjectPtr val9);
int QMetaObject_QMetaObject_InvokeMethod3(QtObjectPtr obj, char* member, int ty, QtObjectPtr val0, QtObjectPtr val1, QtObjectPtr val2, QtObjectPtr val3, QtObjectPtr val4, QtObjectPtr val5, QtObjectPtr val6, QtObjectPtr val7, QtObjectPtr val8, QtObjectPtr val9);
int QMetaObject_QMetaObject_InvokeMethod(QtObjectPtr obj, char* member, int ty, QtObjectPtr ret, QtObjectPtr val0, QtObjectPtr val1, QtObjectPtr val2, QtObjectPtr val3, QtObjectPtr val4, QtObjectPtr val5, QtObjectPtr val6, QtObjectPtr val7, QtObjectPtr val8, QtObjectPtr val9);
int QMetaObject_MethodCount(QtObjectPtr ptr);
int QMetaObject_MethodOffset(QtObjectPtr ptr);
QtObjectPtr QMetaObject_NewInstance(QtObjectPtr ptr, QtObjectPtr val0, QtObjectPtr val1, QtObjectPtr val2, QtObjectPtr val3, QtObjectPtr val4, QtObjectPtr val5, QtObjectPtr val6, QtObjectPtr val7, QtObjectPtr val8, QtObjectPtr val9);
int QMetaObject_PropertyCount(QtObjectPtr ptr);
int QMetaObject_PropertyOffset(QtObjectPtr ptr);
QtObjectPtr QMetaObject_SuperClass(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif