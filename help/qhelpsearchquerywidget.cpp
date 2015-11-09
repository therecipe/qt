#include "qhelpsearchquerywidget.h"
#include <QWidget>
#include <QObject>
#include <QHelpSearchQuery>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QHelpSearchQueryWidget>
#include "_cgo_export.h"

class MyQHelpSearchQueryWidget: public QHelpSearchQueryWidget {
public:
void Signal_Search(){callbackQHelpSearchQueryWidgetSearch(this->objectName().toUtf8().data());};
};

int QHelpSearchQueryWidget_IsCompactMode(void* ptr){
	return static_cast<QHelpSearchQueryWidget*>(ptr)->isCompactMode();
}

void* QHelpSearchQueryWidget_NewQHelpSearchQueryWidget(void* parent){
	return new QHelpSearchQueryWidget(static_cast<QWidget*>(parent));
}

void QHelpSearchQueryWidget_CollapseExtendedSearch(void* ptr){
	static_cast<QHelpSearchQueryWidget*>(ptr)->collapseExtendedSearch();
}

void QHelpSearchQueryWidget_ExpandExtendedSearch(void* ptr){
	static_cast<QHelpSearchQueryWidget*>(ptr)->expandExtendedSearch();
}

void QHelpSearchQueryWidget_ConnectSearch(void* ptr){
	QObject::connect(static_cast<QHelpSearchQueryWidget*>(ptr), static_cast<void (QHelpSearchQueryWidget::*)()>(&QHelpSearchQueryWidget::search), static_cast<MyQHelpSearchQueryWidget*>(ptr), static_cast<void (MyQHelpSearchQueryWidget::*)()>(&MyQHelpSearchQueryWidget::Signal_Search));;
}

void QHelpSearchQueryWidget_DisconnectSearch(void* ptr){
	QObject::disconnect(static_cast<QHelpSearchQueryWidget*>(ptr), static_cast<void (QHelpSearchQueryWidget::*)()>(&QHelpSearchQueryWidget::search), static_cast<MyQHelpSearchQueryWidget*>(ptr), static_cast<void (MyQHelpSearchQueryWidget::*)()>(&MyQHelpSearchQueryWidget::Signal_Search));;
}

void QHelpSearchQueryWidget_DestroyQHelpSearchQueryWidget(void* ptr){
	static_cast<QHelpSearchQueryWidget*>(ptr)->~QHelpSearchQueryWidget();
}

