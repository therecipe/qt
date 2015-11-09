#include "qwidgetaction.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QWidgetAction>
#include "_cgo_export.h"

class MyQWidgetAction: public QWidgetAction {
public:
};

void* QWidgetAction_NewQWidgetAction(void* parent){
	return new QWidgetAction(static_cast<QObject*>(parent));
}

void* QWidgetAction_DefaultWidget(void* ptr){
	return static_cast<QWidgetAction*>(ptr)->defaultWidget();
}

void QWidgetAction_ReleaseWidget(void* ptr, void* widget){
	static_cast<QWidgetAction*>(ptr)->releaseWidget(static_cast<QWidget*>(widget));
}

void* QWidgetAction_RequestWidget(void* ptr, void* parent){
	return static_cast<QWidgetAction*>(ptr)->requestWidget(static_cast<QWidget*>(parent));
}

void QWidgetAction_SetDefaultWidget(void* ptr, void* widget){
	static_cast<QWidgetAction*>(ptr)->setDefaultWidget(static_cast<QWidget*>(widget));
}

void QWidgetAction_DestroyQWidgetAction(void* ptr){
	static_cast<QWidgetAction*>(ptr)->~QWidgetAction();
}

