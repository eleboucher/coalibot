var _0xe06b=['logtime','profil','who','where','./slack_api','./request','client-oauth2','./const','moment','sprintf-js','sprintf','env','INTRA_CLIENT_ID','INTRA_SECRET','https://api.intra.42.fr','credentials','data','access_token','Request-Promise','/v2/coalitions','sort','score','\x20points\x20d\x27avance.\x20:the-alliance:','Nous\x20sommes\x20à\x20la\x20','eme\x20place\x20avec\x20','${name}\x20${score}\x0a','color','https://profile.intra.42.fr/blocs/1/coalitions','name','Powered\x20by\x20Coalibot','format','?page[size]=100&range[begin_at]=','/v2/users/','/locations/','duration','begin_at','&page[number]=','length','end_at','minutes','floor','invalid\x20login','2013','pool_year','2014','cursus_users','level','slug','https://projects.intra.42.fr/projects/graph?login=','A\x20fait\x20son','En\x20cours\x20de','project','status','N\x27a\x20pas\x20fait\x20son','displayname','pool_month','split','Usage:\x20bc\x20logtime\x20login\x20datedebut\x20datefin\x20(date\x20au\x20format\x20\x22Y-M-D\x22)','includes','replace','trimestre','getFullYear','get','month','%02dh%02d','today','isValid','startsWith','/v2/campus/1/locations/?filter[host]=','&filter[active]=true','place\x20invalide','*\x20vide','user','login','*\x20est\x20à\x20la\x20place\x20*','queen','way','elebouch','login\x20invalide','*\x20est\x20hors\x20ligne','host','exports','alliance'];(function(_0x2f1c1a,_0x409058){var _0x2c2baa=function(_0x2544c6){while(--_0x2544c6){_0x2f1c1a['push'](_0x2f1c1a['shift']());}};_0x2c2baa(++_0x409058);}(_0xe06b,0x150));var _0x51f6=function(_0x40ee1f,_0x5a5f61){_0x40ee1f=_0x40ee1f-0x0;var _0x102c2e=_0xe06b[_0x40ee1f];return _0x102c2e;};const {postMessage,postUserMessage,sendReaction,fileUpload,postOnThread,getUsername,postAttachments}=require(_0x51f6('0x0'));const rq=require(_0x51f6('0x1'))['rq'];const ClientOAuth2=require(_0x51f6('0x2'));const {month}=require(_0x51f6('0x3'));const moment=require(_0x51f6('0x4'));const sprintf=require(_0x51f6('0x5'))[_0x51f6('0x6')];const forty2auth=new ClientOAuth2({'clientId':process[_0x51f6('0x7')][_0x51f6('0x8')],'clientSecret':process[_0x51f6('0x7')][_0x51f6('0x9')],'accessTokenUri':'https://api.intra.42.fr/oauth/token'});const request42=async _0x140a33=>{var _0x140a33=_0x51f6('0xa')+_0x140a33;const _0x5dec7c=await forty2auth[_0x51f6('0xb')]['getToken']();var _0x2f4a05={'uri':_0x140a33,'qs':{'access_token':_0x5dec7c[_0x51f6('0xc')][_0x51f6('0xd')]},'headers':{'User-Agent':_0x51f6('0xe')},'json':!![]};try{return await rq(_0x2f4a05);}catch(_0x3192f8){return null;}};const alliance=async _0x573964=>{const _0x301cb4=await request42(_0x51f6('0xf'));_0x301cb4[_0x51f6('0x10')](function(_0x8145e5,_0x3d7574){return _0x8145e5['score']<_0x3d7574[_0x51f6('0x11')];});let _0x5c6e59=0x0;while(_0x301cb4[_0x5c6e59]['id']!==0x2)_0x5c6e59+=0x1;if(_0x5c6e59===0x0){const _0x315755=_0x301cb4[_0x5c6e59][_0x51f6('0x11')]-_0x301cb4[0x1]['score'];postMessage('Felicitations\x20Nous\x20sommes\x20premiers\x20avec\x20'+(_0x5c6e59+0x1)+_0x51f6('0x12'),_0x573964);}else{const _0x3a970c=_0x301cb4[0x0][_0x51f6('0x11')]-_0x301cb4[_0x5c6e59]['score'];postMessage(_0x51f6('0x13')+(_0x5c6e59+0x1)+_0x51f6('0x14')+_0x3a970c+'\x20points\x20de\x20retard.\x20:the-alliance:',_0x573964);}};const score=async _0x231794=>{const _0x323479=await request42(_0x51f6('0xf'));_0x323479[_0x51f6('0x10')](function(_0x17ed83,_0x4c5605){return _0x17ed83[_0x51f6('0x11')]<_0x4c5605[_0x51f6('0x11')];});var _0x11f221='';for(let _0x3ed81f of _0x323479){_0x11f221+=_0x51f6('0x15');}var _0x819c62=[{'fallback':_0x11f221,'color':_0x323479[0x0][_0x51f6('0x16')],'author_link':_0x51f6('0x17'),'fields':[{'title':_0x323479[0x0][_0x51f6('0x18')],'value':_0x323479[0x0][_0x51f6('0x11')],'short':!![]},{'title':_0x323479[0x1]['name'],'value':String(_0x323479[0x1][_0x51f6('0x11')]+'\x20('+Number(_0x323479[0x1][_0x51f6('0x11')]-_0x323479[0x0]['score'])+')'),'short':!![]},{'title':_0x323479[0x2][_0x51f6('0x18')],'value':String(_0x323479[0x2][_0x51f6('0x11')]+'\x20('+Number(_0x323479[0x2][_0x51f6('0x11')]-_0x323479[0x0][_0x51f6('0x11')])+')'),'short':!![]},{'title':_0x323479[0x3][_0x51f6('0x18')],'value':String(_0x323479[0x3]['score']+'\x20('+Number(_0x323479[0x3]['score']-_0x323479[0x0][_0x51f6('0x11')])+')'),'short':!![]}],'footer':_0x51f6('0x19')}];postAttachments('',_0x819c62,_0x231794);};const get_range_logtime=async(_0x39ef83,_0x4af590,_0x278a27)=>{_0x4af590=moment(_0x4af590)[_0x51f6('0x1a')]('YYYY-MM-DD');_0x278a27=moment(_0x278a27)[_0x51f6('0x1a')]('YYYY-MM-DD');range_date=_0x51f6('0x1b')+_0x4af590+','+_0x278a27;url=_0x51f6('0x1c')+_0x39ef83+_0x51f6('0x1d')+range_date;const _0x312e6b=await request42(url);if(_0x4af590===_0x278a27){return moment[_0x51f6('0x1e')](0x0);}try{async function _0x23493b(_0x43482f){let _0x2410fa;let _0x596ba7=0x2;let _0x527745=_0x43482f;do{last_location=moment(_0x527745[_0x527745['length']-0x1][_0x51f6('0x1f')]);if(moment(_0x4af590)['isBefore'](last_location)){_0x2410fa=await request42(url+_0x51f6('0x20')+_0x596ba7);if(_0x2410fa){_0x527745=_0x527745['concat'](_0x2410fa);}_0x596ba7+=0x1;}else{return _0x527745;}}while(_0x2410fa&&_0x2410fa[_0x51f6('0x21')]);return _0x527745;}let _0x519779=await _0x23493b(_0x312e6b);let _0x132109=moment[_0x51f6('0x1e')](0x0);for(let _0x35848c of _0x519779){if(_0x35848c[_0x51f6('0x22')])log_end=moment(_0x35848c[_0x51f6('0x22')]);else log_end=moment();log_start=moment(_0x35848c[_0x51f6('0x1f')]);log_session=log_end-log_start;_0x132109['add'](log_session);}return _0x132109;}catch(_0x37b97f){return moment[_0x51f6('0x1e')](0x0);}};const format_output_datetime=_0x132d85=>{const _0x2df13c=Number(_0x132d85['as'](_0x51f6('0x23')));const _0x29d6cb=Math[_0x51f6('0x24')](_0x2df13c/0x3c);const _0x164ef0=Math[_0x51f6('0x24')](_0x2df13c%0x3c);return[_0x29d6cb,_0x164ef0];};const profil=async(_0x486725,_0x4422c6)=>{url='/v2/users/'+_0x486725;urlcoal=url+'/coalitions/';const _0x3e5e08=await request42(url);let _0x585d8b=0x1;const _0x2a822e=await request42(urlcoal);if(!_0x3e5e08){postMessage(_0x51f6('0x25'),_0x4422c6);return;}let _0xa02bf7=0x0;if(_0x3e5e08['pool_year']===_0x51f6('0x26')||_0x3e5e08[_0x51f6('0x27')]===_0x51f6('0x28')){_0xa02bf7=0x0;}else if(_0x3e5e08[_0x51f6('0x29')]['length']===0x1){_0xa02bf7=_0x3e5e08['cursus_users'][0x0][_0x51f6('0x2a')];_0x585d8b=0x0;}else _0xa02bf7=_0x3e5e08[_0x51f6('0x29')][0x1][_0x51f6('0x2a')];let _0x44bedd='';if(_0x2a822e['length'])_0x44bedd=':'+_0x2a822e[0x0][_0x51f6('0x2b')]+':';range_end=moment();range_begin=moment()['subtract'](0x7,'days');const _0x1c9e49=await get_range_logtime(_0x486725,range_begin,range_end);var _0x4cc29c=format_output_datetime(_0x1c9e49);graph=_0x51f6('0x2c')+_0x486725;const _0x196beb=(_0x2507ed=>{const _0x3d2bda={'finished':_0x51f6('0x2d'),'in_progress':_0x51f6('0x2e')};const _0x580013=_0x2507ed['projects_users']['find'](_0x149e9d=>_0x149e9d[_0x51f6('0x2f')]['id']===0x76);return _0x580013?_0x3d2bda[_0x580013[_0x51f6('0x30')]]:_0x51f6('0x31');})(_0x3e5e08);postMessage(sprintf('%s\x20%s\x0aPhoto:\x20`%s`\x0aTemps\x20de\x20log\x20cette\x20semaine\x20%02d:%02d\x0aNiveau:\x20%.2f\x0aNiveau\x20piscine\x20\x20%.2f\x20%s\x20%s\x0a%s\x20stage\x0aGraph:\x20%s',_0x3e5e08[_0x51f6('0x32')],_0x44bedd,_0x3e5e08['image_url'],_0x4cc29c[0x0],_0x4cc29c[0x1],_0x585d8b===0x0?0x0:_0x3e5e08['cursus_users'][0x0]['level'],_0xa02bf7,_0x3e5e08[_0x51f6('0x33')],_0x3e5e08[_0x51f6('0x27')],_0x196beb,graph),_0x4422c6);};const logtime=async(_0x2e618c,_0x534277,_0x53a1c2)=>{if(_0x2e618c[_0x51f6('0x34')]('\x20')['length']<0x4){postOnThread(_0x51f6('0x35'),_0x534277,_0x53a1c2);}else if(_0x2e618c['split']('\x20')[_0x51f6('0x21')]===0x4&&!isNaN(_0x2e618c[_0x51f6('0x34')]('\x20')[0x3])&&parseInt(_0x2e618c[_0x51f6('0x34')]('\x20')[0x3])>0x7dc){let _0x39dd7f=moment({'y':parseInt(_0x2e618c[_0x51f6('0x34')]('\x20')[0x3]),'M':0x0,'d':0x1});let _0x1d209c=moment({'y':parseInt(_0x2e618c['split']('\x20')[0x3]),'M':0xb,'d':0x1f});const _0x2a757c=await get_range_logtime(_0x2e618c[_0x51f6('0x34')]('\x20')[0x2],_0x39dd7f,_0x1d209c);var _0x45c1e2=format_output_datetime(_0x2a757c);postOnThread(sprintf('%02dh%02d',_0x45c1e2[0x0],_0x45c1e2[0x1]),_0x534277,_0x53a1c2);}else if(_0x2e618c['split']('\x20')[0x3][_0x51f6('0x36')]('trimestre')&&(_0x2e618c[_0x51f6('0x34')]('\x20')[_0x51f6('0x21')]===0x4||_0x2e618c[_0x51f6('0x34')]('\x20')[_0x51f6('0x21')]===0x5&&parseInt(_0x2e618c[_0x51f6('0x34')]('\x20')[0x4])>0x7dc)){let _0x2e25bd=parseInt(_0x2e618c['split']('\x20')[0x3][_0x51f6('0x37')](_0x51f6('0x38'),''))-0x1;let _0x18e5cc;if(_0x2e618c[_0x51f6('0x34')]('\x20')[_0x51f6('0x21')]===0x5&&parseInt(_0x2e618c[_0x51f6('0x34')]('\x20')[0x4])>0x7dc)_0x18e5cc=parseInt(_0x2e618c[_0x51f6('0x34')]('\x20')[0x4]);else _0x18e5cc=new Date()[_0x51f6('0x39')]();let _0x25c618=moment(new Date(_0x18e5cc,_0x2e25bd*0x3,0x1));let _0x450a6a=moment(new Date(_0x18e5cc,_0x25c618[_0x51f6('0x3a')](_0x51f6('0x3b'))+0x3,0x0));const _0xaf8667=await get_range_logtime(_0x2e618c[_0x51f6('0x34')]('\x20')[0x2],_0x25c618,_0x450a6a);var _0x45c1e2=format_output_datetime(_0xaf8667);postOnThread(sprintf(_0x51f6('0x3c'),_0x45c1e2[0x0],_0x45c1e2[0x1]),_0x534277,_0x53a1c2);}else if(_0x2e618c[_0x51f6('0x34')]('\x20')[0x3]in month&&(_0x2e618c[_0x51f6('0x34')]('\x20')[_0x51f6('0x21')]===0x4||_0x2e618c[_0x51f6('0x34')]('\x20')[_0x51f6('0x21')]===0x5&&parseInt(_0x2e618c[_0x51f6('0x34')]('\x20')[0x4])>0x7dc)){if(_0x2e618c[_0x51f6('0x34')]('\x20')[_0x51f6('0x21')]===0x5&&parseInt(_0x2e618c[_0x51f6('0x34')]('\x20')[0x4])>0x7dc)year=parseInt(_0x2e618c[_0x51f6('0x34')]('\x20')[0x4]);else year=new Date()[_0x51f6('0x39')]();let _0x3fac26=moment(new Date(year,month[_0x2e618c['split']('\x20')[0x3]],0x1));let _0xc87dd5=moment(new Date(year,month[_0x2e618c[_0x51f6('0x34')]('\x20')[0x3]]+0x1,0x0));const _0x1ebc47=await get_range_logtime(_0x2e618c['split']('\x20')[0x2],_0x3fac26,_0xc87dd5);var _0x45c1e2=format_output_datetime(_0x1ebc47);postOnThread(sprintf(_0x51f6('0x3c'),_0x45c1e2[0x0],_0x45c1e2[0x1]),_0x534277,_0x53a1c2);}else if(_0x2e618c['split']('\x20')[_0x51f6('0x21')]===0x5){let _0x561fb0;if(_0x2e618c[_0x51f6('0x34')]('\x20')[0x4]===_0x51f6('0x3d'))_0x561fb0=moment();else _0x561fb0=moment(_0x2e618c[_0x51f6('0x34')]('\x20')[0x4]);let _0xf14c49=moment(_0x2e618c[_0x51f6('0x34')]('\x20')[0x3]);if(_0x561fb0[_0x51f6('0x3e')]()&&_0xf14c49[_0x51f6('0x3e')]()){const _0xfdbe99=await get_range_logtime(_0x2e618c[_0x51f6('0x34')]('\x20')[0x2],_0xf14c49,_0x561fb0);var _0x45c1e2=format_output_datetime(_0xfdbe99);postOnThread(sprintf(_0x51f6('0x3c'),_0x45c1e2[0x0],_0x45c1e2[0x1]),_0x534277,_0x53a1c2);}}};const who=async(_0x2358b1,_0x5eb02e)=>{if(!_0x2358b1||_0x2358b1[_0x51f6('0x3f')]('!')||_0x2358b1[_0x51f6('0x3f')]('?'))return;const _0x3fb6c6=_0x51f6('0x40')+_0x2358b1+_0x51f6('0x41');const _0x411132=await request42(_0x3fb6c6);if(!_0x411132){postMessage(_0x51f6('0x42'),_0x5eb02e);return;}if(_0x411132[_0x51f6('0x21')]===0x0)postMessage('Place\x20*'+_0x2358b1+_0x51f6('0x43'),_0x5eb02e);else postMessage('*'+_0x411132[0x0][_0x51f6('0x44')][_0x51f6('0x45')]+_0x51f6('0x46')+_0x2358b1+'*',_0x5eb02e);};const where=async(_0x53e01d,_0x59a0b2)=>{if(!_0x53e01d||_0x53e01d[_0x51f6('0x3f')]('!')||_0x53e01d['startsWith']('?'))return;if(_0x53e01d===_0x51f6('0x47')||_0x53e01d==_0x51f6('0x48')){postMessage('follow\x20me\x20bruddah\x0ai\x27ll\x20show\x20you\x20de\x20way\x20:uganda_knuckles:',_0x59a0b2);return;}if(_0x53e01d==='dieu')_0x53e01d=_0x51f6('0x49');url=_0x51f6('0x1c')+_0x53e01d+'/locations';const _0x39a715=await request42(url);if(!_0x39a715){postMessage(_0x51f6('0x4a'),_0x59a0b2);return;}if(_0x39a715[_0x51f6('0x21')]===0x0||_0x39a715[0x0][_0x51f6('0x22')])postMessage('*'+_0x53e01d+_0x51f6('0x4b'),_0x59a0b2);else postMessage('*'+_0x53e01d+_0x51f6('0x46')+_0x39a715[0x0][_0x51f6('0x4c')]+'*',_0x59a0b2);};module[_0x51f6('0x4d')][_0x51f6('0x4e')]=alliance;module[_0x51f6('0x4d')][_0x51f6('0x4f')]=logtime;module[_0x51f6('0x4d')]['score']=score;module[_0x51f6('0x4d')][_0x51f6('0x50')]=profil;module[_0x51f6('0x4d')][_0x51f6('0x51')]=who;module[_0x51f6('0x4d')][_0x51f6('0x52')]=where;