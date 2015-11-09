#ifdef __cplusplus
extern "C" {
#endif

void* QNdefNfcTextRecord_NewQNdefNfcTextRecord();
void* QNdefNfcTextRecord_NewQNdefNfcTextRecord2(void* other);
int QNdefNfcTextRecord_Encoding(void* ptr);
char* QNdefNfcTextRecord_Locale(void* ptr);
void QNdefNfcTextRecord_SetEncoding(void* ptr, int encoding);
void QNdefNfcTextRecord_SetLocale(void* ptr, char* locale);
void QNdefNfcTextRecord_SetText(void* ptr, char* text);
char* QNdefNfcTextRecord_Text(void* ptr);

#ifdef __cplusplus
}
#endif