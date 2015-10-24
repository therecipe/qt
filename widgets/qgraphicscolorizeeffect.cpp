#include "qgraphicscolorizeeffect.h"
#include <QColor>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsColorizeEffect>
#include "_cgo_export.h"

class MyQGraphicsColorizeEffect: public QGraphicsColorizeEffect {
public:
};

void QGraphicsColorizeEffect_SetColor(QtObjectPtr ptr, QtObjectPtr c){
	QMetaObject::invokeMethod(static_cast<QGraphicsColorizeEffect*>(ptr), "setColor", Q_ARG(QColor, *static_cast<QColor*>(c)));
}

QtObjectPtr QGraphicsColorizeEffect_NewQGraphicsColorizeEffect(QtObjectPtr parent){
	return new QGraphicsColorizeEffect(static_cast<QObject*>(parent));
}

void QGraphicsColorizeEffect_DestroyQGraphicsColorizeEffect(QtObjectPtr ptr){
	static_cast<QGraphicsColorizeEffect*>(ptr)->~QGraphicsColorizeEffect();
}

