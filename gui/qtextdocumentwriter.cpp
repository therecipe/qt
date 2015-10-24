#include "qtextdocumentwriter.h"
#include <QTextDocument>
#include <QTextCodec>
#include <QUrl>
#include <QByteArray>
#include <QTextDocumentFragment>
#include <QIODevice>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QTextDocumentWriter>
#include "_cgo_export.h"

class MyQTextDocumentWriter: public QTextDocumentWriter {
public:
};

QtObjectPtr QTextDocumentWriter_NewQTextDocumentWriter(){
	return new QTextDocumentWriter();
}

QtObjectPtr QTextDocumentWriter_NewQTextDocumentWriter2(QtObjectPtr device, QtObjectPtr format){
	return new QTextDocumentWriter(static_cast<QIODevice*>(device), *static_cast<QByteArray*>(format));
}

QtObjectPtr QTextDocumentWriter_NewQTextDocumentWriter3(char* fileName, QtObjectPtr format){
	return new QTextDocumentWriter(QString(fileName), *static_cast<QByteArray*>(format));
}

QtObjectPtr QTextDocumentWriter_Codec(QtObjectPtr ptr){
	return static_cast<QTextDocumentWriter*>(ptr)->codec();
}

QtObjectPtr QTextDocumentWriter_Device(QtObjectPtr ptr){
	return static_cast<QTextDocumentWriter*>(ptr)->device();
}

char* QTextDocumentWriter_FileName(QtObjectPtr ptr){
	return static_cast<QTextDocumentWriter*>(ptr)->fileName().toUtf8().data();
}

void QTextDocumentWriter_SetCodec(QtObjectPtr ptr, QtObjectPtr codec){
	static_cast<QTextDocumentWriter*>(ptr)->setCodec(static_cast<QTextCodec*>(codec));
}

void QTextDocumentWriter_SetDevice(QtObjectPtr ptr, QtObjectPtr device){
	static_cast<QTextDocumentWriter*>(ptr)->setDevice(static_cast<QIODevice*>(device));
}

void QTextDocumentWriter_SetFileName(QtObjectPtr ptr, char* fileName){
	static_cast<QTextDocumentWriter*>(ptr)->setFileName(QString(fileName));
}

void QTextDocumentWriter_SetFormat(QtObjectPtr ptr, QtObjectPtr format){
	static_cast<QTextDocumentWriter*>(ptr)->setFormat(*static_cast<QByteArray*>(format));
}

int QTextDocumentWriter_Write(QtObjectPtr ptr, QtObjectPtr document){
	return static_cast<QTextDocumentWriter*>(ptr)->write(static_cast<QTextDocument*>(document));
}

int QTextDocumentWriter_Write2(QtObjectPtr ptr, QtObjectPtr fragment){
	return static_cast<QTextDocumentWriter*>(ptr)->write(*static_cast<QTextDocumentFragment*>(fragment));
}

void QTextDocumentWriter_DestroyQTextDocumentWriter(QtObjectPtr ptr){
	static_cast<QTextDocumentWriter*>(ptr)->~QTextDocumentWriter();
}

