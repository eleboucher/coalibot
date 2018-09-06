/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   logtime.js                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: elebouch <elebouch@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/08/20 23:24:22 by elebouch          #+#    #+#             */
/*   Updated: 2018/09/06 18:57:29 by elebouch         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

const {
  getUsername,
  postAttachmentsOnThread,
  postOnThread
} = require('./slack_api')
const { request42 } = require('./request')
const parse = require('csv-parse/lib/sync')
var fs = require('fs')
const { month } = require('./const')
const moment = require('moment')
const sprintf = require('sprintf-js').sprintf

moment.locale('fr')

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

const getRangeLogtime = async (name, start, end) => {
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

const getRangeIntralogtime = async (user, rangeBegin, rangeEnd) => {
  rangeBegin = moment(rangeBegin).format('YYYY-MM-DD')
  rangeEnd = moment(rangeEnd).format('YYYY-MM-DD')
  let rangeDate = `?page[size]=100&range[begin_at]=${rangeBegin},${rangeEnd}`
  let url = `/v2/users/${user}/locations/${rangeDate}`
  const data = await request42(url)
  if (rangeBegin === rangeEnd) {
    return moment.duration(0)
  }
  try {
    async function getMoreData(data) {
      let tmp
      let i = 2
      let ret = data
      let lastLocation = ''
      do {
        lastLocation = moment(ret[ret.length - 1]['begin_at'])
        if (moment(rangeBegin).isBefore(lastLocation)) {
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
    let locations = await getMoreData(data)
    let logtime = moment.duration(0)
    let logEnd = ''
    let logStart = ''
    let logSession = ''
    for (let x of locations) {
      if (x['end_at']) logEnd = moment(x['end_at'])
      else logEnd = moment()
      logStart = moment(x['begin_at'])
      logSession = logEnd - logStart
      logtime.add(logSession)
    }
    return logtime
  } catch (e) {
    return moment.duration(0)
  }
}

const handleYear = (message, option) => {
  option.count += 1
  if (
    !isNaN(message.split(' ')[option.count]) &&
    parseInt(message.split(' ')[option.count]) > 2000
  ) {
    option.date_begin = moment({
      y: parseInt(message.split(' ')[option.count]),
      M: 0,
      d: 1
    })
    option.date_end = moment({
      y: parseInt(message.split(' ')[option.count]),
      M: 11,
      d: 31
    }).endOf('day')
  } else option.error = true
  option.count += 1
  return option
}

const handleDate = (message, option) => {
  option.count += 1
  if (
    moment(message.split(' ')[option.count], [
      'DD-MM-YYYY',
      moment.ISO_8601,
      'YYYY'
    ]).isValid() &&
    (message.split(' ')[option.count + 1] === 'today' ||
      moment(message.split(' ')[option.count + 1], [
        'DD-MM-YYYY',
        moment.ISO_8601,
        'YYYY'
      ]).isValid())
  ) {
    option.date_end =
      message.split(' ')[option.count + 1] === 'today'
        ? moment()
        : moment(message.split(' ')[option.count + 1], [
          'DD-MM-YYYY',
          moment.ISO_8601,
          'YYYY'
        ])
    option.date_begin = moment(message.split(' ')[option.count], [
      'DD-MM-YYYY',
      moment.ISO_8601,
      'YYYY'
    ])
  } else option.error = true
  option.count += 2
  return option
}

const handleMonth = (message, option) => {
  let hasYear = false
  option.count += 1
  if (
    message
      .split(' ')
    [option.count].normalize('NFD')
      .replace(/[\u0300-\u036f]/g, '') in month
  ) {
    let year = ''
    if (
      message.split(' ')[option.count + 1] &&
      parseInt(message.split(' ')[option.count + 1]) > 2000
    ) {
      year = parseInt(message.split(' ')[option.count + 1])
      hasYear = true
    } else year = new Date().getFullYear()
    option.date_begin = moment({
      y: year,
      M: month[message.split(' ')[option.count]],
      d: 1
    })
    option.date_end = moment({
      y: year,
      M: month[message.split(' ')[option.count]],
      d: option.date_begin.daysInMonth()
    }).endOf('day')
  } else if (
    /(\b|^)(0[1-9]|[1-9]|1[012])(\b|$)/i.test(message.split(' ')[option.count])
  ) {
    let year = ''
    if (
      message.split(' ')[option.count + 1] &&
      parseInt(message.split(' ')[option.count + 1]) > 2000
    ) {
      year = parseInt(message.split(' ')[option.count + 1])
      hasYear = true
    } else year = new Date().getFullYear()
    option.date_begin = moment({
      y: year,
      M: parseInt(message.split(' ')[option.count]) - 1,
      d: 1
    })
    option.date_end = moment({
      y: year,
      M: parseInt(message.split(' ')[option.count]) - 1,
      d: option.date_begin.daysInMonth()
    }).endOf('day')
  } else option.error = true
  option.count += hasYear ? 2 : 1
  return option
}

const handleQuarter = (message, option) => {
  option.count += 1
  if (
    message.split(' ')[option.count] &&
    /(\b|^)[1-4](\b|$)/i.test(message.split(' ')[option.count])
  ) {
    let quarter = parseInt(message.split(' ')[option.count]) - 1
    let year = ''
    if (
      message.split(' ')[option.count + 1] &&
      parseInt(message.split(' ')[option.count + 1]) > 2000
    ) {
      year = parseInt(message.split(' ')[option.count + 1])
      option.count += 1
    } else year = new Date().getFullYear()
    option.date_begin = moment(new Date(year, quarter * 3, 1))
    option.date_end = moment(
      new Date(year, option.date_begin.get('month') + 3, 0)
    ).endOf('day')
  } else option.error = true
  option.count += 1
  return option
}

const handleSemester = (message, option) => {
  option.count += 1
  if (
    message.split(' ')[option.count] &&
    /(\b|^)[1-2](\b|$)/i.test(message.split(' ')[option.count])
  ) {
    let semester = parseInt(message.split(' ')[option.count]) - 1
    let year = ''
    if (
      message.split(' ')[option.count + 1] &&
      parseInt(message.split(' ')[option.count + 1]) > 2000
    ) {
      year = parseInt(message.split(' ')[option.count + 1])
      option.count += 1
    } else year = new Date().getFullYear()
    option.date_begin = moment(new Date(year, semester * 6, 1))
    option.date_end = moment(
      new Date(year, option.date_begin.get('month') + 6, 0)
    )
  } else option.error = true
  option.count += 1
  return option
}

const formatOutputDatetime = time => {
  const timem = Number(time.as('minutes'))
  const hours = Math.floor(timem / 60)
  const min = Math.floor(timem % 60)
  return [hours, min]
}

const logtime = async (message, user, channel, ts) => {
  const usage = `\`\`\`Usage:  bc logtime [OPTION] [login]
      Les temps par defaut sont ceux de la badgeuse.
      -i                   temps de l'intra,
      -d debut fin         donne le logtime dans la periode des dates données. format DD/MM/YYYY
      -y année             donne le logtime durant l'année donnée.
      -m mois [année]      donne le logtime durant le mois donné.
      -t trimestre [année] donne le logtime durant le trimestre donné.
      -s semestre [année]  donne le logtime durant le semestre donné.\`\`\``
  let option = {
    count: 2,
    intra: false,
    date_begin: '',
    date_end: '',
    login: '',
    logtime: '',
    error: false
  }

  if (
    message.split(' ')[option.count] === '-i' ||
    message.split(' ')[option.count] === '--intra'
  ) {
    option.intra = true
    option.count += 1
  }

  switch (message.split(' ')[option.count]) {
    case '--date':
    case '-d':
      option = handleDate(message, option)
      break
    case '--year':
    case '-y':
      option = handleYear(message, option)
      break
    case '--month':
    case '-m':
      option = handleMonth(message, option)
      break
    case '--trimestre':
    case '-t':
      option = handleQuarter(message, option)
      break
    case '--semestre':
    case '-s':
      option = handleSemester(message, option)
      break
    case '-h':
    case '--help':
      option.error = true
      break
    default:
      option.date_begin = moment({
        y: new Date().getFullYear(),
        M: 0,
        d: 1
      })
      option.date_end = moment({
        y: new Date().getFullYear(),
        M: 11,
        d: 31
      }).endOf('day')
      break
  }

  if (
    message.split(' ')[option.count] === '-i' ||
    message.split(' ')[option.count] === '--intra'
  ) {
    option.intra = true
    option.count += 1
  }

  if (message.split(' ')[option.count] && !option.error) {
    option.login = message.split(' ')[option.count]
    option.count += 1
  } else {
    let username = await getUsername(user)
    try {
      option.login = username['user']['profile']['email']
        .toString()
        .substr(0, username['user']['profile']['email'].toString().indexOf('@'))
    } catch (err) {
      option.login = username['user']['name']
    }
  }

  if (option.count !== message.split(' ').length) {
    option.error = true
  }
  if (option.date_begin !== '' && option.date_end !== '' && !option.error) {
    switch (option.intra) {
      case false:
        const intradata = await request42('/v2/users/' + option.login)
        let name = {}
        if (intradata && intradata.last_name && intradata.first_name) {
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
        } else {
          option.error = true
        }
        option.logtime = await getRangeLogtime(
          name,
          option.date_begin,
          option.date_end
        )
        break
      case true:
        option.logtime = await getRangeIntralogtime(
          option.login,
          option.date_begin,
          option.date_end
        )
    }
  }
  if (option.error === true) {
    postOnThread(usage, channel, ts)
    return
  }

  if (option.logtime !== '') {
    let logtimeString = option.intra
      ? sprintf(
        '%02dh%02d',
        formatOutputDatetime(option.logtime)[0],
        formatOutputDatetime(option.logtime)[1]
      )
      : option.logtime + 'h'
    let attachment = [
      {
        fields: [
          {
            title: 'Résultat',
            value: logtimeString
          }
        ],
        color: 'good'
      }
    ]
    postAttachmentsOnThread(
      `Logtime ${option.intra === true ? 'intra' : 'badgeuse'} pour *${
      option.login
      }* entre *${option.date_begin.format('LL')}* et *${option.date_end.format(
        'LL'
      )}*`,
      attachment,
      channel,
      ts
    )
  }
}

module.exports = { logtime, getRangeIntralogtime, formatOutputDatetime }
