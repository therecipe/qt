#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDomNode_NewQDomNode();
QtObjectPtr QDomNode_NewQDomNode2(QtObjectPtr n);
void QDomNode_Clear(QtObjectPtr ptr);
int QDomNode_ColumnNumber(QtObjectPtr ptr);
int QDomNode_HasAttributes(QtObjectPtr ptr);
int QDomNode_HasChildNodes(QtObjectPtr ptr);
int QDomNode_IsAttr(QtObjectPtr ptr);
int QDomNode_IsCDATASection(QtObjectPtr ptr);
int QDomNode_IsCharacterData(QtObjectPtr ptr);
int QDomNode_IsComment(QtObjectPtr ptr);
int QDomNode_IsDocument(QtObjectPtr ptr);
int QDomNode_IsDocumentFragment(QtObjectPtr ptr);
int QDomNode_IsDocumentType(QtObjectPtr ptr);
int QDomNode_IsElement(QtObjectPtr ptr);
int QDomNode_IsEntity(QtObjectPtr ptr);
int QDomNode_IsEntityReference(QtObjectPtr ptr);
int QDomNode_IsNotation(QtObjectPtr ptr);
int QDomNode_IsNull(QtObjectPtr ptr);
int QDomNode_IsProcessingInstruction(QtObjectPtr ptr);
int QDomNode_IsSupported(QtObjectPtr ptr, char* feature, char* version);
int QDomNode_IsText(QtObjectPtr ptr);
int QDomNode_LineNumber(QtObjectPtr ptr);
char* QDomNode_LocalName(QtObjectPtr ptr);
char* QDomNode_NamespaceURI(QtObjectPtr ptr);
char* QDomNode_NodeName(QtObjectPtr ptr);
int QDomNode_NodeType(QtObjectPtr ptr);
char* QDomNode_NodeValue(QtObjectPtr ptr);
void QDomNode_Normalize(QtObjectPtr ptr);
char* QDomNode_Prefix(QtObjectPtr ptr);
void QDomNode_Save(QtObjectPtr ptr, QtObjectPtr stream, int indent, int encodingPolicy);
void QDomNode_SetNodeValue(QtObjectPtr ptr, char* v);
void QDomNode_SetPrefix(QtObjectPtr ptr, char* pre);
void QDomNode_DestroyQDomNode(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif