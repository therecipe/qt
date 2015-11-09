#ifdef __cplusplus
extern "C" {
#endif

void* QLinearGradient_NewQLinearGradient3(double x1, double y1, double x2, double y2);
void* QLinearGradient_NewQLinearGradient();
void* QLinearGradient_NewQLinearGradient2(void* start, void* finalStop);
void QLinearGradient_SetFinalStop(void* ptr, void* stop);
void QLinearGradient_SetFinalStop2(void* ptr, double x, double y);
void QLinearGradient_SetStart(void* ptr, void* start);
void QLinearGradient_SetStart2(void* ptr, double x, double y);

#ifdef __cplusplus
}
#endif