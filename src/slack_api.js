/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   slack_api.js                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: elebouch <elebouch@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/19 15:13:10 by elebouch          #+#    #+#             */
/*   Updated: 2018/08/21 10:45:55 by elebouch         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

const Slack = require('slack')
const token = process.env.SLACK_API_TOKEN

const bot = new Slack({ token })

const postMessage = (text, channel) => {
  return bot.chat.postMessage({
    channel: channel,
    text: text
  })
}

const postUserMessage = (text, channel, image, name) => {
  return bot.chat.postMessage({
    channel: channel,
    text: text,
    icon_url: image,
    username: name
  })
}

const sendReaction = (text, channel, ts) => {
  return bot.reactions.add({
    timestamp: ts,
    channel: channel,
    name: text
  })
}

const fileUpload = (fs, channel) => {
  return bot.files.upload({
    channels: channel,
    file: fs
  })
}

const postOnThread = (text, channel, ts) => {
  return bot.chat.postMessage({
    thread_ts: ts,
    channel: channel,
    text: text
  })
}

const postAttachments = (text, attachments, channel) => {
  return bot.chat.postMessage({
    channel: channel,
    text: text,
    attachments: attachments
  })
}

const postAttachmentsOnThread = (text, attachments, channel, ts) => {
  return bot.chat.postMessage({
    thread_ts: ts,
    channel: channel,
    text: text,
    attachments: attachments
  })
}

const getUsername = user => {
  return bot.users.info({
    user: user
  })
}

module.exports = {
  postAttachments,
  postAttachmentsOnThread,
  postMessage,
  postOnThread,
  postUserMessage,
  fileUpload,
  sendReaction,
  getUsername
}
