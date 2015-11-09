#include "qtextdecoder.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextCodec>
#include <QTextDecoder>
#include "_cgo_export.h"

class MyQTextDecoder: public QTextDecoder {
public:
};

void* QTextDecoder_NewQTextDecoder(void* codec){
	return new QTextDecoder(static_cast<QTextCodec*>(codec));
}

void* QTextDecoder_NewQTextDecoder2(void* codec, int flags){
	return new QTextDecoder(static_cast<QTextCodec*>(codec), static_cast<QTextCodec::ConversionFlag>(flags));
}

void QTextDecoder_DestroyQTextDecoder(void* ptr){
	static_cast<QTextDecoder*>(ptr)->~QTextDecoder();
}

