#include "qtextcodec.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QChar>
#include <QByteArray>
#include <QTextCodec>
#include "_cgo_export.h"

class MyQTextCodec: public QTextCodec {
public:
};

int QTextCodec_CanEncode(QtObjectPtr ptr, QtObjectPtr ch){
	return static_cast<QTextCodec*>(ptr)->canEncode(*static_cast<QChar*>(ch));
}

int QTextCodec_CanEncode2(QtObjectPtr ptr, char* s){
	return static_cast<QTextCodec*>(ptr)->canEncode(QString(s));
}

QtObjectPtr QTextCodec_QTextCodec_CodecForHtml2(QtObjectPtr ba){
	return QTextCodec::codecForHtml(*static_cast<QByteArray*>(ba));
}

QtObjectPtr QTextCodec_QTextCodec_CodecForHtml(QtObjectPtr ba, QtObjectPtr defaultCodec){
	return QTextCodec::codecForHtml(*static_cast<QByteArray*>(ba), static_cast<QTextCodec*>(defaultCodec));
}

QtObjectPtr QTextCodec_QTextCodec_CodecForLocale(){
	return QTextCodec::codecForLocale();
}

QtObjectPtr QTextCodec_QTextCodec_CodecForMib(int mib){
	return QTextCodec::codecForMib(mib);
}

QtObjectPtr QTextCodec_QTextCodec_CodecForName(QtObjectPtr name){
	return QTextCodec::codecForName(*static_cast<QByteArray*>(name));
}

QtObjectPtr QTextCodec_QTextCodec_CodecForName2(char* name){
	return QTextCodec::codecForName(const_cast<const char*>(name));
}

QtObjectPtr QTextCodec_QTextCodec_CodecForUtfText2(QtObjectPtr ba){
	return QTextCodec::codecForUtfText(*static_cast<QByteArray*>(ba));
}

QtObjectPtr QTextCodec_QTextCodec_CodecForUtfText(QtObjectPtr ba, QtObjectPtr defaultCodec){
	return QTextCodec::codecForUtfText(*static_cast<QByteArray*>(ba), static_cast<QTextCodec*>(defaultCodec));
}

QtObjectPtr QTextCodec_MakeDecoder(QtObjectPtr ptr, int flags){
	return static_cast<QTextCodec*>(ptr)->makeDecoder(static_cast<QTextCodec::ConversionFlag>(flags));
}

QtObjectPtr QTextCodec_MakeEncoder(QtObjectPtr ptr, int flags){
	return static_cast<QTextCodec*>(ptr)->makeEncoder(static_cast<QTextCodec::ConversionFlag>(flags));
}

int QTextCodec_MibEnum(QtObjectPtr ptr){
	return static_cast<QTextCodec*>(ptr)->mibEnum();
}

void QTextCodec_QTextCodec_SetCodecForLocale(QtObjectPtr c){
	QTextCodec::setCodecForLocale(static_cast<QTextCodec*>(c));
}

