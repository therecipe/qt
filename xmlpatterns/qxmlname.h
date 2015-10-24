#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QXmlName_NewQXmlName();
QtObjectPtr QXmlName_NewQXmlName2(QtObjectPtr namePool, char* localName, char* namespaceURI, char* prefix);
int QXmlName_QXmlName_IsNCName(char* candidate);
int QXmlName_IsNull(QtObjectPtr ptr);
char* QXmlName_LocalName(QtObjectPtr ptr, QtObjectPtr namePool);
char* QXmlName_NamespaceUri(QtObjectPtr ptr, QtObjectPtr namePool);
char* QXmlName_Prefix(QtObjectPtr ptr, QtObjectPtr namePool);
char* QXmlName_ToClarkName(QtObjectPtr ptr, QtObjectPtr namePool);

#ifdef __cplusplus
}
#endif