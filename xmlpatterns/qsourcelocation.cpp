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

void* QSourceLocation_NewQSourceLocation(){
	return new QSourceLocation();
}

void* QSourceLocation_NewQSourceLocation2(void* other){
	return new QSourceLocation(*static_cast<QSourceLocation*>(other));
}

void* QSourceLocation_NewQSourceLocation3(void* u, int l, int c){
	return new QSourceLocation(*static_cast<QUrl*>(u), l, c);
}

int QSourceLocation_IsNull(void* ptr){
	return static_cast<QSourceLocation*>(ptr)->isNull();
}

void QSourceLocation_SetUri(void* ptr, void* newUri){
	static_cast<QSourceLocation*>(ptr)->setUri(*static_cast<QUrl*>(newUri));
}

void QSourceLocation_DestroyQSourceLocation(void* ptr){
	static_cast<QSourceLocation*>(ptr)->~QSourceLocation();
}

