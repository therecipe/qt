#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDnsServiceRecord_NewQDnsServiceRecord();
QtObjectPtr QDnsServiceRecord_NewQDnsServiceRecord2(QtObjectPtr other);
char* QDnsServiceRecord_Name(QtObjectPtr ptr);
void QDnsServiceRecord_Swap(QtObjectPtr ptr, QtObjectPtr other);
char* QDnsServiceRecord_Target(QtObjectPtr ptr);
void QDnsServiceRecord_DestroyQDnsServiceRecord(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif