#include "qcolor.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QColor>
#include "_cgo_export.h"

class MyQColor: public QColor {
public:
};

QtObjectPtr QColor_NewQColor(){
	return new QColor();
}

QtObjectPtr QColor_NewQColor8(int color){
	return new QColor(static_cast<Qt::GlobalColor>(color));
}

QtObjectPtr QColor_NewQColor6(QtObjectPtr color){
	return new QColor(*static_cast<QColor*>(color));
}

QtObjectPtr QColor_NewQColor4(char* name){
	return new QColor(QString(name));
}

QtObjectPtr QColor_NewQColor5(char* name){
	return new QColor(const_cast<const char*>(name));
}

QtObjectPtr QColor_NewQColor2(int r, int g, int b, int a){
	return new QColor(r, g, b, a);
}

int QColor_Alpha(QtObjectPtr ptr){
	return static_cast<QColor*>(ptr)->alpha();
}

int QColor_Black(QtObjectPtr ptr){
	return static_cast<QColor*>(ptr)->black();
}

int QColor_Blue(QtObjectPtr ptr){
	return static_cast<QColor*>(ptr)->blue();
}

char* QColor_QColor_ColorNames(){
	return QColor::colorNames().join("|").toUtf8().data();
}

int QColor_Cyan(QtObjectPtr ptr){
	return static_cast<QColor*>(ptr)->cyan();
}

void QColor_GetCmyk(QtObjectPtr ptr, int c, int m, int y, int k, int a){
	static_cast<QColor*>(ptr)->getCmyk(&c, &m, &y, &k, &a);
}

void QColor_GetHsl(QtObjectPtr ptr, int h, int s, int l, int a){
	static_cast<QColor*>(ptr)->getHsl(&h, &s, &l, &a);
}

void QColor_GetHsv(QtObjectPtr ptr, int h, int s, int v, int a){
	static_cast<QColor*>(ptr)->getHsv(&h, &s, &v, &a);
}

void QColor_GetRgb(QtObjectPtr ptr, int r, int g, int b, int a){
	static_cast<QColor*>(ptr)->getRgb(&r, &g, &b, &a);
}

int QColor_Green(QtObjectPtr ptr){
	return static_cast<QColor*>(ptr)->green();
}

int QColor_HslHue(QtObjectPtr ptr){
	return static_cast<QColor*>(ptr)->hslHue();
}

int QColor_HslSaturation(QtObjectPtr ptr){
	return static_cast<QColor*>(ptr)->hslSaturation();
}

int QColor_HsvHue(QtObjectPtr ptr){
	return static_cast<QColor*>(ptr)->hsvHue();
}

int QColor_HsvSaturation(QtObjectPtr ptr){
	return static_cast<QColor*>(ptr)->hsvSaturation();
}

int QColor_Hue(QtObjectPtr ptr){
	return static_cast<QColor*>(ptr)->hue();
}

int QColor_IsValid(QtObjectPtr ptr){
	return static_cast<QColor*>(ptr)->isValid();
}

int QColor_QColor_IsValidColor(char* name){
	return QColor::isValidColor(QString(name));
}

int QColor_Lightness(QtObjectPtr ptr){
	return static_cast<QColor*>(ptr)->lightness();
}

int QColor_Magenta(QtObjectPtr ptr){
	return static_cast<QColor*>(ptr)->magenta();
}

char* QColor_Name(QtObjectPtr ptr){
	return static_cast<QColor*>(ptr)->name().toUtf8().data();
}

char* QColor_Name2(QtObjectPtr ptr, int format){
	return static_cast<QColor*>(ptr)->name(static_cast<QColor::NameFormat>(format)).toUtf8().data();
}

int QColor_Red(QtObjectPtr ptr){
	return static_cast<QColor*>(ptr)->red();
}

int QColor_Saturation(QtObjectPtr ptr){
	return static_cast<QColor*>(ptr)->saturation();
}

void QColor_SetAlpha(QtObjectPtr ptr, int alpha){
	static_cast<QColor*>(ptr)->setAlpha(alpha);
}

void QColor_SetBlue(QtObjectPtr ptr, int blue){
	static_cast<QColor*>(ptr)->setBlue(blue);
}

void QColor_SetCmyk(QtObjectPtr ptr, int c, int m, int y, int k, int a){
	static_cast<QColor*>(ptr)->setCmyk(c, m, y, k, a);
}

void QColor_SetGreen(QtObjectPtr ptr, int green){
	static_cast<QColor*>(ptr)->setGreen(green);
}

void QColor_SetHsl(QtObjectPtr ptr, int h, int s, int l, int a){
	static_cast<QColor*>(ptr)->setHsl(h, s, l, a);
}

void QColor_SetHsv(QtObjectPtr ptr, int h, int s, int v, int a){
	static_cast<QColor*>(ptr)->setHsv(h, s, v, a);
}

void QColor_SetNamedColor(QtObjectPtr ptr, char* name){
	static_cast<QColor*>(ptr)->setNamedColor(QString(name));
}

void QColor_SetRed(QtObjectPtr ptr, int red){
	static_cast<QColor*>(ptr)->setRed(red);
}

void QColor_SetRgb(QtObjectPtr ptr, int r, int g, int b, int a){
	static_cast<QColor*>(ptr)->setRgb(r, g, b, a);
}

int QColor_Spec(QtObjectPtr ptr){
	return static_cast<QColor*>(ptr)->spec();
}

int QColor_Value(QtObjectPtr ptr){
	return static_cast<QColor*>(ptr)->value();
}

int QColor_Yellow(QtObjectPtr ptr){
	return static_cast<QColor*>(ptr)->yellow();
}

