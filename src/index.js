/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   index.js                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: elebouch <elebouch@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/19 13:39:56 by elebouch          #+#    #+#             */
/*   Updated: 2018/05/04 19:10:59 by elebouch         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

const { RTMClient, CLIENT_EVENTS, RTM_EVENTS } = require('@slack/client')
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
rtm.on('message', function(message) {
  var channel = message['channel']
  if (!BLACKLISTCHAN.includes(channel) && message.text && message.user) {
    var text = message['text']
    var ts = message['ts']
    var user = message['user']
    handleCommand(text, channel, ts, user)
  }
})

// Start the connecting process
rtm.start()
