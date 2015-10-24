#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QNdefNfcTextRecord_NewQNdefNfcTextRecord();
QtObjectPtr QNdefNfcTextRecord_NewQNdefNfcTextRecord2(QtObjectPtr other);
int QNdefNfcTextRecord_Encoding(QtObjectPtr ptr);
char* QNdefNfcTextRecord_Locale(QtObjectPtr ptr);
void QNdefNfcTextRecord_SetEncoding(QtObjectPtr ptr, int encoding);
void QNdefNfcTextRecord_SetLocale(QtObjectPtr ptr, char* locale);
void QNdefNfcTextRecord_SetText(QtObjectPtr ptr, char* text);
char* QNdefNfcTextRecord_Text(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif