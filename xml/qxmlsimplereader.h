#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QXmlSimpleReader_DTDHandler(QtObjectPtr ptr);
QtObjectPtr QXmlSimpleReader_NewQXmlSimpleReader();
QtObjectPtr QXmlSimpleReader_ContentHandler(QtObjectPtr ptr);
QtObjectPtr QXmlSimpleReader_DeclHandler(QtObjectPtr ptr);
QtObjectPtr QXmlSimpleReader_EntityResolver(QtObjectPtr ptr);
QtObjectPtr QXmlSimpleReader_ErrorHandler(QtObjectPtr ptr);
int QXmlSimpleReader_Feature(QtObjectPtr ptr, char* name, int ok);
int QXmlSimpleReader_HasFeature(QtObjectPtr ptr, char* name);
int QXmlSimpleReader_HasProperty(QtObjectPtr ptr, char* name);
QtObjectPtr QXmlSimpleReader_LexicalHandler(QtObjectPtr ptr);
int QXmlSimpleReader_Parse(QtObjectPtr ptr, QtObjectPtr input);
int QXmlSimpleReader_Parse2(QtObjectPtr ptr, QtObjectPtr input);
int QXmlSimpleReader_Parse3(QtObjectPtr ptr, QtObjectPtr input, int incremental);
int QXmlSimpleReader_ParseContinue(QtObjectPtr ptr);
void QXmlSimpleReader_Property(QtObjectPtr ptr, char* name, int ok);
void QXmlSimpleReader_SetContentHandler(QtObjectPtr ptr, QtObjectPtr handler);
void QXmlSimpleReader_SetDTDHandler(QtObjectPtr ptr, QtObjectPtr handler);
void QXmlSimpleReader_SetDeclHandler(QtObjectPtr ptr, QtObjectPtr handler);
void QXmlSimpleReader_SetEntityResolver(QtObjectPtr ptr, QtObjectPtr handler);
void QXmlSimpleReader_SetErrorHandler(QtObjectPtr ptr, QtObjectPtr handler);
void QXmlSimpleReader_SetFeature(QtObjectPtr ptr, char* name, int enable);
void QXmlSimpleReader_SetLexicalHandler(QtObjectPtr ptr, QtObjectPtr handler);
void QXmlSimpleReader_DestroyQXmlSimpleReader(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif