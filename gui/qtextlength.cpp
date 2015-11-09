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

void* QTextLength_NewQTextLength(){
	return new QTextLength();
}

void* QTextLength_NewQTextLength2(int ty, double value){
	return new QTextLength(static_cast<QTextLength::Type>(ty), static_cast<qreal>(value));
}

double QTextLength_RawValue(void* ptr){
	return static_cast<double>(static_cast<QTextLength*>(ptr)->rawValue());
}

int QTextLength_Type(void* ptr){
	return static_cast<QTextLength*>(ptr)->type();
}

double QTextLength_Value(void* ptr, double maximumLength){
	return static_cast<double>(static_cast<QTextLength*>(ptr)->value(static_cast<qreal>(maximumLength)));
}

