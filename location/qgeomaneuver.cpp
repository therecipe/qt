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

void* QGeoManeuver_NewQGeoManeuver(){
	return new QGeoManeuver();
}

void* QGeoManeuver_NewQGeoManeuver2(void* other){
	return new QGeoManeuver(*static_cast<QGeoManeuver*>(other));
}

int QGeoManeuver_Direction(void* ptr){
	return static_cast<QGeoManeuver*>(ptr)->direction();
}

double QGeoManeuver_DistanceToNextInstruction(void* ptr){
	return static_cast<double>(static_cast<QGeoManeuver*>(ptr)->distanceToNextInstruction());
}

char* QGeoManeuver_InstructionText(void* ptr){
	return static_cast<QGeoManeuver*>(ptr)->instructionText().toUtf8().data();
}

int QGeoManeuver_IsValid(void* ptr){
	return static_cast<QGeoManeuver*>(ptr)->isValid();
}

void QGeoManeuver_SetDirection(void* ptr, int direction){
	static_cast<QGeoManeuver*>(ptr)->setDirection(static_cast<QGeoManeuver::InstructionDirection>(direction));
}

void QGeoManeuver_SetDistanceToNextInstruction(void* ptr, double distance){
	static_cast<QGeoManeuver*>(ptr)->setDistanceToNextInstruction(static_cast<qreal>(distance));
}

void QGeoManeuver_SetInstructionText(void* ptr, char* instructionText){
	static_cast<QGeoManeuver*>(ptr)->setInstructionText(QString(instructionText));
}

void QGeoManeuver_SetPosition(void* ptr, void* position){
	static_cast<QGeoManeuver*>(ptr)->setPosition(*static_cast<QGeoCoordinate*>(position));
}

void QGeoManeuver_SetTimeToNextInstruction(void* ptr, int secs){
	static_cast<QGeoManeuver*>(ptr)->setTimeToNextInstruction(secs);
}

void QGeoManeuver_SetWaypoint(void* ptr, void* coordinate){
	static_cast<QGeoManeuver*>(ptr)->setWaypoint(*static_cast<QGeoCoordinate*>(coordinate));
}

int QGeoManeuver_TimeToNextInstruction(void* ptr){
	return static_cast<QGeoManeuver*>(ptr)->timeToNextInstruction();
}

void QGeoManeuver_DestroyQGeoManeuver(void* ptr){
	static_cast<QGeoManeuver*>(ptr)->~QGeoManeuver();
}

