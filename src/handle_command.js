/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   handle_command.js                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: elebouch <elebouch@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/03/02 14:30:21 by elebouch          #+#    #+#             */
/*   Updated: 2018/05/22 23:23:08 by elebouch         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

const { postMessage, postUserMessage, sendReaction, fileUpload, postOnThread, getUsername } = require('./slack_api')
const { score, alliance, intralogtime, logtime, profil, who, where, event } = require('./42_api')
const { citation } = require('./citation')
const { randomgif } = require('./giphy')
const { roll, addmusic, music, meteo, dobby, php, roulette, coin, roulettestat, roulettetop } = require('./miscs')
const fs = require('fs')
const { parrot } = require('./const')
const { choose } = require('./utils')
const { addCommand } = require('./database')

const reply = async (cmd, channel) => {
  try {
    const contents = await fs.readFileSync('./reply.json')
    const json = JSON.parse(contents)
    if (json[cmd]) {
      postMessage(json[cmd], channel)
      return true
    }
  } catch (err) {
    console.log(err)
  }
  return false
}

functions = {
  alliance: (message, channel, ts, user) => alliance(channel),
  score: (message, channel, ts, user) => score(ts, channel),
  help: (message, channel, ts, user) => fileUpload(fs.createReadStream('./featurespic.jpeg'), channel),
  glegendr: (message, channel, ts, user) => randomgif('how about no'.replace(' ', '+'), channel),
  mfranc: (message, channel, ts, user) => postMessage(choose(['>Doucement avec les bots', '>Puuuuuuuuuuuuu']), channel),
  score: (message, channel, ts, user) => score(channel, ts),
  prof: (message, channel, ts, user) => profil(message.toLowerCase(), channel, user),
  logtime: (message, channel, ts, user) => logtime(message, channel, ts),
  intralogtime: (message, channel, ts, user) => intralogtime(message, channel, ts),
  who: (message, channel, ts, user) => who(message.toLowerCase(), channel),
  event: (message, channel, ts, user) => event(message.toLowerCase(), channel),
  roll: (message, channel, ts, user) => roll(message, channel, ts),
  where: (message, channel, ts, user) => where(message.toLowerCase(), channel, user),
  addmusic: (message, channel, ts, user) => addmusic(message, user, channel),
  music: (message, channel, ts, user) => music(channel),
  meteo: (message, channel, ts, user) => meteo(message, channel),
  dobby: (message, channel, ts, user) => dobby(user, channel),
  php: (message, channel, ts, user) => php(message, channel),
  roulette: (message, channel, ts, user) => roulette(channel, user),
  coin: (message, channel, ts, user) => coin(channel, user),
  roulettestat: (message, channel, ts, user) => roulettestat(user, channel),
  roulettetop: (message, channel, ts, user) => roulettetop(channel),
  randomgif: (message, channel, ts, user) =>
    randomgif(
      message
        .split(' ')
        .slice(2)
        .join(),
      channel
    ),
  oss: (message, channel, ts, user) =>
    citation(
      channel,
      './oss.txt',
      'https://static-cdn.jtvnw.net/emoticons/v1/518312/3.0',
      'Hubert Bonisseur de La Bath'
    ),
  parrot: (message, channel, ts, user) =>
    postMessage(':' + parrot[Math.floor(Math.random() * Math.floor(parrot.length))] + ':', channel),
  kaamelott: (message, channel, ts, user) =>
    citation(channel, './kaamelott.txt', 'https://img15.hostingpics.net/pics/4833663350.jpg', 'Perceval')
}

const handleCommand = async (msg, channel, ts, user) => {
  const message = msg.replace(/\s+/g, ' ').trim()
  console.log({ user, message })
  let command
  let option = null
  let isCommand = false
  if (/(\b|^)rip(\b|$)/i.test(message)) sendReaction('rip', channel, ts)
  if (/(\b|^)jpp(\b|$)/i.test(message)) sendReaction('jpp', channel, ts)
  if (/(\b|^)(php|ruby|ror|mongo|mongodb)(\b|$)/i.test(message)) sendReaction('poop', channel, ts)

  if (['coalibot', 'bc', 'cb'].indexOf(message.toLowerCase().split(' ')[0]) > -1 && message.split(' ').length > 1) {
    command = message.split(' ')[1].toLowerCase()
    option = message
      .split(' ')
      .splice(2)
      .join(' ')
    const result = await reply(command, channel)
    if (result === false) {
      if (functions[message.split(' ')[1].toLowerCase()]) {
        functions[command](message, channel, ts, user)
        isCommand = true
      }
    } else {
      isCommand = true
    }
  } else if (message.indexOf('!') === 0) {
    command = message
      .replace('!', 'bc ')
      .split(' ')[1]
      .toLowerCase()
    const result = await reply(command, channel)
    if (result === false) {
      if (functions[command]) {
        functions[command](message.replace('!', 'bc '), channel, ts, user)
        option = message
          .replace('!', 'bc ')
          .split(' ')
          .splice(2)
          .join(' ')
        isCommand = true
      }
    } else {
      isCommand = true
    }
  }
  if (isCommand) addCommand(command, option, channel, ts, user)
}

module.exports.handleCommand = handleCommand
