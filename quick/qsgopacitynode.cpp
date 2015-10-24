#include "qsgopacitynode.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QSGOpacityNode>
#include "_cgo_export.h"

class MyQSGOpacityNode: public QSGOpacityNode {
public:
};

QtObjectPtr QSGOpacityNode_NewQSGOpacityNode(){
	return new QSGOpacityNode();
}

void QSGOpacityNode_DestroyQSGOpacityNode(QtObjectPtr ptr){
	static_cast<QSGOpacityNode*>(ptr)->~QSGOpacityNode();
}

