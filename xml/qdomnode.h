#ifdef __cplusplus
extern "C" {
#endif

void* QDomNode_NewQDomNode();
void* QDomNode_NewQDomNode2(void* n);
void QDomNode_Clear(void* ptr);
int QDomNode_ColumnNumber(void* ptr);
int QDomNode_HasAttributes(void* ptr);
int QDomNode_HasChildNodes(void* ptr);
int QDomNode_IsAttr(void* ptr);
int QDomNode_IsCDATASection(void* ptr);
int QDomNode_IsCharacterData(void* ptr);
int QDomNode_IsComment(void* ptr);
int QDomNode_IsDocument(void* ptr);
int QDomNode_IsDocumentFragment(void* ptr);
int QDomNode_IsDocumentType(void* ptr);
int QDomNode_IsElement(void* ptr);
int QDomNode_IsEntity(void* ptr);
int QDomNode_IsEntityReference(void* ptr);
int QDomNode_IsNotation(void* ptr);
int QDomNode_IsNull(void* ptr);
int QDomNode_IsProcessingInstruction(void* ptr);
int QDomNode_IsSupported(void* ptr, char* feature, char* version);
int QDomNode_IsText(void* ptr);
int QDomNode_LineNumber(void* ptr);
char* QDomNode_LocalName(void* ptr);
char* QDomNode_NamespaceURI(void* ptr);
char* QDomNode_NodeName(void* ptr);
int QDomNode_NodeType(void* ptr);
char* QDomNode_NodeValue(void* ptr);
void QDomNode_Normalize(void* ptr);
char* QDomNode_Prefix(void* ptr);
void QDomNode_Save(void* ptr, void* stream, int indent, int encodingPolicy);
void QDomNode_SetNodeValue(void* ptr, char* v);
void QDomNode_SetPrefix(void* ptr, char* pre);
void QDomNode_DestroyQDomNode(void* ptr);

#ifdef __cplusplus
}
#endif