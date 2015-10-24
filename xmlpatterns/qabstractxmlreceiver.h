#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QAbstractXmlReceiver_Attribute(QtObjectPtr ptr, QtObjectPtr name, QtObjectPtr value);
void QAbstractXmlReceiver_Characters(QtObjectPtr ptr, QtObjectPtr value);
void QAbstractXmlReceiver_Comment(QtObjectPtr ptr, char* value);
void QAbstractXmlReceiver_EndDocument(QtObjectPtr ptr);
void QAbstractXmlReceiver_EndElement(QtObjectPtr ptr);
void QAbstractXmlReceiver_EndOfSequence(QtObjectPtr ptr);
void QAbstractXmlReceiver_NamespaceBinding(QtObjectPtr ptr, QtObjectPtr name);
void QAbstractXmlReceiver_ProcessingInstruction(QtObjectPtr ptr, QtObjectPtr target, char* value);
void QAbstractXmlReceiver_StartDocument(QtObjectPtr ptr);
void QAbstractXmlReceiver_StartElement(QtObjectPtr ptr, QtObjectPtr name);
void QAbstractXmlReceiver_StartOfSequence(QtObjectPtr ptr);
void QAbstractXmlReceiver_DestroyQAbstractXmlReceiver(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif