#ifdef __cplusplus
extern "C" {
#endif

int QProgressBar_Alignment(void* ptr);
char* QProgressBar_Format(void* ptr);
int QProgressBar_InvertedAppearance(void* ptr);
int QProgressBar_IsTextVisible(void* ptr);
int QProgressBar_Maximum(void* ptr);
int QProgressBar_Minimum(void* ptr);
int QProgressBar_Orientation(void* ptr);
void QProgressBar_ResetFormat(void* ptr);
void QProgressBar_SetAlignment(void* ptr, int alignment);
void QProgressBar_SetFormat(void* ptr, char* format);
void QProgressBar_SetInvertedAppearance(void* ptr, int invert);
void QProgressBar_SetMaximum(void* ptr, int maximum);
void QProgressBar_SetMinimum(void* ptr, int minimum);
void QProgressBar_SetOrientation(void* ptr, int v);
void QProgressBar_SetTextDirection(void* ptr, int textDirection);
void QProgressBar_SetTextVisible(void* ptr, int visible);
void QProgressBar_SetValue(void* ptr, int value);
char* QProgressBar_Text(void* ptr);
int QProgressBar_TextDirection(void* ptr);
int QProgressBar_Value(void* ptr);
void* QProgressBar_NewQProgressBar(void* parent);
void QProgressBar_Reset(void* ptr);
void QProgressBar_SetRange(void* ptr, int minimum, int maximum);
void QProgressBar_ConnectValueChanged(void* ptr);
void QProgressBar_DisconnectValueChanged(void* ptr);
void QProgressBar_DestroyQProgressBar(void* ptr);

#ifdef __cplusplus
}
#endif