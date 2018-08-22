/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   index.js                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: elebouch <elebouch@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/19 13:39:56 by elebouch          #+#    #+#             */
/*   Updated: 2018/08/22 15:37:45 by elebouch         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

require('dotenv').config()
const { RTMClient } = require('@slack/client')
const { handleCommand } = require('./handle_command')
const { BLACKLISTCHAN } = require('./const')

/**
 *  Handle Slack RTM
 */

const token = process.env.SLACK_API_TOKEN

const rtm = new RTMClient(token)

rtm.on('authenticated', () => {
  console.log(`Ready`)
})

// Read messages
rtm.on('message', function (message) {
  let channel = message['channel']
  if (!BLACKLISTCHAN.includes(channel) && message.text && message.user) {
    let text = message['text']
    let ts = message['ts']
    let user = message['user']
    handleCommand(text, channel, ts, user)
  }
})

// Start the connecting process
rtm.start()
