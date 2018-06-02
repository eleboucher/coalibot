/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   42_api.js                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: elebouch <elebouch@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/19 21:07:36 by elebouch          #+#    #+#             */
/*   Updated: 2018/06/02 22:34:30 by elebouch         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

const {
  postMessage,
  postUserMessage,
  sendReaction,
  fileUpload,
  postOnThread,
  getUsername,
  postAttachments
} = require('./slack_api')
const rq = require('./request').rq
const ClientOAuth2 = require('client-oauth2')
const parse = require('csv-parse/lib/sync')
var fs = require('fs')
const { month } = require('./const')
const moment = require('moment')
const sprintf = require('sprintf-js').sprintf

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

const alliance = async channel => {
  const json = await request42('/v2/coalitions')
  json.sort(function(a, b) {
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
  const json = await request42('/v2/coalitions')
  json.sort(function(a, b) {
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

const getHourByName = (name, data) => {
  let h = false
  const regex = new RegExp(
    '(\\b|^)' + name.last + '(\\d|), ' + name.first + '(\\d|)(\\b|$)',
    'i'
  )
  for (let o of data) {
    if (regex.test(o.name)) {
      h += parseInt(o.h)
    }
  }
  return h
}

const get_range_logtime = async (name, start, end) => {
  let current = moment(start)
  let logtime = moment.duration(0)
  while (!current.isAfter(end)) {
    try {
      const data = await fs.readFileSync(
        `./logtime/presence-${current.get('year')}-${current.format('MM')}.csv`,
        'utf-8'
      )
      let thisMonth = getHourByName(
        name,
        parse(data, {
          delimiter: ';',
          relax_column_count: true,
          rowDelimiter: ';\r',
          columns: ['name', 'd', 'h'],
          trim: true
        })
      )
      if (thisMonth !== false) logtime.add(parseInt(thisMonth), 'hours')
    } catch (e) {
      logtime.add(0, 'hours')
    }
    current.add(1, 'months')
  }
  return logtime.as('hours')
}

const logtime = async (message, channel, ts) => {
  if (message.split(' ').length < 3) {
    postOnThread(
      'Usage: cb logtime login [annee | mois | trimestre[1-4] [annee] | semestre[1-2] [annee]]',
      channel,
      ts
    )
    return
  }
  const intradata = await request42('/v2/users/' + message.split(' ')[2])
  if (intradata && intradata.last_name && intradata.first_name)
    name = {
      last: intradata.last_name
        .trim()
        .normalize('NFD')
        .replace(/[\u0300-\u036f]/g, ''),
      first: intradata.first_name
        .trim()
        .normalize('NFD')
        .replace(/[\u0300-\u036f]/g, '')
    }
  else {
    postOnThread('login incorrect', channel, ts)
    return
  }
  if (message.split(' ').length === 3) {
    let date_begin = moment({ y: new Date().getFullYear(), M: 0, d: 1 })
    let date_end = moment({ y: new Date().getFullYear(), M: 11, d: 1 })
    const logtime = await get_range_logtime(name, date_begin, date_end)
    postOnThread(logtime + 'h', channel, ts)
  } else if (
    message.split(' ').length === 4 &&
    !isNaN(message.split(' ')[3]) &&
    parseInt(message.split(' ')[3]) > 2016
  ) {
    let date_begin = moment({ y: parseInt(message.split(' ')[3]), M: 0, d: 1 })
    let date_end = moment({ y: parseInt(message.split(' ')[3]), M: 11, d: 1 })
    const logtime = await get_range_logtime(name, date_begin, date_end)
    postOnThread(logtime + 'h', channel, ts)
  } else if (
    /(\b|^)trimestre[1-4](\b|$)/i.test(message.split(' ')[3]) &&
    (message.split(' ').length === 4 ||
      (message.split(' ').length === 5 &&
        parseInt(message.split(' ')[4]) > 2016))
  ) {
    let quarter = parseInt(message.split(' ')[3].replace('trimestre', '')) - 1
    const year =
      message.split(' ').length === 5 && parseInt(message.split(' ')[4]) > 2016
        ? parseInt(message.split(' ')[4])
        : new Date().getFullYear()
    let date_begin = moment(new Date(year, quarter * 3, 1))
    let date_end = moment(new Date(year, date_begin.get('month') + 3, 0))
    const logtime = await get_range_logtime(name, date_begin, date_end)
    postOnThread(logtime + 'h', channel, ts)
  } else if (
    /(\b|^)semestre[1-2](\b|$)/i.test(message.split(' ')[3]) &&
    (message.split(' ').length === 4 ||
      (message.split(' ').length === 5 &&
        parseInt(message.split(' ')[4]) > 2016))
  ) {
    let semestre = parseInt(message.split(' ')[3].replace('semestre', '')) - 1
    const year =
      message.split(' ').length === 5 && parseInt(message.split(' ')[4]) > 2016
        ? parseInt(message.split(' ')[4])
        : new Date().getFullYear()
    let date_begin = moment(new Date(year, semestre * 6, 1))
    let date_end = moment(new Date(year, date_begin.get('month') + 6, 0))
    const logtime = await get_range_logtime(name, date_begin, date_end)
    postOnThread(logtime + 'h', channel, ts)
  } else if (
    message
      .split(' ')[3]
      .normalize('NFD')
      .replace(/[\u0300-\u036f]/g, '') in month &&
    (message.split(' ').length === 4 ||
      (message.split(' ').length === 5 &&
        parseInt(message.split(' ')[4]) > 2016))
  ) {
    const year =
      message.split(' ').length === 5 && parseInt(message.split(' ')[4]) > 2012
        ? parseInt(message.split(' ')[4])
        : new Date().getFullYear()
    let date_begin = moment({ y: year, M: month[message.split(' ')[3]], d: 1 })
    let date_end = moment({ y: year, M: month[message.split(' ')[3]], d: 1 })
    const logtime = await get_range_logtime(name, date_begin, date_end)
    postOnThread(logtime + 'h', channel, ts)
  } else if (
    message.split(' ').length === 5 &&
    moment(message.split(' ')[3], moment.ISO_8601, true).isValid() &&
    (message.split(' ')[4] === 'today' ||
      moment(message.split(' ')[4], moment.ISO_8601, true).isValid())
  ) {
    let date_end =
      message.split(' ')[4] === 'today'
        ? moment()
        : moment(message.split(' ')[4])
    let date_begin = moment(message.split(' ')[3])
    if (date_end.isValid() && date_begin.isValid()) {
      const logtime = await get_range_logtime(name, date_begin, date_end)
      postOnThread(logtime + 'h', channel, ts)
    }
  } else
    postOnThread(
      'Usage: cb logtime login [annee | mois | trimestre[1-4] [annee] | semestre[1-2] [annee]]',
      channel,
      ts
    )
}

const get_range_intralogtime = async (user, range_begin, range_end) => {
  range_begin = moment(range_begin).format('YYYY-MM-DD')
  range_end = moment(range_end).format('YYYY-MM-DD')
  range_date = `?page[size]=100&range[begin_at]=${range_begin},${range_end}`
  url = `/v2/users/${user}/locations/${range_date}`
  const data = await request42(url)
  if (range_begin === range_end) {
    return moment.duration(0)
  }
  try {
    async function get_more(data) {
      let tmp
      let i = 2
      let ret = data
      do {
        last_location = moment(ret[ret.length - 1]['begin_at'])
        if (moment(range_begin).isBefore(last_location)) {
          tmp = await request42(url + '&page[number]=' + i)
          if (tmp) {
            ret = ret.concat(tmp)
          }
          i += 1
        } else {
          return ret
        }
      } while (tmp && tmp.length)
      return ret
    }
    let locations = await get_more(data)
    let logtime = moment.duration(0)
    for (let x of locations) {
      if (x['end_at']) log_end = moment(x['end_at'])
      else log_end = moment()
      log_start = moment(x['begin_at'])
      log_session = log_end - log_start
      logtime.add(log_session)
    }
    return logtime
  } catch (e) {
    return moment.duration(0)
  }
}

const format_output_datetime = time => {
  const timem = Number(time.as('minutes'))
  const hours = Math.floor(timem / 60)
  const min = Math.floor(timem % 60)
  return [hours, min]
}

const profil = async (msg, channel, usr) => {
  let user
  if (msg.split(' ').length > 2) user = msg.split(' ')[2]
  else {
    let username = await getUsername(usr)
    try {
      user = username['user']['email'].strsub(
        0,
        username['user']['email'].indexOf('@')
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
      color: coaldata[0]['color'],
      thumb_url: data['image_url'],
      fields: [
        {
          title: 'Niveau',
          value: `${lvl === 0 ? 0 : data['cursus_users'][0]['level']}`,
          short: true
        },
        {
          title: 'Niveau piscine',
          value: `${lvlpiscine} ${data['pool_month']} ${data['pool_year']}`,
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

const intralogtime = async (message, channel, ts) => {
  if (message.split(' ').length < 3) {
    postOnThread(
      'Usage: cb intralogtime login [datedebut datefin | annee | mois | trimestre[1-4][annee] | semestre[1-2][annee]](date au format "Y-M-D")',
      channel,
      ts
    )
    return
  }
  if (message.split(' ').length === 3) {
    let date_begin = moment().subtract(7, 'days')
    let date_end = moment().add(1, 'days')
    const logtime = await get_range_intralogtime(
      message.split(' ')[2],
      date_begin,
      date_end
    )
    const time = format_output_datetime(logtime)
    postOnThread(sprintf(`%02dh%02d`, time[0], time[1]), channel, ts)
    return
  } else if (
    message.split(' ').length === 4 &&
    !isNaN(message.split(' ')[3]) &&
    parseInt(message.split(' ')[3]) > 2012
  ) {
    let date_begin = moment({
      y: parseInt(message.split(' ')[3]),
      M: 0,
      d: 1
    })
    let date_end = moment({
      y: parseInt(message.split(' ')[3]),
      M: 11,
      d: 31
    })
    const logtime = await get_range_intralogtime(
      message.split(' ')[2],
      date_begin,
      date_end
    )
    const time = format_output_datetime(logtime)
    postOnThread(sprintf(`%02dh%02d`, time[0], time[1]), channel, ts)
    return
  } else if (
    /(\b|^)trimestre[1-4](\b|$)/i.test(message.split(' ')[3]) &&
    (message.split(' ').length === 4 ||
      (message.split(' ').length === 5 &&
        parseInt(message.split(' ')[4]) > 2012))
  ) {
    let quarter = parseInt(message.split(' ')[3].replace('trimestre', '')) - 1
    const year =
      message.split(' ').length === 5 && parseInt(message.split(' ')[4]) > 2012
        ? parseInt(message.split(' ')[4])
        : new Date().getFullYear()
    let date_begin = moment(new Date(year, quarter * 3, 1))
    let date_end = moment(new Date(year, date_begin.get('month') + 3, 0)).add(
      1,
      'days'
    )
    const logtime = await get_range_intralogtime(
      message.split(' ')[2],
      date_begin,
      date_end
    )
    const time = format_output_datetime(logtime)
    postOnThread(sprintf(`%02dh%02d`, time[0], time[1]), channel, ts)
    return
  } else if (
    /(\b|^)semestre[1-2](\b|$)/i.test(message.split(' ')[3]) &&
    (message.split(' ').length === 4 ||
      (message.split(' ').length === 5 &&
        parseInt(message.split(' ')[4]) > 2012))
  ) {
    let semestre = parseInt(message.split(' ')[3].replace('semestre', '')) - 1
    const year =
      message.split(' ').length === 5 && parseInt(message.split(' ')[4]) > 2012
        ? parseInt(message.split(' ')[4])
        : new Date().getFullYear()
    let date_begin = moment(new Date(year, semestre * 6, 1))
    let date_end = moment(new Date(year, date_begin.get('month') + 6, 0)).add(
      1,
      'days'
    )
    const logtime = await get_range_intralogtime(
      message.split(' ')[2],
      date_begin,
      date_end
    )
    const time = format_output_datetime(logtime)
    postOnThread(sprintf(`%02dh%02d`, time[0], time[1]), channel, ts)
    return
  } else if (
    message
      .split(' ')[3]
      .normalize('NFD')
      .replace(/[\u0300-\u036f]/g, '') in month &&
    (message.split(' ').length === 4 ||
      (message.split(' ').length === 5 &&
        parseInt(message.split(' ')[4]) > 2012))
  ) {
    const year =
      message.split(' ').length === 5 && parseInt(message.split(' ')[4]) > 2012
        ? parseInt(message.split(' ')[4])
        : new Date().getFullYear()
    let date_begin = moment(new Date(year, month[message.split(' ')[3]], 1))
    let date_end = moment(
      new Date(year, month[message.split(' ')[3]] + 1, 0)
    ).add(1, 'days')
    const logtime = await get_range_intralogtime(
      message.split(' ')[2],
      date_begin,
      date_end
    )
    let time = format_output_datetime(logtime)
    postOnThread(sprintf(`%02dh%02d`, time[0], time[1]), channel, ts)
    return
  } else if (
    message.split(' ').length === 5 &&
    moment(message.split(' ')[3], moment.ISO_8601, true).isValid() &&
    (message.split(' ')[4] === 'today' ||
      moment(message.split(' ')[4], moment.ISO_8601, true).isValid())
  ) {
    let date_end =
      message.split(' ')[4] === 'today'
        ? moment()
        : moment(message.split(' ')[4])
    let date_begin = moment(message.split(' ')[3])
    if (date_end.isValid() && date_begin.isValid()) {
      const logtime = await get_range_intralogtime(
        message.split(' ')[2],
        date_begin,
        date_end
      )
      const time = format_output_datetime(logtime)
      postOnThread(sprintf(`%02dh%02d`, time[0], time[1]), channel, ts)
      return
    }
  } else
    postOnThread(
      'Usage: cb intralogtime login [datedebut datefin | annee | mois | trimestre [annee]] (date au format "Y-M-D")',
      channel,
      ts
    )
}

const who = async (msg, channel) => {
  if (msg.split(' ').length > 2) place = msg.split(' ')[2]
  else {
    postMessage(`prend une place en parametre`, channel)
    return
  }
  if (!place || place.startsWith('!') || place.startsWith('?')) return
  const url = `/v2/campus/1/locations/?filter[host]=${place}&filter[active]=true`
  const data = await request42(url)
  if (data.length === 0) postMessage(`Place *${place}* vide`, channel)
  else
    postMessage(
      `*${data[0]['user']['login']}* est à la place *${place}*`,
      channel
    )
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
    if (data.length === 0 || data[0]['end_at'])
      postMessage(`*${user}* est hors ligne`, channel)
    else postMessage(`*${user}* est à la place *${data[0]['host']}*`, channel)
    return
  }
  if (msg.split(' ').length > 2) user = msg.split(' ')[2]
  else {
    let username = await getUsername(usr)
    try {
      user = username['user']['email'].strsub(
        0,
        username['user']['email'].indexOf('@')
      )
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
    guardians = [
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
      if (!data || data.length === 0 || data[0]['end_at'])
        postMessage(`*${guardian}* est hors ligne`, channel)
      else
        postMessage(
          `*${guardian}* est à la place *${data[0]['host']}*`,
          channel
        )
    }
    return
  }
  url = `/v2/users/${user}/locations`
  const data = await request42(url)
  if (!data) {
    postMessage(`login invalide`, channel)
    return
  }
  if (data.length === 0 || data[0]['end_at'])
    postMessage(`*${user}* est hors ligne`, channel)
  else postMessage(`*${user}* est à la place *${data[0]['host']}*`, channel)
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
  data.sort(function(a, b) {
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
