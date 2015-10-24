#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QQmlPropertyMap_NewQQmlPropertyMap(QtObjectPtr parent);
void QQmlPropertyMap_Clear(QtObjectPtr ptr, char* key);
int QQmlPropertyMap_Contains(QtObjectPtr ptr, char* key);
int QQmlPropertyMap_Count(QtObjectPtr ptr);
int QQmlPropertyMap_IsEmpty(QtObjectPtr ptr);
char* QQmlPropertyMap_Keys(QtObjectPtr ptr);
int QQmlPropertyMap_Size(QtObjectPtr ptr);
char* QQmlPropertyMap_Value(QtObjectPtr ptr, char* key);
void QQmlPropertyMap_ConnectValueChanged(QtObjectPtr ptr);
void QQmlPropertyMap_DisconnectValueChanged(QtObjectPtr ptr);
void QQmlPropertyMap_DestroyQQmlPropertyMap(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif