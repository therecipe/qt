#include "qgeomaneuver.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGeoCoordinate>
#include <QGeoManeuver>
#include "_cgo_export.h"

class MyQGeoManeuver: public QGeoManeuver {
public:
};

QtObjectPtr QGeoManeuver_NewQGeoManeuver(){
	return new QGeoManeuver();
}

QtObjectPtr QGeoManeuver_NewQGeoManeuver2(QtObjectPtr other){
	return new QGeoManeuver(*static_cast<QGeoManeuver*>(other));
}

int QGeoManeuver_Direction(QtObjectPtr ptr){
	return static_cast<QGeoManeuver*>(ptr)->direction();
}

char* QGeoManeuver_InstructionText(QtObjectPtr ptr){
	return static_cast<QGeoManeuver*>(ptr)->instructionText().toUtf8().data();
}

int QGeoManeuver_IsValid(QtObjectPtr ptr){
	return static_cast<QGeoManeuver*>(ptr)->isValid();
}

void QGeoManeuver_SetDirection(QtObjectPtr ptr, int direction){
	static_cast<QGeoManeuver*>(ptr)->setDirection(static_cast<QGeoManeuver::InstructionDirection>(direction));
}

void QGeoManeuver_SetInstructionText(QtObjectPtr ptr, char* instructionText){
	static_cast<QGeoManeuver*>(ptr)->setInstructionText(QString(instructionText));
}

void QGeoManeuver_SetPosition(QtObjectPtr ptr, QtObjectPtr position){
	static_cast<QGeoManeuver*>(ptr)->setPosition(*static_cast<QGeoCoordinate*>(position));
}

void QGeoManeuver_SetTimeToNextInstruction(QtObjectPtr ptr, int secs){
	static_cast<QGeoManeuver*>(ptr)->setTimeToNextInstruction(secs);
}

void QGeoManeuver_SetWaypoint(QtObjectPtr ptr, QtObjectPtr coordinate){
	static_cast<QGeoManeuver*>(ptr)->setWaypoint(*static_cast<QGeoCoordinate*>(coordinate));
}

int QGeoManeuver_TimeToNextInstruction(QtObjectPtr ptr){
	return static_cast<QGeoManeuver*>(ptr)->timeToNextInstruction();
}

void QGeoManeuver_DestroyQGeoManeuver(QtObjectPtr ptr){
	static_cast<QGeoManeuver*>(ptr)->~QGeoManeuver();
}

