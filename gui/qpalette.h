#ifdef __cplusplus
extern "C" {
#endif

int QPalette_NColorRoles_Type();
void* QPalette_Brush(void* ptr, int group, int role);
int QPalette_IsEqual(void* ptr, int cg1, int cg2);
void QPalette_SetBrush2(void* ptr, int group, int role, void* brush);
void* QPalette_NewQPalette();
void* QPalette_NewQPalette8(void* other);
void* QPalette_NewQPalette3(int button);
void* QPalette_NewQPalette5(void* windowText, void* button, void* light, void* dark, void* mid, void* text, void* bright_text, void* base, void* window);
void* QPalette_NewQPalette2(void* button);
void* QPalette_NewQPalette4(void* button, void* window);
void* QPalette_NewQPalette7(void* p);
void* QPalette_AlternateBase(void* ptr);
void* QPalette_Base(void* ptr);
void* QPalette_BrightText(void* ptr);
void* QPalette_Brush2(void* ptr, int role);
void* QPalette_Button(void* ptr);
void* QPalette_ButtonText(void* ptr);
void* QPalette_Color(void* ptr, int group, int role);
void* QPalette_Color2(void* ptr, int role);
int QPalette_CurrentColorGroup(void* ptr);
void* QPalette_Dark(void* ptr);
void* QPalette_Highlight(void* ptr);
void* QPalette_HighlightedText(void* ptr);
int QPalette_IsBrushSet(void* ptr, int cg, int cr);
int QPalette_IsCopyOf(void* ptr, void* p);
void* QPalette_Light(void* ptr);
void* QPalette_Link(void* ptr);
void* QPalette_LinkVisited(void* ptr);
void* QPalette_Mid(void* ptr);
void* QPalette_Midlight(void* ptr);
void QPalette_SetBrush(void* ptr, int role, void* brush);
void QPalette_SetColor(void* ptr, int group, int role, void* color);
void QPalette_SetColor2(void* ptr, int role, void* color);
void QPalette_SetColorGroup(void* ptr, int cg, void* windowText, void* button, void* light, void* dark, void* mid, void* text, void* bright_text, void* base, void* window);
void QPalette_SetCurrentColorGroup(void* ptr, int cg);
void* QPalette_Shadow(void* ptr);
void QPalette_Swap(void* ptr, void* other);
void* QPalette_Text(void* ptr);
void* QPalette_ToolTipBase(void* ptr);
void* QPalette_ToolTipText(void* ptr);
void* QPalette_Window(void* ptr);
void* QPalette_WindowText(void* ptr);
void QPalette_DestroyQPalette(void* ptr);

#ifdef __cplusplus
}
#endif