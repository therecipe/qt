#ifdef __cplusplus
extern "C" {
#endif

int QXmlStreamWriter_AutoFormattingIndent(void* ptr);
void QXmlStreamWriter_SetAutoFormattingIndent(void* ptr, int spacesOrTabs);
int QXmlStreamWriter_AutoFormatting(void* ptr);
void* QXmlStreamWriter_Codec(void* ptr);
void* QXmlStreamWriter_Device(void* ptr);
int QXmlStreamWriter_HasError(void* ptr);
void QXmlStreamWriter_SetAutoFormatting(void* ptr, int enable);
void QXmlStreamWriter_SetCodec(void* ptr, void* codec);
void QXmlStreamWriter_SetCodec2(void* ptr, char* codecName);
void QXmlStreamWriter_SetDevice(void* ptr, void* device);
void QXmlStreamWriter_WriteAttribute(void* ptr, char* namespaceUri, char* name, char* value);
void QXmlStreamWriter_WriteAttribute2(void* ptr, char* qualifiedName, char* value);
void QXmlStreamWriter_WriteAttribute3(void* ptr, void* attribute);
void QXmlStreamWriter_WriteAttributes(void* ptr, void* attributes);
void QXmlStreamWriter_WriteCDATA(void* ptr, char* text);
void QXmlStreamWriter_WriteCharacters(void* ptr, char* text);
void QXmlStreamWriter_WriteComment(void* ptr, char* text);
void QXmlStreamWriter_WriteCurrentToken(void* ptr, void* reader);
void QXmlStreamWriter_WriteDTD(void* ptr, char* dtd);
void QXmlStreamWriter_WriteDefaultNamespace(void* ptr, char* namespaceUri);
void QXmlStreamWriter_WriteEmptyElement(void* ptr, char* namespaceUri, char* name);
void QXmlStreamWriter_WriteEmptyElement2(void* ptr, char* qualifiedName);
void QXmlStreamWriter_WriteEndDocument(void* ptr);
void QXmlStreamWriter_WriteEndElement(void* ptr);
void QXmlStreamWriter_WriteEntityReference(void* ptr, char* name);
void QXmlStreamWriter_WriteNamespace(void* ptr, char* namespaceUri, char* prefix);
void QXmlStreamWriter_WriteProcessingInstruction(void* ptr, char* target, char* data);
void QXmlStreamWriter_WriteStartDocument3(void* ptr);
void QXmlStreamWriter_WriteStartDocument(void* ptr, char* version);
void QXmlStreamWriter_WriteStartDocument2(void* ptr, char* version, int standalone);
void QXmlStreamWriter_WriteStartElement(void* ptr, char* namespaceUri, char* name);
void QXmlStreamWriter_WriteStartElement2(void* ptr, char* qualifiedName);
void QXmlStreamWriter_WriteTextElement(void* ptr, char* namespaceUri, char* name, char* text);
void QXmlStreamWriter_WriteTextElement2(void* ptr, char* qualifiedName, char* text);
void QXmlStreamWriter_DestroyQXmlStreamWriter(void* ptr);

#ifdef __cplusplus
}
#endif