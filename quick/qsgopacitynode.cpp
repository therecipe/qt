#include "qsgopacitynode.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSGOpacityNode>
#include "_cgo_export.h"

class MyQSGOpacityNode: public QSGOpacityNode {
public:
};

void* QSGOpacityNode_NewQSGOpacityNode(){
	return new QSGOpacityNode();
}

double QSGOpacityNode_Opacity(void* ptr){
	return static_cast<double>(static_cast<QSGOpacityNode*>(ptr)->opacity());
}

void QSGOpacityNode_SetOpacity(void* ptr, double opacity){
	static_cast<QSGOpacityNode*>(ptr)->setOpacity(static_cast<qreal>(opacity));
}

void QSGOpacityNode_DestroyQSGOpacityNode(void* ptr){
	static_cast<QSGOpacityNode*>(ptr)->~QSGOpacityNode();
}

