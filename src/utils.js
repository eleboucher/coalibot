/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   utils.js                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: erwanleb <erwanleb@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/23 11:07:52 by erwanleb          #+#    #+#             */
/*   Updated: 2018/03/02 17:33:20 by elebouch         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

const fs = require('fs')
const { getUsername } = require('./slack_api')


const choose = choices => {
  var index = Math.floor(Math.random() * choices.length)
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
  if ('user' in username && 'profile' in username['user'] && username['user']['profile']['display_name']) {
    username = username['user']['profile']['display_name']
  } else {
    username = ''
  }
  if (json[username]){
    json[username] += 1
  }
  else {
    json[username] = 1
  }
  fs.writeFile('./roulette.json', JSON.stringify(json, null, 4), 'utf8', err => {
    if (err) throw err
  })
}

module.exports.choose = choose
module.exports.handlestat = handlestat
