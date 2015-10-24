#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QXmlReader_DTDHandler(QtObjectPtr ptr);
QtObjectPtr QXmlReader_ContentHandler(QtObjectPtr ptr);
QtObjectPtr QXmlReader_DeclHandler(QtObjectPtr ptr);
QtObjectPtr QXmlReader_EntityResolver(QtObjectPtr ptr);
QtObjectPtr QXmlReader_ErrorHandler(QtObjectPtr ptr);
int QXmlReader_Feature(QtObjectPtr ptr, char* name, int ok);
int QXmlReader_HasFeature(QtObjectPtr ptr, char* name);
int QXmlReader_HasProperty(QtObjectPtr ptr, char* name);
QtObjectPtr QXmlReader_LexicalHandler(QtObjectPtr ptr);
int QXmlReader_Parse(QtObjectPtr ptr, QtObjectPtr input);
void QXmlReader_Property(QtObjectPtr ptr, char* name, int ok);
void QXmlReader_SetContentHandler(QtObjectPtr ptr, QtObjectPtr handler);
void QXmlReader_SetDTDHandler(QtObjectPtr ptr, QtObjectPtr handler);
void QXmlReader_SetDeclHandler(QtObjectPtr ptr, QtObjectPtr handler);
void QXmlReader_SetEntityResolver(QtObjectPtr ptr, QtObjectPtr handler);
void QXmlReader_SetErrorHandler(QtObjectPtr ptr, QtObjectPtr handler);
void QXmlReader_SetFeature(QtObjectPtr ptr, char* name, int value);
void QXmlReader_SetLexicalHandler(QtObjectPtr ptr, QtObjectPtr handler);
void QXmlReader_DestroyQXmlReader(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif