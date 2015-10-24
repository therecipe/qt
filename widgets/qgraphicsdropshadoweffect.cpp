#include "qgraphicsdropshadoweffect.h"
#include <QMetaObject>
#include <QObject>
#include <QPoint>
#include <QVariant>
#include <QModelIndex>
#include <QPointF>
#include <QColor>
#include <QString>
#include <QUrl>
#include <QGraphicsDropShadowEffect>
#include "_cgo_export.h"

class MyQGraphicsDropShadowEffect: public QGraphicsDropShadowEffect {
public:
};

void QGraphicsDropShadowEffect_SetColor(QtObjectPtr ptr, QtObjectPtr color){
	QMetaObject::invokeMethod(static_cast<QGraphicsDropShadowEffect*>(ptr), "setColor", Q_ARG(QColor, *static_cast<QColor*>(color)));
}

void QGraphicsDropShadowEffect_SetOffset(QtObjectPtr ptr, QtObjectPtr ofs){
	QMetaObject::invokeMethod(static_cast<QGraphicsDropShadowEffect*>(ptr), "setOffset", Q_ARG(QPointF, *static_cast<QPointF*>(ofs)));
}

QtObjectPtr QGraphicsDropShadowEffect_NewQGraphicsDropShadowEffect(QtObjectPtr parent){
	return new QGraphicsDropShadowEffect(static_cast<QObject*>(parent));
}

void QGraphicsDropShadowEffect_DestroyQGraphicsDropShadowEffect(QtObjectPtr ptr){
	static_cast<QGraphicsDropShadowEffect*>(ptr)->~QGraphicsDropShadowEffect();
}

