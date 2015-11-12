#include "qsgbasicgeometrynode.h"
#include <QSGGeometry>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSGBasicGeometryNode>
#include "_cgo_export.h"

class MyQSGBasicGeometryNode: public QSGBasicGeometryNode {
public:
};

void* QSGBasicGeometryNode_Geometry2(void* ptr){
	return static_cast<QSGBasicGeometryNode*>(ptr)->geometry();
}

void* QSGBasicGeometryNode_Geometry(void* ptr){
	return const_cast<QSGGeometry*>(static_cast<QSGBasicGeometryNode*>(ptr)->geometry());
}

void QSGBasicGeometryNode_SetGeometry(void* ptr, void* geometry){
	static_cast<QSGBasicGeometryNode*>(ptr)->setGeometry(static_cast<QSGGeometry*>(geometry));
}

void QSGBasicGeometryNode_DestroyQSGBasicGeometryNode(void* ptr){
	static_cast<QSGBasicGeometryNode*>(ptr)->~QSGBasicGeometryNode();
}

