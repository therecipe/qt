#ifdef __cplusplus
extern "C" {
#endif

void* QJsonDocument_NewQJsonDocument();
void* QJsonDocument_NewQJsonDocument3(void* array);
void* QJsonDocument_NewQJsonDocument4(void* other);
void* QJsonDocument_NewQJsonDocument2(void* object);
void* QJsonDocument_Array(void* ptr);
void* QJsonDocument_QJsonDocument_FromBinaryData(void* data, int validation);
void* QJsonDocument_QJsonDocument_FromJson(void* json, void* error);
void* QJsonDocument_QJsonDocument_FromRawData(char* data, int size, int validation);
void* QJsonDocument_QJsonDocument_FromVariant(void* variant);
int QJsonDocument_IsArray(void* ptr);
int QJsonDocument_IsEmpty(void* ptr);
int QJsonDocument_IsNull(void* ptr);
int QJsonDocument_IsObject(void* ptr);
void* QJsonDocument_Object(void* ptr);
void QJsonDocument_SetArray(void* ptr, void* array);
void QJsonDocument_SetObject(void* ptr, void* object);
void* QJsonDocument_ToBinaryData(void* ptr);
void* QJsonDocument_ToJson(void* ptr, int format);
void* QJsonDocument_ToVariant(void* ptr);
void QJsonDocument_DestroyQJsonDocument(void* ptr);

#ifdef __cplusplus
}
#endif