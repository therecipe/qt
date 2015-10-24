#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QShortcut_AutoRepeat(QtObjectPtr ptr);
int QShortcut_Context(QtObjectPtr ptr);
int QShortcut_IsEnabled(QtObjectPtr ptr);
void QShortcut_SetAutoRepeat(QtObjectPtr ptr, int on);
void QShortcut_SetContext(QtObjectPtr ptr, int context);
void QShortcut_SetEnabled(QtObjectPtr ptr, int enable);
void QShortcut_SetKey(QtObjectPtr ptr, QtObjectPtr key);
void QShortcut_SetWhatsThis(QtObjectPtr ptr, char* text);
char* QShortcut_WhatsThis(QtObjectPtr ptr);
QtObjectPtr QShortcut_NewQShortcut(QtObjectPtr parent);
QtObjectPtr QShortcut_NewQShortcut2(QtObjectPtr key, QtObjectPtr parent, char* member, char* ambiguousMember, int context);
void QShortcut_ConnectActivated(QtObjectPtr ptr);
void QShortcut_DisconnectActivated(QtObjectPtr ptr);
void QShortcut_ConnectActivatedAmbiguously(QtObjectPtr ptr);
void QShortcut_DisconnectActivatedAmbiguously(QtObjectPtr ptr);
int QShortcut_Id(QtObjectPtr ptr);
QtObjectPtr QShortcut_ParentWidget(QtObjectPtr ptr);
void QShortcut_DestroyQShortcut(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif