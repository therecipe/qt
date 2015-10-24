#include "qlowenergydescriptor.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QLowEnergyDescriptor>
#include "_cgo_export.h"

class MyQLowEnergyDescriptor: public QLowEnergyDescriptor {
public:
};

QtObjectPtr QLowEnergyDescriptor_NewQLowEnergyDescriptor(){
	return new QLowEnergyDescriptor();
}

QtObjectPtr QLowEnergyDescriptor_NewQLowEnergyDescriptor2(QtObjectPtr other){
	return new QLowEnergyDescriptor(*static_cast<QLowEnergyDescriptor*>(other));
}

int QLowEnergyDescriptor_IsValid(QtObjectPtr ptr){
	return static_cast<QLowEnergyDescriptor*>(ptr)->isValid();
}

char* QLowEnergyDescriptor_Name(QtObjectPtr ptr){
	return static_cast<QLowEnergyDescriptor*>(ptr)->name().toUtf8().data();
}

int QLowEnergyDescriptor_Type(QtObjectPtr ptr){
	return static_cast<QLowEnergyDescriptor*>(ptr)->type();
}

void QLowEnergyDescriptor_DestroyQLowEnergyDescriptor(QtObjectPtr ptr){
	static_cast<QLowEnergyDescriptor*>(ptr)->~QLowEnergyDescriptor();
}

