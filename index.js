// Initial setup
if (process.argv.length != 3)
  console.log('usage: node index.js token');
process.env.NTBA_FIX_319 = 1;
const axios = require('axios'),
  fs = require('fs'),
  TelegramBot = require('node-telegram-bot-api'),
  actions = require('./json/actions.json'),
  groups = fs.existsSync('./json/groups.json') ? require('./json/groups.json') : {},
  memes = require('./json/memes.json'),
  settings = require('./json/settings.json'),
  bot = new TelegramBot(process.argv[2], {polling: true});

// String formatting via placeholders: has troubles with placeholders injections
String.format = function () {
  var s = arguments[0].slice();
  for (var i = 0; i < arguments.length - 1; ++i)
    s = s.replace(new RegExp("\\{" + i + "\\}", "gm"), arguments[i + 1]);
  return s;
}

// Simple messages
function message(msg, text) {
  bot.sendMessage(msg.chat.id, text, settings.messageOptions).catch(e => console.error(e.stack));
}

// Web scraping the timetable -- Lezioni oggi
function trovalezionioggi(msg, url, fallbackText) {
  axios.get(url)
    .then(res => {
      let now = new Date();
      let todayLectures = [];
      for (let i = 0; i < res.data.length; ++i) {
        let start = new Date(res.data[i].start);
        if (start.getFullYear() === now.getFullYear() &&
          start.getMonth() === now.getMonth() &&
          (start.getDate() === now.getDate()))
          todayLectures.push(res.data[i]);
      }

      let text = '<b>Lezioni di oggi:</b>\n';
      todayLectures.sort((a, b) => {
        if (a.start > b.start)
          return 1;
        if (a.start < b.start)
          return -1;
        return 0;
      });
      for (let i = 0; i < todayLectures.length; ++i) {
        text += '🕘 <b>' + '<a href="' + todayLectures[i].teams + '">' + todayLectures[i].title + '</a></b> ' + todayLectures[i].time + '\n';
        text += '🏢 ' + todayLectures[i].aule[0].des_edificio +  ' - ' + todayLectures[i].aule[0].des_piano + '\n';
        text += '📍 '  + todayLectures[i].aule[0].des_indirizzo + '\n';
        text += '〰〰〰〰〰〰〰〰〰〰〰\n';
      }
      if (todayLectures.length !== 0)
        message(msg, text);
      else
        message(msg, fallbackText);
    }).catch(e => console.error(e.stack));
}

// Web scraping the timetable -- Lezioni domani
function trovalezionidomani(msg, url, fallbackText) {
  axios.get(url)
    .then(res => {
      let now = new Date();
	  let domani = new Date(now)
	  domani.setDate(domani.getDate() + 1)
      let tomorrowLectures = [];
      for (let i = 0; i < res.data.length; ++i) {
        let start = new Date(res.data[i].start);
        if (start.getFullYear() === domani.getFullYear() &&
          start.getMonth() === domani.getMonth() &&
          (start.getDate() === domani.getDate()))
          tomorrowLectures.push(res.data[i]);
      }

      let text = '<b>Lezioni di domani:</b>\n';
      tomorrowLectures.sort((a, b) => {
        if (a.start > b.start)
          return 1;
        if (a.start < b.start)
          return -1;
        return 0;
      });
      for (let i = 0; i < tomorrowLectures.length; ++i) {
        text += '🕘 <b>' + '<a href="' + tomorrowLectures[i].teams + '">' + tomorrowLectures[i].title + '</a></b> ' + tomorrowLectures[i].time + '\n';
        text += '🏢 ' + tomorrowLectures[i].aule[0].des_edificio +  ' - ' + tomorrowLectures[i].aule[0].des_piano + '\n';
        text += '📍 '  + tomorrowLectures[i].aule[0].des_indirizzo + '\n';
        text += '〰〰〰〰〰〰〰〰〰〰〰\n';
      }
      if (tomorrowLectures.length !== 0)
        message(msg, text);
      else
        message(msg, fallbackText);
    }).catch(e => console.error(e.stack));
}





// Autogenerated courses info
function course(msg, name, virtuale, teams, website, professors) {
  const emails = professors.join('@unibo.it\n  ') + '@unibo.it';
  message(msg, `<b>${name}</b>
  <a href='https://virtuale.unibo.it/course/view.php?id=${virtuale}'>Virtuale</a>
  <a href='https://teams.microsoft.com/l/meetup-join/19%3ameeting_${teams}%40thread.v2/0?context=%7b%22Tid%22%3a%22e99647dc-1b08-454a-bf8c-699181b389ab%22%2c%22Oid%22%3a%22080683d2-51aa-4842-aa73-291a43203f71%22%7d'>Videolezione</a>
  <a href='https://www.unibo.it/it/didattica/insegnamenti/insegnamento/${website}'>Sito</a>
  <a href='https://www.unibo.it/it/didattica/insegnamenti/insegnamento/${website}/orariolezioni'>Orario</a>
  ${emails}`);
}

// Adding a user to a list
function lookingFor(msg, singularText, pluralText, chatError) {
  if (msg.chat.type !== 'group' && msg.chat.type !== 'supergroup' || settings.lookingForBlackList.includes(msg.chat.id))
    message(msg, chatError);
  else {
    const chatId = msg.chat.id, senderId = msg.from.id;
    if (!(chatId in groups))
      groups[chatId] = [];
    const group = groups[chatId];
    if (!group.includes(senderId))
      group.push(senderId);
    fs.writeFileSync('json/groups.json', JSON.stringify(groups));
    const length = group.length.toString(), promises = Array(length);
    var list = String.format(length == '1' ? singularText : pluralText, msg.chat.title, length);
    group.forEach((e, i) => {
      promises[i] = bot.getChatMember(chatId, e.toString()).then(
        (result) => {
          const user = result.user;
          return `👤 <a href='tg://user?id=${user.id}'>${user.first_name}${user.last_name ? ' ' + user.last_name : ''}</a>\n`;
        }
      );
    });
    Promise.allSettled(promises).then(
      (result) => {
        result.forEach(e => {
          if (e.status === 'fulfilled')
            list += e.value;
        });
        message(msg, list);
      });
  }
}

// Removing a user from a list
function notLookingFor(msg, text, chatError, notFoundError) {
  if (msg.chat.type !== 'group' && msg.chat.type !== 'supergroup' || settings.lookingForBlackList.includes(msg.chat.id))
    message(msg, chatError);
  else {
    const chatId = msg.chat.id, title = msg.chat.title;
    if (!(chatId in groups))
      message(msg, String.format(notFoundError, title));
    else {
      const group = groups[chatId], senderId = msg.from.id;
      if (!group.includes(senderId))
        message(msg, String.format(notFoundError, title));
      else {
        group.splice(group.indexOf(senderId), 1);
        if (group.length == 0)
          delete groups[chatId];
        fs.writeFileSync('json/groups.json', JSON.stringify(groups));
        message(msg, String.format(text, title));
      }
    }
  }
}

// Available actions
function act(msg, action) {
  switch (action.type) {
    case 'alias':
      act(msg, actions[action.command]);
      break;
    case 'course':
      course(msg, action.name, action.virtuale, action.teams, action.website, action.professors);
      break;
    case 'lookingFor':
      lookingFor(msg, action.singularText, action.pluralText, action.chatError);
      break;
    case 'message':
      message(msg, action.text);
      break;
    case 'notLookingFor':
      notLookingFor(msg, action.text, action.chatError, action.notFoundError);
      break;
    case 'lezionioggi':
      trovalezionioggi(msg, action.url, action.fallbackText);
      break;
	case 'lezionidomani':
     trovalezionidomani(msg, action.url, action.fallbackText);
      break;  
    default:
      console.error(`Unknown action type "${action.type}"`);
  }
}

// Parsing
function onMessage(msg) {
  if (msg.text) {
    const text = msg.text.toString()
    if (text[0] == '/') {
      // '/command@bot param0 ... paramN' -> 'command'
      command = text.toLowerCase().split(' ')[0].substring(1);
      if (command.includes('@'))
        command = command.substring(0, command.indexOf('@'));
      if (command in actions)
        act(msg, actions[command]);
      else if (command in memes)
        message(msg, memes[command]);
    }
  }
}

bot.on('message', onMessage);
