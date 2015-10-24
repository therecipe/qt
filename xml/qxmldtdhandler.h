#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QXmlDTDHandler_ErrorString(QtObjectPtr ptr);
int QXmlDTDHandler_NotationDecl(QtObjectPtr ptr, char* name, char* publicId, char* systemId);
int QXmlDTDHandler_UnparsedEntityDecl(QtObjectPtr ptr, char* name, char* publicId, char* systemId, char* notationName);
void QXmlDTDHandler_DestroyQXmlDTDHandler(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif