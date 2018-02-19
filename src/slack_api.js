/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   slack_api.js                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: elebouch <elebouch@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/19 15:13:10 by elebouch          #+#    #+#             */
/*   Updated: 2018/02/19 16:43:24 by elebouch         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

const Slack = require('slack');
const token = process.env.SLACK_API_TOKEN;

const bot = new Slack({ token });

function postMessage(text, channel) {
	bot.chat.postMessage({
		channel: channel,
		text: text
	});
}

function postUserMessage(text, channel, image, name) {
	bot.chat.postMessage({
		channel: channel,
		text: text,
		icon_url: image,
		name: name
	});
}

function sendReaction(text, channel, ts) {
	bot.reactions.add({
		timestamp: ts,
		channel: channel,
		name: text
	});
}

function fileUpload(fs, channel) {
	bot.files.upload({
		channels: channel,
		file: fp
	});
}

function postOnThread(text, channel, ts) {
	bot.postMessage({
		thread_ts: ts,
		channel: channel,
		text: text
	});
}

async function getUsername(user) {
	return await bot.users
		.info({
			user: user
		})
		.then(function(username) {
			if ('user' in username && 'name' in username['user']) {
				return username['user']['name'];
			}
			return null;
		});
}

module.exports.postMessage = postMessage;
module.exports.postUserMessage = postUserMessage;
module.exports.sendReaction = sendReaction;
module.exports.fileUpload = fileUpload;
module.exports.postOnThread = postOnThread;
module.exports.getUsername = getUsername;
