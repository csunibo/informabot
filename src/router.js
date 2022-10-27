const
  { data } = require("./jsons"),
  { tomorrowDate } = require("./util"),
  { message, giveHelp, list } = require("./commands/basics"),
  { lookingFor, notLookingFor } = require("./commands/looking-for"),
  { timetable, course } = require("./commands/uni"),
  { considerUpdating } = require("./commands/update"),
  TelegramBot = require("node-telegram-bot-api");

/**
 * Invokes the right procedure of each action.
 * @param {TelegramBot} bot The bot that should listen for the message.
 * @param {TelegramBot.Message} msg The message that triggered this action.
 * @param {Object} action The action to be interpreted.
 */
module.exports.act = function(bot, msg, action) {
  switch (action.type) {
    case "alias":
      act(bot, msg, data.actions[action.command]);
      break;
    case "course":
      course(
        bot,
        msg,
        action.name,
        action.virtuale,
        action.teams,
        action.website,
        action.professors,
        action.telegram
      );
      break;
    case "help":
      giveHelp(bot, msg);
      break;
    case "lookingFor":
      lookingFor(bot, msg, action.singularText, action.pluralText, action.chatError);
      break;
    case "message":
      message(bot, msg, action.text);
      break;
    case "notLookingFor":
      notLookingFor(bot, msg, action.text, action.chatError, action.notFoundError);
      break;
    case "todayLectures":
      timetable(bot, msg, action.url, new Date(), action.title, action.fallbackText);
      break;
    case "tomorrowLectures":
      timetable(
        bot,
        msg,
        action.url,
        tomorrowDate(),
        action.title,
        action.fallbackText
      );
      break;
    case "yearly":
      yearly(bot, msg, action.command, action.noYear);
      break;
    case "list":
      list(bot, msg, action.header, action.template, action.items);
      break;
    case "update":
      considerUpdating(
        bot,
        msg,
        action.noYear,
        action.noMod,
        action.started,
        action.ended,
        action.failed
      );
      break;
    default:
      console.error(`Unknown action type "${action.type}"`);
  }
}

/**
 * Picks a different command based on the year of the current group.
 * @param {TelegramBot} bot The bot that should send the message.
 * @param {TelegramBot.Message} msg The message that triggered this action.
 * @param {string} command The command name's.
 * @param {string} noYear Error message for missing year.
 */
module.exports.yearly = function(bot, msg, command, noYear) {
  bot
    .getChat(msg.chat.id)
    .then((chat) => {
      const title = chat.title ? chat.title.toLowerCase() : "";
      if (title.includes("primo")) act(bot, msg, data.actions[command + "1"]);
      else if (title.includes("secondo")) act(bot, msg, data.actions[command + "2"]);
      else if (title.includes("terzo")) act(bot, msg, data.actions[command + "3"]);
      else message(bot, msg, noYear);
    })
    .catch((e) => {
      message(bot, msg, `Yearly: "${e}")`);
      console.error(e.stack);
    });
}
