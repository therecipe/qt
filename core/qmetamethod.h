#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QMetaMethod_Access(QtObjectPtr ptr);
int QMetaMethod_Invoke4(QtObjectPtr ptr, QtObjectPtr object, QtObjectPtr val0, QtObjectPtr val1, QtObjectPtr val2, QtObjectPtr val3, QtObjectPtr val4, QtObjectPtr val5, QtObjectPtr val6, QtObjectPtr val7, QtObjectPtr val8, QtObjectPtr val9);
int QMetaMethod_Invoke2(QtObjectPtr ptr, QtObjectPtr object, QtObjectPtr returnValue, QtObjectPtr val0, QtObjectPtr val1, QtObjectPtr val2, QtObjectPtr val3, QtObjectPtr val4, QtObjectPtr val5, QtObjectPtr val6, QtObjectPtr val7, QtObjectPtr val8, QtObjectPtr val9);
int QMetaMethod_Invoke3(QtObjectPtr ptr, QtObjectPtr object, int connectionType, QtObjectPtr val0, QtObjectPtr val1, QtObjectPtr val2, QtObjectPtr val3, QtObjectPtr val4, QtObjectPtr val5, QtObjectPtr val6, QtObjectPtr val7, QtObjectPtr val8, QtObjectPtr val9);
int QMetaMethod_Invoke(QtObjectPtr ptr, QtObjectPtr object, int connectionType, QtObjectPtr returnValue, QtObjectPtr val0, QtObjectPtr val1, QtObjectPtr val2, QtObjectPtr val3, QtObjectPtr val4, QtObjectPtr val5, QtObjectPtr val6, QtObjectPtr val7, QtObjectPtr val8, QtObjectPtr val9);
int QMetaMethod_IsValid(QtObjectPtr ptr);
int QMetaMethod_MethodIndex(QtObjectPtr ptr);
int QMetaMethod_MethodType(QtObjectPtr ptr);
int QMetaMethod_ParameterCount(QtObjectPtr ptr);
int QMetaMethod_ParameterType(QtObjectPtr ptr, int index);
int QMetaMethod_ReturnType(QtObjectPtr ptr);
int QMetaMethod_Revision(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif