/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   citation.js                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: elebouch <elebouch@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/21 18:28:43 by elebouch          #+#    #+#             */
/*   Updated: 2018/04/24 18:37:08 by elebouch         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

const fs = require('fs')
const { postUserMessage } = require('./slack_api')

const citation = async (channel, file, pic, name) => {
  const txt = await fs.readFileSync(file, 'utf-8')
  const citation = txt.split('\n').filter(String)
  postUserMessage('>' + citation[Math.floor(Math.random() * Math.floor(citation.length))], channel, pic, name) //postUserMessage(text, channel, image, name)
}

module.exports.citation = citation
