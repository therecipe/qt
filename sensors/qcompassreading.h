#ifdef __cplusplus
extern "C" {
#endif

double QCompassReading_Azimuth(void* ptr);
double QCompassReading_CalibrationLevel(void* ptr);
void QCompassReading_SetAzimuth(void* ptr, double azimuth);
void QCompassReading_SetCalibrationLevel(void* ptr, double calibrationLevel);

#ifdef __cplusplus
}
#endif