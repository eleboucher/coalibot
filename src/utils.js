/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   utils.js                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: erwanleb <erwanleb@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/23 11:07:52 by erwanleb          #+#    #+#             */
/*   Updated: 2018/02/23 11:08:26 by erwanleboucher   ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

const choose = choices => {
    var index = Math.floor(Math.random() * choices.length);
    return choices[index];
};

module.exports.choose = choose;
