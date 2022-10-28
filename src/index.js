const { startInformabot } = require("./bot");
if (process.argv.length != 3) console.log("usage: node index.js <token>");
else startInformabot(process.argv[2]);
