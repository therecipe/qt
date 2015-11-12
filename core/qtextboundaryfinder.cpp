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

void* QTextBoundaryFinder_NewQTextBoundaryFinder(){
	return new QTextBoundaryFinder();
}

void* QTextBoundaryFinder_NewQTextBoundaryFinder3(int ty, char* stri){
	return new QTextBoundaryFinder(static_cast<QTextBoundaryFinder::BoundaryType>(ty), QString(stri));
}

void* QTextBoundaryFinder_NewQTextBoundaryFinder2(void* other){
	return new QTextBoundaryFinder(*static_cast<QTextBoundaryFinder*>(other));
}

int QTextBoundaryFinder_BoundaryReasons(void* ptr){
	return static_cast<QTextBoundaryFinder*>(ptr)->boundaryReasons();
}

int QTextBoundaryFinder_IsAtBoundary(void* ptr){
	return static_cast<QTextBoundaryFinder*>(ptr)->isAtBoundary();
}

int QTextBoundaryFinder_IsValid(void* ptr){
	return static_cast<QTextBoundaryFinder*>(ptr)->isValid();
}

int QTextBoundaryFinder_Position(void* ptr){
	return static_cast<QTextBoundaryFinder*>(ptr)->position();
}

void QTextBoundaryFinder_SetPosition(void* ptr, int position){
	static_cast<QTextBoundaryFinder*>(ptr)->setPosition(position);
}

char* QTextBoundaryFinder_String(void* ptr){
	return static_cast<QTextBoundaryFinder*>(ptr)->string().toUtf8().data();
}

void QTextBoundaryFinder_ToEnd(void* ptr){
	static_cast<QTextBoundaryFinder*>(ptr)->toEnd();
}

int QTextBoundaryFinder_ToNextBoundary(void* ptr){
	return static_cast<QTextBoundaryFinder*>(ptr)->toNextBoundary();
}

int QTextBoundaryFinder_ToPreviousBoundary(void* ptr){
	return static_cast<QTextBoundaryFinder*>(ptr)->toPreviousBoundary();
}

void QTextBoundaryFinder_ToStart(void* ptr){
	static_cast<QTextBoundaryFinder*>(ptr)->toStart();
}

int QTextBoundaryFinder_Type(void* ptr){
	return static_cast<QTextBoundaryFinder*>(ptr)->type();
}

void QTextBoundaryFinder_DestroyQTextBoundaryFinder(void* ptr){
	static_cast<QTextBoundaryFinder*>(ptr)->~QTextBoundaryFinder();
}

