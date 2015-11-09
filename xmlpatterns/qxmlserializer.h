#ifdef __cplusplus
extern "C" {
#endif

void* QXmlSerializer_NewQXmlSerializer(void* query, void* outputDevice);
void QXmlSerializer_Attribute(void* ptr, void* name, void* value);
void QXmlSerializer_Characters(void* ptr, void* value);
void QXmlSerializer_Comment(void* ptr, char* value);
void QXmlSerializer_EndDocument(void* ptr);
void QXmlSerializer_EndElement(void* ptr);
void* QXmlSerializer_Codec(void* ptr);
void QXmlSerializer_EndOfSequence(void* ptr);
void QXmlSerializer_NamespaceBinding(void* ptr, void* nb);
void* QXmlSerializer_OutputDevice(void* ptr);
void QXmlSerializer_ProcessingInstruction(void* ptr, void* name, char* value);
void QXmlSerializer_SetCodec(void* ptr, void* outputCodec);
void QXmlSerializer_StartDocument(void* ptr);
void QXmlSerializer_StartElement(void* ptr, void* name);
void QXmlSerializer_StartOfSequence(void* ptr);

#ifdef __cplusplus
}
#endif