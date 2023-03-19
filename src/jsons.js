const { existsSync } = require("fs");
const path = require("path");

/** An object storing all of the JSON data. */
const data = {};
module.exports.data = data;

/**
 * Computes the path of a given json file looking into the json/ folder.
 * @param {string} x Name of the json file (w/o .json extension).
 * @returns {string} Absolute path of the desired json file.
 */
const jsonPath = (x) => path.join(__dirname, "..", "json", x + ".json");

/**
 * Requires the path of a given json file looking into the json/ folder.
 * @param {string} x Name of the json file (w/o .json extension).
 * @returns {any} The contents of the file in question.
 */
const requireJsonPath = (x) => require(jsonPath(x));

/**
 * Updates the JSON data.
 */
function readJsons() {
  data.actions = requireJsonPath("actions");
  const groupsFilePath = jsonPath("groups");
  (data.groups = existsSync(groupsFilePath) ? require(groupsFilePath) : {}),
    (data.memes = requireJsonPath("memes"));
  data.settings = requireJsonPath("settings");
  data.autoreply = requireJsonPath("autoreply");
}

module.exports.readJsons = readJsons;
readJsons();
