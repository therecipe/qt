#ifdef __cplusplus
extern "C" {
#endif

void* QDnsServiceRecord_NewQDnsServiceRecord();
void* QDnsServiceRecord_NewQDnsServiceRecord2(void* other);
char* QDnsServiceRecord_Name(void* ptr);
void QDnsServiceRecord_Swap(void* ptr, void* other);
char* QDnsServiceRecord_Target(void* ptr);
void QDnsServiceRecord_DestroyQDnsServiceRecord(void* ptr);

#ifdef __cplusplus
}
#endif