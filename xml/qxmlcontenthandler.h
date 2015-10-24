#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QXmlContentHandler_Characters(QtObjectPtr ptr, char* ch);
int QXmlContentHandler_EndDocument(QtObjectPtr ptr);
int QXmlContentHandler_EndElement(QtObjectPtr ptr, char* namespaceURI, char* localName, char* qName);
int QXmlContentHandler_EndPrefixMapping(QtObjectPtr ptr, char* prefix);
char* QXmlContentHandler_ErrorString(QtObjectPtr ptr);
int QXmlContentHandler_IgnorableWhitespace(QtObjectPtr ptr, char* ch);
int QXmlContentHandler_ProcessingInstruction(QtObjectPtr ptr, char* target, char* data);
void QXmlContentHandler_SetDocumentLocator(QtObjectPtr ptr, QtObjectPtr locator);
int QXmlContentHandler_SkippedEntity(QtObjectPtr ptr, char* name);
int QXmlContentHandler_StartDocument(QtObjectPtr ptr);
int QXmlContentHandler_StartElement(QtObjectPtr ptr, char* namespaceURI, char* localName, char* qName, QtObjectPtr atts);
int QXmlContentHandler_StartPrefixMapping(QtObjectPtr ptr, char* prefix, char* uri);
void QXmlContentHandler_DestroyQXmlContentHandler(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif