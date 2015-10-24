#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QXmlSerializer_NewQXmlSerializer(QtObjectPtr query, QtObjectPtr outputDevice);
void QXmlSerializer_Attribute(QtObjectPtr ptr, QtObjectPtr name, QtObjectPtr value);
void QXmlSerializer_Characters(QtObjectPtr ptr, QtObjectPtr value);
void QXmlSerializer_Comment(QtObjectPtr ptr, char* value);
void QXmlSerializer_EndDocument(QtObjectPtr ptr);
void QXmlSerializer_EndElement(QtObjectPtr ptr);
QtObjectPtr QXmlSerializer_Codec(QtObjectPtr ptr);
void QXmlSerializer_EndOfSequence(QtObjectPtr ptr);
void QXmlSerializer_NamespaceBinding(QtObjectPtr ptr, QtObjectPtr nb);
QtObjectPtr QXmlSerializer_OutputDevice(QtObjectPtr ptr);
void QXmlSerializer_ProcessingInstruction(QtObjectPtr ptr, QtObjectPtr name, char* value);
void QXmlSerializer_SetCodec(QtObjectPtr ptr, QtObjectPtr outputCodec);
void QXmlSerializer_StartDocument(QtObjectPtr ptr);
void QXmlSerializer_StartElement(QtObjectPtr ptr, QtObjectPtr name);
void QXmlSerializer_StartOfSequence(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif