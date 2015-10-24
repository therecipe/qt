#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QGeoManeuver_NewQGeoManeuver();
QtObjectPtr QGeoManeuver_NewQGeoManeuver2(QtObjectPtr other);
int QGeoManeuver_Direction(QtObjectPtr ptr);
char* QGeoManeuver_InstructionText(QtObjectPtr ptr);
int QGeoManeuver_IsValid(QtObjectPtr ptr);
void QGeoManeuver_SetDirection(QtObjectPtr ptr, int direction);
void QGeoManeuver_SetInstructionText(QtObjectPtr ptr, char* instructionText);
void QGeoManeuver_SetPosition(QtObjectPtr ptr, QtObjectPtr position);
void QGeoManeuver_SetTimeToNextInstruction(QtObjectPtr ptr, int secs);
void QGeoManeuver_SetWaypoint(QtObjectPtr ptr, QtObjectPtr coordinate);
int QGeoManeuver_TimeToNextInstruction(QtObjectPtr ptr);
void QGeoManeuver_DestroyQGeoManeuver(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif