#include "qstyleoption.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QStyle>
#include <QString>
#include <QStyleOption>
#include "_cgo_export.h"

class MyQStyleOption: public QStyleOption {
public:
};

int QStyleOption_SO_Slider_Type(){
	return QStyleOption::SO_Slider;
}

int QStyleOption_SO_SpinBox_Type(){
	return QStyleOption::SO_SpinBox;
}

int QStyleOption_SO_ToolButton_Type(){
	return QStyleOption::SO_ToolButton;
}

int QStyleOption_SO_ComboBox_Type(){
	return QStyleOption::SO_ComboBox;
}

int QStyleOption_SO_TitleBar_Type(){
	return QStyleOption::SO_TitleBar;
}

int QStyleOption_SO_GroupBox_Type(){
	return QStyleOption::SO_GroupBox;
}

int QStyleOption_SO_SizeGrip_Type(){
	return QStyleOption::SO_SizeGrip;
}

void* QStyleOption_NewQStyleOption2(void* other){
	return new QStyleOption(*static_cast<QStyleOption*>(other));
}

void* QStyleOption_NewQStyleOption(int version, int ty){
	return new QStyleOption(version, ty);
}

void QStyleOption_InitFrom(void* ptr, void* widget){
	static_cast<QStyleOption*>(ptr)->initFrom(static_cast<QWidget*>(widget));
}

void QStyleOption_DestroyQStyleOption(void* ptr){
	static_cast<QStyleOption*>(ptr)->~QStyleOption();
}

