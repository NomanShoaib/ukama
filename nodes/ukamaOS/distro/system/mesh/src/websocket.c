/**
 * Copyright (c) 2021-present, Ukama Inc.
 * All rights reserved.
 *
 * This source code is licensed under the XXX-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 */

#include <jansson.h>
#include <ulfius.h>
#include <string.h>

#include "mesh.h"
#include "log.h"
#include "work.h"
#include "jserdes.h"
#include "data.h"
#include "map.h"

extern WorkList *Transmit;
extern MapTable *IDTable;

/*
 * clear_response -- free up memory from MResponse.
 *
 */
static void clear_response(MResponse **resp) {

  if (*resp==NULL) return;

  free((*resp)->reqType);
  free((*resp)->serviceInfo);
  if ((*resp)->data) {
    free((*resp)->data);
  }

  free(*resp);
}

/*
 * websocket related callback functions.
 */

void websocket_manager(const URequest *request, WSManager *manager,
		       void *data) {

  WorkList *list;
  WorkItem *work;
  WorkList **transmit = &Transmit;

  if (*transmit == NULL)
    return;

  list = *transmit;

  while (TRUE) {

    pthread_mutex_lock(&(list->mutex));

    if (list->exit) { /* Likely we are closing the socket. */
      break;
    }

    if (list->first == NULL) { /* Empty. Wait. */
      log_debug("Waiting for work to be available ...");
      pthread_cond_wait(&(list->hasWork), &(list->mutex)); /* unlock mutex. */
    }

    /* We have some packet to transmit. */
    work = get_work_to_transmit(list);

    /* Unlock. */
    pthread_mutex_unlock(&(list->mutex));

    if (work == NULL) {
      continue;
    }

    /* We have valid work to do. yaay. */

    /* 1. Any pre-processing. */
    if (work->preFunc) {
      work->preFunc(work->data, work->preArgs);
    }

    /* 2. Send data over the wire. */
    /* Currently, Packet is JSON string. Send it over. */
    if (ulfius_websocket_wait_close(manager, 2000) == U_WEBSOCKET_STATUS_OPEN) {
      if (ulfius_websocket_send_json_message(manager, work->data) != U_OK) {
	log_error("Error sending JSON message.");
      }
    }

    /* 3. Any post-processing. */
    if (work->postFunc) {
      work->postFunc(work->data, work->postArgs);
    }

    /* Free up the memory */
    destroy_work_item(work);
  }

  return;
}

/*
 * websocket_incoming_message -- handle incoming message over websocket.
 *
 */

void websocket_incoming_message(const URequest *request,
				WSManager *manager, WSMessage *message,
				void *data) {
  MRequest *rcvdData=NULL;
  MResponse *rcvdResp=NULL;
  MapItem *item=NULL;
  json_t *json;
  char *str;
  char idStr[36+1];
  int ret;
  Config *config = (Config *)data;

  log_debug("Packet recevied. Data: %s", message->data);

  /* Server incoming packet handling. */
  if (config->mode == MODE_SERVER) {

    /* If we recevied a packet and our proxy is disable log error and reject*/
    if (config->proxy == FALSE) {
      log_error("Recevied packet while reverse-proxy is disabled. Ignoring");
      goto done;
    }

    /* Ignore the rest, for now. */
    if (message->opcode == U_WEBSOCKET_OPCODE_TEXT) {
      /* Convert to JSON and deserialize it. */
      json = json_loads(message->data, JSON_DECODE_ANY, NULL);

      if (json==NULL) {
	log_error("Error loading recevied data into JSON format. Str: %s",
		  message->data);
	goto done;
      }

      /* Convert JSON into request. */
      ret = deserialize_forward_request(&rcvdData, json);
      if (ret==FALSE)
	goto done;

      handle_recevied_data(rcvdData, config);

      /* Free up the memory from deser. */
      clear_request(&rcvdData);
    }

    goto done;
  }

  /* Client incoming packet handling. */
  if (config->mode == MODE_CLIENT) {

    /* Steps are:
     * 1. deserialize the response.
     * 2. lookup ID table for matching client, if any.
     * 3. Copy the data to the matching thread.
     * 4. Trigger conditional variable to enable processing.
     */

    json = json_loads(message->data, JSON_DECODE_ANY, NULL);
    if (json==NULL) {
      log_error("Error loading recevied data into JSON format. Str: %s",
		message->data);
      goto done;
    }

    ret = deserialize_response(&rcvdResp, json);
    if (ret==FALSE) {
      if (rcvdResp != NULL) free(rcvdResp);
      goto done;
    }

    item = lookup_item(IDTable, rcvdResp->serviceInfo->uuid);
    if (item == NULL) { /* No macthing service found in table. Ignore it */
      uuid_unparse(rcvdResp->serviceInfo->uuid, &idStr[0]);
      log_debug("No matching entry found in the table for UUID: %s Ignoring",
		&idStr[0]);
      goto done;
    }

    /* Copy recevied data into item. */
    pthread_mutex_lock(&item->mutex);

    item->size = rcvdResp->size;
    item->data = (void *)calloc(1, item->size);
    if (item->data == NULL) goto done;

    memcpy(item->data, rcvdResp->data, item->size);

    /* set the conditional variable. */
    pthread_cond_broadcast(&(item->hasResp));
    pthread_mutex_unlock(&item->mutex);
  }

 done:
  if (json) json_decref(json);
  clear_response(&rcvdResp);
  return;
}

/*
 * websocket_onclose -- is called when the websocket is closed.
 *
 */

void  websocket_onclose(const URequest *request, WSManager *manager,
			void *data) {

  char idStr[36+1];
  Config *config = (Config *)data;

  if (config == NULL)
    return;

  if (config->mode == MODE_SERVER) {
    if (config->deviceInfo) {
      uuid_unparse(config->deviceInfo->uuid, &idStr[0]);
      if (publish_amqp_event(config->conn, config->amqpExchange, CONN_CLOSE,
			     config->mode, config->deviceInfo->uuid) == FALSE) {
	log_error("Error publish device close msg on AMQP exchange: %s",
		  &idStr[0]);
      } else {
	log_debug("AMQP device close msg successfull for UUID: %s", &idStr[0]);
      }
    }
  }

  return;
}
