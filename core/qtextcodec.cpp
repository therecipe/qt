#include "qtextcodec.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QChar>
#include <QTextCodec>
#include "_cgo_export.h"

class MyQTextCodec: public QTextCodec {
public:
};

int QTextCodec_CanEncode(void* ptr, void* ch){
	return static_cast<QTextCodec*>(ptr)->canEncode(*static_cast<QChar*>(ch));
}

int QTextCodec_CanEncode2(void* ptr, char* s){
	return static_cast<QTextCodec*>(ptr)->canEncode(QString(s));
}

void* QTextCodec_QTextCodec_CodecForHtml2(void* ba){
	return QTextCodec::codecForHtml(*static_cast<QByteArray*>(ba));
}

void* QTextCodec_QTextCodec_CodecForHtml(void* ba, void* defaultCodec){
	return QTextCodec::codecForHtml(*static_cast<QByteArray*>(ba), static_cast<QTextCodec*>(defaultCodec));
}

void* QTextCodec_QTextCodec_CodecForLocale(){
	return QTextCodec::codecForLocale();
}

void* QTextCodec_QTextCodec_CodecForMib(int mib){
	return QTextCodec::codecForMib(mib);
}

void* QTextCodec_QTextCodec_CodecForName(void* name){
	return QTextCodec::codecForName(*static_cast<QByteArray*>(name));
}

void* QTextCodec_QTextCodec_CodecForName2(char* name){
	return QTextCodec::codecForName(const_cast<const char*>(name));
}

void* QTextCodec_QTextCodec_CodecForUtfText2(void* ba){
	return QTextCodec::codecForUtfText(*static_cast<QByteArray*>(ba));
}

void* QTextCodec_QTextCodec_CodecForUtfText(void* ba, void* defaultCodec){
	return QTextCodec::codecForUtfText(*static_cast<QByteArray*>(ba), static_cast<QTextCodec*>(defaultCodec));
}

void* QTextCodec_FromUnicode(void* ptr, char* str){
	return new QByteArray(static_cast<QTextCodec*>(ptr)->fromUnicode(QString(str)));
}

void* QTextCodec_MakeDecoder(void* ptr, int flags){
	return static_cast<QTextCodec*>(ptr)->makeDecoder(static_cast<QTextCodec::ConversionFlag>(flags));
}

void* QTextCodec_MakeEncoder(void* ptr, int flags){
	return static_cast<QTextCodec*>(ptr)->makeEncoder(static_cast<QTextCodec::ConversionFlag>(flags));
}

int QTextCodec_MibEnum(void* ptr){
	return static_cast<QTextCodec*>(ptr)->mibEnum();
}

void* QTextCodec_Name(void* ptr){
	return new QByteArray(static_cast<QTextCodec*>(ptr)->name());
}

void QTextCodec_QTextCodec_SetCodecForLocale(void* c){
	QTextCodec::setCodecForLocale(static_cast<QTextCodec*>(c));
}

