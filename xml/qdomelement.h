#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDomElement_NewQDomElement();
QtObjectPtr QDomElement_NewQDomElement2(QtObjectPtr x);
char* QDomElement_Attribute(QtObjectPtr ptr, char* name, char* defValue);
char* QDomElement_AttributeNS(QtObjectPtr ptr, char* nsURI, char* localName, char* defValue);
int QDomElement_HasAttribute(QtObjectPtr ptr, char* name);
int QDomElement_HasAttributeNS(QtObjectPtr ptr, char* nsURI, char* localName);
int QDomElement_NodeType(QtObjectPtr ptr);
void QDomElement_RemoveAttribute(QtObjectPtr ptr, char* name);
void QDomElement_RemoveAttributeNS(QtObjectPtr ptr, char* nsURI, char* localName);
void QDomElement_SetAttribute(QtObjectPtr ptr, char* name, char* value);
void QDomElement_SetAttribute2(QtObjectPtr ptr, char* name, int value);
void QDomElement_SetAttributeNS(QtObjectPtr ptr, char* nsURI, char* qName, char* value);
void QDomElement_SetAttributeNS2(QtObjectPtr ptr, char* nsURI, char* qName, int value);
void QDomElement_SetTagName(QtObjectPtr ptr, char* name);
char* QDomElement_TagName(QtObjectPtr ptr);
char* QDomElement_Text(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif