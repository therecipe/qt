#ifdef __cplusplus
extern "C" {
#endif

void* QUrlQuery_NewQUrlQuery();
void* QUrlQuery_NewQUrlQuery3(char* queryString);
void* QUrlQuery_NewQUrlQuery2(void* url);
void* QUrlQuery_NewQUrlQuery4(void* other);
void QUrlQuery_AddQueryItem(void* ptr, char* key, char* value);
char* QUrlQuery_AllQueryItemValues(void* ptr, char* key, int encoding);
void QUrlQuery_Clear(void* ptr);
int QUrlQuery_IsEmpty(void* ptr);
char* QUrlQuery_Query(void* ptr, int encoding);
void QUrlQuery_RemoveAllQueryItems(void* ptr, char* key);
void QUrlQuery_RemoveQueryItem(void* ptr, char* key);
void QUrlQuery_SetQuery(void* ptr, char* queryString);
void QUrlQuery_SetQueryDelimiters(void* ptr, void* valueDelimiter, void* pairDelimiter);
void QUrlQuery_Swap(void* ptr, void* other);
char* QUrlQuery_ToString(void* ptr, int encoding);
void QUrlQuery_DestroyQUrlQuery(void* ptr);

#ifdef __cplusplus
}
#endif