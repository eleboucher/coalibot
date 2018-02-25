/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   giphy.js                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: elebouch <elebouch@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/24 15:03:45 by elebouch          #+#    #+#             */
/*   Updated: 2018/02/25 12:10:29 by elebouch         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

var rp = require("request-promise");
const { postMessage, postUserMessage, sendReaction, fileUpload, postOnThread, getUsername } = require("./slack_api");

const randomgif = async (query, channel) => {
    //not my api key :3:
    try {
        const res = await rp(
            `http://api.giphy.com/v1/gifs/search?q=${query}&api_key=aH0B2QT0mjnLZ7xKvHTwURhIHcIiB4MR&limit=30`,
            { json: true }
        );
        const theOnlyOne = res.data[Math.floor(Math.random() * Math.floor(res.data.length - 1))].url;
        postMessage(theOnlyOne, channel);
    } catch (err) {
        postMessage("Aucun gif trouvé", channel);
    }
};

module.exports.randomgif = randomgif;
