#ifdef __cplusplus
extern "C" {
#endif

void* QXmlSimpleReader_DTDHandler(void* ptr);
void* QXmlSimpleReader_NewQXmlSimpleReader();
void* QXmlSimpleReader_ContentHandler(void* ptr);
void* QXmlSimpleReader_DeclHandler(void* ptr);
void* QXmlSimpleReader_EntityResolver(void* ptr);
void* QXmlSimpleReader_ErrorHandler(void* ptr);
int QXmlSimpleReader_Feature(void* ptr, char* name, int ok);
int QXmlSimpleReader_HasFeature(void* ptr, char* name);
int QXmlSimpleReader_HasProperty(void* ptr, char* name);
void* QXmlSimpleReader_LexicalHandler(void* ptr);
int QXmlSimpleReader_Parse(void* ptr, void* input);
int QXmlSimpleReader_Parse2(void* ptr, void* input);
int QXmlSimpleReader_Parse3(void* ptr, void* input, int incremental);
int QXmlSimpleReader_ParseContinue(void* ptr);
void* QXmlSimpleReader_Property(void* ptr, char* name, int ok);
void QXmlSimpleReader_SetContentHandler(void* ptr, void* handler);
void QXmlSimpleReader_SetDTDHandler(void* ptr, void* handler);
void QXmlSimpleReader_SetDeclHandler(void* ptr, void* handler);
void QXmlSimpleReader_SetEntityResolver(void* ptr, void* handler);
void QXmlSimpleReader_SetErrorHandler(void* ptr, void* handler);
void QXmlSimpleReader_SetFeature(void* ptr, char* name, int enable);
void QXmlSimpleReader_SetLexicalHandler(void* ptr, void* handler);
void QXmlSimpleReader_DestroyQXmlSimpleReader(void* ptr);

#ifdef __cplusplus
}
#endif