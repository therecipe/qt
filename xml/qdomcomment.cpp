#include "qdomcomment.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QDomComment>
#include "_cgo_export.h"

class MyQDomComment: public QDomComment {
public:
};

void* QDomComment_NewQDomComment(){
	return new QDomComment();
}

void* QDomComment_NewQDomComment2(void* x){
	return new QDomComment(*static_cast<QDomComment*>(x));
}

int QDomComment_NodeType(void* ptr){
	return static_cast<QDomComment*>(ptr)->nodeType();
}

