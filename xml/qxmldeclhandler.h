#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QXmlDeclHandler_AttributeDecl(QtObjectPtr ptr, char* eName, char* aName, char* ty, char* valueDefault, char* value);
char* QXmlDeclHandler_ErrorString(QtObjectPtr ptr);
int QXmlDeclHandler_ExternalEntityDecl(QtObjectPtr ptr, char* name, char* publicId, char* systemId);
int QXmlDeclHandler_InternalEntityDecl(QtObjectPtr ptr, char* name, char* value);
void QXmlDeclHandler_DestroyQXmlDeclHandler(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif