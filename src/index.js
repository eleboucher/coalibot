/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   index.js                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: elebouch <elebouch@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/19 13:39:56 by elebouch          #+#    #+#             */
/*   Updated: 2018/05/04 15:26:10 by elebouch         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

const { RTMClient, CLIENT_EVENTS, RTM_EVENTS } = require('@slack/client')
const { handleCommand } = require('./handle_command')
const { BLACKLISTCHAN } = require('./blacklist_channel')

/**
 *  Handle Slack RTM
 */

const token = process.env.SLACK_API_TOKEN

// Cache of data
const appData = {}

// Initialize the RTM client with the recommended settings. Using the defaults for these
// settings is deprecated.
const rtm = new RTMClient(token, {
  dataStore: false,
  useRtmConnect: true
})

// The client will emit an RTM.RTM_CONNECTION_OPEN the connection is ready for
// sending and recieving messages
rtm.on('authenticated', () => {
  console.log(`Ready`)
})

// Read messages
rtm.on('message', function(message) {
  var channel = message['channel']
  if (!BLACKLISTCHAN.includes(channel) && message.text) {
    var text = message['text']
    var ts = message['ts']
    var user = message['user']
    handleCommand(text, channel, ts, user)
  }
})

// Start the connecting process
rtm.start()
