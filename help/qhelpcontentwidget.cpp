#include "qhelpcontentwidget.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QHelpContentWidget>
#include "_cgo_export.h"

class MyQHelpContentWidget: public QHelpContentWidget {
public:
void Signal_LinkActivated(const QUrl & link){callbackQHelpContentWidgetLinkActivated(this->objectName().toUtf8().data(), link.toString().toUtf8().data());};
};

QtObjectPtr QHelpContentWidget_IndexOf(QtObjectPtr ptr, char* link){
	return static_cast<QHelpContentWidget*>(ptr)->indexOf(QUrl(QString(link))).internalPointer();
}

void QHelpContentWidget_ConnectLinkActivated(QtObjectPtr ptr){
	QObject::connect(static_cast<QHelpContentWidget*>(ptr), static_cast<void (QHelpContentWidget::*)(const QUrl &)>(&QHelpContentWidget::linkActivated), static_cast<MyQHelpContentWidget*>(ptr), static_cast<void (MyQHelpContentWidget::*)(const QUrl &)>(&MyQHelpContentWidget::Signal_LinkActivated));;
}

void QHelpContentWidget_DisconnectLinkActivated(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QHelpContentWidget*>(ptr), static_cast<void (QHelpContentWidget::*)(const QUrl &)>(&QHelpContentWidget::linkActivated), static_cast<MyQHelpContentWidget*>(ptr), static_cast<void (MyQHelpContentWidget::*)(const QUrl &)>(&MyQHelpContentWidget::Signal_LinkActivated));;
}

