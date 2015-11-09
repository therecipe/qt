#include "qgraphicsanchor.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSize>
#include <QSizePolicy>
#include <QGraphicsAnchor>
#include "_cgo_export.h"

class MyQGraphicsAnchor: public QGraphicsAnchor {
public:
};

void QGraphicsAnchor_SetSizePolicy(void* ptr, int policy){
	static_cast<QGraphicsAnchor*>(ptr)->setSizePolicy(static_cast<QSizePolicy::Policy>(policy));
}

void QGraphicsAnchor_SetSpacing(void* ptr, double spacing){
	static_cast<QGraphicsAnchor*>(ptr)->setSpacing(static_cast<qreal>(spacing));
}

int QGraphicsAnchor_SizePolicy(void* ptr){
	return static_cast<QGraphicsAnchor*>(ptr)->sizePolicy();
}

double QGraphicsAnchor_Spacing(void* ptr){
	return static_cast<double>(static_cast<QGraphicsAnchor*>(ptr)->spacing());
}

void QGraphicsAnchor_UnsetSpacing(void* ptr){
	static_cast<QGraphicsAnchor*>(ptr)->unsetSpacing();
}

void QGraphicsAnchor_DestroyQGraphicsAnchor(void* ptr){
	static_cast<QGraphicsAnchor*>(ptr)->~QGraphicsAnchor();
}

