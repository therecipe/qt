#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QXmlSchema_NewQXmlSchema();
QtObjectPtr QXmlSchema_NewQXmlSchema2(QtObjectPtr other);
char* QXmlSchema_DocumentUri(QtObjectPtr ptr);
int QXmlSchema_IsValid(QtObjectPtr ptr);
int QXmlSchema_Load2(QtObjectPtr ptr, QtObjectPtr source, char* documentUri);
int QXmlSchema_Load3(QtObjectPtr ptr, QtObjectPtr data, char* documentUri);
int QXmlSchema_Load(QtObjectPtr ptr, char* source);
QtObjectPtr QXmlSchema_MessageHandler(QtObjectPtr ptr);
QtObjectPtr QXmlSchema_NetworkAccessManager(QtObjectPtr ptr);
void QXmlSchema_SetMessageHandler(QtObjectPtr ptr, QtObjectPtr handler);
void QXmlSchema_SetNetworkAccessManager(QtObjectPtr ptr, QtObjectPtr manager);
void QXmlSchema_SetUriResolver(QtObjectPtr ptr, QtObjectPtr resolver);
QtObjectPtr QXmlSchema_UriResolver(QtObjectPtr ptr);
void QXmlSchema_DestroyQXmlSchema(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif