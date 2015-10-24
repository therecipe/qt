#include "qdomcomment.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDomComment>
#include "_cgo_export.h"

class MyQDomComment: public QDomComment {
public:
};

QtObjectPtr QDomComment_NewQDomComment(){
	return new QDomComment();
}

QtObjectPtr QDomComment_NewQDomComment2(QtObjectPtr x){
	return new QDomComment(*static_cast<QDomComment*>(x));
}

int QDomComment_NodeType(QtObjectPtr ptr){
	return static_cast<QDomComment*>(ptr)->nodeType();
}

