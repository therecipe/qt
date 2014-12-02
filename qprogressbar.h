#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QProgressBar_New_QWidget(QtObjectPtr parent);
void QProgressBar_Destroy(QtObjectPtr ptr);
int QProgressBar_Alignment(QtObjectPtr ptr);
char* QProgressBar_Format(QtObjectPtr ptr);
int QProgressBar_InvertedAppearance(QtObjectPtr ptr);
int QProgressBar_IsTextVisible(QtObjectPtr ptr);
int QProgressBar_Maximum(QtObjectPtr ptr);
int QProgressBar_Minimum(QtObjectPtr ptr);
int QProgressBar_Orientation(QtObjectPtr ptr);
void QProgressBar_ResetFormat(QtObjectPtr ptr);
void QProgressBar_SetAlignment_AlignmentFlag(QtObjectPtr ptr, int alignment);
void QProgressBar_SetFormat_String(QtObjectPtr ptr, char* format);
void QProgressBar_SetInvertedAppearance_Bool(QtObjectPtr ptr, int invert);
void QProgressBar_SetTextVisible_Bool(QtObjectPtr ptr, int visible);
char* QProgressBar_Text(QtObjectPtr ptr);
int QProgressBar_Value(QtObjectPtr ptr);
//Public Slots
void QProgressBar_ConnectSlotReset(QtObjectPtr ptr);
void QProgressBar_DisconnectSlotReset(QtObjectPtr ptr);
void QProgressBar_Reset(QtObjectPtr ptr);
void QProgressBar_ConnectSlotSetMaximum(QtObjectPtr ptr);
void QProgressBar_DisconnectSlotSetMaximum(QtObjectPtr ptr);
void QProgressBar_SetMaximum_Int(QtObjectPtr ptr, int maximum);
void QProgressBar_ConnectSlotSetMinimum(QtObjectPtr ptr);
void QProgressBar_DisconnectSlotSetMinimum(QtObjectPtr ptr);
void QProgressBar_SetMinimum_Int(QtObjectPtr ptr, int minimum);
void QProgressBar_ConnectSlotSetOrientation(QtObjectPtr ptr);
void QProgressBar_DisconnectSlotSetOrientation(QtObjectPtr ptr);
void QProgressBar_SetOrientation_Orientation(QtObjectPtr ptr, int orientation);
void QProgressBar_ConnectSlotSetRange(QtObjectPtr ptr);
void QProgressBar_DisconnectSlotSetRange(QtObjectPtr ptr);
void QProgressBar_SetRange_Int_Int(QtObjectPtr ptr, int minimum, int maximum);
void QProgressBar_ConnectSlotSetValue(QtObjectPtr ptr);
void QProgressBar_DisconnectSlotSetValue(QtObjectPtr ptr);
void QProgressBar_SetValue_Int(QtObjectPtr ptr, int value);
//Signals
void QProgressBar_ConnectSignalValueChanged(QtObjectPtr ptr);
void QProgressBar_DisconnectSignalValueChanged(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
