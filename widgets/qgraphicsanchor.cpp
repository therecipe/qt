#include "qgraphicsanchor.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSize>
#include <QSizePolicy>
#include <QString>
#include <QGraphicsAnchor>
#include "_cgo_export.h"

class MyQGraphicsAnchor: public QGraphicsAnchor {
public:
};

void QGraphicsAnchor_SetSizePolicy(QtObjectPtr ptr, int policy){
	static_cast<QGraphicsAnchor*>(ptr)->setSizePolicy(static_cast<QSizePolicy::Policy>(policy));
}

int QGraphicsAnchor_SizePolicy(QtObjectPtr ptr){
	return static_cast<QGraphicsAnchor*>(ptr)->sizePolicy();
}

void QGraphicsAnchor_UnsetSpacing(QtObjectPtr ptr){
	static_cast<QGraphicsAnchor*>(ptr)->unsetSpacing();
}

void QGraphicsAnchor_DestroyQGraphicsAnchor(QtObjectPtr ptr){
	static_cast<QGraphicsAnchor*>(ptr)->~QGraphicsAnchor();
}

