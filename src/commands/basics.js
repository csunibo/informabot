const { data } = require("../jsons"),
  { format } = require("../util"),
  { randomInt } = require("crypto"),
  TelegramBot = require("node-telegram-bot-api");

/**
 * Sends a simple text message.
 * @param {TelegramBot} bot The bot that should send the message.
 * @param {TelegramBot.Message} msg The message that triggered this action.
 * @param {string} text The context of the message that should be sent.
 */
function message(bot, msg, text) {
  bot
    .sendMessage(msg.chat.id, text, data.settings.messageOptions)
    .catch((e) => console.error(e.stack));
}
module.exports.message = message;

/**
 * Sends a dice message.
 * @param {TelegramBot} bot The bot that should send the message.
 * @param {TelegramBot.Message} msg The message that triggered this action.
 * @param {'🎲'|'🎯'|'🏀'|'⚽'|'🎳'|'🎰'} [emoji] The emoji that should be sent, must be one of:
 */
function dice(bot, msg, emoji) {
  if (!emoji) {
    const emojis = ["🎲", "🎯", "🏀", "⚽", "🎳", "🎰"];
    // Don't send last sent emoji
    while (emoji === this.last) emoji = emojis[randomInt(emojis.length)];
  }
  this.last = emoji;
  bot
    .sendDice(msg.chat.id, {
      allow_sending_without_reply: true,
      disable_notification: true,
      emoji
    })
    .catch((e) => console.error(e.stack));
}
module.exports.dice = dice;

/**
 * Controlla se ha senso limitare il comando /ludopatico
 * @param {TelegramBot} bot The bot that should send the message.
 * @param {TelegramBot.Message} msg The message that triggered this action.
 * @param {string} importantGroup Text on the importance of the group
 */

module.exports.considerLuck = function (bot, msg, importantGroup)
{
	var isImportant = false;
	const groupChatId = [-1563447632];

	if (groupChatId.includes(msg.chat.id))
		isImportant = true;

	if (!isImportant)
		dice(bot, msg);
	else
		message(bot, msg, importantGroup);
}

/**
 * Sends a help message.
 * @param {TelegramBot} bot The bot that should send the message.
 * @param {TelegramBot.Message} msg The message that triggered this action.
 * @param {Object[]} actions The list of all possible actions.
 */
module.exports.giveHelp = function (bot, msg, actions) {
  answer = "";
  for (command in data.actions)
    if (data.actions[command] && data.actions[command].description)
      answer += `/${command} - ${data.actions[command].description}\n`;
  message(bot, msg, answer);
};

/**
 * Sends a list message.
 * @param {TelegramBot} bot The bot that should send the message.
 * @param {TelegramBot.Message} msg The message that triggered this action.
 * @param {string} header The header for the new message.
 * @param {string} template A string with {0}, {1}, ... placeholders.
 * @param {*[][]} items Each array features the substitutions to the
 *                placholders.
 */
module.exports.list = function (bot, msg, header, template, items) {
  let text = header.slice();
  for (let i = 0; i < items.length; ++i) {
    const params = items[i].slice();
    params.unshift(template);
    text += format.apply(this, params);
  }
  message(bot, msg, text);
};
