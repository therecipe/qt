#ifdef __cplusplus
extern "C" {
#endif

void* QXmlParseException_NewQXmlParseException(char* name, int c, int l, char* p, char* s);
void* QXmlParseException_NewQXmlParseException2(void* other);
int QXmlParseException_ColumnNumber(void* ptr);
int QXmlParseException_LineNumber(void* ptr);
char* QXmlParseException_Message(void* ptr);
char* QXmlParseException_PublicId(void* ptr);
char* QXmlParseException_SystemId(void* ptr);
void QXmlParseException_DestroyQXmlParseException(void* ptr);

#ifdef __cplusplus
}
#endif