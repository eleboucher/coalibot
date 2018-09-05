/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   42_api.js                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: elebouch <elebouch@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/19 21:07:36 by elebouch          #+#    #+#             */
/*   Updated: 2018/09/05 15:55:08 by elebouch         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

const { postMessage, getUsername, postAttachments } = require('./slack_api')
const { getRangeIntralogtime, formatOutputDatetime } = require('./logtime')
const { request42 } = require('./request')
const moment = require('moment')
const sprintf = require('sprintf-js').sprintf

const alliance = async channel => {
  const json = await request42('/v2/blocs/1/coalitions')
  json.sort(function (a, b) {
    return a.score < b.score
  })
  let rang = 0
  while (json[rang]['id'] !== 2) rang += 1
  if (rang === 0) {
    postMessage(
      `Felicitations Nous sommes premiers avec ${rang +
      1} points d'avance. :the-alliance:`,
      channel
    )
  } else {
    const diff = json[0]['score'] - json[rang]['score']
    postMessage(
      `Nous sommes à la ${rang +
      1}eme place avec ${diff} points de retard. :the-alliance:`,
      channel
    )
  }
}

const score = async channel => {
  const json = await request42('/v2/blocs/1/coalitions')
  json.sort(function (a, b) {
    return a.score < b.score
  })
  let reply = ''
  for (let coa of json) {
    reply += `${coa.name} ${coa.score}\n`
  }
  let attachments = [
    {
      fallback: reply,
      color: json[0]['color'],
      author_link: 'https://profile.intra.42.fr/blocs/1/coalitions',
      fields: [
        {
          title: json[0]['name'],
          value: json[0]['score'],
          short: true
        },
        {
          title: json[1]['name'],
          value: String(
            json[1]['score'] +
            ' (' +
            Number(json[1]['score'] - json[0]['score']) +
            ')'
          ),
          short: true
        },
        {
          title: json[2]['name'],
          value: String(
            json[2]['score'] +
            ' (' +
            Number(json[2]['score'] - json[0]['score']) +
            ')'
          ),
          short: true
        },
        {
          title: json[3]['name'],
          value: String(
            json[3]['score'] +
            ' (' +
            Number(json[3]['score'] - json[0]['score']) +
            ')'
          ),
          short: true
        }
      ],
      footer: 'Powered by Coalibot'
    }
  ]
  postAttachments('', attachments, channel)
}

const profil = async (msg, channel, usr) => {
  let user
  if (msg.split(' ').length > 2) user = msg.split(' ')[2]
  else {
    let username = await getUsername(usr)
    try {
      user = username['user']['profile']['email'].substr(
        0,
        username['user']['profile']['email'].toString().indexOf('@')
      )
    } catch (err) {
      user = username['user']['name']
    }
  }
  let url = '/v2/users/' + user
  let urlcoal = url + '/coalitions/'
  const data = await request42(url)
  let lvl = 1
  const coaldata = await request42(urlcoal)
  if (!data) {
    postMessage('invalid login', channel)
    return
  }
  let lvlpiscine = 0
  if (
    !data['pool_year'] ||
    data['pool_year'] === '2013' ||
    data['pool_year'] === '2014'
  ) {
    lvlpiscine = 0
  } else if (data['cursus_users'].length === 1) {
    lvlpiscine = data['cursus_users'][0]['level']
    lvl = 0
  } else lvlpiscine = data['cursus_users'][1]['level']
  let coalslug = ''
  if (coaldata.length) coalslug = ':' + coaldata[0]['slug'] + ':'
  let rangeEnd = moment()
  let rangeBegin = moment().subtract(7, 'days')
  const logtime = await getRangeIntralogtime(user, rangeBegin, rangeEnd)
  const time = formatOutputDatetime(logtime)
  const stage = (data => {
    const ret = {
      finished: ':white_check_mark:',
      in_progress: ':clock1:'
    }
    const u = data.projects_users.find(d => d.project.id === 118)
    const uploaded = data.projects_users.find(d => d.project.id === 119)
    return u && uploaded && uploaded['status'] === 'finished' && u['final_mark'] > 0
      ? ret[u['status']]
      : ':negative_squared_cross_mark:'
  })(data)
  const attachments = [
    {
      title: `${data['displayname']} - ${user} ${coalslug}`,
      title_link: 'https://profile.intra.42.fr/users/' + user,
      color:
        coaldata && coaldata.length !== 0 && coaldata[0]['color']
          ? coaldata[0]['color']
          : '#D40000',
      thumb_url: data['image_url'],
      fields: [
        {
          title: 'Niveau',
          value: `${
            lvl === 0 ? 0 : sprintf('%.2f', data['cursus_users'][0]['level'])
            }`,
          short: true
        },
        {
          title: 'Niveau piscine',
          value: `${sprintf('%.2f', lvlpiscine)} ${data['pool_month']} ${
            data['pool_year']
            }`,
          short: true
        },
        {
          title: 'Temps de log cette semaine',
          value: `${sprintf('%02d:%02d', time[0], time[1])}`,
          short: true
        },
        { title: 'Stage', value: stage, short: true },
        {
          title: 'Location',
          value: `${data.location ? data.location : 'Hors ligne'}`,
          short: true
        }
      ]
    }
  ]
  postAttachments('', attachments, channel)
}

const who = async (msg, channel) => {
  let place = ''
  if (msg.split(' ').length > 2) place = msg.split(' ')[2]
  else {
    postMessage(`prend une place en parametre`, channel)
    return
  }
  if (!place || place.startsWith('!') || place.startsWith('?')) return
  const url = `/v2/campus/1/locations/?filter[host]=${place}`
  const data = await request42(url)
  if (data.length === 0) postMessage(`Place *${place}* vide`, channel)
  if (data[0].endAt === null) {
    postMessage(
      `*${data[0]['user']['login']}* est à la place *${place}*`,
      channel
    )
  } else if (data[0].endAt !== null) {
    const diff = moment(data[0].endAt).fromNow()
    postMessage(
      `Place *${place}* vide, ancien utilisateur: *${
      data[0]['user']['login']
      }* ${diff}`,
      channel
    )
  }
}

const where = async (msg, channel, usr) => {
  let user = ''
  if (
    msg.split(' ').length === 6 &&
    (msg.indexOf('le branle couille') !== -1 ||
      msg.indexOf('la branle couille') !== -1)
  ) {
    let dateBegin = moment().subtract(7, 'days')
    let dateEnd = moment().add(1, 'days')
    user = msg.split(' ')[5]
    const logtime = await getRangeIntralogtime(user, dateBegin, dateEnd)
    const time = formatOutputDatetime(logtime)
    if (time[0] >= 35) {
      postMessage(`*${user}* is not a branle couille`, channel)
      return
    }
    let url = `/v2/users/${user}/locations`
    const data = await request42(url)
    if (!data) {
      postMessage(`login invalide`, channel)
      return
    }
    if (data.length === 0 || data[0]['endAt']) {
      postMessage(`*${user}* est hors ligne`, channel)
    } else {
      postMessage(`*${user}* est à la place *${data[0]['host']}*`, channel)
    }
    return
  }
  if (msg.split(' ').length > 2) user = msg.split(' ')[2]
  else {
    let username = await getUsername(usr)
    try {
      user = username['user']['profile']['email']
        .toString()
        .substr(0, username['user']['profile']['email'].toString().indexOf('@'))
    } catch (err) {
      user = username['user']['name']
    }
  }
  if (!user || user.startsWith('!') || user.startsWith('?')) return
  if (user === 'queen' || user === 'way') {
    postMessage(
      "follow me bruddah\ni'll show you de way :uganda_knuckles:",
      channel
    )
    return
  }
  if (user === 'dieu' || user === 'dobby') user = 'elebouch'
  if (user === 'manager') user = 'vtennero'
  if (user === 'guardians' || user === 'gardiens') {
    let guardians = [
      'dcirlig',
      'vtennero',
      'elebouch',
      'fbabin',
      'tbailly-',
      'mmerabet',
      'aledru',
      'dlavaury'
    ]
    for (let guardian of guardians) {
      let url = `/v2/users/${guardian}/locations?filter[active]=true`
      const data = await request42(url)
      if (!data || data.length === 0 || data[0]['endAt']) {
        postMessage(`*${guardian}* est hors ligne`, channel)
      } else {
        postMessage(
          `*${guardian}* est à la place *${data[0]['host']}*`,
          channel
        )
      }
    }
    return
  }
  let url = `/v2/users/${user}/locations?filter[active]=true`
  const data = await request42(url)
  if (!data) {
    postMessage(`login invalide`, channel)
    return
  }
  if (data.length === 0 || data[0]['endAt']) {
    postMessage(`*${user}* est hors ligne`, channel)
  } else postMessage(`*${user}* est à la place *${data[0]['host']}*`, channel)
}

const event = async (msg, channel) => {
  let beginAt
  let endAt
  if (msg.split(' ').length === 2) {
    beginAt = moment().format('YYYY-MM-DD')
    endAt = moment()
      .add(1, 'days')
      .format('YYYY-MM-DD')
  } else if (msg.split(' ').length === 3) {
    try {
      beginAt = moment(msg.split(' ')[2]).format('YYYY-MM-DD')
    } catch (e) {
      beginAt = moment().format('YYYY-MM-DD')
    }
    endAt = moment(beginAt)
      .add(1, 'days')
      .format('YYYY-MM-DD')
  } else {
    return
  }
  let url = `/v2/campus/1/events?range[beginAt]=${beginAt},${endAt}`
  const data = await request42(url)
  data.sort(function (a, b) {
    return moment(a.beginAt) - moment(b.beginAt)
  })

  for (let event of data) {
    let attachments = [
      {
        fallback: event.name,
        title: event.name,
        title_link: `https://profile.intra.42.fr/events/${event.id}`,
        text: event.description.slice(0, 150) + ' ...',
        footer: `${event.nbr_subscribers}/${event.max_people} Participants`,
        ts: moment(event.beginAt).format('X'),
        color: '#01babc'
      }
    ]
    await postAttachments('', attachments, channel)
  }
  if (data.length === 0) {
    postMessage('Aucun events à cette date!', channel)
  }
}

module.exports = { alliance, score, profil, who, event, where }
