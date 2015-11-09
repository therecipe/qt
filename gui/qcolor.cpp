#include "qcolor.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QColor>
#include "_cgo_export.h"

class MyQColor: public QColor {
public:
};

void* QColor_ConvertTo(void* ptr, int colorSpec){
	return new QColor(static_cast<QColor*>(ptr)->convertTo(static_cast<QColor::Spec>(colorSpec)));
}

void QColor_SetRgbF(void* ptr, double r, double g, double b, double a){
	static_cast<QColor*>(ptr)->setRgbF(static_cast<qreal>(r), static_cast<qreal>(g), static_cast<qreal>(b), static_cast<qreal>(a));
}

void* QColor_NewQColor(){
	return new QColor();
}

void* QColor_NewQColor8(int color){
	return new QColor(static_cast<Qt::GlobalColor>(color));
}

void* QColor_NewQColor6(void* color){
	return new QColor(*static_cast<QColor*>(color));
}

void* QColor_NewQColor4(char* name){
	return new QColor(QString(name));
}

void* QColor_NewQColor5(char* name){
	return new QColor(const_cast<const char*>(name));
}

void* QColor_NewQColor2(int r, int g, int b, int a){
	return new QColor(r, g, b, a);
}

int QColor_Alpha(void* ptr){
	return static_cast<QColor*>(ptr)->alpha();
}

double QColor_AlphaF(void* ptr){
	return static_cast<double>(static_cast<QColor*>(ptr)->alphaF());
}

int QColor_Black(void* ptr){
	return static_cast<QColor*>(ptr)->black();
}

double QColor_BlackF(void* ptr){
	return static_cast<double>(static_cast<QColor*>(ptr)->blackF());
}

int QColor_Blue(void* ptr){
	return static_cast<QColor*>(ptr)->blue();
}

double QColor_BlueF(void* ptr){
	return static_cast<double>(static_cast<QColor*>(ptr)->blueF());
}

char* QColor_QColor_ColorNames(){
	return QColor::colorNames().join("|").toUtf8().data();
}

int QColor_Cyan(void* ptr){
	return static_cast<QColor*>(ptr)->cyan();
}

double QColor_CyanF(void* ptr){
	return static_cast<double>(static_cast<QColor*>(ptr)->cyanF());
}

void* QColor_Darker(void* ptr, int factor){
	return new QColor(static_cast<QColor*>(ptr)->darker(factor));
}

void* QColor_QColor_FromCmyk(int c, int m, int y, int k, int a){
	return new QColor(QColor::fromCmyk(c, m, y, k, a));
}

void* QColor_QColor_FromCmykF(double c, double m, double y, double k, double a){
	return new QColor(QColor::fromCmykF(static_cast<qreal>(c), static_cast<qreal>(m), static_cast<qreal>(y), static_cast<qreal>(k), static_cast<qreal>(a)));
}

void* QColor_QColor_FromHsl(int h, int s, int l, int a){
	return new QColor(QColor::fromHsl(h, s, l, a));
}

void* QColor_QColor_FromHslF(double h, double s, double l, double a){
	return new QColor(QColor::fromHslF(static_cast<qreal>(h), static_cast<qreal>(s), static_cast<qreal>(l), static_cast<qreal>(a)));
}

void* QColor_QColor_FromHsv(int h, int s, int v, int a){
	return new QColor(QColor::fromHsv(h, s, v, a));
}

void* QColor_QColor_FromHsvF(double h, double s, double v, double a){
	return new QColor(QColor::fromHsvF(static_cast<qreal>(h), static_cast<qreal>(s), static_cast<qreal>(v), static_cast<qreal>(a)));
}

void* QColor_QColor_FromRgb2(int r, int g, int b, int a){
	return new QColor(QColor::fromRgb(r, g, b, a));
}

void* QColor_QColor_FromRgbF(double r, double g, double b, double a){
	return new QColor(QColor::fromRgbF(static_cast<qreal>(r), static_cast<qreal>(g), static_cast<qreal>(b), static_cast<qreal>(a)));
}

void QColor_GetCmyk(void* ptr, int c, int m, int y, int k, int a){
	static_cast<QColor*>(ptr)->getCmyk(&c, &m, &y, &k, &a);
}

void QColor_GetHsl(void* ptr, int h, int s, int l, int a){
	static_cast<QColor*>(ptr)->getHsl(&h, &s, &l, &a);
}

void QColor_GetHsv(void* ptr, int h, int s, int v, int a){
	static_cast<QColor*>(ptr)->getHsv(&h, &s, &v, &a);
}

void QColor_GetRgb(void* ptr, int r, int g, int b, int a){
	static_cast<QColor*>(ptr)->getRgb(&r, &g, &b, &a);
}

int QColor_Green(void* ptr){
	return static_cast<QColor*>(ptr)->green();
}

double QColor_GreenF(void* ptr){
	return static_cast<double>(static_cast<QColor*>(ptr)->greenF());
}

int QColor_HslHue(void* ptr){
	return static_cast<QColor*>(ptr)->hslHue();
}

double QColor_HslHueF(void* ptr){
	return static_cast<double>(static_cast<QColor*>(ptr)->hslHueF());
}

int QColor_HslSaturation(void* ptr){
	return static_cast<QColor*>(ptr)->hslSaturation();
}

double QColor_HslSaturationF(void* ptr){
	return static_cast<double>(static_cast<QColor*>(ptr)->hslSaturationF());
}

int QColor_HsvHue(void* ptr){
	return static_cast<QColor*>(ptr)->hsvHue();
}

double QColor_HsvHueF(void* ptr){
	return static_cast<double>(static_cast<QColor*>(ptr)->hsvHueF());
}

int QColor_HsvSaturation(void* ptr){
	return static_cast<QColor*>(ptr)->hsvSaturation();
}

double QColor_HsvSaturationF(void* ptr){
	return static_cast<double>(static_cast<QColor*>(ptr)->hsvSaturationF());
}

int QColor_Hue(void* ptr){
	return static_cast<QColor*>(ptr)->hue();
}

double QColor_HueF(void* ptr){
	return static_cast<double>(static_cast<QColor*>(ptr)->hueF());
}

int QColor_IsValid(void* ptr){
	return static_cast<QColor*>(ptr)->isValid();
}

int QColor_QColor_IsValidColor(char* name){
	return QColor::isValidColor(QString(name));
}

void* QColor_Lighter(void* ptr, int factor){
	return new QColor(static_cast<QColor*>(ptr)->lighter(factor));
}

int QColor_Lightness(void* ptr){
	return static_cast<QColor*>(ptr)->lightness();
}

double QColor_LightnessF(void* ptr){
	return static_cast<double>(static_cast<QColor*>(ptr)->lightnessF());
}

int QColor_Magenta(void* ptr){
	return static_cast<QColor*>(ptr)->magenta();
}

double QColor_MagentaF(void* ptr){
	return static_cast<double>(static_cast<QColor*>(ptr)->magentaF());
}

char* QColor_Name(void* ptr){
	return static_cast<QColor*>(ptr)->name().toUtf8().data();
}

char* QColor_Name2(void* ptr, int format){
	return static_cast<QColor*>(ptr)->name(static_cast<QColor::NameFormat>(format)).toUtf8().data();
}

int QColor_Red(void* ptr){
	return static_cast<QColor*>(ptr)->red();
}

double QColor_RedF(void* ptr){
	return static_cast<double>(static_cast<QColor*>(ptr)->redF());
}

int QColor_Saturation(void* ptr){
	return static_cast<QColor*>(ptr)->saturation();
}

double QColor_SaturationF(void* ptr){
	return static_cast<double>(static_cast<QColor*>(ptr)->saturationF());
}

void QColor_SetAlpha(void* ptr, int alpha){
	static_cast<QColor*>(ptr)->setAlpha(alpha);
}

void QColor_SetAlphaF(void* ptr, double alpha){
	static_cast<QColor*>(ptr)->setAlphaF(static_cast<qreal>(alpha));
}

void QColor_SetBlue(void* ptr, int blue){
	static_cast<QColor*>(ptr)->setBlue(blue);
}

void QColor_SetBlueF(void* ptr, double blue){
	static_cast<QColor*>(ptr)->setBlueF(static_cast<qreal>(blue));
}

void QColor_SetCmyk(void* ptr, int c, int m, int y, int k, int a){
	static_cast<QColor*>(ptr)->setCmyk(c, m, y, k, a);
}

void QColor_SetCmykF(void* ptr, double c, double m, double y, double k, double a){
	static_cast<QColor*>(ptr)->setCmykF(static_cast<qreal>(c), static_cast<qreal>(m), static_cast<qreal>(y), static_cast<qreal>(k), static_cast<qreal>(a));
}

void QColor_SetGreen(void* ptr, int green){
	static_cast<QColor*>(ptr)->setGreen(green);
}

void QColor_SetGreenF(void* ptr, double green){
	static_cast<QColor*>(ptr)->setGreenF(static_cast<qreal>(green));
}

void QColor_SetHsl(void* ptr, int h, int s, int l, int a){
	static_cast<QColor*>(ptr)->setHsl(h, s, l, a);
}

void QColor_SetHslF(void* ptr, double h, double s, double l, double a){
	static_cast<QColor*>(ptr)->setHslF(static_cast<qreal>(h), static_cast<qreal>(s), static_cast<qreal>(l), static_cast<qreal>(a));
}

void QColor_SetHsv(void* ptr, int h, int s, int v, int a){
	static_cast<QColor*>(ptr)->setHsv(h, s, v, a);
}

void QColor_SetHsvF(void* ptr, double h, double s, double v, double a){
	static_cast<QColor*>(ptr)->setHsvF(static_cast<qreal>(h), static_cast<qreal>(s), static_cast<qreal>(v), static_cast<qreal>(a));
}

void QColor_SetNamedColor(void* ptr, char* name){
	static_cast<QColor*>(ptr)->setNamedColor(QString(name));
}

void QColor_SetRed(void* ptr, int red){
	static_cast<QColor*>(ptr)->setRed(red);
}

void QColor_SetRedF(void* ptr, double red){
	static_cast<QColor*>(ptr)->setRedF(static_cast<qreal>(red));
}

void QColor_SetRgb(void* ptr, int r, int g, int b, int a){
	static_cast<QColor*>(ptr)->setRgb(r, g, b, a);
}

int QColor_Spec(void* ptr){
	return static_cast<QColor*>(ptr)->spec();
}

void* QColor_ToCmyk(void* ptr){
	return new QColor(static_cast<QColor*>(ptr)->toCmyk());
}

void* QColor_ToHsl(void* ptr){
	return new QColor(static_cast<QColor*>(ptr)->toHsl());
}

void* QColor_ToHsv(void* ptr){
	return new QColor(static_cast<QColor*>(ptr)->toHsv());
}

void* QColor_ToRgb(void* ptr){
	return new QColor(static_cast<QColor*>(ptr)->toRgb());
}

int QColor_Value(void* ptr){
	return static_cast<QColor*>(ptr)->value();
}

double QColor_ValueF(void* ptr){
	return static_cast<double>(static_cast<QColor*>(ptr)->valueF());
}

int QColor_Yellow(void* ptr){
	return static_cast<QColor*>(ptr)->yellow();
}

double QColor_YellowF(void* ptr){
	return static_cast<double>(static_cast<QColor*>(ptr)->yellowF());
}

