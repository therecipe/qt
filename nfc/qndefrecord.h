#ifdef __cplusplus
extern "C" {
#endif

void* QNdefRecord_NewQNdefRecord();
void* QNdefRecord_NewQNdefRecord2(void* other);
void* QNdefRecord_Id(void* ptr);
int QNdefRecord_IsEmpty(void* ptr);
void* QNdefRecord_Payload(void* ptr);
void QNdefRecord_SetId(void* ptr, void* id);
void QNdefRecord_SetPayload(void* ptr, void* payload);
void QNdefRecord_SetType(void* ptr, void* ty);
void QNdefRecord_SetTypeNameFormat(void* ptr, int typeNameFormat);
void* QNdefRecord_Type(void* ptr);
int QNdefRecord_TypeNameFormat(void* ptr);
void QNdefRecord_DestroyQNdefRecord(void* ptr);

#ifdef __cplusplus
}
#endif