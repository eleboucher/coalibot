/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   request.js                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: elebouch <elebouch@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/21 14:57:57 by elebouch          #+#    #+#             */
/*   Updated: 2018/05/22 13:25:01 by elebouch         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

const pThrottle = require('p-throttle')
let rp = require('request-promise')

const request = async options => await rp(options)

const fetching = pThrottle(request, 2, 1000)

const rq = async options => await fetching(options)

module.exports = { rq }
