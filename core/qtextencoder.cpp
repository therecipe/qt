#include "qtextencoder.h"
#include <QUrl>
#include <QModelIndex>
#include <QTextCodec>
#include <QString>
#include <QVariant>
#include <QTextEncoder>
#include "_cgo_export.h"

class MyQTextEncoder: public QTextEncoder {
public:
};

QtObjectPtr QTextEncoder_NewQTextEncoder(QtObjectPtr codec){
	return new QTextEncoder(static_cast<QTextCodec*>(codec));
}

QtObjectPtr QTextEncoder_NewQTextEncoder2(QtObjectPtr codec, int flags){
	return new QTextEncoder(static_cast<QTextCodec*>(codec), static_cast<QTextCodec::ConversionFlag>(flags));
}

void QTextEncoder_DestroyQTextEncoder(QtObjectPtr ptr){
	static_cast<QTextEncoder*>(ptr)->~QTextEncoder();
}

