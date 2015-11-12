#include "qsgclipnode.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRectF>
#include <QRect>
#include <QSGClipNode>
#include "_cgo_export.h"

class MyQSGClipNode: public QSGClipNode {
public:
};

void* QSGClipNode_NewQSGClipNode(){
	return new QSGClipNode();
}

int QSGClipNode_IsRectangular(void* ptr){
	return static_cast<QSGClipNode*>(ptr)->isRectangular();
}

void QSGClipNode_SetClipRect(void* ptr, void* rect){
	static_cast<QSGClipNode*>(ptr)->setClipRect(*static_cast<QRectF*>(rect));
}

void QSGClipNode_SetIsRectangular(void* ptr, int rectHint){
	static_cast<QSGClipNode*>(ptr)->setIsRectangular(rectHint != 0);
}

void QSGClipNode_DestroyQSGClipNode(void* ptr){
	static_cast<QSGClipNode*>(ptr)->~QSGClipNode();
}

