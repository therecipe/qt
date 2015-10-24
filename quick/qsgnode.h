#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSGNode_ChildAtIndex(QtObjectPtr ptr, int i);
int QSGNode_ChildCount(QtObjectPtr ptr);
QtObjectPtr QSGNode_NewQSGNode();
void QSGNode_AppendChildNode(QtObjectPtr ptr, QtObjectPtr node);
QtObjectPtr QSGNode_FirstChild(QtObjectPtr ptr);
int QSGNode_Flags(QtObjectPtr ptr);
void QSGNode_InsertChildNodeAfter(QtObjectPtr ptr, QtObjectPtr node, QtObjectPtr after);
void QSGNode_InsertChildNodeBefore(QtObjectPtr ptr, QtObjectPtr node, QtObjectPtr before);
int QSGNode_IsSubtreeBlocked(QtObjectPtr ptr);
QtObjectPtr QSGNode_LastChild(QtObjectPtr ptr);
void QSGNode_MarkDirty(QtObjectPtr ptr, int bits);
QtObjectPtr QSGNode_NextSibling(QtObjectPtr ptr);
QtObjectPtr QSGNode_Parent(QtObjectPtr ptr);
void QSGNode_PrependChildNode(QtObjectPtr ptr, QtObjectPtr node);
void QSGNode_Preprocess(QtObjectPtr ptr);
QtObjectPtr QSGNode_PreviousSibling(QtObjectPtr ptr);
void QSGNode_RemoveAllChildNodes(QtObjectPtr ptr);
void QSGNode_RemoveChildNode(QtObjectPtr ptr, QtObjectPtr node);
void QSGNode_SetFlag(QtObjectPtr ptr, int f, int enabled);
void QSGNode_SetFlags(QtObjectPtr ptr, int f, int enabled);
int QSGNode_Type(QtObjectPtr ptr);
void QSGNode_DestroyQSGNode(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif