#include "qapplication.h"
#include <QModelIndex>
#include <QStyle>
#include <QString>
#include <QWidget>
#include <QByteArray>
#include <QEvent>
#include <QVariant>
#include <QMetaObject>
#include <QSize>
#include <QPalette>
#include <QList>
#include <QPoint>
#include <QUrl>
#include <QFont>
#include <QObject>
#include <QIcon>
#include <QApplication>
#include "_cgo_export.h"

class MyQApplication: public QApplication {
public:
void Signal_FocusChanged(QWidget * old, QWidget * now){callbackQApplicationFocusChanged(this->objectName().toUtf8().data(), old, now);};
};

void QApplication_QApplication_Alert(void* widget, int msec){
	QApplication::alert(static_cast<QWidget*>(widget), msec);
}

int QApplication_AutoMaximizeThreshold(void* ptr){
	return QMetaObject::invokeMethod(static_cast<QApplication*>(ptr), "autoMaximizeThreshold");
}

int QApplication_AutoSipEnabled(void* ptr){
	return QMetaObject::invokeMethod(static_cast<QApplication*>(ptr), "autoSipEnabled");
}

void QApplication_QApplication_Beep(){
	QApplication::beep();
}

int QApplication_QApplication_CursorFlashTime(){
	return QApplication::cursorFlashTime();
}

int QApplication_QApplication_DoubleClickInterval(){
	return QApplication::doubleClickInterval();
}

int QApplication_QApplication_IsEffectEnabled(int effect){
	return QApplication::isEffectEnabled(static_cast<Qt::UIEffect>(effect));
}

int QApplication_QApplication_KeyboardInputInterval(){
	return QApplication::keyboardInputInterval();
}

void QApplication_QApplication_SetActiveWindow(void* active){
	QApplication::setActiveWindow(static_cast<QWidget*>(active));
}

void QApplication_SetAutoMaximizeThreshold(void* ptr, int threshold){
	QMetaObject::invokeMethod(static_cast<QApplication*>(ptr), "setAutoMaximizeThreshold", Q_ARG(int, threshold));
}

void QApplication_SetAutoSipEnabled(void* ptr, int enabled){
	QMetaObject::invokeMethod(static_cast<QApplication*>(ptr), "setAutoSipEnabled", Q_ARG(bool, enabled != 0));
}

void QApplication_QApplication_SetCursorFlashTime(int v){
	QApplication::setCursorFlashTime(v);
}

void QApplication_QApplication_SetDoubleClickInterval(int v){
	QApplication::setDoubleClickInterval(v);
}

void QApplication_QApplication_SetEffectEnabled(int effect, int enable){
	QApplication::setEffectEnabled(static_cast<Qt::UIEffect>(effect), enable != 0);
}

void QApplication_QApplication_SetGlobalStrut(void* v){
	QApplication::setGlobalStrut(*static_cast<QSize*>(v));
}

void QApplication_QApplication_SetKeyboardInputInterval(int v){
	QApplication::setKeyboardInputInterval(v);
}

void QApplication_QApplication_SetStartDragDistance(int l){
	QApplication::setStartDragDistance(l);
}

void QApplication_QApplication_SetStartDragTime(int ms){
	QApplication::setStartDragTime(ms);
}

void QApplication_SetStyleSheet(void* ptr, char* sheet){
	QMetaObject::invokeMethod(static_cast<QApplication*>(ptr), "setStyleSheet", Q_ARG(QString, QString(sheet)));
}

void QApplication_QApplication_SetWheelScrollLines(int v){
	QApplication::setWheelScrollLines(v);
}

void QApplication_QApplication_SetWindowIcon(void* icon){
	QApplication::setWindowIcon(*static_cast<QIcon*>(icon));
}

int QApplication_QApplication_StartDragDistance(){
	return QApplication::startDragDistance();
}

int QApplication_QApplication_StartDragTime(){
	return QApplication::startDragTime();
}

char* QApplication_StyleSheet(void* ptr){
	return static_cast<QApplication*>(ptr)->styleSheet().toUtf8().data();
}

void* QApplication_QApplication_TopLevelAt(void* point){
	return QApplication::topLevelAt(*static_cast<QPoint*>(point));
}

int QApplication_QApplication_WheelScrollLines(){
	return QApplication::wheelScrollLines();
}

void* QApplication_QApplication_WidgetAt(void* point){
	return QApplication::widgetAt(*static_cast<QPoint*>(point));
}

void* QApplication_NewQApplication(int argc, char* argv){
	QList<QByteArray> aList = QByteArray(argv).split('|');
	char *argvs[argc];
	static int argcs = argc;
	for (int i = 0; i < argc; i++)
		argvs[i] = aList[i].data();

	return new QApplication(argcs, argvs);
}

void QApplication_QApplication_AboutQt(){
	QMetaObject::invokeMethod(QApplication::instance(), "aboutQt");
}

void* QApplication_QApplication_ActiveModalWidget(){
	return QApplication::activeModalWidget();
}

void* QApplication_QApplication_ActivePopupWidget(){
	return QApplication::activePopupWidget();
}

void* QApplication_QApplication_ActiveWindow(){
	return QApplication::activeWindow();
}

void QApplication_QApplication_CloseAllWindows(){
	QMetaObject::invokeMethod(QApplication::instance(), "closeAllWindows");
}

int QApplication_QApplication_ColorSpec(){
	return QApplication::colorSpec();
}

void* QApplication_QApplication_Desktop(){
	return QApplication::desktop();
}

int QApplication_QApplication_Exec(){
	return QApplication::exec();
}

void QApplication_ConnectFocusChanged(void* ptr){
	QObject::connect(static_cast<QApplication*>(ptr), static_cast<void (QApplication::*)(QWidget *, QWidget *)>(&QApplication::focusChanged), static_cast<MyQApplication*>(ptr), static_cast<void (MyQApplication::*)(QWidget *, QWidget *)>(&MyQApplication::Signal_FocusChanged));;
}

void QApplication_DisconnectFocusChanged(void* ptr){
	QObject::disconnect(static_cast<QApplication*>(ptr), static_cast<void (QApplication::*)(QWidget *, QWidget *)>(&QApplication::focusChanged), static_cast<MyQApplication*>(ptr), static_cast<void (MyQApplication::*)(QWidget *, QWidget *)>(&MyQApplication::Signal_FocusChanged));;
}

void* QApplication_QApplication_FocusWidget(){
	return QApplication::focusWidget();
}

int QApplication_Notify(void* ptr, void* receiver, void* e){
	return static_cast<QApplication*>(ptr)->notify(static_cast<QObject*>(receiver), static_cast<QEvent*>(e));
}

void QApplication_QApplication_SetColorSpec(int spec){
	QApplication::setColorSpec(spec);
}

void QApplication_QApplication_SetFont(void* font, char* className){
	QApplication::setFont(*static_cast<QFont*>(font), const_cast<const char*>(className));
}

void QApplication_QApplication_SetPalette(void* palette, char* className){
	QApplication::setPalette(*static_cast<QPalette*>(palette), const_cast<const char*>(className));
}

void* QApplication_QApplication_SetStyle2(char* style){
	return QApplication::setStyle(QString(style));
}

void QApplication_QApplication_SetStyle(void* style){
	QApplication::setStyle(static_cast<QStyle*>(style));
}

void* QApplication_QApplication_Style(){
	return QApplication::style();
}

void* QApplication_QApplication_TopLevelAt2(int x, int y){
	return QApplication::topLevelAt(x, y);
}

void* QApplication_QApplication_WidgetAt2(int x, int y){
	return QApplication::widgetAt(x, y);
}

void QApplication_DestroyQApplication(void* ptr){
	static_cast<QApplication*>(ptr)->~QApplication();
}

