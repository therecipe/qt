#include "qhelpindexwidget.h"
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QHelpIndexWidget>
#include "_cgo_export.h"

class MyQHelpIndexWidget: public QHelpIndexWidget {
public:
void Signal_LinkActivated(const QUrl & link, const QString & keyword){callbackQHelpIndexWidgetLinkActivated(this->objectName().toUtf8().data(), link.toString().toUtf8().data(), keyword.toUtf8().data());};
};

void QHelpIndexWidget_ActivateCurrentItem(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "activateCurrentItem");
}

void QHelpIndexWidget_FilterIndices(QtObjectPtr ptr, char* filter, char* wildcard){
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "filterIndices", Q_ARG(QString, QString(filter)), Q_ARG(QString, QString(wildcard)));
}

void QHelpIndexWidget_ConnectLinkActivated(QtObjectPtr ptr){
	QObject::connect(static_cast<QHelpIndexWidget*>(ptr), static_cast<void (QHelpIndexWidget::*)(const QUrl &, const QString &)>(&QHelpIndexWidget::linkActivated), static_cast<MyQHelpIndexWidget*>(ptr), static_cast<void (MyQHelpIndexWidget::*)(const QUrl &, const QString &)>(&MyQHelpIndexWidget::Signal_LinkActivated));;
}

void QHelpIndexWidget_DisconnectLinkActivated(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QHelpIndexWidget*>(ptr), static_cast<void (QHelpIndexWidget::*)(const QUrl &, const QString &)>(&QHelpIndexWidget::linkActivated), static_cast<MyQHelpIndexWidget*>(ptr), static_cast<void (MyQHelpIndexWidget::*)(const QUrl &, const QString &)>(&MyQHelpIndexWidget::Signal_LinkActivated));;
}

