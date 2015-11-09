#include "qmimedata.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QMimeData>
#include "_cgo_export.h"

class MyQMimeData: public QMimeData {
public:
};

void* QMimeData_NewQMimeData(){
	return new QMimeData();
}

void QMimeData_Clear(void* ptr){
	static_cast<QMimeData*>(ptr)->clear();
}

void* QMimeData_ColorData(void* ptr){
	return new QVariant(static_cast<QMimeData*>(ptr)->colorData());
}

void* QMimeData_Data(void* ptr, char* mimeType){
	return new QByteArray(static_cast<QMimeData*>(ptr)->data(QString(mimeType)));
}

char* QMimeData_Formats(void* ptr){
	return static_cast<QMimeData*>(ptr)->formats().join("|").toUtf8().data();
}

int QMimeData_HasColor(void* ptr){
	return static_cast<QMimeData*>(ptr)->hasColor();
}

int QMimeData_HasFormat(void* ptr, char* mimeType){
	return static_cast<QMimeData*>(ptr)->hasFormat(QString(mimeType));
}

int QMimeData_HasHtml(void* ptr){
	return static_cast<QMimeData*>(ptr)->hasHtml();
}

int QMimeData_HasImage(void* ptr){
	return static_cast<QMimeData*>(ptr)->hasImage();
}

int QMimeData_HasText(void* ptr){
	return static_cast<QMimeData*>(ptr)->hasText();
}

int QMimeData_HasUrls(void* ptr){
	return static_cast<QMimeData*>(ptr)->hasUrls();
}

char* QMimeData_Html(void* ptr){
	return static_cast<QMimeData*>(ptr)->html().toUtf8().data();
}

void* QMimeData_ImageData(void* ptr){
	return new QVariant(static_cast<QMimeData*>(ptr)->imageData());
}

void QMimeData_RemoveFormat(void* ptr, char* mimeType){
	static_cast<QMimeData*>(ptr)->removeFormat(QString(mimeType));
}

void QMimeData_SetColorData(void* ptr, void* color){
	static_cast<QMimeData*>(ptr)->setColorData(*static_cast<QVariant*>(color));
}

void QMimeData_SetData(void* ptr, char* mimeType, void* data){
	static_cast<QMimeData*>(ptr)->setData(QString(mimeType), *static_cast<QByteArray*>(data));
}

void QMimeData_SetHtml(void* ptr, char* html){
	static_cast<QMimeData*>(ptr)->setHtml(QString(html));
}

void QMimeData_SetImageData(void* ptr, void* image){
	static_cast<QMimeData*>(ptr)->setImageData(*static_cast<QVariant*>(image));
}

void QMimeData_SetText(void* ptr, char* text){
	static_cast<QMimeData*>(ptr)->setText(QString(text));
}

char* QMimeData_Text(void* ptr){
	return static_cast<QMimeData*>(ptr)->text().toUtf8().data();
}

void QMimeData_DestroyQMimeData(void* ptr){
	static_cast<QMimeData*>(ptr)->~QMimeData();
}

