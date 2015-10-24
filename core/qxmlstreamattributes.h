#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QXmlStreamAttributes_NewQXmlStreamAttributes();
void QXmlStreamAttributes_Append(QtObjectPtr ptr, char* namespaceUri, char* name, char* value);
void QXmlStreamAttributes_Append2(QtObjectPtr ptr, char* qualifiedName, char* value);
int QXmlStreamAttributes_HasAttribute2(QtObjectPtr ptr, QtObjectPtr qualifiedName);
int QXmlStreamAttributes_HasAttribute3(QtObjectPtr ptr, char* namespaceUri, char* name);
int QXmlStreamAttributes_HasAttribute(QtObjectPtr ptr, char* qualifiedName);

#ifdef __cplusplus
}
#endif