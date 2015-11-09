#ifdef __cplusplus
extern "C" {
#endif

int QDial_NotchSize(void* ptr);
double QDial_NotchTarget(void* ptr);
int QDial_NotchesVisible(void* ptr);
void QDial_SetNotchesVisible(void* ptr, int visible);
void QDial_SetWrapping(void* ptr, int on);
int QDial_Wrapping(void* ptr);
void* QDial_NewQDial(void* parent);
void QDial_DestroyQDial(void* ptr);

#ifdef __cplusplus
}
#endif