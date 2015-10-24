#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QXmlStreamAttribute_NewQXmlStreamAttribute();
QtObjectPtr QXmlStreamAttribute_NewQXmlStreamAttribute3(char* namespaceUri, char* name, char* value);
QtObjectPtr QXmlStreamAttribute_NewQXmlStreamAttribute2(char* qualifiedName, char* value);
QtObjectPtr QXmlStreamAttribute_NewQXmlStreamAttribute4(QtObjectPtr other);
int QXmlStreamAttribute_IsDefault(QtObjectPtr ptr);
void QXmlStreamAttribute_DestroyQXmlStreamAttribute(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif