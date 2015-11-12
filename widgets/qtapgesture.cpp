#include "qtapgesture.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPointF>
#include <QPoint>
#include <QTapGesture>
#include "_cgo_export.h"

class MyQTapGesture: public QTapGesture {
public:
};

void QTapGesture_SetPosition(void* ptr, void* pos){
	static_cast<QTapGesture*>(ptr)->setPosition(*static_cast<QPointF*>(pos));
}

void QTapGesture_DestroyQTapGesture(void* ptr){
	static_cast<QTapGesture*>(ptr)->~QTapGesture();
}

