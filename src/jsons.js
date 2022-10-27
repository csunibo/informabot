const { existsSync } = require("fs");

/** An object storing all of the JSON data. */
const data = {};
module.exports.data = data;

/**
 * Updates the JSON data.
 */
function readJsons() {
  data.actions = require("../json/actions.json");
  data.groups = existsSync("../json/groups.json")
    ? require("../json/groups.json")
    : {},
  data.memes = require("../json/memes.json");
  data.settings = require("../json/settings.json");
}
module.exports.readJsons = readJsons();
readJsons();
