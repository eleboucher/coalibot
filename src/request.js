/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   request.js                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: elebouch <elebouch@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/21 14:57:57 by elebouch          #+#    #+#             */
/*   Updated: 2018/03/02 17:33:19 by elebouch         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

const pThrottle = require('p-throttle')
var rp = require('request-promise')

const request = async options => await rp(options)

const fetching = pThrottle(request, 2, 1000)

const rq = async options => await fetching(options)

module.exports = { rq }
