#include "qgraphicsopacityeffect.h"
#include <QBrush>
#include <QObject>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsOpacityEffect>
#include "_cgo_export.h"

class MyQGraphicsOpacityEffect: public QGraphicsOpacityEffect {
public:
};

void QGraphicsOpacityEffect_SetOpacityMask(QtObjectPtr ptr, QtObjectPtr mask){
	QMetaObject::invokeMethod(static_cast<QGraphicsOpacityEffect*>(ptr), "setOpacityMask", Q_ARG(QBrush, *static_cast<QBrush*>(mask)));
}

QtObjectPtr QGraphicsOpacityEffect_NewQGraphicsOpacityEffect(QtObjectPtr parent){
	return new QGraphicsOpacityEffect(static_cast<QObject*>(parent));
}

void QGraphicsOpacityEffect_DestroyQGraphicsOpacityEffect(QtObjectPtr ptr){
	static_cast<QGraphicsOpacityEffect*>(ptr)->~QGraphicsOpacityEffect();
}

