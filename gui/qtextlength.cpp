#include "qtextlength.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextLength>
#include "_cgo_export.h"

class MyQTextLength: public QTextLength {
public:
};

QtObjectPtr QTextLength_NewQTextLength(){
	return new QTextLength();
}

int QTextLength_Type(QtObjectPtr ptr){
	return static_cast<QTextLength*>(ptr)->type();
}

