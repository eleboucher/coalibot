/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   miscs.js                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: elebouch <elebouch@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/22 14:27:58 by elebouch          #+#    #+#             */
/*   Updated: 2018/02/22 16:57:43 by elebouch         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

const {
	postMessage,
	postUserMessage,
	sendReaction,
	fileUpload,
	postOnThread,
	getUsername,
} = require('./slack_api');
const fs = require('fs');
var rp = require('request-promise');

const roll = (message, channel) => {
	console.log(message.split(' ').length);
	if (
		message.split(' ').length !== 4 ||
		isNaN(message.split(' ')[2]) ||
		isNaN(message.split(' ')[3])
	) {
		postMessage('Usage: bc roll nbde tailledude', channel);
		return;
	}
	let str = '';
	let length = parseInt(message.split(' ')[2]);
	let max = parseInt(message.split(' ')[3]);
	if (length > 1000 || max > 1000000 || length < 0 || max < 0) {
		postMessage('nbde max == 100 et tailledude max == 1000000', channel);
		return;
	}
	for (let i = 0; i <= length; i++) {
		str += ' ' + Math.floor(Math.random() * Math.floor(max));
	}
	postMessage(str, channel);
};

const youtubeChecker = async link => {
	var options = {
		uri: link,
	};
	try {
		const request = await rp(options);
		return 1;
	} catch (err) {
		return 0;
	}
};

const addmusic = async (link, user, channel) => {
	let json = await fs.readFileSync('./music.json', 'utf-8');
	json = JSON.parse(json);
	const checker = /(?:youtube(?:-nocookie)?\.com\/(?:[^\/\n\s]+\/\S+\/|(?:v|e(?:mbed)?)\/|\S*?[?&]v=)|youtu\.be\/)([a-zA-Z0-9_-]{11})/;
	if (
		(link.includes('youtu') && youtubeChecker(link)) ||
		link.includes('soundcloud')
	) {
		let username = await getUsername(user);
		if ('user' in username && 'name' in username['user']) {
			username = username['user']['name'];
		} else {
			username = '';
		}
		const checkduplicate = (link, json) => {
			for (let u of json) {
				if (u['link'] === link) return false;
			}
			return true;
		};
		if (checkduplicate(link, json) === true) {
			console.log('oui');
			const info = {
				login: username,
				link: link,
			};
			json = json.concat(info);
			fs.writeFile('./music.json', JSON.stringify(json), 'utf8', err => {
				if (err) throw err;
			});
			postMessage('Musique ajoutée', channel);
			return;
		} else {
			postMessage('Lien déjà enregistré', channel);
		}
	} else postMessage('Lien incorrect', channel);
};
const music = async channel => {
	let json = await fs.readFileSync('./music.json', 'utf-8');
	json = JSON.parse(json);
	const music = json[Math.floor(Math.random() * Math.floor(json.length))];
	let login;
	if (music.login === 'pk') login = 'p/k';
	else login = music.login;
	postMessage(`${login} ${music.link}`, channel);
};

module.exports.roll = roll;
module.exports.addmusic = addmusic;
module.exports.music = music;
