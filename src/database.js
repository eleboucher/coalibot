/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   database.js                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: elebouch <elebouch@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/03/09 16:47:42 by elebouch          #+#    #+#             */
/*   Updated: 2018/08/22 15:37:20 by elebouch         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

require('dotenv').config()

const Sequelize = require('sequelize')
const { getUsername } = require('./slack_api')

const sequelize = new Sequelize(
  'coalibot',
  process.env.POSTGRES_USER,
  process.env.POSTGRES_PASSWORD,
  {
    host: process.env.DB_IP,
    dialect: 'postgres',
    logging: false,
    operatorsAliases: false,
    pool: {
      max: 5,
      min: 0,
      acquire: 30000,
      idle: 10000
    }
  }
)

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

module.exports = { addCommand }
