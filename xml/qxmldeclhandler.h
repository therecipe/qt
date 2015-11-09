#ifdef __cplusplus
extern "C" {
#endif

int QXmlDeclHandler_AttributeDecl(void* ptr, char* eName, char* aName, char* ty, char* valueDefault, char* value);
char* QXmlDeclHandler_ErrorString(void* ptr);
int QXmlDeclHandler_ExternalEntityDecl(void* ptr, char* name, char* publicId, char* systemId);
int QXmlDeclHandler_InternalEntityDecl(void* ptr, char* name, char* value);
void QXmlDeclHandler_DestroyQXmlDeclHandler(void* ptr);

#ifdef __cplusplus
}
#endif