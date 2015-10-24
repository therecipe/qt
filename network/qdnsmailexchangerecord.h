#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDnsMailExchangeRecord_NewQDnsMailExchangeRecord();
QtObjectPtr QDnsMailExchangeRecord_NewQDnsMailExchangeRecord2(QtObjectPtr other);
char* QDnsMailExchangeRecord_Exchange(QtObjectPtr ptr);
char* QDnsMailExchangeRecord_Name(QtObjectPtr ptr);
void QDnsMailExchangeRecord_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QDnsMailExchangeRecord_DestroyQDnsMailExchangeRecord(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif