#ifdef __cplusplus
extern "C" {
#endif

void* QDnsDomainNameRecord_NewQDnsDomainNameRecord();
void* QDnsDomainNameRecord_NewQDnsDomainNameRecord2(void* other);
char* QDnsDomainNameRecord_Name(void* ptr);
void QDnsDomainNameRecord_Swap(void* ptr, void* other);
char* QDnsDomainNameRecord_Value(void* ptr);
void QDnsDomainNameRecord_DestroyQDnsDomainNameRecord(void* ptr);

#ifdef __cplusplus
}
#endif