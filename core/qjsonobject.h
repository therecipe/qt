#ifdef __cplusplus
extern "C" {
#endif

int QJsonObject_Contains(void* ptr, char* key);
int QJsonObject_Count(void* ptr);
int QJsonObject_Empty(void* ptr);
int QJsonObject_IsEmpty(void* ptr);
char* QJsonObject_Keys(void* ptr);
int QJsonObject_Length(void* ptr);
int QJsonObject_Size(void* ptr);
void QJsonObject_DestroyQJsonObject(void* ptr);

#ifdef __cplusplus
}
#endif