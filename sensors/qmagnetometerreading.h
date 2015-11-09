#ifdef __cplusplus
extern "C" {
#endif

double QMagnetometerReading_CalibrationLevel(void* ptr);
double QMagnetometerReading_X(void* ptr);
double QMagnetometerReading_Y(void* ptr);
double QMagnetometerReading_Z(void* ptr);
void QMagnetometerReading_SetCalibrationLevel(void* ptr, double calibrationLevel);
void QMagnetometerReading_SetX(void* ptr, double x);
void QMagnetometerReading_SetY(void* ptr, double y);
void QMagnetometerReading_SetZ(void* ptr, double z);

#ifdef __cplusplus
}
#endif