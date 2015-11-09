#ifdef __cplusplus
extern "C" {
#endif

void* QSGNode_ChildAtIndex(void* ptr, int i);
int QSGNode_ChildCount(void* ptr);
void* QSGNode_NewQSGNode();
void QSGNode_AppendChildNode(void* ptr, void* node);
void* QSGNode_FirstChild(void* ptr);
int QSGNode_Flags(void* ptr);
void QSGNode_InsertChildNodeAfter(void* ptr, void* node, void* after);
void QSGNode_InsertChildNodeBefore(void* ptr, void* node, void* before);
int QSGNode_IsSubtreeBlocked(void* ptr);
void* QSGNode_LastChild(void* ptr);
void QSGNode_MarkDirty(void* ptr, int bits);
void* QSGNode_NextSibling(void* ptr);
void* QSGNode_Parent(void* ptr);
void QSGNode_PrependChildNode(void* ptr, void* node);
void QSGNode_Preprocess(void* ptr);
void* QSGNode_PreviousSibling(void* ptr);
void QSGNode_RemoveAllChildNodes(void* ptr);
void QSGNode_RemoveChildNode(void* ptr, void* node);
void QSGNode_SetFlag(void* ptr, int f, int enabled);
void QSGNode_SetFlags(void* ptr, int f, int enabled);
int QSGNode_Type(void* ptr);
void QSGNode_DestroyQSGNode(void* ptr);

#ifdef __cplusplus
}
#endif