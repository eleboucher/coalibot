/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   42_api.js                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: elebouch <elebouch@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/19 21:07:36 by elebouch          #+#    #+#             */
/*   Updated: 2018/02/20 15:12:38 by elebouch         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

const {
	postMessage,
	postUserMessage,
	sendReaction,
	fileUpload,
	postOnThread,
	getUsername,
	postAttachments
} = require('./slack_api');
var ClientOAuth2 = require('client-oauth2');
var rp = require('request-promise');
var moment = require('moment');
var util = require('util');
const pThrottle = require('p-throttle');

var forty2auth = new ClientOAuth2({
	clientId: process.env.INTRA_CLIENT_ID,
	clientSecret: process.env.INTRA_SECRET,
	accessTokenUri: 'https://api.intra.42.fr/oauth/token'
});

function request42(url) {
	var url = 'https://api.intra.42.fr' + url;
	return forty2auth.credentials.getToken().then(
		pThrottle(
			function(token) {
				var options = {
					uri: url,
					qs: {
						access_token: token.data.access_token
					},
					headers: {
						'User-Agent': 'Request-Promise'
					},
					json: true // Automatically parses the JSON string in the response
				};
				return rp(options)
					.then(function(json) {
						return json;
					})
					.catch(function(err) {
						return NULL;
					});
			},
			2,
			1000
		)
	);
}

function score(channel) {
	request42('/v2/coalitions').then(function(json) {
		console.log(json);
		json.sort(function(a, b) {
			return a.score < b.score;
		});
		var reply = '';
		for (let coa of json) {
			reply += '${name} ${score}\n';
		}
		var attachments = [
			{
				fallback: reply,
				color: json[0]['color'],
				author_link: 'https://profile.intra.42.fr/blocs/1/coalitions',
				fields: [
					{
						title: json[0]['name'],
						value: json[0]['score'],
						short: true
					},
					{
						title: json[1]['name'],
						value: String(
							json[1]['score'] +
								' (' +
								Number(json[1]['score'] - json[0]['score']) +
								')'
						),
						short: true
					},
					{
						title: json[2]['name'],
						value: String(
							json[2]['score'] +
								' (' +
								Number(json[2]['score'] - json[0]['score']) +
								')'
						),
						short: true
					},
					{
						title: json[3]['name'],
						value: String(
							json[3]['score'] +
								' (' +
								Number(json[3]['score'] - json[0]['score']) +
								')'
						),
						short: true
					}
				],
				footer: 'Powered by Coalibot'
			}
		];
		postAttachments('', attachments, channel);
	});
}

function get_range_logtime(user, range_begin, range_end) {
	range_begin = moment(range_begin).format('YYYY-MM-DD');
	range_end = moment(range_end).format('YYYY-MM-DD');
	range_date =
		'?page[size]=100&range[begin_at]=' +
		String(range_begin) +
		',' +
		String(range_end);
	url = '/v2/users/' + user + '/locations' + range_date;

	return request42(url).then(function(data) {
		var logtime = moment.duration(0);
		if (range_begin !== range_end) {
			let tmp = {};
			let i = 2;
			(async () => {
				while (tmp !== null) {
					last_location = moment(
						data[data.length - 1]['begin_at'].slice(0, 10)
					);
					if (moment(range_begin).isBefore(last_location)) {
						console.log("c'est partie");
						tmp = await request42(url + '&page[number]=' + i);
						console.log(tmp);
						data.push(tmp);
						i += 1;
					} else return;
				}
				return;
			})();
		}
		for (let x of data) {
			if (x['end_at']) log_end = moment(x['end_at']);
			else log_end = moment();
			console.log(x['begin_at']);
			log_start = moment(x['begin_at']);
			log_session = log_end - log_start;
			logtime.add(log_session);
		}
		return logtime;
	});
}
function format_output_datetime(time) {
	time = Number(time.as('minutes'));
	hours = Math.floor(time / 60);
	min = Math.floor(time % 60);
	console.log(min);
	return [hours, min];
}

function profile(user, channel) {
	url = '/v2/users/' + user;
	urlcoal = url + '/coalitions/';
	request42(url).then(function(data) {
		let lvl = 1;
		request42(urlcoal).then(function(coaldata) {
			let lvlpiscine = 0;
			if (data['pool_year'] === '2013' || data['pool_year'] === '2014') {
				lvlpiscine = 0;
			} else if (data['cursus_users'].length === 1) {
				lvlpiscine = data['cursus_users'][0]['level'];
				lvl = 0;
			} else lvlpiscine = data['cursus_users'][1]['level'];
			let coalslug = '';
			if (coaldata) coalslug = ':' + coaldata[0]['slug'] + ':';
			range_end = moment();
			range_begin = moment().subtract(7, 'days');
			get_range_logtime(user, range_begin, range_end).then(function(
				logtime
			) {
				var time = format_output_datetime(logtime);
				graph =
					'https://projects.intra.42.fr/projects/graph?login=' + user;
				var stage = (function() {
					for (u of data['projects_users']) {
						if (u['project']['id'] === 118) {
							if (u['status'] === 'finished') return 'A fait son';
							else if (u['status'] === 'in_progress')
								return 'En cours de';
						}
					}
					return "N'a pas fait son";
				})();
				postMessage(
					util.format(
						'%s %s\nPhoto: `%s`\nTemps de log cette semaine %d:%d\nNiveau: %d\nNiveau piscine  %d %s %s\n%s stage\nGraph: %s',
						data['displayname'],
						coalslug,
						data['image_url'],
						time[0],
						time[1],
						lvl === 0 ? 0 : data['cursus_users'][0]['level'],
						lvlpiscine,
						data['pool_month'],
						data['pool_year'],
						stage,
						graph
					),
					channel
				);
			});
		});
	});
}

module.exports.score = score;
module.exports.profile = profile;
