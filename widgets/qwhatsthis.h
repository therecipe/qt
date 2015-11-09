#ifdef __cplusplus
extern "C" {
#endif

void* QWhatsThis_QWhatsThis_CreateAction(void* parent);
void QWhatsThis_QWhatsThis_EnterWhatsThisMode();
void QWhatsThis_QWhatsThis_HideText();
int QWhatsThis_QWhatsThis_InWhatsThisMode();
void QWhatsThis_QWhatsThis_LeaveWhatsThisMode();
void QWhatsThis_QWhatsThis_ShowText(void* pos, char* text, void* w);

#ifdef __cplusplus
}
#endif