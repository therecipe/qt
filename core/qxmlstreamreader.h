#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QXmlStreamReader_NamespaceProcessing(QtObjectPtr ptr);
void QXmlStreamReader_SetNamespaceProcessing(QtObjectPtr ptr, int v);
QtObjectPtr QXmlStreamReader_NewQXmlStreamReader();
QtObjectPtr QXmlStreamReader_NewQXmlStreamReader2(QtObjectPtr device);
QtObjectPtr QXmlStreamReader_NewQXmlStreamReader3(QtObjectPtr data);
QtObjectPtr QXmlStreamReader_NewQXmlStreamReader4(char* data);
QtObjectPtr QXmlStreamReader_NewQXmlStreamReader5(char* data);
void QXmlStreamReader_AddData(QtObjectPtr ptr, QtObjectPtr data);
void QXmlStreamReader_AddData2(QtObjectPtr ptr, char* data);
void QXmlStreamReader_AddData3(QtObjectPtr ptr, char* data);
void QXmlStreamReader_AddExtraNamespaceDeclaration(QtObjectPtr ptr, QtObjectPtr extraNamespaceDeclaration);
int QXmlStreamReader_AtEnd(QtObjectPtr ptr);
void QXmlStreamReader_Clear(QtObjectPtr ptr);
QtObjectPtr QXmlStreamReader_Device(QtObjectPtr ptr);
QtObjectPtr QXmlStreamReader_EntityResolver(QtObjectPtr ptr);
int QXmlStreamReader_Error(QtObjectPtr ptr);
char* QXmlStreamReader_ErrorString(QtObjectPtr ptr);
int QXmlStreamReader_HasError(QtObjectPtr ptr);
int QXmlStreamReader_IsCDATA(QtObjectPtr ptr);
int QXmlStreamReader_IsCharacters(QtObjectPtr ptr);
int QXmlStreamReader_IsComment(QtObjectPtr ptr);
int QXmlStreamReader_IsDTD(QtObjectPtr ptr);
int QXmlStreamReader_IsEndDocument(QtObjectPtr ptr);
int QXmlStreamReader_IsEndElement(QtObjectPtr ptr);
int QXmlStreamReader_IsEntityReference(QtObjectPtr ptr);
int QXmlStreamReader_IsProcessingInstruction(QtObjectPtr ptr);
int QXmlStreamReader_IsStandaloneDocument(QtObjectPtr ptr);
int QXmlStreamReader_IsStartDocument(QtObjectPtr ptr);
int QXmlStreamReader_IsStartElement(QtObjectPtr ptr);
int QXmlStreamReader_IsWhitespace(QtObjectPtr ptr);
void QXmlStreamReader_RaiseError(QtObjectPtr ptr, char* message);
char* QXmlStreamReader_ReadElementText(QtObjectPtr ptr, int behaviour);
int QXmlStreamReader_ReadNext(QtObjectPtr ptr);
int QXmlStreamReader_ReadNextStartElement(QtObjectPtr ptr);
void QXmlStreamReader_SetDevice(QtObjectPtr ptr, QtObjectPtr device);
void QXmlStreamReader_SetEntityResolver(QtObjectPtr ptr, QtObjectPtr resolver);
void QXmlStreamReader_SkipCurrentElement(QtObjectPtr ptr);
char* QXmlStreamReader_TokenString(QtObjectPtr ptr);
int QXmlStreamReader_TokenType(QtObjectPtr ptr);
void QXmlStreamReader_DestroyQXmlStreamReader(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif