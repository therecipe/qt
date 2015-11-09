#ifdef __cplusplus
extern "C" {
#endif

void* QConicalGradient_NewQConicalGradient();
void* QConicalGradient_NewQConicalGradient2(void* center, double angle);
void* QConicalGradient_NewQConicalGradient3(double cx, double cy, double angle);
double QConicalGradient_Angle(void* ptr);
void QConicalGradient_SetAngle(void* ptr, double angle);
void QConicalGradient_SetCenter(void* ptr, void* center);
void QConicalGradient_SetCenter2(void* ptr, double x, double y);

#ifdef __cplusplus
}
#endif