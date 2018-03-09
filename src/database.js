/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   database.js                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: elebouch <elebouch@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/03/09 16:47:42 by elebouch          #+#    #+#             */
/*   Updated: 2018/03/09 18:08:54 by elebouch         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

const Sequelize = require('sequelize')
const moment = require('moment')
const { getUsername } = require('./slack_api')

const sequelize = new Sequelize('genesixx_testcoal', process.env.DB_USERNAME, process.env.DB_PASSWORD, {
  host: process.env.DB_HOST,
  dialect: 'postgres',

  pool: {
    max: 5,
    min: 0,
    acquire: 30000,
    idle: 10000
  }
})

const Command = sequelize.define('command', {
  command_name: Sequelize.STRING,
  user: Sequelize.STRING,
  option: Sequelize.STRING,
  date: { type: Sequelize.DATE, defaultValue: Sequelize.NOW }
})

const addCommand = async (cmd, message, channel, ts, user) => {
  await sequelize.sync()
  let username = await getUsername(user)
  if ('user' in username && 'name' in username['user']) {
    username = username['user']['name']
  } else {
    username = ''
  }
  Command.create({
    command_name: cmd,
    user: username,
    option: message
  })
}

module.exports.addCommand = addCommand
