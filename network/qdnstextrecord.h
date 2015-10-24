#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDnsTextRecord_NewQDnsTextRecord();
QtObjectPtr QDnsTextRecord_NewQDnsTextRecord2(QtObjectPtr other);
char* QDnsTextRecord_Name(QtObjectPtr ptr);
void QDnsTextRecord_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QDnsTextRecord_DestroyQDnsTextRecord(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif