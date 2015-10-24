#include "qwidgetaction.h"
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QWidgetAction>
#include "_cgo_export.h"

class MyQWidgetAction: public QWidgetAction {
public:
};

QtObjectPtr QWidgetAction_NewQWidgetAction(QtObjectPtr parent){
	return new QWidgetAction(static_cast<QObject*>(parent));
}

QtObjectPtr QWidgetAction_DefaultWidget(QtObjectPtr ptr){
	return static_cast<QWidgetAction*>(ptr)->defaultWidget();
}

void QWidgetAction_ReleaseWidget(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QWidgetAction*>(ptr)->releaseWidget(static_cast<QWidget*>(widget));
}

QtObjectPtr QWidgetAction_RequestWidget(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QWidgetAction*>(ptr)->requestWidget(static_cast<QWidget*>(parent));
}

void QWidgetAction_SetDefaultWidget(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QWidgetAction*>(ptr)->setDefaultWidget(static_cast<QWidget*>(widget));
}

void QWidgetAction_DestroyQWidgetAction(QtObjectPtr ptr){
	static_cast<QWidgetAction*>(ptr)->~QWidgetAction();
}

