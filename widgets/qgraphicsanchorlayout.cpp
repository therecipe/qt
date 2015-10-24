#include "qgraphicsanchorlayout.h"
#include <QString>
#include <QVariant>
#include <QRect>
#include <QRectF>
#include <QGraphicsLayout>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsLayoutItem>
#include <QGraphicsAnchor>
#include <QGraphicsAnchorLayout>
#include "_cgo_export.h"

class MyQGraphicsAnchorLayout: public QGraphicsAnchorLayout {
public:
};

QtObjectPtr QGraphicsAnchorLayout_NewQGraphicsAnchorLayout(QtObjectPtr parent){
	return new QGraphicsAnchorLayout(static_cast<QGraphicsLayoutItem*>(parent));
}

QtObjectPtr QGraphicsAnchorLayout_AddAnchor(QtObjectPtr ptr, QtObjectPtr firstItem, int firstEdge, QtObjectPtr secondItem, int secondEdge){
	return static_cast<QGraphicsAnchorLayout*>(ptr)->addAnchor(static_cast<QGraphicsLayoutItem*>(firstItem), static_cast<Qt::AnchorPoint>(firstEdge), static_cast<QGraphicsLayoutItem*>(secondItem), static_cast<Qt::AnchorPoint>(secondEdge));
}

void QGraphicsAnchorLayout_AddAnchors(QtObjectPtr ptr, QtObjectPtr firstItem, QtObjectPtr secondItem, int orientations){
	static_cast<QGraphicsAnchorLayout*>(ptr)->addAnchors(static_cast<QGraphicsLayoutItem*>(firstItem), static_cast<QGraphicsLayoutItem*>(secondItem), static_cast<Qt::Orientation>(orientations));
}

void QGraphicsAnchorLayout_AddCornerAnchors(QtObjectPtr ptr, QtObjectPtr firstItem, int firstCorner, QtObjectPtr secondItem, int secondCorner){
	static_cast<QGraphicsAnchorLayout*>(ptr)->addCornerAnchors(static_cast<QGraphicsLayoutItem*>(firstItem), static_cast<Qt::Corner>(firstCorner), static_cast<QGraphicsLayoutItem*>(secondItem), static_cast<Qt::Corner>(secondCorner));
}

QtObjectPtr QGraphicsAnchorLayout_Anchor(QtObjectPtr ptr, QtObjectPtr firstItem, int firstEdge, QtObjectPtr secondItem, int secondEdge){
	return static_cast<QGraphicsAnchorLayout*>(ptr)->anchor(static_cast<QGraphicsLayoutItem*>(firstItem), static_cast<Qt::AnchorPoint>(firstEdge), static_cast<QGraphicsLayoutItem*>(secondItem), static_cast<Qt::AnchorPoint>(secondEdge));
}

int QGraphicsAnchorLayout_Count(QtObjectPtr ptr){
	return static_cast<QGraphicsAnchorLayout*>(ptr)->count();
}

void QGraphicsAnchorLayout_Invalidate(QtObjectPtr ptr){
	static_cast<QGraphicsAnchorLayout*>(ptr)->invalidate();
}

QtObjectPtr QGraphicsAnchorLayout_ItemAt(QtObjectPtr ptr, int index){
	return static_cast<QGraphicsAnchorLayout*>(ptr)->itemAt(index);
}

void QGraphicsAnchorLayout_RemoveAt(QtObjectPtr ptr, int index){
	static_cast<QGraphicsAnchorLayout*>(ptr)->removeAt(index);
}

void QGraphicsAnchorLayout_SetGeometry(QtObjectPtr ptr, QtObjectPtr geom){
	static_cast<QGraphicsAnchorLayout*>(ptr)->setGeometry(*static_cast<QRectF*>(geom));
}

void QGraphicsAnchorLayout_DestroyQGraphicsAnchorLayout(QtObjectPtr ptr){
	static_cast<QGraphicsAnchorLayout*>(ptr)->~QGraphicsAnchorLayout();
}

