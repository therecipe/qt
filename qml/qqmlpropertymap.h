#ifdef __cplusplus
extern "C" {
#endif

void* QQmlPropertyMap_NewQQmlPropertyMap(void* parent);
void QQmlPropertyMap_Clear(void* ptr, char* key);
int QQmlPropertyMap_Contains(void* ptr, char* key);
int QQmlPropertyMap_Count(void* ptr);
int QQmlPropertyMap_IsEmpty(void* ptr);
char* QQmlPropertyMap_Keys(void* ptr);
int QQmlPropertyMap_Size(void* ptr);
void* QQmlPropertyMap_Value(void* ptr, char* key);
void QQmlPropertyMap_ConnectValueChanged(void* ptr);
void QQmlPropertyMap_DisconnectValueChanged(void* ptr);
void QQmlPropertyMap_DestroyQQmlPropertyMap(void* ptr);

#ifdef __cplusplus
}
#endif