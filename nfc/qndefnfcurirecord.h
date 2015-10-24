#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QNdefNfcUriRecord_NewQNdefNfcUriRecord();
QtObjectPtr QNdefNfcUriRecord_NewQNdefNfcUriRecord2(QtObjectPtr other);
void QNdefNfcUriRecord_SetUri(QtObjectPtr ptr, char* uri);
char* QNdefNfcUriRecord_Uri(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif