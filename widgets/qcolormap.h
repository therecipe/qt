#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QColormap_NewQColormap(QtObjectPtr colormap);
int QColormap_Depth(QtObjectPtr ptr);
int QColormap_Mode(QtObjectPtr ptr);
int QColormap_Size(QtObjectPtr ptr);
void QColormap_DestroyQColormap(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif