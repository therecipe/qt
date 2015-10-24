#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QXmlStreamWriter_AutoFormattingIndent(QtObjectPtr ptr);
void QXmlStreamWriter_SetAutoFormattingIndent(QtObjectPtr ptr, int spacesOrTabs);
int QXmlStreamWriter_AutoFormatting(QtObjectPtr ptr);
QtObjectPtr QXmlStreamWriter_Codec(QtObjectPtr ptr);
QtObjectPtr QXmlStreamWriter_Device(QtObjectPtr ptr);
int QXmlStreamWriter_HasError(QtObjectPtr ptr);
void QXmlStreamWriter_SetAutoFormatting(QtObjectPtr ptr, int enable);
void QXmlStreamWriter_SetCodec(QtObjectPtr ptr, QtObjectPtr codec);
void QXmlStreamWriter_SetCodec2(QtObjectPtr ptr, char* codecName);
void QXmlStreamWriter_SetDevice(QtObjectPtr ptr, QtObjectPtr device);
void QXmlStreamWriter_WriteAttribute(QtObjectPtr ptr, char* namespaceUri, char* name, char* value);
void QXmlStreamWriter_WriteAttribute2(QtObjectPtr ptr, char* qualifiedName, char* value);
void QXmlStreamWriter_WriteAttribute3(QtObjectPtr ptr, QtObjectPtr attribute);
void QXmlStreamWriter_WriteAttributes(QtObjectPtr ptr, QtObjectPtr attributes);
void QXmlStreamWriter_WriteCDATA(QtObjectPtr ptr, char* text);
void QXmlStreamWriter_WriteCharacters(QtObjectPtr ptr, char* text);
void QXmlStreamWriter_WriteComment(QtObjectPtr ptr, char* text);
void QXmlStreamWriter_WriteCurrentToken(QtObjectPtr ptr, QtObjectPtr reader);
void QXmlStreamWriter_WriteDTD(QtObjectPtr ptr, char* dtd);
void QXmlStreamWriter_WriteDefaultNamespace(QtObjectPtr ptr, char* namespaceUri);
void QXmlStreamWriter_WriteEmptyElement(QtObjectPtr ptr, char* namespaceUri, char* name);
void QXmlStreamWriter_WriteEmptyElement2(QtObjectPtr ptr, char* qualifiedName);
void QXmlStreamWriter_WriteEndDocument(QtObjectPtr ptr);
void QXmlStreamWriter_WriteEndElement(QtObjectPtr ptr);
void QXmlStreamWriter_WriteEntityReference(QtObjectPtr ptr, char* name);
void QXmlStreamWriter_WriteNamespace(QtObjectPtr ptr, char* namespaceUri, char* prefix);
void QXmlStreamWriter_WriteProcessingInstruction(QtObjectPtr ptr, char* target, char* data);
void QXmlStreamWriter_WriteStartDocument3(QtObjectPtr ptr);
void QXmlStreamWriter_WriteStartDocument(QtObjectPtr ptr, char* version);
void QXmlStreamWriter_WriteStartDocument2(QtObjectPtr ptr, char* version, int standalone);
void QXmlStreamWriter_WriteStartElement(QtObjectPtr ptr, char* namespaceUri, char* name);
void QXmlStreamWriter_WriteStartElement2(QtObjectPtr ptr, char* qualifiedName);
void QXmlStreamWriter_WriteTextElement(QtObjectPtr ptr, char* namespaceUri, char* name, char* text);
void QXmlStreamWriter_WriteTextElement2(QtObjectPtr ptr, char* qualifiedName, char* text);
void QXmlStreamWriter_DestroyQXmlStreamWriter(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif