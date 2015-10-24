#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QProgressBar_Alignment(QtObjectPtr ptr);
char* QProgressBar_Format(QtObjectPtr ptr);
int QProgressBar_InvertedAppearance(QtObjectPtr ptr);
int QProgressBar_IsTextVisible(QtObjectPtr ptr);
int QProgressBar_Maximum(QtObjectPtr ptr);
int QProgressBar_Minimum(QtObjectPtr ptr);
int QProgressBar_Orientation(QtObjectPtr ptr);
void QProgressBar_ResetFormat(QtObjectPtr ptr);
void QProgressBar_SetAlignment(QtObjectPtr ptr, int alignment);
void QProgressBar_SetFormat(QtObjectPtr ptr, char* format);
void QProgressBar_SetInvertedAppearance(QtObjectPtr ptr, int invert);
void QProgressBar_SetMaximum(QtObjectPtr ptr, int maximum);
void QProgressBar_SetMinimum(QtObjectPtr ptr, int minimum);
void QProgressBar_SetOrientation(QtObjectPtr ptr, int v);
void QProgressBar_SetTextDirection(QtObjectPtr ptr, int textDirection);
void QProgressBar_SetTextVisible(QtObjectPtr ptr, int visible);
void QProgressBar_SetValue(QtObjectPtr ptr, int value);
char* QProgressBar_Text(QtObjectPtr ptr);
int QProgressBar_TextDirection(QtObjectPtr ptr);
int QProgressBar_Value(QtObjectPtr ptr);
QtObjectPtr QProgressBar_NewQProgressBar(QtObjectPtr parent);
void QProgressBar_Reset(QtObjectPtr ptr);
void QProgressBar_SetRange(QtObjectPtr ptr, int minimum, int maximum);
void QProgressBar_ConnectValueChanged(QtObjectPtr ptr);
void QProgressBar_DisconnectValueChanged(QtObjectPtr ptr);
void QProgressBar_DestroyQProgressBar(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif