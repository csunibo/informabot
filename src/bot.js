const { data } = require("./jsons"),
  { act } = require("./router"),
  { message } = require("./commands/basics"),
  TelegramBot = require("node-telegram-bot-api");

/**
 * Parses a message.
 * @param {TelegramBot} bot The bot that should listen for the message.
 * @param {TelegramBot.Message} msg The message that triggered this action.
 */
function onMessage(bot, msg) {
  if (!msg.text) return; // no text
  const text = msg.text.toString();
  if (text[0] !== "/") {
    Object.entries(data.autoreply).forEach(([regexp, value]) => {
      const indexOfAt = text.search(new RegExp(regexp, "i")); //case insensitive search
      if (indexOfAt != -1) message(bot, msg, value);
    });
    return; // no command
  }
  // '/command@bot param0 ... paramN' -> 'command@bot'
  let command = text.split(" ")[0].substring(1);
  const indexOfAt = command.indexOf("@");
  if (indexOfAt != -1) {
    if (command.substring(indexOfAt + 1) !== bot.username) return; // command issued to another bot
    // 'command@bot' -> 'command'
    command = command.substring(0, command.indexOf("@"));
  }
  try {
    if (command in data.actions)
      // action
      act(bot, msg, data.actions[command]);
    else if (command in data.memes)
      // meme
      message(bot, msg, data.memes[command]);
    // unkown command
    else act(bot, msg, data.actions["unknown"]);
  } catch (e) {
    console.error(e);
  }
}

/**
 * Initializes a bot from his Telegram API user object.
 * @param {TelegramBot} bot The bot to be initialized.
 * @param {TelegramBot.User} botUser The user object by the Telegram API.
 */
function init(bot, botUser) {
  bot.username = botUser.username;
  bot.on("message", (msg) => onMessage(bot, msg));
  bot.on("error", console.error);
  bot.on("polling_error", console.error);
}

/**
 * Build a new Informabot and return it.
 * @param {string} token The Telegram Bot API token to be used.
 */
module.exports.startInformabot = function (token) {
  process.env.NTBA_FIX_319 = 1;
  let bot = new TelegramBot(token, { polling: true });
  bot
    .getMe()
    .then((botUser) => init(bot, botUser))
    .catch(console.error);
  return bot;
};
