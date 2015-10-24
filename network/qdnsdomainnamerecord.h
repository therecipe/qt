#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDnsDomainNameRecord_NewQDnsDomainNameRecord();
QtObjectPtr QDnsDomainNameRecord_NewQDnsDomainNameRecord2(QtObjectPtr other);
char* QDnsDomainNameRecord_Name(QtObjectPtr ptr);
void QDnsDomainNameRecord_Swap(QtObjectPtr ptr, QtObjectPtr other);
char* QDnsDomainNameRecord_Value(QtObjectPtr ptr);
void QDnsDomainNameRecord_DestroyQDnsDomainNameRecord(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif