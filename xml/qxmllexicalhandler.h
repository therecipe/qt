#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QXmlLexicalHandler_Comment(QtObjectPtr ptr, char* ch);
int QXmlLexicalHandler_EndCDATA(QtObjectPtr ptr);
int QXmlLexicalHandler_EndDTD(QtObjectPtr ptr);
int QXmlLexicalHandler_EndEntity(QtObjectPtr ptr, char* name);
char* QXmlLexicalHandler_ErrorString(QtObjectPtr ptr);
int QXmlLexicalHandler_StartCDATA(QtObjectPtr ptr);
int QXmlLexicalHandler_StartDTD(QtObjectPtr ptr, char* name, char* publicId, char* systemId);
int QXmlLexicalHandler_StartEntity(QtObjectPtr ptr, char* name);
void QXmlLexicalHandler_DestroyQXmlLexicalHandler(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif