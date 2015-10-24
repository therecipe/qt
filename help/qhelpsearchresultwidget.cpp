#include "qhelpsearchresultwidget.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QPoint>
#include <QString>
#include <QHelpSearchResultWidget>
#include "_cgo_export.h"

class MyQHelpSearchResultWidget: public QHelpSearchResultWidget {
public:
void Signal_RequestShowLink(const QUrl & link){callbackQHelpSearchResultWidgetRequestShowLink(this->objectName().toUtf8().data(), link.toString().toUtf8().data());};
};

char* QHelpSearchResultWidget_LinkAt(QtObjectPtr ptr, QtObjectPtr point){
	return static_cast<QHelpSearchResultWidget*>(ptr)->linkAt(*static_cast<QPoint*>(point)).toString().toUtf8().data();
}

void QHelpSearchResultWidget_ConnectRequestShowLink(QtObjectPtr ptr){
	QObject::connect(static_cast<QHelpSearchResultWidget*>(ptr), static_cast<void (QHelpSearchResultWidget::*)(const QUrl &)>(&QHelpSearchResultWidget::requestShowLink), static_cast<MyQHelpSearchResultWidget*>(ptr), static_cast<void (MyQHelpSearchResultWidget::*)(const QUrl &)>(&MyQHelpSearchResultWidget::Signal_RequestShowLink));;
}

void QHelpSearchResultWidget_DisconnectRequestShowLink(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QHelpSearchResultWidget*>(ptr), static_cast<void (QHelpSearchResultWidget::*)(const QUrl &)>(&QHelpSearchResultWidget::requestShowLink), static_cast<MyQHelpSearchResultWidget*>(ptr), static_cast<void (MyQHelpSearchResultWidget::*)(const QUrl &)>(&MyQHelpSearchResultWidget::Signal_RequestShowLink));;
}

void QHelpSearchResultWidget_DestroyQHelpSearchResultWidget(QtObjectPtr ptr){
	static_cast<QHelpSearchResultWidget*>(ptr)->~QHelpSearchResultWidget();
}

