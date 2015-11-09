#include "qsgtransformnode.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMatrix4x4>
#include <QSGTransformNode>
#include "_cgo_export.h"

class MyQSGTransformNode: public QSGTransformNode {
public:
};

void* QSGTransformNode_NewQSGTransformNode(){
	return new QSGTransformNode();
}

void QSGTransformNode_SetMatrix(void* ptr, void* matrix){
	static_cast<QSGTransformNode*>(ptr)->setMatrix(*static_cast<QMatrix4x4*>(matrix));
}

void QSGTransformNode_DestroyQSGTransformNode(void* ptr){
	static_cast<QSGTransformNode*>(ptr)->~QSGTransformNode();
}

