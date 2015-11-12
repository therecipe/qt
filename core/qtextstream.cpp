#include "qtextstream.h"
#include <QChar>
#include <QIODevice>
#include <QTextCodec>
#include <QLocale>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextStream>
#include "_cgo_export.h"

class MyQTextStream: public QTextStream {
public:
};

int QTextStream_AtEnd(void* ptr){
	return static_cast<QTextStream*>(ptr)->atEnd();
}

int QTextStream_AutoDetectUnicode(void* ptr){
	return static_cast<QTextStream*>(ptr)->autoDetectUnicode();
}

void* QTextStream_Codec(void* ptr){
	return static_cast<QTextStream*>(ptr)->codec();
}

void* QTextStream_Device(void* ptr){
	return static_cast<QTextStream*>(ptr)->device();
}

int QTextStream_FieldAlignment(void* ptr){
	return static_cast<QTextStream*>(ptr)->fieldAlignment();
}

int QTextStream_FieldWidth(void* ptr){
	return static_cast<QTextStream*>(ptr)->fieldWidth();
}

void QTextStream_Flush(void* ptr){
	static_cast<QTextStream*>(ptr)->flush();
}

int QTextStream_GenerateByteOrderMark(void* ptr){
	return static_cast<QTextStream*>(ptr)->generateByteOrderMark();
}

int QTextStream_IntegerBase(void* ptr){
	return static_cast<QTextStream*>(ptr)->integerBase();
}

int QTextStream_NumberFlags(void* ptr){
	return static_cast<QTextStream*>(ptr)->numberFlags();
}

char* QTextStream_ReadAll(void* ptr){
	return static_cast<QTextStream*>(ptr)->readAll().toUtf8().data();
}

int QTextStream_RealNumberNotation(void* ptr){
	return static_cast<QTextStream*>(ptr)->realNumberNotation();
}

int QTextStream_RealNumberPrecision(void* ptr){
	return static_cast<QTextStream*>(ptr)->realNumberPrecision();
}

void QTextStream_Reset(void* ptr){
	static_cast<QTextStream*>(ptr)->reset();
}

void QTextStream_ResetStatus(void* ptr){
	static_cast<QTextStream*>(ptr)->resetStatus();
}

void QTextStream_SetAutoDetectUnicode(void* ptr, int enabled){
	static_cast<QTextStream*>(ptr)->setAutoDetectUnicode(enabled != 0);
}

void QTextStream_SetCodec(void* ptr, void* codec){
	static_cast<QTextStream*>(ptr)->setCodec(static_cast<QTextCodec*>(codec));
}

void QTextStream_SetCodec2(void* ptr, char* codecName){
	static_cast<QTextStream*>(ptr)->setCodec(const_cast<const char*>(codecName));
}

void QTextStream_SetDevice(void* ptr, void* device){
	static_cast<QTextStream*>(ptr)->setDevice(static_cast<QIODevice*>(device));
}

void QTextStream_SetFieldAlignment(void* ptr, int mode){
	static_cast<QTextStream*>(ptr)->setFieldAlignment(static_cast<QTextStream::FieldAlignment>(mode));
}

void QTextStream_SetFieldWidth(void* ptr, int width){
	static_cast<QTextStream*>(ptr)->setFieldWidth(width);
}

void QTextStream_SetGenerateByteOrderMark(void* ptr, int generate){
	static_cast<QTextStream*>(ptr)->setGenerateByteOrderMark(generate != 0);
}

void QTextStream_SetIntegerBase(void* ptr, int base){
	static_cast<QTextStream*>(ptr)->setIntegerBase(base);
}

void QTextStream_SetLocale(void* ptr, void* locale){
	static_cast<QTextStream*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QTextStream_SetNumberFlags(void* ptr, int flags){
	static_cast<QTextStream*>(ptr)->setNumberFlags(static_cast<QTextStream::NumberFlag>(flags));
}

void QTextStream_SetPadChar(void* ptr, void* ch){
	static_cast<QTextStream*>(ptr)->setPadChar(*static_cast<QChar*>(ch));
}

void QTextStream_SetRealNumberNotation(void* ptr, int notation){
	static_cast<QTextStream*>(ptr)->setRealNumberNotation(static_cast<QTextStream::RealNumberNotation>(notation));
}

void QTextStream_SetRealNumberPrecision(void* ptr, int precision){
	static_cast<QTextStream*>(ptr)->setRealNumberPrecision(precision);
}

void QTextStream_SetStatus(void* ptr, int status){
	static_cast<QTextStream*>(ptr)->setStatus(static_cast<QTextStream::Status>(status));
}

void QTextStream_SetString(void* ptr, char* stri, int openMode){
	static_cast<QTextStream*>(ptr)->setString(new QString(stri), static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QTextStream_SkipWhiteSpace(void* ptr){
	static_cast<QTextStream*>(ptr)->skipWhiteSpace();
}

int QTextStream_Status(void* ptr){
	return static_cast<QTextStream*>(ptr)->status();
}

char* QTextStream_String(void* ptr){
	return static_cast<QTextStream*>(ptr)->string()->toUtf8().data();
}

void QTextStream_DestroyQTextStream(void* ptr){
	static_cast<QTextStream*>(ptr)->~QTextStream();
}

