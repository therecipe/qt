#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord();
QtObjectPtr QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord3(QtObjectPtr other);
QtObjectPtr QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord2(QtObjectPtr other);
int QNdefNfcSmartPosterRecord_Action(QtObjectPtr ptr);
void QNdefNfcSmartPosterRecord_AddIcon2(QtObjectPtr ptr, QtObjectPtr ty, QtObjectPtr data);
int QNdefNfcSmartPosterRecord_AddTitle(QtObjectPtr ptr, QtObjectPtr text);
int QNdefNfcSmartPosterRecord_AddTitle2(QtObjectPtr ptr, char* text, char* locale, int encoding);
int QNdefNfcSmartPosterRecord_HasAction(QtObjectPtr ptr);
int QNdefNfcSmartPosterRecord_HasIcon(QtObjectPtr ptr, QtObjectPtr mimetype);
int QNdefNfcSmartPosterRecord_HasSize(QtObjectPtr ptr);
int QNdefNfcSmartPosterRecord_HasTitle(QtObjectPtr ptr, char* locale);
int QNdefNfcSmartPosterRecord_HasTypeInfo(QtObjectPtr ptr);
int QNdefNfcSmartPosterRecord_IconCount(QtObjectPtr ptr);
int QNdefNfcSmartPosterRecord_RemoveIcon2(QtObjectPtr ptr, QtObjectPtr ty);
int QNdefNfcSmartPosterRecord_RemoveTitle(QtObjectPtr ptr, QtObjectPtr text);
int QNdefNfcSmartPosterRecord_RemoveTitle2(QtObjectPtr ptr, char* locale);
void QNdefNfcSmartPosterRecord_SetAction(QtObjectPtr ptr, int act);
void QNdefNfcSmartPosterRecord_SetTypeInfo(QtObjectPtr ptr, QtObjectPtr ty);
void QNdefNfcSmartPosterRecord_SetUri(QtObjectPtr ptr, QtObjectPtr url);
void QNdefNfcSmartPosterRecord_SetUri2(QtObjectPtr ptr, char* url);
char* QNdefNfcSmartPosterRecord_Title(QtObjectPtr ptr, char* locale);
int QNdefNfcSmartPosterRecord_TitleCount(QtObjectPtr ptr);
char* QNdefNfcSmartPosterRecord_Uri(QtObjectPtr ptr);
void QNdefNfcSmartPosterRecord_DestroyQNdefNfcSmartPosterRecord(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif