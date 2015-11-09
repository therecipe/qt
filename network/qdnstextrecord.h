#ifdef __cplusplus
extern "C" {
#endif

void* QDnsTextRecord_NewQDnsTextRecord();
void* QDnsTextRecord_NewQDnsTextRecord2(void* other);
char* QDnsTextRecord_Name(void* ptr);
void QDnsTextRecord_Swap(void* ptr, void* other);
void QDnsTextRecord_DestroyQDnsTextRecord(void* ptr);

#ifdef __cplusplus
}
#endif