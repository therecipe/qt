#ifdef __cplusplus
extern "C" {
#endif

void* QDnsMailExchangeRecord_NewQDnsMailExchangeRecord();
void* QDnsMailExchangeRecord_NewQDnsMailExchangeRecord2(void* other);
char* QDnsMailExchangeRecord_Exchange(void* ptr);
char* QDnsMailExchangeRecord_Name(void* ptr);
void QDnsMailExchangeRecord_Swap(void* ptr, void* other);
void QDnsMailExchangeRecord_DestroyQDnsMailExchangeRecord(void* ptr);

#ifdef __cplusplus
}
#endif