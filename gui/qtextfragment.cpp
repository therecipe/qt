#include "qtextfragment.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextFragment>
#include "_cgo_export.h"

class MyQTextFragment: public QTextFragment {
public:
};

void* QTextFragment_NewQTextFragment(){
	return new QTextFragment();
}

void* QTextFragment_NewQTextFragment3(void* other){
	return new QTextFragment(*static_cast<QTextFragment*>(other));
}

int QTextFragment_CharFormatIndex(void* ptr){
	return static_cast<QTextFragment*>(ptr)->charFormatIndex();
}

int QTextFragment_Contains(void* ptr, int position){
	return static_cast<QTextFragment*>(ptr)->contains(position);
}

int QTextFragment_IsValid(void* ptr){
	return static_cast<QTextFragment*>(ptr)->isValid();
}

int QTextFragment_Length(void* ptr){
	return static_cast<QTextFragment*>(ptr)->length();
}

int QTextFragment_Position(void* ptr){
	return static_cast<QTextFragment*>(ptr)->position();
}

char* QTextFragment_Text(void* ptr){
	return static_cast<QTextFragment*>(ptr)->text().toUtf8().data();
}

