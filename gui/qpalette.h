#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QPalette_NColorRoles_Type();
int QPalette_IsEqual(QtObjectPtr ptr, int cg1, int cg2);
void QPalette_SetBrush2(QtObjectPtr ptr, int group, int role, QtObjectPtr brush);
QtObjectPtr QPalette_NewQPalette();
QtObjectPtr QPalette_NewQPalette8(QtObjectPtr other);
QtObjectPtr QPalette_NewQPalette3(int button);
QtObjectPtr QPalette_NewQPalette5(QtObjectPtr windowText, QtObjectPtr button, QtObjectPtr light, QtObjectPtr dark, QtObjectPtr mid, QtObjectPtr text, QtObjectPtr bright_text, QtObjectPtr base, QtObjectPtr window);
QtObjectPtr QPalette_NewQPalette2(QtObjectPtr button);
QtObjectPtr QPalette_NewQPalette4(QtObjectPtr button, QtObjectPtr window);
QtObjectPtr QPalette_NewQPalette7(QtObjectPtr p);
int QPalette_CurrentColorGroup(QtObjectPtr ptr);
int QPalette_IsBrushSet(QtObjectPtr ptr, int cg, int cr);
int QPalette_IsCopyOf(QtObjectPtr ptr, QtObjectPtr p);
void QPalette_SetBrush(QtObjectPtr ptr, int role, QtObjectPtr brush);
void QPalette_SetColor(QtObjectPtr ptr, int group, int role, QtObjectPtr color);
void QPalette_SetColor2(QtObjectPtr ptr, int role, QtObjectPtr color);
void QPalette_SetColorGroup(QtObjectPtr ptr, int cg, QtObjectPtr windowText, QtObjectPtr button, QtObjectPtr light, QtObjectPtr dark, QtObjectPtr mid, QtObjectPtr text, QtObjectPtr bright_text, QtObjectPtr base, QtObjectPtr window);
void QPalette_SetCurrentColorGroup(QtObjectPtr ptr, int cg);
void QPalette_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QPalette_DestroyQPalette(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif