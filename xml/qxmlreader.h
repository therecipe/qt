#ifdef __cplusplus
extern "C" {
#endif

void* QXmlReader_DTDHandler(void* ptr);
void* QXmlReader_ContentHandler(void* ptr);
void* QXmlReader_DeclHandler(void* ptr);
void* QXmlReader_EntityResolver(void* ptr);
void* QXmlReader_ErrorHandler(void* ptr);
int QXmlReader_Feature(void* ptr, char* name, int ok);
int QXmlReader_HasFeature(void* ptr, char* name);
int QXmlReader_HasProperty(void* ptr, char* name);
void* QXmlReader_LexicalHandler(void* ptr);
int QXmlReader_Parse(void* ptr, void* input);
void* QXmlReader_Property(void* ptr, char* name, int ok);
void QXmlReader_SetContentHandler(void* ptr, void* handler);
void QXmlReader_SetDTDHandler(void* ptr, void* handler);
void QXmlReader_SetDeclHandler(void* ptr, void* handler);
void QXmlReader_SetEntityResolver(void* ptr, void* handler);
void QXmlReader_SetErrorHandler(void* ptr, void* handler);
void QXmlReader_SetFeature(void* ptr, char* name, int value);
void QXmlReader_SetLexicalHandler(void* ptr, void* handler);
void QXmlReader_DestroyQXmlReader(void* ptr);

#ifdef __cplusplus
}
#endif