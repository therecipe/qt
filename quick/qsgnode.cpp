#include "qsgnode.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSGNode>
#include "_cgo_export.h"

class MyQSGNode: public QSGNode {
public:
};

void* QSGNode_ChildAtIndex(void* ptr, int i){
	return static_cast<QSGNode*>(ptr)->childAtIndex(i);
}

int QSGNode_ChildCount(void* ptr){
	return static_cast<QSGNode*>(ptr)->childCount();
}

void* QSGNode_NewQSGNode(){
	return new QSGNode();
}

void QSGNode_AppendChildNode(void* ptr, void* node){
	static_cast<QSGNode*>(ptr)->appendChildNode(static_cast<QSGNode*>(node));
}

void* QSGNode_FirstChild(void* ptr){
	return static_cast<QSGNode*>(ptr)->firstChild();
}

int QSGNode_Flags(void* ptr){
	return static_cast<QSGNode*>(ptr)->flags();
}

void QSGNode_InsertChildNodeAfter(void* ptr, void* node, void* after){
	static_cast<QSGNode*>(ptr)->insertChildNodeAfter(static_cast<QSGNode*>(node), static_cast<QSGNode*>(after));
}

void QSGNode_InsertChildNodeBefore(void* ptr, void* node, void* before){
	static_cast<QSGNode*>(ptr)->insertChildNodeBefore(static_cast<QSGNode*>(node), static_cast<QSGNode*>(before));
}

int QSGNode_IsSubtreeBlocked(void* ptr){
	return static_cast<QSGNode*>(ptr)->isSubtreeBlocked();
}

void* QSGNode_LastChild(void* ptr){
	return static_cast<QSGNode*>(ptr)->lastChild();
}

void QSGNode_MarkDirty(void* ptr, int bits){
	static_cast<QSGNode*>(ptr)->markDirty(static_cast<QSGNode::DirtyStateBit>(bits));
}

void* QSGNode_NextSibling(void* ptr){
	return static_cast<QSGNode*>(ptr)->nextSibling();
}

void* QSGNode_Parent(void* ptr){
	return static_cast<QSGNode*>(ptr)->parent();
}

void QSGNode_PrependChildNode(void* ptr, void* node){
	static_cast<QSGNode*>(ptr)->prependChildNode(static_cast<QSGNode*>(node));
}

void QSGNode_Preprocess(void* ptr){
	static_cast<QSGNode*>(ptr)->preprocess();
}

void* QSGNode_PreviousSibling(void* ptr){
	return static_cast<QSGNode*>(ptr)->previousSibling();
}

void QSGNode_RemoveAllChildNodes(void* ptr){
	static_cast<QSGNode*>(ptr)->removeAllChildNodes();
}

void QSGNode_RemoveChildNode(void* ptr, void* node){
	static_cast<QSGNode*>(ptr)->removeChildNode(static_cast<QSGNode*>(node));
}

void QSGNode_SetFlag(void* ptr, int f, int enabled){
	static_cast<QSGNode*>(ptr)->setFlag(static_cast<QSGNode::Flag>(f), enabled != 0);
}

void QSGNode_SetFlags(void* ptr, int f, int enabled){
	static_cast<QSGNode*>(ptr)->setFlags(static_cast<QSGNode::Flag>(f), enabled != 0);
}

int QSGNode_Type(void* ptr){
	return static_cast<QSGNode*>(ptr)->type();
}

void QSGNode_DestroyQSGNode(void* ptr){
	static_cast<QSGNode*>(ptr)->~QSGNode();
}

