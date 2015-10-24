#include "qpdfwriter.h"
#include <QIODevice>
#include <QString>
#include <QMargins>
#include <QPageLayout>
#include <QMarginsF>
#include <QPageSize>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPdfWriter>
#include "_cgo_export.h"

class MyQPdfWriter: public QPdfWriter {
public:
};

QtObjectPtr QPdfWriter_NewQPdfWriter2(QtObjectPtr device){
	return new QPdfWriter(static_cast<QIODevice*>(device));
}

QtObjectPtr QPdfWriter_NewQPdfWriter(char* filename){
	return new QPdfWriter(QString(filename));
}

char* QPdfWriter_Creator(QtObjectPtr ptr){
	return static_cast<QPdfWriter*>(ptr)->creator().toUtf8().data();
}

int QPdfWriter_NewPage(QtObjectPtr ptr){
	return static_cast<QPdfWriter*>(ptr)->newPage();
}

int QPdfWriter_Resolution(QtObjectPtr ptr){
	return static_cast<QPdfWriter*>(ptr)->resolution();
}

void QPdfWriter_SetCreator(QtObjectPtr ptr, char* creator){
	static_cast<QPdfWriter*>(ptr)->setCreator(QString(creator));
}

int QPdfWriter_SetPageLayout(QtObjectPtr ptr, QtObjectPtr newPageLayout){
	return static_cast<QPdfWriter*>(ptr)->setPageLayout(*static_cast<QPageLayout*>(newPageLayout));
}

int QPdfWriter_SetPageMargins(QtObjectPtr ptr, QtObjectPtr margins){
	return static_cast<QPdfWriter*>(ptr)->setPageMargins(*static_cast<QMarginsF*>(margins));
}

int QPdfWriter_SetPageMargins2(QtObjectPtr ptr, QtObjectPtr margins, int units){
	return static_cast<QPdfWriter*>(ptr)->setPageMargins(*static_cast<QMarginsF*>(margins), static_cast<QPageLayout::Unit>(units));
}

int QPdfWriter_SetPageOrientation(QtObjectPtr ptr, int orientation){
	return static_cast<QPdfWriter*>(ptr)->setPageOrientation(static_cast<QPageLayout::Orientation>(orientation));
}

int QPdfWriter_SetPageSize(QtObjectPtr ptr, QtObjectPtr pageSize){
	return static_cast<QPdfWriter*>(ptr)->setPageSize(*static_cast<QPageSize*>(pageSize));
}

void QPdfWriter_SetResolution(QtObjectPtr ptr, int resolution){
	static_cast<QPdfWriter*>(ptr)->setResolution(resolution);
}

void QPdfWriter_SetTitle(QtObjectPtr ptr, char* title){
	static_cast<QPdfWriter*>(ptr)->setTitle(QString(title));
}

char* QPdfWriter_Title(QtObjectPtr ptr){
	return static_cast<QPdfWriter*>(ptr)->title().toUtf8().data();
}

void QPdfWriter_DestroyQPdfWriter(QtObjectPtr ptr){
	static_cast<QPdfWriter*>(ptr)->~QPdfWriter();
}

