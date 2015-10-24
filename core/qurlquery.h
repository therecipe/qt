#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QUrlQuery_NewQUrlQuery();
QtObjectPtr QUrlQuery_NewQUrlQuery3(char* queryString);
QtObjectPtr QUrlQuery_NewQUrlQuery2(char* url);
QtObjectPtr QUrlQuery_NewQUrlQuery4(QtObjectPtr other);
void QUrlQuery_AddQueryItem(QtObjectPtr ptr, char* key, char* value);
char* QUrlQuery_AllQueryItemValues(QtObjectPtr ptr, char* key, int encoding);
void QUrlQuery_Clear(QtObjectPtr ptr);
int QUrlQuery_IsEmpty(QtObjectPtr ptr);
char* QUrlQuery_Query(QtObjectPtr ptr, int encoding);
void QUrlQuery_RemoveAllQueryItems(QtObjectPtr ptr, char* key);
void QUrlQuery_RemoveQueryItem(QtObjectPtr ptr, char* key);
void QUrlQuery_SetQuery(QtObjectPtr ptr, char* queryString);
void QUrlQuery_SetQueryDelimiters(QtObjectPtr ptr, QtObjectPtr valueDelimiter, QtObjectPtr pairDelimiter);
void QUrlQuery_Swap(QtObjectPtr ptr, QtObjectPtr other);
char* QUrlQuery_ToString(QtObjectPtr ptr, int encoding);
void QUrlQuery_DestroyQUrlQuery(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif