#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QWhatsThis_QWhatsThis_CreateAction(QtObjectPtr parent);
void QWhatsThis_QWhatsThis_EnterWhatsThisMode();
void QWhatsThis_QWhatsThis_HideText();
int QWhatsThis_QWhatsThis_InWhatsThisMode();
void QWhatsThis_QWhatsThis_LeaveWhatsThisMode();
void QWhatsThis_QWhatsThis_ShowText(QtObjectPtr pos, char* text, QtObjectPtr w);

#ifdef __cplusplus
}
#endif