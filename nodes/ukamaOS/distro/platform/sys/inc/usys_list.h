/**
 * Copyright (c) 2020-present, Ukama.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 */

#ifndef UTILS_LIST_H_
#define UTILS_LIST_H_

#include <stdbool.h>
#include <stdint.h>

// a common function used to free malloc'd objects
typedef void (*FreeFxn)(void *);
typedef int (*CompareFxn)(void*, void*);
typedef void* (*CopyDataFxn)(void*);

typedef int (*listIterator)(void *);


typedef struct  __attribute__((__packed__)) listNode {
  void *data;
  struct listNode *next;
} ListNode;

typedef struct {
  uint16_t logicalLength;
  uint16_t elementSize;
  ListNode *head;
  ListNode *tail;
  FreeFxn free;
  CompareFxn cmp;
  CopyDataFxn datacpy;
} ListInfo;

int usys_list_size(ListInfo *list);
int usys_list_if_element_found(ListInfo *list, void *element);
int usys_list_remove(ListInfo *list, void *element);
int usys_list_update(ListInfo *list, void *element);
void usys_list_append(ListInfo *list, void *element);
void usys_list_next(ListInfo *list, ListNode **node);
void usys_list_copy(ListInfo *list, void *data);
void usys_list_for_each(ListInfo *list, listIterator iterator);
void usys_list_new(ListInfo *list, uint16_t elementSize, FreeFxn freeFn,
    CompareFxn cmpFn, CopyDataFxn copyFn);
void usys_list_destroy(ListInfo *list);
void usys_list_prepend(ListInfo *list, void *element);
void* usys_list_search(ListInfo *list, void *element);
void* usys_list_tail(ListInfo *list );
void* usys_list_head(ListInfo *list,  uint8_t removeFromList);

#endif /* UTILS_LIST_H_ */