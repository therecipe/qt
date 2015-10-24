#include "qdockwidget.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QObject>
#include <QDockWidget>
#include "_cgo_export.h"

class MyQDockWidget: public QDockWidget {
public:
void Signal_AllowedAreasChanged(Qt::DockWidgetAreas allowedAreas){callbackQDockWidgetAllowedAreasChanged(this->objectName().toUtf8().data(), allowedAreas);};
void Signal_DockLocationChanged(Qt::DockWidgetArea area){callbackQDockWidgetDockLocationChanged(this->objectName().toUtf8().data(), area);};
void Signal_FeaturesChanged(QDockWidget::DockWidgetFeatures features){callbackQDockWidgetFeaturesChanged(this->objectName().toUtf8().data(), features);};
void Signal_TopLevelChanged(bool topLevel){callbackQDockWidgetTopLevelChanged(this->objectName().toUtf8().data(), topLevel);};
void Signal_VisibilityChanged(bool visible){callbackQDockWidgetVisibilityChanged(this->objectName().toUtf8().data(), visible);};
};

int QDockWidget_AllowedAreas(QtObjectPtr ptr){
	return static_cast<QDockWidget*>(ptr)->allowedAreas();
}

int QDockWidget_Features(QtObjectPtr ptr){
	return static_cast<QDockWidget*>(ptr)->features();
}

void QDockWidget_SetAllowedAreas(QtObjectPtr ptr, int areas){
	static_cast<QDockWidget*>(ptr)->setAllowedAreas(static_cast<Qt::DockWidgetArea>(areas));
}

void QDockWidget_SetFeatures(QtObjectPtr ptr, int features){
	static_cast<QDockWidget*>(ptr)->setFeatures(static_cast<QDockWidget::DockWidgetFeature>(features));
}

void QDockWidget_SetFloating(QtObjectPtr ptr, int floating){
	static_cast<QDockWidget*>(ptr)->setFloating(floating != 0);
}

QtObjectPtr QDockWidget_NewQDockWidget2(QtObjectPtr parent, int flags){
	return new QDockWidget(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

QtObjectPtr QDockWidget_NewQDockWidget(char* title, QtObjectPtr parent, int flags){
	return new QDockWidget(QString(title), static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void QDockWidget_ConnectAllowedAreasChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QDockWidget*>(ptr), static_cast<void (QDockWidget::*)(Qt::DockWidgetAreas)>(&QDockWidget::allowedAreasChanged), static_cast<MyQDockWidget*>(ptr), static_cast<void (MyQDockWidget::*)(Qt::DockWidgetAreas)>(&MyQDockWidget::Signal_AllowedAreasChanged));;
}

void QDockWidget_DisconnectAllowedAreasChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QDockWidget*>(ptr), static_cast<void (QDockWidget::*)(Qt::DockWidgetAreas)>(&QDockWidget::allowedAreasChanged), static_cast<MyQDockWidget*>(ptr), static_cast<void (MyQDockWidget::*)(Qt::DockWidgetAreas)>(&MyQDockWidget::Signal_AllowedAreasChanged));;
}

void QDockWidget_ConnectDockLocationChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QDockWidget*>(ptr), static_cast<void (QDockWidget::*)(Qt::DockWidgetArea)>(&QDockWidget::dockLocationChanged), static_cast<MyQDockWidget*>(ptr), static_cast<void (MyQDockWidget::*)(Qt::DockWidgetArea)>(&MyQDockWidget::Signal_DockLocationChanged));;
}

void QDockWidget_DisconnectDockLocationChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QDockWidget*>(ptr), static_cast<void (QDockWidget::*)(Qt::DockWidgetArea)>(&QDockWidget::dockLocationChanged), static_cast<MyQDockWidget*>(ptr), static_cast<void (MyQDockWidget::*)(Qt::DockWidgetArea)>(&MyQDockWidget::Signal_DockLocationChanged));;
}

void QDockWidget_ConnectFeaturesChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QDockWidget*>(ptr), static_cast<void (QDockWidget::*)(QDockWidget::DockWidgetFeatures)>(&QDockWidget::featuresChanged), static_cast<MyQDockWidget*>(ptr), static_cast<void (MyQDockWidget::*)(QDockWidget::DockWidgetFeatures)>(&MyQDockWidget::Signal_FeaturesChanged));;
}

void QDockWidget_DisconnectFeaturesChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QDockWidget*>(ptr), static_cast<void (QDockWidget::*)(QDockWidget::DockWidgetFeatures)>(&QDockWidget::featuresChanged), static_cast<MyQDockWidget*>(ptr), static_cast<void (MyQDockWidget::*)(QDockWidget::DockWidgetFeatures)>(&MyQDockWidget::Signal_FeaturesChanged));;
}

int QDockWidget_IsAreaAllowed(QtObjectPtr ptr, int area){
	return static_cast<QDockWidget*>(ptr)->isAreaAllowed(static_cast<Qt::DockWidgetArea>(area));
}

int QDockWidget_IsFloating(QtObjectPtr ptr){
	return static_cast<QDockWidget*>(ptr)->isFloating();
}

void QDockWidget_SetTitleBarWidget(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QDockWidget*>(ptr)->setTitleBarWidget(static_cast<QWidget*>(widget));
}

void QDockWidget_SetWidget(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QDockWidget*>(ptr)->setWidget(static_cast<QWidget*>(widget));
}

QtObjectPtr QDockWidget_TitleBarWidget(QtObjectPtr ptr){
	return static_cast<QDockWidget*>(ptr)->titleBarWidget();
}

QtObjectPtr QDockWidget_ToggleViewAction(QtObjectPtr ptr){
	return static_cast<QDockWidget*>(ptr)->toggleViewAction();
}

void QDockWidget_ConnectTopLevelChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QDockWidget*>(ptr), static_cast<void (QDockWidget::*)(bool)>(&QDockWidget::topLevelChanged), static_cast<MyQDockWidget*>(ptr), static_cast<void (MyQDockWidget::*)(bool)>(&MyQDockWidget::Signal_TopLevelChanged));;
}

void QDockWidget_DisconnectTopLevelChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QDockWidget*>(ptr), static_cast<void (QDockWidget::*)(bool)>(&QDockWidget::topLevelChanged), static_cast<MyQDockWidget*>(ptr), static_cast<void (MyQDockWidget::*)(bool)>(&MyQDockWidget::Signal_TopLevelChanged));;
}

void QDockWidget_ConnectVisibilityChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QDockWidget*>(ptr), static_cast<void (QDockWidget::*)(bool)>(&QDockWidget::visibilityChanged), static_cast<MyQDockWidget*>(ptr), static_cast<void (MyQDockWidget::*)(bool)>(&MyQDockWidget::Signal_VisibilityChanged));;
}

void QDockWidget_DisconnectVisibilityChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QDockWidget*>(ptr), static_cast<void (QDockWidget::*)(bool)>(&QDockWidget::visibilityChanged), static_cast<MyQDockWidget*>(ptr), static_cast<void (MyQDockWidget::*)(bool)>(&MyQDockWidget::Signal_VisibilityChanged));;
}

QtObjectPtr QDockWidget_Widget(QtObjectPtr ptr){
	return static_cast<QDockWidget*>(ptr)->widget();
}

void QDockWidget_DestroyQDockWidget(QtObjectPtr ptr){
	static_cast<QDockWidget*>(ptr)->~QDockWidget();
}

