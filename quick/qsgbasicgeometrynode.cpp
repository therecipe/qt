#include "qsgbasicgeometrynode.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSGGeometry>
#include <QSGBasicGeometryNode>
#include "_cgo_export.h"

class MyQSGBasicGeometryNode: public QSGBasicGeometryNode {
public:
};

QtObjectPtr QSGBasicGeometryNode_Geometry2(QtObjectPtr ptr){
	return static_cast<QSGBasicGeometryNode*>(ptr)->geometry();
}

QtObjectPtr QSGBasicGeometryNode_Geometry(QtObjectPtr ptr){
	return const_cast<QSGGeometry*>(static_cast<QSGBasicGeometryNode*>(ptr)->geometry());
}

void QSGBasicGeometryNode_SetGeometry(QtObjectPtr ptr, QtObjectPtr geometry){
	static_cast<QSGBasicGeometryNode*>(ptr)->setGeometry(static_cast<QSGGeometry*>(geometry));
}

void QSGBasicGeometryNode_DestroyQSGBasicGeometryNode(QtObjectPtr ptr){
	static_cast<QSGBasicGeometryNode*>(ptr)->~QSGBasicGeometryNode();
}

