#include "qsvgwidget.h"
#include <QWidget>
#include <QByteArray>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSvgWidget>
#include "_cgo_export.h"

class MyQSvgWidget: public QSvgWidget {
public:
};

QtObjectPtr QSvgWidget_NewQSvgWidget(QtObjectPtr parent){
	return new QSvgWidget(static_cast<QWidget*>(parent));
}

QtObjectPtr QSvgWidget_NewQSvgWidget2(char* file, QtObjectPtr parent){
	return new QSvgWidget(QString(file), static_cast<QWidget*>(parent));
}

void QSvgWidget_Load2(QtObjectPtr ptr, QtObjectPtr contents){
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "load", Q_ARG(QByteArray, *static_cast<QByteArray*>(contents)));
}

void QSvgWidget_Load(QtObjectPtr ptr, char* file){
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "load", Q_ARG(QString, QString(file)));
}

QtObjectPtr QSvgWidget_Renderer(QtObjectPtr ptr){
	return static_cast<QSvgWidget*>(ptr)->renderer();
}

void QSvgWidget_DestroyQSvgWidget(QtObjectPtr ptr){
	static_cast<QSvgWidget*>(ptr)->~QSvgWidget();
}

