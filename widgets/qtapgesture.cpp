#include "qtapgesture.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QPointF>
#include <QTapGesture>
#include "_cgo_export.h"

class MyQTapGesture: public QTapGesture {
public:
};

void QTapGesture_SetPosition(QtObjectPtr ptr, QtObjectPtr pos){
	static_cast<QTapGesture*>(ptr)->setPosition(*static_cast<QPointF*>(pos));
}

void QTapGesture_DestroyQTapGesture(QtObjectPtr ptr){
	static_cast<QTapGesture*>(ptr)->~QTapGesture();
}

