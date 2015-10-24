#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QXmlFormatter_NewQXmlFormatter(QtObjectPtr query, QtObjectPtr outputDevice);
void QXmlFormatter_Attribute(QtObjectPtr ptr, QtObjectPtr name, QtObjectPtr value);
void QXmlFormatter_Characters(QtObjectPtr ptr, QtObjectPtr value);
void QXmlFormatter_Comment(QtObjectPtr ptr, char* value);
void QXmlFormatter_EndDocument(QtObjectPtr ptr);
void QXmlFormatter_EndElement(QtObjectPtr ptr);
void QXmlFormatter_EndOfSequence(QtObjectPtr ptr);
int QXmlFormatter_IndentationDepth(QtObjectPtr ptr);
void QXmlFormatter_ProcessingInstruction(QtObjectPtr ptr, QtObjectPtr name, char* value);
void QXmlFormatter_SetIndentationDepth(QtObjectPtr ptr, int depth);
void QXmlFormatter_StartDocument(QtObjectPtr ptr);
void QXmlFormatter_StartElement(QtObjectPtr ptr, QtObjectPtr name);
void QXmlFormatter_StartOfSequence(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif