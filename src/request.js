/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   request.js                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: elebouch <elebouch@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/21 14:57:57 by elebouch          #+#    #+#             */
/*   Updated: 2018/08/21 10:47:39 by elebouch         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

const pThrottle = require('p-throttle')
let rp = require('request-promise')
const ClientOAuth2 = require('client-oauth2')

const request = async options => rp(options)

const fetching = pThrottle(request, 2, 1000)

const rq = async options => fetching(options)

const forty2auth = new ClientOAuth2({
  clientId: process.env.INTRA_CLIENT_ID,
  clientSecret: process.env.INTRA_SECRET,
  accessTokenUri: 'https://api.intra.42.fr/oauth/token'
})

const request42 = async url => {
  let uri = 'https://api.intra.42.fr' + url
  const token = await forty2auth.credentials.getToken()
  let options = {
    uri: uri,
    qs: {
      access_token: token.data.access_token
    },
    headers: {
      'User-Agent': 'Request-Promise'
    },
    json: true // Automatically parses the JSON string in the response
  }
  try {
    return await rq(options)
  } catch (err) {
    return null
  }
}

module.exports = { request42 }
