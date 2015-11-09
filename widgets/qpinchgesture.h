#ifdef __cplusplus
extern "C" {
#endif

int QPinchGesture_ChangeFlags(void* ptr);
double QPinchGesture_LastRotationAngle(void* ptr);
double QPinchGesture_LastScaleFactor(void* ptr);
double QPinchGesture_RotationAngle(void* ptr);
double QPinchGesture_ScaleFactor(void* ptr);
void QPinchGesture_SetCenterPoint(void* ptr, void* value);
void QPinchGesture_SetChangeFlags(void* ptr, int value);
void QPinchGesture_SetLastCenterPoint(void* ptr, void* value);
void QPinchGesture_SetLastRotationAngle(void* ptr, double value);
void QPinchGesture_SetLastScaleFactor(void* ptr, double value);
void QPinchGesture_SetRotationAngle(void* ptr, double value);
void QPinchGesture_SetScaleFactor(void* ptr, double value);
void QPinchGesture_SetStartCenterPoint(void* ptr, void* value);
void QPinchGesture_SetTotalChangeFlags(void* ptr, int value);
void QPinchGesture_SetTotalRotationAngle(void* ptr, double value);
void QPinchGesture_SetTotalScaleFactor(void* ptr, double value);
int QPinchGesture_TotalChangeFlags(void* ptr);
double QPinchGesture_TotalRotationAngle(void* ptr);
double QPinchGesture_TotalScaleFactor(void* ptr);
void QPinchGesture_DestroyQPinchGesture(void* ptr);

#ifdef __cplusplus
}
#endif