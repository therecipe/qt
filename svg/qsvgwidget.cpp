#include "qsvgwidget.h"
#include <QModelIndex>
#include <QMetaObject>
#include <QWidget>
#include <QByteArray>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSvgWidget>
#include "_cgo_export.h"

class MyQSvgWidget: public QSvgWidget {
public:
};

void* QSvgWidget_NewQSvgWidget(void* parent){
	return new QSvgWidget(static_cast<QWidget*>(parent));
}

void* QSvgWidget_NewQSvgWidget2(char* file, void* parent){
	return new QSvgWidget(QString(file), static_cast<QWidget*>(parent));
}

void QSvgWidget_Load2(void* ptr, void* contents){
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "load", Q_ARG(QByteArray, *static_cast<QByteArray*>(contents)));
}

void QSvgWidget_Load(void* ptr, char* file){
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "load", Q_ARG(QString, QString(file)));
}

void* QSvgWidget_Renderer(void* ptr){
	return static_cast<QSvgWidget*>(ptr)->renderer();
}

void QSvgWidget_DestroyQSvgWidget(void* ptr){
	static_cast<QSvgWidget*>(ptr)->~QSvgWidget();
}

