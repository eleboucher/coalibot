/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   handle_command.js                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: elebouch <elebouch@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/19 14:39:11 by elebouch          #+#    #+#             */
/*   Updated: 2018/02/24 14:43:21 by elebouch         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

const { postMessage, postUserMessage, sendReaction, fileUpload, postOnThread, getUsername } = require("./slack_api");
const { score, alliance, logtime, profil, who, where } = require("./42_api");
const { citation } = require("./citation");
const { roll, addmusic, music, meteo } = require("./miscs");
const fs = require("fs");
const { parrot, blExcMark } = require("./const");
const { choose } = require("./utils");
const reply = {
    home:
        "Si `Disk Not Ejected Properly ??` suivre : https://42born2code.slack.com/archives/C7P0Z4F3L/p1510233807000241",
    brew:
        "```rm -rf $HOME/.brew && git clone --depth=1 https://github.com/Homebrew/brew $HOME/.brew && echo 'export PATH=$HOME/.brew/bin:$PATH' >> $HOME/.zshrc && source $HOME/.zshrc && brew update```",
    halp:
        "Bonjour\n Je t'invite à taper `iscsictl list targets` dans ton terminal\n à copier la ligne contenant ton login mais sans la partie entre <>\n puis à taper `iscsictl <la ligne copiée>`",
    source: "`https://github.com/genesixx/coalibot`",
    elebouch: "Dodge les BH tel un moine shaolin",
    jcharloi: "fais tes 9h!",
    glegendr: "/giphy how about no",
    makefile: "`https://forum.intra.42.fr/topics/85/messages`",
    sygnano: "https://youtu.be/V2UGfj2qPCw?t=8s",
    nestor:
        "Pour commander sur nestor utilise le code `NESTOR42` ! tu peux utiliser le code de parrainage `cZ44h` pour avoir 5e gratuitement",
    fpons: "Fais quelque chose !!!"
};

functions = {
    alliance: (message, channel, ts, user) => alliance(channel),
    home: (message, channel, ts, user) => postOnThread(reply["home"], channel, ts),
    brew: (message, channel, ts, user) => postOnThread(reply["brew"], channel, ts),
    halp: (message, channel, ts, user) => postOnThread(reply["halp"], channel, ts),
    source: (message, channel, ts, user) => postOnThread(reply["source"], channel, ts),
    score: (message, channel, ts, user) => score(ts, channel),
    help: (message, channel, ts, user) => fileUpload(fs.createReadStream("./featurespic.jpeg"), channel),
    elebouch: (message, channel, ts, user) => postMessage(reply["elebouch"], channel),
    jcharloi: (message, channel, ts, user) => postMessage(reply["jcharloi"], channel),
    glegendr: (message, channel, ts, user) => postMessage(reply["glegendr"], channel),
    makefile: (message, channel, ts, user) => postOnThread(reply["makefile"], channel, ts),
    sygnano: (message, channel, ts, user) => postMessage(reply["sygnano"], channel),
    nestor: (message, channel, ts, user) => postOnThread(reply["nestor"], channel, ts),
    fpons: (message, channel, ts, user) => postMessage(reply["fpons"], channel),
    mfranc: (message, channel, ts, user) =>
        postMessage(choose([">Doucement avec les bots", ">Puuuuuuuuuuuuu"]), channel),
    score: (message, channel, ts, user) => score(channel, ts),
    prof: (message, channel, ts, user) => profil(message.split(" ")[2].toLowerCase(), channel, ts),

    logtime: (message, channel, ts, user) => logtime(message, channel, ts),
    who: (message, channel, ts, user) => who(message.split(" ")[2].toLowerCase(), channel),
    roll: (message, channel, ts, user) => roll(message, channel),
    where: (message, channel, ts, user) => where(message.split(" ")[2].toLowerCase(), channel),
    addmusic: (message, channel, ts, user) => addmusic(message.split(" ")[2], user, channel),
    music: (message, channel, ts, user) => music(channel),
    meteo: (message, channel, ts, user) => meteo(channel),
    oss: (message, channel, ts, user) =>
        citation(
            channel,
            "./oss.txt",
            "https://static-cdn.jtvnw.net/emoticons/v1/518312/3.0",
            "Hubert Bonisseur de La Bath"
        ),
    parrot: (message, channel, ts, user) =>
        postMessage(":" + parrot[Math.floor(Math.random() * Math.floor(parrot.length - 1))] + ":", channel),
    kaamelott: (message, channel, ts, user) =>
        citation(channel, "./kaamelott.txt", "https://img15.hostingpics.net/pics/4833663350.jpg", "Perceval")
};

function handleCommand(msg, channel, ts, user) {
    const message = msg.replace(/\s+/g, " ").trim();
    if (/(\b|^)rip(\b|$)/i.test(message)) sendReaction("rip", channel, ts);
    if (/(\b|^)jpp(\b|$)/i.test(message)) sendReaction("jpp", channel, ts);
    if (/(\b|^)(php|Ruby|RoR|mongo|mongodb)(\b|$)/i.test(message)) sendReaction("poop", channel, ts);
    if (["coalibot", "bc", "cb"].indexOf(message.toLowerCase().split(" ")[0]) > -1) {
        if (message.split(" ")[1].toLowerCase() in functions)
            functions[message.split(" ")[1].toLowerCase()](message, channel, ts, user);
    }
    if (
        message.indexOf("!") === 0 &&
        blExcMark.indexOf(message.replace("!", "").split(" ")[0]) === -1 &&
        message.replace("!", "").split(" ")[0] in functions
    ) {
        functions[message.replace("!", "").split(" ")[0]](message.replace("!", "bc "), channel, ts, user);
    }
}

module.exports.handleCommand = handleCommand;
