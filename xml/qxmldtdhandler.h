#ifdef __cplusplus
extern "C" {
#endif

char* QXmlDTDHandler_ErrorString(void* ptr);
int QXmlDTDHandler_NotationDecl(void* ptr, char* name, char* publicId, char* systemId);
int QXmlDTDHandler_UnparsedEntityDecl(void* ptr, char* name, char* publicId, char* systemId, char* notationName);
void QXmlDTDHandler_DestroyQXmlDTDHandler(void* ptr);

#ifdef __cplusplus
}
#endif