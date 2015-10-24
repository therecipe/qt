#include "qgraphicsvideoitem.h"
#include <QWidget>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QPointF>
#include <QSize>
#include <QStyleOption>
#include <QSizeF>
#include <QGraphicsItem>
#include <QString>
#include <QVariant>
#include <QStyleOptionGraphicsItem>
#include <QPainter>
#include <QPoint>
#include <QGraphicsVideoItem>
#include "_cgo_export.h"

class MyQGraphicsVideoItem: public QGraphicsVideoItem {
public:
};

QtObjectPtr QGraphicsVideoItem_NewQGraphicsVideoItem(QtObjectPtr parent){
	return new QGraphicsVideoItem(static_cast<QGraphicsItem*>(parent));
}

int QGraphicsVideoItem_AspectRatioMode(QtObjectPtr ptr){
	return static_cast<QGraphicsVideoItem*>(ptr)->aspectRatioMode();
}

QtObjectPtr QGraphicsVideoItem_MediaObject(QtObjectPtr ptr){
	return static_cast<QGraphicsVideoItem*>(ptr)->mediaObject();
}

void QGraphicsVideoItem_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget){
	static_cast<QGraphicsVideoItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsVideoItem_SetAspectRatioMode(QtObjectPtr ptr, int mode){
	static_cast<QGraphicsVideoItem*>(ptr)->setAspectRatioMode(static_cast<Qt::AspectRatioMode>(mode));
}

void QGraphicsVideoItem_SetOffset(QtObjectPtr ptr, QtObjectPtr offset){
	static_cast<QGraphicsVideoItem*>(ptr)->setOffset(*static_cast<QPointF*>(offset));
}

void QGraphicsVideoItem_SetSize(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QGraphicsVideoItem*>(ptr)->setSize(*static_cast<QSizeF*>(size));
}

void QGraphicsVideoItem_DestroyQGraphicsVideoItem(QtObjectPtr ptr){
	static_cast<QGraphicsVideoItem*>(ptr)->~QGraphicsVideoItem();
}

