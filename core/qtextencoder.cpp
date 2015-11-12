#include "qtextencoder.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextCodec>
#include <QByteArray>
#include <QChar>
#include <QTextEncoder>
#include "_cgo_export.h"

class MyQTextEncoder: public QTextEncoder {
public:
};

void* QTextEncoder_NewQTextEncoder(void* codec){
	return new QTextEncoder(static_cast<QTextCodec*>(codec));
}

void* QTextEncoder_NewQTextEncoder2(void* codec, int flags){
	return new QTextEncoder(static_cast<QTextCodec*>(codec), static_cast<QTextCodec::ConversionFlag>(flags));
}

void* QTextEncoder_FromUnicode2(void* ptr, void* uc, int len){
	return new QByteArray(static_cast<QTextEncoder*>(ptr)->fromUnicode(static_cast<QChar*>(uc), len));
}

void* QTextEncoder_FromUnicode(void* ptr, char* str){
	return new QByteArray(static_cast<QTextEncoder*>(ptr)->fromUnicode(QString(str)));
}

void QTextEncoder_DestroyQTextEncoder(void* ptr){
	static_cast<QTextEncoder*>(ptr)->~QTextEncoder();
}

