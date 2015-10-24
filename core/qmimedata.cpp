#include "qmimedata.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QString>
#include <QMimeData>
#include "_cgo_export.h"

class MyQMimeData: public QMimeData {
public:
};

QtObjectPtr QMimeData_NewQMimeData(){
	return new QMimeData();
}

void QMimeData_Clear(QtObjectPtr ptr){
	static_cast<QMimeData*>(ptr)->clear();
}

char* QMimeData_ColorData(QtObjectPtr ptr){
	return static_cast<QMimeData*>(ptr)->colorData().toString().toUtf8().data();
}

char* QMimeData_Formats(QtObjectPtr ptr){
	return static_cast<QMimeData*>(ptr)->formats().join("|").toUtf8().data();
}

int QMimeData_HasColor(QtObjectPtr ptr){
	return static_cast<QMimeData*>(ptr)->hasColor();
}

int QMimeData_HasFormat(QtObjectPtr ptr, char* mimeType){
	return static_cast<QMimeData*>(ptr)->hasFormat(QString(mimeType));
}

int QMimeData_HasHtml(QtObjectPtr ptr){
	return static_cast<QMimeData*>(ptr)->hasHtml();
}

int QMimeData_HasImage(QtObjectPtr ptr){
	return static_cast<QMimeData*>(ptr)->hasImage();
}

int QMimeData_HasText(QtObjectPtr ptr){
	return static_cast<QMimeData*>(ptr)->hasText();
}

int QMimeData_HasUrls(QtObjectPtr ptr){
	return static_cast<QMimeData*>(ptr)->hasUrls();
}

char* QMimeData_Html(QtObjectPtr ptr){
	return static_cast<QMimeData*>(ptr)->html().toUtf8().data();
}

char* QMimeData_ImageData(QtObjectPtr ptr){
	return static_cast<QMimeData*>(ptr)->imageData().toString().toUtf8().data();
}

void QMimeData_RemoveFormat(QtObjectPtr ptr, char* mimeType){
	static_cast<QMimeData*>(ptr)->removeFormat(QString(mimeType));
}

void QMimeData_SetColorData(QtObjectPtr ptr, char* color){
	static_cast<QMimeData*>(ptr)->setColorData(QVariant(color));
}

void QMimeData_SetData(QtObjectPtr ptr, char* mimeType, QtObjectPtr data){
	static_cast<QMimeData*>(ptr)->setData(QString(mimeType), *static_cast<QByteArray*>(data));
}

void QMimeData_SetHtml(QtObjectPtr ptr, char* html){
	static_cast<QMimeData*>(ptr)->setHtml(QString(html));
}

void QMimeData_SetImageData(QtObjectPtr ptr, char* image){
	static_cast<QMimeData*>(ptr)->setImageData(QVariant(image));
}

void QMimeData_SetText(QtObjectPtr ptr, char* text){
	static_cast<QMimeData*>(ptr)->setText(QString(text));
}

char* QMimeData_Text(QtObjectPtr ptr){
	return static_cast<QMimeData*>(ptr)->text().toUtf8().data();
}

void QMimeData_DestroyQMimeData(QtObjectPtr ptr){
	static_cast<QMimeData*>(ptr)->~QMimeData();
}

