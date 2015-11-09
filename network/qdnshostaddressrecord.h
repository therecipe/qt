#ifdef __cplusplus
extern "C" {
#endif

void* QDnsHostAddressRecord_NewQDnsHostAddressRecord();
void* QDnsHostAddressRecord_NewQDnsHostAddressRecord2(void* other);
char* QDnsHostAddressRecord_Name(void* ptr);
void QDnsHostAddressRecord_Swap(void* ptr, void* other);
void QDnsHostAddressRecord_DestroyQDnsHostAddressRecord(void* ptr);

#ifdef __cplusplus
}
#endif