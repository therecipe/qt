#pragma once

#ifndef GO_QZXING_H
#define GO_QZXING_H

#ifdef __cplusplus
extern "C" {
#endif

void QZXing_registerQMLTypes();
void QZXing_registerQMLImageProvider(void* ptr);

#ifdef __cplusplus
}
#endif

#endif