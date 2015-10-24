#include "qtextdecoder.h"
#include <QModelIndex>
#include <QTextCodec>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QTextDecoder>
#include "_cgo_export.h"

class MyQTextDecoder: public QTextDecoder {
public:
};

QtObjectPtr QTextDecoder_NewQTextDecoder(QtObjectPtr codec){
	return new QTextDecoder(static_cast<QTextCodec*>(codec));
}

QtObjectPtr QTextDecoder_NewQTextDecoder2(QtObjectPtr codec, int flags){
	return new QTextDecoder(static_cast<QTextCodec*>(codec), static_cast<QTextCodec::ConversionFlag>(flags));
}

void QTextDecoder_DestroyQTextDecoder(QtObjectPtr ptr){
	static_cast<QTextDecoder*>(ptr)->~QTextDecoder();
}

