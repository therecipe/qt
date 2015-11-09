#ifdef __cplusplus
extern "C" {
#endif

void* QXmlFormatter_NewQXmlFormatter(void* query, void* outputDevice);
void QXmlFormatter_Attribute(void* ptr, void* name, void* value);
void QXmlFormatter_Characters(void* ptr, void* value);
void QXmlFormatter_Comment(void* ptr, char* value);
void QXmlFormatter_EndDocument(void* ptr);
void QXmlFormatter_EndElement(void* ptr);
void QXmlFormatter_EndOfSequence(void* ptr);
int QXmlFormatter_IndentationDepth(void* ptr);
void QXmlFormatter_ProcessingInstruction(void* ptr, void* name, char* value);
void QXmlFormatter_SetIndentationDepth(void* ptr, int depth);
void QXmlFormatter_StartDocument(void* ptr);
void QXmlFormatter_StartElement(void* ptr, void* name);
void QXmlFormatter_StartOfSequence(void* ptr);

#ifdef __cplusplus
}
#endif