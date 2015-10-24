#include "qtextstream.h"
#include <QLocale>
#include <QIODevice>
#include <QTextCodec>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QChar>
#include <QTextStream>
#include "_cgo_export.h"

class MyQTextStream: public QTextStream {
public:
};

int QTextStream_AtEnd(QtObjectPtr ptr){
	return static_cast<QTextStream*>(ptr)->atEnd();
}

int QTextStream_AutoDetectUnicode(QtObjectPtr ptr){
	return static_cast<QTextStream*>(ptr)->autoDetectUnicode();
}

QtObjectPtr QTextStream_Codec(QtObjectPtr ptr){
	return static_cast<QTextStream*>(ptr)->codec();
}

QtObjectPtr QTextStream_Device(QtObjectPtr ptr){
	return static_cast<QTextStream*>(ptr)->device();
}

int QTextStream_FieldAlignment(QtObjectPtr ptr){
	return static_cast<QTextStream*>(ptr)->fieldAlignment();
}

int QTextStream_FieldWidth(QtObjectPtr ptr){
	return static_cast<QTextStream*>(ptr)->fieldWidth();
}

void QTextStream_Flush(QtObjectPtr ptr){
	static_cast<QTextStream*>(ptr)->flush();
}

int QTextStream_GenerateByteOrderMark(QtObjectPtr ptr){
	return static_cast<QTextStream*>(ptr)->generateByteOrderMark();
}

int QTextStream_IntegerBase(QtObjectPtr ptr){
	return static_cast<QTextStream*>(ptr)->integerBase();
}

int QTextStream_NumberFlags(QtObjectPtr ptr){
	return static_cast<QTextStream*>(ptr)->numberFlags();
}

char* QTextStream_ReadAll(QtObjectPtr ptr){
	return static_cast<QTextStream*>(ptr)->readAll().toUtf8().data();
}

int QTextStream_RealNumberNotation(QtObjectPtr ptr){
	return static_cast<QTextStream*>(ptr)->realNumberNotation();
}

int QTextStream_RealNumberPrecision(QtObjectPtr ptr){
	return static_cast<QTextStream*>(ptr)->realNumberPrecision();
}

void QTextStream_Reset(QtObjectPtr ptr){
	static_cast<QTextStream*>(ptr)->reset();
}

void QTextStream_ResetStatus(QtObjectPtr ptr){
	static_cast<QTextStream*>(ptr)->resetStatus();
}

void QTextStream_SetAutoDetectUnicode(QtObjectPtr ptr, int enabled){
	static_cast<QTextStream*>(ptr)->setAutoDetectUnicode(enabled != 0);
}

void QTextStream_SetCodec(QtObjectPtr ptr, QtObjectPtr codec){
	static_cast<QTextStream*>(ptr)->setCodec(static_cast<QTextCodec*>(codec));
}

void QTextStream_SetCodec2(QtObjectPtr ptr, char* codecName){
	static_cast<QTextStream*>(ptr)->setCodec(const_cast<const char*>(codecName));
}

void QTextStream_SetDevice(QtObjectPtr ptr, QtObjectPtr device){
	static_cast<QTextStream*>(ptr)->setDevice(static_cast<QIODevice*>(device));
}

void QTextStream_SetFieldAlignment(QtObjectPtr ptr, int mode){
	static_cast<QTextStream*>(ptr)->setFieldAlignment(static_cast<QTextStream::FieldAlignment>(mode));
}

void QTextStream_SetFieldWidth(QtObjectPtr ptr, int width){
	static_cast<QTextStream*>(ptr)->setFieldWidth(width);
}

void QTextStream_SetGenerateByteOrderMark(QtObjectPtr ptr, int generate){
	static_cast<QTextStream*>(ptr)->setGenerateByteOrderMark(generate != 0);
}

void QTextStream_SetIntegerBase(QtObjectPtr ptr, int base){
	static_cast<QTextStream*>(ptr)->setIntegerBase(base);
}

void QTextStream_SetLocale(QtObjectPtr ptr, QtObjectPtr locale){
	static_cast<QTextStream*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QTextStream_SetNumberFlags(QtObjectPtr ptr, int flags){
	static_cast<QTextStream*>(ptr)->setNumberFlags(static_cast<QTextStream::NumberFlag>(flags));
}

void QTextStream_SetPadChar(QtObjectPtr ptr, QtObjectPtr ch){
	static_cast<QTextStream*>(ptr)->setPadChar(*static_cast<QChar*>(ch));
}

void QTextStream_SetRealNumberNotation(QtObjectPtr ptr, int notation){
	static_cast<QTextStream*>(ptr)->setRealNumberNotation(static_cast<QTextStream::RealNumberNotation>(notation));
}

void QTextStream_SetRealNumberPrecision(QtObjectPtr ptr, int precision){
	static_cast<QTextStream*>(ptr)->setRealNumberPrecision(precision);
}

void QTextStream_SetStatus(QtObjectPtr ptr, int status){
	static_cast<QTextStream*>(ptr)->setStatus(static_cast<QTextStream::Status>(status));
}

void QTextStream_SetString(QtObjectPtr ptr, char* stri, int openMode){
	static_cast<QTextStream*>(ptr)->setString(new QString(stri), static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QTextStream_SkipWhiteSpace(QtObjectPtr ptr){
	static_cast<QTextStream*>(ptr)->skipWhiteSpace();
}

int QTextStream_Status(QtObjectPtr ptr){
	return static_cast<QTextStream*>(ptr)->status();
}

char* QTextStream_String(QtObjectPtr ptr){
	return static_cast<QTextStream*>(ptr)->string()->toUtf8().data();
}

void QTextStream_DestroyQTextStream(QtObjectPtr ptr){
	static_cast<QTextStream*>(ptr)->~QTextStream();
}

