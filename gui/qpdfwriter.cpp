#include "qpdfwriter.h"
#include <QString>
#include <QVariant>
#include <QPageSize>
#include <QIODevice>
#include <QPageLayout>
#include <QMargins>
#include <QUrl>
#include <QModelIndex>
#include <QMarginsF>
#include <QPdfWriter>
#include "_cgo_export.h"

class MyQPdfWriter: public QPdfWriter {
public:
};

void* QPdfWriter_NewQPdfWriter2(void* device){
	return new QPdfWriter(static_cast<QIODevice*>(device));
}

void* QPdfWriter_NewQPdfWriter(char* filename){
	return new QPdfWriter(QString(filename));
}

char* QPdfWriter_Creator(void* ptr){
	return static_cast<QPdfWriter*>(ptr)->creator().toUtf8().data();
}

int QPdfWriter_NewPage(void* ptr){
	return static_cast<QPdfWriter*>(ptr)->newPage();
}

int QPdfWriter_Resolution(void* ptr){
	return static_cast<QPdfWriter*>(ptr)->resolution();
}

void QPdfWriter_SetCreator(void* ptr, char* creator){
	static_cast<QPdfWriter*>(ptr)->setCreator(QString(creator));
}

int QPdfWriter_SetPageLayout(void* ptr, void* newPageLayout){
	return static_cast<QPdfWriter*>(ptr)->setPageLayout(*static_cast<QPageLayout*>(newPageLayout));
}

int QPdfWriter_SetPageMargins(void* ptr, void* margins){
	return static_cast<QPdfWriter*>(ptr)->setPageMargins(*static_cast<QMarginsF*>(margins));
}

int QPdfWriter_SetPageMargins2(void* ptr, void* margins, int units){
	return static_cast<QPdfWriter*>(ptr)->setPageMargins(*static_cast<QMarginsF*>(margins), static_cast<QPageLayout::Unit>(units));
}

int QPdfWriter_SetPageOrientation(void* ptr, int orientation){
	return static_cast<QPdfWriter*>(ptr)->setPageOrientation(static_cast<QPageLayout::Orientation>(orientation));
}

int QPdfWriter_SetPageSize(void* ptr, void* pageSize){
	return static_cast<QPdfWriter*>(ptr)->setPageSize(*static_cast<QPageSize*>(pageSize));
}

void QPdfWriter_SetResolution(void* ptr, int resolution){
	static_cast<QPdfWriter*>(ptr)->setResolution(resolution);
}

void QPdfWriter_SetTitle(void* ptr, char* title){
	static_cast<QPdfWriter*>(ptr)->setTitle(QString(title));
}

char* QPdfWriter_Title(void* ptr){
	return static_cast<QPdfWriter*>(ptr)->title().toUtf8().data();
}

void QPdfWriter_DestroyQPdfWriter(void* ptr){
	static_cast<QPdfWriter*>(ptr)->~QPdfWriter();
}

