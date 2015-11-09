#include "qlowenergycharacteristic.h"
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QString>
#include <QVariant>
#include <QLowEnergyCharacteristic>
#include "_cgo_export.h"

class MyQLowEnergyCharacteristic: public QLowEnergyCharacteristic {
public:
};

void* QLowEnergyCharacteristic_NewQLowEnergyCharacteristic(){
	return new QLowEnergyCharacteristic();
}

void* QLowEnergyCharacteristic_NewQLowEnergyCharacteristic2(void* other){
	return new QLowEnergyCharacteristic(*static_cast<QLowEnergyCharacteristic*>(other));
}

int QLowEnergyCharacteristic_IsValid(void* ptr){
	return static_cast<QLowEnergyCharacteristic*>(ptr)->isValid();
}

char* QLowEnergyCharacteristic_Name(void* ptr){
	return static_cast<QLowEnergyCharacteristic*>(ptr)->name().toUtf8().data();
}

int QLowEnergyCharacteristic_Properties(void* ptr){
	return static_cast<QLowEnergyCharacteristic*>(ptr)->properties();
}

void* QLowEnergyCharacteristic_Value(void* ptr){
	return new QByteArray(static_cast<QLowEnergyCharacteristic*>(ptr)->value());
}

void QLowEnergyCharacteristic_DestroyQLowEnergyCharacteristic(void* ptr){
	static_cast<QLowEnergyCharacteristic*>(ptr)->~QLowEnergyCharacteristic();
}

