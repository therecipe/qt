#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QXmlDefaultHandler_NewQXmlDefaultHandler();
void QXmlDefaultHandler_DestroyQXmlDefaultHandler(QtObjectPtr ptr);
int QXmlDefaultHandler_AttributeDecl(QtObjectPtr ptr, char* eName, char* aName, char* ty, char* valueDefault, char* value);
int QXmlDefaultHandler_Characters(QtObjectPtr ptr, char* ch);
int QXmlDefaultHandler_Comment(QtObjectPtr ptr, char* ch);
int QXmlDefaultHandler_EndCDATA(QtObjectPtr ptr);
int QXmlDefaultHandler_EndDTD(QtObjectPtr ptr);
int QXmlDefaultHandler_EndDocument(QtObjectPtr ptr);
int QXmlDefaultHandler_EndElement(QtObjectPtr ptr, char* namespaceURI, char* localName, char* qName);
int QXmlDefaultHandler_EndEntity(QtObjectPtr ptr, char* name);
int QXmlDefaultHandler_EndPrefixMapping(QtObjectPtr ptr, char* prefix);
int QXmlDefaultHandler_Error(QtObjectPtr ptr, QtObjectPtr exception);
char* QXmlDefaultHandler_ErrorString(QtObjectPtr ptr);
int QXmlDefaultHandler_ExternalEntityDecl(QtObjectPtr ptr, char* name, char* publicId, char* systemId);
int QXmlDefaultHandler_FatalError(QtObjectPtr ptr, QtObjectPtr exception);
int QXmlDefaultHandler_IgnorableWhitespace(QtObjectPtr ptr, char* ch);
int QXmlDefaultHandler_InternalEntityDecl(QtObjectPtr ptr, char* name, char* value);
int QXmlDefaultHandler_NotationDecl(QtObjectPtr ptr, char* name, char* publicId, char* systemId);
int QXmlDefaultHandler_ProcessingInstruction(QtObjectPtr ptr, char* target, char* data);
void QXmlDefaultHandler_SetDocumentLocator(QtObjectPtr ptr, QtObjectPtr locator);
int QXmlDefaultHandler_SkippedEntity(QtObjectPtr ptr, char* name);
int QXmlDefaultHandler_StartCDATA(QtObjectPtr ptr);
int QXmlDefaultHandler_StartDTD(QtObjectPtr ptr, char* name, char* publicId, char* systemId);
int QXmlDefaultHandler_StartDocument(QtObjectPtr ptr);
int QXmlDefaultHandler_StartElement(QtObjectPtr ptr, char* namespaceURI, char* localName, char* qName, QtObjectPtr atts);
int QXmlDefaultHandler_StartEntity(QtObjectPtr ptr, char* name);
int QXmlDefaultHandler_StartPrefixMapping(QtObjectPtr ptr, char* prefix, char* uri);
int QXmlDefaultHandler_UnparsedEntityDecl(QtObjectPtr ptr, char* name, char* publicId, char* systemId, char* notationName);
int QXmlDefaultHandler_Warning(QtObjectPtr ptr, QtObjectPtr exception);

#ifdef __cplusplus
}
#endif