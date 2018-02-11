# coding=utf-8

import time
from slackclient import SlackClient
from configparser import ConfigParser
from sanction import Client
from datetime import datetime, timedelta, date
import random
import sys
import time
import requests
import json
import commands
import os
from binance.client import Client as binanceClient
from currency_converter import CurrencyConverter
import calendar
try:
    from urllib import quote_plus
except:
    from urllib.parse import quote_plus

reload(sys)
sys.setdefaultencoding('utf8')
sc = SlackClient(os.environ["SLACK_API_TOKEN"])
parrot = ["parrot", "middleparrot", "rightparrot", "aussieparrot", "gothparrot", "oldtimeyparrot", "boredparrot", "shuffleparrot", "shufflefurtherparrot", "congaparrot", "reversecongaparrot", "partyparrot", "sadparrot", "parrotcop", "fastparrot", "slowparrot", "parrotdad", "dealwithitparrot", "fiestaparrot", "pizzaparrot", "hamburgerparrot", "bananaparrot", "chillparrot", "explodyparrot", "shufflepartyparrot", "icecreamparrot", "sassyparrot", "confusedparrot", "aussiecongaparrot", "aussiereversecongaparrot", "parrotwave1", "parrotwave2", "parrotwave3", "parrotwave4", "parrotwave5", "parrotwave6", "parrotwave7", "congapartyparrot", "moonwalkingparrot", "thumbsupparrot", "coffeeparrot", "parrotwithmustache", "christmasparrot", "witnessprotectionparrot", "parrotsleep", "parrotbeer", "darkbeerparrot", "blondesassyparrot", "bluescluesparrot", "gentlemanparrot", "margaritaparrot", "oriolesparrot", "dreidelparrot", "harrypotterparrot", "fieriparrot", "upvotepartyparrot", "twinsparrot", "tripletsparrot", "stableparrot", "shipitparrot", "skiparrot", "loveparrot", "halalparrot", "nyanparrot", "wendyparrot", "popcornparrot", "donutparrot", "evilparrot", "discoparrot", "matrixparrot", "papalparrot", "stalkerparrot", "scienceparrot", "prideparrot", "revolutionparrot", "fidgetparrot", "beretparrot", "tacoparrot", "ryangoslingparrot", "john_francis_parrot", "mexa_parrot", "moneyparrot", "moneyparrot2", "parrothd" ]

coalitions = ['the-alliance','the-order','the-federation','the-assembly']
blacklist = ["C0CPCFLJ1", "C7PSZQK7G", "C7P52T7U5", "C7P0YTCTU", "C8URKBK8B", "C7P536D5K", "C7NF60E0Z", "C4LF6DS82", "C8U8Z88TB", "C7NBD3J5N", "C7P0Z4F3L", "C13DH4BCY", "C7NB7PZ6U", "C04CVLBKB", "C8C0EGZUY", "C0F7G79EE", "C04CTPB5X", "C04C5F93J", "C03V5RBN6", "C6VQ9L2TW", "C03BN553Z", "C03EZET39", "C06UVQGFL", "C06ANAELW", "C7X7V8HV5", "C074DJDEC", "C039P7U6E", "C0FEKP3P1", "C04RFBY3K", "C0BS44KGV", "C4VMBA7P1", "C3T7LJB1S", "C3SDZFG8Y", "C3SES169F", "C03DVJ5EF", "C0JEKEL0K", "C04GT8U3Y", "C7UNAUYUF", "C03CRTSUC", "C91FEPML4", "C0G2AQLN9", "C7DKMEZKQ", "C4W29G7N2", "C03DHLN8D", "C076RB5DG", "C03KS21SB", "C03AV99HK", "C5KPM4NE4", "C7TQJJHCG", "C03DJS74G", "C355S7GUA", "C719JSE8M", "C2T0SGY9J", "C736Z2BNZ", "C0HTYJR99", "C7C27F1ED", "C85T5U5S8", "C0G5PTGUA", "C7801A41Y", "C0FGRSTCH", "C8BC7CZ09", "C0HHL3538", "C0GUZ071A", "C1A52QDKR", "C3B3B500N", "C517FS1V1", "C5F7H6Y93", "C3QG85SG6", "C0G68K3KR", "C03MP8LEY", "C0GLLKY4V", "C78HL1T4L", "C16G9SSHM", "C90S130F2", "C7BQK9KLY", "C2R6PKH5F", "C1LG33P8U", "C03C9658F", "C70RCSGHY", "C39V02M4G", "C03KTTYVC", "C125TKHDJ", "C0K85597Z", "C0F8Z1L74", "C7LA1QKGE", "C8Z3JD82J", "C16GM3ZFE", "C2S0TPT5L", "C0NS5FBU4", "C8TLMERMJ", "C04J5PVEN", "C03SP1BEQ", "C098N4V4P", "C043W9TDS", "C03AV9ADR", "C7MAB363D", "C3Q44MHV5", "C2XGT903X", "C5052G9ED", "C0HEJU3C1", "C03D8415Z", "C0G07RYJD", "C1XKANLK0", "C04GGCC5B", "C03S3GBH2", "C0J5XDZ29", "C04CSJB3V", "C0B10J7HQ", "C0G0YV3DZ", "C03D7R0KL", "C33NZAYEP", "C0GMR7E2D", "C6Z0Y3680", "C03KTQ821", "C800Y85PA", "C7LHBB0UV", "C0G8M3HKJ", "C18LHA1AB", "C6L5QQKTK", "C7CK14RGC", "C44EZQEQP", "C2T0ZB1K7", "C6G4V2BL2", "C8CLBD8KC", "C1XTF6PUN", "C4BEGUM2P", "C0678LWES", "C4MEJLKQF", "C3BFMN2S0", "C0F8X2DFS", "C7NFQKJF9"]
trimestre = {"trimestre1" : 1, "trimestre2" : 2, "trimestre3" : 3, "trimestre4" : 4}
month = {"janvier": 1, "fevrier": 2, "mars": 3, "avril": 4, "mai": 5, "juin": 6, "juillet": 7, "aout": 8, "septembre": 9, "octobre": 10,"novembre": 11,"decembre": 2 }
def get_token(grant_type):
    client = Client(token_endpoint = "https://api.intra.42.fr/oauth/token",
            resource_endpoint = "https://api.intra.42.fr",
            client_id = os.environ["INTRA_CLIENT_ID"],
            client_secret =  os.environ["INTRA_SECRET"])
    client.request_token(grant_type=grant_type)
    time.sleep(1)
    return client

def get_username(user):
    username = sc.api_call(
                "users.info",
                user = user
                )

    if ('user' in username and 'name' in username['user']):
        return (username['user']['name'])
    else :
        return "null"

def get_first_day_of_the_quarter(quarter, year):
    return datetime.strptime(str((datetime(year, 3 * quarter - 2, 1)).date()), '%Y-%m-%d').date()

def get_last_day_of_the_quarter(quarter, year):
    month = 3 * quarter
    remaining = month / 12
    return datetime.strptime(str((datetime(year + remaining, month % 12 + 1, 1) +timedelta(days=-1)).date()), '%Y-%m-%d').date()

def get_more_location(client, request, locations, range_begin):
    try:
        tmp = True
        i = 2;
        while tmp:
            last_location = datetime.strptime((locations[-1]['begin_at'])[:10], "%Y-%m-%d")
            if range_begin < last_location:
                tmp = client.request(request + "&page[number]=" + str(i))
                locations += tmp
                i += 1
                time.sleep(1)
            else:
                return
    except :
        return

def get_range_logtime (user, range_begin, range_end):
    range_begin = datetime.strptime(str(range_begin), '%Y-%m-%d')
    range_end = datetime.strptime(str(range_end), '%Y-%m-%d')
    range_date = "?page[size]=100&range[begin_at]=" + str(range_begin) + "," + str(range_end)
    url = "/v2/users/" + user + "/locations" + range_date
    client = get_token("client_credentials")
    try:
        data = client.request(url)
    except IOError:
        return 0;
    logtime = timedelta()
    if range_begin != range_end :
        get_more_location(client, url, data, range_begin)
    for x in data :
        if x['end_at']:
            log_end = datetime.strptime((x['end_at'])[:19], "%Y-%m-%dT%H:%M:%S")
        else:
            log_end = datetime.strptime((str(datetime.utcnow()))[:19], "%Y-%m-%d %H:%M:%S")
        log_start = datetime.strptime((x['begin_at'])[:19], "%Y-%m-%dT%H:%M:%S")
        log_session = log_end - log_start
        logtime += log_session
    return logtime

def logtime(message, ts, channel):
    reply = ""
    if len(message.split( )) < 4 :
        reply = "Usage: bc logtime login datedebut datefin (date au format \"Y-M-D\")"
    elif len(message.split( )) == 4 and message.split( )[3].isdigit() and (int(message.split( )[3]) > 2012 and int(message.split( )[3]) < 2030) :
        date_begin = date(int(message.split( )[3]), 1, 1)
        date_end = date(int(message.split( )[3]), 12, 31)
        logtime = get_range_logtime(message.split( )[2], date_begin, date_end)
        try:
            (h, m) = format_output_datetime(logtime.days * 86400 + logtime.seconds)
        except :
            h = 0
            m = 0
        reply = "{:02d}h{:02d}".format(h,m)
    elif "trimestre" in message.split( )[3] and (len(message.split( )) == 4 or len(message.split( )) == 5 and (int(message.split( )[4]) > 2012 and int(message.split( )[4]) < 2030)):
            quarter = int(message.split( )[3].replace("trimestre", ""))
            if len(message.split( )) == 5 and (int(message.split( )[4]) > 2012 and int(message.split( )[4]) < 2030):
                year = int(message.split( )[4])
            else:
                year = datetime.now().year
            if quarter <= 4 and quarter > 0:
                date_begin = str(get_first_day_of_the_quarter(quarter, year))
                date_end = str(get_last_day_of_the_quarter(quarter, year))
                logtime = get_range_logtime(message.split( )[2], date_begin, date_end)
                try:
                    (h, m) = format_output_datetime(logtime.days * 86400 + logtime.seconds)
                except :
                    h = 0
                    m = 0
                reply = "{:02d}h{:02d}".format(h,m)
    elif message.split( )[3] in month.keys() and (len(message.split( )) == 4 or (len(message.split( )) == 5 and int(message.split( )[4]) > 2012 and int(message.split( )[4]) < 2030)):
        if len(message.split( )) == 5 and (int(message.split( )[4]) > 2012 and int(message.split( )[4]) < 2030):
            year = int(message.split( )[4])
        else:
            year = datetime.now().year
        _, num_days = calendar.monthrange(year, month.get(message.split( )[3]))
        date_begin = date(year, month.get(message.split( )[3]), 1)
        date_end = date(year, month.get(message.split( )[3]), num_days)
        logtime = get_range_logtime(message.split( )[2], date_begin, date_end)
        try:
            (h, m) = format_output_datetime(logtime.days * 86400 + logtime.seconds)
        except :
            h = 0
            m = 0
        reply = "{:02d}h{:02d}".format(h,m)
    elif len(message.split( )) == 5:
        if "today" in message.split( )[4]:
            date_end = str(date.today())
        else :
            date_end = message.split( )[4]
        if validate(date_end) == 1 and validate(message.split( )[3]) == 1 and "!" not in message.split( )[2] :
            logtime = get_range_logtime(message.split( )[2], message.split( )[3], date_end)
            try:
                (h, m) = format_output_datetime(logtime.days * 86400 + logtime.seconds)
            except :
                h = 0
                m = 0
            reply = "{:02d}h{:02d}".format(h,m)
    sc.api_call(
        "chat.postMessage",
        thread_ts = ts,
        channel = channel,
        text = reply
        )

def format_output_datetime(duration_timedelta):
    hours, remainder = divmod(duration_timedelta, 3600)
    minutes, seconds = divmod(remainder, 60)
    return (hours, minutes)

def hasdoneintern(user):
    client=get_token("client_credentials")
    url = "/v2/users/"+ user
    data = client.request(url)
    for u in data['projects_users']:
        if u['project']['id'] is 118:
            if u['status'] == "finished" :
                return "A fait son"
            elif u['status'] == "in_progress" :
                return "En cours de"
    return "N'a pas fait son"

def profile(user):
    url = "/v2/users/" + user
    urlcoal = url + "/coalitions/"
    client = get_token("client_credentials")
    lvl = 1
    try:
        data = client.request(url)
    except IOError:
        return ":3:"
    coal = client.request(urlcoal)
    if user == None or user == "" or "?" in user :
        return ":3:"
    if data['pool_year'] == "2013" or data['pool_year'] == "2014":
        lvlpiscine = 0
    elif len(data['cursus_users']) == 1:
        lvlpiscine = data['cursus_users'][0]['level']
        lvl = 0
    else:
        lvlpiscine = data['cursus_users'][1]['level']
    if not coal:
        coalslug = ""
    else:
        coalslug = ":" + coal[0]['slug'] + ":"
    range_end = str(date.today())
    range_begin = str(date.today() - timedelta(days=7))
    logtime = get_range_logtime(user, range_begin, range_end)
    (h, m) = format_output_datetime(logtime.days * 86400 + logtime.seconds)
    graph = "https://projects.intra.42.fr/projects/graph?login=" + user
    return "{} {}\nPhoto: {}\nTemps de log cette semaine {:02d}:{:02d}\nNiveau: {:.02f} {}\nNiveau piscine {:0.2f} {} {}\n{} stage\nGraph: {}".format(data['displayname'], coalslug,data['image_url'],h ,m, 0 if lvl == 0 else data['cursus_users'][0]['level'], ('' if lvl == 0 else data['cursus_users'][0]['grade']), lvlpiscine , data['pool_month'], data['pool_year'], hasdoneintern(user), graph)

def score(ts, channel):
    url = "/v2/coalitions"
    client = get_token("client_credentials")
    reply = ""
    try:
        data = client.request(url)
    except IOError:
        return ":3:"
    data.sort(key=lambda k: k['score'], reverse=True)
    for coa in data :
        reply += "{} {}\n".format(coa['name'],coa['score'])
    sc.api_call(
        "chat.postMessage",
        channel=channel,
        text="",
        ts = ts,
        attachments = [
                {
                    'fallback': reply,
                    'color' : data[0]['color'],
                    'author_link': "https:\/\/profile.intra.42.fr\/blocs\/1\/coalitions",
                    'fields': [
                        {
                            "title": data[0]['name'],
                            "value": data[0]['score'],
                            "short": True
                        },
                        {
                            "title": data[1]['name'],
                            "value": str(data[1]['score']) + " ("+ str(int(data[1]['score']) - int(data[0]['score']))  +")",
                            "short": True
                        },
                        {
                            "title": data[2]['name'],
                            "value": str(data[2]['score']) + " ("+ str(int(data[2]['score']) - int(data[0]['score']))  +")",
                            "short": True
                        },
                        {
                            "title": data[3]['name'],
                            "value": str(data[3]['score']) + " ("+ str(int(data[3]['score']) - int(data[0]['score']))  +")",
                            "short": True
                        }
                    ],
                    'footer' : 'Powered by Coalibot'
                }
            ]
        )


def checkduplicate(content, link):
    for x in content:
        if link in x['link']:
            return True
    return False

def youtube_url_validation(url):
    youtube_regex = (
        r'(https?://)?(www\.)?'
        '(youtube|youtu|youtube-nocookie)\.(com|be)/'
        '(watch\?v=|embed/|v/|.+\?v=)?([^&=%\?]{11})')

    youtube_regex_match = re.match(youtube_regex, url)
    if youtube_regex_match:
        return 1
    else :
        return 0

def addmusic(link, user):
    with open('music.json', 'r') as fp:
        content = json.load(fp)
    if (("youtube" in link or "youtu.be" in link) or "soundcloud.com" in link) and  checkduplicate(content, link) == False :
        info = {
                "login": get_username(user),
                "link": link
                }
        content.append(info)
        with open("music.json", "w") as output:
            json.dump(content, output, indent=2)
        return "Musique ajoutée"
    else :
        return "Lien incorrect"

def musique():
    try:
        with open('music.json', 'r') as fp:
           content = json.load(fp)
    except IOError:
        return "erreur de fichier"
    choice = content[random.randint(0, len(content) - 1)]
    if "pk" in choice['login']:
        login = "p/k"
    else:
        login = choice['login']
    return "{} {}".format(login, choice['link'])

def who(place):
    if "!" in place :
        return
    if "?" in place :
        return
    url = "/v2/campus/1/locations/?filter[host]=" + place + "&filter[active]=true"
    client = get_token("client_credentials")
    try:
        data = client.request(url)
    except IOError:
        return "saisie incorrecte"
    if len (data) == 0:
        return "Place *{}* vide".format(place)
    else:
        return "*{}* est à la place *{}*".format(data[0]['user']['login'], place)

def where(user):
    if "!" in user:
        return
    if "?" in user:
        return
    if user == "queen" or user == "way":
	return "follow me bruddah\ni'll show you de way :uganda_knuckles:"
    url = "/v2/users/" + user + "/locations"
    client = get_token("client_credentials")
    try:
        data = client.request(url)
    except IOError:
        return "login incorrect"
    if len (data) == 0 or data[0]['end_at'] is not None:
        return "*{}* est hors ligne".format(user)
    else:
        return "*{}* est à la place *{}*".format(user, data[0]['host'])

def alliance():
    url = "/v2/coalitions"
    client = get_token("client_credentials")
    try:
        data = client.request(url)
    except IOError:
        return ":3:"
    data.sort(key=lambda k: k['score'], reverse=True)
    rang = 0
    diff_score = 0
    while data[rang]['id'] is not 2:
        rang += 1
    if rang is not 0:
        diff_score = data[0]['score'] - data[rang]['score']
        return "Nous sommes à la {}eme place avec {} points de retard. :the-alliance:".format((rang + 1), diff_score)
    else :
        diff_score = data[rang]['score'] - data[1]['score']
        return "Felicitations Nous sommes premiers avec {} points d'avance. :the-alliance:".format((diff_score))

def roll(nb, taille):
    if nb.isdigit() == False and nb.isdigit() == False:
        return "Usage: bc roll nbde tailledude"
        return
    string = ""
    if int(nb) <= 100 and int(taille) <= 10000 and int(taille) > 0:
        for i in range (1,int(nb) + 1):
            string += str(random.randint(0, int(taille))) + " "
        return string
    else:
        return ":3:"

def weather(channel):
    try:
        commands.getstatusoutput("curl fr.wttr.in/48.90,2.32?T0 > meteo.txt")
    except IOError:
        return
    with open('./meteo.txt', 'r') as fp:
       sc.api_call(
               "files.upload",
               channels=channel,
               file=fp,
               )

def oss(channel):
    with open('./oss.txt', 'r') as fp:
        citation = fp.read()
    message = citation.split('\n')[random.randint(0, len(citation.split('\n'))- 1)]
    if message == "" :
        oss(channel)
        return
    sc.api_call(
            "chat.postMessage",
            channel= channel,
            icon_url= "https://static-cdn.jtvnw.net/emoticons/v1/518312/3.0",
            username = "Hubert Bonisseur de La Bath",
            text= ">" + message
        )
    return

def kaamelott(channel):
    with open('./kaamelott.txt', 'r') as fp:
        citation = fp.read()
    message = citation.split('\n')[random.randint(0, len(citation.split('\n'))- 1)]
    if message == "":
        kaamelott(channel)
        return
    sc.api_call(
            "chat.postMessage",
            channel= channel,
            icon_url= "https://img15.hostingpics.net/pics/4833663350.jpg",
            username = "Perceval",
            text= ">" + message
        )

def list(bh):
    if bh not in ["bh1","bh5","bh7"]:
            return "list = bh1 bh5 bh7"
    url = "http://42.erwanleboucher.fr/v2/" + bh
    response = requests.request("GET",url)
    return "Il reste {} personnes dans le {}".format(len(response.json()), bh)


def validate(date_text):
   if  "1970-01-01" in date_text:
        return (0)
   try:
        datetime.strptime(date_text, '%Y-%m-%d')
   except ValueError:
        return (0)
   return (1)

def listban(bannedfile):
    banned = []
    with open(bannedfile, 'r') as fp:
        banned = fp.read().split('\n')
    return banned

def addban(bannedfile, user):
    with open(bannedfile, 'a') as fp:
        fp.write(user +'\n')

def post_message(text, channel):
    sc.api_call(
        "chat.postMessage",
        channel=channel,
        text=text
        )

def post_reaction(text, channel, ts):
    sc.api_call(
        "reactions.add",
        timestamp=ts,
        channel=channel,
        name = text
        )

def crypto(cryptoname,currency, ts, channel):
    clientbinance = binanceClient("", "")
    c = CurrencyConverter()
    reply = ""
    try:
        if cryptoname != "BTC":
            info = clientbinance.get_ticker(symbol=(cryptoname + "BTC"))
        elif currency != "BTC":
            info = clientbinance.get_ticker(symbol=(cryptoname + "USDT"))
        else:
            info = None
        if currency in ['USD', 'EUR'] and info is not None:
            btcprice = clientbinance.get_ticker(symbol='BTCUSDT')
            if cryptoname != "BTC":
                if currency == 'EUR':
                    price = c.convert(float(info['lastPrice']) * float(btcprice['lastPrice']), 'USD', 'EUR')
                else:
                    price = float(info['lastPrice']) * float(btcprice['lastPrice'])
            else:
                if currency == 'EUR':
                    price = c.convert(float(info['lastPrice']), 'USD', 'EUR')
                else:
                    price = float(info['lastPrice'])
        else :
            if currency == 'BTC':
                price = float(info['lastPrice'])
            else:
                reply =  "Erreur"
    except :
        reply =  "Erreur"
    if reply != "Erreur" :
        reply = "{} : *{:,.2f} {}* ({:,.2f}%)".format(cryptoname, price, currency, float(info['priceChangePercent']))
    sc.api_call(
            "chat.postMessage",
            thread_ts = ts,
            channel = channel,
            text = reply
        )

def help(channel):
    with open('featurespic.jpeg', 'r') as fp:
       sc.api_call(
               "files.upload",
               channels=channel,
               file=fp,
               )


def handle_command(message, channel, ts, user):
    reply = ""
    if "jpp" in message.lower():
   	post_reaction("jpp", channel, ts)
    if "rip" in message.lower():
	post_reaction("rip", channel, ts)
    if  any(coal in message.lower() for coal in coalitions):
        post_reaction("the-alliance", channel, ts)
    if len(message.split( )) > 1 and message.split( )[0].lower() in ["coalibot", "cb", "bc"] :
        if len(message.split( )) > 2:
            if message.split( )[1].lower() == "prof":
                reply = profile(message.split( )[2].lower())
            if get_username(user) == "elebouch" and message.split( )[1].lower() == "banmusic":
                addban('banmusic.txt', message.split( )[2].lower())
                reply = "Utilisateur banni de bc musique"
            if message.split( )[1].lower() == "where":
                reply = where(message.split( )[2].lower())
            if message.split( )[1].lower() == "who":
                reply = who(message.split( )[2].lower())
            if message.split( )[1].lower() == "crypto":
                if (len(message.split( )) == 4):
                    currency = message.split( )[3].upper()
                else :
                    currency = "EUR"
                crypto(message.split( )[2].upper(),currency, ts, channel)
            if message.split( )[1].lower() == "list":
                reply = list(message.split( )[2].lower())
            if message.split( )[1].lower() == "addmusic" and get_username(user) not in listban('banmusic.txt'):
               reply = addmusic(message.split( )[2], user)
            if message.split( )[1].lower() == "logtime" :
                logtime(message, ts, channel)
            if message.split( )[1].lower() == "roll":
               if len(message.split( )) >= 4:
                    reply = roll(message.split( )[2], message.split( )[3])
               else :
                    reply = "Usage: bc roll nbdé tailledudé"
            if message.split( )[1].lower() == "lmgtfy" :
                try:
                    reply = "http://lmgtfy.com/?q=" + quote_plus(' '.join(message.split('\n')[0].split(' ')[2:]))
                except:
                    pass
        elif message.split( )[1].lower() == "score" :
            score(ts, channel)
        elif message.split( )[1].lower() == "alliance" :
            reply = alliance()
        elif message.split( )[1].lower() == "music" :
            reply = musique()
        elif message.split( )[1].lower() == "help" :
            help(channel)
        elif message.split( )[1].lower() == "source" :
            reply = "`https://github.com/genesixx/coalibot`"
        elif message.split( )[1].lower() == "brew" :
            reply = "```rm -rf $HOME/.brew && git clone --depth=1 https://github.com/Homebrew/brew $HOME/.brew && echo \'export PATH=$HOME/.brew/bin:$PATH\' >> $HOME/.zshrc && source $HOME/.zshrc && brew update```"
        elif message.split( )[1].lower() == "halp":
            reply = "Bonjours\n Je t'invite à taper `iscsictl list targets` dans ton terminal\n à copier la ligne contenant ton login mais sans la partie entre <>\n puis à taper `iscsictl <la ligne copiée>`"
	elif message.split( )[1].lower() == "home":
            reply = "si `Disk Not Ejected Properly ??` suivre : https://42born2code.slack.com/archives/C7P0Z4F3L/p1510233807000241"
	elif message.split( )[1].lower() == "meteo" :
            weather(channel)
    elif message.split( )[0].lower() == "!anais":
        post_reaction(parrot[random.randint(0, len(parrot)) - 1], channel, ts)
    elif message.split( )[0].lower() == "!elebouch":
        reply = "Charbonne pour finir son `ft_ls`"
    elif message.split( )[0].lower() == "!lain":
        post_reaction("francais", channel, ts)
    elif message.split( )[0].lower() == "!jcharloi":
        reply = "fais tes 9h!"
    elif message.split( )[0].lower() == "!glegendr":
	reply = "Corrige ton `fillit`!"
    elif message.split( )[0].lower() == "!makefile":
        reply = "`https://forum.intra.42.fr/topics/85/messages`"
    elif message.split( )[0].lower() == "!sygnano":
        reply = "https://youtu.be/V2UGfj2qPCw?t=8s"
    elif message.split( )[0].lower() == "!fpons":
        reply = "t'as pas une video à finir ?"
    elif message.split( )[0].lower() == "!oss":
        oss(channel)
    elif message.split( )[0].lower() == "!kaamelott":
        kaamelott(channel)
    elif message.split( )[0].lower() == "!nestor":
        reply =  "Pour commander sur nestor utilise le code `NESTOR42` ! tu peux utiliser le code de parrainage `cZ44h` pour avoir 5e gratuitement"
    elif message.split( )[0].lower() == "!test":
        reply =  "oui je suis osi un beauteux "
    elif message.split( )[0].lower() == "!mfranc":
        reply = random.choice([">Doucement avec les bots", ">Puuuuuuuuuuuuu"])
    elif message.split( )[0].lower() in ["!parrot", "!perroquet", "!perruche", "!parakeet"] :
        reply =  ":"+parrot[random.randint(0, len(parrot)) - 1]+":"
    if reply != "" or reply is not None :
        post_message(reply, channel)

def coalibot():
    if sc.rtm_connect():
        while True:
      		events = sc.rtm_read()
      		if len(events) > 0:
      			for event in events:
      				if ('channel' in event and 'text' in event and event.get('type') == 'message' and 'user' in event):
      					channel = event['channel']
      					message = event['text']
      					ts = event['ts']
      					user = event['user']
                        		if channel not in blacklist:
     					    handle_command(message, channel, ts, user)
        	time.sleep(0.5)

    else:
       print ("connection failed")

if __name__ == '__main__':
    try:
        coalibot()
    except Exception as e:
        print e
        os.execv(sys.executable, ['python'] + sys.argv)
