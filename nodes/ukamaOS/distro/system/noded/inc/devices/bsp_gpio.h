/**
 * Copyright (c) 2021-present, Ukama Inc.
 * All rights reserved.
 *
 * This source code is licensed under the XXX-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 */

#ifndef BSP_GPIO_H_
#define BSP_GPIO_H_

#ifdef __cplusplus
extern "C" {
#endif

#include "device.h"

/* EDGE */
#define GPIO_EDGE_NONE				0x00
#define GPIO_EDGE_RISING			0x01
#define GPIO_EDGE_FALLING			0x02
#define GPIO_EDGE_BOTH				0x03

/* Direction */
#define GPIO_DIRECTION_INPUT		0x00
#define GPIO_DIRECTION_OUTPUT		0x01

/*Active Low */
#define GPIO_NORMAL					0x00
#define GPIO_INVERT					0x01

/**
 * @fn      int bsp_gpio_configure(void*, void*, void*)
 * @brief   passes configuration request to driver wrapper layer.
 *
 * @param   p_dev
 * @param   prop
 * @param   data
 * @return  On success, 0
 *          On failure, non zero value
 */
int bsp_gpio_configure(void* p_dev, void* prop, void* data );

/**
 * @fn      int bsp_gpio_disable(void*, void*, void*)
 * @brief   passes sensor disable request to driver wrapper layer
 *
 * @param   p_dev
 * @param   prop
 * @param   data
 * @return  On success, 0
 *          On failure, non zero value
 */
int bsp_gpio_disable(void* p_dev, void* prop, void* data);

/**
 * @fn      int bsp_gpio_enable(void*, void*, void*)
 * @brief   passes sensor enable request to driver wrapper layer
 *
 * @param   p_dev
 * @param   prop
 * @param   data
 * @return  On success, 0
 *          On failure, non zero value
 */
int bsp_gpio_enable(void* p_dev, void* prop, void* data);

/**
 * @fn      int bsp_gpio_init()
 * @brief   Read sensor properties from the property config parsed during
 *          startup and then initialization request to driver wrapper layer.
 *
 * @return  On success, 0
 *          On failure, non zero value
 */
int bsp_gpio_init ();

/**
 * @fn      int bsp_gpio_registration(Device*)
 * @brief   passes registration request for driver wrapper layer.
 *
 * @param   p_dev
 * @return  On success, 0
 *          On failure, non zero value
 */
int bsp_gpio_registration(Device* p_dev);

/**
 * @fn      int bsp_gpio_read_prop_count(Device*, uint16_t*)
 * @brief   Reads the number properties available for sensor.
 *
 * @param   p_dev
 * @param   count
 * @return  On success, 0
 *          On failure, non zero value
 */
int bsp_gpio_read_prop_count(Device* p_dev, uint16_t * count);

/**
 * @fn      int bsp_gpio_read_properties(Device*, void*)
 * @brief   Reads the properties available for sensor.
 *
 * @param   p_dev
 * @param   prop
 * @return  On success, 0
 *          On failure, non zero value
 */
int bsp_gpio_read_properties(Device* p_dev, void* prop);

/**
 * @fn      int bsp_gpio_read(void*, void*, void*)
 * @brief   passes read request to driver wrapper layer.
 *
 * @param   p_dev
 * @param   prop
 * @param   data
 * @return  On success, 0
 *          On failure, non zero value
 */
int bsp_gpio_read(void* p_dev, void* prop, void* data);

/**
 * @fn      int bsp_gpio_write(void*, void*, void*)
 * @brief   passes write request to driver wrapper layer.
 *
 * @param   p_dev
 * @param   prop
 * @param   data
 * @return  On success, 0
 *          On failure, non zero value
 */
int bsp_gpio_write(void* p_dev, void* prop, void* data);

#ifdef __cplusplus
}
#endif

#endif /* BSP_GPIO_H_ */