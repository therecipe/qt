#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QXmlQuery_NewQXmlQuery();
QtObjectPtr QXmlQuery_NewQXmlQuery4(int queryLanguage, QtObjectPtr np);
QtObjectPtr QXmlQuery_NewQXmlQuery3(QtObjectPtr np);
QtObjectPtr QXmlQuery_NewQXmlQuery2(QtObjectPtr other);
void QXmlQuery_BindVariable5(QtObjectPtr ptr, char* localName, QtObjectPtr device);
void QXmlQuery_BindVariable4(QtObjectPtr ptr, char* localName, QtObjectPtr value);
void QXmlQuery_BindVariable6(QtObjectPtr ptr, char* localName, QtObjectPtr query);
void QXmlQuery_BindVariable2(QtObjectPtr ptr, QtObjectPtr name, QtObjectPtr device);
void QXmlQuery_BindVariable(QtObjectPtr ptr, QtObjectPtr name, QtObjectPtr value);
void QXmlQuery_BindVariable3(QtObjectPtr ptr, QtObjectPtr name, QtObjectPtr query);
int QXmlQuery_IsValid(QtObjectPtr ptr);
QtObjectPtr QXmlQuery_MessageHandler(QtObjectPtr ptr);
QtObjectPtr QXmlQuery_NetworkAccessManager(QtObjectPtr ptr);
int QXmlQuery_QueryLanguage(QtObjectPtr ptr);
int QXmlQuery_SetFocus3(QtObjectPtr ptr, QtObjectPtr document);
int QXmlQuery_SetFocus4(QtObjectPtr ptr, char* focus);
int QXmlQuery_SetFocus2(QtObjectPtr ptr, char* documentURI);
void QXmlQuery_SetFocus(QtObjectPtr ptr, QtObjectPtr item);
void QXmlQuery_SetInitialTemplateName2(QtObjectPtr ptr, char* localName);
void QXmlQuery_SetInitialTemplateName(QtObjectPtr ptr, QtObjectPtr name);
void QXmlQuery_SetMessageHandler(QtObjectPtr ptr, QtObjectPtr aMessageHandler);
void QXmlQuery_SetNetworkAccessManager(QtObjectPtr ptr, QtObjectPtr newManager);
void QXmlQuery_SetQuery(QtObjectPtr ptr, QtObjectPtr sourceCode, char* documentURI);
void QXmlQuery_SetQuery3(QtObjectPtr ptr, char* sourceCode, char* documentURI);
void QXmlQuery_SetQuery2(QtObjectPtr ptr, char* queryURI, char* baseURI);
void QXmlQuery_SetUriResolver(QtObjectPtr ptr, QtObjectPtr resolver);
QtObjectPtr QXmlQuery_UriResolver(QtObjectPtr ptr);
void QXmlQuery_DestroyQXmlQuery(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif