const { data, readJsons } = require("../jsons"),
  { message } = require("./basics"),
  git = require("simple-git")(),
  TelegramBot = require("node-telegram-bot-api");

/**
 * Updates the whole bots, but rereads the JSONs only.
 * @param {TelegramBot} bot The bot that should send the message.
 * @param {TelegramBot.Message} msg The message that triggered this action.
 * @param {string} started Text on update start.
 * @param {string} ended Text on update end.
 * @param {string} failed Text on update failure.
 */
function update(bot, msg, started, ended, failed) {
  message(bot, msg, started);
  git
    .pull()
    .then((_) => {
      readJsons();
      message(bot, msg, ended);
    })
    .catch(e => {
      message(bot, msg, failed);
      console.error(e);
    });
}

/**
 * Updates the whole bots, but rereads the JSONs only.
 * @param {TelegramBot} bot The bot that should send the message.
 * @param {TelegramBot.Message} msg The message that triggered this action.
 * @param {string} noYear Error message on non-general group.
 * @param {string} noMode Error message on non-moderator sender.
 * @param {string} started Text on update start.
 * @param {string} ended Text on update end.
 * @param {string} failed Text on update failure.
 */
module.exports.considerUpdating = function (
  bot,
  msg,
  noYear,
  noMod,
  started,
  ended,
  failed
) {
  if (
    (msg.chat.type !== "group" && msg.chat.type !== "supergroup") ||
    !data.settings.generalGroups.includes(msg.chat.id)
  )
    message(bot, msg, noYear);
  else
    bot
      .getChatAdministrators(msg.chat.id)
      .then((admins) => {
        if (admins.map((x) => x.user.id).includes(msg.from.id))
          update(bot, msg, started, ended, failed);
        else message(bot, msg, noMod);
      })
      .catch(console.error);
};
