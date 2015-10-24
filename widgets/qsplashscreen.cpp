#include "qsplashscreen.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QColor>
#include <QPixmap>
#include <QWidget>
#include <QSplashScreen>
#include "_cgo_export.h"

class MyQSplashScreen: public QSplashScreen {
public:
void Signal_MessageChanged(const QString & message){callbackQSplashScreenMessageChanged(this->objectName().toUtf8().data(), message.toUtf8().data());};
};

QtObjectPtr QSplashScreen_NewQSplashScreen2(QtObjectPtr parent, QtObjectPtr pixmap, int f){
	return new QSplashScreen(static_cast<QWidget*>(parent), *static_cast<QPixmap*>(pixmap), static_cast<Qt::WindowType>(f));
}

QtObjectPtr QSplashScreen_NewQSplashScreen(QtObjectPtr pixmap, int f){
	return new QSplashScreen(*static_cast<QPixmap*>(pixmap), static_cast<Qt::WindowType>(f));
}

void QSplashScreen_ClearMessage(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QSplashScreen*>(ptr), "clearMessage");
}

void QSplashScreen_Finish(QtObjectPtr ptr, QtObjectPtr mainWin){
	static_cast<QSplashScreen*>(ptr)->finish(static_cast<QWidget*>(mainWin));
}

char* QSplashScreen_Message(QtObjectPtr ptr){
	return static_cast<QSplashScreen*>(ptr)->message().toUtf8().data();
}

void QSplashScreen_ConnectMessageChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSplashScreen*>(ptr), static_cast<void (QSplashScreen::*)(const QString &)>(&QSplashScreen::messageChanged), static_cast<MyQSplashScreen*>(ptr), static_cast<void (MyQSplashScreen::*)(const QString &)>(&MyQSplashScreen::Signal_MessageChanged));;
}

void QSplashScreen_DisconnectMessageChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSplashScreen*>(ptr), static_cast<void (QSplashScreen::*)(const QString &)>(&QSplashScreen::messageChanged), static_cast<MyQSplashScreen*>(ptr), static_cast<void (MyQSplashScreen::*)(const QString &)>(&MyQSplashScreen::Signal_MessageChanged));;
}

void QSplashScreen_Repaint(QtObjectPtr ptr){
	static_cast<QSplashScreen*>(ptr)->repaint();
}

void QSplashScreen_SetPixmap(QtObjectPtr ptr, QtObjectPtr pixmap){
	static_cast<QSplashScreen*>(ptr)->setPixmap(*static_cast<QPixmap*>(pixmap));
}

void QSplashScreen_ShowMessage(QtObjectPtr ptr, char* message, int alignment, QtObjectPtr color){
	QMetaObject::invokeMethod(static_cast<QSplashScreen*>(ptr), "showMessage", Q_ARG(QString, QString(message)), Q_ARG(int, alignment), Q_ARG(QColor, *static_cast<QColor*>(color)));
}

void QSplashScreen_DestroyQSplashScreen(QtObjectPtr ptr){
	static_cast<QSplashScreen*>(ptr)->~QSplashScreen();
}

