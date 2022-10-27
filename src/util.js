/**
 * String formatting via placeholders: has troubles with placeholders
 * injections.
 * @param {string} s - The template string (featuring {0}, {1}, ...)
 * @param {...*} x - A value whose string conversion is to replace the
 *                   corresponding placeholder.
 * @returns {string} The formatted string.
 * @todo Make this work with placeholders injections.
 */
module.exports.format = function(s, ...x) {
  s = s.slice();
  for (let i = 0; i < arguments.length - 1; ++i)
    s = s.replace(new RegExp("\\{" + i + "\\}", "gm"), x[i]);
  return s;
};

/**
 * Returns tomorrow's date.
 * @returns {Date} Tomorrow's date.
 */
module.exports.tomorrowDate = function() {
  const d = new Date();
  d.setDate(d.getDate() + 1);
  return d;
}

/* convert a string into kebab case
 * useful for GitHub repository
 *
 * example:
 * string = "Logica per l'informatica"
 * converted_string = toOurCase(string); = "logica-per-informatica" (sic!)
 */
module.exports.toKebabCase = function(str) {
  return str &&
  str
    .normalize("NFD")
    .replace(/[\u0300-\u036f]/g, "")
    .match(
      /(?:[A-Z]{2,}(?=[A-Z][a-z]+[0-9]*|\b)|[A-Z]?[a-z]+[0-9]*|[A-Z]|[0-9]+)'?/g
    )
    .filter((value) => !value.endsWith("'"))
    .map((x) => x.toLowerCase())
    .join("-");
}
