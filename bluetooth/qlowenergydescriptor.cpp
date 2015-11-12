#include "qlowenergydescriptor.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QLowEnergyDescriptor>
#include "_cgo_export.h"

class MyQLowEnergyDescriptor: public QLowEnergyDescriptor {
public:
};

void* QLowEnergyDescriptor_NewQLowEnergyDescriptor(){
	return new QLowEnergyDescriptor();
}

void* QLowEnergyDescriptor_NewQLowEnergyDescriptor2(void* other){
	return new QLowEnergyDescriptor(*static_cast<QLowEnergyDescriptor*>(other));
}

int QLowEnergyDescriptor_IsValid(void* ptr){
	return static_cast<QLowEnergyDescriptor*>(ptr)->isValid();
}

char* QLowEnergyDescriptor_Name(void* ptr){
	return static_cast<QLowEnergyDescriptor*>(ptr)->name().toUtf8().data();
}

int QLowEnergyDescriptor_Type(void* ptr){
	return static_cast<QLowEnergyDescriptor*>(ptr)->type();
}

void* QLowEnergyDescriptor_Value(void* ptr){
	return new QByteArray(static_cast<QLowEnergyDescriptor*>(ptr)->value());
}

void QLowEnergyDescriptor_DestroyQLowEnergyDescriptor(void* ptr){
	static_cast<QLowEnergyDescriptor*>(ptr)->~QLowEnergyDescriptor();
}

