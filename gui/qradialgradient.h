#ifdef __cplusplus
extern "C" {
#endif

void* QRadialGradient_NewQRadialGradient();
void* QRadialGradient_NewQRadialGradient6(void* center, double centerRadius, void* focalPoint, double focalRadius);
void* QRadialGradient_NewQRadialGradient4(void* center, double radius);
void* QRadialGradient_NewQRadialGradient2(void* center, double radius, void* focalPoint);
void* QRadialGradient_NewQRadialGradient7(double cx, double cy, double centerRadius, double fx, double fy, double focalRadius);
void* QRadialGradient_NewQRadialGradient5(double cx, double cy, double radius);
void* QRadialGradient_NewQRadialGradient3(double cx, double cy, double radius, double fx, double fy);
double QRadialGradient_CenterRadius(void* ptr);
double QRadialGradient_FocalRadius(void* ptr);
double QRadialGradient_Radius(void* ptr);
void QRadialGradient_SetCenter(void* ptr, void* center);
void QRadialGradient_SetCenter2(void* ptr, double x, double y);
void QRadialGradient_SetCenterRadius(void* ptr, double radius);
void QRadialGradient_SetFocalPoint(void* ptr, void* focalPoint);
void QRadialGradient_SetFocalPoint2(void* ptr, double x, double y);
void QRadialGradient_SetFocalRadius(void* ptr, double radius);
void QRadialGradient_SetRadius(void* ptr, double radius);

#ifdef __cplusplus
}
#endif