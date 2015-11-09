#ifdef __cplusplus
extern "C" {
#endif

void* QXmlQuery_NewQXmlQuery();
void* QXmlQuery_NewQXmlQuery4(int queryLanguage, void* np);
void* QXmlQuery_NewQXmlQuery3(void* np);
void* QXmlQuery_NewQXmlQuery2(void* other);
void QXmlQuery_BindVariable5(void* ptr, char* localName, void* device);
void QXmlQuery_BindVariable4(void* ptr, char* localName, void* value);
void QXmlQuery_BindVariable6(void* ptr, char* localName, void* query);
void QXmlQuery_BindVariable2(void* ptr, void* name, void* device);
void QXmlQuery_BindVariable(void* ptr, void* name, void* value);
void QXmlQuery_BindVariable3(void* ptr, void* name, void* query);
int QXmlQuery_IsValid(void* ptr);
void* QXmlQuery_MessageHandler(void* ptr);
void* QXmlQuery_NetworkAccessManager(void* ptr);
int QXmlQuery_QueryLanguage(void* ptr);
int QXmlQuery_SetFocus3(void* ptr, void* document);
int QXmlQuery_SetFocus4(void* ptr, char* focus);
int QXmlQuery_SetFocus2(void* ptr, void* documentURI);
void QXmlQuery_SetFocus(void* ptr, void* item);
void QXmlQuery_SetInitialTemplateName2(void* ptr, char* localName);
void QXmlQuery_SetInitialTemplateName(void* ptr, void* name);
void QXmlQuery_SetMessageHandler(void* ptr, void* aMessageHandler);
void QXmlQuery_SetNetworkAccessManager(void* ptr, void* newManager);
void QXmlQuery_SetQuery(void* ptr, void* sourceCode, void* documentURI);
void QXmlQuery_SetQuery3(void* ptr, char* sourceCode, void* documentURI);
void QXmlQuery_SetQuery2(void* ptr, void* queryURI, void* baseURI);
void QXmlQuery_SetUriResolver(void* ptr, void* resolver);
void* QXmlQuery_UriResolver(void* ptr);
void QXmlQuery_DestroyQXmlQuery(void* ptr);

#ifdef __cplusplus
}
#endif