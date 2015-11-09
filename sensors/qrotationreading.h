#ifdef __cplusplus
extern "C" {
#endif

double QRotationReading_X(void* ptr);
double QRotationReading_Y(void* ptr);
double QRotationReading_Z(void* ptr);
void QRotationReading_SetFromEuler(void* ptr, double x, double y, double z);

#ifdef __cplusplus
}
#endif