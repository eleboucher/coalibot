/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   42_api.js                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: elebouch <elebouch@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/19 21:07:36 by elebouch          #+#    #+#             */
/*   Updated: 2018/08/21 00:11:53 by elebouch         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

const { postMessage, getUsername, postAttachments } = require('./slack_api')
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
    const diff = json[rang]['score'] - json[1]['score']
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
  url = '/v2/users/' + user
  urlcoal = url + '/coalitions/'
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
  range_end = moment()
  range_begin = moment().subtract(7, 'days')
  const logtime = await get_range_intralogtime(user, range_begin, range_end)
  const time = format_output_datetime(logtime)
  graph = 'https://projects.intra.42.fr/projects/graph?login=' + user
  const stage = (data => {
    const ret = {
      finished: ':white_check_mark:',
      in_progress: ':clock1:'
    }
    const u = data.projects_users.find(d => d.project.id === 118)
    const uploaded = data.projects_users.find(d => d.project.id === 119)
    return u && uploaded && uploaded['status'] === 'finished'
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
  if (msg.split(' ').length > 2) place = msg.split(' ')[2]
  else {
    postMessage(`prend une place en parametre`, channel)
    return
  }
  if (!place || place.startsWith('!') || place.startsWith('?')) return
  const url = `/v2/campus/1/locations/?filter[host]=${place}`
  const data = await request42(url)
  if (data.length === 0) postMessage(`Place *${place}* vide`, channel)
  if (data[0].end_at === null) {
    postMessage(
      `*${data[0]['user']['login']}* est à la place *${place}*`,
      channel
    )
  } else if (data[0].end_at !== null) {
    const diff = moment(data[0].end_at).fromNow()
    postMessage(
      `Place *${place}* vide, ancien utilisateur: *${
        data[0]['user']['login']
      }* ${diff}`,
      channel
    )
  }
}

const where = async (msg, channel, usr) => {
  if (
    msg.split(' ').length === 6 &&
    (msg.indexOf('le branle couille') !== -1 ||
      msg.indexOf('la branle couille') !== -1)
  ) {
    let date_begin = moment().subtract(7, 'days')
    let date_end = moment().add(1, 'days')
    const user = msg.split(' ')[5]
    const logtime = await get_range_intralogtime(user, date_begin, date_end)
    const time = format_output_datetime(logtime)
    if (time[0] >= 35) {
      postMessage(`*${user}* is not a branle couille`, channel)
      return
    }
    url = `/v2/users/${user}/locations`
    const data = await request42(url)
    if (!data) {
      postMessage(`login invalide`, channel)
      return
    }
    if (data.length === 0 || data[0]['end_at']) {
      postMessage(`*${user}* est hors ligne`, channel)
    } else postMessage(`*${user}* est à la place *${data[0]['host']}*`, channel)
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
  if (user === 'queen' || user == 'way') {
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
      url = `/v2/users/${guardian}/locations`
      const data = await request42(url)
      if (!data || data.length === 0 || data[0]['end_at']) {
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
  url = `/v2/users/${user}/locations`
  const data = await request42(url)
  if (!data) {
    postMessage(`login invalide`, channel)
    return
  }
  if (data.length === 0 || data[0]['end_at']) {
    postMessage(`*${user}* est hors ligne`, channel)
  } else postMessage(`*${user}* est à la place *${data[0]['host']}*`, channel)
}

const event = async (msg, channel) => {
  let begin_at
  let end_at
  if (msg.split(' ').length === 2) {
    begin_at = moment().format('YYYY-MM-DD')
    end_at = moment()
      .add(1, 'days')
      .format('YYYY-MM-DD')
  } else if (msg.split(' ').length === 3) {
    try {
      begin_at = moment(msg.split(' ')[2]).format('YYYY-MM-DD')
    } catch (e) {
      begin_at = moment().format('YYYY-MM-DD')
    }
    end_at = moment(begin_at)
      .add(1, 'days')
      .format('YYYY-MM-DD')
  } else {
    return
  }
  url = `/v2/campus/1/events?range[begin_at]=${begin_at},${end_at}`
  const data = await request42(url)
  data.sort(function (a, b) {
    return moment(a.begin_at) - moment(b.begin_at)
  })

  for (let event of data) {
    attachments = [
      {
        fallback: event.name,
        title: event.name,
        title_link: `https://profile.intra.42.fr/events/${event.id}`,
        text: event.description.slice(0, 150) + ' ...',
        footer: `${event.nbr_subscribers}/${event.max_people} Participants`,
        ts: moment(event.begin_at).format('X'),
        color: '#01babc'
      }
    ]
    await postAttachments('', attachments, channel)
  }
  if (data.length == 0) {
    postMessage('Aucun events à cette date!', channel)
  }
}

module.exports.alliance = alliance
module.exports.logtime = logtime
module.exports.intralogtime = intralogtime
module.exports.score = score
module.exports.profil = profil
module.exports.who = who
module.exports.where = where
module.exports.event = event
