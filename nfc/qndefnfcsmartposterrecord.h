#ifdef __cplusplus
extern "C" {
#endif

void* QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord();
void* QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord3(void* other);
void* QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord2(void* other);
int QNdefNfcSmartPosterRecord_Action(void* ptr);
void QNdefNfcSmartPosterRecord_AddIcon2(void* ptr, void* ty, void* data);
int QNdefNfcSmartPosterRecord_AddTitle(void* ptr, void* text);
int QNdefNfcSmartPosterRecord_AddTitle2(void* ptr, char* text, char* locale, int encoding);
int QNdefNfcSmartPosterRecord_HasAction(void* ptr);
int QNdefNfcSmartPosterRecord_HasIcon(void* ptr, void* mimetype);
int QNdefNfcSmartPosterRecord_HasSize(void* ptr);
int QNdefNfcSmartPosterRecord_HasTitle(void* ptr, char* locale);
int QNdefNfcSmartPosterRecord_HasTypeInfo(void* ptr);
void* QNdefNfcSmartPosterRecord_Icon(void* ptr, void* mimetype);
int QNdefNfcSmartPosterRecord_IconCount(void* ptr);
int QNdefNfcSmartPosterRecord_RemoveIcon2(void* ptr, void* ty);
int QNdefNfcSmartPosterRecord_RemoveTitle(void* ptr, void* text);
int QNdefNfcSmartPosterRecord_RemoveTitle2(void* ptr, char* locale);
void QNdefNfcSmartPosterRecord_SetAction(void* ptr, int act);
void QNdefNfcSmartPosterRecord_SetTypeInfo(void* ptr, void* ty);
void QNdefNfcSmartPosterRecord_SetUri(void* ptr, void* url);
void QNdefNfcSmartPosterRecord_SetUri2(void* ptr, void* url);
char* QNdefNfcSmartPosterRecord_Title(void* ptr, char* locale);
int QNdefNfcSmartPosterRecord_TitleCount(void* ptr);
void* QNdefNfcSmartPosterRecord_TypeInfo(void* ptr);
void QNdefNfcSmartPosterRecord_DestroyQNdefNfcSmartPosterRecord(void* ptr);

#ifdef __cplusplus
}
#endif