#ifdef __cplusplus
extern "C" {
#endif

void* QDomElement_NewQDomElement();
void* QDomElement_NewQDomElement2(void* x);
char* QDomElement_Attribute(void* ptr, char* name, char* defValue);
char* QDomElement_AttributeNS(void* ptr, char* nsURI, char* localName, char* defValue);
int QDomElement_HasAttribute(void* ptr, char* name);
int QDomElement_HasAttributeNS(void* ptr, char* nsURI, char* localName);
int QDomElement_NodeType(void* ptr);
void QDomElement_RemoveAttribute(void* ptr, char* name);
void QDomElement_RemoveAttributeNS(void* ptr, char* nsURI, char* localName);
void QDomElement_SetAttribute(void* ptr, char* name, char* value);
void QDomElement_SetAttribute2(void* ptr, char* name, int value);
void QDomElement_SetAttributeNS(void* ptr, char* nsURI, char* qName, char* value);
void QDomElement_SetAttributeNS2(void* ptr, char* nsURI, char* qName, int value);
void QDomElement_SetTagName(void* ptr, char* name);
char* QDomElement_TagName(void* ptr);
char* QDomElement_Text(void* ptr);

#ifdef __cplusplus
}
#endif