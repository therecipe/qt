#ifdef __cplusplus
extern "C" {
#endif

int QQmlNdefRecord_TypeNameFormat(void* ptr);
void* QQmlNdefRecord_NewQQmlNdefRecord(void* parent);
void* QQmlNdefRecord_NewQQmlNdefRecord2(void* record, void* parent);
void QQmlNdefRecord_ConnectRecordChanged(void* ptr);
void QQmlNdefRecord_DisconnectRecordChanged(void* ptr);
void QQmlNdefRecord_SetRecord(void* ptr, void* record);
void QQmlNdefRecord_SetType(void* ptr, char* newtype);
void QQmlNdefRecord_SetTypeNameFormat(void* ptr, int newTypeNameFormat);
char* QQmlNdefRecord_Type(void* ptr);
void QQmlNdefRecord_ConnectTypeChanged(void* ptr);
void QQmlNdefRecord_DisconnectTypeChanged(void* ptr);
void QQmlNdefRecord_ConnectTypeNameFormatChanged(void* ptr);
void QQmlNdefRecord_DisconnectTypeNameFormatChanged(void* ptr);

#ifdef __cplusplus
}
#endif