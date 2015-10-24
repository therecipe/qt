#include "qsggeometrynode.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSGGeometry>
#include <QSGMaterial>
#include <QSGGeometryNode>
#include "_cgo_export.h"

class MyQSGGeometryNode: public QSGGeometryNode {
public:
};

QtObjectPtr QSGGeometryNode_NewQSGGeometryNode(){
	return new QSGGeometryNode();
}

QtObjectPtr QSGGeometryNode_Material(QtObjectPtr ptr){
	return static_cast<QSGGeometryNode*>(ptr)->material();
}

QtObjectPtr QSGGeometryNode_OpaqueMaterial(QtObjectPtr ptr){
	return static_cast<QSGGeometryNode*>(ptr)->opaqueMaterial();
}

void QSGGeometryNode_SetMaterial(QtObjectPtr ptr, QtObjectPtr material){
	static_cast<QSGGeometryNode*>(ptr)->setMaterial(static_cast<QSGMaterial*>(material));
}

void QSGGeometryNode_SetOpaqueMaterial(QtObjectPtr ptr, QtObjectPtr material){
	static_cast<QSGGeometryNode*>(ptr)->setOpaqueMaterial(static_cast<QSGMaterial*>(material));
}

void QSGGeometryNode_DestroyQSGGeometryNode(QtObjectPtr ptr){
	static_cast<QSGGeometryNode*>(ptr)->~QSGGeometryNode();
}

