/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   citation.js                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: elebouch <elebouch@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/21 18:28:43 by elebouch          #+#    #+#             */
/*   Updated: 2018/02/22 21:21:23 by elebouch         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

const fs = require('fs');
const { postUserMessage } = require('./slack_api');

const citation = async (channel, file, pic, name) => {
    const txt = await fs.readFileSync(file, 'utf-8');
    const citation = txt.split('\n');
    postUserMessage(
        '>' +
            citation[
                Math.floor(Math.random() * Math.floor(citation.length - 1))
            ],
        channel,
        pic,
        name
    ); //postUserMessage(text, channel, image, name)
};

module.exports.citation = citation;
