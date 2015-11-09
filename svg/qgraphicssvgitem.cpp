#include "qgraphicssvgitem.h"
#include <QString>
#include <QStyleOptionGraphicsItem>
#include <QSize>
#include <QSvgRenderer>
#include <QStyle>
#include <QPainter>
#include <QStyleOption>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QGraphicsSvgItem>
#include "_cgo_export.h"

class MyQGraphicsSvgItem: public QGraphicsSvgItem {
public:
};

char* QGraphicsSvgItem_ElementId(void* ptr){
	return static_cast<QGraphicsSvgItem*>(ptr)->elementId().toUtf8().data();
}

void QGraphicsSvgItem_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsSvgItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void* QGraphicsSvgItem_Renderer(void* ptr){
	return static_cast<QGraphicsSvgItem*>(ptr)->renderer();
}

void QGraphicsSvgItem_SetElementId(void* ptr, char* id){
	static_cast<QGraphicsSvgItem*>(ptr)->setElementId(QString(id));
}

void QGraphicsSvgItem_SetMaximumCacheSize(void* ptr, void* size){
	static_cast<QGraphicsSvgItem*>(ptr)->setMaximumCacheSize(*static_cast<QSize*>(size));
}

void QGraphicsSvgItem_SetSharedRenderer(void* ptr, void* renderer){
	static_cast<QGraphicsSvgItem*>(ptr)->setSharedRenderer(static_cast<QSvgRenderer*>(renderer));
}

int QGraphicsSvgItem_Type(void* ptr){
	return static_cast<QGraphicsSvgItem*>(ptr)->type();
}

