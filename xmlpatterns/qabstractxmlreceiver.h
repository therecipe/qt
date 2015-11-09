#ifdef __cplusplus
extern "C" {
#endif

void QAbstractXmlReceiver_Attribute(void* ptr, void* name, void* value);
void QAbstractXmlReceiver_Characters(void* ptr, void* value);
void QAbstractXmlReceiver_Comment(void* ptr, char* value);
void QAbstractXmlReceiver_EndDocument(void* ptr);
void QAbstractXmlReceiver_EndElement(void* ptr);
void QAbstractXmlReceiver_EndOfSequence(void* ptr);
void QAbstractXmlReceiver_NamespaceBinding(void* ptr, void* name);
void QAbstractXmlReceiver_ProcessingInstruction(void* ptr, void* target, char* value);
void QAbstractXmlReceiver_StartDocument(void* ptr);
void QAbstractXmlReceiver_StartElement(void* ptr, void* name);
void QAbstractXmlReceiver_StartOfSequence(void* ptr);
void QAbstractXmlReceiver_DestroyQAbstractXmlReceiver(void* ptr);

#ifdef __cplusplus
}
#endif