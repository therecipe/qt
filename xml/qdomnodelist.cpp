#include "qdomnodelist.h"
#include <QList>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDomNode>
#include <QList>
#include <QString>
#include <QDomNodeList>
#include "_cgo_export.h"

class MyQDomNodeList: public QDomNodeList {
public:
};

QtObjectPtr QDomNodeList_NewQDomNodeList(){
	return new QDomNodeList();
}

QtObjectPtr QDomNodeList_NewQDomNodeList2(QtObjectPtr n){
	return new QDomNodeList(*static_cast<QDomNodeList*>(n));
}

int QDomNodeList_Count(QtObjectPtr ptr){
	return static_cast<QDomNodeList*>(ptr)->count();
}

int QDomNodeList_IsEmpty(QtObjectPtr ptr){
	return static_cast<QDomNodeList*>(ptr)->isEmpty();
}

int QDomNodeList_Length(QtObjectPtr ptr){
	return static_cast<QDomNodeList*>(ptr)->length();
}

int QDomNodeList_Size(QtObjectPtr ptr){
	return static_cast<QDomNodeList*>(ptr)->size();
}

void QDomNodeList_DestroyQDomNodeList(QtObjectPtr ptr){
	static_cast<QDomNodeList*>(ptr)->~QDomNodeList();
}

