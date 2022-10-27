const
  { message } = require("./basics"),
  { toKebabCase } = require("../util"),
  fetch = import("node-fetch"),
  TelegramBot = require("node-telegram-bot-api");

/**
 * Webscrape the lessons timetable.
 * @param {TelegramBot} bot The bot that should send the message.
 * @param {TelegramBot.Message} msg The message that triggered this action.
 * @param {string} url URL to webscrape from.
 * @param {Date} date Date to consider.
 * @param {string} title Calendar title.
 * @param {string} fallbackText Text for empty calendar.
 */
module.exports.timetable = function(bot, msg, url, date, title, fallbackText) {
  fetch(url)
    .then((res) => {
      let lectures = [];
      for (let i = 0; i < res.data.length; ++i) {
        let start = new Date(res.data[i].start);
        if (
          start.getFullYear() === date.getFullYear() &&
          start.getMonth() === date.getMonth() &&
          start.getDate() === date.getDate()
        )
          lectures.push(res.data[i]);
      }
      let text = title;
      lectures.sort((a, b) => a.start - b.start);
      for (let i = 0; i < lectures.length; ++i)
        text += `  🕘 <b><a href="${lectures[i].teams}">${lectures[i].title}</a></b> ${lectures[i].time}
  🏢 ${lectures[i].aule[0].des_edificio} - ${lectures[i].aule[0].des_piano}
  📍 ${lectures[i].aule[0].des_indirizzo}
  〰〰〰〰〰〰〰〰〰〰〰
`;
      message(bot, msg, lectures.length !== 0 ? text : fallbackText);
    })
    .catch((e) => {
      message(bot, msg, `Timetable: "${e}")`);
      console.error(e.stack);
    });
}

/**
 * Describe a given course.
 * @param {TelegramBot} bot The bot that should send the message.
 * @param {TelegramBot.Message} msg The message that triggered this action.
 * @param {string} name Course's name.
 * @param {string} virtuale Course's ID on virtuale.unibo.it.
 * @param {string} teams Course's ID on teams.microsoft.com.
 * @param {string[]} professor List of professors' usernames.
 * @param {string} professor List of professors' usernames.
 * @param {string} telegram Telegram group's ID.
 */
module.exports.course = function(bot, msg, name, virtuale, teams, website, professors, telegram) {
  const emails = professors
    ? professors.join("@unibo.it\n  ") + "@unibo.it\n  "
    : "";
  message(
    bot,
    msg,
    (name ? `<b>${name}</b>\n` : ``) +
      (virtuale
        ? `  <a href='https://virtuale.unibo.it/course/view.php?id=${virtuale}'>Virtuale</a>\n`
        : ``) +
      (teams
        ? `  <a href='https://teams.microsoft.com/l/meetup-join/19%3ameeting_${teams}%40thread.v2/0?context=%7b%22Tid%22%3a%22e99647dc-1b08-454a-bf8c-699181b389ab%22%2c%22Oid%22%3a%22080683d2-51aa-4842-aa73-291a43203f71%22%7d'>Videolezione</a>\n`
        : ``) +
      (website
        ? `  <a href='https://www.unibo.it/it/didattica/insegnamenti/insegnamento/${website}'>Sito</a>\n  <a href='https://www.unibo.it/it/didattica/insegnamenti/insegnamento/${website}/orariolezioni'>Orario</a>\n`
        : ``) +
      (emails ? `  ${emails}` : ``) +
      (name
        ? `  <a href='https://csunibo.github.io/${toKebabCase(
            name
          )}/'>📚 Risorse: materiali, libri, prove</a>\n  <a href='https://github.com/csunibo/${toKebabCase(
            name
          )}/'>📂 Repository GitHub delle risorse</a>\n`
        : ``) +
      (telegram ? `  <a href='t.me/${telegram}'>👥 Gruppo Studenti</a>\n` : ``)
  );
}
