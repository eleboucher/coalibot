/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   giphy.js                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: elebouch <elebouch@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/24 15:03:45 by elebouch          #+#    #+#             */
/*   Updated: 2018/08/21 10:49:02 by elebouch         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

let rp = require('request-promise')
const { postMessage } = require('./slack_api')

const randomgif = async (query, channel) => {
  // not my api key :3:
  try {
    const res = await rp(
      `http://api.giphy.com/v1/gifs/search?q=${query}&api_key=aH0B2QT0mjnLZ7xKvHTwURhIHcIiB4MR&limit=15`,
      { json: true }
    )
    const theOnlyOne =
      res.data[Math.floor(Math.random() * res.data.length)].url
    postMessage(theOnlyOne, channel)
  } catch (err) {
    postMessage('Aucun gif trouvé', channel)
  }
}

module.exports = { randomgif }
