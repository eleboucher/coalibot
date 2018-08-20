/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   utils.js                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: elebouch <elebouch@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/23 11:07:52 by erwanleb          #+#    #+#             */
/*   Updated: 2018/08/20 23:40:27 by elebouch         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

const fs = require('fs')
const { getUsername } = require('./slack_api')

const choose = choices => {
  let index = Math.floor(Math.random() * choices.length)
  return choices[index]
}

const handlestat = async user => {
  let json
  try {
    json = await fs.readFileSync('./roulette.json', 'utf-8')
    json = JSON.parse(json)
  } catch (err) {
    json = {}
  }
  let username = await getUsername(user)
  if ('user' in username && 'name' in username['user']) {
    username = username['user']['name']
  } else {
    username = ''
  }
  if (json[username]) {
    json[username] += 1
  } else {
    json[username] = 1
  }
  fs.writeFile(
    './roulette.json',
    JSON.stringify(json, null, 4),
    'utf8',
    err => {
      if (err) throw err
    }
  )
}

module.exports.choose = choose
module.exports.handlestat = handlestat
