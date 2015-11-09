#ifdef __cplusplus
extern "C" {
#endif

int QXmlLexicalHandler_Comment(void* ptr, char* ch);
int QXmlLexicalHandler_EndCDATA(void* ptr);
int QXmlLexicalHandler_EndDTD(void* ptr);
int QXmlLexicalHandler_EndEntity(void* ptr, char* name);
char* QXmlLexicalHandler_ErrorString(void* ptr);
int QXmlLexicalHandler_StartCDATA(void* ptr);
int QXmlLexicalHandler_StartDTD(void* ptr, char* name, char* publicId, char* systemId);
int QXmlLexicalHandler_StartEntity(void* ptr, char* name);
void QXmlLexicalHandler_DestroyQXmlLexicalHandler(void* ptr);

#ifdef __cplusplus
}
#endif