#ifdef __cplusplus
extern "C" {
#endif

void* QXmlDefaultHandler_NewQXmlDefaultHandler();
void QXmlDefaultHandler_DestroyQXmlDefaultHandler(void* ptr);
int QXmlDefaultHandler_AttributeDecl(void* ptr, char* eName, char* aName, char* ty, char* valueDefault, char* value);
int QXmlDefaultHandler_Characters(void* ptr, char* ch);
int QXmlDefaultHandler_Comment(void* ptr, char* ch);
int QXmlDefaultHandler_EndCDATA(void* ptr);
int QXmlDefaultHandler_EndDTD(void* ptr);
int QXmlDefaultHandler_EndDocument(void* ptr);
int QXmlDefaultHandler_EndElement(void* ptr, char* namespaceURI, char* localName, char* qName);
int QXmlDefaultHandler_EndEntity(void* ptr, char* name);
int QXmlDefaultHandler_EndPrefixMapping(void* ptr, char* prefix);
int QXmlDefaultHandler_Error(void* ptr, void* exception);
char* QXmlDefaultHandler_ErrorString(void* ptr);
int QXmlDefaultHandler_ExternalEntityDecl(void* ptr, char* name, char* publicId, char* systemId);
int QXmlDefaultHandler_FatalError(void* ptr, void* exception);
int QXmlDefaultHandler_IgnorableWhitespace(void* ptr, char* ch);
int QXmlDefaultHandler_InternalEntityDecl(void* ptr, char* name, char* value);
int QXmlDefaultHandler_NotationDecl(void* ptr, char* name, char* publicId, char* systemId);
int QXmlDefaultHandler_ProcessingInstruction(void* ptr, char* target, char* data);
void QXmlDefaultHandler_SetDocumentLocator(void* ptr, void* locator);
int QXmlDefaultHandler_SkippedEntity(void* ptr, char* name);
int QXmlDefaultHandler_StartCDATA(void* ptr);
int QXmlDefaultHandler_StartDTD(void* ptr, char* name, char* publicId, char* systemId);
int QXmlDefaultHandler_StartDocument(void* ptr);
int QXmlDefaultHandler_StartElement(void* ptr, char* namespaceURI, char* localName, char* qName, void* atts);
int QXmlDefaultHandler_StartEntity(void* ptr, char* name);
int QXmlDefaultHandler_StartPrefixMapping(void* ptr, char* prefix, char* uri);
int QXmlDefaultHandler_UnparsedEntityDecl(void* ptr, char* name, char* publicId, char* systemId, char* notationName);
int QXmlDefaultHandler_Warning(void* ptr, void* exception);

#ifdef __cplusplus
}
#endif