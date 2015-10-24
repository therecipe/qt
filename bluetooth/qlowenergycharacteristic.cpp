#include "qlowenergycharacteristic.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QLowEnergyCharacteristic>
#include "_cgo_export.h"

class MyQLowEnergyCharacteristic: public QLowEnergyCharacteristic {
public:
};

QtObjectPtr QLowEnergyCharacteristic_NewQLowEnergyCharacteristic(){
	return new QLowEnergyCharacteristic();
}

QtObjectPtr QLowEnergyCharacteristic_NewQLowEnergyCharacteristic2(QtObjectPtr other){
	return new QLowEnergyCharacteristic(*static_cast<QLowEnergyCharacteristic*>(other));
}

int QLowEnergyCharacteristic_IsValid(QtObjectPtr ptr){
	return static_cast<QLowEnergyCharacteristic*>(ptr)->isValid();
}

char* QLowEnergyCharacteristic_Name(QtObjectPtr ptr){
	return static_cast<QLowEnergyCharacteristic*>(ptr)->name().toUtf8().data();
}

int QLowEnergyCharacteristic_Properties(QtObjectPtr ptr){
	return static_cast<QLowEnergyCharacteristic*>(ptr)->properties();
}

void QLowEnergyCharacteristic_DestroyQLowEnergyCharacteristic(QtObjectPtr ptr){
	static_cast<QLowEnergyCharacteristic*>(ptr)->~QLowEnergyCharacteristic();
}

