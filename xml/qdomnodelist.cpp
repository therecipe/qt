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

void* QDomNodeList_NewQDomNodeList(){
	return new QDomNodeList();
}

void* QDomNodeList_NewQDomNodeList2(void* n){
	return new QDomNodeList(*static_cast<QDomNodeList*>(n));
}

int QDomNodeList_Count(void* ptr){
	return static_cast<QDomNodeList*>(ptr)->count();
}

int QDomNodeList_IsEmpty(void* ptr){
	return static_cast<QDomNodeList*>(ptr)->isEmpty();
}

int QDomNodeList_Length(void* ptr){
	return static_cast<QDomNodeList*>(ptr)->length();
}

int QDomNodeList_Size(void* ptr){
	return static_cast<QDomNodeList*>(ptr)->size();
}

void QDomNodeList_DestroyQDomNodeList(void* ptr){
	static_cast<QDomNodeList*>(ptr)->~QDomNodeList();
}

