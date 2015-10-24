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

QtObjectPtr QSGNode_ChildAtIndex(QtObjectPtr ptr, int i){
	return static_cast<QSGNode*>(ptr)->childAtIndex(i);
}

int QSGNode_ChildCount(QtObjectPtr ptr){
	return static_cast<QSGNode*>(ptr)->childCount();
}

QtObjectPtr QSGNode_NewQSGNode(){
	return new QSGNode();
}

void QSGNode_AppendChildNode(QtObjectPtr ptr, QtObjectPtr node){
	static_cast<QSGNode*>(ptr)->appendChildNode(static_cast<QSGNode*>(node));
}

QtObjectPtr QSGNode_FirstChild(QtObjectPtr ptr){
	return static_cast<QSGNode*>(ptr)->firstChild();
}

int QSGNode_Flags(QtObjectPtr ptr){
	return static_cast<QSGNode*>(ptr)->flags();
}

void QSGNode_InsertChildNodeAfter(QtObjectPtr ptr, QtObjectPtr node, QtObjectPtr after){
	static_cast<QSGNode*>(ptr)->insertChildNodeAfter(static_cast<QSGNode*>(node), static_cast<QSGNode*>(after));
}

void QSGNode_InsertChildNodeBefore(QtObjectPtr ptr, QtObjectPtr node, QtObjectPtr before){
	static_cast<QSGNode*>(ptr)->insertChildNodeBefore(static_cast<QSGNode*>(node), static_cast<QSGNode*>(before));
}

int QSGNode_IsSubtreeBlocked(QtObjectPtr ptr){
	return static_cast<QSGNode*>(ptr)->isSubtreeBlocked();
}

QtObjectPtr QSGNode_LastChild(QtObjectPtr ptr){
	return static_cast<QSGNode*>(ptr)->lastChild();
}

void QSGNode_MarkDirty(QtObjectPtr ptr, int bits){
	static_cast<QSGNode*>(ptr)->markDirty(static_cast<QSGNode::DirtyStateBit>(bits));
}

QtObjectPtr QSGNode_NextSibling(QtObjectPtr ptr){
	return static_cast<QSGNode*>(ptr)->nextSibling();
}

QtObjectPtr QSGNode_Parent(QtObjectPtr ptr){
	return static_cast<QSGNode*>(ptr)->parent();
}

void QSGNode_PrependChildNode(QtObjectPtr ptr, QtObjectPtr node){
	static_cast<QSGNode*>(ptr)->prependChildNode(static_cast<QSGNode*>(node));
}

void QSGNode_Preprocess(QtObjectPtr ptr){
	static_cast<QSGNode*>(ptr)->preprocess();
}

QtObjectPtr QSGNode_PreviousSibling(QtObjectPtr ptr){
	return static_cast<QSGNode*>(ptr)->previousSibling();
}

void QSGNode_RemoveAllChildNodes(QtObjectPtr ptr){
	static_cast<QSGNode*>(ptr)->removeAllChildNodes();
}

void QSGNode_RemoveChildNode(QtObjectPtr ptr, QtObjectPtr node){
	static_cast<QSGNode*>(ptr)->removeChildNode(static_cast<QSGNode*>(node));
}

void QSGNode_SetFlag(QtObjectPtr ptr, int f, int enabled){
	static_cast<QSGNode*>(ptr)->setFlag(static_cast<QSGNode::Flag>(f), enabled != 0);
}

void QSGNode_SetFlags(QtObjectPtr ptr, int f, int enabled){
	static_cast<QSGNode*>(ptr)->setFlags(static_cast<QSGNode::Flag>(f), enabled != 0);
}

int QSGNode_Type(QtObjectPtr ptr){
	return static_cast<QSGNode*>(ptr)->type();
}

void QSGNode_DestroyQSGNode(QtObjectPtr ptr){
	static_cast<QSGNode*>(ptr)->~QSGNode();
}

