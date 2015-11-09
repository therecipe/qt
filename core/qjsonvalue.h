#ifdef __cplusplus
extern "C" {
#endif

void* QJsonValue_NewQJsonValue5(void* s);
void* QJsonValue_NewQJsonValue(int ty);
void* QJsonValue_NewQJsonValue2(int b);
void* QJsonValue_NewQJsonValue7(void* a);
void* QJsonValue_NewQJsonValue8(void* o);
void* QJsonValue_NewQJsonValue9(void* other);
void* QJsonValue_NewQJsonValue4(char* s);
void* QJsonValue_NewQJsonValue6(char* s);
void* QJsonValue_NewQJsonValue12(int n);
int QJsonValue_IsArray(void* ptr);
int QJsonValue_IsBool(void* ptr);
int QJsonValue_IsDouble(void* ptr);
int QJsonValue_IsNull(void* ptr);
int QJsonValue_IsObject(void* ptr);
int QJsonValue_IsString(void* ptr);
int QJsonValue_IsUndefined(void* ptr);
void* QJsonValue_ToArray2(void* ptr);
void* QJsonValue_ToArray(void* ptr, void* defaultValue);
int QJsonValue_ToBool(void* ptr, int defaultValue);
int QJsonValue_ToInt(void* ptr, int defaultValue);
void* QJsonValue_ToObject2(void* ptr);
void* QJsonValue_ToObject(void* ptr, void* defaultValue);
char* QJsonValue_ToString(void* ptr, char* defaultValue);
void* QJsonValue_ToVariant(void* ptr);
int QJsonValue_Type(void* ptr);
void QJsonValue_DestroyQJsonValue(void* ptr);

#ifdef __cplusplus
}
#endif