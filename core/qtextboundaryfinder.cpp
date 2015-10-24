#include "qtextboundaryfinder.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextBoundaryFinder>
#include "_cgo_export.h"

class MyQTextBoundaryFinder: public QTextBoundaryFinder {
public:
};

QtObjectPtr QTextBoundaryFinder_NewQTextBoundaryFinder(){
	return new QTextBoundaryFinder();
}

QtObjectPtr QTextBoundaryFinder_NewQTextBoundaryFinder3(int ty, char* stri){
	return new QTextBoundaryFinder(static_cast<QTextBoundaryFinder::BoundaryType>(ty), QString(stri));
}

QtObjectPtr QTextBoundaryFinder_NewQTextBoundaryFinder2(QtObjectPtr other){
	return new QTextBoundaryFinder(*static_cast<QTextBoundaryFinder*>(other));
}

int QTextBoundaryFinder_BoundaryReasons(QtObjectPtr ptr){
	return static_cast<QTextBoundaryFinder*>(ptr)->boundaryReasons();
}

int QTextBoundaryFinder_IsAtBoundary(QtObjectPtr ptr){
	return static_cast<QTextBoundaryFinder*>(ptr)->isAtBoundary();
}

int QTextBoundaryFinder_IsValid(QtObjectPtr ptr){
	return static_cast<QTextBoundaryFinder*>(ptr)->isValid();
}

int QTextBoundaryFinder_Position(QtObjectPtr ptr){
	return static_cast<QTextBoundaryFinder*>(ptr)->position();
}

void QTextBoundaryFinder_SetPosition(QtObjectPtr ptr, int position){
	static_cast<QTextBoundaryFinder*>(ptr)->setPosition(position);
}

char* QTextBoundaryFinder_String(QtObjectPtr ptr){
	return static_cast<QTextBoundaryFinder*>(ptr)->string().toUtf8().data();
}

void QTextBoundaryFinder_ToEnd(QtObjectPtr ptr){
	static_cast<QTextBoundaryFinder*>(ptr)->toEnd();
}

int QTextBoundaryFinder_ToNextBoundary(QtObjectPtr ptr){
	return static_cast<QTextBoundaryFinder*>(ptr)->toNextBoundary();
}

int QTextBoundaryFinder_ToPreviousBoundary(QtObjectPtr ptr){
	return static_cast<QTextBoundaryFinder*>(ptr)->toPreviousBoundary();
}

void QTextBoundaryFinder_ToStart(QtObjectPtr ptr){
	static_cast<QTextBoundaryFinder*>(ptr)->toStart();
}

int QTextBoundaryFinder_Type(QtObjectPtr ptr){
	return static_cast<QTextBoundaryFinder*>(ptr)->type();
}

void QTextBoundaryFinder_DestroyQTextBoundaryFinder(QtObjectPtr ptr){
	static_cast<QTextBoundaryFinder*>(ptr)->~QTextBoundaryFinder();
}

