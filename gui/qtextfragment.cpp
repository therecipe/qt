#include "qtextfragment.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QTextFragment>
#include "_cgo_export.h"

class MyQTextFragment: public QTextFragment {
public:
};

QtObjectPtr QTextFragment_NewQTextFragment(){
	return new QTextFragment();
}

QtObjectPtr QTextFragment_NewQTextFragment3(QtObjectPtr other){
	return new QTextFragment(*static_cast<QTextFragment*>(other));
}

int QTextFragment_CharFormatIndex(QtObjectPtr ptr){
	return static_cast<QTextFragment*>(ptr)->charFormatIndex();
}

int QTextFragment_Contains(QtObjectPtr ptr, int position){
	return static_cast<QTextFragment*>(ptr)->contains(position);
}

int QTextFragment_IsValid(QtObjectPtr ptr){
	return static_cast<QTextFragment*>(ptr)->isValid();
}

int QTextFragment_Length(QtObjectPtr ptr){
	return static_cast<QTextFragment*>(ptr)->length();
}

int QTextFragment_Position(QtObjectPtr ptr){
	return static_cast<QTextFragment*>(ptr)->position();
}

char* QTextFragment_Text(QtObjectPtr ptr){
	return static_cast<QTextFragment*>(ptr)->text().toUtf8().data();
}

