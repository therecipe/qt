#include "qtextdocumentwriter.h"
#include <QString>
#include <QModelIndex>
#include <QTextDocument>
#include <QByteArray>
#include <QTextCodec>
#include <QVariant>
#include <QUrl>
#include <QIODevice>
#include <QTextDocumentFragment>
#include <QTextDocumentWriter>
#include "_cgo_export.h"

class MyQTextDocumentWriter: public QTextDocumentWriter {
public:
};

void* QTextDocumentWriter_NewQTextDocumentWriter(){
	return new QTextDocumentWriter();
}

void* QTextDocumentWriter_NewQTextDocumentWriter2(void* device, void* format){
	return new QTextDocumentWriter(static_cast<QIODevice*>(device), *static_cast<QByteArray*>(format));
}

void* QTextDocumentWriter_NewQTextDocumentWriter3(char* fileName, void* format){
	return new QTextDocumentWriter(QString(fileName), *static_cast<QByteArray*>(format));
}

void* QTextDocumentWriter_Codec(void* ptr){
	return static_cast<QTextDocumentWriter*>(ptr)->codec();
}

void* QTextDocumentWriter_Device(void* ptr){
	return static_cast<QTextDocumentWriter*>(ptr)->device();
}

char* QTextDocumentWriter_FileName(void* ptr){
	return static_cast<QTextDocumentWriter*>(ptr)->fileName().toUtf8().data();
}

void* QTextDocumentWriter_Format(void* ptr){
	return new QByteArray(static_cast<QTextDocumentWriter*>(ptr)->format());
}

void QTextDocumentWriter_SetCodec(void* ptr, void* codec){
	static_cast<QTextDocumentWriter*>(ptr)->setCodec(static_cast<QTextCodec*>(codec));
}

void QTextDocumentWriter_SetDevice(void* ptr, void* device){
	static_cast<QTextDocumentWriter*>(ptr)->setDevice(static_cast<QIODevice*>(device));
}

void QTextDocumentWriter_SetFileName(void* ptr, char* fileName){
	static_cast<QTextDocumentWriter*>(ptr)->setFileName(QString(fileName));
}

void QTextDocumentWriter_SetFormat(void* ptr, void* format){
	static_cast<QTextDocumentWriter*>(ptr)->setFormat(*static_cast<QByteArray*>(format));
}

int QTextDocumentWriter_Write(void* ptr, void* document){
	return static_cast<QTextDocumentWriter*>(ptr)->write(static_cast<QTextDocument*>(document));
}

int QTextDocumentWriter_Write2(void* ptr, void* fragment){
	return static_cast<QTextDocumentWriter*>(ptr)->write(*static_cast<QTextDocumentFragment*>(fragment));
}

void QTextDocumentWriter_DestroyQTextDocumentWriter(void* ptr){
	static_cast<QTextDocumentWriter*>(ptr)->~QTextDocumentWriter();
}

