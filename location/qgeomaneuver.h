#ifdef __cplusplus
extern "C" {
#endif

void* QGeoManeuver_NewQGeoManeuver();
void* QGeoManeuver_NewQGeoManeuver2(void* other);
int QGeoManeuver_Direction(void* ptr);
double QGeoManeuver_DistanceToNextInstruction(void* ptr);
char* QGeoManeuver_InstructionText(void* ptr);
int QGeoManeuver_IsValid(void* ptr);
void QGeoManeuver_SetDirection(void* ptr, int direction);
void QGeoManeuver_SetDistanceToNextInstruction(void* ptr, double distance);
void QGeoManeuver_SetInstructionText(void* ptr, char* instructionText);
void QGeoManeuver_SetPosition(void* ptr, void* position);
void QGeoManeuver_SetTimeToNextInstruction(void* ptr, int secs);
void QGeoManeuver_SetWaypoint(void* ptr, void* coordinate);
int QGeoManeuver_TimeToNextInstruction(void* ptr);
void QGeoManeuver_DestroyQGeoManeuver(void* ptr);

#ifdef __cplusplus
}
#endif