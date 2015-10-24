#include "qpushbutton.h"
#include <QWidget>
#include <QMenu>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QIcon>
#include <QPushButton>
#include "_cgo_export.h"

class MyQPushButton: public QPushButton {
public:
};

int QPushButton_AutoDefault(QtObjectPtr ptr){
	return static_cast<QPushButton*>(ptr)->autoDefault();
}

int QPushButton_IsDefault(QtObjectPtr ptr){
	return static_cast<QPushButton*>(ptr)->isDefault();
}

int QPushButton_IsFlat(QtObjectPtr ptr){
	return static_cast<QPushButton*>(ptr)->isFlat();
}

void QPushButton_SetAutoDefault(QtObjectPtr ptr, int v){
	static_cast<QPushButton*>(ptr)->setAutoDefault(v != 0);
}

void QPushButton_SetDefault(QtObjectPtr ptr, int v){
	static_cast<QPushButton*>(ptr)->setDefault(v != 0);
}

void QPushButton_SetFlat(QtObjectPtr ptr, int v){
	static_cast<QPushButton*>(ptr)->setFlat(v != 0);
}

QtObjectPtr QPushButton_NewQPushButton(QtObjectPtr parent){
	return new QPushButton(static_cast<QWidget*>(parent));
}

QtObjectPtr QPushButton_NewQPushButton3(QtObjectPtr icon, char* text, QtObjectPtr parent){
	return new QPushButton(*static_cast<QIcon*>(icon), QString(text), static_cast<QWidget*>(parent));
}

QtObjectPtr QPushButton_NewQPushButton2(char* text, QtObjectPtr parent){
	return new QPushButton(QString(text), static_cast<QWidget*>(parent));
}

QtObjectPtr QPushButton_Menu(QtObjectPtr ptr){
	return static_cast<QPushButton*>(ptr)->menu();
}

void QPushButton_SetMenu(QtObjectPtr ptr, QtObjectPtr menu){
	static_cast<QPushButton*>(ptr)->setMenu(static_cast<QMenu*>(menu));
}

void QPushButton_ShowMenu(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QPushButton*>(ptr), "showMenu");
}

void QPushButton_DestroyQPushButton(QtObjectPtr ptr){
	static_cast<QPushButton*>(ptr)->~QPushButton();
}

