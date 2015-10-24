#include "qstyleoption.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QStyle>
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

QtObjectPtr QStyleOption_NewQStyleOption2(QtObjectPtr other){
	return new QStyleOption(*static_cast<QStyleOption*>(other));
}

QtObjectPtr QStyleOption_NewQStyleOption(int version, int ty){
	return new QStyleOption(version, ty);
}

void QStyleOption_InitFrom(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QStyleOption*>(ptr)->initFrom(static_cast<QWidget*>(widget));
}

void QStyleOption_DestroyQStyleOption(QtObjectPtr ptr){
	static_cast<QStyleOption*>(ptr)->~QStyleOption();
}

