/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   handlecommand.js                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: elebouch <elebouch@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/19 14:39:11 by elebouch          #+#    #+#             */
/*   Updated: 2018/02/19 16:30:18 by elebouch         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

const {
	postMessage,
	postUserMessage,
	sendReaction,
	fileUpload,
	postOnThread,
	getUsername
} = require('./slack_api');

function handleCommand(message, channel, ts, user) {
	username = getUsername(user);
	console.log(username);
	if (username === 'elebouch') postMessage('hello ' + username, channel);
}

module.exports.handleCommand = handleCommand;
