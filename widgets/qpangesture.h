#ifdef __cplusplus
extern "C" {
#endif

double QPanGesture_Acceleration(void* ptr);
void QPanGesture_SetAcceleration(void* ptr, double value);
void QPanGesture_SetLastOffset(void* ptr, void* value);
void QPanGesture_SetOffset(void* ptr, void* value);
void QPanGesture_DestroyQPanGesture(void* ptr);

#ifdef __cplusplus
}
#endif