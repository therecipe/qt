#include "qsgclipnode.h"
#include <QRect>
#include <QRectF>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSGClipNode>
#include "_cgo_export.h"

class MyQSGClipNode: public QSGClipNode {
public:
};

QtObjectPtr QSGClipNode_NewQSGClipNode(){
	return new QSGClipNode();
}

int QSGClipNode_IsRectangular(QtObjectPtr ptr){
	return static_cast<QSGClipNode*>(ptr)->isRectangular();
}

void QSGClipNode_SetClipRect(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QSGClipNode*>(ptr)->setClipRect(*static_cast<QRectF*>(rect));
}

void QSGClipNode_SetIsRectangular(QtObjectPtr ptr, int rectHint){
	static_cast<QSGClipNode*>(ptr)->setIsRectangular(rectHint != 0);
}

void QSGClipNode_DestroyQSGClipNode(QtObjectPtr ptr){
	static_cast<QSGClipNode*>(ptr)->~QSGClipNode();
}

