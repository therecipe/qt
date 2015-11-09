#ifdef __cplusplus
extern "C" {
#endif

int QUuid_Variant(void* ptr);
int QUuid_Version(void* ptr);
void* QUuid_NewQUuid();
void* QUuid_NewQUuid5(void* text);
void* QUuid_NewQUuid3(char* text);
int QUuid_IsNull(void* ptr);
void* QUuid_ToByteArray(void* ptr);
void* QUuid_ToRfc4122(void* ptr);
char* QUuid_ToString(void* ptr);

#ifdef __cplusplus
}
#endif