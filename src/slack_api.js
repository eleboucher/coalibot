/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   slack_api.js                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: elebouch <elebouch@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/19 15:13:10 by elebouch          #+#    #+#             */
/*   Updated: 2018/03/02 17:33:20 by elebouch         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

const Slack = require('slack')
const token = process.env.SLACK_API_TOKEN

const bot = new Slack({ token })

function postMessage(text, channel) {
  return bot.chat.postMessage({
    channel: channel,
    text: text
  })
}

function postUserMessage(text, channel, image, name) {
  return bot.chat.postMessage({
    channel: channel,
    text: text,
    icon_url: image,
    username: name
  })
}

function sendReaction(text, channel, ts) {
  return bot.reactions.add({
    timestamp: ts,
    channel: channel,
    name: text
  })
}

function fileUpload(fs, channel) {
  return bot.files.upload({
    channels: channel,
    file: fs
  })
}

function postOnThread(text, channel, ts) {
  return bot.chat.postMessage({
    thread_ts: ts,
    channel: channel,
    text: text
  })
}

function postAttachments(text, attachments, channel) {
  return bot.chat.postMessage({
    channel: channel,
    text: text,
    attachments: attachments
  })
}

function getUsername(user) {
  return bot.users.info({
    user: user
  })
}

module.exports.postMessage = postMessage
module.exports.postUserMessage = postUserMessage
module.exports.sendReaction = sendReaction
module.exports.fileUpload = fileUpload
module.exports.postOnThread = postOnThread
module.exports.getUsername = getUsername
module.exports.postAttachments = postAttachments
