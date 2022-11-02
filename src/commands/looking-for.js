const { message } = require("./basics"),
  { format } = require("../util"),
  { data } = require("../jsons"),
  fs = require("fs"),
  path = require("path"),
  TelegramBot = require("node-telegram-bot-api");

/**
 * Add a user to the "looking for" list.
 * @param {TelegramBot} bot The bot that should send the message.
 * @param {TelegramBot.Message} msg The message that triggered this action.
 * @param {string} singularText Singular header for singlet lists.
 * @param {string} pluralText Plural header for non-singlet lists.
 * @param {string} chatError Error message for wrong chats.
 */
module.exports.lookingFor = function (bot, msg, singularText, pluralText, chatError) {
  if (
    (msg.chat.type !== "group" && msg.chat.type !== "supergroup") ||
    settings.lookingForBlackList.includes(msg.chat.id)
  )
    message(bot, msg, chatError);
  else {
    const chatId = msg.chat.id,
      senderId = msg.from.id;
    if (!(chatId in data.groups)) data.groups[chatId] = [];
    const group = data.groups[chatId];
    if (!group.includes(senderId)) group.push(senderId);
    fs.writeFileSync("json/groups.json", JSON.stringify(data.groups));
    const length = group.length.toString(),
      promises = Array(length);
    group.forEach((e, i) => {
      promises[i] = bot
        .getChatMember(chatId, e.toString())
        .then(
          (result) => {
            const user = result.user;
            return `👤 <a href='tg://user?id=${user.id}'>${user.first_name}${
              user.last_name ? " " + user.last_name : ""
            }</a>\n`;
          },
          (reason) => console.error(reason)
        )
        .catch((e) => {
          message(bot, msg, `Looking for: "${e}")`);
          console.error(e.stack);
        });
    });
    Promise.allSettled(promises).then((result) => {
      let list = format(
        length == "1" ? singularText : pluralText,
        msg.chat.title,
        length
      );
      result.forEach((e, i) => {
        list +=
          e.status === "fulfilled" && e.value
            ? e.value
            : `👤 <a href='tg://user?id=${group[i]}'>??? ???</a>\n`;
      });
      message(bot, msg, list);
    });
  }
};

/**
 * Remove a user to the "looking for" list.
 * @param {TelegramBot} bot The bot that should send the message.
 * @param {TelegramBot.Message} msg The message that triggered this action.
 * @param {string} text Content of the action acknowledgement message.
 * @param {string} chatError Error message for wrong chats.
 * @param {string} notFoundError Error message for users not in list.
 */
module.exports.notLookingFor = function (
  bot,
  msg,
  text,
  chatError,
  notFoundError
) {
  if (
    (msg.chat.type !== "group" && msg.chat.type !== "supergroup") ||
    settings.lookingForBlackList.includes(msg.chat.id)
  )
    message(bot, msg, chatError);
  else {
    const chatId = msg.chat.id,
      title = msg.chat.title;
    if (!(chatId in data.groups))
      message(bot, msg, format(notFoundError, title));
    else {
      const group = data.groups[chatId],
        senderId = msg.from.id;
      if (!group.includes(senderId))
        message(bot, msg, format(notFoundError, title));
      else {
        group.splice(group.indexOf(senderId), 1);
        if (group.length == 0) delete data.groups[chatId];
        fs.writeFileSync(
          path.join(__dirname, "..", "..", "json", "groups.json"),
          JSON.stringify(data.groups)
        );
        message(bot, msg, format(text, title));
      }
    }
  }
};
