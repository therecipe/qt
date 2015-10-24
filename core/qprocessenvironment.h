#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QProcessEnvironment_NewQProcessEnvironment();
QtObjectPtr QProcessEnvironment_NewQProcessEnvironment2(QtObjectPtr other);
void QProcessEnvironment_Clear(QtObjectPtr ptr);
int QProcessEnvironment_Contains(QtObjectPtr ptr, char* name);
int QProcessEnvironment_IsEmpty(QtObjectPtr ptr);
char* QProcessEnvironment_Keys(QtObjectPtr ptr);
void QProcessEnvironment_Swap(QtObjectPtr ptr, QtObjectPtr other);
char* QProcessEnvironment_ToStringList(QtObjectPtr ptr);
char* QProcessEnvironment_Value(QtObjectPtr ptr, char* name, char* defaultValue);
void QProcessEnvironment_DestroyQProcessEnvironment(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif