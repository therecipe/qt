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

QtObjectPtr QSGTransformNode_NewQSGTransformNode(){
	return new QSGTransformNode();
}

void QSGTransformNode_SetMatrix(QtObjectPtr ptr, QtObjectPtr matrix){
	static_cast<QSGTransformNode*>(ptr)->setMatrix(*static_cast<QMatrix4x4*>(matrix));
}

void QSGTransformNode_DestroyQSGTransformNode(QtObjectPtr ptr){
	static_cast<QSGTransformNode*>(ptr)->~QSGTransformNode();
}

