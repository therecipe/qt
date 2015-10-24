#include "qhelpsearchquerywidget.h"
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QHelpSearchQuery>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QHelpSearchQueryWidget>
#include "_cgo_export.h"

class MyQHelpSearchQueryWidget: public QHelpSearchQueryWidget {
public:
void Signal_Search(){callbackQHelpSearchQueryWidgetSearch(this->objectName().toUtf8().data());};
};

int QHelpSearchQueryWidget_IsCompactMode(QtObjectPtr ptr){
	return static_cast<QHelpSearchQueryWidget*>(ptr)->isCompactMode();
}

QtObjectPtr QHelpSearchQueryWidget_NewQHelpSearchQueryWidget(QtObjectPtr parent){
	return new QHelpSearchQueryWidget(static_cast<QWidget*>(parent));
}

void QHelpSearchQueryWidget_CollapseExtendedSearch(QtObjectPtr ptr){
	static_cast<QHelpSearchQueryWidget*>(ptr)->collapseExtendedSearch();
}

void QHelpSearchQueryWidget_ExpandExtendedSearch(QtObjectPtr ptr){
	static_cast<QHelpSearchQueryWidget*>(ptr)->expandExtendedSearch();
}

void QHelpSearchQueryWidget_ConnectSearch(QtObjectPtr ptr){
	QObject::connect(static_cast<QHelpSearchQueryWidget*>(ptr), static_cast<void (QHelpSearchQueryWidget::*)()>(&QHelpSearchQueryWidget::search), static_cast<MyQHelpSearchQueryWidget*>(ptr), static_cast<void (MyQHelpSearchQueryWidget::*)()>(&MyQHelpSearchQueryWidget::Signal_Search));;
}

void QHelpSearchQueryWidget_DisconnectSearch(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QHelpSearchQueryWidget*>(ptr), static_cast<void (QHelpSearchQueryWidget::*)()>(&QHelpSearchQueryWidget::search), static_cast<MyQHelpSearchQueryWidget*>(ptr), static_cast<void (MyQHelpSearchQueryWidget::*)()>(&MyQHelpSearchQueryWidget::Signal_Search));;
}

void QHelpSearchQueryWidget_DestroyQHelpSearchQueryWidget(QtObjectPtr ptr){
	static_cast<QHelpSearchQueryWidget*>(ptr)->~QHelpSearchQueryWidget();
}

