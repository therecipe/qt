#include "qsourcelocation.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSourceLocation>
#include "_cgo_export.h"

class MyQSourceLocation: public QSourceLocation {
public:
};

QtObjectPtr QSourceLocation_NewQSourceLocation(){
	return new QSourceLocation();
}

QtObjectPtr QSourceLocation_NewQSourceLocation2(QtObjectPtr other){
	return new QSourceLocation(*static_cast<QSourceLocation*>(other));
}

QtObjectPtr QSourceLocation_NewQSourceLocation3(char* u, int l, int c){
	return new QSourceLocation(QUrl(QString(u)), l, c);
}

int QSourceLocation_IsNull(QtObjectPtr ptr){
	return static_cast<QSourceLocation*>(ptr)->isNull();
}

void QSourceLocation_SetUri(QtObjectPtr ptr, char* newUri){
	static_cast<QSourceLocation*>(ptr)->setUri(QUrl(QString(newUri)));
}

char* QSourceLocation_Uri(QtObjectPtr ptr){
	return static_cast<QSourceLocation*>(ptr)->uri().toString().toUtf8().data();
}

void QSourceLocation_DestroyQSourceLocation(QtObjectPtr ptr){
	static_cast<QSourceLocation*>(ptr)->~QSourceLocation();
}

