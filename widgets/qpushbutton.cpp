#include "qpushbutton.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QWidget>
#include <QIcon>
#include <QMenu>
#include <QString>
#include <QPushButton>
#include "_cgo_export.h"

class MyQPushButton: public QPushButton {
public:
};

int QPushButton_AutoDefault(void* ptr){
	return static_cast<QPushButton*>(ptr)->autoDefault();
}

int QPushButton_IsDefault(void* ptr){
	return static_cast<QPushButton*>(ptr)->isDefault();
}

int QPushButton_IsFlat(void* ptr){
	return static_cast<QPushButton*>(ptr)->isFlat();
}

void QPushButton_SetAutoDefault(void* ptr, int v){
	static_cast<QPushButton*>(ptr)->setAutoDefault(v != 0);
}

void QPushButton_SetDefault(void* ptr, int v){
	static_cast<QPushButton*>(ptr)->setDefault(v != 0);
}

void QPushButton_SetFlat(void* ptr, int v){
	static_cast<QPushButton*>(ptr)->setFlat(v != 0);
}

void* QPushButton_NewQPushButton(void* parent){
	return new QPushButton(static_cast<QWidget*>(parent));
}

void* QPushButton_NewQPushButton3(void* icon, char* text, void* parent){
	return new QPushButton(*static_cast<QIcon*>(icon), QString(text), static_cast<QWidget*>(parent));
}

void* QPushButton_NewQPushButton2(char* text, void* parent){
	return new QPushButton(QString(text), static_cast<QWidget*>(parent));
}

void* QPushButton_Menu(void* ptr){
	return static_cast<QPushButton*>(ptr)->menu();
}

void QPushButton_SetMenu(void* ptr, void* menu){
	static_cast<QPushButton*>(ptr)->setMenu(static_cast<QMenu*>(menu));
}

void QPushButton_ShowMenu(void* ptr){
	QMetaObject::invokeMethod(static_cast<QPushButton*>(ptr), "showMenu");
}

void QPushButton_DestroyQPushButton(void* ptr){
	static_cast<QPushButton*>(ptr)->~QPushButton();
}

