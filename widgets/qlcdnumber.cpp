#include "qlcdnumber.h"
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QLCDNumber>
#include "_cgo_export.h"

class MyQLCDNumber: public QLCDNumber {
public:
void Signal_Overflow(){callbackQLCDNumberOverflow(this->objectName().toUtf8().data());};
};

int QLCDNumber_IntValue(QtObjectPtr ptr){
	return static_cast<QLCDNumber*>(ptr)->intValue();
}

int QLCDNumber_Mode(QtObjectPtr ptr){
	return static_cast<QLCDNumber*>(ptr)->mode();
}

int QLCDNumber_SegmentStyle(QtObjectPtr ptr){
	return static_cast<QLCDNumber*>(ptr)->segmentStyle();
}

void QLCDNumber_SetMode(QtObjectPtr ptr, int v){
	static_cast<QLCDNumber*>(ptr)->setMode(static_cast<QLCDNumber::Mode>(v));
}

void QLCDNumber_SetSegmentStyle(QtObjectPtr ptr, int v){
	static_cast<QLCDNumber*>(ptr)->setSegmentStyle(static_cast<QLCDNumber::SegmentStyle>(v));
}

void QLCDNumber_SetSmallDecimalPoint(QtObjectPtr ptr, int v){
	QMetaObject::invokeMethod(static_cast<QLCDNumber*>(ptr), "setSmallDecimalPoint", Q_ARG(bool, v != 0));
}

int QLCDNumber_SmallDecimalPoint(QtObjectPtr ptr){
	return static_cast<QLCDNumber*>(ptr)->smallDecimalPoint();
}

QtObjectPtr QLCDNumber_NewQLCDNumber(QtObjectPtr parent){
	return new QLCDNumber(static_cast<QWidget*>(parent));
}

int QLCDNumber_CheckOverflow2(QtObjectPtr ptr, int num){
	return static_cast<QLCDNumber*>(ptr)->checkOverflow(num);
}

int QLCDNumber_DigitCount(QtObjectPtr ptr){
	return static_cast<QLCDNumber*>(ptr)->digitCount();
}

void QLCDNumber_Display(QtObjectPtr ptr, char* s){
	QMetaObject::invokeMethod(static_cast<QLCDNumber*>(ptr), "display", Q_ARG(QString, QString(s)));
}

void QLCDNumber_Display3(QtObjectPtr ptr, int num){
	QMetaObject::invokeMethod(static_cast<QLCDNumber*>(ptr), "display", Q_ARG(int, num));
}

void QLCDNumber_ConnectOverflow(QtObjectPtr ptr){
	QObject::connect(static_cast<QLCDNumber*>(ptr), static_cast<void (QLCDNumber::*)()>(&QLCDNumber::overflow), static_cast<MyQLCDNumber*>(ptr), static_cast<void (MyQLCDNumber::*)()>(&MyQLCDNumber::Signal_Overflow));;
}

void QLCDNumber_DisconnectOverflow(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QLCDNumber*>(ptr), static_cast<void (QLCDNumber::*)()>(&QLCDNumber::overflow), static_cast<MyQLCDNumber*>(ptr), static_cast<void (MyQLCDNumber::*)()>(&MyQLCDNumber::Signal_Overflow));;
}

void QLCDNumber_SetBinMode(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QLCDNumber*>(ptr), "setBinMode");
}

void QLCDNumber_SetDecMode(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QLCDNumber*>(ptr), "setDecMode");
}

void QLCDNumber_SetDigitCount(QtObjectPtr ptr, int numDigits){
	static_cast<QLCDNumber*>(ptr)->setDigitCount(numDigits);
}

void QLCDNumber_SetHexMode(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QLCDNumber*>(ptr), "setHexMode");
}

void QLCDNumber_SetOctMode(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QLCDNumber*>(ptr), "setOctMode");
}

void QLCDNumber_DestroyQLCDNumber(QtObjectPtr ptr){
	static_cast<QLCDNumber*>(ptr)->~QLCDNumber();
}

