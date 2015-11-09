#ifdef __cplusplus
extern "C" {
#endif

int QShortcut_AutoRepeat(void* ptr);
int QShortcut_Context(void* ptr);
int QShortcut_IsEnabled(void* ptr);
void QShortcut_SetAutoRepeat(void* ptr, int on);
void QShortcut_SetContext(void* ptr, int context);
void QShortcut_SetEnabled(void* ptr, int enable);
void QShortcut_SetKey(void* ptr, void* key);
void QShortcut_SetWhatsThis(void* ptr, char* text);
char* QShortcut_WhatsThis(void* ptr);
void* QShortcut_NewQShortcut(void* parent);
void* QShortcut_NewQShortcut2(void* key, void* parent, char* member, char* ambiguousMember, int context);
void QShortcut_ConnectActivated(void* ptr);
void QShortcut_DisconnectActivated(void* ptr);
void QShortcut_ConnectActivatedAmbiguously(void* ptr);
void QShortcut_DisconnectActivatedAmbiguously(void* ptr);
int QShortcut_Id(void* ptr);
void* QShortcut_ParentWidget(void* ptr);
void QShortcut_DestroyQShortcut(void* ptr);

#ifdef __cplusplus
}
#endif