#include "qsgsimplerectnode.h"
#include <QRect>
#include <QRectF>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QColor>
#include <QSGSimpleRectNode>
#include "_cgo_export.h"

class MyQSGSimpleRectNode: public QSGSimpleRectNode {
public:
};

QtObjectPtr QSGSimpleRectNode_NewQSGSimpleRectNode2(){
	return new QSGSimpleRectNode();
}

QtObjectPtr QSGSimpleRectNode_NewQSGSimpleRectNode(QtObjectPtr rect, QtObjectPtr color){
	return new QSGSimpleRectNode(*static_cast<QRectF*>(rect), *static_cast<QColor*>(color));
}

void QSGSimpleRectNode_SetColor(QtObjectPtr ptr, QtObjectPtr color){
	static_cast<QSGSimpleRectNode*>(ptr)->setColor(*static_cast<QColor*>(color));
}

void QSGSimpleRectNode_SetRect(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QSGSimpleRectNode*>(ptr)->setRect(*static_cast<QRectF*>(rect));
}

