#ifdef __cplusplus
extern "C" {
#endif

int QXmlContentHandler_Characters(void* ptr, char* ch);
int QXmlContentHandler_EndDocument(void* ptr);
int QXmlContentHandler_EndElement(void* ptr, char* namespaceURI, char* localName, char* qName);
int QXmlContentHandler_EndPrefixMapping(void* ptr, char* prefix);
char* QXmlContentHandler_ErrorString(void* ptr);
int QXmlContentHandler_IgnorableWhitespace(void* ptr, char* ch);
int QXmlContentHandler_ProcessingInstruction(void* ptr, char* target, char* data);
void QXmlContentHandler_SetDocumentLocator(void* ptr, void* locator);
int QXmlContentHandler_SkippedEntity(void* ptr, char* name);
int QXmlContentHandler_StartDocument(void* ptr);
int QXmlContentHandler_StartElement(void* ptr, char* namespaceURI, char* localName, char* qName, void* atts);
int QXmlContentHandler_StartPrefixMapping(void* ptr, char* prefix, char* uri);
void QXmlContentHandler_DestroyQXmlContentHandler(void* ptr);

#ifdef __cplusplus
}
#endif