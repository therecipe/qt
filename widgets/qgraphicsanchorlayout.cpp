#include "qgraphicsanchorlayout.h"
#include <QRect>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QGraphicsLayoutItem>
#include <QGraphicsLayout>
#include <QModelIndex>
#include <QGraphicsAnchor>
#include <QRectF>
#include <QGraphicsAnchorLayout>
#include "_cgo_export.h"

class MyQGraphicsAnchorLayout: public QGraphicsAnchorLayout {
public:
};

void* QGraphicsAnchorLayout_NewQGraphicsAnchorLayout(void* parent){
	return new QGraphicsAnchorLayout(static_cast<QGraphicsLayoutItem*>(parent));
}

void* QGraphicsAnchorLayout_AddAnchor(void* ptr, void* firstItem, int firstEdge, void* secondItem, int secondEdge){
	return static_cast<QGraphicsAnchorLayout*>(ptr)->addAnchor(static_cast<QGraphicsLayoutItem*>(firstItem), static_cast<Qt::AnchorPoint>(firstEdge), static_cast<QGraphicsLayoutItem*>(secondItem), static_cast<Qt::AnchorPoint>(secondEdge));
}

void QGraphicsAnchorLayout_AddAnchors(void* ptr, void* firstItem, void* secondItem, int orientations){
	static_cast<QGraphicsAnchorLayout*>(ptr)->addAnchors(static_cast<QGraphicsLayoutItem*>(firstItem), static_cast<QGraphicsLayoutItem*>(secondItem), static_cast<Qt::Orientation>(orientations));
}

void QGraphicsAnchorLayout_AddCornerAnchors(void* ptr, void* firstItem, int firstCorner, void* secondItem, int secondCorner){
	static_cast<QGraphicsAnchorLayout*>(ptr)->addCornerAnchors(static_cast<QGraphicsLayoutItem*>(firstItem), static_cast<Qt::Corner>(firstCorner), static_cast<QGraphicsLayoutItem*>(secondItem), static_cast<Qt::Corner>(secondCorner));
}

void* QGraphicsAnchorLayout_Anchor(void* ptr, void* firstItem, int firstEdge, void* secondItem, int secondEdge){
	return static_cast<QGraphicsAnchorLayout*>(ptr)->anchor(static_cast<QGraphicsLayoutItem*>(firstItem), static_cast<Qt::AnchorPoint>(firstEdge), static_cast<QGraphicsLayoutItem*>(secondItem), static_cast<Qt::AnchorPoint>(secondEdge));
}

int QGraphicsAnchorLayout_Count(void* ptr){
	return static_cast<QGraphicsAnchorLayout*>(ptr)->count();
}

double QGraphicsAnchorLayout_HorizontalSpacing(void* ptr){
	return static_cast<double>(static_cast<QGraphicsAnchorLayout*>(ptr)->horizontalSpacing());
}

void QGraphicsAnchorLayout_Invalidate(void* ptr){
	static_cast<QGraphicsAnchorLayout*>(ptr)->invalidate();
}

void* QGraphicsAnchorLayout_ItemAt(void* ptr, int index){
	return static_cast<QGraphicsAnchorLayout*>(ptr)->itemAt(index);
}

void QGraphicsAnchorLayout_RemoveAt(void* ptr, int index){
	static_cast<QGraphicsAnchorLayout*>(ptr)->removeAt(index);
}

void QGraphicsAnchorLayout_SetGeometry(void* ptr, void* geom){
	static_cast<QGraphicsAnchorLayout*>(ptr)->setGeometry(*static_cast<QRectF*>(geom));
}

void QGraphicsAnchorLayout_SetHorizontalSpacing(void* ptr, double spacing){
	static_cast<QGraphicsAnchorLayout*>(ptr)->setHorizontalSpacing(static_cast<qreal>(spacing));
}

void QGraphicsAnchorLayout_SetSpacing(void* ptr, double spacing){
	static_cast<QGraphicsAnchorLayout*>(ptr)->setSpacing(static_cast<qreal>(spacing));
}

void QGraphicsAnchorLayout_SetVerticalSpacing(void* ptr, double spacing){
	static_cast<QGraphicsAnchorLayout*>(ptr)->setVerticalSpacing(static_cast<qreal>(spacing));
}

double QGraphicsAnchorLayout_VerticalSpacing(void* ptr){
	return static_cast<double>(static_cast<QGraphicsAnchorLayout*>(ptr)->verticalSpacing());
}

void QGraphicsAnchorLayout_DestroyQGraphicsAnchorLayout(void* ptr){
	static_cast<QGraphicsAnchorLayout*>(ptr)->~QGraphicsAnchorLayout();
}

