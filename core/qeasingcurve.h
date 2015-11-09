#ifdef __cplusplus
extern "C" {
#endif

void* QEasingCurve_NewQEasingCurve3(void* other);
void* QEasingCurve_NewQEasingCurve(int ty);
void* QEasingCurve_NewQEasingCurve2(void* other);
void QEasingCurve_AddCubicBezierSegment(void* ptr, void* c1, void* c2, void* endPoint);
void QEasingCurve_AddTCBSegment(void* ptr, void* nextPoint, double t, double c, double b);
double QEasingCurve_Amplitude(void* ptr);
double QEasingCurve_Overshoot(void* ptr);
double QEasingCurve_Period(void* ptr);
void QEasingCurve_SetAmplitude(void* ptr, double amplitude);
void QEasingCurve_SetOvershoot(void* ptr, double overshoot);
void QEasingCurve_SetPeriod(void* ptr, double period);
void QEasingCurve_SetType(void* ptr, int ty);
void QEasingCurve_Swap(void* ptr, void* other);
int QEasingCurve_Type(void* ptr);
double QEasingCurve_ValueForProgress(void* ptr, double progress);
void QEasingCurve_DestroyQEasingCurve(void* ptr);

#ifdef __cplusplus
}
#endif