#ifdef __cplusplus
extern "C" {
#endif

int QXmlStreamReader_NamespaceProcessing(void* ptr);
void QXmlStreamReader_SetNamespaceProcessing(void* ptr, int v);
void* QXmlStreamReader_NewQXmlStreamReader();
void* QXmlStreamReader_NewQXmlStreamReader2(void* device);
void* QXmlStreamReader_NewQXmlStreamReader3(void* data);
void* QXmlStreamReader_NewQXmlStreamReader4(char* data);
void* QXmlStreamReader_NewQXmlStreamReader5(char* data);
void QXmlStreamReader_AddData(void* ptr, void* data);
void QXmlStreamReader_AddData2(void* ptr, char* data);
void QXmlStreamReader_AddData3(void* ptr, char* data);
void QXmlStreamReader_AddExtraNamespaceDeclaration(void* ptr, void* extraNamespaceDeclaration);
int QXmlStreamReader_AtEnd(void* ptr);
void QXmlStreamReader_Clear(void* ptr);
void* QXmlStreamReader_Device(void* ptr);
void* QXmlStreamReader_DocumentEncoding(void* ptr);
void* QXmlStreamReader_DocumentVersion(void* ptr);
void* QXmlStreamReader_DtdName(void* ptr);
void* QXmlStreamReader_DtdPublicId(void* ptr);
void* QXmlStreamReader_DtdSystemId(void* ptr);
void* QXmlStreamReader_EntityResolver(void* ptr);
int QXmlStreamReader_Error(void* ptr);
char* QXmlStreamReader_ErrorString(void* ptr);
int QXmlStreamReader_HasError(void* ptr);
int QXmlStreamReader_IsCDATA(void* ptr);
int QXmlStreamReader_IsCharacters(void* ptr);
int QXmlStreamReader_IsComment(void* ptr);
int QXmlStreamReader_IsDTD(void* ptr);
int QXmlStreamReader_IsEndDocument(void* ptr);
int QXmlStreamReader_IsEndElement(void* ptr);
int QXmlStreamReader_IsEntityReference(void* ptr);
int QXmlStreamReader_IsProcessingInstruction(void* ptr);
int QXmlStreamReader_IsStandaloneDocument(void* ptr);
int QXmlStreamReader_IsStartDocument(void* ptr);
int QXmlStreamReader_IsStartElement(void* ptr);
int QXmlStreamReader_IsWhitespace(void* ptr);
void* QXmlStreamReader_Name(void* ptr);
void* QXmlStreamReader_NamespaceUri(void* ptr);
void* QXmlStreamReader_Prefix(void* ptr);
void* QXmlStreamReader_ProcessingInstructionData(void* ptr);
void* QXmlStreamReader_ProcessingInstructionTarget(void* ptr);
void* QXmlStreamReader_QualifiedName(void* ptr);
void QXmlStreamReader_RaiseError(void* ptr, char* message);
char* QXmlStreamReader_ReadElementText(void* ptr, int behaviour);
int QXmlStreamReader_ReadNext(void* ptr);
int QXmlStreamReader_ReadNextStartElement(void* ptr);
void QXmlStreamReader_SetDevice(void* ptr, void* device);
void QXmlStreamReader_SetEntityResolver(void* ptr, void* resolver);
void QXmlStreamReader_SkipCurrentElement(void* ptr);
void* QXmlStreamReader_Text(void* ptr);
char* QXmlStreamReader_TokenString(void* ptr);
int QXmlStreamReader_TokenType(void* ptr);
void QXmlStreamReader_DestroyQXmlStreamReader(void* ptr);

#ifdef __cplusplus
}
#endif