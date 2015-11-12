#include "qsgsimplerectnode.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRectF>
#include <QColor>
#include <QRect>
#include <QString>
#include <QSGSimpleRectNode>
#include "_cgo_export.h"

class MyQSGSimpleRectNode: public QSGSimpleRectNode {
public:
};

void* QSGSimpleRectNode_NewQSGSimpleRectNode2(){
	return new QSGSimpleRectNode();
}

void* QSGSimpleRectNode_NewQSGSimpleRectNode(void* rect, void* color){
	return new QSGSimpleRectNode(*static_cast<QRectF*>(rect), *static_cast<QColor*>(color));
}

void* QSGSimpleRectNode_Color(void* ptr){
	return new QColor(static_cast<QSGSimpleRectNode*>(ptr)->color());
}

void QSGSimpleRectNode_SetColor(void* ptr, void* color){
	static_cast<QSGSimpleRectNode*>(ptr)->setColor(*static_cast<QColor*>(color));
}

void QSGSimpleRectNode_SetRect(void* ptr, void* rect){
	static_cast<QSGSimpleRectNode*>(ptr)->setRect(*static_cast<QRectF*>(rect));
}

void QSGSimpleRectNode_SetRect2(void* ptr, double x, double y, double w, double h){
	static_cast<QSGSimpleRectNode*>(ptr)->setRect(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(w), static_cast<qreal>(h));
}

