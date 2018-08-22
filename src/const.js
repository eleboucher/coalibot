/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   const.js                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: elebouch <elebouch@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/21 18:07:28 by elebouch          #+#    #+#             */
/*   Updated: 2018/08/22 22:11:42 by elebouch         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

const month = {
  janvier: 0,
  fevrier: 1,
  mars: 2,
  avril: 3,
  mai: 4,
  juin: 5,
  juillet: 6,
  aout: 7,
  septembre: 8,
  octobre: 9,
  novembre: 10,
  decembre: 11
}

const helper = [
  {
    title: 'Coalibot Helper',
    title_link: 'https://github.com/genesixx/coalibot/blob/master/README.md',
    color: '#0c87e1',
    fields: [
      {
        title: 'alliance',
        value: "Stats de l'Alliance",
        short: true
      },
      {
        title: 'addmusic',
        value: 'ajoute une musique a la playlist',
        short: true
      },
      {
        title: 'brew',
        value: 'Commande pour installer brew',
        short: true
      },
      {
        title: 'event',
        value: 'Event de ce jour',
        short: true
      },
      {
        title: 'halp',
        value: 'Instructions pour les problèmes liés à iscsi',
        short: true
      },
      {
        title: 'logtime',
        value: "Pour plus d'info `bc logtime --help`",
        short: true
      },
      {
        title: 'meteo',
        value: "Donne le temps qu'il fait à 42",
        short: true
      },
      {
        title: 'music',
        value: 'Donne une musique aléatoire',
        short: true
      },
      {
        title: 'prof',
        value: " Donne les infos de l'étudiant",
        short: true
      },
      {
        title: 'score',
        value: 'Donne le score des coalitions',
        short: true
      },
      {
        title: 'source',
        value: 'Donne le repo de Coalibot',
        short: true
      },
      {
        title: 'stat',
        value: 'Stat du bot en direct',
        short: true
      },
      {
        title: 'where',
        value: "Donne la position de l'étudiant",
        short: true
      },
      {
        title: 'who',
        value: 'Donne qui est à cette place',
        short: true
      },
      {
        title: 'Bonus',
        short: false
      },
      {
        title: 'coin',
        value: 'pile ou face',
        short: true
      },
      {
        title: 'kaamelott',
        value: 'Citation aléatoire de Kaamelott',
        short: true
      },
      {
        title: 'oss',
        value: 'Citation aléatoire de OSS 117',
        short: true
      },
      {
        title: 'roll',
        value: 'random',
        short: true
      },
      {
        title: 'roulette',
        value: 'Roulette Russe',
        short: true
      }
    ],
    footer: 'Powered by Coalibot'
  }
]

const BLACKLISTCHAN = [
  'CAW4XU70T',
  'G2XKVFENT',
  'C0CPCFLJ1',
  'C7PSZQK7G',
  'C7P52T7U5',
  'C7P0YTCTU',
  'C8URKBK8B',
  'C7P536D5K',
  'C7NF60E0Z',
  'C4LF6DS82',
  'C8U8Z88TB',
  'C7NBD3J5N',
  'C7P0Z4F3L',
  'C13DH4BCY',
  'C7NB7PZ6U',
  'C04CVLBKB',
  'C8C0EGZUY',
  'C0F7G79EE',
  'C04CTPB5X',
  'C04C5F93J',
  'C03V5RBN6',
  'C6VQ9L2TW',
  'C03BN553Z',
  'C03EZET39',
  'C06UVQGFL',
  'C06ANAELW',
  'C7X7V8HV5',
  'C074DJDEC',
  'C039P7U6E',
  'C0FEKP3P1',
  'C04RFBY3K',
  'C0BS44KGV',
  'C4VMBA7P1',
  'C3T7LJB1S',
  'C3SDZFG8Y',
  'C3SES169F',
  'C03DVJ5EF',
  'C0JEKEL0K',
  'C04GT8U3Y',
  'C7UNAUYUF',
  'C03CRTSUC',
  'C91FEPML4',
  'C0G2AQLN9',
  'C7DKMEZKQ',
  'C4W29G7N2',
  'C03DHLN8D',
  'C076RB5DG',
  'C03AV99HK',
  'C5KPM4NE4',
  'C7TQJJHCG',
  'C03DJS74G',
  'C355S7GUA',
  'C719JSE8M',
  'C2T0SGY9J',
  'C736Z2BNZ',
  'C0HTYJR99',
  'C7C27F1ED',
  'C85T5U5S8',
  'C0G5PTGUA',
  'C7801A41Y',
  'C0FGRSTCH',
  'C8BC7CZ09',
  'C0HHL3538',
  'C0GUZ071A',
  'C1A52QDKR',
  'C3B3B500N',
  'C517FS1V1',
  'C5F7H6Y93',
  'C3QG85SG6',
  'C0G68K3KR',
  'C03MP8LEY',
  'C0GLLKY4V',
  'C78HL1T4L',
  'C16G9SSHM',
  'C90S130F2',
  'C7BQK9KLY',
  'C2R6PKH5F',
  'C1LG33P8U',
  'C03C9658F',
  'C70RCSGHY',
  'C39V02M4G',
  'C03KTTYVC',
  'C125TKHDJ',
  'C0K85597Z',
  'C0F8Z1L74',
  'C7LA1QKGE',
  'C8Z3JD82J',
  'C16GM3ZFE',
  'C2S0TPT5L',
  'C0NS5FBU4',
  'C8TLMERMJ',
  'C04J5PVEN',
  'C03SP1BEQ',
  'C098N4V4P',
  'C043W9TDS',
  'C03AV9ADR',
  'C7MAB363D',
  'C3Q44MHV5',
  'C2XGT903X',
  'C5052G9ED',
  'C0HEJU3C1',
  'C03D8415Z',
  'C0G07RYJD',
  'C1XKANLK0',
  'C04GGCC5B',
  'C03S3GBH2',
  'C0J5XDZ29',
  'C04CSJB3V',
  'C0B10J7HQ',
  'C0G0YV3DZ',
  'C03D7R0KL',
  'C33NZAYEP',
  'C0GMR7E2D',
  'C6Z0Y3680',
  'C03KTQ821',
  'C800Y85PA',
  'C7LHBB0UV',
  'C0G8M3HKJ',
  'C18LHA1AB',
  'C6L5QQKTK',
  'C7CK14RGC',
  'C44EZQEQP',
  'C2T0ZB1K7',
  'C6G4V2BL2',
  'C8CLBD8KC',
  'C1XTF6PUN',
  'C4BEGUM2P',
  'C0678LWES',
  'C4MEJLKQF',
  'C3BFMN2S0',
  'C0F8X2DFS',
  'C7NFQKJF9'
]

const parrot = [
  'parrot',
  'middleparrot',
  'rightparrot',
  'aussieparrot',
  'gothparrot',
  'oldtimeyparrot',
  'boredparrot',
  'shuffleparrot',
  'shufflefurtherparrot',
  'congaparrot',
  'reversecongaparrot',
  'partyparrot',
  'sadparrot',
  'parrotcop',
  'fastparrot',
  'slowparrot',
  'parrotdad',
  'dealwithitparrot',
  'fiestaparrot',
  'pizzaparrot',
  'hamburgerparrot',
  'bananaparrot',
  'chillparrot',
  'explodyparrot',
  'shufflepartyparrot',
  'icecreamparrot',
  'sassyparrot',
  'confusedparrot',
  'aussiecongaparrot',
  'aussiereversecongaparrot',
  'parrotwave1',
  'parrotwave2',
  'parrotwave3',
  'parrotwave4',
  'parrotwave5',
  'parrotwave6',
  'parrotwave7',
  'congapartyparrot',
  'moonwalkingparrot',
  'thumbsupparrot',
  'coffeeparrot',
  'parrotwithmustache',
  'christmasparrot',
  'witnessprotectionparrot',
  'parrotsleep',
  'parrotbeer',
  'darkbeerparrot',
  'blondesassyparrot',
  'bluescluesparrot',
  'gentlemanparrot',
  'margaritaparrot',
  'oriolesparrot',
  'dreidelparrot',
  'harrypotterparrot',
  'fieriparrot',
  'upvotepartyparrot',
  'twinsparrot',
  'tripletsparrot',
  'stableparrot',
  'shipitparrot',
  'skiparrot',
  'loveparrot',
  'halalparrot',
  'nyanparrot',
  'wendyparrot',
  'popcornparrot',
  'donutparrot',
  'evilparrot',
  'discoparrot',
  'matrixparrot',
  'papalparrot',
  'stalkerparrot',
  'scienceparrot',
  'prideparrot',
  'revolutionparrot',
  'fidgetparrot',
  'beretparrot',
  'tacoparrot',
  'ryangoslingparrot',
  'john_francis_parrot',
  'mexa_parrot',
  'moneyparrot',
  'moneyparrot2',
  'parrothd'
]

module.exports = { month, helper, parrot, BLACKLISTCHAN }
